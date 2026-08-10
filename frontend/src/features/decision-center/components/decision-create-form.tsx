/**
 * DecisionCreateForm — 结构化模板字段录入表单
 *
 * §5.5 最小结构化模板字段：
 * - phase06 draft-first 创建必填：title / choice / reason
 * - 创建可选：context / problem / alternatives / impact
 * - status 默认预填 proposed，可由用户改写
 *
 * §5.5 alternatives 冻结为按顺序保留的文本条目集合，不引入嵌套对象结构。
 *
 * phase03-05 组件树冻结：
 * - DecisionCreateForm 只承接结构化模板字段录入
 * - DecisionCreateActions（提交与取消）内联于此组件
 *
 * 布局降级（phase03-05）：
 * - PC：表单字段多列网格
 * - 移动：单列垂直布局
 */
import { useState } from 'react'
import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import { Textarea } from '@/components/ui/textarea'
import { Label } from '@/components/ui/label'
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/select'
import { Plus, Trash2 } from 'lucide-react'
import type { DecisionStatus, CreateDecisionInput } from '../types'

interface DecisionCreateFormProps {
  submitting: boolean
  onSubmit: (input: CreateDecisionInput) => void
  submitError?: string
  /** 来源 Module 标识（从搜索参数传入，用于提交时携带） */
  sourceModuleId?: string
}

/** 初始草稿状态 */
const INITIAL_DRAFT = {
  title: '',
  context: '',
  problem: '',
  alternatives: [] as string[],
  choice: '',
  reason: '',
  impact: '',
  status: 'proposed' as DecisionStatus,
}

const STATUS_OPTIONS: { value: DecisionStatus; label: string }[] = [
  { value: 'proposed', label: 'Proposed' },
  { value: 'active', label: 'Active' },
  { value: 'superseded', label: 'Superseded' },
  { value: 'archived', label: 'Archived' },
]

export function DecisionCreateForm({ submitting, onSubmit, submitError, sourceModuleId }: DecisionCreateFormProps) {
  const [draft, setDraft] = useState(INITIAL_DRAFT)
  const [newAlternative, setNewAlternative] = useState('')

  /** 更新草稿字段 */
  const update = <K extends keyof typeof draft>(key: K, value: (typeof draft)[K]) => {
    setDraft((prev) => ({ ...prev, [key]: value }))
  }

  /** 添加备选方案 */
  const addAlternative = () => {
    const trimmed = newAlternative.trim()
    if (trimmed) {
      update('alternatives', [...draft.alternatives, trimmed])
      setNewAlternative('')
    }
  }

  /** 删除备选方案 */
  const removeAlternative = (index: number) => {
    update('alternatives', draft.alternatives.filter((_, i) => i !== index))
  }

  /** 提交表单 */
  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault()
    const title = draft.title.trim()
    const choice = draft.choice.trim()
    const reason = draft.reason.trim()
    if (!title || !choice || !reason) {
      return
    }

    onSubmit({
      title,
      choice,
      reason,
      context: draft.context.trim() === '' ? undefined : draft.context.trim(),
      problem: draft.problem.trim() === '' ? undefined : draft.problem.trim(),
      alternatives: draft.alternatives.length > 0 ? draft.alternatives : undefined,
      impact: draft.impact.trim() === '' ? undefined : draft.impact.trim(),
      status: draft.status,
      source_module_id: sourceModuleId || '',
    })
  }

  return (
    <form onSubmit={handleSubmit} className="space-y-4">
      {/* 提交错误反馈 — 停留在表单上下文内 */}
      {submitError && (
        <div className="rounded-lg border border-destructive/50 bg-destructive/10 p-3">
          <p className="text-sm text-destructive">{submitError}</p>
        </div>
      )}

      {/* 标题 — 必填 */}
      <div className="space-y-2">
        <Label htmlFor="title">标题 *</Label>
        <Input
          id="title"
          value={draft.title}
          onChange={(e) => update('title', e.target.value)}
          placeholder="简述决策内容"
          required
        />
      </div>

      {/* 上下文 — 可选 */}
      <div className="space-y-2">
        <Label htmlFor="context">上下文（可选）</Label>
        <Textarea
          id="context"
          value={draft.context}
          onChange={(e) => update('context', e.target.value)}
          placeholder="决策发生的背景与前提"
        />
      </div>

      {/* 问题 — 可选 */}
      <div className="space-y-2">
        <Label htmlFor="problem">问题（可选）</Label>
        <Textarea
          id="problem"
          value={draft.problem}
          onChange={(e) => update('problem', e.target.value)}
          placeholder="需要解决的核心问题"
        />
      </div>

      {/* 备选方案 — 可选，按顺序保留的文本条目集合 */}
      <div className="space-y-2">
        <Label>备选方案（可选）</Label>
        <div className="flex gap-2">
          <Input
            value={newAlternative}
            onChange={(e) => setNewAlternative(e.target.value)}
            onKeyDown={(e) => {
              if (e.key === 'Enter') {
                e.preventDefault()
                addAlternative()
              }
            }}
            placeholder="输入备选方案后回车添加"
          />
          <Button type="button" variant="outline" size="icon" onClick={addAlternative}>
            <Plus className="h-4 w-4" />
          </Button>
        </div>
        {draft.alternatives.length > 0 && (
          <ol className="space-y-1">
            {draft.alternatives.map((alt, i) => (
              <li key={i} className="flex items-center justify-between rounded-md border px-3 py-1.5">
                <span className="text-sm">
                  <span className="text-muted-foreground mr-2">{i + 1}.</span>
                  {alt}
                </span>
                <Button
                  type="button"
                  variant="ghost"
                  size="icon"
                  className="h-6 w-6"
                  onClick={() => removeAlternative(i)}
                >
                  <Trash2 className="h-3 w-3" />
                </Button>
              </li>
            ))}
          </ol>
        )}
      </div>

      {/* 选择 — 必填 */}
      <div className="space-y-2">
        <Label htmlFor="choice">选择 *</Label>
        <Textarea
          id="choice"
          value={draft.choice}
          onChange={(e) => update('choice', e.target.value)}
          placeholder="最终选择的方案"
          required
        />
      </div>

      {/* 理由 — 必填 */}
      <div className="space-y-2">
        <Label htmlFor="reason">理由 *</Label>
        <Textarea
          id="reason"
          value={draft.reason}
          onChange={(e) => update('reason', e.target.value)}
          placeholder="做出该选择的理由"
          required
        />
      </div>

      {/* 影响 — 可选 */}
      <div className="space-y-2">
        <Label htmlFor="impact">影响（可选）</Label>
        <Textarea
          id="impact"
          value={draft.impact}
          onChange={(e) => update('impact', e.target.value)}
          placeholder="该决策带来的影响"
        />
      </div>

      {/* 状态 — 默认 proposed，可选调整 */}
      <div className="space-y-2">
        <Label>状态</Label>
        <Select
          value={draft.status}
          onValueChange={(value) => update('status', value as DecisionStatus)}
        >
          <SelectTrigger>
            <SelectValue />
          </SelectTrigger>
          <SelectContent>
            {STATUS_OPTIONS.map((opt) => (
              <SelectItem key={opt.value} value={opt.value}>
                {opt.label}
              </SelectItem>
            ))}
          </SelectContent>
        </Select>
      </div>

      {/* 动作区 — DecisionCreateActions 内联于此 */}
      <div className="flex gap-2 pt-2">
        <Button type="submit" disabled={submitting}>
          {submitting ? '提交中...' : '记录决策'}
        </Button>
      </div>
    </form>
  )
}
