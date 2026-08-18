/**
 * StandardReadonlySummary — Repository detail 挂载的 Standard 只读摘要
 *
 * phase14-05 §ADDED-6（画像区让位承接位）：
 * - 数据源：use-repository-standards-read（GetProjectBrief.standards[] 投影，与 agent 消费同源）
 * - 每条 Standard：name（链接 /standards/:id）+ status Badge + compact 树形只读
 * - 纯只读：无任何绑定操作入口（裁决⑦——绑定发起控件仅存在于 StandardDetailPage）
 * - 过渡态说明：phase14-08 落地至 phase14-09 迁移完成之间，旧画像数据尚未迁入时
 *   显示空态，属设计内过渡态
 */
import { Link } from '@tanstack/react-router'
import { useRepositoryStandardsRead } from '../data/use-repository-standards-read'
import { StandardTreeView } from './standard-tree-view'
import { Badge } from '@/components/ui/badge'

interface StandardReadonlySummaryProps {
  repositoryId: string
}

export function StandardReadonlySummary({ repositoryId }: StandardReadonlySummaryProps) {
  const { data, isLoading, isError } = useRepositoryStandardsRead(repositoryId)

  return (
    <section className="min-w-0 space-y-2">
      <p className="text-xs font-medium text-muted-foreground">关联 Standard</p>
      {isLoading ? (
        <p className="p-4 text-xs text-muted-foreground">加载中...</p>
      ) : isError ? (
        <p className="p-4 text-xs text-destructive">关联 Standard 读取失败</p>
      ) : !data || data.length === 0 ? (
        <p className="p-4 text-xs text-muted-foreground">该仓库尚未关联 Standard</p>
      ) : (
        <div className="space-y-3">
          {data.map((standard) => (
            <div key={standard.id} className="min-w-0 space-y-1.5">
              <div className="flex min-w-0 items-center gap-2">
                <Link
                  to="/standards/$standardId"
                  params={{ standardId: standard.id }}
                  className="min-w-0 truncate text-sm font-medium hover:underline"
                >
                  {standard.name}
                </Link>
                <Badge variant={standard.status === 'active' ? 'default' : 'secondary'}>
                  {standard.status}
                </Badge>
              </div>
              <StandardTreeView tree={standard.directory_tree} compact />
            </div>
          ))}
        </div>
      )}
    </section>
  )
}
