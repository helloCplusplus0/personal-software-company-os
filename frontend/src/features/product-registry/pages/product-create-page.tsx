import { useSearch, useNavigate } from '@tanstack/react-router'
import { toast } from 'sonner'
import { useCreateDraftProduct } from '../application/use-create-draft-product'
import { useProductCreateFormState } from '../application/use-product-create-form-state'
import { useProductCreateTemplateHandoff } from '../application/use-product-create-template-handoff'
import type { CreateProductInput } from '../types'
import { ProductCreateForm } from '../components/product-create-form'
import { Button } from '@/components/ui/button'
import { ArrowLeft, AlertTriangle, RefreshCw } from 'lucide-react'
import { BackToDashboardButton } from '@/features/dashboard/components/back-to-dashboard-button'
import { useDashboardBackButton } from '@/features/dashboard/lib/dashboard-source'

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
 * phase09-09 新增第四种来源上下文：
 * - fromTemplateReuse 存在 → 来自模板候选，承接 templateCandidateId / templateSource
 * - fromTemplateReuse 优先级高于 fromList / fromModuleDetail
 * - 取消返回按 templateSource 决定目的地
 *
 * 布局降级（phase04-05）：
 * - PC / 移动：单列垂直布局，主动作按钮无需横向滚动即可见
 */
export function ProductCreatePage() {
  const search = useSearch({ from: '/products/new' })
  const navigate = useNavigate()
  const { showBackButton: isFromDashboard } = useDashboardBackButton()

  // phase09-09：模板 handoff 编排
  const templateHandoff = useProductCreateTemplateHandoff({
    fromTemplateReuse: (search as any).fromTemplateReuse,
    templateCandidateId: (search as any).templateCandidateId,
    templateSource: (search as any).templateSource,
    templateSourceProductId: (search as any).templateSourceProductId,
    fromDashboard: (search as any).fromDashboard,
    dashboardSection: (search as any).dashboardSection,
    dashboardReturnTo: (search as any).dashboardReturnTo,
  })

  // phase09-09：正式 form state owner，接受模板预填初始值
  const formState = useProductCreateFormState(templateHandoff.prefillInitialValues)

  // phase04-06 来源上下文单值判定
  const fromList = search.fromList === true
  const fromModuleDetail = search.fromModuleDetail === true
  const hasSourceModule = Boolean(search.moduleId && search.moduleName)

  // phase06-15 §"既有 create 页面回收"：
  // 正式 create 主线已回收到 application owner，页面只保留表单编排、toast 与导航消费
  const mutation = useCreateDraftProduct()

  // phase09-09 取消返回路径 — 模板来源优先
  const handleReturn = () => {
    if (templateHandoff.isFromTemplate) {
      templateHandoff.handleReturn()
      return
    }
    if (fromList) {
      navigate({
        to: '/products',
        search: {
          queryText: search.queryText,
          statusFilter: search.statusFilter ?? 'all',
        },
      })
    } else if (fromModuleDetail && search.moduleId) {
      navigate({
        to: '/modules/$moduleId',
        params: { moduleId: search.moduleId },
      })
    } else {
      navigate({
        to: '/products',
        search: { statusFilter: 'all' },
      })
    }
  }

  // phase06-15：提交成功 / 失败的 toast 与导航由页面在 call-site 承接
  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault()
    const input = formState.buildSubmitInput()
    if (!input.name) return

    const submitInput: CreateProductInput = {
      name: input.name,
      description: input.description || undefined,
      status: input.status,
    }

    mutation.mutate(submitInput, {
      onSuccess: (response) => {
        toast.success('产品创建成功')

        const detailSearch: Record<string, unknown> = {}

        // phase09-09 模板来源回流优先
        if (templateHandoff.isFromTemplate) {
          const templateParams = templateHandoff.buildSuccessSearch()
          Object.assign(detailSearch, templateParams)
        }

        // phase04-06 回流时必须继续携带创建页已有的来源标记与必要上下文参数
        if (fromList) {
          detailSearch.fromList = true
          detailSearch.queryText = search.queryText
          detailSearch.statusFilter = search.statusFilter ?? 'all'
        } else if (fromModuleDetail) {
          detailSearch.fromModuleDetail = true
          detailSearch.moduleId = search.moduleId
          detailSearch.moduleName = search.moduleName
        }
        // phase05-13 提交成功后进入 Detail 页时必须继续保留 fromDashboard 等参数
        if (isFromDashboard) {
          detailSearch.fromDashboard = true
          detailSearch.dashboardSection = (search as any).dashboardSection
          detailSearch.dashboardReturnTo = (search as any).dashboardReturnTo
        }

        navigate({
          to: '/products/$productId',
          params: { productId: response.product_id },
          search: detailSearch,
        })
      },
      onError: (error: Error) => {
        toast.error('创建失败：' + error.message)
      },
    })
  }

  return (
    <div className="max-w-2xl space-y-4">
      <div className="flex items-center gap-3">
        {isFromDashboard ? (
          <BackToDashboardButton />
        ) : (
          <Button variant="ghost" size="sm" onClick={handleReturn}>
            <ArrowLeft className="mr-2 h-4 w-4" />
            {templateHandoff.isFromTemplate
              ? `返回${templateHandoff.templateSourceLabel}`
              : fromModuleDetail
                ? '返回模块详情'
                : '返回列表'}
          </Button>
        )}
        <h1 className="text-2xl font-bold">
          {templateHandoff.isFromTemplate ? '基于模板创建产品' : '新建产品'}
        </h1>
      </div>

      {/* phase09-09 模板来源摘要区 */}
      {templateHandoff.isFromTemplate && templateHandoff.templateSummary && (
        <div className="border-t pt-2">
          <div className="text-xs text-muted-foreground mb-1">
            来源：{templateHandoff.templateSummary.sourceLabel}
          </div>
          <h3 className="text-sm font-medium">{templateHandoff.templateSummary.templateTitle}</h3>
          <p className="text-xs text-muted-foreground mt-0.5">
            {templateHandoff.templateSummary.templateDescription}
          </p>
          {templateHandoff.templateSummary.moduleNames.length > 0 && (
            <div className="flex flex-wrap gap-1 mt-1">
              {templateHandoff.templateSummary.moduleNames.map((name) => (
                <span
                  key={name}
                  className="inline-flex items-center rounded-md bg-muted px-2 py-0.5 text-xs font-medium"
                >
                  {name}
                </span>
              ))}
            </div>
          )}
        </div>
      )}

      {/* phase09-09 模板 unavailable 成功态 */}
      {templateHandoff.isFromTemplate && templateHandoff.resolutionStatus === 'unavailable' && (
        <div className="flex items-start gap-2 rounded-lg border border-amber-200 bg-amber-50 p-3 text-sm">
          <AlertTriangle className="h-4 w-4 text-amber-600 mt-0.5 shrink-0" />
          <div>
            <p className="font-medium text-amber-800">模板来源已失效</p>
            <p className="text-amber-700 mt-0.5">
              模板来源已不可复读，但仍可继续手动创建产品。
            </p>
          </div>
        </div>
      )}

      {/* phase09-09 模板预填请求失败态 */}
      {templateHandoff.isFromTemplate && templateHandoff.resolutionStatus === 'error' && (
        <div className="flex items-start gap-2 rounded-lg border border-red-200 bg-red-50 p-3 text-sm">
          <AlertTriangle className="h-4 w-4 text-red-600 mt-0.5 shrink-0" />
          <div>
            <p className="font-medium text-red-800">模板预填加载失败</p>
            <p className="text-red-700 mt-0.5">
              无法加载模板预填数据，您仍可手动创建产品。
            </p>
            <Button
              variant="outline"
              size="sm"
              className="mt-2 h-7 px-2 text-xs"
              onClick={() => window.location.reload()}
            >
              <RefreshCw className="mr-1 h-3 w-3" />
              重试
            </Button>
          </div>
        </div>
      )}

      {/* phase04-06 来源上下文展示 — 从 Module Detail 带上下文进入时 */}
      {hasSourceModule && !templateHandoff.isFromTemplate && (
        <div className="rounded-lg border bg-muted/50 p-3 text-sm">
          <span className="text-muted-foreground">来源模块：</span>
          <span className="font-medium">{search.moduleName}</span>
        </div>
      )}

      <ProductCreateForm
        name={formState.name}
        description={formState.description}
        status={formState.status}
        onChangeName={formState.setName}
        onChangeDescription={formState.setDescription}
        onChangeStatus={formState.setStatus}
        isFromTemplate={templateHandoff.isFromTemplate}
        submitting={mutation.isPending}
        onSubmit={handleSubmit}
        submitError={mutation.isError ? (mutation.error as Error).message : undefined}
        isDirty={formState.isDirty}
      />
    </div>
  )
}
