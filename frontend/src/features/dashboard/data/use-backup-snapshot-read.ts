/**
 * useBackupSnapshotRead — Backup Snapshot 只读 query owner。
 *
 * phase07-10 §5.4：query 层纯只读，唯一 owner。
 * 替换 dashboard/data/sovereignty-api-adapter.ts 的 fetchBackupSnapshot。
 */
import { useQuery, type UseQueryResult } from '@tanstack/react-query'
import { timestampDate } from '@bufbuild/protobuf/wkt'
import { backupClient } from './connect-client'
import type { BackupSnapshotReadResult } from '../types'

export type UseBackupSnapshotRead = UseQueryResult<BackupSnapshotReadResult, Error>

export function useBackupSnapshotRead(): UseBackupSnapshotRead {
  return useQuery<BackupSnapshotReadResult, Error>({
    queryKey: ['backup-snapshot'],
    queryFn: async (): Promise<BackupSnapshotReadResult> => {
      const res = await backupClient.getBackupSnapshot({})
      const s = res.snapshot
      if (!s) return { snapshot: null }
      return {
        snapshot: {
          created_at: s.createdAt ? timestampDate(s.createdAt).toISOString() : '',
          manifest_summary: s.manifestSummary ? {
            manifest_version: s.manifestSummary.manifestVersion ?? '',
            total_asset_entries: s.manifestSummary.totalAssetEntries ?? 0,
            covered_asset_entries: s.manifestSummary.coveredAssetEntries ?? 0,
          } : null,
          asset_coverage: (s.assetCoverage ?? []).map((c) => ({
            asset_scope: c.assetScope ?? '',
            covered: c.covered ?? false,
          })),
          schema_version_prerequisite: s.schemaVersionPrerequisite ? {
            schema_version: s.schemaVersionPrerequisite.schemaVersion ?? '',
            instance_version: s.schemaVersionPrerequisite.instanceVersion ?? '',
            prerequisite_checkable: s.schemaVersionPrerequisite.prerequisiteCheckable ?? false,
          } : null,
          verified_status: s.verifiedStatus,
          verify_failure_code: s.verifyFailureCode,
        },
      } as unknown as BackupSnapshotReadResult
    },
  })
}