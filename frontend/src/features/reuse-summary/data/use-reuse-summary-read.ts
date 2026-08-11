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
import { timestampDate } from '@bufbuild/protobuf/wkt'
import { reuseSummaryClient } from './connect-client'
import type { ReuseSummaryReadResult, ReuseSummaryQueryParams } from '../types'
import { ReuseSummaryScope } from '@/gen/proto/psco/reuse_summary/v1/reuse_summary_pb'

export type UseReuseSummaryRead = UseQueryResult<ReuseSummaryReadResult, Error>

function mapScopeToProto(scope: ReuseSummaryQueryParams['scope']): ReuseSummaryScope {
  switch (scope) {
    case 'dashboard':
      return ReuseSummaryScope.DASHBOARD
    case 'module_detail':
      return ReuseSummaryScope.MODULE_DETAIL
    case 'product_detail':
      return ReuseSummaryScope.PRODUCT_DETAIL
  }
}

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
    queryFn: async (): Promise<ReuseSummaryReadResult> => {
      const res = await reuseSummaryClient.getReuseSummary({
        scope: mapScopeToProto(params.scope),
        moduleId: params.module_id ?? '',
        productId: params.product_id ?? '',
      })
      return {
        module_reuse_summary: (res.moduleReuseSummary ?? []).map((m) => ({
          module_id: m.moduleId ?? '',
          module_name: '',
          reuse_count: 0,
          reuse_product_count: m.reuseProductCount ?? 0,
          latest_reuse_at: m.latestReuseAt ? timestampDate(m.latestReuseAt).toISOString() : '',
          explanation_text: m.explanationText ?? '',
        })),
        capability_summary: (res.capabilitySummary ?? []).map((c) => ({
          capability_key: c.capabilityKey ?? '',
          capability_label: c.capabilityLabel ?? '',
          supporting_module_count: c.supportingModuleCount ?? 0,
          latest_capability_update_at: c.latestCapabilityUpdateAt ? timestampDate(c.latestCapabilityUpdateAt).toISOString() : '',
          empty_state_text: c.emptyStateText ?? '',
        })),
      } as unknown as ReuseSummaryReadResult
    },
    enabled: options?.enabled ?? true,
  })
}
