/**
 * ReviewPageShell — Review 页面壳层。
 *
 * phase08-08 §"Daily / Weekly Review 最小 enablement"：
 *   ReviewPageShell 承接头部、离开入口、页面级状态与底部动作区。
 *
 * phase08-08 UI 对齐 Dashboard 基线（验收后调整）：
 *   - 壳层结构对齐 dashboard：space-y-4 整体收紧
 *   - 头部支持 subtitle 副标题，传达 review 会话语义
 *   - 头部移动端响应式：flex-col → sm:flex-row
 *   - 错误态/加载态样式与 dashboard 完全一致
 *
 * 职责：
 *   - 只消费 read owner 派生状态与 action owner 结果，不直连底层 hooks
 *   - 不额外持有 review read owner、review action owner、或页面级状态机
 */
import type { ReactNode } from 'react'
import { Button } from '@/components/ui/button'
import { Skeleton } from '@/components/ui/skeleton'
import { ArrowLeft } from 'lucide-react'
import type { PageStatus } from '../data/use-daily-review-read'

interface ReviewPageShellProps {
  /** 页面标题（如 "Daily Review" / "Weekly Review"） */
  title: string
  /** 副标题：传达当前 review 会话语义（如 "今日要处理的焦点与决策"） */
  subtitle?: string
  /** 页面级状态 */
  pageStatus: PageStatus
  /** 返回 Dashboard 的回调 */
  onBackToDashboard: () => void
  /** 整页重试回调（由 read owner 暴露），page-error 时显示重试按钮 */
  onRetry?: () => void
  /** 页面主体内容 */
  children: ReactNode
  /** 底部动作区 */
  actionFooter: ReactNode
}

export function ReviewPageShell({
  title,
  subtitle,
  pageStatus,
  onBackToDashboard,
  onRetry,
  children,
  actionFooter,
}: ReviewPageShellProps) {
  return (
    <div className="space-y-4">
      {/* 头部：返回按钮 + 标题 + 副标题
          移动端纵向堆叠，桌面端横向排列 */}
      <div className="flex flex-col gap-2 sm:flex-row sm:items-center sm:justify-between">
        <div className="flex flex-col gap-1 min-w-0">
          <div className="flex items-center gap-2 min-w-0">
            <Button
              variant="ghost"
              size="sm"
              onClick={onBackToDashboard}
              className="h-8 shrink-0 px-2"
            >
              <ArrowLeft className="h-4 w-4" />
              <span className="text-xs">Dashboard</span>
            </Button>
            <h1 className="text-xl font-bold truncate">{title}</h1>
          </div>
          {subtitle && (
            <p className="text-xs text-muted-foreground truncate pl-10 sm:pl-0">
              {subtitle}
            </p>
          )}
        </div>
      </div>

      {/* 页面级加载态：骨架对齐 dashboard */}
      {pageStatus === 'initial-loading' && (
        <div className="space-y-4">
          <Skeleton className="h-7 w-48" />
          <Skeleton className="h-32 w-full" />
          <Skeleton className="h-32 w-full" />
        </div>
      )}

      {/* 页面级错误态：对齐 dashboard 错误态样式 */}
      {pageStatus === 'page-error' && (
        <div className="rounded-lg border border-destructive/50 bg-destructive/10 p-4">
          <p className="text-sm text-destructive mb-3">
            Review 上下文加载失败，请稍后重试
          </p>
          <div className="flex items-center gap-2">
            {onRetry && (
              <Button
                variant="default"
                size="sm"
                onClick={onRetry}
                className="h-7"
              >
                重试
              </Button>
            )}
            <Button
              variant="outline"
              size="sm"
              onClick={onBackToDashboard}
              className="h-7"
            >
              返回 Dashboard
            </Button>
          </div>
        </div>
      )}

      {/* 页面内容 + 底部动作区 */}
      {pageStatus === 'ready' && (
        <>
          <div className="space-y-4">
            {children}
          </div>

          {/* 底部动作区 */}
          {actionFooter}
        </>
      )}
    </div>
  )
}
