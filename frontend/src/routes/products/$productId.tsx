import { createFileRoute } from '@tanstack/react-router'
import { z } from 'zod'
import { ProductDetailPage } from '@/features/product-registry/pages/product-detail-page'
import { dashboardSourceSearchSchema } from '@/features/dashboard/lib/dashboard-source-schema'
import { onboardingSourceSearchSchema } from '@/features/onboarding/lib/onboarding-source-schema'
import { reviewSourceSearchSchema } from '@/features/review/lib/review-source-schema'

/**
 * ProductDetailRoute — /products/:productId
 *
 * phase04-06 来源上下文承接（由路由搜索参数派生，只允许三种之一）：
 * - fromList: 由 ProductListPage 导航到本页、或由 ProductCreatePage 创建成功回流时显式置为 true，
 *   承接 queryText / statusFilter
 * - fromModuleDetail: 由 Module Detail 兼容入口跳入时置为 true，承接 moduleId / moduleName
 *   （用于预填 ProductModuleBindingPanel 的候选 Module 选择）
 * - 无来源参数 → direct-entry
 *
 * phase04-06 刷新恢复规则：
 * - 刷新后必须继续恢复来源标记，不得静默丢失
 * - 从 ProductCreatePage 成功创建后进入时，来源上下文继承自创建页
 * - 无来源列表上下文时返回列表落默认筛选参数
 *
 * phase04-05 上下文承接路由参数：
 * - moduleId / moduleName / fromModuleDetail 作为上下文搜索参数传递
 * - productId 作为目标页身份参数，由 URL 路径参数承接
 */
const productDetailSearchSchema = z.object({
  // fromList 来源标记
  fromList: z.boolean().optional(),
  queryText: z.string().optional(),
  statusFilter: z.enum(['all', 'active', 'archived']).optional(),
  // fromModuleDetail 来源标记（phase04-05 路由参数命名）
  fromModuleDetail: z.boolean().optional(),
  moduleId: z.string().optional(),
  moduleName: z.string().optional(),
  // phase09-09 模板来源标记
  fromTemplateReuse: z.boolean().optional(),
  templateCandidateId: z.string().optional(),
  templateSource: z.enum(['weekly-review', 'dashboard', 'product-detail']).optional(),
  // phase05-13 Dashboard 来源参数
  ...dashboardSourceSearchSchema,
  // phase10-10 Review 来源参数
  ...reviewSourceSearchSchema,
  // phase06-15 Onboarding 来源参数
  ...onboardingSourceSearchSchema,
})

export const Route = createFileRoute('/products/$productId')({
  validateSearch: productDetailSearchSchema,
  component: ProductDetailPage,
})
