import { useParams, useNavigate, useSearch } from '@tanstack/react-router'
import { useModuleDetailRead } from '../data/use-module-detail-read'
import { ModuleSummaryCard } from '../components/module-summary-card'
import { ModuleReleaseListSection } from '../components/module-release-list-section'
import { ModuleBindingPanel } from '../components/module-binding-panel'
import { ModuleDecisionEntryPanel } from '../components/module-decision-entry-panel'
import { Button } from '@/components/ui/button'
import { ArrowLeft } from 'lucide-react'
import { Skeleton } from '@/components/ui/skeleton'
import { useModuleListSearchStore } from '../stores/module-list-search-store'
import { BackToDashboardButton } from '@/features/dashboard/components/back-to-dashboard-button'
import { mergeCurrentDashboardSource } from '@/features/dashboard/lib/dashboard-source'
import {
  shouldReturnToOnboarding,
  buildOnboardingReturnSearch,
} from '@/features/onboarding/lib/onboarding-return'
import { useReuseSummaryRead } from '@/features/reuse-summary/data/use-reuse-summary-read'
import { ReuseSummaryInline } from '@/features/reuse-summary/components/reuse-summary-inline'
import { ModuleNextActionBar } from '../components/module-next-action-bar'
import { buildReviewReturnSearch, shouldReturnToReview } from '@/features/review/lib/review-source'

/**
 * ModuleDetailPage — Module Detail
 *
 * §8.4 统一回流宿主：承接创建完成后的落点与版本登记后的回流
 * §5.7 详情读取：核心字段、版本列表、产品绑定、仓库映射与 Decision 入口
 * §6.3 Decision 读取内嵌于 ModuleDetailRead，不设独立读接口组
 *
 * phase04-13 收敛：
 * - ModuleBindingPanel 已从直接写入承接位回落为只读摘要展示与兼容跳转入口
 * - 绑定写入统一迁移到 ProductDetailPage 与 RepositoryBindingDetailPage 承接
 * - 本页不再承接 BindModuleToProduct / MapModuleToRepository 的正式前端写入流程
 *
 * 布局降级（§8.5）：
 * - PC：分区式详情布局，摘要、版本与关联入口可同时可见
 * - 移动：摘要、版本、关联、Decision 入口按垂直顺序重排
 */
export function ModuleDetailPage() {
  const { moduleId } = useParams({ from: '/modules/$moduleId' })
  const detailSearch = useSearch({ from: '/modules/$moduleId' })
  const navigate = useNavigate()
  // §7.4 从 store 读取最后一次列表搜索上下文，返回列表时恢复
  const lastSearch = useModuleListSearchStore((s) => s.lastSearch)
  // phase06-15 §"detail 页来源优先级"：fromOnboarding 优先级高于其他来源
  const fromOnboarding = shouldReturnToOnboarding(detailSearch)
  const fromReview = shouldReturnToReview(detailSearch)
  // phase09-10 模板复用提示返回链：returnTo 优先级高于 fromList
  const fromTemplateReturn = (detailSearch as any).returnTo === 'weekly-review' || (detailSearch as any).returnTo === 'product-create'
  const returnLabel = fromOnboarding
    ? '返回首轮录入'
    : fromReview
      ? `返回 ${detailSearch.reviewKind === 'weekly' ? 'Weekly Review' : 'Daily Review'}`
    : fromTemplateReturn
      ? ((detailSearch as any).returnTo === 'weekly-review' ? '返回 Weekly Review' : '返回创建产品')
      : '返回列表'
  const returnSearch = fromOnboarding
    ? (buildOnboardingReturnSearch(detailSearch) as Record<string, unknown>)
    : fromReview
      ? (buildReviewReturnSearch(detailSearch) as Record<string, unknown>)
    : fromTemplateReturn
      ? ((): Record<string, unknown> => {
          const s: Record<string, unknown> = {}
          if ((detailSearch as any).returnCandidateId) {
            s.returnCandidateId = (detailSearch as any).returnCandidateId
          }
          if ((detailSearch as any).fromTemplateReuse) {
            s.fromTemplateReuse = true
            s.templateCandidateId = (detailSearch as any).templateCandidateId
            s.templateSource = (detailSearch as any).templateSource
          }
          if ((detailSearch as any).fromDashboard) {
            s.fromDashboard = true
            s.dashboardSection = (detailSearch as any).dashboardSection
            s.dashboardReturnTo = (detailSearch as any).dashboardReturnTo
          }
          return s
        })()
      : (mergeCurrentDashboardSource(
          detailSearch.fromList
            ? {
                queryText: detailSearch.queryText,
                statusFilter: detailSearch.statusFilter ?? 'all',
              }
            : lastSearch,
          detailSearch,
        ) as unknown as Record<string, unknown>)
  const decisionDetailSearch = fromOnboarding
    ? (buildOnboardingReturnSearch(detailSearch) as Record<string, unknown>)
    : fromReview
      ? (buildReviewReturnSearch(detailSearch) as Record<string, unknown>)
      : detailSearch.fromDashboard === true
        ? (mergeCurrentDashboardSource({}, detailSearch) as Record<string, unknown>)
        : undefined

  const { data, isLoading, isError, error } = useModuleDetailRead(moduleId)

  // phase06-15 §"Module Detail 与 Product Detail 挂接位"：
  // Module Detail 只新增一个页面级 ReuseSummaryRead query（scope=module_detail）
  // 失败不回退整页，只影响复用摘要内联组件
  const reuseSummaryQuery = useReuseSummaryRead(
    { scope: 'module_detail', module_id: moduleId },
    { enabled: Boolean(moduleId) },
  )

  const reuseSummaryStatus: 'loading' | 'ready' | 'empty' | 'error' = reuseSummaryQuery.isLoading
    ? 'loading'
    : reuseSummaryQuery.isError
      ? 'error'
      : (reuseSummaryQuery.data?.module_reuse_summary?.length ?? 0) === 0 &&
          (reuseSummaryQuery.data?.capability_summary?.length ?? 0) === 0
        ? 'empty'
        : 'ready'

  // phase06-15：返回按钮统一通过 handleReturn 承接，支持 fromOnboarding 优先级
  // phase09-10：模板复用提示返回链（returnTo=weekly-review/product-create）优先级高于 fromList
  const handleReturn = () => {
    if (fromOnboarding) {
      navigate({ to: '/onboarding', search: buildOnboardingReturnSearch(detailSearch) })
      return
    }
    if (fromReview) {
      navigate({
        to: detailSearch.reviewReturnTo ?? '/reviews/daily',
        search: buildReviewReturnSearch(detailSearch) as Record<string, unknown>,
      })
      return
    }
    if (fromTemplateReturn) {
      const rs = (detailSearch as any).returnTo as string
      if (rs === 'weekly-review') {
        navigate({ to: '/reviews/weekly', search: returnSearch })
        return
      }
      if (rs === 'product-create') {
        navigate({
          to: '/products/new',
          search: returnSearch,
        })
        return
      }
    }
    navigate({ to: '/modules', search: returnSearch })
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
        <Skeleton className="h-32 w-full" />
        <Skeleton className="h-48 w-full" />
        <Skeleton className="h-48 w-full" />
      </div>
    )
  }

  return (
    <div className="space-y-4">
      {/* 返回列表 — §7.4 恢复原有搜索参数 */}
      {/* phase05-13：从 Dashboard 进入时同时展示"返回 Dashboard"与"返回列表" */}
      {/* phase06-15：fromOnboarding=true 时返回 /onboarding 并恢复 onboardingStep */}
      <div className="flex items-center gap-2">
        <BackToDashboardButton />
        <Button variant="ghost" size="sm" onClick={handleReturn}>
          <ArrowLeft className="mr-2 h-4 w-4" />
          {returnLabel}
        </Button>
      </div>

      {/* phase10-10：页面级下一步动作区 */}
      <ModuleNextActionBar
        hasProductBinding={(data.product_bindings?.length ?? 0) > 0}
        hasRepositoryMapping={(data.repository_mappings?.length ?? 0) > 0}
        moduleId={moduleId}
        moduleName={data.module.name}
        decisionLinks={data.decision_links}
        decisionDetailSearch={decisionDetailSearch}
      />

      {/* PC：分区式布局；移动端：垂直顺序重排 */}
      <div className="grid gap-4 lg:grid-cols-3">
        {/* 摘要主区 — 占 1 列（PC）/ 全宽（移动） */}
        <div className="space-y-4 lg:col-span-1">
          <ModuleSummaryCard module={data.module} />
          {/*
            phase06-15 §"Module Detail 与 Product Detail 挂接位"：
            在 Module Summary 邻近区域挂接复用摘要内联组件

            phase06-16 §"紧凑型优化"：
            移除重型卡片包裹（rounded-lg border bg-muted/20 p-3），
            改用轻量顶部分隔线（border-t pt-2），避免与 ModuleSummaryCard 形成双层卡片嵌套。
          */}
          <div className="border-t pt-2">
            <ReuseSummaryInline
              status={reuseSummaryStatus}
              moduleReuseSummary={reuseSummaryQuery.data?.module_reuse_summary ?? []}
              capabilitySummary={reuseSummaryQuery.data?.capability_summary ?? []}
              error={reuseSummaryQuery.error as Error | null}
              invalidateQueryKey={['reuse-summary', 'module_detail', moduleId]}
              title="复用摘要"
            />
          </div>
        </div>

        {/* 版本列表区 + 关联动作区 — 占 2 列（PC）/ 全宽（移动） */}
        <div className="space-y-4 lg:col-span-2">
          <ModuleReleaseListSection
            releases={data.releases}
            moduleId={moduleId}
          />
          <ModuleBindingPanel
            moduleId={moduleId}
            moduleName={data.module.name}
            productBindings={data.product_bindings}
            repositoryMappings={data.repository_mappings}
          />
          {/* §6.3 Decision 入口：phase03-13 升级为正式入口触点 */}
          <ModuleDecisionEntryPanel
            moduleId={data.module.id}
            moduleName={data.module.name}
            decisionLinks={data.decision_links}
            decisionDetailSearch={decisionDetailSearch}
          />
        </div>
      </div>
    </div>
  )
}
