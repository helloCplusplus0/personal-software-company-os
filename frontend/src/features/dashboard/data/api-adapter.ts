/**
 * Dashboard API 适配层
 *
 * phase05-13 §"Dashboard 类型与 API 适配层必须从 proto envelope 单向派生"
 * phase05-12 已落地三个 GET endpoint：
 *   - GET /api/dashboard/overview
 *   - GET /api/dashboard/feedback-signals
 *   - GET /api/dashboard/recent-activities
 *
 * 复用 module-registry/data/api-adapter.ts 的 request<T> 封装模式：
 *   - GET 请求不设 Content-Type，避免触发 CORS preflight
 *   - 非 2xx 响应解析后端 { error } 结构并抛出 ApiError
 *   - 204 No Content 返回 undefined
 *
 * 只读边界（phase05-13 §"Dashboard 前端必须只读不写"）：
 *   - 只导出 fetch* 读函数
 *   - 不得导出 create / update / delete / bind / link 等写入函数
 */
import type {
  DashboardOverview,
  DashboardOverviewResponse,
  FeedbackSignalsResponse,
  RecentActivitiesResponse,
} from '../types'

/**
 * API 基础 URL。
 *
 * 与 module-registry/data/api-adapter.ts 一致：
 * - 开发期间通过 Vite proxy 转发 /api 到后端，前端走同源 (localhost:5173)，此值为空字符串
 * - 生产环境由 Caddy 反代统一接入，也可通过构建时 env 注入实际地址
 */
const API_BASE_URL = import.meta.env.VITE_API_BASE_URL ?? ''

/**
 * API 错误类型，承载后端返回的 error 字段与 HTTP 状态码。
 *
 * 后端错误响应统一结构：{ "error": "<message>" }
 * 适配层将其抛出，由页面层在区块上下文或整页上下文中展示。
 */
export class ApiError extends Error {
  status: number
  constructor(status: number, message: string) {
    super(message)
    this.name = 'ApiError'
    this.status = status
  }
}

/**
 * 通用 JSON 请求封装。
 *
 * - GET 请求不带 body，不设 Content-Type（避免触发不必要的 CORS preflight）
 * - POST/PUT 请求带 JSON body，设 Content-Type: application/json
 * - 非 2xx 响应解析后端 { error } 结构并抛出 ApiError
 * - 204 No Content 返回 undefined
 *
 * 与 module-registry/data/api-adapter.ts 的 request<T> 实现保持一致，
 * 不引入第二套 fetch 封装。
 */
async function request<T>(path: string, init?: RequestInit): Promise<T> {
  // 仅有 body 的请求（POST/PUT/PATCH）才设 Content-Type，避免 GET 触发 CORS preflight
  const headers: Record<string, string> = { ...(init?.headers as Record<string, string> ?? {}) }
  if (init?.body && !headers['Content-Type']) {
    headers['Content-Type'] = 'application/json'
  }

  const res = await fetch(`${API_BASE_URL}${path}`, {
    ...init,
    headers,
  })

  // 204 No Content
  if (res.status === 204) {
    return undefined as T
  }

  // 非 2xx：解析错误体
  if (!res.ok) {
    let message = `HTTP ${res.status}`
    try {
      const body = await res.json()
      if (body?.error) {
        message = body.error
      }
    } catch {
      // 响应体非 JSON，保留默认 message
    }
    throw new ApiError(res.status, message)
  }

  return res.json() as Promise<T>
}

/**
 * DashboardOverviewRead — 主聚合读取
 * phase05-10 §7.1 主聚合读取，失败触发整页失败
 *
 * GET /api/dashboard/overview
 *
 * 返回 DashboardOverview（解包 envelope.overview）。
 * 后端保证 overview 字段不为 nil（空态返回零计数结构）。
 */
export async function fetchDashboardOverview(): Promise<DashboardOverview> {
  const response = await request<DashboardOverviewResponse>(`/api/dashboard/overview`)
  // 后端保证 overview 不为 nil；若异常情况为 undefined，回退到零计数结构以保证前端不崩
  return response.overview ?? {
    module_count: 0,
    product_count: 0,
    repository_count: 0,
    decision_count: 0,
    product_with_repository_count: 0,
    product_with_module_count: 0,
  }
}

/**
 * FeedbackSignalRead — 附属聚合读取
 * phase05-10 §7.1 附属聚合读取，失败只触发局部失败
 *
 * GET /api/dashboard/feedback-signals
 *
 * 返回 FeedbackSignalsResponse（current_focus_signals + asset_feedback_summary）。
 * 后端保证空态返回完整结构与零计数、代表项为空列表。
 */
export async function fetchFeedbackSignals(): Promise<FeedbackSignalsResponse> {
  return request<FeedbackSignalsResponse>(`/api/dashboard/feedback-signals`)
}

/**
 * RecentActivityRead — 附属聚合读取
 * phase05-10 §7.1 附属聚合读取，失败只触发局部失败
 *
 * GET /api/dashboard/recent-activities
 *
 * 返回 RecentActivitiesResponse（activities 列表，后端已按 activity_at DESC 排序）。
 * 后端保证空态返回空列表（非 nil）。
 */
export async function fetchRecentActivities(): Promise<RecentActivitiesResponse> {
  return request<RecentActivitiesResponse>(`/api/dashboard/recent-activities`)
}
