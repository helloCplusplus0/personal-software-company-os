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
import type { CreateReleaseInput, ReleaseStatus } from '../types'

interface ReleaseCreateFormProps {
  moduleId: string
  submitting: boolean
  submitError?: string
  onSubmit: (input: CreateReleaseInput) => void
}

/**
 * ReleaseCreateForm — 版本登记表单
 * §5.7 版本写入承接 CreateRelease（最小字段 version / status / released_at，module_id 由上下文隐式承接）
 * §8.4 状态：idle / dirty / submitting / submit-success / submit-error
 * §8.4 提交失败时停留当前页，保留输入，不得跳出当前 moduleId 上下文
 */
export function ReleaseCreateForm({ moduleId, submitting, submitError, onSubmit }: ReleaseCreateFormProps) {
  // §8.4 草稿状态优先归属于当前页面（§8.6）
  const [version, setVersion] = useState('')
  const [status, setStatus] = useState<ReleaseStatus>('active')
  const [releasedAt, setReleasedAt] = useState(new Date().toISOString().slice(0, 10))

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault()
    if (!version) return
    onSubmit({
      moduleId,
      version,
      status,
      releasedAt: new Date(releasedAt).toISOString(),
    })
  }

  return (
    <form onSubmit={handleSubmit} className="space-y-4 rounded-lg border p-6">
      <div className="space-y-2">
        <Label htmlFor="version">
          版本号 <span className="text-destructive">*</span>
        </Label>
        <Input
          id="version"
          value={version}
          onChange={(e) => setVersion(e.target.value)}
          placeholder="如 1.0.0"
          required
        />
      </div>

      <div className="space-y-2">
        <Label>状态</Label>
        <Select value={status} onValueChange={(v) => setStatus(v as ReleaseStatus)}>
          <SelectTrigger>
            <SelectValue />
          </SelectTrigger>
          <SelectContent>
            <SelectItem value="active">Active</SelectItem>
            <SelectItem value="archived">Archived</SelectItem>
          </SelectContent>
        </Select>
      </div>

      <div className="space-y-2">
        <Label htmlFor="releasedAt">发布日期</Label>
        <Input
          id="releasedAt"
          type="date"
          value={releasedAt}
          onChange={(e) => setReleasedAt(e.target.value)}
        />
      </div>

      {/* §8.4 提交失败时错误显示在表单上下文 */}
      {submitError && (
        <p className="text-sm text-destructive">{submitError}</p>
      )}

      <Button type="submit" disabled={submitting || !version}>
        {submitting ? '提交中...' : '登记版本'}
      </Button>
    </form>
  )
}
