import { Outlet, createFileRoute, useRouterState } from '@tanstack/react-router'
import { StandardDetailPage } from '@/features/standard/pages/standard-detail-page'

/**
 * StandardsDetailRoute — /standards/:standardId
 *
 * phase14-05 §ADDED-1 URL 语义冻结：规范详情（树展示 + 绑定管理 + revision 回看）。
 * route param standardId → 页面 props 接线（页面组件不直接依赖路由）。
 *
 * 嵌套让位（TanStack Router 文件约定）：`$standardId.edit.tsx` 是本路由的 child，
 * URL `/standards/:id/edit` 匹配时 parent component 必须渲染 `<Outlet />` 才会挂载
 * 编辑页。此处采用 layout-less 模式：edit child 匹配时整页让位渲染 Outlet，
 * 其余情况渲染详情页——两个视图互斥，不产生 detail + edit 叠加。
 */
function StandardsDetailRoute() {
  const { standardId } = Route.useParams()
  const isEditChild = useRouterState({
    select: (s) => s.matches.some((m) => m.routeId === '/standards/$standardId/edit'),
  })
  if (isEditChild) return <Outlet />
  return <StandardDetailPage standardId={standardId} />
}

export const Route = createFileRoute('/standards/$standardId')({
  component: StandardsDetailRoute,
})
