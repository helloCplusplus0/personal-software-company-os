/**
 * Dashboard CTA 优先级矩阵命中计算
 *
 * phase05-13 §"DashboardPrimaryActionPanel 必须按 CTA 矩阵命中"
 * phase05-04 / phase05-10 §7.2 已冻结 CTA 优先级矩阵：
 *
 * 顺序 | 命中条件                                          | 主 CTA 目标
 * -----|---------------------------------------------------|------------------------------
 * CTA1 | module=0 && product=0 && repository=0 && decision=0 | Module Registry / Create
 * CTA2 | module=0 && (product>0 || repository>0 || decision>0) | Module Registry / Create
 * CTA3 | module>0 && product=0                              | Product Registry / Create
 * CTA4 | module>0 && product>0 && repository=0              | Repository Binding / Create
 * CTA5 | 存在 pending_decision 信号                         | 最高优先级决策信号落点
 * CTA6 | 存在 product_missing_both_bindings 信号            | 对应 Product Detail
 * CTA7 | 存在 product_missing_repository_binding 信号       | 对应 Product Detail
 * CTA8 | 存在 product_missing_module_binding 信号           | 对应 Product Detail
 * CTA9 | 无缺口且有活动数据                                  | 无主 CTA（hidden）
 *
 * 约束：
 * - 同一时刻只允许展示一个主 CTA
 * - CTA 1-4 是创建导向（overview 即可命中，不依赖 feedback query）
 * - CTA 5-8 是缺口导向（依赖 feedback query）
 * - CTA 5-8 内部按优先级顺序命中，不并排展示
 */
import type {
  DashboardOverview,
  FeedbackSignal,
  FeedbackSignalCode,
  DashboardTargetType,
  DashboardSection,
} from '../types'

/**
 * CTA 编号标识，对齐 phase05-10 §7.2 优先级矩阵。
 */
export type CtaId = 'CTA_1' | 'CTA_2' | 'CTA_3' | 'CTA_4' | 'CTA_5' | 'CTA_6' | 'CTA_7' | 'CTA_8' | 'CTA_9'

/**
 * CTA 创建导向目标路由（CTA 1-4 命中时使用）。
 */
export type CreateTargetRoute = '/modules/new' | '/products/new' | '/repositories/new'

/**
 * ComputePrimaryCtaResult — computePrimaryCta 返回值。
 *
 * - state: 'ready' 表示命中 CTA 1-8，需要渲染主 CTA
 * - state: 'hidden' 表示命中 CTA 9（无缺口），不渲染主 CTA
 * - state: 'suppressed' 表示 feedback query 失败导致 CTA 5-8 无法判定
 *   （此状态由面板组件根据 query 状态推断，computePrimaryCta 不直接返回）
 */
export type ComputePrimaryCtaResult =
  | {
      state: 'ready'
      ctaId: CtaId
      // CTA 1-4：创建导向目标
      createTarget?: CreateTargetRoute
      // CTA 5-8：缺口导向目标（FeedbackSignal 已包含 target_type / target_id）
      feedbackSignal?: FeedbackSignal
      // 主 CTA 按钮文案，由 cta-matrix 单值化，避免组件层各自硬编码
      actionLabel: string
      // 跳转目标路由（创建导向为 Create 路由，缺口导向为 Detail 路由）
      to: string
      // 跳转目标类型，用于组件层决定 params 形状
      targetType: DashboardTargetType
      // dashboardSection：CTA 1-4 为 'empty-state'，CTA 5-8 为 'current-focus'
      dashboardSection: DashboardSection
    }
  | {
      state: 'hidden'
      ctaId: 'CTA_9'
    }

/**
 * 从 FeedbackSignal 列表中按 signal_code 找到第一条匹配信号。
 * 用于 CTA 5-8 命中时取出对应缺口信号。
 */
function findSignalByCode(
  signals: FeedbackSignal[],
  code: FeedbackSignalCode,
): FeedbackSignal | undefined {
  return signals.find((s) => s.signal_code === code)
}

/**
 * 计算主 CTA 命中结果。
 *
 * 入参：
 * - overview: DashboardOverview（必填，CTA 1-4 判定前提）
 * - currentFocusSignals: FeedbackSignal[]（CTA 5 判定前提）
 * - assetFeedbackRepresentativeSignals: FeedbackSignal[]（CTA 6-8 判定前提，
 *   来自 ProductAssetCoverageSummary.representative_signals）
 *
 * 返回：
 * - 命中 CTA 1-8：{ state: 'ready', ... }
 * - 命中 CTA 9：{ state: 'hidden', ctaId: 'CTA_9' }
 *
 * 注意：本函数不承担 feedback query 失败判定。query 失败时由面板组件
 * 直接进入 'suppressed' 状态，不调用本函数。
 */
export function computePrimaryCta(
  overview: DashboardOverview,
  currentFocusSignals: FeedbackSignal[],
  assetFeedbackRepresentativeSignals: FeedbackSignal[],
): ComputePrimaryCtaResult {
  const {
    module_count,
    product_count,
    repository_count,
    decision_count,
  } = overview

  // CTA 1：全空冷启动
  if (module_count === 0 && product_count === 0 && repository_count === 0 && decision_count === 0) {
    return {
      state: 'ready',
      ctaId: 'CTA_1',
      createTarget: '/modules/new',
      actionLabel: '登记首个模块',
      to: '/modules/new',
      targetType: 'module_detail', // CTA 1-4 是创建导向，targetType 仅用于组件层区分
      dashboardSection: 'empty-state',
    }
  }

  // CTA 2：非空但无 Module
  if (module_count === 0 && (product_count > 0 || repository_count > 0 || decision_count > 0)) {
    return {
      state: 'ready',
      ctaId: 'CTA_2',
      createTarget: '/modules/new',
      actionLabel: '补登记模块',
      to: '/modules/new',
      targetType: 'module_detail',
      dashboardSection: 'empty-state',
    }
  }

  // CTA 3：有 Module 但无 Product
  if (module_count > 0 && product_count === 0) {
    return {
      state: 'ready',
      ctaId: 'CTA_3',
      createTarget: '/products/new',
      actionLabel: '登记首个产品',
      to: '/products/new',
      targetType: 'product_detail',
      dashboardSection: 'empty-state',
    }
  }

  // CTA 4：有 Module 与 Product 但无 Repository
  if (module_count > 0 && product_count > 0 && repository_count === 0) {
    return {
      state: 'ready',
      ctaId: 'CTA_4',
      createTarget: '/repositories/new',
      actionLabel: '登记首个仓库',
      to: '/repositories/new',
      targetType: 'repository_detail',
      dashboardSection: 'empty-state',
    }
  }

  // CTA 5：存在 pending_decision 信号（最高优先级缺口）
  // currentFocusSignals 已由后端按优先级排序，第一条 pending_decision 即最高优先级
  const pendingDecisionSignal = findSignalByCode(currentFocusSignals, 'pending_decision')
  if (pendingDecisionSignal) {
    return {
      state: 'ready',
      ctaId: 'CTA_5',
      feedbackSignal: pendingDecisionSignal,
      actionLabel: pendingDecisionSignal.action_label,
      to:
        pendingDecisionSignal.target_type === 'decision_detail'
          ? '/decisions/$decisionId'
          : '/decisions',
      targetType: pendingDecisionSignal.target_type,
      dashboardSection: 'current-focus',
    }
  }

  // CTA 6-8：从 asset_feedback_summary.representative_signals 中按优先级顺序查找
  // representative_signals 已由后端按优先级排序，按 CTA 6 → 7 → 8 顺序命中
  const cta6Signal = findSignalByCode(assetFeedbackRepresentativeSignals, 'product_missing_both_bindings')
  if (cta6Signal) {
    return {
      state: 'ready',
      ctaId: 'CTA_6',
      feedbackSignal: cta6Signal,
      actionLabel: cta6Signal.action_label,
      to: '/products/$productId',
      targetType: cta6Signal.target_type,
      dashboardSection: 'current-focus',
    }
  }

  const cta7Signal = findSignalByCode(assetFeedbackRepresentativeSignals, 'product_missing_repository_binding')
  if (cta7Signal) {
    return {
      state: 'ready',
      ctaId: 'CTA_7',
      feedbackSignal: cta7Signal,
      actionLabel: cta7Signal.action_label,
      to: '/products/$productId',
      targetType: cta7Signal.target_type,
      dashboardSection: 'current-focus',
    }
  }

  const cta8Signal = findSignalByCode(assetFeedbackRepresentativeSignals, 'product_missing_module_binding')
  if (cta8Signal) {
    return {
      state: 'ready',
      ctaId: 'CTA_8',
      feedbackSignal: cta8Signal,
      actionLabel: cta8Signal.action_label,
      to: '/products/$productId',
      targetType: cta8Signal.target_type,
      dashboardSection: 'current-focus',
    }
  }

  // CTA 9：无缺口，无主 CTA
  return {
    state: 'hidden',
    ctaId: 'CTA_9',
  }
}
