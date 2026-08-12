/**
 * useDerivedInsightHintsRead — 派生提示只读 owner。
 *
 * phase09-08 spec §"template reuse query options 与 read owner 的正式职责"：
 *   本 hook 以 TemplateReuseService generated client 为正式 transport 入口，
 *   只承接 queryKey、只读请求、响应解包与 empty/unavailable/error 派生。
 *
 * 一对一映射：use-derived-insight-hints-read -> GetDerivedInsightHints
 */
import { useQuery } from '@tanstack/react-query'
import {
  derivedInsightHintsQueryOptions,
  DERIVED_INSIGHT_HINTS_QUERY_KEY,
} from './template-reuse-query-options'
import type {
  DerivedInsightHint,
  TemplateConsumerSurface,
} from '@/gen/proto/psco/template_reuse/v1/template_reuse_pb'
import { useMemo } from 'react'

// ============================================================================
// 只读状态模型
// ============================================================================

export type DerivedInsightHintsPageStatus = 'initial-loading' | 'resolved' | 'unavailable' | 'error'

export interface DerivedInsightHintsReadModel {
  hints: DerivedInsightHint[]
  resolutionStatus: 'RESOLVED' | 'UNAVAILABLE' | 'UNSPECIFIED'
  unavailableReasonText: string
}

export interface UseDerivedInsightHintsReadResult {
  data: DerivedInsightHintsReadModel | undefined
  isLoading: boolean
  isError: boolean
  error: Error | null
  pageStatus: DerivedInsightHintsPageStatus
}

// ============================================================================
// Hook
// ============================================================================

export function useDerivedInsightHintsRead(
  templateCandidateId: string,
  consumerSurface: TemplateConsumerSurface,
  reviewScopeKey: string,
): UseDerivedInsightHintsReadResult {
  const enabled = templateCandidateId !== ''

  const query = useQuery({
    ...derivedInsightHintsQueryOptions(templateCandidateId, consumerSurface, reviewScopeKey),
    queryKey: [...DERIVED_INSIGHT_HINTS_QUERY_KEY, templateCandidateId, consumerSurface, reviewScopeKey],
    enabled,
  })

  const pageStatus = useMemo((): DerivedInsightHintsPageStatus => {
    if (!enabled) return 'unavailable'
    if (query.isLoading && !query.data) {
      return 'initial-loading'
    }
    if (query.isError) {
      return 'error'
    }
    const data = query.data
    if (!data) return 'unavailable'
    if (data.resolutionStatus === 2) return 'unavailable' // UNAVAILABLE
    return 'resolved' // RESOLVED or empty hints
  }, [enabled, query.isLoading, query.isError, query.data])

  const data = useMemo((): DerivedInsightHintsReadModel | undefined => {
    const res = query.data
    if (!res) return undefined
    return {
      hints: res.hints,
      resolutionStatus: res.resolutionStatus === 1 ? 'RESOLVED'
        : res.resolutionStatus === 2 ? 'UNAVAILABLE'
        : 'UNSPECIFIED',
      unavailableReasonText: res.unavailableReasonText,
    }
  }, [query.data])

  return {
    data,
    isLoading: query.isLoading,
    isError: query.isError,
    error: query.error as Error | null,
    pageStatus,
  }
}