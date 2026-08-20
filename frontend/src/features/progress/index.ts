/**
 * frontend/src/features/progress/ — Progress 切片受控导出入口（barrel）。
 *
 * phase15-05 §"切片结构必须冻结"：仅导出 ProgressSection
 * （Repository detail 挂载位唯一出口）；不导出内部组件 / query / mutation
 * owner，防止页面在切片外散装拼装写路径（project_rules §2.5）。
 */
export { ProgressSection } from './components/progress-section'
