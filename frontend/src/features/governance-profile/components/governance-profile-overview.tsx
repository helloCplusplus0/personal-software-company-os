/**
 * GovernanceProfileOverview — 治理画像概览层（只读）。
 *
 * phase13-06 冻结：
 *   - 只承接 project_profile_version / track_type / docs_workflow_layout /
 *     template_source / current_phase_* 的轻量只读展示
 *   - backend / database / frontend / proto 顶层目录矩阵只作为当前项目范式 v1
 *     的前端只读基线表达（来自前端常量，不来自后端字段）
 *   - current_phase_ref 是定位入口，以轻量 locator 呈现，可复制
 *   - 保持轻量，不做大块解释型文案
 */
import { useState } from 'react'
import { Button } from '@/components/ui/button'
import { Badge } from '@/components/ui/badge'
import type { GovernanceProfile } from '@/gen/proto/psco/governance_profile/v1/governance_profile_pb'
import {
  TOP_LEVEL_DIRECTORY_BASELINE,
  trackTypeLabel,
  phaseStatusLabel,
} from '../data/governance-profile-baseline'

interface GovernanceProfileOverviewProps {
  profile: GovernanceProfile
}

export function GovernanceProfileOverview({ profile }: GovernanceProfileOverviewProps) {
  const [copiedPhaseRef, setCopiedPhaseRef] = useState(false)

  const handleCopyPhaseRef = async () => {
    if (!profile.currentPhaseRef || typeof navigator === 'undefined' || !navigator.clipboard) {
      return
    }
    try {
      await navigator.clipboard.writeText(profile.currentPhaseRef)
      setCopiedPhaseRef(true)
    } catch {
      setCopiedPhaseRef(false)
    }
  }

  return (
    <div className="space-y-2">
      {/* 只读标量字段 — 紧凑两列（移动端单列） */}
      <dl className="grid grid-cols-1 gap-x-6 gap-y-1 sm:grid-cols-2">
        <div className="flex items-baseline gap-2 text-xs">
          <dt className="shrink-0 text-muted-foreground">画像版本</dt>
          <dd className="min-w-0 truncate font-medium">{profile.projectProfileVersion}</dd>
        </div>
        <div className="flex items-baseline gap-2 text-xs">
          <dt className="shrink-0 text-muted-foreground">技术路线</dt>
          <dd className="min-w-0 truncate font-medium">{trackTypeLabel(profile.trackType)}</dd>
        </div>
        <div className="flex items-baseline gap-2 text-xs">
          <dt className="shrink-0 text-muted-foreground">docs 布局</dt>
          <dd className="min-w-0 truncate font-medium">{profile.docsWorkflowLayout}</dd>
        </div>
        <div className="flex items-baseline gap-2 text-xs">
          <dt className="shrink-0 text-muted-foreground">模板来源</dt>
          <dd className="min-w-0 truncate font-medium">
            {profile.templateSource ? profile.templateSource : '未声明'}
          </dd>
        </div>
      </dl>

      {/* 当前阶段 — 只读回看，phase ref 以轻量 locator 呈现 */}
      <div className="space-y-1 text-xs">
        <div className="flex flex-wrap items-baseline gap-2">
          <span className="text-muted-foreground">当前阶段</span>
          <span className="font-medium">{profile.currentPhaseName}</span>
          <Badge variant="outline" className="h-4 min-w-4 px-1 text-[10px]">
            {phaseStatusLabel(profile.currentPhaseStatus)}
          </Badge>
        </div>
        <div className="flex flex-wrap items-center gap-2">
          <span className="text-muted-foreground">阶段入口</span>
          <code className="break-all rounded bg-muted px-1.5 py-0.5">{profile.currentPhaseRef}</code>
          <Button
            type="button"
            variant="outline"
            size="sm"
            className="h-6 px-1.5 text-[10px]"
            onClick={() => void handleCopyPhaseRef()}
          >
            {copiedPhaseRef ? '已复制' : '复制'}
          </Button>
        </div>
      </div>

      {/* 顶层目录矩阵 — 当前项目范式 v1 前端只读基线表达（非后端字段） */}
      <div className="flex flex-wrap items-center gap-1.5">
        <span className="text-xs text-muted-foreground">顶层目录基线</span>
        {TOP_LEVEL_DIRECTORY_BASELINE.map((dir) => (
          <Badge key={dir} variant="secondary" className="h-4 min-w-4 px-1.5 text-[10px] font-normal">
            {dir}/
          </Badge>
        ))}
      </div>
    </div>
  )
}
