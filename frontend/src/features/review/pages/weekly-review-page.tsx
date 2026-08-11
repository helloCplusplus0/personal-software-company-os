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
import { useNavigate } from '@tanstack/react-router'
import { ReviewPageShell } from '../components/review-page-shell'
import { ReviewActionFooter } from '../components/review-action-footer'
import { useWeeklyReviewRead } from '../data/use-weekly-review-read'
import { useReviewAction } from '../application/use-review-action'
import { useNavigateBackToDashboard } from '@/features/dashboard/lib/dashboard-source'
import { FeedbackSignalCard } from '@/features/dashboard/components/feedback-signal-card'
import { RecentActivityItemCard } from '@/features/dashboard/components/recent-activity-item-card'
import { ReuseSnapshotSection } from '@/features/dashboard/components/reuse-snapshot-section'
import { Skeleton } from '@/components/ui/skeleton'
import type { FeedbackSignal, RecentActivityItem, DashboardOverview } from '@/features/dashboard/types'
import type {
  ModuleReuseSummaryEntry,
  CapabilitySummaryEntry,
} from '@/features/reuse-summary/types'
import type { ReviewActionInput } from '../application/review-action-types'
import { toast } from 'sonner'

export function WeeklyReviewPage() {
  const navigate = useNavigate()
  const navigateBack = useNavigateBackToDashboard()
  const review = useWeeklyReviewRead()
  const action = useReviewAction()

  const handleBackToDashboard = () => {
    navigateBack('overview')
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
      subtitle="本周资产盘点与整理"
      pageStatus={review.pageState.pageStatus}
      onBackToDashboard={handleBackToDashboard}
      actionFooter={
        <ReviewActionFooter
          isPending={action.isPending}
          hasError={action.isError}
          errorMessage={action.error?.message}
          onSubmitAction={handleSubmitAction}
          onReset={action.reset}
          reviewKind="weekly"
        />
      }
    >
      {/* Overview 区块 — 改为紧凑 stat bar 风格，对齐 DashboardStatBar
          移动端 flex-wrap 不撑破容器 */}
      <section className="space-y-2" aria-label="Overview">
        <h2 className="text-base font-semibold">系统概览</h2>
        <OverviewStatBar overview={overview ?? null} status={review.pageState.overviewSectionStatus} />
      </section>

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

      {/* Reuse Snapshot 区块 — 直接复用 dashboard 的 ReuseSnapshotSection */}
      <section className="space-y-2" aria-label="Reuse Snapshot">
        <h2 className="text-base font-semibold">复用感知快照</h2>
        {/* ReuseSnapshotSection 内部已建立 loading/ready/empty/error 四态样式
            作为 review 顶层区块用 border bg-card 包裹，对齐 dashboard 紧凑列表的容器风格 */}
          <div className="rounded-lg border bg-card p-3">
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
        />
      ))}
    </div>
  )
}
