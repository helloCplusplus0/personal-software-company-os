import { useQueries } from '@tanstack/react-query'
import { moduleRegistryClient } from './connect-client'

export interface ModuleRelatedDecisionLink {
  decision_id: string
  decision_title: string
}

/**
 * 复用 Module Detail 既有读模型，按一组 moduleId 聚合相关 decision links。
 * 当前用于 Product / Repository Detail 补齐“相关决策”展示，不新增第二套后端合同。
 */
export function useModuleDecisionLinksByModuleIds(moduleIds: string[]) {
  const normalizedModuleIds = Array.isArray(moduleIds) ? moduleIds : []
  const uniqueModuleIds = Array.from(new Set(normalizedModuleIds.filter(Boolean)))

  const queries = useQueries({
    queries: uniqueModuleIds.map((moduleId) => ({
      queryKey: ['module-detail', moduleId],
      queryFn: async (): Promise<ModuleRelatedDecisionLink[]> => {
        const res = await moduleRegistryClient.getModuleDetail({ moduleId })
        const detail = res.moduleDetail
        if (!detail) {
          return []
        }

        return (detail.decisionLinks ?? []).map((link) => ({
          decision_id: link.decisionId ?? '',
          decision_title: link.decisionTitle ?? '',
        }))
      },
      enabled: Boolean(moduleId),
    })),
  })

  const decisionMap = new Map<string, ModuleRelatedDecisionLink>()
  for (const query of Array.isArray(queries) ? queries : []) {
    const links = Array.isArray(query.data) ? query.data : []
    for (const link of links) {
      if (!link.decision_id || decisionMap.has(link.decision_id)) {
        continue
      }
      decisionMap.set(link.decision_id, link)
    }
  }

  return {
    decisionLinks: Array.from(decisionMap.values()),
    isLoading: queries.some((query) => query.isLoading),
    isError: queries.some((query) => query.isError),
  }
}
