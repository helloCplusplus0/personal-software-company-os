/**
 * Module Registry 数据适配层 — 切换入口
 *
 * phase02-11 后端与数据主线完成后，前端可在 mock 适配层与真实 API 适配层之间切换。
 * 切换由 Vite 环境变量 VITE_USE_REAL_API 控制：
 *   - VITE_USE_REAL_API=true  → 使用真实后端 API（api-adapter.ts）
 *   - 未设置 / false          → 使用 mock 数据（mock-adapter.ts，phase02-10 默认行为）
 *
 * 切换不改变任何函数签名，页面与组件代码无需修改。
 * 上游规格：phase02-11 spec §"前端临时适配层必须能切换到真实后端"
 * 约束：不得借机引入第二套对象字段语义、第二套返回路径或第二套数据主线。
 *
 * 使用方式（本地联调）：
 *   1. 启动后端：./backend/bin/psco-server（需 .env 配置 DATABASE_URL）
 *   2. 在 frontend/.env 设置 VITE_USE_REAL_API=true
 *   3. 重启 Vite dev server：npm run dev
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

// 切换开关：由 Vite 环境变量决定
const USE_REAL_API = import.meta.env.VITE_USE_REAL_API === 'true'

// 按需导入对应实现
// 使用动态条件导入会引入代码分割复杂度，phase02 阶段采用静态导入 + 运行时分支
import * as mockAdapter from './mock-adapter'
import * as apiAdapter from './api-adapter'

/**
 * 适配层统一导出函数。
 *
 * 根据 USE_REAL_API 标志在运行时选择 mock 或真实 API 实现。
 * 函数签名与 mock-adapter / api-adapter 完全一致。
 */
export const fetchModuleList = (search: ModuleListSearch): Promise<ModuleListItem[]> =>
  USE_REAL_API ? apiAdapter.fetchModuleList(search) : mockAdapter.fetchModuleList(search)

export const fetchModuleDetail = (moduleId: string): Promise<ModuleDetail> =>
  USE_REAL_API ? apiAdapter.fetchModuleDetail(moduleId) : mockAdapter.fetchModuleDetail(moduleId)

export const createModule = (input: CreateModuleInput): Promise<Module> =>
  USE_REAL_API ? apiAdapter.createModule(input) : mockAdapter.createModule(input)

export const createRelease = (input: CreateReleaseInput): Promise<Release> =>
  USE_REAL_API ? apiAdapter.createRelease(input) : mockAdapter.createRelease(input)

export const bindModuleToProduct = (input: BindModuleToProductInput): Promise<void> =>
  USE_REAL_API ? apiAdapter.bindModuleToProduct(input) : mockAdapter.bindModuleToProduct(input)

export const mapModuleToRepository = (input: MapModuleToRepositoryInput): Promise<void> =>
  USE_REAL_API ? apiAdapter.mapModuleToRepository(input) : mockAdapter.mapModuleToRepository(input)

export const fetchProductCandidates = (): Promise<ProductCandidate[]> =>
  USE_REAL_API ? apiAdapter.fetchProductCandidates() : mockAdapter.fetchProductCandidates()

export const fetchRepositoryCandidates = (): Promise<RepositoryCandidate[]> =>
  USE_REAL_API ? apiAdapter.fetchRepositoryCandidates() : mockAdapter.fetchRepositoryCandidates()

// 导出 ApiError 以便页面层区分错误类型（仅在真实 API 模式下抛出）
export { ApiError } from './api-adapter'
