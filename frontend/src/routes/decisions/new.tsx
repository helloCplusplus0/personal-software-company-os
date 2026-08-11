import { createFileRoute } from '@tanstack/react-router'
import { z } from 'zod'
import { DecisionCreatePage } from '@/features/decision-center/pages/decision-create-page'
import { dashboardSourceSearchSchema } from '@/features/dashboard/lib/dashboard-source-schema'

/**
 * DecisionCreateRoute — /decisions/new
 *
 * 承接 RecordDecision（§5.5）。
 *
 * 搜索参数承接来源上下文（§5.11）：
 * - sourceModuleId: 从 Module Detail 带入的来源 Module 标识
 * - sourceModuleName: 从 Module Detail 带入的来源 Module 名称（仅用于展示）
 *
 * 这两个参数为可选，直接进入 /decisions/new 时不携带。
 *
 * fromList（§9.1 列表上下文跨页面恢复的单值化标记）：
 * - 由 DecisionListPage 在导航到本页时显式置为 true
 * - 从 Module Detail 或外部直达进入时不存在（undefined）
 * - 用于返回列表时单值化判断“来源列表上下文存在 / 不存在”：
 *   存在则恢复 lastSearch，不存在则落到默认参数，不恢复历史筛选
 */
const decisionCreateSearchSchema = z.object({
  sourceModuleId: z.string().optional(),
  sourceModuleName: z.string().optional(),
  fromList: z.boolean().optional(),
  ...dashboardSourceSearchSchema,
})

export const Route = createFileRoute('/decisions/new')({
  validateSearch: decisionCreateSearchSchema,
  component: DecisionCreatePage,
})
