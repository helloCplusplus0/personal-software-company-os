/**
 * useDecisionDetailRead — Decision Detail 只读 query owner。
 *
 * phase07-10 §5.4：query 层纯只读，唯一 owner。
 * 替换 decision-center/data/api-adapter.ts 的 fetchDecisionDetail。
 */
import { useQuery, type UseQueryResult } from '@tanstack/react-query'
import { timestampDate } from '@bufbuild/protobuf/wkt'
import { decisionCenterClient } from './connect-client'
import type { DecisionDetail } from '../types'

export type UseDecisionDetailRead = UseQueryResult<DecisionDetail, Error>

export function useDecisionDetailRead(decisionId: string): UseDecisionDetailRead {
  return useQuery<DecisionDetail, Error>({
    queryKey: ['decision-detail', decisionId],
    queryFn: async () => {
      const res = await decisionCenterClient.getDecisionDetail({ decisionId })
      const d = res.decisionDetail
      if (!d) throw new Error('decision not found')
      const dec = d.decision
      return {
        decision: {
          id: dec?.id ?? '',
          title: dec?.title ?? '',
          context: dec?.context ?? '',
          problem: dec?.problem ?? '',
          alternatives: dec?.alternatives ?? [],
          choice: dec?.choice ?? '',
          reason: dec?.reason ?? '',
          impact: dec?.impact ?? '',
          status: (dec?.status === 1 ? 'proposed' : dec?.status === 2 ? 'active' : dec?.status === 3 ? 'superseded' : dec?.status === 4 ? 'archived' : '') as DecisionDetail['decision']['status'],
          created_at: dec?.createdAt ? timestampDate(dec.createdAt).toISOString() : '',
        },
        linked_modules: (d.linkedModules ?? []).map((m) => ({
          module_id: m.moduleId ?? '',
          module_name: m.moduleName ?? '',
        })),
        source_context: {
          source_module_id: d.sourceContext?.sourceModuleId ?? '',
          source_module_name: d.sourceContext?.sourceModuleName ?? '',
        },
      }
    },
    enabled: !!decisionId,
  })
}