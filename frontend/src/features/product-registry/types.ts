/**
 * Product Registry 类型定义
 * 对齐 phase04-10 正式规格正文与 phase04-12 后端 types.go
 *
 * 字段语义全部从 proto/psco/product_registry/v1/product_registry.proto 单向派生，
 * 不在前端新增后端未定义的业务字段或第二套状态枚举。
 */

/** Product 状态 — phase04-04 冻结为 active / archived */
export type ProductStatus = 'active' | 'archived'

/** Product 核心对象 — 对齐 proto Product */
export interface Product {
  id: string
  name: string
  description: string
  status: ProductStatus
  created_at: string
}

/** 列表读取模型 — 对齐 proto ProductListItem */
export interface ProductListItem {
  id: string
  name: string
  description: string
  status: ProductStatus
  created_at: string
  module_bind_count: number
  repository_bind_count: number
}

/** 已绑定 Module 摘要 — 对齐 proto BoundModuleSummary */
export interface BoundModuleSummary {
  module_id: string
  module_name: string
  module_status: 'active' | 'archived'
}

/** 已绑定 Repository 摘要 — 对齐 proto BoundRepositorySummary */
export interface BoundRepositorySummary {
  repository_id: string
  repository_name: string
  provider: string
  // repository_status 描述的是 Repository 实体的状态，使用字面量类型避免跨模块类型别名语义混淆
  repository_status: 'active' | 'archived'
}

/** 详情读取模型 — 对齐 proto ProductDetail */
export interface ProductDetail {
  product: Product
  bound_modules: BoundModuleSummary[]
  bound_repositories: BoundRepositorySummary[]
}

/** Module 候选读取返回项 — 对齐 proto ProductModuleCandidate */
export interface ProductModuleCandidate {
  module_id: string
  module_name: string
  module_status: 'active' | 'archived'
}

/** 列表查询参数（路由搜索参数） — phase04-06 冻结 */
export interface ProductListSearch {
  queryText?: string
  statusFilter?: ProductStatus | 'all'
}

/** CreateProduct 写入参数 — 对齐 proto CreateProductRequest */
export interface CreateProductInput {
  name: string
  description?: string
  status?: ProductStatus
}

/** CreateProduct 响应 — 对齐 proto CreateProductResponse */
export interface CreateProductResponse {
  product_id: string
}

/** BindModuleToProduct 写入参数 — productId 由 URL 路径参数承接 */
export interface BindModuleToProductInput {
  productId: string
  moduleId: string
}

/**
 * 详情页来源上下文 — phase04-06 冻结
 *
 * 只允许以下三种之一：
 * - fromList: 来自 Product List，承接 queryText / statusFilter
 * - fromModuleDetail: 来自 Module Detail，承接 moduleId / moduleName
 * - direct-entry: 无来源参数
 */
export interface ProductDetailSourceContext {
  fromList?: boolean
  queryText?: string
  statusFilter?: ProductStatus | 'all'
  fromModuleDetail?: boolean
  moduleId?: string
  moduleName?: string
}
