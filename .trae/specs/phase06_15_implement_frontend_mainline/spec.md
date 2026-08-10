# Phase06-15 Onboarding + Sovereignty + Reuse 前端主线实现 Spec

## Why

`phase06-06` 已冻结 Onboarding 前端页面、路由与回流语义，`phase06-07` 已冻结四类 canonical create 的固定 mutation 承接位，`phase06-09` 已冻结 Reuse Summary 在 Dashboard / Detail 的页面级 query 与挂接位，`phase06-12` 已将这些结论收口为 `phase06` 正式规格正文，`phase06-14` 又把后端、数据与脚本主线推进为真实可运行实现。但截至当前，仓库中仍缺少 `/onboarding`、Export / Backup 用户入口、Reuse Summary 页面挂接与四类 create 页面的 owner 回收实现。

如果不在这一阶段把前端主线真正落到仓库，`phase06` 的关键闭环仍然只停留在“后端可用、规格已冻”的状态：用户无法从前端走通首轮录入，既有 create 页面仍会保留与 Onboarding 冲突的第二套写入语义，Dashboard / Detail 也无法展示最小复用反馈，更无法把 Export / Backup 暴露为正式用户入口。`phase06-15` 的目标，就是把这些已冻结结论推进为可编译、可运行、可联调的单一前端主线。

## What Changes

- 新增 `frontend/src/routes/onboarding.tsx`，承接 `/onboarding` 唯一路由入口
- 新增 `frontend/src/routes/index.tsx`，承接根级 first-run 默认进入路径
- 新增 `frontend/src/features/onboarding/`，承接 Onboarding 页面、组件、只读 query owner 与步骤编排
- 新增 `frontend/src/features/reuse-summary/`，承接 `ReuseSummaryRead` 的前端类型、API 适配层与页面级 query owner
- 新增四个 feature slice 的 `application/` 目录与正式 create owner：
  - `frontend/src/features/product-registry/application/`
  - `frontend/src/features/repository-binding/application/`
  - `frontend/src/features/module-registry/application/`
  - `frontend/src/features/decision-center/application/`
- 修改 `frontend/src/features/dashboard/pages/dashboard-home-page.tsx` 与相关组件，接入 `first_run_state`、Export / Backup 入口与 Dashboard 内复用快照子区域
- 修改 `frontend/src/routes/products/$productId.tsx`、`repositories/$repositoryId.tsx`、`modules/$moduleId.tsx`、`decisions/$decisionId.tsx`，承接 `fromOnboarding` 回流参数
- 修改 `frontend/src/features/product-registry/pages/product-create-page.tsx`
- 修改 `frontend/src/features/repository-binding/pages/repository-create-page.tsx`
- 修改 `frontend/src/features/module-registry/pages/module-create-page.tsx`
- 修改 `frontend/src/features/decision-center/pages/decision-create-page.tsx`
- 修改 `frontend/src/features/product-registry/pages/product-detail-page.tsx`
- 修改 `frontend/src/features/repository-binding/pages/repository-binding-detail-page.tsx`
- 修改 `frontend/src/features/module-registry/pages/module-detail-page.tsx`
- 修改 `frontend/src/features/decision-center/pages/decision-detail-page.tsx`
- **BREAKING**：`phase06` 覆盖范围内的四个 create 页面完成回收后，不得继续内联 page-level `useMutation` 作为正式 create 主线

## Impact

- Affected specs:
  - `phase06_06_design_onboarding_frontend_route_interaction_flow`
  - `phase06_07_design_frontend_write_path_mutation_owners`
  - `phase06_09_design_reuse_summary_query_dashboard_detail_integration`
  - `phase06_12_onboarding_sovereignty_reuse_formal_spec`
  - `phase06_14_implement_backend_data_script_mainline`
  - 后续 `phase06` 联调验收
- Affected code:
  - `frontend/src/routes/index.tsx`
  - `frontend/src/routes/onboarding.tsx`
  - `frontend/src/routes/__root.tsx`
  - `frontend/src/routes/products/$productId.tsx`
  - `frontend/src/routes/repositories/$repositoryId.tsx`
  - `frontend/src/routes/modules/$moduleId.tsx`
  - `frontend/src/routes/decisions/$decisionId.tsx`
  - `frontend/src/features/onboarding/`
  - `frontend/src/features/reuse-summary/`
  - `frontend/src/features/dashboard/pages/dashboard-home-page.tsx`
  - `frontend/src/features/dashboard/components/dashboard-primary-action-panel.tsx`
  - `frontend/src/features/dashboard/components/asset-feedback-section.tsx`
  - `frontend/src/features/dashboard/components/dashboard-stat-bar.tsx`
  - `frontend/src/features/product-registry/application/`
  - `frontend/src/features/repository-binding/application/`
  - `frontend/src/features/module-registry/application/`
  - `frontend/src/features/decision-center/application/`
  - `frontend/src/features/product-registry/pages/product-create-page.tsx`
  - `frontend/src/features/repository-binding/pages/repository-create-page.tsx`
  - `frontend/src/features/module-registry/pages/module-create-page.tsx`
  - `frontend/src/features/decision-center/pages/decision-create-page.tsx`
  - `frontend/src/features/product-registry/pages/product-detail-page.tsx`
  - `frontend/src/features/repository-binding/pages/repository-binding-detail-page.tsx`
  - `frontend/src/features/module-registry/pages/module-detail-page.tsx`
  - `frontend/src/features/decision-center/pages/decision-detail-page.tsx`

## ADDED Requirements

### Requirement: Onboarding 前端主线必须落地为唯一正式首轮入口

系统 SHALL 将 first-run onboarding 落地为仓库中真实存在、可运行的前端主线，并继续以 `/onboarding` 作为唯一正式首轮业务入口。

#### Scenario: 路由与根级默认进入路径

- **WHEN** 实现 `phase06-15` 的前端入口
- **THEN** 必须新增 `frontend/src/routes/onboarding.tsx`
- **AND** 必须新增 `frontend/src/routes/index.tsx`
- **AND** `/` 的 `beforeLoad` 必须读取 `GetFirstRunState`
- **AND** `first_run_state = not_started` 时必须默认进入 `/onboarding`
- **AND** `first_run_state = in_progress` 时必须默认进入 `/dashboard`
- **AND** `first_run_state = completed` 时不得默认劫持到 `/onboarding`

#### Scenario: Onboarding 页面与步骤编排

- **WHEN** 实现 `OnboardingPage`
- **THEN** 必须完整承接 `welcome / product / repository / module / decision / complete` 六段步骤语义
- **AND** Product / Repository / Module / Decision 四个步骤的最小人工必填字段必须继续对齐：
  - `Product`: `name`
  - `Repository`: `name + url`
  - `Module`: `name`
  - `Decision`: `title + choice + reason`
- **AND** 页面组件不得内联正式 mutation 主线
- **AND** 步骤表单必须通过各自 feature slice 的 `application` owner 提交
- **AND** 四类对象都完成最小持久化后，页面必须进入 `complete` 步骤

#### Scenario: Dashboard 到 Onboarding 的继续入口

- **WHEN** Dashboard 读取 `first_run_state`
- **THEN** `not_started` 时必须展示 `Start Onboarding`
- **AND** `in_progress` 时必须展示 `Continue Onboarding`
- **AND** `completed` 时不得继续展示首轮录入 CTA
- **AND** 当前阶段不得在 `/dashboard` 外再发明第二个首轮录入入口

### Requirement: 四类 canonical create 必须收敛到固定 mutation 承接位

系统 SHALL 将 `Product / Repository / Module / Decision` 的 create 主线统一回收到各自 feature slice 的 `application` 目录，不再允许 page-level `useMutation` 与 Onboarding 并行存在。

#### Scenario: application owner 物理落点

- **WHEN** 实现四类 create 主线
- **THEN** 必须存在以下正式承接位：
  - `frontend/src/features/product-registry/application/use-create-draft-product.ts`
  - `frontend/src/features/repository-binding/application/use-create-draft-repository.ts`
  - `frontend/src/features/module-registry/application/use-create-draft-module.ts`
  - `frontend/src/features/decision-center/application/use-create-draft-decision.ts`
- **AND** 每个 owner 都必须导出统一的 `mutate / mutateAsync / isPending / isError / error / data`
- **AND** owner 内部必须承接默认补值、错误归一化与 query 失效
- **AND** 页面、表单与展示组件不得自行拼装第二套正式 create 语义

#### Scenario: 既有 create 页面回收

- **WHEN** 修改四类既有 create 页面
- **THEN** `product-create-page.tsx`、`repository-create-page.tsx`、`module-create-page.tsx`、`decision-create-page.tsx` 必须改为消费各自 `application` owner
- **AND** 现有 page-level `useMutation` 必须被移除
- **AND** 成功回流、错误归一化与 query 失效语义必须与 Onboarding 共享同一套 owner 行为
- **AND** 当前阶段不得继续保留与 Onboarding 冲突的第二套 create 主线

#### Scenario: query 层继续保持纯只读

- **WHEN** 实现 `phase06-15` 的前端写路径
- **THEN** `data/` 或 `query owner` 层只允许承接读取、缓存键与响应解包
- **AND** `create / update / bind / link` 不得混入 `query` 层
- **AND** 当前阶段不得为了方便把 Onboarding 写入逻辑塞回 `useOnboardingRead` 或 `ReuseSummaryRead`

### Requirement: Reuse Summary 必须作为页面级只读快照挂接到 Dashboard 与 Detail

系统 SHALL 将 `ReuseSummaryRead` 落地为页面级只读 query，并把最小复用反馈挂接到 Dashboard / Module Detail / Product Detail，而不是新开一级导航或第二套经营工作台。

#### Scenario: Reuse Summary 前端切片落地

- **WHEN** 实现复用感知前端主线
- **THEN** 必须新增 `frontend/src/features/reuse-summary/`
- **AND** 至少包含：
  - `types.ts`
  - `data/api-adapter.ts`
  - `data/use-reuse-summary-read.ts` 或等价 query owner 文件
- **AND** 字段语义必须单向承接 `phase06-14` 已落地的后端响应
- **AND** 当前阶段不得为 `module_reuse_summary` 与 `capability_summary` 长出两套平行 query owner

#### Scenario: Dashboard 复用快照挂接位

- **WHEN** Dashboard 展示 `Asset Feedback`
- **THEN** 必须在 `Asset Feedback` 区块内部增加独立 `Reuse Snapshot` 子区域
- **AND** Dashboard 页面必须额外新增一个页面级 `ReuseSummaryRead` query
- **AND** 该 query 失败不得把整页打回 `page-error`
- **AND** `module_reuse_summary` 与 `capability_summary` 都必须在 Dashboard 下按“数量优先、时间次级”排序并最多展示前 `5` 条
- **AND** 复用反馈必须可见且可解释，不得只显示数字而缺少解释文本

#### Scenario: Module Detail 与 Product Detail 挂接位

- **WHEN** Module Detail 或 Product Detail 进入 `ready`
- **THEN** Module Detail 必须在 `Module Summary` 邻近区域挂接复用摘要
- **AND** Product Detail 必须在已绑定模块相关区域附近挂接复用摘要
- **AND** 每个详情页都必须只新增一个页面级 `ReuseSummaryRead` query
- **AND** 详情页中的复用反馈不得承接绑定写入、解绑写入或候选筛选逻辑

### Requirement: Export / Backup 必须作为 Dashboard 内正式用户入口

系统 SHALL 将 `Export / Backup` 暴露为 Dashboard 内正式可点击、可读取、可触发的用户入口，而不是只停留在后端路由存在。

#### Scenario: Dashboard 用户入口落地

- **WHEN** 实现 Dashboard 用户入口
- **THEN** Dashboard 中必须存在稳定可见的 Export 与 Backup 入口
- **AND** 用户可以读取当前 `ExportSnapshot` 与 `BackupSnapshot`
- **AND** 用户可以从前端触发导出与备份动作
- **AND** Backup 的三类失败语义必须继续可区分：
  - `manifest_missing`
  - `coverage_incomplete`
  - `schema_mismatch`
- **AND** 当前阶段不得把 Export / Backup 做成隐藏路由、开发者入口或仅测试按钮

#### Scenario: Snapshot 消费边界

- **WHEN** 前端消费 `ExportSnapshot` 与 `BackupSnapshot`
- **THEN** 必须继续从 `.proto -> HTTP DTO` 单向派生字段语义
- **AND** 不得在前端 view model 中私自新增第二套业务字段
- **AND** `backup verified` 与三类失败语义必须继续与后端单值语义一致

## MODIFIED Requirements

### Requirement: Dashboard 主线必须并入 phase06 的 Onboarding / Sovereignty / Reuse 入口

系统 SHALL 在保留 `phase05` Dashboard 读主线的前提下，将 `Onboarding CTA / Export / Backup / Reuse Snapshot` 并入当前 Dashboard 页面，而不是另起第二套 dashboard 变体。

#### Scenario: Dashboard 页面编排升级

- **WHEN** `DashboardHomePage` 在 `phase06-15` 进入实现
- **THEN** 必须继续保留 `overview / feedback / recent-activity` 三个既有查询
- **AND** 必须新增 `OnboardingRead` 与 `ReuseSummaryRead` 的页面级读取承接
- **AND** `Export / Backup` 入口必须在现有 Dashboard 页面内集成
- **AND** 当前阶段不得把这些 `phase06` 能力拆成新的一级导航页面

### Requirement: canonical detail 页返回优先级必须正式并入 fromOnboarding

系统 SHALL 将 `fromOnboarding=true + onboardingStep` 正式并入四类 canonical detail 页的返回优先级中，避免用户从 Onboarding 进入详情补全后丢失回流路径。

#### Scenario: detail 页来源优先级

- **WHEN** 用户从 Onboarding 草稿摘要进入任一 canonical detail 页面
- **THEN** detail 页返回路径必须优先回到 `/onboarding`
- **AND** 必须优先恢复 `onboardingStep` 指定步骤
- **AND** `fromOnboarding` 的返回优先级必须高于 `fromList / fromDashboard / fromProductDetail / fromModuleDetail`
- **AND** 当前阶段不得让四个 detail 页分别自由决定回流优先级

## REMOVED Requirements

### Requirement: 四类 create 页面的 page-level useMutation 作为正式主线
**Reason**: 该模式会让既有 create 页面与 Onboarding 并存两套正式写入语义，违反 `phase06-07` 已冻结的固定 mutation 承接位规则。
**Migration**: 四类 create 页面统一改为消费各自 feature slice 的 `application` owner；页面只保留表单编排、toast 与导航消费，不再直接持有正式 mutation。
