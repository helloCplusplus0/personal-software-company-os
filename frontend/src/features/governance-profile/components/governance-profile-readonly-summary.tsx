/**
 * GovernanceProfileReadonlySummary — 治理画像摘要回看层。
 *
 * phase13-06 冻结：
 *   - 承接 canonical 根级文件角色与全局规范资产的结构化摘要回看
 *   - structured_summary 优先于真实路径成为主阅读内容
 *   - entry_ref 只作为轻量 locator / secondary metadata 呈现，可复制
 *   - markdown_resolvable 只表达“是否允许正文回源”的能力状态
 *   - 使用紧凑化规范（text-xs / divide-y），不做大块路径说明区
 */
import { useState } from 'react'
import { Button } from '@/components/ui/button'
import { Badge } from '@/components/ui/badge'
import type { GovernanceProfile } from '@/gen/proto/psco/governance_profile/v1/governance_profile_pb'
import { GLOBAL_ASSET_MATRIX } from '../data/governance-profile-baseline'

interface GovernanceProfileReadonlySummaryProps {
  profile: GovernanceProfile
}

export function GovernanceProfileReadonlySummary({ profile }: GovernanceProfileReadonlySummaryProps) {
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

  // 按冻结矩阵顺序回看已承接资产；未承接（未持久化）的资产不进入回看列表
  const bindingsByName = new Map(profile.globalAssetBindings.map((b) => [b.name, b]))
  const boundAssets = GLOBAL_ASSET_MATRIX.flatMap((entry) => {
    const binding = bindingsByName.get(entry.name)
    return binding ? [{ matrix: entry, binding }] : []
  })

  return (
    <div className="space-y-3">
      {/* canonical 根级文件 — 角色回看 */}
      <div>
        <h4 className="mb-1.5 text-xs font-medium text-muted-foreground">canonical 根级文件</h4>
        <ul className="divide-y rounded-md border">
          {profile.canonicalRootFiles.map((file) => (
            <li key={file.fileName} className="flex flex-wrap items-baseline gap-2 px-3 py-1.5 text-xs">
              <span className="break-all font-medium">{file.fileName}</span>
              <span className="text-muted-foreground">{file.role}</span>
              {file.required && (
                <Badge variant="outline" className="h-4 min-w-4 px-1 text-[10px]">
                  必需
                </Badge>
              )}
            </li>
          ))}
        </ul>
      </div>

      {/* 全局规范资产 — 结构化摘要优先，entry_ref 轻量 locator */}
      <div>
        <h4 className="mb-1.5 text-xs font-medium text-muted-foreground">全局规范资产</h4>
        <ul className="divide-y rounded-md border">
          {boundAssets.map(({ matrix, binding }) => (
            <li key={binding.name} className="space-y-1 px-3 py-2 text-xs">
              <div className="flex flex-wrap items-baseline gap-2">
                <span className="break-all font-medium">{binding.name}</span>
                <span className="text-muted-foreground">{binding.role}</span>
                {binding.markdownResolvable === true && (
                  <Badge variant="secondary" className="h-4 min-w-4 px-1 text-[10px] font-normal">
                    可回源
                  </Badge>
                )}
              </div>
              {matrix.summaryRequired && binding.structuredSummary && (
                <p className="break-words text-muted-foreground">{binding.structuredSummary}</p>
              )}
              <div className="flex flex-wrap items-center gap-2">
                <span className="text-muted-foreground">入口</span>
                <code className="break-all rounded bg-muted px-1.5 py-0.5">{binding.entryRef}</code>
                <Button
                  type="button"
                  variant="outline"
                  size="sm"
                  className="h-6 px-1.5 text-[10px]"
                  onClick={() => void handleCopyEntryRef(binding.entryRef)}
                >
                  {copiedEntryRef === binding.entryRef ? '已复制' : '复制'}
                </Button>
              </div>
            </li>
          ))}
        </ul>
      </div>
    </div>
  )
}
