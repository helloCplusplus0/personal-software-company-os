/**
 * StandardBindingPanel — Standard 绑定管理区（裁决⑦，仅 StandardDetailPage 挂载）
 *
 * phase14-05 §ADDED-5：
 * - 现有绑定列表：target_type 标签 + target 名称 + role 标签 + note + created_at + 解绑按钮
 * - target 名称解析：四个实体 owning 切片 list query owner 常驻调用（hook 不能条件调用）
 *   + select 数据做 id→name 映射——本文件选型为"四 hook 常驻"方案（最直接实现）；
 *   未命中缓存显示 id 前 8 位
 * - 发起绑定 inline 表单（非弹窗）：target_type → 目标（全量下拉，单用户数量级不新建检索 RPC）
 *   → role 联动禁用（target_type ≠ repository 时 template_source 禁用，八格矩阵 UI 投影）
 *   → note 可选 → 提交；invalid_argument（含 already bound）行内回显
 * - 解绑：window.confirm 确认后四元组调用 use-unbind-standard（note 不参与）
 * - 切片纪律：只调用 application owner hooks，不内联 mutation hook
 */
import { useState } from 'react'
import type { BindingRole, BindingTargetType, StandardBinding } from '../types'
import { useBindStandard } from '../application/use-bind-standard'
import { useUnbindStandard } from '../application/use-unbind-standard'
import { useModuleListRead } from '@/features/module-registry/data/use-module-list-read'
import { useRepositoryListRead } from '@/features/repository-binding/data/use-repository-list-read'
import { useProductListRead } from '@/features/product-registry/data/use-product-list-read'
import { useDecisionListRead } from '@/features/decision-center/data/use-decision-list-read'
import { Badge } from '@/components/ui/badge'
import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import { Label } from '@/components/ui/label'
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/select'

interface StandardBindingPanelProps {
  standardId: string
  bindings: StandardBinding[]
}

/** 全量检索参数 — 模块级常量保持 queryKey 引用稳定（列表页同参数时可直接复用缓存） */
const LIST_SEARCH = { queryText: '', statusFilter: 'all' } as const

const TARGET_TYPE_OPTIONS: BindingTargetType[] = ['repository', 'product', 'decision', 'module']

interface TargetOption {
  id: string
  name: string
}

export function StandardBindingPanel({ standardId, bindings }: StandardBindingPanelProps) {
  // 四实体 list query owner 常驻调用（hook 不能条件调用；数据量单用户级）
  const repositoryList = useRepositoryListRead(LIST_SEARCH)
  const productList = useProductListRead(LIST_SEARCH)
  const decisionList = useDecisionListRead(LIST_SEARCH)
  const moduleList = useModuleListRead(LIST_SEARCH)

  // 目标下拉数据：按 target_type 投影 id→name（decision 用 title 命名）
  const optionsByType: Record<BindingTargetType, TargetOption[]> = {
    repository: (repositoryList.data ?? []).map((r) => ({ id: r.id, name: r.name })),
    product: (productList.data ?? []).map((p) => ({ id: p.id, name: p.name })),
    decision: (decisionList.data ?? []).map((d) => ({ id: d.id, name: d.title })),
    module: (moduleList.data ?? []).map((m) => ({ id: m.id, name: m.name })),
  }
  const listLoadingByType: Record<BindingTargetType, boolean> = {
    repository: repositoryList.isLoading,
    product: productList.isLoading,
    decision: decisionList.isLoading,
    module: moduleList.isLoading,
  }

  /** 绑定行 target 名称解析：命中缓存取 name，未命中显示 id 前 8 位 */
  const resolveTargetName = (binding: StandardBinding): string =>
    optionsByType[binding.target_type].find((o) => o.id === binding.target_id)?.name
      ?? binding.target_id.slice(0, 8)

  // ---- 发起绑定 inline 表单状态 ----
  const [targetType, setTargetType] = useState<BindingTargetType>('repository')
  const [targetId, setTargetId] = useState('')
  const [role, setRole] = useState<BindingRole>('adopts')
  const [note, setNote] = useState('')
  const [formError, setFormError] = useState<string | undefined>(undefined)

  const bindMutation = useBindStandard()
  const unbindMutation = useUnbindStandard()

  /** target_type 切换：清空已选目标；非 repository 时 template_source 不合法，回落 adopts */
  const handleTargetTypeChange = (next: BindingTargetType) => {
    setTargetType(next)
    setTargetId('')
    if (next !== 'repository' && role === 'template_source') setRole('adopts')
  }

  /** 提交绑定 — invalid_argument（含 already bound）经 owner 归一化后行内回显 */
  const handleBind = () => {
    setFormError(undefined)
    if (!targetId) {
      setFormError('请选择绑定目标')
      return
    }
    bindMutation.mutate(
      { standard_id: standardId, form: { target_type: targetType, target_id: targetId, role, note } },
      {
        onSuccess: () => {
          setTargetId('')
          setNote('')
        },
        onError: (error: Error) => setFormError(error.message),
      },
    )
  }

  /** 解绑 — window.confirm 确认后四元组调用（note 不参与） */
  const handleUnbind = (binding: StandardBinding) => {
    if (!window.confirm(`确认解除与 ${resolveTargetName(binding)} 的绑定？`)) return
    unbindMutation.mutate({
      standard_id: standardId,
      target_type: binding.target_type,
      target_id: binding.target_id,
      role: binding.role,
    })
  }

  const targetOptions = optionsByType[targetType]

  return (
    <div className="space-y-2">
      <p className="text-xs font-medium text-muted-foreground">绑定管理</p>

      {/* 现有绑定列表 — flex flex-wrap 行布局 */}
      {bindings.length === 0 ? (
        <p className="text-xs text-muted-foreground">暂无绑定</p>
      ) : (
        <div className="divide-y">
          {bindings.map((binding) => (
            <div key={binding.id} className="flex flex-wrap items-center gap-x-2 gap-y-1 py-1.5">
              <Badge variant="outline">{binding.target_type}</Badge>
              <span className="min-w-0 truncate text-xs font-medium">{resolveTargetName(binding)}</span>
              <Badge variant="secondary">{binding.role}</Badge>
              {binding.note ? (
                <span className="min-w-0 max-w-xs truncate text-xs text-muted-foreground">
                  {binding.note}
                </span>
              ) : null}
              <span className="text-xs text-muted-foreground">
                {new Date(binding.created_at).toLocaleString()}
              </span>
              <Button
                type="button"
                variant="outline"
                className="h-7 px-2 text-xs"
                disabled={unbindMutation.isPending}
                onClick={() => handleUnbind(binding)}
              >
                解绑
              </Button>
            </div>
          ))}
        </div>
      )}

      {/* 发起绑定 inline 表单 — 移动端单列 grid-cols-1 sm:grid-cols-2 */}
      <div className="border-t pt-2">
        <div className="grid grid-cols-1 gap-2 sm:grid-cols-2">
          <div className="space-y-1">
            <Label className="text-xs">目标类型</Label>
            <Select value={targetType} onValueChange={(v) => handleTargetTypeChange(v as BindingTargetType)}>
              <SelectTrigger size="sm" className="w-full text-xs">
                <SelectValue />
              </SelectTrigger>
              <SelectContent>
                {TARGET_TYPE_OPTIONS.map((t) => (
                  <SelectItem key={t} value={t}>
                    {t}
                  </SelectItem>
                ))}
              </SelectContent>
            </Select>
          </div>

          <div className="space-y-1">
            <Label className="text-xs">绑定目标</Label>
            {listLoadingByType[targetType] ? (
              <p className="text-xs text-muted-foreground">目标列表加载中...</p>
            ) : targetOptions.length === 0 ? (
              <p className="text-xs text-muted-foreground">该类型暂无可选目标</p>
            ) : (
              <Select value={targetId} onValueChange={setTargetId}>
                <SelectTrigger size="sm" className="w-full text-xs">
                  <SelectValue placeholder="选择目标" />
                </SelectTrigger>
                <SelectContent>
                  {targetOptions.map((o) => (
                    <SelectItem key={o.id} value={o.id}>
                      {o.name}
                    </SelectItem>
                  ))}
                </SelectContent>
              </Select>
            )}
          </div>

          <div className="space-y-1">
            <Label className="text-xs">绑定角色</Label>
            <Select value={role} onValueChange={(v) => setRole(v as BindingRole)}>
              <SelectTrigger size="sm" className="w-full text-xs">
                <SelectValue />
              </SelectTrigger>
              <SelectContent>
                {/* 八格矩阵联动：target_type ≠ repository 时 template_source 禁用 */}
                <SelectItem value="adopts">adopts</SelectItem>
                <SelectItem value="template_source" disabled={targetType !== 'repository'}>
                  template_source
                </SelectItem>
              </SelectContent>
            </Select>
          </div>

          <div className="space-y-1">
            <Label className="text-xs">备注（可选）</Label>
            <Input
              className="h-8 text-xs"
              value={note}
              onChange={(e) => setNote(e.target.value)}
              placeholder="绑定说明"
            />
          </div>
        </div>

        {/* invalid_argument（含 already bound）行内回显 */}
        {formError ? <p className="mt-2 text-xs text-destructive">{formError}</p> : null}

        <Button
          type="button"
          className="mt-2 h-7 px-2 text-xs"
          disabled={!targetId || bindMutation.isPending}
          onClick={handleBind}
        >
          {bindMutation.isPending ? '提交中...' : '发起绑定'}
        </Button>
      </div>
    </div>
  )
}
