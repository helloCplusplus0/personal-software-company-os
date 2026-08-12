/**
 * Template Reuse slice-local Connect client。
 *
 * phase09-08 spec §"template reuse slice-local connect client 的正式 transport 路径"：
 *   template reuse 前端 read owner 必须以 TemplateReuseService generated client 为正式 transport 入口，
 *   而不是在页面层直接 createClient(TemplateReuseService, ...)。
 */
import { createClient } from '@connectrpc/connect'
import { TemplateReuseService } from '@/gen/proto/psco/template_reuse/v1/template_reuse_pb'
import { transport } from '@/shared/rpc/connect-transport'

export const templateReuseClient = createClient(TemplateReuseService, transport)