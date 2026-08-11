/**
 * Dashboard slice-local Connect client。
 *
 * phase07-10 §5.3：Dashboard 切片同时承接 Dashboard / Export / Backup 三个 service 的 client。
 * Export 与 Backup 在 SovereigntyPanel 内作为显式过渡位保留。
 */
import { createClient } from '@connectrpc/connect'
import { DashboardService } from '@/gen/proto/psco/dashboard/v1/dashboard_pb'
import { ExportService } from '@/gen/proto/psco/export/v1/export_pb'
import { BackupService } from '@/gen/proto/psco/backup/v1/backup_pb'
import { transport } from '@/shared/rpc/connect-transport'

export const dashboardClient = createClient(DashboardService, transport)
export const exportClient = createClient(ExportService, transport)
export const backupClient = createClient(BackupService, transport)