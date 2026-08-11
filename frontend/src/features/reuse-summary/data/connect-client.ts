/**
 * Reuse Summary slice-local Connect client。
 *
 * phase07-10 §5.3：slice-local generated client 唯一承接位。
 */
import { createClient } from '@connectrpc/connect'
import { ReuseSummaryService } from '@/gen/proto/psco/reuse_summary/v1/reuse_summary_pb'
import { transport } from '@/shared/rpc/connect-transport'

export const reuseSummaryClient = createClient(ReuseSummaryService, transport)