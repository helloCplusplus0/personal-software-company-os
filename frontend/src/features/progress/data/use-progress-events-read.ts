/**
 * useProgressEventsRead — Repository 推进事件流只读 query owner。
 *
 * phase15-05 §"切片结构必须冻结"：
 *   - ListProgressEvents → ProgressEvent[]（结果即后端三键链倒序，前端不重排序）
 *   - filter：'all' 不传过滤（UNSPECIFIED = 三轨全量，phase15-04 合同决策 2）；
 *     'phase' / 'audit' / 'fix' 传对应枚举
 *   - DP-1 封死：结果仅用于时间轴原始事件展示，不得参与当前卡派生计算
 */
import { useQuery, type UseQueryResult } from '@tanstack/react-query'
import { progressClient } from './connect-client'
import { pbToProgressEvent, workflowTypeToPb } from '../types'
import type { ProgressEvent } from '../types'

/** 三轨过滤参数 — 'all' = 三轨全量（不过滤） */
export type ProgressFilter = 'all' | 'phase' | 'audit' | 'fix'

export type UseProgressEventsRead = UseQueryResult<ProgressEvent[], Error>

export function useProgressEventsRead(
  repositoryId: string,
  filter: ProgressFilter,
): UseProgressEventsRead {
  return useQuery<ProgressEvent[], Error>({
    queryKey: ['progress-events', repositoryId, filter],
    queryFn: async (): Promise<ProgressEvent[]> => {
      const res = await progressClient.listProgressEvents({
        repositoryId,
        // 'all' 不传过滤参数（UNSPECIFIED = 三轨全量）
        workflowType: filter === 'all' ? undefined : workflowTypeToPb(filter),
      })
      return (res.events ?? [])
        .map((e) => pbToProgressEvent(e))
        .filter((e): e is ProgressEvent => e !== null)
    },
    enabled: !!repositoryId,
  })
}
