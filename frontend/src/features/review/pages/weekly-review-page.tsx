/**
 * WeeklyReviewPage — Weekly Review 正式页面。
 *
 * phase08-08 §"Weekly Review 最小 enablement"：
 *   - 通过 ReviewPageShell 承接头部、离开入口、页面级状态与底部动作区
 *   - 至少必须显示 overview / recent activity / representative signals / reuse snapshot 的最小只读组织
 *   - 必须接上 ReviewActionFooter 或等价完成区，能够正式触发 useReviewAction()
 *
 * phase08-08 UI 对齐 Dashboard 基线（验收后调整）：
 *   - 移除 Card 重型卡片，改用 dashboard 的 section + divide-y 紧凑列表
 *   - overview 改为紧凑 stat bar 风格（flex flex-wrap divide-x，对齐 DashboardStatBar）
 *   - recent activity 直接复用既有 RecentActivityItemCard
 *   - representative signals 直接复用既有 FeedbackSignalCard
 *   - reuse snapshot 直接复用 dashboard 的 ReuseSnapshotSection
 *   - 移动端响应式：min-w-0 + truncate + flex-wrap，不溢出
 *
 * 约束：
 *   - 页面只能消费 useWeeklyReviewRead / useReviewAction
 *   - 不得直接 import 底层 canonical query hook
 *   - 不得在页面内直接 createClient()、createConnectTransport()、useMutation()
 */
import { useNavigate, useSearch } from '@tanstack/react-router'
import { ReviewPageShell } from '../components/review-page-shell'
import { ReviewActionFooter } from '../components/review-action-footer'
import { useWeeklyReviewRead } from '../data/use-weekly-review-read'
import { useReviewAction } from '../application/use-review-action'
import { useNavigateBackToDashboard } from '@/features/dashboard/lib/dashboard-source'
import { buildReviewSourceParams } from '../lib/review-source'
import { FeedbackSignalCard } from '@/features/dashboard/components/feedback-signal-card'
import { RecentActivityItemCard } from '@/features/dashboard/components/recent-activity-item-card'
import { ReuseSnapshotSection } from '@/features/dashboard/components/reuse-snapshot-section'
import { Button } from '@/components/ui/button'
import { Skeleton } from '@/components/ui/skeleton'
import { Layers, ArrowRight, LayoutTemplate, Lightbulb, AlertCircle, ExternalLink } from 'lucide-react'
import type { FeedbackSignal, RecentActivityItem, DashboardOverview } from '@/features/dashboard/types'
import type {
  ModuleReuseSummaryEntry,
  CapabilitySummaryEntry,
} from '@/features/reuse-summary/types'
import type { ReviewActionInput } from '../application/review-action-types'
import type { TemplateCandidateSummary, DerivedInsightHint } from '@/gen/proto/psco/template_reuse/v1/template_reuse_pb'
import { DerivedInsightHintType } from '@/gen/proto/psco/template_reuse/v1/template_reuse_pb'
import { useDerivedHintHandoff } from '@/features/template-reuse/application/use-derived-hint-handoff'
import { toast } from 'sonner'

const WEEKLY_REVIEW_ACTION_SECTIONS = {
  decision: 'overview',
  product: 'asset-feedback',
  module: 'asset-feedback',
  repository: 'asset-feedback',
  submitNextStep: 'overview',
} as const

export function WeeklyReviewPage() {
  const navigate = useNavigate()
  const search = useSearch({ from: '/reviews/weekly' })
  const navigateBack = useNavigateBackToDashboard()
  const review = useWeeklyReviewRead((search as any).returnCandidateId)
  const action = useReviewAction()

  const handleBackToDashboard = () => {
    navigateBack('overview')
  }

  const handleNavigateToCreateFromTemplate = (candidateId: string) => {
    navigate({
      to: '/products/new',
      search: {
        fromTemplateReuse: true as any,
        templateCandidateId: candidateId,
        templateSource: 'weekly-review',
      },
    })
  }

  const handleSubmitAction = (input: ReviewActionInput) => {
    action.mutate(input, {
      onSuccess: (envelope) => {
        navigate({
          to: envelope.navigateTo as any,
          params: envelope.params as any,
          search: envelope.search as any,
        })
        if (envelope.successMessage) {
          toast.success(envelope.successMessage)
        }
      },
    })
  }

  const overview = review.data?.overview

  // 适配 ReuseSnapshotSection 期望的 ModuleReuseSummaryEntry 类型
  // review.proto 的 ModuleReuseSummary 不含 module_name 字段，
  // ReuseSnapshotSection 内部也不展示 module_name，传空字符串即可
  const moduleReuseSummary: ModuleReuseSummaryEntry[] =
    review.data?.moduleReuseSummary.map((item) => ({
      module_id: item.module_id,
      module_name: '',
      reuse_product_count: item.reuse_product_count,
      latest_reuse_at: item.latest_reuse_at,
      explanation_text: item.explanation_text,
    })) ?? []

  const capabilitySummary: CapabilitySummaryEntry[] =
    review.data?.capabilitySummary.map((item) => ({
      capability_key: item.capability_key,
      capability_label: item.capability_label,
      supporting_module_count: item.supporting_module_count,
      latest_capability_update_at: item.latest_capability_update_at,
      empty_state_text: item.empty_state_text,
    })) ?? []

  return (
    <ReviewPageShell
      title="Weekly Review"
      subtitle="本周经营回顾与能力整理"
      pageStatus={review.pageState.pageStatus}
      onBackToDashboard={handleBackToDashboard}
      onRetry={review.retry}
      actionFooter={
        <ReviewActionFooter
          isPending={action.isPending}
          hasError={action.isError}
          errorMessage={action.error?.message}
          onSubmitAction={handleSubmitAction}
          onReset={action.reset}
          reviewKind="weekly"
          actionSections={WEEKLY_REVIEW_ACTION_SECTIONS}
        />
      }
    >
      {/* Overview 区块 — 改为紧凑 stat bar 风格，对齐 DashboardStatBar
          移动端 flex-wrap 不撑破容器 */}
      <section className="space-y-2" aria-label="Overview">
        <h2 className="text-base font-semibold">系统概览</h2>
        <OverviewStatBar overview={overview ?? null} status={review.pageState.overviewSectionStatus} />
      </section>

      {/* phase09-09 模板候选选择区 — 位于 Overview 与 Recent Activity 之间 */}
      <section className="space-y-2" aria-label="Template Candidates">
        <h2 className="text-base font-semibold">模板候选</h2>
        <TemplateCandidateSection
          candidates={review.templateCandidates}
          activeCandidateId={review.activeCandidateId}
          onSelectCandidate={review.setActiveCandidateId}
          status={review.pageState.templateSectionStatus}
          onCreateProduct={handleNavigateToCreateFromTemplate}
        />
      </section>

      {/* phase09-10 派生提示展示区 — 位于模板候选区与 Recent Activity 之间 */}
      <DerivedHintsSection
        hints={review.hints}
        activeCandidateId={review.activeCandidateId}
        hintsSectionStatus={review.pageState.hintsSectionStatus}
        onRetry={review.retryHints}
      />

      {/* Recent Activity 区块 — 复用 RecentActivityItemCard
          section="recent-activity"，跳转后回到 dashboard 的 recent-activity 区块 */}
      <section className="space-y-2" aria-label="Recent Activity">
        <h2 className="text-base font-semibold">最近活动</h2>
        <RecentActivityList
          status={review.pageState.recentActivitySectionStatus}
          activities={review.data?.recentActivities ?? []}
        />
      </section>

      {/* Representative Signals 区块 — 复用 FeedbackSignalCard
          section="asset-feedback"，跳转后回到 dashboard 的 asset-feedback 区块 */}
      <section className="space-y-2" aria-label="Representative Signals">
        <h2 className="text-base font-semibold">代表性反馈信号</h2>
        <RepresentativeSignalsList
          status={review.pageState.representativeSignalsSectionStatus}
          signals={review.data?.representativeSignals ?? []}
        />
      </section>

      {/* Reuse Snapshot 区块 — 直接复用 dashboard 的 ReuseSnapshotSection
          phase09-10 基线对齐：容器从 rounded-lg border bg-card p-3 改为 border-t pt-2，
          对齐 dashboard AssetFeedbackSection L118 与 ProductDetailPage L221 的轻量分隔方式，
          避免在 review 顶层 section 内嵌套双层卡片 */}
      <section className="space-y-2" aria-label="Reuse Snapshot">
        <h2 className="text-base font-semibold">复用感知快照</h2>
        <div className="border-t pt-2">
          {/* 重试能力由 read owner 暴露，页面不直接触碰 queryClient / queryKey。 */}
          <ReuseSnapshotSection
            status={review.pageState.reuseSnapshotSectionStatus}
            moduleReuseSummary={moduleReuseSummary}
            capabilitySummary={capabilitySummary}
            error={review.error}
            onRetry={review.retry}
          />
        </div>
      </section>
    </ReviewPageShell>
  )
}

// ============================================================================
// Overview Stat Bar — 紧凑 stat bar 风格，对齐 DashboardStatBar 但只含 overview 组
// ============================================================================

interface OverviewStat {
  label: string
  value: number
}

function OverviewStatBar({
  overview,
  status,
}: {
  overview: DashboardOverview | null
  status: 'ready' | 'empty' | 'error'
}) {
  if (status === 'error') {
    return (
      <div className="rounded-lg border border-destructive/50 bg-destructive/10 p-3">
        <p className="text-xs text-destructive">概览数据加载失败</p>
      </div>
    )
  }

  if (!overview) {
    return (
      <div className="flex flex-wrap items-stretch overflow-hidden rounded-lg border bg-card divide-x divide-border">
        {Array.from({ length: 6 }).map((_, i) => (
          <div key={i} className="flex min-w-[68px] flex-col justify-center px-3 py-2">
            <Skeleton className="h-5 w-8" />
            <Skeleton className="mt-1 h-2.5 w-10" />
          </div>
        ))}
      </div>
    )
  }

  const stats: OverviewStat[] = [
    { label: '模块', value: overview.module_count },
    { label: '产品', value: overview.product_count },
    { label: '仓库', value: overview.repository_count },
    { label: '决策', value: overview.decision_count },
    { label: '已绑仓', value: overview.product_with_repository_count },
    { label: '已绑模', value: overview.product_with_module_count },
  ]

  return (
    <div className="flex flex-wrap items-stretch overflow-hidden rounded-lg border bg-card divide-x divide-border">
      {stats.map((stat) => (
        <div
          key={stat.label}
          className="flex min-w-[68px] flex-col justify-center px-3 py-2"
        >
          <span className="text-lg font-bold leading-none tabular-nums">
            {stat.value}
          </span>
          <span className="mt-1 text-[10px] text-muted-foreground">{stat.label}</span>
        </div>
      ))}
    </div>
  )
}

// ============================================================================
// Recent Activity 列表 — 复用 RecentActivityItemCard
// ============================================================================

function RecentActivityList({
  status,
  activities,
}: {
  status: 'ready' | 'empty' | 'error'
  activities: RecentActivityItem[]
}) {
  if (status === 'empty') {
    return (
      <div className="rounded-lg border border-dashed p-4 text-center">
        <p className="text-xs text-muted-foreground">暂无最近活动</p>
      </div>
    )
  }

  if (status === 'error') {
    return (
      <div className="rounded-lg border border-destructive/50 bg-destructive/10 p-3">
        <p className="text-xs text-destructive">最近活动加载失败</p>
      </div>
    )
  }

  if (activities.length === 0) {
    return (
      <div className="divide-y divide-border overflow-hidden rounded-lg border">
        {Array.from({ length: 4 }).map((_, i) => (
          <div key={i} className="px-3 py-2">
            <Skeleton className="h-5 w-full" />
          </div>
        ))}
      </div>
    )
  }

  return (
    <div className="max-h-[320px] divide-y divide-border overflow-y-auto rounded-lg border md:max-h-[420px]">
      {activities.map((activity, index) => (
        <RecentActivityItemCard
          key={`${activity.activity_type}-${activity.target_id}-${index}`}
          activity={activity}
          section="recent-activity"
          getSearch={() => buildReviewSourceParams('weekly') as unknown as Record<string, unknown>}
        />
      ))}
    </div>
  )
}

// ============================================================================
// Representative Signals 列表 — 复用 FeedbackSignalCard
// ============================================================================

function RepresentativeSignalsList({
  status,
  signals,
}: {
  status: 'ready' | 'empty' | 'error'
  signals: FeedbackSignal[]
}) {
  if (status === 'empty') {
    return (
      <div className="rounded-lg border border-dashed p-4 text-center">
        <p className="text-xs text-muted-foreground">暂无代表性反馈信号</p>
      </div>
    )
  }

  if (status === 'error') {
    return (
      <div className="rounded-lg border border-destructive/50 bg-destructive/10 p-3">
        <p className="text-xs text-destructive">代表性反馈信号加载失败</p>
      </div>
    )
  }

  if (signals.length === 0) {
    return (
      <div className="divide-y divide-border overflow-hidden rounded-lg border">
        {Array.from({ length: 2 }).map((_, i) => (
          <div key={i} className="px-3 py-2">
            <Skeleton className="h-5 w-full" />
          </div>
        ))}
      </div>
    )
  }

  return (
    <div className="divide-y divide-border overflow-hidden rounded-lg border">
      {signals.map((signal, index) => (
        <FeedbackSignalCard
          key={`${signal.signal_code}-${signal.target_id}-${index}`}
          signal={signal}
          section="asset-feedback"
          getSearch={() => buildReviewSourceParams('weekly') as unknown as Record<string, unknown>}
        />
      ))}
    </div>
  )
}

// ============================================================================
// Template Candidate 选择区 — phase09-09 新增
// ============================================================================

function TemplateCandidateSection({
  candidates,
  activeCandidateId,
  onSelectCandidate,
  status,
  onCreateProduct,
}: {
  candidates: TemplateCandidateSummary[]
  activeCandidateId: string
  onSelectCandidate: (id: string) => void
  status: 'ready' | 'empty' | 'error'
  onCreateProduct: (candidateId: string) => void
}) {
  if (status === 'error') {
    return (
      <div className="rounded-lg border border-destructive/50 bg-destructive/10 p-3">
        <p className="text-xs text-destructive">模板候选加载失败，不影响其他数据</p>
      </div>
    )
  }

  if (status === 'empty' || candidates.length === 0) {
    return (
      <div className="rounded-lg border border-dashed p-4 text-center">
        <LayoutTemplate className="mx-auto h-6 w-6 text-muted-foreground/50 mb-2" />
        <p className="text-xs text-muted-foreground">当前没有可复用模板候选</p>
        <p className="text-[10px] text-muted-foreground/70 mt-1">
          当有多个产品共享相同模块组合时，将自动生成模板候选
        </p>
      </div>
    )
  }

  const activeCandidate = candidates.find((c) => c.templateCandidateId === activeCandidateId)

  return (
    <div className="space-y-2">
      {/* 候选卡片列表 */}
      <div className="space-y-1">
        {candidates.map((candidate) => {
          const isActive = candidate.templateCandidateId === activeCandidateId
          return (
            <button
              key={candidate.templateCandidateId}
              type="button"
              onClick={() => onSelectCandidate(candidate.templateCandidateId)}
              className={`w-full rounded-lg border p-2 text-left transition-colors ${
                isActive
                  ? 'ring-2 ring-primary border-primary bg-primary/5'
                  : 'border-border bg-card hover:bg-muted/50'
              }`}
            >
              <div className="flex items-center justify-between gap-2">
                <div className="min-w-0 flex-1">
                  <p className="text-sm font-medium truncate">{candidate.templateTitle}</p>
                  <p className="text-xs text-muted-foreground truncate mt-0.5">
                    {candidate.templateDescription}
                  </p>
                </div>
                <div className="flex items-center gap-2 shrink-0">
                  <span className="inline-flex items-center gap-1 text-[10px] text-muted-foreground">
                    <Layers className="h-3 w-3" />
                    {candidate.modules.length} 模块
                  </span>
                  <span className="inline-flex items-center gap-1 text-[10px] text-muted-foreground">
                    {candidate.sourceProductCount} 产品
                  </span>
                </div>
              </div>
              {/* 模块列表 — 仅在 active 时展示 */}
              {isActive && candidate.modules.length > 0 && (
                <div className="flex flex-wrap gap-1 mt-2">
                  {candidate.modules.map((m) => (
                    <span
                      key={m.moduleId}
                      className="inline-flex items-center rounded-md bg-muted px-2 py-0.5 text-[10px] font-medium"
                    >
                      {m.moduleName}
                    </span>
                  ))}
                </div>
              )}
            </button>
          )
        })}
      </div>

      {/* Active candidate CTA — 紧凑化按钮规格 h-7 px-2 text-xs，
          对齐 DerivedHintsSection 提示 CTA 与 dashboard 紧凑化基线 */}
      {activeCandidate && (
        <div className="border-t pt-2">
          <Button
            onClick={() => onCreateProduct(activeCandidate.templateCandidateId)}
            className="w-full h-7 px-2 text-xs"
            size="sm"
          >
            以该模板创建产品
            <ArrowRight className="ml-2 h-3 w-3" />
          </Button>
        </div>
      )}
    </div>
  )
}

// ============================================================================
// Derived Hints Section — phase09-10 派生提示展示区
// ============================================================================

function DerivedHintsSection({
  hints,
  activeCandidateId,
  hintsSectionStatus,
  onRetry,
}: {
  hints: DerivedInsightHint[]
  activeCandidateId: string
  hintsSectionStatus: 'ready' | 'empty' | 'error'
  onRetry: () => Promise<unknown>
}) {
  const navigate = useNavigate()

  // 单一提示 handoff owner
  const handoff = useDerivedHintHandoff({
    sourceSurface: 'weekly-review',
    activeCandidateId,
    templateSource: 'weekly-review',
  })

  // 过滤出合法提示（四元组完整）
  const validHints = hints.filter((h) => handoff.isValidHint(h))

  // 按类型分组
  const reuseHints = validHints.filter(
    (h) => h.hintType === DerivedInsightHintType.REUSE_OPPORTUNITY,
  )
  const gapHints = validHints.filter(
    (h) => h.hintType === DerivedInsightHintType.CAPABILITY_GAP,
  )

  if (hintsSectionStatus === 'error') {
    return (
      <section className="space-y-2" aria-label="Derived Hints">
        <h2 className="text-base font-semibold">派生智能提示</h2>
        <div className="rounded-lg border border-destructive/50 bg-destructive/10 p-3">
          <div className="flex items-start gap-2">
            <AlertCircle className="h-4 w-4 text-destructive mt-0.5 shrink-0" />
            <div>
              <p className="text-xs text-destructive">提示数据加载失败，不影响其他功能</p>
              <Button
                variant="outline"
                size="sm"
                className="mt-2 h-7 px-2 text-xs"
                onClick={() => {
                  void onRetry()
                }}
              >
                重试
              </Button>
            </div>
          </div>
        </div>
      </section>
    )
  }

  // 无 active candidate 或无有效提示时退回成功空态
  if (hintsSectionStatus === 'empty' || !activeCandidateId || validHints.length === 0) {
    return null
  }

  const handleHintCTA = (hint: DerivedInsightHint) => {
    const result = handoff.computeHandoff(hint)
    if (!result) {
      toast.error('无法处理该提示动作')
      return
    }
    navigate({
      to: result.to as any,
      params: result.params as any,
      search: result.search as any,
    })
  }

  return (
    <section className="space-y-2" aria-label="Derived Hints">
      <h2 className="text-base font-semibold">派生智能提示</h2>

      {/* 复用机会提示 — phase09-10 基线对齐：p-3 → p-2，text-sm → text-xs */}
      {reuseHints.map((hint) => (
        <div
          key={`reuse-${hint.templateCandidateId}`}
          className="rounded-lg border border-emerald-200 bg-emerald-50/50 p-2"
        >
          <div className="flex items-start gap-2">
            <Lightbulb className="h-4 w-4 text-emerald-600 mt-0.5 shrink-0" />
            <div className="min-w-0 flex-1">
              <p className="text-xs font-medium text-emerald-800">{hint.title}</p>
              <p className="text-[10px] text-emerald-700 mt-0.5">{hint.explanationText}</p>
              <Button
                variant="outline"
                size="sm"
                className="mt-2 h-7 px-2 text-xs border-emerald-300 text-emerald-700 hover:bg-emerald-100"
                onClick={() => handleHintCTA(hint)}
              >
                基于模板创建产品
                <ExternalLink className="ml-1 h-3 w-3" />
              </Button>
            </div>
          </div>
        </div>
      ))}

      {/* 能力缺口提示 — phase09-10 基线对齐：p-3 → p-2，text-sm → text-xs */}
      {gapHints.map((hint) => (
        <div
          key={`gap-${hint.templateCandidateId}-${hint.capabilityKey ?? ''}`}
          className="rounded-lg border border-amber-200 bg-amber-50/50 p-2"
        >
          <div className="flex items-start gap-2">
            <AlertCircle className="h-4 w-4 text-amber-600 mt-0.5 shrink-0" />
            <div className="min-w-0 flex-1">
              <p className="text-xs font-medium text-amber-800">{hint.title}</p>
              <p className="text-[10px] text-amber-700 mt-0.5">{hint.explanationText}</p>
              <Button
                variant="outline"
                size="sm"
                className="mt-2 h-7 px-2 text-xs border-amber-300 text-amber-700 hover:bg-amber-100"
                onClick={() => handleHintCTA(hint)}
              >
                {hint.moduleId ? '查看模块详情' : '前往 Module Registry'}
                <ExternalLink className="ml-1 h-3 w-3" />
              </Button>
            </div>
          </div>
        </div>
      ))}
    </section>
  )
}
