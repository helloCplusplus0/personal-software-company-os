import { createFileRoute } from '@tanstack/react-router'
import { z } from 'zod'
import { dashboardSourceSearchSchema } from '@/features/dashboard/lib/dashboard-source-schema'
import { DailyReviewPage } from '@/features/review/pages/daily-review-page'

/**
 * DailyReviewRoute — /reviews/daily
 *
 * phase08-08 §"review route 必须以两条独立文件路由落地"：
 *   - 必须定义 validateSearch
 *   - validateSearch 必须继续复用 dashboardSourceSearchSchema
 *   - 正式路径固定为 /reviews/daily
 *
 * 路由约束：
 *   - 路由文件只承接 DailyReviewPage
 *   - 路由文件不得直接拼 query、mutation、toast、section-level retry
 */
const dailyReviewSearchSchema = z.object({
  ...dashboardSourceSearchSchema,
})

export const Route = createFileRoute('/reviews/daily')({
  validateSearch: dailyReviewSearchSchema,
  component: DailyReviewPage,
})