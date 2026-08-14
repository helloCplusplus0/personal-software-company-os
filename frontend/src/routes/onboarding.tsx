/**
 * /onboarding 路由 — 首轮录入唯一正式入口
 *
 * phase06-06 §"路由与根级默认进入路径"
 * phase06-15 §"Onboarding 前端主线必须落地为唯一正式首轮入口"
 *            §"canonical detail 页返回优先级必须正式并入 fromOnboarding"
 * phase10-08 §"canonical handoff 返回合同必须正式切换到 onboardingProductId"
 *
 * 约束：
 *   - /onboarding 是唯一正式首轮业务入口
 *   - 不承接业务搜索参数（步骤状态由 first_run_state + chain_state 服务端派生）
 *   - 仅承接可选 onboardingStep：由 detail 页返回时携带，用于服务端步骤未到位时本地兜底
 *   - phase10-08: 新增 onboardingProductId 作为正式恢复锚点
 *   - 不得在 /onboarding 外再发明第二个首轮录入入口
 */
import { createFileRoute } from '@tanstack/react-router'
import { z } from 'zod'
import { OnboardingPage } from '@/features/onboarding/pages/onboarding-page'
import { onboardingSourceSearchSchema } from '@/features/onboarding/lib/onboarding-source-schema'

const onboardingSearchSchema = z.object({
  // phase10-08 §"canonical handoff 返回合同"：
  // fromOnboarding 由 detail 页返回时显式置为 true
  fromOnboarding: onboardingSourceSearchSchema.fromOnboarding,
  // phase10-08 新增：正式恢复锚点
  onboardingProductId: onboardingSourceSearchSchema.onboardingProductId,
  // phase06-15 §"detail 页来源优先级"：
  // 仅在从 canonical detail 页返回时携带，作为服务端步骤未到位时的本地兜底
  onboardingStep: onboardingSourceSearchSchema.onboardingStep,
  // phase06-15 草稿摘要字段（phase10-08 起已降级为兼容字段，不再作为正式恢复主线）
  productDraftId: onboardingSourceSearchSchema.productDraftId,
  productDraftLabel: onboardingSourceSearchSchema.productDraftLabel,
  repositoryDraftId: onboardingSourceSearchSchema.repositoryDraftId,
  repositoryDraftLabel: onboardingSourceSearchSchema.repositoryDraftLabel,
  moduleDraftId: onboardingSourceSearchSchema.moduleDraftId,
  moduleDraftLabel: onboardingSourceSearchSchema.moduleDraftLabel,
  decisionDraftId: onboardingSourceSearchSchema.decisionDraftId,
  decisionDraftLabel: onboardingSourceSearchSchema.decisionDraftLabel,
})

export const Route = createFileRoute('/onboarding')({
  validateSearch: onboardingSearchSchema,
  component: OnboardingPage,
})
