import { useSearch, useNavigate, Link } from '@tanstack/react-router'
import { useQuery } from '@tanstack/react-query'
import { useEffect } from 'react'
import { fetchModuleList } from '../data/module-registry-adapter'
import { ModuleListToolbar } from '../components/module-list-toolbar'
import { ModuleListContent } from '../components/module-list-content'
import { Button } from '@/components/ui/button'
import { Plus } from 'lucide-react'
import { useModuleListSearchStore } from '../stores/module-list-search-store'
import { BackToDashboardButton } from '@/features/dashboard/components/back-to-dashboard-button'
import { mergeCurrentDashboardSource } from '@/features/dashboard/lib/dashboard-source'

/**
 * ModuleListPage — Module Registry / List
 *
 * §8.4 状态模型：
 * - 查询条件冻结到路由搜索参数层：queryText / statusFilter
 * - 读取状态：pending / success / error
 * - 派生视图状态：initial-loading / ready / empty / error
 * - 从 ModuleCreatePage 或 ModuleDetailPage 返回时恢复列表上下文
 * - 错误反馈停留在内容区域，不跳转独立错误页
 */
export function ModuleListPage() {
  const search = useSearch({ from: '/modules/' })
  const navigate = useNavigate({ from: '/modules/' })
  const setLastSearch = useModuleListSearchStore((s) => s.setLastSearch)

  // §7.4 同步当前列表搜索上下文到 store，供创建页/详情页返回时恢复
  useEffect(() => {
    setLastSearch({
      queryText: search.queryText,
      statusFilter: search.statusFilter,
    })
  }, [search.queryText, search.statusFilter, setLastSearch])

  const { data, isLoading, isError, error, refetch } = useQuery({
    queryKey: ['module-list', search],
    queryFn: () => fetchModuleList(search),
  })

  const isFiltered = Boolean(search.queryText) || (search.statusFilter !== 'all')
  const isEmpty = !isLoading && !isError && (data?.length ?? 0) === 0
  const detailSearch: Record<string, unknown> = mergeCurrentDashboardSource(
    {
      fromList: true,
      queryText: search.queryText,
      statusFilter: search.statusFilter,
    },
    search,
  ) as unknown as Record<string, unknown>
  const createSearch: Record<string, unknown> = mergeCurrentDashboardSource({}, search) as unknown as Record<string, unknown>

  return (
    <div className="space-y-4">
      {/* phase05-13：从 Dashboard 进入时展示"返回 Dashboard"按钮 */}
      <BackToDashboardButton />

      {/* 页面标题与创建入口 */}
      <div className="flex items-center justify-between">
        <h1 className="text-2xl font-bold">Module Registry</h1>
        <Button asChild>
            <Link to="/modules/new" search={createSearch}>
            <Plus className="mr-2 h-4 w-4" />
            新建模块
          </Link>
        </Button>
      </div>

      {/* 工具栏：搜索参数 */}
      <ModuleListToolbar
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
        // §8.4 错误反馈停留在内容区域
        <div className="rounded-lg border border-destructive/50 bg-destructive/10 p-4">
          <p className="text-sm text-destructive">列表读取失败：{(error as Error).message}</p>
          <Button variant="outline" size="sm" className="mt-2" onClick={() => refetch()}>
            重试
          </Button>
        </div>
      ) : isEmpty && !isFiltered ? (
        // §7.1 空状态：主动作直接进入 Module Create
        <div className="rounded-lg border border-dashed p-8 text-center">
          <p className="text-muted-foreground mb-4">系统中尚无任何模块，先完成首个模块登记</p>
          <Button asChild>
              <Link to="/modules/new" search={createSearch}>
              <Plus className="mr-2 h-4 w-4" />
              完成首个模块登记
            </Link>
          </Button>
        </div>
      ) : isEmpty && isFiltered ? (
        <div className="rounded-lg border border-dashed p-8 text-center">
          <p className="text-muted-foreground">没有匹配筛选条件的模块</p>
        </div>
      ) : (
          <ModuleListContent items={data ?? []} isLoading={isLoading} detailSearch={detailSearch} />
      )}
    </div>
  )
}
