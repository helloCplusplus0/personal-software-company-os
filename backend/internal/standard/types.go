// Package standard 承载全局规范实体结构化写读能力。
//
// 分层语义（对齐 phase14-04 已冻结结论）：
//   - connect/     入口层：Connect handler，proto request 解包 → service 调用 → proto response 组装
//   - service/     业务编排层：写读编排、status 归一、树校验触发与八格绑定矩阵校验
//   - repository/  持久化层：PostgreSQL 读写（standards / standard_revisions / standard_bindings 三表）
//   - candidate/   外部连接层：跨模块 target 存在性校验 reader（standard 自拥有）
//
// 本文件定义跨层共享的领域模型与受控枚举。
// 约束：消息结构从 proto/psco/standard/v1/standard.proto 单向派生或显式对齐，
// 不直接暴露存储模型，不在本文件 import 生成 pb 包（proto ↔ domain 转换收敛在 connect 层）。
package standard

import "time"

// ============================================================================
// 受控枚举
// ============================================================================

// StandardStatus 规范生命周期（受控枚举，对齐 DDL status CHECK）。
type StandardStatus string

const (
	StandardStatusDraft   StandardStatus = "draft"
	StandardStatusActive  StandardStatus = "active"
	StandardStatusRetired StandardStatus = "retired"
)

// NodeType 目录树节点类型（受控枚举）。
// 与 pb 包的 int32 枚举 NodeType 同名不同包，合法共存；string 形态用于
// DirectoryTreeNode jsonb 与领域层判定。
type NodeType string

const (
	NodeTypeDirectory NodeType = "directory"
	NodeTypeFile      NodeType = "file"
)

// BindingTargetType 绑定目标类型（受控枚举，可扩展：扩 enum + 登记 phase14-02 八格矩阵）。
type BindingTargetType string

const (
	BindingTargetRepository BindingTargetType = "repository"
	BindingTargetProduct    BindingTargetType = "product"
	BindingTargetDecision   BindingTargetType = "decision"
	BindingTargetModule     BindingTargetType = "module"
)

// BindingRole 绑定角色（受控枚举，可扩展；组合合法性按 phase14-02 八格矩阵）。
type BindingRole string

const (
	BindingRoleTemplateSource BindingRole = "template_source"
	BindingRoleAdopts         BindingRole = "adopts"
)

// ============================================================================
// 目录树结构（DirectoryTreeNode jsonb 单值映射，phase14-03 冻结结构）
// ============================================================================

// DirectoryTreeNode 目录树节点，与 standards.directory_tree jsonb 单值映射。
// 结构逐字冻结自 phase14-03 §ADDED-2，不得改动任何 tag。
type DirectoryTreeNode struct {
	Name     string               `json:"name"`
	NodeType string               `json:"node_type"`         // "directory" | "file"
	Role     string               `json:"role,omitempty"`    // file 必填（校验层保证）；directory 可空
	Summary  string               `json:"summary,omitempty"` // 结构化摘要（裁决①承接位）
	Ref      string               `json:"ref,omitempty"`     // 定位引用：/ 开头树内路径或 https:// URL
	Children []*DirectoryTreeNode `json:"children"`
}

// ============================================================================
// 读模型
// ============================================================================

// StandardReadResult 规范主记录读取结果（standards 表投影 + 整树）。
type StandardReadResult struct {
	ID            string             `json:"id"`
	Name          string             `json:"name"`
	Description   string             `json:"description"`
	Status        StandardStatus     `json:"status"`
	DirectoryTree *DirectoryTreeNode `json:"directory_tree"`
	CreatedAt     time.Time          `json:"created_at"`
	UpdatedAt     time.Time          `json:"updated_at"`
}

// StandardBindingReadResult 绑定关系读取结果（standard_bindings 表投影）。
type StandardBindingReadResult struct {
	ID         string            `json:"id"`
	StandardID string            `json:"standard_id"`
	TargetType BindingTargetType `json:"target_type"`
	TargetID   string            `json:"target_id"`
	Role       BindingRole       `json:"role"`
	Note       string            `json:"note,omitempty"`
	CreatedAt  time.Time         `json:"created_at"`
}

// StandardRevisionReadResult 演进留痕读取结果（standard_revisions 表投影，只追加）。
type StandardRevisionReadResult struct {
	ID            string    `json:"id"`
	StandardID    string    `json:"standard_id"`
	ChangeSummary string    `json:"change_summary"`
	CreatedAt     time.Time `json:"created_at"`
}

// ============================================================================
// 写模型
// ============================================================================

// CreateStandardInput 创建规范输入。
// Status 已由 service 归一化：空值归 draft。
type CreateStandardInput struct {
	Name          string             `json:"name"`
	Description   string             `json:"description"`
	DirectoryTree *DirectoryTreeNode `json:"directory_tree"`
	Status        StandardStatus     `json:"status"`
}

// UpdateStandardInput 更新规范输入（整树原子替换语义，无部分更新）。
// Name / Description / Status 为 nil 时不变更（description 非 nil 含空串时写入清空）。
type UpdateStandardInput struct {
	StandardID    string             `json:"standard_id"`
	Name          *string            `json:"name"`
	Description   *string            `json:"description"`
	Status        *StandardStatus    `json:"status"`
	DirectoryTree *DirectoryTreeNode `json:"directory_tree"`
	ChangeSummary string             `json:"change_summary"`
}

// BindStandardInput 建立绑定输入。
type BindStandardInput struct {
	StandardID string            `json:"standard_id"`
	TargetType BindingTargetType `json:"target_type"`
	TargetID   string            `json:"target_id"`
	Role       BindingRole       `json:"role"`
	Note       *string           `json:"note"`
}

// UnbindStandardInput 解除绑定输入（四元组定位，note 不参与）。
type UnbindStandardInput struct {
	StandardID string            `json:"standard_id"`
	TargetType BindingTargetType `json:"target_type"`
	TargetID   string            `json:"target_id"`
	Role       BindingRole       `json:"role"`
}
