/**
 * useTemplateCandidatesRead — 模板候选列表只读 owner。
 *
 * phase09-08 spec §"template reuse query options 与 read owner 的正式职责"：
 *   本 hook 以 TemplateReuseService generated client 为正式 transport 入口，
 *   只承接 queryKey、只读请求、响应解包与 empty/unavailable/error 派生。
 *
 * 一对一映射：use-template-candidates-read -> ListTemplateCandidates
 */
import { useQuery } from '@tanstack/react-query'
import {
  templateCandidatesQueryOptions,
  TEMPLATE_CANDIDATES_QUERY_KEY,
} from './template-reuse-query-options'
import type {
  TemplateCandidateSummary,
  TemplateConsumerSurface,
} from '@/gen/proto/psco/template_reuse/v1/template_reuse_pb'
import { useMemo } from 'react'

// ============================================================================
// 只读状态模型
// ============================================================================

export type TemplateCandidatesPageStatus = 'initial-loading' | 'ready' | 'empty' | 'error'

export interface TemplateCandidatesReadModel {
  candidates: TemplateCandidateSummary[]
  defaultActiveCandidateId: string
}

export interface UseTemplateCandidatesReadResult {
  data: TemplateCandidatesReadModel | undefined
  isLoading: boolean
  isError: boolean
  error: Error | null
  pageStatus: TemplateCandidatesPageStatus
}

// ============================================================================
// Hook
// ============================================================================

export function useTemplateCandidatesRead(
  consumerSurface: TemplateConsumerSurface,
): UseTemplateCandidatesReadResult {
  const query = useQuery({
    ...templateCandidatesQueryOptions(consumerSurface),
    queryKey: [...TEMPLATE_CANDIDATES_QUERY_KEY, consumerSurface],
  })

  const pageStatus = useMemo((): TemplateCandidatesPageStatus => {
    if (query.isLoading && !query.data) {
      return 'initial-loading'
    }
    if (query.isError) {
      return 'error'
    }
    const data = query.data
    if (!data || data.candidates.length === 0) {
      return 'empty'
    }
    return 'ready'
  }, [query.isLoading, query.isError, query.data])

  const data = useMemo((): TemplateCandidatesReadModel | undefined => {
    const res = query.data
    if (!res) return undefined
    return {
      candidates: res.candidates,
      defaultActiveCandidateId: res.defaultActiveCandidateId,
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