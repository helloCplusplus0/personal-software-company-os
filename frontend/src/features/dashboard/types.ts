/**
 * Dashboard 前端类型定义
 *
 * phase05-13 §"Dashboard 类型与 API 适配层必须从 proto envelope 单向派生"
 *
 * 字段语义对齐：
 * - 后端 HTTP JSON：snake_case（backend/internal/dashboard/types.go 的 json tag）
 * - 前端类型：snake_case，与后端 JSON tag 直接对齐
 * - 不直接消费 frontend/src/gen/proto/psco/dashboard/v1/dashboard_pb.ts 的 camelCase，
 *   因为 HTTP JSON 是后端响应的实际形状，proto TS 类型仅作为合同源追溯
 *
 * 三个响应 envelope 直接对齐 phase05-12 后端 handler 返回结构：
 *   - DashboardOverviewReadResult     { overview: DashboardOverview }
 *   - FeedbackSignalReadResult        { current_focus_signals, asset_feedback_summary }
 *   - RecentActivityReadResult        { activities: RecentActivityItem[] }
 */

/**
 * DashboardTargetType — 反馈信号与活动项的 canonical 跳转目标类型。
 * 对齐 backend DashboardTargetType 常量与 proto DashboardTargetType 枚举。
 */
export type DashboardTargetType =
  | 'decision_detail'
  | 'decision_list'
  | 'product_detail'
  | 'module_detail'
  | 'repository_detail'

/**
 * FeedbackSignalFamily — 反馈信号家族分组。
 * 对齐 backend FeedbackSignalFamily 常量。
 */
export type FeedbackSignalFamily =
  | 'pending_decision'
  | 'product_asset_coverage'

/**
 * FeedbackSignalCode — 反馈信号具体业务码。
 * product_missing_both_bindings 作为独立 code，不回退为隐式组合。
 * 对齐 backend FeedbackSignalCode 常量。
 */
export type FeedbackSignalCode =
  | 'pending_decision'
  | 'product_missing_both_bindings'
  | 'product_missing_repository_binding'
  | 'product_missing_module_binding'

/**
 * FeedbackSignalPriority — 反馈信号优先级字符串形式。
 * 后端 service 层通过 PriorityString() 输出，对齐 proto 枚举的小写形式。
 * 数值越小优先级越高（P1 > P2 > P3 > P4）。
 */
export type FeedbackSignalPriority =
  | 'p1_pending_decision'
  | 'p2_product_missing_both_bindings'
  | 'p3_product_missing_repository_binding'
  | 'p4_product_missing_module_binding'

/**
 * RecentActivityType — 最近活动项的业务类型。
 * binding 必须拆分为三类，不得以笼统 'binding' 保留歧义。
 * 对齐 backend RecentActivityType 常量。
 */
export type RecentActivityType =
  | 'module'
  | 'release'
  | 'product'
  | 'repository'
  | 'decision'
  | 'product_module_binding'
  | 'product_repository_binding'
  | 'module_repository_binding'

/**
 * DashboardOverview — 概览卡片主聚合读模型。
 * 对齐 backend DashboardOverview 与 proto DashboardOverview。
 * 只服务概览卡片与系统状态判定前提，不混入 FeedbackSignal 或 RecentActivityItem。
 */
export interface DashboardOverview {
  module_count: number
  product_count: number
  repository_count: number
  decision_count: number
  product_with_repository_count: number
  product_with_module_count: number
}

/**
 * FeedbackSignal — 统一反馈主队列的单值卡片模型。
 * 对齐 backend FeedbackSignal 与 proto FeedbackSignal。
 * 同时承接"解释缺口"与"导航到 canonical owner"两类语义，不得减少字段。
 *
 * 注意：后端 JSON 中 priority 字段是 PriorityString() 输出的字符串形式
 * （如 'p1_pending_decision'），而非 int 数值。后端 struct 内部 Priority int 字段
 * 标记为 `json:"-"` 不直接序列化。
 */
export interface FeedbackSignal {
  signal_family: FeedbackSignalFamily
  signal_code: FeedbackSignalCode
  priority: FeedbackSignalPriority
  title: string
  summary: string
  action_label: string
  target_type: DashboardTargetType
  target_id: string
  target_label: string
}

/**
 * ProductAssetCoverageSummary — 资产缺口补充摘要。
 * 对齐 backend ProductAssetCoverageSummary 与 proto ProductAssetCoverageSummary。
 * missing_both_bindings_count 作为独立计数字段，不回退为隐式组合。
 */
export interface ProductAssetCoverageSummary {
  fully_bound_product_count: number
  missing_both_bindings_count: number
  missing_repository_binding_count: number
  missing_module_binding_count: number
  representative_signals: FeedbackSignal[]
}

/**
 * RecentActivityItem — 独立活动流单值项。
 * 对齐 backend RecentActivityItem 与 proto RecentActivityItem。
 * activity_at 使用 ISO 8601 时间字符串，由后端 time.Time 序列化得到。
 */
export interface RecentActivityItem {
  activity_type: RecentActivityType
  activity_at: string
  target_type: DashboardTargetType
  target_id: string
  target_label: string
}

// ============================================================================
// 响应 envelope 类型（对齐 phase05-12 后端 handler 响应结构）
// ============================================================================

/**
 * DashboardOverviewResponse — GET /api/dashboard/overview 响应包络。
 * 对齐 backend DashboardOverviewReadResult 与 proto GetDashboardOverviewResponse。
 */
export interface DashboardOverviewResponse {
  overview: DashboardOverview
}

/**
 * FeedbackSignalsResponse — GET /api/dashboard/feedback-signals 响应包络。
 * 对齐 backend FeedbackSignalReadResult 与 proto GetFeedbackSignalsResponse。
 *
 * 空态语义：
 * - current_focus_signals 为空列表（非 nil）表示无主队列信号
 * - asset_feedback_summary 在空态下返回完整结构与零计数、代表项为空列表
 */
export interface FeedbackSignalsResponse {
  current_focus_signals: FeedbackSignal[]
  asset_feedback_summary: ProductAssetCoverageSummary
}

/**
 * RecentActivitiesResponse — GET /api/dashboard/recent-activities 响应包络。
 * 对齐 backend RecentActivityReadResult 与 proto GetRecentActivitiesResponse。
 * activities 为空列表（非 nil）表示无活动项，不视为错误。
 */
export interface RecentActivitiesResponse {
  activities: RecentActivityItem[]
}

// ============================================================================
// Dashboard 来源参数（路由搜索参数层）
// ============================================================================

/**
 * DashboardSection — Dashboard 来源区块标记。
 * 对齐 phase05-03 / phase05-10 §8.2 已冻结的允许取值。
 */
export type DashboardSection =
  | 'overview'
  | 'current-focus'
  | 'asset-feedback'
  | 'recent-activity'
  | 'empty-state'

/**
 * DashboardSourceSearch — Dashboard 来源参数在路由搜索参数层的形状。
 * 三字段均为可选，由既有 canonical 路由的 validateSearch 扩展承接。
 */
export interface DashboardSourceSearch {
  fromDashboard?: boolean
  dashboardSection?: DashboardSection
  dashboardReturnTo?: string
}

// ============================================================================
// Export / Backup 类型（对齐 proto export/v1/export.proto 与 backup/v1/backup.proto）
// 从 sovereignty-api-adapter.ts 迁移到此处，作为 canonical 类型定义源。
// ============================================================================

export type ExportAssetScope =
  | 'products'
  | 'modules'
  | 'releases'
  | 'repositories'
  | 'decisions'
  | 'decision_links'
  | 'product_modules'
  | 'product_repositories'
  | 'module_repositories'

export type ExportResultStatus = 'success' | 'in_progress' | 'failed'

export interface ExportSnapshot {
  asset_scope: ExportAssetScope[]
  created_at: string
  result_status: ExportResultStatus
  result_summary: string
}

export interface ExportSnapshotReadResult {
  snapshot: ExportSnapshot | null
}

export type BackupVerifiedStatus = 'unverified' | 'verified' | 'verify_failed'

export type VerifyFailureCode = 'manifest_missing' | 'coverage_incomplete' | 'schema_mismatch'

export interface ManifestSummary {
  manifest_version: string
  total_asset_entries: number
  covered_asset_entries: number
}

export interface AssetCoverageEntry {
  asset_scope: string
  covered: boolean
}

export interface SchemaVersionPrerequisite {
  schema_version: string
  instance_version: string
  prerequisite_checkable: boolean
}

export interface BackupSnapshot {
  created_at: string
  manifest_summary: ManifestSummary | null
  asset_coverage: AssetCoverageEntry[]
  schema_version_prerequisite: SchemaVersionPrerequisite | null
  verified_status: BackupVerifiedStatus
  /** 仅在 verified_status = verify_failed 时有值 */
  verify_failure_code?: VerifyFailureCode
}

export interface BackupSnapshotReadResult {
  snapshot: BackupSnapshot | null
}
