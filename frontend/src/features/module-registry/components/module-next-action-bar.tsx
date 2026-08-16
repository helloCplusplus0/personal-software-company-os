import { Button } from '@/components/ui/button'
import { ArrowRight, Boxes, GitBranch, ArrowLeft } from 'lucide-react'
import { Link } from '@tanstack/react-router'
import { MODULE_SEMANTIC_LABEL, REPOSITORY_SEMANTIC_LABEL } from '@/features/project-context/data/shared-semantic-constants'

interface ModuleNextActionBarProps {
  /** 是否有已绑定 Product */
  hasProductBinding: boolean
  /** 是否有已映射 Repository */
  hasRepositoryMapping: boolean
  /** 当前 Module 标识 */
  moduleId: string
  moduleName: string
  decisionLinks?: { decision_id: string; decision_title: string }[]
  decisionDetailSearch?: Record<string, unknown>
}

/**
 * ModuleNextActionBar — Module Detail 的页面级下一步动作区
 *
 * phase10-10 §"Module Detail 下一步动作承接矩阵"：
 * - Product 绑定 > Repository 映射 > 返回 Dashboard
 * - 主 CTA 指向 canonical binding owner（Product Detail / Repository Detail）
 */
export function ModuleNextActionBar({
  hasProductBinding,
  hasRepositoryMapping,
  moduleId,
  moduleName,
  decisionLinks = [],
  decisionDetailSearch,
}: ModuleNextActionBarProps) {
  // 优先级 1：Product 绑定缺口
  if (!hasProductBinding) {
    return (
      <div className="flex items-center gap-3 rounded-lg border bg-primary/5 p-3">
        <Boxes className="h-4 w-4 text-primary shrink-0" />
        <div className="flex-1 min-w-0">
          <p className="text-sm font-medium">下一步：绑定产品</p>
          <p className="text-xs text-muted-foreground truncate">
            将{MODULE_SEMANTIC_LABEL} "{moduleName}" 绑定到所属经营目标（Product），建立资产归属关系
          </p>
        </div>
        <Button size="sm" asChild className="shrink-0">
          <Link
            to="/products"
            search={{ fromModuleDetail: true, moduleId, moduleName } as any}
          >
            去绑定产品
            <ArrowRight className="ml-1 h-3 w-3" />
          </Link>
        </Button>
      </div>
    )
  }

  // 优先级 2：Repository 映射缺口
  if (!hasRepositoryMapping) {
    return (
      <div className="flex items-center gap-3 rounded-lg border bg-primary/5 p-3">
        <GitBranch className="h-4 w-4 text-primary shrink-0" />
        <div className="flex-1 min-w-0">
          <p className="text-sm font-medium">下一步：映射仓库</p>
          <p className="text-xs text-muted-foreground truncate">
            将{MODULE_SEMANTIC_LABEL} "{moduleName}" 映射到{REPOSITORY_SEMANTIC_LABEL}（Repository），建立可追溯关联
          </p>
        </div>
        <Button size="sm" asChild className="shrink-0">
          <Link
            to="/repositories"
            search={{ fromModuleDetail: true, moduleId, moduleName } as any}
          >
            去映射仓库
            <ArrowRight className="ml-1 h-3 w-3" />
          </Link>
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
            继续处理与该{MODULE_SEMANTIC_LABEL}相关的规则与决策记录
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
        <p className="text-sm text-muted-foreground">该{MODULE_SEMANTIC_LABEL}的关系已完整</p>
      </div>
      <Button size="sm" variant="outline" asChild className="shrink-0">
        <Link to="/dashboard">返回 Dashboard</Link>
      </Button>
    </div>
  )
}
