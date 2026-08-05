import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import { Badge } from '@/components/ui/badge'
import type { Module } from '../types'

interface ModuleSummaryCardProps {
  module: Module
}

/**
 * ModuleSummaryCard — 模块摘要卡片
 * §8.3 只承接模块核心字段与状态表达
 * §8.3.1 默认归属于 ModuleDetailPage
 */
export function ModuleSummaryCard({ module }: ModuleSummaryCardProps) {
  return (
    <Card>
      <CardHeader>
        <CardTitle className="flex items-center justify-between gap-2">
          <span className="truncate">{module.name}</span>
          <Badge variant={module.status === 'active' ? 'default' : 'secondary'}>
            {module.status}
          </Badge>
        </CardTitle>
      </CardHeader>
      <CardContent className="space-y-2">
        <div>
          <p className="text-sm text-muted-foreground">描述</p>
          <p className="text-sm">{module.description}</p>
        </div>
        <div>
          <p className="text-sm text-muted-foreground">创建时间</p>
          <p className="text-sm">{new Date(module.created_at).toLocaleDateString()}</p>
        </div>
      </CardContent>
    </Card>
  )
}
