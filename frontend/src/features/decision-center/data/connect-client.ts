/**
 * Decision Center slice-local Connect client。
 *
 * phase07-10 §5.3：slice-local generated client 唯一承接位。
 */
import { createClient } from '@connectrpc/connect'
import { DecisionCenterService } from '@/gen/proto/psco/decision_center/v1/decision_center_pb'
import { transport } from '@/shared/rpc/connect-transport'

export const decisionCenterClient = createClient(DecisionCenterService, transport)