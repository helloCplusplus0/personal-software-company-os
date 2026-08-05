import { createFileRoute } from '@tanstack/react-router'
import { ReleaseCreatePage } from '@/features/module-registry/pages/release-create-page'

/**
 * ReleaseCreateRoute — /modules/:moduleId/releases/new
 * 依附当前模块上下文，承接 CreateRelease（§3.1）
 * moduleId 来自路由参数，不得复制可写全局状态（§8.4）
 */
export const Route = createFileRoute('/modules/$moduleId/releases/new')({
  component: ReleaseCreatePage,
})
