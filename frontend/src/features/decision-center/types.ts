/**
 * Decision Center 类型定义
 *
 * 对齐 phase03-10 decision_center_spec_v0.1.md 与 phase03-12 后端实现。
 * 字段语义从 phase03-11 .proto 合同源派生，不形成第二套合同源。
 *
 * 后端返回 JSON 使用 snake_case，前端类型直接承接，无需转换。
 */

/** Decision 状态 — §5.6 冻结为 proposed / active / superseded / archived */
export type DecisionStatus = 'proposed' | 'active' | 'superseded' | 'archived'

/** Decision 关联目标类型 — 当前阶段唯一允许值为 module */
export type DecisionLinkTargetType = 'module'

/**
 * Decision 核心对象 — §5.4 / §5.5 最小结构化模板字段。
 * alternatives 建模为 string[]，按输入顺序保留。
 */
export interface Decision {
  id: string
  title: string
  context: string
  problem: string
  alternatives: string[]
  choice: string
  reason: string
  impact: string
  status: DecisionStatus
  created_at: string
}

/** 列表读取模型 — §5.9 */
export interface DecisionListItem {
  id: string
  title: string
  status: DecisionStatus
  created_at: string
  /** 已建立 Decision -> Module 有效关联数，无关联返回 0 */
  link_count: number
  /** 按 module_name 升序取前 3 + "+N"，无关联返回空字符串 */
  linked_module_summary: string
}

/** 已关联 Module — §5.8 DecisionDetailRead.linked_modules */
export interface LinkedModule {
  module_id: string
  module_name: string
}

/** 来源上下文 — §5.11 入口上下文与正式关联结果边界 */
export interface SourceContext {
  /** 来源 Module 标识，无来源时为空字符串 */
  source_module_id: string
  /** 来源 Module 名称，无来源时为空字符串 */
  source_module_name: string
}

/** 详情读取模型 — §5.8 */
export interface DecisionDetail {
  decision: Decision
  /** 已关联 Module 列表 */
  linked_modules: LinkedModule[]
  /** 入口上下文来源（§5.11 持续到正式关联完成，当前阶段不提供主动放弃出口） */
  source_context: SourceContext
}

/** 候选 Module — §5.10 DecisionModuleCandidateRead */
export interface DecisionModuleCandidate {
  module_id: string
  module_name: string
  /** 复用 ModuleRegistry.ModuleStatus，跨包不重定义 */
  status: 'active' | 'archived'
}

/** 列表查询参数（路由搜索参数） — §9.1 */
export interface DecisionListSearch {
  queryText?: string
  statusFilter?: DecisionStatus | 'all'
}

/** CreateDecision 写入参数 — §5.5 / §5.11 */
export interface CreateDecisionInput {
  title: string
  context: string
  problem: string
  alternatives: string[]
  choice: string
  reason: string
  impact: string
  status: DecisionStatus
  /** 入口上下文来源 Module 标识（可选，§5.11），空字符串表示无来源 */
  source_module_id?: string
}

/** CreateDecision 响应 — §6.4 只返回 decision_id */
export interface CreateDecisionResponse {
  decision_id: string
}

/** LinkDecisionToTarget 写入参数 — decision_id 由 URL 路径参数承接 */
export interface LinkDecisionToTargetInput {
  decisionId: string
  target_type: DecisionLinkTargetType
  module_id: string
}

/** 列表搜索上下文快照（Zustand + sessionStorage 缓存） */
export interface DecisionListSearchContext {
  queryText?: string
  statusFilter: DecisionStatus | 'all'
}
