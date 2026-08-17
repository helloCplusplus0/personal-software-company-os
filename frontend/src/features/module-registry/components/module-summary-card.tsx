import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import { Badge } from '@/components/ui/badge'
import type { Module } from '../types'
import { MODULE_SEMANTIC_LABEL } from '@/features/project-context/data/shared-semantic-constants'

interface ModuleSummaryCardProps {
  module: Module
}

/**
 * ModuleSummaryCard — 模块摘要卡片
 * §8.3 只承接模块核心字段与状态表达
 * §8.3.1 默认归属于 ModuleDetailPage
 * phase12-08：增加"可复用能力资产"语义标签，字段文案对齐
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
        <div className="flex items-center gap-2">
          <span className="text-[10px] font-medium text-muted-foreground uppercase tracking-wide">
            {MODULE_SEMANTIC_LABEL}
          </span>
        </div>
        <div>
          <p className="text-xs text-muted-foreground">能力描述</p>
          <p className="text-sm">{module.description}</p>
        </div>
        <div>
          <p className="text-xs text-muted-foreground">登记时间</p>
          <p className="text-sm">{new Date(module.created_at).toLocaleDateString()}</p>
        </div>
      </CardContent>
    </Card>
  )
}
