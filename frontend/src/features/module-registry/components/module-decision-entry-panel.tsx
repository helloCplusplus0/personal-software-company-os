import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import { Badge } from '@/components/ui/badge'
import { Button } from '@/components/ui/button'
import { Link } from '@tanstack/react-router'
import { Plus, List, ExternalLink } from 'lucide-react'
import type { DecisionLink } from '../types'

interface ModuleDecisionEntryPanelProps {
  /** 当前 Module 标识 — 用于带上下文导航到 Decision Create */
  moduleId: string
  /** 当前 Module 名称 — 用于来源上下文展示 */
  moduleName: string
  /** 当前 Module 已关联的决策列表 */
  decisionLinks: DecisionLink[]
}

/**
 * ModuleDecisionEntryPanel — Decision 入口面板
 *
 * phase03-13 spec §"Module Detail 入口组件必须升级为正式触点"：
 * - 从只读展示升级为 Decision Center 的正式入口触点
 * - "为当前 Module 记录决策" -> 导航到带 sourceModuleId / sourceModuleName 的 /decisions/new
 * - "查看当前 Module 相关决策" -> 导航到 /decisions
 * - 当前已展示的相关决策列表项应可直接进入对应的 DecisionDetailPage
 * - 不得在 Module Detail 侧新增中间路由或中间分发组件
 *
 * 保留当前 Module Detail 作为单一宿主页面。
 */
export function ModuleDecisionEntryPanel({ moduleId, moduleName, decisionLinks }: ModuleDecisionEntryPanelProps) {
  return (
    <Card>
      <CardHeader>
        <div className="flex items-center justify-between">
          <CardTitle>相关决策</CardTitle>
          {/* 两个正式入口动作 */}
          <div className="flex gap-2">
            <Button size="sm" variant="outline" asChild>
              <Link
                to="/decisions/new"
                search={{
                  sourceModuleId: moduleId,
                  sourceModuleName: moduleName,
                }}
              >
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
        {decisionLinks.length === 0 ? (
          <p className="text-sm text-muted-foreground py-2">暂无相关决策</p>
        ) : (
          <div className="space-y-2">
            {decisionLinks.map((dl) => (
              <Link
                key={dl.decision_id}
                to="/decisions/$decisionId"
                params={{ decisionId: dl.decision_id }}
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
