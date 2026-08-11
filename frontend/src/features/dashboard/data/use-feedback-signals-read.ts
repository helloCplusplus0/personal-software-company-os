/**
 * useFeedbackSignalsRead — Feedback Signals 只读 query owner。
 *
 * phase07-10 §5.4：query 层纯只读，唯一 owner。
 * 替换 dashboard/data/api-adapter.ts 的 fetchFeedbackSignals。
 */
import { useQuery, type UseQueryResult } from '@tanstack/react-query'
import { dashboardClient } from './connect-client'
import type { FeedbackSignalsResponse } from '../types'

export type UseFeedbackSignalsRead = UseQueryResult<FeedbackSignalsResponse, Error>

export function useFeedbackSignalsRead(): UseFeedbackSignalsRead {
  return useQuery<FeedbackSignalsResponse, Error>({
    queryKey: ['dashboard-feedback-signals'],
    queryFn: async (): Promise<FeedbackSignalsResponse> => {
      const res = await dashboardClient.getFeedbackSignals({})
      // eslint-disable-next-line @typescript-eslint/no-explicit-any
      const mapSignal = (s: any) => ({
        signal_family: s.signalFamily ?? '',
        signal_code: s.signalCode ?? '',
        priority: s.priority,
        title: s.title ?? '',
        summary: s.summary ?? '',
        action_label: s.actionLabel ?? '',
        target_type: s.targetType,
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