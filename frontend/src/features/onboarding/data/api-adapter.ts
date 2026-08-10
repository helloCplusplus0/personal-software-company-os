/**
 * Onboarding API 适配层
 *
 * 对接 phase06-14 后端 onboarding 模块：
 *   GET /api/onboarding/state → GetFirstRunState
 *
 * 字段语义从 .proto -> HTTP DTO 单向派生（phase06-05 / phase06-13 合同约束）。
 * 后端已使用 snake_case，与前端 types.ts 一致，无需转换。
 */
import type { OnboardingReadResult } from '../types'

const API_BASE_URL = import.meta.env.VITE_API_BASE_URL ?? ''

/**
 * fetchFirstRunState — GetFirstRunState 读组
 *
 * GET /api/onboarding/state
 *
 * 返回 first_run_state，包含 status / is_first_entry / current_step / completion_progress。
 */
export async function fetchFirstRunState(): Promise<OnboardingReadResult> {
  const res = await fetch(`${API_BASE_URL}/api/onboarding/state`, {
    headers: { Accept: 'application/json' },
  })

  if (!res.ok) {
    let message = `HTTP ${res.status}`
    try {
      const body = await res.json()
      if (body?.error) {
        message = body.error
      }
    } catch {
      // 响应体非 JSON，保留默认 message
    }
    throw new Error(message)
  }

  return res.json() as Promise<OnboardingReadResult>
}
