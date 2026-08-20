/**
 * Progress 切片类型定义
 *
 * 对齐 phase15-05 §"切片结构必须冻结"（types 对齐后端 snake_case）与
 * phase15-04 冻结的后端合同（3 RPC envelope）。
 * 前端消费模型从 pb 生成产物单向派生；pb ↔ domain 转换集中在本文件，
 * data / application 层不得散写 pb 映射（沿 standard/types.ts 模式）。
 */
import { timestampDate, timestampFromDate } from '@bufbuild/protobuf/wkt'
import type { Timestamp } from '@bufbuild/protobuf/wkt'
import {
  WorkflowType as WorkflowTypePb,
  EventKind as EventKindPb,
  ProgressSource as ProgressSourcePb,
} from '@/gen/proto/psco/progress/v1/progress_pb'
import type { ProgressEvent as ProgressEventPb } from '@/gen/proto/psco/progress/v1/progress_pb'

/** 三轨 workflow — phase / audit / fix（裁决④） */
export type WorkflowType = 'phase' | 'audit' | 'fix'

/** 事件类型 — phase 边界标记 + 任务完成 + 自由标注（裁决⑤） */
export type EventKind = 'phase_started' | 'phase_completed' | 'task_completed' | 'note'

/** 事件来源 — 预留三值；本阶段创建入口仅 manual（裁决⑧） */
export type ProgressSource = 'manual' | 'git' | 'agent'

/** ProgressEvent — progress_events 表投影（11 字段，List / Create / brief 统一消费模型） */
export interface ProgressEvent {
  id: string
  repository_id: string
  workflow_type: WorkflowType
  event_kind: EventKind
  task_key: string
  title: string
  detail: string
  evidence_ref: string
  source: ProgressSource
  /** 用户声明发生时间（允许补录历史，与 created_at 分离） */
  occurred_at: Date | null
  /** 系统录入时间（不在时间轴行内展示） */
  created_at: Date | null
}

/** ProgressSummary — BriefProgress 投影（DP-1 当前卡通道；空态恒有值） */
export interface ProgressSummary {
  /** 当前 phase 的 task_key（phaseNN）；空串含从未开始 / 全部完结两情形（同型零值） */
  current_phase_key: string
  current_phase_label: string
  latest_task_completed: ProgressEvent | null
  recent_events: ProgressEvent[]
}

/** CreateProgressEvent 写入参数 — 表单值模型，owner 内转换为 pb 请求 */
export interface CreateProgressEventInput {
  repository_id: string
  workflow_type: WorkflowType
  event_kind: EventKind
  task_key: string
  title: string
  detail: string
  evidence_ref: string
  /** datetime-local 字符串（浏览器本地时区语义，DP-3） */
  occurred_at: string
}

// ============================================================================
// K-1~K-4 正则常量（phase15-02 冻结；联动必填与 placeholder / 轻量提示依据）
// ============================================================================

/** K-1：phase 轨边界事件 task_key 格式（phaseNN） */
export const K1_TASK_KEY_PHASE = /^phase[0-9]{2,}$/

/** K-2：phase 轨任务项 task_key 格式（phaseNN-MM） */
export const K2_TASK_KEY_PHASE_TASK = /^phase[0-9]{2,}-[0-9]{2,}$/

/** K-3：audit 轨任务项 task_key 格式（audit_NNN） */
export const K3_TASK_KEY_AUDIT = /^audit_[0-9]{3,}$/

/** K-4：fix 轨任务项 task_key 格式（fix_NNN） */
export const K4_TASK_KEY_FIX = /^fix_[0-9]{3,}$/

// ============================================================================
// 合法组合判定（联动禁用依据 — phase15-02 合法矩阵 12 格的 UI 投影）
// ============================================================================

/** audit / fix 轨禁止 phase 边界标记（规则 7）；其余组合全部合法 */
export function isEventKindAllowed(workflowType: WorkflowType, eventKind: EventKind): boolean {
  if (workflowType === 'audit' || workflowType === 'fix') {
    return eventKind === 'task_completed' || eventKind === 'note'
  }
  return true
}

/**
 * 按矩阵取 task_key 格式正则；note 轨 task_key 可选 → null（规则 8）
 * 其余格子必填且格式固定（K-1~K-4 一一对应，不收窄不放宽）
 */
export function taskKeyPatternFor(
  workflowType: WorkflowType,
  eventKind: EventKind,
): RegExp | null {
  if (eventKind === 'note') return null
  if (eventKind === 'task_completed') {
    if (workflowType === 'phase') return K2_TASK_KEY_PHASE_TASK
    if (workflowType === 'audit') return K3_TASK_KEY_AUDIT
    return K4_TASK_KEY_FIX
  }
  return K1_TASK_KEY_PHASE
}

// ============================================================================
// 枚举映射（pb enum ↔ string union，沿 standard/types.ts 防御性回退模式）
// ============================================================================

/** pb → domain：未知值（UNSPECIFIED）防御性落 'phase'（后端 V1a 拒绝零值，理论不出现） */
export function pbToWorkflowType(v: WorkflowTypePb): WorkflowType {
  switch (v) {
    case WorkflowTypePb.AUDIT:
      return 'audit'
    case WorkflowTypePb.FIX:
      return 'fix'
    default:
      return 'phase'
  }
}

/** domain → pb */
export function workflowTypeToPb(t: WorkflowType): WorkflowTypePb {
  switch (t) {
    case 'audit':
      return WorkflowTypePb.AUDIT
    case 'fix':
      return WorkflowTypePb.FIX
    default:
      return WorkflowTypePb.PHASE
  }
}

/** pb → domain：未知值（UNSPECIFIED）防御性落 'note'（最中性值；后端 V1b 拒绝零值） */
export function pbToEventKind(v: EventKindPb): EventKind {
  switch (v) {
    case EventKindPb.PHASE_STARTED:
      return 'phase_started'
    case EventKindPb.PHASE_COMPLETED:
      return 'phase_completed'
    case EventKindPb.TASK_COMPLETED:
      return 'task_completed'
    default:
      return 'note'
  }
}

/** domain → pb */
export function eventKindToPb(k: EventKind): EventKindPb {
  switch (k) {
    case 'phase_started':
      return EventKindPb.PHASE_STARTED
    case 'phase_completed':
      return EventKindPb.PHASE_COMPLETED
    case 'task_completed':
      return EventKindPb.TASK_COMPLETED
    default:
      return EventKindPb.NOTE
  }
}

/** pb → domain：未知值（UNSPECIFIED）防御性落 'manual'（创建入口仅 manual，后端归一） */
export function pbToProgressSource(v: ProgressSourcePb): ProgressSource {
  switch (v) {
    case ProgressSourcePb.GIT:
      return 'git'
    case ProgressSourcePb.AGENT:
      return 'agent'
    default:
      return 'manual'
  }
}

// ============================================================================
// 消息转换与时间辅助（pb ↔ domain；DP-3 时区口径）
// ============================================================================

/** pb ProgressEvent → domain ProgressEvent；缺主体（理论不发生）返回 null */
export function pbToProgressEvent(pb: ProgressEventPb | undefined): ProgressEvent | null {
  if (!pb) return null
  return {
    id: pb.id ?? '',
    repository_id: pb.repositoryId ?? '',
    workflow_type: pbToWorkflowType(pb.workflowType),
    event_kind: pbToEventKind(pb.eventKind),
    task_key: pb.taskKey ?? '',
    title: pb.title ?? '',
    detail: pb.detail ?? '',
    evidence_ref: pb.evidenceRef ?? '',
    source: pbToProgressSource(pb.source),
    occurred_at: pb.occurredAt ? timestampDate(pb.occurredAt) : null,
    created_at: pb.createdAt ? timestampDate(pb.createdAt) : null,
  }
}

/** 浏览器本地当前时刻 → datetime-local 值（YYYY-MM-DDTHH:mm，DP-3 默认值） */
export function nowDatetimeLocal(): string {
  const now = new Date()
  const pad = (n: number) => String(n).padStart(2, '0')
  return `${now.getFullYear()}-${pad(now.getMonth() + 1)}-${pad(now.getDate())}T${pad(now.getHours())}:${pad(now.getMinutes())}`
}

/** datetime-local 值（浏览器本地时区语义）→ pb Timestamp（UTC）；补录历史同路径（DP-3） */
export function datetimeLocalToPbTimestamp(value: string): Timestamp {
  return timestampFromDate(new Date(value))
}

/** 零值摘要 — brief progress 块防御 undefined 时的空态恒有值承接（DP-1） */
export function emptyProgressSummary(): ProgressSummary {
  return {
    current_phase_key: '',
    current_phase_label: '',
    latest_task_completed: null,
    recent_events: [],
  }
}
