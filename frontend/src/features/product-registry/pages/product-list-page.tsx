import { useSearch, useNavigate, Link } from '@tanstack/react-router'
import { useProductListRead } from '../data/use-product-list-read'
import { ProductListToolbar } from '../components/product-list-toolbar'
import { ProductListContent } from '../components/product-list-content'
import { Button } from '@/components/ui/button'
import { Plus } from 'lucide-react'
import { BackToDashboardButton } from '@/features/dashboard/components/back-to-dashboard-button'
import { mergeCurrentDashboardSource } from '@/features/dashboard/lib/dashboard-source'
import type { ProductDetailSearch } from '../components/product-list-content'

/**
 * ProductListPage — Product Registry / List
 *
 * phase04-06 状态模型：
 * - 查询条件冻结到路由搜索参数层：queryText / statusFilter
 * - 读取状态：pending / success / error
 * - 派生视图状态：initial-loading / ready / empty / error
 * - 错误反馈停留在内容区域，不跳转独立错误页
 *
 * phase04-06 列表查询条件承接策略：
 * - 路由搜索参数是列表查询条件的唯一事实源
 * - 不使用 sessionStorage 等持久化层作为缺参回退源
 *
 * phase04-05 上下文入口：
 * - 从 Module Detail 带上下文进入时（fromModuleDetail + moduleId / moduleName），
 *   展示来源上下文提示，点击列表项时继续携带上下文到详情页
 *
 * 布局降级（phase04-05）：
 * - PC：高信息密度列表布局，工具栏与列表内容同屏可见
 * - 移动：单列列表或卡片重排
 */
export function ProductListPage() {
  const search = useSearch({ from: '/products/' })
  const navigate = useNavigate({ from: '/products/' })

  const { data, isLoading, isError, error, refetch } = useProductListRead(search)

  const isFiltered = Boolean(search.queryText) || (search.statusFilter !== 'all')
  const isEmpty = !isLoading && !isError && (data?.length ?? 0) === 0

  // phase04-06 来源上下文：从 Module Detail 带上下文进入
  const hasModuleContext = Boolean(search.fromModuleDetail && search.moduleId && search.moduleName)

  // phase04-05 / 06 上下文入口：点击列表项时继续携带上下文搜索参数到详情页
  // - 从 Module Detail 进入 → 继续携带 fromModuleDetail + moduleId / moduleName
  // - 普通列表进入 → 继续携带 fromList + 原 queryText / statusFilter
  const detailSearch: ProductDetailSearch = mergeCurrentDashboardSource(
    hasModuleContext
      ? { fromModuleDetail: true, moduleId: search.moduleId, moduleName: search.moduleName }
      : { fromList: true, queryText: search.queryText, statusFilter: search.statusFilter },
    search,
  ) as unknown as ProductDetailSearch

  // phase04-05 上下文入口：新建产品时也继续携带上下文搜索参数
  const createSearch = mergeCurrentDashboardSource(
    hasModuleContext
      ? { fromModuleDetail: true, moduleId: search.moduleId, moduleName: search.moduleName }
      : { fromList: true, queryText: search.queryText, statusFilter: search.statusFilter },
    search,
  ) as unknown as Record<string, unknown>

  return (
    <div className="space-y-4">
      {/* phase05-13：从 Dashboard 进入时展示"返回 Dashboard"按钮 */}
      <BackToDashboardButton />

      {/* 页面标题与创建入口 */}
      <div className="flex items-center justify-between">
        <h1 className="text-2xl font-bold">Product Registry</h1>
        <Button asChild>
          <Link
            to="/products/new"
            search={createSearch}
          >
            <Plus className="mr-2 h-4 w-4" />
            新建产品
          </Link>
        </Button>
      </div>

      {/* phase04-06 来源上下文展示 — 从 Module Detail 带上下文进入时 */}
      {hasModuleContext && (
        <div className="rounded-lg border bg-muted/50 p-3 text-sm">
          <span className="text-muted-foreground">来源模块：</span>
          <span className="font-medium">{search.moduleName}</span>
          <span className="ml-2 text-muted-foreground">（选择目标产品后进入详情完成绑定）</span>
        </div>
      )}

      {/* 工具栏：搜索参数 — phase04-06 路由搜索参数层 */}
      <ProductListToolbar
        queryText={search.queryText ?? ''}
        statusFilter={search.statusFilter ?? 'all'}
          onChange={(queryText, statusFilter) =>
            navigate({
              search: (prev) => ({
                ...prev,
                queryText: queryText || undefined,
                statusFilter,
              }),
            })
          }
      />

      {/* 内容区：根据读取结果派生视图状态 */}
      {isError ? (
        // phase04-06 错误反馈停留在内容区域
        <div className="rounded-lg border border-destructive/50 bg-destructive/10 p-4">
          <p className="text-sm text-destructive">列表读取失败：{(error as Error).message}</p>
          <Button variant="outline" size="sm" className="mt-2" onClick={() => refetch()}>
            重试
          </Button>
        </div>
      ) : isEmpty && !isFiltered ? (
        // phase04-06 空状态：主动作直接进入 Product Create
        <div className="rounded-lg border border-dashed p-8 text-center">
          <p className="text-muted-foreground mb-4">系统中尚无任何产品，先完成首个产品登记</p>
          <Button asChild>
            <Link to="/products/new" search={createSearch}>
              <Plus className="mr-2 h-4 w-4" />
              完成首个产品登记
            </Link>
          </Button>
        </div>
      ) : isEmpty && isFiltered ? (
        <div className="rounded-lg border border-dashed p-8 text-center">
          <p className="text-muted-foreground">没有匹配筛选条件的产品</p>
        </div>
      ) : (
        <ProductListContent items={data ?? []} isLoading={isLoading} detailSearch={detailSearch} />
      )}
    </div>
  )
}
