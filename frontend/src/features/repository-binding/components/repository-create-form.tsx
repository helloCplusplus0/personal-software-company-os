import { useState } from 'react'
import { Input } from '@/components/ui/input'
import { Label } from '@/components/ui/label'
import { Button } from '@/components/ui/button'
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/select'
import type { CreateRepositoryInput, RepositoryStatus } from '../types'

interface RepositoryCreateFormProps {
  submitting: boolean
  submitError?: string
  onSubmit: (input: CreateRepositoryInput) => void
}

/**
 * RepositoryCreateForm — 仓库创建表单
 * phase04-05 组件树冻结：只承接 name / url / provider / status 录入
 * phase06 draft-first：name + url 为最小人工必填，provider 可留空并默认补 manual，status 默认预填 active
 * phase04-06 提交失败时停留当前页，保留草稿，错误显示在表单上下文
 *
 * 布局降级（phase04-05）：
 * - PC / 移动：单列垂直布局，主动作按钮无需横向滚动即可见
 */
export function RepositoryCreateForm({ submitting, submitError, onSubmit }: RepositoryCreateFormProps) {
  // phase04-06 草稿状态优先归属于当前页面
  const [name, setName] = useState('')
  const [url, setUrl] = useState('')
  const [provider, setProvider] = useState('')
  const [status, setStatus] = useState<RepositoryStatus>('active')
  const isDirty = name !== '' || url !== '' || provider !== ''

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault()
    const trimmedName = name.trim()
    const trimmedURL = url.trim()
    const trimmedProvider = provider.trim()
    if (!trimmedName || !trimmedURL) return
    onSubmit({
      name: trimmedName,
      url: trimmedURL,
      provider: trimmedProvider === '' ? undefined : trimmedProvider,
      status,
    })
  }

  return (
    <form onSubmit={handleSubmit} className="space-y-4 rounded-lg border p-6">
      <div className="space-y-2">
        <Label htmlFor="name">
          仓库名称 <span className="text-destructive">*</span>
        </Label>
        <Input
          id="name"
          value={name}
          onChange={(e) => setName(e.target.value)}
          placeholder="如 psco-backend"
          required
        />
      </div>

      <div className="space-y-2">
        <Label htmlFor="url">
          仓库 URL <span className="text-destructive">*</span>
        </Label>
        <Input
          id="url"
          value={url}
          onChange={(e) => setUrl(e.target.value)}
          placeholder="https://github.com/org/repo"
          required
        />
      </div>

      <div className="space-y-2">
        <Label htmlFor="provider">提供商（可选，默认 manual）</Label>
        <Input
          id="provider"
          value={provider}
          onChange={(e) => setProvider(e.target.value)}
          placeholder="如 github / gitlab"
        />
      </div>

      <div className="space-y-2">
        <Label>状态</Label>
        <Select value={status} onValueChange={(v) => setStatus(v as RepositoryStatus)}>
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
        <Button type="submit" disabled={submitting || !name.trim() || !url.trim()}>
          {submitting ? '提交中...' : '创建仓库'}
        </Button>
      </div>

      {isDirty && (
        <p className="text-xs text-muted-foreground">草稿未保存</p>
      )}
    </form>
  )
}
