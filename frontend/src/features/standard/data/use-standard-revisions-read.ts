/**
 * useStandardRevisionsRead — Standard Revisions 只读 query owner。
 *
 * phase14-05 §"切片结构必须冻结"（project_rules §2.5）：query 层纯只读，唯一 owner。
 * ListStandardRevisions 不分页，按 created_at DESC 直接投影。
 */
import { useQuery, type UseQueryResult } from '@tanstack/react-query'
import { standardClient } from './connect-client'
import { pbToRevision } from '../types'
import type { StandardRevision } from '../types'

export type UseStandardRevisionsRead = UseQueryResult<StandardRevision[], Error>

export function useStandardRevisionsRead(standardId: string): UseStandardRevisionsRead {
  return useQuery<StandardRevision[], Error>({
    queryKey: ['standard-revisions', standardId],
    queryFn: async (): Promise<StandardRevision[]> => {
      const res = await standardClient.listStandardRevisions({ standardId })
      return (res.revisions ?? []).map((r) => pbToRevision(r))
    },
    enabled: !!standardId,
  })
}
