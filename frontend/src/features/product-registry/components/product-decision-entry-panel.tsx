import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import { Button } from '@/components/ui/button'
import { Link } from '@tanstack/react-router'
import { Plus, List, ExternalLink } from 'lucide-react'
import { Badge } from '@/components/ui/badge'

interface ProductDecisionEntryPanelProps {
  decisionLinks: { decision_id: string; decision_title: string }[]
  decisionDetailSearch?: Record<string, unknown>
  isLoading?: boolean
  isError?: boolean
}

/**
 * ProductDecisionEntryPanel — Product Detail 的 Decision 入口面板
 *
 * phase10-10 §"Product Detail Decision 入口面板"：
 * - 承接"为当前 Product 记录决策"与"查看当前 Product 相关决策"两类正式入口
 * - "记录决策" → 导航到带 sourceProductId / sourceProductName 的 /decisions/new
 * - "查看全部" → 导航到 /decisions
 * - 已关联决策列表可直接进入对应 DecisionDetailPage
 */
export function ProductDecisionEntryPanel({
  decisionLinks,
  decisionDetailSearch,
  isLoading = false,
  isError = false,
}: ProductDecisionEntryPanelProps) {
  return (
    <Card>
      <CardHeader>
        <div className="flex items-center justify-between">
          <CardTitle>相关决策</CardTitle>
          <div className="flex gap-2">
            <Button size="sm" variant="outline" asChild>
              <Link to="/decisions/new">
                <Plus className="mr-1 h-3 w-3" />
                记录决策
              </Link>
            </Button>
            <Button size="sm" variant="ghost" asChild>
              <Link to="/decisions">
                <List className="mr-1 h-3 w-3" />
                查看全部
              </Link>
            </Button>
          </div>
        </div>
      </CardHeader>
      <CardContent>
        {isLoading ? (
          <p className="text-sm text-muted-foreground py-2">加载相关决策...</p>
        ) : isError ? (
          <p className="text-sm text-destructive py-2">相关决策加载失败</p>
        ) : decisionLinks.length === 0 ? (
          <p className="text-sm text-muted-foreground py-2">暂无相关决策</p>
        ) : (
          <div className="space-y-2">
            {decisionLinks.map((dl) => (
              <Link
                key={dl.decision_id}
                to="/decisions/$decisionId"
                params={{ decisionId: dl.decision_id }}
                search={decisionDetailSearch}
                className="flex items-center justify-between rounded-md border p-3 hover:bg-accent transition-colors"
              >
                <span className="text-sm">{dl.decision_title}</span>
                <Badge variant="outline" className="gap-1">
                  <ExternalLink className="h-3 w-3" />
                  查看
                </Badge>
              </Link>
            ))}
          </div>
        )}
      </CardContent>
    </Card>
  )
}
