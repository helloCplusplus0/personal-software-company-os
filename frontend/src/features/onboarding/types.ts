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
