import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import { Badge } from '@/components/ui/badge'
import type { Repository } from '../types'

interface RepositorySummaryCardProps {
  repository: Repository
}

/**
 * RepositorySummaryCard — 仓库摘要卡片
 * phase04-05 组件树冻结：只承接 Repository 核心字段 id / name / url / provider / status / created_at
 * phase04-05 默认归属于 RepositoryBindingDetailPage
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
        <div>
          <p className="text-sm text-muted-foreground">URL</p>
          <p className="break-all text-sm">{repository.url}</p>
        </div>
        <div>
          <p className="text-sm text-muted-foreground">提供商</p>
          <p className="text-sm">{repository.provider}</p>
        </div>
        <div>
          <p className="text-sm text-muted-foreground">创建时间</p>
          <p className="text-sm">{new Date(repository.created_at).toLocaleDateString()}</p>
        </div>
      </CardContent>
    </Card>
  )
}
