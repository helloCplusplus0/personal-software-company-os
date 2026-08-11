/**
 * useDecisionModuleCandidatesRead — Decision Module Candidates 只读 query owner。
 *
 * phase07-10 §5.4：query 层纯只读，唯一 owner。
 * 替换 decision-center/data/api-adapter.ts 的 fetchDecisionModuleCandidates。
 */
import { useQuery, type UseQueryResult } from '@tanstack/react-query'
import { decisionCenterClient } from './connect-client'
import type { DecisionModuleCandidate } from '../types'

export type UseDecisionModuleCandidatesRead = UseQueryResult<DecisionModuleCandidate[], Error>

export function useDecisionModuleCandidatesRead(decisionId: string): UseDecisionModuleCandidatesRead {
  return useQuery<DecisionModuleCandidate[], Error>({
    queryKey: ['decision-module-candidates', decisionId],
    queryFn: async () => {
      const res = await decisionCenterClient.listDecisionModuleCandidates({ decisionId })
      return (res.candidates ?? []).map((c) => ({
        module_id: c.moduleId ?? '',
        module_name: c.moduleName ?? '',
        status: (c.status === 1 ? 'active' : c.status === 2 ? 'archived' : 'active') as DecisionModuleCandidate['status'],
      }))
    },
    enabled: !!decisionId,
  })
}