/**
 * 前端唯一共享 Connect transport。
 *
 * phase07-10 §5.2：前端正式传输主线唯一落点。
 * 所有业务切片通过 slice-local generated client 消费本 transport，
 * 页面与组件不得直接 import 本文件或 createConnectTransport()。
 *
 * 默认仍使用单一 /api 基址：
 *   - 开发环境：Vite proxy /api → localhost:8081
 *   - 部署环境：Caddy 反代 /api → 后端 :8081
 *
 * 当本地需要隔离验收环境时，允许通过 VITE_API_BASE_URL 显式指定同一条
 * canonical API 根路径，例如 http://127.0.0.1:8082 或 http://127.0.0.1:8082/api。
 * 该变量只改变环境落点，不引入第二套 transport。
 */
import { createConnectTransport } from '@connectrpc/connect-web'

function resolveApiBaseUrl() {
  const raw = import.meta.env.VITE_API_BASE_URL?.trim()
  if (!raw) {
    return '/api'
  }

  // 允许传 origin 或显式 /api 根路径，两者都统一收敛到 canonical /api。
  return new URL('/api', raw.endsWith('/') ? raw : `${raw}/`).toString().replace(/\/$/, '')
}

export const transport = createConnectTransport({ baseUrl: resolveApiBaseUrl() })
