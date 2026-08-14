# phase10-01 冻结 Asset-Action Closure 范围边界、成功标准与非目标 Spec

## Why

`phase10` 已经在三件套中完成第一轮规划冻结，但如果不先把当前 phase 的唯一主交付能力、最小成功标准与明确非目标冻结为单值入口，后续 `/spec -> 实现 -> 验收` 很容易把 `Agent Consumption Layer`、跨项目资产、第五态状态机或局部体验改造偷渡成当前阶段并列主线。
因此，`phase10-01` 必须先把 `Asset-Action Closure` 正式收敛为“让现有资产与动作闭环”的唯一中心主线，作为 `phase10-02+` 的强制边界上游。

## What Changes

- 冻结 `phase10` 的唯一主交付能力为 `Asset-Action Closure`
- 冻结 `phase10` 的阶段成功标准与最小完成条件
- 冻结 `phase10` 与 `Agent Consumption Layer / Cross-Project Convention Asset / 第五态状态机 / 工作流引擎化 / AI 工作台` 的边界
- 冻结 `phase10` 后续 `/spec` 必须直接承接的上游与禁止回退项
- 冻结 `phase10` 的正式消费面与明确排除面
- 冻结 `phase10-02+` 后续 `/spec` 的准入前提

## Impact

- Affected specs:
  - `docs/phase/phase10_asset_action_closure_foundation_architecture_plan.md`
  - `docs/phase/phase10_asset_action_closure_foundation_dev_plan.md`
  - `docs/phase/phase10_asset_action_closure_foundation_shared_baseline.md`
  - 后续 `phase10-02 ~ phase10-12` 的 `/spec`、实现与验收规格
- Affected code:
  - 当前无直接源码改动
  - 后续会影响 `Onboarding / Dashboard / Daily Review / Decision Detail / Product Detail / Module Detail / Repository Detail` 相关切片的实现边界、合同承接与验收口径

## ADDED Requirements

### Requirement: `phase10` 必须保持单一主交付能力

系统 SHALL 冻结 `phase10` 的唯一主交付能力为 `Asset-Action Closure`，并要求后续 `/spec`、实现与验收只围绕“让现有资产主线、review 主线与 detail 主线共同承接真实下一步动作”展开。

#### Scenario: 判断当前 phase 的主交付是否单值

- **WHEN** 后续文档、实现或验收描述 `phase10` 的主交付能力
- **THEN** 必须只承接 `Asset-Action Closure`
- **AND** 必须将 `Onboarding` 首轮建链引导、`Decision` 最小真实生命周期、`Dashboard / Review / Detail pages` 下一步动作承接、`Current Focus / pending signals` 真实经营语义回归视为该主线的组成部分
- **AND** 不得把 `Agent Consumption Layer`、`Cross-Project Convention Asset` 或新的长期业务对象写成当前阶段并列主交付

#### Scenario: 判断当前阶段正式消费面是否被错误扩张

- **WHEN** 后续 `/spec`、实现或验收定义 `phase10` 的正式消费面
- **THEN** 必须只承接 `Onboarding / Dashboard Home / Daily Review / Decision Detail / Product Detail / Module Detail / Repository Detail`
- **AND** 必须明确 `Weekly Review` 当前不是本阶段中心承接位，只允许被动继承“动作语义更清晰”的结果
- **AND** 不得把 `Weekly Review` 扩写为新的复杂动作编排入口

### Requirement: `phase10` 必须冻结最小成功标准与阶段完成条件

系统 SHALL 在 `phase10-01` 中显式冻结阶段成功标准，用于约束后续 `/spec` 不得把“页面上出现新按钮”“局部提示更清晰”或“单点路径可用”误判为阶段完成。

#### Scenario: 判断最小成功标准是否成立

- **WHEN** 后续 `/spec`、实现或验收定义 `phase10` 的成功标准
- **THEN** 必须至少同时覆盖以下结果：
  - `Onboarding` 已从逐段登记升级为首轮建链引导
  - `Decision` 已在既有四态内形成最小但真实的生命周期
  - `Dashboard / Daily Review / Detail pages` 已形成统一的下一步动作承接矩阵
  - `Current Focus / pending signals` 已回到真实 canonical facts
  - 页面、合同、服务、返回链与 reread 规则已收敛到单值结构
  - `Browser Validation -> Root Sync` 已作为阶段收口口径的一部分被正式承接
- **AND** 不得把“单页局部优化”或“单个 CTA 可点击”单独视为阶段成功

#### Scenario: 判断当前阶段的最小闭环是否完整

- **WHEN** 后续 `/spec`、实现或验收描述 `phase10` 的主链闭环
- **THEN** 必须至少同时覆盖以下动作链：
  - `Dashboard -> Onboarding`
  - `Onboarding -> Product / Repository / Module / Decision`
  - `Dashboard / Daily Review -> Decision Detail -> Status Advance`
  - `Detail Page -> Next Step CTA -> Canonical Owner`
  - `Current Focus / pending signals -> reread`
  - `Browser Validation -> Root Sync`
- **AND** 不得把其中任一动作链后移为“后续再补”的可选项

### Requirement: `phase10` 必须冻结正式非目标边界

系统 SHALL 在 `phase10-01` 中显式冻结当前阶段非目标，避免 `mvp0.4` 的后续方向或长期候选能力被提前写成当前阶段实现承诺。

#### Scenario: 判断是否越界到后续阶段或并列主线

- **WHEN** 后续 `/spec`、实现设计或验收用例描述 `phase10` 范围
- **THEN** 必须明确以下内容属于当前阶段非目标：
  - `Agent Consumption Layer`
  - `Cross-Project Convention Asset`
  - 新的长期核心业务实体
  - 第五个 `DecisionStatus`
  - `Onboarding` 工作流引擎化
  - `Dashboard / Review` 任务管理器化
  - `Weekly Review` 作为新的复杂动作编排主入口
  - AI 工作台 / 对话式主入口
  - agent 写入主线
  - 真实连接的重型 GitHub 集成
- **AND** 不得把这些内容作为当前阶段实现承诺、并列子任务或 Done 标准

#### Scenario: 判断当前阶段是否错误回写为局部体验重做

- **WHEN** 后续方案试图把 `phase10` 收窄为局部 UI 调整、按钮补齐或页面美化
- **THEN** 必须判定为偏离当前阶段边界
- **AND** 必须要求动作继续围绕 canonical facts、正式写路径与 reread 结果展开

### Requirement: `phase10` 必须保持上游承接单值且不得回退既有主线

系统 SHALL 要求 `phase10` 后续所有 `/spec` 与实现，直接承接已冻结的 `phase03 / phase04 / phase05 / phase06 / phase08 / phase09` 上游能力与约束，不得重新长出第二套合同、第二套业务主线或对 `Decision`、`Product Create`、`.proto + ConnectRPC` 的地位作重新解释。

#### Scenario: 判断后续 `/spec` 是否正确承接上游

- **WHEN** 后续 `/spec` 描述合同、动作承接、页面职责、返回链或 reread 边界
- **THEN** 必须直接承接：
  - `phase06` 已冻结的 `Onboarding` 与 `first_run_state` 主线
  - `phase03` 已冻结的 `Decision` 中心地位与状态写链
  - `phase04` 已冻结的 `Product / Repository / Module` canonical owner
  - `phase05` 已冻结的 `Dashboard + Feedback` 聚合能力
  - `phase08` 已冻结的 `Operating Review Loop`
  - `phase09` 已冻结的 `Template Reuse + Derived Hints` 支撑结果
  - `phase07` 已冻结的 `.proto + ConnectRPC` 正式传输主线
- **AND** 不得回退为手写 JSON canonical contract
- **AND** 不得新增第二事实源、影子状态表或页面局部“已处理”真相源

#### Scenario: 判断后续 `/spec` 是否继续直接消费 `phase10` 三件套细化矩阵

- **WHEN** 任一 `phase10-02+` 规格进入页面矩阵、动作矩阵、CTA inventory、数据矩阵、caller / owner 或验收前提设计
- **THEN** 除本规格外，还必须继续直接消费：
  - `phase10_asset_action_closure_foundation_shared_baseline.md` 中已冻结的逐步建链矩阵、页面矩阵、下一步动作矩阵、CTA inventory、数据矩阵与验收前提
  - `phase10_asset_action_closure_foundation_architecture_plan.md` 中已冻结的正式消费面、前端交付策略、后端承接策略与业务边界原则
  - `phase10_asset_action_closure_foundation_dev_plan.md` 中已冻结的子任务依赖、DoD 与执行顺序
- **AND** 不得只满足本规格的高层边界，却在具体矩阵层重新长出第二套解释口径

### Requirement: `phase10-02+` 后续 `/spec` 必须以本规格作为强制边界上游

系统 SHALL 要求 `phase10-02+` 的后续 `/spec`，必须以本规格冻结的唯一主交付能力、成功标准、非目标矩阵与上游承接要求作为准入前提，而不是只“参考”三件套。

#### Scenario: 判断后续 `/spec` 是否具备进入实施设计的前提

- **WHEN** 任一 `phase10-02+` 规格尝试进入页面、合同、交互、数据或验收设计
- **THEN** 必须先满足本规格已单值冻结：
  - 当前 phase 的唯一主交付能力
  - 阶段成功标准与最小闭环
  - 正式非目标边界
  - 上游承接与禁止回退项
- **AND** 必须继续直接消费 `phase10` 三件套中已冻结的页面矩阵、动作矩阵、CTA inventory、数据矩阵与验收前提
- **AND** 不得在这些前提未冻结前提前进入并列主线扩写

## MODIFIED Requirements

### Requirement: `phase10` 后续 `/spec` 的入口解释

`phase10-01` 修改了对当前阶段后续 `/spec` 入口的解释：后续任务不再只需要“参考 `phase10` 三件套”，而必须以本规格冻结的范围边界、成功标准与非目标矩阵为强制上游。

#### Scenario: 判断后续 `/spec` 是否可以进入更细设计

- **WHEN** 任一 `phase10-02+` 规格进入动作矩阵、caller / owner、合同、返回链或验收设计
- **THEN** 必须先满足本规格已冻结的边界与非目标前提
- **AND** 不得把后续阶段能力写成当前阶段既成事实

## REMOVED Requirements

### Requirement: 将 `Agent Consumption Layer`、跨项目资产或状态机扩写直接视为 `phase10` 当前实现承诺

**Reason**: `PSCO-mvp04-summarize-feedback.md` 与 `phase10` 三件套都已明确冻结：当前阶段只交付 `Asset-Action Closure`，agent 可消费层与跨项目资产属于后续阶段或候选探索，第五态状态机不进入本阶段承诺。
**Migration**: 后续文档若需要引用这些能力，只允许以“后续依赖、后续进入条件、后续候选方向”表达，不得写成 `phase10` 当前并列主交付、当前阶段 Done 标准或当前 `/spec` 的实现目标。
