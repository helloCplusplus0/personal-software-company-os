/**
 * template-reuse-query-options — Template Reuse 只读 query key 与 query 配置。
 *
 * phase09-08 spec §"template reuse query options 与 read owner 的正式职责"：
 *   - query options 只承接 queryKey、只读请求与响应解包
 *   - 四个 read owner 各承接一个 RPC 的一对一映射
 */
import { queryOptions } from '@tanstack/react-query'
import { templateReuseClient } from './connect-client'
import type { TemplateConsumerSurface, TemplateSource } from '@/gen/proto/psco/template_reuse/v1/template_reuse_pb'

/**
 * 模板候选列表的 query key。
 */
export const TEMPLATE_CANDIDATES_QUERY_KEY = ['template-candidates'] as const

/**
 * 模板预填详情的 query key。
 */
export const TEMPLATE_PREFILL_QUERY_KEY = ['template-prefill'] as const

/**
 * 派生提示的 query key。
 */
export const DERIVED_INSIGHT_HINTS_QUERY_KEY = ['derived-insight-hints'] as const

/**
 * 模板来源复读的 query key。
 */
export const TEMPLATE_SOURCE_SUMMARY_QUERY_KEY = ['template-source-summary'] as const

/**
 * templateCandidatesQueryOptions — 模板候选列表的 query options。
 */
export function templateCandidatesQueryOptions(consumerSurface: TemplateConsumerSurface) {
  return queryOptions({
    queryKey: [...TEMPLATE_CANDIDATES_QUERY_KEY, consumerSurface],
    queryFn: async ({ signal }) => {
      const res = await templateReuseClient.listTemplateCandidates(
        { consumerSurface },
        { signal },
      )
      return res
    },
  })
}

/**
 * templatePrefillQueryOptions — 模板预填详情的 query options。
 */
export function templatePrefillQueryOptions(
  templateCandidateId: string,
  consumerSurface: TemplateConsumerSurface,
) {
  return queryOptions({
    queryKey: [...TEMPLATE_PREFILL_QUERY_KEY, templateCandidateId, consumerSurface],
    queryFn: async ({ signal }) => {
      const res = await templateReuseClient.getTemplateCandidatePrefill(
        { templateCandidateId, consumerSurface },
        { signal },
      )
      return res.prefill
    },
  })
}

/**
 * derivedInsightHintsQueryOptions — 派生提示的 query options。
 */
export function derivedInsightHintsQueryOptions(
  templateCandidateId: string,
  consumerSurface: TemplateConsumerSurface,
  reviewScopeKey: string,
) {
  return queryOptions({
    queryKey: [...DERIVED_INSIGHT_HINTS_QUERY_KEY, templateCandidateId, consumerSurface, reviewScopeKey],
    queryFn: async ({ signal }) => {
      const res = await templateReuseClient.getDerivedInsightHints(
        { templateCandidateId, consumerSurface, reviewScopeKey },
        { signal },
      )
      return res
    },
  })
}

/**
 * templateSourceSummaryQueryOptions — 模板来源复读的 query options。
 */
export function templateSourceSummaryQueryOptions(
  templateCandidateId: string,
  templateSource: TemplateSource,
  consumerSurface: TemplateConsumerSurface,
) {
  return queryOptions({
    queryKey: [...TEMPLATE_SOURCE_SUMMARY_QUERY_KEY, templateCandidateId, templateSource, consumerSurface],
    queryFn: async ({ signal }) => {
      const res = await templateReuseClient.getTemplateSourceSummary(
        { templateCandidateId, templateSource, consumerSurface },
        { signal },
      )
      return res.sourceSummary
    },
  })
}