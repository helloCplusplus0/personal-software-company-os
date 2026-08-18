/**
 * useCreateStandard — Standard Create 的固定 mutation 承接位。
 *
 * phase14-05 §"切片结构必须冻结"（project_rules §2.5）：
 *   - 承接 CreateStandard 的表单值 → pb 请求组装（枚举映射 + 整树组装）、
 *     错误归一化与 query 失效；页面与组件不得内联第二套 useMutation
 *   - 暴露 onSuccess 回调位，页面在此承接成功回流（导航到详情页）
 *
 * Query 失效（失效矩阵逐字）：
 *   - 成功后失效 standard-list
 */
import { useMutation, useQueryClient, type UseMutationResult } from '@tanstack/react-query'
import { ConnectError } from '@connectrpc/connect'
import { standardClient } from '../data/connect-client'
import { pbToStandard, treeToPb, standardStatusToPb } from '../types'
import type { Standard, CreateStandardInput } from '../types'

/** 错误归一化：Connect 错误提取原始 message（去除 code 前缀），供页面行内回显 */
function normalizeError(err: unknown): Error {
  if (err instanceof ConnectError) {
    return new Error(err.rawMessage || err.message)
  }
  if (err instanceof Error) {
    return err
  }
  return new Error('Standard 创建失败，请稍后重试')
}

export type UseCreateStandard = UseMutationResult<Standard, Error, CreateStandardInput, unknown>

export function useCreateStandard(onSuccess?: (standard: Standard) => void): UseCreateStandard {
  const queryClient = useQueryClient()

  return useMutation<Standard, Error, CreateStandardInput, unknown>({
    mutationFn: async (input: CreateStandardInput): Promise<Standard> => {
      try {
        const res = await standardClient.createStandard({
          name: input.name,
          description: input.description ?? '',
          directoryTree: treeToPb(input.directory_tree),
          // status 未提供时不设置，由后端默认 DRAFT（phase14-04 合同）
          status: input.status ? standardStatusToPb(input.status) : undefined,
        })
        const standard = pbToStandard(res.standard)
        if (!standard) throw new Error('Standard 创建失败：未返回标准数据')
        return standard
      } catch (err) {
        throw normalizeError(err)
      }
    },
    onSuccess: (standard) => {
      queryClient.invalidateQueries({ queryKey: ['standard-list'] })
      onSuccess?.(standard)
    },
  })
}
