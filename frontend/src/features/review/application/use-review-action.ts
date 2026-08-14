/**
 * useReviewAction — Review 动作的正式 application owner。
 *
 * phase08-06 §"review 写路径必须收敛到单一 Review action application owner"：
 *   本 hook 是 review 完成区唯一正式写路径 owner，
 *   统一承接 success envelope、query invalidation 与错误归一化。
 *
 * phase08-08 §"review action owner 的正式职责"：
 *   - 纯 decision / entity handoff 路径只返回 success envelope 并导航到既有 canonical 页面
 *   - next-step result 路径必须正式调用 review service 的 SubmitReviewResult
 *   - 所有成功 envelope 必须继续透传 fromDashboard / dashboardSection / dashboardReturnTo
 *   - query invalidation 必须统一收敛在 onSuccess 内通过 queryClient.invalidateQueries() 触发
 *
 * 约束：
 *   - route / page / section / card 不得直接持有第二套 useMutation
 *   - 本 hook 不得直接持有页面局部 UI 状态
 */
import { useMutation, useQueryClient, type UseMutationResult } from '@tanstack/react-query'
import { reviewClient } from '../data/connect-client'
import { ReviewKind, ReviewResultKind } from '@/gen/proto/psco/review/v1/review_pb'
import { buildReviewSourceParams } from '../lib/review-source'
import { DAILY_REVIEW_QUERY_KEY, WEEKLY_REVIEW_QUERY_KEY } from '../data/review-query-options'
import type {
  ReviewActionInput,
  ReviewActionSuccess,
} from './review-action-types'


function normalizeError(err: unknown): Error {
  if (err instanceof Error) {
    return err
  }
  return new Error('Review 动作执行失败，请稍后重试')
}

function buildSuccessEnvelope(
  input: ReviewActionInput,
): ReviewActionSuccess {
  const sourceParams = buildReviewSourceParams(input.reviewKind ?? 'daily')

  switch (input.actionType) {
    case 'create_decision':
      return {
        resultKind: 'decision_handoff',
        navigateTo: '/decisions/new',
        search: sourceParams,
        successMessage: undefined,
      }

    case 'go_to_decision':
      return {
        resultKind: 'decision_handoff',
        navigateTo: input.targetId
          ? `/decisions/${input.targetId}`
          : '/decisions',
        params: input.targetId ? { decisionId: input.targetId } : undefined,
        search: sourceParams,
        successMessage: undefined,
      }

    case 'go_to_product':
      return {
        resultKind: 'entity_handoff',
        navigateTo: input.targetId
          ? `/products/${input.targetId}`
          : '/products',
        params: input.targetId ? { productId: input.targetId } : undefined,
        search: sourceParams,
        successMessage: undefined,
      }

    case 'go_to_module':
      return {
        resultKind: 'entity_handoff',
        navigateTo: input.targetId
          ? `/modules/${input.targetId}`
          : '/modules',
        params: input.targetId ? { moduleId: input.targetId } : undefined,
        search: sourceParams,
        successMessage: undefined,
      }

    case 'go_to_repository':
      return {
        resultKind: 'entity_handoff',
        navigateTo: input.targetId
          ? `/repositories/${input.targetId}`
          : '/repositories',
        params: input.targetId ? { repositoryId: input.targetId } : undefined,
        search: sourceParams,
        successMessage: undefined,
      }

    case 'submit_next_step':
      return {
        resultKind: 'next_step_result',
        navigateTo: '/dashboard',
        search: sourceParams,
        successMessage: 'Review 结果已记录',
      }

    default:
      return {
        resultKind: 'entity_handoff',
        navigateTo: '/dashboard',
        search: sourceParams,
        successMessage: undefined,
      }
  }
}

export type UseReviewAction = UseMutationResult<
  ReviewActionSuccess,
  Error,
  ReviewActionInput,
  unknown
>

/**
 * useReviewAction — Review 动作正式 owner。
 *
 * 职责：
 *   - 纯 route handoff 动作（decision/entity handoff, navigate_only）直接返回 success envelope
 *   - next-step result 动作调用 reviewClient.submitReviewResult 正式持久化
 *   - 成功后统一失效 review 与 dashboard 相关 query
 *   - 错误归一化
 */
export function useReviewAction(): UseReviewAction {
  const queryClient = useQueryClient()

  return useMutation<ReviewActionSuccess, Error, ReviewActionInput, unknown>({
    mutationFn: async (input: ReviewActionInput) => {
      try {
        // 纯 route handoff 路径：无需 mutation，直接返回 success envelope
        if (input.actionType !== 'submit_next_step') {
          return buildSuccessEnvelope(input)
        }

        // next-step result 路径：调用 review service 持久化
        const now = new Date()
        await reviewClient.submitReviewResult({
          reviewKind: input.reviewKind === 'weekly' ? ReviewKind.WEEKLY : ReviewKind.DAILY,
          resultKind: ReviewResultKind.NEXT_STEP_RESULT,
          decisionId: input.targetId ?? '',
          summaryText: input.summaryText ?? '',
          startedAt: { seconds: BigInt(Math.floor(now.getTime() / 1000)), nanos: 0 },
          completedAt: { seconds: BigInt(Math.floor(now.getTime() / 1000)), nanos: 0 },
        })

        return buildSuccessEnvelope(input)
      } catch (err) {
        throw normalizeError(err)
      }
    },
      onSuccess: async () => {
        // 统一等待相关读取完成失效，避免页面拿到旧的 review/dashboard 快照。
        await Promise.all([
          queryClient.invalidateQueries({ queryKey: DAILY_REVIEW_QUERY_KEY }),
          queryClient.invalidateQueries({ queryKey: WEEKLY_REVIEW_QUERY_KEY }),
          queryClient.invalidateQueries({ queryKey: ['dashboard-overview'] }),
          queryClient.invalidateQueries({ queryKey: ['dashboard-feedback-signals'] }),
          queryClient.invalidateQueries({ queryKey: ['dashboard-recent-activities'] }),
        ])
    },
  })
}
