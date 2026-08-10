/**
 * ReuseSnapshotSection — Dashboard Asset Feedback 内的复用快照子区域
 *
 * phase06-15 §"Dashboard 复用快照挂接位"
 * phase06-09 §"Dashboard 复用快照挂接位"
 *
 * 约束：
 *   - 独立页面级 ReuseSummaryRead query（由 DashboardHomePage 传入 data）
 *   - 该 query 失败不得把整页打回 page-error
 *   - module_reuse_summary 与 capability_summary 都按"数量优先、时间次级"排序，最多展示前 5 条
 *   - 复用反馈必须可见且可解释，不得只显示数字而缺少解释文本
 */
import { Skeleton } from '@/components/ui/skeleton'
import type {
  ModuleReuseSummaryEntry,
  CapabilitySummaryEntry,
} from '@/features/reuse-summary/types'

interface ReuseSnapshotSectionProps {
  status: 'loading' | 'ready' | 'empty' | 'error'
  moduleReuseSummary: ModuleReuseSummaryEntry[]
  capabilitySummary: CapabilitySummaryEntry[]
  error: Error | null
  onRetry: () => void
}

export function ReuseSnapshotSection({
  status,
  moduleReuseSummary,
  capabilitySummary,
  error,
  onRetry,
}: ReuseSnapshotSectionProps) {
  if (status === 'loading') {
    return (
      <div className="space-y-2">
        <h4 className="text-sm font-semibold text-muted-foreground">Reuse Snapshot</h4>
        <Skeleton className="h-16 w-full" />
      </div>
    )
  }

  if (status === 'error') {
    return (
      <div className="space-y-2">
        <h4 className="text-sm font-semibold text-muted-foreground">Reuse Snapshot</h4>
        <div className="rounded-md border border-destructive/30 bg-destructive/5 p-3 text-sm">
          <p className="text-muted-foreground mb-1">复用摘要读取失败</p>
          <p className="text-xs text-destructive mb-2">{error?.message ?? '未知错误'}</p>
          <button
            type="button"
            onClick={onRetry}
            className="text-xs text-primary hover:underline"
          >
            重试
          </button>
        </div>
      </div>
    )
  }

  if (status === 'empty') {
    return (
      <div className="space-y-2">
        <h4 className="text-sm font-semibold text-muted-foreground">Reuse Snapshot</h4>
        <p className="text-sm text-muted-foreground">
          暂无复用反馈。创建模块并绑定到多个产品后，这里将展示复用摘要。
        </p>
      </div>
    )
  }

  return (
    <div className="space-y-3">
      <h4 className="text-sm font-semibold text-muted-foreground">Reuse Snapshot</h4>

      {/* 模块复用摘要（最多 5 条） */}
      {moduleReuseSummary.length > 0 && (
        <div className="space-y-1">
          <p className="text-xs font-medium text-muted-foreground">模块复用</p>
          {moduleReuseSummary.slice(0, 5).map((entry) => (
            <div key={entry.module_id} className="flex items-start gap-2 text-sm">
              <span className="inline-flex h-5 min-w-5 items-center justify-center rounded bg-primary/10 px-1 text-xs font-medium text-primary">
                {entry.reuse_product_count}
              </span>
              <span className="text-foreground">{entry.explanation_text}</span>
            </div>
          ))}
        </div>
      )}

      {/* 能力摘要（最多 5 条） */}
      {capabilitySummary.length > 0 && (
        <div className="space-y-1">
          <p className="text-xs font-medium text-muted-foreground">能力分布</p>
          {capabilitySummary.slice(0, 5).map((entry) => (
            <div key={entry.capability_key} className="flex items-start gap-2 text-sm">
              <span className="inline-flex h-5 min-w-5 items-center justify-center rounded bg-primary/10 px-1 text-xs font-medium text-primary">
                {entry.supporting_module_count}
              </span>
              <span className="text-foreground">
                {entry.capability_label}
                {entry.empty_state_text ? ` — ${entry.empty_state_text}` : ''}
              </span>
            </div>
          ))}
        </div>
      )}
    </div>
  )
}
