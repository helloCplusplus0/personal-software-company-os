/**
 * Onboarding 来源参数 Zod schema 片段
 *
 * phase06-15 §"canonical detail 页返回优先级必须正式并入 fromOnboarding"
 * phase10-08 §"canonical handoff 返回合同必须正式切换到 onboardingProductId"
 *
 * 为四类 canonical detail 路由提供统一的 Onboarding 来源参数 schema，
 * 供各路由通过 spread 操作符扩展到既有 validateSearch schema。
 *
 * phase10-08 新合同：
 *   - fromOnboarding + onboardingProductId + onboardingStep 是正式来源合同
 *   - 草稿摘要字段（productDraftId 等）保留为兼容但不再作为正式恢复主线
 *
 * 返回优先级约束（phase06-15 spec）：
 *   - fromOnboarding 的返回优先级高于 fromList / fromDashboard / fromProductDetail / fromModuleDetail
 *   - fromOnboarding=true 时返回 /onboarding，并恢复 onboardingStep 指定步骤
 */
import { z } from 'zod'

export type OnboardingDraftKey = 'product' | 'repository' | 'module' | 'decision'

export interface OnboardingDraftSummary {
  id: string
  label: string
}

export type OnboardingDraftMap = Partial<Record<OnboardingDraftKey, OnboardingDraftSummary>>

/**
 * Onboarding 来源参数 Zod schema 片段。
 *
 * 通过 spread 扩展到既有 detail 路由 schema：
 * ```ts
 * const xxxDetailSearchSchema = z.object({
 *   ...原生字段,
 *   ...dashboardSourceSearchSchema,
 *   ...onboardingSourceSearchSchema,
 * })
 * ```
 */
export const onboardingSourceSearchSchema = {
  fromOnboarding: z.boolean().optional(),
  onboardingStep: z
    .enum(['welcome', 'product', 'repository', 'module', 'decision', 'complete'])
    .optional(),
  // phase10-08 新增：正式恢复锚点
  onboardingProductId: z.string().optional(),
  // phase06-15 草稿摘要字段（phase10-08 起已降级为兼容字段，不再作为正式恢复主线）
  /** @deprecated phase10-08 起不再作为正式恢复主线 */
  productDraftId: z.string().optional(),
  /** @deprecated phase10-08 起不再作为正式恢复主线 */
  productDraftLabel: z.string().optional(),
  /** @deprecated phase10-08 起不再作为正式恢复主线 */
  repositoryDraftId: z.string().optional(),
  /** @deprecated phase10-08 起不再作为正式恢复主线 */
  repositoryDraftLabel: z.string().optional(),
  /** @deprecated phase10-08 起不再作为正式恢复主线 */
  moduleDraftId: z.string().optional(),
  /** @deprecated phase10-08 起不再作为正式恢复主线 */
  moduleDraftLabel: z.string().optional(),
  /** @deprecated phase10-08 起不再作为正式恢复主线 */
  decisionDraftId: z.string().optional(),
  /** @deprecated phase10-08 起不再作为正式恢复主线 */
  decisionDraftLabel: z.string().optional(),
} as const

/**
 * Onboarding 来源参数的 TypeScript 推断类型。
 */
export type OnboardingSourceSearchParsed = {
  fromOnboarding?: boolean
  onboardingStep?: 'welcome' | 'product' | 'repository' | 'module' | 'decision' | 'complete'
  /** phase10-08 新增：正式恢复锚点 */
  onboardingProductId?: string
  /** @deprecated phase10-08 起不再作为正式恢复主线 */
  productDraftId?: string
  /** @deprecated phase10-08 起不再作为正式恢复主线 */
  productDraftLabel?: string
  /** @deprecated phase10-08 起不再作为正式恢复主线 */
  repositoryDraftId?: string
  /** @deprecated phase10-08 起不再作为正式恢复主线 */
  repositoryDraftLabel?: string
  /** @deprecated phase10-08 起不再作为正式恢复主线 */
  moduleDraftId?: string
  /** @deprecated phase10-08 起不再作为正式恢复主线 */
  moduleDraftLabel?: string
  /** @deprecated phase10-08 起不再作为正式恢复主线 */
  decisionDraftId?: string
  /** @deprecated phase10-08 起不再作为正式恢复主线 */
  decisionDraftLabel?: string
}

export function buildOnboardingDraftSearch(
  drafts: OnboardingDraftMap,
): Pick<
  OnboardingSourceSearchParsed,
  | 'productDraftId'
  | 'productDraftLabel'
  | 'repositoryDraftId'
  | 'repositoryDraftLabel'
  | 'moduleDraftId'
  | 'moduleDraftLabel'
  | 'decisionDraftId'
  | 'decisionDraftLabel'
> {
  return {
    productDraftId: drafts.product?.id,
    productDraftLabel: drafts.product?.label,
    repositoryDraftId: drafts.repository?.id,
    repositoryDraftLabel: drafts.repository?.label,
    moduleDraftId: drafts.module?.id,
    moduleDraftLabel: drafts.module?.label,
    decisionDraftId: drafts.decision?.id,
    decisionDraftLabel: drafts.decision?.label,
  }
}

export function parseOnboardingDraftSearch(
  search: OnboardingSourceSearchParsed,
): OnboardingDraftMap {
  return {
    product:
      search.productDraftId && search.productDraftLabel
        ? { id: search.productDraftId, label: search.productDraftLabel }
        : undefined,
    repository:
      search.repositoryDraftId && search.repositoryDraftLabel
        ? { id: search.repositoryDraftId, label: search.repositoryDraftLabel }
        : undefined,
    module:
      search.moduleDraftId && search.moduleDraftLabel
        ? { id: search.moduleDraftId, label: search.moduleDraftLabel }
        : undefined,
    decision:
      search.decisionDraftId && search.decisionDraftLabel
        ? { id: search.decisionDraftId, label: search.decisionDraftLabel }
        : undefined,
  }
}
