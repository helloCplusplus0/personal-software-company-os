import { createFileRoute } from '@tanstack/react-router'
import { z } from 'zod'
import { ProductListPage } from '@/features/product-registry/pages/product-list-page'
import { dashboardSourceSearchSchema } from '@/features/dashboard/lib/dashboard-source-schema'

/**
 * ProductListRoute — /products
 *
 * phase04-06 列表查询条件承接策略：
 * - queryText 与 statusFilter 必须冻结到路由搜索参数层，作为列表筛选的唯一事实源
 * - queryText 允许空字符串
 * - statusFilter 只允许 all / active / archived（承接 phase04-02 冻结结论）
 * - 列表默认查询条件为 queryText = 空、statusFilter = all
 * - 路由搜索参数是列表查询条件的唯一事实源，不使用 sessionStorage 等持久化层
 *
 * 从 ProductCreatePage 或 ProductDetailPage 返回时，
 * 必须按原有搜索参数恢复列表上下文（phase04-06 返回列表恢复规则）。
 */
const productListSearchSchema = z.object({
  queryText: z.string().optional(),
  // .default('all') 使导航到 /products 时 search 可选（输入层 statusFilter 可省略）
  // .catch('all') 保证 URL 中出现非法值时优雅降级为 'all'
  statusFilter: z.enum(['all', 'active', 'archived']).catch('all').default('all'),
  // phase04-05 上下文入口：从 Module Detail 进入时携带
  fromModuleDetail: z.boolean().optional(),
  moduleId: z.string().optional(),
  moduleName: z.string().optional(),
  // phase05-13 Dashboard 来源参数
  ...dashboardSourceSearchSchema,
})

export const Route = createFileRoute('/products/')({
  validateSearch: productListSearchSchema,
  component: ProductListPage,
})
