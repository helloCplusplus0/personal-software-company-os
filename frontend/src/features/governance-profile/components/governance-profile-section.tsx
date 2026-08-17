/**
 * GovernanceProfileSection — Repository detail 内唯一治理画像正式承接区。
 *
 * phase13-09 冻结：
 *   - 治理画像只出现在 Repository detail，本组件是该承接区的唯一页面消费入口
 *   - 三层结构（phase13-06）：概览层（只读）/ 结构化维护层（编辑表单）/ 摘要回看层
 *     同属一个治理画像区，不拆成多个并列卡片系统
 *   - 读取：唯一 query owner（repository_id 锚点）；画像未创建（NotFound）是合法空态，
 *     提供初始化入口（预填当前项目范式 v1 基线默认值）
 *   - 保存：唯一 mutation owner；成功后精准刷新当前 repository_id 的读取结果并回到回看态；
 *     失败停留表单上下文，保留草稿与错误可见
 *   - 治理画像区是 secondary governance 区，不抢占仓库业务主内容主视觉
 */
import { useState } from 'react'
import { Code, ConnectError } from '@connectrpc/connect'
import { Button } from '@/components/ui/button'
import { Skeleton } from '@/components/ui/skeleton'
import { useGovernanceProfileRead } from '../data/use-governance-profile-read'
import { useUpdateGovernanceProfile } from '../application/use-update-governance-profile'
import { GovernanceProfileOverview } from './governance-profile-overview'
import { GovernanceProfileReadonlySummary } from './governance-profile-readonly-summary'
import { GovernanceProfileForm } from './governance-profile-form'
import type { UpdateGovernanceProfileInitialSource } from '../types'

interface GovernanceProfileSectionProps {
  repositoryId: string
}

/** 判断读取失败是否为“画像未创建”合法空态（repository 不存在时页面级详情已先行报错） */
function isProfileNotFoundError(error: unknown): boolean {
  return error instanceof ConnectError && error.code === Code.NotFound
}

/** 从已保存画像构造表单初始值（只包含第一版正式可写集合） */
function toFormInitial(profile: NonNullable<ReturnType<typeof useGovernanceProfileRead>['data']>): UpdateGovernanceProfileInitialSource {
  return {
    templateSource: profile.templateSource,
    canonicalRootFiles: profile.canonicalRootFiles.map((f) => ({
      fileName: f.fileName,
      role: f.role,
      required: f.required,
    })),
    globalAssetBindings: profile.globalAssetBindings.map((b) => ({
      name: b.name,
      entryRef: b.entryRef,
      role: b.role,
      structuredSummary: b.structuredSummary,
    })),
  }
}

export function GovernanceProfileSection({ repositoryId }: GovernanceProfileSectionProps) {
  const profileQuery = useGovernanceProfileRead(repositoryId)
  const updateMutation = useUpdateGovernanceProfile()
  // 编辑态：已保存画像的“维护”或空态的“初始化”
  const [editMode, setEditMode] = useState(false)

  const handleSubmit = (request: Parameters<typeof updateMutation.mutate>[0]['request']) => {
    updateMutation.mutate(
      { repositoryId, request },
      {
        // 保存成功：精准失效已由 mutation owner 承接，这里只回到回看态
        onSuccess: () => setEditMode(false),
      },
    )
  }

  return (
    <div className="border-t pt-4">
      {/* 区域头 — secondary governance 区，保持轻量 */}
      <div className="mb-3 flex flex-wrap items-center justify-between gap-2">
        <div>
          <h3 className="text-sm font-medium">项目治理画像</h3>
          <p className="text-xs text-muted-foreground">
            项目级治理事实：范式、根级文件、全局规范资产与当前阶段
          </p>
        </div>
        {!editMode && (profileQuery.data || (profileQuery.isError && isProfileNotFoundError(profileQuery.error))) && (
          <Button
            type="button"
            variant="outline"
            size="sm"
            className="h-7 px-2 text-xs"
            onClick={() => setEditMode(true)}
          >
            {profileQuery.data ? '维护治理信息' : '初始化治理画像'}
          </Button>
        )}
      </div>

      {/* 读取中 — 轻量骨架 */}
      {profileQuery.isLoading && (
        <div className="space-y-2">
          <Skeleton className="h-16 w-full" />
          <Skeleton className="h-32 w-full" />
        </div>
      )}

      {/* 读取失败 — 区分“画像未创建”空态与真实错误 */}
      {profileQuery.isError && !isProfileNotFoundError(profileQuery.error) && (
        <div className="rounded-lg border border-destructive/50 bg-destructive/10 p-4">
          <p className="text-sm text-destructive">
            治理画像读取失败：{profileQuery.error instanceof Error ? profileQuery.error.message : '未知错误'}
          </p>
          <Button
            variant="outline"
            size="sm"
            className="mt-2 h-7 px-2 text-xs"
            onClick={() => void profileQuery.refetch()}
          >
            重试
          </Button>
        </div>
      )}

      {/* 空态 — 画像未创建是合法状态，不阻断页面其余内容 */}
      {profileQuery.isError && isProfileNotFoundError(profileQuery.error) && !editMode && (
        <p className="text-xs text-muted-foreground">
          该 Repository 尚未建立治理画像；点击“初始化治理画像”按当前项目范式 v1 基线开始维护。
        </p>
      )}

      {/* 回看态 — 概览层 + 摘要回看层 */}
      {profileQuery.data && !editMode && (
        <div className="space-y-3">
          <GovernanceProfileOverview profile={profileQuery.data} />
          <GovernanceProfileReadonlySummary profile={profileQuery.data} />
        </div>
      )}

      {/* 编辑态 — 结构化维护层（空态初始化时 initial 为 null，预填范式基线默认值） */}
      {editMode && (
        <GovernanceProfileForm
          initial={profileQuery.data ? toFormInitial(profileQuery.data) : null}
          submitting={updateMutation.isPending}
          submitError={updateMutation.isError ? (updateMutation.error as Error).message : undefined}
          onSubmit={handleSubmit}
          onCancel={() => {
            updateMutation.reset()
            setEditMode(false)
          }}
        />
      )}
    </div>
  )
}
