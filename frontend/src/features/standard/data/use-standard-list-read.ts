/**
 * useStandardListRead — Standard List 只读 query owner。
 *
 * phase14-05 §"切片结构必须冻结"（project_rules §2.5）：query 层纯只读，唯一 owner。
 * ListStandards 无参数不分页，直接投影为 domain Standard[]（按 updated_at DESC）。
 */
import { useQuery, type UseQueryResult } from '@tanstack/react-query'
import { standardClient } from './connect-client'
import { pbToStandard } from '../types'
import type { Standard } from '../types'

export type UseStandardListRead = UseQueryResult<Standard[], Error>

export function useStandardListRead(): UseStandardListRead {
  return useQuery<Standard[], Error>({
    queryKey: ['standard-list'],
    queryFn: async (): Promise<Standard[]> => {
      const res = await standardClient.listStandards({})
      return (res.standards ?? [])
        .map((s) => pbToStandard(s))
        .filter((s): s is Standard => s !== null)
    },
  })
}
