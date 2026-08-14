import { useMemo } from 'react'
import { useDecisionDetailRead } from './use-decision-detail-read'
import type { DecisionStatus } from '../types'
import {
  buildOnboardingReturnSearch,
  shouldReturnToOnboarding,
} from '@/features/onboarding/lib/onboarding-return'
import {
  buildReviewReturnSearch,
  shouldReturnToReview,
} from '@/features/review/lib/review-source'
import { mergeCurrentDashboardSource } from '@/features/dashboard/lib/dashboard-source'
import type { DashboardSourceSearch } from '@/features/dashboard/types'
import type { OnboardingSourceSearchParsed } from '@/features/onboarding/lib/onboarding-source-schema'
import type { ReviewSourceSearchParsed } from '@/features/review/lib/review-source-schema'
import type { DecisionListSearchContext } from '../types'

type DecisionDetailRouteSearch = DashboardSourceSearch &
  OnboardingSourceSearchParsed &
  ReviewSourceSearchParsed & {
    fromList?: boolean
    queryText?: string
    statusFilter?: 'all' | 'proposed' | 'active' | 'superseded' | 'archived'
  }

export interface DecisionDetailStatusAction {
  label: string
  target: DecisionStatus
}

export interface DecisionDetailPageReadModel {
  decisionDetailQuery: ReturnType<typeof useDecisionDetailRead>
  returnLabel: string
  returnTo:
    | { kind: 'onboarding'; search: Record<string, unknown> }
    | { kind: 'review'; to: string; search: Record<string, unknown> }
    | { kind: 'decision-list'; search: Record<string, unknown> }
  statusActions: DecisionDetailStatusAction[]
}

const STATUS_TRANSITIONS: Record<DecisionStatus, DecisionDetailStatusAction[]> = {
  proposed: [
    { label: 'Mark Active', target: 'active' },
    { label: 'Mark Superseded', target: 'superseded' },
    { label: 'Archive', target: 'archived' },
  ],
  active: [
    { label: 'Mark Superseded', target: 'superseded' },
    { label: 'Archive', target: 'archived' },
  ],
  superseded: [],
  archived: [],
}

export function useDecisionDetailPageRead(
  decisionId: string,
  detailSearch: DecisionDetailRouteSearch,
  lastSearch: DecisionListSearchContext,
): DecisionDetailPageReadModel {
  const decisionDetailQuery = useDecisionDetailRead(decisionId)

  const returnTarget = useMemo<DecisionDetailPageReadModel['returnTo']>(() => {
    if (shouldReturnToOnboarding(detailSearch)) {
      return {
        kind: 'onboarding',
        search: buildOnboardingReturnSearch(detailSearch) as unknown as Record<string, unknown>,
      }
    }

    if (shouldReturnToReview(detailSearch)) {
      return {
        kind: 'review',
        to: detailSearch.reviewReturnTo ?? '/reviews/daily',
        search: buildReviewReturnSearch(detailSearch) as unknown as Record<string, unknown>,
      }
    }

    return {
      kind: 'decision-list',
      search: mergeCurrentDashboardSource(
        detailSearch.fromList
          ? {
              queryText: detailSearch.queryText ?? lastSearch.queryText,
              statusFilter: detailSearch.statusFilter ?? lastSearch.statusFilter,
            }
          : { statusFilter: 'all' as const },
        detailSearch,
      ) as unknown as Record<string, unknown>,
    }
  }, [detailSearch, lastSearch])

  const statusActions = useMemo<DecisionDetailStatusAction[]>(() => {
    const status = decisionDetailQuery.data?.decision.status
    if (!status) {
      return []
    }
    return STATUS_TRANSITIONS[status]
  }, [decisionDetailQuery.data?.decision.status])

  return {
    decisionDetailQuery,
    returnLabel:
      returnTarget.kind === 'onboarding'
        ? '返回首轮录入'
        : returnTarget.kind === 'review'
          ? `返回 ${detailSearch.reviewKind === 'weekly' ? 'Weekly Review' : 'Daily Review'}`
          : '返回列表',
    returnTo: returnTarget,
    statusActions,
  }
}
