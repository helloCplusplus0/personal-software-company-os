import { createFileRoute } from '@tanstack/react-router'
import { z } from 'zod'
import { RepositoryBindingListPage } from '@/features/repository-binding/pages/repository-binding-list-page'

/**
 * RepositoryBindingListRoute — /repositories
 *
 * phase04-06 列表查询条件承接策略：
 * - queryText 与 statusFilter 必须冻结到路由搜索参数层，作为列表筛选的唯一事实源
 * - 路由搜索参数是列表查询条件的唯一事实源，不使用 sessionStorage 等持久化层
 *
 * phase04-05 上下文入口路由参数：
 * - 从 Product Detail 进入时携带 fromProductDetail + productId / productName
 * - 从 Module Detail 进入时携带 fromModuleDetail + moduleId / moduleName
 * - 这些上下文参数在用户选择目标仓库后继续传递到 /repositories/$repositoryId
 *
 * phase04-06 来源上下文透传（Product Detail 原始来源）：
 * - 从 Product Detail 进入时，必须继续携带 Product Detail 自己的来源上下文（使用 product 前缀区分）
 * - 返回 Product Detail 时，基于这些参数恢复 Product Detail 的来源标记，不得退化为 direct-entry
 * - productFromList + productQueryText / productStatusFilter → Product Detail 来源是 Product List
 * - productFromModuleDetail + productModuleId / productModuleName → Product Detail 来源是 Module Detail
 *
 * 从 RepositoryCreatePage 或 RepositoryBindingDetailPage 返回时，
 * 必须按原有搜索参数恢复列表上下文（phase04-06 返回列表恢复规则）。
 */
const repositoryListSearchSchema = z.object({
  queryText: z.string().optional(),
  statusFilter: z.enum(['all', 'active', 'archived']).catch('all').default('all'),
  // phase04-05 上下文入口：从 Product Detail 进入时携带
  fromProductDetail: z.boolean().optional(),
  productId: z.string().optional(),
  productName: z.string().optional(),
  // phase04-05 上下文入口：从 Module Detail 进入时携带
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
})

export const Route = createFileRoute('/repositories/')({
  validateSearch: repositoryListSearchSchema,
  component: RepositoryBindingListPage,
})
