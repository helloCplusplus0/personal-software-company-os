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
import { useQuery } from '@tanstack/react-query'
import { fetchDecisionDetail } from '../data/decision-center-adapter'
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
import {
  shouldReturnToOnboarding,
  buildOnboardingReturnSearch,
} from '@/features/onboarding/lib/onboarding-return'

export function DecisionDetailPage() {
  const { decisionId } = useParams({ from: '/decisions/$decisionId' })
  const navigate = useNavigate()
  // §9.1 fromList 单值化“来源列表上下文存在 / 不存在”
  const detailSearch = useSearch({ from: '/decisions/$decisionId' })
  // §9.1 从 store 读取最后一次列表搜索上下文
  const lastSearch = useDecisionListSearchStore((s) => s.lastSearch)
  // phase06-15 §"detail 页来源优先级"：fromOnboarding 优先级高于其他来源
  const fromOnboarding = shouldReturnToOnboarding(detailSearch)
  const returnLabel = fromOnboarding ? '返回首轮录入' : '返回列表'
  // §9.1 单值化返回参数：
  // - fromOnboarding=true → 返回 /onboarding 并恢复 onboardingStep
  // - fromList === true（从 DecisionListPage 进入）：返回列表恢复 lastSearch
  // - fromList 不存在（从 Module Detail 入口或外部直达进入）：返回列表落默认参数，不恢复历史筛选
  const returnSearch = fromOnboarding
    ? (buildOnboardingReturnSearch(detailSearch) as Record<string, unknown>)
    : (mergeCurrentDashboardSource(
        detailSearch.fromList
          ? {
              queryText: detailSearch.queryText ?? lastSearch.queryText,
              statusFilter: detailSearch.statusFilter ?? lastSearch.statusFilter,
            }
          : { statusFilter: 'all' as const },
        detailSearch,
      ) as unknown as Record<string, unknown>)

  const { data, isLoading, isError, error } = useQuery({
    queryKey: ['decision-detail', decisionId],
    queryFn: () => fetchDecisionDetail(decisionId),
    enabled: Boolean(decisionId),
  })

  // phase06-15：返回按钮统一通过 handleReturn 承接，支持 fromOnboarding 优先级
  const handleReturn = () => {
    if (fromOnboarding) {
      navigate({ to: '/onboarding', search: buildOnboardingReturnSearch(detailSearch) })
      return
    }
    navigate({ to: '/decisions', search: returnSearch })
  }

  if (isError) {
    return (
      <div className="space-y-4">
        {/* phase05-13：从 Dashboard 进入时同时展示"返回 Dashboard"与"返回列表" */}
        <div className="flex items-center gap-2">
          <BackToDashboardButton />
          <Button variant="ghost" size="sm" onClick={handleReturn}>
            <ArrowLeft className="mr-2 h-4 w-4" />
            {returnLabel}
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
            {returnLabel}
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
          {returnLabel}
        </Button>
      </div>

      {/* PC：分区式布局；移动端：垂直顺序重排 */}
      <div className="grid gap-4 lg:grid-cols-3">
        {/* 概要主区 — 占 1 列（PC）/ 全宽（移动） */}
        <div className="lg:col-span-1">
          <DecisionDetailSummaryCard
            decision={data.decision}
            sourceContext={data.source_context}
          />
        </div>

        {/* 已关联目标 + 待关联目标 + 候选关联区 — 占 2 列（PC）/ 全宽（移动） */}
        <div className="space-y-4 lg:col-span-2">
          {/* §5.8 已关联目标展示 */}
          <DecisionLinkedTargetsSection linkedModules={data.linked_modules} />

          {/* §5.11 待关联目标展示 — 仅在有待关联目标时展示，正式关联后由 reread 驱动消失 */}
          {hasPendingLinkTarget && (
            <DecisionPendingLinkTargetCard
              sourceModuleId={pendingTargetModuleId}
              sourceModuleName={data.source_context.source_module_name}
            />
          )}

          {/* §5.10 候选读取与最小目标关联 */}
          <DecisionModuleCandidatePanel decisionId={decisionId} />
        </div>
      </div>
    </div>
  )
}
