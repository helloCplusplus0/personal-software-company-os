/**
 * DecisionLinkedTargetsSection — 已关联目标区
 *
 * §5.8 详情读取中的 linked_modules 展示。
 * phase03-05 组件树冻结：只承接已建立的 Decision -> Module 关联结果展示。
 */
import { Link } from '@tanstack/react-router'
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import { Badge } from '@/components/ui/badge'
import type { LinkedModule } from '../types'

interface DecisionLinkedTargetsSectionProps {
  linkedModules: LinkedModule[]
  moduleDetailSearch: Record<string, unknown>
}

export function DecisionLinkedTargetsSection({
  linkedModules,
  moduleDetailSearch,
}: DecisionLinkedTargetsSectionProps) {
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
                <Link
                  key={lm.module_id}
                  to="/modules/$moduleId"
                  params={{ moduleId: lm.module_id }}
                  search={moduleDetailSearch}
                  className="inline-block"
                >
                  <Badge variant="secondary" className="cursor-pointer hover:bg-accent transition-colors">
                    {lm.module_name}
                  </Badge>
                </Link>
            ))}
          </div>
        )}
      </CardContent>
    </Card>
  )
}
