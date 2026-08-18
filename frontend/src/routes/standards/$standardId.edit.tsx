import { createFileRoute } from '@tanstack/react-router'
import { StandardEditPage } from '@/features/standard/pages/standard-edit-page'

/**
 * StandardsEditRoute — /standards/:standardId/edit
 *
 * phase14-05 §ADDED-1 URL 语义冻结：编辑规范（整树替换 + change_summary）。
 * route param standardId → 页面 props 接线（页面组件不直接依赖路由）。
 */
function StandardsEditRoute() {
  const { standardId } = Route.useParams()
  return <StandardEditPage standardId={standardId} />
}

export const Route = createFileRoute('/standards/$standardId/edit')({
  component: StandardsEditRoute,
})
