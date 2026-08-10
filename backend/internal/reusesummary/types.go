// Package reusesummary 承载复用感知（Reuse Summary）后端模块的全部业务实现。
//
// 分层语义（对齐 phase06-14 spec §"Phase06 后端模块必须按现有主线结构落地"）：
//   - handler/     入口层：只负责承接 HTTP 请求与返回结果
//   - service/     业务编排层：复用聚合、排序、裁剪与最新状态语义
//   - candidate/   外部连接层：跨模块复用 reader（ReuseSummary 自拥有）
//
// 当前阶段 ReuseSummary 实现为读时聚合，不引入异步统计表或离线聚合作业。
//
// 本文件定义跨层共享的 API 消息结构。
// 约束：消息结构从 proto/psco/reuse_summary/v1/reuse_summary.proto 单向派生或显式对齐。
package reusesummary

import "time"

// ============================================================================
// 枚举类型
// ============================================================================

// ReuseSummaryScope 复用感知读取的页面作用域。
// 对齐 proto ReuseSummaryScope。
type ReuseSummaryScope string

const (
	ReuseSummaryScopeUnspecified   ReuseSummaryScope = ""
	ReuseSummaryScopeDashboard     ReuseSummaryScope = "dashboard"
	ReuseSummaryScopeModuleDetail  ReuseSummaryScope = "module_detail"
	ReuseSummaryScopeProductDetail ReuseSummaryScope = "product_detail"
)

// ============================================================================
// 核心消息 DTO
// ============================================================================

// ModuleReuseSummary 模块复用摘要单值项。
// 对齐 proto ModuleReuseSummary。
// 统计口径冻结为"一个 Module 当前被多少 Product 直接复用"。
type ModuleReuseSummary struct {
	ModuleID          string     `json:"module_id"`
	ReuseProductCount int        `json:"reuse_product_count"`
	LatestReuseAt     *time.Time `json:"latest_reuse_at"`
	ExplanationText   string     `json:"explanation_text"`
}

// CapabilitySummary 能力聚合摘要单值项。
// 对齐 proto CapabilitySummary。
// 事实来源冻结为 Module.capability_key + capability_label 映射。
type CapabilitySummary struct {
	CapabilityKey           string     `json:"capability_key"`
	CapabilityLabel         string     `json:"capability_label"`
	SupportingModuleCount   int        `json:"supporting_module_count"`
	LatestCapabilityUpdateAt *time.Time `json:"latest_capability_update_at"`
	EmptyStateText          string     `json:"empty_state_text"`
}

// ============================================================================
// 响应 DTO
// ============================================================================

// ReuseSummaryReadResult GetReuseSummary 的响应结构。
// 对齐 proto GetReuseSummaryResponse：module_reuse_summary[] + capability_summary[]。
// handler 必须返回此包络结构。
type ReuseSummaryReadResult struct {
	ModuleReuseSummary []ModuleReuseSummary `json:"module_reuse_summary"`
	CapabilitySummary  []CapabilitySummary  `json:"capability_summary"`
}

// NewEmptyReuseSummaryReadResult 构造空态响应。
// 空结果表现为正常空列表响应，不映射为错误。
func NewEmptyReuseSummaryReadResult() *ReuseSummaryReadResult {
	return &ReuseSummaryReadResult{
		ModuleReuseSummary: []ModuleReuseSummary{},
		CapabilitySummary:  []CapabilitySummary{},
	}
}
