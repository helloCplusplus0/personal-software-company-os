/**
 * StandardListPage — Standards / List
 *
 * phase14-05 §ADDED-3 四页面组件树冻结（列表页，逐字落地）：
 * - 页壳：标题行 flex-col gap-2 sm:flex-row sm:items-center sm:justify-between
 *   （主标题 text-xl "Standards" + 导语 text-xs text-muted-foreground）+ 新建 CTA（h-9）
 * - 错误区（p-3 text-xs text-destructive + 重试）→ 加载态 → 空态（引导新建）→ 摘要卡列表
 * - 摘要卡：p-3 space-y-2 hover:bg-muted/50（hover 限可交互整卡 Link）；整卡 Link 进详情
 * - 切片纪律：仅消费 query owner，页面零写动作（project_rules §2.5）
 */
import { Link } from '@tanstack/react-router'
import { useStandardListRead } from '../data/use-standard-list-read'
import { Badge } from '@/components/ui/badge'
import { Button } from '@/components/ui/button'
import { Skeleton } from '@/components/ui/skeleton'
import { Plus } from 'lucide-react'

export function StandardListPage() {
  const { data, isLoading, isError, error, refetch } = useStandardListRead()

  return (
    <div className="space-y-4">
      {/* 标题行 — phase14-05 §ADDED-7 移动端基线：移动纵向堆叠，sm+ 单行对齐 */}
      <div className="flex flex-col gap-2 sm:flex-row sm:items-center sm:justify-between">
        <div className="space-y-1">
          <h1 className="text-xl font-bold">Standards</h1>
          <p className="text-xs text-muted-foreground">全局规范资产：结构化目录与跨实体绑定复用</p>
        </div>
        <Button asChild className="w-full sm:w-auto">
          <Link to="/standards/new">
            <Plus className="mr-2 h-4 w-4" />
            新建 Standard
          </Link>
        </Button>
      </div>

      {/* 错误区 — 对齐既有列表页紧凑错误基线 */}
      {isError ? (
        <div className="rounded-lg border border-destructive/50 bg-destructive/10 p-3">
          <p className="text-xs text-destructive">列表读取失败：{(error as Error).message}</p>
          <Button variant="outline" size="sm" className="mt-2 h-7" onClick={() => refetch()}>
            重试
          </Button>
        </div>
      ) : isLoading ? (
        <div className="space-y-2">
          <Skeleton className="h-16 w-full" />
          <Skeleton className="h-16 w-full" />
          <Skeleton className="h-16 w-full" />
        </div>
      ) : (data?.length ?? 0) === 0 ? (
        // 空态 — 引导新建（对齐 Dashboard 紧凑空态基线）
        <div className="rounded-lg border border-dashed p-4 text-center">
          <p className="mb-3 text-xs text-muted-foreground">暂无登记的全局规范资产</p>
          <Button asChild size="sm">
            <Link to="/standards/new">
              <Plus className="mr-2 h-4 w-4" />
              登记首个 Standard
            </Link>
          </Button>
        </div>
      ) : (
        // 摘要卡列表 — 整卡 Link 可交互，hover 仅限卡本身
        <div className="divide-y rounded-lg border">
          {(data ?? []).map((standard) => (
            <Link
              key={standard.id}
              to="/standards/$standardId"
              params={{ standardId: standard.id }}
              className="block space-y-2 p-3 hover:bg-muted/50"
            >
              <div className="flex min-w-0 items-center gap-2">
                <span className="min-w-0 truncate text-sm font-medium">{standard.name}</span>
                <Badge variant={standard.status === 'active' ? 'default' : 'secondary'}>
                  {standard.status}
                </Badge>
              </div>
              {standard.description ? (
                <p className="min-w-0 truncate text-xs text-muted-foreground">
                  {standard.description}
                </p>
              ) : null}
              <p className="text-xs text-muted-foreground">
                更新于 {new Date(standard.updated_at).toLocaleString()}
              </p>
            </Link>
          ))}
        </div>
      )}
    </div>
  )
}
