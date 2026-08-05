import { createFileRoute } from '@tanstack/react-router'
import { z } from 'zod'
import { ModuleListPage } from '@/features/module-registry/pages/module-list-page'

/**
 * ModuleListRoute — /modules
 *
 * 搜索参数冻结到路由搜索参数层（§8.4）：
 * - queryText: 文本筛选
 * - statusFilter: 状态筛选
 *
 * 从 ModuleCreatePage 或 ModuleDetailPage 返回时，
 * 必须按原有搜索参数恢复列表上下文（§7.4）
 */
const moduleListSearchSchema = z.object({
  queryText: z.string().optional(),
  // .default('all') 使导航到 /modules 时 search 可选（输入层 statusFilter 可省略）
  // .catch('all') 保证 URL 中出现非法值时优雅降级为 'all'
  statusFilter: z.enum(['all', 'active', 'archived']).catch('all').default('all'),
})

export const Route = createFileRoute('/modules/')({
  validateSearch: moduleListSearchSchema,
  component: ModuleListPage,
})
