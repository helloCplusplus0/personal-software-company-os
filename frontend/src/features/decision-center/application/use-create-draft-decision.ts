/**
 * useCreateDraftDecision — Decision Create 的固定 mutation 承接位
 *
 * phase06-07 §"四类 canonical create 必须收敛到固定 mutation 承接位"
 * phase06-15 §"application owner 物理落点"
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
import { createDecision, ApiError } from '../data/api-adapter'
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
  if (err instanceof ApiError) {
    return new Error(err.message)
  }
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
        return await createDecision(applyDefaults(input))
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
