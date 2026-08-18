/**
 * useUpdateStandard — Standard Update 的固定 mutation 承接位。
 *
 * phase14-05 §"切片结构必须冻结"（project_rules §2.5）：
 *   - 承接 UpdateStandard 的表单值 → pb 请求组装（optional 字段未设置不变更、
 *     整树原子替换 + change_summary 必填）、错误归一化与 query 失效
 *
 * Query 失效（失效矩阵逐字）：
 *   - 成功后失效 standard-list 与 standard-detail（当前 id）
 */
import { useMutation, useQueryClient, type UseMutationResult } from '@tanstack/react-query'
import { ConnectError } from '@connectrpc/connect'
import { standardClient } from '../data/connect-client'
import { pbToStandard, treeToPb, standardStatusToPb } from '../types'
import type { Standard, UpdateStandardInput } from '../types'

/** 错误归一化：Connect 错误提取原始 message（去除 code 前缀），供页面行内回显 */
function normalizeError(err: unknown): Error {
  if (err instanceof ConnectError) {
    return new Error(err.rawMessage || err.message)
  }
  if (err instanceof Error) {
    return err
  }
  return new Error('Standard 更新失败，请稍后重试')
}

export type UseUpdateStandard = UseMutationResult<Standard, Error, UpdateStandardInput, unknown>

export function useUpdateStandard(): UseUpdateStandard {
  const queryClient = useQueryClient()

  return useMutation<Standard, Error, UpdateStandardInput, unknown>({
    mutationFn: async (input: UpdateStandardInput): Promise<Standard> => {
      try {
        const res = await standardClient.updateStandard({
          standardId: input.standard_id,
          // optional 三字段：未设置（undefined）即不变更（phase14-04 合同）
          name: input.name,
          description: input.description,
          status: input.status ? standardStatusToPb(input.status) : undefined,
          directoryTree: treeToPb(input.directory_tree),
          changeSummary: input.change_summary,
        })
        const standard = pbToStandard(res.standard)
        if (!standard) throw new Error('Standard 更新失败：未返回标准数据')
        return standard
      } catch (err) {
        throw normalizeError(err)
      }
    },
    onSuccess: (_data, variables) => {
      queryClient.invalidateQueries({ queryKey: ['standard-list'] })
      queryClient.invalidateQueries({ queryKey: ['standard-detail', variables.standard_id] })
    },
  })
}
