import { useSearch, useNavigate } from '@tanstack/react-router'
import { useMutation, useQueryClient } from '@tanstack/react-query'
import { toast } from 'sonner'
import { createProduct } from '../data/product-registry-adapter'
import { ProductCreateForm } from '../components/product-create-form'
import { Button } from '@/components/ui/button'
import { ArrowLeft } from 'lucide-react'

/**
 * ProductCreatePage — Product Create
 *
 * phase04-06 状态模型：
 * - 草稿状态：idle / dirty（至少承接 name / description / status）
 * - 提交状态：submitting / submit-success / submit-error
 * - 提交失败时停留当前页，保留草稿，错误显示在表单上下文
 * - 提交成功默认回流到 ProductDetailPage
 *
 * phase04-06 来源上下文承接（由路由搜索参数派生，只允许三种之一）：
 * - fromList 存在 → 来自 Product List，承接 queryText / statusFilter
 * - fromModuleDetail 存在 → 来自 Module Detail，承接 moduleId / moduleName
 * - 无来源参数 → direct-entry
 *
 * phase04-06 提交成功回流：
 * - 回流时必须继续携带创建页已有的来源标记与必要上下文参数
 * - fromList → 继续携带 fromList + queryText / statusFilter
 * - fromModuleDetail → 继续携带 fromModuleDetail + moduleId / moduleName
 * - direct-entry → 不携带来源标记
 *
 * phase04-06 主动取消返回：
 * - fromList → 回 Product List + 原 queryText / statusFilter
 * - fromModuleDetail → 回原 ModuleDetailPage
 * - direct-entry → 回 Product List 默认筛选参数
 *
 * 布局降级（phase04-05）：
 * - PC / 移动：单列垂直布局，主动作按钮无需横向滚动即可见
 */
export function ProductCreatePage() {
  const search = useSearch({ from: '/products/new' })
  const navigate = useNavigate()
  const queryClient = useQueryClient()

  // phase04-06 来源上下文单值判定
  const fromList = search.fromList === true
  const fromModuleDetail = search.fromModuleDetail === true
  const hasSourceModule = Boolean(search.moduleId && search.moduleName)

  // phase04-06 主动取消返回路径 — 按真实来源决定
  const handleReturn = () => {
    if (fromList) {
      // fromList → 回 Product List + 原 queryText / statusFilter
      navigate({
        to: '/products',
        search: {
          queryText: search.queryText,
          statusFilter: search.statusFilter ?? 'all',
        },
      })
    } else if (fromModuleDetail && search.moduleId) {
      // fromModuleDetail → 回原 ModuleDetailPage
      navigate({
        to: '/modules/$moduleId',
        params: { moduleId: search.moduleId },
      })
    } else {
      // direct-entry → 回 Product List 默认筛选参数
      navigate({
        to: '/products',
        search: { statusFilter: 'all' },
      })
    }
  }

  const mutation = useMutation({
    mutationFn: createProduct,
    onSuccess: (response) => {
      // phase04-06 提交成功默认回流到 ProductDetailPage
      queryClient.invalidateQueries({ queryKey: ['product-list'] })
      toast.success('产品创建成功')

      // phase04-06 回流时必须继续携带创建页已有的来源标记与必要上下文参数
      const detailSearch: Record<string, unknown> = {}
      if (fromList) {
        detailSearch.fromList = true
        detailSearch.queryText = search.queryText
        detailSearch.statusFilter = search.statusFilter ?? 'all'
      } else if (fromModuleDetail) {
        detailSearch.fromModuleDetail = true
        detailSearch.moduleId = search.moduleId
        detailSearch.moduleName = search.moduleName
      }
      // direct-entry → 不携带来源标记

      navigate({
        to: '/products/$productId',
        params: { productId: response.product_id },
        search: detailSearch,
      })
    },
    onError: (error: Error) => {
      // phase04-06 提交失败时停留当前页，错误显示在表单上下文
      toast.error('创建失败：' + error.message)
    },
  })

  return (
    <div className="max-w-2xl space-y-4">
      <div className="flex items-center gap-3">
        <Button variant="ghost" size="sm" onClick={handleReturn}>
          <ArrowLeft className="mr-2 h-4 w-4" />
          {fromModuleDetail ? '返回模块详情' : '返回列表'}
        </Button>
        <h1 className="text-2xl font-bold">新建产品</h1>
      </div>

      {/* phase04-06 来源上下文展示 — 从 Module Detail 带上下文进入时 */}
      {hasSourceModule && (
        <div className="rounded-lg border bg-muted/50 p-3 text-sm">
          <span className="text-muted-foreground">来源模块：</span>
          <span className="font-medium">{search.moduleName}</span>
        </div>
      )}

      <ProductCreateForm
        submitting={mutation.isPending}
        onSubmit={(input) => mutation.mutate(input)}
        submitError={mutation.isError ? (mutation.error as Error).message : undefined}
      />
    </div>
  )
}
