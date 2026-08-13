/**
 * OnboardingPage — 首轮录入主线页面
 *
 * phase06-06 §"Onboarding 页面与步骤编排"
 * phase06-15 §"Onboarding 前端主线必须落地为唯一正式首轮入口"
 *
 * 六段步骤语义：
 *   welcome → product → repository → module → decision → complete
 *
 * 关键约束：
 *   - 页面组件不得内联正式 mutation 主线
 *   - 步骤表单必须通过各自 feature slice 的 application owner 提交
 *   - 四类对象都完成最小持久化后，页面必须进入 complete 步骤
 *   - Product / Repository / Module / Decision 最小人工必填字段：
 *     Product: name
 *     Repository: name + url
 *     Module: name
 *     Decision: title + choice + reason
 *
 * 步骤推进策略：
 *   - 读取 first_run_state 确定当前步骤
 *   - 每个 create 步骤成功后失效 onboarding-state query，触发 refetch
 *   - refetch 后 first_run_state 更新，页面自动推进到下一个未完成步骤
 *   - status = completed 时进入 complete 步骤，提供"进入 Dashboard"入口
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
import { useCreateDraftProduct } from '@/features/product-registry/application/use-create-draft-product'
import { useCreateDraftRepository } from '@/features/repository-binding/application/use-create-draft-repository'
import { useCreateDraftModule } from '@/features/module-registry/application/use-create-draft-module'
import { useCreateDraftDecision } from '@/features/decision-center/application/use-create-draft-decision'
import type { OnboardingStep } from '../types'
import {
  buildOnboardingDraftSearch,
  parseOnboardingDraftSearch,
  type OnboardingDraftKey,
  type OnboardingDraftMap,
  type OnboardingDraftSummary,
} from '../lib/onboarding-source-schema'

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

  // phase06-15 §"detail 页来源优先级"：
  // 从 canonical detail 页返回时携带 onboardingStep，并把对应步骤恢复为一次性焦点。
  const returnSearch = useSearch({ from: '/onboarding' })

  const [startStep, setStartStep] = useState<OnboardingStep | null>(null)
  const [focusedStep, setFocusedStep] = useState<OnboardingStep | null>(returnSearch.onboardingStep ?? null)
  const [drafts, setDrafts] = useState<OnboardingDraftMap>(() => parseOnboardingDraftSearch(returnSearch))

  const firstRunState = onboardingQuery.data?.first_run_state
  const serverStep = firstRunState?.current_step
  const status = firstRunState?.status

  // 当前步骤优先级（fix_001 修正）：
  // 1. detail 页返回的一次性 focusedStep
  // 2. welcome 首次点击后的一次性 startStep（本地起步兜底）
  // 3. 服务端 first_run_state.current_step（长期事实源）
  // 4. 最终兜底 welcome
  const currentStep: OnboardingStep =
    focusedStep ?? startStep ?? serverStep ?? 'welcome'

  useEffect(() => {
    if (!returnSearch.onboardingStep) {
      return
    }

    // 仅在本次回流中消费一次 onboardingStep，然后从 URL 中移除。
    // 这样既能恢复目标步骤，又不会把用户永久锁在旧步骤上。
    setFocusedStep(returnSearch.onboardingStep)
    navigate({
      to: '/onboarding',
        search: buildOnboardingDraftSearch(
          parseOnboardingDraftSearch({
            productDraftId: returnSearch.productDraftId,
            productDraftLabel: returnSearch.productDraftLabel,
            repositoryDraftId: returnSearch.repositoryDraftId,
            repositoryDraftLabel: returnSearch.repositoryDraftLabel,
            moduleDraftId: returnSearch.moduleDraftId,
            moduleDraftLabel: returnSearch.moduleDraftLabel,
            decisionDraftId: returnSearch.decisionDraftId,
            decisionDraftLabel: returnSearch.decisionDraftLabel,
          }),
        ),
      replace: true,
    })
  }, [
    navigate,
    returnSearch.onboardingStep,
    returnSearch.productDraftId,
    returnSearch.productDraftLabel,
    returnSearch.repositoryDraftId,
    returnSearch.repositoryDraftLabel,
    returnSearch.moduleDraftId,
    returnSearch.moduleDraftLabel,
    returnSearch.decisionDraftId,
    returnSearch.decisionDraftLabel,
  ])

  // fix_001 收敛规则：当服务端步骤已追平或超过本地起步步骤时，清空 startStep，
  // 让页面重新回到服务端 first_run_state.current_step 驱动。
  useEffect(() => {
    if (startStep && serverStep && serverStep !== 'welcome') {
      setStartStep(null)
    }
  }, [serverStep, startStep])

  // status = completed 时默认进入 complete 步骤；
  // 若当前存在从 detail 页回流的一次性 focusedStep，则保留该步骤优先级。
  useEffect(() => {
    if (status === 'completed') {
      setStartStep('complete')
    }
  }, [status])

  const persistDrafts = (nextDrafts: OnboardingDraftMap) => {
    navigate({
      to: '/onboarding',
      search: buildOnboardingDraftSearch(nextDrafts),
      replace: true,
    })
  }

  // 写操作成功后同步持久化草稿摘要并失效 onboarding-state，触发 refetch 推进步骤
  const handleStepSuccess = (step: OnboardingDraftKey, draft: OnboardingDraftSummary) => {
    const nextDrafts = { ...drafts, [step]: draft }
    setDrafts(nextDrafts)
    setFocusedStep(null)
    persistDrafts(nextDrafts)
    queryClient.invalidateQueries({ queryKey: ONBOARDING_STATE_QUERY_KEY })
  }

  const buildOnboardingSourceSearch = (step: OnboardingStep) => ({
    fromOnboarding: true,
    onboardingStep: step,
    ...buildOnboardingDraftSearch(drafts),
  })

  const navigateToDetail = (step: OnboardingStep, draft?: OnboardingDraftSummary) => {
    if (!draft) {
      return
    }

    const search = buildOnboardingSourceSearch(step)
    if (step === 'product') {
      navigate({
        to: '/products/$productId',
        params: { productId: draft.id },
        search,
      })
      return
    }
    if (step === 'repository') {
      navigate({
        to: '/repositories/$repositoryId',
        params: { repositoryId: draft.id },
        search,
      })
      return
    }
    if (step === 'module') {
      navigate({
        to: '/modules/$moduleId',
        params: { moduleId: draft.id },
        search,
      })
      return
    }
    if (step === 'decision') {
      navigate({
        to: '/decisions/$decisionId',
        params: { decisionId: draft.id },
        search,
      })
    }
  }

  const resumeServerProgress = () => {
    setFocusedStep(null)
    navigate({
      to: '/onboarding',
      search: buildOnboardingDraftSearch(drafts),
      replace: true,
    })
  }

  const buildResumeProps = (step: OnboardingDraftKey) => {
    if (!focusedStep || focusedStep !== step || !serverStep || serverStep === step) {
      return {}
    }

    return {
      resumeLabel: serverStep === 'complete' ? '返回完成页' : `继续到${STEP_LABELS[serverStep]}`,
      onResume: resumeServerProgress,
    }
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
        <h1 className="text-2xl font-bold">首轮录入</h1>
        <div className="rounded-lg border border-destructive/50 bg-destructive/10 p-4">
          <p className="text-sm text-destructive mb-3">
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
    <div className="max-w-2xl mx-auto space-y-6">
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
          draft={drafts.product}
          onEdit={() => navigateToDetail('product', drafts.product)}
          onSuccess={(draft) => {
              handleStepSuccess('product', draft)
          }}
            {...buildResumeProps('product')}
        />
      )}

      {currentStep === 'repository' && (
        <RepositoryStep
          draft={drafts.repository}
          onEdit={() => navigateToDetail('repository', drafts.repository)}
          onSuccess={(draft) => {
              handleStepSuccess('repository', draft)
          }}
            {...buildResumeProps('repository')}
        />
      )}

      {currentStep === 'module' && (
        <ModuleStep
          draft={drafts.module}
          onEdit={() => navigateToDetail('module', drafts.module)}
          onSuccess={(draft) => {
              handleStepSuccess('module', draft)
          }}
            {...buildResumeProps('module')}
        />
      )}

      {currentStep === 'decision' && (
        <DecisionStep
          draft={drafts.decision}
          onEdit={() => navigateToDetail('decision', drafts.decision)}
          onSuccess={(draft) => {
              handleStepSuccess('decision', draft)
          }}
            {...buildResumeProps('decision')}
        />
      )}

      {currentStep === 'complete' && (
        <CompleteStep
          drafts={drafts}
          onEdit={navigateToDetail}
          onGoToDashboard={() => navigate({ to: '/dashboard' })}
        />
      )}
    </div>
  )
}

// ============================================================================
// 步骤进度条
// ============================================================================

function StepProgress({ currentStep }: { currentStep: OnboardingStep }) {
  const currentIndex = STEP_ORDER.indexOf(currentStep)

  return (
    <div className="flex items-center gap-2">
      {STEP_ORDER.map((step, index) => (
        <div key={step} className="flex items-center gap-2">
          <div
            className={`flex h-8 w-8 items-center justify-center rounded-full text-sm font-medium ${
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
            className={`text-sm ${
              index === currentIndex ? 'font-semibold text-foreground' : 'text-muted-foreground'
            }`}
          >
            {STEP_LABELS[step]}
          </span>
          {index < STEP_ORDER.length - 1 && (
            <div className={`h-px w-8 ${index < currentIndex ? 'bg-primary' : 'bg-muted'}`} />
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
        <h1 className="text-2xl font-bold">欢迎使用 PSCO</h1>
        <p className="mt-2 text-muted-foreground">
          Personal Software Company OS 帮助你登记软件资产、记录决策并追踪复用反馈。
        </p>
      </div>
      <div className="rounded-lg border bg-muted/50 p-4 space-y-2 text-sm">
        <p className="font-medium">首轮录入需要完成以下四步：</p>
        <ol className="list-decimal list-inside space-y-1 text-muted-foreground">
          <li>创建一个产品（Product）</li>
          <li>创建一个仓库（Repository）</li>
          <li>创建一个模块（Module）</li>
          <li>记录一条决策（Decision）</li>
        </ol>
        <p className="text-muted-foreground">
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
// Product 步骤
// ============================================================================

function ProductStep({
  draft,
  onEdit,
  onSuccess,
  onResume,
  resumeLabel,
}: {
  draft?: OnboardingDraftSummary
  onEdit: () => void
  onSuccess: (draft: OnboardingDraftSummary) => void
  onResume?: () => void
  resumeLabel?: string
}) {
  const mutation = useCreateDraftProduct()
  const [name, setName] = useState('')

  if (draft) {
    return (
      <DraftSummaryCard
        title="产品已创建"
        description={`已创建产品：${draft.label}`}
        editLabel="继续编辑产品"
        onEdit={onEdit}
          onResume={onResume}
          resumeLabel={resumeLabel}
      />
    )
  }

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault()
    if (!name.trim()) {
      toast.error('请输入产品名称')
      return
    }
    mutation.mutate(
      { name: name.trim() },
      {
        onSuccess: (response) => {
          toast.success('产品创建成功')
          onSuccess({ id: response.product_id, label: name.trim() })
          setName('')
        },
        onError: (err: Error) => {
          toast.error('创建失败：' + err.message)
        },
      },
    )
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
  draft,
  onEdit,
  onSuccess,
  onResume,
  resumeLabel,
}: {
  draft?: OnboardingDraftSummary
  onEdit: () => void
  onSuccess: (draft: OnboardingDraftSummary) => void
  onResume?: () => void
  resumeLabel?: string
}) {
  const mutation = useCreateDraftRepository()
  const [name, setName] = useState('')
  const [url, setUrl] = useState('')

  if (draft) {
    return (
      <DraftSummaryCard
        title="仓库已创建"
        description={`已创建仓库：${draft.label}`}
        editLabel="继续编辑仓库"
        onEdit={onEdit}
          onResume={onResume}
          resumeLabel={resumeLabel}
      />
    )
  }

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault()
    if (!name.trim()) {
      toast.error('请输入仓库名称')
      return
    }
    if (!url.trim()) {
      toast.error('请输入仓库 URL')
      return
    }
    mutation.mutate(
      { name: name.trim(), url: url.trim() },
      {
        onSuccess: (response) => {
          toast.success('仓库创建成功')
          onSuccess({ id: response.repository_id, label: name.trim() })
          setName('')
          setUrl('')
        },
        onError: (err: Error) => {
          toast.error('创建失败：' + err.message)
        },
      },
    )
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
  draft,
  onEdit,
  onSuccess,
  onResume,
  resumeLabel,
}: {
  draft?: OnboardingDraftSummary
  onEdit: () => void
  onSuccess: (draft: OnboardingDraftSummary) => void
  onResume?: () => void
  resumeLabel?: string
}) {
  const mutation = useCreateDraftModule()
  const [name, setName] = useState('')

  if (draft) {
    return (
      <DraftSummaryCard
        title="模块已创建"
        description={`已创建模块：${draft.label}`}
        editLabel="继续编辑模块"
        onEdit={onEdit}
          onResume={onResume}
          resumeLabel={resumeLabel}
      />
    )
  }

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault()
    if (!name.trim()) {
      toast.error('请输入模块名称')
      return
    }
    mutation.mutate(
      { name: name.trim() },
      {
        onSuccess: (module) => {
          toast.success('模块创建成功')
          onSuccess({ id: module.id, label: name.trim() })
          setName('')
        },
        onError: (err: Error) => {
          toast.error('创建失败：' + err.message)
        },
      },
    )
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
  draft,
  onEdit,
  onSuccess,
  onResume,
  resumeLabel,
}: {
  draft?: OnboardingDraftSummary
  onEdit: () => void
  onSuccess: (draft: OnboardingDraftSummary) => void
  onResume?: () => void
  resumeLabel?: string
}) {
  const mutation = useCreateDraftDecision()
  const [title, setTitle] = useState('')
  const [choice, setChoice] = useState('')
  const [reason, setReason] = useState('')

  if (draft) {
    return (
      <DraftSummaryCard
        title="决策已记录"
        description={`已创建决策：${draft.label}`}
        editLabel="继续编辑决策"
        onEdit={onEdit}
          onResume={onResume}
          resumeLabel={resumeLabel}
      />
    )
  }

  const handleSubmit = (e: React.FormEvent) => {
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
    mutation.mutate(
      { title: title.trim(), choice: choice.trim(), reason: reason.trim() },
      {
        onSuccess: (response) => {
          toast.success('决策记录成功')
          onSuccess({ id: response.decision_id, label: title.trim() })
          setTitle('')
          setChoice('')
          setReason('')
        },
        onError: (err: Error) => {
          toast.error('创建失败：' + err.message)
        },
      },
    )
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
  drafts,
  onEdit,
  onGoToDashboard,
}: {
    drafts: OnboardingDraftMap
    onEdit: (step: OnboardingStep, draft?: OnboardingDraftSummary) => void
  onGoToDashboard: () => void
}) {
  return (
    <div className="space-y-4 text-center">
      <div className="mx-auto flex h-16 w-16 items-center justify-center rounded-full bg-primary text-primary-foreground text-3xl">
        ✓
      </div>
      <h1 className="text-2xl font-bold">首轮录入完成</h1>
      <p className="text-muted-foreground">
        你已完成产品、仓库、模块和决策的最小登记。现在可以进入 Dashboard 查看系统概览与复用反馈。
      </p>
      <div className="grid gap-3 text-left md:grid-cols-2">
        <CompleteDraftCard title="产品" draft={drafts.product} onEdit={() => onEdit('product', drafts.product)} />
        <CompleteDraftCard title="仓库" draft={drafts.repository} onEdit={() => onEdit('repository', drafts.repository)} />
        <CompleteDraftCard title="模块" draft={drafts.module} onEdit={() => onEdit('module', drafts.module)} />
        <CompleteDraftCard title="决策" draft={drafts.decision} onEdit={() => onEdit('decision', drafts.decision)} />
      </div>
      <Button onClick={onGoToDashboard} size="lg">
        进入 Dashboard
      </Button>
    </div>
  )
}

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
        <CardTitle className="text-xl">{title}</CardTitle>
      </CardHeader>
      <CardContent className="space-y-3">
        <p className="text-sm text-muted-foreground">{description}</p>
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

function CompleteDraftCard({
  title,
  draft,
  onEdit,
}: {
  title: string
  draft?: OnboardingDraftSummary
  onEdit: () => void
}) {
  return (
    <Card>
      <CardHeader className="pb-2">
        <CardTitle className="text-base">{title}</CardTitle>
      </CardHeader>
      <CardContent className="space-y-2">
        <p className="text-sm text-muted-foreground">
          {draft?.label ?? '当前会话未保留摘要，仍可通过对应列表或详情继续补全。'}
        </p>
        {draft && (
          <Button variant="outline" size="sm" onClick={onEdit}>
            继续编辑
          </Button>
        )}
      </CardContent>
    </Card>
  )
}
