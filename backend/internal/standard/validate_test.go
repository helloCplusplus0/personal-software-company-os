// Standard 树校验单元测试（R1-R8 + R2 三时机 + 序列化等价）。
//
// 覆盖冻结要求（phase14-07 Task 4）：
//   - 9 个稳定错误码逐码覆盖（R7 含 INVALID_FILE_NODE / INVALID_SUMMARY_LENGTH 两码）
//   - 每条规则至少 1 个非法用例（断言错误码 + 节点路径）+ 1 个合法边界用例
//   - phase14-03 示例树（根 "." + AGENTS.md + docs/phase/README.md，ref 两种形态）全规则通过
//   - children nil / [] 序列化反序列化等价；role / summary / ref omitempty 语义
//
// 文件落点：backend/internal/standard/validate_test.go
package standard

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"testing"
)

// validTree 构造 phase14-03 示例合法完整树：
// 根 "." + AGENTS.md file 节点（ref 树内路径形态）+ docs/phase/README.md
// 嵌套 directory + file（ref 树内路径形态）。
func validTree() *DirectoryTreeNode {
	return &DirectoryTreeNode{
		Name:     ".",
		NodeType: "directory",
		Children: []*DirectoryTreeNode{
			{
				Name:     "AGENTS.md",
				NodeType: "file",
				Role:     "全局上下文入口",
				Summary:  "项目定位与接手提醒入口",
				Ref:      "/AGENTS.md",
			},
			{
				Name:     "docs",
				NodeType: "directory",
				Children: []*DirectoryTreeNode{
					{
						Name:     "phase",
						NodeType: "directory",
						Children: []*DirectoryTreeNode{
							{
								Name:     "README.md",
								NodeType: "file",
								Role:     "docs workflow 入口",
								Ref:      "/docs/phase/README.md",
							},
						},
					},
				},
			},
		},
	}
}

// emptyRootTree 构造单根空树（children 为 nil）。
func emptyRootTree() *DirectoryTreeNode {
	return &DirectoryTreeNode{Name: ".", NodeType: "directory"}
}

// TestValidateTreeStructure 表驱动覆盖结构规则 R1/R3/R4/R5/R6/R7/R8。
func TestValidateTreeStructure(t *testing.T) {
	tests := []struct {
		name        string
		tree        func() *DirectoryTreeNode
		wantErrCode string // 空串表示合法用例
		wantNodeIn  string // 期望错误信息包含的节点路径
	}{
		// --- 合法用例 ---
		{
			name: "合法：phase14-03 示例完整树（ref 两种形态路径一致）",
			tree: validTree,
		},
		{
			name: "合法：单根空树（结构层面合法，R2 时机另行判定）",
			tree: emptyRootTree,
		},
		{
			name: "合法边界：ref 为 https:// URL 形态",
			tree: func() *DirectoryTreeNode {
				tree := validTree()
				tree.Children[0].Ref = "https://example.com/AGENTS.md"
				return tree
			},
		},
		{
			name: "合法边界：非根 name 恰好 64 字符",
			tree: func() *DirectoryTreeNode {
				name := strings.Repeat("a", 64)
				return &DirectoryTreeNode{
					Name:     ".",
					NodeType: "directory",
					Children: []*DirectoryTreeNode{{Name: name, NodeType: "file", Role: "r"}},
				}
			},
		},
		{
			name: "合法边界：深度恰好 6 层且第 6 层为 file",
			tree: func() *DirectoryTreeNode {
				// 根(.) → a(2) → b(3) → c(4) → d(5) → leaf.md(6, file)
				node := &DirectoryTreeNode{Name: "leaf.md", NodeType: "file", Role: "r"}
				for _, seg := range []string{"d", "c", "b", "a"} {
					node = &DirectoryTreeNode{Name: seg, NodeType: "directory", Children: []*DirectoryTreeNode{node}}
				}
				return &DirectoryTreeNode{Name: ".", NodeType: "directory", Children: []*DirectoryTreeNode{node}}
			},
		},
		{
			name: "合法边界：file role 恰好 32 字符",
			tree: func() *DirectoryTreeNode {
				return &DirectoryTreeNode{
					Name:     ".",
					NodeType: "directory",
					Children: []*DirectoryTreeNode{{
						Name: "AGENTS.md", NodeType: "file",
						Role: strings.Repeat("角", 32),
					}},
				}
			},
		},
		{
			name: "合法边界：summary 恰好 2000 字符",
			tree: func() *DirectoryTreeNode {
				return &DirectoryTreeNode{
					Name:     ".",
					NodeType: "directory",
					Children: []*DirectoryTreeNode{{
						Name: "AGENTS.md", NodeType: "file", Role: "r",
						Summary: strings.Repeat("摘", 2000),
					}},
				}
			},
		},

		// --- R1 INVALID_TREE_ROOT ---
		{
			name:        "R1 非法：tree 为 nil",
			tree:        func() *DirectoryTreeNode { return nil },
			wantErrCode: "INVALID_TREE_ROOT",
			wantNodeIn:  "/",
		},
		{
			name: "R1 非法：根 node_type 为 file",
			tree: func() *DirectoryTreeNode {
				return &DirectoryTreeNode{Name: ".", NodeType: "file", Role: "r"}
			},
			wantErrCode: "INVALID_TREE_ROOT",
			wantNodeIn:  "/",
		},
		{
			name: "R1 非法：根 name 非 \".\"",
			tree: func() *DirectoryTreeNode {
				return &DirectoryTreeNode{Name: "root", NodeType: "directory"}
			},
			wantErrCode: "INVALID_TREE_ROOT",
			wantNodeIn:  "/",
		},
		{
			name: "R1 非法：根 role 非空",
			tree: func() *DirectoryTreeNode {
				return &DirectoryTreeNode{Name: ".", NodeType: "directory", Role: "root-role"}
			},
			wantErrCode: "INVALID_TREE_ROOT",
			wantNodeIn:  "/",
		},

		// --- R3 DUPLICATE_SIBLING_NAME ---
		{
			name: "R3 非法：同层兄弟 name 重复（大小写敏感场景另测）",
			tree: func() *DirectoryTreeNode {
				return &DirectoryTreeNode{
					Name:     ".",
					NodeType: "directory",
					Children: []*DirectoryTreeNode{
						{Name: "AGENTS.md", NodeType: "file", Role: "r"},
						{Name: "AGENTS.md", NodeType: "file", Role: "r2"},
					},
				}
			},
			wantErrCode: "DUPLICATE_SIBLING_NAME",
			wantNodeIn:  "/AGENTS.md",
		},
		{
			name: "R3 合法边界：大小写不同不视为重复",
			tree: func() *DirectoryTreeNode {
				return &DirectoryTreeNode{
					Name:     ".",
					NodeType: "directory",
					Children: []*DirectoryTreeNode{
						{Name: "README.md", NodeType: "file", Role: "r"},
						{Name: "readme.md", NodeType: "file", Role: "r2"},
					},
				}
			},
		},

		// --- R4 INVALID_NODE_NAME ---
		{
			name: "R4 非法：name 含空格不匹配形态",
			tree: func() *DirectoryTreeNode {
				return &DirectoryTreeNode{
					Name:     ".",
					NodeType: "directory",
					Children: []*DirectoryTreeNode{{Name: "bad name", NodeType: "file", Role: "r"}},
				}
			},
			wantErrCode: "INVALID_NODE_NAME",
			wantNodeIn:  "/bad name",
		},
		{
			name: "R4 非法：name 含中文不匹配形态",
			tree: func() *DirectoryTreeNode {
				return &DirectoryTreeNode{
					Name:     ".",
					NodeType: "directory",
					Children: []*DirectoryTreeNode{{Name: "文档", NodeType: "file", Role: "r"}},
				}
			},
			wantErrCode: "INVALID_NODE_NAME",
			wantNodeIn:  "/文档",
		},

		// --- R5 TREE_TOO_DEEP ---
		{
			name: "R5 非法：第 6 层出现 directory 节点",
			tree: func() *DirectoryTreeNode {
				node := &DirectoryTreeNode{Name: "e", NodeType: "directory"}
				for _, seg := range []string{"d", "c", "b", "a"} {
					node = &DirectoryTreeNode{Name: seg, NodeType: "directory", Children: []*DirectoryTreeNode{node}}
				}
				return &DirectoryTreeNode{Name: ".", NodeType: "directory", Children: []*DirectoryTreeNode{node}}
			},
			wantErrCode: "TREE_TOO_DEEP",
			wantNodeIn:  "/a/b/c/d/e",
		},

		// --- R6 TREE_TOO_LARGE ---
		{
			name: "R6 非法：整树 JSON 序列化超 65536 字节",
			tree: func() *DirectoryTreeNode {
				tree := &DirectoryTreeNode{Name: ".", NodeType: "directory"}
				for i := 0; i < 40; i++ {
					tree.Children = append(tree.Children, &DirectoryTreeNode{
						Name:     fmt.Sprintf("file_%02d.md", i),
						NodeType: "file",
						Role:     "r",
						Summary:  strings.Repeat("a", 2000), // 单节点合法，总量超限
					})
				}
				return tree
			},
			wantErrCode: "TREE_TOO_LARGE",
			wantNodeIn:  "/",
		},

		// --- R7 INVALID_FILE_NODE / INVALID_SUMMARY_LENGTH ---
		{
			name: "R7 非法：file 节点携带 children",
			tree: func() *DirectoryTreeNode {
				return &DirectoryTreeNode{
					Name:     ".",
					NodeType: "directory",
					Children: []*DirectoryTreeNode{{
						Name: "AGENTS.md", NodeType: "file", Role: "r",
						Children: []*DirectoryTreeNode{{Name: "child.md", NodeType: "file", Role: "r"}},
					}},
				}
			},
			wantErrCode: "INVALID_FILE_NODE",
			wantNodeIn:  "/AGENTS.md",
		},
		{
			name: "R7 非法：file 节点 role 为空",
			tree: func() *DirectoryTreeNode {
				return &DirectoryTreeNode{
					Name:     ".",
					NodeType: "directory",
					Children: []*DirectoryTreeNode{{Name: "AGENTS.md", NodeType: "file"}},
				}
			},
			wantErrCode: "INVALID_FILE_NODE",
			wantNodeIn:  "/AGENTS.md",
		},
		{
			name: "R7 非法：file 节点 role 超 32 字符",
			tree: func() *DirectoryTreeNode {
				return &DirectoryTreeNode{
					Name:     ".",
					NodeType: "directory",
					Children: []*DirectoryTreeNode{{
						Name: "AGENTS.md", NodeType: "file", Role: strings.Repeat("a", 33),
					}},
				}
			},
			wantErrCode: "INVALID_FILE_NODE",
			wantNodeIn:  "/AGENTS.md",
		},
		{
			name: "R7 非法：summary 超 2000 字符",
			tree: func() *DirectoryTreeNode {
				return &DirectoryTreeNode{
					Name:     ".",
					NodeType: "directory",
					Children: []*DirectoryTreeNode{{
						Name: "AGENTS.md", NodeType: "file", Role: "r",
						Summary: strings.Repeat("a", 2001),
					}},
				}
			},
			wantErrCode: "INVALID_SUMMARY_LENGTH",
			wantNodeIn:  "/AGENTS.md",
		},

		// --- R8 INVALID_REF ---
		{
			name: "R8 非法：ref 不以 / 或 https:// 开头",
			tree: func() *DirectoryTreeNode {
				tree := validTree()
				tree.Children[0].Ref = "docs/AGENTS.md"
				return tree
			},
			wantErrCode: "INVALID_REF",
			wantNodeIn:  "/AGENTS.md",
		},
		{
			name: "R8 非法：ref 树内路径与实际路径不一致（/docs/phase 节点）",
			tree: func() *DirectoryTreeNode {
				tree := validTree()
				tree.Children[1].Children[0].Ref = "/docs"
				return tree
			},
			wantErrCode: "INVALID_REF",
			wantNodeIn:  "/docs/phase",
		},
		{
			name: "R8 非法：file 节点 ref 指向其他节点路径",
			tree: func() *DirectoryTreeNode {
				tree := validTree()
				tree.Children[0].Ref = "/docs/phase/README.md"
				return tree
			},
			wantErrCode: "INVALID_REF",
			wantNodeIn:  "/AGENTS.md",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateTreeStructure(tt.tree())
			if tt.wantErrCode == "" {
				if err != nil {
					t.Fatalf("期望合法，实际报错: %v", err)
				}
				return
			}
			if err == nil {
				t.Fatalf("期望错误码 %s，实际无错误", tt.wantErrCode)
			}
			if !strings.Contains(err.Error(), "["+tt.wantErrCode+"]") {
				t.Errorf("错误信息缺少稳定错误码 %s: %v", tt.wantErrCode, err)
			}
			if !strings.Contains(err.Error(), "(node: "+tt.wantNodeIn+")") && !strings.Contains(err.Error(), tt.wantNodeIn) {
				t.Errorf("错误信息缺少节点路径 %s: %v", tt.wantNodeIn, err)
			}
			if !errors.Is(err, ErrInvalidInput) {
				t.Errorf("错误未包装 ErrInvalidInput: %v", err)
			}
		})
	}
}

// TestValidateTreeForStatus 覆盖 R2 三时机语义。
func TestValidateTreeForStatus(t *testing.T) {
	tests := []struct {
		name        string
		tree        *DirectoryTreeNode
		status      StandardStatus
		wantErrCode string // 空串表示合法
	}{
		{
			name:   "R2 合法：draft + 单根空树",
			tree:   emptyRootTree(),
			status: StandardStatusDraft,
		},
		{
			name:        "R2 非法：active + 单根空树",
			tree:        emptyRootTree(),
			status:      StandardStatusActive,
			wantErrCode: "EMPTY_TREE_NOT_ALLOWED",
		},
		{
			name:        "R2 非法：retired + 单根空树",
			tree:        emptyRootTree(),
			status:      StandardStatusRetired,
			wantErrCode: "EMPTY_TREE_NOT_ALLOWED",
		},
		{
			name:   "R2 合法：active + 含 file 节点",
			tree:   validTree(),
			status: StandardStatusActive,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateTreeForStatus(tt.tree, tt.status)
			if tt.wantErrCode == "" {
				if err != nil {
					t.Fatalf("期望合法，实际报错: %v", err)
				}
				return
			}
			if err == nil {
				t.Fatalf("期望错误码 %s，实际无错误", tt.wantErrCode)
			}
			if !strings.Contains(err.Error(), "["+tt.wantErrCode+"]") {
				t.Errorf("错误信息缺少稳定错误码 %s: %v", tt.wantErrCode, err)
			}
			if !errors.Is(err, ErrInvalidInput) {
				t.Errorf("错误未包装 ErrInvalidInput: %v", err)
			}
		})
	}
}

// TestCountFileNodes 覆盖 file 节点计数（R2 判定输入）。
func TestCountFileNodes(t *testing.T) {
	if got := CountFileNodes(nil); got != 0 {
		t.Errorf("nil 树 file 计数 = %d, 期望 0", got)
	}
	if got := CountFileNodes(emptyRootTree()); got != 0 {
		t.Errorf("单根空树 file 计数 = %d, 期望 0", got)
	}
	if got := CountFileNodes(validTree()); got != 2 {
		t.Errorf("示例树 file 计数 = %d, 期望 2", got)
	}
}

// normalizeTree 将空 children 切片归一为 nil，用于序列化等价断言。
func normalizeTree(node *DirectoryTreeNode) {
	if node == nil {
		return
	}
	if len(node.Children) == 0 {
		node.Children = nil
	}
	for _, child := range node.Children {
		normalizeTree(child)
	}
}

// TestDirectoryTreeNodeJSONSerialization 覆盖序列化等价规则：
// children nil 与 [] 反序列化后等价；role / summary / ref omitempty。
func TestDirectoryTreeNodeJSONSerialization(t *testing.T) {
	// 1. children nil 与 [] 序列化后反序列化等价
	nilChildren := &DirectoryTreeNode{Name: ".", NodeType: "directory"}
	emptyChildren := &DirectoryTreeNode{Name: ".", NodeType: "directory", Children: []*DirectoryTreeNode{}}

	dataNil, err := json.Marshal(nilChildren)
	if err != nil {
		t.Fatalf("marshal nil-children 失败: %v", err)
	}
	dataEmpty, err := json.Marshal(emptyChildren)
	if err != nil {
		t.Fatalf("marshal empty-children 失败: %v", err)
	}

	var roundNil, roundEmpty DirectoryTreeNode
	if err := json.Unmarshal(dataNil, &roundNil); err != nil {
		t.Fatalf("unmarshal nil-children 失败: %v", err)
	}
	if err := json.Unmarshal(dataEmpty, &roundEmpty); err != nil {
		t.Fatalf("unmarshal empty-children 失败: %v", err)
	}
	normalizeTree(&roundNil)
	normalizeTree(&roundEmpty)
	if !reflect.DeepEqual(roundNil, roundEmpty) {
		t.Errorf("children nil 与 [] 反序列化后不等价: %+v vs %+v", roundNil, roundEmpty)
	}

	// 2. omitempty：role / summary / ref 空串序列化省略
	node := &DirectoryTreeNode{Name: "AGENTS.md", NodeType: "file", Role: "r"}
	data, err := json.Marshal(node)
	if err != nil {
		t.Fatalf("marshal 失败: %v", err)
	}
	payload := string(data)
	for _, key := range []string{`"summary"`, `"ref"`} {
		if strings.Contains(payload, key) {
			t.Errorf("空串字段未被 omitempty 省略：%s 出现在 %s", key, payload)
		}
	}
	// role 非空必须保留
	if !strings.Contains(payload, `"role"`) {
		t.Errorf("非空 role 被错误省略: %s", payload)
	}

	// 3. 反序列化缺失字段视为空字符串
	var got DirectoryTreeNode
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatalf("unmarshal 失败: %v", err)
	}
	if got.Summary != "" || got.Ref != "" {
		t.Errorf("缺失字段未归空字符串: summary=%q ref=%q", got.Summary, got.Ref)
	}

	// 4. directory 节点 role 空串同样省略
	dirNode := &DirectoryTreeNode{Name: ".", NodeType: "directory"}
	dirData, err := json.Marshal(dirNode)
	if err != nil {
		t.Fatalf("marshal directory 失败: %v", err)
	}
	if strings.Contains(string(dirData), `"role"`) {
		t.Errorf("directory 空 role 未被省略: %s", dirData)
	}
}
