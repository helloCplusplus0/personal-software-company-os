/**
 * useProductModuleCandidatesRead — Product Module Candidates 只读 query owner。
 *
 * phase07-10 §5.4：query 层纯只读，唯一 owner。
 * 替换 product-registry/data/api-adapter.ts 的 fetchProductModuleCandidates。
 */
import { useQuery, type UseQueryResult } from '@tanstack/react-query'
import { productRegistryClient } from './connect-client'
import type { ProductModuleCandidate } from '../types'

export type UseProductModuleCandidatesRead = UseQueryResult<ProductModuleCandidate[], Error>

export function useProductModuleCandidatesRead(productId: string): UseProductModuleCandidatesRead {
  return useQuery<ProductModuleCandidate[], Error>({
    queryKey: ['product-module-candidates', productId],
    queryFn: async () => {
      const res = await productRegistryClient.listProductModuleCandidates({ productId })
      return (res.candidates ?? []).map((c) => ({
        module_id: c.moduleId ?? '',
        module_name: c.moduleName ?? '',
        module_status: (c.moduleStatus === 1 ? 'active' : c.moduleStatus === 2 ? 'archived' : 'active') as ProductModuleCandidate['module_status'],
      }))
    },
    enabled: !!productId,
  })
}