/**
 * StandardTreeView — Standard 目录树只读展示组件
 *
 * phase14-05 §ADDED-4（裁决⑥只读侧）与 §ADDED-6（readonly-summary 紧凑树规格）：
 * - 两种模式：常规（text-sm，StandardDetailPage 主树）与 compact（text-xs + divide-y，
 *   Repository detail 只读摘要挂载）
 * - 纯只读展示，无任何交互操作入口；min-w-0 truncate 防横向溢出
 * - 根节点为 directory 且 name 固定渲染为 "."（phase14-03 根规范）
 */
import type { DirectoryTreeNode } from '../types'
import { Badge } from '@/components/ui/badge'

interface StandardTreeViewProps {
  tree: DirectoryTreeNode | null
  /** compact 模式：text-xs + divide-y + 更紧凑密度（readonly-summary 场景） */
  compact?: boolean
}

interface TreeNodeRowProps {
  node: DirectoryTreeNode
  depth: number
  compact: boolean
}

/**
 * ref 展示：https:// 开头渲染为外部链接；其余（仓库内路径 / 开头）渲染为链接色文本
 * —— 对齐任务规格"小字链接色文本或链接"
 */
function TreeRefText({ ref: nodeRef, compact }: { ref: string; compact: boolean }) {
  const sizeCls = compact ? 'text-[10px]' : 'text-xs'
  if (nodeRef.startsWith('https://')) {
    return (
      <a
        href={nodeRef}
        target="_blank"
        rel="noreferrer"
        className={`min-w-0 truncate text-primary hover:underline ${sizeCls}`}
      >
        {nodeRef}
      </a>
    )
  }
  return <span className={`min-w-0 truncate text-primary ${sizeCls}`}>{nodeRef}</span>
}

/** 单节点行：directory 只显示名称；file 附带 role 标签 / summary / ref */
function TreeNodeRow({ node, depth, compact }: TreeNodeRowProps) {
  const isRoot = depth === 0
  const isDirectory = node.node_type === 'directory'

  return (
    <div className="min-w-0">
      <div className={`flex min-w-0 items-center gap-1.5 ${compact ? 'py-0.5' : 'py-1'}`}>
        {isDirectory ? (
          // directory 节点：名称 + 可选 summary（裁决①结构化摘要承接位——编辑器允许
          // 编辑 directory summary，有值即展示，避免"编辑了却看不到"的断链）
          <>
            <span
              className={`min-w-0 truncate font-medium ${isRoot ? 'text-muted-foreground' : ''}`}
            >
              {isRoot ? '.' : node.name}
            </span>
            {node.summary ? (
              <span className="min-w-0 flex-1 truncate text-muted-foreground">{node.summary}</span>
            ) : null}
          </>
        ) : (
          <>
            <span className="min-w-0 shrink-0 truncate">{node.name}</span>
            {node.role ? (
              <Badge variant="outline" className="shrink-0 px-1.5">
                {node.role}
              </Badge>
            ) : null}
            {node.summary ? (
              <span className="min-w-0 flex-1 truncate text-muted-foreground">{node.summary}</span>
            ) : null}
            {node.ref ? <TreeRefText ref={node.ref} compact={compact} /> : null}
          </>
        )}
      </div>
      {node.children.length > 0 ? (
        // 每层缩进 pl-4；compact 模式兄弟行间 divide-y
        <div className={`min-w-0 pl-4 ${compact ? 'divide-y' : ''}`}>
          {node.children.map((child, index) => (
            <TreeNodeRow
              key={`${child.name}-${index}`}
              node={child}
              depth={depth + 1}
              compact={compact}
            />
          ))}
        </div>
      ) : null}
    </div>
  )
}

export function StandardTreeView({ tree, compact = false }: StandardTreeViewProps) {
  // 空树判定：null 或根无 children（phase14-05 §ADDED-6 空态规格）
  if (!tree || tree.children.length === 0) {
    return <p className="text-xs text-muted-foreground">暂无目录结构</p>
  }
  return (
    <div className={`min-w-0 ${compact ? 'text-xs divide-y' : 'text-sm'}`}>
      <TreeNodeRow node={tree} depth={0} compact={compact} />
    </div>
  )
}
