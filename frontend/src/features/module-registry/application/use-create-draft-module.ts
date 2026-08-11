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
import { timestampDate } from '@bufbuild/protobuf/wkt'
import { moduleRegistryClient } from '../data/connect-client'
import type { CreateModuleInput, Module } from '../types'
import { ONBOARDING_STATE_QUERY_KEY } from '@/features/onboarding/data/use-onboarding-read'

function applyDefaults(input: CreateModuleInput): CreateModuleInput {
  return {
    name: input.name,
    description: input.description,
    status: input.status ?? 'active',
  }
}

export type UseCreateDraftModule = UseMutationResult<Module, Error, CreateModuleInput, unknown>

export function useCreateDraftModule(): UseCreateDraftModule {
  const queryClient = useQueryClient()

  return useMutation<Module, Error, CreateModuleInput, unknown>({
    mutationFn: async (input: CreateModuleInput) => {
      const defaults = applyDefaults(input)
      const res = await moduleRegistryClient.createModule({
        name: defaults.name,
        description: defaults.description ?? '',
        status: (defaults.status === 'active' ? 1 : defaults.status === 'archived' ? 2 : 0),
      })
      const m = res.module
      if (!m) throw new Error('模块创建失败：未返回模块数据')
      return {
        id: m.id ?? '',
        name: m.name ?? '',
        description: m.description ?? '',
        status: (m.status === 1 ? 'active' : m.status === 2 ? 'archived' : '') as Module['status'],
        created_at: m.createdAt ? timestampDate(m.createdAt).toISOString() : '',
      }
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['module-list'] })
        queryClient.invalidateQueries({ queryKey: ONBOARDING_STATE_QUERY_KEY })
    },
  })
}
