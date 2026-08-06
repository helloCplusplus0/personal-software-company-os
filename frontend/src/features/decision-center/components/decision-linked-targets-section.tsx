/**
 * DecisionLinkedTargetsSection — 已关联目标区
 *
 * §5.8 详情读取中的 linked_modules 展示。
 * phase03-05 组件树冻结：只承接已建立的 Decision -> Module 关联结果展示。
 */
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import { Badge } from '@/components/ui/badge'
import type { LinkedModule } from '../types'

interface DecisionLinkedTargetsSectionProps {
  linkedModules: LinkedModule[]
}

export function DecisionLinkedTargetsSection({ linkedModules }: DecisionLinkedTargetsSectionProps) {
  return (
    <Card>
      <CardHeader>
        <CardTitle className="text-base">已关联模块</CardTitle>
      </CardHeader>
      <CardContent>
        {linkedModules.length === 0 ? (
          <p className="text-sm text-muted-foreground py-2">暂无已关联模块</p>
        ) : (
          <div className="flex flex-wrap gap-2">
            {linkedModules.map((lm) => (
              <Badge key={lm.module_id} variant="secondary">
                {lm.module_name}
              </Badge>
            ))}
          </div>
        )}
      </CardContent>
    </Card>
  )
}
