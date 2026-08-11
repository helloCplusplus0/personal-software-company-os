/**
 * review-action-types — Review 动作类型定义。
 *
 * phase08-06 §"Review action owner 的最小接口"：
 *   ReviewActionSuccess 必须是稳定的 review-facing 成功 envelope，
 *   而不是直接把底层 mutation response 裸传给页面。
 *
 * 成功 envelope 至少包含：
 *   - resultKind
 *   - navigateTo
 *   - params
 *   - search
 *   - 可选的 successMessage
 */
import type { DashboardSection } from '@/features/dashboard/types'

/**
 * ReviewResultKind — review 结果类型，对齐 .proto ReviewResultKind 枚举。
 */
export type ReviewResultKind =
  | 'decision_handoff'
  | 'entity_handoff'
  | 'next_step_result'

/**
 * ReviewActionType — review 动作类型。
 *
 * - create_decision：进入 Decision Create，形成新的正式经营判断
 * - decision_handoff：进入既有 Decision canonical 页面
 * - entity_handoff：进入既有 Product / Module / Repository canonical 页面
 * - next_step_result：提交 review 过程记录
 */
export type ReviewActionType =
  | 'create_decision'
  | 'go_to_decision'
  | 'go_to_product'
  | 'go_to_module'
  | 'go_to_repository'
  | 'submit_next_step'

/**
 * ReviewActionInput — review 动作的统一输入。
 */
export interface ReviewActionInput {
  actionType: ReviewActionType
  /** 可选：目标实体 ID（decision / product / module / repository） */
  targetId?: string
  /** 可选：摘要文本（next_step_result 时必填） */
  summaryText?: string
  /** 当前 Dashboard 来源区块 */
  dashboardSection: DashboardSection
  /** 当前 review 会话类型，用于 submit_next_step 时区分 daily/weekly */
  reviewKind?: 'daily' | 'weekly'
}

/**
 * ReviewActionSuccess — review 动作的成功 envelope。
 *
 * 页面只负责消费该 success envelope 执行 navigate() 与展示 toast。
 */
export interface ReviewActionSuccess {
  resultKind: ReviewResultKind
  navigateTo: string
  params?: Record<string, string>
  search: {
    fromDashboard?: boolean
    dashboardSection?: DashboardSection
    dashboardReturnTo?: string
  }
  successMessage?: string
}
