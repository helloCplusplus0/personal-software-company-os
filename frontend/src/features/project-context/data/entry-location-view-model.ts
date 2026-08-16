/**
 * entry-location-view-model — 规则 / phase 入口定位受控 adapter。
 *
 * phase12-09：将 entry_ref / entry_kind / label / summary 统一裁剪为
 * detail 页共享上下文区可消费的入口定位视图。
 *
 * 当前 consumer：仅 repository detail 页在接入 use-project-context-read 后使用。
 * 启用条件：页面已通过唯一 repositoryId 接入共享只读主线。
 */
import type { RuleEntry, PhaseEntry } from '@/gen/proto/psco/project_context/v1/project_context_pb'

export interface EntryLocationView {
  entryRef: string
  entryKind: string
  title: string
  summary: string
}

export function toEntryLocationView(entry: RuleEntry | PhaseEntry): EntryLocationView {
  return {
    entryRef: entry.entryRef,
    entryKind: entry.entryKind,
    title: entry.label,
    summary: 'statusSummary' in entry ? entry.statusSummary : entry.summary,
  }
}

export function toEntryLocationViews(entries: (RuleEntry | PhaseEntry)[]): EntryLocationView[] {
  return entries.map(toEntryLocationView)
}