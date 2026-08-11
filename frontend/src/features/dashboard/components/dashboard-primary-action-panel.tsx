/**
 * DashboardPrimaryActionPanel — 主 CTA 面板（phase08-08 收敛为 dual review launcher）。
 *
 * phase08-08 §"Dashboard 标题行动区必须在本阶段完成 dual review launcher 收敛"：
 *   - 本面板从旧的单 CTA 命中器正式收敛为 dual review launcher
 *   - 稳定渲染 Daily Review 与 Weekly Review 两个显式入口
 *   - 只负责组装 /reviews/daily 或 /reviews/weekly 的 route search 并透传
 *     buildDashboardSourceParams('empty-state')
 *   - 不得继续保留 computePrimaryCta() 驱动的旧单主 CTA 作为 formal review 入口 owner
 *
 * phase08-06 §"DashboardPrimaryActionPanel 的前端职责解释"：
 *   - 本面板被解释为 review route launcher caller，而不是 review 读层或写层的编排 owner
 *   - 只能负责组装 route search 并导航到对应 review route
 *   - 不得直接读取 pending decisions、reuse snapshot 或持有 Review action application owner
 */
import { useNavigate } from '@tanstack/react-router'
import { Button } from '@/components/ui/button'
import { Calendar, Clock } from 'lucide-react'
import { buildDashboardSourceParams } from '../lib/dashboard-source'

/**
 * DashboardPrimaryActionPanel — dual review launcher。
 *
 * 始终渲染 Daily Review 与 Weekly Review 两个显式入口按钮，
 * 不再依赖 overview / feedback 的命中状态决定是否显示。
 *
 * phase08-05 §"Dashboard 标题行动区的信息密度约束"：
 *   - 不得因为新增 daily / weekly review 双入口而把首屏改造成大块工作台
 *   - PC 下允许双按钮并排或主次按钮组合
 *   - 移动浏览器下必须降级为紧凑双按钮，不得压缩成难以点击的微型入口
 *
 * 视觉风格对齐 dashboard 紧凑化基线：
 *   - 采用 segmented control 风格（共享边框 + ghost 按钮 + divide-x 分隔）
 *   - 按钮高度 h-8、字号 text-xs，与 dashboard stat bar 视觉重量协调
 *   - 不再使用 default 紫主色 + outline 灰色的双独立大按钮组合
 *   - 移动端可水平一行展示，不换行、不撑破 header
 */
export function DashboardPrimaryActionPanel() {
  const navigate = useNavigate()

  const handleDailyReview = () => {
    const sourceParams = buildDashboardSourceParams('empty-state')
    navigate({
      to: '/reviews/daily',
      search: sourceParams,
    })
  }

  const handleWeeklyReview = () => {
    const sourceParams = buildDashboardSourceParams('empty-state')
    navigate({
      to: '/reviews/weekly',
      search: sourceParams,
    })
  }

  return (
    <div className="inline-flex items-center rounded-md border bg-muted/30 divide-x divide-border overflow-hidden">
      <Button
        onClick={handleDailyReview}
        size="sm"
        variant="ghost"
        className="h-8 px-2.5 sm:px-3 text-xs rounded-none gap-1 sm:gap-1.5 hover:bg-background hover:shadow-sm"
      >
        <Clock className="h-3.5 w-3.5 text-amber-500" />
        {/* 移动端文案简化为 "Daily" 节省宽度；桌面端展示完整 "Daily Review" 表达 review 会话语义 */}
        <span>Daily</span>
        <span className="hidden sm:inline">&nbsp;Review</span>
      </Button>
      <Button
        onClick={handleWeeklyReview}
        size="sm"
        variant="ghost"
        className="h-8 px-2.5 sm:px-3 text-xs rounded-none gap-1 sm:gap-1.5 hover:bg-background hover:shadow-sm"
      >
        <Calendar className="h-3.5 w-3.5 text-blue-500" />
        <span>Weekly</span>
        <span className="hidden sm:inline">&nbsp;Review</span>
      </Button>
    </div>
  )
}