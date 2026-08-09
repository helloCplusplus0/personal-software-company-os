/**
 * CurrentFocusSection — Current Focus / Next Action 区块
 *
 * phase05-13 体验修复：
 * - 区块间距收紧（space-y-2），标题降级（text-base font-semibold）
 * - 反馈信号列表从「space-y-2 + 完整 Card」改为「divide-y + 紧凑行」容器
 * - 骨架尺寸收紧，匹配紧凑行高度
 *
 * phase05-13 §"四个区块组件必须按 phase05-06 状态模型实现"
 * phase05-10 §4.2 Current Focus / Next Action 区块：
 *   - 承接主行动队列区块容器与标题语义
 *   - 只消费归一化后的反馈信号卡片
 *   - 包含组件：CurrentFocusSection → FeedbackSignalCardList → FeedbackSignalCard
 *   - 最多展示 5 条主队列反馈卡片
 *
 * 状态模型（phase05-06）：
 *   - loading：骨架或加载提示
 *   - ready：以 FeedbackSignalCardList 展示 current_focus_signals（最多 5 条）
 *   - empty：展示"暂无待处理反馈信号"成功空态
 *   - error：区块内容区域内展示局部错误与重试入口
 *
 * 与 AssetFeedbackSection 的关系（phase05-10 §4.3）：
 *   - 两者共享同一个 feedbackQueryState，但不共用卡片列表
 *   - CurrentFocus 消费 current_focus_signals
 *   - AssetFeedback 消费 asset_feedback_summary.representative_signals
 *   - AssetFeedback 不形成第二条独立优先级队列
 */
import { Button } from '@/components/ui/button'
import { Skeleton } from '@/components/ui/skeleton'
import type { FeedbackSignal } from '../types'
import { FeedbackSignalCard } from './feedback-signal-card'

interface CurrentFocusSectionProps {
  // feedback query 状态
  status: 'loading' | 'ready' | 'empty' | 'error'
  signals: FeedbackSignal[]
  error: Error | null
  onRetry: () => void
}

/**
 * CurrentFocusSection — 主行动队列区块容器。
 *
 * 移动浏览器顺序（phase05-10 §3.3 / §9.x）：
 * - CurrentFocus 在第一屏优先位置
 */
export function CurrentFocusSection({
  status,
  signals,
  error,
  onRetry,
}: CurrentFocusSectionProps) {
  return (
    <section className="space-y-2" aria-label="Current Focus / Next Action">
      <h2 className="text-base font-semibold">Current Focus</h2>

      {status === 'loading' && (
        <div className="divide-y divide-border overflow-hidden rounded-lg border">
          {Array.from({ length: 3 }).map((_, i) => (
            <div key={i} className="px-3 py-2">
              <Skeleton className="h-5 w-full" />
            </div>
          ))}
        </div>
      )}

      {status === 'error' && (
        <div className="rounded-lg border border-destructive/50 bg-destructive/10 p-3">
          <p className="text-xs text-destructive">
            反馈信号读取失败：{error?.message ?? '未知错误'}
          </p>
          <Button variant="outline" size="sm" className="mt-2 h-7" onClick={onRetry}>
            重试
          </Button>
        </div>
      )}

      {status === 'empty' && (
        <div className="rounded-lg border border-dashed p-4 text-center">
          <p className="text-xs text-muted-foreground">暂无待处理反馈信号</p>
        </div>
      )}

      {status === 'ready' && <FeedbackSignalCardList signals={signals} />}
    </section>
  )
}

/**
 * FeedbackSignalCardList — 反馈信号紧凑行列表。
 *
 * 后端已按优先级排序并截断到最多 5 条，前端直接渲染。
 * 使用 divide-y 容器承接行间分隔，行本身不渲染边框。
 * 卡片 section 固定为 'current-focus'。
 */
function FeedbackSignalCardList({ signals }: { signals: FeedbackSignal[] }) {
  if (signals.length === 0) {
    return (
      <div className="rounded-lg border border-dashed p-4 text-center">
        <p className="text-xs text-muted-foreground">暂无待处理反馈信号</p>
      </div>
    )
  }

  return (
    <div className="divide-y divide-border overflow-hidden rounded-lg border">
      {signals.map((signal, index) => (
        <FeedbackSignalCard
          key={`${signal.signal_code}-${signal.target_id}-${index}`}
          signal={signal}
          section="current-focus"
        />
      ))}
    </div>
  )
}
