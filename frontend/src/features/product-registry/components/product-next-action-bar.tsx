import { Button } from '@/components/ui/button'
import { ArrowRight, Package, Database, ArrowLeft } from 'lucide-react'
import { Link } from '@tanstack/react-router'

interface ProductNextActionBarProps {
  /** 是否有已绑定 Repository */
  hasRepository: boolean
  /** 是否有已绑定 Module */
  hasModule: boolean
  /** 当前 Product 标识 */
  productId: string
  productName: string
  decisionLinks?: { decision_id: string; decision_title: string }[]
  decisionDetailSearch?: Record<string, unknown>
}

/**
 * ProductNextActionBar — Product Detail 的页面级下一步动作区
 *
 * phase10-10 §"Product Detail 下一步动作承接矩阵"：
 * - Repository 缺口 > Module 缺口 > 返回 Dashboard
 * - 主 CTA 以主按钮样式展示，指向 canonical path
 */
export function ProductNextActionBar({
  hasRepository,
  hasModule,
  productId,
  productName,
  decisionLinks = [],
  decisionDetailSearch,
}: ProductNextActionBarProps) {
  // 优先级 1：Repository 缺口
  if (!hasRepository) {
    return (
      <div className="flex items-center gap-3 rounded-lg border bg-primary/5 p-3">
        <Database className="h-4 w-4 text-primary shrink-0" />
        <div className="flex-1 min-w-0">
          <p className="text-sm font-medium">下一步：绑定仓库</p>
          <p className="text-xs text-muted-foreground truncate">
            为 "{productName}" 绑定代码仓库，建立可追溯的资产关联
          </p>
        </div>
        <Button size="sm" asChild className="shrink-0">
          <Link
            to="/repositories"
            search={{ fromProductDetail: true, productId, productName } as any}
          >
            绑定仓库
            <ArrowRight className="ml-1 h-3 w-3" />
          </Link>
        </Button>
      </div>
    )
  }

  // 优先级 2：Module 缺口
  if (!hasModule) {
    return (
      <div className="flex items-center gap-3 rounded-lg border bg-primary/5 p-3">
        <Package className="h-4 w-4 text-primary shrink-0" />
        <div className="flex-1 min-w-0">
          <p className="text-sm font-medium">下一步：绑定模块</p>
          <p className="text-xs text-muted-foreground truncate">
            为 "{productName}" 绑定软件模块，完善产品结构
          </p>
        </div>
        <Button
          size="sm"
          className="shrink-0"
          onClick={() => {
            document.getElementById('product-module-binding')?.scrollIntoView({ behavior: 'smooth' })
          }}
        >
          绑定模块
          <ArrowRight className="ml-1 h-3 w-3" />
        </Button>
      </div>
    )
  }

  if (decisionLinks.length > 0) {
    const primaryDecision = decisionLinks[0]

    return (
      <div className="flex items-center gap-3 rounded-lg border bg-primary/5 p-3">
        <Database className="h-4 w-4 text-primary shrink-0" />
        <div className="flex-1 min-w-0">
          <p className="text-sm font-medium">下一步：查看相关决策</p>
          <p className="text-xs text-muted-foreground truncate">
            继续处理与 "{productName}" 相关的决策记录
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
        <p className="text-sm text-muted-foreground">产品结构已完整</p>
      </div>
      <Button size="sm" variant="outline" asChild className="shrink-0">
        <Link to="/dashboard">返回 Dashboard</Link>
      </Button>
    </div>
  )
}
