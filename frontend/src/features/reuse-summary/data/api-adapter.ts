/**
 * Reuse Summary API 适配层
 *
 * 对接 phase06-14 后端 reusesummary 模块：
 *   GET /api/reuse-summary?scope=...&module_id=...&product_id=...
 *
 * 字段语义从 .proto -> HTTP DTO 单向派生（phase06-05 / phase06-13 合同约束）。
 * 后端已使用 snake_case，与前端 types.ts 一致，无需转换。
 */
import type { ReuseSummaryReadResult, ReuseSummaryQueryParams } from '../types'

const API_BASE_URL = import.meta.env.VITE_API_BASE_URL ?? ''

/**
 * fetchReuseSummary — ReuseSummaryRead 读组
 *
 * GET /api/reuse-summary?scope=dashboard|module_detail|product_detail&module_id=...&product_id=...
 *
 * 返回 module_reuse_summary 与 capability_summary。
 * dashboard 作用域返回前 5 条；module_detail / product_detail 作用域围绕指定 ID 返回。
 */
export async function fetchReuseSummary(
  params: ReuseSummaryQueryParams,
): Promise<ReuseSummaryReadResult> {
  const searchParams = new URLSearchParams()
  searchParams.set('scope', params.scope)
  if (params.module_id) {
    searchParams.set('module_id', params.module_id)
  }
  if (params.product_id) {
    searchParams.set('product_id', params.product_id)
  }

  const res = await fetch(
    `${API_BASE_URL}/api/reuse-summary?${searchParams.toString()}`,
    { headers: { Accept: 'application/json' } },
  )

  if (!res.ok) {
    let message = `HTTP ${res.status}`
    try {
      const body = await res.json()
      if (body?.error) {
        message = body.error
      }
    } catch {
      // 响应体非 JSON，保留默认 message
    }
    throw new Error(message)
  }

  return res.json() as Promise<ReuseSummaryReadResult>
}
