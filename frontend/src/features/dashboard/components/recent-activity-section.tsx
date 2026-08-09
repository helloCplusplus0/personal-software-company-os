/**
 * RecentActivitySection — Recent Activity 区块
 *
 * phase05-13 体验修复：
 * - 限高 + 内部滚动：列表容器 max-h 约束（PC 420px / 移动 320px），
 *   超出部分内部滚动，活动流永不撑开页面，保住 dashboard 一屏特性
 * - 区块间距收紧，标题降级
 * - 活动项从完整 Card 改为紧凑行（divide-y 容器）
 *
 * phase05-13 §"四个区块组件必须按 phase05-06 状态模型实现"
 * phase05-10 §4.4 Recent Activity 区块：
 *   - 承接活动流区块容器与标题语义
 *   - 只按活动时间倒序排序，不参与反馈优先级竞争
 *   - 包含组件：RecentActivitySection → RecentActivityList → RecentActivityItemCard
 *   - 最多展示 10 条活动项
 *
 * 状态模型（phase05-06）：
 *   - loading：骨架或加载提示
 *   - ready：以 RecentActivityList 展示 activities（最多 10 条，后端已排序）
 *   - empty：展示"暂无最近活动"成功空态
 *   - error：区块内容区域内展示局部错误与重试入口
 */
import { Button } from '@/components/ui/button'
import { Skeleton } from '@/components/ui/skeleton'
import type { RecentActivityItem } from '../types'
import { RecentActivityItemCard } from './recent-activity-item-card'

interface RecentActivitySectionProps {
  // activity query 状态
  status: 'loading' | 'ready' | 'empty' | 'error'
  activities: RecentActivityItem[]
  error: Error | null
  onRetry: () => void
}

/**
 * RecentActivitySection — 独立活动流区块容器。
 *
 * 区块布局：
 * - 标题：Recent Activity
 * - 内容：最多 10 条活动项紧凑行，限高内部滚动
 *
 * 约束（phase05-10 §4.4）：
 * - 不与反馈信号共用排序逻辑
 * - 只按活动时间倒序排序（后端已承接）
 * - 限高滚动仅是视觉约束，不影响后端返回的 10 条上限语义
 */
export function RecentActivitySection({
  status,
  activities,
  error,
  onRetry,
}: RecentActivitySectionProps) {
  return (
    <section className="space-y-2" aria-label="Recent Activity">
      <h2 className="text-base font-semibold">Recent Activity</h2>

      {status === 'loading' && (
        // 限高容器内渲染骨架，与 ready 态高度一致
        <div className="max-h-[320px] divide-y divide-border overflow-y-auto rounded-lg border md:max-h-[420px]">
          {Array.from({ length: 4 }).map((_, i) => (
            <div key={i} className="px-3 py-2">
              <Skeleton className="h-5 w-full" />
            </div>
          ))}
        </div>
      )}

      {status === 'error' && (
        <div className="rounded-lg border border-destructive/50 bg-destructive/10 p-3">
          <p className="text-xs text-destructive">
            最近活动读取失败：{error?.message ?? '未知错误'}
          </p>
          <Button variant="outline" size="sm" className="mt-2 h-7" onClick={onRetry}>
            重试
          </Button>
        </div>
      )}

      {status === 'empty' && (
        <div className="rounded-lg border border-dashed p-4 text-center">
          <p className="text-xs text-muted-foreground">暂无最近活动</p>
        </div>
      )}

      {status === 'ready' && <RecentActivityList activities={activities} />}
    </section>
  )
}

/**
 * RecentActivityList — 活动项紧凑行列表（限高内部滚动）。
 *
 * 后端已按 activity_at DESC 排序并截断到最多 10 条，前端直接渲染。
 * 限高滚动：PC 420px / 移动 320px，超出部分内部滚动，永不撑开页面。
 */
function RecentActivityList({ activities }: { activities: RecentActivityItem[] }) {
  if (activities.length === 0) {
    return (
      <div className="rounded-lg border border-dashed p-4 text-center">
        <p className="text-xs text-muted-foreground">暂无最近活动</p>
      </div>
    )
  }

  return (
    <div className="max-h-[320px] divide-y divide-border overflow-y-auto rounded-lg border md:max-h-[420px]">
      {activities.map((activity, index) => (
        <RecentActivityItemCard
          key={`${activity.activity_type}-${activity.target_id}-${index}`}
          activity={activity}
        />
      ))}
    </div>
  )
}
