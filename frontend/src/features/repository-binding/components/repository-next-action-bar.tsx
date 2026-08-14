import { Button } from '@/components/ui/button'
import { ArrowRight, Boxes, GitBranch, ArrowLeft } from 'lucide-react'
import { Link } from '@tanstack/react-router'

interface RepositoryNextActionBarProps {
  /** 是否有已绑定 Product */
  hasProductBinding: boolean
  /** 是否有已映射 Module */
  hasModuleMapping: boolean
  /** 当前 Repository 标识 */
  repositoryId: string
  repositoryName: string
  decisionLinks?: { decision_id: string; decision_title: string }[]
  decisionDetailSearch?: Record<string, unknown>
  onOpenProductBinding?: () => void
  onOpenModuleMapping?: () => void
}

/**
 * RepositoryNextActionBar — Repository Detail 的页面级下一步动作区
 *
 * phase10-10 §"Repository Detail 下一步动作承接矩阵"：
 * - Product 绑定 > Module 映射 > 返回 Dashboard
 * - 主 CTA 以主按钮样式展示，指向当前页面已有的绑定面板
 */
export function RepositoryNextActionBar({
  hasProductBinding,
  hasModuleMapping,
  repositoryId: _repositoryId,
  repositoryName,
  decisionLinks = [],
  decisionDetailSearch,
  onOpenProductBinding,
  onOpenModuleMapping,
}: RepositoryNextActionBarProps) {
  // 优先级 1：Product 绑定缺口
  if (!hasProductBinding) {
    return (
      <div className="flex items-center gap-3 rounded-lg border bg-primary/5 p-3">
        <Boxes className="h-4 w-4 text-primary shrink-0" />
        <div className="flex-1 min-w-0">
          <p className="text-sm font-medium">下一步：绑定产品</p>
          <p className="text-xs text-muted-foreground truncate">
            将仓库 "{repositoryName}" 绑定到所属产品
          </p>
        </div>
        <Button
          size="sm"
          className="shrink-0"
          onClick={() => {
            onOpenProductBinding?.()
            document.getElementById('repository-product-binding')?.scrollIntoView({ behavior: 'smooth', block: 'start' })
          }}
        >
          绑定产品
          <ArrowRight className="ml-1 h-3 w-3" />
        </Button>
      </div>
    )
  }

  // 优先级 2：Module 映射缺口
  if (!hasModuleMapping) {
    return (
      <div className="flex items-center gap-3 rounded-lg border bg-primary/5 p-3">
        <GitBranch className="h-4 w-4 text-primary shrink-0" />
        <div className="flex-1 min-w-0">
          <p className="text-sm font-medium">下一步：映射模块</p>
          <p className="text-xs text-muted-foreground truncate">
            将仓库 "{repositoryName}" 映射到软件模块
          </p>
        </div>
        <Button
          size="sm"
          className="shrink-0"
          onClick={() => {
            onOpenModuleMapping?.()
            document.getElementById('repository-module-mapping')?.scrollIntoView({ behavior: 'smooth', block: 'start' })
          }}
        >
          映射模块
          <ArrowRight className="ml-1 h-3 w-3" />
        </Button>
      </div>
    )
  }

  if (decisionLinks.length > 0) {
    const primaryDecision = decisionLinks[0]

    return (
      <div className="flex items-center gap-3 rounded-lg border bg-primary/5 p-3">
        <GitBranch className="h-4 w-4 text-primary shrink-0" />
        <div className="flex-1 min-w-0">
          <p className="text-sm font-medium">下一步：查看相关决策</p>
          <p className="text-xs text-muted-foreground truncate">
            继续处理与 "{repositoryName}" 相关的决策记录
          </p>
        </div>
        <Button size="sm" asChild className="shrink-0">
          <Link
            to="/decisions/$decisionId"
            params={{ decisionId: primaryDecision.decision_id }}
            search={decisionDetailSearch}
          >
            查看决策
            <ArrowRight className="ml-1 h-3 w-3" />
          </Link>
        </Button>
      </div>
    )
  }

  // 兜底：返回 Dashboard
  return (
    <div className="flex items-center gap-3 rounded-lg border bg-muted/30 p-3">
      <ArrowLeft className="h-4 w-4 text-muted-foreground shrink-0" />
      <div className="flex-1 min-w-0">
        <p className="text-sm text-muted-foreground">仓库结构已完整</p>
      </div>
      <Button size="sm" variant="outline" asChild className="shrink-0">
        <Link to="/dashboard">返回 Dashboard</Link>
      </Button>
    </div>
  )
}
