/**
 * SovereigntyPanel — Export / Backup 用户入口面板（紧凑型）
 *
 * phase06-15 §"Export / Backup 必须作为 Dashboard 内正式用户入口"
 * phase06-16 §"Dashboard 紧凑型优化"：延续 phase05-13 紧凑化规范
 *
 * 约束：
 *   - Dashboard 中必须存在稳定可见的 Export 与 Backup 入口
 *   - 用户可以读取当前 ExportSnapshot 与 BackupSnapshot
 *   - 用户可以从前端触发导出与备份动作
 *   - Backup 的三类失败语义必须继续可区分：
 *     manifest_missing / coverage_incomplete / schema_mismatch
 *   - 当前阶段不得把 Export / Backup 做成隐藏路由、开发者入口或仅测试按钮
 *
 * 紧凑型优化（phase06-16）：
 *   - 移除大段描述文字，标题行内联简短说明
 *   - 快照信息用紧凑 text-xs 列表，不用 border/bg 卡片包裹
 *   - 状态 + 按钮单行排列，两子区域用 divide-x 分隔
 *   - 整体高度从约 200px 压缩到约 80-100px，不推高首屏
 *
 * 字段语义从 .proto -> HTTP DTO 单向派生（phase06-05 / phase06-13 合同约束）。
 */
import { useMutation, useQueryClient } from '@tanstack/react-query'
import { Button } from '@/components/ui/button'
import { Skeleton } from '@/components/ui/skeleton'
import { Download, ShieldCheck, RefreshCw } from 'lucide-react'
import { useExportSnapshotRead } from '../data/use-export-snapshot-read'
import { useBackupSnapshotRead } from '../data/use-backup-snapshot-read'
import { exportClient, backupClient } from '../data/connect-client'
import type { ExportSnapshot, BackupSnapshot, VerifyFailureCode } from '../types'

const EXPORT_SNAPSHOT_QUERY_KEY = ['export-snapshot'] as const
const BACKUP_SNAPSHOT_QUERY_KEY = ['backup-snapshot'] as const

/** Backup 三类失败语义的简短中文标签（phase06-15 §"Backup 三类失败语义"） */
const VERIFY_FAILURE_CODE_LABELS: Record<VerifyFailureCode, string> = {
  manifest_missing: 'manifest 缺失',
  coverage_incomplete: '覆盖不完整',
  schema_mismatch: 'schema 不匹配',
}

/** Backup verified 状态的简短中文标签 */
const BACKUP_VERIFIED_STATUS_LABELS: Record<BackupSnapshot['verified_status'], string> = {
  unverified: '未校验',
  verified: '已校验',
  verify_failed: '校验失败',
}

/**
 * SovereigntyPanel — Export / Backup 紧凑入口面板。
 *
 * 布局：标题行 + 双列（Export | Backup），每列含状态摘要 + 触发按钮。
 * 两个子区域各自独立 query / mutation，互不阻塞。
 */
export function SovereigntyPanel() {
  const queryClient = useQueryClient()

  // ============================================================================
  // Export Snapshot read + mutation
  // ============================================================================
  const exportQuery = useExportSnapshotRead()

  const exportMutation = useMutation({
    mutationFn: () => exportClient.exportCoreAssets({}),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: EXPORT_SNAPSHOT_QUERY_KEY })
    },
  })

  // ============================================================================
  // Backup Snapshot read + mutation
  // ============================================================================
  const backupQuery = useBackupSnapshotRead()

  const backupMutation = useMutation({
    mutationFn: () => backupClient.createInstanceBackup({}),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: BACKUP_SNAPSHOT_QUERY_KEY })
    },
  })

  return (
    <section
      className="rounded-lg border bg-card p-3"
      aria-label="数据主权：导出与备份"
    >
      <div className="grid grid-cols-1 divide-y divide-border md:grid-cols-2 md:divide-x md:divide-y-0">
        {/* Export 子区域 */}
        <div className="space-y-1.5 md:pr-3">
          <div className="flex items-center justify-between gap-2">
            <div className="flex items-center gap-1.5 min-w-0">
              <Download className="h-3.5 w-3.5 shrink-0 text-muted-foreground" />
              <span className="text-xs font-medium">导出</span>
              <ExportStatusBadge snapshot={exportQuery.data?.snapshot ?? null} loading={exportQuery.isLoading} />
            </div>
            <Button
              size="sm"
              variant="outline"
              className="h-7 shrink-0 px-2 text-xs"
              onClick={() => exportMutation.mutate()}
              disabled={exportMutation.isPending}
            >
              {exportMutation.isPending ? '导出中...' : '触发导出'}
            </Button>
          </div>
          <ExportSummaryLine
            snapshot={exportQuery.data?.snapshot ?? null}
            loading={exportQuery.isLoading}
            error={exportQuery.error as Error | null}
            onRetry={() => queryClient.invalidateQueries({ queryKey: EXPORT_SNAPSHOT_QUERY_KEY })}
          />
          {exportMutation.isError && (
            <p className="text-[10px] text-destructive truncate">
              {(exportMutation.error as Error).message}
            </p>
          )}
        </div>

        {/* Backup 子区域 */}
        <div className="space-y-1.5 md:pl-3 md:pt-0 pt-3">
          <div className="flex items-center justify-between gap-2">
            <div className="flex items-center gap-1.5 min-w-0">
              <ShieldCheck className="h-3.5 w-3.5 shrink-0 text-muted-foreground" />
              <span className="text-xs font-medium">备份</span>
              <BackupStatusBadge
                snapshot={backupQuery.data?.snapshot ?? null}
                loading={backupQuery.isLoading}
              />
            </div>
            <Button
              size="sm"
              variant="outline"
              className="h-7 shrink-0 px-2 text-xs"
              onClick={() => backupMutation.mutate()}
              disabled={backupMutation.isPending}
            >
              {backupMutation.isPending ? '备份中...' : '触发备份'}
            </Button>
          </div>
          <BackupSummaryLine
            snapshot={backupQuery.data?.snapshot ?? null}
            loading={backupQuery.isLoading}
            error={backupQuery.error as Error | null}
            onRetry={() => queryClient.invalidateQueries({ queryKey: BACKUP_SNAPSHOT_QUERY_KEY })}
          />
          {backupMutation.isError && (
            <p className="text-[10px] text-destructive truncate">
              {(backupMutation.error as Error).message}
            </p>
          )}
        </div>
      </div>
    </section>
  )
}

// ============================================================================
// Export 子组件
// ============================================================================

/** Export 状态内联标签（紧凑型，单行展示在标题旁） */
function ExportStatusBadge({
  snapshot,
  loading,
}: {
  snapshot: ExportSnapshot | null
  loading: boolean
}) {
  if (loading) return <Skeleton className="h-3.5 w-10" />
  if (!snapshot) return <span className="text-[10px] text-muted-foreground">暂无</span>
  return (
    <span className="text-[10px] text-muted-foreground">
      {snapshot.result_status}
    </span>
  )
}

/** Export 摘要行（单行紧凑展示，可截断） */
function ExportSummaryLine({
  snapshot,
  loading,
  error,
  onRetry,
}: {
  snapshot: ExportSnapshot | null
  loading: boolean
  error: Error | null
  onRetry: () => void
}) {
  if (loading) return <Skeleton className="h-3 w-full" />

  if (error) {
    return (
      <div className="flex items-center gap-1 text-[10px]">
        <span className="text-destructive truncate">读取失败</span>
        <button
          type="button"
          onClick={onRetry}
          className="inline-flex items-center text-primary hover:underline shrink-0"
        >
          <RefreshCw className="h-2.5 w-2.5" />
        </button>
      </div>
    )
  }

  if (!snapshot) {
    return <p className="text-[10px] text-muted-foreground">点击触发首次导出</p>
  }

  return (
    <p className="text-[10px] text-muted-foreground truncate">
      {snapshot.result_summary ?? `${snapshot.result_status} · ${formatTime(snapshot.created_at)}`}
    </p>
  )
}

// ============================================================================
// Backup 子组件
// ============================================================================

/** Backup 状态内联标签（含校验状态 + 可选失败原因，紧凑型） */
function BackupStatusBadge({
  snapshot,
  loading,
}: {
  snapshot: BackupSnapshot | null
  loading: boolean
}) {
  if (loading) return <Skeleton className="h-3.5 w-12" />
  if (!snapshot) return <span className="text-[10px] text-muted-foreground">暂无</span>

  const label = BACKUP_VERIFIED_STATUS_LABELS[snapshot.verified_status]
  const isFailed = snapshot.verified_status === 'verify_failed'

  return (
    <span className={`text-[10px] ${isFailed ? 'text-destructive' : 'text-muted-foreground'}`}>
      {label}
      {isFailed && snapshot.verify_failure_code
        ? ` · ${VERIFY_FAILURE_CODE_LABELS[snapshot.verify_failure_code]}`
        : ''}
    </span>
  )
}

/** Backup 摘要行（单行紧凑展示，含资产条目计数） */
function BackupSummaryLine({
  snapshot,
  loading,
  error,
  onRetry,
}: {
  snapshot: BackupSnapshot | null
  loading: boolean
  error: Error | null
  onRetry: () => void
}) {
  if (loading) return <Skeleton className="h-3 w-full" />

  if (error) {
    return (
      <div className="flex items-center gap-1 text-[10px]">
        <span className="text-destructive truncate">读取失败</span>
        <button
          type="button"
          onClick={onRetry}
          className="inline-flex items-center text-primary hover:underline shrink-0"
        >
          <RefreshCw className="h-2.5 w-2.5" />
        </button>
      </div>
    )
  }

  if (!snapshot) {
    return <p className="text-[10px] text-muted-foreground">点击触发首次备份</p>
  }

  const coverage = snapshot.manifest_summary
    ? `${snapshot.manifest_summary.covered_asset_entries}/${snapshot.manifest_summary.total_asset_entries} 条目`
    : ''

  return (
    <p className="text-[10px] text-muted-foreground truncate">
      {coverage}{coverage && ' · '}{formatTime(snapshot.created_at)}
    </p>
  )
}

// ============================================================================
// 工具函数
// ============================================================================

/** 时间戳格式化为简短的 MM/DD HH:mm 格式 */
function formatTime(isoString: string | undefined | null): string {
  if (!isoString) return '—'
  const d = new Date(isoString)
  if (isNaN(d.getTime())) return '—'
  const mm = String(d.getMonth() + 1).padStart(2, '0')
  const dd = String(d.getDate()).padStart(2, '0')
  const hh = String(d.getHours()).padStart(2, '0')
  const mi = String(d.getMinutes()).padStart(2, '0')
  return `${mm}/${dd} ${hh}:${mi}`
}
