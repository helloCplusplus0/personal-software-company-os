/**
 * useRepositoryProgressRead — Repository 进度摘要只读 query owner（DP-1 通道）。
 *
 * phase15-05 §"DP-1 裁决"：
 *   - 调 projectContextClient.getProjectBrief 仅投影 progress 块 → ProgressSummary，
 *     与 agent 消费主路径同源（GetProjectBrief.progress），不新增 RPC、不建第二通道
 *   - 空态恒有值——后端恒构造块；防御 undefined → 零值摘要
 *   - 沿 use-repository-standards-read 的"切片自有 brief 投影 owner"模式
 */
import { useQuery, type UseQueryResult } from '@tanstack/react-query'
import { projectContextClient } from './connect-client'
import { pbToProgressEvent, emptyProgressSummary } from '../types'
import type { ProgressSummary, ProgressEvent } from '../types'

export type UseRepositoryProgressRead = UseQueryResult<ProgressSummary, Error>

export function useRepositoryProgressRead(repositoryId: string): UseRepositoryProgressRead {
  return useQuery<ProgressSummary, Error>({
    queryKey: ['repository-progress', repositoryId],
    queryFn: async (): Promise<ProgressSummary> => {
      const res = await projectContextClient.getProjectBrief({ repositoryId })
      const progress = res.progress
      // 后端恒构造 progress 块；防御 undefined → 零值摘要（DP-1 空态恒有值）
      if (!progress) return emptyProgressSummary()
      return {
        current_phase_key: progress.currentPhaseKey ?? '',
        current_phase_label: progress.currentPhaseLabel ?? '',
        latest_task_completed: progress.latestTaskCompleted
          ? pbToProgressEvent(progress.latestTaskCompleted)
          : null,
        recent_events: (progress.recentEvents ?? [])
          .map((e) => pbToProgressEvent(e))
          .filter((e): e is ProgressEvent => e !== null),
      }
    },
    enabled: !!repositoryId,
  })
}
