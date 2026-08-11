/**
 * useLinkDecisionToTarget — Decision->Module 关联的固定 mutation 承接位。
 *
 * phase07-10 §5.5：canonical 写动作单一正式 owner。
 */
import { useMutation, useQueryClient, type UseMutationResult } from '@tanstack/react-query'
import { decisionCenterClient } from '../data/connect-client'
import { DecisionLinkTargetType } from '@/gen/proto/psco/decision_center/v1/decision_center_pb'

export function useLinkDecisionToTarget(): UseMutationResult<void, Error, { decisionId: string; moduleId: string }, unknown> {
  const queryClient = useQueryClient()
  return useMutation<void, Error, { decisionId: string; moduleId: string }, unknown>({
    mutationFn: async ({ decisionId, moduleId }) => {
      await decisionCenterClient.linkDecisionToTarget({
        decisionId,
        targetType: DecisionLinkTargetType.MODULE,
        moduleId,
      })
    },
    onSuccess: (_data, variables) => {
      queryClient.invalidateQueries({ queryKey: ['decision-detail', variables.decisionId] })
      queryClient.invalidateQueries({ queryKey: ['decision-module-candidates', variables.decisionId] })
    },
  })
}