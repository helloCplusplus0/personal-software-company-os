/**
 * useRepositoryListRead — Repository List 只读 query owner。
 *
 * phase07-10 §5.4：query 层纯只读，唯一 owner。
 * 替换 repository-binding/data/api-adapter.ts 的 fetchRepositoryList。
 */
import { useQuery, type UseQueryResult } from '@tanstack/react-query'
import { timestampDate } from '@bufbuild/protobuf/wkt'
import { repositoryBindingClient } from './connect-client'
import type { RepositoryListItem, RepositoryListSearch, RepositoryStatus } from '../types'

export type UseRepositoryListRead = UseQueryResult<RepositoryListItem[], Error>

export function useRepositoryListRead(search: RepositoryListSearch): UseRepositoryListRead {
  return useQuery<RepositoryListItem[], Error>({
    queryKey: ['repository-list', search],
    queryFn: async (): Promise<RepositoryListItem[]> => {
      const res = await repositoryBindingClient.listRepositories({
        queryText: search.queryText ?? '',
        statusFilter: mapStatusToProto(search.statusFilter ?? 'all'),
      })
      return (res.repositories ?? []).map((r) => ({
        id: r.id ?? '',
        name: r.name ?? '',
        url: r.url ?? '',
        provider: r.provider ?? '',
        status: mapProtoToStatus(r.status) as RepositoryStatus,
        created_at: r.createdAt ? timestampDate(r.createdAt).toISOString() : '',
        product_bind_count: r.productBindCount ?? 0,
        module_bind_count: r.moduleBindCount ?? 0,
      }))
    },
  })
}

function mapStatusToProto(s: string): number {
  switch (s) {
    case 'active': return 1
    case 'archived': return 2
    default: return 0
  }
}

function mapProtoToStatus(v: number): string {
  switch (v) {
    case 1: return 'active'
    case 2: return 'archived'
    default: return ''
  }
}