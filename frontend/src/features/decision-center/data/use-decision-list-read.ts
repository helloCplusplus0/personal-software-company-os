/**
 * useDecisionListRead — Decision List 只读 query owner。
 *
 * phase07-10 §5.4：query 层纯只读，唯一 owner。
 * 替换 decision-center/data/api-adapter.ts 的 fetchDecisionList。
 */
import { useQuery, type UseQueryResult } from '@tanstack/react-query'
import { timestampDate } from '@bufbuild/protobuf/wkt'
import { decisionCenterClient } from './connect-client'
import type { DecisionListItem, DecisionListSearch, DecisionStatus } from '../types'

export type UseDecisionListRead = UseQueryResult<DecisionListItem[], Error>

export function useDecisionListRead(search: DecisionListSearch): UseDecisionListRead {
  return useQuery<DecisionListItem[], Error>({
    queryKey: ['decision-list', search],
    queryFn: async (): Promise<DecisionListItem[]> => {
      const res = await decisionCenterClient.listDecisions({
        queryText: search.queryText ?? '',
        statusFilter: mapDecisionStatusToProto(search.statusFilter ?? 'all'),
      })
      return (res.decisions ?? []).map((d) => ({
        id: d.id ?? '',
        title: d.title ?? '',
        status: mapProtoToDecisionStatus(d.status) as DecisionStatus,
        created_at: d.createdAt ? timestampDate(d.createdAt).toISOString() : '',
        link_count: d.linkCount ?? 0,
        linked_module_summary: d.linkedModuleSummary ?? '',
      }))
    },
  })
}

function mapDecisionStatusToProto(s: string): number {
  switch (s) {
    case 'proposed': return 1
    case 'active': return 2
    case 'superseded': return 3
    case 'archived': return 4
    default: return 0
  }
}

function mapProtoToDecisionStatus(v: number): string {
  switch (v) {
    case 1: return 'proposed'
    case 2: return 'active'
    case 3: return 'superseded'
    case 4: return 'archived'
    default: return ''
  }
}