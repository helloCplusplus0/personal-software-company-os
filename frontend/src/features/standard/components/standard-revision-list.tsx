/**
 * StandardRevisionList — Standard revision 轻量列表
 *
 * phase14-05 §"四页面组件树必须冻结"（StandardDetailPage 尾区）：
 * - 每行 change_summary + created_at（text-xs text-muted-foreground）；divide-y 分隔
 * - 空态"暂无 revision 记录"；纯只读展示
 */
import type { StandardRevision } from '../types'

interface StandardRevisionListProps {
  revisions: StandardRevision[]
}

export function StandardRevisionList({ revisions }: StandardRevisionListProps) {
  if (revisions.length === 0) {
    return <p className="text-xs text-muted-foreground">暂无 revision 记录</p>
  }
  return (
    <div className="divide-y">
      {revisions.map((revision) => (
        <div key={revision.id} className="flex flex-wrap items-baseline justify-between gap-x-2 py-1.5">
          <p className="min-w-0 flex-1 truncate text-xs">{revision.change_summary}</p>
          <span className="shrink-0 text-xs text-muted-foreground">
            {new Date(revision.created_at).toLocaleString()}
          </span>
        </div>
      ))}
    </div>
  )
}
