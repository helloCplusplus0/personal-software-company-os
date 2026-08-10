/**
 * useCreateDraftModule — Module Create 的固定 mutation 承接位
 *
 * phase06-07 §"四类 canonical create 必须收敛到固定 mutation 承接位"
 * phase06-15 §"application owner 物理落点"
 *
 * 职责：
 *   - 承接 Module create 的默认补值、错误归一化与 query 失效
 *   - 导出统一的 mutate / mutateAsync / isPending / isError / error / data
 *
 * 默认补值（phase06 draft-first 最小输入口径）：
 *   - status 未提供时默认 'active'
 *   - description 不再阻断提交（phase06-07 字段级放宽）
 *
 * Query 失效：
 *   - 成功后失效 module-list
 *   - 成功后失效 onboarding-state（首轮录入状态与 Dashboard CTA 重新派生）
 */
import { useMutation, useQueryClient, type UseMutationResult } from '@tanstack/react-query'
import { createModule, ApiError } from '../data/api-adapter'
import type { CreateModuleInput, Module } from '../types'
import { ONBOARDING_STATE_QUERY_KEY } from '@/features/onboarding/data/use-onboarding-read'

function applyDefaults(input: CreateModuleInput): CreateModuleInput {
  return {
    name: input.name,
    description: input.description,
    status: input.status ?? 'active',
  }
}

function normalizeError(err: unknown): Error {
  if (err instanceof ApiError) {
    return new Error(err.message)
  }
  if (err instanceof Error) {
    return err
  }
  return new Error('模块创建失败，请稍后重试')
}

export type UseCreateDraftModule = UseMutationResult<Module, Error, CreateModuleInput, unknown>

export function useCreateDraftModule(): UseCreateDraftModule {
  const queryClient = useQueryClient()

  return useMutation<Module, Error, CreateModuleInput, unknown>({
    mutationFn: async (input: CreateModuleInput) => {
      try {
        return await createModule(applyDefaults(input))
      } catch (err) {
        throw normalizeError(err)
      }
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['module-list'] })
        queryClient.invalidateQueries({ queryKey: ONBOARDING_STATE_QUERY_KEY })
    },
  })
}
