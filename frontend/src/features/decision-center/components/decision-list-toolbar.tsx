/**
 * DecisionListToolbar — 列表搜索工具栏
 *
 * §9.1 承接搜索输入、状态筛选与进入 Decision Create 的入口。
 * 搜索参数冻结到路由搜索参数层，工具栏只承接用户输入并回调 onChange。
 *
 * 与 module-registry/components/module-list-toolbar.tsx 同构。
 */
import { Input } from '@/components/ui/input'
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/select'
import type { DecisionStatus } from '../types'

interface DecisionListToolbarProps {
  queryText: string
  statusFilter: DecisionStatus | 'all'
  onChange: (queryText: string, statusFilter: DecisionStatus | 'all') => void
}

/** 状态筛选选项 — 对齐 §5.6 冻结的 status 枚举 */
const STATUS_OPTIONS: { value: DecisionStatus | 'all'; label: string }[] = [
  { value: 'all', label: '全部状态' },
  { value: 'proposed', label: 'Proposed' },
  { value: 'active', label: 'Active' },
  { value: 'superseded', label: 'Superseded' },
  { value: 'archived', label: 'Archived' },
]

export function DecisionListToolbar({ queryText, statusFilter, onChange }: DecisionListToolbarProps) {
  return (
    <div className="flex flex-col gap-3 sm:flex-row sm:items-center">
      <Input
        placeholder="搜索决策标题..."
        value={queryText}
        onChange={(e) => onChange(e.target.value, statusFilter)}
        className="sm:max-w-xs"
      />
      <Select
        value={statusFilter}
        onValueChange={(value) => onChange(queryText, value as DecisionStatus | 'all')}
      >
        <SelectTrigger className="sm:w-40">
          <SelectValue />
        </SelectTrigger>
        <SelectContent>
          {STATUS_OPTIONS.map((opt) => (
            <SelectItem key={opt.value} value={opt.value}>
              {opt.label}
            </SelectItem>
          ))}
        </SelectContent>
      </Select>
    </div>
  )
}
