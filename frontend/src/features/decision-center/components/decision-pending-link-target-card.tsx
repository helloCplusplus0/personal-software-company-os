/**
 * DecisionPendingLinkTargetCard — 待关联目标卡片
 *
 * §5.11 入口上下文与正式关联结果边界：
 * - 当 source_context.source_module_id 存在且当前 linked_modules 中尚无该 Module 时，
 *   必须显式展示该待关联目标
 * - 不得在进入详情页后静默丢失该待关联目标
 * - 待关联目标仅在正式 LinkDecisionToTarget 写入后消失（由页面 reread 驱动）
 *
 * phase03-05 组件树冻结：
 * - 只承接从 Module Detail 带入的入口上下文中尚未完成正式关联的待关联 Module
 * - 必须作为显式待关联目标持续展示，直到用户完成正式 LinkDecisionToTarget
 *
 * phase03-13 spec 非目标（结束条件收敛）：
 * - 当前阶段不提供“主动放弃关联”出口
 * - source_context 作为入口历史记录保留，后端无清除接口，
 *   待关联目标持续展示直到正式关联完成，不因用户放弃而清除
 */
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import { Badge } from '@/components/ui/badge'
import { Link2 } from 'lucide-react'

interface DecisionPendingLinkTargetCardProps {
  /** 待关联 Module 标识 */
  sourceModuleId: string
  /** 待关联 Module 名称 */
  sourceModuleName: string
}

export function DecisionPendingLinkTargetCard({ sourceModuleId, sourceModuleName }: DecisionPendingLinkTargetCardProps) {
  return (
    <Card className="border-primary/30 bg-primary/5">
      <CardHeader>
        <CardTitle className="flex items-center gap-2 text-base">
          <Link2 className="h-4 w-4 text-primary" />
          待关联来源
        </CardTitle>
      </CardHeader>
      <CardContent>
        <p className="text-sm text-muted-foreground mb-2">
          本决策从以下 Module 发起，尚未建立正式关联：
        </p>
        <div className="flex items-center gap-2">
          <Badge variant="outline" className="border-primary/50">
            {sourceModuleName}
          </Badge>
          <span className="text-xs text-muted-foreground font-mono">{sourceModuleId}</span>
        </div>
        <p className="mt-2 text-xs text-muted-foreground">
          可在下方候选列表中完成正式关联，关联后此卡片将自动消失。
        </p>
      </CardContent>
    </Card>
  )
}
