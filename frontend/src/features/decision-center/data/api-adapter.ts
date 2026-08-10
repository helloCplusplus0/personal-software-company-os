/**
 * Decision Center 真实 API 适配层
 *
 * phase03-12 后端与数据主线已完成，本适配层直接通过真实 HTTP 调用对接 Go 后端。
 * 不提供并列 mock 主线（phase03-13 spec §"前端数据适配必须直接消费 phase03-12 真实 API"）。
 *
 * 字段映射说明：
 *   - 响应字段：后端已使用 snake_case，与前端 types.ts 一致，无需转换
 *   - 请求体字段：CreateDecisionInput 使用 camelCase（sourceModuleId），
 *     后端期望 snake_case（source_module_id），在本适配层完成转换
 *   - decisionId：通过 URL 路径参数传递，不放在请求体
 *
 * 上游规格：phase03-12 spec §"应用装配必须将 Decision Center 挂载到 /api"
 *           phase03-10 §7.7 RPC → HTTP 映射矩阵
 */
import type {
  DecisionListItem,
  DecisionDetail,
  DecisionModuleCandidate,
  CreateDecisionInput,
  CreateDecisionResponse,
  LinkDecisionToTargetInput,
  DecisionListSearch,
} from '../types'

/**
 * API 基础 URL。
 *
 * 开发期间通过 Vite proxy 转发 /api 到后端，前端走同源 (localhost:5173)，
 * 此值为空字符串。生产环境由 Caddy 反代统一接入，也可通过构建时 env 注入实际地址。
 */
const API_BASE_URL = import.meta.env.VITE_API_BASE_URL ?? ''

/**
 * API 错误类型，承载后端返回的 error 字段与 HTTP 状态码。
 *
 * 后端错误响应统一结构：{ "error": "<message>" }
 * 适配层将其抛出，由页面层在表单上下文或面板上下文中展示。
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
 * - POST 请求带 JSON body，设 Content-Type: application/json
 * - 204 No Content 直接返回 undefined
 * - 非 2xx 响应解析后端 { error } 结构并抛出 ApiError
 *
 * phase03-14 修复：GET 请求不再统一设 Content-Type。
 * 之前对所有请求都加 Content-Type: application/json，使 GET 变成非 simple request，
 * 每次都触发 CORS preflight。结合后端未设 Access-Control-Max-Age（Chrome 默认仅 5 秒），
 * 导致 POST 在 preflight 缓存过期后被中止 (net::ERR_ABORTED) 再重新 preflight。
 * 现在仅在有 body 时设 Content-Type，GET 请求恢复为 simple request，无需 preflight。
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

  // 204 No Content（关联成功无返回体）
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
 * DecisionListRead — 列表读取
 * §5.9 列表读取至少承接 title / status / created_at / link_count / linked_module_summary
 *
 * GET /api/decisions?queryText=...&statusFilter=...
 */
export async function fetchDecisionList(search: DecisionListSearch): Promise<DecisionListItem[]> {
  const params = new URLSearchParams()
  if (search.queryText) {
    params.set('queryText', search.queryText)
  }
  if (search.statusFilter && search.statusFilter !== 'all') {
    params.set('statusFilter', search.statusFilter)
  }
  const qs = params.toString()
  return request<DecisionListItem[]>(`/api/decisions${qs ? `?${qs}` : ''}`)
}

/**
 * DecisionDetailRead — 详情读取（统一读模型宿主）
 * §5.8 详情读取承接核心对象字段、结构化模板字段、已关联目标列表与最小来源上下文
 *
 * GET /api/decisions/{decisionId}
 */
export async function fetchDecisionDetail(decisionId: string): Promise<DecisionDetail> {
  return request<DecisionDetail>(`/api/decisions/${encodeURIComponent(decisionId)}`)
}

/**
 * DecisionModuleCandidateRead — 候选 Module 读取
 * §5.10 候选读取为 Decision Detail 的附属读取
 *
 * GET /api/decisions/{decisionId}/candidates/modules
 */
export async function fetchDecisionModuleCandidates(decisionId: string): Promise<DecisionModuleCandidate[]> {
  return request<DecisionModuleCandidate[]>(`/api/decisions/${encodeURIComponent(decisionId)}/candidates/modules`)
}

/**
 * DecisionWrite — 创建决策
 * §5.5 创建写入承接 RecordDecision（最小结构化模板字段）
 * §5.11 source_module_id 为可选来源上下文
 * 返回新建决策标识以支持前端回流到 DecisionDetailPage
 *
 * POST /api/decisions
 *
 * 字段映射：sourceModuleId → source_module_id
 */
export async function createDecision(input: CreateDecisionInput): Promise<CreateDecisionResponse> {
  const payload: Record<string, string | string[]> = {
    title: input.title,
    choice: input.choice,
    reason: input.reason,
    source_module_id: input.source_module_id ?? '',
  }
  if (input.context !== undefined) {
    payload.context = input.context
  }
  if (input.problem !== undefined) {
    payload.problem = input.problem
  }
  if (input.alternatives !== undefined) {
    payload.alternatives = input.alternatives
  }
  if (input.impact !== undefined) {
    payload.impact = input.impact
  }
  if (input.status !== undefined) {
    payload.status = input.status
  }

  return request<CreateDecisionResponse>(`/api/decisions`, {
    method: 'POST',
    body: JSON.stringify(payload),
  })
}

/**
 * DecisionLinkWrite — 关联目标
 * §5.10 关联写入承接 LinkDecisionToTarget
 * decision_id 由 URL 路径参数承接，不放在请求体
 *
 * POST /api/decisions/{decisionId}/links
 *
 * @returns 204 No Content（无返回体）
 */
export async function linkDecisionToTarget(input: LinkDecisionToTargetInput): Promise<void> {
  await request<void>(`/api/decisions/${encodeURIComponent(input.decisionId)}/links`, {
    method: 'POST',
    body: JSON.stringify({
      target_type: input.target_type,
      module_id: input.module_id,
    }),
  })
}
