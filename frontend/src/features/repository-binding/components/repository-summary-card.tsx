import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import { Badge } from '@/components/ui/badge'
import type { Repository } from '../types'
import { REPOSITORY_SEMANTIC_LABEL } from '@/features/project-context/data/shared-semantic-constants'

interface RepositorySummaryCardProps {
  repository: Repository
}

/**
 * RepositorySummaryCard — 仓库摘要卡片
 * phase04-05 组件树冻结：只承接 Repository 核心字段 id / name / url / provider / status / created_at
 * phase04-05 默认归属于 RepositoryBindingDetailPage
 * phase12-08：增加"代码仓库身份对象与项目锚点"语义标签
 */
export function RepositorySummaryCard({ repository }: RepositorySummaryCardProps) {
  return (
    <Card>
      <CardHeader>
        <CardTitle className="flex items-center justify-between gap-2">
          <span className="truncate">{repository.name}</span>
          <Badge variant={repository.status === 'active' ? 'default' : 'secondary'}>
            {repository.status}
          </Badge>
        </CardTitle>
      </CardHeader>
      <CardContent className="space-y-2">
        <div className="flex items-center gap-2">
          <span className="text-[10px] font-medium text-muted-foreground uppercase tracking-wide">
            {REPOSITORY_SEMANTIC_LABEL}
          </span>
        </div>
        <div>
          <p className="text-xs text-muted-foreground">URL</p>
          <p className="break-all text-sm">{repository.url}</p>
        </div>
        <div>
          <p className="text-xs text-muted-foreground">提供商</p>
          <p className="text-sm">{repository.provider}</p>
        </div>
        <div>
          <p className="text-xs text-muted-foreground">创建时间</p>
          <p className="text-sm">{new Date(repository.created_at).toLocaleDateString()}</p>
        </div>
      </CardContent>
    </Card>
  )
}
