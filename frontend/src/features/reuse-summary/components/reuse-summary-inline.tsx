/**
 * ReuseSummaryInline — 详情页用的最小复用摘要内联组件
 *
 * phase06-15 §"Module Detail 与 Product Detail 挂接位"
 * phase06-09 §"Module Detail 与 Product Detail 挂接位"
 *
 * 约束：
 *   - Module Detail 在 Module Summary 邻近区域挂接
 *   - Product Detail 在已绑定模块相关区域附近挂接
 *   - 每个详情页只新增一个页面级 ReuseSummaryRead query
 *   - 详情页中的复用反馈不得承接绑定写入、解绑写入或候选筛选逻辑
 *   - query 失败不回退整页，只展示局部错误与重试入口
 *
 * 与 Dashboard 的 ReuseSnapshotSection 区别：
 *   - Dashboard 版本同时展示 module_reuse_summary 与 capability_summary
 *   - 详情页版本围绕指定 module_id / product_id 返回，展示紧凑摘要
 */
import { Button } from '@/components/ui/button'
import { Skeleton } from '@/components/ui/skeleton'
import { useQueryClient } from '@tanstack/react-query'

interface ReuseSummaryInlineProps {
  // 页面级 query 状态
  status: 'loading' | 'ready' | 'empty' | 'error'
  // 模块复用摘要条目（围绕指定 module_id / product_id）
  moduleReuseSummary: {
    module_id: string
    module_name: string
    reuse_product_count: number
    latest_reuse_at: string
    explanation_text: string
  }[]
  // 能力摘要条目
  capabilitySummary: {
    capability_key: string
    capability_label: string
    supporting_module_count: number
    latest_capability_update_at: string
    empty_state_text: string
  }[]
  error: Error | null
  // 失效用的 query key，由调用方传入
  invalidateQueryKey: readonly unknown[]
  // 标题（区分 Module / Product 上下文）
  title?: string
}

/**
 * ReuseSummaryInline — 详情页内联复用摘要组件。
 *
 * 用于 ModuleDetailPage / ProductDetailPage 在 ready 状态下挂接最小复用反馈。
 * 不承接任何写入逻辑，只承接只读展示与局部重试。
 */
export function ReuseSummaryInline({
  status,
  moduleReuseSummary,
  capabilitySummary,
  error,
  invalidateQueryKey,
  title = '复用摘要',
}: ReuseSummaryInlineProps) {
  const queryClient = useQueryClient()

  const handleRetry = () => {
    queryClient.invalidateQueries({ queryKey: invalidateQueryKey })
  }

  return (
    <section className="space-y-2" aria-label={title}>
      <h3 className="text-sm font-semibold text-muted-foreground">{title}</h3>

      {status === 'loading' && <Skeleton className="h-16 w-full" />}

      {status === 'error' && (
        <div className="rounded-md border border-destructive/30 bg-destructive/5 p-3 text-sm">
          <p className="text-xs text-destructive mb-1">
            复用摘要读取失败：{error?.message ?? '未知错误'}
          </p>
          <Button variant="outline" size="sm" className="h-7" onClick={handleRetry}>
            重试
          </Button>
        </div>
      )}

      {status === 'empty' && (
        <div className="rounded-md border border-dashed p-3 text-center">
          <p className="text-xs text-muted-foreground">暂无复用反馈</p>
        </div>
      )}

      {status === 'ready' && (
        <div className="space-y-2">
          {moduleReuseSummary.length > 0 && (
            <div className="space-y-1">
              {moduleReuseSummary.slice(0, 5).map((entry) => (
                <div
                  key={entry.module_id}
                  className="flex items-start gap-2 text-sm"
                >
                  <span className="inline-flex h-5 min-w-5 items-center justify-center rounded bg-primary/10 px-1 text-xs font-medium text-primary">
                    {entry.reuse_product_count}
                  </span>
                  <span className="text-foreground">{entry.explanation_text}</span>
                </div>
              ))}
            </div>
          )}

          {capabilitySummary.length > 0 && (
            <div className="space-y-1">
              {capabilitySummary.slice(0, 5).map((entry) => (
                <div
                  key={entry.capability_key}
                  className="flex items-start gap-2 text-sm"
                >
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

          {moduleReuseSummary.length === 0 && capabilitySummary.length === 0 && (
            <p className="text-xs text-muted-foreground">暂无复用反馈</p>
          )}
        </div>
      )}
    </section>
  )
}
