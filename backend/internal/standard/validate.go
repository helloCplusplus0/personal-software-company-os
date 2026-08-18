// DirectoryTreeNode 树校验（R1-R8，phase14-03 冻结规则）。
//
// 本文件承接写路径统一执行的 8 条树结构校验：
//   - R1 INVALID_TREE_ROOT        根结构（单根 directory、name="."、role 空）
//   - R2 EMPTY_TREE_NOT_ALLOWED   非 draft 状态树内必须含至少一个 file 节点
//   - R3 DUPLICATE_SIBLING_NAME   同层兄弟 name 查重（大小写敏感，根豁免）
//   - R4 INVALID_NODE_NAME        非根 name 匹配 ^[A-Za-z0-9._-]{1,64}$
//   - R5 TREE_TOO_DEEP            深度 ≤6；第 6 层只允许 file
//   - R6 TREE_TOO_LARGE           整树 JSON 序列化 ≤65536 字节
//   - R7 INVALID_FILE_NODE        file 节点无 children、role 长度 1-32
//     INVALID_SUMMARY_LENGTH      任意节点 summary ≤2000 字符
//   - R8 INVALID_REF              ref 以 "/" 或 "https://" 开头；
//     "/" 开头时必须与节点实际路径完全一致
//
// 所有校验错误统一包装为 ErrInvalidInput（InvalidArgument），
// 错误信息携带稳定错误码与自根起 "/" 连接的节点路径，供前端定位展示。
//
// 文件落点：backend/internal/standard/validate.go
package standard

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
	"unicode/utf8"
)

// 树校验稳定错误码（逐字冻结自 phase14-03，前端按码定位）。
const (
	errCodeInvalidTreeRoot      = "INVALID_TREE_ROOT"
	errCodeEmptyTreeNotAllowed  = "EMPTY_TREE_NOT_ALLOWED"
	errCodeDuplicateSiblingName = "DUPLICATE_SIBLING_NAME"
	errCodeInvalidNodeName      = "INVALID_NODE_NAME"
	errCodeTreeTooDeep          = "TREE_TOO_DEEP"
	errCodeTreeTooLarge         = "TREE_TOO_LARGE"
	errCodeInvalidFileNode      = "INVALID_FILE_NODE"
	errCodeInvalidSummaryLength = "INVALID_SUMMARY_LENGTH"
	errCodeInvalidRef           = "INVALID_REF"
)

// 树校验冻结阈值（phase14-03）。
const (
	maxTreeDepth     = 6     // 根为第 1 层
	maxTreeJSONBytes = 65536 // 整树 JSON 序列化字节上限
	maxSummaryRunes  = 2000  // summary 字符上限
	maxRoleRunes     = 32    // file 节点 role 字符上限
)

// nodeNamePattern 非根节点 name 合法形态（R4）。
var nodeNamePattern = regexp.MustCompile(`^[A-Za-z0-9._-]{1,64}$`)

// ValidateTreeStructure 执行与 status 无关的结构校验（R1/R3/R4/R5/R6/R7/R8）。
// DFS 按 children 切片原序遍历，报第一个错误。
func ValidateTreeStructure(tree *DirectoryTreeNode) error {
	// R1：根结构（nil / 根类型 / 根名 / 根 role）
	if tree == nil {
		return fmt.Errorf("%w: [%s] directory_tree root is required (node: /)", ErrInvalidInput, errCodeInvalidTreeRoot)
	}
	if tree.NodeType != string(NodeTypeDirectory) {
		return fmt.Errorf("%w: [%s] root node_type must be directory (node: /)", ErrInvalidInput, errCodeInvalidTreeRoot)
	}
	if tree.Name != "." {
		return fmt.Errorf("%w: [%s] root name must be %q (node: /)", ErrInvalidInput, errCodeInvalidTreeRoot, ".")
	}
	if tree.Role != "" {
		return fmt.Errorf("%w: [%s] root role must be empty (node: /)", ErrInvalidInput, errCodeInvalidTreeRoot)
	}

	// R6：整树 JSON 序列化字节上限
	data, err := json.Marshal(tree)
	if err != nil {
		return fmt.Errorf("%w: [%s] directory_tree serialization failed: %v (node: /)", ErrInvalidInput, errCodeTreeTooLarge, err)
	}
	if len(data) > maxTreeJSONBytes {
		return fmt.Errorf("%w: [%s] directory_tree serialized size %d exceeds %d bytes (node: /)", ErrInvalidInput, errCodeTreeTooLarge, len(data), maxTreeJSONBytes)
	}

	// R3/R4/R5/R7/R8：自根 DFS
	return validateNode(tree, "/", 1)
}

// ValidateTreeForStatus 执行 R2 状态时机校验：
// status 非 draft 时树内必须含至少一个 file 节点。
func ValidateTreeForStatus(tree *DirectoryTreeNode, status StandardStatus) error {
	if status == StandardStatusDraft {
		return nil
	}
	if CountFileNodes(tree) == 0 {
		return fmt.Errorf("%w: [%s] standard with status %q must contain at least one file node (node: /)", ErrInvalidInput, errCodeEmptyTreeNotAllowed, status)
	}
	return nil
}

// CountFileNodes 统计树内 file 节点数（R2 判定输入）。
func CountFileNodes(tree *DirectoryTreeNode) int {
	if tree == nil {
		return 0
	}
	count := 0
	if tree.NodeType == string(NodeTypeFile) {
		count = 1
	}
	for _, child := range tree.Children {
		count += CountFileNodes(child)
	}
	return count
}

// validateNode 递归校验单个节点。path 为该节点自根起的实际路径（根为 "/"），
// depth 为该节点所在层（根为第 1 层）。
func validateNode(node *DirectoryTreeNode, path string, depth int) error {
	// R5：深度上限；第 6 层只允许 file
	if depth > maxTreeDepth {
		return fmt.Errorf("%w: [%s] tree exceeds maximum depth %d (node: %s)", ErrInvalidInput, errCodeTreeTooDeep, maxTreeDepth, path)
	}
	if depth == maxTreeDepth && node.NodeType == string(NodeTypeDirectory) {
		return fmt.Errorf("%w: [%s] level %d only allows file nodes (node: %s)", ErrInvalidInput, errCodeTreeTooDeep, maxTreeDepth, path)
	}

	// R4：非根 name 形态
	if depth > 1 && !nodeNamePattern.MatchString(node.Name) {
		return fmt.Errorf("%w: [%s] node name %q must match ^[A-Za-z0-9._-]{1,64}$ (node: %s)", ErrInvalidInput, errCodeInvalidNodeName, node.Name, path)
	}

	// R7：file 节点约束（无 children、role 长度 1-32）
	if node.NodeType == string(NodeTypeFile) {
		if len(node.Children) > 0 {
			return fmt.Errorf("%w: [%s] file node must not have children (node: %s)", ErrInvalidInput, errCodeInvalidFileNode, path)
		}
		roleLen := utf8.RuneCountInString(node.Role)
		if roleLen < 1 || roleLen > maxRoleRunes {
			return fmt.Errorf("%w: [%s] file node role length must be 1-%d (node: %s)", ErrInvalidInput, errCodeInvalidFileNode, maxRoleRunes, path)
		}
	}

	// R7：summary 长度（任意节点）
	if utf8.RuneCountInString(node.Summary) > maxSummaryRunes {
		return fmt.Errorf("%w: [%s] summary exceeds %d characters (node: %s)", ErrInvalidInput, errCodeInvalidSummaryLength, maxSummaryRunes, path)
	}

	// R8：ref 形态与路径一致性
	if node.Ref != "" {
		if !strings.HasPrefix(node.Ref, "/") && !strings.HasPrefix(node.Ref, "https://") {
			return fmt.Errorf("%w: [%s] ref must start with %q or %q (node: %s)", ErrInvalidInput, errCodeInvalidRef, "/", "https://", path)
		}
		if strings.HasPrefix(node.Ref, "/") && node.Ref != path {
			return fmt.Errorf("%w: [%s] ref %q must equal node path %q (node: %s)", ErrInvalidInput, errCodeInvalidRef, node.Ref, path, path)
		}
	}

	// R3：同层兄弟 name 查重（大小写敏感；根单节点天然豁免）
	seen := make(map[string]struct{}, len(node.Children))
	for _, child := range node.Children {
		if _, dup := seen[child.Name]; dup {
			return fmt.Errorf("%w: [%s] duplicate sibling name %q (node: %s)", ErrInvalidInput, errCodeDuplicateSiblingName, child.Name, childPath(path, child.Name))
		}
		seen[child.Name] = struct{}{}
	}

	// DFS：按 children 切片原序递归
	for _, child := range node.Children {
		if err := validateNode(child, childPath(path, child.Name), depth+1); err != nil {
			return err
		}
	}
	return nil
}

// childPath 计算子节点实际路径：根下第一段为 "/name"，此后逐段追加。
func childPath(parentPath, name string) string {
	if parentPath == "/" {
		return "/" + name
	}
	return parentPath + "/" + name
}
