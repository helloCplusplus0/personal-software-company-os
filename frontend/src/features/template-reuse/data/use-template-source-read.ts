/**
 * useTemplateSourceRead — 模板来源复读只读 owner。
 *
 * phase09-08 spec §"template reuse query options 与 read owner 的正式职责"：
 *   本 hook 以 TemplateReuseService generated client 为正式 transport 入口，
 *   只承接 queryKey、只读请求、响应解包与 empty/unavailable/error 派生。
 *
 * 一对一映射：use-template-source-read -> GetTemplateSourceSummary
 */
import { useQuery } from '@tanstack/react-query'
import {
  templateSourceSummaryQueryOptions,
  TEMPLATE_SOURCE_SUMMARY_QUERY_KEY,
} from './template-reuse-query-options'
import type {
  TemplateConsumerSurface,
  TemplateSource,
  TemplateSourceSummary,
} from '@/gen/proto/psco/template_reuse/v1/template_reuse_pb'
import { useMemo } from 'react'

// ============================================================================
// 只读状态模型
// ============================================================================

export type TemplateSourcePageStatus = 'initial-loading' | 'resolved' | 'unavailable' | 'error'

export interface UseTemplateSourceReadResult {
  sourceSummary: TemplateSourceSummary | undefined
  isLoading: boolean
  isError: boolean
  error: Error | null
  pageStatus: TemplateSourcePageStatus
}

// ============================================================================
// Hook
// ============================================================================

export function useTemplateSourceRead(
  templateCandidateId: string,
  templateSource: TemplateSource,
  consumerSurface: TemplateConsumerSurface,
): UseTemplateSourceReadResult {
  const enabled = templateCandidateId !== ''

  const query = useQuery({
    ...templateSourceSummaryQueryOptions(templateCandidateId, templateSource, consumerSurface),
    queryKey: [...TEMPLATE_SOURCE_SUMMARY_QUERY_KEY, templateCandidateId, templateSource, consumerSurface],
    enabled,
  })

  const pageStatus = useMemo((): TemplateSourcePageStatus => {
    if (!enabled) return 'unavailable'
    if (query.isLoading && !query.data) {
      return 'initial-loading'
    }
    if (query.isError) {
      return 'error'
    }
    const summary = query.data
    if (!summary) return 'unavailable'
    if (summary.resolutionStatus === 2) return 'unavailable' // UNAVAILABLE
    return 'resolved' // RESOLVED
  }, [enabled, query.isLoading, query.isError, query.data])

  return {
    sourceSummary: query.data,
    isLoading: query.isLoading,
    isError: query.isError,
    error: query.error as Error | null,
    pageStatus,
  }
}