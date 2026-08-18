/**
 * useBindStandard — Standard Bind 的固定 mutation 承接位。
 *
 * phase14-05 §"切片结构必须冻结"（project_rules §2.5）：
 *   - 承接 BindStandard 的表单值 → pb 请求组装（枚举映射）、
 *     错误归一化与 query 失效
 *   - invalid_argument（含重复绑定 "already bound"）经归一化 message 由页面行内回显
 *
 * Query 失效（失效矩阵逐字）：
 *   - 成功后失效 standard-detail（当前 id）
 *   - 成功后失效 repository-standards 前缀（前缀失效全部 repositoryId）
 */
import { useMutation, useQueryClient, type UseMutationResult } from '@tanstack/react-query'
import { ConnectError } from '@connectrpc/connect'
import { standardClient } from '../data/connect-client'
import { pbToBinding, bindingTargetTypeToPb, bindingRoleToPb } from '../types'
import type { StandardBinding, BindStandardInput } from '../types'

/** 错误归一化：Connect 错误提取原始 message（去除 code 前缀），供页面行内回显 */
function normalizeError(err: unknown): Error {
  if (err instanceof ConnectError) {
    return new Error(err.rawMessage || err.message)
  }
  if (err instanceof Error) {
    return err
  }
  return new Error('Standard 绑定失败，请稍后重试')
}

export type UseBindStandard = UseMutationResult<StandardBinding, Error, BindStandardInput, unknown>

export function useBindStandard(): UseBindStandard {
  const queryClient = useQueryClient()

  return useMutation<StandardBinding, Error, BindStandardInput, unknown>({
    mutationFn: async (input: BindStandardInput): Promise<StandardBinding> => {
      try {
        const res = await standardClient.bindStandard({
          standardId: input.standard_id,
          targetType: bindingTargetTypeToPb(input.form.target_type),
          targetId: input.form.target_id,
          role: bindingRoleToPb(input.form.role),
          note: input.form.note ?? '',
        })
        const binding = res.binding
        if (!binding) throw new Error('Standard 绑定失败：未返回绑定数据')
        return pbToBinding(binding)
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
