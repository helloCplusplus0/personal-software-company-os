/**
 * SovereigntyPanel — Export / Backup 用户入口面板
 *
 * phase06-15 §"Export / Backup 必须作为 Dashboard 内正式用户入口"
 *
 * 约束：
 *   - Dashboard 中必须存在稳定可见的 Export 与 Backup 入口
 *   - 用户可以读取当前 ExportSnapshot 与 BackupSnapshot
 *   - 用户可以从前端触发导出与备份动作
 *   - Backup 的三类失败语义必须继续可区分：
 *     manifest_missing / coverage_incomplete / schema_mismatch
 *   - 当前阶段不得把 Export / Backup 做成隐藏路由、开发者入口或仅测试按钮
 *
 * 字段语义从 .proto -> HTTP DTO 单向派生（phase06-05 / phase06-13 合同约束）。
 */
import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query'
import { Button } from '@/components/ui/button'
import { Skeleton } from '@/components/ui/skeleton'
import { Download, ShieldCheck, RefreshCw } from 'lucide-react'
import {
  fetchExportSnapshot,
  triggerExportCoreAssets,
  fetchBackupSnapshot,
  triggerCreateInstanceBackup,
  type ExportSnapshot,
  type BackupSnapshot,
  type VerifyFailureCode,
} from '../data/sovereignty-api-adapter'

const EXPORT_SNAPSHOT_QUERY_KEY = ['export-snapshot'] as const
const BACKUP_SNAPSHOT_QUERY_KEY = ['backup-snapshot'] as const

/** Backup 三类失败语义的中文映射（phase06-15 §"Backup 三类失败语义"） */
const VERIFY_FAILURE_CODE_LABELS: Record<VerifyFailureCode, string> = {
  manifest_missing: 'manifest 缺失',
  coverage_incomplete: '资产覆盖不完整',
  schema_mismatch: 'schema 版本不匹配',
}

/** Backup verified 状态的中文映射 */
const BACKUP_VERIFIED_STATUS_LABELS: Record<BackupSnapshot['verified_status'], string> = {
  unverified: '未校验',
  verified: '已校验',
  verify_failed: '校验失败',
}

/**
 * SovereigntyPanel — Export / Backup 入口面板。
 *
 * 包含两个并列子区域：
 *   1. Export Snapshot：展示最近一次导出快照 + 触发新导出
 *   2. Backup Snapshot：展示最近一次备份快照 + 触发新备份 + 校验状态
 *
 * 两个子区域各自独立 query / mutation，互不阻塞。
 * 任一 query 失败不回退整页，只展示局部错误。
 */
export function SovereigntyPanel() {
  const queryClient = useQueryClient()

  // ============================================================================
  // Export Snapshot query + mutation
  // ============================================================================
  const exportQuery = useQuery({
    queryKey: EXPORT_SNAPSHOT_QUERY_KEY,
    queryFn: fetchExportSnapshot,
  })

  const exportMutation = useMutation({
    mutationFn: triggerExportCoreAssets,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: EXPORT_SNAPSHOT_QUERY_KEY })
    },
  })

  // ============================================================================
  // Backup Snapshot query + mutation
  // ============================================================================
  const backupQuery = useQuery({
    queryKey: BACKUP_SNAPSHOT_QUERY_KEY,
    queryFn: fetchBackupSnapshot,
  })

  const backupMutation = useMutation({
    mutationFn: triggerCreateInstanceBackup,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: BACKUP_SNAPSHOT_QUERY_KEY })
    },
  })

  return (
    <section
      className="rounded-lg border bg-card p-4 space-y-4"
      aria-label="数据主权：导出与备份"
    >
      <div>
        <h2 className="text-base font-semibold">数据主权</h2>
        <p className="text-xs text-muted-foreground mt-1">
          导出核心资产快照或创建完整实例备份，确保数据所有权始终掌握在你手中。
        </p>
      </div>

      <div className="grid gap-4 md:grid-cols-2">
        {/* Export 子区域 */}
        <div className="space-y-2">
          <div className="flex items-center gap-2">
            <Download className="h-4 w-4 text-muted-foreground" />
            <h3 className="text-sm font-medium">导出（Export）</h3>
          </div>
          <ExportSnapshotView
            snapshot={exportQuery.data?.snapshot ?? null}
            loading={exportQuery.isLoading}
            error={exportQuery.error as Error | null}
            onRetry={() =>
              queryClient.invalidateQueries({ queryKey: EXPORT_SNAPSHOT_QUERY_KEY })
            }
          />
          <Button
            size="sm"
            className="w-full"
            onClick={() => exportMutation.mutate()}
            disabled={exportMutation.isPending}
          >
            {exportMutation.isPending ? '导出中...' : '触发导出'}
          </Button>
          {exportMutation.isError && (
            <p className="text-xs text-destructive">
              导出失败：{(exportMutation.error as Error).message}
            </p>
          )}
        </div>

        {/* Backup 子区域 */}
        <div className="space-y-2">
          <div className="flex items-center gap-2">
            <ShieldCheck className="h-4 w-4 text-muted-foreground" />
            <h3 className="text-sm font-medium">备份（Backup）</h3>
          </div>
          <BackupSnapshotView
            snapshot={backupQuery.data?.snapshot ?? null}
            loading={backupQuery.isLoading}
            error={backupQuery.error as Error | null}
            onRetry={() =>
              queryClient.invalidateQueries({ queryKey: BACKUP_SNAPSHOT_QUERY_KEY })
            }
          />
          <Button
            size="sm"
            className="w-full"
            onClick={() => backupMutation.mutate()}
            disabled={backupMutation.isPending}
          >
            {backupMutation.isPending ? '备份中...' : '触发备份'}
          </Button>
          {backupMutation.isError && (
            <p className="text-xs text-destructive">
              备份失败：{(backupMutation.error as Error).message}
            </p>
          )}
        </div>
      </div>
    </section>
  )
}

/**
 * ExportSnapshotView — Export 快照展示子组件。
 */
function ExportSnapshotView({
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
  if (loading) {
    return <Skeleton className="h-20 w-full" />
  }

  if (error) {
    return (
      <div className="rounded-md border border-destructive/30 bg-destructive/5 p-2 text-xs">
        <p className="text-destructive mb-1">读取失败：{error.message}</p>
        <button
          type="button"
          onClick={onRetry}
          className="inline-flex items-center gap-1 text-primary hover:underline"
        >
          <RefreshCw className="h-3 w-3" />
          重试
        </button>
      </div>
    )
  }

  if (!snapshot) {
    return (
      <div className="rounded-md border border-dashed p-3 text-xs text-muted-foreground">
        暂无导出记录。点击下方按钮触发首次导出。
      </div>
    )
  }

  return (
    <div className="rounded-md border bg-muted/20 p-3 text-xs space-y-1">
      <div className="flex items-center justify-between">
        <span className="text-muted-foreground">状态</span>
        <span className="font-medium">{snapshot.result_status}</span>
      </div>
      <div className="flex items-center justify-between">
        <span className="text-muted-foreground">创建时间</span>
        <span>{snapshot.created_at ? new Date(snapshot.created_at).toLocaleString() : '—'}</span>
      </div>
      {snapshot.result_summary && (
        <p className="text-muted-foreground pt-1 border-t">{snapshot.result_summary}</p>
      )}
    </div>
  )
}

/**
 * BackupSnapshotView — Backup 快照展示子组件。
 *
 * 必须可区分三类校验失败语义（phase06-15 §"Backup 三类失败语义"）：
 *   - manifest_missing
 *   - coverage_incomplete
 *   - schema_mismatch
 */
function BackupSnapshotView({
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
  if (loading) {
    return <Skeleton className="h-20 w-full" />
  }

  if (error) {
    return (
      <div className="rounded-md border border-destructive/30 bg-destructive/5 p-2 text-xs">
        <p className="text-destructive mb-1">读取失败：{error.message}</p>
        <button
          type="button"
          onClick={onRetry}
          className="inline-flex items-center gap-1 text-primary hover:underline"
        >
          <RefreshCw className="h-3 w-3" />
          重试
        </button>
      </div>
    )
  }

  if (!snapshot) {
    return (
      <div className="rounded-md border border-dashed p-3 text-xs text-muted-foreground">
        暂无备份记录。点击下方按钮触发首次备份。
      </div>
    )
  }

  return (
    <div className="rounded-md border bg-muted/20 p-3 text-xs space-y-1">
      <div className="flex items-center justify-between">
        <span className="text-muted-foreground">校验状态</span>
        <span className="font-medium">
          {BACKUP_VERIFIED_STATUS_LABELS[snapshot.verified_status]}
        </span>
      </div>
      {/* phase06-15 §"Backup 三类失败语义"：verify_failed 时显式展示 failure_code */}
      {snapshot.verified_status === 'verify_failed' && snapshot.verify_failure_code && (
        <div className="flex items-center justify-between">
          <span className="text-muted-foreground">失败原因</span>
          <span className="font-medium text-destructive">
            {VERIFY_FAILURE_CODE_LABELS[snapshot.verify_failure_code]}
          </span>
        </div>
      )}
      <div className="flex items-center justify-between">
        <span className="text-muted-foreground">创建时间</span>
        <span>{snapshot.created_at ? new Date(snapshot.created_at).toLocaleString() : '—'}</span>
      </div>
      {snapshot.manifest_summary && (
        <div className="flex items-center justify-between">
          <span className="text-muted-foreground">资产条目</span>
          <span>
            {snapshot.manifest_summary.covered_asset_entries} /{' '}
            {snapshot.manifest_summary.total_asset_entries}
          </span>
        </div>
      )}
    </div>
  )
}
