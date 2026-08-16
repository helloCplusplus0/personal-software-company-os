/**
 * useProjectContextRead — 项目上下文只读 query owner。
 *
 * phase12-09：唯一跨切片共享 GetProjectContext query options。
 * 各页面通过本 hook 消费共享 project-context 事实源，不得各自直接创建 transport 或 client。
 *
 * 输入：repository_id（唯一结构化输入锚点）
 * 输出：UseQueryResult<ProjectContext>
 * 缓存键：['project-context', repositoryId]
 * 失败语义：以 UseQueryResult.error 暴露，不回退到页面级错误
 */
import { useQuery, type UseQueryResult } from '@tanstack/react-query'
import { projectContextClient } from './connect-client'
import type {
  RepositorySummary,
  ProductSummary,
  ModuleSummary,
  DecisionSummary,
  RuleEntry,
  PhaseEntry,
  BoundaryEntry,
} from '@/gen/proto/psco/project_context/v1/project_context_pb'

export interface ProjectContext {
  repository: RepositorySummary | undefined
  product: ProductSummary | undefined
  modules: ModuleSummary[]
  decisions: DecisionSummary[]
  rules: RuleEntry[]
  phases: PhaseEntry[]
  boundaries: BoundaryEntry[]
}

export type UseProjectContextRead = UseQueryResult<ProjectContext, Error>

export function useProjectContextRead(repositoryId: string): UseProjectContextRead {
  return useQuery<ProjectContext, Error>({
    queryKey: ['project-context', repositoryId],
    queryFn: async () => {
      const res = await projectContextClient.getProjectContext({ repositoryId })
      return {
        repository: res.repository,
        product: res.product,
        modules: res.modules,
        decisions: res.decisions,
        rules: res.rules,
        phases: res.phases,
        boundaries: res.boundaries,
      }
    },
    enabled: !!repositoryId,
  })
}