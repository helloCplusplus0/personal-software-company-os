/**
 * useRecentActivitiesRead — Recent Activities 只读 query owner。
 *
 * phase07-10 §5.4：query 层纯只读，唯一 owner。
 * 替换 dashboard/data/api-adapter.ts 的 fetchRecentActivities。
 */
import { useQuery, type UseQueryResult } from '@tanstack/react-query'
import { timestampDate } from '@bufbuild/protobuf/wkt'
import { dashboardClient } from './connect-client'
import { DashboardTargetType, RecentActivityType } from '@/gen/proto/psco/dashboard/v1/dashboard_pb'
import type {
  DashboardTargetType as DashboardTargetTypeValue,
  RecentActivitiesResponse,
  RecentActivityType as RecentActivityTypeValue,
} from '../types'

export type UseRecentActivitiesRead = UseQueryResult<RecentActivitiesResponse, Error>

function mapActivityType(v: RecentActivityType): RecentActivityTypeValue {
  switch (v) {
    case RecentActivityType.MODULE:
      return 'module'
    case RecentActivityType.RELEASE:
      return 'release'
    case RecentActivityType.PRODUCT:
      return 'product'
    case RecentActivityType.REPOSITORY:
      return 'repository'
    case RecentActivityType.DECISION:
      return 'decision'
    case RecentActivityType.PRODUCT_MODULE_BINDING:
      return 'product_module_binding'
    case RecentActivityType.PRODUCT_REPOSITORY_BINDING:
      return 'product_repository_binding'
    case RecentActivityType.MODULE_REPOSITORY_BINDING:
      return 'module_repository_binding'
    default:
      return 'module'
  }
}

function mapTargetType(v: DashboardTargetType): DashboardTargetTypeValue {
  switch (v) {
    case DashboardTargetType.DECISION_DETAIL:
      return 'decision_detail'
    case DashboardTargetType.DECISION_LIST:
      return 'decision_list'
    case DashboardTargetType.PRODUCT_DETAIL:
      return 'product_detail'
    case DashboardTargetType.MODULE_DETAIL:
      return 'module_detail'
    case DashboardTargetType.REPOSITORY_DETAIL:
      return 'repository_detail'
    default:
      return 'decision_list'
  }
}

export function useRecentActivitiesRead(): UseRecentActivitiesRead {
  return useQuery<RecentActivitiesResponse, Error>({
    queryKey: ['dashboard-recent-activities'],
    queryFn: async (): Promise<RecentActivitiesResponse> => {
      const res = await dashboardClient.getRecentActivities({})
      return {
        activities: (res.activities ?? []).map((a) => ({
          activity_type: mapActivityType(a.activityType),
          activity_at: a.activityAt ? timestampDate(a.activityAt).toISOString() : '',
          target_type: mapTargetType(a.targetType),
          target_id: a.targetId ?? '',
          target_label: a.targetLabel ?? '',
        })),
      } as unknown as RecentActivitiesResponse
    },
  })
}
