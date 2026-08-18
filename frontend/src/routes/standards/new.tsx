import { createFileRoute } from '@tanstack/react-router'
import { StandardCreatePage } from '@/features/standard/pages/standard-create-page'

/**
 * StandardsNewRoute — /standards/new
 *
 * phase14-05 §ADDED-1 URL 语义冻结：创建规范（基本信息 + 目录结构整树）。
 */
export const Route = createFileRoute('/standards/new')({
  component: StandardCreatePage,
})
