import { createFileRoute } from '@tanstack/react-router'
import { StandardListPage } from '@/features/standard/pages/standard-list-page'

/**
 * StandardsIndexRoute — /standards
 *
 * phase14-05 §ADDED-1 URL 语义冻结：规范列表。
 * ListStandards 无参数不分页，第一版不做筛选（无搜索参数）。
 */
export const Route = createFileRoute('/standards/')({
  component: StandardListPage,
})
