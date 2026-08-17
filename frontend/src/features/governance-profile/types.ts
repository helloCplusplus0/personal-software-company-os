/**
 * governance-profile 切片共享类型。
 *
 * phase13-09：治理画像保存 draft 与表单初始源均为切片内 plain 结构，
 * 不直接暴露 proto Message 类型给表单层（Message 构造需 Schema），
 * 与传输层的转换收敛在 mutation owner 的 PartialMessage 调用点。
 */

/** 治理画像保存 draft — 第一版只包含正式可写集合 */
export interface GovernanceProfileSaveDraft {
  /** 模板来源（optional，为空表示尚未声明） */
  templateSource?: string | undefined
  /** canonical 根级文件绑定集合 */
  canonicalRootFiles: {
    fileName: string
    role: string
    required: boolean
  }[]
  /** 已承接的全局规范资产集合（entry_ref 为空的行不进入提交负载） */
  globalAssetBindings: {
    name: string
    kind: string
    entryRef: string
    role: string
    structuredSummary?: string | undefined
  }[]
}

/** 维护表单初始值来源（已保存画像的可写字段子集；null 表示空态初始化） */
export interface UpdateGovernanceProfileInitialSource {
  templateSource?: string | undefined
  canonicalRootFiles: {
    fileName: string
    role: string
    required: boolean
  }[]
  globalAssetBindings: {
    name: string
    entryRef: string
    role: string
    structuredSummary?: string | undefined
  }[]
}
