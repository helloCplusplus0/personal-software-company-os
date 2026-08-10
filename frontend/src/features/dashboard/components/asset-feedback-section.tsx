/**
 * AssetFeedbackSection — Asset Feedback 区块（代表性缺口项 + 复用快照）
 *
 * phase05-13 体验修复：
 * - 移除原 AssetCoverageCountGrid（四类缺口计数已并入 DashboardStatBar）
 * - 本区块只保留代表性缺口项列表（紧凑行，复用 FeedbackSignalCard）
 * - 区块间距收紧，标题降级
 * - status 现基于 representative_signals 独立派生（不再与 CurrentFocus 共享 empty 判定）
 *
 * phase05-13 §"四个区块组件必须按 phase05-06 状态模型实现"
 * phase05-10 §4.3 Asset Feedback 区块：
 *   - 承接补充摘要区块容器与标题语义
 *   - 只消费 product_asset_coverage 的代表性缺口项
 *   - 最多展示 3 条代表性缺口项
 *   - 不得与 Current Focus 复用成一个无语义差别的区块容器
 *   - 不得在页面级升级为并列主 CTA
 *
 * phase06-15 §"Dashboard 复用快照挂接位"：
 *   - Asset Feedback 区块内部增加独立 Reuse Snapshot 子区域
 *   - Reuse Snapshot 由页面级 ReuseSummaryRead query 提供数据
 *   - 该 query 失败不得把整页打回 page-error，也不得让 Asset Feedback 整体回退
 *   - module_reuse_summary 与 capability_summary 按"数量优先、时间次级"排序，最多 5 条
 *
 * 状态模型（phase05-06）：
 *   - loading：骨架
 *   - ready：以 AssetFeedbackList 展示 representative_signals（最多 3 条）
 *   - empty：展示"暂无代表性缺口项"成功空态
 *   - error：区块内容区域内展示局部错误与重试入口
 */
import { Button } from '@/components/ui/button'
import { Skeleton } from '@/components/ui/skeleton'
import type { FeedbackSignal } from '../types'
import { FeedbackSignalCard } from './feedback-signal-card'
import { ReuseSnapshotSection } from './reuse-snapshot-section'
import type {
  ModuleReuseSummaryEntry,
  CapabilitySummaryEntry,
} from '@/features/reuse-summary/types'

interface AssetFeedbackSectionProps {
  // 资产缺口代表性信号 query 状态（基于 representative_signals 独立派生）
  status: 'loading' | 'ready' | 'empty' | 'error'
  // representative_signals，由调用方从 asset_feedback_summary 取出传入
  signals: FeedbackSignal[]
  error: Error | null
  onRetry: () => void
  // phase06-15：复用快照子区域 props（独立页面级 query，失败不回退整区块）
  reuseSnapshotStatus: 'loading' | 'ready' | 'empty' | 'error'
  moduleReuseSummary: ModuleReuseSummaryEntry[]
  capabilitySummary: CapabilitySummaryEntry[]
  reuseSnapshotError: Error | null
  onReuseSnapshotRetry: () => void
}

/**
 * AssetFeedbackSection — 资产代表性缺口项区块容器。
 *
 * 区块布局：
 * - 标题：Asset Feedback
 * - 内容：最多 3 条代表性缺口项紧凑行
 * - phase06-15：内容底部追加独立 Reuse Snapshot 子区域
 *
 * 计数展示已移至 DashboardStatBar 的「资产覆盖」分组，本区块不再渲染计数网格。
 */
export function AssetFeedbackSection({
  status,
  signals,
  error,
  onRetry,
  reuseSnapshotStatus,
  moduleReuseSummary,
  capabilitySummary,
  reuseSnapshotError,
  onReuseSnapshotRetry,
}: AssetFeedbackSectionProps) {
  return (
    <section className="space-y-2" aria-label="Asset Feedback">
      <h2 className="text-base font-semibold">Asset Feedback</h2>

      {status === 'loading' && (
        <div className="divide-y divide-border overflow-hidden rounded-lg border">
          {Array.from({ length: 2 }).map((_, i) => (
            <div key={i} className="px-3 py-2">
              <Skeleton className="h-5 w-full" />
            </div>
          ))}
        </div>
      )}

      {status === 'error' && (
        <div className="rounded-lg border border-destructive/50 bg-destructive/10 p-3">
          <p className="text-xs text-destructive">
            资产缺口读取失败：{error?.message ?? '未知错误'}
          </p>
          <Button variant="outline" size="sm" className="mt-2 h-7" onClick={onRetry}>
            重试
          </Button>
        </div>
      )}

      {status === 'empty' && (
        <div className="rounded-lg border border-dashed p-4 text-center">
          <p className="text-xs text-muted-foreground">暂无代表性缺口项</p>
        </div>
      )}

      {status === 'ready' && <AssetFeedbackList signals={signals} />}

        {/*
          phase06-15 §"Dashboard 复用快照挂接位"：
          Reuse Snapshot 作为 Asset Feedback 内独立子区域，必须始终独立渲染自己的
          loading / ready / empty / error 状态，不能被代表性缺口项的 loading 覆盖。
        */}
        <div className="mt-3 rounded-lg border bg-muted/20 p-3">
          <ReuseSnapshotSection
            status={reuseSnapshotStatus}
            moduleReuseSummary={moduleReuseSummary}
            capabilitySummary={capabilitySummary}
            error={reuseSnapshotError}
            onRetry={onReuseSnapshotRetry}
          />
        </div>
    </section>
  )
}

/**
 * AssetFeedbackList — 代表性缺口项紧凑行列表。
 *
 * 后端已按优先级排序并截断到最多 3 条，前端直接渲染。
 * 复用 FeedbackSignalCard，但 section 固定为 'asset-feedback'。
 * 使用 divide-y 容器承接行间分隔。
 *
 * 若 representative_signals 为空列表，展示"暂无代表性缺口项"成功空态。
 */
function AssetFeedbackList({ signals }: { signals: FeedbackSignal[] }) {
  if (signals.length === 0) {
    return (
      <div className="rounded-lg border border-dashed p-4 text-center">
        <p className="text-xs text-muted-foreground">暂无代表性缺口项</p>
      </div>
    )
  }

  return (
    <div className="divide-y divide-border overflow-hidden rounded-lg border">
      {signals.map((signal, index) => (
        <FeedbackSignalCard
          key={`${signal.signal_code}-${signal.target_id}-${index}`}
          signal={signal}
          section="asset-feedback"
        />
      ))}
    </div>
  )
}
