import { useParams, Link, useNavigate } from '@tanstack/react-router'
import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query'
import { ReleaseCreateForm } from '../components/release-create-form'
import { Button } from '@/components/ui/button'
import { ArrowLeft } from 'lucide-react'
import { Skeleton } from '@/components/ui/skeleton'
import { createRelease, fetchModuleDetail } from '../data/module-registry-adapter'
import { toast } from 'sonner'
import type { CreateReleaseInput } from '../types'
import { useModuleListSearchStore } from '../stores/module-list-search-store'

/**
 * ReleaseCreatePage — Release Create
 *
 * §8.4 状态模型：
 * - 当前模块标识来自路由参数 moduleId，不得复制可写全局状态
 * - 状态：idle / dirty / submitting / submit-success / submit-error
 * - 提交成功默认回流到当前模块的 ModuleDetailPage
 * - 提交成功回流后 ModuleDetailPage 必须承接最新版本列表读取
 *
 * §7.4 返回路径：从此页主动返回，默认返回当前模块的 ModuleDetailPage
 *
 * 模块上下文校验：
 * - Release Create 必须依附有效当前模块上下文，不得在无效 moduleId 下创建孤儿 release
 * - 加载时通过 ModuleDetailRead 校验 moduleId 存在性
 * - 模块不存在时阻止提交，引导返回列表
 */
export function ReleaseCreatePage() {
  const { moduleId } = useParams({ from: '/modules/$moduleId/releases/new' })
  const navigate = useNavigate()
  const queryClient = useQueryClient()
  const lastSearch = useModuleListSearchStore((s) => s.lastSearch)

  // 校验当前模块是否存在 — Release Create 必须依附有效当前模块上下文
  const { data: moduleData, isLoading, isError } = useQuery({
    queryKey: ['module-detail', moduleId],
    queryFn: () => fetchModuleDetail(moduleId),
    enabled: Boolean(moduleId),
  })

  const mutation = useMutation({
    mutationFn: (input: CreateReleaseInput) => createRelease(input),
    onSuccess: () => {
      // §8.4 提交成功回流后 ModuleDetailPage 必须承接最新版本列表读取
      queryClient.invalidateQueries({ queryKey: ['module-detail', moduleId] })
      toast.success('版本登记成功')
      navigate({ to: '/modules/$moduleId', params: { moduleId } })
    },
    onError: (error: Error) => {
      toast.error('版本登记失败：' + error.message)
    },
  })

  // 加载中：校验模块存在性
  if (isLoading) {
    return (
      <div className="max-w-2xl space-y-4">
        <Button variant="ghost" size="sm" asChild>
          <Link to="/modules/$moduleId" params={{ moduleId }}>
            <ArrowLeft className="mr-2 h-4 w-4" />
            返回详情
          </Link>
        </Button>
        <Skeleton className="h-12 w-full" />
        <Skeleton className="h-64 w-full" />
      </div>
    )
  }

  // 模块不存在：阻止提交，引导返回列表
  if (isError || !moduleData) {
    return (
      <div className="max-w-2xl space-y-4">
        <Button variant="ghost" size="sm" asChild>
          <Link to="/modules" search={lastSearch}>
            <ArrowLeft className="mr-2 h-4 w-4" />
            返回列表
          </Link>
        </Button>
        <div className="rounded-lg border border-destructive/50 bg-destructive/10 p-4">
          <p className="text-sm text-destructive">
            模块 {moduleId} 不存在，无法登记版本。
          </p>
        </div>
      </div>
    )
  }

  return (
    <div className="max-w-2xl space-y-4">
      <div className="flex items-center gap-3">
        <Button variant="ghost" size="sm" asChild>
          <Link to="/modules/$moduleId" params={{ moduleId }}>
            <ArrowLeft className="mr-2 h-4 w-4" />
            返回详情
          </Link>
        </Button>
        <h1 className="text-2xl font-bold">登记新版本 — {moduleData.module.name}</h1>
      </div>

      <ReleaseCreateForm
        moduleId={moduleId}
        submitting={mutation.isPending}
        onSubmit={(input) => mutation.mutate(input)}
        submitError={mutation.isError ? (mutation.error as Error).message : undefined}
      />
    </div>
  )
}
