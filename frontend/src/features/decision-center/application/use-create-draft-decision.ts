/**
 * useCreateDraftDecision — Decision Create 的固定 mutation 承接位。
 *
 * phase07-10 §5.5：canonical 写动作单一正式 owner。
 *
 * 职责：
 *   - 承接 Decision create 的默认补值、错误归一化与 query 失效
 *   - 导出统一的 mutate / mutateAsync / isPending / isError / error / data
 *
 * 默认补值（phase06 draft-first 最小输入口径）：
 *   - status 未提供时默认 'proposed'（phase06-07 系统补位）
 *   - context / problem 不再前端必填（phase06-07 字段级放宽）
 *   - source_module_id 未提供时默认空字符串
 *
 * Query 失效：
 *   - 成功后失效 decision-list
 *   - 成功后失效 onboarding-state（首轮录入状态与 Dashboard CTA 重新派生）
 */
import { useMutation, useQueryClient, type UseMutationResult } from '@tanstack/react-query'
import { decisionCenterClient } from '../data/connect-client'
import { DecisionStatus } from '@/gen/proto/psco/decision_center/v1/decision_center_pb'
import type { CreateDecisionInput, CreateDecisionResponse } from '../types'
import { ONBOARDING_STATE_QUERY_KEY } from '@/features/onboarding/data/use-onboarding-read'

function applyDefaults(input: CreateDecisionInput): CreateDecisionInput {
  return {
    title: input.title,
    context: input.context,
    problem: input.problem,
    alternatives: input.alternatives,
    choice: input.choice,
    reason: input.reason,
    impact: input.impact,
    status: input.status ?? 'proposed',
    source_module_id: input.source_module_id ?? '',
  }
}

function normalizeError(err: unknown): Error {
  if (err instanceof Error) {
    return err
  }
  return new Error('决策创建失败，请稍后重试')
}

export type UseCreateDraftDecision = UseMutationResult<
  CreateDecisionResponse,
  Error,
  CreateDecisionInput,
  unknown
>

export function useCreateDraftDecision(): UseCreateDraftDecision {
  const queryClient = useQueryClient()

  return useMutation<CreateDecisionResponse, Error, CreateDecisionInput, unknown>({
    mutationFn: async (input: CreateDecisionInput) => {
      try {
        const defaults = applyDefaults(input)
        const res = await decisionCenterClient.createDecision({
          title: defaults.title,
          context: defaults.context ?? '',
          problem: defaults.problem ?? '',
          alternatives: defaults.alternatives ?? [],
          choice: defaults.choice,
          reason: defaults.reason,
          impact: defaults.impact ?? '',
          status: defaults.status === 'proposed' ? DecisionStatus.PROPOSED : defaults.status === 'active' ? DecisionStatus.ACTIVE : DecisionStatus.SUPERSEDED,
          sourceModuleId: defaults.source_module_id ?? '',
        })
        return { decision_id: res.decisionId ?? '' }
      } catch (err) {
        throw normalizeError(err)
      }
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['decision-list'] })
      queryClient.invalidateQueries({ queryKey: ONBOARDING_STATE_QUERY_KEY })
    },
  })
}