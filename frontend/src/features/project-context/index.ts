/**
 * frontend/src/features/project-context/ — 四实体共享语义常量受控导出入口。
 *
 * phase13-09：phase12“项目上下文 / 共享项目上下文”前端区块已按退出规则整体移除，
 * 本切片只保留 phase12-02 冻结的四实体语义标签（仍被各页面语义导语消费），
 * 不再导出 project-context 读取、入口视图或摘要面板。
 */
export {
  PRODUCT_SEMANTIC_LABEL,
  REPOSITORY_SEMANTIC_LABEL,
  MODULE_SEMANTIC_LABEL,
  DECISION_SEMANTIC_LABEL,
  DECISION_SEMANTIC_CORE,
  ENTITY_SEMANTIC_LABEL_MAP,
} from './data/shared-semantic-constants'
