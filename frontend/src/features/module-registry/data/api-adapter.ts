/**
 * Module Registry 真实 API 适配层
 *
 * phase02-11 后端与数据主线完成后，通过真实 HTTP 调用对接 Go 后端。
 * 函数签名与 mock-adapter.ts 完全一致，页面与组件代码无需修改。
 *
 * 切换机制：由 module-registry-adapter.ts 根据 VITE_USE_REAL_API 环境变量决定
 * 导出 mock 实现还是真实 API 实现。
 *
 * 字段映射说明：
 *   - 响应字段：后端已使用 snake_case，与前端 types.ts 一致，无需转换
 *   - 请求体字段：前端 input 类型使用 camelCase（releasedAt / productId / repositoryId），
 *     后端期望 snake_case，在本适配层完成转换
 *   - moduleId：通过 URL 路径参数传递，不放在请求体
 *
 * 上游规格：phase02-11 spec §"前端临时适配层必须能切换到真实后端"
 */
import type {
  Module,
  ModuleListItem,
  ModuleDetail,
  ModuleListSearch,
  CreateModuleInput,
  CreateReleaseInput,
  Release,
  BindModuleToProductInput,
  MapModuleToRepositoryInput,
  ProductCandidate,
  RepositoryCandidate,
} from '../types'

/**
 * API 基础 URL。
 *
 * 通过 Vite 环境变量 VITE_API_BASE_URL 配置，缺省指向本地后端 http://localhost:8080。
 * 生产环境由 Caddy 反代统一接入，可通过构建时 env 注入实际地址。
 */
const API_BASE_URL = import.meta.env.VITE_API_BASE_URL ?? 'http://localhost:8080'

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
 * - GET 请求不带 body
 * - POST/PUT 请求带 JSON body
 * - 非 2xx 响应解析后端 { error } 结构并抛出 ApiError
 */
async function request<T>(path: string, init?: RequestInit): Promise<T> {
  const res = await fetch(`${API_BASE_URL}${path}`, {
    headers: {
      'Content-Type': 'application/json',
      ...(init?.headers ?? {}),
    },
    ...init,
  })

  // 204 No Content（绑定/映射成功无返回体）
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
 * ModuleListRead — 列表读取
 * §5.7 列表读取至少承接 name / description / status / latest_release / product_bind_count / repository_bind_count
 *
 * GET /api/modules?queryText=...&statusFilter=...
 */
export async function fetchModuleList(search: ModuleListSearch): Promise<ModuleListItem[]> {
  const params = new URLSearchParams()
  if (search.queryText) {
    params.set('queryText', search.queryText)
  }
  if (search.statusFilter && search.statusFilter !== 'all') {
    params.set('statusFilter', search.statusFilter)
  }
  const qs = params.toString()
  return request<ModuleListItem[]>(`/api/modules${qs ? `?${qs}` : ''}`)
}

/**
 * ModuleDetailRead — 详情读取（统一读模型宿主）
 * §5.7 详情读取至少承接核心字段、版本列表、产品绑定、仓库映射与 Decision 入口
 * §6.3 Decision 读取内嵌于此，不设独立读接口
 *
 * GET /api/modules/{moduleId}
 */
export async function fetchModuleDetail(moduleId: string): Promise<ModuleDetail> {
  return request<ModuleDetail>(`/api/modules/${encodeURIComponent(moduleId)}`)
}

/**
 * ModuleCreateWrite — 创建模块
 * §5.7 创建写入承接 CreateModule（最小字段 name / description / status）
 * 返回新建模块标识以支持前端回流到 ModuleDetailPage
 *
 * POST /api/modules
 */
export async function createModule(input: CreateModuleInput): Promise<Module> {
  return request<Module>(`/api/modules`, {
    method: 'POST',
    body: JSON.stringify({
      name: input.name,
      description: input.description,
      status: input.status,
    }),
  })
}

/**
 * ModuleReleaseWrite — 版本登记
 * §5.7 版本写入承接 CreateRelease（最小字段 version / status / released_at，module_id 由上下文隐式承接）
 *
 * POST /api/modules/{moduleId}/releases
 *
 * 字段映射：releasedAt → released_at
 */
export async function createRelease(input: CreateReleaseInput): Promise<Release> {
  return request<Release>(`/api/modules/${encodeURIComponent(input.moduleId)}/releases`, {
    method: 'POST',
    body: JSON.stringify({
      version: input.version,
      status: input.status,
      released_at: input.releasedAt,
    }),
  })
}

/**
 * ModuleBindingWrite — 绑定产品
 * §4.1 BindModuleToProduct 归属 Module Registry 后端模块
 *
 * POST /api/modules/{moduleId}/bindings/products
 *
 * 字段映射：productId → product_id
 */
export async function bindModuleToProduct(input: BindModuleToProductInput): Promise<void> {
  await request<void>(`/api/modules/${encodeURIComponent(input.moduleId)}/bindings/products`, {
    method: 'POST',
    body: JSON.stringify({
      product_id: input.productId,
    }),
  })
}

/**
 * ModuleBindingWrite — 映射仓库
 * §4.1 MapModuleToRepository 归属 Module Registry 后端模块
 *
 * POST /api/modules/{moduleId}/bindings/repositories
 *
 * 字段映射：repositoryId → repository_id
 */
export async function mapModuleToRepository(input: MapModuleToRepositoryInput): Promise<void> {
  await request<void>(`/api/modules/${encodeURIComponent(input.moduleId)}/bindings/repositories`, {
    method: 'POST',
    body: JSON.stringify({
      repository_id: input.repositoryId,
    }),
  })
}

/**
 * ProductBindingCandidateRead — Product 候选读取
 * §6.2 候选读取（phase02 由 Module Registry 临时承接）
 *
 * GET /api/candidates/products
 */
export async function fetchProductCandidates(): Promise<ProductCandidate[]> {
  return request<ProductCandidate[]>(`/api/candidates/products`)
}

/**
 * RepositoryBindingCandidateRead — Repository 候选读取
 * §6.2 候选读取（phase02 由 Module Registry 临时承接）
 *
 * GET /api/candidates/repositories
 */
export async function fetchRepositoryCandidates(): Promise<RepositoryCandidate[]> {
  return request<RepositoryCandidate[]>(`/api/candidates/repositories`)
}
