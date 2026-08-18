/**
 * Standard 切片类型定义
 *
 * 对齐 phase14-05 §"切片结构必须冻结"与 phase14-04 冻结的后端合同（8 RPC）。
 * 前端消费模型从 pb 生成产物单向派生（snake_case 对齐后端）；
 * pb ↔ domain 转换集中在本文件，data / application 层不得散写 pb 映射。
 */
import { create } from '@bufbuild/protobuf'
import { timestampDate } from '@bufbuild/protobuf/wkt'
import {
  StandardStatus as StandardStatusPb,
  NodeType as NodeTypePb,
  BindingTargetType as BindingTargetTypePb,
  BindingRole as BindingRolePb,
  DirectoryTreeNodeSchema,
} from '@/gen/proto/psco/standard/v1/standard_pb'
import type {
  Standard as StandardPb,
  DirectoryTreeNode as DirectoryTreeNodePb,
  StandardBinding as StandardBindingPb,
  StandardRevision as StandardRevisionPb,
} from '@/gen/proto/psco/standard/v1/standard_pb'

/** Standard 状态 — 生命周期 draft / active / retired */
export type StandardStatus = 'draft' | 'active' | 'retired'

/** 目录树节点类型 — directory / file */
export type NodeType = 'directory' | 'file'

/** 绑定目标类型 — phase14-02 八格矩阵四实体 */
export type BindingTargetType = 'repository' | 'product' | 'decision' | 'module'

/** 绑定角色 — template_source 仅 repository 合法（八格矩阵） */
export type BindingRole = 'template_source' | 'adopts'

/** 目录树节点 — 与 directory_tree jsonb 单值映射（phase14-03 冻结 6 字段结构） */
export interface DirectoryTreeNode {
  name: string
  node_type: NodeType
  role?: string
  summary?: string
  ref?: string
  children: DirectoryTreeNode[]
}

/** Standard 核心对象 — 主表投影 + 整树（List/Get/brief 统一消费模型） */
export interface Standard {
  id: string
  name: string
  description: string
  status: StandardStatus
  directory_tree: DirectoryTreeNode | null
  created_at: string
  updated_at: string
}

/** Standard 多态绑定关系 — standard_bindings 表投影 */
export interface StandardBinding {
  id: string
  standard_id: string
  target_type: BindingTargetType
  target_id: string
  role: BindingRole
  note?: string
  created_at: string
}

/** Standard 演进留痕 — standard_revisions 表投影（只追加） */
export interface StandardRevision {
  id: string
  standard_id: string
  change_summary: string
  created_at: string
}

/** 详情读取模型 — GetStandard 响应投影（绑定管理区直接消费） */
export interface StandardDetail {
  standard: Standard
  bindings: StandardBinding[]
}

/** 绑定表单模型 — StandardBindingPanel 发起绑定输入 */
export interface StandardBindingFormModel {
  target_type: BindingTargetType
  target_id: string
  role: BindingRole
  note: string
}

/** CreateStandard 写入参数 — 表单值模型，owner 内转换为 pb 请求 */
export interface CreateStandardInput {
  name: string
  description?: string
  status?: StandardStatus
  directory_tree: DirectoryTreeNode
}

/** UpdateStandard 写入参数 — 整树原子替换 + change_summary 必填；optional 字段未设置即不变更 */
export interface UpdateStandardInput {
  standard_id: string
  name?: string
  description?: string
  status?: StandardStatus
  directory_tree: DirectoryTreeNode
  change_summary: string
}

/** BindStandard 写入参数 — standard_id + 绑定表单模型 */
export interface BindStandardInput {
  standard_id: string
  form: StandardBindingFormModel
}

/** UnbindStandard 写入参数 — 四元组定位，note 不参与 */
export interface UnbindStandardInput {
  standard_id: string
  target_type: BindingTargetType
  target_id: string
  role: BindingRole
}

// ============================================================================
// 枚举映射（pb enum ↔ string union）
// ============================================================================

/** pb → domain：未知值（UNSPECIFIED）防御性落 'draft' */
export function pbToStandardStatus(v: StandardStatusPb): StandardStatus {
  switch (v) {
    case StandardStatusPb.ACTIVE:
      return 'active'
    case StandardStatusPb.RETIRED:
      return 'retired'
    default:
      return 'draft'
  }
}

/** domain → pb */
export function standardStatusToPb(s: StandardStatus): StandardStatusPb {
  switch (s) {
    case 'active':
      return StandardStatusPb.ACTIVE
    case 'retired':
      return StandardStatusPb.RETIRED
    default:
      return StandardStatusPb.DRAFT
  }
}

/** pb → domain：未知值（UNSPECIFIED）防御性落 'directory' */
export function pbToNodeType(v: NodeTypePb): NodeType {
  return v === NodeTypePb.FILE ? 'file' : 'directory'
}

/** domain → pb */
export function nodeTypeToPb(t: NodeType): NodeTypePb {
  return t === 'file' ? NodeTypePb.FILE : NodeTypePb.DIRECTORY
}

/** pb → domain：未知值（UNSPECIFIED）防御性落 'repository'（八格矩阵第一合法目标） */
export function pbToBindingTargetType(v: BindingTargetTypePb): BindingTargetType {
  switch (v) {
    case BindingTargetTypePb.PRODUCT:
      return 'product'
    case BindingTargetTypePb.DECISION:
      return 'decision'
    case BindingTargetTypePb.MODULE:
      return 'module'
    case BindingTargetTypePb.REPOSITORY:
      return 'repository'
    default:
      return 'repository'
  }
}

/** domain → pb */
export function bindingTargetTypeToPb(t: BindingTargetType): BindingTargetTypePb {
  switch (t) {
    case 'product':
      return BindingTargetTypePb.PRODUCT
    case 'decision':
      return BindingTargetTypePb.DECISION
    case 'module':
      return BindingTargetTypePb.MODULE
    default:
      return BindingTargetTypePb.REPOSITORY
  }
}

/** pb → domain：未知值（UNSPECIFIED）防御性落 'adopts'（全目标合法角色） */
export function pbToBindingRole(v: BindingRolePb): BindingRole {
  return v === BindingRolePb.TEMPLATE_SOURCE ? 'template_source' : 'adopts'
}

/** domain → pb */
export function bindingRoleToPb(r: BindingRole): BindingRolePb {
  return r === 'template_source' ? BindingRolePb.TEMPLATE_SOURCE : BindingRolePb.ADOPTS
}

// ============================================================================
// 消息转换器（pb → domain）
// ============================================================================

/** pb 树节点 → domain 树节点（递归整树） */
function toDomainTreeNode(node: DirectoryTreeNodePb): DirectoryTreeNode {
  return {
    name: node.name ?? '',
    node_type: pbToNodeType(node.nodeType),
    role: node.role || undefined,
    summary: node.summary || undefined,
    ref: node.ref || undefined,
    children: (node.children ?? []).map(toDomainTreeNode),
  }
}

/** pb 整树 → domain 整树；缺树（理论不发生）返回 null */
export function pbToTree(node: DirectoryTreeNodePb | undefined): DirectoryTreeNode | null {
  return node ? toDomainTreeNode(node) : null
}

/** pb Standard → domain Standard；缺主体（理论不发生）返回 null */
export function pbToStandard(pb: StandardPb | undefined): Standard | null {
  if (!pb) return null
  return {
    id: pb.id ?? '',
    name: pb.name ?? '',
    description: pb.description ?? '',
    status: pbToStandardStatus(pb.status),
    directory_tree: pbToTree(pb.directoryTree),
    created_at: pb.createdAt ? timestampDate(pb.createdAt).toISOString() : '',
    updated_at: pb.updatedAt ? timestampDate(pb.updatedAt).toISOString() : '',
  }
}

/** pb StandardBinding → domain StandardBinding */
export function pbToBinding(pb: StandardBindingPb): StandardBinding {
  return {
    id: pb.id ?? '',
    standard_id: pb.standardId ?? '',
    target_type: pbToBindingTargetType(pb.targetType),
    target_id: pb.targetId ?? '',
    role: pbToBindingRole(pb.role),
    note: pb.note || undefined,
    created_at: pb.createdAt ? timestampDate(pb.createdAt).toISOString() : '',
  }
}

/** pb StandardRevision → domain StandardRevision */
export function pbToRevision(pb: StandardRevisionPb): StandardRevision {
  return {
    id: pb.id ?? '',
    standard_id: pb.standardId ?? '',
    change_summary: pb.changeSummary ?? '',
    created_at: pb.createdAt ? timestampDate(pb.createdAt).toISOString() : '',
  }
}

// ============================================================================
// 请求组装（domain → pb）
// ============================================================================

/** domain 整树 → pb DirectoryTreeNode（Create / Update 整树组装，递归） */
export function treeToPb(node: DirectoryTreeNode): DirectoryTreeNodePb {
  return create(DirectoryTreeNodeSchema, {
    name: node.name,
    nodeType: nodeTypeToPb(node.node_type),
    role: node.role ?? '',
    summary: node.summary ?? '',
    ref: node.ref ?? '',
    children: node.children.map((child) => treeToPb(child)),
  })
}
