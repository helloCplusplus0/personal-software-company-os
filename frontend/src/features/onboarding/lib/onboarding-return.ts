/**
 * Onboarding 返回优先级 helper
 *
 * phase06-15 §"canonical detail 页返回优先级必须正式并入 fromOnboarding"
 * phase10-08 §"canonical handoff 返回合同必须正式切换到 onboardingProductId"
 *
 * 约束：
 *   - fromOnboarding 的返回优先级高于 fromList / fromDashboard / fromProductDetail / fromModuleDetail
 *   - fromOnboarding=true 时返回 /onboarding，并恢复 onboardingStep 指定步骤
 *   - onboardingStep 缺省时仍返回 /onboarding（由 OnboardingPage 内部派生当前步骤）
 *   - phase10-08: 正式返回合同已从草稿摘要 search 切换为 fromOnboarding + onboardingProductId + onboardingStep
 *
 * 本 helper 只承接"返回 /onboarding 时所需的 search 对象构造"，
 * 不接管各 detail 页对 fromList / fromDashboard / fromProductDetail / fromModuleDetail 的原生处理。
 */
import {
  buildOnboardingDraftSearch,
  type OnboardingSourceSearchParsed,
  parseOnboardingDraftSearch,
} from './onboarding-source-schema'
import type { OnboardingStep } from '../types'

/**
 * 判定是否应优先返回 /onboarding。
 *
 * 仅当 fromOnboarding === true 时返回 true；
 * 其他来源标记（fromList / fromDashboard / fromProductDetail / fromModuleDetail）不影响本判定。
 */
export function shouldReturnToOnboarding(
  search: OnboardingSourceSearchParsed,
): boolean {
  return search.fromOnboarding === true
}

/**
 * 构造返回 /onboarding 时的 search 对象（兼容旧草稿摘要合同）。
 *
 * - onboardingStep 存在 → 携带该步骤，由 OnboardingPage 在服务端步骤未到位时作为本地兜底
 * - onboardingProductId 存在 → 携带正式恢复锚点（phase10-08 新增）
 * - onboardingStep 缺省 → 返回空对象，OnboardingPage 内部派生当前步骤
 *
 * @deprecated phase10-08 起，正式返回合同已切换为 buildOnboardingChainReturnSearch。
 * 本函数保留仅用于兼容旧 detail 页面中已存在的 fromOnboarding 返回逻辑。
 * phase10-08 已更新：同时携带 onboardingProductId 作为正式恢复锚点。
 */
export function buildOnboardingReturnSearch(
  search: OnboardingSourceSearchParsed,
): { onboardingStep?: OnboardingStep; onboardingProductId?: string } & ReturnType<typeof buildOnboardingDraftSearch> {
  return {
    ...(search.onboardingStep ? { onboardingStep: search.onboardingStep } : {}),
    ...(search.onboardingProductId ? { onboardingProductId: search.onboardingProductId } : {}),
    ...buildOnboardingDraftSearch(parseOnboardingDraftSearch(search)),
  }
}

/**
 * 构造返回 /onboarding 时的正式 search 对象（phase10-08 新合同）。
 *
 * 正式来源合同冻结为：
 *   - fromOnboarding=true
 *   - onboardingProductId=<current_product_id>
 *   - onboardingStep=<step>
 *
 * 不再依赖草稿摘要 search（productDraftId 等）。
 */
export function buildOnboardingChainReturnSearch(
  step: OnboardingStep,
  productId: string,
): Record<string, unknown> {
  return {
    fromOnboarding: true,
    onboardingProductId: productId,
    onboardingStep: step,
  }
}
