import { Link } from '@tanstack/react-router'
import { ArrowRight } from 'lucide-react'
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import { Button } from '@/components/ui/button'
import {
  PRODUCT_SEMANTIC_LABEL,
  REPOSITORY_SEMANTIC_LABEL,
  MODULE_SEMANTIC_LABEL,
  DECISION_SEMANTIC_LABEL,
} from '../data/shared-semantic-constants'

interface SemanticEntry {
  entityLabel: string
  semanticLabel: string
  helperText: string
  to: '/products' | '/repositories' | '/modules' | '/decisions'
}

const SEMANTIC_ENTRIES: SemanticEntry[] = [
  {
    entityLabel: 'Product',
    semanticLabel: PRODUCT_SEMANTIC_LABEL,
    helperText: '回看经营目标、交付对象与当前项目承接范围。',
    to: '/products',
  },
  {
    entityLabel: 'Repository',
    semanticLabel: REPOSITORY_SEMANTIC_LABEL,
    helperText: '回看代码锚点、绑定关系与共享项目上下文。',
    to: '/repositories',
  },
  {
    entityLabel: 'Module',
    semanticLabel: MODULE_SEMANTIC_LABEL,
    helperText: '回看可复用能力资产及其产品、仓库与决策关系。',
    to: '/modules',
  },
  {
    entityLabel: 'Decision',
    semanticLabel: DECISION_SEMANTIC_LABEL,
    helperText: '回看规则、约束、选择与依据的留痕结果。',
    to: '/decisions',
  },
]

interface ProjectSemanticSummaryPanelProps {
  title?: string
  description?: string
}

export function ProjectSemanticSummaryPanel({
  title = '当前项目四实体角色',
  description = '通过同一套共享语义与固定入口回看当前项目。',
}: ProjectSemanticSummaryPanelProps) {
  return (
    <Card>
      <CardHeader>
        <CardTitle className="text-base">{title}</CardTitle>
        <p className="text-sm text-muted-foreground">{description}</p>
      </CardHeader>
      <CardContent className="space-y-3">
        {SEMANTIC_ENTRIES.map((entry) => (
          <div
            key={entry.entityLabel}
            className="flex flex-col gap-2 rounded-md border bg-muted/20 px-3 py-2 sm:flex-row sm:items-center sm:justify-between"
          >
            <div className="min-w-0 space-y-1">
              <div className="text-sm font-medium">
                {entry.entityLabel}: {entry.semanticLabel}
              </div>
              <p className="text-xs text-muted-foreground">{entry.helperText}</p>
            </div>
            <Button variant="outline" size="sm" asChild>
              <Link to={entry.to}>
                <ArrowRight className="mr-2 h-4 w-4" />
                查看{entry.entityLabel}
              </Link>
            </Button>
          </div>
        ))}
      </CardContent>
    </Card>
  )
}
