/**
 * useModuleListRead — Module List 只读 query owner。
 *
 * phase07-10 §5.4：query 层纯只读，唯一 owner。
 * 替换 module-registry/data/api-adapter.ts 的 fetchModuleList。
 */
import { useQuery, type UseQueryResult } from '@tanstack/react-query'
import { moduleRegistryClient } from './connect-client'
import type { ModuleListItem, ModuleListSearch, ModuleStatus } from '../types'

export type UseModuleListRead = UseQueryResult<ModuleListItem[], Error>

export function useModuleListRead(search: ModuleListSearch): UseModuleListRead {
  return useQuery<ModuleListItem[], Error>({
    queryKey: ['module-list', search],
    queryFn: async (): Promise<ModuleListItem[]> => {
      const res = await moduleRegistryClient.listModules({
        queryText: search.queryText ?? '',
        statusFilter: mapModuleStatusToProto(search.statusFilter ?? 'all'),
      })
      return (res.modules ?? []).map((m) => ({
        id: m.id ?? '',
        name: m.name ?? '',
        description: m.description ?? '',
        status: mapProtoToModuleStatus(m.status) as ModuleStatus,
        latest_release: m.latestRelease ?? null,
        product_bind_count: m.productBindCount ?? 0,
        repository_bind_count: m.repositoryBindCount ?? 0,
      }))
    },
  })
}

// Helper: map string status to proto enum
function mapModuleStatusToProto(s: string): number {
  switch (s) {
    case 'active': return 1
    case 'archived': return 2
    default: return 0
  }
}

function mapProtoToModuleStatus(v: number): string {
  switch (v) {
    case 1: return 'active'
    case 2: return 'archived'
    default: return ''
  }
}