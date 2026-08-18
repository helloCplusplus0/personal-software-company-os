/**
 * frontend/src/features/standard/ — Standard 切片受控导出入口（barrel）。
 *
 * phase14-05 §"切片结构必须冻结"：仅导出 4 页面与 StandardReadonlySummary
 * （Repository detail 挂载位）；不导出其他内部组件 / query / mutation owner，
 * 防止页面在切片外散装拼装写路径（project_rules §2.5）。
 */
export * from './pages/standard-list-page'
export * from './pages/standard-detail-page'
export * from './pages/standard-create-page'
export * from './pages/standard-edit-page'
export { StandardReadonlySummary } from './components/standard-readonly-summary'
