import { useCallback, useMemo } from 'react'
import { useUpdateDecisionStatus } from './use-update-decision-status'
import type { DecisionStatus } from '../types'

export interface DecisionDetailActionSuccess {
  targetStatus: DecisionStatus
}

export function useDecisionDetailActions(decisionId: string) {
  const updateStatusMutation = useUpdateDecisionStatus()

  const advanceDecisionStatus = useCallback(
    async (nextStatus: DecisionStatus): Promise<DecisionDetailActionSuccess> => {
      await updateStatusMutation.mutateAsync({
        decisionId,
        status: nextStatus,
      })

      return {
        targetStatus: nextStatus,
      }
    },
    [decisionId, updateStatusMutation],
  )

  return useMemo(() => ({
    advanceDecisionStatus,
    isSubmitting: updateStatusMutation.isPending,
    error: updateStatusMutation.error,
    reset: updateStatusMutation.reset,
  }), [
    advanceDecisionStatus,
    updateStatusMutation.error,
    updateStatusMutation.isPending,
    updateStatusMutation.reset,
  ])
}
