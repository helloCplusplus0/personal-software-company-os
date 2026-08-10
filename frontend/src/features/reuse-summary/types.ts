/**
 * Reuse Summary 前端类型定义
 *
 * 对齐 phase06-14 后端 reusesummary 模块响应结构。
 * 字段语义从 .proto -> HTTP DTO 单向派生（phase06-05 / phase06-13 合同约束）。
 *
 * 后端响应示例（GET /api/reuse-summary?scope=dashboard）：
 *   {
 *     "module_reuse_summary": [
 *       {
 *         "module_id": "...",
 *         "module_name": "auth-service",
 *         "reuse_product_count": 3,
 *         "latest_reuse_at": "2026-08-10T...",
 *         "explanation_text": "模块「auth-service」当前被 3 个 Product 复用"
 *       }
 *     ],
 *     "capability_summary": [
 *       {
 *         "capability_key": "auth",
 *         "capability_label": "Authentication",
 *         "supporting_module_count": 1,
 *         "latest_capability_update_at": "2026-08-10T...",
 *         "empty_state_text": ""
 *       }
 *     ]
 *   }
 */

/** Reuse Summary 读取作用域 */
export type ReuseSummaryScope = 'dashboard' | 'module_detail' | 'product_detail'

/** 模块复用摘要项（对齐 proto ModuleReuseSummaryEntry） */
export interface ModuleReuseSummaryEntry {
  module_id: string
  module_name: string
  reuse_product_count: number
  latest_reuse_at: string
  explanation_text: string
}

/** 能力摘要项（对齐 proto CapabilitySummaryEntry） */
export interface CapabilitySummaryEntry {
  capability_key: string
  capability_label: string
  supporting_module_count: number
  latest_capability_update_at: string
  empty_state_text: string
}

/** ReuseSummaryRead 响应（对齐 proto GetReuseSummaryResponse） */
export interface ReuseSummaryReadResult {
  module_reuse_summary: ModuleReuseSummaryEntry[]
  capability_summary: CapabilitySummaryEntry[]
}

/** ReuseSummaryRead 查询参数 */
export interface ReuseSummaryQueryParams {
  scope: ReuseSummaryScope
  module_id?: string
  product_id?: string
}
