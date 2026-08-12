/**
 * useTemplatePrefillRead — 模板预填详情只读 owner。
 *
 * phase09-08 spec §"template reuse query options 与 read owner 的正式职责"：
 *   本 hook 以 TemplateReuseService generated client 为正式 transport 入口，
 *   只承接 queryKey、只读请求、响应解包与 empty/unavailable/error 派生。
 *
 * 一对一映射：use-template-prefill-read -> GetTemplateCandidatePrefill
 */
import { useQuery } from '@tanstack/react-query'
import {
  templatePrefillQueryOptions,
  TEMPLATE_PREFILL_QUERY_KEY,
} from './template-reuse-query-options'
import type {
  TemplateCandidatePrefill,
  TemplateConsumerSurface,
} from '@/gen/proto/psco/template_reuse/v1/template_reuse_pb'
import { useMemo } from 'react'

// ============================================================================
// 只读状态模型
// ============================================================================

export type TemplatePrefillPageStatus = 'initial-loading' | 'resolved' | 'unavailable' | 'error'

export interface UseTemplatePrefillReadResult {
  prefill: TemplateCandidatePrefill | undefined
  isLoading: boolean
  isError: boolean
  error: Error | null
  /** 派生状态：resolved（可用）、unavailable（模板已失效但可继续创建）、error（请求失败） */
  pageStatus: TemplatePrefillPageStatus
}

// ============================================================================
// Hook
// ============================================================================

export function useTemplatePrefillRead(
  templateCandidateId: string,
  consumerSurface: TemplateConsumerSurface,
): UseTemplatePrefillReadResult {
  const enabled = templateCandidateId !== ''

  const query = useQuery({
    ...templatePrefillQueryOptions(templateCandidateId, consumerSurface),
    queryKey: [...TEMPLATE_PREFILL_QUERY_KEY, templateCandidateId, consumerSurface],
    enabled,
  })

  const pageStatus = useMemo((): TemplatePrefillPageStatus => {
    if (!enabled) return 'unavailable'
    if (query.isLoading && !query.data) {
      return 'initial-loading'
    }
    if (query.isError) {
      return 'error'
    }
    const prefill = query.data
    if (!prefill) return 'unavailable'
    if (prefill.resolutionStatus === 1) return 'resolved' // RESOLVED
    if (prefill.resolutionStatus === 2) return 'unavailable' // UNAVAILABLE
    return 'unavailable'
  }, [enabled, query.isLoading, query.isError, query.data])

  return {
    prefill: query.data,
    isLoading: query.isLoading,
    isError: query.isError,
    error: query.error as Error | null,
    pageStatus,
  }
}