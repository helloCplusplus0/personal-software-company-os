/**
 * governance-profile-baseline — 当前项目范式 v1 的前端只读基线表达。
 *
 * phase13-06 冻结：
 *   - backend / database / frontend / proto 顶层目录矩阵只作为当前项目范式 v1
 *     的前端只读基线表达，不来自后端字段、运行时目录扫描或 agent 输入
 *   - docs_workflow_layout 当前项目正式取值固定为 phase/fix/audit/review（phase13-04）
 *   - 8 项全局规范资产矩阵冻结自 phase13-05，name/kind 为受控值，
 *     不得让用户自由新增第 9 项资产
 *
 * 本文件同时承接画像未创建（空态初始化）时的默认预填值：
 * 预填只是编辑起点，用户可改可删，不构成强制模板规则（phase13-04 补充冻结）。
 */

// ============================================================================
// 顶层目录矩阵（当前项目范式 v1 只读基线，非后端字段）
// ============================================================================

export const TOP_LEVEL_DIRECTORY_BASELINE = [
  'backend',
  'database',
  'frontend',
  'proto',
] as const

// ============================================================================
// docs workflow 布局基线
// ============================================================================

/** 当前项目范式 v1 的 docs workflow 正式取值（phase13-04 冻结） */
export const DOCS_WORKFLOW_LAYOUT_BASELINE = 'phase/fix/audit/review'

// ============================================================================
// 8 项全局规范资产冻结矩阵（phase13-05 逐项承接策略）
// ============================================================================

export interface GlobalAssetMatrixEntry {
  /** 资产名（受控值，不得新增第 9 项） */
  name: string
  /** 资产分类（受控展示值） */
  kind: string
  /** 是否必须携带 structured_summary（前 5 项摘要型资产必填） */
  summaryRequired: boolean
}

export const GLOBAL_ASSET_MATRIX: readonly GlobalAssetMatrixEntry[] = [
  { name: 'project_rules.md', kind: 'rules', summaryRequired: true },
  { name: 'TECH_STACK_BASELINE.md', kind: 'tech_baseline', summaryRequired: true },
  { name: 'AGENTS.md', kind: 'agents', summaryRequired: true },
  { name: 'architecture_map.md', kind: 'architecture_map', summaryRequired: true },
  { name: 'plan.md', kind: 'plan', summaryRequired: true },
  { name: 'README.md', kind: 'readme', summaryRequired: false },
  { name: 'global_skills.md', kind: 'skills', summaryRequired: false },
  { name: 'project_skills.md', kind: 'skills', summaryRequired: false },
] as const

// ============================================================================
// 空态初始化默认值（当前项目范式 v1 的根级 canonical 文件集合，phase13-04）
// ============================================================================

export interface CanonicalRootFileDefault {
  fileName: string
  role: string
  required: boolean
}

export const CANONICAL_ROOT_FILE_DEFAULTS: readonly CanonicalRootFileDefault[] = [
  { fileName: '.env', role: 'env', required: true },
  { fileName: 'AGENTS.md', role: 'agents', required: true },
  { fileName: 'architecture_map.md', role: 'architecture_map', required: true },
  { fileName: 'plan.md', role: 'plan', required: true },
  { fileName: 'project_rules.md', role: 'rules', required: true },
  { fileName: 'README.md', role: 'readme', required: true },
  { fileName: 'TECH_STACK_BASELINE.md', role: 'tech_baseline', required: true },
  { fileName: 'global_skills.md', role: 'global_skills', required: false },
  { fileName: 'project_skills.md', role: 'project_skills', required: false },
] as const

// ============================================================================
// 受控枚举的展示映射（只读回看用）
// ============================================================================

import type { PhaseStatus, TrackType } from '@/gen/proto/psco/governance_profile/v1/governance_profile_pb'

export function trackTypeLabel(trackType: TrackType): string {
  switch (trackType) {
    case 1:
      return 'Product Track'
    case 2:
      return 'Durable System Track'
    default:
      return '未声明'
  }
}

export function phaseStatusLabel(status: PhaseStatus): string {
  switch (status) {
    case 1:
      return '已规划'
    case 2:
      return '进行中'
    case 3:
      return '已完成'
    case 4:
      return '受阻'
    default:
      return '未声明'
  }
}
