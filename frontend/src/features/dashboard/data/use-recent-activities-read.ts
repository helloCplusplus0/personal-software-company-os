/**
 * useRecentActivitiesRead — Recent Activities 只读 query owner。
 *
 * phase07-10 §5.4：query 层纯只读，唯一 owner。
 * 替换 dashboard/data/api-adapter.ts 的 fetchRecentActivities。
 */
import { useQuery, type UseQueryResult } from '@tanstack/react-query'
import { timestampDate } from '@bufbuild/protobuf/wkt'
import { dashboardClient } from './connect-client'
import type { RecentActivitiesResponse } from '../types'

export type UseRecentActivitiesRead = UseQueryResult<RecentActivitiesResponse, Error>

export function useRecentActivitiesRead(): UseRecentActivitiesRead {
  return useQuery<RecentActivitiesResponse, Error>({
    queryKey: ['dashboard-recent-activities'],
    queryFn: async (): Promise<RecentActivitiesResponse> => {
      const res = await dashboardClient.getRecentActivities({})
      return {
        activities: (res.activities ?? []).map((a) => ({
          activity_type: a.activityType ?? '',
          activity_at: a.activityAt ? timestampDate(a.activityAt).toISOString() : '',
          target_type: a.targetType,
          target_id: a.targetId ?? '',
          target_label: a.targetLabel ?? '',
        })),
      } as unknown as RecentActivitiesResponse
    },
  })
}