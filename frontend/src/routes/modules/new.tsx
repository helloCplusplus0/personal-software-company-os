import { createFileRoute } from '@tanstack/react-router'
import { z } from 'zod'
import { ModuleCreatePage } from '@/features/module-registry/pages/module-create-page'
import { dashboardSourceSearchSchema } from '@/features/dashboard/lib/dashboard-source-schema'

/**
 * ModuleCreateRoute — /modules/new
 * 承接 CreateModule（§3.1）
 *
 * phase05-13：扩展承接 Dashboard 来源参数，
 * 用于从 Dashboard 空状态 CTA 跳转后保留返回 Dashboard 上下文。
 */
const moduleCreateSearchSchema = z.object({
  ...dashboardSourceSearchSchema,
})

export const Route = createFileRoute('/modules/new')({
  validateSearch: moduleCreateSearchSchema,
  component: ModuleCreatePage,
})
