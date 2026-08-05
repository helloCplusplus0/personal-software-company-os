/**
 * Module Registry 类型定义
 * 对齐 phase02-09 正式规格正文 module_registry_spec_v0.1.md
 */

/** Module 状态 — §5.5 冻结为 active / archived */
export type ModuleStatus = 'active' | 'archived'

/** Release 状态 */
export type ReleaseStatus = 'active' | 'archived'

/** Module 核心对象 — §5.4 */
export interface Module {
  id: string
  name: string
  description: string
  status: ModuleStatus
  created_at: string
}

/** Release 核心对象 — §5.4 */
export interface Release {
  id: string
  module_id: string
  version: string
  status: ReleaseStatus
  released_at: string
}

/** 列表读取模型 — §5.7 */
export interface ModuleListItem {
  id: string
  name: string
  description: string
  status: ModuleStatus
  latest_release: string | null
  product_bind_count: number
  repository_bind_count: number
}

/** 产品绑定关系 */
export interface ProductBinding {
  product_id: string
  product_name: string
}

/** 仓库映射关系 */
export interface RepositoryMapping {
  repository_id: string
  repository_name: string
}

/** Decision 入口（只读展示，不设独立读接口组） — §6.3 */
export interface DecisionLink {
  decision_id: string
  decision_title: string
}

/** 详情读取模型 — §5.7 */
export interface ModuleDetail {
  module: Module
  releases: Release[]
  product_bindings: ProductBinding[]
  repository_mappings: RepositoryMapping[]
  decision_links: DecisionLink[]
}

/** 候选 Product — §6.2 ProductBindingCandidateRead */
export interface ProductCandidate {
  id: string
  name: string
}

/** 候选 Repository — §6.2 RepositoryBindingCandidateRead */
export interface RepositoryCandidate {
  id: string
  name: string
}

/** 列表查询参数（路由搜索参数） — §8.4 */
export interface ModuleListSearch {
  queryText?: string
  statusFilter?: ModuleStatus | 'all'
}

/** CreateModule 写入参数 — §5.7 */
export interface CreateModuleInput {
  name: string
  description: string
  status: ModuleStatus
}

/** CreateRelease 写入参数 — §5.7 */
export interface CreateReleaseInput {
  moduleId: string
  version: string
  status: ReleaseStatus
  releasedAt: string
}

/** BindModuleToProduct 写入参数 */
export interface BindModuleToProductInput {
  moduleId: string
  productId: string
}

/** MapModuleToRepository 写入参数 */
export interface MapModuleToRepositoryInput {
  moduleId: string
  repositoryId: string
}
