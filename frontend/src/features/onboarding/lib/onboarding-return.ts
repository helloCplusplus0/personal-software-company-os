/**
 * Onboarding 返回优先级 helper
 *
 * phase06-15 §"canonical detail 页返回优先级必须正式并入 fromOnboarding"
 *
 * 约束：
 *   - fromOnboarding 的返回优先级高于 fromList / fromDashboard / fromProductDetail / fromModuleDetail
 *   - fromOnboarding=true 时返回 /onboarding，并恢复 onboardingStep 指定步骤
 *   - onboardingStep 缺省时仍返回 /onboarding（由 OnboardingPage 内部派生当前步骤）
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
 * 构造返回 /onboarding 时的 search 对象。
 *
 * - onboardingStep 存在 → 携带该步骤，由 OnboardingPage 在服务端步骤未到位时作为本地兜底
 * - onboardingStep 缺省 → 返回空对象，OnboardingPage 内部派生当前步骤
 */
export function buildOnboardingReturnSearch(
  search: OnboardingSourceSearchParsed,
): { onboardingStep?: OnboardingStep } & ReturnType<typeof buildOnboardingDraftSearch> {
  return {
    ...(search.onboardingStep ? { onboardingStep: search.onboardingStep } : {}),
    ...buildOnboardingDraftSearch(parseOnboardingDraftSearch(search)),
  }
}
