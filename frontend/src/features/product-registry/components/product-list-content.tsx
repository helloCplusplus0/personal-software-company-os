import { Link, useNavigate } from '@tanstack/react-router'
import { Badge } from '@/components/ui/badge'
import { Skeleton } from '@/components/ui/skeleton'
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from '@/components/ui/table'
import type { ProductListItem } from '../types'

/** 点击列表项导航到详情页时携带的搜索参数 */
export type ProductDetailSearch = Record<string, unknown>

interface ProductListContentProps {
  items: ProductListItem[]
  isLoading: boolean
  /** phase04-05 上下文入口：从 Module Detail 带入的上下文参数，点击列表项时继续传递到详情页 */
  detailSearch: ProductDetailSearch
}

/**
 * ProductListContent — 列表内容区
 * phase04-05 组件树冻结：只承接产品列表读取结果展示
 *
 * phase04-05 上下文入口承接：
 * - 从 Module Detail 带上下文进入列表页后，
 *   点击列表项时必须继续携带上下文搜索参数到详情页
 * - 无上下文时点击列表项携带 fromList: true
 *
 * 布局降级（phase04-05）：
 * - PC：高信息密度表格，展示 name / description / status / created_at / module_bind_count / repository_bind_count
 * - 移动：单列卡片重排，保留 name / status / module_bind_count / repository_bind_count 核心展示
 */
export function ProductListContent({ items, isLoading, detailSearch }: ProductListContentProps) {
  const navigate = useNavigate()

  if (isLoading) {
    return (
      <div className="space-y-2">
        {Array.from({ length: 3 }).map((_, i) => (
          <Skeleton key={i} className="h-12 w-full" />
        ))}
      </div>
    )
  }

  return (
    <>
      {/* PC：高信息密度表格 — phase04-05 */}
      <div className="hidden md:block">
        <Table>
          <TableHeader>
            <TableRow>
              <TableHead>名称</TableHead>
              <TableHead>描述</TableHead>
              <TableHead>状态</TableHead>
              <TableHead>创建时间</TableHead>
              <TableHead className="text-center">模块绑定</TableHead>
              <TableHead className="text-center">仓库绑定</TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            {items.map((item) => (
              <TableRow
                key={item.id}
                className="cursor-pointer hover:bg-muted/50"
                onClick={() =>
                  navigate({
                    to: '/products/$productId',
                    params: { productId: item.id },
                    search: detailSearch,
                  })
                }
              >
                <TableCell className="font-medium">{item.name}</TableCell>
                <TableCell className="max-w-xs truncate text-muted-foreground">
                  {item.description}
                </TableCell>
                <TableCell>
                  <StatusBadge status={item.status} />
                </TableCell>
                <TableCell className="text-muted-foreground">
                  {new Date(item.created_at).toLocaleDateString()}
                </TableCell>
                <TableCell className="text-center">{item.module_bind_count}</TableCell>
                <TableCell className="text-center">{item.repository_bind_count}</TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </div>

      {/* 移动：单列卡片重排 — phase04-05 保留核心展示（对齐 Dashboard 紧凑密度基线） */}
      <div className="space-y-2 md:hidden">
        {items.map((item) => (
          <Link
            key={item.id}
            to="/products/$productId"
            params={{ productId: item.id }}
            search={detailSearch}
            className="block rounded-lg border p-3 transition-colors hover:bg-muted/50"
          >
            <div className="flex items-center justify-between">
              <span className="font-medium">{item.name}</span>
              <StatusBadge status={item.status} />
            </div>
            <div className="mt-2 flex items-center gap-4 text-xs text-muted-foreground">
              <span>模块：{item.module_bind_count}</span>
              <span>仓库：{item.repository_bind_count}</span>
            </div>
          </Link>
        ))}
      </div>
    </>
  )
}

function StatusBadge({ status }: { status: 'active' | 'archived' }) {
  return (
    <Badge variant={status === 'active' ? 'default' : 'secondary'}>
      {status}
    </Badge>
  )
}
