import { z } from 'zod'

/**
 * Review 来源参数 Zod schema 片段。
 *
 * 用于 canonical detail 页在来自 Review 入口时，保留正式返回链。
 * 当前 phase10-09 只要求 Daily Review 闭环，但 schema 预留 weekly，
 * 避免后续继续扩展时再长出第二套来源字段。
 */
export const reviewSourceSearchSchema = {
  fromReview: z.boolean().optional(),
  reviewKind: z.enum(['daily', 'weekly']).optional(),
  reviewReturnTo: z.string().optional(),
} as const

export type ReviewSourceSearchParsed = {
  fromReview?: boolean
  reviewKind?: 'daily' | 'weekly'
  reviewReturnTo?: string
}
