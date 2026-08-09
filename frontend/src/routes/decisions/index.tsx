import { createFileRoute } from '@tanstack/react-router'
import { z } from 'zod'
import { DecisionListPage } from '@/features/decision-center/pages/decision-list-page'
import { dashboardSourceSearchSchema } from '@/features/dashboard/lib/dashboard-source-schema'

/**
 * DecisionListRoute — /decisions
 *
 * 搜索参数冻结到路由搜索参数层（§9.1）：
 * - queryText: 文本筛选
 * - statusFilter: 状态筛选
 *
 * 从 DecisionCreatePage 或 DecisionDetailPage 返回时，
 * 必须按原有搜索参数恢复列表上下文（§9.1 返回列表恢复规则）。
 */
const decisionListSearchSchema = z.object({
  queryText: z.string().optional(),
  // .default('all') 使导航到 /decisions 时 search 可选
  // .catch('all') 保证 URL 中出现非法值时优雅降级为 'all'
  statusFilter: z.enum(['all', 'proposed', 'active', 'superseded', 'archived']).catch('all').default('all'),
  // phase05-13 Dashboard 来源参数
  ...dashboardSourceSearchSchema,
})

export const Route = createFileRoute('/decisions/')({
  validateSearch: decisionListSearchSchema,
  component: DecisionListPage,
})
