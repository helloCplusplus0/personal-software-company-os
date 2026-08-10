# phase06_onboarding_sovereignty_reuse_foundation_shared_baseline

## 1. 文档定位

本文档用于集中冻结 `phase06` 的共享基线，避免相同结论在 `architecture_plan`、`dev_plan`、后续 `/spec` 与根级真相源中重复发散。

> 收口说明：`phase06` 已完成 `/plan -> /spec -> 实现 -> 验收 -> 收口`。本文保留为该阶段 `/plan` 的共享基线与冻结记录，不承担根级当前阶段状态说明；当前阶段状态以 `AGENTS.md`、`plan.md` 与 `docs/README.md` 为准。

## 2. 当前单值基线

### 2.1 项目路线

- 当前项目：`PSCO`
- 当前 phase：`phase06_onboarding_sovereignty_reuse_foundation`
- 当前技术路线：`Durable System Track`
- 当前根级阶段状态：`phase06` 已完成 `/plan -> /spec -> 实现 -> 验收 -> 收口`，当前处于收口后根级同步完成状态
- `phase06` 的规划上游统一以 `PSCO-mvp02-summarize-feedback.md` 与 `phase05` 已交付边界为准

### 2.2 当前阶段唯一直接执行层上游

- 直接执行层上游：
  - `PSCO-mvp02-summarize-feedback.md`
  - `.trae/specs/phase05_10_dashboard_feedback_formal_spec/dashboard_feedback_spec_v0.1.md`
  - `.trae/specs/phase05_11_dashboard_feedback_proto_mainline/`
  - `.trae/specs/phase05_14_dashboard_feedback_integration_validation_acceptance/acceptance_report.md`
- 当前阶段只承接 `phase05` 已冻结并验收的 Dashboard / Feedback、Product / Repository / Module / Decision 主线
- 当前阶段不反向重写 `phase05` 已冻结的 Dashboard 路由、反馈信号与 canonical owner 边界

### 2.3 当前阶段正式技术主线

- Web：`React + Vite + TypeScript`
- Frontend Delivery：单一 `React Web` 客户端，同时覆盖 `PC` 与移动浏览器 UI
- Router：`TanStack Router`
- Data Fetching：`TanStack Query`
- UI：`Tailwind CSS + shadcn/ui`
- Backend：`Go + chi + net/http`
- Database：`PostgreSQL`
- Contract：`Protocol Buffers`
- Contract Tooling：`buf build / lint / generate / breaking`
- Deployment：`Caddy + systemd`
- Runtime Policy：`Single Server First`

### 2.4 当前阶段特别约束

- 当前阶段新增前端写路径必须遵守：
  - 业务写路径唯一 `application` 入口
  - `query` 层纯只读
  - mutation 固定承接位
  - 切片优先与 `shared` 延迟晋升
- 当前阶段新增后端接口必须遵守：
  - `.proto` 是唯一长期合同源
  - `chi + HTTP JSON` 只承担传输适配职责
  - HTTP DTO 只能从 `.proto` 单向派生或显式映射
- 当前阶段不得重新引入 `Feature / Opportunity / Experiment`
- 当前阶段不得引入独立 `AI Assistant` 一级导航
- 当前阶段不得把导出 / 备份写成依赖第三方平台的能力
- 当前阶段不得把 `Capability` 扩写为独立重实体
- 当前阶段不得把模板级复用提前扩写为完整模板系统

### 2.5 当前阶段交付模式

- `phase06` 是交付型 phase，不是纯文档冻结阶段
- 当前 `/plan` 只负责建立阶段上游、任务拆分与共享基线
- 当前阶段后续必须继续进入 `/spec`、源代码实现、验证验收与根级同步
- 当前阶段结束时必须新增可运行、可验收的 Onboarding / Data Sovereignty / Reuse Awareness 最小主线代码

## 3. 当前阶段动作矩阵

`phase06` 最少需要直接承接：

- `StartFirstRunOnboarding`
- `CreateDraftProduct`
- `CreateDraftRepository`
- `CreateDraftModule`
- `CreateDraftDecision`
- `ExportCoreAssets`
- `CreateInstanceBackup`
- `ReadModuleReuseSummary`
- `ReadCapabilitySummary`
- `ViewReuseFeedbackFromDashboardOrDetail`

当前阶段必须打通的最小闭环：

- `first-run -> 最小录入 -> 返回 Dashboard / Detail`
- `existing assets -> export -> verify data can be carried away`
- `current instance -> backup -> verify restore prerequisites`
- `structured assets -> reuse summary -> visible feedback`

允许以最小连接位承接但不扩写为独立主线：

- `Decision` 的低摩擦 capture
- `Capability` 的派生视角
- 后续模板级复用的准备性数据结构

## 4. 当前阶段页面矩阵

- `Onboarding Home / Flow`
- `Dashboard Home`
- `Module Detail`
- `Product Detail`
- `Repository Binding Detail / Workspace`
- `Decision Center / Detail`

### 4.1 当前阶段交互归属矩阵

- `Onboarding`：承接首轮引导、最小录入与继续补全入口
- `Dashboard Home`：继续承接反馈展示、经营入口与复用反馈挂接位
- `Product / Repository / Module / Decision` 的详情与编辑：继续承接 canonical owner 写入
- `Export / Backup`：承接用户带走资产与实例保留入口，不扩写为完整运维中心

补充冻结：

- `Onboarding` 正式业务入口路由：`/onboarding`
- `Export` 正式执行路由：`/dashboard/export`
- `Backup` 正式执行路由：`/dashboard/backup`
- 冷启动空系统主 CTA：`/onboarding`
- 首轮未完成用户回访主 CTA：`Continue Onboarding -> /onboarding`
- `Export / Backup` 允许从 `Dashboard` 动作区进入，但不得在 `Dashboard Home` 主内容区内联完成全部操作
- 首次进入应用时的 cold-start 判定固定由前端根级路由入口守卫承接（`beforeLoad` 或等价根级 loader），不得分散到页面组件 `useEffect` 中各自判断
- `first_run_state = not_started` 时，根级默认进入路径必须回落到 `/onboarding`
- `first_run_state = in_progress` 时，不要求劫持所有 canonical detail 路由；根级默认进入路径与 `Dashboard` 必须提供 `Continue Onboarding` 入口

## 5. 当前阶段数据矩阵

直接承接：

- `modules`
- `releases`
- `products`
- `repositories`
- `decisions`
- `decision_links`
- `product_modules`
- `product_repositories`
- `module_repositories`

当前阶段必须正式消费或新增的派生读取：

- `module_reuse_summary`
- `capability_summary`
- `first_run_state`
- `export_snapshot`
- `backup_snapshot`

### 5.1 最小读写模型

- `first_run_state` 至少承接：
  - `status`：`not_started | in_progress | completed`
  - 是否首次进入
  - 当前引导步骤
  - 首轮完成条件
- `draft-first` 最小写模型至少承接：
  - 草稿创建
  - 后续补全
  - 成功回流落点
- `module_reuse_summary` 至少承接：
  - `module_id`
  - module 被多少 Product 复用
  - `latest_reuse_at`
  - 最小解释文案
- `capability_summary` 至少承接：
  - `capability_key`
  - `capability_label`
  - `supporting_module_count`
  - `latest_capability_update_at`
  - 空状态语义
- `export_snapshot / backup_snapshot` 至少承接：
  - 资产范围
  - 创建时间
  - 创建结果
  - 恢复前提或校验结果

当前阶段关于首轮成功会话的单值定义如下：

- 推荐执行顺序：`Product -> Repository -> Module -> Decision`
- 成功会话成立条件：
  - 一次连续会话内至少各创建 `1` 条已持久化记录：`Product / Repository / Module / Decision`
  - 允许当前阶段先以 `draft-first / partial-entry` 形式存在
  - 会话结束时用户能够回到 `Dashboard` 或任一 canonical owner 页面继续补全
- 当前阶段首轮成功会话不强制要求：
  - `Product` 已完成绑定
  - `Repository` 已绑定 `Product`
  - `Module` 已映射 `Repository`
  - `Decision` 已完成对象链接
- `first_run_state` 的最小状态跃迁冻结为：
  - 尚未开始任何首轮对象写入：`not_started`
  - 已至少创建 `1` 条首轮对象记录、但四类对象未全部持久化：`in_progress`
  - 四类对象均已持久化并满足首轮成功会话条件：`completed`

### 5.2 最小接口归属前提

- `OnboardingRead / OnboardingWrite` 由当前阶段新增 owner 承接
- `ReuseSummaryRead` 由当前阶段新增或扩展 query owner 承接
- `ExportRead / ExportWrite / BackupWrite` 由当前阶段新增 owner 承接
- 当前阶段允许保留 `chi + HTTP JSON` 作为传输层，但不得形成与 `.proto` 并列的第二套合同源

补充冻结：

- `module_reuse_summary` 由当前 canonical 绑定关系读时聚合得到，不引入独立统计表作为唯一事实源
- `module_reuse_summary` 的最小统计口径冻结为“一个 Module 当前被多少 Product 直接复用”
- `capability_summary` 由当前 Module 派生信息读时聚合得到，不引入独立 Capability 重实体
- `capability_summary` 的最小事实来源冻结为 Module 写模型中的轻量 `capability_key`（可空）与系统内置 `capability_label` 映射；未填写 `capability_key` 的 Module 不参与当前阶段 capability 聚合，但不阻断首轮成功会话
- 当前阶段复用感知的新鲜度口径冻结为“读取时反映最新已提交状态”，不引入异步离线刷新前提

## 6. 当前阶段合同与演进基线

- `.proto` 是当前阶段新增接口的唯一合同源
- `buf` 校验链至少覆盖：`build`、`lint`、`generate`、`breaking`
- `.proto` 字段语义必须与 HTTP DTO、前端消费模型保持单值一致
- 合同演进必须遵守兼容性约束；删除字段后必须保留 `reserved` 字段号
- 当前阶段不在共享基线中提前冻结导出文件格式与备份存储介质

当前阶段关于导出 / 备份的最小覆盖矩阵如下：

- `Export` 最小必须覆盖：
  - `products`
  - `modules`
  - `releases`
  - `repositories`
  - `decisions`
  - `decision_links`
  - `product_modules`
  - `product_repositories`
  - `module_repositories`
- `Backup` 的最小覆盖范围不得小于 `Export`
- `Backup` 除核心资产外，至少还必须带出：
  - 当前备份清单或 manifest
  - 备份创建时间
  - 当前实例恢复所需的 schema / 版本前提
- 当前阶段 `backup verified` 的最小成立条件冻结为：
  - 已生成可读取的备份产物
  - 可重新读取并校验备份 manifest
  - manifest 中可见核心资产覆盖矩阵与 schema / 版本恢复前提
  - 不要求当前阶段完成真正 restore 写回

## 7. 当前阶段冷启动与验收基线

- 首轮必须允许用户在一次会话内完成最小 `Product + Repository + Module + Decision` 录入
- 首轮必须允许用户中途退出并继续补全
- 当前阶段必须允许用户导出核心资产并验证导出成功
- 当前阶段必须允许用户触发基础备份并验证恢复前提
- 当前阶段必须允许用户在 Dashboard 或相关详情页看见最小复用 / 能力反馈
- 当前阶段验收不得依赖手工补 SQL 才能建立最小联调环境

补充约束：

- 验收时不得把“只创建了部分对象”算作首轮成功会话
- 验收时不得把“缺少绑定关系导出”算作已完成数据主权闭合
- 验收时必须验证 `module_reuse_summary / capability_summary` 读取的是最新已提交状态
- 验收时必须分别验证：
  - cold-start 用户会被正确导向 `/onboarding`
  - `in_progress` 回访用户不会被强制劫持离开 canonical detail，但能稳定看到 `Continue Onboarding`
  - `backup verified` 基于备份产物与 manifest 读取校验成立，而不是仅以“写出文件成功”代替

## 8. 非目标矩阵

- `Opportunity / Feature / Experiment` 主线
- `Capability` 重实体 CRUD
- 完整模板系统
- GitHub OAuth / 自动导入
- AI 一级工作台
- 自动扫描 / 知识图谱
- Rust Intelligence Layer
- 真实项目 `dry-run` 的正式收口
