/**
 * useDeleteProgressEvent — Progress Delete 的固定 mutation 承接位。
 *
 * phase15-05 §"时间轴列表与误录删除交互规格必须冻结"：
 *   - 误录删除（append-only 语义，裁决⑨：整条删除是唯一修正路径，无软删除）
 *   - mutate 参数 { id, repositoryId }：id 定位删除目标；repositoryId 用于失效矩阵
 *     （DeleteProgressEvent 请求仅需 id，失效需要 repositoryId 定位缓存键）
 *
 * Query 失效（失效矩阵逐字，与 Create 相同）：
 *   - ['progress-events', repositoryId] 前缀（覆盖全部过滤变体）
 *   - ['repository-progress', repositoryId]（删除同样改变派生摘要）
 */
import { useMutation, useQueryClient, type UseMutationResult } from '@tanstack/react-query'
import { ConnectError } from '@connectrpc/connect'
import { progressClient } from '../data/connect-client'

/** 错误归一化：Connect 错误提取原始 message（去除 code 前缀），供列表行内回显 */
function normalizeError(err: unknown): Error {
  if (err instanceof ConnectError) {
    return new Error(err.rawMessage || err.message)
  }
  if (err instanceof Error) {
    return err
  }
  return new Error('推进事件删除失败，请稍后重试')
}

export type UseDeleteProgressEvent = UseMutationResult<
  void,
  Error,
  { id: string; repositoryId: string },
  unknown
>

export function useDeleteProgressEvent(): UseDeleteProgressEvent {
  const queryClient = useQueryClient()

  return useMutation<void, Error, { id: string; repositoryId: string }, unknown>({
    mutationFn: async ({ id }: { id: string; repositoryId: string }): Promise<void> => {
      try {
        await progressClient.deleteProgressEvent({ id })
      } catch (err) {
        throw normalizeError(err)
      }
    },
    onSuccess: (_data, variables) => {
      // 失效矩阵：删除同样改变事件流与派生摘要（当前卡同步刷新）
      queryClient.invalidateQueries({ queryKey: ['progress-events', variables.repositoryId] })
      queryClient.invalidateQueries({ queryKey: ['repository-progress', variables.repositoryId] })
    },
  })
}
