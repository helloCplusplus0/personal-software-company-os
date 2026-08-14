import type { ReviewSourceSearchParsed } from './review-source-schema'

export interface ReviewSourceSearch {
  fromReview?: boolean
  reviewKind?: 'daily' | 'weekly'
  reviewReturnTo?: string
}

/**
 * 构造 Review 来源参数。
 *
 * phase10-09 当前只依赖 detail 页正确返回 Daily Review，
 * 这里仍保留 weekly 以避免后续再造第二套来源合同。
 */
export function buildReviewSourceParams(
  reviewKind: 'daily' | 'weekly',
): ReviewSourceSearch {
  return {
    fromReview: true,
    reviewKind,
    reviewReturnTo: reviewKind === 'daily' ? '/reviews/daily' : '/reviews/weekly',
  }
}

export function shouldReturnToReview(
  search: ReviewSourceSearchParsed,
): boolean {
  return search.fromReview === true && !!search.reviewReturnTo
}

export function buildReviewReturnSearch(
  search: ReviewSourceSearchParsed,
): ReviewSourceSearch {
  if (!shouldReturnToReview(search)) {
    return {}
  }

  return {
    fromReview: true,
    reviewKind: search.reviewKind,
    reviewReturnTo: search.reviewReturnTo,
  }
}
