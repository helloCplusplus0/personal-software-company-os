/**
 * useOnboardingRead — Onboarding 只读 query owner
 *
 * phase06-15 §"query 层继续保持纯只读"
 * phase06-06 §"Onboarding 前端只读 query owner"
 *
 * 职责：
 *   - 只承接读取、缓存键与响应解包
 *   - 不得混入 create / update / bind / link 写动作
 *
 * 缓存键：['onboarding-state']
 * 失效策略：由消费方（OnboardingPage / DashboardHomePage / index route）在写操作成功后显式失效
 */
import { useQuery, type UseQueryResult } from '@tanstack/react-query'
import { fetchFirstRunState } from './api-adapter'
import type { OnboardingReadResult } from '../types'

export type UseOnboardingRead = UseQueryResult<OnboardingReadResult, Error>

/** Onboarding state 的缓存键，供消费方显式失效使用 */
export const ONBOARDING_STATE_QUERY_KEY = ['onboarding-state'] as const

export function useOnboardingRead(): UseOnboardingRead {
  return useQuery<OnboardingReadResult, Error>({
    queryKey: ONBOARDING_STATE_QUERY_KEY,
    queryFn: fetchFirstRunState,
  })
}
