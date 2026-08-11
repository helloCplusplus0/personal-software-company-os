/**
 * Review slice-local Connect client。
 *
 * phase08-08 §"review read owner 的正式 transport 路径"：
 *   review 前端 read owner 必须以 ReviewService generated client 为正式 transport 入口，
 *   而不是在页面层直接并排消费 dashboard / decision-center / reuse-summary 的底层 query hook。
 */
import { createClient } from '@connectrpc/connect'
import { ReviewService } from '@/gen/proto/psco/review/v1/review_pb'
import { transport } from '@/shared/rpc/connect-transport'

export const reviewClient = createClient(ReviewService, transport)