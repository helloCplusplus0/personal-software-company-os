// 四实体冻结语义标签（phase12-02 冻结）
// 唯一跨切片共享语义常量源，回收重复硬编码的四实体解释
export const PRODUCT_SEMANTIC_LABEL = "经营目标与交付容器"
export const REPOSITORY_SEMANTIC_LABEL = "代码仓库身份对象与项目锚点"
export const MODULE_SEMANTIC_LABEL = "可复用能力资产"
export const DECISION_SEMANTIC_LABEL = "规则、约束、选择与依据的索引对象"
export const DECISION_SEMANTIC_CORE = "规则、约束、选择与依据"

// 实体类型-标签映射
export const ENTITY_SEMANTIC_LABEL_MAP: Record<string, string> = {
  Product: PRODUCT_SEMANTIC_LABEL,
  Repository: REPOSITORY_SEMANTIC_LABEL,
  Module: MODULE_SEMANTIC_LABEL,
  Decision: DECISION_SEMANTIC_LABEL,
}

// phase13-09：phase12 project-context 叙事入口常量（rules / phases / boundaries）
// 已随“项目上下文”前端区块一并退出，不再保留
