/**
 * ProgressTimelineList — 推进事件时间轴（三轨过滤 + 事件行 + 误录删除）
 *
 * phase15-05 §"时间轴列表与误录删除交互规格必须冻结"：
 * - 三轨过滤：纯组件 state（无路由参数，URL 不变），切换即变更 query key
 *   第三段自动重新查询
 * - 前端不重排序（后端三键链倒序直渲）、不过滤（过滤唯一经 query 参数）
 * - 误录删除：window.confirm 确认（文案逐字冻结，沿 standard-detail-page 先例，
 *   不引入 Dialog 组件）后经 use-delete-progress-event（append-only 语义，
 *   整条删除是唯一修正路径，裁决⑨）
 */
import { useState } from 'react'
import { useProgressEventsRead } from '../data/use-progress-events-read'
import type { ProgressFilter } from '../data/use-progress-events-read'
import { useDeleteProgressEvent } from '../application/use-delete-progress-event'
import type { ProgressEvent, WorkflowType } from '../types'
import { Badge } from '@/components/ui/badge'
import { Button } from '@/components/ui/button'
import { Trash2 } from 'lucide-react'

interface ProgressTimelineListProps {
  repositoryId: string
}

/** 三轨过滤按钮组（全部 / Phase / Audit / Fix） */
const FILTER_OPTIONS: { value: ProgressFilter; label: string }[] = [
  { value: 'all', label: '全部' },
  { value: 'phase', label: 'Phase' },
  { value: 'audit', label: 'Audit' },
  { value: 'fix', label: 'Fix' },
]

/** event_kind Badge variant 按 workflow_type 区分：phase=default / audit=secondary / fix=outline */
function eventKindBadgeVariant(workflowType: WorkflowType): 'default' | 'secondary' | 'outline' {
  if (workflowType === 'audit') return 'secondary'
  if (workflowType === 'fix') return 'outline'
  return 'default'
}

export function ProgressTimelineList({ repositoryId }: ProgressTimelineListProps) {
  const [filter, setFilter] = useState<ProgressFilter>('all')
  const { data, isLoading, isError } = useProgressEventsRead(repositoryId, filter)
  const deleteMutation = useDeleteProgressEvent()

  /** 误录删除确认 — 文案逐字冻结；确认后经 owner（无软删除、无二次修正入口） */
  const handleDelete = (event: ProgressEvent) => {
    if (!window.confirm(`确认删除事件「${event.title}」？删除仅用于修正误录，操作不可撤销。`)) return
    deleteMutation.mutate({ id: event.id, repositoryId })
  }

  const events = data ?? []

  return (
    <div className="space-y-2">
      {/* 三轨过滤按钮组 — active 高亮 / inactive ghost；纯组件 state，URL 不变 */}
      <div className="flex flex-wrap items-center gap-1.5">
        {FILTER_OPTIONS.map((opt) => (
          <Button
            key={opt.value}
            type="button"
            variant={filter === opt.value ? 'default' : 'ghost'}
            className="h-7 px-2 text-xs"
            onClick={() => setFilter(opt.value)}
          >
            {opt.label}
          </Button>
        ))}
      </div>

      {/* 删除失败行内错误回显 */}
      {deleteMutation.isError ? (
        <p className="text-xs text-destructive">{(deleteMutation.error as Error).message}</p>
      ) : null}

      {isLoading ? (
        <p className="text-xs text-muted-foreground">加载中...</p>
      ) : isError ? (
        <p className="text-xs text-destructive">进度事件读取失败</p>
      ) : events.length === 0 ? (
        <p className="text-xs text-muted-foreground">
          {filter === 'all' ? '暂无推进事件，从上方表单录入第一条。' : '该轨暂无事件。'}
        </p>
      ) : (
        // 2026-08-19 用户第三轮 UI 反馈：列表限高滚动，超过约 6 条事件后
        // 垂直滚动预览，不再无限下撑页面
        <div className="max-h-80 divide-y overflow-y-auto">
          {events.map((event) => (
            <div key={event.id} className="min-w-0 space-y-1 p-2">
              {/* 第一行：event_kind Badge + workflow_type + task_key + occurred_at + source + 删除 */}
              <div className="flex flex-wrap items-center gap-1.5 text-xs">
                <Badge variant={eventKindBadgeVariant(event.workflow_type)}>{event.event_kind}</Badge>
                <span className="text-muted-foreground">{event.workflow_type}</span>
                {event.task_key ? <code className="text-[10px]">{event.task_key}</code> : null}
                <span className="ml-auto text-muted-foreground">
                  {event.occurred_at?.toLocaleString()}
                </span>
                <Badge variant="outline" className="text-[10px]">
                  {event.source}
                </Badge>
                <Button
                  type="button"
                  variant="ghost"
                  size="icon-xs"
                  aria-label={`删除事件 ${event.title}`}
                  disabled={deleteMutation.isPending}
                  onClick={() => handleDelete(event)}
                >
                  <Trash2 />
                </Button>
              </div>
              {/* 第二行：title */}
              <p className="truncate text-xs font-medium">{event.title}</p>
              {/* 第三行（可选，detail 非空时） */}
              {event.detail ? (
                <p className="line-clamp-2 text-xs text-muted-foreground">{event.detail}</p>
              ) : null}
              {/* 第四行（可选，evidence_ref 非空时）——2026-08-19 用户第三轮反馈：
                  仅指明仓库目录大概位置；https:// 真实外链保留可点击，
                  `/` 开头仓库内路径前端无对应路由，渲染为纯文本标注不提供跳转 */}
              {event.evidence_ref ? (
                event.evidence_ref.startsWith('https://') ? (
                  <a
                    href={event.evidence_ref}
                    target="_blank"
                    rel="noreferrer"
                    className="block truncate text-xs underline"
                  >
                    {event.evidence_ref}
                  </a>
                ) : (
                  <p className="block truncate text-xs text-muted-foreground">
                    {event.evidence_ref}
                  </p>
                )
              ) : null}
            </div>
          ))}
        </div>
      )}
    </div>
  )
}
