/**
 * useUpdateGovernanceProfile — 治理画像保存的固定 mutation 承接位。
 *
 * phase13-09 冻结：
 *   - 治理画像唯一正式写路径 owner，页面 / 表单 / 展示组件不得内联 useMutation
 *   - 保存成功后精准刷新当前 repository_id 对应的治理画像读取结果
 *   - 提交负载只允许 template_source / canonical_root_files[] / global_asset_bindings[]；
 *     docs_workflow_layout / track_type / current_phase_* 都不进入提交负载
 *
 * 失败语义：错误经 UseMutationResult.error 暴露，由表单层保持可重试与错误可见。
 * 传输层转换：draft plain 结构在唯一调用点以 PartialMessage 语义传给 Connect client。
 */
import { useMutation, useQueryClient, type UseMutationResult } from '@tanstack/react-query'
import { governanceProfileClient } from '../data/connect-client'
import type { GovernanceProfile } from '@/gen/proto/psco/governance_profile/v1/governance_profile_pb'
import type { GovernanceProfileSaveDraft } from '../types'

export interface UpdateGovernanceProfileVariables {
  repositoryId: string
  request: GovernanceProfileSaveDraft
}

export function useUpdateGovernanceProfile(): UseMutationResult<
  GovernanceProfile,
  Error,
  UpdateGovernanceProfileVariables,
  unknown
> {
  const queryClient = useQueryClient()
  return useMutation<GovernanceProfile, Error, UpdateGovernanceProfileVariables, unknown>({
    mutationFn: async ({ repositoryId, request }) => {
      const res = await governanceProfileClient.updateGovernanceProfile({
        repositoryId,
        templateSource: request.templateSource,
        canonicalRootFiles: request.canonicalRootFiles,
        globalAssetBindings: request.globalAssetBindings,
      })
      if (!res.profile) {
        throw new Error('治理画像保存响应缺少 profile 字段')
      }
      return res.profile
    },
    onSuccess: (_data, variables) => {
      queryClient.invalidateQueries({ queryKey: ['governance-profile', variables.repositoryId] })
    },
  })
}
