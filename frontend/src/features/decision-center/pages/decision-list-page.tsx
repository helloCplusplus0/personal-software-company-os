/**
 * DecisionListPage — Decision Center / List
 *
 * §9.1 状态模型：
 * - 查询条件冻结到路由搜索参数层：queryText / statusFilter
 * - 读取状态：pending / success / error
 * - 派生视图状态：initial-loading / ready / empty / error
 * - 从 DecisionCreatePage 或 DecisionDetailPage 返回时恢复列表上下文
 * - 错误反馈停留在内容区域，不跳转独立错误页
 *
 * 布局降级（phase03-05 §"布局降级策略"）：
 * - PC：高信息密度表格布局
 * - 移动浏览器：单列卡片重排
 */
import { useSearch, useNavigate, Link } from '@tanstack/react-router'
import { useEffect } from 'react'
import { useDecisionListRead } from '../data/use-decision-list-read'
import { DecisionListToolbar } from '../components/decision-list-toolbar'
import { DecisionListContent } from '../components/decision-list-content'
import { Button } from '@/components/ui/button'
import { Plus } from 'lucide-react'
import { useDecisionListSearchStore } from '../stores/decision-list-search-store'
import { BackToDashboardButton } from '@/features/dashboard/components/back-to-dashboard-button'
import { mergeCurrentDashboardSource } from '@/features/dashboard/lib/dashboard-source'

export function DecisionListPage() {
  const search = useSearch({ from: '/decisions/' })
  const navigate = useNavigate({ from: '/decisions/' })
  const setLastSearch = useDecisionListSearchStore((s) => s.setLastSearch)

  // §9.1 同步当前列表搜索上下文到 store，供创建页/详情页返回时恢复
  useEffect(() => {
    setLastSearch({
      queryText: search.queryText,
      statusFilter: search.statusFilter,
    })
  }, [search.queryText, search.statusFilter, setLastSearch])

  const { data, isLoading, isError, error, refetch } = useDecisionListRead(search)

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

  return (
    <div className="space-y-4">
      {/* phase05-13：从 Dashboard 进入时展示"返回 Dashboard"按钮 */}
      <BackToDashboardButton />

      {/* 页面标题与创建入口 — 对齐 Dashboard 基线：标题 text-xl、移动端纵向堆叠 */}
      <div className="flex flex-col gap-2 sm:flex-row sm:items-center sm:justify-between">
        <h1 className="text-xl font-bold">Decision Center</h1>
        <Button asChild>
            <Link
              to="/decisions/new"
              search={mergeCurrentDashboardSource({ fromList: true }, search)}
            >
            <Plus className="mr-2 h-4 w-4" />
            记录决策
          </Link>
        </Button>
      </div>

      {/* 工具栏：搜索参数 */}
      <DecisionListToolbar
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
        // §9.1 错误反馈停留在内容区域 — 对齐 Dashboard 紧凑错误区基线
        <div className="rounded-lg border border-destructive/50 bg-destructive/10 p-3">
          <p className="text-xs text-destructive">列表读取失败：{(error as Error).message}</p>
          <Button variant="outline" size="sm" className="mt-2 h-7" onClick={() => refetch()}>
            重试
          </Button>
        </div>
      ) : isEmpty && !isFiltered ? (
        // 空状态：主动作直接进入 Decision Create — 对齐 Dashboard 紧凑空态基线
        <div className="rounded-lg border border-dashed p-4 text-center">
          <p className="text-xs text-muted-foreground mb-3">暂无规则与决策记录</p>
          <Button asChild size="sm">
              <Link
                to="/decisions/new"
                search={mergeCurrentDashboardSource({ fromList: true }, search)}
              >
              <Plus className="mr-2 h-4 w-4" />
              记录首条决策
            </Link>
          </Button>
        </div>
      ) : isEmpty && isFiltered ? (
        <div className="rounded-lg border border-dashed p-4 text-center">
          <p className="text-xs text-muted-foreground">没有匹配筛选条件的决策</p>
        </div>
      ) : (
          <DecisionListContent items={data ?? []} isLoading={isLoading} detailSearch={detailSearch} />
      )}
    </div>
  )
}
