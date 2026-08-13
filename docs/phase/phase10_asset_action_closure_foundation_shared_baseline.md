# phase10_asset_action_closure_foundation_shared_baseline

## 1. 文档定位

本文档用于集中冻结 `phase10` 的共享基线，避免相同结论在 `architecture_plan`、`dev_plan`、后续 `/spec` 与根级真相源中重复发散。

`phase10` 当前处于 `/plan` 阶段。本文档只承接当前阶段的单值基线、能力矩阵、动作矩阵与验收前提，不替代后续 `/spec`、实现或根级状态正文。

## 2. 当前单值基线

### 2.1 项目路线

- 当前项目：`PSCO`
- 当前 phase：`phase10_asset_action_closure_foundation`
- 当前技术路线：`Durable System Track`
- 当前根级阶段状态：`phase10` 已建立正式 `/plan` 入口，作为 `mvp0.4` 的首个正式业务 phase
- 当前阶段规划上游统一以 `PSCO-mvp04-summarize-feedback.md`、`fix_001_003` 收口记录、`phase06 / phase08 / phase09` 已交付主线为准

### 2.2 当前阶段唯一直接执行层上游

- 直接执行层上游：
  - `PSCO-mvp04-summarize-feedback.md`
  - `docs/fix/fix_001_003_mvp03_dry_run_rerun_closure.md`
  - `docs/fix/fix_001_onboarding_cold_start_state_analysis.md`
  - `docs/fix/fix_002_decision_pending_signal_semantics_analysis.md`
  - `docs/fix/fix_003_decision_detail_status_advance_analysis.md`
  - `.trae/specs/phase06_12_onboarding_sovereignty_reuse_formal_spec/spec.md`
  - `.trae/specs/phase03_10_decision_center_formal_spec/decision_center_spec_v0.1.md`
  - `.trae/specs/phase04_10_product_repository_binding_formal_spec/product_repository_binding_spec_v0.1.md`
  - `.trae/specs/phase05_10_dashboard_feedback_formal_spec/dashboard_feedback_spec_v0.1.md`
  - `.trae/specs/phase08_11_validate_review_loop_integration_browser_regression_acceptance/acceptance_report.md`
  - `.trae/specs/phase09_11_validate_template_reuse_derived_hint_integration_browser_regression_acceptance/acceptance_report.md`
- 当前阶段只承接 `Asset-Action Closure`
- 当前阶段不反向重写 `phase08 / phase09` 已完成的正式交付

### 2.3 当前阶段正式技术主线

- Web：`React + Vite + TypeScript`
- Frontend Delivery：单一 `React Web`
- Router：`TanStack Router`
- Data Fetching：`TanStack Query`
- UI：`Tailwind CSS + shadcn/ui`
- Backend：`Go + chi + net/http + ConnectRPC`
- Database：`PostgreSQL`
- Contract：`Protocol Buffers`
- Contract Tooling：`buf build / lint / generate / breaking`
- Deployment：`Caddy + systemd`
- Runtime Policy：`Single Server First`

### 2.4 当前阶段特别约束

- 当前阶段新增或演进后的业务接口必须遵守：
  - `.proto` 是唯一长期合同源
  - `ConnectRPC` 是业务接口正式传输层
  - `chi` 只承担 router shell、middleware 与非业务端点承载职责
- 当前阶段新增或演进后的前端业务动作必须遵守：
  - 写路径唯一 `application` 入口
  - `query` 层纯只读
  - mutation 固定承接位
- 当前阶段不允许：
  - 把 `Onboarding` 做成工作流引擎
  - 把 `Decision` 扩写为第五态状态机
  - 把 `Dashboard / Review` 写成任务管理器
  - 把 `Agent Consumption Layer` 偷渡为并列主交付
  - 引入新的长期核心实体

### 2.5 当前阶段交付模式

- `phase10` 是交付型 phase，不是纯文档冻结阶段
- 当前 `/plan` 只负责建立阶段上游、任务拆分与共享基线
- 当前阶段后续必须继续进入 `/spec`、源代码实现、验证验收与根级同步
- 当前阶段结束时必须新增可运行、可验收的 `Asset-Action Closure` 主线代码

## 3. 当前阶段能力矩阵

### 3.1 Asset-Action Closure 单值定义

`Asset-Action Closure` 在当前阶段冻结为：

- `Onboarding` 的首轮建链引导
- `Decision` 的最小真实生命周期
- Dashboard / Review / Detail pages 的下一步动作承接
- `Current Focus / pending signals` 的真实经营语义回归

当前阶段不把以下内容解释为 `Asset-Action Closure`：

- `Agent Consumption Layer`
- `Cross-Project Convention Asset`
- 新的长期业务对象
- AI 增强工作台

### 3.2 Onboarding 单值定义

`Onboarding` 在当前阶段冻结为：

- 保留既有六段式主线：`welcome / product / repository / module / decision / complete`
- 在创建过程中尽量承接最小建链动作，而不是把所有关系留给成功后手动补链
- 成功后仍回到既有 canonical `Product / Repository / Module / Decision` 详情页与列表页主线

当前阶段不把以下内容解释为 `Onboarding`：

- 可配置工作流引擎
- 第二套独立 draft 实体
- 与 canonical owner 并列的新写路径

### 3.3 Decision 生命周期单值定义

当前阶段 `Decision` 生命周期继续冻结为既有四态：

- `proposed`
- `active`
- `superseded`
- `archived`

最小迁移矩阵冻结如下：

- `proposed -> active`
- `proposed -> superseded`
- `proposed -> archived`
- `active -> superseded`
- `active -> archived`
- `superseded` 与 `archived` 为当前阶段终态

补充冻结：

- `Decision Detail` 是正式状态推进承接位
- `pending decision` 的判定继续完全锚定 `Decision.status`
- `decision_links` 与 `review_records` 继续不是退出 pending 的代理条件

### 3.4 当前阶段必须直接承接的最小闭环

当前阶段最少需要直接承接：

- `Dashboard -> Onboarding`
- `Onboarding -> Product / Repository / Module / Decision`
- `Dashboard / Daily Review -> Decision Detail -> Status Advance`
- `Detail Page -> Next Step CTA -> Canonical Owner`
- `Current Focus / pending signals -> reread`
- `Browser Validation -> Root Sync`

允许以最小连接位承接但不扩写为独立主线：

- `Weekly Review` 的动作语义继承
- `Template Reuse / Derived Hints` 的既有支撑能力结果
- 后续 `Agent Consumption Layer` 的依赖表达

## 4. 当前阶段页面矩阵

- `Dashboard Home`
- `Onboarding`
- `Daily Review`
- `Decision Detail`
- `Product Detail`
- `Module Detail`
- `Repository Detail`

### 4.1 当前阶段交互归属矩阵

- `Dashboard Home`：承接首轮进入、当前焦点与跨页面“下一步动作”入口
- `Onboarding`：承接首轮建链引导与成功后 handoff
- `Daily Review`：承接“今天需要推进什么”的最小动作分诊
- `Decision Detail`：承接 `Decision` 生命周期推进
- `Product Detail / Module Detail / Repository Detail`：承接各自实体的下一步动作 CTA 与返回 reread

补充冻结：

- `Weekly Review` 当前不承担新的复杂动作编排主入口
- 页面不得各自拼装第二套 mutation、失效刷新与成功回流语义
- 各 detail page 若没有正式下一步动作，不得用“返回列表页”作为长期兜底

### 4.2 当前阶段下一步动作矩阵

当前阶段所有关键页面都必须至少回答一个问题：**“下一步做什么？”**

最小动作矩阵冻结如下：

- `Dashboard Home`
  - 空态或近空态：进入 `Onboarding`
  - 存在 pending decision：进入 `Decision Detail`
  - 存在明确结构缺口：进入对应 canonical detail / list owner

- `Onboarding`
  - 当前步完成：进入下一步建链动作
  - 所有步骤完成：进入 `Dashboard`

- `Daily Review`
  - current focus 命中 `Decision`：进入 `Decision Detail`
  - current focus 命中实体结构缺口：进入对应 canonical owner

- `Decision Detail`
  - `proposed`：可推进为 `active / superseded / archived`
  - `active`：可推进为 `superseded / archived`
  - 终态：不再展示状态推进 CTA，只保留结果消费与回流动作

- `Product / Module / Repository Detail`
  - 必须存在至少一个可执行的“下一步动作”CTA
  - CTA 只能指向既有 canonical path，不得长出局部旁路

## 5. 当前阶段数据矩阵

直接承接：

- `first_run_state`
- `dashboard overview`
- `feedback signals`
- `daily review context`
- `decision detail`
- `product / module / repository detail`

当前阶段必须正式消费或新增的动作相关读取：

- `next-step action signals`
- `Onboarding` 建链引导上下文
- `Decision` 状态推进后的 reread 结果
- 各 detail page 的 CTA 支撑读模型

### 5.1 最小合同与读写模型前提

- `Onboarding`
  - 继续复用既有 `first_run_state` 与相关 create / bind canonical path
  - 若需新增合同，只允许直接服务首轮建链引导

- `Decision`
  - 继续复用既有 `Decision.status`
  - 状态推进继续走 `.proto -> Connect -> service -> store` 正式写链

- `Dashboard / Review / Current Focus`
  - 读侧只消费 canonical status 与既有事实源
  - 不新增第二套“已处理”局部状态字段

- `Detail pages`
  - CTA 只导向既有 canonical path
  - 若需新增读模型，只承接“下一步动作”解释与回流辅助

### 5.2 Reread 单值规则

当前阶段所有动作完成后，都必须能在 reread 中回答：

1. 页面当前是否还需要推进这件事？
2. 如果不需要，系统是否已经停止误报？
3. 如果仍需要，是否给出新的正式动作入口？

补充冻结：

- reread 语义只能基于 canonical facts
- 不允许通过页面局部隐藏、前端临时过滤或 toast 假象掩盖未完成状态

## 6. 当前阶段验收前提

当前阶段浏览器验收至少必须能机械验证：

1. 用户从空态或近空态进入时，能否顺畅完成首轮建链，而不是完成一堆孤立登记
2. `Decision` 从 `proposed` 推进后，`Dashboard / Daily Review / Current Focus` 是否与 detail 状态一致
3. `Product / Module / Repository Detail` 是否都能明确回答“下一步做什么”
4. 任一关键动作完成后，返回 Dashboard / Review / Detail reread 时，是否能看懂刚刚发生了什么
5. 阶段结果是否仍保持：
   - 无第五态
   - 无第二事实源
   - 无 AI 工作台
   - 无第二套路由 / 状态管理 / UI 事实源
