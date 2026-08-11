/**
 * useDashboardOverviewRead — Dashboard Overview 只读 query owner。
 *
 * phase07-10 §5.4：query 层纯只读，唯一 owner。
 * 替换 dashboard/data/api-adapter.ts 的 fetchDashboardOverview。
 */
import { useQuery, type UseQueryResult } from '@tanstack/react-query'
import { dashboardClient } from './connect-client'
import type { DashboardOverview } from '../types'

export type UseDashboardOverviewRead = UseQueryResult<DashboardOverview, Error>

export function useDashboardOverviewRead(): UseDashboardOverviewRead {
  return useQuery<DashboardOverview, Error>({
    queryKey: ['dashboard-overview'],
    queryFn: async () => {
      const res = await dashboardClient.getDashboardOverview({})
      const overview = res.overview
      return {
        module_count: overview?.moduleCount ?? 0,
        product_count: overview?.productCount ?? 0,
        repository_count: overview?.repositoryCount ?? 0,
        decision_count: overview?.decisionCount ?? 0,
        product_with_repository_count: overview?.productWithRepositoryCount ?? 0,
        product_with_module_count: overview?.productWithModuleCount ?? 0,
      }
    },
  })
}