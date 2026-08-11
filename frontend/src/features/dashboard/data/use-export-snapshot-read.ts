/**
 * useExportSnapshotRead — Export Snapshot 只读 query owner。
 *
 * phase07-10 §5.4：query 层纯只读，唯一 owner。
 * 替换 dashboard/data/sovereignty-api-adapter.ts 的 fetchExportSnapshot。
 */
import { useQuery, type UseQueryResult } from '@tanstack/react-query'
import { timestampDate } from '@bufbuild/protobuf/wkt'
import { exportClient } from './connect-client'
import type { ExportSnapshotReadResult } from '../types'

export type UseExportSnapshotRead = UseQueryResult<ExportSnapshotReadResult, Error>

export function useExportSnapshotRead(): UseExportSnapshotRead {
  return useQuery<ExportSnapshotReadResult, Error>({
    queryKey: ['export-snapshot'],
    queryFn: async (): Promise<ExportSnapshotReadResult> => {
      const res = await exportClient.getExportSnapshot({})
      const s = res.snapshot
      if (!s) return { snapshot: null }
      return {
        snapshot: {
          asset_scope: s.assetScope ?? [],
          created_at: s.createdAt ? timestampDate(s.createdAt).toISOString() : '',
          result_status: s.resultStatus,
          result_summary: s.resultSummary ?? '',
        },
      } as unknown as ExportSnapshotReadResult
    },
  })
}