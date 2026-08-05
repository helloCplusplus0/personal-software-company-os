import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import { Badge } from '@/components/ui/badge'
import { ExternalLink } from 'lucide-react'
import type { DecisionLink } from '../types'

interface ModuleDecisionEntryPanelProps {
  decisionLinks: DecisionLink[]
}

/**
 * ModuleDecisionEntryPanel — Decision 入口面板
 * §6.3 Decision 在当前阶段作为 ModuleDetailRead 的附属读取承接，不设独立读接口组
 * §4.2 只作为只读展示或跳转入口，不扩写为当前阶段独立写入主线
 * §8.3 默认归属于 ModuleDetailPage
 */
export function ModuleDecisionEntryPanel({ decisionLinks }: ModuleDecisionEntryPanelProps) {
  return (
    <Card>
      <CardHeader>
        <CardTitle>相关决策</CardTitle>
      </CardHeader>
      <CardContent>
        {decisionLinks.length === 0 ? (
          <p className="text-sm text-muted-foreground py-2">暂无相关决策</p>
        ) : (
          <div className="space-y-2">
            {decisionLinks.map((dl) => (
              <div
                key={dl.decision_id}
                className="flex items-center justify-between rounded-md border p-3"
              >
                <span className="text-sm">{dl.decision_title}</span>
                <Badge variant="outline" className="gap-1">
                  <ExternalLink className="h-3 w-3" />
                  跳转
                </Badge>
              </div>
            ))}
          </div>
        )}
      </CardContent>
    </Card>
  )
}
