/**
 * / 路由 — 根级 first-run 默认进入路径
 *
 * phase06-06 §"路由与根级默认进入路径"
 * phase06-15 §"Onboarding 前端主线必须落地为唯一正式首轮入口"
 *
 * 行为：
 *   - beforeLoad 读取 GetFirstRunState
 *   - first_run_state = not_started → 默认进入 /onboarding
 *   - first_run_state = in_progress → 默认进入 /dashboard
 *   - first_run_state = completed → 默认进入 /dashboard
 *   - 读取失败 → 降级进入 /dashboard（不得因读取失败阻断用户访问）
 *
 * 约束：
 *   - completed 时不得默认劫持到 /onboarding
 *   - 不得在 / 路由内内联正式 mutation 或 create 语义
 */
import { createFileRoute, redirect } from '@tanstack/react-router'
import { fetchFirstRunState } from '@/features/onboarding/data/api-adapter'

export const Route = createFileRoute('/')({
  beforeLoad: async () => {
    let target = '/dashboard'

    try {
      const result = await fetchFirstRunState()
      const status = result.first_run_state?.status

      if (status === 'not_started') {
        target = '/onboarding'
      }
      // in_progress / completed → /dashboard
      // 读取失败 → 降级 /dashboard
    } catch {
      // 读取失败时降级进入 /dashboard，不得阻断用户访问
      target = '/dashboard'
    }

    throw redirect({ to: target })
  },
  component: () => null,
})
