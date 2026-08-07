import { Input } from '@/components/ui/input'
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/select'
import type { RepositoryStatus } from '../types'

interface RepositoryBindingListToolbarProps {
  queryText: string
  statusFilter: RepositoryStatus | 'all'
  onChange: (queryText: string, statusFilter: RepositoryStatus | 'all') => void
}

/**
 * RepositoryBindingListToolbar — 列表工具栏
 * phase04-05 组件树冻结：承接 queryText 搜索输入、statusFilter 状态筛选
 * phase04-06 筛选维度冻结：只冻结 queryText / statusFilter，不引入 providerFilter
 * phase04-06 查询条件冻结到路由搜索参数层
 *
 * 布局降级（phase04-05）：
 * - PC：工具栏横向排列
 * - 移动：单列垂直排列
 */
export function RepositoryBindingListToolbar({ queryText, statusFilter, onChange }: RepositoryBindingListToolbarProps) {
  return (
    <div className="flex flex-col gap-2 sm:flex-row sm:items-center">
      <Input
        placeholder="搜索仓库名称..."
        value={queryText}
        onChange={(e) => onChange(e.target.value, statusFilter)}
        className="sm:max-w-xs"
      />
      <Select
        value={statusFilter}
        onValueChange={(v) => onChange(queryText, v as RepositoryStatus | 'all')}
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
