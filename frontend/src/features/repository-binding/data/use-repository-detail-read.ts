/**
 * useRepositoryDetailRead — Repository Detail 只读 query owner。
 *
 * phase07-10 §5.4：query 层纯只读，唯一 owner。
 * 替换 repository-binding/data/api-adapter.ts 的 fetchRepositoryDetail。
 */
import { useQuery, type UseQueryResult } from '@tanstack/react-query'
import { timestampDate } from '@bufbuild/protobuf/wkt'
import { repositoryBindingClient } from './connect-client'
import type { RepositoryDetail } from '../types'

export type UseRepositoryDetailRead = UseQueryResult<RepositoryDetail, Error>

export function useRepositoryDetailRead(repositoryId: string): UseRepositoryDetailRead {
  return useQuery<RepositoryDetail, Error>({
    queryKey: ['repository-detail', repositoryId],
    queryFn: async () => {
      const res = await repositoryBindingClient.getRepositoryDetail({ repositoryId })
      const d = res.repositoryDetail
      if (!d) throw new Error('repository not found')
      const r = d.repository
      return {
        repository: {
          id: r?.id ?? '',
          name: r?.name ?? '',
          url: r?.url ?? '',
          provider: r?.provider ?? '',
          status: (r?.status === 1 ? 'active' : r?.status === 2 ? 'archived' : '') as RepositoryDetail['repository']['status'],
          created_at: r?.createdAt ? timestampDate(r.createdAt).toISOString() : '',
        },
        bound_products: (d.boundProducts ?? []).map((p) => ({
          product_id: p.productId ?? '',
          product_name: p.productName ?? '',
          product_status: (p.productStatus === 1 ? 'active' : p.productStatus === 2 ? 'archived' : 'active') as RepositoryDetail['bound_products'][number]['product_status'],
        })),
        mapped_modules: (d.mappedModules ?? []).map((m) => ({
          module_id: m.moduleId ?? '',
          module_name: m.moduleName ?? '',
          module_status: (m.moduleStatus === 1 ? 'active' : m.moduleStatus === 2 ? 'archived' : 'active') as RepositoryDetail['mapped_modules'][number]['module_status'],
        })),
      }
    },
    enabled: !!repositoryId,
  })
}