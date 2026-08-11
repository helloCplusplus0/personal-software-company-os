/**
 * ReviewActionFooter — Review 页面底部动作区。
 *
 * phase08-08 §"Daily / Weekly Review 最小 enablement"：
 *   ReviewActionFooter 承接 review 完成区的最小动作触发与 useReviewAction() 的正式消费。
 *
 * phase08-08 UI 对齐 Dashboard 基线（验收后调整）：
 *   - 分层为主行动区 + 完成区，对齐 phase08-05 spec §"review 完成与主动离开的交互流必须单值化"
 *   - 主行动区：进入决策（主按钮）+ 实体 handoff 次按钮组（产品/模块/仓库）
 *   - 完成区：完成 Review（submit_next_step，secondary 风格）+ 说明文案
 *   - 移动端按钮纵向堆叠 w-full，桌面端紧凑排列
 *   - 错误态样式对齐 dashboard 错误态
 *
 * 职责：
 *   - 只承接 review 完成区的动作按钮编排
 *   - 不得额外持有 review read owner、页面级状态机或独立的 mutation 编排器
 */
import { Button } from '@/components/ui/button'
import { ArrowRight, CheckCircle } from 'lucide-react'
import type { ReviewActionInput } from '../application/review-action-types'

interface ReviewActionFooterProps {
  /** 是否正在提交 */
  isPending: boolean
  /** 是否有错误 */
  hasError: boolean
  /** 错误消息 */
  errorMessage?: string
  /** 触发 review 动作 */
  onSubmitAction: (input: ReviewActionInput) => void
  /** 重置错误状态 */
  onReset: () => void
  /** 当前 review 会话类型，用于 submit_next_step 时区分 daily/weekly */
  reviewKind?: 'daily' | 'weekly'
}

export function ReviewActionFooter({
  isPending,
  hasError,
  errorMessage,
  onSubmitAction,
  onReset,
  reviewKind,
}: ReviewActionFooterProps) {
  return (
    <div className="border-t pt-3 space-y-3">
      {/* 错误态：对齐 dashboard 错误态样式 */}
      {hasError && (
        <div className="rounded-lg border border-destructive/50 bg-destructive/10 p-3">
          <p className="text-xs text-destructive">
            {errorMessage ?? '操作失败，请稍后重试'}
          </p>
          <Button
            variant="outline"
            size="sm"
            onClick={onReset}
            className="mt-2 h-7"
          >
            重试
          </Button>
        </div>
      )}

      {/* 主行动区：进入决策（主按钮）+ 实体 handoff 次按钮组
          移动端纵向堆叠 w-full，桌面端紧凑排列 */}
      <div className="flex flex-col gap-2 sm:flex-row sm:flex-wrap sm:items-center">
        <Button
          size="sm"
          variant="default"
          disabled={isPending}
          onClick={() => onSubmitAction({
            actionType: 'go_to_decision',
            dashboardSection: 'current-focus',
          })}
          className="w-full sm:w-auto h-9 shrink-0"
        >
          进入决策
          <ArrowRight className="ml-1.5 h-4 w-4" />
        </Button>

        <div className="flex flex-col gap-2 sm:flex-row sm:flex-wrap sm:items-center">
          <Button
            size="sm"
            variant="outline"
            disabled={isPending}
            onClick={() => onSubmitAction({
              actionType: 'go_to_product',
              dashboardSection: 'asset-feedback',
            })}
            className="w-full sm:w-auto h-8 shrink-0 text-xs"
          >
            产品
            <ArrowRight className="ml-1.5 h-3.5 w-3.5" />
          </Button>
          <Button
            size="sm"
            variant="outline"
            disabled={isPending}
            onClick={() => onSubmitAction({
              actionType: 'go_to_module',
              dashboardSection: 'current-focus',
            })}
            className="w-full sm:w-auto h-8 shrink-0 text-xs"
          >
            模块
            <ArrowRight className="ml-1.5 h-3.5 w-3.5" />
          </Button>
          <Button
            size="sm"
            variant="outline"
            disabled={isPending}
            onClick={() => onSubmitAction({
              actionType: 'go_to_repository',
              dashboardSection: 'asset-feedback',
            })}
            className="w-full sm:w-auto h-8 shrink-0 text-xs"
          >
            仓库
            <ArrowRight className="ml-1.5 h-3.5 w-3.5" />
          </Button>
        </div>
      </div>

      {/* 完成区：完成 Review（submit_next_step）
          phase08-05 spec §"review 形成 next-step result"：
          next-step result 必须落到轻量 review_records，不能只停留在页面局部状态 */}
      <div className="flex flex-col gap-2 sm:flex-row sm:items-center sm:justify-between">
        <p className="text-xs text-muted-foreground">
          完成会记录一条 next-step 结果到 review_records
        </p>
        <Button
          size="sm"
          variant="secondary"
          disabled={isPending}
          onClick={() => onSubmitAction({
            actionType: 'submit_next_step',
            dashboardSection: 'current-focus',
            summaryText: 'Review 完成',
            reviewKind,
          })}
          className="w-full sm:w-auto h-9 shrink-0"
        >
          <CheckCircle className="mr-1.5 h-4 w-4" />
          {isPending ? '提交中...' : '完成 Review'}
        </Button>
      </div>
    </div>
  )
}
