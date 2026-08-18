/**
 * useRepositoryStandardsRead — Repository 关联 Standard 只读 query owner。
 *
 * phase14-05 §"Repository detail 画像区让位方案"：
 *   调 projectContextClient.getProjectBrief 仅投影 standards[]，
 *   与 agent 消费主路径同源（GetProjectBrief.standards[]），不新增 RPC。
 * brief 中 standards 复用 psco.standard.v1.Standard 同一 pb 消息，pbToStandard 直接适用。
 */
import { useQuery, type UseQueryResult } from '@tanstack/react-query'
import { projectContextClient } from './connect-client'
import { pbToStandard } from '../types'
import type { Standard } from '../types'

export type UseRepositoryStandardsRead = UseQueryResult<Standard[], Error>

export function useRepositoryStandardsRead(repositoryId: string): UseRepositoryStandardsRead {
  return useQuery<Standard[], Error>({
    queryKey: ['repository-standards', repositoryId],
    queryFn: async (): Promise<Standard[]> => {
      const res = await projectContextClient.getProjectBrief({ repositoryId })
      return (res.standards ?? [])
        .map((s) => pbToStandard(s))
        .filter((s): s is Standard => s !== null)
    },
    enabled: !!repositoryId,
  })
}
