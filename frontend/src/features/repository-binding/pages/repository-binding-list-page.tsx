import { useSearch, useNavigate, Link } from '@tanstack/react-router'
import { useQuery } from '@tanstack/react-query'
import { fetchRepositoryList } from '../data/repository-binding-adapter'
import { RepositoryBindingListToolbar } from '../components/repository-binding-list-toolbar'
import { RepositoryBindingListContent } from '../components/repository-binding-list-content'
import { Button } from '@/components/ui/button'
import { Plus } from 'lucide-react'
import {
  extractProductSourceTransit,
  type RepositoryBindingSearch,
} from '../utils/product-source-transit'

/**
 * RepositoryBindingListPage — Repository Binding / List
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
 * phase04-06 来源上下文保留规则：
 * - 用户修改筛选条件时，必须保留来源上下文参数（fromProductDetail / fromModuleDetail 等）
 * - 不得因修改筛选条件而丢失来源标记导致退化为 direct-entry
 *
 * 布局降级（phase04-05）：
 * - PC：高信息密度列表布局，工具栏与列表内容同屏可见
 * - 移动：单列列表或卡片重排
 *
 * phase04-06 来源上下文入口：
 * - 从 Product Detail 带上下文进入时（fromProductDetail + productId / productName），
 *   展示来源上下文提示并提供创建仓库入口
 * - 从 Module Detail 带上下文进入时（fromModuleDetail + moduleId / moduleName），
 *   展示来源上下文提示并提供创建仓库入口
 */
export function RepositoryBindingListPage() {
  const search = useSearch({ from: '/repositories/' })
  const navigate = useNavigate({ from: '/repositories/' })

  const { data, isLoading, isError, error, refetch } = useQuery({
    queryKey: ['repository-list', search],
    queryFn: () => fetchRepositoryList(search),
  })

  const isFiltered = Boolean(search.queryText) || (search.statusFilter !== 'all')
  const isEmpty = !isLoading && !isError && (data?.length ?? 0) === 0

  // phase04-06 来源上下文：从 Product Detail 带上下文进入
  const hasProductContext = Boolean(search.fromProductDetail && search.productId && search.productName)
  // phase04-06 来源上下文：从 Module Detail 带上下文进入
  const hasModuleContext = Boolean(search.fromModuleDetail && search.moduleId && search.moduleName)

  // phase04-06 Product Detail 来源上下文透传参数（用于继续传递到详情页/创建页）
  const productTransit = extractProductSourceTransit(search as RepositoryBindingSearch)

  // phase04-05 / 06 上下文入口：点击列表项时继续携带上下文搜索参数到详情页
  // - 从 Product Detail 进入 → 继续携带 fromProductDetail + productId / productName + Product Detail 来源透传
  // - 从 Module Detail 进入 → 继续携带 fromModuleDetail + moduleId / moduleName
  // - 普通列表进入 → 继续携带 fromList + 原 queryText / statusFilter
  const detailSearch: Record<string, unknown> = hasProductContext
    ? { fromProductDetail: true, productId: search.productId, productName: search.productName, ...productTransit }
    : hasModuleContext
      ? { fromModuleDetail: true, moduleId: search.moduleId, moduleName: search.moduleName }
      : { fromList: true, queryText: search.queryText, statusFilter: search.statusFilter }

  // phase04-05 上下文入口：新建仓库时也继续携带上下文搜索参数
  const createSearch: Record<string, unknown> = hasProductContext
    ? { fromProductDetail: true, productId: search.productId, productName: search.productName, ...productTransit }
    : hasModuleContext
      ? { fromModuleDetail: true, moduleId: search.moduleId, moduleName: search.moduleName }
      : { fromList: true, queryText: search.queryText, statusFilter: search.statusFilter }

  // phase04-06 工具栏 onChange：修改筛选条件时必须保留来源上下文参数
  // 不得因修改筛选条件而丢失 fromProductDetail / fromModuleDetail 等来源标记
  const handleFilterChange = (queryText: string, statusFilter: string) => {
    const newSearch: Record<string, unknown> = {
      queryText: queryText || undefined,
      statusFilter,
    }
    // 保留来源上下文参数
    if (hasProductContext) {
      newSearch.fromProductDetail = true
      newSearch.productId = search.productId
      newSearch.productName = search.productName
      // 继续携带 Product Detail 来源透传参数
      Object.assign(newSearch, productTransit)
    } else if (hasModuleContext) {
      newSearch.fromModuleDetail = true
      newSearch.moduleId = search.moduleId
      newSearch.moduleName = search.moduleName
    }
    navigate({ search: newSearch })
  }

  return (
    <div className="space-y-4">
      {/* 页面标题与创建入口 */}
      <div className="flex items-center justify-between">
        <h1 className="text-2xl font-bold">Repository Binding</h1>
        <Button asChild>
          <Link
            to="/repositories/new"
            search={createSearch}
          >
            <Plus className="mr-2 h-4 w-4" />
            新建仓库
          </Link>
        </Button>
      </div>

      {/* phase04-06 来源上下文展示 — 从 Product Detail / Module Detail 带上下文进入时 */}
      {hasProductContext && (
        <div className="rounded-lg border bg-muted/50 p-3 text-sm">
          <span className="text-muted-foreground">来源产品：</span>
          <span className="font-medium">{search.productName}</span>
          <span className="ml-2 text-muted-foreground">（选择目标仓库后进入详情完成绑定）</span>
        </div>
      )}
      {hasModuleContext && (
        <div className="rounded-lg border bg-muted/50 p-3 text-sm">
          <span className="text-muted-foreground">来源模块：</span>
          <span className="font-medium">{search.moduleName}</span>
          <span className="ml-2 text-muted-foreground">（选择目标仓库后进入详情完成映射）</span>
        </div>
      )}

      {/* 工具栏：搜索参数 — phase04-06 路由搜索参数层 */}
      <RepositoryBindingListToolbar
        queryText={search.queryText ?? ''}
        statusFilter={search.statusFilter ?? 'all'}
        onChange={handleFilterChange}
      />

      {/* 内容区：根据读取结果派生视图状态 */}
      {isError ? (
        <div className="rounded-lg border border-destructive/50 bg-destructive/10 p-4">
          <p className="text-sm text-destructive">列表读取失败：{(error as Error).message}</p>
          <Button variant="outline" size="sm" className="mt-2" onClick={() => refetch()}>
            重试
          </Button>
        </div>
      ) : isEmpty && !isFiltered ? (
        // phase04-06 空状态：主动作直接进入 Repository Create
        <div className="rounded-lg border border-dashed p-8 text-center">
          <p className="text-muted-foreground mb-4">系统中尚无任何仓库，先完成首个仓库登记</p>
          <Button asChild>
            <Link to="/repositories/new" search={createSearch}>
              <Plus className="mr-2 h-4 w-4" />
              完成首个仓库登记
            </Link>
          </Button>
        </div>
      ) : isEmpty && isFiltered ? (
        <div className="rounded-lg border border-dashed p-8 text-center">
          <p className="text-muted-foreground">没有匹配筛选条件的仓库</p>
        </div>
      ) : (
        <RepositoryBindingListContent items={data ?? []} isLoading={isLoading} detailSearch={detailSearch} />
      )}
    </div>
  )
}
