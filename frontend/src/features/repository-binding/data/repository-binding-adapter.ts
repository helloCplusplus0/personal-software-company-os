/**
 * Repository Binding 数据适配层 — 直接导出真实 API 实现
 *
 * phase04-13 spec §"前端数据适配必须直接消费 phase04-12 真实 API"：
 * 不提供并列 mock-adapter.ts，本文件直接 re-export api-adapter.ts 的真实 API 实现。
 *
 * 上游规格：phase04-13 spec §"前端数据适配必须直接消费 phase04-12 真实 API"
 */
export {
  fetchRepositoryList,
  fetchRepositoryDetail,
  fetchRepositoryProductCandidates,
  fetchRepositoryModuleCandidates,
  createRepository,
  bindRepositoryToProduct,
  mapModuleToRepository,
  ApiError,
} from './api-adapter'
