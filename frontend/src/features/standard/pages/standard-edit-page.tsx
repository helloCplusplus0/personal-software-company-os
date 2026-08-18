/**
 * StandardEditPage — Standard Edit
 *
 * phase14-05 §ADDED-3 四页面组件树冻结（编辑页，逐字落地）：
 * - use-standard-detail-read 预填（加载完成前显示加载态）；表单：
 *   name / description / status（draft|active|retired 全量可选）+
 *   change_summary 必填输入 + 整树 draft state（预填全树深拷贝）
 * - 保存经 use-update-standard（整树原子替换 + change_summary；
 *   optional name/description/status 全量透传 — owner 未设置即不变更，全量传入即以表单为准）
 * - 成功 → 详情页；后端 R1-R8 校验失败同 create 页 errorPaths 映射 + 顶部错误条
 * - 切片纪律：页面不内联 mutation hook（project_rules §2.5）
 */
import { useEffect, useState } from 'react'
import { Link, useNavigate } from '@tanstack/react-router'
import { useStandardDetailRead } from '../data/use-standard-detail-read'
import { useUpdateStandard } from '../application/use-update-standard'
import { StandardTreeEditor } from '../components/standard-tree-editor'
import type { DirectoryTreeNode, StandardStatus } from '../types'
import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import { Label } from '@/components/ui/label'
import { Skeleton } from '@/components/ui/skeleton'
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/select'

/** 兜底初始树 = 单根 directory name="."（理论不触发：detail 必带 directory_tree） */
function makeInitialTree(): DirectoryTreeNode {
  return { name: '.', node_type: 'directory', role: '', summary: '', ref: '', children: [] }
}

/**
 * 后端权威层校验错误 → 节点路径集合（与 create 页同规则）。
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

interface StandardEditPageProps {
  /** 路由文件取参传入（/standards/:standardId/edit） */
  standardId: string
}

export function StandardEditPage({ standardId }: StandardEditPageProps) {
  const navigate = useNavigate()
  const detailQuery = useStandardDetailRead(standardId)
  const updateMutation = useUpdateStandard()

  // ---- 表单 state：基本信息 + change_summary + 整树 draft state ----
  const [initialized, setInitialized] = useState(false)
  const [name, setName] = useState('')
  const [description, setDescription] = useState('')
  const [status, setStatus] = useState<StandardStatus>('draft')
  const [changeSummary, setChangeSummary] = useState('')
  const [tree, setTree] = useState<DirectoryTreeNode>(makeInitialTree)
  const [submitError, setSubmitError] = useState<string | undefined>(undefined)
  const [errorPaths, setErrorPaths] = useState<Set<string>>(() => new Set())

  // 预填：detail 读取完成后一次性初始化表单（整树深拷贝，不污染 query 缓存对象）
  useEffect(() => {
    if (!detailQuery.data || initialized) return
    const standard = detailQuery.data.standard
    setName(standard.name)
    setDescription(standard.description)
    setStatus(standard.status)
    setTree(standard.directory_tree ? structuredClone(standard.directory_tree) : makeInitialTree())
    setInitialized(true)
  }, [detailQuery.data, initialized])

  /** 返回详情页 — 取消与保存成功共用回流目标 */
  const goDetail = () => {
    navigate({
      to: '/standards/$standardId',
      params: { standardId },
    })
  }

  const handleSave = () => {
    setSubmitError(undefined)
    setErrorPaths(new Set())
    if (!name.trim()) {
      setSubmitError('name 不能为空')
      return
    }
    if (!changeSummary.trim()) {
      setSubmitError('change_summary 不能为空')
      return
    }
    updateMutation.mutate(
      {
        standard_id: standardId,
        // optional 三字段全量透传：表单值即最终值（owner 未设置才不变更）
        name: name.trim(),
        description: description.trim(),
        status,
        directory_tree: tree,
        change_summary: changeSummary.trim(),
      },
      {
        onSuccess: () => goDetail(),
        // 后端 R1-R8 权威校验失败：错误信息回显 + 节点路径映射编辑器高亮
        onError: (error: Error) => {
          setSubmitError(error.message)
          setErrorPaths(extractErrorPaths(error.message))
        },
      },
    )
  }

  // 错误态 — 含 not found
  if (detailQuery.isError) {
    return (
      <div className="space-y-4">
        <div className="rounded-lg border border-destructive/50 bg-destructive/10 p-3">
          <p className="text-xs text-destructive">
            {(detailQuery.error as Error).message === 'standard not found'
              ? '该 Standard 不存在或已被删除'
              : `详情读取失败：${(detailQuery.error as Error).message}`}
          </p>
        </div>
        <Button variant="outline" size="sm" asChild>
          <Link to="/standards">返回列表</Link>
        </Button>
      </div>
    )
  }

  // 加载态 — 预填完成前禁用表单（显示加载态）
  if (detailQuery.isLoading || !detailQuery.data || !initialized) {
    return (
      <div className="space-y-4">
        <Skeleton className="h-24 w-full" />
        <Skeleton className="h-48 w-full" />
        <Skeleton className="h-16 w-full" />
      </div>
    )
  }

  return (
    <div className="space-y-4">
      {/* 标题行 — phase14-05 §ADDED-7 移动端基线 */}
      <div className="flex flex-col gap-2 sm:flex-row sm:items-center sm:justify-between">
        <div className="min-w-0 space-y-1">
          <h1 className="min-w-0 truncate text-xl font-bold">编辑 {name || 'Standard'}</h1>
          <p className="text-xs text-muted-foreground">
            整树原子替换 + change_summary 留痕（保存后生成一条 revision）
          </p>
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
          <Label className="text-xs">描述</Label>
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
              {/* 编辑全量可选：draft | active | retired（Retire 即经此切换） */}
              <SelectItem value="draft">draft</SelectItem>
              <SelectItem value="active">active</SelectItem>
              <SelectItem value="retired">retired</SelectItem>
            </SelectContent>
          </Select>
        </div>
        <div className="space-y-1">
          {/* change_summary 必填 — UpdateStandard 合同约束 */}
          <Label className="text-xs">
            变更摘要<span className="text-destructive">*</span>
          </Label>
          <Input
            className="h-8 text-xs"
            value={changeSummary}
            onChange={(e) => setChangeSummary(e.target.value)}
            placeholder="本次变更摘要（必填，将记入 revision）"
            aria-label="变更摘要"
          />
        </div>
      </div>

      {/* 目录结构 — StandardTreeEditor 预填全树（深拷贝 draft，受控组装） */}
      <section className="space-y-2 border-t pt-2">
        <p className="text-xs font-medium text-muted-foreground">目录结构</p>
        <StandardTreeEditor tree={tree} onChange={setTree} errorPaths={errorPaths} />
      </section>

      {/* 操作行 — 保存 h-9，移动端 w-full sm:w-auto */}
      <div className="flex flex-col gap-2 border-t pt-2 sm:flex-row">
        <Button
          className="w-full sm:w-auto"
          disabled={updateMutation.isPending}
          onClick={handleSave}
        >
          {updateMutation.isPending ? '保存中...' : '保存'}
        </Button>
        <Button
          variant="outline"
          className="w-full sm:w-auto"
          disabled={updateMutation.isPending}
          onClick={goDetail}
        >
          取消
        </Button>
      </div>
    </div>
  )
}
