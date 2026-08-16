// Package projectcontext 承载最小只读项目上下文聚合读取能力。
//
// 分层语义（对齐 phase11-07 已冻结结论）：
//   - service/     业务编排层：只读聚合编排
//   - candidate/   外部连接层：跨模块 reader 接口定义与实现（projectcontext 自拥有）
//   - connect/     入口层：Connect handler 实现，负责 proto request 解包 → service 调用 → proto response 组装
//
// 本文件定义跨层共享的 API 消息结构。
// 约束：消息结构从 proto/psco/project_context/v1/project_context.proto 单向派生或显式对齐，
// 不直接暴露存储模型，不在 types.go 中新增 .proto 中不存在的业务字段语义。
package projectcontext

// ============================================================================
// 核心消息 DTO
// ============================================================================

// RepositorySummary 当前仓库身份摘要。
// 对齐 proto RepositorySummary。
type RepositorySummary struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Provider    string `json:"provider"`
	URL         string `json:"url"`
	Description string `json:"description"`
}

// ProductSummary 关联产品摘要。
// 对齐 proto ProductSummary。
type ProductSummary struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

// ModuleSummary 关联模块摘要。
// 对齐 proto ModuleSummary。
type ModuleSummary struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

// DecisionSummary 关联决策摘要。
// 对齐 proto DecisionSummary。
// HitSources 标识该 Decision 通过哪类 canonical 关系被命中。
type DecisionSummary struct {
	ID         string   `json:"id"`
	Title      string   `json:"title"`
	Status     string   `json:"status"`
	Context    string   `json:"context"`
	HitSources []string `json:"hit_sources"`
	CreatedAt  string   `json:"created_at"`
}

// RuleEntry 规则与约束入口。
// 对齐 proto RuleEntry。
type RuleEntry struct {
	Key       string `json:"key"`
	Label     string `json:"label"`
	Summary   string `json:"summary"`
	EntryRef  string `json:"entry_ref"`
	EntryKind string `json:"entry_kind"`
}

// PhaseEntry 当前 phase 相关的文档入口。
// 对齐 proto PhaseEntry。
type PhaseEntry struct {
	Phase         string `json:"phase"`
	Label         string `json:"label"`
	StatusSummary string `json:"status_summary"`
	EntryRef      string `json:"entry_ref"`
	EntryKind     string `json:"entry_kind"`
}

// ============================================================================
// 响应 DTO
// ============================================================================

// ProjectContextReadResult GetProjectContext 的响应结构。
// 对齐 proto GetProjectContextResponse。
type ProjectContextReadResult struct {
	Repository *RepositorySummary `json:"repository"`
	Product    *ProductSummary    `json:"product"`
	Modules    []ModuleSummary    `json:"modules"`
	Decisions  []DecisionSummary  `json:"decisions"`
	Rules      []RuleEntry        `json:"rules"`
	Phases     []PhaseEntry       `json:"phases"`
}
