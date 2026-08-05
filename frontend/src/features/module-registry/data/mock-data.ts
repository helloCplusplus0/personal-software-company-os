/**
 * Mock 数据 — 仅服务 phase02-10 前端主线运行与演示
 * phase02-11 后端主线完成后，此数据将被真实 API 替换
 * 数据结构严格对齐 phase02-09 正式规格正文
 */
import type {
  Module,
  Release,
  ProductBinding,
  RepositoryMapping,
  DecisionLink,
  ProductCandidate,
  RepositoryCandidate,
} from '../types'

/**
 * Mock 内部绑定关系类型
 *
 * ProductBinding / RepositoryMapping 是详情读取模型（不含 module_id），
 * 但 mock 层需要 module_id 来建立模块与绑定之间的关系，
 * 供列表读取的 product_bind_count / repository_bind_count 与详情读取的按模块过滤使用。
 * 此类型仅在 mock 层内部使用，不泄漏到前端类型契约。
 */
interface MockProductBinding extends ProductBinding {
  module_id: string
}

interface MockRepositoryMapping extends RepositoryMapping {
  module_id: string
}

/**
 * Mock 内部 Decision 关联类型
 *
 * DecisionLink 是详情读取模型（不含 module_id），
 * 但 mock 层需要 module_id 来建立模块与决策之间的关系，
 * 供详情读取的按模块过滤使用。
 */
interface MockDecisionLink extends DecisionLink {
  module_id: string
}

// --- 直接承接数据 ---

export const mockModules: Module[] = [
  { id: 'mod-1', name: 'auth-service', description: '认证服务模块，提供 OAuth2 与 JWT 支持', status: 'active', created_at: '2026-07-01T00:00:00Z' },
  { id: 'mod-2', name: 'payment-gateway', description: '支付网关模块，对接多渠道支付', status: 'active', created_at: '2026-07-15T00:00:00Z' },
  { id: 'mod-3', name: 'legacy-importer', description: '旧版数据导入器', status: 'archived', created_at: '2026-06-01T00:00:00Z' },
]

export const mockReleases: Release[] = [
  { id: 'rel-1', module_id: 'mod-1', version: '1.0.0', status: 'active', released_at: '2026-07-01T00:00:00Z' },
  { id: 'rel-2', module_id: 'mod-1', version: '1.1.0', status: 'active', released_at: '2026-07-10T00:00:00Z' },
  { id: 'rel-3', module_id: 'mod-2', version: '0.9.0', status: 'active', released_at: '2026-07-15T00:00:00Z' },
]

// --- 绑定关系（含 module_id 关联） ---

export const mockProductBindings: MockProductBinding[] = [
  { module_id: 'mod-1', product_id: 'prod-1', product_name: 'Product A' },
]

export const mockRepositoryMappings: MockRepositoryMapping[] = [
  { module_id: 'mod-1', repository_id: 'repo-1', repository_name: 'main-repo' },
]

// --- Decision 只读入口（附属读取，不设独立读接口组，含 module_id 关联） ---

export const mockDecisionLinks: MockDecisionLink[] = [
  { module_id: 'mod-1', decision_id: 'dec-1', decision_title: '关于 auth-service 技术选型的决策' },
]

// --- 候选读取数据（phase02 临时承接，只读） ---

export const mockProductCandidates: ProductCandidate[] = [
  { id: 'prod-1', name: 'Product A' },
  { id: 'prod-2', name: 'Product B' },
  { id: 'prod-3', name: 'Product C' },
]

export const mockRepositoryCandidates: RepositoryCandidate[] = [
  { id: 'repo-1', name: 'main-repo' },
  { id: 'repo-2', name: 'mirror-repo' },
]
