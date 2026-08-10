/**
 * DecisionCreatePage — Decision Create
 *
 * §9.1 状态模型：
 * - 草稿状态：idle / dirty（至少承接 title / context / problem / choice / reason / status）
 * - 提交状态：submitting / submit-success / submit-error
 * - 提交失败时停留当前页，保留草稿，错误显示在表单上下文
 * - 提交成功默认回流到 DecisionDetailPage（§5.11 从 Module Detail 带上下文创建后也回流到详情页）
 *
 * §5.11 来源上下文承接：
 * - 从 Module Detail 带上下文进入时，通过路由搜索参数承接 sourceModuleId / sourceModuleName
 * - DecisionContextSourcePanel 展示该来源 Module
 * - 提交时通过 source_module_id 持久化到后端
 *
 * §9.1 列表上下文跨页面恢复（单值化）：
 * - fromList === true：用户从 DecisionListPage 进入，返回列表时恢复 lastSearch
 * - fromList 不存在：用户从 Module Detail 或外部直达进入，返回列表时落到默认参数，
 *   不恢复历史筛选，避免错误继承旧上下文
 *
 * phase06-15 §"既有 create 页面回收"：
 * - 正式 create 主线已回收到 application owner，页面只保留表单编排、toast 与导航消费
 *
 * 布局降级（phase03-05）：
 * - PC：来源上下文区、表单区与动作区同屏可见
 * - 移动：单列垂直布局
 */
import { useSearch, useNavigate, Link } from '@tanstack/react-router'
import { toast } from 'sonner'
import { useCreateDraftDecision } from '../application/use-create-draft-decision'
import { DecisionContextSourcePanel } from '../components/decision-context-source-panel'
import { DecisionCreateForm } from '../components/decision-create-form'
import { Button } from '@/components/ui/button'
import { ArrowLeft } from 'lucide-react'
import { useDecisionListSearchStore } from '../stores/decision-list-search-store'
import type { CreateDecisionInput } from '../types'

export function DecisionCreatePage() {
  // §5.11 从路由搜索参数承接来源上下文；§9.1 fromList 承接列表上下文标记
  const search = useSearch({ from: '/decisions/new' })
  const navigate = useNavigate()
  // §9.1 从 store 读取最后一次列表搜索上下文
  const lastSearch = useDecisionListSearchStore((s) => s.lastSearch)
  // §9.1 单值化"来源列表上下文存在 / 不存在"：
  // - fromList === true（从 DecisionListPage 进入）：返回列表恢复 lastSearch
  // - fromList 不存在（从 Module Detail 或外部直达进入）：返回列表落默认参数，不恢复历史筛选
  const returnSearch = search.fromList ? lastSearch : { statusFilter: 'all' as const }

  const hasSourceContext = Boolean(search.sourceModuleId && search.sourceModuleName)

  // phase06-15 §"既有 create 页面回收"：
  // 正式 create 主线已回收到 application owner，页面只保留表单编排、toast 与导航消费
  const mutation = useCreateDraftDecision()

  const handleSubmit = (input: CreateDecisionInput) => {
    mutation.mutate(input, {
      onSuccess: (response) => {
        // §9.1 提交成功默认回流到 DecisionDetailPage
        // §6.4 只返回 decision_id，不返回完整 Decision 对象
        // §9.1 透传 fromList：从 List 发起的创建回流到 Detail 后，返回列表仍恢复原上下文；
        //     从 Module Detail 发起的创建不透传，返回列表落默认参数
        toast.success('决策创建成功')
        navigate({
          to: '/decisions/$decisionId',
          params: { decisionId: response.decision_id },
          search: { fromList: search.fromList },
        })
      },
      onError: (error: Error) => {
        // §9.1 提交失败时停留当前页，错误显示在表单上下文
        toast.error('创建失败：' + error.message)
      },
    })
  }

  return (
    <div className="max-w-2xl space-y-4">
      {/* 返回列表 — §9.1 按 fromList 单值化决定恢复 lastSearch 或落默认参数 */}
      <div className="flex items-center gap-3">
        <Button variant="ghost" size="sm" asChild>
          <Link to="/decisions" search={returnSearch}>
            <ArrowLeft className="mr-2 h-4 w-4" />
            返回列表
          </Link>
        </Button>
        <h1 className="text-2xl font-bold">记录决策</h1>
      </div>

      {/* §5.11 来源上下文展示 — 仅在从 Module Detail 带上下文进入时展示 */}
      {hasSourceContext && (
        <DecisionContextSourcePanel
          sourceModuleId={search.sourceModuleId!}
          sourceModuleName={search.sourceModuleName!}
        />
      )}

      {/* 结构化模板字段录入表单 */}
      <DecisionCreateForm
        submitting={mutation.isPending}
        onSubmit={handleSubmit}
        submitError={mutation.isError ? (mutation.error as Error).message : undefined}
        sourceModuleId={search.sourceModuleId}
      />
    </div>
  )
}
