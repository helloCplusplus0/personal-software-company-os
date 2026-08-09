import { createFileRoute } from '@tanstack/react-router'
import { DashboardHomePage } from '@/features/dashboard/pages/dashboard-home-page'

/**
 * DashboardRoute — /dashboard
 *
 * phase05-13 §"Dashboard 路由与主导航接入必须按 phase05-05 落地"
 * phase05-05 / phase05-10 §3.1 已冻结：
 *   - Dashboard Home 的正式业务入口路由冻结为 /dashboard
 *   - Dashboard 在当前阶段作为既有主导航中的一级入口新增，不替代根级布局宿主本身
 *   - 当前阶段不把 / 单值解释为 Dashboard Home
 *
 * 路由约束（phase05-13 spec）：
 *   - 必须使用 createFileRoute('/dashboard') 注册
 *   - 路由组件只承接 DashboardHomePage
 *   - 当前阶段不得为 /dashboard 引入 validateSearch（/dashboard 不承接搜索参数）
 *   - 不得把 /dashboard 拆成 overview / activity / feedback 等子路由
 */
export const Route = createFileRoute('/dashboard')({
  component: DashboardHomePage,
})
