import { createFileRoute } from '@tanstack/react-router'
import { z } from 'zod'
import { dashboardSourceSearchSchema } from '@/features/dashboard/lib/dashboard-source-schema'
import { WeeklyReviewPage } from '@/features/review/pages/weekly-review-page'

/**
 * WeeklyReviewRoute — /reviews/weekly
 *
 * phase08-08 §"review route 必须以两条独立文件路由落地"：
 *   - 必须定义 validateSearch
 *   - validateSearch 必须继续复用 dashboardSourceSearchSchema
 *   - 正式路径固定为 /reviews/weekly
 *
 * 路由约束：
 *   - 路由文件只承接 WeeklyReviewPage
 *   - 路由文件不得直接拼 query、mutation、toast、section-level retry
 */
const weeklyReviewSearchSchema = z.object({
  ...dashboardSourceSearchSchema,
  // phase09-10 模板复用提示返回链参数（用于 reread active candidate 恢复）
  returnCandidateId: z.string().optional(),
})

export const Route = createFileRoute('/reviews/weekly')({
  validateSearch: weeklyReviewSearchSchema,
  component: WeeklyReviewPage,
})