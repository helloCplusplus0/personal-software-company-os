import { createFileRoute } from '@tanstack/react-router'
import { z } from 'zod'
import { RepositoryCreatePage } from '@/features/repository-binding/pages/repository-create-page'
import { dashboardSourceSearchSchema } from '@/features/dashboard/lib/dashboard-source-schema'

/**
 * RepositoryCreateRoute — /repositories/new
 *
 * phase04-06 来源上下文承接（由路由搜索参数派生，只允许四种之一）：
 * - fromList: 由 RepositoryBindingListPage 导航到本页时显式置为 true，承接 queryText / statusFilter
 * - fromProductDetail: 由 Product Detail 上下文入口跳入时置为 true，承接 productId / productName
 * - fromModuleDetail: 由 Module Detail 兼容入口跳入时置为 true，承接 moduleId / moduleName
 * - 无来源参数 → direct-entry
 *
 * phase04-06 来源上下文透传（Product Detail 原始来源）：
 * - 从 Product Detail 进入时，必须继续携带 Product Detail 自己的来源上下文（使用 product 前缀区分）
 * - 返回 Product Detail 时，基于这些参数恢复 Product Detail 的来源标记，不得退化为 direct-entry
 *
 * phase04-06 刷新恢复规则：
 * - 刷新后必须继续恢复来源标记，不得静默丢失
 * - 无来源列表上下文时返回列表落默认筛选参数
 */
const repositoryCreateSearchSchema = z.object({
  // fromList 来源标记
  fromList: z.boolean().optional(),
  queryText: z.string().optional(),
  statusFilter: z.enum(['all', 'active', 'archived']).optional(),
  // fromProductDetail 来源标记（phase04-05 路由参数命名）
  fromProductDetail: z.boolean().optional(),
  productId: z.string().optional(),
  productName: z.string().optional(),
  // fromModuleDetail 来源标记（phase04-05 路由参数命名）
  fromModuleDetail: z.boolean().optional(),
  moduleId: z.string().optional(),
  moduleName: z.string().optional(),
  // phase04-06 Product Detail 来源上下文透传（使用 product 前缀区分，避免与自身来源参数冲突）
  productFromList: z.boolean().optional(),
  productQueryText: z.string().optional(),
  productStatusFilter: z.enum(['all', 'active', 'archived']).optional(),
  productFromModuleDetail: z.boolean().optional(),
  productModuleId: z.string().optional(),
  productModuleName: z.string().optional(),
  // phase05-13 Dashboard 来源参数
  ...dashboardSourceSearchSchema,
})

export const Route = createFileRoute('/repositories/new')({
  validateSearch: repositoryCreateSearchSchema,
  component: RepositoryCreatePage,
})
