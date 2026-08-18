/**
 * useDeleteStandard — Standard Delete 的固定 mutation 承接位。
 *
 * phase14-05 §"切片结构必须冻结"（project_rules §2.5）：
 *   - 承接 DeleteStandard（后端 CASCADE 连带删除 bindings / revisions）、
 *     错误归一化与 query 失效
 *   - ACTIVE 态删除按钮禁用（先经 Update 置 RETIRED）属页面交互约束，不在本层
 *
 * Query 失效（失效矩阵逐字）：
 *   - 成功后失效 standard-list
 */
import { useMutation, useQueryClient, type UseMutationResult } from '@tanstack/react-query'
import { ConnectError } from '@connectrpc/connect'
import { standardClient } from '../data/connect-client'

/** 错误归一化：Connect 错误提取原始 message（去除 code 前缀），供页面行内回显 */
function normalizeError(err: unknown): Error {
  if (err instanceof ConnectError) {
    return new Error(err.rawMessage || err.message)
  }
  if (err instanceof Error) {
    return err
  }
  return new Error('Standard 删除失败，请稍后重试')
}

export type UseDeleteStandard = UseMutationResult<void, Error, { standardId: string }, unknown>

export function useDeleteStandard(): UseDeleteStandard {
  const queryClient = useQueryClient()

  return useMutation<void, Error, { standardId: string }, unknown>({
    mutationFn: async ({ standardId }: { standardId: string }): Promise<void> => {
      try {
        await standardClient.deleteStandard({ standardId })
      } catch (err) {
        throw normalizeError(err)
      }
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['standard-list'] })
    },
  })
}
