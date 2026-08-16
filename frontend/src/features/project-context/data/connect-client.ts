/**
 * project-context slice-local Connect client。
 *
 * phase12-09：唯一跨切片共享只读 Connect client 承接位。
 * 只承接 ProjectContextService 的只读 RPC，不承接写入。
 */
import { createClient } from '@connectrpc/connect'
import { ProjectContextService } from '@/gen/proto/psco/project_context/v1/project_context_pb'
import { transport } from '@/shared/rpc/connect-transport'

export const projectContextClient = createClient(ProjectContextService, transport)