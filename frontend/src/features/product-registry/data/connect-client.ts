/**
 * Product Registry slice-local Connect client。
 *
 * phase07-10 §5.3：slice-local generated client 唯一承接位。
 */
import { createClient } from '@connectrpc/connect'
import { ProductRegistryService } from '@/gen/proto/psco/product_registry/v1/product_registry_pb'
import { transport } from '@/shared/rpc/connect-transport'

export const productRegistryClient = createClient(ProductRegistryService, transport)