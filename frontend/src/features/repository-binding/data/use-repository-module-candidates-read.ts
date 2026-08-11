/**
 * useRepositoryModuleCandidatesRead — Repository Module Candidates 只读 query owner。
 *
 * phase07-10 §5.4：query 层纯只读，唯一 owner。
 * 替换 repository-binding/data/api-adapter.ts 的 fetchRepositoryModuleCandidates。
 */
import { useQuery, type UseQueryResult } from '@tanstack/react-query'
import { repositoryBindingClient } from './connect-client'
import type { RepositoryModuleCandidate } from '../types'

export type UseRepositoryModuleCandidatesRead = UseQueryResult<RepositoryModuleCandidate[], Error>

export function useRepositoryModuleCandidatesRead(repositoryId: string): UseRepositoryModuleCandidatesRead {
  return useQuery<RepositoryModuleCandidate[], Error>({
    queryKey: ['repository-module-candidates', repositoryId],
    queryFn: async () => {
      const res = await repositoryBindingClient.listRepositoryModuleCandidates({ repositoryId })
      return (res.candidates ?? []).map((c) => ({
        module_id: c.moduleId ?? '',
        module_name: c.moduleName ?? '',
        module_status: (c.moduleStatus === 1 ? 'active' : c.moduleStatus === 2 ? 'archived' : 'active') as RepositoryModuleCandidate['module_status'],
      }))
    },
    enabled: !!repositoryId,
  })
}