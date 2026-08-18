/**
 * StandardCreatePage — Standard Create
 *
 * phase14-05 §ADDED-3 四页面组件树冻结（创建页，逐字落地）：
 * - 基本信息：name 必填 / description 可选 / status select 默认 draft
 *   （选项仅 draft|active — 对齐 CreateStandard 约束，不含 retired）
 * - 整树 draft state：初始单根 directory name="."（phase14-03 根规范），
 *   StandardTreeEditor 受控组装，提交时整树随 CreateStandard 单次发出
 * - 提交经 use-create-standard（onSuccess 回调位导航到 /standards/:newId）；
 *   后端 R1-R8 校验失败时将错误信息中的节点路径 `(node: /...)` 提取为 errorPaths
 *   传给编辑器高亮（权威层反馈）+ 顶部错误条
 * - 切片纪律：页面不内联 mutation hook（project_rules §2.5）
 */
import { useState } from 'react'
import { useNavigate } from '@tanstack/react-router'
import { useCreateStandard } from '../application/use-create-standard'
import { StandardTreeEditor } from '../components/standard-tree-editor'
import type { DirectoryTreeNode, StandardStatus } from '../types'
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

/** 初始树 = 单根 directory name="."（R1 根规范；role/ref 必空） */
function makeInitialTree(): DirectoryTreeNode {
  return { name: '.', node_type: 'directory', role: '', summary: '', ref: '', children: [] }
}

/**
 * 后端权威层校验错误 → 节点路径集合。
 * 错误信息中节点路径形如 `(node: /docs/phase)`，提取后映射回编辑器命中行高亮。
 */
function extractErrorPaths(message: string): Set<string> {
  const paths = new Set<string>()
  const re = /\(node: (\/[^)]*)\)/g
  let match: RegExpExecArray | null
  while ((match = re.exec(message)) !== null) {
    paths.add(match[1])
  }
  return paths
}

export function StandardCreatePage() {
  const navigate = useNavigate()

  // ---- 表单 state：基本信息 + 整树 draft state（编辑器受控） ----
  const [name, setName] = useState('')
  const [description, setDescription] = useState('')
  const [status, setStatus] = useState<StandardStatus>('draft')
  const [tree, setTree] = useState<DirectoryTreeNode>(makeInitialTree)
  const [submitError, setSubmitError] = useState<string | undefined>(undefined)
  const [errorPaths, setErrorPaths] = useState<Set<string>>(() => new Set())

  // 创建 — onSuccess 回调位承接成功回流（导航到新 Standard 详情页）
  const createMutation = useCreateStandard((standard) => {
    navigate({
      to: '/standards/$standardId',
      params: { standardId: standard.id },
    })
  })

  const handleSubmit = () => {
    setSubmitError(undefined)
    setErrorPaths(new Set())
    if (!name.trim()) {
      setSubmitError('name 不能为空')
      return
    }
    createMutation.mutate(
      {
        name: name.trim(),
        description: description.trim() || undefined,
        status,
        directory_tree: tree,
      },
      {
        // 后端 R1-R8 权威校验失败：错误信息回显 + 节点路径映射编辑器高亮
        onError: (error: Error) => {
          setSubmitError(error.message)
          setErrorPaths(extractErrorPaths(error.message))
        },
      },
    )
  }

  return (
    <div className="space-y-4">
      {/* 标题行 — phase14-05 §ADDED-7 移动端基线 */}
      <div className="flex flex-col gap-2 sm:flex-row sm:items-center sm:justify-between">
        <div className="space-y-1">
          <h1 className="text-xl font-bold">新建 Standard</h1>
          <p className="text-xs text-muted-foreground">基本信息 + 目录结构整树（创建后可继续绑定实体）</p>
        </div>
      </div>

      {/* 顶部错误条 — 前端必填拦截与后端权威层校验共用回显位 */}
      {submitError ? (
        <div className="rounded-lg border border-destructive/50 bg-destructive/10 p-3">
          <p className="text-xs text-destructive">{submitError}</p>
        </div>
      ) : null}

      {/* 基本信息表单 — 移动端单列 grid-cols-1 sm:grid-cols-2 */}
      <div className="grid grid-cols-1 gap-2 sm:grid-cols-2">
        <div className="space-y-1">
          <Label className="text-xs">
            名称<span className="text-destructive">*</span>
          </Label>
          <Input
            className="h-8 text-xs"
            value={name}
            onChange={(e) => setName(e.target.value)}
            placeholder="Standard 名称"
            aria-label="Standard 名称"
          />
        </div>
        <div className="space-y-1">
          <Label className="text-xs">描述（可选）</Label>
          <Input
            className="h-8 text-xs"
            value={description}
            onChange={(e) => setDescription(e.target.value)}
            placeholder="一句话说明该规范资产"
            aria-label="Standard 描述"
          />
        </div>
        <div className="space-y-1">
          <Label className="text-xs">状态</Label>
          <Select value={status} onValueChange={(v) => setStatus(v as StandardStatus)}>
            <SelectTrigger size="sm" className="w-full text-xs" aria-label="Standard 状态">
              <SelectValue />
            </SelectTrigger>
            <SelectContent>
              {/* 创建约束：仅 draft|active，不含 retired（phase14-04 合同） */}
              <SelectItem value="draft">draft</SelectItem>
              <SelectItem value="active">active</SelectItem>
            </SelectContent>
          </Select>
        </div>
      </div>

      {/* 目录结构 — StandardTreeEditor 受控组装（整树 draft state，无网络请求） */}
      <section className="space-y-2 border-t pt-2">
        <p className="text-xs font-medium text-muted-foreground">目录结构</p>
        <StandardTreeEditor tree={tree} onChange={setTree} errorPaths={errorPaths} />
      </section>

      {/* 操作行 — 提交 h-9，移动端 w-full sm:w-auto */}
      <div className="flex flex-col gap-2 border-t pt-2 sm:flex-row">
        <Button
          className="w-full sm:w-auto"
          disabled={createMutation.isPending}
          onClick={handleSubmit}
        >
          {createMutation.isPending ? '提交中...' : '创建'}
        </Button>
        <Button
          variant="outline"
          className="w-full sm:w-auto"
          onClick={() => navigate({ to: '/standards' })}
        >
          取消
        </Button>
      </div>
    </div>
  )
}
