# Phase05-01 Dashboard 页面边界与信息结构 Spec

## Why

`phase05` 要把 `Dashboard + Feedback` 从 `phase04` 已交付的 `Product / Repository / Binding` 主线之上推进为新的最小闭环，但第一步不能直接进入读模型、合同或前端实现，而必须先把 `Dashboard Home` 的页面边界、区块归属、首页入口语义与 `PC / 移动浏览器` 下的信息层级冻结成单值结论。只有先收住页面职责与首页语义，后续 `Feedback Signal Card`、路由/状态设计、后端聚合读取与验收路径才不会继续漂移。

## What Changes

- 冻结 `Dashboard Home` 为 `phase05-01` 当前阶段唯一页面主线
- 冻结 `dashboard_overview / Current Focus（内部承接 Next Action 主队列） / Asset Feedback / Recent Activity` 四个区块的固定归属关系
- 冻结第一屏信息层级与区块优先级，明确 `Current Focus / Next Action` 是唯一主行动队列
- 冻结 `Dashboard Home` 的正式业务入口路由为 `/dashboard`
- 冻结 `Dashboard` 与既有主导航、根级布局宿主及 `/` 的关系
- 冻结 `PC / 移动浏览器` 双场景下的信息层级与布局降级策略
- 冻结 `Dashboard Home` 四类区块的可点击入口热区为"整卡可点"统一规则，空状态主 CTA 例外为按钮模式
- 明确当前阶段不形成复杂驾驶舱、不形成第二套首页体系、不引入第二套移动端 UI

## Impact

- Affected specs: `phase05_dashboard_feedback_foundation`
- Affected code: 后续 `frontend/src/routes/__root.tsx`、新增 `frontend/src/routes/dashboard.tsx` 或等价路由文件、`frontend/src/features/dashboard/` 页面与区块组件分层、相关导航与返回上下文承接逻辑

## ADDED Requirements

### Requirement: Phase05 页面主线冻结

系统 SHALL 将 `phase05-01` 的页面主线冻结为单一 `Dashboard Home`，并明确它只承接聚合读取、空状态引导与跳转入口，不承接新的业务写入入口。

#### Scenario: 页面范围判定

- **WHEN** 接手者阅读 `phase05-01` 页面规格
- **THEN** 必须得到 `Dashboard Home` 是当前子任务唯一页面主线的单值结论
- **AND** 不得在本子任务并行引入 `Dashboard Detail`、独立 `Feedback` 页面、复杂驾驶舱页或第二套首页体系
- **AND** 不得把 `Product Detail`、`Repository Binding Detail / Workspace`、`Decision Detail / Decision Center`、`Module Detail` 的既有写入 owner 迁入 Dashboard

### Requirement: Dashboard Home 页面职责冻结

系统 SHALL 将 `Dashboard Home` 冻结为当前阶段 `Dashboard` 主线的默认进入页，承接概览读取、反馈信号主队列、资产缺口补充摘要、最近活动展示与空状态主 CTA。

#### Scenario: Dashboard Home 职责判定

- **WHEN** 用户进入 `Dashboard Home`
- **THEN** 页面必须承接 `dashboard_overview` 的概览读取
- **AND** 必须承接 `Current Focus / Next Action` 主行动队列
- **AND** 必须承接 `Asset Feedback` 补充摘要区块
- **AND** 必须承接 `Recent Activity` 活动流区块
- **AND** 必须承接空状态主 CTA 与既有 canonical owner 页面跳转入口
- **AND** 当前阶段不得在本页直接提交 `BindModuleToProduct`、`BindRepositoryToProduct`、`MapModuleToRepository` 或决策补录写入

### Requirement: Dashboard 四区块归属冻结

系统 SHALL 将 `Dashboard Home` 的最小页面级区块固定为 `dashboard_overview / Current Focus（内部承接 Next Action 主队列） / Asset Feedback / Recent Activity` 四个区块。

#### Scenario: 四区块归属判定

- **WHEN** 接手者判断 `Dashboard Home` 的最小区块组成
- **THEN** 必须得到上述四个区块的单值结论
- **AND** `dashboard_overview` 只承接概览卡片区块，不进入反馈优先级队列
- **AND** `Current Focus` 内部承接统一的 `Next Action` 主队列，不单独拆成与 `Current Focus` 并列的第二个页面区块
- **AND** `Asset Feedback` 只承接 `product_asset_coverage` 的补充摘要，不形成第二条独立优先级队列
- **AND** `Recent Activity` 只承接独立活动流，不与反馈信号混合排序

### Requirement: 第一屏信息层级冻结

系统 SHALL 将 `Current Focus / Next Action` 冻结为 `Dashboard Home` 第一屏优先承接的唯一主行动队列，并明确其他区块的降级关系。

#### Scenario: 第一屏优先级判定

- **WHEN** 用户首次进入 `Dashboard Home`
- **THEN** 第一屏必须优先暴露 `Current Focus / Next Action`
- **AND** `dashboard_overview` 作为概览摘要存在，但不得压过主行动队列成为第一优先区块
- **AND** `Asset Feedback` 只能作为主行动队列之后的补充摘要
- **AND** `Recent Activity` 只能作为独立活动流存在，不得替代主行动区块

### Requirement: Dashboard 正式入口与首页语义冻结

系统 SHALL 将 `Dashboard Home` 的正式业务入口冻结为 `/dashboard`，并明确当前阶段不把 `/` 单值解释为 `Dashboard Home`。

#### Scenario: 路由入口判定

- **WHEN** 接手者判断 `Dashboard` 的正式业务入口
- **THEN** 必须得到 `/dashboard` 的单值结论
- **AND** `Dashboard` 必须作为既有主导航中的一级入口新增
- **AND** 根级布局宿主继续只承接全局导航与页面容器语义，不自动等同于 `Dashboard Home`
- **AND** 当前阶段不得把 `/` 直接重写为 `Dashboard Home`

### Requirement: Dashboard 与既有主导航关系冻结

系统 SHALL 冻结 `Dashboard` 与既有主导航的关系，避免后续前端实现阶段出现“替换首页”与“新增导航项”两套并行解释。

#### Scenario: 主导航归属判定

- **WHEN** 用户在现有主导航中进入 `Dashboard`
- **THEN** 必须通过新增一级导航入口进入 `Dashboard Home`
- **AND** 既有 `Modules / Decisions / Products / Repositories` 导航主线继续保留
- **AND** 当前阶段不得把 `Dashboard` 设计成隐藏在其他页面中的附属区块

### Requirement: PC 与移动浏览器信息密度策略冻结

系统 SHALL 在单一 `React Web` 前端交付策略下，同时定义 `PC` 与移动浏览器的基础信息密度和布局降级规则。

#### Scenario: 桌面端信息密度

- **WHEN** 页面在 `PC` 桌面环境展示
- **THEN** `Dashboard Home` 必须优先承接较高信息密度
- **AND** 允许同屏展示概览、主行动队列、补充摘要与最近活动四类区块

#### Scenario: 移动浏览器信息密度

- **WHEN** 页面在移动浏览器窄屏环境展示
- **THEN** 必须保持与桌面端同一套页面语义与动作体系
- **AND** 必须通过区块垂直重排、信息裁剪、折叠或延后展示降低拥挤度
- **AND** 不得引入第二套独立移动端 UI 方案

### Requirement: Dashboard 可点击入口热区冻结

系统 SHALL 将 `Dashboard Home` 四类区块的可点击入口热区冻结为单值规则，避免后续前端实现阶段对“整卡可点”与“仅 `action_label` 可点”做出不同解释。

#### Scenario: 概览卡片点击热区判定

- **WHEN** 用户在 `dashboard_overview` 区块点击概览卡片（如 `module_count` + 标签组合）
- **THEN** 概览卡片整体作为可点击热区
- **AND** 仅以下四类一级概览卡片在当前阶段作为独立可点击卡片存在：`module_count`、`product_count`、`repository_count`、`decision_count`
- **AND** 点击后分别跳转到对应 List 页面：`module_count` → `/modules`、`product_count` → `/products`、`repository_count` → `/repositories`、`decision_count` → `/decisions`
- **AND** 不得仅将数字或标签单独作为可点击元素

#### Scenario: 派生产品覆盖计数的展示与点击判定

- **WHEN** `dashboard_overview` 区块展示 `product_with_repository_count` 或 `product_with_module_count`
- **THEN** 它们在当前阶段只作为概览辅助指标展示
- **AND** 不得在 `phase05-01` 中形成独立可点击概览卡片
- **AND** 不得为它们临时发明新的 `Product List` 筛选态、路由参数或第二套落点语义
- **AND** 如后续需要把它们升级为可点击入口，必须在后续 `phase05-02 / phase05-03 / phase05-05` 中单独冻结其 canonical 落点与筛选态

#### Scenario: Next Action 反馈信号卡片点击热区判定

- **WHEN** 用户在 `Current Focus / Next Action` 区块点击反馈信号卡片
- **THEN** 反馈信号卡片整体作为可点击热区
- **AND** `action_label` 作为视觉锚点（视觉高亮但不限制点击热区）
- **AND** 点击后跳转到 `target_type` 对应的 canonical owner 页面
- **AND** 不得仅将 `action_label` 单独作为可点击元素

#### Scenario: Asset Feedback 缺口项点击热区判定

- **WHEN** 用户在 `Asset Feedback` 区块点击缺口项
- **THEN** 缺口项整体作为可点击热区
- **AND** `action_label` 作为视觉锚点
- **AND** 点击后跳转到对应 `Product Detail`
- **AND** 不得仅将 `action_label` 单独作为可点击元素

#### Scenario: Recent Activity 活动项点击热区判定

- **WHEN** 用户在 `Recent Activity` 区块点击活动项
- **THEN** 活动项整体作为可点击热区
- **AND** 对象名称作为视觉锚点
- **AND** 点击后跳转到对应对象详情页
- **AND** 不得仅将对象名称单独作为可点击元素

#### Scenario: 空状态主 CTA 点击热区判定

- **WHEN** 用户在空状态下点击主 CTA
- **THEN** 主 CTA 作为按钮可点
- **AND** 不采用整卡模式
- **AND** 点击后跳转到对应 Create / Detail 页面

#### Scenario: 统一性约束

- **WHEN** 接手者判断 `Dashboard Home` 各区块的点击热区
- **THEN** 必须得到“整卡可点”是四类区块（`dashboard_overview` / `Current Focus / Next Action` / `Asset Feedback` / `Recent Activity`）的统一基础规则
- **AND** 空状态主 CTA 是唯一例外（按钮模式）
- **AND** 不得在同一区块内混合“整卡可点”与“仅按钮可点”两种模式

## MODIFIED Requirements

### Requirement: 根级布局宿主与业务首页的关系

现有根级布局宿主 SHALL 继续作为全局导航与页面容器宿主，而不是在 `phase05-01` 中被重新解释为 `Dashboard Home` 的默认业务首页。

#### Scenario: 判定根级布局含义

- **WHEN** 接手者阅读 `phase05-01` 后续前端路由与布局设计
- **THEN** 必须明确根级布局宿主只承接导航与页面容器语义
- **AND** 必须明确 `Dashboard Home` 通过 `/dashboard` 进入，而不是通过重新解释根级宿主的业务含义进入

## REMOVED Requirements

### Requirement: 将 `/` 隐式解释为 Dashboard Home

**Reason**: 该解释会与既有根级布局宿主、现有主导航结构以及 `Dashboard` 新增一级入口的冻结结论冲突，导致 `/` 与 `/dashboard` 并存两套首页语义。

**Migration**: 当前阶段统一以 `/dashboard` 作为 `Dashboard Home` 的正式业务入口；如后续确需把 `/` 重定向到 `/dashboard`，必须在新的 `phase / fix` 中单独冻结并同步更新相关规格与导航语义。
