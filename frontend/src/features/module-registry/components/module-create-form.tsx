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
import type { CreateModuleInput, ModuleStatus } from '../types'

interface ModuleCreateFormProps {
  submitting: boolean
  submitError?: string
  onSubmit: (input: CreateModuleInput) => void
}

/**
 * ModuleCreateForm — 模块创建表单
 * §5.7 创建写入承接 CreateModule（phase06 draft-first：name 为最小人工必填）
 * §8.4 草稿状态：idle / dirty
 * §8.4 提交失败时停留当前页，保留草稿，错误显示在表单上下文
 */
export function ModuleCreateForm({ submitting, submitError, onSubmit }: ModuleCreateFormProps) {
  // §8.4 草稿状态优先归属于当前页面（§8.6）
  const [name, setName] = useState('')
  const [description, setDescription] = useState('')
  const [status, setStatus] = useState<ModuleStatus>('active')
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
          模块名称 <span className="text-destructive">*</span>
        </Label>
        <Input
          id="name"
          value={name}
          onChange={(e) => setName(e.target.value)}
          placeholder="如 auth-service"
          required
        />
      </div>

      <div className="space-y-2">
        <Label htmlFor="description">模块描述（可选）</Label>
        <Textarea
          id="description"
          value={description}
          onChange={(e) => setDescription(e.target.value)}
          placeholder="简要描述模块职责"
        />
      </div>

      <div className="space-y-2">
        <Label>状态</Label>
        <Select value={status} onValueChange={(v) => setStatus(v as ModuleStatus)}>
          <SelectTrigger>
            <SelectValue />
          </SelectTrigger>
          <SelectContent>
            <SelectItem value="active">Active</SelectItem>
            <SelectItem value="archived">Archived</SelectItem>
          </SelectContent>
        </Select>
      </div>

      {/* §8.4 提交失败时错误显示在表单上下文 */}
      {submitError && (
        <p className="text-sm text-destructive">{submitError}</p>
      )}

      <div className="flex gap-2">
        <Button type="submit" disabled={submitting || !name.trim()}>
          {submitting ? '提交中...' : '创建模块'}
        </Button>
      </div>

      {isDirty && (
        <p className="text-xs text-muted-foreground">草稿未保存</p>
      )}
    </form>
  )
}
