/**
 * StandardDetailPage — Standard Detail
 *
 * phase14-05 §ADDED-3 四页面组件树冻结（详情页）+ phase14-10 前置 UI 一致性反馈对齐：
 * - 返回行：ghost sm「返回列表」（对齐 module / product / repository / decision 详情页基线；
 *   Standard 无 dashboard / onboarding 来源链，仅承接返回列表）
 * - 语义导语：text-xs text-muted-foreground（对齐 phase12-08 页面级语义导语基线）
 * - 分区布局：grid lg:grid-cols-3（对齐既有详情页"左摘要 1 列 + 右内容 2 列"模式）
 *   - 左列：摘要 Card（name + status Badge + 语义标签 + description + updated_at + 操作组）
 *   - 右列：目录结构 → 绑定管理（border-t pt-2）→ Revision 历史（border-t pt-2）
 * - 操作组：编辑 Link /standards/:id/edit；删除 — ACTIVE 态禁用并提示先 Retire，
 *   非 ACTIVE 经 window.confirm 后调 use-delete-standard，成功 → /standards
 * - 右列内容区无 Card 重型嵌套；子区域 border-t pt-2 分隔；区块标题统一 text-xs 小节风格
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
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import { Skeleton } from '@/components/ui/skeleton'
import { ArrowLeft } from 'lucide-react'

interface StandardDetailPageProps {
  /** 路由文件取参传入（/standards/:standardId） */
  standardId: string
}

export function StandardDetailPage({ standardId }: StandardDetailPageProps) {
  const navigate = useNavigate()
  const detailQuery = useStandardDetailRead(standardId)
  const revisionsQuery = useStandardRevisionsRead(standardId)
  const deleteMutation = useDeleteStandard()

  /** 返回列表 — Standard 无跨页来源链，固定回 /standards */
  const handleReturn = () => navigate({ to: '/standards' })

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
        <Button variant="ghost" size="sm" onClick={handleReturn}>
          <ArrowLeft className="mr-2 h-4 w-4" />
          返回列表
        </Button>
        <div className="rounded-lg border border-destructive/50 bg-destructive/10 p-3">
          <p className="text-xs text-destructive">
            {(detailQuery.error as Error).message === 'standard not found'
              ? '该 Standard 不存在或已被删除'
              : `详情读取失败：${(detailQuery.error as Error).message}`}
          </p>
        </div>
      </div>
    )
  }

  // 加载态 — 骨架对齐分区布局（左摘要 + 右内容）
  if (detailQuery.isLoading || !detailQuery.data) {
    return (
      <div className="space-y-4">
        <Button variant="ghost" size="sm" onClick={handleReturn}>
          <ArrowLeft className="mr-2 h-4 w-4" />
          返回列表
        </Button>
        <div className="grid gap-4 lg:grid-cols-3">
          <div className="space-y-2 lg:col-span-1">
            <Skeleton className="h-40 w-full" />
          </div>
          <div className="space-y-2 lg:col-span-2">
            <Skeleton className="h-48 w-full" />
            <Skeleton className="h-32 w-full" />
          </div>
        </div>
      </div>
    )
  }

  const { standard, bindings } = detailQuery.data
  // ACTIVE 态删除禁用（phase14-05：提示先 Retire / 先经编辑改为 draft/retired 再删除）
  const isActive = standard.status === 'active'

  return (
    <div className="space-y-4">
      {/* 返回行 — 对齐既有详情页基线（ghost sm + ArrowLeft） */}
      <Button variant="ghost" size="sm" onClick={handleReturn}>
        <ArrowLeft className="mr-2 h-4 w-4" />
        返回列表
      </Button>

      {/* 语义导语 — 对齐 phase12-08 页面级语义导语基线 */}
      <p className="text-xs text-muted-foreground">
        该 Standard 代表一条全局规范资产：以结构化目录树登记关键全局文档，并跨实体绑定复用
      </p>

      {/* PC：左摘要 1 列 + 右内容 2 列；移动端按垂直顺序重排（对齐既有详情页分区布局） */}
      <div className="grid gap-4 lg:grid-cols-3">
        {/* 摘要主区 — 摘要 Card（对齐 ModuleSummaryCard 卡式基线） */}
        <div className="space-y-2 lg:col-span-1">
          <Card>
            <CardHeader>
              <CardTitle className="flex items-center justify-between gap-2">
                <span className="min-w-0 truncate">{standard.name}</span>
                <Badge variant={standard.status === 'active' ? 'default' : 'secondary'}>
                  {standard.status}
                </Badge>
              </CardTitle>
            </CardHeader>
            <CardContent className="space-y-2">
              <div className="flex items-center gap-2">
                <span className="text-[10px] font-medium uppercase tracking-wide text-muted-foreground">
                  全局规范资产
                </span>
              </div>
              <div>
                <p className="text-xs text-muted-foreground">描述</p>
                <p className="text-sm">{standard.description || '—'}</p>
              </div>
              <div>
                <p className="text-xs text-muted-foreground">更新时间</p>
                <p className="text-sm">{new Date(standard.updated_at).toLocaleString()}</p>
              </div>

              {/* 操作组 — 卡内纵向铺满（编辑 / 删除，紧凑基线 h-8 text-xs） */}
              <div className="space-y-2 border-t pt-2">
                <Button asChild variant="outline" className="h-8 w-full text-xs">
                  <Link to="/standards/$standardId/edit" params={{ standardId }}>
                    编辑
                  </Link>
                </Button>
                <Button
                  variant="destructive"
                  className="h-8 w-full text-xs"
                  disabled={isActive}
                  title={isActive ? '先经编辑改为 draft/retired 再删除' : undefined}
                  onClick={handleDelete}
                >
                  删除
                </Button>
                {/* ACTIVE 态删除禁用文案提示（phase14-05 表述：提示先 Retire） */}
                {isActive ? (
                  <p className="text-xs text-muted-foreground">
                    active 状态不可直接删除，先经编辑改为 draft / retired（Retire）后再删除
                  </p>
                ) : null}
                {/* 删除失败行内回显 */}
                {deleteMutation.isError ? (
                  <p className="text-xs text-destructive">
                    删除失败：{(deleteMutation.error as Error).message}
                  </p>
                ) : null}
              </div>
            </CardContent>
          </Card>
        </div>

        {/* 内容主区 — 目录结构 → 绑定管理 → Revision 历史（无 Card 嵌套，border-t 分隔） */}
        <div className="space-y-4 lg:col-span-2">
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
      </div>
    </div>
  )
}
