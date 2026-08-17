import { useParams, useSearch, useNavigate } from '@tanstack/react-router'
import { useQueryClient } from '@tanstack/react-query'
import { useProductDetailRead } from '../data/use-product-detail-read'
import { ProductSummaryCard } from '../components/product-summary-card'
import { ProductModuleBindingPanel } from '../components/product-module-binding-panel'
import { ProductBoundRepositoryListSection } from '../components/product-bound-repository-list-section'
import { Button } from '@/components/ui/button'
import { ArrowLeft } from 'lucide-react'
import { Skeleton } from '@/components/ui/skeleton'
import { BackToDashboardButton } from '@/features/dashboard/components/back-to-dashboard-button'
import { mergeCurrentDashboardSource } from '@/features/dashboard/lib/dashboard-source'
import {
  shouldReturnToOnboarding,
  buildOnboardingReturnSearch,
} from '@/features/onboarding/lib/onboarding-return'
import { useReuseSummaryRead } from '@/features/reuse-summary/data/use-reuse-summary-read'
import { ReuseSummaryInline } from '@/features/reuse-summary/components/reuse-summary-inline'
import { useTemplateSourceRead } from '@/features/template-reuse/data/use-template-source-read'
import { TemplateSource, TemplateConsumerSurface } from '@/gen/proto/psco/template_reuse/v1/template_reuse_pb'
import { ArrowDown } from 'lucide-react'
import { ProductNextActionBar } from '../components/product-next-action-bar'
import { ProductDecisionEntryPanel } from '../components/product-decision-entry-panel'
import { DAILY_REVIEW_QUERY_KEY, WEEKLY_REVIEW_QUERY_KEY } from '@/features/review/data/review-query-options'
import { buildReviewReturnSearch, shouldReturnToReview } from '@/features/review/lib/review-source'
import { useModuleDecisionLinksByModuleIds } from '@/features/module-registry/data/use-module-decision-links-by-module-ids'
import { PRODUCT_SEMANTIC_LABEL } from '@/features/project-context/data/shared-semantic-constants'

/**
 * ProductDetailPage — Product Detail
 *
 * phase04-06 状态模型：
 * - 详情读取状态：pending / success / error
 * - 资源不存在时派生 not-found 视图状态
 * - 错误反馈停留在详情页内容区域，不跳转独立错误页
 *
 * phase04-06 来源上下文（由路由搜索参数派生，只允许三种之一）：
 * - fromList 存在 → 来自 Product List，承接 queryText / statusFilter
 * - fromModuleDetail 存在 → 来自 Module Detail，承接 moduleId / moduleName（用于预填绑定面板）
 * - 无来源参数 → direct-entry
 * - 从 ProductCreatePage 成功创建后进入时，来源上下文继承自创建页
 *
 * phase04-06 主动返回路径：
 * - fromList → 回 Product List + 原 queryText / statusFilter
 * - fromModuleDetail → 回原 ModuleDetailPage
 * - direct-entry → 回 Product List 默认筛选参数
 * - 刷新后必须恢复来源标记
 *
 * phase04-06 BindModuleToProduct reread：
 * - 绑定成功后停留在当前 ProductDetailPage 并重新读取详情结果
 * - 不得只靠 toast 作为成功依据
 *
 * 布局降级（phase04-05）：
 * - PC：分区式详情布局，摘要、已绑定模块、已绑定仓库与绑定入口可同时可见
 * - 移动：摘要、已绑定模块、已绑定仓库按垂直顺序重排
 */
export function ProductDetailPage() {
  const { productId } = useParams({ from: '/products/$productId' })
  const search = useSearch({ from: '/products/$productId' })
  const navigate = useNavigate()
  const queryClient = useQueryClient()

  // phase04-06 来源上下文单值判定
  const fromList = search.fromList === true
  const fromModuleDetail = search.fromModuleDetail === true
  // phase06-15 §"detail 页来源优先级"：fromOnboarding 优先级高于其他来源
  const fromOnboarding = shouldReturnToOnboarding(search)
  const fromReview = shouldReturnToReview(search)

  const { data, isLoading, isError, error } = useProductDetailRead(productId)
  const relatedDecisionLinksQuery = useModuleDecisionLinksByModuleIds(
    data?.bound_modules?.map((module) => module.module_id) ?? [],
  )
  // phase06-15 §"Module Detail 与 Product Detail 挂接位"：
  // Product Detail 只新增一个页面级 ReuseSummaryRead query（scope=product_detail）
  // 在已绑定模块相关区域附近挂接；失败不回退整页，只影响复用摘要内联组件
  const reuseSummaryQuery = useReuseSummaryRead(
    { scope: 'product_detail', product_id: productId },
    { enabled: Boolean(productId) },
  )

  const reuseSummaryStatus: 'loading' | 'ready' | 'empty' | 'error' = reuseSummaryQuery.isLoading
    ? 'loading'
    : reuseSummaryQuery.isError
      ? 'error'
      : (reuseSummaryQuery.data?.module_reuse_summary?.length ?? 0) === 0 &&
          (reuseSummaryQuery.data?.capability_summary?.length ?? 0) === 0
        ? 'empty'
        : 'ready'

  // phase09-09：模板来源复读
  const fromTemplateReuse = search.fromTemplateReuse === true
  const templateCandidateId = (search as any).templateCandidateId ?? ''
  const templateSourceStr = (search as any).templateSource ?? ''

  const templateSourceQuery = useTemplateSourceRead(
    fromTemplateReuse && templateCandidateId !== '' ? templateCandidateId : '',
    templateSourceToEnum(templateSourceStr),
    TemplateConsumerSurface.PRODUCT_DETAIL,
  )

  // phase04-06 BindModuleToProduct 成功后重新读取详情结果（reread）
  // phase10-10：补齐 Dashboard / Review query 失效，确保返回后 reread 正确
  // phase13-09：phase12 project-context 前端消费已退出，不再失效该缓存键
  const invalidateDetail = () => {
    queryClient.invalidateQueries({ queryKey: ['product-detail', productId] })
    queryClient.invalidateQueries({ queryKey: ['product-list'] })
    queryClient.invalidateQueries({ queryKey: ['product-module-candidates', productId] })
    queryClient.invalidateQueries({ queryKey: ['dashboard-feedback-signals'] })
    queryClient.invalidateQueries({ queryKey: ['dashboard-overview'] })
    queryClient.invalidateQueries({ queryKey: DAILY_REVIEW_QUERY_KEY })
    queryClient.invalidateQueries({ queryKey: WEEKLY_REVIEW_QUERY_KEY })
  }

  // phase04-06 主动返回路径 — 按真实来源决定，刷新后恢复来源标记
  // phase06-15：fromOnboarding 优先级最高，先于 fromList / fromModuleDetail / direct-entry
  const handleReturn = () => {
    if (fromOnboarding) {
      navigate({
        to: '/onboarding',
        search: buildOnboardingReturnSearch(search),
      })
      return
    }
    if (fromReview) {
      navigate({
        to: search.reviewReturnTo ?? '/reviews/daily',
        search: buildReviewReturnSearch(search) as Record<string, unknown>,
      })
      return
    }
    if (fromList) {
      navigate({
        to: '/products',
          search: mergeCurrentDashboardSource(
            {
              queryText: search.queryText,
              statusFilter: search.statusFilter ?? 'all',
            },
            search,
          ),
      })
    } else if (fromModuleDetail && search.moduleId) {
      navigate({
        to: '/modules/$moduleId',
        params: { moduleId: search.moduleId },
          search: mergeCurrentDashboardSource({}, search) as Record<string, unknown>,
      })
    } else {
      // direct-entry → 回 Product List 默认筛选参数
      navigate({
        to: '/products',
          search: mergeCurrentDashboardSource({ statusFilter: 'all' as const }, search),
      })
    }
  }

  // 返回按钮文案：fromOnboarding 优先展示"返回首轮录入"
  const returnLabel = fromOnboarding
    ? '返回首轮录入'
    : fromReview
      ? `返回 ${search.reviewKind === 'weekly' ? 'Weekly Review' : 'Daily Review'}`
    : fromModuleDetail
      ? '返回模块详情'
      : '返回列表'
  const decisionDetailSearch = fromOnboarding
    ? (buildOnboardingReturnSearch(search) as Record<string, unknown>)
    : fromReview
      ? (buildReviewReturnSearch(search) as Record<string, unknown>)
      : search.fromDashboard === true
        ? (mergeCurrentDashboardSource({}, search) as Record<string, unknown>)
        : undefined

  if (isError) {
    return (
      <div className="space-y-4">
        {/* phase05-13：从 Dashboard 进入时同时展示"返回 Dashboard"与原生返回 */}
        <div className="flex items-center gap-2">
          <BackToDashboardButton />
          <Button variant="ghost" size="sm" onClick={handleReturn}>
            <ArrowLeft className="mr-2 h-4 w-4" />
            {returnLabel}
          </Button>
        </div>
        <div className="rounded-lg border border-destructive/50 bg-destructive/10 p-4">
          <p className="text-sm text-destructive">
            {error instanceof Error && 'status' in error && (error as { status: number }).status === 404
              ? '产品不存在'
              : `详情读取失败：${(error as Error).message}`}
          </p>
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
      {/* 返回 — phase04-06 按真实来源决定 */}
      {/* phase05-13：从 Dashboard 进入时同时展示"返回 Dashboard"与原生返回 */}
      <div className="flex items-center gap-2">
        <BackToDashboardButton />
        <Button variant="ghost" size="sm" onClick={handleReturn}>
          <ArrowLeft className="mr-2 h-4 w-4" />
          {returnLabel}
        </Button>
      </div>

      {/* phase04-06 来源上下文展示 — 从 Module Detail 带上下文进入时 */}
      {fromModuleDetail && search.moduleName && (
        <div className="rounded-lg border bg-muted/50 p-3 text-sm">
          <span className="text-muted-foreground">来源模块：</span>
          <span className="font-medium">{search.moduleName}</span>
        </div>
      )}

      {/* phase12-08：页面级语义导语 */}
      <p className="text-sm text-muted-foreground">
        该 Product 代表一个{PRODUCT_SEMANTIC_LABEL}
      </p>

      {/* phase10-10：页面级下一步动作区 */}
      <ProductNextActionBar
        hasRepository={(data.bound_repositories?.length ?? 0) > 0}
        hasModule={(data.bound_modules?.length ?? 0) > 0}
        productId={productId}
        productName={data.product.name}
        decisionLinks={relatedDecisionLinksQuery.decisionLinks}
        decisionDetailSearch={decisionDetailSearch}
      />

      {/* PC：分区式详情布局；移动端：垂直顺序重排 — phase04-05 */}
      <div className="grid gap-4 lg:grid-cols-3">
        {/* 摘要主区 — 占 1 列（PC）/ 全宽（移动） */}
        <div className="space-y-4 lg:col-span-1">
          <ProductSummaryCard product={data.product} />
          {/* phase09-09：模板来源摘要 — 位于 ProductSummaryCard 与 ReuseSummaryInline 之间 */}
          {fromTemplateReuse && (
            <TemplateSourceSummarySection
              sourceSummary={templateSourceQuery.sourceSummary}
              pageStatus={templateSourceQuery.pageStatus}
            />
          )}
          {/*
            phase06-15 §"Module Detail 与 Product Detail 挂接位"：
            在已绑定模块相关区域附近挂接复用摘要内联组件

            phase06-16 §"紧凑型优化"：
            移除重型卡片包裹（rounded-lg border bg-muted/20 p-3），
            改用轻量顶部分隔线（border-t pt-2），避免与 ProductSummaryCard 形成双层卡片嵌套。
          */}
          <div className="border-t pt-2">
            <ReuseSummaryInline
              status={reuseSummaryStatus}
              moduleReuseSummary={reuseSummaryQuery.data?.module_reuse_summary ?? []}
              capabilitySummary={reuseSummaryQuery.data?.capability_summary ?? []}
              error={reuseSummaryQuery.error as Error | null}
              invalidateQueryKey={['reuse-summary', 'product_detail', productId]}
              title="复用摘要"
            />
          </div>
        </div>

        {/* 绑定区 — 占 2 列（PC）/ 全宽（移动） */}
        <div className="space-y-4 lg:col-span-2">
          <div id="product-module-binding">
          <ProductModuleBindingPanel
            productId={productId}
            boundModules={data.bound_modules}
            prefillModuleId={fromModuleDetail ? search.moduleId : undefined}
            onBindingSuccess={invalidateDetail}
          />
          </div>
          <ProductBoundRepositoryListSection
            product={data.product}
            boundRepositories={data.bound_repositories}
            // phase04-06 透传 Product Detail 的来源上下文，使 Repository Binding 返回时能恢复来源标记
            productSource={{
              fromList,
              queryText: search.queryText,
              statusFilter: search.statusFilter,
              fromModuleDetail,
              moduleId: search.moduleId,
              moduleName: search.moduleName,
            }}
          />
          {/* phase10-10：Product Detail Decision 入口面板 */}
          <ProductDecisionEntryPanel
            decisionLinks={relatedDecisionLinksQuery.decisionLinks}
            decisionDetailSearch={decisionDetailSearch}
            isLoading={relatedDecisionLinksQuery.isLoading}
            isError={relatedDecisionLinksQuery.isError}
          />
        </div>
      </div>
    </div>
  )
}

// ============================================================================
// 辅助函数
// ============================================================================

/** 字符串 → TemplateSource 枚举 */
function templateSourceToEnum(s: string): TemplateSource {
  switch (s) {
    case 'weekly-review': return TemplateSource.WEEKLY_REVIEW
    case 'dashboard': return TemplateSource.DASHBOARD
    case 'product-detail': return TemplateSource.PRODUCT_DETAIL
    default: return TemplateSource.UNSPECIFIED
  }
}

// ============================================================================
// 模板来源摘要组件 — phase09-09 新增
// ============================================================================

function TemplateSourceSummarySection({
  sourceSummary,
  pageStatus,
}: {
  sourceSummary: {
    templateTitle: string
    templateDescription: string
    modules: { moduleId: string; moduleName: string }[]
    templateSource: number
    resolutionStatus: number
    unavailableReasonText: string
  } | undefined
  pageStatus: 'initial-loading' | 'resolved' | 'unavailable' | 'error'
}) {
  if (pageStatus === 'initial-loading') {
    return (
      <div className="border-t pt-2">
        <Skeleton className="h-4 w-24 mb-1" />
        <Skeleton className="h-5 w-40" />
        <div className="flex gap-1 mt-1">
          <Skeleton className="h-4 w-16" />
          <Skeleton className="h-4 w-20" />
        </div>
      </div>
    )
  }

  if (pageStatus === 'error') {
    return (
      <div className="border-t pt-2">
        <div className="rounded-lg border border-red-200 bg-red-50 p-2 text-xs">
          <p className="text-red-700">模板来源加载失败</p>
        </div>
      </div>
    )
  }

  if (pageStatus === 'unavailable' || !sourceSummary) {
    return (
      <div className="border-t pt-2">
        <div className="rounded-lg border border-amber-200 bg-amber-50 p-2 text-xs">
          <p className="text-amber-700">模板来源已不可复读</p>
        </div>
        {/* canonical binding CTA — unavailable 状态下仍必须可见 */}
        <div className="mt-2">
          <a
            href="#product-module-binding"
            className="inline-flex items-center gap-1 text-xs text-primary hover:underline"
          >
            为模板模块绑定仓库
            <ArrowDown className="h-3 w-3" />
          </a>
        </div>
      </div>
    )
  }

  const sourceLabel = sourceSummary.templateSource === 1
    ? 'Weekly Review'
    : sourceSummary.templateSource === 2
      ? 'Dashboard'
      : 'Product Detail'

  return (
    <div className="border-t pt-2">
      {/* phase09-10 基线对齐：字号对齐 ReuseSnapshotSection 紧凑化规范
          - 来源标签：text-[10px]（元信息）
          - 模板标题：text-xs font-semibold（主信息，对齐 ReuseSnapshotSection 标题）
          - 模板描述：text-[10px]（次要信息） */}
      <div className="text-[10px] text-muted-foreground mb-1">
        来源：{sourceLabel}
      </div>
      <h4 className="text-xs font-semibold">{sourceSummary.templateTitle}</h4>
      <p className="text-[10px] text-muted-foreground mt-0.5">
        {sourceSummary.templateDescription}
      </p>
      {sourceSummary.modules.length > 0 && (
        <div className="flex flex-wrap gap-1 mt-1">
          {sourceSummary.modules.map((m) => (
            <span
              key={m.moduleId}
              className="inline-flex items-center rounded-md bg-muted px-2 py-0.5 text-[10px] font-medium"
            >
              {m.moduleName}
            </span>
          ))}
        </div>
      )}
      {/* canonical binding CTA */}
      <div className="mt-2">
        <a
          href="#product-module-binding"
          className="inline-flex items-center gap-1 text-xs text-primary hover:underline"
        >
          为模板模块绑定仓库
          <ArrowDown className="h-3 w-3" />
        </a>
      </div>
    </div>
  )
}
