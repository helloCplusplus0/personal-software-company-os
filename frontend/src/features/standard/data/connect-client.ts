/**
 * Standard slice-local Connect client。
 *
 * phase14-05 §"切片结构必须冻结"：slice-local generated client 唯一承接位。
 * standardClient 承接 StandardService 8 RPC（写读正式主线）；
 * projectContextClient 承接 GetProjectBrief（Repository standards 摘要数据源）。
 * 页面与组件不得直接 createClient() 或 import transport。
 */
import { createClient } from '@connectrpc/connect'
import { StandardService } from '@/gen/proto/psco/standard/v1/standard_pb'
import { ProjectContextService } from '@/gen/proto/psco/project_context/v1/project_context_pb'
import { transport } from '@/shared/rpc/connect-transport'

export const standardClient = createClient(StandardService, transport)

export const projectContextClient = createClient(ProjectContextService, transport)
