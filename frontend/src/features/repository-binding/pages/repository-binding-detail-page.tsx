import { useState } from 'react'
import { useParams, useSearch, useNavigate } from '@tanstack/react-router'
import { useQueryClient } from '@tanstack/react-query'
import { useRepositoryDetailRead } from '../data/use-repository-detail-read'
import { RepositorySummaryCard } from '../components/repository-summary-card'
import { RepositoryProductBindingPanel } from '../components/repository-product-binding-panel'
import { RepositoryModuleMappingPanel } from '../components/repository-module-mapping-panel'
import { Button } from '@/components/ui/button'
import { ArrowLeft } from 'lucide-react'
import { Skeleton } from '@/components/ui/skeleton'
import {
  buildProductDetailSearchFromTransit,
  type RepositoryBindingSearch,
} from '../utils/product-source-transit'
import { BackToDashboardButton } from '@/features/dashboard/components/back-to-dashboard-button'
import { mergeCurrentDashboardSource } from '@/features/dashboard/lib/dashboard-source'
import {
  shouldReturnToOnboarding,
  buildOnboardingReturnSearch,
} from '@/features/onboarding/lib/onboarding-return'
import { RepositoryNextActionBar } from '../components/repository-next-action-bar'
import { RepositoryDecisionEntryPanel } from '../components/repository-decision-entry-panel'
import { DAILY_REVIEW_QUERY_KEY, WEEKLY_REVIEW_QUERY_KEY } from '@/features/review/data/review-query-options'
import { buildReviewReturnSearch, shouldReturnToReview } from '@/features/review/lib/review-source'
import { useModuleDecisionLinksByModuleIds } from '@/features/module-registry/data/use-module-decision-links-by-module-ids'
import { REPOSITORY_SEMANTIC_LABEL } from '@/features/project-context/data/shared-semantic-constants'
import { useProjectContextRead } from '@/features/project-context/data/use-project-context-read'
import { toEntryLocationViews } from '@/features/project-context/data/entry-location-view-model'
import { RULE_ENTRY_LABEL, PHASE_ENTRY_LABEL, BOUNDARY_ENTRY_LABEL } from '@/features/project-context/data/shared-semantic-constants'

/**
 * panelMode — phase04-06 互斥展开状态
 * 同一时刻只允许一个绑定面板处于打开态
 */
type PanelMode = 'closed' | 'product' | 'module'

/**
 * RepositoryBindingDetailPage — Repository Binding Detail / Workspace
 *
 * phase04-06 状态模型：
 * - 详情读取状态：pending / success / error
 * - 资源不存在时派生 not-found 视图状态
 * - 错误反馈停留在工作台内容区域，不跳转独立错误页
 *
 * phase04-06 来源上下文（由路由搜索参数派生，只允许四种之一）：
 * - fromList 存在 → 来自 Repository Binding / List，承接 queryText / statusFilter
 * - fromProductDetail 存在 → 来自 Product Detail，承接 productId / productName（用于预填 Product 绑定面板）
 * - fromModuleDetail 存在 → 来自 Module Detail，承接 moduleId / moduleName（用于预填 Module 映射面板）
 * - 无来源参数 → direct-entry
 * - 从 RepositoryCreatePage 成功创建后进入时，来源上下文继承自创建页
 *
 * phase04-06 两类绑定面板互斥展开：
 * - 同一时刻只允许一个绑定面板处于打开态
 * - BindRepositoryToProduct / MapModuleToRepository 成功后停留在当前页并重新读取详情结果
 * - 绑定失败时停留在面板上下文，保留当前已选候选目标
 *
 * phase04-06 主动返回路径：
 * - fromList → 回 Repository Binding / List + 原 queryText / statusFilter
 * - fromProductDetail → 回原 ProductDetailPage
 * - fromModuleDetail → 回原 ModuleDetailPage
 * - direct-entry → 回 Repository Binding / List 默认筛选参数
 * - 刷新后必须恢复来源标记
 *
 * 布局降级（phase04-05）：
 * - PC：分区式详情布局，摘要、已绑定产品、已映射模块与绑定工作台区可同时可见
 * - 移动：摘要、已绑定产品、已映射模块与绑定面板按垂直顺序重排
 */
export function RepositoryBindingDetailPage() {
  const { repositoryId } = useParams({ from: '/repositories/$repositoryId' })
  const search = useSearch({ from: '/repositories/$repositoryId' })
  const navigate = useNavigate()
  const queryClient = useQueryClient()

  // phase04-06 互斥展开状态
  const [panelMode, setPanelMode] = useState<PanelMode>('closed')
  const [copiedEntryRef, setCopiedEntryRef] = useState<string | null>(null)

  // phase04-06 来源上下文单值判定
  const fromList = search.fromList === true
  const fromProductDetail = search.fromProductDetail === true
  const fromModuleDetail = search.fromModuleDetail === true
  // phase06-15 §"detail 页来源优先级"：fromOnboarding 优先级高于其他来源
  const fromOnboarding = shouldReturnToOnboarding(search)
  const fromReview = shouldReturnToReview(search)

  const { data, isLoading, isError, error } = useRepositoryDetailRead(repositoryId)
  const mappedModules = Array.isArray(data?.mapped_modules) ? data.mapped_modules : []
  const relatedDecisionLinksQuery = useModuleDecisionLinksByModuleIds(
    mappedModules.map((module) => module.module_id),
  )

  // phase12-09：直接 repository-scoped 页面接入共享只读主线
  const projectContextQuery = useProjectContextRead(repositoryId)

  // phase04-06 绑定成功后重新读取详情结果（reread）
  // phase10-10：补齐 Dashboard / Review query 失效，确保返回后 reread 正确
  // phase12-09/10：补齐 project-context 共享只读失效，确保共享摘要与入口定位不滞后
  const invalidateDetail = () => {
    queryClient.invalidateQueries({ queryKey: ['repository-detail', repositoryId] })
    queryClient.invalidateQueries({ queryKey: ['project-context', repositoryId] })
    queryClient.invalidateQueries({ queryKey: ['repository-list'] })
    queryClient.invalidateQueries({ queryKey: ['repository-product-candidates', repositoryId] })
    queryClient.invalidateQueries({ queryKey: ['repository-module-candidates', repositoryId] })
    queryClient.invalidateQueries({ queryKey: ['dashboard-feedback-signals'] })
    queryClient.invalidateQueries({ queryKey: ['dashboard-overview'] })
    queryClient.invalidateQueries({ queryKey: DAILY_REVIEW_QUERY_KEY })
    queryClient.invalidateQueries({ queryKey: WEEKLY_REVIEW_QUERY_KEY })
  }

  const handleCopyEntryRef = async (entryRef: string) => {
    if (!entryRef || typeof navigator === 'undefined' || !navigator.clipboard) {
      return
    }
    try {
      await navigator.clipboard.writeText(entryRef)
      setCopiedEntryRef(entryRef)
    } catch {
      setCopiedEntryRef(null)
    }
  }

  const renderEntryLocationList = (entries: ReturnType<typeof toEntryLocationViews>) => (
    <div className="space-y-2">
      {entries.map((entry, index) => (
        <div
          key={`${entry.entryKind}-${entry.entryRef}-${index}`}
          className="rounded-md border bg-muted/20 px-3 py-2"
        >
          <div className="flex items-start justify-between gap-3">
            <div className="min-w-0 space-y-1">
              <div className="flex items-start gap-2 text-sm">
                <span className="text-muted-foreground shrink-0">[{entry.entryKind}]</span>
                <span className="font-medium break-words">{entry.title}</span>
              </div>
              <p className="text-sm text-muted-foreground break-words">{entry.summary}</p>
              <div className="flex flex-wrap items-center gap-2 text-xs text-muted-foreground">
                <span>入口引用</span>
                <code className="rounded bg-muted px-1.5 py-0.5 break-all">{entry.entryRef}</code>
              </div>
            </div>
            <Button
              type="button"
              variant="outline"
              size="sm"
              className="shrink-0"
              onClick={() => void handleCopyEntryRef(entry.entryRef)}
            >
              {copiedEntryRef === entry.entryRef ? '已复制' : '复制入口'}
            </Button>
          </div>
        </div>
      ))}
    </div>
  )

  // phase04-06 互斥展开：打开一个面板时关闭另一个
  const handleProductPanelOpenChange = (open: boolean) => {
    setPanelMode(open ? 'product' : 'closed')
  }
  const handleModulePanelOpenChange = (open: boolean) => {
    setPanelMode(open ? 'module' : 'closed')
  }

  // phase04-06 主动返回路径 — 按真实来源决定，刷新后恢复来源标记
  // phase06-15：fromOnboarding 优先级最高，先于 fromList / fromProductDetail / fromModuleDetail / direct-entry
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
        to: '/repositories',
          search: mergeCurrentDashboardSource(
            {
              queryText: search.queryText,
              statusFilter: search.statusFilter ?? 'all',
            },
            search,
          ),
      })
    } else if (fromProductDetail && search.productId) {
      // phase04-06 返回 Product Detail 时，必须恢复 Product Detail 的来源标记
      // 基于透传参数（product 前缀）恢复 fromList / fromModuleDetail + 相应参数
      // 不得退化为 direct-entry
      const productDetailSearch = buildProductDetailSearchFromTransit(search as RepositoryBindingSearch)
      navigate({
        to: '/products/$productId',
        params: { productId: search.productId },
          search: mergeCurrentDashboardSource(productDetailSearch, search),
      })
    } else if (fromModuleDetail && search.moduleId) {
      navigate({
        to: '/modules/$moduleId',
        params: { moduleId: search.moduleId },
          search: mergeCurrentDashboardSource({}, search) as Record<string, unknown>,
      })
    } else {
      // direct-entry → 回 Repository Binding / List 默认筛选参数
      navigate({
        to: '/repositories',
          search: mergeCurrentDashboardSource({ statusFilter: 'all' as const }, search),
      })
    }
  }

  // 返回按钮文案根据来源决定
  // phase06-15：fromOnboarding 优先展示"返回首轮录入"
  const returnLabel = fromOnboarding
    ? '返回首轮录入'
    : fromReview
      ? `返回 ${search.reviewKind === 'weekly' ? 'Weekly Review' : 'Daily Review'}`
    : fromProductDetail
      ? '返回产品详情'
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
              ? '仓库不存在'
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

      {/* phase04-06 来源上下文展示 */}
      {fromProductDetail && search.productName && (
        <div className="rounded-lg border bg-muted/50 p-3 text-sm">
          <span className="text-muted-foreground">来源产品：</span>
          <span className="font-medium">{search.productName}</span>
        </div>
      )}
      {fromModuleDetail && search.moduleName && (
        <div className="rounded-lg border bg-muted/50 p-3 text-sm">
          <span className="text-muted-foreground">来源模块：</span>
          <span className="font-medium">{search.moduleName}</span>
        </div>
      )}

      {/* phase12-08：页面级语义导语 */}
      <p className="text-sm text-muted-foreground">
        该 Repository 代表一个{REPOSITORY_SEMANTIC_LABEL}
      </p>

      {/* phase10-10：页面级下一步动作区 */}
      <RepositoryNextActionBar
        hasProductBinding={(data.bound_products?.length ?? 0) > 0}
        hasModuleMapping={mappedModules.length > 0}
        repositoryId={repositoryId}
        repositoryName={data.repository.name}
        decisionLinks={relatedDecisionLinksQuery.decisionLinks}
        decisionDetailSearch={decisionDetailSearch}
        onOpenProductBinding={() => setPanelMode('product')}
        onOpenModuleMapping={() => setPanelMode('module')}
      />

      {/* PC：分区式布局；移动端：垂直顺序重排 — phase04-05 */}
      <div className="grid gap-4 lg:grid-cols-3">
        {/* 摘要主区 — 占 1 列（PC）/ 全宽（移动） */}
        <div className="lg:col-span-1">
          <RepositorySummaryCard repository={data.repository} />
        </div>

        {/* 绑定工作台区 — 占 2 列（PC）/ 全宽（移动） */}
        <div className="space-y-4 lg:col-span-2">
          {/* phase04-06 互斥展开：product 面板与 module 面板同一时刻只允许一个打开 */}
          <div id="repository-product-binding">
            <RepositoryProductBindingPanel
              repositoryId={repositoryId}
              boundProducts={data.bound_products}
              open={panelMode === 'product'}
              onOpenChange={handleProductPanelOpenChange}
              prefillProductId={fromProductDetail ? search.productId : undefined}
              onBindingSuccess={invalidateDetail}
            />
          </div>
          <div id="repository-module-mapping">
            <RepositoryModuleMappingPanel
              repositoryId={repositoryId}
              mappedModules={mappedModules}
              open={panelMode === 'module'}
              onOpenChange={handleModulePanelOpenChange}
              prefillModuleId={fromModuleDetail ? search.moduleId : undefined}
              onBindingSuccess={invalidateDetail}
            />
          </div>
          {/* phase10-10：Repository Detail Decision 入口面板 */}
          <RepositoryDecisionEntryPanel
            decisionLinks={relatedDecisionLinksQuery.decisionLinks}
            decisionDetailSearch={decisionDetailSearch}
            isLoading={relatedDecisionLinksQuery.isLoading}
            isError={relatedDecisionLinksQuery.isError}
          />
        </div>
      </div>

      {/* phase12-09：共享只读上下文 — 规则 / phase / 边界入口 */}
      <div className="border-t pt-4 space-y-3">
        <div className="flex items-center justify-between gap-2">
          <div>
            <h3 className="text-sm font-medium">项目上下文</h3>
            <p className="text-xs text-muted-foreground">共享只读摘要、规则入口与阶段定位</p>
          </div>
          {projectContextQuery.isLoading && (
            <span className="text-xs text-muted-foreground">读取中...</span>
          )}
        </div>

        {projectContextQuery.isError ? (
          <div className="rounded-lg border border-destructive/50 bg-destructive/10 p-4">
            <p className="text-sm text-destructive">
              项目上下文读取失败：{projectContextQuery.error?.message ?? '未知错误'}
            </p>
            <Button variant="outline" size="sm" className="mt-2" onClick={() => void projectContextQuery.refetch()}>
              重试
            </Button>
          </div>
        ) : projectContextQuery.data ? (
          <>
            {projectContextQuery.data.rules.length > 0 && (
              <div>
                <h4 className="mb-2 text-xs font-medium text-muted-foreground">{RULE_ENTRY_LABEL}</h4>
                {renderEntryLocationList(toEntryLocationViews(projectContextQuery.data.rules))}
              </div>
            )}
            {projectContextQuery.data.phases.length > 0 && (
              <div>
                <h4 className="mb-2 text-xs font-medium text-muted-foreground">{PHASE_ENTRY_LABEL}</h4>
                {renderEntryLocationList(toEntryLocationViews(projectContextQuery.data.phases))}
              </div>
            )}
            {projectContextQuery.data.boundaries.length > 0 && (
              <div>
                <h4 className="mb-2 text-xs font-medium text-muted-foreground">{BOUNDARY_ENTRY_LABEL}</h4>
                <div className="space-y-1">
                  {projectContextQuery.data.boundaries.map((entry, i) => (
                    <div key={i} className="text-sm">
                      <span className="font-medium">{entry.label}</span>
                      <span className="text-muted-foreground"> — {entry.summary}</span>
                    </div>
                  ))}
                </div>
              </div>
            )}
          </>
        ) : null}
      </div>
    </div>
  )
}
