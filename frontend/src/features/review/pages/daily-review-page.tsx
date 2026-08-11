/**
 * DailyReviewPage — Daily Review 正式页面。
 *
 * phase08-08 §"Daily Review 最小 enablement"：
 *   - 通过 ReviewPageShell 承接头部、离开入口、页面级状态与底部动作区
 *   - 至少必须显示 current focus / pending decisions / representative signals 的最小只读组织
 *   - 必须接上 ReviewActionFooter 或等价完成区，能够正式触发 useReviewAction()
 *
 * phase08-08 UI 对齐 Dashboard 基线（验收后调整）：
 *   - 移除 Card 重型卡片，改用 dashboard 的 section + divide-y 紧凑列表
 *   - current focus / representative signals 直接复用既有 FeedbackSignalCard
 *   - pending decisions 行可点击进入既有 Decision Detail，携带 dashboardSection=current-focus
 *   - 区块标题样式对齐 dashboard：text-base font-semibold
 *   - 空态/错误态/骨架样式对齐 dashboard
 *
 * 约束：
 *   - 页面只能消费 useDailyReviewRead / useReviewAction
 *   - 不得直接 import 底层 canonical query hook
 *   - 不得在页面内直接 createClient()、createConnectTransport()、useMutation()
 */
import { useNavigate } from '@tanstack/react-router'
import { ReviewPageShell } from '../components/review-page-shell'
import { ReviewActionFooter } from '../components/review-action-footer'
import { useDailyReviewRead } from '../data/use-daily-review-read'
import { useReviewAction } from '../application/use-review-action'
import { useNavigateBackToDashboard, buildDashboardSourceParams } from '@/features/dashboard/lib/dashboard-source'
import { FeedbackSignalCard } from '@/features/dashboard/components/feedback-signal-card'
import { Skeleton } from '@/components/ui/skeleton'
import { ArrowRight } from 'lucide-react'
import type { FeedbackSignal } from '@/features/dashboard/types'
import type { ReviewActionInput } from '../application/review-action-types'
import type { PendingDecision } from '../data/use-daily-review-read'
import { toast } from 'sonner'

const DAILY_REVIEW_ACTION_SECTIONS = {
  decision: 'current-focus',
  product: 'asset-feedback',
  module: 'current-focus',
  repository: 'asset-feedback',
  submitNextStep: 'current-focus',
} as const

export function DailyReviewPage() {
  const navigate = useNavigate()
  const navigateBack = useNavigateBackToDashboard()
  const review = useDailyReviewRead()
  const action = useReviewAction()

  const handleBackToDashboard = () => {
    navigateBack('current-focus')
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

  return (
    <ReviewPageShell
        title="Daily Review"
        subtitle="今日要处理的焦点与决策"
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
          reviewKind="daily"
          actionSections={DAILY_REVIEW_ACTION_SECTIONS}
        />
      }
    >
      {/* Current Focus 区块 — 对齐 dashboard CurrentFocusSection 样式
          复用 FeedbackSignalCard，跳转后回到 dashboard 的 current-focus 区块 */}
      <section className="space-y-2" aria-label="Current Focus">
        <h2 className="text-base font-semibold">Current Focus</h2>
        <CurrentFocusList
          status={review.pageState.currentFocusSectionStatus}
          signals={review.data?.currentFocusSignals ?? []}
        />
      </section>

      {/* Pending Decisions 区块 — 对齐 dashboard 紧凑列表样式
          行可点击进入既有 Decision Detail，携带 dashboardSection=current-focus */}
      <section className="space-y-2" aria-label="Pending Decisions">
        <h2 className="text-base font-semibold">
          待处理决策
          {review.data && review.data.pendingDecisions.length > 0 && (
            <span className="ml-1.5 text-xs text-muted-foreground font-normal">
              ({review.data.pendingDecisions.length})
            </span>
          )}
        </h2>
        <PendingDecisionList
          status={review.pageState.pendingDecisionsSectionStatus}
          decisions={review.data?.pendingDecisions ?? []}
        />
      </section>

      {/* Representative Signals 区块 — 对齐 dashboard AssetFeedbackSection 样式
          复用 FeedbackSignalCard，跳转后回到 dashboard 的 asset-feedback 区块 */}
      <section className="space-y-2" aria-label="Representative Signals">
        <h2 className="text-base font-semibold">代表性反馈信号</h2>
        <RepresentativeSignalsList
          status={review.pageState.representativeSignalsSectionStatus}
          signals={review.data?.representativeSignals ?? []}
        />
      </section>
    </ReviewPageShell>
  )
}

// ============================================================================
// Current Focus 列表 — 复用 FeedbackSignalCard，section="current-focus"
// ============================================================================

function CurrentFocusList({
  status,
  signals,
}: {
  status: 'ready' | 'empty' | 'error'
  signals: FeedbackSignal[]
}) {
  if (status === 'empty') {
    return (
      <div className="rounded-lg border border-dashed p-4 text-center">
        <p className="text-xs text-muted-foreground">暂无待处理焦点信号</p>
      </div>
    )
  }

  if (status === 'error') {
    return (
      <div className="rounded-lg border border-destructive/50 bg-destructive/10 p-3">
        <p className="text-xs text-destructive">当前焦点加载失败</p>
      </div>
    )
  }

  if (signals.length === 0) {
    return (
      <div className="divide-y divide-border overflow-hidden rounded-lg border">
        {Array.from({ length: 3 }).map((_, i) => (
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
          section="current-focus"
        />
      ))}
    </div>
  )
}

// ============================================================================
// Pending Decision 列表 — 紧凑行可点击进入 Decision Detail
// ============================================================================

function PendingDecisionList({
  status,
  decisions,
}: {
  status: 'ready' | 'empty' | 'error'
  decisions: PendingDecision[]
}) {
  if (status === 'empty') {
    return (
      <div className="rounded-lg border border-dashed p-4 text-center">
        <p className="text-xs text-muted-foreground">暂无待处理决策</p>
      </div>
    )
  }

  if (status === 'error') {
    return (
      <div className="rounded-lg border border-destructive/50 bg-destructive/10 p-3">
        <p className="text-xs text-destructive">待处理决策加载失败</p>
      </div>
    )
  }

  if (decisions.length === 0) {
    return (
      <div className="divide-y divide-border overflow-hidden rounded-lg border">
        {Array.from({ length: 3 }).map((_, i) => (
          <div key={i} className="px-3 py-2">
            <Skeleton className="h-5 w-full" />
          </div>
        ))}
      </div>
    )
  }

  return (
    <div className="divide-y divide-border overflow-hidden rounded-lg border">
      {decisions.map((decision) => (
        <PendingDecisionCard key={decision.id} decision={decision} />
      ))}
    </div>
  )
}

/**
 * PendingDecisionCard — 待处理决策单值紧凑行。
 *
 * 整行可点击跳转，进入既有 /decisions/$decisionId canonical 路径，
 * 携带 buildDashboardSourceParams('current-focus') 来源参数，
 * 确保跳转后 BackToDashboardButton 仍然可用。
 */
function PendingDecisionCard({ decision }: { decision: PendingDecision }) {
  const navigate = useNavigate()

  const handleClick = () => {
    const sourceParams = buildDashboardSourceParams('current-focus')
    navigate({
      to: '/decisions/$decisionId',
      params: { decisionId: decision.id },
      search: sourceParams,
    })
  }

  return (
    <button
      type="button"
      onClick={handleClick}
      title={decision.linked_module_summary || decision.title}
      className="flex w-full items-center gap-2 px-3 py-2 text-left transition-colors hover:bg-muted/40"
    >
      {/* 状态 badge — 紧凑，蓝色调（proposed 状态） */}
      <span className="inline-flex h-5 min-w-[28px] shrink-0 items-center justify-center rounded bg-blue-500/10 px-1 text-[10px] font-semibold text-blue-600 dark:text-blue-400">
        {decision.status}
      </span>
      {/* 标题 — 主信息，允许占满剩余空间 */}
      <span className="min-w-0 flex-1 truncate text-sm font-medium">
        {decision.title}
      </span>
      {/* linked module summary — 次要信息，sm+ 屏幕展示 */}
      {decision.linked_module_summary && (
        <span className="hidden max-w-[140px] shrink-0 truncate text-xs text-muted-foreground sm:inline">
          {decision.linked_module_summary}
        </span>
      )}
      <ArrowRight className="h-4 w-4 shrink-0 text-muted-foreground" />
    </button>
  )
}

// ============================================================================
// Representative Signals 列表 — 复用 FeedbackSignalCard，section="asset-feedback"
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
