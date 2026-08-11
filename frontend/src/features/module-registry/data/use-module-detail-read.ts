/**
 * useModuleDetailRead — Module Detail 只读 query owner。
 *
 * phase07-10 §5.4：query 层纯只读，唯一 owner。
 * 替换 module-registry/data/api-adapter.ts 的 fetchModuleDetail。
 */
import { useQuery, type UseQueryResult } from '@tanstack/react-query'
import { timestampDate } from '@bufbuild/protobuf/wkt'
import { moduleRegistryClient } from './connect-client'
import type { ModuleDetail } from '../types'

export type UseModuleDetailRead = UseQueryResult<ModuleDetail, Error>

export function useModuleDetailRead(moduleId: string): UseModuleDetailRead {
  return useQuery<ModuleDetail, Error>({
    queryKey: ['module-detail', moduleId],
    queryFn: async () => {
      const res = await moduleRegistryClient.getModuleDetail({ moduleId })
      const d = res.moduleDetail
      if (!d) throw new Error('module not found')
      return {
        module: {
          id: d.module?.id ?? '',
          name: d.module?.name ?? '',
          description: d.module?.description ?? '',
          status: (d.module?.status === 1 ? 'active' : d.module?.status === 2 ? 'archived' : '') as ModuleDetail['module']['status'],
          created_at: d.module?.createdAt ? timestampDate(d.module.createdAt).toISOString() : '',
        },
        releases: (d.releases ?? []).map((r) => ({
          id: r.id ?? '',
          module_id: r.moduleId ?? '',
          version: r.version ?? '',
          status: (r.status === 1 ? 'active' : r.status === 2 ? 'archived' : '') as ModuleDetail['releases'][number]['status'],
          released_at: r.releasedAt ? timestampDate(r.releasedAt).toISOString() : '',
        })),
        product_bindings: (d.productBindings ?? []).map((b) => ({
          product_id: b.productId ?? '',
          product_name: b.productName ?? '',
        })),
        repository_mappings: (d.repositoryMappings ?? []).map((m) => ({
          repository_id: m.repositoryId ?? '',
          repository_name: m.repositoryName ?? '',
        })),
        decision_links: (d.decisionLinks ?? []).map((l) => ({
          decision_id: l.decisionId ?? '',
          decision_title: l.decisionTitle ?? '',
        })),
      }
    },
    enabled: !!moduleId,
  })
}