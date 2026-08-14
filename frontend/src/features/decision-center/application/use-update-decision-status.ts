/**
 * useUpdateDecisionStatus — Decision.status 正式状态推进的固定 mutation 承接位。
 *
 * fix_002_003：canonical 状态推进写动作唯一 owner。
 * 状态推进成功后统一失效 detail / list / review / dashboard 读取。
 */
import { useMutation, useQueryClient, type UseMutationResult } from '@tanstack/react-query'
import { decisionCenterClient } from '../data/connect-client'
import { DecisionStatus } from '@/gen/proto/psco/decision_center/v1/decision_center_pb'
import type { DecisionStatus as DomainStatus } from '../types'
import { DAILY_REVIEW_QUERY_KEY, WEEKLY_REVIEW_QUERY_KEY } from '@/features/review/data/review-query-options'

const DOMAIN_TO_PROTO: Record<DomainStatus, DecisionStatus> = {
  proposed: DecisionStatus.PROPOSED,
  active: DecisionStatus.ACTIVE,
  superseded: DecisionStatus.SUPERSEDED,
  archived: DecisionStatus.ARCHIVED,
}

export function useUpdateDecisionStatus(): UseMutationResult<
  void,
  Error,
  { decisionId: string; status: DomainStatus },
  unknown
> {
  const queryClient = useQueryClient()
  return useMutation<void, Error, { decisionId: string; status: DomainStatus }, unknown>({
    mutationFn: async ({ decisionId, status }) => {
      await decisionCenterClient.updateDecisionStatus({
        decisionId,
        status: DOMAIN_TO_PROTO[status],
      })
    },
    onSuccess: async (_data, variables) => {
      // 让 mutation 在相关 reread 完成前保持 pending，
      // 避免来源页先返回、再晚一步收口的状态抖动。
      await Promise.all([
        queryClient.invalidateQueries({ queryKey: ['decision-detail', variables.decisionId] }),
        queryClient.invalidateQueries({ queryKey: ['decision-list'] }),
        queryClient.invalidateQueries({ queryKey: DAILY_REVIEW_QUERY_KEY }),
        queryClient.invalidateQueries({ queryKey: WEEKLY_REVIEW_QUERY_KEY }),
        queryClient.invalidateQueries({ queryKey: ['dashboard-feedback-signals'] }),
        queryClient.invalidateQueries({ queryKey: ['dashboard-overview'] }),
        queryClient.invalidateQueries({ queryKey: ['dashboard-recent-activities'] }),
      ])
    },
  })
}
