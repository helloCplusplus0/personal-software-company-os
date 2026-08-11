/**
 * useDailyReviewRead — Daily Review 只读 query owner。
 *
 * phase08-08 §"review read owner 的正式 transport 路径"：
 *   本 hook 以 ReviewService generated client 为正式 transport 入口，
 *   不再在页面层直接并排消费 dashboard / decision-center 的底层 query hook。
 *
 * phase08-06 §"review read layer 必须导出稳定的页面级状态模型"：
 *   导出 pageStatus 与 sectionStatus，而不是把多个 query 的原始状态直接暴露给页面逐段拼接。
 *
 * 数据模型：
 *   - current_focus_signals：来自 phase05 Dashboard FeedbackSignalRead
 *   - representative_signals：来自 phase05 Dashboard FeedbackSignalRead 的 asset_feedback_summary
 *   - pending_decisions：来自 phase03 Decision Center 的 proposed 决策 top N 摘要
 */
import { useQuery } from '@tanstack/react-query'
import { dailyReviewQueryOptions, DAILY_REVIEW_QUERY_KEY } from './review-query-options'
import type { DailyReviewContext } from '@/gen/proto/psco/review/v1/review_pb'
import type { Timestamp } from '@bufbuild/protobuf/wkt'
import type { FeedbackSignal } from '@/features/dashboard/types'
import { useCallback, useMemo } from 'react'

// ============================================================================
// 页面级状态模型
// ============================================================================

export type PageStatus = 'initial-loading' | 'ready' | 'page-error'

export type SectionStatus = 'ready' | 'empty' | 'error'

export interface DailyReviewPageState {
  pageStatus: PageStatus
  currentFocusSectionStatus: SectionStatus
  representativeSignalsSectionStatus: SectionStatus
  pendingDecisionsSectionStatus: SectionStatus
}

// ============================================================================
// 只读数据模型
// ============================================================================

export interface DailyReviewReadModel {
  currentFocusSignals: FeedbackSignal[]
  representativeSignals: FeedbackSignal[]
  pendingDecisions: PendingDecision[]
}

export interface PendingDecision {
  id: string
  title: string
  status: string
  created_at: string
  link_count: number
  linked_module_summary: string
}

// ============================================================================
// 类型转换
// ============================================================================

/** proto Timestamp → ISO 8601 字符串 */
function timestampToISO(ts: Timestamp | undefined): string {
  if (!ts) return ''
  const ms = Number(ts.seconds) * 1000 + Math.floor(ts.nanos / 1_000_000)
  return new Date(ms).toISOString()
}

/** 从 proto FeedbackSignal（camelCase）转换为前端 FeedbackSignal（snake_case） */
function protoToFeedbackSignal(signal: DailyReviewContext['currentFocusSignals'][number]): FeedbackSignal {
  return {
    signal_family: protoSignalFamilyToString(signal.signalFamily),
    signal_code: protoSignalCodeToString(signal.signalCode),
    priority: protoSignalPriorityToString(signal.priority),
    title: signal.title,
    summary: signal.summary,
    action_label: signal.actionLabel,
    target_type: protoTargetTypeToString(signal.targetType),
    target_id: signal.targetId,
    target_label: signal.targetLabel,
  }
}

/** 从 proto DecisionListItem（camelCase）转换为前端 PendingDecision */
function protoToPendingDecision(item: DailyReviewContext['pendingDecisions'][number]): PendingDecision {
  return {
    id: item.id,
    title: item.title,
    status: protoDecisionStatusToString(item.status),
    created_at: timestampToISO(item.createdAt),
    link_count: item.linkCount,
    linked_module_summary: item.linkedModuleSummary,
  }
}

function protoSignalFamilyToString(f: number): FeedbackSignal['signal_family'] {
  switch (f) {
    case 1: return 'pending_decision'
    case 2: return 'product_asset_coverage'
    default: return 'pending_decision'
  }
}

function protoSignalCodeToString(c: number): FeedbackSignal['signal_code'] {
  switch (c) {
    case 1: return 'pending_decision'
    case 2: return 'product_missing_both_bindings'
    case 3: return 'product_missing_repository_binding'
    case 4: return 'product_missing_module_binding'
    default: return 'pending_decision'
  }
}

function protoSignalPriorityToString(p: number): FeedbackSignal['priority'] {
  switch (p) {
    case 1: return 'p1_pending_decision'
    case 2: return 'p2_product_missing_both_bindings'
    case 3: return 'p3_product_missing_repository_binding'
    case 4: return 'p4_product_missing_module_binding'
    default: return 'p1_pending_decision'
  }
}

function protoTargetTypeToString(t: number): FeedbackSignal['target_type'] {
  switch (t) {
    case 1: return 'decision_detail'
    case 2: return 'decision_list'
    case 3: return 'product_detail'
    case 4: return 'module_detail'
    case 5: return 'repository_detail'
    default: return 'decision_detail'
  }
}

function protoDecisionStatusToString(s: number): string {
  switch (s) {
    case 1: return 'proposed'
    case 2: return 'active'
    case 3: return 'superseded'
    case 4: return 'archived'
    default: return 'proposed'
  }
}

// ============================================================================
// Hook
// ============================================================================

export interface UseDailyReviewReadResult {
  data: DailyReviewReadModel | undefined
  isLoading: boolean
  isError: boolean
  error: Error | null
  pageState: DailyReviewPageState
  /** 整页重试：重新触发 daily review context 查询 */
  retry: () => void
}

export function useDailyReviewRead(): UseDailyReviewReadResult {
  const query = useQuery({
    ...dailyReviewQueryOptions(),
    queryKey: DAILY_REVIEW_QUERY_KEY,
  })

  const retry = useCallback(() => {
    void query.refetch()
  }, [query])

  const pageState = useMemo((): DailyReviewPageState => {
    if (query.isLoading && !query.data) {
      return {
        pageStatus: 'initial-loading',
        currentFocusSectionStatus: 'ready',
        representativeSignalsSectionStatus: 'ready',
        pendingDecisionsSectionStatus: 'ready',
      }
    }
    if (query.isError) {
      return {
        pageStatus: 'page-error',
        currentFocusSectionStatus: 'error',
        representativeSignalsSectionStatus: 'error',
        pendingDecisionsSectionStatus: 'error',
      }
    }
    const data = query.data
    if (!data) {
      return {
        pageStatus: 'ready',
        currentFocusSectionStatus: 'empty',
        representativeSignalsSectionStatus: 'empty',
        pendingDecisionsSectionStatus: 'empty',
      }
    }
    return {
      pageStatus: 'ready',
      currentFocusSectionStatus: data.currentFocusSignals.length > 0 ? 'ready' : 'empty',
      representativeSignalsSectionStatus: data.representativeSignals.length > 0 ? 'ready' : 'empty',
      pendingDecisionsSectionStatus: data.pendingDecisions.length > 0 ? 'ready' : 'empty',
    }
  }, [query.isLoading, query.isError, query.data])

  const data = useMemo((): DailyReviewReadModel | undefined => {
    const ctx = query.data
    if (!ctx) return undefined
    return {
      currentFocusSignals: ctx.currentFocusSignals.map(protoToFeedbackSignal),
      representativeSignals: ctx.representativeSignals.map(protoToFeedbackSignal),
      pendingDecisions: ctx.pendingDecisions.map(protoToPendingDecision),
    }
  }, [query.data])

  return {
    data,
    isLoading: query.isLoading,
    isError: query.isError,
    error: query.error as Error | null,
    pageState,
    retry,
  }
}