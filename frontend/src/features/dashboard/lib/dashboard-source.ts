/**
 * Dashboard 来源参数工具
 *
 * phase05-13 §"既有路由必须承接 Dashboard 来源参数"
 * phase05-03 / phase05-10 §8.2 已冻结三个来源参数：
 *   - fromDashboard: boolean — Dashboard 外层来源标记
 *   - dashboardSection: 'overview' | 'current-focus' | 'asset-feedback' | 'recent-activity' | 'empty-state'
 *   - dashboardReturnTo: string — 返回 Dashboard 路径（当前阶段固定为 '/dashboard'）
 *
 * 约束：
 * - 来源参数是路由搜索参数层的可观察字段，刷新后必须恢复
 * - 主动返回 Dashboard 时，dashboardSection 通过 TanStack Router 路由 state 一次性承接
 *   （phase05-13 §"Dashboard 主动返回的一次性状态承接必须用路由 state"）
 * - 不得使用 sessionStorage / localStorage / Zustand store 承接来源恢复标记
 * - 不得把 dashboardSection 一次性恢复标记提升为 /dashboard 的搜索参数
 */
import { useNavigate, useRouterState, type HistoryState } from '@tanstack/react-router'
import { useCallback } from 'react'
import type { DashboardSection, DashboardSourceSearch } from '../types'

/**
 * 构造 Dashboard 来源参数对象，用于导航到 canonical owner 页面时透传。
 *
 * 返回的对象可直接作为 TanStack Router navigate / Link 的 search 参数补充字段。
 *
 * @param section 来源区块标记
 * @returns 三字段来源参数对象
 */
export function buildDashboardSourceParams(
  section: DashboardSection,
): DashboardSourceSearch {
  return {
    fromDashboard: true,
    dashboardSection: section,
    dashboardReturnTo: '/dashboard',
  }
}

/**
 * useDashboardSource — 从当前路由搜索参数读取 Dashboard 来源上下文。
 *
 * 泛型 TRoute 允许调用方指定具体路由 from，以获得类型安全的 search schema。
 * 若不指定，默认从根路由读取（适用于工具函数式调用）。
 *
 * 返回：
 * - isFromDashboard: 是否从 Dashboard 进入（fromDashboard === true）
 * - dashboardSection: 来源区块（若 isFromDashboard 为 false 则为 undefined）
 * - dashboardReturnTo: 返回路径（默认 '/dashboard'）
 *
 * 注意：本 hook 必须在路由组件内调用，不能在普通工具函数中调用。
 */
export function useDashboardSource(): {
  isFromDashboard: boolean
  dashboardSection: DashboardSection | undefined
  dashboardReturnTo: string
} {
  // 通过 useRouterState 读取当前激活路由的 search 参数
  // 不使用 useSearch({ strict: false })，避免泛型推断与宽松读取的类型冲突
  // 字段存在性由调用方路由的 validateSearch 保证
  const search = useRouterState({
    select: (s) => s.location.search,
  }) as DashboardSourceSearch

  const isFromDashboard = search.fromDashboard === true
  const dashboardSection = isFromDashboard ? search.dashboardSection : undefined
  const dashboardReturnTo = search.dashboardReturnTo ?? '/dashboard'

  return {
    isFromDashboard,
    dashboardSection,
    dashboardReturnTo,
  }
}

/**
 * navigateBackToDashboard — 主动返回 Dashboard 导航工具。
 *
 * phase05-13 §"Dashboard 主动返回的一次性状态承接必须用路由 state"：
 * - 使用 TanStack Router 的 navigate({ to: '/dashboard', state: { dashboardSection } })
 * - dashboardSection 作为一次性恢复标记，通过路由 state 承接
 * - 不持久化，刷新 /dashboard 后不保留
 *
 * @param section 当前所在区块（用于 Dashboard 恢复时定位滚动或焦点）
 *
 * 注意：本工具是一个 hook（内部使用 useNavigate），必须在组件内调用。
 */
export function useNavigateBackToDashboard() {
  const navigate = useNavigate()

  return useCallback(
    (section: DashboardSection) => {
      // HistoryState 是 TanStack Router 预留的空可扩展接口，
      // dashboardSection 作为一次性恢复标记挂在路由 state 上
      // （phase05-13 §"一次性路由 state 承接"）
      const routeState = { dashboardSection: section } as HistoryState
      navigate({
        to: '/dashboard',
        state: routeState,
      })
    },
    [navigate],
  )
}

/**
 * useDashboardReturnSection — DashboardHomePage 读取一次性路由 state 中的 dashboardSection。
 *
 * phase05-13 §"Dashboard 主动返回的一次性状态承接必须用路由 state"：
 * - 通过 useRouterState().location.state 读取一次性 dashboardSection
 * - 读取后该 state 不持久化，刷新 /dashboard 后不保留
 * - 不使用 sessionStorage / localStorage / Zustand store
 *
 * 返回 dashboardSection 或 undefined（无一次性恢复标记时）。
 */
export function useDashboardReturnSection(): DashboardSection | undefined {
  const routerState = useRouterState()
  const locationState = routerState.location.state as { dashboardSection?: DashboardSection } | undefined
  return locationState?.dashboardSection
}

/**
 * 合并 Dashboard 来源参数到既有 search 对象。
 *
 * 用于既有 canonical 路由在导航到下一跳时，需要同时保留原生来源参数
 * （fromList / fromModuleDetail 等）与 Dashboard 来源参数。
 *
 * @param baseSearch 既有原生 search 对象
 * @param section Dashboard 来源区块
 * @returns 合并后的 search 对象
 */
export function mergeDashboardSource<T extends object>(
  baseSearch: T,
  section: DashboardSection,
): T & DashboardSourceSearch {
  return {
    ...baseSearch,
    ...buildDashboardSourceParams(section),
  }
}

/**
 * mergeCurrentDashboardSource — 将当前路由中的 Dashboard 外层来源参数合并到下一跳 search。
 *
 * 适用于 List / Detail / Create 之间继续保留 Dashboard 外层返回链：
 * - 当前页面已经携带 fromDashboard 时，继续透传三字段
 * - 当前页面不带 fromDashboard 时，保持 baseSearch 原样返回
 */
export function mergeCurrentDashboardSource<T extends object>(
  baseSearch: T,
  currentSearch: DashboardSourceSearch,
): T & DashboardSourceSearch {
  if (currentSearch.fromDashboard !== true) {
    return baseSearch as T & DashboardSourceSearch
  }

  return {
    ...baseSearch,
    fromDashboard: true,
    dashboardSection: currentSearch.dashboardSection ?? 'overview',
    dashboardReturnTo: currentSearch.dashboardReturnTo ?? '/dashboard',
  }
}

/**
 * useDashboardBackButton — 组合 hook，提供"返回 Dashboard"按钮所需的状态与回调。
 *
 * 返回：
 * - showBackButton: 是否应展示"返回 Dashboard"按钮（isFromDashboard）
 * - handleBack: 点击回调，触发主动返回 Dashboard 导航
 * - dashboardSection: 当前区块（用于按钮 tooltip 或无障碍标签）
 *
 * 此 hook 封装了 useDashboardSource + useNavigateBackToDashboard，
 * 供既有 List/Detail/Create 页面组件直接消费，最小化侵入。
 */
export function useDashboardBackButton(): {
  showBackButton: boolean
  handleBack: () => void
  dashboardSection: DashboardSection | undefined
} {
  const { isFromDashboard, dashboardSection } = useDashboardSource()
  const navigateBack = useNavigateBackToDashboard()

  const handleBack = useCallback(() => {
    // 优先使用路由搜索参数中的 dashboardSection；
    // 若缺失（异常情况），回退到 'overview' 作为安全默认
    navigateBack(dashboardSection ?? 'overview')
  }, [navigateBack, dashboardSection])

  return {
    showBackButton: isFromDashboard,
    handleBack,
    dashboardSection,
  }
}
