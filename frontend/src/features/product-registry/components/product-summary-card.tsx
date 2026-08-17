import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import { Badge } from '@/components/ui/badge'
import type { Product } from '../types'
import { PRODUCT_SEMANTIC_LABEL } from '@/features/project-context/data/shared-semantic-constants'

interface ProductSummaryCardProps {
  product: Product
}

/**
 * ProductSummaryCard — 产品摘要卡片
 * phase04-05 组件树冻结：只承接 Product 核心字段 id / name / description / status / created_at
 * phase04-05 默认归属于 ProductDetailPage
 * phase12-08：增加"经营目标与交付容器"语义标签
 */
export function ProductSummaryCard({ product }: ProductSummaryCardProps) {
  return (
    <Card>
      <CardHeader>
        <CardTitle className="flex items-center justify-between gap-2">
          <span className="truncate">{product.name}</span>
          <Badge variant={product.status === 'active' ? 'default' : 'secondary'}>
            {product.status}
          </Badge>
        </CardTitle>
      </CardHeader>
      <CardContent className="space-y-2">
        <div className="flex items-center gap-2">
          <span className="text-[10px] font-medium text-muted-foreground uppercase tracking-wide">
            {PRODUCT_SEMANTIC_LABEL}
          </span>
        </div>
        <div>
          <p className="text-xs text-muted-foreground">描述</p>
          <p className="text-sm">{product.description}</p>
        </div>
        <div>
          <p className="text-xs text-muted-foreground">创建时间</p>
          <p className="text-sm">{new Date(product.created_at).toLocaleDateString()}</p>
        </div>
      </CardContent>
    </Card>
  )
}
