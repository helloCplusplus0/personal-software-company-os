import { useSearch, useNavigate, Link } from '@tanstack/react-router'
import { useEffect, useMemo } from 'react'
import { useModuleListRead } from '../data/use-module-list-read'
import { ModuleListToolbar } from '../components/module-list-toolbar'
import { ModuleListContent } from '../components/module-list-content'
import { Button } from '@/components/ui/button'
import { Plus, ArrowLeft } from 'lucide-react'
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

  const { data, isLoading, isError, error, refetch } = useModuleListRead(search)

  const isFiltered = Boolean(search.queryText) || (search.statusFilter !== 'all')
  const isEmpty = !isLoading && !isError && (data?.length ?? 0) === 0

  // phase09-10 模板复用提示返回链
  const fromTemplateReturn = (search as any).returnTo === 'weekly-review' || (search as any).returnTo === 'product-create'
  const templateReturnLabel = (search as any).returnTo === 'weekly-review' ? '返回 Weekly Review' : '返回创建产品'

  const handleTemplateReturn = useMemo(() => {
    if (!fromTemplateReturn) return undefined
    return () => {
      const rs = (search as any).returnTo as string
      const returnSearch: Record<string, unknown> = {}
      if ((search as any).returnCandidateId) {
        returnSearch.returnCandidateId = (search as any).returnCandidateId
      }
      if ((search as any).fromTemplateReuse) {
        returnSearch.fromTemplateReuse = true
        returnSearch.templateCandidateId = (search as any).templateCandidateId
        returnSearch.templateSource = (search as any).templateSource
      }
      if ((search as any).fromDashboard) {
        returnSearch.fromDashboard = true
        returnSearch.dashboardSection = (search as any).dashboardSection
        returnSearch.dashboardReturnTo = (search as any).dashboardReturnTo
      }
      if (rs === 'weekly-review') {
        navigate({ to: '/reviews/weekly', search: returnSearch })
      } else if (rs === 'product-create') {
        navigate({ to: '/products/new', search: returnSearch })
      }
    }
  }, [fromTemplateReturn, navigate, search])

  const detailSearch: Record<string, unknown> = mergeCurrentDashboardSource(
    {
      fromList: true,
      queryText: search.queryText,
      statusFilter: search.statusFilter,
    },
    search,
  ) as unknown as Record<string, unknown>

  // phase09-10 模板返回链参数透传到 detail 页
  if (fromTemplateReturn) {
    detailSearch.returnTo = (search as any).returnTo
    if ((search as any).returnCandidateId) detailSearch.returnCandidateId = (search as any).returnCandidateId
    if ((search as any).fromTemplateReuse) {
      detailSearch.fromTemplateReuse = true
      detailSearch.templateCandidateId = (search as any).templateCandidateId
      detailSearch.templateSource = (search as any).templateSource
    }
    if ((search as any).fromDashboard) {
      detailSearch.fromDashboard = true
      detailSearch.dashboardSection = (search as any).dashboardSection
      detailSearch.dashboardReturnTo = (search as any).dashboardReturnTo
    }
  }

  const createSearch: Record<string, unknown> = mergeCurrentDashboardSource({}, search) as unknown as Record<string, unknown>

  return (
    <div className="space-y-4">
      {/* phase05-13：从 Dashboard 进入时展示"返回 Dashboard"按钮 */}
      {/* phase09-10：模板复用提示返回链按钮 */}
      <div className="flex items-center gap-2">
        <BackToDashboardButton />
        {fromTemplateReturn && handleTemplateReturn && (
          <Button variant="ghost" size="sm" onClick={handleTemplateReturn}>
            <ArrowLeft className="mr-2 h-4 w-4" />
            {templateReturnLabel}
          </Button>
        )}
      </div>

      {/* 页面标题与创建入口 — 对齐 Dashboard 基线：标题 text-xl、移动端纵向堆叠 */}
      <div className="flex flex-col gap-2 sm:flex-row sm:items-center sm:justify-between">
        <h1 className="text-xl font-bold">Module Registry</h1>
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
        // §8.4 错误反馈停留在内容区域 — 对齐 Dashboard 紧凑错误区基线
        <div className="rounded-lg border border-destructive/50 bg-destructive/10 p-3">
          <p className="text-xs text-destructive">列表读取失败：{(error as Error).message}</p>
          <Button variant="outline" size="sm" className="mt-2 h-7" onClick={() => refetch()}>
            重试
          </Button>
        </div>
      ) : isEmpty && !isFiltered ? (
        // §7.1 空状态：主动作直接进入 Module Create — 对齐 Dashboard 紧凑空态基线
        <div className="rounded-lg border border-dashed p-4 text-center">
          <p className="text-xs text-muted-foreground mb-3">暂无已登记的可复用能力资产</p>
          <Button asChild size="sm">
              <Link to="/modules/new" search={createSearch}>
              <Plus className="mr-2 h-4 w-4" />
              完成首个模块登记
            </Link>
          </Button>
        </div>
      ) : isEmpty && isFiltered ? (
        <div className="rounded-lg border border-dashed p-4 text-center">
          <p className="text-xs text-muted-foreground">没有匹配筛选条件的模块</p>
        </div>
      ) : (
          <ModuleListContent items={data ?? []} isLoading={isLoading} detailSearch={detailSearch} />
      )}
    </div>
  )
}
