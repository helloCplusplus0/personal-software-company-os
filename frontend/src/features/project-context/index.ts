/**
 * frontend/src/features/project-context/ — Web 唯一跨切片共享只读入口。
 *
 * phase12-09：受控导出入口，只暴露共享只读 query、adapter 与语义来源。
 * 不承接写路径、页面私有状态或第二套 canonical facts。
 */
export { useProjectContextRead } from './data/use-project-context-read'
export type { ProjectContext, UseProjectContextRead } from './data/use-project-context-read'
export { toEntryLocationView, toEntryLocationViews } from './data/entry-location-view-model'
export type { EntryLocationView } from './data/entry-location-view-model'
export {
  PRODUCT_SEMANTIC_LABEL,
  REPOSITORY_SEMANTIC_LABEL,
  MODULE_SEMANTIC_LABEL,
  DECISION_SEMANTIC_LABEL,
  ENTITY_SEMANTIC_LABEL_MAP,
  RULE_ENTRY_LABEL,
  PHASE_ENTRY_LABEL,
  BOUNDARY_ENTRY_LABEL,
} from './data/shared-semantic-constants'