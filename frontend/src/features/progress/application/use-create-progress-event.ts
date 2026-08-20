/**
 * useCreateProgressEvent — Progress Create 的固定 mutation 承接位。
 *
 * phase15-05 §"切片结构必须冻结"（project_rules §2.5）：
 *   - 承接 CreateProgressEvent 的表单值 → pb 请求组装（枚举映射 +
 *     occurred_at 本地解析转 UTC Timestamp + 可空文本空串直传）、
 *     错误归一化与 query 失效；组件不得内联第二套 useMutation
 *   - source 不设置 → 后端归一 manual（裁决⑧；请求体不含 source 字段）
 *   - 暴露 onSuccess 回调位，表单在此承接成功回流（重置 + event_kind 记忆写入）
 *
 * Query 失效（失效矩阵逐字）：
 *   - ['progress-events', repositoryId] 前缀（覆盖全部过滤变体）
 *   - ['repository-progress', repositoryId]（当前卡派生摘要）
 */
import { useMutation, useQueryClient, type UseMutationResult } from '@tanstack/react-query'
import { ConnectError } from '@connectrpc/connect'
import { progressClient } from '../data/connect-client'
import {
  pbToProgressEvent,
  workflowTypeToPb,
  eventKindToPb,
  datetimeLocalToPbTimestamp,
} from '../types'
import type { ProgressEvent, CreateProgressEventInput } from '../types'

/** 错误归一化：Connect 错误提取原始 message（去除 code 前缀），供表单行内回显 */
function normalizeError(err: unknown): Error {
  if (err instanceof ConnectError) {
    return new Error(err.rawMessage || err.message)
  }
  if (err instanceof Error) {
    return err
  }
  return new Error('推进事件创建失败，请稍后重试')
}

export type UseCreateProgressEvent = UseMutationResult<
  ProgressEvent,
  Error,
  CreateProgressEventInput,
  unknown
>

export function useCreateProgressEvent(onSuccess?: (event: ProgressEvent) => void): UseCreateProgressEvent {
  const queryClient = useQueryClient()

  return useMutation<ProgressEvent, Error, CreateProgressEventInput, unknown>({
    mutationFn: async (input: CreateProgressEventInput): Promise<ProgressEvent> => {
      try {
        const res = await progressClient.createProgressEvent({
          repositoryId: input.repository_id,
          workflowType: workflowTypeToPb(input.workflow_type),
          eventKind: eventKindToPb(input.event_kind),
          taskKey: input.task_key,
          title: input.title,
          detail: input.detail,
          evidenceRef: input.evidence_ref,
          // DP-3：datetime-local 按浏览器本地时区语义解析 → UTC pb Timestamp
          occurredAt: datetimeLocalToPbTimestamp(input.occurred_at),
          // source 不设置（裁决⑧）→ 后端归一 manual
        })
        const event = pbToProgressEvent(res.event)
        if (!event) throw new Error('推进事件创建失败：未返回事件数据')
        return event
      } catch (err) {
        throw normalizeError(err)
      }
    },
    onSuccess: (event) => {
      // 失效矩阵：事件流前缀失效（覆盖全部过滤变体）+ 当前卡派生摘要
      queryClient.invalidateQueries({ queryKey: ['progress-events', event.repository_id] })
      queryClient.invalidateQueries({ queryKey: ['repository-progress', event.repository_id] })
      onSuccess?.(event)
    },
  })
}
