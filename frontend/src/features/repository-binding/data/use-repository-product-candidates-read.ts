/**
 * useRepositoryProductCandidatesRead — Repository Product Candidates 只读 query owner。
 *
 * phase07-10 §5.4：query 层纯只读，唯一 owner。
 * 替换 repository-binding/data/api-adapter.ts 的 fetchRepositoryProductCandidates。
 */
import { useQuery, type UseQueryResult } from '@tanstack/react-query'
import { repositoryBindingClient } from './connect-client'
import type { RepositoryProductCandidate } from '../types'

export type UseRepositoryProductCandidatesRead = UseQueryResult<RepositoryProductCandidate[], Error>

export function useRepositoryProductCandidatesRead(repositoryId: string): UseRepositoryProductCandidatesRead {
  return useQuery<RepositoryProductCandidate[], Error>({
    queryKey: ['repository-product-candidates', repositoryId],
    queryFn: async () => {
      const res = await repositoryBindingClient.listRepositoryProductCandidates({ repositoryId })
      return (res.candidates ?? []).map((c) => ({
        product_id: c.productId ?? '',
        product_name: c.productName ?? '',
        product_status: (c.productStatus === 1 ? 'active' : c.productStatus === 2 ? 'archived' : 'active') as RepositoryProductCandidate['product_status'],
      }))
    },
    enabled: !!repositoryId,
  })
}