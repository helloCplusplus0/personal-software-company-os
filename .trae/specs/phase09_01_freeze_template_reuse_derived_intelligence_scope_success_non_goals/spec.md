# phase09-01 冻结 Template Reuse + Derived Intelligence 范围边界、成功标准与非目标 Spec

## Why

`phase09` 已经在三件套中完成第一轮规划冻结，但如果不先把当前 phase 的单一主交付能力、最小成功会话、成功标准与明确非目标冻结为单值入口，后续 `/spec -> 实现 -> 验收` 很容易重新把模板平台、AI 工作台或 `dry-run` 偷渡为当前阶段并列主线。
因此，`phase09-01` 必须先把 `Template Reuse + Derived Intelligence Deepening` 作为“下一次创造的加速支撑层”这一定位正式收敛，作为 `phase09-02+` 的强制边界上游。

## What Changes

- 冻结 `phase09` 的单一主交付能力为 `Next-Creation Acceleration Support`
- 冻结 `Template Reuse` 与 `Derived Intelligence Deepening` 在当前阶段的正式职责与相互关系
- 冻结 `phase09` 的最小成功会话与阶段成功标准
- 冻结 `phase09` 与 `Operating Review Loop`、`Real-Project Dry-Run`、`AI Context Enhancement`、完整模板平台的边界
- 冻结 `phase09` 后续 `/spec` 必须直接承接的上游、禁止回退项与非目标矩阵

## Impact

- Affected specs:
  - `docs/phase/phase09_template_reuse_derived_intelligence_foundation_architecture_plan.md`
  - `docs/phase/phase09_template_reuse_derived_intelligence_foundation_dev_plan.md`
  - `docs/phase/phase09_template_reuse_derived_intelligence_foundation_shared_baseline.md`
  - 后续 `phase09-02 ~ phase09-12` 的 `/spec`、实现与验收规格
- Affected code:
  - 当前无直接源码改动
  - 后续会影响 `Review / Product Create / Product Detail / Dashboard / ReuseSummary / Module Registry` 相关切片的实现边界与验收口径

## ADDED Requirements

### Requirement: `phase09` 必须保持单一主交付能力

系统 SHALL 冻结 `phase09` 的单一主交付能力为 `Next-Creation Acceleration Support`，并要求后续 `/spec`、实现与验收只围绕“让用户更快开始下一次 `Product` 创建，并获得最小动作支撑提示”展开。

#### Scenario: 判断当前 phase 的主交付是否单值

- **WHEN** 后续文档、实现或验收描述 `phase09` 的主交付能力
- **THEN** 必须只承接 `Template Reuse + Derived Intelligence Deepening` 组成的 `Next-Creation Acceleration Support`
- **AND** 必须将 `Template Reuse` 视为低摩擦创建起点
- **AND** 必须将 `Derived Intelligence Deepening` 视为动作支撑提示层
- **AND** 不得把 `Operating Review Loop`、`Real-Project Dry-Run`、`AI Context Enhancement` 或完整模板平台写成当前阶段并列主交付

### Requirement: `phase09` 必须冻结两类支撑能力的正式职责

系统 SHALL 在 `phase09-01` 中显式冻结 `Template Reuse` 与 `Derived Intelligence Deepening` 的最小职责，避免后续 `/spec` 与实现把它们扩写成新的平台型系统。

#### Scenario: 判断 `Template Reuse` 的职责边界

- **WHEN** 后续 `/spec`、实现或验收描述 `Template Reuse`
- **THEN** 必须至少同时覆盖以下职责：
  - `Module` 组合快照
  - 面向 `Product Create` 的预填辅助
  - 预填后继续编辑并完成创建
- **AND** 不得把模板解释为独立模板实体、模板平台、模板版本系统或参数化模板体系

#### Scenario: 判断 `Derived Intelligence Deepening` 的职责边界

- **WHEN** 后续 `/spec`、实现或验收描述 `Derived Intelligence Deepening`
- **THEN** 必须至少同时覆盖以下职责：
  - `reuse opportunity` 最小提示
  - `capability gap` 最小提示
  - 与 `review / create` 直接相连的解释性指标与动作文案
- **AND** 不得把派生提示扩写为独立智能中心、AI 工作台、长期策略生成器或第二套任务系统

### Requirement: `phase09` 必须冻结最小成功会话与阶段成功标准

系统 SHALL 在 `phase09-01` 中显式冻结最小成功会话，用于约束后续 `/spec` 不得把 `phase09` 误做成“多了几条统计摘要”或“只是能跳转到创建页”的弱结果。

#### Scenario: 判断最小成功会话是否成立

- **WHEN** 后续 `/spec` 描述 `phase09` 的成功标准
- **THEN** 必须至少同时覆盖以下会话：
  - 用户在 `Weekly Review` 中看到模板候选或成功空态
  - 用户在 `Weekly Review` 中看到派生提示或成功空态
  - 用户从模板候选进入 `Product Create` 预填
  - 用户在预填基础上继续编辑并创建成功
  - 用户在 `Product Detail` 中继续读取模板来源摘要并拿到 canonical 下一步动作
- **AND** 不得把“出现候选卡片”或“出现提示文案”单独视为阶段成功

#### Scenario: 判断当前阶段成功标准是否可机械验收

- **WHEN** 后续 `/spec`、实现或验收定义 `phase09` 的成功标准
- **THEN** 必须明确当前阶段成功不能停留在展示增强或单次跳转
- **AND** 必须要求后续 `phase09-02 ~ phase09-04` 把模板候选、模板 handoff、提示触发条件、读写 owner 与成功回流链继续冻结为单值口径
- **AND** 不得保留“实现时再决定”的并列解释口径

### Requirement: `phase09` 必须冻结正式非目标边界

系统 SHALL 在 `phase09-01` 中显式冻结当前阶段非目标，避免 `mvp0.3` 的后续收口或 `mvp0.4+` 候选方向被提前写成当前阶段既成事实。

#### Scenario: 判断是否越界到后续阶段或平台型能力

- **WHEN** 后续 `/spec`、实现设计或验收用例描述 `phase09` 范围
- **THEN** 必须明确以下内容属于当前阶段非目标：
  - `Real-Project Dry-Run`
  - `Venture`
  - `Decision Intelligence`
  - `AI Context Enhancement`
  - 完整模板平台
  - 参数化模板版本体系
  - 独立智能工作台
  - 自动扫描 / 知识图谱
  - 新的长期核心业务实体
- **AND** 不得把这些内容作为本阶段实现承诺、并列子任务或 Done 标准

#### Scenario: 判断当前阶段是否错误回写 `Operating Review Loop`

- **WHEN** 后续方案试图把 `phase09` 解释为重做 `phase08` 主线
- **THEN** 必须判定为偏离当前阶段边界
- **AND** 必须要求 `phase09` 只在 `Weekly Review / Product Create / Product Detail / Dashboard reread` 上增加支撑能力消费位
- **AND** 不得重写 `phase08` 已冻结的 review 会话边界与 `Feedback -> Decision -> Update` 主闭环

### Requirement: `phase09` 必须保持上游承接单值且不回退既有主线

系统 SHALL 要求 `phase09` 后续所有 `/spec` 与实现，直接承接已冻结的 `phase03 ~ phase08` 上游能力与约束，不得重新长出第二套合同、第二套创建主线或对 `Decision` 地位作重新解释。

#### Scenario: 判断后续 `/spec` 是否正确承接上游

- **WHEN** 后续 `/spec` 描述合同、动作承接、页面职责或结果回流边界
- **THEN** 必须直接承接：
  - `phase08` 已冻结并验收的 `Operating Review Loop`
  - `phase06` 已交付的 `reuse_summary / capability_summary`
  - `phase04` 已冻结的 `Product Create` canonical 路径
  - `phase03` 已冻结的 `Decision` 中心地位
  - `phase07` 已冻结的 `.proto + ConnectRPC` 正式传输主线
- **AND** 不得回退为手写 JSON canonical contract
- **AND** 不得把 `Product Create` 替换成第二套模板创建宿主

### Requirement: `phase09-02+` 后续 `/spec` 必须以本规格作为强制边界上游

系统 SHALL 要求 `phase09-02+` 的后续 `/spec`，必须以本规格冻结的单一主交付能力、成功标准、非目标矩阵与上游承接要求作为准入前提，而不是只“参考”三件套。

#### Scenario: 判断后续 `/spec` 是否具备进入实施设计的前提

- **WHEN** 任一 `phase09-02+` 规格尝试进入页面、合同、交互、数据或验收设计
- **THEN** 必须先满足本规格已单值冻结：
  - 当前 phase 的单一主交付能力
  - `Template Reuse / Derived Intelligence Deepening` 的正式职责
  - 最小成功会话
  - 阶段成功标准
  - 正式非目标边界
  - 上游承接与禁止回退项
- **AND** 不得在这些前提未冻结前提前进入并列主线扩写

## MODIFIED Requirements

### Requirement: `phase09` 后续 `/spec` 的入口解释

`phase09-01` 修改了对当前阶段后续 `/spec` 入口的解释：后续任务不再只需要“参考 `phase09` 三件套”，而必须以本规格冻结的范围边界、成功标准与非目标矩阵为强制上游。

#### Scenario: 判断后续 `/spec` 是否可以进入更细设计

- **WHEN** 任一 `phase09-02+` 规格进入模板候选、提示矩阵、读写 owner、页面流或验收设计
- **THEN** 必须先满足本规格已冻结的边界与非目标前提
- **AND** 不得把未来阶段能力写成当前阶段既成事实

## REMOVED Requirements

### Requirement: 将 `dry-run`、AI 增强或平台型模板能力直接视为 `phase09` 当前实现承诺

**Reason**: `PSCO-mvp03-summarize-feedback.md` 与 `phase09` 三件套都已明确冻结：当前阶段只交付 `Template Reuse + Derived Intelligence Deepening` 的最小支撑能力，`dry-run` 仍是后续独立验收闸，AI 与平台化能力属于后移方向。
**Migration**: 后续文档若需要引用这些能力，只允许以“后续依赖、后续进入条件、后续候选方向”表达，不得写成 `phase09` 当前并列主交付或当前阶段 Done 标准。
