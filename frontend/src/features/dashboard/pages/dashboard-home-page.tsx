/**
 * DashboardHomePage — Dashboard 主页面
 *
 * phase05-13 §"DashboardHomePage 必须编排三个独立查询并派生整页状态"
 * phase05-06 已冻结整页查询状态、整页视图状态、分区状态模型
 *
 * 三路独立查询（phase05-10 §7.1）：
 *   - DashboardOverviewRead（主聚合读取，失败触发整页失败）
 *   - FeedbackSignalRead（附属聚合读取，失败只触发局部失败）
 *   - RecentActivityRead（附属聚合读取，失败只触发局部失败）
 *
 * 整页视图状态派生（phase05-06 / phase05-13 spec）：
 *   - initial-loading：只允许出现在 overview query 首次 pending 时
 *   - page-error：只允许由 overview query 失败触发
 *   - ready：一旦 overview query 成功，整页进入 ready，
 *     即使 feedback 或 recent-activity 局部失败也不得回退
 *
 * 整页重试（phase05-13 spec）：
 *   - 整页处于 page-error 时提供整页级重试入口
 *   - 整页重试必须同时重新触发 overview / feedback / recent-activity 三个 query
 *
 * phase05-13 体验修复数据流调整：
 *   - 资产概览数字 + 缺口计数合并到 DashboardStatBar（取代原 DashboardOverviewSection）
 *   - AssetFeedback 区块状态基于 representative_signals 独立派生（不再与 CurrentFocus 共享）
 */
import { useQuery, useQueryClient } from '@tanstack/react-query'
import { useNavigate, type HistoryState } from '@tanstack/react-router'
import { useEffect, useRef } from 'react'
import { Button } from '@/components/ui/button'
import { Skeleton } from '@/components/ui/skeleton'
import {
  fetchDashboardOverview,
  fetchFeedbackSignals,
  fetchRecentActivities,
} from '../data/api-adapter'
import { DashboardHomePageShell } from '../components/dashboard-home-page-shell'
import { DashboardStatBar } from '../components/dashboard-stat-bar'
import { CurrentFocusSection } from '../components/current-focus-section'
import { AssetFeedbackSection } from '../components/asset-feedback-section'
import { RecentActivitySection } from '../components/recent-activity-section'
import { DashboardPrimaryActionPanel } from '../components/dashboard-primary-action-panel'
import { SovereigntyPanel } from '../components/sovereignty-panel'
import { OnboardingCtaButton } from '../components/onboarding-cta-button'
import { useDashboardReturnSection } from '../lib/dashboard-source'
import { useReuseSummaryRead } from '@/features/reuse-summary/data/use-reuse-summary-read'
import type { DashboardSection } from '../types'

const DASHBOARD_SECTION_LABELS: Record<DashboardSection, string> = {
  overview: '系统概览',
  'current-focus': 'Current Focus',
  'asset-feedback': 'Asset Feedback',
  'recent-activity': 'Recent Activity',
  'empty-state': '主 CTA',
}

/**
 * DashboardHomePage — Dashboard 主页面组件。
 *
 * 职责：
 * 1. 编排三个独立 TanStack Query 查询
 * 2. 派生整页视图状态（initial-loading / page-error / ready）
 * 3. 将分区状态分发给 stat bar 与四个区块组件
 * 4. 渲染主 CTA 面板（内联到标题行）
 *
 * 不承接任何业务写入（phase05-13 §"Dashboard 前端必须只读不写"）。
 */
export function DashboardHomePage() {
  const navigate = useNavigate({ from: '/dashboard' })
  const queryClient = useQueryClient()
  const sectionRefs = useRef<Record<DashboardSection, HTMLDivElement | null>>({
    overview: null,
    'current-focus': null,
    'asset-feedback': null,
    'recent-activity': null,
    'empty-state': null,
  })

  // 一次性路由 state 中的 dashboardSection 恢复标记（phase05-13 §"一次性路由 state 承接"）
  // 用于主动返回 Dashboard 后的区块定位；读取后立刻清空，避免变成持久事实源。
  const returnSection = useDashboardReturnSection()

  useEffect(() => {
    if (!returnSection) {
      return
    }

    const target = sectionRefs.current[returnSection]
    if (!target) {
      navigate({
        to: '/dashboard',
        replace: true,
        state: {} as HistoryState,
      })
      return
    }

    const frameId = window.requestAnimationFrame(() => {
      target.scrollIntoView({ block: 'start', behavior: 'auto' })
      target.focus({ preventScroll: true })
      navigate({
        to: '/dashboard',
        replace: true,
        state: {} as HistoryState,
      })
    })

    return () => {
      window.cancelAnimationFrame(frameId)
    }
  }, [navigate, returnSection])

  // ============================================================================
  // 三路独立查询（phase05-10 §7.1）
  // ============================================================================

  // 1. DashboardOverviewRead — 主聚合读取
  const overviewQuery = useQuery({
    queryKey: ['dashboard-overview'],
    queryFn: fetchDashboardOverview,
  })

  // 2. FeedbackSignalRead — 附属聚合读取（current_focus_signals + asset_feedback_summary）
  const feedbackQuery = useQuery({
    queryKey: ['dashboard-feedback-signals'],
    queryFn: fetchFeedbackSignals,
  })

  // 3. RecentActivityRead — 附属聚合读取
  const activityQuery = useQuery({
    queryKey: ['dashboard-recent-activities'],
    queryFn: fetchRecentActivities,
  })

  // 4. ReuseSummaryRead — phase06-15 §"Dashboard 复用快照挂接位"
  //    独立页面级只读 query，失败不回退整页，只影响 Asset Feedback 内的 Reuse Snapshot 子区域
  const reuseSummaryQuery = useReuseSummaryRead({ scope: 'dashboard' })

  // ============================================================================
  // 整页视图状态派生（phase05-06 / phase05-13 spec）
  // ============================================================================

  // initial-loading：只允许出现在 overview query 首次 pending 时
  const isInitialLoading = overviewQuery.isLoading

  // page-error：只允许由 overview query 失败触发
  const isPageError = overviewQuery.isError

  // ready：overview query 成功后整页进入 ready
  // （由下方渲染分支兜底：非 initial-loading 且非 page-error 即为 ready）

  // ============================================================================
  // 整页重试（phase05-13 §"整页重试"）
  // ============================================================================

  const handleFullRetry = () => {
    // 整页重试必须同时重新触发三个 query
    queryClient.invalidateQueries({ queryKey: ['dashboard-overview'] })
    queryClient.invalidateQueries({ queryKey: ['dashboard-feedback-signals'] })
    queryClient.invalidateQueries({ queryKey: ['dashboard-recent-activities'] })
  }

  // ============================================================================
  // 分区状态派生
  // ============================================================================

  // feedback query 统一状态（供 stat bar coverage 与 CurrentFocus / AssetFeedback 共用底层判定）
  const feedbackLoading = feedbackQuery.isLoading
  const feedbackError = feedbackQuery.isError

  // stat bar 的 coverage 状态：loading / error / ready
  const coverageStatus: 'loading' | 'ready' | 'error' = feedbackLoading
    ? 'loading'
    : feedbackError
      ? 'error'
      : 'ready'

  // Current Focus 分区状态：基于 current_focus_signals
  const currentFocusStatus: 'loading' | 'ready' | 'empty' | 'error' = feedbackLoading
    ? 'loading'
    : feedbackError
      ? 'error'
      : (feedbackQuery.data?.current_focus_signals?.length ?? 0) === 0
        ? 'empty'
        : 'ready'

  // Asset Feedback 分区状态：基于 representative_signals（独立派生，不再与 CurrentFocus 共享 empty 判定）
  const assetFeedbackStatus: 'loading' | 'ready' | 'empty' | 'error' = feedbackLoading
    ? 'loading'
    : feedbackError
      ? 'error'
      : (feedbackQuery.data?.asset_feedback_summary?.representative_signals?.length ?? 0) === 0
        ? 'empty'
        : 'ready'

  // activity section 状态
  const activitySectionStatus: 'loading' | 'ready' | 'empty' | 'error' = activityQuery.isLoading
    ? 'loading'
    : activityQuery.isError
      ? 'error'
      : (activityQuery.data?.activities?.length ?? 0) === 0
        ? 'empty'
        : 'ready'

  // 主 CTA 面板所需的 feedback 状态（loading / ready / error，不含 empty）
  const ctaFeedbackStatus: 'loading' | 'ready' | 'error' = feedbackLoading
    ? 'loading'
    : feedbackError
      ? 'error'
      : 'ready'

  // phase06-15：ReuseSummaryRead 局部状态派生
  // 该 query 失败不回退整页，只影响 Asset Feedback 内的 Reuse Snapshot 子区域
  const reuseSnapshotStatus: 'loading' | 'ready' | 'empty' | 'error' = reuseSummaryQuery.isLoading
    ? 'loading'
    : reuseSummaryQuery.isError
      ? 'error'
      : (reuseSummaryQuery.data?.module_reuse_summary?.length ?? 0) === 0 &&
          (reuseSummaryQuery.data?.capability_summary?.length ?? 0) === 0
        ? 'empty'
        : 'ready'

  // ============================================================================
  // 渲染
  // ============================================================================

  // 整页 initial-loading：骨架
  if (isInitialLoading) {
    return (
      <div className="space-y-4">
        <div className="flex items-center justify-between">
          <Skeleton className="h-7 w-28" />
        </div>
        <Skeleton className="h-16 w-full" />
        <div className="grid gap-4 md:grid-cols-2">
          <div className="space-y-2">
            <Skeleton className="h-5 w-28" />
            <Skeleton className="h-32 w-full" />
          </div>
          <div className="space-y-2">
            <Skeleton className="h-5 w-28" />
            <Skeleton className="h-32 w-full" />
          </div>
        </div>
      </div>
    )
  }

  // 整页 page-error：overview query 失败
  if (isPageError) {
    return (
      <div className="space-y-4">
        <h1 className="text-xl font-bold">Dashboard</h1>
        <div className="rounded-lg border border-destructive/50 bg-destructive/10 p-4">
          <p className="text-sm text-destructive mb-3">
            Dashboard 概览读取失败：{(overviewQuery.error as Error)?.message ?? '未知错误'}
          </p>
          <Button variant="outline" size="sm" onClick={handleFullRetry}>
            重试
          </Button>
        </div>
      </div>
    )
  }

  // 整页 ready：渲染 stat bar + 四区块 + 主 CTA 面板
  return (
    <>
      {returnSection && (
        <div className="sr-only" aria-live="polite">
          已恢复到 {DASHBOARD_SECTION_LABELS[returnSection]} 区块
        </div>
      )}

      <DashboardHomePageShell
        primaryActionPanel={
          <div className="flex items-center gap-2">
            {/*
              phase06-15 §"Dashboard 到 Onboarding 的继续入口"：
              OnboardingCtaButton 独立于 DashboardPrimaryActionPanel，
              作为 phase06 首轮录入的正式入口（not_started / in_progress 时可见）。
            */}
            <OnboardingCtaButton />
            <div
              ref={(node) => {
                sectionRefs.current['empty-state'] = node
              }}
              data-dashboard-section="empty-state"
              tabIndex={-1}
              className="rounded-md focus:outline-none focus:ring-2 focus:ring-ring focus:ring-offset-2"
            >
              <DashboardPrimaryActionPanel
                overviewStatus="ready"
                overview={overviewQuery.data}
                feedbackStatus={ctaFeedbackStatus}
                currentFocusSignals={feedbackQuery.data?.current_focus_signals ?? []}
                assetFeedbackRepresentativeSignals={
                  feedbackQuery.data?.asset_feedback_summary?.representative_signals ?? []
                }
              />
            </div>
          </div>
        }
        sovereigntyPanel={<SovereigntyPanel />}
        statBar={
          overviewQuery.data && (
            <div
              ref={(node) => {
                sectionRefs.current.overview = node
              }}
              data-dashboard-section="overview"
              tabIndex={-1}
              className="rounded-lg focus:outline-none focus:ring-2 focus:ring-ring focus:ring-offset-2"
            >
              <DashboardStatBar
                overview={overviewQuery.data}
                coverageStatus={coverageStatus}
                summary={feedbackQuery.data?.asset_feedback_summary}
              />
            </div>
          )
        }
        currentFocusSection={
          <div
            ref={(node) => {
              sectionRefs.current['current-focus'] = node
            }}
            data-dashboard-section="current-focus"
            tabIndex={-1}
            className="rounded-lg focus:outline-none focus:ring-2 focus:ring-ring focus:ring-offset-2"
          >
            <CurrentFocusSection
              status={currentFocusStatus}
              signals={feedbackQuery.data?.current_focus_signals ?? []}
              error={feedbackQuery.error as Error | null}
              onRetry={() => queryClient.invalidateQueries({ queryKey: ['dashboard-feedback-signals'] })}
            />
          </div>
        }
        assetFeedbackSection={
          <div
            ref={(node) => {
              sectionRefs.current['asset-feedback'] = node
            }}
            data-dashboard-section="asset-feedback"
            tabIndex={-1}
            className="rounded-lg focus:outline-none focus:ring-2 focus:ring-ring focus:ring-offset-2"
          >
            <AssetFeedbackSection
              status={assetFeedbackStatus}
              signals={
                feedbackQuery.data?.asset_feedback_summary?.representative_signals ?? []
              }
              error={feedbackQuery.error as Error | null}
              onRetry={() => queryClient.invalidateQueries({ queryKey: ['dashboard-feedback-signals'] })}
              reuseSnapshotStatus={reuseSnapshotStatus}
              moduleReuseSummary={reuseSummaryQuery.data?.module_reuse_summary ?? []}
              capabilitySummary={reuseSummaryQuery.data?.capability_summary ?? []}
              reuseSnapshotError={reuseSummaryQuery.error as Error | null}
              onReuseSnapshotRetry={() =>
                queryClient.invalidateQueries({ queryKey: ['reuse-summary', 'dashboard'] })
              }
            />
          </div>
        }
        recentActivitySection={
          <div
            ref={(node) => {
              sectionRefs.current['recent-activity'] = node
            }}
            data-dashboard-section="recent-activity"
            tabIndex={-1}
            className="rounded-lg focus:outline-none focus:ring-2 focus:ring-ring focus:ring-offset-2"
          >
            <RecentActivitySection
              status={activitySectionStatus}
              activities={activityQuery.data?.activities ?? []}
              error={activityQuery.error as Error | null}
              onRetry={() => queryClient.invalidateQueries({ queryKey: ['dashboard-recent-activities'] })}
            />
          </div>
        }
      />
    </>
  )
}
