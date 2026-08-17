/**
 * GovernanceProfileForm — 治理画像结构化维护层（编辑态表单）。
 *
 * phase13-09 冻结：
 *   - 表单范围收敛为 template_source / canonical_root_files[] / global_asset_bindings[]
 *   - 8 项全局规范资产矩阵为受控展示：name / kind 只读，不得新增第 9 项资产
 *   - 前 5 项摘要型资产必须提供 structured_summary 输入位（后 3 项允许不填）
 *   - docs_workflow_layout 留在只读概览层显示，不进入表单提交负载
 *   - 不提供 markdown 正文编辑器；track_type / current_phase_* 不进入表单
 *   - 保存失败停留表单上下文，保留草稿与错误可见；取消丢弃未保存修改
 *
 * 写路径约束：本组件只承接本地草稿状态与提交前预校验，
 * 保存请求与失效刷新统一由 application 层 mutation owner 承接。
 */
import { useState } from 'react'
import { Input } from '@/components/ui/input'
import { Label } from '@/components/ui/label'
import { Button } from '@/components/ui/button'
import { Badge } from '@/components/ui/badge'
import type {
  GovernanceProfileSaveDraft,
  UpdateGovernanceProfileInitialSource,
} from '../types'
import {
  CANONICAL_ROOT_FILE_DEFAULTS,
  GLOBAL_ASSET_MATRIX,
} from '../data/governance-profile-baseline'

// ============================================================================
// 表单草稿模型
// ============================================================================

interface RootFileDraft {
  key: number
  fileName: string
  role: string
  required: boolean
}

interface AssetDraft {
  name: string
  kind: string
  summaryRequired: boolean
  entryRef: string
  role: string
  structuredSummary: string
}

interface GovernanceProfileFormProps {
  /** 已保存画像的可写字段子集（编辑场景）；null 表示画像未创建（空态初始化，预填范式基线默认值） */
  initial: UpdateGovernanceProfileInitialSource | null
  submitting: boolean
  submitError?: string
  onSubmit: (request: GovernanceProfileSaveDraft) => void
  onCancel: () => void
}

// ============================================================================
// 组件
// ============================================================================

export function GovernanceProfileForm({
  initial,
  submitting,
  submitError,
  onSubmit,
  onCancel,
}: GovernanceProfileFormProps) {
  const nextKey = () => Date.now() + Math.random()

  const [templateSource, setTemplateSource] = useState(initial?.templateSource ?? '')

  const [rootFiles, setRootFiles] = useState<RootFileDraft[]>(() => {
    const source = initial?.canonicalRootFiles ?? CANONICAL_ROOT_FILE_DEFAULTS
    return source.map((f) => ({ key: nextKey(), fileName: f.fileName, role: f.role, required: f.required }))
  })

  const [assetRows, setAssetRows] = useState<AssetDraft[]>(() => {
    const saved = new Map((initial?.globalAssetBindings ?? []).map((b) => [b.name, b]))
    return GLOBAL_ASSET_MATRIX.map((entry) => {
      const binding = saved.get(entry.name)
      return {
        name: entry.name,
        kind: entry.kind,
        summaryRequired: entry.summaryRequired,
        entryRef: binding?.entryRef ?? '',
        role: binding?.role ?? '',
        structuredSummary: binding?.structuredSummary ?? '',
      }
    })
  })

  const [validationError, setValidationError] = useState<string | null>(null)

  // —— canonical 根级文件行操作 ——
  const updateRootFile = (key: number, patch: Partial<Omit<RootFileDraft, 'key'>>) => {
    setRootFiles((rows) => rows.map((r) => (r.key === key ? { ...r, ...patch } : r)))
  }
  const removeRootFile = (key: number) => {
    setRootFiles((rows) => rows.filter((r) => r.key !== key))
  }
  const addRootFile = () => {
    setRootFiles((rows) => [...rows, { key: nextKey(), fileName: '', role: '', required: false }])
  }

  // —— 资产行操作（name/kind 受控，只维护 entryRef / role / structuredSummary）——
  const updateAssetRow = (name: string, patch: Partial<Pick<AssetDraft, 'entryRef' | 'role' | 'structuredSummary'>>) => {
    setAssetRows((rows) => rows.map((r) => (r.name === name ? { ...r, ...patch } : r)))
  }

  /**
   * 提交前预校验 — 与后端 phase13-08 校验规则保持一致：
   * canonical_root_files 至少 1 项 / 字段非空 / 不重复；
   * 已承接资产（entry_ref 非空）至少 1 项 / role 非空 / 前 5 项 summary 必填。
   */
  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault()
    setValidationError(null)

    if (rootFiles.length === 0) {
      setValidationError('canonical 根级文件至少保留 1 项')
      return
    }
    const fileNames = new Set<string>()
    for (const file of rootFiles) {
      if (!file.fileName.trim() || !file.role.trim()) {
        setValidationError('canonical 根级文件的文件名与角色均为必填')
        return
      }
      if (fileNames.has(file.fileName)) {
        setValidationError(`canonical 根级文件存在重复文件名：${file.fileName}`)
        return
      }
      fileNames.add(file.fileName)
    }

    // entry_ref 为空的资产行视为“未承接”，不进入提交负载
    const boundAssets = assetRows.filter((row) => row.entryRef.trim() !== '')
    if (boundAssets.length === 0) {
      setValidationError('全局规范资产至少承接 1 项（填写入口引用即视为承接）')
      return
    }
    for (const row of boundAssets) {
      if (!row.role.trim()) {
        setValidationError(`资产 ${row.name} 的角色为必填`)
        return
      }
      if (row.summaryRequired && !row.structuredSummary.trim()) {
        setValidationError(`资产 ${row.name} 需要填写结构化摘要`)
        return
      }
    }

    const trimmedTemplateSource = templateSource.trim()
    onSubmit({
      templateSource: trimmedTemplateSource === '' ? undefined : trimmedTemplateSource,
      canonicalRootFiles: rootFiles.map((f) => ({
        fileName: f.fileName.trim(),
        role: f.role.trim(),
        required: f.required,
      })),
      globalAssetBindings: boundAssets.map((row) => ({
        name: row.name,
        kind: row.kind,
        entryRef: row.entryRef.trim(),
        role: row.role.trim(),
        structuredSummary: row.structuredSummary.trim() === '' ? undefined : row.structuredSummary.trim(),
      })),
    })
  }

  return (
    <form onSubmit={handleSubmit} className="space-y-4">
      {/* template_source — optional 手工字段 */}
      <div className="space-y-1">
        <Label htmlFor="governance-template-source" className="text-xs">
          模板来源（可选）
        </Label>
        <Input
          id="governance-template-source"
          value={templateSource}
          onChange={(e) => setTemplateSource(e.target.value)}
          placeholder="如 project-governance-template-v1；为空表示尚未声明"
          className="h-8 text-xs"
        />
      </div>

      {/* canonical 根级文件 — 可增删行 */}
      <div className="space-y-1.5">
        <div className="flex items-center justify-between">
          <Label className="text-xs">
            canonical 根级文件 <span className="text-destructive">*</span>
          </Label>
          <Button type="button" variant="outline" size="sm" className="h-7 px-2 text-xs" onClick={addRootFile}>
            添加文件
          </Button>
        </div>
        <div className="divide-y rounded-md border">
          {rootFiles.map((file) => (
            <div key={file.key} className="flex flex-wrap items-center gap-2 px-2.5 py-2">
              <Input
                value={file.fileName}
                onChange={(e) => updateRootFile(file.key, { fileName: e.target.value })}
                placeholder="文件名，如 project_rules.md"
                className="h-7 min-w-40 flex-1 text-xs"
                aria-label="根级文件名"
              />
              <Input
                value={file.role}
                onChange={(e) => updateRootFile(file.key, { role: e.target.value })}
                placeholder="角色，如 rules"
                className="h-7 w-36 text-xs"
                aria-label="根级文件角色"
              />
              <label className="flex shrink-0 items-center gap-1 text-xs text-muted-foreground">
                <input
                  type="checkbox"
                  checked={file.required}
                  onChange={(e) => updateRootFile(file.key, { required: e.target.checked })}
                  className="h-3.5 w-3.5"
                />
                必需
              </label>
              <Button
                type="button"
                variant="ghost"
                size="sm"
                className="h-7 px-2 text-xs text-destructive"
                onClick={() => removeRootFile(file.key)}
              >
                移除
              </Button>
            </div>
          ))}
          {rootFiles.length === 0 && (
            <p className="px-3 py-2 text-xs text-muted-foreground">尚未登记任何根级文件</p>
          )}
        </div>
      </div>

      {/* 全局规范资产 — 8 项受控矩阵行 */}
      <div className="space-y-1.5">
        <Label className="text-xs">
          全局规范资产承接 <span className="text-destructive">*</span>
        </Label>
        <p className="text-[10px] text-muted-foreground">
          资产名与分类为受控矩阵；填写入口引用即视为承接，留空表示暂不承接
        </p>
        <div className="divide-y rounded-md border">
          {assetRows.map((row) => (
            <div key={row.name} className="space-y-1.5 px-2.5 py-2">
              <div className="flex flex-wrap items-center gap-2">
                <span className="break-all text-xs font-medium">{row.name}</span>
                <Badge variant="secondary" className="h-4 min-w-4 px-1 text-[10px] font-normal">
                  {row.kind}
                </Badge>
                {row.summaryRequired && (
                  <Badge variant="outline" className="h-4 min-w-4 px-1 text-[10px]">
                    摘要必填
                  </Badge>
                )}
              </div>
              <div className="flex flex-wrap items-center gap-2">
                <Input
                  value={row.entryRef}
                  onChange={(e) => updateAssetRow(row.name, { entryRef: e.target.value })}
                  placeholder="入口引用，如 docs/rules/project_rules.md"
                  className="h-7 min-w-48 flex-1 text-xs"
                  aria-label={`${row.name} 入口引用`}
                />
                <Input
                  value={row.role}
                  onChange={(e) => updateAssetRow(row.name, { role: e.target.value })}
                  placeholder="角色，如 全局规则源"
                  className="h-7 w-40 text-xs"
                  aria-label={`${row.name} 角色`}
                />
              </div>
              {row.summaryRequired && (
                <textarea
                  value={row.structuredSummary}
                  onChange={(e) => updateAssetRow(row.name, { structuredSummary: e.target.value })}
                  placeholder="结构化摘要（供人类回看与 agent 消费的结构化事实，不是正文全文）"
                  rows={2}
                  className="w-full rounded-md border border-input bg-transparent px-2 py-1 text-xs shadow-xs placeholder:text-muted-foreground focus-visible:border-ring focus-visible:ring-[3px] focus-visible:ring-ring/50"
                  aria-label={`${row.name} 结构化摘要`}
                />
              )}
            </div>
          ))}
        </div>
      </div>

      {/* 校验与提交错误 — 停留在表单上下文，保留草稿可重试 */}
      {validationError && <p className="text-xs text-destructive">{validationError}</p>}
      {submitError && <p className="text-xs text-destructive">保存失败：{submitError}</p>}

      <div className="flex gap-2">
        <Button type="submit" size="sm" className="h-8 px-3 text-xs" disabled={submitting}>
          {submitting ? '保存中...' : '保存治理画像'}
        </Button>
        <Button type="button" variant="outline" size="sm" className="h-8 px-3 text-xs" onClick={onCancel} disabled={submitting}>
          取消
        </Button>
      </div>
    </form>
  )
}
