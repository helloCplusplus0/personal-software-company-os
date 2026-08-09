/**
 * DashboardStatBar — 资产概览与缺口计数合并的单行紧凑条带
 *
 * phase05-13 体验修复：
 * - 将原 DashboardOverviewSection 的 6 个概览数字与 AssetFeedbackSection 的 4 个缺口计数
 *   合并为单行紧凑 stat bar，让用户一行看完全局资产盘子与绑定健康度
 * - 概览组前 4 个数字（模块/产品/仓库/决策）可点击跳转到对应 List
 * - 缺口组按严重度着色：双缺 destructive、缺仓/缺模 amber
 * - aria-label 仍区分「系统概览」与「资产覆盖」两个 section，保留无障碍语义
 *
 * 数据来源：
 * - overview 来自 overviewQuery（页面渲染 stat bar 时 overview 必为 ready）
 * - coverage 来自 feedbackQuery（可能 loading/error/ready）
 *
 * 布局降级：
 * - PC 桌面：10 个 cell 单行排列，概览组与缺口组之间用更明显的分隔区分
 * - 移动浏览器：flex-wrap 折行，仍保持单容器不撑开页面
 */
import { useNavigate } from '@tanstack/react-router'
import { Skeleton } from '@/components/ui/skeleton'
import type { DashboardOverview, ProductAssetCoverageSummary } from '../types'
import { buildDashboardSourceParams } from '../lib/dashboard-source'

interface DashboardStatBarProps {
  // overview query（页面渲染本组件时 overview 已 ready）
  overview: DashboardOverview
  // coverage 来自 feedbackQuery 的 asset_feedback_summary
  coverageStatus: 'loading' | 'ready' | 'error'
  summary: ProductAssetCoverageSummary | undefined
}

/**
 * 概览组单值 cell 数据。
 * clickable 的 cell 点击跳转到对应 canonical List，携带 dashboardSection=overview。
 */
interface OverviewStat {
  label: string
  value: number
  clickable: boolean
  to?: '/modules' | '/products' | '/repositories' | '/decisions'
}

/**
 * 缺口组单值 cell 数据，按严重度着色。
 */
interface CoverageStat {
  label: string
  value: number
  // tone：destructive 用于「双缺」，amber 用于「缺仓/缺模」，muted 用于「完整绑定」
  tone: 'destructive' | 'amber' | 'muted'
}

/**
 * DashboardStatBar — 资产概览 + 缺口计数的紧凑单行条带。
 *
 * 视觉分组：
 *   [模块 | 产品 | 仓库 | 决策 | 已绑仓 | 已绑模]  ‖  [完整 | 双缺 | 缺仓 | 缺模]
 *   ─────────── 概览组（前 4 可点击）───────────      ───── 缺口组（着色）─────
 */
export function DashboardStatBar({
  overview,
  coverageStatus,
  summary,
}: DashboardStatBarProps) {
  const navigate = useNavigate()

  const overviewStats: OverviewStat[] = [
    { label: '模块', value: overview.module_count, clickable: true, to: '/modules' },
    { label: '产品', value: overview.product_count, clickable: true, to: '/products' },
    { label: '仓库', value: overview.repository_count, clickable: true, to: '/repositories' },
    { label: '决策', value: overview.decision_count, clickable: true, to: '/decisions' },
    { label: '已绑仓', value: overview.product_with_repository_count, clickable: false },
    { label: '已绑模', value: overview.product_with_module_count, clickable: false },
  ]

  const coverageStats: CoverageStat[] = summary
    ? [
        { label: '完整', value: summary.fully_bound_product_count, tone: 'muted' },
        { label: '双缺', value: summary.missing_both_bindings_count, tone: 'destructive' },
        { label: '缺仓', value: summary.missing_repository_binding_count, tone: 'amber' },
        { label: '缺模', value: summary.missing_module_binding_count, tone: 'amber' },
      ]
    : []

  const handleOverviewClick = (to: NonNullable<OverviewStat['to']>) => {
    navigate({ to, search: buildDashboardSourceParams('overview') })
  }

  return (
    <div className="flex flex-wrap items-stretch overflow-hidden rounded-lg border bg-card divide-x divide-border">
      {/* 概览组 — aria-label 保留「系统概览」section 语义 */}
      <section aria-label="系统概览" className="flex flex-wrap divide-x divide-border">
        {overviewStats.map((stat) => {
          // 可点击 cell 用 <button>（整格可点击 + 进 tab 序列）；纯展示 cell 用 <div>（不进 tab 序列）
          const cellClassName = [
            'flex min-w-[68px] flex-col justify-center px-3 py-2 text-left',
            stat.clickable
              ? 'cursor-pointer hover:bg-muted/50 transition-colors'
              : 'cursor-default',
          ].join(' ')
          const inner = (
            <>
              <span className="text-lg font-bold leading-none tabular-nums">
                {stat.value}
              </span>
              <span className="mt-1 text-[10px] text-muted-foreground">{stat.label}</span>
            </>
          )
          return stat.clickable && stat.to ? (
            <button
              key={stat.label}
              type="button"
              onClick={() => handleOverviewClick(stat.to!)}
              className={cellClassName}
            >
              {inner}
            </button>
          ) : (
            <div key={stat.label} className={cellClassName}>
              {inner}
            </div>
          )
        })}
      </section>

      {/* 缺口组 — 与概览组之间用更明显的左边距 + 边框区分；aria-label 保留「资产覆盖」section 语义 */}
      <section
        aria-label="资产覆盖"
        className="flex flex-wrap border-l-2 border-border/60 divide-x divide-border"
      >
        {coverageStats.length === 0 && coverageStatus === 'loading' ? (
          // coverage 加载中：4 个骨架 cell
          Array.from({ length: 4 }).map((_, i) => (
            <div key={i} className="flex min-w-[68px] flex-col justify-center px-3 py-2">
              <Skeleton className="h-5 w-6" />
              <Skeleton className="mt-1 h-2.5 w-8" />
            </div>
          ))
        ) : coverageStats.length === 0 && coverageStatus === 'error' ? (
          // coverage 读取失败：4 个占位 cell
          ['完整', '双缺', '缺仓', '缺模'].map((label) => (
            <div
              key={label}
              className="flex min-w-[68px] flex-col justify-center px-3 py-2"
            >
              <span className="text-lg font-bold leading-none text-muted-foreground">–</span>
              <span className="mt-1 text-[10px] text-muted-foreground">{label}</span>
            </div>
          ))
        ) : (
          coverageStats.map((stat) => (
            <div
              key={stat.label}
              className="flex min-w-[68px] flex-col justify-center px-3 py-2"
            >
              <span
                className={[
                  'text-lg font-bold leading-none tabular-nums',
                  stat.tone === 'destructive'
                    ? 'text-destructive'
                    : stat.tone === 'amber'
                      ? 'text-amber-600 dark:text-amber-400'
                      : 'text-foreground',
                ].join(' ')}
              >
                {stat.value}
              </span>
              <span className="mt-1 text-[10px] text-muted-foreground">{stat.label}</span>
            </div>
          ))
        )}
      </section>
    </div>
  )
}
