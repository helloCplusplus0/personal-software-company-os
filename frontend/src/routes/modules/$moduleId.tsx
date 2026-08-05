import { createFileRoute } from '@tanstack/react-router'
import { ModuleDetailPage } from '@/features/module-registry/pages/module-detail-page'

/**
 * ModuleDetailRoute — /modules/:moduleId
 * 统一详情读模型宿主，承接详情读取、版本登记入口、绑定面板与 Decision 只读入口（§3.1）
 */
export const Route = createFileRoute('/modules/$moduleId')({
  component: ModuleDetailPage,
})
