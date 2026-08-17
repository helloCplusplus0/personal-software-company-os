/**
 * OnboardingPage — 首轮建链引导页面
 *
 * phase10-08 §"OnboardingPage 必须从页面级草稿协调改为链路状态驱动"
 *
 * 六段步骤语义：
 *   welcome → product → repository → module → decision → complete
 *
 * 关键约束：
 *   - 页面组件只保留页面壳、步骤 UI 与 owner 消费，不承担正式业务编排
 *   - 读取通过 useOnboardingRead 统一承接
 *   - 写动作通过 useOnboardingAction 统一承接
 *   - 不再依赖 URL search 保存草稿摘要
 *   - 不再直接持有四个 create owner
 *   - Product / Repository / Module / Decision 最小人工必填字段：
 *     Product: name
 *     Repository: name + url
 *     Module: name
 *     Decision: title + choice + reason
 */
import { useEffect, useState } from 'react'
import { useNavigate, useSearch } from '@tanstack/react-router'
import { useQueryClient } from '@tanstack/react-query'
import { Button } from '@/components/ui/button'
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import { Input } from '@/components/ui/input'
import { Label } from '@/components/ui/label'
import { Textarea } from '@/components/ui/textarea'
import { Skeleton } from '@/components/ui/skeleton'
import { toast } from 'sonner'
import { useOnboardingRead, ONBOARDING_STATE_QUERY_KEY } from '../data/use-onboarding-read'
import { useOnboardingAction } from '../application/use-onboarding-action'
import type { OnboardingActionSuccess } from '../application/use-onboarding-action'
import type { OnboardingStep } from '../types'
import {
  PRODUCT_SEMANTIC_LABEL,
  REPOSITORY_SEMANTIC_LABEL,
  MODULE_SEMANTIC_LABEL,
  DECISION_SEMANTIC_LABEL,
} from '@/features/project-context/data/shared-semantic-constants'

/** 步骤顺序与标签 */
const STEP_LABELS: Record<OnboardingStep, string> = {
  welcome: '欢迎',
  product: '创建产品',
  repository: '创建仓库',
  module: '创建模块',
  decision: '记录决策',
  complete: '完成',
}

const STEP_ORDER: OnboardingStep[] = ['welcome', 'product', 'repository', 'module', 'decision', 'complete']

export function OnboardingPage() {
  const navigate = useNavigate()
  const queryClient = useQueryClient()
  const onboardingQuery = useOnboardingRead()
  const action = useOnboardingAction()

  // phase10-08：从 detail 页返回时携带 onboardingStep，作为一次性恢复焦点
  const returnSearch = useSearch({ from: '/onboarding' })

  const [startStep, setStartStep] = useState<OnboardingStep | null>(null)
  const [focusedStep, setFocusedStep] = useState<OnboardingStep | null>(
    returnSearch.onboardingStep ?? null,
  )

  const chainState = onboardingQuery.data?.chain_state
  const firstRunState = onboardingQuery.data?.first_run_state
  const serverStep = chainState?.current_step ?? firstRunState?.current_step
  const currentProductId = chainState?.current_product_id ?? ''

  // 当前步骤优先级：
  // 1. detail 页返回的一次性 focusedStep
  // 2. welcome 首次点击后的一次性 startStep（本地起步兜底）
  // 3. 服务端 chain_state.current_step（长期事实源）
  // 4. 最终兜底 welcome
  const currentStep: OnboardingStep =
    focusedStep ?? startStep ?? serverStep ?? 'welcome'

  // 消费 detail 返回的一次性 onboardingStep，然后从 URL 移除
  useEffect(() => {
    if (!returnSearch.onboardingStep) {
      return
    }
    setFocusedStep(returnSearch.onboardingStep)
    navigate({
      to: '/onboarding',
      search: {},
      replace: true,
    })
  }, [navigate, returnSearch.onboardingStep])

  // detail 返回带来的 focusedStep 只是一层一次性兜底；
  // 一旦服务端 step 追平或前进，就必须让位给服务端链路状态。
  useEffect(() => {
    if (!focusedStep || !serverStep) {
      return
    }
    if (focusedStep !== serverStep) {
      setFocusedStep(null)
    }
  }, [focusedStep, serverStep])

  // 当服务端步骤已追平或超过本地起步步骤时，清空 startStep
  useEffect(() => {
    if (startStep && serverStep && serverStep !== 'welcome' && serverStep !== 'complete') {
      setStartStep(null)
    }
  }, [serverStep, startStep])

  // phase10-08：complete 只允许由链路状态驱动，不能再用全局 first_run_state.completed
  // 覆盖当前锚点产品尚未补链完成的主线。
  useEffect(() => {
    if (chainState?.resume_status === 'completed') {
      setStartStep('complete')
    }
  }, [chainState?.resume_status])

  // 写动作成功后的统一处理
  // 注意：缓存失效由 useOnboardingAction 内部统一处理，页面层不再重复失效
  const handleActionSuccess = (result: OnboardingActionSuccess) => {
    if (result.successMessage) {
      toast.success(result.successMessage)
    }

    if (result.resultKind === 'handoff' && result.navigateTo && result.params) {
      // canonical handoff：跳转到 detail 页
      navigate({
        to: result.navigateTo as any,
        params: result.params as any,
        search: (result.search ?? {}) as any,
      })
      return
    }

    if (result.resultKind === 'complete') {
      setStartStep('complete')
      return
    }

    // advance：前进到下一步
    setFocusedStep(null)
    setStartStep(result.nextStep)
  }

  // ============================================================================
  // 渲染
  // ============================================================================

  if (onboardingQuery.isLoading) {
    return (
      <div className="max-w-2xl mx-auto space-y-4">
        <Skeleton className="h-8 w-48" />
        <Skeleton className="h-32 w-full" />
      </div>
    )
  }

  if (onboardingQuery.isError) {
    return (
      <div className="max-w-2xl mx-auto space-y-4">
        <h1 className="text-xl font-bold">首轮录入</h1>
        <div className="rounded-lg border border-destructive/50 bg-destructive/10 p-3">
          <p className="text-xs text-destructive mb-2">
            读取首轮状态失败：{(onboardingQuery.error as Error)?.message ?? '未知错误'}
          </p>
          <Button
            variant="outline"
            size="sm"
            onClick={() => queryClient.invalidateQueries({ queryKey: ONBOARDING_STATE_QUERY_KEY })}
          >
            重试
          </Button>
        </div>
      </div>
    )
  }

  return (
    <div className="max-w-2xl mx-auto space-y-4">
      {/* 步骤进度条 */}
      <StepProgress currentStep={currentStep} />

      {/* 步骤内容 */}
      {currentStep === 'welcome' && (
        <WelcomeStep
          onStart={() => {
            setStartStep('product')
          }}
        />
      )}

      {currentStep === 'product' && (
        <ProductStep
          action={action}
          onSuccess={handleActionSuccess}
          chainState={chainState}
        />
      )}

      {currentStep === 'repository' && (
        <RepositoryStep
          action={action}
          onSuccess={handleActionSuccess}
          currentProductId={currentProductId}
        />
      )}

      {currentStep === 'module' && (
        <ModuleStep
          action={action}
          onSuccess={handleActionSuccess}
          currentProductId={currentProductId}
        />
      )}

      {currentStep === 'decision' && (
        <DecisionStep
          action={action}
          onSuccess={handleActionSuccess}
        />
      )}

      {currentStep === 'complete' && (
        <CompleteStep
          chainState={chainState}
          onGoToDashboard={() => navigate({ to: '/dashboard' })}
        />
      )}
    </div>
  )
}

// ============================================================================
// 步骤进度条
// 移动端适配基线：移动端仅显示圆点 + 弹性连接线（等分铺满），隐藏文字标签避免横向溢出；
// sm 及以上恢复完整「圆点 + 标签 + 固定连接线」形态
// ============================================================================

function StepProgress({ currentStep }: { currentStep: OnboardingStep }) {
  const currentIndex = STEP_ORDER.indexOf(currentStep)

  return (
    <div className="flex items-center gap-1 sm:gap-2">
      {STEP_ORDER.map((step, index) => (
        <div
          key={step}
          className="flex min-w-0 flex-1 items-center gap-1 sm:flex-none sm:gap-2"
        >
          <div
            className={`flex h-7 w-7 shrink-0 items-center justify-center rounded-full text-xs font-medium sm:h-8 sm:w-8 sm:text-sm ${
              index < currentIndex
                ? 'bg-primary text-primary-foreground'
                : index === currentIndex
                  ? 'bg-primary text-primary-foreground ring-2 ring-primary ring-offset-2'
                  : 'bg-muted text-muted-foreground'
            }`}
          >
            {index < currentIndex ? '✓' : index + 1}
          </div>
          <span
            className={`hidden truncate text-xs sm:inline ${
              index === currentIndex ? 'font-semibold text-foreground' : 'text-muted-foreground'
            }`}
          >
            {STEP_LABELS[step]}
          </span>
          {index < STEP_ORDER.length - 1 && (
            <div
              className={`h-px min-w-2 flex-1 sm:w-8 sm:flex-none sm:min-w-0 ${
                index < currentIndex ? 'bg-primary' : 'bg-muted'
              }`}
            />
          )}
        </div>
      ))}
    </div>
  )
}

// ============================================================================
// Welcome 步骤
// ============================================================================

function WelcomeStep({ onStart }: { onStart: () => void }) {
  return (
    <div className="space-y-4">
      <div>
        <h1 className="text-xl font-bold">欢迎使用 PSCO</h1>
        <p className="mt-2 text-sm text-muted-foreground">
            Personal Software Company OS 帮助你管理{PRODUCT_SEMANTIC_LABEL}、登记
            {REPOSITORY_SEMANTIC_LABEL}、盘点{MODULE_SEMANTIC_LABEL}并记录
            {DECISION_SEMANTIC_LABEL}。
        </p>
      </div>
      <div className="rounded-lg border bg-muted/50 p-3 space-y-2 text-sm">
        <p className="font-medium">首轮录入需要完成以下四步：</p>
        <ol className="list-decimal list-inside space-y-1 text-xs text-muted-foreground">
          <li>登记一个{PRODUCT_SEMANTIC_LABEL}（Product）</li>
          <li>登记一个{REPOSITORY_SEMANTIC_LABEL}（Repository）</li>
          <li>登记一个{MODULE_SEMANTIC_LABEL}（Module）</li>
          <li>记录一条{DECISION_SEMANTIC_LABEL}（Decision）</li>
        </ol>
        <p className="text-xs text-muted-foreground">
          每一步只需填写最小必填字段，其余可在后续补充。
        </p>
      </div>
      <Button onClick={onStart} size="lg">
        开始首轮录入
      </Button>
    </div>
  )
}

// ============================================================================
// Product 步骤（phase10-08：通过 useOnboardingAction 写入）
// ============================================================================

function ProductStep({
  action,
  onSuccess,
  chainState,
}: {
  action: ReturnType<typeof useOnboardingAction>
  onSuccess: (result: OnboardingActionSuccess) => void
  chainState: import('../types').OnboardingChainState | undefined
}) {
  const [name, setName] = useState('')
  const mutation = action.productMutation

  // 如果已有产品（chainState 显示 product 步骤已完成），展示摘要卡片
  const isCompleted = chainState && chainState.current_step !== 'product' && chainState.current_step !== 'welcome'
  if (isCompleted && chainState.current_product_id) {
    return (
      <DraftSummaryCard
        title="产品已创建"
        description={`当前产品 ID: ${chainState.current_product_id}`}
        editLabel="继续编辑产品"
        onEdit={() => {
          // 跳转到产品详情
          window.location.href = `/products/${chainState.current_product_id}`
        }}
      />
    )
  }

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()
    if (!name.trim()) {
      toast.error('请输入产品名称')
      return
    }
    try {
      const result = await action.submitProduct({ name: name.trim() })
      setName('')
      onSuccess(result)
    } catch (err) {
      toast.error('创建失败：' + (err as Error).message)
    }
  }

  return (
    <div className="space-y-4">
      <div>
        <h2 className="text-xl font-bold">创建产品</h2>
        <p className="text-sm text-muted-foreground">最小必填：产品名称</p>
      </div>
      <form onSubmit={handleSubmit} className="space-y-4">
        <div className="space-y-2">
          <Label htmlFor="product-name">产品名称 *</Label>
          <Input
            id="product-name"
            value={name}
            onChange={(e) => setName(e.target.value)}
            placeholder="例如：我的第一个产品"
            disabled={mutation.isPending}
          />
        </div>
        {mutation.isError && (
          <p className="text-sm text-destructive">{(mutation.error as Error).message}</p>
        )}
        <Button type="submit" disabled={mutation.isPending || !name.trim()}>
          {mutation.isPending ? '创建中...' : '创建并继续'}
        </Button>
      </form>
    </div>
  )
}

// ============================================================================
// Repository 步骤
// ============================================================================

function RepositoryStep({
  action,
  onSuccess,
  currentProductId,
}: {
  action: ReturnType<typeof useOnboardingAction>
  onSuccess: (result: OnboardingActionSuccess) => void
  currentProductId: string
}) {
  const [name, setName] = useState('')
  const [url, setUrl] = useState('')
  const mutation = action.repositoryMutation

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()
    if (!name.trim()) {
      toast.error('请输入仓库名称')
      return
    }
    if (!url.trim()) {
      toast.error('请输入仓库 URL')
      return
    }
    try {
      const result = await action.submitRepository(
        { name: name.trim(), url: url.trim() },
        currentProductId,
      )
      setName('')
      setUrl('')
      onSuccess(result)
    } catch (err) {
      toast.error('创建失败：' + (err as Error).message)
    }
  }

  return (
    <div className="space-y-4">
      <div>
        <h2 className="text-xl font-bold">创建仓库</h2>
        <p className="text-sm text-muted-foreground">最小必填：仓库名称 + URL</p>
      </div>
      <form onSubmit={handleSubmit} className="space-y-4">
        <div className="space-y-2">
          <Label htmlFor="repo-name">仓库名称 *</Label>
          <Input
            id="repo-name"
            value={name}
            onChange={(e) => setName(e.target.value)}
            placeholder="例如：main-repo"
            disabled={mutation.isPending}
          />
        </div>
        <div className="space-y-2">
          <Label htmlFor="repo-url">仓库 URL *</Label>
          <Input
            id="repo-url"
            value={url}
            onChange={(e) => setUrl(e.target.value)}
            placeholder="例如：https://github.com/user/repo"
            disabled={mutation.isPending}
          />
        </div>
        {mutation.isError && (
          <p className="text-sm text-destructive">{(mutation.error as Error).message}</p>
        )}
        <Button type="submit" disabled={mutation.isPending || !name.trim() || !url.trim()}>
          {mutation.isPending ? '创建中...' : '创建并继续'}
        </Button>
      </form>
    </div>
  )
}

// ============================================================================
// Module 步骤
// ============================================================================

function ModuleStep({
  action,
  onSuccess,
  currentProductId,
}: {
  action: ReturnType<typeof useOnboardingAction>
  onSuccess: (result: OnboardingActionSuccess) => void
  currentProductId: string
}) {
  const [name, setName] = useState('')
  const mutation = action.moduleMutation

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()
    if (!name.trim()) {
      toast.error('请输入模块名称')
      return
    }
    try {
      const result = await action.submitModule(
        { name: name.trim() },
        currentProductId,
      )
      setName('')
      onSuccess(result)
    } catch (err) {
      toast.error('创建失败：' + (err as Error).message)
    }
  }

  return (
    <div className="space-y-4">
      <div>
        <h2 className="text-xl font-bold">创建模块</h2>
        <p className="text-sm text-muted-foreground">最小必填：模块名称</p>
      </div>
      <form onSubmit={handleSubmit} className="space-y-4">
        <div className="space-y-2">
          <Label htmlFor="module-name">模块名称 *</Label>
          <Input
            id="module-name"
            value={name}
            onChange={(e) => setName(e.target.value)}
            placeholder="例如：auth-service"
            disabled={mutation.isPending}
          />
        </div>
        {mutation.isError && (
          <p className="text-sm text-destructive">{(mutation.error as Error).message}</p>
        )}
        <Button type="submit" disabled={mutation.isPending || !name.trim()}>
          {mutation.isPending ? '创建中...' : '创建并继续'}
        </Button>
      </form>
    </div>
  )
}

// ============================================================================
// Decision 步骤
// ============================================================================

function DecisionStep({
  action,
  onSuccess,
}: {
  action: ReturnType<typeof useOnboardingAction>
  onSuccess: (result: OnboardingActionSuccess) => void
}) {
  const [title, setTitle] = useState('')
  const [choice, setChoice] = useState('')
  const [reason, setReason] = useState('')
  const mutation = action.decisionMutation

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()
    if (!title.trim()) {
      toast.error('请输入决策标题')
      return
    }
    if (!choice.trim()) {
      toast.error('请输入决策选择')
      return
    }
    if (!reason.trim()) {
      toast.error('请输入决策理由')
      return
    }
    try {
      const result = await action.submitDecision({
        title: title.trim(),
        choice: choice.trim(),
        reason: reason.trim(),
      })
      setTitle('')
      setChoice('')
      setReason('')
      onSuccess(result)
    } catch (err) {
      toast.error('创建失败：' + (err as Error).message)
    }
  }

  return (
    <div className="space-y-4">
      <div>
        <h2 className="text-xl font-bold">记录决策</h2>
        <p className="text-sm text-muted-foreground">最小必填：标题 + 选择 + 理由</p>
      </div>
      <form onSubmit={handleSubmit} className="space-y-4">
        <div className="space-y-2">
          <Label htmlFor="decision-title">决策标题 *</Label>
          <Input
            id="decision-title"
            value={title}
            onChange={(e) => setTitle(e.target.value)}
            placeholder="例如：技术选型决策"
            disabled={mutation.isPending}
          />
        </div>
        <div className="space-y-2">
          <Label htmlFor="decision-choice">选择 *</Label>
          <Input
            id="decision-choice"
            value={choice}
            onChange={(e) => setChoice(e.target.value)}
            placeholder="例如：使用 PostgreSQL"
            disabled={mutation.isPending}
          />
        </div>
        <div className="space-y-2">
          <Label htmlFor="decision-reason">理由 *</Label>
          <Textarea
            id="decision-reason"
            value={reason}
            onChange={(e) => setReason(e.target.value)}
            placeholder="例如：关系型数据库满足当前数据模型需求"
            disabled={mutation.isPending}
            rows={3}
          />
        </div>
        {mutation.isError && (
          <p className="text-sm text-destructive">{(mutation.error as Error).message}</p>
        )}
        <Button
          type="submit"
          disabled={mutation.isPending || !title.trim() || !choice.trim() || !reason.trim()}
        >
          {mutation.isPending ? '创建中...' : '创建并完成'}
        </Button>
      </form>
    </div>
  )
}

// ============================================================================
// Complete 步骤
// ============================================================================

function CompleteStep({
  chainState,
  onGoToDashboard,
}: {
  chainState: import('../types').OnboardingChainState | undefined
  onGoToDashboard: () => void
}) {
  return (
    <div className="space-y-4 text-center">
      <div className="mx-auto flex h-16 w-16 items-center justify-center rounded-full bg-primary text-primary-foreground text-3xl">
        ✓
      </div>
      <h1 className="text-xl font-bold">首轮录入完成</h1>
      <p className="text-sm text-muted-foreground">
        你已完成产品、仓库、模块和决策的最小登记。现在可以进入 Dashboard 查看系统概览与复用反馈。
      </p>
      {chainState?.current_product_id && (
        <p className="text-xs text-muted-foreground">
          当前产品 ID: {chainState.current_product_id}
        </p>
      )}
      <Button onClick={onGoToDashboard} size="lg">
        进入 Dashboard
      </Button>
    </div>
  )
}

// ============================================================================
// 共享组件
// ============================================================================

function DraftSummaryCard({
  title,
  description,
  editLabel,
  onEdit,
  onResume,
  resumeLabel,
}: {
  title: string
  description: string
  editLabel: string
  onEdit: () => void
  onResume?: () => void
  resumeLabel?: string
}) {
  return (
    <Card>
      <CardHeader>
        <CardTitle>{title}</CardTitle>
      </CardHeader>
      <CardContent className="space-y-3">
        <p className="text-xs text-muted-foreground">{description}</p>
        <div className="flex flex-wrap gap-2">
          <Button variant="outline" onClick={onEdit}>
            {editLabel}
          </Button>
          {onResume && resumeLabel && (
            <Button variant="ghost" onClick={onResume}>
              {resumeLabel}
            </Button>
          )}
        </div>
      </CardContent>
    </Card>
  )
}
