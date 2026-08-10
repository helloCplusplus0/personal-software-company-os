/**
 * Export / Backup API 适配层
 *
 * 对接 phase06-14 后端 export / backup 模块：
 *   GET  /api/dashboard/export  → GetExportSnapshot
 *   POST /api/dashboard/export  → ExportCoreAssets
 *   GET  /api/dashboard/backup  → GetBackupSnapshot (read / verify)
 *   POST /api/dashboard/backup  → CreateInstanceBackup
 *
 * 字段语义从 .proto -> HTTP DTO 单向派生（phase06-05 / phase06-13 合同约束）。
 * 后端已使用 snake_case，与前端类型一致，无需转换。
 */

const API_BASE_URL = import.meta.env.VITE_API_BASE_URL ?? ''

// ============================================================================
// Export 类型（对齐 proto export/v1/export.proto）
// ============================================================================

export type ExportAssetScope =
  | 'products'
  | 'modules'
  | 'releases'
  | 'repositories'
  | 'decisions'
  | 'decision_links'
  | 'product_modules'
  | 'product_repositories'
  | 'module_repositories'

export type ExportResultStatus = 'success' | 'in_progress' | 'failed'

export interface ExportSnapshot {
  asset_scope: ExportAssetScope[]
  created_at: string
  result_status: ExportResultStatus
  result_summary: string
}

export interface ExportSnapshotReadResult {
  snapshot: ExportSnapshot | null
}

// ============================================================================
// Backup 类型（对齐 proto backup/v1/backup.proto）
// ============================================================================

export type BackupAssetScope =
  | 'products'
  | 'modules'
  | 'releases'
  | 'repositories'
  | 'decisions'
  | 'decision_links'
  | 'product_modules'
  | 'product_repositories'
  | 'module_repositories'

export type BackupVerifiedStatus = 'unverified' | 'verified' | 'verify_failed'

export type VerifyFailureCode = 'manifest_missing' | 'coverage_incomplete' | 'schema_mismatch'

export interface ManifestSummary {
  manifest_version: string
  total_asset_entries: number
  covered_asset_entries: number
}

export interface AssetCoverageEntry {
  asset_scope: BackupAssetScope
  covered: boolean
}

export interface SchemaVersionPrerequisite {
  schema_version: string
  instance_version: string
  prerequisite_checkable: boolean
}

export interface BackupSnapshot {
  created_at: string
  manifest_summary: ManifestSummary | null
  asset_coverage: AssetCoverageEntry[]
  schema_version_prerequisite: SchemaVersionPrerequisite | null
  verified_status: BackupVerifiedStatus
  /** 仅在 verified_status = verify_failed 时有值 */
  verify_failure_code?: VerifyFailureCode
}

export interface BackupSnapshotReadResult {
  snapshot: BackupSnapshot | null
}

// ============================================================================
// API 函数
// ============================================================================

/**
 * fetchExportSnapshot — GetExportSnapshot 读组
 * GET /api/dashboard/export
 * 无历史记录时返回预览态快照（result_status = success, result_summary 含 preview 标识）
 */
export async function fetchExportSnapshot(): Promise<ExportSnapshotReadResult> {
  const res = await fetch(`${API_BASE_URL}/api/dashboard/export`, {
    headers: { Accept: 'application/json' },
  })
  if (!res.ok) {
    const msg = await extractErrorMessage(res)
    throw new Error(msg)
  }
  return res.json() as Promise<ExportSnapshotReadResult>
}

/**
 * triggerExportCoreAssets — ExportCoreAssets 写组
 * POST /api/dashboard/export
 * 装配 9 类核心资产并持久化到 instance_exports，返回正式导出快照
 */
export async function triggerExportCoreAssets(): Promise<ExportSnapshotReadResult> {
  const res = await fetch(`${API_BASE_URL}/api/dashboard/export`, {
    method: 'POST',
    headers: { Accept: 'application/json' },
  })
  if (!res.ok) {
    const msg = await extractErrorMessage(res)
    throw new Error(msg)
  }
  return res.json() as Promise<ExportSnapshotReadResult>
}

/**
 * fetchBackupSnapshot — GetBackupSnapshot 读组（read / verify 子路径）
 * GET /api/dashboard/backup
 * 返回最新备份快照并附带校验结果（verified / verify_failed + failure_code）
 * 无历史备份记录时返回 snapshot: null
 */
export async function fetchBackupSnapshot(): Promise<BackupSnapshotReadResult> {
  const res = await fetch(`${API_BASE_URL}/api/dashboard/backup`, {
    headers: { Accept: 'application/json' },
  })
  if (!res.ok) {
    const msg = await extractErrorMessage(res)
    throw new Error(msg)
  }
  return res.json() as Promise<BackupSnapshotReadResult>
}

/**
 * triggerCreateInstanceBackup — CreateInstanceBackup 写组
 * POST /api/dashboard/backup
 * 装配 9 类资产 + manifest + coverage + schema_version，持久化到 instance_backups
 */
export async function triggerCreateInstanceBackup(): Promise<BackupSnapshotReadResult> {
  const res = await fetch(`${API_BASE_URL}/api/dashboard/backup`, {
    method: 'POST',
    headers: { Accept: 'application/json' },
  })
  if (!res.ok) {
    const msg = await extractErrorMessage(res)
    throw new Error(msg)
  }
  return res.json() as Promise<BackupSnapshotReadResult>
}

// ============================================================================
// 辅助
// ============================================================================

async function extractErrorMessage(res: Response): Promise<string> {
  let message = `HTTP ${res.status}`
  try {
    const body = await res.json()
    if (body?.error) {
      message = body.error
    }
  } catch {
    // 响应体非 JSON，保留默认 message
  }
  return message
}
