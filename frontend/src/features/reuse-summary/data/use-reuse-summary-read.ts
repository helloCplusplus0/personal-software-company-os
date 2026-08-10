/**
 * useReuseSummaryRead — Reuse Summary 只读 query owner
 *
 * phase06-15 §"ReuseSummaryRead 以前端单一页面级 query owner 落地"
 * phase06-09 §"Dashboard 复用快照挂接位"
 *
 * 职责：
 *   - 只承接读取、缓存键与响应解包
 *   - 不得混入 create / update / bind / link 写动作
 *   - 不得为 module_reuse_summary 与 capability_summary 长出两套平行 query owner
 *
 * 缓存键：['reuse-summary', scope, module_id?, product_id?]
 */
import { useQuery, type UseQueryResult } from '@tanstack/react-query'
import { fetchReuseSummary } from './api-adapter'
import type { ReuseSummaryReadResult, ReuseSummaryQueryParams } from '../types'

export type UseReuseSummaryRead = UseQueryResult<ReuseSummaryReadResult, Error>

/**
 * useReuseSummaryRead — 页面级只读 query owner。
 *
 * 参数：
 *   - scope: 'dashboard' | 'module_detail' | 'product_detail'
 *   - moduleId: module_detail 作用域必填
 *   - productId: product_detail 作用域必填
 *   - enabled: 可控制是否启用查询（如详情页未 ready 时延迟查询）
 */
export function useReuseSummaryRead(
  params: ReuseSummaryQueryParams,
  options?: { enabled?: boolean },
): UseReuseSummaryRead {
  return useQuery<ReuseSummaryReadResult, Error>({
    queryKey: ['reuse-summary', params.scope, params.module_id, params.product_id],
    queryFn: () => fetchReuseSummary(params),
    enabled: options?.enabled ?? true,
  })
}
