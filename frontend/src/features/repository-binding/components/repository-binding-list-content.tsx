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
import type { RepositoryListItem } from '../types'

/** 点击列表项导航到详情页时携带的搜索参数 */
export type RepositoryDetailSearch = Record<string, unknown>

interface RepositoryBindingListContentProps {
  items: RepositoryListItem[]
  isLoading: boolean
  /** phase04-05 上下文入口：从 Product Detail / Module Detail 带入的上下文参数，点击列表项时继续传递到详情页 */
  detailSearch: RepositoryDetailSearch
}

/**
 * RepositoryBindingListContent — 列表内容区
 * phase04-05 组件树冻结：只承接仓库列表读取结果展示
 *
 * phase04-05 上下文入口承接：
 * - 从 Product Detail / Module Detail 带上下文进入列表页后，
 *   点击列表项时必须继续携带上下文搜索参数到详情页
 * - 无上下文时点击列表项携带 fromList: true
 *
 * 布局降级（phase04-05）：
 * - PC：高信息密度表格，展示 name / url / provider / status / created_at / product_bind_count / module_bind_count
 * - 移动：单列卡片重排，保留 name / provider / status / product_bind_count / module_bind_count 核心展示
 */
export function RepositoryBindingListContent({ items, isLoading, detailSearch }: RepositoryBindingListContentProps) {
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
              <TableHead>URL</TableHead>
              <TableHead>提供商</TableHead>
              <TableHead>状态</TableHead>
              <TableHead>创建时间</TableHead>
              <TableHead className="text-center">产品绑定</TableHead>
              <TableHead className="text-center">模块映射</TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            {items.map((item) => (
              <TableRow
                key={item.id}
                className="cursor-pointer hover:bg-muted/50"
                onClick={() =>
                  navigate({
                    to: '/repositories/$repositoryId',
                    params: { repositoryId: item.id },
                    search: detailSearch,
                  })
                }
              >
                <TableCell className="font-medium">{item.name}</TableCell>
                <TableCell className="max-w-xs truncate text-muted-foreground">
                  {item.url}
                </TableCell>
                <TableCell className="text-muted-foreground">{item.provider}</TableCell>
                <TableCell>
                  <StatusBadge status={item.status} />
                </TableCell>
                <TableCell className="text-muted-foreground">
                  {new Date(item.created_at).toLocaleDateString()}
                </TableCell>
                <TableCell className="text-center">{item.product_bind_count}</TableCell>
                <TableCell className="text-center">{item.module_bind_count}</TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </div>

      {/* 移动：单列卡片重排 — phase04-05 保留核心展示 */}
      <div className="space-y-3 md:hidden">
        {items.map((item) => (
          <Link
            key={item.id}
            to="/repositories/$repositoryId"
            params={{ repositoryId: item.id }}
            search={detailSearch}
            className="block rounded-lg border p-4 hover:bg-accent transition-colors"
          >
            <div className="flex items-center justify-between">
              <span className="font-medium">{item.name}</span>
              <StatusBadge status={item.status} />
            </div>
            <p className="mt-1 text-xs text-muted-foreground">{item.provider}</p>
            <div className="mt-2 flex items-center gap-4 text-xs text-muted-foreground">
              <span>产品：{item.product_bind_count}</span>
              <span>模块：{item.module_bind_count}</span>
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
