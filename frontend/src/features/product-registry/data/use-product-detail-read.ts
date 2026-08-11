/**
 * useProductDetailRead — Product Detail 只读 query owner。
 *
 * phase07-10 §5.4：query 层纯只读，唯一 owner。
 * 替换 product-registry/data/api-adapter.ts 的 fetchProductDetail。
 */
import { useQuery, type UseQueryResult } from '@tanstack/react-query'
import { timestampDate } from '@bufbuild/protobuf/wkt'
import { productRegistryClient } from './connect-client'
import type { ProductDetail } from '../types'

export type UseProductDetailRead = UseQueryResult<ProductDetail, Error>

export function useProductDetailRead(productId: string): UseProductDetailRead {
  return useQuery<ProductDetail, Error>({
    queryKey: ['product-detail', productId],
    queryFn: async () => {
      const res = await productRegistryClient.getProductDetail({ productId })
      const d = res.productDetail
      if (!d) throw new Error('product not found')
      const p = d.product
      return {
        product: {
          id: p?.id ?? '',
          name: p?.name ?? '',
          description: p?.description ?? '',
          status: (p?.status === 1 ? 'active' : p?.status === 2 ? 'archived' : '') as ProductDetail['product']['status'],
          created_at: p?.createdAt ? timestampDate(p.createdAt).toISOString() : '',
        },
        bound_modules: (d.boundModules ?? []).map((m) => ({
          module_id: m.moduleId ?? '',
          module_name: m.moduleName ?? '',
          module_status: (m.moduleStatus === 1 ? 'active' : m.moduleStatus === 2 ? 'archived' : 'active') as ProductDetail['bound_modules'][number]['module_status'],
        })),
        bound_repositories: (d.boundRepositories ?? []).map((r) => ({
          repository_id: r.repositoryId ?? '',
          repository_name: r.repositoryName ?? '',
          provider: r.provider ?? '',
          repository_status: (r.repositoryStatus === 1 ? 'active' : r.repositoryStatus === 2 ? 'archived' : 'active') as ProductDetail['bound_repositories'][number]['repository_status'],
        })),
      }
    },
    enabled: !!productId,
  })
}