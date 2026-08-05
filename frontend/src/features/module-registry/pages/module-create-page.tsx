import { useNavigate, Link } from '@tanstack/react-router'
import { ModuleCreateForm } from '../components/module-create-form'
import { Button } from '@/components/ui/button'
import { ArrowLeft } from 'lucide-react'
import { createModule } from '../data/module-registry-adapter'
import { useMutation, useQueryClient } from '@tanstack/react-query'
import { toast } from 'sonner'
import type { CreateModuleInput } from '../types'
import { useModuleListSearchStore } from '../stores/module-list-search-store'

/**
 * ModuleCreatePage — Module Create
 *
 * §8.4 状态模型：
 * - 草稿状态：idle / dirty（至少承接 name / description / status）
 * - 提交状态：submitting / submit-success / submit-error
 * - 提交失败时停留当前页，保留草稿，错误显示在表单上下文
 * - 提交成功默认回流到 ModuleDetailPage
 * - 从此页主动返回：保留原搜索参数上下文的 ModuleListPage（§7.4）
 */
export function ModuleCreatePage() {
  const navigate = useNavigate()
  const queryClient = useQueryClient()
  // §7.4 从 store 读取最后一次列表搜索上下文，返回列表时恢复
  const lastSearch = useModuleListSearchStore((s) => s.lastSearch)

  const mutation = useMutation({
    mutationFn: (input: CreateModuleInput) => createModule(input),
    onSuccess: (module) => {
      // §8.4 提交成功默认回流到 ModuleDetailPage
      queryClient.invalidateQueries({ queryKey: ['module-list'] })
      toast.success('模块创建成功')
      navigate({ to: '/modules/$moduleId', params: { moduleId: module.id } })
    },
    onError: (error: Error) => {
      // §8.4 提交失败时停留当前页，错误显示在表单上下文
      toast.error('创建失败：' + error.message)
    },
  })

  return (
    <div className="max-w-2xl space-y-4">
      <div className="flex items-center gap-3">
        <Button variant="ghost" size="sm" asChild>
          <Link to="/modules" search={lastSearch}>
            <ArrowLeft className="mr-2 h-4 w-4" />
            返回列表
          </Link>
        </Button>
        <h1 className="text-2xl font-bold">新建模块</h1>
      </div>

      <ModuleCreateForm
        submitting={mutation.isPending}
        onSubmit={(input) => mutation.mutate(input)}
        submitError={mutation.isError ? (mutation.error as Error).message : undefined}
      />
    </div>
  )
}
