/**
 * Product Registry 真实 API 适配层
 *
 * phase04-12 后端与数据主线已完成，本适配层直接通过真实 HTTP 调用对接 Go 后端。
 * 不提供并列 mock 主线（phase04-13 spec §"前端数据适配必须直接消费 phase04-12 真实 API"）。
 *
 * 字段映射说明：
 *   - 响应字段：后端已使用 snake_case，与前端 types.ts 一致，无需转换
 *   - 请求体字段：CreateProductInput 使用 camelCase，后端期望 snake_case，
 *     在本适配层完成转换（当前 name / description / status 无需转换）
 *   - productId：通过 URL 路径参数传递，不放在请求体
 *   - BindModuleToProduct 请求体：module_id（snake_case）
 *
 * 上游规格：phase04-12 spec §"路由装配必须沿用 chi 子路由主线"
 *           phase04-10 §RPC→HTTP 映射矩阵
 *
 * 路由矩阵（phase04-12 router.go）：
 *   GET    /api/products                                 ProductListRead
 *   GET    /api/products/{productId}                     ProductDetailRead
 *   GET    /api/products/{productId}/candidates/modules  ProductModuleCandidateRead
 *   POST   /api/products                                 ProductCreateWrite
 *   POST   /api/products/{productId}/bindings/modules    ProductModuleBindingWrite
 */
import type {
  ProductListItem,
  ProductDetail,
  ProductListSearch,
  CreateProductInput,
  CreateProductResponse,
  ProductModuleCandidate,
  BindModuleToProductInput,
} from '../types'

/**
 * API 基础 URL。
 *
 * 开发期间通过 Vite proxy 转发 /api 到后端，前端走同源 (localhost:5173)，
 * 此值为空字符串。生产环境由 Caddy 反代统一接入。
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
 * - 204 No Content（绑定成功无返回体）直接返回 undefined
 * - 非 2xx 响应解析后端 { error } 结构并抛出 ApiError
 */
async function request<T>(path: string, init?: RequestInit): Promise<T> {
  const headers: Record<string, string> = { ...(init?.headers as Record<string, string> ?? {}) }
  if (init?.body && !headers['Content-Type']) {
    headers['Content-Type'] = 'application/json'
  }

  const res = await fetch(`${API_BASE_URL}${path}`, {
    ...init,
    headers,
  })

  // 204 No Content（绑定成功无返回体）
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
 * ProductListRead — 列表读取
 * phase04-04 列表读取至少承接 name / description / status / created_at / module_bind_count / repository_bind_count
 *
 * GET /api/products?queryText=...&statusFilter=...
 */
export async function fetchProductList(search: ProductListSearch): Promise<ProductListItem[]> {
  const params = new URLSearchParams()
  if (search.queryText) {
    params.set('queryText', search.queryText)
  }
  if (search.statusFilter && search.statusFilter !== 'all') {
    params.set('statusFilter', search.statusFilter)
  }
  const qs = params.toString()
  return request<ProductListItem[]>(`/api/products${qs ? `?${qs}` : ''}`)
}

/**
 * ProductDetailRead — 详情读取（统一读模型宿主）
 * phase04-04 详情读取承接核心字段、已绑定模块列表与已绑定仓库列表
 *
 * GET /api/products/{productId}
 */
export async function fetchProductDetail(productId: string): Promise<ProductDetail> {
  return request<ProductDetail>(`/api/products/${encodeURIComponent(productId)}`)
}

/**
 * ProductModuleCandidateRead — 候选 Module 读取
 * phase04-04 候选读取为 Product Detail 的附属读取，独立于详情读取
 * 无可关联候选时返回空列表，不返回错误
 *
 * GET /api/products/{productId}/candidates/modules
 */
export async function fetchProductModuleCandidates(productId: string): Promise<ProductModuleCandidate[]> {
  return request<ProductModuleCandidate[]>(`/api/products/${encodeURIComponent(productId)}/candidates/modules`)
}

/**
 * ProductCreateWrite — 创建产品
 * phase04-04 创建写入承接 CreateProduct（最小字段 name / description / status）
 * 返回新建产品标识以支持前端回流到 ProductDetailPage
 *
 * POST /api/products
 */
export async function createProduct(input: CreateProductInput): Promise<CreateProductResponse> {
  const payload: Record<string, string> = {
    name: input.name,
  }
  if (input.description !== undefined) {
    payload.description = input.description
  }
  if (input.status !== undefined) {
    payload.status = input.status
  }

  return request<CreateProductResponse>(`/api/products`, {
    method: 'POST',
    body: JSON.stringify(payload),
  })
}

/**
 * ProductModuleBindingWrite — 绑定 Module 到 Product
 * phase04-04 BindModuleToProduct 归属 Product Registry 后端模块
 * productId 由 URL 路径参数承接
 *
 * POST /api/products/{productId}/bindings/modules
 *
 * 字段映射：moduleId → module_id
 */
export async function bindModuleToProduct(input: BindModuleToProductInput): Promise<void> {
  await request<void>(`/api/products/${encodeURIComponent(input.productId)}/bindings/modules`, {
    method: 'POST',
    body: JSON.stringify({
      module_id: input.moduleId,
    }),
  })
}
