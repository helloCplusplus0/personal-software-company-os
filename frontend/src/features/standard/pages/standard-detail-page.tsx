/**
 * StandardDetailPage — Standard Detail
 *
 * phase14-05 §ADDED-3 四页面组件树冻结（详情页，逐字落地）：
 * - 头部：name（text-xl）+ status Badge + description + updated_at + 操作组
 *   （编辑 Link /standards/:id/edit（h-9）；删除 — ACTIVE 态禁用并提示先 Retire，
 *   非 ACTIVE 经 window.confirm 后调 use-delete-standard，成功 → /standards）
 * - StandardTreeView（常规模式，区块标题"目录结构"）→ StandardBindingPanel（"绑定管理"）
 *   → StandardRevisionList（border-t pt-2，区块标题"Revision 历史"）
 * - 页面壳无 Card 重型嵌套；子区域 border-t pt-2 分隔；区块标题统一 text-xs 小节风格
 * - 切片纪律：删除仅经 application owner，页面不内联 mutation hook（project_rules §2.5）
 */
import { Link, useNavigate } from '@tanstack/react-router'
import { useStandardDetailRead } from '../data/use-standard-detail-read'
import { useStandardRevisionsRead } from '../data/use-standard-revisions-read'
import { useDeleteStandard } from '../application/use-delete-standard'
import { StandardTreeView } from '../components/standard-tree-view'
import { StandardBindingPanel } from '../components/standard-binding-panel'
import { StandardRevisionList } from '../components/standard-revision-list'
import { Badge } from '@/components/ui/badge'
import { Button } from '@/components/ui/button'
import { Skeleton } from '@/components/ui/skeleton'

interface StandardDetailPageProps {
  /** 路由文件取参传入（/standards/:standardId） */
  standardId: string
}

export function StandardDetailPage({ standardId }: StandardDetailPageProps) {
  const navigate = useNavigate()
  const detailQuery = useStandardDetailRead(standardId)
  const revisionsQuery = useStandardRevisionsRead(standardId)
  const deleteMutation = useDeleteStandard()

  /** 删除 — 非 ACTIVE 态确认后经 owner 调用；成功回流列表页 */
  const handleDelete = () => {
    if (!window.confirm('确认删除该 Standard？绑定与 revision 历史将连带删除。')) return
    deleteMutation.mutate(
      { standardId },
      {
        onSuccess: () => {
          navigate({ to: '/standards' })
        },
      },
    )
  }

  // 错误态 — 含 not found（owner 将缺主体抛为 'standard not found'）
  if (detailQuery.isError) {
    return (
      <div className="space-y-4">
        <div className="rounded-lg border border-destructive/50 bg-destructive/10 p-3">
          <p className="text-xs text-destructive">
            {(detailQuery.error as Error).message === 'standard not found'
              ? '该 Standard 不存在或已被删除'
              : `详情读取失败：${(detailQuery.error as Error).message}`}
          </p>
        </div>
        <Button variant="outline" size="sm" asChild>
          <Link to="/standards">返回列表</Link>
        </Button>
      </div>
    )
  }

  // 加载态
  if (detailQuery.isLoading || !detailQuery.data) {
    return (
      <div className="space-y-4">
        <Skeleton className="h-24 w-full" />
        <Skeleton className="h-48 w-full" />
        <Skeleton className="h-32 w-full" />
      </div>
    )
  }

  const { standard, bindings } = detailQuery.data
  // ACTIVE 态删除禁用（phase14-05：提示先 Retire / 先经编辑改为 draft/retired 再删除）
  const isActive = standard.status === 'active'

  return (
    <div className="space-y-4">
      {/* 头部：name + status + description + updated_at + 操作组 */}
      <div className="flex flex-col gap-2 sm:flex-row sm:items-start sm:justify-between">
        <div className="min-w-0 space-y-1">
          <div className="flex min-w-0 items-center gap-2">
            <h1 className="min-w-0 truncate text-xl font-bold">{standard.name}</h1>
            <Badge variant={standard.status === 'active' ? 'default' : 'secondary'}>
              {standard.status}
            </Badge>
          </div>
          {standard.description ? (
            <p className="text-xs text-muted-foreground">{standard.description}</p>
          ) : null}
          <p className="text-xs text-muted-foreground">
            更新于 {new Date(standard.updated_at).toLocaleString()}
          </p>
        </div>

        {/* 操作组 — 移动端纵向铺满，sm+ 单行；编辑 h-9 / 删除 h-9 */}
        <div className="flex w-full flex-col gap-2 sm:w-auto sm:flex-row sm:items-center">
          <Button asChild className="w-full sm:w-auto">
            <Link
              to="/standards/$standardId/edit"
              params={{ standardId }}
            >
              编辑
            </Link>
          </Button>
          <Button
            variant="destructive"
            className="w-full sm:w-auto"
            disabled={isActive}
            title={isActive ? '先经编辑改为 draft/retired 再删除' : undefined}
            onClick={handleDelete}
          >
            删除
          </Button>
        </div>
      </div>

      {/* ACTIVE 态删除禁用文案提示（phase14-05 表述：提示先 Retire） */}
      {isActive ? (
        <p className="text-xs text-muted-foreground">
          active 状态不可直接删除，先经编辑改为 draft / retired（Retire）后再删除
        </p>
      ) : null}

      {/* 删除失败行内回显 */}
      {deleteMutation.isError ? (
        <p className="text-xs text-destructive">删除失败：{(deleteMutation.error as Error).message}</p>
      ) : null}

      {/* 目录结构 — StandardTreeView 常规模式 */}
      <section className="space-y-2">
        <p className="text-xs font-medium text-muted-foreground">目录结构</p>
        <StandardTreeView tree={standard.directory_tree} />
      </section>

      {/* 绑定管理 — StandardBindingPanel（组件内自带区块标题；全站唯一绑定发起位） */}
      <section className="border-t pt-2">
        <StandardBindingPanel standardId={standardId} bindings={bindings} />
      </section>

      {/* Revision 历史 — border-t pt-2 分隔 */}
      <section className="space-y-2 border-t pt-2">
        <p className="text-xs font-medium text-muted-foreground">Revision 历史</p>
        {revisionsQuery.isLoading ? (
          <p className="text-xs text-muted-foreground">加载中...</p>
        ) : revisionsQuery.isError ? (
          <p className="text-xs text-destructive">
            revision 读取失败：{(revisionsQuery.error as Error).message}
          </p>
        ) : (
          <StandardRevisionList revisions={revisionsQuery.data ?? []} />
        )}
      </section>
    </div>
  )
}
