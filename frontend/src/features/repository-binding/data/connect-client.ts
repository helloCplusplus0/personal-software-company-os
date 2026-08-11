/**
 * Repository Binding slice-local Connect client。
 *
 * phase07-10 §5.3：slice-local generated client 唯一承接位。
 */
import { createClient } from '@connectrpc/connect'
import { RepositoryBindingService } from '@/gen/proto/psco/repository_binding/v1/repository_binding_pb'
import { transport } from '@/shared/rpc/connect-transport'

export const repositoryBindingClient = createClient(RepositoryBindingService, transport)