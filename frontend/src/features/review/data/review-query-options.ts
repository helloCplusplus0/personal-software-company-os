/**
 * review-query-options — Review 只读 query key 与 query 配置。
 *
 * phase08-06 / phase08-08 §"review read layer 的缓存口径仍然是轻量组合层"：
 *   - 当前阶段不把 ['review'] 冻结成必需的长期 query key
 *   - 若后续实现确实引入 review slice-local query key，它的身份也只能是组合层缓存别名
 */
import { queryOptions } from '@tanstack/react-query'
import { reviewClient } from './connect-client'

/**
 * Daily Review Context 的 query key。
 * 当前阶段作为组合层缓存别名，不是新的事实主缓存。
 */
export const DAILY_REVIEW_QUERY_KEY = ['daily-review-context'] as const

/**
 * Weekly Review Context 的 query key。
 * 当前阶段作为组合层缓存别名，不是新的事实主缓存。
 */
export const WEEKLY_REVIEW_QUERY_KEY = ['weekly-review-context'] as const

/**
 * dailyReviewQueryOptions — Daily Review Context 的 query options。
 */
export function dailyReviewQueryOptions() {
  return queryOptions({
    queryKey: DAILY_REVIEW_QUERY_KEY,
    queryFn: async ({ signal }) => {
      const res = await reviewClient.getDailyReviewContext({}, { signal })
      return res.context
    },
  })
}

/**
 * weeklyReviewQueryOptions — Weekly Review Context 的 query options。
 */
export function weeklyReviewQueryOptions() {
  return queryOptions({
    queryKey: WEEKLY_REVIEW_QUERY_KEY,
    queryFn: async ({ signal }) => {
      const res = await reviewClient.getWeeklyReviewContext({}, { signal })
      return res.context
    },
  })
}
