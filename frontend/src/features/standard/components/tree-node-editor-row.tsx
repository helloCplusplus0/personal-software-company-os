/**
 * TreeNodeEditorRow — 树编辑器单节点行
 *
 * phase14-05 §ADDED-4（裁决⑥）节点行规格：
 * - 一行内 5 输入 name / node_type / role / summary / ref + 操作组 4 按钮
 * - 移动端折行：grid-cols-1 → sm:grid-cols-2 → lg 全展开，操作组 col-span-full flex-wrap 独立成行
 * - 禁用态规则表逐行落地（根只读 / file 无子节点 / 含 children 禁切 file /
 *   第 6 层只允许 file / 同层首尾禁上移下移 / 删除含子节点确认弹窗）
 * - node_type 切换：file→directory 直接允许（children 置 []）；directory→file 仅当无 children
 * - 校验反馈双层模型之前端轻量层：行内警告 text-xs text-destructive，不阻断输入
 * - 后端权威层：errorPaths 命中的节点行 ring-1 ring-destructive 高亮（R1-R8 响应错误路径映射）
 */
import type { DirectoryTreeNode, NodeType } from '../types'
import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/select'

export interface TreeNodeEditorRowProps {
  node: DirectoryTreeNode
  /** 根为 depth 0（第 1 层）；depth 5 即第 6 层（最大深度） */
  depth: number
  isFirst: boolean
  isLast: boolean
  /** 节点路径：根 "/"，子节点 "/{name}" 逐级拼接（与后端错误路径对齐） */
  path: string
  onChange: (node: DirectoryTreeNode) => void
  onDelete: () => void
  onMoveUp: () => void
  onMoveDown: () => void
  /** 后端校验错误命中的节点路径集合（如 "/docs/phase"）；命中行高亮 */
  errorPaths?: Set<string>
}

/** 添加子节点默认值 — phase14-05 裁决⑥：默认 directory、字段全空（editor 添加根级节点复用） */
export function makeDefaultChildNode(): DirectoryTreeNode {
  return { name: '', node_type: 'directory', role: '', summary: '', ref: '', children: [] }
}

/** name 前端轻量校验（R4 字符集 + 长度 64 + 非空）——仅警告不阻断 */
function validateName(name: string): string | undefined {
  if (name === '') return '名称不能为空'
  if (name.includes('/')) return '名称不能包含 "/"'
  if (/\s/.test(name)) return '名称不能包含空白字符'
  if (name.length > 64) return '名称不能超过 64 字符'
  if (/[^A-Za-z0-9._-]/.test(name)) return '名称仅允许字母、数字与 . _ -'
  return undefined
}

/** file 节点 role 校验（必填 + 长度 32）——仅警告不阻断 */
function validateRole(role: string | undefined, isFile: boolean): string | undefined {
  if (!isFile) return undefined
  if (!role || role === '') return 'file 节点 role 不能为空'
  if (role.length > 32) return 'role 不能超过 32 字符'
  return undefined
}

/** summary 长度校验（2000）——仅警告不阻断 */
function validateSummary(summary: string | undefined): string | undefined {
  if (summary && summary.length > 2000) return 'summary 不能超过 2000 字符'
  return undefined
}

/** ref 格式校验（非空需以 / 或 https:// 开头）——仅警告不阻断 */
function validateRef(ref: string | undefined): string | undefined {
  if (ref && ref !== '' && !ref.startsWith('/') && !ref.startsWith('https://')) {
    return 'ref 需以 / 或 https:// 开头'
  }
  return undefined
}

export function TreeNodeEditorRow({
  node,
  depth,
  isFirst,
  isLast,
  path,
  onChange,
  onDelete,
  onMoveUp,
  onMoveDown,
  errorPaths,
}: TreeNodeEditorRowProps) {
  const isRoot = depth === 0
  const isFile = node.node_type === 'file'
  const hasChildren = node.children.length > 0
  // 第 6 层（depth 5，根为第 1 层）：只允许 file（R5），禁添加子节点、禁 directory 选项
  const atMaxDepth = depth >= 5

  // ---- 禁用态规则表（phase14-05 §ADDED-4）----
  // 根节点：删除 / 上移 / 下移禁用（无父层语义，R1）
  const canDelete = !isRoot
  const canMoveUp = !isRoot && !isFirst
  const canMoveDown = !isRoot && !isLast
  // file 节点添加子节点禁用（R7）；第 6 层添加子节点禁用（R5）
  const canAddChild = !isFile && !atMaxDepth
  // directory 含 children 时禁切 file（须先清空子节点，R7）；第 6 层禁 directory 选项（R5）
  const fileOptionDisabled = hasChildren
  const directoryOptionDisabled = atMaxDepth

  // ---- 前端轻量校验（根 name 固定 "." 不校验；根不提供 role / ref 输入位）----
  const warnings = [
    !isRoot ? validateName(node.name) : undefined,
    validateRole(node.role, isFile),
    validateSummary(node.summary),
    !isRoot ? validateRef(node.ref) : undefined,
  ].filter((w): w is string => w !== undefined)

  const errorHit = errorPaths?.has(path) ?? false

  /** 字段更新：不可变替换节点后经 onChange 上抛 */
  const patchNode = (patch: Partial<DirectoryTreeNode>) => onChange({ ...node, ...patch })

  /** node_type 切换：file→directory 直接允许；directory→file 仅当无 children（children 置 []） */
  const handleNodeTypeChange = (next: NodeType) => {
    if (next === node.node_type) return
    if (next === 'file' && hasChildren) return // 选项已禁用，此处双保险拦截
    patchNode({ node_type: next, children: [] })
  }

  /** 添加子节点：插入该层 children 末尾，默认 directory 空节点 */
  const handleAddChild = () => {
    if (!canAddChild) return
    patchNode({ children: [...node.children, makeDefaultChildNode()] })
  }

  /** 删除 directory 含 children：确认弹窗提示连带删除全部后代 */
  const handleDelete = () => {
    if (!isFile && hasChildren) {
      if (!window.confirm('该目录含子节点，删除将连带删除全部后代，确认删除？')) return
    }
    onDelete()
  }

  return (
    <div
      className={`grid grid-cols-1 items-center gap-1.5 rounded-md p-2 sm:grid-cols-2 lg:grid-cols-[minmax(0,1fr)_auto_minmax(0,8rem)_minmax(0,1fr)_minmax(0,10rem)] ${
        errorHit ? 'ring-1 ring-destructive' : ''
      }`}
    >
      {/* name — 根固定 "." 只读（R1） */}
      <Input
        className="h-8 text-xs"
        value={isRoot ? '.' : node.name}
        onChange={(e) => patchNode({ name: e.target.value })}
        disabled={isRoot}
        placeholder="节点名称"
        aria-label="节点名称"
      />

      {/* node_type — 根固定 directory 只读（R1） */}
      <Select
        value={node.node_type}
        onValueChange={(v) => handleNodeTypeChange(v as NodeType)}
        disabled={isRoot}
      >
        <SelectTrigger size="sm" className="w-full text-xs" aria-label="节点类型">
          <SelectValue />
        </SelectTrigger>
        <SelectContent>
          <SelectItem value="directory" disabled={directoryOptionDisabled}>
            directory
          </SelectItem>
          <SelectItem value="file" disabled={fileOptionDisabled}>
            file
          </SelectItem>
        </SelectContent>
      </Select>

      {/* role — 根不提供输入位（禁用占位保持行网格稳定） */}
      <Input
        className="h-8 text-xs"
        value={node.role ?? ''}
        onChange={(e) => patchNode({ role: e.target.value })}
        disabled={isRoot}
        placeholder="role"
        aria-label="role"
      />

      {/* summary — 根亦可编辑（phase14-05 禁用表：根 summary 可编辑） */}
      <Input
        className="h-8 text-xs"
        value={node.summary ?? ''}
        onChange={(e) => patchNode({ summary: e.target.value })}
        placeholder="summary"
        aria-label="summary"
      />

      {/* ref — 根不提供输入位（phase14-05 禁用表） */}
      {!isRoot ? (
        <Input
          className="h-8 text-xs"
          value={node.ref ?? ''}
          onChange={(e) => patchNode({ ref: e.target.value })}
          placeholder="/path 或 https://"
          aria-label="ref"
        />
      ) : null}

      {/* 操作组 — 独立成行 flex-wrap（phase14-05 移动端基线） */}
      <div className="col-span-full flex flex-wrap items-center gap-1">
        <Button
          type="button"
          variant="outline"
          className="h-7 px-2 text-xs"
          disabled={!canAddChild}
          onClick={handleAddChild}
        >
          添加子节点
        </Button>
        <Button
          type="button"
          variant="outline"
          className="h-7 px-2 text-xs"
          disabled={!canDelete}
          onClick={handleDelete}
        >
          删除
        </Button>
        <Button
          type="button"
          variant="outline"
          className="h-7 px-2 text-xs"
          disabled={!canMoveUp}
          onClick={onMoveUp}
        >
          上移
        </Button>
        <Button
          type="button"
          variant="outline"
          className="h-7 px-2 text-xs"
          disabled={!canMoveDown}
          onClick={onMoveDown}
        >
          下移
        </Button>
      </div>

      {/* 前端轻量校验行内警告 — 不阻断输入（校验双层模型第一层） */}
      {warnings.length > 0 ? (
        <div className="col-span-full space-y-0.5">
          {warnings.map((w) => (
            <p key={w} className="text-xs text-destructive">
              {w}
            </p>
          ))}
        </div>
      ) : null}
    </div>
  )
}
