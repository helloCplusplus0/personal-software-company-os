/**
 * DecisionContextSourcePanel — 来源上下文展示面板
 *
 * §5.11 入口上下文与正式关联结果边界：
 * - 从 Module Detail 带上下文进入 Decision Create 时，展示该来源 Module
 * - 该来源上下文只表示"待关联来源"，不等于已建立正式 decision_links
 * - 来源 Module 标识将通过 CreateDecisionRequest.source_module_id 持久化
 *
 * phase03-05 组件树冻结：只承接从 Module Detail 带入的来源 Module 展示。
 */
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import { Badge } from '@/components/ui/badge'
import { Info } from 'lucide-react'

interface DecisionContextSourcePanelProps {
  /** 来源 Module 标识 */
  sourceModuleId: string
  /** 来源 Module 名称 */
  sourceModuleName: string
}

export function DecisionContextSourcePanel({ sourceModuleId, sourceModuleName }: DecisionContextSourcePanelProps) {
  return (
    <Card>
      <CardHeader>
        <CardTitle className="flex items-center gap-2 text-base">
          <Info className="h-4 w-4" />
          来源上下文
        </CardTitle>
      </CardHeader>
      <CardContent>
        <p className="text-sm text-muted-foreground mb-2">
          本决策从以下 Module 发起，创建后将作为待关联目标继续承接：
        </p>
        <div className="flex items-center gap-2">
          <Badge variant="secondary">{sourceModuleName}</Badge>
          <span className="text-xs text-muted-foreground font-mono">{sourceModuleId}</span>
        </div>
        <p className="mt-2 text-xs text-muted-foreground">
          该来源只表示"待关联来源"，不等于已建立正式关联。可在决策详情页完成正式关联。
        </p>
      </CardContent>
    </Card>
  )
}
