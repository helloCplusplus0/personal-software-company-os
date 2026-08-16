import { useState } from 'react'
import { Button } from '@/components/ui/button'
import { RULE_ENTRY_LABEL, PHASE_ENTRY_LABEL, BOUNDARY_ENTRY_LABEL } from '../data/shared-semantic-constants'
import { toEntryLocationViews } from '../data/entry-location-view-model'
import type { UseProjectContextRead } from '../data/use-project-context-read'

interface ProjectContextSectionProps {
  query: UseProjectContextRead
  title?: string
  description?: string
}

export function ProjectContextSection({
  query,
  title = '项目上下文',
  description = '共享只读摘要、规则入口与阶段定位。',
}: ProjectContextSectionProps) {
  const [copiedEntryRef, setCopiedEntryRef] = useState<string | null>(null)

  const handleCopyEntryRef = async (entryRef: string) => {
    if (!entryRef || typeof navigator === 'undefined' || !navigator.clipboard) {
      return
    }

    try {
      await navigator.clipboard.writeText(entryRef)
      setCopiedEntryRef(entryRef)
    } catch {
      setCopiedEntryRef(null)
    }
  }

  const renderEntryLocationList = (entries: ReturnType<typeof toEntryLocationViews>) => (
    <div className="space-y-2">
      {entries.map((entry, index) => (
        <div
          key={`${entry.entryKind}-${entry.entryRef}-${index}`}
          className="rounded-md border bg-muted/20 px-3 py-2"
        >
          <div className="flex items-start justify-between gap-3">
            <div className="min-w-0 space-y-1">
              <div className="flex items-start gap-2 text-sm">
                <span className="shrink-0 text-muted-foreground">[{entry.entryKind}]</span>
                <span className="break-words font-medium">{entry.title}</span>
              </div>
              <p className="break-words text-sm text-muted-foreground">{entry.summary}</p>
              <div className="flex flex-wrap items-center gap-2 text-xs text-muted-foreground">
                <span>入口引用</span>
                <code className="break-all rounded bg-muted px-1.5 py-0.5">{entry.entryRef}</code>
              </div>
            </div>
            <Button
              type="button"
              variant="outline"
              size="sm"
              className="shrink-0"
              onClick={() => void handleCopyEntryRef(entry.entryRef)}
            >
              {copiedEntryRef === entry.entryRef ? '已复制' : '复制入口'}
            </Button>
          </div>
        </div>
      ))}
    </div>
  )

  return (
    <div className="space-y-3 border-t pt-4">
      <div className="flex items-center justify-between gap-2">
        <div>
          <h3 className="text-sm font-medium">{title}</h3>
          <p className="text-xs text-muted-foreground">{description}</p>
        </div>
        {query.isLoading && <span className="text-xs text-muted-foreground">读取中...</span>}
      </div>

      {query.isError ? (
        <div className="rounded-lg border border-destructive/50 bg-destructive/10 p-4">
          <p className="text-sm text-destructive">
            项目上下文读取失败：{query.error?.message ?? '未知错误'}
          </p>
          <Button variant="outline" size="sm" className="mt-2" onClick={() => void query.refetch()}>
            重试
          </Button>
        </div>
      ) : query.data ? (
        <>
          {query.data.rules.length > 0 && (
            <div>
              <h4 className="mb-2 text-xs font-medium text-muted-foreground">{RULE_ENTRY_LABEL}</h4>
              {renderEntryLocationList(toEntryLocationViews(query.data.rules))}
            </div>
          )}
          {query.data.phases.length > 0 && (
            <div>
              <h4 className="mb-2 text-xs font-medium text-muted-foreground">{PHASE_ENTRY_LABEL}</h4>
              {renderEntryLocationList(toEntryLocationViews(query.data.phases))}
            </div>
          )}
          {query.data.boundaries.length > 0 && (
            <div>
              <h4 className="mb-2 text-xs font-medium text-muted-foreground">{BOUNDARY_ENTRY_LABEL}</h4>
              <div className="space-y-1">
                {query.data.boundaries.map((entry, index) => (
                  <div key={index} className="text-sm">
                    <span className="font-medium">{entry.label}</span>
                    <span className="text-muted-foreground"> — {entry.summary}</span>
                  </div>
                ))}
              </div>
            </div>
          )}
        </>
      ) : null}
    </div>
  )
}
