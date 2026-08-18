/**
 * useUnbindStandard — Standard Unbind 的固定 mutation 承接位。
 *
 * phase14-05 §"切片结构必须冻结"（project_rules §2.5）：
 *   - 承接 UnbindStandard（四元组定位，note 不参与）、错误归一化与 query 失效
 *
 * Query 失效（与 bind 相同矩阵，逐字）：
 *   - 成功后失效 standard-detail（当前 id）
 *   - 成功后失效 repository-standards 前缀（前缀失效全部 repositoryId）
 */
import { useMutation, useQueryClient, type UseMutationResult } from '@tanstack/react-query'
import { ConnectError } from '@connectrpc/connect'
import { standardClient } from '../data/connect-client'
import { bindingTargetTypeToPb, bindingRoleToPb } from '../types'
import type { UnbindStandardInput } from '../types'

/** 错误归一化：Connect 错误提取原始 message（去除 code 前缀），供页面行内回显 */
function normalizeError(err: unknown): Error {
  if (err instanceof ConnectError) {
    return new Error(err.rawMessage || err.message)
  }
  if (err instanceof Error) {
    return err
  }
  return new Error('Standard 解绑失败，请稍后重试')
}

export type UseUnbindStandard = UseMutationResult<void, Error, UnbindStandardInput, unknown>

export function useUnbindStandard(): UseUnbindStandard {
  const queryClient = useQueryClient()

  return useMutation<void, Error, UnbindStandardInput, unknown>({
    mutationFn: async (input: UnbindStandardInput): Promise<void> => {
      try {
        await standardClient.unbindStandard({
          standardId: input.standard_id,
          targetType: bindingTargetTypeToPb(input.target_type),
          targetId: input.target_id,
          role: bindingRoleToPb(input.role),
        })
      } catch (err) {
        throw normalizeError(err)
      }
    },
    onSuccess: (_data, variables) => {
      queryClient.invalidateQueries({ queryKey: ['standard-detail', variables.standard_id] })
      queryClient.invalidateQueries({ queryKey: ['repository-standards'] })
    },
  })
}
