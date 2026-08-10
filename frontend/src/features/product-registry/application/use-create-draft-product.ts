/**
 * useCreateDraftProduct — Product Create 的固定 mutation 承接位
 *
 * phase06-07 §"四类 canonical create 必须收敛到固定 mutation 承接位"
 * phase06-15 §"application owner 物理落点"
 *
 * 职责：
 *   - 承接 Product create 的默认补值、错误归一化与 query 失效
 *   - 导出统一的 mutate / mutateAsync / isPending / isError / error / data
 *   - 页面、表单与展示组件不得自行拼装第二套正式 create 语义
 *
 * 默认补值（phase06 draft-first 最小输入口径）：
 *   - status 未提供时默认 'active'
 *   - description 不再阻断提交（phase06-07 字段级放宽）
 *
 * 错误归一化：
 *   - ApiError 的 message 直接承接后端错误语义
 *   - 网络错误 fallback 为可读提示
 *
 * Query 失效：
 *   - 成功后失效 product-list（列表读取会重新拉取）
 *   - 成功后失效 onboarding-state（首轮录入状态与 Dashboard CTA 重新派生）
 */
import { useMutation, useQueryClient, type UseMutationResult } from '@tanstack/react-query'
import { createProduct, ApiError } from '../data/api-adapter'
import type { CreateProductInput, CreateProductResponse } from '../types'
import { ONBOARDING_STATE_QUERY_KEY } from '@/features/onboarding/data/use-onboarding-read'

/**
 * 对输入做默认补值，确保后端收到合法的最小 draft-first payload。
 */
function applyDefaults(input: CreateProductInput): CreateProductInput {
  return {
    name: input.name,
    description: input.description,
    status: input.status ?? 'active',
  }
}

/**
 * 归一化错误为可读 message。
 */
function normalizeError(err: unknown): Error {
  if (err instanceof ApiError) {
    return new Error(err.message)
  }
  if (err instanceof Error) {
    return err
  }
  return new Error('产品创建失败，请稍后重试')
}

export type UseCreateDraftProduct = UseMutationResult<
  CreateProductResponse,
  Error,
  CreateProductInput,
  unknown
>

/**
 * useCreateDraftProduct — Product Create 的正式 mutation owner。
 *
 * 使用方：
 *   - OnboardingPage product 步骤
 *   - ProductCreatePage（回收后）
 *
 * 返回统一的 UseMutationResult 接口。
 */
export function useCreateDraftProduct(): UseCreateDraftProduct {
  const queryClient = useQueryClient()

  return useMutation<CreateProductResponse, Error, CreateProductInput, unknown>({
    mutationFn: async (input: CreateProductInput) => {
      try {
        return await createProduct(applyDefaults(input))
      } catch (err) {
        throw normalizeError(err)
      }
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['product-list'] })
        queryClient.invalidateQueries({ queryKey: ONBOARDING_STATE_QUERY_KEY })
    },
  })
}
