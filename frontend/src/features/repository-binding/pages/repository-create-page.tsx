import { useSearch, useNavigate } from '@tanstack/react-router'
import { useMutation, useQueryClient } from '@tanstack/react-query'
import { toast } from 'sonner'
import { createRepository } from '../data/repository-binding-adapter'
import { RepositoryCreateForm } from '../components/repository-create-form'
import { Button } from '@/components/ui/button'
import { ArrowLeft } from 'lucide-react'
import {
  buildProductDetailSearchFromTransit,
  extractProductSourceTransit,
  type RepositoryBindingSearch,
} from '../utils/product-source-transit'

/**
 * RepositoryCreatePage — Repository Create
 *
 * phase04-06 状态模型：
 * - 草稿状态：idle / dirty（至少承接 name / url / provider / status）
 * - 提交状态：submitting / submit-success / submit-error
 * - 提交失败时停留当前页，保留草稿，错误显示在表单上下文
 * - 提交成功默认回流到 RepositoryBindingDetailPage
 *
 * phase04-06 来源上下文承接（由路由搜索参数派生，只允许四种之一）：
 * - fromList 存在 → 来自 Repository Binding / List，承接 queryText / statusFilter
 * - fromProductDetail 存在 → 来自 Product Detail，承接 productId / productName
 * - fromModuleDetail 存在 → 来自 Module Detail，承接 moduleId / moduleName
 * - 无来源参数 → direct-entry
 *
 * phase04-06 来源上下文透传（Product Detail 原始来源）：
 * - 从 Product Detail 进入时，必须继续携带 Product Detail 自己的来源上下文（product 前缀参数）
 * - 返回 Product Detail 时，基于透传参数恢复 Product Detail 的来源标记，不得退化为 direct-entry
 * - 成功回流到 RepositoryBindingDetailPage 时，继续携带透传参数
 *
 * phase04-06 提交成功回流：
 * - 回流时必须继续携带创建页已有的来源标记与必要上下文参数
 * - fromList → 继续携带 fromList + queryText / statusFilter
 * - fromProductDetail → 继续携带 fromProductDetail + productId / productName + Product Detail 来源透传
 * - fromModuleDetail → 继续携带 fromModuleDetail + moduleId / moduleName
 * - direct-entry → 不携带来源标记
 *
 * phase04-06 主动取消返回：
 * - fromList → 回 Repository Binding / List + 原 queryText / statusFilter
 * - fromProductDetail → 回原 ProductDetailPage（恢复 Product Detail 来源标记）
 * - fromModuleDetail → 回原 ModuleDetailPage
 * - direct-entry → 回 Repository Binding / List 默认筛选参数
 *
 * 布局降级（phase04-05）：
 * - PC / 移动：单列垂直布局，主动作按钮无需横向滚动即可见
 */
export function RepositoryCreatePage() {
  const search = useSearch({ from: '/repositories/new' })
  const navigate = useNavigate()
  const queryClient = useQueryClient()

  // phase04-06 来源上下文单值判定
  const fromList = search.fromList === true
  const fromProductDetail = search.fromProductDetail === true
  const fromModuleDetail = search.fromModuleDetail === true
  const hasSourceProduct = Boolean(search.productId && search.productName)
  const hasSourceModule = Boolean(search.moduleId && search.moduleName)

  // phase04-06 Product Detail 来源上下文透传参数
  const productTransit = extractProductSourceTransit(search as RepositoryBindingSearch)

  // phase04-06 主动取消返回路径 — 按真实来源决定
  const handleReturn = () => {
    if (fromList) {
      navigate({
        to: '/repositories',
        search: {
          queryText: search.queryText,
          statusFilter: search.statusFilter ?? 'all',
        },
      })
    } else if (fromProductDetail && search.productId) {
      // phase04-06 返回 Product Detail 时，恢复 Product Detail 的来源标记
      const productDetailSearch = buildProductDetailSearchFromTransit(search as RepositoryBindingSearch)
      navigate({
        to: '/products/$productId',
        params: { productId: search.productId },
        search: productDetailSearch,
      })
    } else if (fromModuleDetail && search.moduleId) {
      navigate({
        to: '/modules/$moduleId',
        params: { moduleId: search.moduleId },
      })
    } else {
      // direct-entry → 回 Repository Binding / List 默认筛选参数
      navigate({
        to: '/repositories',
        search: { statusFilter: 'all' },
      })
    }
  }

  const mutation = useMutation({
    mutationFn: createRepository,
    onSuccess: (response) => {
      // phase04-06 提交成功默认回流到 RepositoryBindingDetailPage
      queryClient.invalidateQueries({ queryKey: ['repository-list'] })
      toast.success('仓库创建成功')

      // phase04-06 回流时必须继续携带创建页已有的来源标记与必要上下文参数
      const detailSearch: Record<string, unknown> = {}
      if (fromList) {
        detailSearch.fromList = true
        detailSearch.queryText = search.queryText
        detailSearch.statusFilter = search.statusFilter ?? 'all'
      } else if (fromProductDetail) {
        detailSearch.fromProductDetail = true
        detailSearch.productId = search.productId
        detailSearch.productName = search.productName
        // phase04-06 继续携带 Product Detail 来源透传参数
        Object.assign(detailSearch, productTransit)
      } else if (fromModuleDetail) {
        detailSearch.fromModuleDetail = true
        detailSearch.moduleId = search.moduleId
        detailSearch.moduleName = search.moduleName
      }
      // direct-entry → 不携带来源标记

      navigate({
        to: '/repositories/$repositoryId',
        params: { repositoryId: response.repository_id },
        search: detailSearch,
      })
    },
    onError: (error: Error) => {
      // phase04-06 提交失败时停留当前页，错误显示在表单上下文
      toast.error('创建失败：' + error.message)
    },
  })

  // 返回按钮文案根据来源决定
  const returnLabel = fromProductDetail
    ? '返回产品详情'
    : fromModuleDetail
      ? '返回模块详情'
      : '返回列表'

  return (
    <div className="max-w-2xl space-y-4">
      <div className="flex items-center gap-3">
        <Button variant="ghost" size="sm" onClick={handleReturn}>
          <ArrowLeft className="mr-2 h-4 w-4" />
          {returnLabel}
        </Button>
        <h1 className="text-2xl font-bold">新建仓库</h1>
      </div>

      {/* phase04-06 来源上下文展示 */}
      {hasSourceProduct && (
        <div className="rounded-lg border bg-muted/50 p-3 text-sm">
          <span className="text-muted-foreground">来源产品：</span>
          <span className="font-medium">{search.productName}</span>
        </div>
      )}
      {hasSourceModule && (
        <div className="rounded-lg border bg-muted/50 p-3 text-sm">
          <span className="text-muted-foreground">来源模块：</span>
          <span className="font-medium">{search.moduleName}</span>
        </div>
      )}

      <RepositoryCreateForm
        submitting={mutation.isPending}
        onSubmit={(input) => mutation.mutate(input)}
        submitError={mutation.isError ? (mutation.error as Error).message : undefined}
      />
    </div>
  )
}
