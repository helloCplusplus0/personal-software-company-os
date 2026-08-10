/**
 * Onboarding 来源参数 Zod schema 片段
 *
 * phase06-15 §"canonical detail 页返回优先级必须正式并入 fromOnboarding"
 *
 * 为四类 canonical detail 路由提供统一的 Onboarding 来源参数 schema，
 * 供各路由通过 spread 操作符扩展到既有 validateSearch schema。
 *
 * 当前不仅需要恢复 onboardingStep，还需要把当前会话中的草稿摘要一并带回
 * /onboarding，避免用户从 canonical detail 返回后丢失已创建草稿身份。
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
  productDraftId: z.string().optional(),
  productDraftLabel: z.string().optional(),
  repositoryDraftId: z.string().optional(),
  repositoryDraftLabel: z.string().optional(),
  moduleDraftId: z.string().optional(),
  moduleDraftLabel: z.string().optional(),
  decisionDraftId: z.string().optional(),
  decisionDraftLabel: z.string().optional(),
} as const

/**
 * Onboarding 来源参数的 TypeScript 推断类型。
 */
export type OnboardingSourceSearchParsed = {
  fromOnboarding?: boolean
  onboardingStep?: 'welcome' | 'product' | 'repository' | 'module' | 'decision' | 'complete'
  productDraftId?: string
  productDraftLabel?: string
  repositoryDraftId?: string
  repositoryDraftLabel?: string
  moduleDraftId?: string
  moduleDraftLabel?: string
  decisionDraftId?: string
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
