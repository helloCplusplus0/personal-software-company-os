/**
 * RecentActivityItemCard — 最近活动单值紧凑行
 *
 * phase05-13 体验修复：
 * - 从完整 Card 降级为单行紧凑行
 * - 行内：[类型 badge] [目标 label truncate] [时间 muted] [箭头]
 * - 列表容器通过 divide-y 提供行间分隔，行本身不渲染边框
 *
 * phase05-13 §"RecentActivityItemCard 必须按活动类型跳转"
 * phase05-10 §8.1 活动项跳转矩阵：
 *   - module_detail              → /modules/$moduleId
 *   - product_detail             → /products/$productId
 *   - repository_detail          → /repositories/$repositoryId
 *   - decision_detail            → /decisions/$decisionId
 *
 * 携带来源参数：fromDashboard=true / dashboardSection=recent-activity / dashboardReturnTo=/dashboard
 *
 * 注意：activity_type 与 target_type 是两个独立字段：
 *   - activity_type 描述"发生了什么"（module / release / product / ...）
 *   - target_type 描述"跳转到哪里"（已由后端归一化为 canonical owner target_type）
 *   - release 活动后端统一回落到 module_detail target_type
 *   - product_module_binding 后端统一落到 product_detail target_type
 *   - product_repository_binding / module_repository_binding 后端统一落到 repository_detail target_type
 */
import { useNavigate } from '@tanstack/react-router'
import { ArrowRight } from 'lucide-react'
import type { RecentActivityItem, DashboardSection } from '../types'
import { buildDashboardSourceParams } from '../lib/dashboard-source'

interface RecentActivityItemCardProps {
  activity: RecentActivityItem
  /**
   * 来源区块：决定跳转时携带的 dashboardSection 值。
   * - Dashboard RecentActivitySection 调用方保持默认 'recent-activity'
   * - Weekly Review 复用时也传入 'recent-activity'，保持返回 dashboard 后定位一致
   */
  section?: DashboardSection
}

/**
 * 将 activity_type 映射为中文短标签，用于行内 badge 展示。
 */
function activityTypeLabel(activityType: RecentActivityItem['activity_type']): string {
  const labelMap: Record<RecentActivityItem['activity_type'], string> = {
    module: '模块',
    release: '版本',
    product: '产品',
    repository: '仓库',
    decision: '决策',
    product_module_binding: '产品-模块',
    product_repository_binding: '产品-仓库',
    module_repository_binding: '模块-仓库',
  }
  return labelMap[activityType] ?? activityType
}

/**
 * 将 ISO 时间字符串格式化为简短展示（月-日 时:分），用于行内时间列。
 * 去掉年份以节省行内空间，dashboard 活动流默认看近期。
 */
function formatActivityAtShort(isoString: string): string {
  try {
    const date = new Date(isoString)
    return date.toLocaleString('zh-CN', {
      month: '2-digit',
      day: '2-digit',
      hour: '2-digit',
      minute: '2-digit',
    })
  } catch {
    return isoString
  }
}

/**
 * RecentActivityItemCard — 渲染单条最近活动项紧凑行。
 *
 * 整行可点击跳转，按 target_type 决定跳转目标。
 */
export function RecentActivityItemCard({ activity, section = 'recent-activity' }: RecentActivityItemCardProps) {
  const navigate = useNavigate()

  const handleClick = () => {
    // 活动项跳转统一携带传入的 section 作为 dashboardSection（默认 'recent-activity'）
    const sourceParams = buildDashboardSourceParams(section)

    switch (activity.target_type) {
      case 'module_detail':
        navigate({
          to: '/modules/$moduleId',
          params: { moduleId: activity.target_id },
          search: sourceParams,
        })
        break
      case 'product_detail':
        navigate({
          to: '/products/$productId',
          params: { productId: activity.target_id },
          search: sourceParams,
        })
        break
      case 'repository_detail':
        navigate({
          to: '/repositories/$repositoryId',
          params: { repositoryId: activity.target_id },
          search: sourceParams,
        })
        break
      case 'decision_detail':
        navigate({
          to: '/decisions/$decisionId',
          params: { decisionId: activity.target_id },
          search: sourceParams,
        })
        break
      // decision_list 不在活动项跳转矩阵中，保留兜底
      case 'decision_list':
        navigate({
          to: '/decisions',
          search: sourceParams,
        })
        break
    }
  }

  return (
    <button
      type="button"
      onClick={handleClick}
      className="flex w-full items-center gap-2 px-3 py-2 text-left transition-colors hover:bg-muted/40"
    >
      {/* 类型 badge — 紧凑 */}
      <span className="inline-flex shrink-0 items-center rounded bg-muted px-1.5 py-0.5 text-[10px] text-muted-foreground">
        {activityTypeLabel(activity.activity_type)}
      </span>
      {/* 目标 label — 主信息 */}
      <span className="min-w-0 flex-1 truncate text-sm">
        {activity.target_label}
      </span>
      {/* 时间 — 次要信息，sm+ 屏幕展示 */}
      <span className="hidden shrink-0 text-xs text-muted-foreground sm:inline">
        {formatActivityAtShort(activity.activity_at)}
      </span>
      <ArrowRight className="h-4 w-4 shrink-0 text-muted-foreground" />
    </button>
  )
}
