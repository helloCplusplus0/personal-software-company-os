/**
 * DecisionListContent — 列表内容展示
 *
 * §9.1 / §9.2 承接列表读取结果展示。
 *
 * 布局降级（§9.2 / phase03-05 §"布局降级策略"）：
 * - PC：高信息密度表格布局
 * - 移动浏览器：单列卡片重排
 *
 * 使用 Tailwind 响应式断点实现双场景布局，不引入第二套移动端 UI 架构。
 */
import { Link } from '@tanstack/react-router'
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
import type { DecisionListItem, DecisionStatus } from '../types'

interface DecisionListContentProps {
  items: DecisionListItem[]
  isLoading: boolean
}

/** 状态对应的 Badge variant */
const STATUS_VARIANT: Record<DecisionStatus, 'default' | 'secondary' | 'outline' | 'destructive'> = {
  proposed: 'secondary',
  active: 'default',
  superseded: 'outline',
  archived: 'outline',
}

/** 状态中文标签 */
const STATUS_LABEL: Record<DecisionStatus, string> = {
  proposed: 'Proposed',
  active: 'Active',
  superseded: 'Superseded',
  archived: 'Archived',
}

export function DecisionListContent({ items, isLoading }: DecisionListContentProps) {
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
      {/* PC：表格布局（md 以上） */}
      <div className="hidden rounded-md border md:block">
        <Table>
          <TableHeader>
            <TableRow>
              <TableHead>标题</TableHead>
              <TableHead className="w-28">状态</TableHead>
              <TableHead className="w-32">已关联模块</TableHead>
              <TableHead className="w-40">创建时间</TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            {items.map((item) => (
              <TableRow key={item.id}>
                <TableCell className="font-medium">
                  <Link
                    to="/decisions/$decisionId"
                    params={{ decisionId: item.id }}
                    search={{ fromList: true }}
                    className="text-primary hover:underline"
                  >
                    {item.title}
                  </Link>
                </TableCell>
                <TableCell>
                  <Badge variant={STATUS_VARIANT[item.status]}>
                    {STATUS_LABEL[item.status]}
                  </Badge>
                </TableCell>
                <TableCell>
                  {item.link_count > 0 ? (
                    <span className="text-sm text-muted-foreground">
                      {item.linked_module_summary || `${item.link_count} 个模块`}
                    </span>
                  ) : (
                    <span className="text-sm text-muted-foreground">—</span>
                  )}
                </TableCell>
                <TableCell className="text-sm text-muted-foreground">
                  {new Date(item.created_at).toLocaleDateString('zh-CN')}
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </div>

      {/* 移动浏览器：单列卡片布局（md 以下） */}
      <div className="space-y-3 md:hidden">
        {items.map((item) => (
          <Link
            key={item.id}
            to="/decisions/$decisionId"
            params={{ decisionId: item.id }}
            search={{ fromList: true }}
            className="block rounded-md border p-3 hover:bg-accent transition-colors"
          >
            <div className="flex items-center justify-between gap-2">
              <span className="font-medium">{item.title}</span>
              <Badge variant={STATUS_VARIANT[item.status]}>
                {STATUS_LABEL[item.status]}
              </Badge>
            </div>
            <div className="mt-1 text-sm text-muted-foreground">
              {item.link_count > 0
                ? item.linked_module_summary || `${item.link_count} 个模块`
                : '暂无关联模块'}
            </div>
            <div className="mt-1 text-xs text-muted-foreground">
              {new Date(item.created_at).toLocaleDateString('zh-CN')}
            </div>
          </Link>
        ))}
      </div>
    </>
  )
}
