/**
 * OnboardingCtaButton — Dashboard 内的 Start / Continue Onboarding 入口
 *
 * phase06-15 §"Dashboard 到 Onboarding 的继续入口"
 *
 * 约束：
 *   - not_started → 展示「Start Onboarding」
 *   - in_progress → 展示「Continue Onboarding」
 *   - completed → 不展示首轮录入 CTA
 *   - 当前阶段不得在 /dashboard 外再发明第二个首轮录入入口
 *
 * 该按钮独立于 DashboardPrimaryActionPanel，作为 phase06-15 的正式 onboarding 入口。
 * 不承接 mutation，只承接读取 first_run_state 与导航到 /onboarding。
 */
import { useNavigate } from '@tanstack/react-router'
import { useQueryClient } from '@tanstack/react-query'
import { Button } from '@/components/ui/button'
import { Skeleton } from '@/components/ui/skeleton'
import { Sparkles, ArrowRight } from 'lucide-react'
import { useOnboardingRead, ONBOARDING_STATE_QUERY_KEY } from '@/features/onboarding/data/use-onboarding-read'

/**
 * OnboardingCtaButton — Dashboard 内的 onboarding CTA 按钮。
 *
 * 读取 first_run_state 决定按钮文案与可见性。
 * completed 时不渲染；not_started / in_progress 时渲染对应 CTA。
 */
export function OnboardingCtaButton() {
  const navigate = useNavigate()
  const queryClient = useQueryClient()

  const onboardingStateQuery = useOnboardingRead()

  // loading：渲染骨架占位，保持布局稳定
  if (onboardingStateQuery.isLoading) {
    return <Skeleton className="h-9 w-40" />
  }

  // error：降级不渲染（不阻塞 Dashboard）
  if (onboardingStateQuery.isError) {
    return null
  }

  const status = onboardingStateQuery.data?.first_run_state?.status

  // completed：不展示首轮录入 CTA
  if (status === 'completed') {
    return null
  }

  // not_started → Start Onboarding；in_progress → Continue Onboarding
  const label = status === 'not_started' ? '开始首轮录入' : '继续首轮录入'

  const handleClick = () => {
    // 失效 onboarding-state query，确保进入 /onboarding 时拿到最新状态
    queryClient.invalidateQueries({ queryKey: ONBOARDING_STATE_QUERY_KEY })
    navigate({ to: '/onboarding' })
  }

  return (
    <Button onClick={handleClick} size="sm" variant="default">
      <Sparkles className="mr-1.5 h-4 w-4" />
      {label}
      <ArrowRight className="ml-1.5 h-4 w-4" />
    </Button>
  )
}
