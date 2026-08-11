/**
 * 前端唯一共享 Connect transport。
 *
 * phase07-10 §5.2：前端正式传输主线唯一落点。
 * 所有业务切片通过 slice-local generated client 消费本 transport，
 * 页面与组件不得直接 import 本文件或 createConnectTransport()。
 *
 * 单一 /api 基址：
 *   - 开发环境：Vite proxy /api → localhost:8081
 *   - 验收环境：本地后端 :8081
 *   - 部署环境：Caddy 反代 /api → 后端 :8081
 */
import { createConnectTransport } from '@connectrpc/connect-web'

export const transport = createConnectTransport({ baseUrl: '/api' })