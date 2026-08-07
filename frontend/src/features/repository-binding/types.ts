/**
 * Repository Binding 类型定义
 * 对齐 phase04-10 正式规格正文与 phase04-12 后端 types.go
 *
 * 字段语义全部从 proto/psco/repository_binding/v1/repository_binding.proto 单向派生，
 * 不在前端新增后端未定义的业务字段或第二套状态枚举。
 */

/** Repository 状态 — phase04-04 冻结为 active / archived */
export type RepositoryStatus = 'active' | 'archived'

/** Repository 核心对象 — 对齐 proto Repository */
export interface Repository {
  id: string
  name: string
  url: string
  provider: string
  status: RepositoryStatus
  created_at: string
}

/** 列表读取模型 — 对齐 proto RepositoryListItem */
export interface RepositoryListItem {
  id: string
  name: string
  url: string
  provider: string
  status: RepositoryStatus
  created_at: string
  product_bind_count: number
  module_bind_count: number
}

/** 已绑定 Product 摘要 — 对齐 proto BoundProductSummary */
export interface BoundProductSummary {
  product_id: string
  product_name: string
  // product_status 描述的是 Product 实体的状态，使用字面量类型避免跨模块类型别名语义混淆
  product_status: 'active' | 'archived'
}

/** 已映射 Module 摘要 — 对齐 proto MappedModuleSummary */
export interface MappedModuleSummary {
  module_id: string
  module_name: string
  module_status: 'active' | 'archived'
}

/** 详情读取模型 — 对齐 proto RepositoryDetail */
export interface RepositoryDetail {
  repository: Repository
  bound_products: BoundProductSummary[]
  mapped_modules: MappedModuleSummary[]
}

/** Product 候选读取返回项 — 对齐 proto RepositoryProductCandidate */
export interface RepositoryProductCandidate {
  product_id: string
  product_name: string
  // product_status 描述的是 Product 实体的状态，使用字面量类型避免跨模块类型别名语义混淆
  product_status: 'active' | 'archived'
}

/** Module 候选读取返回项 — 对齐 proto RepositoryModuleCandidate */
export interface RepositoryModuleCandidate {
  module_id: string
  module_name: string
  module_status: 'active' | 'archived'
}

/** 列表查询参数（路由搜索参数） — phase04-06 冻结 */
export interface RepositoryListSearch {
  queryText?: string
  statusFilter?: RepositoryStatus | 'all'
}

/** CreateRepository 写入参数 — 对齐 proto CreateRepositoryRequest */
export interface CreateRepositoryInput {
  name: string
  url: string
  provider: string
  status: RepositoryStatus
}

/** CreateRepository 响应 — 对齐 proto CreateRepositoryResponse */
export interface CreateRepositoryResponse {
  repository_id: string
}

/** BindRepositoryToProduct 写入参数 — repositoryId 由 URL 路径参数承接 */
export interface BindRepositoryToProductInput {
  repositoryId: string
  productId: string
}

/** MapModuleToRepository 写入参数 — repositoryId 由 URL 路径参数承接 */
export interface MapModuleToRepositoryInput {
  repositoryId: string
  moduleId: string
}

/**
 * 详情页来源上下文 — phase04-06 冻结
 *
 * 只允许以下四种之一：
 * - fromList: 来自 Repository Binding / List，承接 queryText / statusFilter
 * - fromProductDetail: 来自 Product Detail，承接 productId / productName
 * - fromModuleDetail: 来自 Module Detail，承接 moduleId / moduleName
 * - direct-entry: 无来源参数
 */
export interface RepositoryDetailSourceContext {
  fromList?: boolean
  queryText?: string
  statusFilter?: RepositoryStatus | 'all'
  fromProductDetail?: boolean
  productId?: string
  productName?: string
  fromModuleDetail?: boolean
  moduleId?: string
  moduleName?: string
}
