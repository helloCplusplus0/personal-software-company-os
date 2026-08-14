/**
 * FeedbackSignalCard — 反馈信号单值紧凑行
 *
 * phase05-13 体验修复：
 * - 从完整 Card（CardHeader + CardContent + 内嵌 Button）降级为单行紧凑行
 * - 行内：[优先级 badge] [标题 truncate] [目标 label muted] [箭头]
 * - summary 作为 hover title（title 属性），不在行内展开，保持 dashboard 信息密度
 * - 列表容器通过 divide-y 提供行间分隔，行本身不渲染边框
 *
 * phase05-13 §"FeedbackSignalCard 必须按信号类型跳转"
 * phase05-10 §8.1 反馈信号卡片跳转矩阵：
 *   - target_type=decision_detail → /decisions/$decisionId
 *   - target_type=decision_list   → /decisions
 *   - target_type=product_detail  → /products/$productId
 *
 * 携带来源参数：fromDashboard=true / dashboardReturnTo=/dashboard
 * dashboardSection 由调用方传入（current-focus 或 asset-feedback）
 */
import { useNavigate } from '@tanstack/react-router'
import { ArrowRight } from 'lucide-react'
import type { FeedbackSignal, FeedbackSignalPriority } from '../types'
import { buildDashboardSourceParams } from '../lib/dashboard-source'

interface FeedbackSignalCardProps {
  signal: FeedbackSignal
  // 来源区块：current-focus 或 asset-feedback
  // 决定跳转时携带的 dashboardSection 值
  section: 'current-focus' | 'asset-feedback'
  // 可选：由调用方覆盖 search 参数，用于非 Dashboard 宿主保留自己的返回链。
  getSearch?: (signal: FeedbackSignal) => Record<string, unknown>
}

/**
 * 将优先级映射为短标签（P1/P2/P3/P4）与语义色调。
 * - P1 / P2：destructive（需立即处理）
 * - P3 / P4：amber（需关注）
 */
function priorityBadge(priority: FeedbackSignalPriority): {
  label: string
  className: string
} {
  switch (priority) {
    case 'p1_pending_decision':
      return {
        label: 'P1',
        className: 'bg-destructive/10 text-destructive',
      }
    case 'p2_product_missing_both_bindings':
      return {
        label: 'P2',
        className: 'bg-destructive/10 text-destructive',
      }
    case 'p3_product_missing_repository_binding':
      return {
        label: 'P3',
        className: 'bg-amber-500/10 text-amber-600 dark:text-amber-400',
      }
    case 'p4_product_missing_module_binding':
      return {
        label: 'P4',
        className: 'bg-amber-500/10 text-amber-600 dark:text-amber-400',
      }
  }
}

/**
 * FeedbackSignalCard — 渲染单条反馈信号紧凑行。
 *
 * 整行可点击跳转，按 target_type 决定跳转目标。
 */
export function FeedbackSignalCard({ signal, section, getSearch }: FeedbackSignalCardProps) {
  const navigate = useNavigate()
  const badge = priorityBadge(signal.priority)

  const handleClick = () => {
    const sourceParams = getSearch?.(signal) ?? buildDashboardSourceParams(section)

    // 按 target_type 决定跳转目标与 params
    switch (signal.target_type) {
      case 'decision_detail':
        navigate({
          to: '/decisions/$decisionId',
          params: { decisionId: signal.target_id },
          search: sourceParams,
        })
        break
      case 'decision_list':
        navigate({
          to: '/decisions',
          search: sourceParams,
        })
        break
      case 'product_detail':
        navigate({
          to: '/products/$productId',
          params: { productId: signal.target_id },
          search: sourceParams,
        })
        break
      // module_detail / repository_detail 不在反馈信号跳转矩阵中，
      // 但保留兜底以避免后端异常数据导致前端崩溃
      case 'module_detail':
        navigate({
          to: '/modules/$moduleId',
          params: { moduleId: signal.target_id },
          search: sourceParams,
        })
        break
      case 'repository_detail':
        navigate({
          to: '/repositories/$repositoryId',
          params: { repositoryId: signal.target_id },
          search: sourceParams,
        })
        break
    }
  }

  // title 属性承接 summary，hover 可见，不在行内展开以保持密度
  return (
    <button
      type="button"
      onClick={handleClick}
      title={signal.summary}
      className="flex w-full items-center gap-2 px-3 py-2 text-left transition-colors hover:bg-muted/40"
    >
      {/* 优先级 badge — 紧凑 */}
      <span
        className={[
          'inline-flex h-5 min-w-[28px] shrink-0 items-center justify-center rounded px-1 text-[10px] font-semibold',
          badge.className,
        ].join(' ')}
      >
        {badge.label}
      </span>
      {/* 标题 — 主信息，允许占满剩余空间 */}
      <span className="min-w-0 flex-1 truncate text-sm font-medium">
        {signal.title}
      </span>
      {/* 目标 label — 次要信息，sm+ 屏幕展示 */}
      <span className="hidden max-w-[140px] shrink-0 truncate text-xs text-muted-foreground sm:inline">
        {signal.target_label}
      </span>
      <ArrowRight className="h-4 w-4 shrink-0 text-muted-foreground" />
    </button>
  )
}
