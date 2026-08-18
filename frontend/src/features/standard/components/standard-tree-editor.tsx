/**
 * StandardTreeEditor — 树编辑器外壳与递归渲染
 *
 * phase14-05 §ADDED-4（裁决⑥）：
 * - 受控组件：draft state 由父页面持有，编辑器不持有 useState 存树、不发起任何网络请求
 * - 整树变更经不可变更新（复制路径）后 onChange 上抛；提交时整树随 Create/Update 发出
 * - 工具行：左"添加根级节点"（根 children 追加默认 directory 节点）+ 右节点计数
 * - errorPaths：后端 R1-R8 校验错误命中的节点路径集合，透传至命中行高亮（权威层反馈）
 * - 无拖拽交互（不引入拖拽排序依赖，不绑定任何拖拽事件处理器）
 */
import type { DirectoryTreeNode } from '../types'
import { Button } from '@/components/ui/button'
import { TreeNodeEditorRow, makeDefaultChildNode } from './tree-node-editor-row'

interface StandardTreeEditorProps {
  tree: DirectoryTreeNode
  onChange: (tree: DirectoryTreeNode) => void
  /** 后端校验错误命中的节点路径集合（如 "/docs/phase"），命中行 ring-destructive 高亮 */
  errorPaths?: Set<string>
}

/** 整树节点计数（含根）— 工具行轻量规模提示 */
function countNodes(node: DirectoryTreeNode): number {
  return 1 + node.children.reduce((acc, child) => acc + countNodes(child), 0)
}

/** 子节点路径拼接：根 "/" → "/docs" → "/docs/phase"（与后端错误路径规则一致） */
function childPath(parentPath: string, childName: string): string {
  return `${parentPath === '/' ? '' : parentPath}/${childName}`
}

/** 不可变替换 index 处子节点 */
function replaceChild(node: DirectoryTreeNode, index: number, next: DirectoryTreeNode): DirectoryTreeNode {
  const children = [...node.children]
  children[index] = next
  return { ...node, children }
}

/** 不可变移除 index 处子节点 */
function removeChild(node: DirectoryTreeNode, index: number): DirectoryTreeNode {
  return { ...node, children: node.children.filter((_, i) => i !== index) }
}

/** 不可变交换同层两个子节点（上移 / 下移） */
function swapChildren(node: DirectoryTreeNode, index: number, other: number): DirectoryTreeNode {
  const children = [...node.children]
  ;[children[index], children[other]] = [children[other], children[index]]
  return { ...node, children }
}

interface EditorSubtreeProps {
  node: DirectoryTreeNode
  depth: number
  path: string
  isFirst: boolean
  isLast: boolean
  /** 替换自身（含整棵子树）的回调，由父级递归层以不可变更新实现 */
  onSelfChange: (node: DirectoryTreeNode) => void
  onSelfDelete: () => void
  onSelfMoveUp: () => void
  onSelfMoveDown: () => void
  errorPaths?: Set<string>
}

/** 递归渲染层：渲染自身行 + children 容器（每层缩进 pl-4），并为子行装配兄弟操作回调 */
function EditorSubtree({
  node,
  depth,
  path,
  isFirst,
  isLast,
  onSelfChange,
  onSelfDelete,
  onSelfMoveUp,
  onSelfMoveDown,
  errorPaths,
}: EditorSubtreeProps) {
  return (
    <>
      <TreeNodeEditorRow
        node={node}
        depth={depth}
        isFirst={isFirst}
        isLast={isLast}
        path={path}
        onChange={onSelfChange}
        onDelete={onSelfDelete}
        onMoveUp={onSelfMoveUp}
        onMoveDown={onSelfMoveDown}
        errorPaths={errorPaths}
      />
      {node.children.length > 0 ? (
        <div className="min-w-0 pl-4">
          {node.children.map((child, index) => (
            <EditorSubtree
              // key 必须稳定：若含 child.name，每敲一个字母 key 变化会导致
              // React 卸载重建子树、input 焦点丢失（phase14-10 前置反馈缺陷修复）。
              // 行组件为纯受控（状态全部在父级 draft tree），index key 安全。
              key={index}
              node={child}
              depth={depth + 1}
              path={childPath(path, child.name)}
              isFirst={index === 0}
              isLast={index === node.children.length - 1}
              onSelfChange={(next) => onSelfChange(replaceChild(node, index, next))}
              onSelfDelete={() => onSelfChange(removeChild(node, index))}
              onSelfMoveUp={() => onSelfChange(swapChildren(node, index, index - 1))}
              onSelfMoveDown={() => onSelfChange(swapChildren(node, index, index + 1))}
              errorPaths={errorPaths}
            />
          ))}
        </div>
      ) : null}
    </>
  )
}

export function StandardTreeEditor({ tree, onChange, errorPaths }: StandardTreeEditorProps) {
  /** 添加根级节点：等同对根"添加子节点"（根 children 追加默认 directory 节点） */
  const handleAddRootChild = () => {
    onChange({ ...tree, children: [...tree.children, makeDefaultChildNode()] })
  }

  return (
    <div className="space-y-2">
      {/* 工具行：左添加根级节点 + 右节点计数（轻量规模提示） */}
      <div className="flex items-center justify-between gap-2">
        <Button
          type="button"
          variant="outline"
          className="h-7 px-2 text-xs"
          onClick={handleAddRootChild}
        >
          添加根级节点
        </Button>
        <span className="text-xs text-muted-foreground">节点 {countNodes(tree)} 个</span>
      </div>

      {/* 根行（depth 0，路径 "/"）：行内按 depth===0 禁用删除/上移/下移，兄弟回调为占位 no-op */}
      <EditorSubtree
        node={tree}
        depth={0}
        path="/"
        isFirst
        isLast
        onSelfChange={onChange}
        onSelfDelete={() => {}}
        onSelfMoveUp={() => {}}
        onSelfMoveDown={() => {}}
        errorPaths={errorPaths}
      />
    </div>
  )
}
