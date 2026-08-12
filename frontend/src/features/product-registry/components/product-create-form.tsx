import { Input } from '@/components/ui/input'
import { Label } from '@/components/ui/label'
import { Textarea } from '@/components/ui/textarea'
import { Button } from '@/components/ui/button'
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/select'
import type { ProductStatus } from '../types'

interface ProductCreateFormProps {
  /** 表单字段值 */
  name: string
  description: string
  status: ProductStatus
  /** 表单字段变更 */
  onChangeName: (value: string) => void
  onChangeDescription: (value: string) => void
  onChangeStatus: (value: ProductStatus) => void
  /** 是否来自模板预填（控制预填标记展示） */
  isFromTemplate?: boolean
  /** 提交状态 */
  submitting: boolean
  submitError?: string
  onSubmit: (e: React.FormEvent) => void
  isDirty: boolean
}

/**
 * ProductCreateForm — 产品创建表单（props-driven）。
 *
 * phase09-09 重构：组件本地 useState 已回收至 use-product-create-form-state。
 * 表单字段值统一由父组件通过 props 传入，更新通过 onChange 回调。
 *
 * 布局降级（phase04-05）：
 * - PC / 移动：单列垂直布局，主动作按钮无需横向滚动即可见
 */
export function ProductCreateForm({
  name,
  description,
  status,
  onChangeName,
  onChangeDescription,
  onChangeStatus,
  isFromTemplate = false,
  submitting,
  submitError,
  onSubmit,
  isDirty,
}: ProductCreateFormProps) {
  return (
    <form onSubmit={onSubmit} className="space-y-4 rounded-lg border p-6">
      <div className="space-y-2">
        <Label htmlFor="name">
          产品名称 <span className="text-destructive">*</span>
          {isFromTemplate && name !== '' && (
            <span className="text-xs text-muted-foreground ml-2">（来自模板）</span>
          )}
        </Label>
        <Input
          id="name"
          value={name}
          onChange={(e) => onChangeName(e.target.value)}
          placeholder="如 auth-gateway"
          required
        />
      </div>

      <div className="space-y-2">
        <Label htmlFor="description">
          产品描述（可选）
          {isFromTemplate && description !== '' && (
            <span className="text-xs text-muted-foreground ml-2">（来自模板）</span>
          )}
        </Label>
        <Textarea
          id="description"
          value={description}
          onChange={(e) => onChangeDescription(e.target.value)}
          placeholder="简要描述产品职责与边界"
        />
      </div>

      <div className="space-y-2">
        <Label>状态</Label>
        <Select value={status} onValueChange={(v) => onChangeStatus(v as ProductStatus)}>
          <SelectTrigger>
            <SelectValue />
          </SelectTrigger>
          <SelectContent>
            <SelectItem value="active">Active</SelectItem>
            <SelectItem value="archived">Archived</SelectItem>
          </SelectContent>
        </Select>
      </div>

      {/* phase04-06 提交失败时错误显示在表单上下文 */}
      {submitError && (
        <p className="text-sm text-destructive">{submitError}</p>
      )}

      <div className="flex gap-2">
        <Button type="submit" disabled={submitting || !name.trim()}>
          {submitting ? '提交中...' : '创建产品'}
        </Button>
      </div>

      {isDirty && (
        <p className="text-xs text-muted-foreground">草稿未保存</p>
      )}
    </form>
  )
}