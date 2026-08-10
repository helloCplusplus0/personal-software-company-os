import { useState } from 'react'
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
import type { CreateProductInput, ProductStatus } from '../types'

interface ProductCreateFormProps {
  submitting: boolean
  submitError?: string
  onSubmit: (input: CreateProductInput) => void
}

/**
 * ProductCreateForm — 产品创建表单
 * phase04-05 组件树冻结：只承接 name / description / status 录入
 * phase06 draft-first：name 为最小人工必填，description 可留空，status 默认预填 active
 * phase04-06 提交失败时停留当前页，保留草稿，错误显示在表单上下文
 *
 * 布局降级（phase04-05）：
 * - PC / 移动：单列垂直布局，主动作按钮无需横向滚动即可见
 */
export function ProductCreateForm({ submitting, submitError, onSubmit }: ProductCreateFormProps) {
  // phase04-06 草稿状态优先归属于当前页面
  const [name, setName] = useState('')
  const [description, setDescription] = useState('')
  const [status, setStatus] = useState<ProductStatus>('active')
  const isDirty = name !== '' || description !== ''

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault()
    const trimmedName = name.trim()
    const trimmedDescription = description.trim()
    if (!trimmedName) return
    onSubmit({
      name: trimmedName,
      description: trimmedDescription === '' ? undefined : trimmedDescription,
      status,
    })
  }

  return (
    <form onSubmit={handleSubmit} className="space-y-4 rounded-lg border p-6">
      <div className="space-y-2">
        <Label htmlFor="name">
          产品名称 <span className="text-destructive">*</span>
        </Label>
        <Input
          id="name"
          value={name}
          onChange={(e) => setName(e.target.value)}
          placeholder="如 auth-gateway"
          required
        />
      </div>

      <div className="space-y-2">
        <Label htmlFor="description">产品描述（可选）</Label>
        <Textarea
          id="description"
          value={description}
          onChange={(e) => setDescription(e.target.value)}
          placeholder="简要描述产品职责与边界"
        />
      </div>

      <div className="space-y-2">
        <Label>状态</Label>
        <Select value={status} onValueChange={(v) => setStatus(v as ProductStatus)}>
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
