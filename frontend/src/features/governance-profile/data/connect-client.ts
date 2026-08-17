/**
 * Governance Profile slice-local Connect client。
 *
 * phase13-09：治理画像前端唯一 transport 消费承接位。
 * 页面与组件不得直接 import 生成 client 或 transport。
 */
import { createClient } from '@connectrpc/connect'
import { GovernanceProfileService } from '@/gen/proto/psco/governance_profile/v1/governance_profile_pb'
import { transport } from '@/shared/rpc/connect-transport'

export const governanceProfileClient = createClient(GovernanceProfileService, transport)
