/**
 * DecisionDetailPage — Decision Detail
 *
 * §5.8 统一详情读模型宿主：承接详情读取、已关联目标展示、
 * 待关联目标承接、候选读取与最小目标关联动作。
 *
 * §5.11 入口上下文与正式关联结果边界：
 * - source_context.source_module_id 存在且 linked_modules 中尚无该 Module 时，
 *   DecisionPendingLinkTargetCard 必须显式展示该待关联目标
 * - 不得在进入详情页后静默丢失该待关联目标
 * - 待关联目标仅在正式 LinkDecisionToTarget 写入后消失（由页面 reread 驱动）
 *   当前阶段不提供“主动放弃关联”出口，source_context 作为入口历史记录保留
 *
 * §9.1 列表上下文跨页面恢复（单值化）：
 * - fromList === true：用户从 DecisionListPage 进入，返回列表时恢复 lastSearch
 * - fromList 不存在：用户从 Module Detail 入口或外部直达进入，返回列表时落到默认参数，
 *   不恢复历史筛选，避免错误继承旧上下文
 *
 * 布局降级（phase03-05 §"布局降级策略"）：
 * - PC：分区式详情布局，概要、已关联目标、待关联目标与候选关联区同页可见
 * - 移动浏览器：按概要、已关联目标、待关联目标、候选读取与目标关联的垂直顺序重排
 */
import { useParams, useNavigate, useSearch } from '@tanstack/react-router'
import { useDecisionDetailPageRead } from '../data/use-decision-detail-page-read'
import { useDecisionDetailActions } from '../application/use-decision-detail-actions'
import { DecisionDetailSummaryCard } from '../components/decision-detail-summary-card'
import { DecisionLinkedTargetsSection } from '../components/decision-linked-targets-section'
import { DecisionPendingLinkTargetCard } from '../components/decision-pending-link-target-card'
import { DecisionModuleCandidatePanel } from '../components/decision-module-candidate-panel'
import { Button } from '@/components/ui/button'
import { ArrowLeft } from 'lucide-react'
import { Skeleton } from '@/components/ui/skeleton'
import { useDecisionListSearchStore } from '../stores/decision-list-search-store'
import { BackToDashboardButton } from '@/features/dashboard/components/back-to-dashboard-button'
import { mergeCurrentDashboardSource } from '@/features/dashboard/lib/dashboard-source'
import type { DecisionStatus } from '../types'

export function DecisionDetailPage() {
  const { decisionId } = useParams({ from: '/decisions/$decisionId' })
  const navigate = useNavigate()
  // §9.1 fromList 单值化“来源列表上下文存在 / 不存在”
  const detailSearch = useSearch({ from: '/decisions/$decisionId' })
  // §9.1 从 store 读取最后一次列表搜索上下文
  const lastSearch = useDecisionListSearchStore((s) => s.lastSearch)
  const moduleDetailSearch = mergeCurrentDashboardSource(
    {},
    detailSearch,
  ) as unknown as Record<string, unknown>

  const pageRead = useDecisionDetailPageRead(decisionId, detailSearch, lastSearch)
  const detailActions = useDecisionDetailActions(decisionId)
  const { data, isLoading, isError, error } = pageRead.decisionDetailQuery

  // fix_002_003：状态推进 handler，单值化承接 canonical 状态变更
  const handleStatusChange = async (status: DecisionStatus) => {
    await detailActions.advanceDecisionStatus(status)
  }

  // 返回按钮统一通过 page read owner 产出的单值返回目标承接。
  const handleReturn = () => {
    switch (pageRead.returnTo.kind) {
      case 'onboarding':
        navigate({ to: '/onboarding', search: pageRead.returnTo.search })
        return
      case 'review':
        navigate({
          to: pageRead.returnTo.to as '/reviews/daily' | '/reviews/weekly',
          search: pageRead.returnTo.search,
        })
        return
      case 'decision-list':
        navigate({ to: '/decisions', search: pageRead.returnTo.search })
        return
    }
  }

  if (isError) {
    return (
      <div className="space-y-4">
        {/* phase05-13：从 Dashboard 进入时同时展示"返回 Dashboard"与"返回列表" */}
        <div className="flex items-center gap-2">
          <BackToDashboardButton />
          <Button variant="ghost" size="sm" onClick={handleReturn}>
            <ArrowLeft className="mr-2 h-4 w-4" />
            {pageRead.returnLabel}
          </Button>
        </div>
        <div className="rounded-lg border border-destructive/50 bg-destructive/10 p-4">
          <p className="text-sm text-destructive">详情读取失败：{(error as Error).message}</p>
        </div>
      </div>
    )
  }

  if (isLoading || !data) {
    return (
      <div className="space-y-4">
        <div className="flex items-center gap-2">
          <BackToDashboardButton />
          <Button variant="ghost" size="sm" onClick={handleReturn}>
            <ArrowLeft className="mr-2 h-4 w-4" />
            {pageRead.returnLabel}
          </Button>
        </div>
        <Skeleton className="h-48 w-full" />
        <Skeleton className="h-32 w-full" />
        <Skeleton className="h-32 w-full" />
      </div>
    )
  }

  // §5.11 派生待关联目标：
  // source_context.source_module_id 存在且 linked_modules 中尚无该 Module 时，
  // 需要显式展示为待关联目标
  const pendingTargetModuleId = data.source_context.source_module_id
  const hasPendingLinkTarget =
    pendingTargetModuleId !== '' &&
    !data.linked_modules.some((lm) => lm.module_id === pendingTargetModuleId)

  return (
    <div className="space-y-4">
      {/* 返回列表 — §9.1 按 fromList 单值化决定恢复 lastSearch 或落默认参数 */}
      {/* phase05-13：从 Dashboard 进入时同时展示"返回 Dashboard"与"返回列表" */}
      {/* phase06-15：fromOnboarding=true 时返回 /onboarding 并恢复 onboardingStep */}
      <div className="flex items-center gap-2">
        <BackToDashboardButton />
        <Button variant="ghost" size="sm" onClick={handleReturn}>
          <ArrowLeft className="mr-2 h-4 w-4" />
          {pageRead.returnLabel}
        </Button>
      </div>

      {/* PC：分区式布局；移动端：垂直顺序重排 */}
      <div className="grid gap-4 lg:grid-cols-3">
        {/* 概要主区 — 占 1 列（PC）/ 全宽（移动） */}
        <div className="lg:col-span-1">
          <DecisionDetailSummaryCard
            decision={data.decision}
            sourceContext={data.source_context}
            statusActions={pageRead.statusActions}
            onStatusChange={handleStatusChange}
            isUpdating={detailActions.isSubmitting}
          />
        </div>

        {/* 已关联目标 + 待关联目标 + 候选关联区 — 占 2 列（PC）/ 全宽（移动） */}
        <div className="space-y-4 lg:col-span-2">
          {/* §5.8 已关联目标展示 */}
          <DecisionLinkedTargetsSection
            linkedModules={data.linked_modules}
            moduleDetailSearch={moduleDetailSearch}
          />

          {/* §5.11 待关联目标展示 — 仅在有待关联目标时展示，正式关联后由 reread 驱动消失 */}
          {hasPendingLinkTarget && (
            <DecisionPendingLinkTargetCard
              sourceModuleId={pendingTargetModuleId}
              sourceModuleName={data.source_context.source_module_name}
              moduleDetailSearch={moduleDetailSearch}
            />
          )}

          {/* §5.10 候选读取与最小目标关联 */}
          <DecisionModuleCandidatePanel
            decisionId={decisionId}
            moduleDetailSearch={moduleDetailSearch}
          />
        </div>
      </div>
    </div>
  )
}
