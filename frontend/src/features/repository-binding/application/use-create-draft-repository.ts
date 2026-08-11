/**
 * useCreateDraftRepository — Repository Create 的固定 mutation 承接位。
 *
 * phase07-10 §5.5：canonical 写动作单一正式 owner。
 *
 * 职责：
 *   - 承接 Repository create 的默认补值、错误归一化与 query 失效
 *   - 导出统一的 mutate / mutateAsync / isPending / isError / error / data
 *
 * 默认补值（phase06 draft-first 最小输入口径）：
 *   - status 未提供时默认 'active'
 *   - provider 不再阻断提交（phase06-07 字段级放宽）
 *
 * Query 失效：
 *   - 成功后失效 repository-list
 *   - 成功后失效 onboarding-state（首轮录入状态与 Dashboard CTA 重新派生）
 */
import { useMutation, useQueryClient, type UseMutationResult } from '@tanstack/react-query'
import { repositoryBindingClient } from '../data/connect-client'
import { ActiveArchivedStatus } from '@/gen/proto/psco/common/v1/common_pb'
import type { CreateRepositoryInput, CreateRepositoryResponse } from '../types'
import { ONBOARDING_STATE_QUERY_KEY } from '@/features/onboarding/data/use-onboarding-read'

function applyDefaults(input: CreateRepositoryInput): CreateRepositoryInput {
  return {
    name: input.name,
    url: input.url,
    provider: input.provider,
    status: input.status ?? 'active',
  }
}

function normalizeError(err: unknown): Error {
  if (err instanceof Error) {
    return err
  }
  return new Error('仓库创建失败，请稍后重试')
}

export type UseCreateDraftRepository = UseMutationResult<
  CreateRepositoryResponse,
  Error,
  CreateRepositoryInput,
  unknown
>

export function useCreateDraftRepository(): UseCreateDraftRepository {
  const queryClient = useQueryClient()

  return useMutation<CreateRepositoryResponse, Error, CreateRepositoryInput, unknown>({
    mutationFn: async (input: CreateRepositoryInput) => {
      try {
        const defaults = applyDefaults(input)
        const res = await repositoryBindingClient.createRepository({
          name: defaults.name,
          url: defaults.url,
          provider: defaults.provider ?? '',
          status: defaults.status === 'active' ? ActiveArchivedStatus.ACTIVE : ActiveArchivedStatus.ARCHIVED,
        })
        return { repository_id: res.repositoryId ?? '' }
      } catch (err) {
        throw normalizeError(err)
      }
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['repository-list'] })
      queryClient.invalidateQueries({ queryKey: ONBOARDING_STATE_QUERY_KEY })
    },
  })
}