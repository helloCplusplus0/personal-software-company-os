/**
 * frontend/src/features/governance-profile/ — 治理画像前端切片受控导出入口。
 *
 * phase13-09：唯一页面消费入口是 GovernanceProfileSection（挂在 Repository detail）。
 * 只暴露该承接区组件，不导出 query / mutation owner 供页面散装拼装。
 */
export { GovernanceProfileSection } from './components/governance-profile-section'
