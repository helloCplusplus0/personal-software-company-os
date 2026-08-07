import { Link } from '@tanstack/react-router'
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import { Badge } from '@/components/ui/badge'
import { Button } from '@/components/ui/button'
import { ArrowRight } from 'lucide-react'
import type { BoundRepositorySummary, Product } from '../types'
import type { ProductSourceContext } from '@/features/repository-binding/utils/product-source-transit'
import { buildProductSourceTransit } from '@/features/repository-binding/utils/product-source-transit'

interface ProductBoundRepositoryListSectionProps {
  product: Product
  boundRepositories: BoundRepositorySummary[]
  /**
   * phase04-06 Product Detail 来源上下文
   * 用于从 Product Detail 进入 Repository Binding 时继续携带来源标记，
   * 使 Repository Binding 返回时能恢复 Product Detail 的来源，不退化为 direct-entry
   */
  productSource: ProductSourceContext
}

/**
 * ProductBoundRepositoryListSection — 已绑定仓库列表区 + Repository Binding 上下文入口
 *
 * phase04-05 组件树冻结：
 * - ProductBoundRepositoryListSection 承接已绑定 Repository 列表展示（只读摘要）
 * - ProductRepositoryBindingEntry 承接进入 Repository Binding Detail / Workspace 的上下文跳转入口
 *
 * phase04-05 上下文入口路由参数冻结：
 * - 进入 Repository Binding 时必须携带 productId / productName / fromProductDetail 作为上下文搜索参数
 * - 若目标 Repository 尚未确定，先导航到 /repositories 并携带上下文搜索参数
 *
 * phase04-06 来源上下文透传：
 * - 从本组件进入 Repository Binding 时，必须继续携带 Product Detail 自己的来源上下文（使用 product 前缀区分）
 * - 这样 Repository Binding 返回到 Product Detail 时，后者才能按真实来源返回，不退化为 direct-entry
 *
 * 布局降级（phase04-05）：
 * - PC / 移动：分区式展示，已绑定列表与入口按钮同区可见
 */
export function ProductBoundRepositoryListSection({
  product,
  boundRepositories,
  productSource,
}: ProductBoundRepositoryListSectionProps) {
  // phase04-06 构造进入 Repository Binding 的完整搜索参数：
  // - 自身来源标记 fromProductDetail + productId / productName
  // - Product Detail 来源上下文透传（product 前缀）
  const repositorySearch: Record<string, unknown> = {
    fromProductDetail: true,
    productId: product.id,
    productName: product.name,
    ...buildProductSourceTransit(productSource),
  }

  return (
    <Card>
      <CardHeader>
        <CardTitle className="flex items-center justify-between gap-2">
          <span>已绑定仓库</span>
          {/* phase04-05 ProductRepositoryBindingEntry — 上下文跳转入口 */}
          <Button variant="outline" size="sm" asChild>
            <Link to="/repositories" search={repositorySearch}>
              <ArrowRight className="mr-2 h-4 w-4" />
              进入仓库绑定
            </Link>
          </Button>
        </CardTitle>
      </CardHeader>
      <CardContent className="space-y-1">
        {boundRepositories.length === 0 ? (
          <p className="text-xs text-muted-foreground">未绑定仓库</p>
        ) : (
          boundRepositories.map((br) => (
            <div key={br.repository_id} className="flex items-center gap-2">
              <Badge variant="outline">{br.repository_name}</Badge>
              <span className="text-xs text-muted-foreground">{br.provider}</span>
              <Badge variant={br.repository_status === 'active' ? 'default' : 'secondary'}>
                {br.repository_status}
              </Badge>
            </div>
          ))
        )}
      </CardContent>
    </Card>
  )
}
