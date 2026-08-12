/**
 * useWeeklyReviewRead — Weekly Review 只读 query owner。
 *
 * phase08-08 §"review read owner 的正式 transport 路径"：
 *   本 hook 以 ReviewService generated client 为正式 transport 入口，
 *   不再在页面层直接并排消费 dashboard / reuse-summary 的底层 query hook。
 *
 * phase08-06 §"review read layer 必须导出稳定的页面级状态模型"：
 *   导出 pageStatus 与 sectionStatus，而不是把多个 query 的原始状态直接暴露给页面逐段拼接。
 *
 * 数据模型：
 *   - overview：来自 phase05 Dashboard DashboardOverview
 *   - recent_activities：来自 phase05 Dashboard RecentActivityRead
 *   - representative_signals：来自 phase05 Dashboard FeedbackSignalRead 的 asset_feedback_summary
 *   - module_reuse_summary：来自 phase06 ReuseSummary
 *   - capability_summary：来自 phase06 ReuseSummary
 */
import { useQuery } from '@tanstack/react-query'
import { weeklyReviewQueryOptions, WEEKLY_REVIEW_QUERY_KEY } from './review-query-options'
import type { WeeklyReviewContext } from '@/gen/proto/psco/review/v1/review_pb'
import type { Timestamp } from '@bufbuild/protobuf/wkt'
import type { DashboardOverview, FeedbackSignal, RecentActivityItem } from '@/features/dashboard/types'
import { useCallback, useMemo, useState } from 'react'
import { useTemplateCandidatesRead } from '@/features/template-reuse/data/use-template-candidates-read'
import { useDerivedInsightHintsRead } from '@/features/template-reuse/data/use-derived-insight-hints-read'
import { TemplateConsumerSurface } from '@/gen/proto/psco/template_reuse/v1/template_reuse_pb'
import type { TemplateCandidateSummary, DerivedInsightHint } from '@/gen/proto/psco/template_reuse/v1/template_reuse_pb'

// ============================================================================
// 页面级状态模型
// ============================================================================

export type PageStatus = 'initial-loading' | 'ready' | 'page-error'

export type SectionStatus = 'ready' | 'empty' | 'error'

export interface WeeklyReviewPageState {
  pageStatus: PageStatus
  overviewSectionStatus: SectionStatus
  recentActivitySectionStatus: SectionStatus
  representativeSignalsSectionStatus: SectionStatus
  reuseSnapshotSectionStatus: SectionStatus
  templateSectionStatus: SectionStatus
  hintsSectionStatus: SectionStatus
}

// ============================================================================
// 只读数据模型
// ============================================================================

export interface WeeklyReviewReadModel {
  overview: DashboardOverview | null
  recentActivities: RecentActivityItem[]
  representativeSignals: FeedbackSignal[]
  moduleReuseSummary: ModuleReuseSummaryItem[]
  capabilitySummary: CapabilitySummaryItem[]
}

export interface ModuleReuseSummaryItem {
  module_id: string
  reuse_product_count: number
  latest_reuse_at: string
  explanation_text: string
}

export interface CapabilitySummaryItem {
  capability_key: string
  capability_label: string
  supporting_module_count: number
  latest_capability_update_at: string
  empty_state_text: string
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

/** 从 proto DashboardOverview（camelCase）转换为前端 DashboardOverview（snake_case） */
function protoToOverview(o: WeeklyReviewContext['overview']): DashboardOverview | null {
  if (!o) return null
  return {
    module_count: o.moduleCount,
    product_count: o.productCount,
    repository_count: o.repositoryCount,
    decision_count: o.decisionCount,
    product_with_repository_count: o.productWithRepositoryCount,
    product_with_module_count: o.productWithModuleCount,
  }
}

/** 从 proto FeedbackSignal（camelCase）转换为前端 FeedbackSignal（snake_case） */
function protoToFeedbackSignal(signal: WeeklyReviewContext['representativeSignals'][number]): FeedbackSignal {
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

/** 从 proto RecentActivityItem（camelCase）转换为前端 RecentActivityItem（snake_case） */
function protoToRecentActivity(item: WeeklyReviewContext['recentActivities'][number]): RecentActivityItem {
  return {
    activity_type: protoActivityTypeToString(item.activityType),
    activity_at: timestampToISO(item.activityAt),
    target_type: protoTargetTypeToString(item.targetType),
    target_id: item.targetId,
    target_label: item.targetLabel,
  }
}

/** 从 proto ModuleReuseSummary（camelCase）转换为前端 ModuleReuseSummaryItem */
function protoToModuleReuseSummary(item: WeeklyReviewContext['moduleReuseSummary'][number]): ModuleReuseSummaryItem {
  return {
    module_id: item.moduleId,
    reuse_product_count: item.reuseProductCount,
    latest_reuse_at: timestampToISO(item.latestReuseAt),
    explanation_text: item.explanationText,
  }
}

/** 从 proto CapabilitySummary（camelCase）转换为前端 CapabilitySummaryItem */
function protoToCapabilitySummary(item: WeeklyReviewContext['capabilitySummary'][number]): CapabilitySummaryItem {
  return {
    capability_key: item.capabilityKey,
    capability_label: item.capabilityLabel,
    supporting_module_count: item.supportingModuleCount,
    latest_capability_update_at: timestampToISO(item.latestCapabilityUpdateAt),
    empty_state_text: item.emptyStateText,
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

function protoActivityTypeToString(t: number): RecentActivityItem['activity_type'] {
  switch (t) {
    case 1: return 'module'
    case 2: return 'release'
    case 3: return 'product'
    case 4: return 'repository'
    case 5: return 'decision'
    case 6: return 'product_module_binding'
    case 7: return 'product_repository_binding'
    case 8: return 'module_repository_binding'
    default: return 'module'
  }
}

// ============================================================================
// Hook
// ============================================================================

export interface UseWeeklyReviewReadResult {
  data: WeeklyReviewReadModel | undefined
  isLoading: boolean
  isError: boolean
  error: Error | null
  pageState: WeeklyReviewPageState
  retry: () => void
  /** phase09-09 模板候选列表 */
  templateCandidates: TemplateCandidateSummary[]
  /** phase09-09 默认激活候选 ID */
  defaultActiveCandidateId: string
  /** phase09-09 当前激活候选 ID */
  activeCandidateId: string
  /** phase09-09 设置激活候选 */
  setActiveCandidateId: (id: string) => void
  /** phase09-09 派生提示列表 */
  hints: DerivedInsightHint[]
}

export function useWeeklyReviewRead(): UseWeeklyReviewReadResult {
  const query = useQuery({
    ...weeklyReviewQueryOptions(),
    queryKey: WEEKLY_REVIEW_QUERY_KEY,
  })

  // phase09-09：模板候选只读
  const templateCandidatesQuery = useTemplateCandidatesRead(
    TemplateConsumerSurface.WEEKLY_REVIEW,
  )

  // phase09-09：模板候选默认选中第一个
  const candidates = templateCandidatesQuery.data?.candidates ?? []
  const defaultActiveId = templateCandidatesQuery.data?.defaultActiveCandidateId ?? ''
  const [activeCandidateId, setActiveCandidateId] = useState<string>(defaultActiveId)

  // 当候选列表更新时，如果 activeCandidateId 不在新列表中，重置为默认
  useMemo(() => {
    if (candidates.length > 0 && !candidates.find((c) => c.templateCandidateId === activeCandidateId)) {
      setActiveCandidateId(defaultActiveId)
    }
  }, [candidates, defaultActiveId, activeCandidateId])

  // phase09-09：派生提示只读
  const hintsQuery = useDerivedInsightHintsRead(
    activeCandidateId,
    TemplateConsumerSurface.WEEKLY_REVIEW,
    'weekly-review-scope',
  )

  const retry = useCallback(() => {
    void query.refetch()
  }, [query])

  const pageState = useMemo((): WeeklyReviewPageState => {
    if (query.isLoading && !query.data) {
      return {
        pageStatus: 'initial-loading',
        overviewSectionStatus: 'ready',
        recentActivitySectionStatus: 'ready',
        representativeSignalsSectionStatus: 'ready',
        reuseSnapshotSectionStatus: 'ready',
        templateSectionStatus: 'ready',
        hintsSectionStatus: 'ready',
      }
    }
    if (query.isError) {
      return {
        pageStatus: 'page-error',
        overviewSectionStatus: 'error',
        recentActivitySectionStatus: 'error',
        representativeSignalsSectionStatus: 'error',
        reuseSnapshotSectionStatus: 'error',
        templateSectionStatus: 'error',
        hintsSectionStatus: 'error',
      }
    }
    const data = query.data
    if (!data) {
      return {
        pageStatus: 'ready',
        overviewSectionStatus: 'empty',
        recentActivitySectionStatus: 'empty',
        representativeSignalsSectionStatus: 'empty',
        reuseSnapshotSectionStatus: 'empty',
        templateSectionStatus: candidates.length > 0 ? 'ready' : 'empty',
        hintsSectionStatus: 'empty',
      }
    }
    return {
      pageStatus: 'ready',
      overviewSectionStatus: data.overview ? 'ready' : 'empty',
      recentActivitySectionStatus: data.recentActivities.length > 0 ? 'ready' : 'empty',
      representativeSignalsSectionStatus: data.representativeSignals.length > 0 ? 'ready' : 'empty',
      reuseSnapshotSectionStatus: (data.moduleReuseSummary.length > 0 || data.capabilitySummary.length > 0) ? 'ready' : 'empty',
      templateSectionStatus: candidates.length > 0 ? 'ready' : 'empty',
      hintsSectionStatus: 'ready',
    }
  }, [query.isLoading, query.isError, query.data, candidates.length])

  const data = useMemo((): WeeklyReviewReadModel | undefined => {
    const ctx = query.data
    if (!ctx) return undefined
    return {
      overview: protoToOverview(ctx.overview),
      recentActivities: ctx.recentActivities.map(protoToRecentActivity),
      representativeSignals: ctx.representativeSignals.map(protoToFeedbackSignal),
      moduleReuseSummary: ctx.moduleReuseSummary.map(protoToModuleReuseSummary),
      capabilitySummary: ctx.capabilitySummary.map(protoToCapabilitySummary),
    }
  }, [query.data])

  return {
    data,
    isLoading: query.isLoading,
    isError: query.isError,
    error: query.error as Error | null,
    pageState,
    retry,
    templateCandidates: candidates,
    defaultActiveCandidateId: defaultActiveId,
    activeCandidateId,
    setActiveCandidateId,
    hints: hintsQuery.data?.hints ?? [],
  }
}
