/**
 * useOnboardingRead — Onboarding 只读 query owner
 *
 * phase06-15 §"query 层继续保持纯只读"
 * phase06-06 §"Onboarding 前端只读 query owner"
 * phase10-08 §"useOnboardingRead 必须升级为六段式链路读取 owner"
 *
 * 职责：
 *   - 组合 GetFirstRunState + GetOnboardingChainState 产出统一 view model
 *   - 只承接读取、缓存键与响应解包
 *   - 不得混入 create / update / bind / link 写动作
 *
 * 缓存键：['onboarding-state']
 * 失效策略：由消费方（OnboardingPage / DashboardHomePage / index route）在写操作成功后显式失效
 */
import { useQuery, type UseQueryResult } from '@tanstack/react-query'
import { onboardingClient } from './connect-client'
import { FirstRunStatus as ProtoFirstRunStatus } from '@/gen/proto/psco/onboarding/v1/onboarding_pb'
import type {
  OnboardingFullReadResult,
  FirstRunStatus,
  OnboardingStep,
  OnboardingChainState,
} from '../types'

export type UseOnboardingRead = UseQueryResult<OnboardingFullReadResult, Error>

/** Onboarding state 的缓存键，供消费方显式失效使用 */
export const ONBOARDING_STATE_QUERY_KEY = ['onboarding-state'] as const

/** Onboarding chain state 的缓存键，供消费方显式失效使用（phase10-08 新增） */
export const ONBOARDING_CHAIN_STATE_QUERY_KEY = ['onboarding-chain-state'] as const

/** Proto FirstRunStatus 枚举值映射到前端字符串 */
function mapProtoFirstRunStatus(v: ProtoFirstRunStatus): FirstRunStatus {
  switch (v) {
    case ProtoFirstRunStatus.NOT_STARTED: return 'not_started'
    case ProtoFirstRunStatus.IN_PROGRESS: return 'in_progress'
    case ProtoFirstRunStatus.COMPLETED: return 'completed'
    default: return 'not_started'
  }
}

/** Proto OnboardingStep 枚举值映射到前端字符串 */
function mapProtoOnboardingStep(v: number): OnboardingStep {
  switch (v) {
    case 1: return 'welcome'
    case 2: return 'product'
    case 3: return 'repository'
    case 4: return 'module'
    case 5: return 'decision'
    case 6: return 'complete'
    default: return 'welcome'
  }
}

/**
 * fetchOnboardingRead — Onboarding 共享只读 helper。
 *
 * 供 React Query hook 与根路由 `/` 共同复用，避免长出第二条读取主线。
 * phase10-08: 组合 GetFirstRunState + GetOnboardingChainState。
 */
export async function fetchOnboardingRead(): Promise<OnboardingFullReadResult> {
  const [firstRunRes, chainStateRes] = await Promise.all([
    onboardingClient.getFirstRunState({}),
    onboardingClient.getOnboardingChainState({}),
  ])

  const firstRunState = {
    status: mapProtoFirstRunStatus(firstRunRes.firstRunState?.status ?? ProtoFirstRunStatus.UNSPECIFIED),
    is_first_entry: firstRunRes.firstRunState?.isFirstEntry ?? false,
    current_step: mapProtoOnboardingStep(firstRunRes.firstRunState?.currentStep ?? 0),
    completion_progress: firstRunRes.firstRunState?.completionProgress ?? 0,
  }

  const chainState: OnboardingChainState = {
    current_product_id: chainStateRes.currentProductId ?? '',
    current_step: mapProtoOnboardingStep(chainStateRes.currentStep ?? 0),
    resume_status: (chainStateRes.resumeStatus as OnboardingChainState['resume_status']) ?? 'cold_start',
    next_step_kind: (chainStateRes.nextStepKind as OnboardingChainState['next_step_kind']) ?? 'create',
    canonical_handoff_target: chainStateRes.canonicalHandoffTarget ?? undefined,
    return_hint: chainStateRes.returnHint ?? undefined,
  }

  return {
    first_run_state: firstRunState,
    chain_state: chainState,
  }
}

/**
 * fetchOnboardingChainStateOnly — 只读 GetOnboardingChainState helper。
 *
 * phase10-08 新增：供 detail 页返回后或写操作成功后单独刷新链状态使用。
 */
export async function fetchOnboardingChainStateOnly(): Promise<OnboardingChainState> {
  const res = await onboardingClient.getOnboardingChainState({})
  return {
    current_product_id: res.currentProductId ?? '',
    current_step: mapProtoOnboardingStep(res.currentStep ?? 0),
    resume_status: (res.resumeStatus as OnboardingChainState['resume_status']) ?? 'cold_start',
    next_step_kind: (res.nextStepKind as OnboardingChainState['next_step_kind']) ?? 'create',
    canonical_handoff_target: res.canonicalHandoffTarget ?? undefined,
    return_hint: res.returnHint ?? undefined,
  }
}

export function useOnboardingRead(): UseOnboardingRead {
  return useQuery<OnboardingFullReadResult, Error>({
    queryKey: ONBOARDING_STATE_QUERY_KEY,
    queryFn: fetchOnboardingRead,
  })
}
