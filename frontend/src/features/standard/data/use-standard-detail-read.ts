/**
 * useStandardDetailRead — Standard Detail 只读 query owner。
 *
 * phase14-05 §"切片结构必须冻结"（project_rules §2.5）：query 层纯只读，唯一 owner。
 * GetStandard 一次返回主体与绑定集合，StandardDetailPage 与绑定管理区直接消费。
 */
import { useQuery, type UseQueryResult } from '@tanstack/react-query'
import { standardClient } from './connect-client'
import { pbToStandard, pbToBinding } from '../types'
import type { StandardDetail } from '../types'

export type UseStandardDetailRead = UseQueryResult<StandardDetail, Error>

export function useStandardDetailRead(standardId: string): UseStandardDetailRead {
  return useQuery<StandardDetail, Error>({
    queryKey: ['standard-detail', standardId],
    queryFn: async (): Promise<StandardDetail> => {
      const res = await standardClient.getStandard({ standardId })
      const standard = pbToStandard(res.standard)
      if (!standard) throw new Error('standard not found')
      return {
        standard,
        bindings: (res.bindings ?? []).map((b) => pbToBinding(b)),
      }
    },
    enabled: !!standardId,
  })
}
