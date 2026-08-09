import { useParams, useSearch, useNavigate } from '@tanstack/react-router'
import { useQuery, useQueryClient } from '@tanstack/react-query'
import { fetchProductDetail } from '../data/product-registry-adapter'
import { ProductSummaryCard } from '../components/product-summary-card'
import { ProductModuleBindingPanel } from '../components/product-module-binding-panel'
import { ProductBoundRepositoryListSection } from '../components/product-bound-repository-list-section'
import { Button } from '@/components/ui/button'
import { ArrowLeft } from 'lucide-react'
import { Skeleton } from '@/components/ui/skeleton'
import { BackToDashboardButton } from '@/features/dashboard/components/back-to-dashboard-button'
import { mergeCurrentDashboardSource } from '@/features/dashboard/lib/dashboard-source'

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

  const { data, isLoading, isError, error } = useQuery({
    queryKey: ['product-detail', productId],
    queryFn: () => fetchProductDetail(productId),
    enabled: Boolean(productId),
  })

  // phase04-06 BindModuleToProduct 成功后重新读取详情结果（reread）
  const invalidateDetail = () => {
    queryClient.invalidateQueries({ queryKey: ['product-detail', productId] })
    queryClient.invalidateQueries({ queryKey: ['product-list'] })
    queryClient.invalidateQueries({ queryKey: ['product-module-candidates', productId] })
  }

  // phase04-06 主动返回路径 — 按真实来源决定，刷新后恢复来源标记
  const handleReturn = () => {
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

  if (isError) {
    return (
      <div className="space-y-4">
        {/* phase05-13：从 Dashboard 进入时同时展示"返回 Dashboard"与原生返回 */}
        <div className="flex items-center gap-2">
          <BackToDashboardButton />
          <Button variant="ghost" size="sm" onClick={handleReturn}>
            <ArrowLeft className="mr-2 h-4 w-4" />
            {fromModuleDetail ? '返回模块详情' : '返回列表'}
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
            {fromModuleDetail ? '返回模块详情' : '返回列表'}
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
          {fromModuleDetail ? '返回模块详情' : '返回列表'}
        </Button>
      </div>

      {/* phase04-06 来源上下文展示 — 从 Module Detail 带上下文进入时 */}
      {fromModuleDetail && search.moduleName && (
        <div className="rounded-lg border bg-muted/50 p-3 text-sm">
          <span className="text-muted-foreground">来源模块：</span>
          <span className="font-medium">{search.moduleName}</span>
        </div>
      )}

      {/* PC：分区式布局；移动端：垂直顺序重排 — phase04-05 */}
      <div className="grid gap-4 lg:grid-cols-3">
        {/* 摘要主区 — 占 1 列（PC）/ 全宽（移动） */}
        <div className="lg:col-span-1">
          <ProductSummaryCard product={data.product} />
        </div>

        {/* 绑定区 — 占 2 列（PC）/ 全宽（移动） */}
        <div className="space-y-4 lg:col-span-2">
          <ProductModuleBindingPanel
            productId={productId}
            boundModules={data.bound_modules}
            prefillModuleId={fromModuleDetail ? search.moduleId : undefined}
            onBindingSuccess={invalidateDetail}
          />
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
        </div>
      </div>
    </div>
  )
}
