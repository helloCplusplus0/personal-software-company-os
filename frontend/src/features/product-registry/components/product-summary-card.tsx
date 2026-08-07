import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import { Badge } from '@/components/ui/badge'
import type { Product } from '../types'

interface ProductSummaryCardProps {
  product: Product
}

/**
 * ProductSummaryCard — 产品摘要卡片
 * phase04-05 组件树冻结：只承接 Product 核心字段 id / name / description / status / created_at
 * phase04-05 默认归属于 ProductDetailPage
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
        <div>
          <p className="text-sm text-muted-foreground">描述</p>
          <p className="text-sm">{product.description}</p>
        </div>
        <div>
          <p className="text-sm text-muted-foreground">创建时间</p>
          <p className="text-sm">{new Date(product.created_at).toLocaleDateString()}</p>
        </div>
      </CardContent>
    </Card>
  )
}
