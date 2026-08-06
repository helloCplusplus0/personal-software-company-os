/**
 * Decision Center 数据适配层 — 统一导出入口
 *
 * phase03-12 后端与数据主线已完成，本适配层直接导出真实 API 实现。
 * 不提供并列 mock 主线，不提供运行时切换开关（phase03-13 spec §"前端数据适配必须直接消费 phase03-12 真实 API"）。
 *
 * 与 module-registry/data/module-registry-adapter.ts 的区别：
 *   - module-registry 保留了 mock / real 双轨切换（phase02-10 遗留）
 *   - decision-center 从一开始就只导出真实 API，不引入第二套数据主线
 *
 * 上游规格：phase03-13 spec §"前端数据适配必须直接消费 phase03-12 真实 API"
 */
export {
  fetchDecisionList,
  fetchDecisionDetail,
  fetchDecisionModuleCandidates,
  createDecision,
  linkDecisionToTarget,
  ApiError,
} from './api-adapter'
