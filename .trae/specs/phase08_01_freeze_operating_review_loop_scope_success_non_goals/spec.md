# phase08-01 冻结 Operating Review Loop 范围边界、成功标准与非目标 Spec

## Why

`phase08` 是 `mvp0.3` 的首个正式业务 phase，但当前仓库刚从 `phase07` 的传输主线收口进入业务阶段，最容易发生的问题不是“做不出功能”，而是把 `Operating Review Loop` 做散、做宽，或者把后续支撑能力提前偷渡进来。  
因此，`phase08-01` 必须先把当前 phase 的唯一中心主线、最小成功会话、阶段成功标准与明确非目标冻结为单值入口，作为后续 `/spec -> 实现 -> 验收` 的唯一边界上游。

## What Changes

- 冻结 `phase08` 的唯一中心主线为 `Operating Review Loop`
- 冻结 `phase08` 的最小成功会话与阶段成功标准
- 冻结 `phase08` 与 `Template Reuse / Derived Intelligence Deepening / Real-Project Dry-Run` 的边界
- 冻结 `phase08` 不得演化为通用任务管理器的约束
- 冻结 `phase08` 后续 `/spec` 必须直接承接的上游与禁止回退项

## Impact

- Affected specs:
  - `phase08_operating_review_loop_foundation_architecture_plan.md`
  - `phase08_operating_review_loop_foundation_dev_plan.md`
  - `phase08_operating_review_loop_foundation_shared_baseline.md`
  - 后续 `phase08-02 ~ phase08-12` 的 `/spec` 与验收规格
- Affected code:
  - 当前无直接源码改动
  - 后续会影响 `Dashboard / Review / Decision / Product / Module / Repository` 相关切片的实现边界与验收口径

## ADDED Requirements

### Requirement: `phase08` 必须保持单一中心主线

系统 SHALL 冻结 `phase08` 的唯一中心主线为 `Operating Review Loop`，并要求后续 `/spec`、实现与验收只围绕 `Dashboard -> Review -> Decision -> Update` 最小经营回路展开。

#### Scenario: 判断当前 phase 的主交付是否单值

- **WHEN** 后续文档或实现描述 `phase08` 的主交付能力
- **THEN** 必须只承接 `Operating Review Loop`
- **AND** 必须将 `Dashboard Review Flow`、`Feedback -> Decision -> Update Action Chain`、`Review Result Writeback` 视为该主线的组成部分
- **AND** 不得把 `Template Reuse / Derived Intelligence Deepening / Real-Project Dry-Run` 写成当前阶段的并列主交付

### Requirement: `phase08` 必须冻结最小成功会话与阶段成功标准

系统 SHALL 在 `phase08-01` 中显式冻结最小成功会话，用于约束后续 `/spec` 不得只停留在“多了一个 review 区块”的展示层结果。

#### Scenario: 判断最小成功会话是否成立

- **WHEN** 后续 `/spec` 描述 `phase08` 的成功标准
- **THEN** 必须至少同时覆盖以下会话：
  - 用户从 `Dashboard` 进入 review
  - review 汇总当前焦点、代表性反馈与待处理决策形成统一动作入口
  - 用户从 review 正式进入 `Decision`
  - review 结果回流既有实体或既有 canonical action handoff
- **AND** 不得把“页面已出现 review 入口”单独视为阶段成功

#### Scenario: 判断 daily / weekly 是否被当成真实双路径

- **WHEN** 后续 `/spec`、实现或验收定义 daily / weekly review
- **THEN** 必须将两者视为同一主线下的两条独立成功会话
- **AND** 不得用同一套数据装配、标题语义、完成定义与验收口径冒充双路径

#### Scenario: 判断当前阶段是否正式消费既有反馈与复用感知

- **WHEN** 后续 `/spec`、实现或验收定义 `phase08` 的成功标准
- **THEN** 必须把 `phase05` Feedback 与 `phase06` Reuse Awareness 的正式消费纳入成功判定
- **AND** 必须明确 weekly review 至少承接 `reuse snapshot`
- **AND** 不得把 `module_reuse_summary / capability_summary` 继续留在“可接可不接”的可选状态

### Requirement: `phase08` 必须冻结正式非目标边界

系统 SHALL 在 `phase08-01` 中显式冻结当前阶段非目标，避免 `mvp0.3` 的后续支撑能力被提前写成当前阶段既成事实。

#### Scenario: 判断是否越界到后续支撑能力

- **WHEN** 后续 `/spec`、实现设计或验收用例描述 `phase08` 范围
- **THEN** 必须明确以下内容属于当前阶段非目标：
  - `Template Reuse`
  - `Derived Intelligence Deepening`
  - `Real-Project Dry-Run`
  - 新的长期核心实体
  - 通用任务系统 / backlog 系统
- **AND** 不得把这些内容作为本阶段实现承诺、并列子任务或 Done 标准

#### Scenario: 判断 review loop 是否被扩写成通用任务管理器

- **WHEN** 后续方案试图在 review 中引入独立任务池、并列状态体系或 review-local 实体写路径
- **THEN** 必须判定为偏离当前阶段边界
- **AND** 必须要求动作继续围绕 `Feedback / Decision / Update` 闭环

### Requirement: `phase08` 必须保持上游承接单值且不回退既有主线

系统 SHALL 要求 `phase08` 后续所有 `/spec` 与实现，直接承接已冻结的 `phase03 ~ phase07` 上游能力与约束，不得重新长出第二套合同、第二套业务主线或对 `Decision` 地位作重新解释。

#### Scenario: 判断后续 `/spec` 是否正确承接上游

- **WHEN** 后续 `/spec` 描述合同、动作承接或实体回流边界
- **THEN** 必须直接承接：
  - `phase07` 已冻结的 `.proto + ConnectRPC` 正式传输主线
  - `phase05` 的 Dashboard + Feedback 基础能力
  - `phase06` 的 Reuse Awareness 已交付能力
  - `phase03 / phase04` 已冻结的 `Decision / Product / Repository / Module` canonical owner
- **AND** 不得回退为手写 JSON canonical contract
- **AND** 不得弱化 `Decision` 在经营回路中的中心地位

### Requirement: `phase08` 后续 `/spec` 必须直接消费已冻结的真实 inventory

系统 SHALL 要求 `phase08-02+` 后续 `/spec`，直接消费 `phase08_operating_review_loop_foundation_shared_baseline.md` 已冻结的真实 `route / page / caller / query owner / application owner` inventory，而不是只停留在理想化页面与动作描述。

#### Scenario: 判断后续 `/spec` 是否强制消费真实调用方

- **WHEN** 任一 `phase08-02+` 规格进入页面、路由、caller、owner 或回流设计
- **THEN** 必须逐项说明：
  - 哪些 caller 被复用
  - 哪些 caller 被升级为 review 正式入口
  - 哪些 caller 保持只读跳转
  - 哪些 owner 被扩展
  - 哪些 owner 明确禁止改成页面内联编排
- **AND** 必须把 `DashboardRoute / DashboardHomePage / DashboardPrimaryActionPanel / FeedbackSignalCard / RecentActivityItemCard / DashboardStatBar / Decision route 宿主 / Decision 写入 owner / dashboard-source` 等已冻结 inventory 视为后续 `/spec` 的强制输入
- **AND** 不得重新绕开这份 inventory 另写一套理想化入口

## MODIFIED Requirements

### Requirement: `phase08` 后续 `/spec` 的收敛入口解释

`phase08-01` 修改了对当前阶段后续 `/spec` 入口的解释：后续任务不再只需要“参考三件套”，而必须以本规格冻结的主线、成功标准与非目标为强制边界。

#### Scenario: 判断后续 `/spec` 是否可进入实施设计

- **WHEN** 任一 `phase08-02+` 规格尝试进入页面、合同、交互或验收设计
- **THEN** 必须先满足本规格已单值冻结：
  - `phase05` Feedback 与 `phase06` Reuse Awareness 的正式消费要求
  - 当前 phase 唯一中心主线
  - 最小成功会话
  - daily / weekly 双路径口径
  - 正式非目标边界
  - 真实 caller / route / owner inventory 的强制消费要求
  - 上游承接与禁止回退项

## REMOVED Requirements

### Requirement: 将 `mvp0.3` 其余主轴直接视为 `phase08` 当前实现承诺

**Reason**: `PSCO-mvp03-summarize-feedback.md` 已明确冻结 `Operating Review Loop` 为当前中心主线，`Template Reuse / Derived Intelligence Deepening / Real-Project Dry-Run` 只允许作为后续支撑能力或独立验收闸承接。  
**Migration**: 后续文档若需要引用这些能力，只允许以“后续依赖、后续进入条件、后续消费场景”表达，不得写成 `phase08` 当前并列主交付。
