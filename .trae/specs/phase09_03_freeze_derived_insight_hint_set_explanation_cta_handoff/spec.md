# phase09-03 冻结派生智能提示集、解释口径与动作承接 Spec

## Why

`phase09-01` 已冻结当前阶段的单一主交付能力与非目标，`phase09-02` 又把模板候选、`Product Create` 预填与成功回流链收敛成单值口径。接下来如果不把派生智能提示的最小集合、解释语义与动作承接继续冻结，后续实现仍会滑向“多几条统计文案”或“后面再决定提示类型与 CTA”的灰区。

## What Changes

- 冻结当前阶段正式提示只包含 `reuse_opportunity_hint` 与 `capability_gap_hint`
- 冻结每类提示的 `trigger -> explanation -> CTA -> target owner` 单值矩阵
- 冻结 `Weekly Review` 与 `Product Create` 各自承接的提示粒度与职责边界
- 冻结 `capability_gap_hint` 对 active template candidate 的依赖关系与无 active candidate 时的成功空态
- 冻结没有稳定 canonical CTA 的提示不得进入 `phase09`
- 冻结提示与既有 `Decision / Product / Module / Review` 动作链的正式对接方式

## Impact

- Affected specs:
  - `phase09_01_freeze_template_reuse_derived_intelligence_scope_success_non_goals`
  - `phase09_02_freeze_template_reuse_assets_candidates_product_create_prefill_chain`
  - `docs/phase/phase09_template_reuse_derived_intelligence_foundation_dev_plan.md`
  - `docs/phase/phase09_template_reuse_derived_intelligence_foundation_architecture_plan.md`
  - `docs/phase/phase09_template_reuse_derived_intelligence_foundation_shared_baseline.md`
  - 后续 `phase09-04 ~ phase09-07` 的合同、owner、页面与验收规格
- Affected code:
  - 后续会影响 `Weekly Review`、`Product Create`、`Module Registry / Module Detail`、模板 handoff 与提示消费位
  - 当前无直接源码改动

## ADDED Requirements

### Requirement: 当前阶段正式提示集合必须保持最小且单值

系统 SHALL 冻结 `phase09` 当前阶段的正式派生提示集合，只允许包含 `reuse_opportunity_hint` 与 `capability_gap_hint` 两类提示。

#### Scenario: 判断当前阶段允许进入实现的提示类型

- **WHEN** 后续 `/spec`、实现或验收描述 `phase09` 的派生智能提示
- **THEN** 只允许存在以下两类正式提示：
  - `reuse_opportunity_hint`
  - `capability_gap_hint`
- **AND** 当前阶段不得新增第三类“泛化智能提示”
- **AND** 不得保留“后面实现时再决定提示类型”的灰区

### Requirement: reuse_opportunity_hint 必须冻结为模板复用机会提示

系统 SHALL 将 `reuse_opportunity_hint` 冻结为“基于既有模板候选与复用事实，指向模板预填创建路径”的最小提示。

#### Scenario: 判断 reuse_opportunity_hint 的事实来源与触发条件

- **WHEN** 系统计算 `reuse_opportunity_hint`
- **THEN** 事实来源必须是模板候选聚合结果与 `module_reuse_summary`
- **AND** 触发条件必须是“当前存在可用模板候选，且候选对应的组合已被至少一个已持久化 Product 证明可复用”

#### Scenario: 判断 reuse_opportunity_hint 的解释与动作承接

- **WHEN** `reuse_opportunity_hint` 在正式消费面中展示
- **THEN** 必须同时给出：
  - 说明该组合为何值得复用的解释文案
  - 进入模板预填创建路径的正式 CTA
  - 指向 `TemplateReuseRead + ProductCreate canonical path` 的目标承接位
- **AND** 不得把该提示退化成纯统计卡片或无动作摘要

#### Scenario: 判断 reuse_opportunity_hint 的成功空态

- **WHEN** 当前不存在可用模板候选，或候选对应组合尚未满足“已被至少一个已持久化 Product 证明可复用”的触发前提
- **THEN** `reuse_opportunity_hint` 必须返回成功空态
- **AND** 不得把“当前没有可复用模板机会”解释成页面错误
- **AND** 不得保留“实现时再决定是否展示提示”的灰区

### Requirement: capability_gap_hint 必须冻结为能力缺口提示

系统 SHALL 将 `capability_gap_hint` 冻结为“基于当前 active template candidate 与 review 作用域能力摘要，提示尚未覆盖的高频 capability”的最小提示。

#### Scenario: 判断 capability_gap_hint 的事实来源与触发条件

- **WHEN** 系统计算 `capability_gap_hint`
- **THEN** 事实来源必须是 `capability_summary` 与当前 active template candidate
- **AND** 触发条件必须是“当前存在 active template candidate，且当前 review 作用域内存在高频 capability 未被该模板候选覆盖”
- **AND** 不得把提示建立在未冻结的 generic review focus 或页面本地猜测上下文上

#### Scenario: 判断 Product Create 中 capability_gap_hint 的模板上下文

- **WHEN** `capability_gap_hint` 在 `Product Create` 中计算
- **THEN** 模板上下文必须来自当前 handoff 进入的 `templateCandidateId` 对应模板候选
- **AND** 必须将该模板候选视为 create 场景下的 selected template candidate
- **AND** 不得继续依赖 `Weekly Review` 会话内的 active candidate 本地状态

#### Scenario: 判断无 active template candidate 时的语义

- **WHEN** 当前不存在 active template candidate
- **THEN** `capability_gap_hint` 必须返回成功空态
- **AND** 不得回退到另一套 generic review focus 口径
- **AND** 不得保留“实现时再决定是否降级”的灰区

#### Scenario: 判断 capability_gap_hint 的解释与动作承接

- **WHEN** `capability_gap_hint` 在正式消费面中展示
- **THEN** 必须同时给出：
  - 缺口来自哪里、为何会阻碍下一次创造的解释文案
  - 进入既有 `Module Registry / Module Detail` canonical 路径继续补齐的正式 CTA
  - 指向 `DerivedInsightRead + Module Registry canonical path` 的目标承接位
- **AND** 不得在提示层内联第二套写动作

### Requirement: 每类提示都必须满足 trigger-explanation-CTA-target owner 四元组

系统 SHALL 要求进入 `phase09` 正式范围的每类提示，都必须具备完整的 `trigger -> explanation -> CTA -> target owner` 四元组。

#### Scenario: 判断某类提示是否有资格进入正式范围

- **WHEN** 后续 `/spec`、实现或验收尝试纳入新的提示文案或提示卡片
- **THEN** 只有同时满足以下条件时，才允许进入正式范围：
  - 触发事实清晰
  - 解释文案清晰
  - canonical CTA 清晰
  - target owner 清晰
- **AND** 若某类提示只有解释、没有稳定 canonical CTA，则该提示不得进入 `phase09` 正式范围

### Requirement: Weekly Review 与 Product Create 的提示职责边界必须单值

系统 SHALL 冻结 `Weekly Review` 与 `Product Create` 的提示职责边界，避免两个页面各自演化出一套提示主线。

#### Scenario: 判断 Weekly Review 的提示职责

- **WHEN** 提示出现在 `Weekly Review`
- **THEN** 页面必须承担模板机会诊断与能力缺口诊断的主要入口职责
- **AND** `reuse_opportunity_hint` 必须以进入模板预填创建路径为主要动作
- **AND** `capability_gap_hint` 必须以引导用户判断下一步补齐方向为主要动作

#### Scenario: 判断 Product Create 的提示职责

- **WHEN** 提示出现在 `Product Create`
- **THEN** 页面只允许消费与当前已选模板相关的解释性提示
- **AND** 当前阶段 `Product Create` 只允许承接 `capability_gap_hint` 这一类提示
- **AND** `Product Create` 不得成长为第二套提示诊断中心或独立智能宿主

### Requirement: 提示必须与既有 canonical 动作链对接

系统 SHALL 冻结派生提示与既有 `Decision / Product / Module / Review` 动作链的对接方式，确保提示不会停留在展示层。

#### Scenario: 判断提示到既有动作链的正式对接

- **WHEN** 用户消费 `reuse_opportunity_hint`
- **THEN** 系统必须把用户导向既有模板 handoff 与 `Product Create` canonical 路径

#### Scenario: 判断能力缺口提示的正式对接

- **WHEN** 用户消费 `capability_gap_hint`
- **THEN** 系统必须把用户导向既有 `Module Registry / Module Detail` canonical 路径继续补齐
- **AND** 后续若需要进入 `Decision`、`Product` 或 `Review reread`，也只能通过既有 canonical 路径继续承接

## MODIFIED Requirements

### Requirement: phase09 成功会话中的 Derived Intelligence 解释

`phase09-03` 修改了 `phase09-01` 中对最小成功会话的解释：当前阶段的 `Derived Intelligence Deepening` 不再只是“页面上多一层解释性文本”，而是必须通过 `reuse_opportunity_hint / capability_gap_hint` 两类最小提示，把解释语义与正式下一步动作稳定绑定起来。

#### Scenario: 判断 Derived Intelligence 是否成立

- **WHEN** 后续 `/spec`、实现或验收判断 `Derived Intelligence Deepening` 是否成立
- **THEN** 必须同时验证提示类型、触发事实、解释文案、CTA 与 target owner 已冻结
- **AND** 不得只用“展示了更多摘要”代替派生智能成立

## REMOVED Requirements

### Requirement: 允许保留未定型提示、泛化智能提示或无稳定 CTA 的提示进入 phase09

**Reason**: 当前阶段目标是交付可运行、可验收的最小动作支撑层，而不是保留一组未来可能成立的智能提示候选池。只要提示类型、触发条件或正式 CTA 仍不稳定，就会把实现阶段重新带回范围漂移。
**Migration**: 后续若需要新增提示类型，必须在新的 `/spec` 中重新证明其 trigger、explanation、CTA 与 target owner 都已单值冻结；否则只允许保留在非正式候选讨论中，不进入 `phase09` 当前实现范围。
