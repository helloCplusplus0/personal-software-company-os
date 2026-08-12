/**
 * use-derived-hint-handoff — 派生提示的单一 application handoff owner。
 *
 * phase09-10 spec §"提示动作 handoff 必须收敛到单一 application owner"：
 *   本 owner 统一承接提示点击后的 search 参数拼装、返回链恢复与 target owner 跳转，
 *   避免导航逻辑散落在多个页面和卡片组件里。
 *
 * 职责：
 *   - reuse_opportunity_hint -> Product Create 的 handoff 参数拼装
 *   - capability_gap_hint -> Module Registry / Module Detail 的 handoff 参数拼装
 *   - Weekly Review active candidate 返回恢复
 *   - Product Create 会话返回恢复
 *   - 非法提示参数回退与错误归一化
 *
 * 不承接：
 *   - 页面级提示展示逻辑（由各页面组件负责）
 *   - 提示数据读取（由 use-derived-insight-hints-read 负责）
 */
import { useCallback } from 'react'
import {
  DerivedInsightHintType,
  CTAKind,
} from '@/gen/proto/psco/template_reuse/v1/template_reuse_pb'
import type { DerivedInsightHint } from '@/gen/proto/psco/template_reuse/v1/template_reuse_pb'

// ============================================================================
// 类型
// ============================================================================

/** 提示消费面 */
export type HintConsumerSurface = 'weekly-review' | 'product-create'

/** 提示 handoff 的返回链上下文 */
export interface HintReturnContext {
  /** 发起提示消费的页面 */
  sourceSurface: HintConsumerSurface
  /** Weekly Review 的 active candidate ID（用于返回恢复） */
  activeCandidateId?: string
  /** Product Create 的模板来源参数（用于返回恢复） */
  fromTemplateReuse?: boolean
  templateCandidateId?: string
  templateSource?: string
  /** Dashboard 返回链元数据 */
  fromDashboard?: boolean
  dashboardSection?: string
  dashboardReturnTo?: string
}

/** 单个提示的 handoff 结果 */
export interface HintHandoffResult {
  /** 导航目标路径 */
  to: string
  /** 路由参数 */
  params?: Record<string, string>
  /** 搜索参数 */
  search: Record<string, unknown>
}

// ============================================================================
// Hook
// ============================================================================

export interface UseDerivedHintHandoffResult {
  /** 计算单个提示的导航目标 */
  computeHandoff: (hint: DerivedInsightHint) => HintHandoffResult | null
  /** 是否为合法提示（四元组完整） */
  isValidHint: (hint: DerivedInsightHint) => boolean
}

export function useDerivedHintHandoff(
  returnContext: HintReturnContext,
): UseDerivedHintHandoffResult {

  /**
   * 判断提示是否满足四元组（trigger / explanation / CTA / target owner）
   */
  const isValidHint = useCallback((hint: DerivedInsightHint): boolean => {
    // 必须有明确的提示类型
    if (
      hint.hintType !== DerivedInsightHintType.REUSE_OPPORTUNITY &&
      hint.hintType !== DerivedInsightHintType.CAPABILITY_GAP
    ) {
      return false
    }
    // 必须有解释文案
    if (!hint.title || !hint.explanationText) {
      return false
    }
    // 必须有稳定 CTA
    if (
      hint.ctaKind !== CTAKind.CTA_KIND_CREATE_PRODUCT_FROM_TEMPLATE &&
      hint.ctaKind !== CTAKind.CTA_KIND_GO_TO_MODULE_DETAIL
    ) {
      return false
    }
    // 必须有 template_candidate_id
    if (!hint.templateCandidateId) {
      return false
    }
    return true
  }, [])

  /** 计算单个提示的导航目标 */
  const computeHandoff = useCallback(
    (hint: DerivedInsightHint): HintHandoffResult | null => {
      if (!isValidHint(hint)) return null

      const candidateId = hint.templateCandidateId

      // reuse_opportunity_hint -> Product Create
      if (
        hint.hintType === DerivedInsightHintType.REUSE_OPPORTUNITY &&
        hint.ctaKind === CTAKind.CTA_KIND_CREATE_PRODUCT_FROM_TEMPLATE
      ) {
        const search: Record<string, unknown> = {
          fromTemplateReuse: true,
          templateCandidateId: candidateId,
          templateSource: returnContext.templateSource ?? 'weekly-review',
        }
        // 携带 Dashboard 返回链元数据
        if (returnContext.fromDashboard) {
          search.fromDashboard = true
          search.dashboardSection = returnContext.dashboardSection
          search.dashboardReturnTo = returnContext.dashboardReturnTo
        }
        return {
          to: '/products/new',
          search,
        }
      }

      // capability_gap_hint -> Module Registry / Module Detail
      if (
        hint.hintType === DerivedInsightHintType.CAPABILITY_GAP &&
        hint.ctaKind === CTAKind.CTA_KIND_GO_TO_MODULE_DETAIL
      ) {
        // 构建返回链上下文
        const returnSearch: Record<string, unknown> = {
          returnTo: returnContext.sourceSurface,
          returnCandidateId: returnContext.activeCandidateId ?? candidateId,
        }
        // Product Create 会话返回恢复
        if (returnContext.sourceSurface === 'product-create') {
          returnSearch.fromTemplateReuse = returnContext.fromTemplateReuse ?? true
          returnSearch.templateCandidateId = returnContext.templateCandidateId ?? candidateId
          returnSearch.templateSource = returnContext.templateSource ?? 'weekly-review'
        }
        // Dashboard 返回链
        if (returnContext.fromDashboard) {
          returnSearch.fromDashboard = true
          returnSearch.dashboardSection = returnContext.dashboardSection
          returnSearch.dashboardReturnTo = returnContext.dashboardReturnTo
        }

        // 如果提示指定了具体 moduleId，跳转到 Module Detail
        if (hint.moduleId) {
          return {
            to: '/modules/$moduleId',
            params: { moduleId: hint.moduleId },
            search: returnSearch,
          }
        }

        // 否则跳转到 Module Registry 列表
        return {
          to: '/modules',
          search: returnSearch,
        }
      }

      return null
    },
    [isValidHint, returnContext],
  )

  return { computeHandoff, isValidHint }
}