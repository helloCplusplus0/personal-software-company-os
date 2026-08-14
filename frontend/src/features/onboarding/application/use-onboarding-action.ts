/**
 * useOnboardingAction — Onboarding 唯一正式写动作承接位
 *
 * phase10-08 §"前端必须落地单一 useOnboardingAction"
 *
 * 职责：
 *   - 复用既有 canonical create / bind owner，不新增第二套 transport
 *   - 为 product / repository / module / decision 四步各自提供单一 action 函数
 *   - 返回统一 success envelope：{ resultKind, nextStep, navigateTo?, params?, search?, successMessage? }
 *   - 成功后失效 ONBOARDING_STATE_QUERY_KEY 与链路读模型 query key
 *
 * 约束：
 *   - 页面层不得直接消费底层 mutation response
 *   - 页面层不得自己决定下一步去哪里
 */
import { useCallback } from 'react'
import { useQueryClient } from '@tanstack/react-query'
import { useCreateDraftProduct } from '@/features/product-registry/application/use-create-draft-product'
import { useCreateDraftRepository } from '@/features/repository-binding/application/use-create-draft-repository'
import { useCreateDraftModule } from '@/features/module-registry/application/use-create-draft-module'
import { useCreateDraftDecision } from '@/features/decision-center/application/use-create-draft-decision'
import { useBindRepositoryToProduct } from '@/features/repository-binding/application/use-bind-repository-to-product'
import { useBindModuleToProduct } from '@/features/product-registry/application/use-bind-module-to-product'
import { ONBOARDING_STATE_QUERY_KEY, ONBOARDING_CHAIN_STATE_QUERY_KEY } from '../data/use-onboarding-read'
import type { OnboardingStep } from '../types'
import { buildOnboardingChainReturnSearch } from '../lib/onboarding-return'
import { onboardingClient } from '../data/connect-client'

// ============================================================================
// 类型定义
// ============================================================================

/** 写动作结果类型 */
export type ActionResultKind = 'advance' | 'handoff' | 'complete'

/** 统一的 success envelope */
export interface OnboardingActionSuccess {
  resultKind: ActionResultKind
  nextStep: OnboardingStep
  navigateTo?: string
  params?: Record<string, string>
  search?: Record<string, unknown>
  successMessage?: string
}

/** 各步骤的输入 */
export interface ProductStepInput {
  name: string
}

export interface RepositoryStepInput {
  name: string
  url: string
}

export interface ModuleStepInput {
  name: string
}

export interface DecisionStepInput {
  title: string
  choice: string
  reason: string
}

/** useOnboardingAction 的返回值 */
export interface UseOnboardingAction {
  /** product 步骤：创建 Product → 冻结锚点 → 前进到 repository */
  submitProduct: (input: ProductStepInput) => Promise<OnboardingActionSuccess>
  /** repository 步骤：创建 Repository → 尝试绑定到 Product */
  submitRepository: (input: RepositoryStepInput, currentProductId: string) => Promise<OnboardingActionSuccess>
  /** module 步骤：创建 Module → 尝试绑定到 Product */
  submitModule: (input: ModuleStepInput, currentProductId: string) => Promise<OnboardingActionSuccess>
  /** decision 步骤：创建 Decision */
  submitDecision: (input: DecisionStepInput) => Promise<OnboardingActionSuccess>
  /** 各步骤的 mutation 状态 */
  isPending: boolean
  productMutation: ReturnType<typeof useCreateDraftProduct>
  repositoryMutation: ReturnType<typeof useCreateDraftRepository>
  moduleMutation: ReturnType<typeof useCreateDraftModule>
  decisionMutation: ReturnType<typeof useCreateDraftDecision>
}

// ============================================================================
// Hook
// ============================================================================

export function useOnboardingAction(): UseOnboardingAction {
  const queryClient = useQueryClient()
  const productMutation = useCreateDraftProduct()
  const repositoryMutation = useCreateDraftRepository()
  const moduleMutation = useCreateDraftModule()
  const decisionMutation = useCreateDraftDecision()
  const bindRepositoryToProduct = useBindRepositoryToProduct()
  const bindModuleToProduct = useBindModuleToProduct()

  /** 统一的缓存失效 */
  const invalidateOnboardingQueries = useCallback(async () => {
    await Promise.all([
      queryClient.invalidateQueries({ queryKey: ONBOARDING_STATE_QUERY_KEY }),
      queryClient.invalidateQueries({ queryKey: ONBOARDING_CHAIN_STATE_QUERY_KEY }),
    ])
  }, [queryClient])

  /**
   * product 步骤：创建 Product → 前进到 repository。
   *
   * phase10-08 修复：必须在同一成功路径上显式冻结刚创建的 Product，
   * 不再依赖 ReadOnboardingChainState 的自动猜测。
   */
  const submitProduct = useCallback(async (input: ProductStepInput): Promise<OnboardingActionSuccess> => {
    const result = await productMutation.mutateAsync({ name: input.name })
    if (result.product_id) {
      await onboardingClient.freezeProductAnchor({ productId: result.product_id })
    }
    await invalidateOnboardingQueries()
    return {
      resultKind: 'advance',
      nextStep: 'repository',
      successMessage: `产品「${input.name}」创建成功`,
    }
  }, [productMutation, invalidateOnboardingQueries])

  /**
   * repository 步骤：创建 Repository → 尝试绑定到 Product。
   *
   * 若能同页闭合（绑定成功），前进到 module；否则进入 canonical handoff。
   */
  const submitRepository = useCallback(async (
    input: RepositoryStepInput,
    currentProductId: string,
  ): Promise<OnboardingActionSuccess> => {
    const result = await repositoryMutation.mutateAsync({
      name: input.name,
      url: input.url,
    })

    const repositoryId = result.repository_id

    // 尝试绑定到当前 Product
    if (currentProductId) {
      try {
        await bindRepositoryToProduct.mutateAsync({
          repositoryId,
          productId: currentProductId,
        })
          await invalidateOnboardingQueries()
        return {
          resultKind: 'advance',
          nextStep: 'module',
          successMessage: `仓库「${input.name}」创建并绑定成功`,
        }
      } catch {
        // 绑定失败，进入 canonical handoff
      }
    }

      await invalidateOnboardingQueries()

    // 进入 canonical handoff：跳转到 Repository Detail 完成绑定
    return {
      resultKind: 'handoff',
      nextStep: 'repository',
      navigateTo: '/repositories/$repositoryId',
      params: { repositoryId },
      search: buildOnboardingChainReturnSearch('repository', currentProductId),
      successMessage: `仓库「${input.name}」创建成功，请在详情页完成绑定`,
    }
    }, [repositoryMutation, bindRepositoryToProduct, invalidateOnboardingQueries])

  /**
   * module 步骤：创建 Module → 尝试绑定到 Product。
   */
  const submitModule = useCallback(async (
    input: ModuleStepInput,
    currentProductId: string,
  ): Promise<OnboardingActionSuccess> => {
    const result = await moduleMutation.mutateAsync({ name: input.name })

    const moduleId = result.id

    // 尝试绑定到当前 Product
    if (currentProductId) {
      try {
        await bindModuleToProduct.mutateAsync({
          productId: currentProductId,
          moduleId,
        })
          await invalidateOnboardingQueries()
        return {
          resultKind: 'advance',
          nextStep: 'decision',
          successMessage: `模块「${input.name}」创建并绑定成功`,
        }
      } catch {
        // 绑定失败，进入 canonical handoff
      }
    }

      await invalidateOnboardingQueries()

    return {
      resultKind: 'handoff',
      nextStep: 'module',
      navigateTo: '/modules/$moduleId',
      params: { moduleId },
      search: buildOnboardingChainReturnSearch('module', currentProductId),
      successMessage: `模块「${input.name}」创建成功，请在详情页完成绑定`,
    }
    }, [moduleMutation, bindModuleToProduct, invalidateOnboardingQueries])

  /**
   * decision 步骤：创建 Decision。
   */
  const submitDecision = useCallback(async (input: DecisionStepInput): Promise<OnboardingActionSuccess> => {
    await decisionMutation.mutateAsync({
      title: input.title,
      choice: input.choice,
      reason: input.reason,
    })
      await invalidateOnboardingQueries()
    return {
      resultKind: 'complete',
      nextStep: 'complete',
      successMessage: `决策「${input.title}」记录成功`,
    }
    }, [decisionMutation, invalidateOnboardingQueries])

  const isPending =
    productMutation.isPending ||
    repositoryMutation.isPending ||
    moduleMutation.isPending ||
    decisionMutation.isPending ||
    bindRepositoryToProduct.isPending ||
    bindModuleToProduct.isPending

  return {
    submitProduct,
    submitRepository,
    submitModule,
    submitDecision,
    isPending,
    productMutation,
    repositoryMutation,
    moduleMutation,
    decisionMutation,
  }
}
