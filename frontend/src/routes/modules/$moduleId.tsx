import { createFileRoute } from '@tanstack/react-router'
import { z } from 'zod'
import { ModuleDetailPage } from '@/features/module-registry/pages/module-detail-page'
import { dashboardSourceSearchSchema } from '@/features/dashboard/lib/dashboard-source-schema'
import { onboardingSourceSearchSchema } from '@/features/onboarding/lib/onboarding-source-schema'

/**
 * ModuleDetailRoute — /modules/:moduleId
 * 统一详情读模型宿主，承接详情读取、版本登记入口、绑定面板与 Decision 只读入口（§3.1）
 *
 * phase05-13：新增 validateSearch 承接 Dashboard 来源参数，
 * 用于从 Dashboard 活动项跳转后保留返回 Dashboard 上下文。
 */
const moduleDetailSearchSchema = z.object({
  fromList: z.boolean().optional(),
  queryText: z.string().optional(),
  statusFilter: z.enum(['all', 'active', 'archived']).optional(),
  ...dashboardSourceSearchSchema,
  // phase06-15 Onboarding 来源参数
  ...onboardingSourceSearchSchema,
})

export const Route = createFileRoute('/modules/$moduleId')({
  validateSearch: moduleDetailSearchSchema,
  component: ModuleDetailPage,
})
