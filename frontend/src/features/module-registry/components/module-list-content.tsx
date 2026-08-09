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
import type { ModuleListItem } from '../types'

interface ModuleListContentProps {
  items: ModuleListItem[]
  isLoading: boolean
  detailSearch: Record<string, unknown>
}

/**
 * ModuleListContent — 列表内容区
 * §8.5 布局降级：
 * - PC：高信息密度表格
 * - 移动：卡片重排（通过 hidden/block 响应式切换）
 */
export function ModuleListContent({ items, isLoading, detailSearch }: ModuleListContentProps) {
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
      {/* PC：高信息密度表格 — §8.5 */}
      <div className="hidden md:block">
        <Table>
          <TableHeader>
            <TableRow>
              <TableHead>名称</TableHead>
              <TableHead>描述</TableHead>
              <TableHead>状态</TableHead>
              <TableHead>最新版本</TableHead>
              <TableHead className="text-center">产品绑定</TableHead>
              <TableHead className="text-center">仓库映射</TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            {items.map((item) => (
              <TableRow
                key={item.id}
                className="cursor-pointer hover:bg-muted/50"
                onClick={() =>
                  navigate({
                    to: '/modules/$moduleId',
                    params: { moduleId: item.id },
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
                  {item.latest_release ?? '—'}
                </TableCell>
                <TableCell className="text-center">{item.product_bind_count}</TableCell>
                <TableCell className="text-center">{item.repository_bind_count}</TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </div>

      {/* 移动：单列卡片重排 — §8.5 */}
      <div className="space-y-3 md:hidden">
        {items.map((item) => (
          <Link
            key={item.id}
            to="/modules/$moduleId"
            params={{ moduleId: item.id }}
              search={detailSearch}
            className="block rounded-lg border p-4 hover:bg-accent transition-colors"
          >
            <div className="flex items-center justify-between">
              <span className="font-medium">{item.name}</span>
              <StatusBadge status={item.status} />
            </div>
            <p className="mt-1 text-sm text-muted-foreground line-clamp-2">{item.description}</p>
            <div className="mt-2 flex items-center gap-4 text-xs text-muted-foreground">
              <span>版本：{item.latest_release ?? '—'}</span>
              <span>产品：{item.product_bind_count}</span>
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
