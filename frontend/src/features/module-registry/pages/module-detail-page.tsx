import { useParams, Link, useNavigate, useSearch } from '@tanstack/react-router'
import { useQuery } from '@tanstack/react-query'
import { fetchModuleDetail } from '../data/module-registry-adapter'
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
  const returnSearch = mergeCurrentDashboardSource(
    detailSearch.fromList
      ? {
          queryText: detailSearch.queryText,
          statusFilter: detailSearch.statusFilter ?? 'all',
        }
      : lastSearch,
    detailSearch,
  ) as unknown as Record<string, unknown>

  const { data, isLoading, isError, error } = useQuery({
    queryKey: ['module-detail', moduleId],
    queryFn: () => fetchModuleDetail(moduleId),
    enabled: Boolean(moduleId),
  })

  if (isError) {
    return (
      <div className="space-y-4">
        {/* phase05-13：从 Dashboard 进入时同时展示"返回 Dashboard"与"返回列表" */}
        <div className="flex items-center gap-2">
          <BackToDashboardButton />
            <Button variant="ghost" size="sm" onClick={() => navigate({ to: '/modules', search: returnSearch })}>
            <ArrowLeft className="mr-2 h-4 w-4" />
            返回列表
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
          <Button variant="ghost" size="sm" asChild>
              <Link to="/modules" search={returnSearch}>
              <ArrowLeft className="mr-2 h-4 w-4" />
              返回列表
            </Link>
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
      <div className="flex items-center gap-2">
        <BackToDashboardButton />
        <Button variant="ghost" size="sm" asChild>
            <Link to="/modules" search={returnSearch}>
            <ArrowLeft className="mr-2 h-4 w-4" />
            返回列表
          </Link>
        </Button>
      </div>

      {/* PC：分区式布局；移动端：垂直顺序重排 */}
      <div className="grid gap-4 lg:grid-cols-3">
        {/* 摘要主区 — 占 1 列（PC）/ 全宽（移动） */}
        <div className="lg:col-span-1">
          <ModuleSummaryCard module={data.module} />
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
          />
        </div>
      </div>
    </div>
  )
}
