/**
 * useProductListRead — Product List 只读 query owner。
 *
 * phase07-10 §5.4：query 层纯只读，唯一 owner。
 * 替换 product-registry/data/api-adapter.ts 的 fetchProductList。
 */
import { useQuery, type UseQueryResult } from '@tanstack/react-query'
import { timestampDate } from '@bufbuild/protobuf/wkt'
import { productRegistryClient } from './connect-client'
import type { ProductListItem, ProductListSearch, ProductStatus } from '../types'

export type UseProductListRead = UseQueryResult<ProductListItem[], Error>

export function useProductListRead(search: ProductListSearch): UseProductListRead {
  return useQuery<ProductListItem[], Error>({
    queryKey: ['product-list', search],
    queryFn: async (): Promise<ProductListItem[]> => {
      const res = await productRegistryClient.listProducts({
        queryText: search.queryText ?? '',
        statusFilter: mapStatusToProto(search.statusFilter ?? 'all'),
      })
      return (res.products ?? []).map((p) => ({
        id: p.id ?? '',
        name: p.name ?? '',
        description: p.description ?? '',
        status: mapProtoToStatus(p.status) as ProductStatus,
        created_at: p.createdAt ? timestampDate(p.createdAt).toISOString() : '',
        module_bind_count: p.moduleBindCount ?? 0,
        repository_bind_count: p.repositoryBindCount ?? 0,
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