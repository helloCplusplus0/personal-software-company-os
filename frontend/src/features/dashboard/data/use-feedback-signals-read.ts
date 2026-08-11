/**
 * useFeedbackSignalsRead — Feedback Signals 只读 query owner。
 *
 * phase07-10 §5.4：query 层纯只读，唯一 owner。
 * 替换 dashboard/data/api-adapter.ts 的 fetchFeedbackSignals。
 */
import { useQuery, type UseQueryResult } from '@tanstack/react-query'
import { dashboardClient } from './connect-client'
import {
  DashboardTargetType,
  FeedbackSignalCode,
  FeedbackSignalFamily,
  FeedbackSignalPriority,
} from '@/gen/proto/psco/dashboard/v1/dashboard_pb'
import type {
  DashboardTargetType as DashboardTargetTypeValue,
  FeedbackSignalCode as FeedbackSignalCodeValue,
  FeedbackSignalFamily as FeedbackSignalFamilyValue,
  FeedbackSignalPriority as FeedbackSignalPriorityValue,
  FeedbackSignalsResponse,
} from '../types'

export type UseFeedbackSignalsRead = UseQueryResult<FeedbackSignalsResponse, Error>

function mapSignalFamily(v: FeedbackSignalFamily): FeedbackSignalFamilyValue {
  switch (v) {
    case FeedbackSignalFamily.PENDING_DECISION:
      return 'pending_decision'
    case FeedbackSignalFamily.PRODUCT_ASSET_COVERAGE:
      return 'product_asset_coverage'
    default:
      return 'pending_decision'
  }
}

function mapSignalCode(v: FeedbackSignalCode): FeedbackSignalCodeValue {
  switch (v) {
    case FeedbackSignalCode.PENDING_DECISION:
      return 'pending_decision'
    case FeedbackSignalCode.PRODUCT_MISSING_BOTH_BINDINGS:
      return 'product_missing_both_bindings'
    case FeedbackSignalCode.PRODUCT_MISSING_REPOSITORY_BINDING:
      return 'product_missing_repository_binding'
    case FeedbackSignalCode.PRODUCT_MISSING_MODULE_BINDING:
      return 'product_missing_module_binding'
    default:
      return 'pending_decision'
  }
}

function mapSignalPriority(v: FeedbackSignalPriority): FeedbackSignalPriorityValue {
  switch (v) {
    case FeedbackSignalPriority.P1_PENDING_DECISION:
      return 'p1_pending_decision'
    case FeedbackSignalPriority.P2_PRODUCT_MISSING_BOTH_BINDINGS:
      return 'p2_product_missing_both_bindings'
    case FeedbackSignalPriority.P3_PRODUCT_MISSING_REPOSITORY_BINDING:
      return 'p3_product_missing_repository_binding'
    case FeedbackSignalPriority.P4_PRODUCT_MISSING_MODULE_BINDING:
      return 'p4_product_missing_module_binding'
    default:
      return 'p1_pending_decision'
  }
}

function mapTargetType(v: DashboardTargetType): DashboardTargetTypeValue {
  switch (v) {
    case DashboardTargetType.DECISION_DETAIL:
      return 'decision_detail'
    case DashboardTargetType.DECISION_LIST:
      return 'decision_list'
    case DashboardTargetType.PRODUCT_DETAIL:
      return 'product_detail'
    case DashboardTargetType.MODULE_DETAIL:
      return 'module_detail'
    case DashboardTargetType.REPOSITORY_DETAIL:
      return 'repository_detail'
    default:
      return 'decision_list'
  }
}

export function useFeedbackSignalsRead(): UseFeedbackSignalsRead {
  return useQuery<FeedbackSignalsResponse, Error>({
    queryKey: ['dashboard-feedback-signals'],
    queryFn: async (): Promise<FeedbackSignalsResponse> => {
      const res = await dashboardClient.getFeedbackSignals({})
      const mapSignal = (s: NonNullable<typeof res.currentFocusSignals>[number]) => ({
        signal_family: mapSignalFamily(s.signalFamily),
        signal_code: mapSignalCode(s.signalCode),
        priority: mapSignalPriority(s.priority),
        title: s.title ?? '',
        summary: s.summary ?? '',
        action_label: s.actionLabel ?? '',
        target_type: mapTargetType(s.targetType),
        target_id: s.targetId ?? '',
        target_label: s.targetLabel ?? '',
      })
      return {
        current_focus_signals: (res.currentFocusSignals ?? []).map(mapSignal),
        asset_feedback_summary: {
          fully_bound_product_count: res.assetFeedbackSummary?.fullyBoundProductCount ?? 0,
          missing_both_bindings_count: res.assetFeedbackSummary?.missingBothBindingsCount ?? 0,
          missing_repository_binding_count: res.assetFeedbackSummary?.missingRepositoryBindingCount ?? 0,
          missing_module_binding_count: res.assetFeedbackSummary?.missingModuleBindingCount ?? 0,
          representative_signals: (res.assetFeedbackSummary?.representativeSignals ?? []).map(mapSignal),
        },
      } as unknown as FeedbackSignalsResponse
    },
  })
}
