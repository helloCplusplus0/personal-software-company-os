/**
 * Dashboard 来源参数 Zod schema 工厂
 *
 * phase05-13 §"既有路由必须承接 Dashboard 来源参数"
 *
 * 为避免在四类 List / 四类 Detail / 三类 Create 路由文件中重复定义相同的
 * Dashboard 来源参数 schema，统一在此处导出 `dashboardSourceSearchSchema`，
 * 供各路由通过 spread 操作符扩展到既有 validateSearch schema。
 *
 * 三字段均为可选：
 *   - fromDashboard: z.boolean().optional()
 *   - dashboardSection: z.enum([...]).optional()
 *   - dashboardReturnTo: z.string().optional()
 *
 * 约束：
 *   - 不得移除各路由原生的搜索参数（如 queryText / statusFilter / fromList）
 *   - 不得改变原生搜索参数的默认值与 catch 行为
 *   - dashboardSection 只允许 phase05-03 / phase05-10 §8.2 已冻结的五个取值
 */
import { z } from 'zod'

/**
 * Dashboard 来源参数 Zod schema 片段。
 *
 * 通过 spread 扩展到既有路由 schema：
 * ```ts
 * const xxxSearchSchema = z.object({
 *   ...原生字段,
 *   ...dashboardSourceSearchSchema,
 * })
 * ```
 */
export const dashboardSourceSearchSchema = {
  fromDashboard: z.boolean().optional(),
  dashboardSection: z
    .enum(['overview', 'current-focus', 'asset-feedback', 'recent-activity', 'empty-state'])
    .optional(),
  dashboardReturnTo: z.string().optional(),
} as const

/**
 * Dashboard 来源参数的 TypeScript 推断类型（与 types.ts 中 DashboardSourceSearch 对齐）。
 *
 * 供页面组件从 useSearch 读取后获得类型安全。
 */
export type DashboardSourceSearchParsed = {
  fromDashboard?: boolean
  dashboardSection?: 'overview' | 'current-focus' | 'asset-feedback' | 'recent-activity' | 'empty-state'
  dashboardReturnTo?: string
}
