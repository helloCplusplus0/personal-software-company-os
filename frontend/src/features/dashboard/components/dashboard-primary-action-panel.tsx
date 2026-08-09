/**
 * DashboardPrimaryActionPanel — 主 CTA 面板
 *
 * phase05-13 体验修复：
 * - 主 CTA 内联到 Dashboard 标题行右侧，不再独占一行带边框面板
 * - CTA 1-4（创建导向，无副标题）：渲染单个紧凑 Button
 * - CTA 5-8（有 target_label 副标题）：在 Button 前追加 muted 前缀文本（sm+ 屏幕可见）
 * - computing / hidden / suppressed 状态仍返回 null（不渲染）
 *
 * phase05-10 §4.5 DashboardPrimaryActionPanel 语义不变：
 *   - 独立于四区块，直接挂载在 DashboardHomePageShell 标题行
 *   - 只承接主 CTA 优先级矩阵的命中与展示
 *   - 同一时刻只展示一个主 CTA
 *
 * 状态机（phase05-13 §"主 CTA 状态机"）：
 *   - computing：overview query 未成功，或 overview 成功未命中 CTA 1-4 且 feedback query 未成功 → 不渲染
 *   - ready：命中 CTA 1-8 → 渲染对应主 CTA 按钮
 *   - hidden：命中 CTA 9（无缺口）→ 不渲染
 *   - suppressed：overview 成功未命中 CTA 1-4，feedback query 失败 → 不渲染
 *
 * CTA 1-4 跳转参数（phase05-13 §"CTA 1-4 跳转参数"）：
 *   - 携带 fromDashboard=true / dashboardSection=empty-state / dashboardReturnTo=/dashboard
 */
import { useNavigate } from '@tanstack/react-router'
import { Button } from '@/components/ui/button'
import { ArrowRight } from 'lucide-react'
import type {
  DashboardOverview,
  FeedbackSignal,
} from '../types'
import {
  computePrimaryCta,
  type ComputePrimaryCtaResult,
} from '../lib/cta-matrix'
import { buildDashboardSourceParams } from '../lib/dashboard-source'

interface DashboardPrimaryActionPanelProps {
  // overview query 状态
  overviewStatus: 'loading' | 'ready' | 'error'
  overview: DashboardOverview | undefined
  // feedback query 状态
  feedbackStatus: 'loading' | 'ready' | 'error'
  currentFocusSignals: FeedbackSignal[]
  assetFeedbackRepresentativeSignals: FeedbackSignal[]
}

/**
 * DashboardPrimaryActionPanel — 主 CTA 面板（内联到标题行）。
 *
 * 面板状态派生规则：
 * 1. overview 未成功 → computing（不渲染）
 * 2. overview 成功，调用 computePrimaryCta 初步判定：
 *    - 命中 CTA 1-4 → ready（不依赖 feedback query）
 *    - 命中 CTA 9 → hidden
 *    - 需要判定 CTA 5-8 → 进入步骤 3
 * 3. 需要 CTA 5-8 判定时：
 *    - feedback loading → computing
 *    - feedback error → suppressed
 *    - feedback ready → 调用 computePrimaryCta 完整判定
 *
 * 注意：CTA 1-4 命中时直接 ready，不等 feedback query；
 * 这样空状态冷启动时能立即展示创建导向 CTA，不被 feedback query 阻塞。
 */
export function DashboardPrimaryActionPanel({
  overviewStatus,
  overview,
  feedbackStatus,
  currentFocusSignals,
  assetFeedbackRepresentativeSignals,
}: DashboardPrimaryActionPanelProps) {
  const navigate = useNavigate()

  // 步骤 1：overview 未成功 → computing
  if (overviewStatus !== 'ready' || !overview) {
    return null
  }

  // 步骤 2：先用 overview 判定 CTA 1-4 或 CTA 9
  // 此时 currentFocusSignals / representativeSignals 传空数组，
  // computePrimaryCta 只会命中 CTA 1-4 或回退到 CTA 5-8 检查（空数组找不到信号，最终 CTA 9）
  const overviewOnlyResult = computePrimaryCta(overview, [], [])

  // 命中 CTA 1-4：直接 ready，不等 feedback query
  if (overviewOnlyResult.state === 'ready' && overviewOnlyResult.ctaId !== 'CTA_9') {
    // overviewOnlyResult 此时只可能是 CTA 1-4（传空 signals 时 computePrimaryCta
    // 只能命中 CTA 1-4 或回退 CTA 9；CTA 9 已被上方 !== 'CTA_9' 守卫排除）
    return <ReadyCtaButton result={overviewOnlyResult} navigate={navigate} />
  }

  // 步骤 3：需要 CTA 5-8 判定
  // feedback loading → computing
  if (feedbackStatus === 'loading') {
    return null
  }

  // feedback error → suppressed
  if (feedbackStatus === 'error') {
    return null
  }

  // feedback ready → 完整判定
  if (feedbackStatus === 'ready') {
    const fullResult = computePrimaryCta(
      overview,
      currentFocusSignals,
      assetFeedbackRepresentativeSignals,
    )

    // 命中 CTA 5-8 → ready
    if (fullResult.state === 'ready') {
      return <ReadyCtaButton result={fullResult} navigate={navigate} />
    }

    // 命中 CTA 9 → hidden
    return null
  }

  // 兜底：不应到达
  return null
}

/**
 * ReadyCtaButton — 命中 CTA 1-8 时渲染的紧凑主 CTA（内联到标题行右侧）。
 *
 * 渲染形态：
 * - CTA 1-4（无 feedbackSignal）：单个紧凑 Button
 * - CTA 5-8（有 feedbackSignal.target_label）：muted 前缀文本 + Button（前缀在 sm+ 屏幕可见）
 *
 * 跳转参数：
 * - CTA 1-4：dashboardSection=empty-state
 * - CTA 5-8：dashboardSection=current-focus（由 computePrimaryCta 已设置）
 */
function ReadyCtaButton({
  result,
  navigate,
}: {
  result: Extract<ComputePrimaryCtaResult, { state: 'ready' }>
  navigate: ReturnType<typeof useNavigate>
}) {
  const handleClick = () => {
    const sourceParams = buildDashboardSourceParams(result.dashboardSection)

    // 按 result.to 路由模式决定 params
    if (result.to === '/modules/new' || result.to === '/products/new' || result.to === '/repositories/new') {
      // CTA 1-4：创建导向，无 params
      navigate({
        to: result.to,
        search: sourceParams,
      })
    } else if (result.to === '/decisions/$decisionId' && result.feedbackSignal) {
      // CTA 5：单项决策信号
      navigate({
        to: '/decisions/$decisionId',
        params: { decisionId: result.feedbackSignal.target_id },
        search: sourceParams,
      })
    } else if (result.to === '/decisions') {
      // CTA 5：聚合决策信号
      navigate({
        to: '/decisions',
        search: sourceParams,
      })
    } else if (result.to === '/products/$productId' && result.feedbackSignal) {
      // CTA 6-8：产品缺口
      navigate({
        to: '/products/$productId',
        params: { productId: result.feedbackSignal.target_id },
        search: sourceParams,
      })
    }
  }

  return (
    <div className="flex items-center gap-2">
      {/* CTA 5-8 副标题：目标 label，仅在 sm+ 屏幕展示，避免窄屏拥挤 */}
      {result.feedbackSignal && (
        <span className="hidden max-w-[160px] truncate text-xs text-muted-foreground sm:inline">
          {result.feedbackSignal.target_label}
        </span>
      )}
      <Button onClick={handleClick} size="sm" className="shrink-0">
        {result.actionLabel}
        <ArrowRight className="ml-1.5 h-4 w-4" />
      </Button>
    </div>
  )
}
