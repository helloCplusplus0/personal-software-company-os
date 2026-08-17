/**
 * useGovernanceProfileRead — 治理画像只读 query owner。
 *
 * phase13-09 冻结：
 *   - 唯一治理画像前端读取承接位，以 repository_id 为唯一正式锚点
 *   - query 层纯只读：只承接读取、缓存键与只读解包，不混入写动作
 *   - 页面与展示组件不得各自拼装读取逻辑
 *
 * 失败语义（对齐 phase13-08 后端合同）：
 *   - repository 不存在或画像未创建 → Connect Code.NotFound
 *     （画像未创建是合法空态，由组件层区分，本层不吞错）
 *   - 其他读取失败 → 以 UseQueryResult.error 暴露
 */
import { useQuery, type UseQueryResult } from '@tanstack/react-query'
import { governanceProfileClient } from './connect-client'
import type { GovernanceProfile } from '@/gen/proto/psco/governance_profile/v1/governance_profile_pb'

export type UseGovernanceProfileRead = UseQueryResult<GovernanceProfile, Error>

export function useGovernanceProfileRead(repositoryId: string): UseGovernanceProfileRead {
  return useQuery<GovernanceProfile, Error>({
    queryKey: ['governance-profile', repositoryId],
    queryFn: async () => {
      const res = await governanceProfileClient.getGovernanceProfile({ repositoryId })
      // 后端合同保证成功响应携带 profile；缺失视为异常合同状态，不静默降级
      if (!res.profile) {
        throw new Error('治理画像读取响应缺少 profile 字段')
      }
      return res.profile
    },
    enabled: !!repositoryId,
  })
}
