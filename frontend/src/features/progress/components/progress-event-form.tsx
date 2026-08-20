/**
 * ProgressEventForm — 推进事件录入表单
 *
 * phase15-05 §"录入表单交互规格必须冻结"：
 * - 最小摩擦：occurred_at 默认 now（DP-3）、event_kind 记住上次选择（localStorage）、
 *   workflow_type × event_kind 非法组合联动禁用矩阵、task_key 联动必填与
 *   placeholder 矩阵（K-1~K-4 正则的 UI 投影，不收窄不放宽）
 * - 校验反馈双层模型：前端轻量层行内即时提示（仅 UX 提示不阻断输入）；
 *   后端权威层错误经 owner normalizeError 行内回显（不自行翻译错误码）
 * - source 无输入位（裁决⑧）——请求不设置 source 字段，后端归一 manual
 * - 成功后不弹窗不跳转（内嵌区就地刷新）：表单按重置语义复位
 *   （title / task_key / detail / evidence_ref 清空 + occurred_at 重置 now；
 *   workflow_type / event_kind 保持——连续录入同轨同类事件是主场景）
 */
import { useState, type FormEvent } from 'react'
import { useCreateProgressEvent } from '../application/use-create-progress-event'
import { isEventKindAllowed, taskKeyPatternFor, nowDatetimeLocal } from '../types'
import type { WorkflowType, EventKind } from '../types'
import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import { Label } from '@/components/ui/label'
import { Textarea } from '@/components/ui/textarea'
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/select'

interface ProgressEventFormProps {
  repositoryId: string
}

/** localStorage key — event_kind 记忆（纯 UX 偏好默认值，不构成数据事实源） */
const LAST_EVENT_KIND_KEY = 'psco.progress.last-event-kind'

const WORKFLOW_TYPE_OPTIONS: WorkflowType[] = ['phase', 'audit', 'fix']
const EVENT_KIND_OPTIONS: EventKind[] = ['phase_started', 'phase_completed', 'task_completed', 'note']

/** 表单挂载初始化：读取 localStorage 记忆值；无记忆或非法回退 task_completed */
function initialEventKind(): EventKind {
  try {
    const remembered = window.localStorage.getItem(LAST_EVENT_KIND_KEY)
    return EVENT_KIND_OPTIONS.includes(remembered as EventKind)
      ? (remembered as EventKind)
      : 'task_completed'
  } catch {
    return 'task_completed'
  }
}

/** task_key placeholder 矩阵（K-1~K-4 正则的 UI 投影，提示语与格式一一对应） */
function taskKeyPlaceholder(workflowType: WorkflowType, eventKind: EventKind): string {
  if (eventKind === 'note') return '自由标注，可留空'
  if (eventKind === 'task_completed') {
    if (workflowType === 'phase') return 'phaseNN-MM（如 phase15-05）'
    if (workflowType === 'audit') return 'audit_NNN（如 audit_001）'
    return 'fix_NNN（如 fix_001）'
  }
  // phase × phase_started / phase_completed
  return 'phaseNN（如 phase15）'
}

export function ProgressEventForm({ repositoryId }: ProgressEventFormProps) {
  // ---- form state（默认值冻结：workflow_type=phase / event_kind=记忆值 / occurred_at=now）----
  const [workflowType, setWorkflowType] = useState<WorkflowType>('phase')
  const [eventKind, setEventKind] = useState<EventKind>(initialEventKind)
  const [taskKey, setTaskKey] = useState('')
  const [title, setTitle] = useState('')
  const [detail, setDetail] = useState('')
  const [evidenceRef, setEvidenceRef] = useState('')
  const [occurredAt, setOccurredAt] = useState(nowDatetimeLocal)

  // ---- 前端轻量提示（仅 UX 提示不阻断输入；权威校验在后端，经 normalizeError 回显）----
  const taskKeyPattern = taskKeyPatternFor(workflowType, eventKind)
  const titleHint = !title.trim() ? '必填' : title.length > 200 ? '上限 200 字符' : undefined
  const detailHint = detail.length > 2000 ? '上限 2000 字符' : undefined
  const evidenceRefHint =
    evidenceRef && !evidenceRef.startsWith('/') && !evidenceRef.startsWith('https://')
      ? '需以 / 或 https:// 开头'
      : undefined
  const taskKeyHint = taskKeyPattern
    ? !taskKey.trim()
      ? '必填'
      : taskKeyPattern.test(taskKey)
        ? undefined
        : '格式不符'
    : undefined

  const createMutation = useCreateProgressEvent((event) => {
    // 成功回流：重置语义冻结——title / task_key / detail / evidence_ref 清空 +
    // occurred_at 重置为 now；workflow_type / event_kind 保持当前选择
    setTaskKey('')
    setTitle('')
    setDetail('')
    setEvidenceRef('')
    setOccurredAt(nowDatetimeLocal())
    // event_kind 记忆：Create 成功后写入当次 event_kind（纯 UX 偏好，不构成数据事实源）
    try {
      window.localStorage.setItem(LAST_EVENT_KIND_KEY, event.event_kind)
    } catch {
      // localStorage 不可用时静默忽略（纯 UX 偏好，不影响数据事实）
    }
  })

  /** 联动禁用矩阵：切换 audit/fix 后 phase 边界标记不合法 → 当前值自动重置 task_completed */
  const handleWorkflowTypeChange = (next: WorkflowType) => {
    setWorkflowType(next)
    if (!isEventKindAllowed(next, eventKind)) setEventKind('task_completed')
  }

  const handleSubmit = (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault()
    createMutation.mutate({
      repository_id: repositoryId,
      workflow_type: workflowType,
      event_kind: eventKind,
      task_key: taskKey,
      title,
      detail,
      evidence_ref: evidenceRef,
      occurred_at: occurredAt,
    })
  }

  return (
    // 2026-08-19 用户 UI 反馈：短值字段紧凑化——PC 四列（四个短字段一行 +
    // title/evidence_ref 并排），移动端两列；detail 全宽
    <form className="space-y-2" onSubmit={handleSubmit}>
      <div className="grid grid-cols-2 gap-2 lg:grid-cols-4">
        <div className="space-y-1">
          <Label className="text-xs" htmlFor="progress-workflow-type">
            workflow_type
          </Label>
          <Select value={workflowType} onValueChange={(v) => handleWorkflowTypeChange(v as WorkflowType)}>
            <SelectTrigger id="progress-workflow-type" size="sm" className="w-full text-xs">
              <SelectValue />
            </SelectTrigger>
            <SelectContent>
              {WORKFLOW_TYPE_OPTIONS.map((t) => (
                <SelectItem key={t} value={t}>
                  {t}
                </SelectItem>
              ))}
            </SelectContent>
          </Select>
        </div>

        <div className="space-y-1">
          <Label className="text-xs" htmlFor="progress-event-kind">
            event_kind
          </Label>
          <Select value={eventKind} onValueChange={(v) => setEventKind(v as EventKind)}>
            <SelectTrigger id="progress-event-kind" size="sm" className="w-full text-xs">
              <SelectValue />
            </SelectTrigger>
            <SelectContent>
              {/* 联动禁用矩阵：audit / fix 时 phase_started / phase_completed disabled */}
              {EVENT_KIND_OPTIONS.map((k) => (
                <SelectItem key={k} value={k} disabled={!isEventKindAllowed(workflowType, k)}>
                  {k}
                </SelectItem>
              ))}
            </SelectContent>
          </Select>
        </div>

        <div className="space-y-1">
          <Label className="text-xs" htmlFor="progress-task-key">
            task_key
          </Label>
          <Input
            id="progress-task-key"
            className="h-8 text-xs"
            value={taskKey}
            onChange={(e) => setTaskKey(e.target.value)}
            placeholder={taskKeyPlaceholder(workflowType, eventKind)}
          />
          {taskKeyHint ? <p className="text-xs text-muted-foreground">{taskKeyHint}</p> : null}
        </div>

        <div className="space-y-1">
          <Label className="text-xs" htmlFor="progress-occurred-at">
            occurred_at
          </Label>
          {/* DP-3：浏览器原生 datetime-local 控件（分钟粒度，本地时区语义） */}
          <Input
            id="progress-occurred-at"
            type="datetime-local"
            className="h-8 text-xs"
            value={occurredAt}
            onChange={(e) => setOccurredAt(e.target.value)}
          />
        </div>

        <div className="col-span-2 space-y-1">
          <Label className="text-xs" htmlFor="progress-title">
            title
          </Label>
          <Input
            id="progress-title"
            className="h-8 text-xs"
            value={title}
            onChange={(e) => setTitle(e.target.value)}
            placeholder="一句话标题"
          />
          {titleHint ? <p className="text-xs text-muted-foreground">{titleHint}</p> : null}
        </div>

        <div className="col-span-2 space-y-1">
          <Label className="text-xs" htmlFor="progress-evidence-ref">
            evidence_ref
          </Label>
          <Input
            id="progress-evidence-ref"
            className="h-8 text-xs"
            value={evidenceRef}
            onChange={(e) => setEvidenceRef(e.target.value)}
            placeholder="https://… 或 /仓库内路径"
          />
          {evidenceRefHint ? <p className="text-xs text-muted-foreground">{evidenceRefHint}</p> : null}
        </div>

        <div className="col-span-2 space-y-1 lg:col-span-4">
          <Label className="text-xs" htmlFor="progress-detail">
            detail
          </Label>
          <Textarea
            id="progress-detail"
            rows={2}
            className="text-xs"
            value={detail}
            onChange={(e) => setDetail(e.target.value)}
          />
          {detailHint ? <p className="text-xs text-muted-foreground">{detailHint}</p> : null}
        </div>
      </div>

      {/* 后端权威层错误行内回显（失败停留表单上下文，保留已填值） */}
      {createMutation.isError ? (
        <p className="text-xs text-destructive">{(createMutation.error as Error).message}</p>
      ) : null}

      <Button type="submit" className="h-8 px-3 text-xs" disabled={createMutation.isPending}>
        {createMutation.isPending ? '提交中...' : '录入推进事件'}
      </Button>
    </form>
  )
}
