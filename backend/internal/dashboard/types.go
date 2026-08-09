// Package dashboard 承载 Dashboard + Feedback 后端模块的全部业务实现。
//
// 分层语义（对齐 phase05-07 已冻结结论）：
//   - handler/     入口层：只负责承接 HTTP 请求与返回结果
//   - service/     业务编排层：三类聚合读编排与 Feedback Signal Card 归一化
//   - candidate/   外部连接层：跨模块 reader 接口定义与实现（Dashboard 自拥有）
//
// 本文件定义跨层共享的 API 消息结构。
// 约束：消息结构从 proto/psco/dashboard/v1/dashboard.proto 单向派生或显式对齐，
// 不直接暴露存储模型，不在 types.go 或 handler DTO 中新增 .proto 中不存在的业务字段语义。
package dashboard

import "time"

// ============================================================================
// 枚举类型
// ============================================================================
//
// 所有枚举使用 Go string 类型，常量值使用 snake_case，
// 对齐现有 ProductStatus 模式与 .proto 枚举的 JSON 小写形式。
// 编号语义对齐 phase05-08 §"当前版本字段编号方案必须冻结为机械映射"。

// FeedbackSignalFamily 反馈信号家族分组。
// 对齐 proto FeedbackSignalFamily：UNSPECIFIED / PENDING_DECISION / PRODUCT_ASSET_COVERAGE。
type FeedbackSignalFamily string

const (
	FeedbackSignalFamilyUnspecified             FeedbackSignalFamily = ""
	FeedbackSignalFamilyPendingDecision         FeedbackSignalFamily = "pending_decision"
	FeedbackSignalFamilyProductAssetCoverage    FeedbackSignalFamily = "product_asset_coverage"
)

// FeedbackSignalCode 反馈信号具体业务码。
// 对齐 proto FeedbackSignalCode。
// PRODUCT_MISSING_BOTH_BINDINGS 作为独立 code，不回退为隐式组合。
type FeedbackSignalCode string

const (
	FeedbackSignalCodeUnspecified                   FeedbackSignalCode = ""
	FeedbackSignalCodePendingDecision               FeedbackSignalCode = "pending_decision"
	FeedbackSignalCodeProductMissingBothBindings    FeedbackSignalCode = "product_missing_both_bindings"
	FeedbackSignalCodeProductMissingRepositoryBinding FeedbackSignalCode = "product_missing_repository_binding"
	FeedbackSignalCodeProductMissingModuleBinding   FeedbackSignalCode = "product_missing_module_binding"
)

// FeedbackSignalPriority 反馈信号优先级。
// 对齐 proto FeedbackSignalPriority。
// 编号顺序直接表达 P1 > P2 > P3 > P4。
// 用于 Go 层排序比较，数值越小优先级越高。
type FeedbackSignalPriority int

const (
	FeedbackSignalPriorityUnspecified                  FeedbackSignalPriority = 0
	FeedbackSignalPriorityP1PendingDecision            FeedbackSignalPriority = 1
	FeedbackSignalPriorityP2ProductMissingBothBindings FeedbackSignalPriority = 2
	FeedbackSignalPriorityP3ProductMissingRepository   FeedbackSignalPriority = 3
	FeedbackSignalPriorityP4ProductMissingModule       FeedbackSignalPriority = 4
)

// PriorityString 返回优先级的 JSON 字符串形式（对齐 proto 枚举的小写形式）。
func (p FeedbackSignalPriority) PriorityString() string {
	switch p {
	case FeedbackSignalPriorityP1PendingDecision:
		return "p1_pending_decision"
	case FeedbackSignalPriorityP2ProductMissingBothBindings:
		return "p2_product_missing_both_bindings"
	case FeedbackSignalPriorityP3ProductMissingRepository:
		return "p3_product_missing_repository_binding"
	case FeedbackSignalPriorityP4ProductMissingModule:
		return "p4_product_missing_module_binding"
	default:
		return ""
	}
}

// DashboardTargetType 反馈信号与活动项的 canonical 跳转目标类型。
// 对齐 proto DashboardTargetType。
type DashboardTargetType string

const (
	DashboardTargetTypeUnspecified        DashboardTargetType = ""
	DashboardTargetTypeDecisionDetail     DashboardTargetType = "decision_detail"
	DashboardTargetTypeDecisionList       DashboardTargetType = "decision_list"
	DashboardTargetTypeProductDetail      DashboardTargetType = "product_detail"
	DashboardTargetTypeModuleDetail       DashboardTargetType = "module_detail"
	DashboardTargetTypeRepositoryDetail   DashboardTargetType = "repository_detail"
)

// RecentActivityType 最近活动项的业务类型。
// 对齐 proto RecentActivityType。
type RecentActivityType string

const (
	RecentActivityTypeUnspecified                RecentActivityType = ""
	RecentActivityTypeModule                     RecentActivityType = "module"
	RecentActivityTypeRelease                    RecentActivityType = "release"
	RecentActivityTypeProduct                    RecentActivityType = "product"
	RecentActivityTypeRepository                 RecentActivityType = "repository"
	RecentActivityTypeDecision                   RecentActivityType = "decision"
	RecentActivityTypeProductModuleBinding       RecentActivityType = "product_module_binding"
	RecentActivityTypeProductRepositoryBinding   RecentActivityType = "product_repository_binding"
	RecentActivityTypeModuleRepositoryBinding    RecentActivityType = "module_repository_binding"
)

// ============================================================================
// 核心消息 DTO
// ============================================================================

// DashboardOverview 概览卡片主聚合读模型。
// 对齐 proto DashboardOverview。
// 只服务概览卡片与系统状态判定，不混入 FeedbackSignal 或 RecentActivityItem。
type DashboardOverview struct {
	ModuleCount               int `json:"module_count"`
	ProductCount              int `json:"product_count"`
	RepositoryCount           int `json:"repository_count"`
	DecisionCount             int `json:"decision_count"`
	ProductWithRepositoryCount int `json:"product_with_repository_count"`
	ProductWithModuleCount    int `json:"product_with_module_count"`
}

// FeedbackSignal 统一反馈主队列的单值卡片模型。
// 对齐 proto FeedbackSignal。
// 同时承接"解释缺口"与"导航到 canonical owner"两类语义。
// Priority 使用 int 类型用于排序，JSON 输出时由 service 层确保 PriorityString 对齐。
type FeedbackSignal struct {
	SignalFamily  FeedbackSignalFamily  `json:"signal_family"`
	SignalCode    FeedbackSignalCode    `json:"signal_code"`
	Priority      FeedbackSignalPriority `json:"-"` // 不直接序列化，由 PriorityLabel 承接
	PriorityLabel string                `json:"priority"`
	Title         string                `json:"title"`
	Summary       string                `json:"summary"`
	ActionLabel   string                `json:"action_label"`
	TargetType    DashboardTargetType   `json:"target_type"`
	TargetID      string                `json:"target_id"`
	TargetLabel   string                `json:"target_label"`
	// CreatedAt 用于同优先级内排序回退（phase05-02 冻结），不暴露到 JSON
	CreatedAt     time.Time             `json:"-"`
}

// ProductAssetCoverageSummary 资产缺口补充摘要。
// 对齐 proto ProductAssetCoverageSummary。
// missing_both_bindings_count 作为独立计数字段，不回退为隐式组合。
type ProductAssetCoverageSummary struct {
	FullyBoundProductCount      int               `json:"fully_bound_product_count"`
	MissingBothBindingsCount    int               `json:"missing_both_bindings_count"`
	MissingRepositoryBindingCount int             `json:"missing_repository_binding_count"`
	MissingModuleBindingCount   int               `json:"missing_module_binding_count"`
	RepresentativeSignals       []FeedbackSignal  `json:"representative_signals"`
}

// RecentActivityItem 独立活动流单值项。
// 对齐 proto RecentActivityItem。
// activity_at 使用 time.Time，对应 proto google.protobuf.Timestamp。
type RecentActivityItem struct {
	ActivityType RecentActivityType   `json:"activity_type"`
	ActivityAt   time.Time            `json:"activity_at"`
	TargetType   DashboardTargetType  `json:"target_type"`
	TargetID     string               `json:"target_id"`
	TargetLabel  string               `json:"target_label"`
}

// ============================================================================
// 响应 DTO
// ============================================================================

// DashboardOverviewReadResult DashboardOverviewRead 的响应结构。
// 对齐 proto GetDashboardOverviewResponse：单一 overview 字段包装主读模型。
// handler 必须返回此包络结构，不得直接返回裸 DashboardOverview，
// 以保证 HTTP JSON 形状与 dashboard.proto 冻结的唯一合同源一致。
type DashboardOverviewReadResult struct {
	Overview *DashboardOverview `json:"overview"`
}

// FeedbackSignalReadResult FeedbackSignalRead 的响应结构。
// 对齐 proto GetFeedbackSignalsResponse：current_focus_signals + asset_feedback_summary。
type FeedbackSignalReadResult struct {
	CurrentFocusSignals   []FeedbackSignal             `json:"current_focus_signals"`
	AssetFeedbackSummary  ProductAssetCoverageSummary  `json:"asset_feedback_summary"`
}

// RecentActivityReadResult RecentActivityRead 的响应结构。
// 对齐 proto GetRecentActivitiesResponse：单一 activities 字段包装活动项列表。
// handler 必须返回此包络结构，不得直接返回裸 []RecentActivityItem，
// 以保证 HTTP JSON 形状与 dashboard.proto 冻结的唯一合同源一致。
type RecentActivityReadResult struct {
	Activities []RecentActivityItem `json:"activities"`
}

// Ensure FeedbackSignalReadResult 的空态语义：
// 空主队列时 current_focus_signals 为空列表（非 nil），
// 空缺口时 asset_feedback_summary 返回完整结构与零计数。
func NewEmptyFeedbackSignalReadResult() *FeedbackSignalReadResult {
	return &FeedbackSignalReadResult{
		CurrentFocusSignals: []FeedbackSignal{},
		AssetFeedbackSummary: ProductAssetCoverageSummary{
			RepresentativeSignals: []FeedbackSignal{},
		},
	}
}

// Ensure RecentActivityReadResult 的空态语义：
// 空活动流时 activities 为空列表（非 nil），不映射为错误。
func NewEmptyRecentActivityReadResult() *RecentActivityReadResult {
	return &RecentActivityReadResult{
		Activities: []RecentActivityItem{},
	}
}
