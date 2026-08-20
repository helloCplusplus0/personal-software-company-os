/**
 * Progress slice-local Connect client。
 *
 * phase15-05 §"切片结构必须冻结"：slice-local generated client 唯一承接位。
 * progressClient 承接 ProgressService 3 RPC（事件流写读正式主线）；
 * projectContextClient 承接 GetProjectBrief（DP-1 当前卡 brief 投影数据源）。
 * 页面与组件不得直接 createClient() 或 import transport。
 */
import { createClient } from '@connectrpc/connect'
import { ProgressService } from '@/gen/proto/psco/progress/v1/progress_pb'
import { ProjectContextService } from '@/gen/proto/psco/project_context/v1/project_context_pb'
import { transport } from '@/shared/rpc/connect-transport'

export const progressClient = createClient(ProgressService, transport)

export const projectContextClient = createClient(ProjectContextService, transport)
