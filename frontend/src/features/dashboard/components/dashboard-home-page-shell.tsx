/**
 * DashboardHomePageShell — Dashboard 页面壳层
 *
 * phase05-13 体验修复后的布局原则：
 * - 一屏概览：主 CTA、资产 stat bar、Current Focus 待办、Recent Activity 在 PC 桌面同屏可见
 * - 信息密度优先：区块间距收紧（space-y-4），区块标题降级（text-base），不再滥用卡片
 * - Recent Activity 限高 + 内部滚动，永不撑开页面（破坏 dashboard 特性的最大元凶）
 * - 主 CTA 内联到标题行，CTA 5-8 有副标题时降级为标题下紧凑条带
 *
 * phase05-10 §3.3 仍冻结四个区块归属与语义：
 *   - dashboard_overview（数字已并入 stat bar，aria-label 保留「系统概览」）
 *   - Current Focus / Next Action（左栏主行动队列）
 *   - Asset Feedback（左栏代表性缺口项；计数已并入 stat bar）
 *   - Recent Activity（右栏独立活动流，限高滚动）
 *   - DashboardPrimaryActionPanel（主 CTA，独立于四区块，内联到标题行）
 *
 * 布局降级（phase05-13 §"PC 与移动浏览器布局降级必须用 Tailwind 响应式"）：
 *   - PC 桌面（md: 以上）：左右两栏，左 CurrentFocus + AssetFeedback，右 RecentActivity
 *   - 移动浏览器（md: 以下）：单列垂直重排，顺序保持
 *     CurrentFocus / StatBar(overview+coverage) / AssetFeedback / RecentActivity
 *
 * 移动端顺序约束（phase05-10 §9.x / phase05-13 spec）：
 *   - CurrentFocus 必须在第一屏优先位置
 */
import type { ReactNode } from 'react'

interface DashboardHomePageShellProps {
  // 主 CTA（内联到标题行右侧；CTA 5-8 时由面板自身降级为标题下条带）
  primaryActionPanel: ReactNode
  // phase06-15：数据主权面板（Export / Backup 入口），独立全宽区块
  sovereigntyPanel?: ReactNode
  // phase12-10：共享语义摘要 / 固定入口说明
  semanticSummaryPanel?: ReactNode
  // 资产概览 + 缺口计数合并条带
  statBar: ReactNode
  // 左栏：待办与代表性缺口
  currentFocusSection: ReactNode
  assetFeedbackSection: ReactNode
  // 右栏：活动流（限高滚动）
  recentActivitySection: ReactNode
}

/**
 * DashboardHomePageShell — 页面壳层。
 *
 * 布局结构（PC 桌面，一屏目标）：
 * ┌─────────────────────────────────────────────┐
 * │ Dashboard               [主 CTA 按钮 →]      │  ← 标题行 + 主CTA内联
 * ├─────────────────────────────────────────────┤
 * │ 数据主权：Export / Backup 入口面板           │  ← phase06-15 新增 sovereigntyPanel
 * ├─────────────────────────────────────────────┤
 * │ [模块 n | 产品 n | ... | 完整 n | 双缺 n]    │  ← stat bar 单行
 * ├──────────────────────────┬──────────────────┤
 * │ Current Focus（待办）     │ Recent Activity  │
 * │ ...紧凑行...              │ ...限高滚动...    │
 * │ Asset Feedback（缺口）    │                  │
 * │  └ Reuse Snapshot         │                  │  ← phase06-15 复用快照子区域
 * │ ...紧凑行...              │                  │
 * └──────────────────────────┴──────────────────┘
 *
 * 移动端（md: 以下）单列顺序（phase05-10 §9.8 字面对齐 + phase06-15 sovereigntyPanel 置顶）：
 * 标题行 → sovereigntyPanel → CurrentFocus → statBar(overview) → AssetFeedback → RecentActivity
 */
export function DashboardHomePageShell({
  primaryActionPanel,
  sovereigntyPanel,
  semanticSummaryPanel,
  statBar,
  currentFocusSection,
  assetFeedbackSection,
  recentActivitySection,
}: DashboardHomePageShellProps) {
  return (
    // 整体收紧：space-y-4（原 space-y-6）
    <div className="space-y-4">
      {/* 标题行 + 主 CTA 内联右侧 */}
      <div className="flex items-center justify-between gap-4">
        <h1 className="text-xl font-bold">Dashboard</h1>
        {primaryActionPanel}
      </div>

      {/*
        phase06-15 §"Dashboard 用户入口落地"：
        sovereigntyPanel 作为独立全宽区块，置于标题行与 stat bar 之间。
        Export / Backup 入口稳定可见，不得做成隐藏路由或仅测试按钮。
      */}
      {sovereigntyPanel}


      {semanticSummaryPanel}
      {/*
        主体网格：扁平化放置 statBar + 四区块，用 CSS grid 显式定位同时满足
        - PC 桌面（md+）：statBar 全宽置顶，左列 CurrentFocus + AssetFeedback，右列 RecentActivity 跨两行
        - PC 桌面（md+）：statBar 全宽置顶，左列 CurrentFocus + AssetFeedback，右列 RecentActivity 跨两行
        - 移动浏览器：按 phase05-10 §9.8 字面顺序 CurrentFocus → statBar(overview) → AssetFeedback → RecentActivity
        移动端通过 order 工具类重排（statBar order-2、CurrentFocus order-1），
        桌面端用 md:col-start / md:row-start 显式定位覆盖 order。

        phase05-13 移动端溢出修复：
        - 必须显式声明 grid-cols-1（= minmax(0,1fr)），否则 md 以下落入隐式 auto 列，
          列宽按内容 max-content 计算（stat bar 10 cell 单行 ≈ 352px 撑大列宽），
          超过 main 内容区并撑出横向滚动，与 header 错位。
        - grid-cols-1 的 minmax(0,1fr) 强制列宽=容器宽度且 min=0，不受内容影响。
        - 各子项额外加 min-w-0，防御内部 flex/grid 内容用 min-width:auto 撑破列。
      */}
      <div className="grid grid-cols-1 gap-4 md:grid-cols-2">
        {/* CurrentFocus — 移动端 order-1（第一屏优先），桌面端左列第二行 */}
        <div className="order-1 min-w-0 md:order-none md:col-start-1 md:row-start-2">
          {currentFocusSection}
        </div>
        {/* statBar — 移动端 order-2（CurrentFocus 之后），桌面端全宽第一行 */}
        <div className="order-2 min-w-0 md:order-none md:col-span-2 md:row-start-1">
          {statBar}
        </div>
        {/* AssetFeedback — 移动端 order-3，桌面端左列第三行 */}
        <div className="order-3 min-w-0 md:order-none md:col-start-1 md:row-start-3">
          {assetFeedbackSection}
        </div>
        {/* RecentActivity — 移动端 order-4（最后），桌面端右列跨第二三行 */}
        <div className="order-4 min-w-0 md:order-none md:col-start-2 md:row-start-2 md:row-span-2">
          {recentActivitySection}
        </div>
      </div>
    </div>
  )
}
