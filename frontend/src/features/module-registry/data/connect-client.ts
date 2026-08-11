/**
 * Module Registry slice-local Connect client。
 *
 * phase07-10 §5.3：slice-local generated client 唯一承接位。
 * 页面与组件不得直接 createClient() 或 import transport。
 */
import { createClient } from '@connectrpc/connect'
import { ModuleRegistryService } from '@/gen/proto/psco/module_registry/v1/module_registry_pb'
import { transport } from '@/shared/rpc/connect-transport'

export const moduleRegistryClient = createClient(ModuleRegistryService, transport)