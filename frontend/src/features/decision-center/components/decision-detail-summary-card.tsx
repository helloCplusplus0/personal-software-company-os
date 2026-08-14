/**
 * DecisionDetailSummaryCard — 决策概要卡片
 *
 * §5.8 详情读取承接核心对象字段、结构化模板字段与最小来源上下文展示。
 * phase03-05 组件树冻结：只承接决策核心字段、结构化模板字段与 source_context 展示。
 * fix_002_003：新增状态推进 CTA，承接 canonical 状态推进动作。
 *
 * 布局降级（phase03-05）：
 * - PC：字段分区展示
 * - 移动：单列垂直展示
 */
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import { Badge } from '@/components/ui/badge'
import { Button } from '@/components/ui/button'
import type { Decision, DecisionStatus, SourceContext } from '../types'
import type { DecisionDetailStatusAction } from '../data/use-decision-detail-page-read'

interface DecisionDetailSummaryCardProps {
  decision: Decision
  sourceContext: SourceContext
  statusActions?: DecisionDetailStatusAction[]
  /** 状态推进 CTA 回调，终态时不展示 */
  onStatusChange?: (status: DecisionStatus) => void
  /** 状态推进是否正在执行 */
  isUpdating?: boolean
}

const STATUS_LABEL: Record<DecisionStatus, string> = {
  proposed: 'Proposed',
  active: 'Active',
  superseded: 'Superseded',
  archived: 'Archived',
}

const STATUS_VARIANT: Record<DecisionStatus, 'default' | 'secondary' | 'outline' | 'destructive'> = {
  proposed: 'secondary',
  active: 'default',
  superseded: 'outline',
  archived: 'outline',
}

export function DecisionDetailSummaryCard({
  decision,
  sourceContext,
  statusActions = [],
  onStatusChange,
  isUpdating,
}: DecisionDetailSummaryCardProps) {
  const hasTransitions = statusActions.length > 0

  return (
    <Card>
      <CardHeader>
        <div className="flex items-start justify-between gap-2">
          <CardTitle className="text-lg">{decision.title}</CardTitle>
          <Badge variant={STATUS_VARIANT[decision.status]}>
            {STATUS_LABEL[decision.status]}
          </Badge>
        </div>
      </CardHeader>
      <CardContent className="space-y-4">
        {/* 结构化模板字段 */}
        <div>
          <h4 className="text-sm font-medium text-muted-foreground mb-1">上下文</h4>
          <p className="text-sm">{decision.context}</p>
        </div>
        <div>
          <h4 className="text-sm font-medium text-muted-foreground mb-1">问题</h4>
          <p className="text-sm">{decision.problem}</p>
        </div>
        {decision.alternatives.length > 0 && (
          <div>
            <h4 className="text-sm font-medium text-muted-foreground mb-1">备选方案</h4>
            <ol className="space-y-1">
              {decision.alternatives.map((alt, i) => (
                <li key={i} className="text-sm">
                  <span className="text-muted-foreground mr-2">{i + 1}.</span>
                  {alt}
                </li>
              ))}
            </ol>
          </div>
        )}
        <div>
          <h4 className="text-sm font-medium text-muted-foreground mb-1">选择</h4>
          <p className="text-sm">{decision.choice}</p>
        </div>
        <div>
          <h4 className="text-sm font-medium text-muted-foreground mb-1">理由</h4>
          <p className="text-sm">{decision.reason}</p>
        </div>
        {decision.impact && (
          <div>
            <h4 className="text-sm font-medium text-muted-foreground mb-1">影响</h4>
            <p className="text-sm">{decision.impact}</p>
          </div>
        )}

        {/* 来源上下文 — §5.11 */}
        {sourceContext.source_module_id && (
          <div className="border-t pt-3">
            <h4 className="text-sm font-medium text-muted-foreground mb-1">来源上下文</h4>
            <p className="text-sm">
              从 <span className="font-medium">{sourceContext.source_module_name}</span> 发起
            </p>
          </div>
        )}

        {/* fix_002_003：状态推进 CTA，仅在非终态时展示 */}
        {hasTransitions && onStatusChange && (
          <div className="border-t pt-3">
            <h4 className="text-sm font-medium text-muted-foreground mb-2">Status</h4>
            <div className="flex flex-wrap gap-2">
              {statusActions.map((t) => (
                <Button
                  key={t.target}
                  variant="outline"
                  size="sm"
                  disabled={isUpdating}
                  onClick={() => onStatusChange(t.target)}
                >
                  {t.label}
                </Button>
              ))}
            </div>
          </div>
        )}

        {/* 创建时间 */}
        <div className="border-t pt-3 text-xs text-muted-foreground">
          创建于 {new Date(decision.created_at).toLocaleString('zh-CN')}
        </div>
      </CardContent>
    </Card>
  )
}
