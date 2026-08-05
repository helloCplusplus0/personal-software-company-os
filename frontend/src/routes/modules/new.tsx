import { createFileRoute } from '@tanstack/react-router'
import { ModuleCreatePage } from '@/features/module-registry/pages/module-create-page'

/**
 * ModuleCreateRoute — /modules/new
 * 承接 CreateModule（§3.1）
 */
export const Route = createFileRoute('/modules/new')({
  component: ModuleCreatePage,
})
