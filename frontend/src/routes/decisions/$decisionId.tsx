import { createFileRoute } from '@tanstack/react-router'
import { z } from 'zod'
import { DecisionDetailPage } from '@/features/decision-center/pages/decision-detail-page'

/**
 * DecisionDetailRoute — /decisions/:decisionId
 *
 * 统一详情读模型宿主，承接详情读取、已关联目标展示、
 * 待关联目标承接、候选读取与最小目标关联动作（§5.8 / §5.10 / §5.11）。
 *
 * fromList（§9.1 列表上下文跨页面恢复的单值化标记）：
 * - 由 DecisionListPage 导航到本页、或由 DecisionCreatePage 创建成功回流时显式置为 true
 * - 从 Module Detail 入口或外部直达进入时不存在（undefined）
 * - 用于返回列表时单值化判断“来源列表上下文存在 / 不存在”：
 *   存在则恢复 lastSearch，不存在则落到默认参数，不恢复历史筛选
 */
const decisionDetailSearchSchema = z.object({
  fromList: z.boolean().optional(),
})

export const Route = createFileRoute('/decisions/$decisionId')({
  validateSearch: decisionDetailSearchSchema,
  component: DecisionDetailPage,
})
