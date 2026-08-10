# Phase06-12 Onboarding + Sovereignty + Reuse 正式规格正文 Spec

## Why

`phase06-01 ~ 11` 已分别冻结了 first-run 入口、draft-first、Export / Backup 语义、Reuse Summary 读模型、前端交互流、写路径 owner、后端边界、最小 Proto 合同与验收基线，但这些结论仍散落在多个子规格中。若不在 `phase06-12` 把这些冻结结论收敛为首份正式规格正文，后续 `phase06-13` 合同主线、实现、联调验收与下一阶段继续承接时，仍会把 `phase06-01 ~ 11` 当作并列上游使用，重新长出第二套页面边界、第二套写路径语义与第二套验收口径。

## What Changes

- 产出 `phase06` 首份正式规格正文，作为 `Onboarding + Data Sovereignty + Reuse Awareness` 的唯一直接上游规格来源
- 收敛 `phase06-01 ~ 11` 已冻结结论，不另立第二套页面、交互、写路径、合同、验收或非目标边界
- 正式覆盖当前阶段的页面矩阵、路由语义、`first_run_state`、draft-first 写模型、导出 / 备份、复用读模型、合同、验收基线、非目标与 Done 标准
- 明确 `phase06-13` 及后续实现只允许承接本文档，不再并列消费 `phase06-01 ~ 11`
- **BREAKING**：`phase06-01 ~ 11` 在本文档生效后退为追溯来源与证据链，不再承担并列直接执行层入口职责

## Impact

- Affected specs:
  - `phase06_onboarding_sovereignty_reuse_foundation`
  - `phase06-01` first-run onboarding 边界、入口与最小完成条件
  - `phase06-02` draft-first / partial-entry 写路径
  - `phase06-03` Export / Backup / restore prerequisites
  - `phase06-04` `module_reuse_summary / capability_summary` 最小读模型与页面挂接位
  - `phase06-05` 合同、传输与源码约束
  - `phase06-06` Onboarding 前端页面、路由与交互流
  - `phase06-07` 前端写路径收敛与 mutation 承接位
  - `phase06-08` Export / Backup 后端模块边界
  - `phase06-09` 复用感知 query / Dashboard / Detail 集成
  - `phase06-10` 最小 Protocol Buffers 合同设计
  - `phase06-11` 联调验收环境、fixture 与恢复基线
- Affected code:
  - 后续 `frontend/src/routes/onboarding.tsx`、`frontend/src/routes/index.tsx`、`frontend/src/routes/__root.tsx`
  - 后续 `frontend/src/features/onboarding/`、`frontend/src/features/export/`、`frontend/src/features/backup/`、`frontend/src/features/reuse-summary/`
  - 后续 `frontend/src/features/product-registry/`、`repository-binding/`、`module-registry/`、`decision-center/` 中的 `application/`
  - 后续 `backend/internal/export/`、`backend/internal/backup/`、`backend/internal/reusesummary/` 与 Onboarding 读承接位
  - 后续 `proto/psco/onboarding/v1/`、`export/v1/`、`backup/v1/`、`reuse_summary/v1/`
  - 后续 `database/scripts/reset_phase06_acceptance.sh` 与 `database/seeds/seed_phase06_*`

## ADDED Requirements

### Requirement: `phase06` 正式规格正文角色冻结

系统 SHALL 将 `phase06-12` 产出的正式规格正文冻结为 `phase06_onboarding_sovereignty_reuse_foundation` 的唯一直接执行层规格入口。

#### Scenario: 正文角色判定

- **WHEN** 接手者进入 `phase06` 后续合同落地、实现、联调验收或下一阶段承接
- **THEN** 必须以本正式规格正文作为唯一直接上游规格来源
- **AND** `phase06-01 ~ 11` 只保留为追溯来源与证据链
- **AND** 当前阶段不得继续把 `phase06-01 ~ 11` 并列当作直接执行层入口

### Requirement: 技术路线与当前阶段主线冻结

系统 SHALL 在本正式规格正文中继续维持 `phase06` 已冻结的技术路线与阶段目标，不得重新解释为其他技术或产品路线。

#### Scenario: 技术路线判定

- **WHEN** 接手者阅读 `phase06` 正式规格正文
- **THEN** 必须得到以下单值结论：
  - 项目路线：`Durable System Track`
  - 正式运行主线：`React Web + Go Backend + PostgreSQL`
  - 前端：`React + Vite + TypeScript + TanStack Router + TanStack Query + Tailwind CSS + shadcn/ui`
  - 后端：`Go + chi + net/http`
  - 合同：`Protocol Buffers`
  - 合同工具链：`buf build / lint / generate / breaking`
- **AND** 当前阶段不得重新解释为 `Product Track`
- **AND** 当前阶段不得引入独立 `React Native` 客户端、AI 一级导航、完整模板系统或独立 Capability 重实体

### Requirement: 页面与路由矩阵冻结

系统 SHALL 在正式规格正文中完整覆盖 `phase06` 的页面矩阵与正式业务路由。

#### Scenario: 页面范围判定

- **WHEN** 接手者判断 `phase06` 的正式页面范围
- **THEN** 必须完整覆盖：
  - `Onboarding Home / Flow`
  - `Dashboard Home`
  - `Module Detail`
  - `Product Detail`
  - `Repository Binding Detail / Workspace`
  - `Decision Center / Detail`
  - `Export`
  - `Backup`
- **AND** 当前阶段不得新建独立 `Reuse Center` 或独立运维中心页面

#### Scenario: 正式业务路由判定

- **WHEN** 接手者判断 `phase06` 的正式业务路由
- **THEN** 必须得到以下单值结论：
  - `Onboarding`：`/onboarding`
  - `Export`：`/dashboard/export`
  - `Backup`：`/dashboard/backup`
  - cold-start 主 CTA：`/onboarding`
  - 未完成回访主 CTA：`Continue Onboarding -> /onboarding`
- **AND** `Export / Backup` 允许从 `Dashboard` 动作区进入
- **AND** `Export / Backup` 不得在 `Dashboard Home` 主内容区内联完成全部操作

#### Scenario: Onboarding 到 canonical detail 的回流链判定

- **WHEN** 用户从 `Onboarding` 内的草稿摘要进入任一 canonical detail 页面继续补全
- **THEN** 正式语义必须继续承接 `fromOnboarding=true` 与必要的 `onboardingStep` 返回参数
- **AND** `Product / Repository / Module / Decision` 对应 detail route 的 `validateSearch` 必须承接该返回参数
- **AND** canonical detail 页面检测到 `fromOnboarding=true` 时，返回路径必须优先回到 `/onboarding`
- **AND** 若同时存在 `fromList / fromDashboard / fromProductDetail / fromModuleDetail` 等既有来源上下文，`fromOnboarding` 的回流优先级必须按上游冻结口径显式并入，而不是由各页面自由决定
- **AND** 返回到 `/onboarding` 时必须优先恢复 `onboardingStep` 指定步骤；若未携带该参数，才允许回退为由 `Onboarding` 页面自动定位

### Requirement: `first_run_state` 与 Onboarding 交互主线冻结

系统 SHALL 在正式规格正文中完整覆盖 `first_run_state`、根级入口判定、Onboarding 步骤主线与回访继续语义。

#### Scenario: 状态机判定

- **WHEN** 接手者判断 `first_run_state`
- **THEN** 必须冻结为 `not_started / in_progress / completed`
- **AND** 状态跃迁必须继续对齐：
  - 尚未开始任何首轮对象写入：`not_started`
  - 已至少创建 `1` 条首轮对象记录但四类对象未全部持久化：`in_progress`
  - 四类对象均已持久化并满足首轮成功会话条件：`completed`

#### Scenario: 根级默认进入路径

- **WHEN** 应用启动并由根级路由入口守卫判定默认进入路径
- **THEN** `first_run_state = not_started` 时必须回落到 `/onboarding`
- **AND** `first_run_state = in_progress` 时根级默认进入 `/dashboard`
- **AND** `Dashboard` 必须提供稳定可见的 `Continue Onboarding`
- **AND** `first_run_state = completed` 时不得再默认劫持到 `/onboarding`

#### Scenario: Onboarding 步骤与完成条件

- **WHEN** 接手者判断 Onboarding 的首轮录入主线
- **THEN** 推荐执行顺序必须冻结为 `Product -> Repository -> Module -> Decision`
- **AND** `welcome / product / repository / module / decision / complete` 六段语义必须在正式正文中被完整承接
- **AND** 首轮成功会话成立条件必须继续冻结为：同一个 first-run onboarding run 中，`Product / Repository / Module / Decision` 四类对象都已完成最小持久化
- **AND** 当前阶段不得把“只创建了部分对象”判定为首轮成功会话成立

### Requirement: Draft-First 写模型与前端写路径边界冻结

系统 SHALL 在正式规格正文中完整覆盖 draft-first / partial-entry 写模型、四类 canonical create 合同复用关系与前端 `application` 单入口约束。

#### Scenario: draft-first 最小写模型

- **WHEN** 接手者判断 `phase06` 的四类首轮创建动作
- **THEN** 必须继续冻结为“先创建最小可持久化记录，再允许后补完整信息”
- **AND** 四类最小人工必填字段必须继续对齐：
  - `Product`：`name`
  - `Repository`：`name + url`
  - `Module`：`name`
  - `Decision`：`title + choice + reason`
- **AND** `description / provider / capability_key / target_link` 等不得重新拉回首轮强制前置

#### Scenario: canonical create 合同复用

- **WHEN** 接手者判断四类 draft-first 写入的正式合同归属
- **THEN** 必须继续复用既有 canonical create contract：
  - `CreateProduct`
  - `CreateRepository`
  - `CreateModule`
  - `CreateDecision`
- **AND** `Onboarding` 不得长出第二套 `CreateDraft*` canonical 合同
- **AND** 当前阶段不得发明 `/api/onboarding/drafts/*` 第二套路由分组

#### Scenario: 前端写路径 owner 边界

- **WHEN** 接手者判断 `phase06` 的前端写路径边界
- **THEN** 每个核心对象必须只有一个 feature-slice `application` 承接位
- **AND** `query` 层必须保持纯只读
- **AND** 页面、表单、展示组件不得各自内联正式 mutation 主线
- **AND** `Onboarding` 页面与既有 create 页面必须共享同一套成功回流、错误归一化与 query 失效语义

### Requirement: Export / Backup 正式语义与边界冻结

系统 SHALL 在正式规格正文中完整覆盖 `Export / Backup / backup verified / restore prerequisites` 的正式语义、覆盖矩阵、入口位与错误边界。

#### Scenario: Export 语义与覆盖矩阵

- **WHEN** 接手者判断 `Export` 的正式语义
- **THEN** 必须得到“面向用户带走核心资产数据”的单值结论
- **AND** 最小覆盖矩阵必须至少包含：
  - `products`
  - `modules`
  - `releases`
  - `repositories`
  - `decisions`
  - `decision_links`
  - `product_modules`
  - `product_repositories`
  - `module_repositories`
- **AND** 不得把“只导出主实体，不导出绑定关系”判定为完成数据主权闭合

#### Scenario: Backup 语义与 `backup verified`

- **WHEN** 接手者判断 `Backup` 的正式语义
- **THEN** 必须得到“面向当前实例保留与恢复前提校验”的单值结论
- **AND** `Backup` 的最小覆盖范围不得小于 `Export`
- **AND** `Backup` 还必须带出 `manifest`、备份创建时间与 `schema / version` 前提
- **AND** `backup verified` 的最小成立条件必须继续冻结为：
  - 已生成可读取的备份产物
  - 可重新读取并解析 `manifest`
  - `manifest` 中可见核心资产覆盖矩阵
  - `schema / version` 恢复前提可校验
- **AND** 当前阶段不得把“产物写出成功”直接等价为 `backup verified`

#### Scenario: 非目标边界

- **WHEN** 接手者扩展 `Export / Backup`
- **THEN** 当前阶段不得扩写为连续备份、复杂灾备、完整 restore 写回或依赖第三方平台的能力

### Requirement: Reuse Summary 正式读模型与页面挂接位冻结

系统 SHALL 在正式规格正文中完整覆盖 `module_reuse_summary / capability_summary` 的事实来源、页面级 query 组织、挂接位与新鲜度语义。

#### Scenario: 最小读模型判定

- **WHEN** 接手者判断复用感知的最小读模型
- **THEN** 必须完整覆盖：
  - `module_reuse_summary`：`module_id / reuse_product_count / latest_reuse_at / explanation_text`
  - `capability_summary`：`capability_key / capability_label / supporting_module_count / latest_capability_update_at / empty_state_text`
- **AND** `module_reuse_summary` 的统计口径必须继续冻结为“一个 Module 当前被多少 Product 直接复用”
- **AND** `capability_summary` 的事实来源必须继续冻结为 `Module.capability_key + capability_label` 映射

#### Scenario: 页面级 query 与挂接位

- **WHEN** 接手者判断 `ReuseSummaryRead` 的正式消费方式
- **THEN** Dashboard / Module Detail / Product Detail 每页必须只新增一个页面级 `ReuseSummaryRead` `useQuery`
- **AND** 该结果对象必须同时承接 `module_reuse_summary[] + capability_summary[]`
- **AND** Dashboard 挂接位必须位于 `Asset Feedback` 内的复用快照子区域
- **AND** Module Detail 挂接位必须位于 `Module Summary` 邻近区域
- **AND** Product Detail 挂接位必须位于已绑定模块相关区域附近
- **AND** 当前阶段不得新建独立一级复用导航

#### Scenario: Reuse Summary 的排序、裁剪与展示范围

- **WHEN** 接手者判断 `ReuseSummaryRead` 在各页面的正式展示边界
- **THEN** Dashboard 中的 `module_reuse_summary` 必须按 `reuse_product_count DESC` 排序，同值时按 `latest_reuse_at DESC` 排序，且最多展示前 `5` 条
- **AND** Dashboard 中的 `capability_summary` 必须按 `supporting_module_count DESC` 排序，同值时按 `latest_capability_update_at DESC` 排序，且最多展示前 `5` 条
- **AND** Product Detail 中的 `module_reuse_summary` 必须先限定在当前 Product 已绑定模块作用域内，再按同样的数量优先、时间次级排序规则全量展示
- **AND** Product Detail 中的 `capability_summary` 必须先限定在当前 Product 已绑定且填写了 `capability_key` 的 Module 作用域内，再按同样的数量优先、时间次级排序规则全量展示
- **AND** Module Detail 必须继续展示当前 `module_id` 的直接复用反馈，并在存在 `capability_key` 时展示对应 `capability_summary` 最小摘要

#### Scenario: 新鲜度与空态

- **WHEN** 相关绑定或已提交状态变化影响复用感知
- **THEN** 再次读取时必须反映最新已提交状态
- **AND** 当前阶段不得依赖异步统计表或离线批处理才能展示复用反馈
- **AND** 成功空态与读取失败态必须严格区分

### Requirement: 合同、传输与演进基线冻结

系统 SHALL 在正式规格正文中完整覆盖 `phase06` 的 `.proto` 唯一合同源、HTTP 映射、包名版本与 breaking 演进规则。

#### Scenario: 合同源与服务矩阵

- **WHEN** 接手者判断 `phase06` 的正式合同基线
- **THEN** 必须继续冻结 `.proto` 为唯一长期合同源
- **AND** 必须完整承接以下最小合同分组：
  - `psco.onboarding.v1`
  - `psco.export.v1`
  - `psco.backup.v1`
  - `psco.reuse_summary.v1`
- **AND** 必须完整承接以下最小服务矩阵：
  - `OnboardingService.GetFirstRunState`
  - `ExportService.GetExportSnapshot / ExportCoreAssets`
  - `BackupService.GetBackupSnapshot / CreateInstanceBackup`
  - `ReuseSummaryService.GetReuseSummary`

#### Scenario: HTTP 映射与 DTO 单向派生

- **WHEN** 接手者判断 `phase06` 的 HTTP 适配层
- **THEN** 必须继续对齐以下正式映射：
  - `GET /api/onboarding/state`
  - `POST /api/products`
  - `POST /api/repositories`
  - `POST /api/modules`
  - `POST /api/decisions`
  - `GET /api/dashboard/export`
  - `POST /api/dashboard/export`
  - `GET /api/dashboard/backup`
  - `POST /api/dashboard/backup`
  - `GET /api/reuse-summary`
- **AND** HTTP DTO、后端 `types.go` 与前端 `types.ts` 只能从 `.proto` 单向派生或显式对齐映射
- **AND** 不得再形成第二套 HTTP canonical contract

#### Scenario: breaking 规则

- **WHEN** 接手者演进 `phase06` 的 `.proto`
- **THEN** 必须继续遵守：
  - 枚举首值为 `*_UNSPECIFIED = 0`
  - 删除字段或枚举值后必须 `reserved`
  - 不得复用 tag
  - 不得随意修改字段类型、重复性或语义
  - 不得新增 `required`

### Requirement: 联调验收环境与 fixture 基线冻结

系统 SHALL 在正式规格正文中完整覆盖 `phase06` 的统一 reset 入口、fixture 白名单、阶段完成矩阵与合同一致性验收门禁。

#### Scenario: 统一验收入口与 fixture 白名单

- **WHEN** 接手者判断 `phase06` 的联调验收基线
- **THEN** 必须继续冻结 `reset_phase06_acceptance.sh` 为唯一统一入口
- **AND** 统一入口最少必须支持默认、`--clean-only`、`--restore-only`、`--fixture <name>`
- **AND** fixture 白名单必须继续冻结为：
  - `cold-start-empty`
  - `in-progress-partial-entry`
  - `completed-unbound`
  - `completed-bound`
  - `export-ready`
  - `backup-verified`
  - `backup-manifest-missing`
  - `backup-coverage-incomplete`
  - `backup-schema-mismatch`
  - `reuse-latest`
  - `reuse-latest-after-binding`

#### Scenario: 阶段完成矩阵

- **WHEN** 验收人员判断 `phase06` 是否达到阶段完成条件
- **THEN** 必须至少独立验证：
  - `cold-start-empty`：入口判定与 `/onboarding`
  - `in-progress-partial-entry`：回访继续与 `Continue Onboarding`
  - `completed-unbound`：缺少绑定关系仍完成首轮会话
  - `completed-bound`：绑定补全后再次验证
  - `export-ready`：数据主权导出闭合
  - `backup-verified`：恢复前提校验成立
  - `reuse-latest-after-binding`：复用感知读取最新已提交状态
- **AND** Onboarding / Export / Backup / Reuse Summary 的 `.proto -> HTTP DTO -> 前端消费模型` 合同一致性必须纳入验收矩阵
- **AND** `backup_snapshot` 读取侧合同一致性必须由 `BackupWrite.read_verify` 或等价上游冻结承接位独立验证
- **AND** 当前阶段不得依赖手工补 SQL、手工改 `first_run_state`、手工补绑定关系或手工伪造备份结果

### Requirement: 非目标与 Done 标准冻结

系统 SHALL 在正式规格正文中显式承接 `phase06` 的非目标矩阵与 Done 标准，避免后续实现继续向外扩写。

#### Scenario: 非目标矩阵

- **WHEN** 接手者判断 `phase06` 当前阶段明确不做什么
- **THEN** 必须显式列出并维持以下非目标：
  - `Capability` 重实体 CRUD
  - 完整模板系统
  - `Opportunity / Feature / Experiment` 主线
  - GitHub OAuth / 自动导入
  - AI 一级工作台
  - 自动扫描 / 知识图谱
  - 连续备份 / 复杂灾备 / 完整 restore
- **AND** 当前阶段不得把上述内容写成既成事实或实现前置

#### Scenario: Done 标准

- **WHEN** 接手者判断 `phase06` 正式规格何时可进入稳定实现
- **THEN** 文档必须完整覆盖：
  - 页面与路由
  - `first_run_state` 与交互流
  - draft-first 写路径与 owner 边界
  - Export / Backup 语义与错误边界
  - Reuse Summary 读模型与挂接位
  - 合同、传输与演进规则
  - 联调验收环境、fixture 与合同一致性门禁
  - 非目标矩阵
- **AND** 上述条目缺一不可

## MODIFIED Requirements

### Requirement: `phase06` 直接执行层上游

`phase06` 在进入正式规格正文阶段后 SHALL 从“多份冻结子规格并行提供设计结论”推进为“由单一正式规格正文承接并向后游交付”。

#### Scenario: 上游职责收口

- **WHEN** 接手者继续推进 `phase06-13` 及后续实现
- **THEN** 必须先读取本正式规格正文
- **AND** `phase06-01 ~ 11` 只作为本文的追溯来源与证据链
- **AND** 不得再跳过本文直接拼接多个子规格形成第二套执行入口

## REMOVED Requirements

### Requirement: `phase06-01 ~ 11` 继续并列承担直接执行层入口职责

**Reason**: 这种做法会让页面边界、写路径、合同和验收口径继续分散在多个子规格中，直接破坏 `phase06-12` 作为首份正式规格正文的职责。

**Migration**: `phase06-12` 生效后，后续合同主线、实现与验收统一以本正式规格正文为唯一直接上游；`phase06-01 ~ 11` 退为追溯来源、设计证据与复核依据。
