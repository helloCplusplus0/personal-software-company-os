/**
 * Onboarding 前端类型定义
 *
 * 对齐 phase06-14 后端 onboarding 模块响应结构。
 * 字段语义从 .proto -> HTTP DTO 单向派生（phase06-05 / phase06-13 合同约束）。
 *
 * 后端响应示例（GET /api/onboarding/state）：
 *   {
 *     "first_run_state": {
 *       "status": "completed",
 *       "is_first_entry": false,
 *       "current_step": "complete",
 *       "completion_progress": 100
 *     }
 *   }
 */

/** 首轮运行状态（对齐 proto FirstRunStatus） */
export type FirstRunStatus = 'not_started' | 'in_progress' | 'completed'

/** Onboarding 步骤（对齐 proto OnboardingStep） */
export type OnboardingStep =
  | 'welcome'
  | 'product'
  | 'repository'
  | 'module'
  | 'decision'
  | 'complete'

/** first_run_state 读模型（对齐 proto FirstRunState） */
export interface FirstRunState {
  status: FirstRunStatus
  is_first_entry: boolean
  current_step: OnboardingStep
  completion_progress: number
}

/** GetFirstRunState 响应（对齐 proto GetFirstRunStateResponse） */
export interface OnboardingReadResult {
  first_run_state: FirstRunState
}

// ============================================================================
// 建链状态类型（phase10-08 新增）
// ============================================================================

/** 建链恢复状态（对齐 proto resume_status） */
export type ChainStateResumeStatus = 'cold_start' | 'resuming' | 'completed'

/** 下一步类型（对齐 proto next_step_kind） */
export type ChainStateNextStepKind = 'create' | 'bind' | 'handoff' | 'complete'

/** 建链状态读模型（对齐 proto GetOnboardingChainStateResponse） */
export interface OnboardingChainState {
  current_product_id: string
  current_step: OnboardingStep
  resume_status: ChainStateResumeStatus
  next_step_kind: ChainStateNextStepKind
  canonical_handoff_target?: string
  return_hint?: string
}

/** 组合后的 Onboarding 读模型（phase10-08 组合 GetFirstRunState + GetOnboardingChainState） */
export interface OnboardingFullReadResult {
  first_run_state: FirstRunState
  chain_state: OnboardingChainState
}
