import { Input } from '@/components/ui/input'
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/select'
import type { ModuleStatus } from '../types'

interface ModuleListToolbarProps {
  queryText: string
  statusFilter: ModuleStatus | 'all'
  onChange: (queryText: string, statusFilter: ModuleStatus | 'all') => void
}

/**
 * ModuleListToolbar — 列表工具栏
 * §8.3 承接筛选入口与创建入口
 * §8.4 查询条件冻结到路由搜索参数层
 */
export function ModuleListToolbar({ queryText, statusFilter, onChange }: ModuleListToolbarProps) {
  return (
    <div className="flex flex-col gap-2 sm:flex-row sm:items-center">
      <Input
        placeholder="搜索模块名称或描述..."
        value={queryText}
        onChange={(e) => onChange(e.target.value, statusFilter)}
        className="sm:max-w-xs"
      />
      <Select
        value={statusFilter}
        onValueChange={(v) => onChange(queryText, v as ModuleStatus | 'all')}
      >
        <SelectTrigger className="sm:w-40">
          <SelectValue placeholder="状态筛选" />
        </SelectTrigger>
        <SelectContent>
          <SelectItem value="all">全部状态</SelectItem>
          <SelectItem value="active">Active</SelectItem>
          <SelectItem value="archived">Archived</SelectItem>
        </SelectContent>
      </Select>
    </div>
  )
}
