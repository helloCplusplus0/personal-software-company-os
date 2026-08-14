# phase10-02 冻结 Onboarding 从首轮登记到首轮建链引导的语义与动作矩阵 Spec

## Why

`phase10-01` 已经冻结了 `Asset-Action Closure` 的范围边界、成功标准与非目标，但如果不继续把 `Onboarding` 从 `phase06` 的“首轮登记入口”升级为 `phase10` 的“首轮建链引导”，后续实现仍会回到“先孤立创建实体，再让用户去多个 detail page 手工补链”的旧路径。  
因此，`phase10-02` 必须先把六段式主线、逐步建链矩阵、canonical handoff、恢复语义与唯一主上下文锚点冻结成单值结论，作为 `phase10-04 / 06 / 07 / 08` 的直接上游。

## What Changes

- 冻结 `Onboarding` 六段式主线在 `phase10` 中的保留边界
- 冻结 `welcome / product / repository / module / decision / complete` 的逐步建链矩阵
- 冻结哪些关系必须在创建过程中直接承接，哪些关系允许延后到 canonical detail handoff
- 冻结 `Onboarding` 与既有 `Product / Repository / Module / Decision` canonical 写路径的承接关系
- 冻结 `current_product_id` 作为唯一主上下文锚点的恢复语义
- 冻结中途中断、返回再进入时的单值恢复规则
- 冻结当前阶段不引入第二套草稿系统、不把 `Onboarding` 演化为工作流引擎

## Impact

- Affected specs:
  - `phase10_01_freeze_asset_action_closure_scope_success_non_goals`
  - `docs/phase/phase10_asset_action_closure_foundation_dev_plan.md`
  - `docs/phase/phase10_asset_action_closure_foundation_shared_baseline.md`
  - `docs/phase/phase10_asset_action_closure_foundation_architecture_plan.md`
  - 后续 `phase10-04 / phase10-06 / phase10-07 / phase10-08` 的 `/spec` 与实现规格
  - 既有 `phase06_01_define_first_run_onboarding_boundary_entry`
  - 既有 `phase06_06_design_onboarding_frontend_route_interaction_flow`
- Affected code:
  - 当前无直接源码改动
  - 后续会影响 `frontend/src/features/onboarding/`
  - 后续会影响 `frontend/src/routes/onboarding.tsx` 与相关返回链
  - 后续会影响 `backend/internal/onboarding/*`
  - 后续会影响 `Product / Repository / Module / Decision` 相关 canonical create / bind handoff

## ADDED Requirements

### Requirement: `Onboarding` 在 `phase10` 中必须保持六段式单一主线

系统 SHALL 将 `Onboarding` 在 `phase10` 中冻结为单一六段式主线：`welcome / product / repository / module / decision / complete`，并要求后续 `/spec` 与实现继续在这一条主线上完成首轮建链引导，而不是长出并列步骤体系或第二条 onboarding 宿主。

#### Scenario: 判断 `Onboarding` 主线是否仍为单值

- **WHEN** 后续 `/spec`、实现或验收描述 `Onboarding` 的页面与步骤主线
- **THEN** 必须只承接 `welcome / product / repository / module / decision / complete`
- **AND** 不得并行新增第二套 `draft workspace`、可配置流程分支或多宿主 onboarding 体系
- **AND** 不得把 `Onboarding` 回退为只解释字段录入顺序的六段说明页

### Requirement: `Product` 必须作为首轮建链的唯一主上下文锚点

系统 SHALL 冻结 `Product` 为 `Onboarding` 首轮建链的唯一主上下文锚点，用于统一解释后续 `repository / module / decision` 三步的默认承接、恢复语义与 reread 结果。

#### Scenario: 判断首轮建链的主上下文归属

- **WHEN** 后续 `/spec`、实现或验收判断 `Onboarding` 当前主上下文属于哪个 canonical owner
- **THEN** 必须得到 `Product` 的单值结论
- **AND** `current_product_id` 必须作为 `Onboarding` 恢复语义的唯一主上下文锚点
- **AND** `Repository / Module / Decision` 只允许作为 step 级结果与辅助恢复线索
- **AND** 不得把主上下文解释为“当前最新创建的任意实体”

### Requirement: `welcome` 步骤只负责正式进入首轮建链主线

系统 SHALL 冻结 `welcome` 为无实体创建、无并列写路径的引导步骤，它只负责让用户正式进入 `product` 步，而不是在欢迎页预创建实体或并行选择多条 onboarding 路径。

#### Scenario: 判断 `welcome` 步骤职责

- **WHEN** 用户进入 `Onboarding` 的 `welcome` 步骤
- **THEN** 页面只允许展示首轮建链引导与进入下一步动作
- **AND** 不得在本步直接创建任何 canonical 实体
- **AND** 成功动作必须单值进入 `product`

### Requirement: `product` 步骤必须冻结首轮建链主上下文

系统 SHALL 冻结 `product` 步骤的 canonical owner 为 `Product`，其最小动作是创建首个 `Product`，并将该 `Product` 冻结为当前首轮建链主上下文。

#### Scenario: 判断 `product` 步骤的最小动作与默认关系

- **WHEN** 用户完成 `product` 步骤
- **THEN** 系统必须创建首个 `Product`
- **AND** 必须将当前 `Product` 冻结为首轮建链主上下文
- **AND** 不要求在本步完成全部跨实体关系
- **AND** 成功后默认下一步必须进入 `repository`

### Requirement: `repository` 步骤必须优先承接与当前 `Product` 的最小正式绑定

系统 SHALL 冻结 `repository` 步骤的 canonical owner 为 `Repository`，其最小动作是创建或登记首个 `Repository`，并优先承接与当前 `Product` 主上下文的最小正式绑定。

#### Scenario: `repository` 可在当前承接位直绑

- **WHEN** 用户完成 `repository` 步骤，且当前承接位允许同页直绑到当前 `Product`
- **THEN** 系统必须在本步直接承接 `Repository -> Product` 的最小正式绑定
- **AND** 成功后默认下一步必须进入 `module`

#### Scenario: `repository` 无法在当前承接位直绑

- **WHEN** 用户完成 `repository` 步骤，但当前承接位不允许同页直绑到当前 `Product`
- **THEN** 系统必须提供单值 canonical handoff 到正式 binding path
- **AND** 不得沉默地把关系留给用户日后猜测
- **AND** 成功后默认下一步仍必须进入 `module`

### Requirement: `module` 步骤必须优先承接与当前 `Product` 的最小正式关系

系统 SHALL 冻结 `module` 步骤的 canonical owner 为 `Module`，其最小动作是创建首个 `Module`，并优先承接与当前 `Product` 主上下文的最小正式关系。

#### Scenario: 判断 `module` 步骤的默认关系与延后关系

- **WHEN** 用户完成 `module` 步骤
- **THEN** 系统必须优先承接 `Module -> Product` 的最小正式关系
- **AND** 若 `Repository` 映射尚未闭合，只允许将其保留为后续 detail CTA 或 canonical handoff
- **AND** 不得在本步并列展开第二条以 `Repository` 为中心的动作主线
- **AND** 成功后默认下一步必须进入 `decision`

#### Scenario: `module` 无法在当前承接位完成最小正式关系

- **WHEN** 用户完成 `module` 步骤，但当前承接位不允许同页完成 `Module -> Product` 的最小正式关系
- **THEN** 系统必须提供单值 canonical handoff 到正式 binding path
- **AND** 不得沉默地将该关系留为未解释的悬空状态
- **AND** `Repository` 映射若同时未闭合，只允许继续保持为从属于该步的 detail CTA 或 canonical handoff，不得因此长出并列第二主线
- **AND** 成功后默认下一步仍必须进入 `decision`

### Requirement: `decision` 步骤必须与当前 `Product` 形成最小正式承接

系统 SHALL 冻结 `decision` 步骤的 canonical owner 为 `Decision`，其最小动作是创建首个 `Decision`，并与当前 `Product` 主上下文形成最小正式承接，使该决策在后续 reread 中可被解释。

#### Scenario: `decision` 可在当前承接位完成最小正式承接

- **WHEN** 用户完成 `decision` 步骤，且当前承接位允许同页完成最小正式承接
- **THEN** 系统必须将该 `Decision` 与当前 `Product` 主上下文形成最小正式承接
- **AND** 成功后默认下一步必须进入 `complete`

#### Scenario: `decision` 无法在当前承接位完成最小正式承接

- **WHEN** 用户完成 `decision` 步骤，但当前承接位不允许同页完成最小正式承接
- **THEN** 系统必须提供单值 canonical handoff
- **AND** 该 handoff 必须足以保证该 `Decision` 能在后续 reread 中被解释
- **AND** 成功后默认下一步仍必须进入 `complete`

### Requirement: `complete` 步骤只承接 onboarding 结束与返回 `Dashboard`

系统 SHALL 冻结 `complete` 为 `Onboarding` 的结束步骤，它只承接首轮建链结束与返回 `Dashboard`，不再在此长出新的并列写路径或第二条动作主线。

#### Scenario: 判断 `complete` 步骤职责

- **WHEN** 用户进入 `complete`
- **THEN** 页面必须只承接 onboarding 结束、结果回看与返回 `Dashboard`
- **AND** 不得在该步骤新增并列写路径

#### Scenario: 判断 `complete` 是否允许存在未即时完成的 handoff

- **WHEN** 用户从 `repository / module / decision` 步骤进入 `complete`
- **THEN** 系统允许仍存在尚未在当前步骤即时完成的 canonical handoff
- **AND** 前提是每一项未被当前步骤直承接的关系，都已经被单值解释为明确的 canonical handoff 或 detail CTA
- **AND** 不得仍然存在未被解释的悬空关系
- **AND** 不得把“未来再补”但当前没有明确落点的关系视为允许进入 `complete` 的状态

### Requirement: `Onboarding` 必须继续复用既有 canonical 写路径

系统 SHALL 冻结 `Onboarding` 与既有 `Product / Repository / Module / Decision` canonical 写路径的承接关系，要求后续实现继续复用既有 create / bind canonical path，而不是在 onboarding 内长出第二套并列写路径。

#### Scenario: 判断 `Onboarding` 的写路径归属

- **WHEN** 后续 `/spec` 或实现设计 `Onboarding` 的写动作承接位
- **THEN** 必须继续回到既有 canonical `Product / Repository / Module / Decision` owner
- **AND** 不得为 onboarding 单独新增并列的长期写路径
- **AND** 不得让页面组件自行拼装第二套 mutation 语义

### Requirement: 无法即时完成的关系必须通过单值 canonical handoff 承接

系统 SHALL 冻结一个统一规则：凡是当前步骤不能安全完成的正式绑定关系，都必须以单值 canonical handoff 承接，而不是留给用户后续自行猜测。

#### Scenario: 判断某一步关系是否允许延后

- **WHEN** 某一建链关系无法在当前步骤安全完成
- **THEN** 系统只允许将其延后到单值 canonical handoff 或后续 detail CTA
- **AND** 必须显式回答“下一步去哪里补”
- **AND** 不得把“未来可能再补”作为长期兜底解释

### Requirement: 中途中断恢复必须依赖 canonical facts 与单值恢复读模型

系统 SHALL 冻结 `Onboarding` 的恢复语义：中途中断、返回再进入时，只允许依据 canonical facts 与单值 onboarding 恢复读模型恢复当前主线，不允许长出第二套草稿系统。

#### Scenario: 判断恢复时的真相源

- **WHEN** 用户在 `Onboarding` 中途退出后再次进入
- **THEN** 恢复逻辑必须依赖 canonical facts 与单值 onboarding 恢复读模型
- **AND** 不得引入第二套持久化 draft 真相源
- **AND** 不得按页面局部缓存、前端临时状态或“最近点过哪个按钮”来判断当前主线

#### Scenario: 显式 step 返回线索与恢复读模型冲突时的优先级

- **WHEN** 用户从 canonical detail handoff 返回 `Onboarding`，且携带显式 step 返回线索（如 `onboardingStep` 或等价 hint）
- **AND** 该线索与恢复读模型推导出的“最近未完成步骤”不一致
- **THEN** 系统必须优先采用显式 step 返回线索，前提是它仍隶属于当前 `current_product_id` 所锚定的主上下文
- **AND** 若该显式线索已失效、越过必需前置步骤或不再隶属于当前主上下文，才允许回退为恢复读模型推导的最近未完成步骤
- **AND** 不得在两者之间做不透明的隐式猜测

#### Scenario: 多实体并存时恢复当前 onboarding 主线

- **WHEN** 当前系统中存在多个 `Product / Repository / Module` 并存
- **THEN** 恢复逻辑必须优先回到 `current_product_id` 所锚定的最近未完成步骤
- **AND** 不得按“全局最新实体”猜测当前 onboarding 主线

### Requirement: step 级辅助恢复线索必须保持从属地位

系统 SHALL 冻结 `current_repository_id / current_module_id / current_decision_id` 只允许作为 step 级辅助恢复线索，而不能反向替代 `current_product_id` 的主上下文地位。

#### Scenario: 判断辅助恢复线索的使用边界

- **WHEN** 后续 `/spec` 或实现需要用 `Repository / Module / Decision` 辅助恢复当前步骤
- **THEN** 这些字段只允许帮助定位最近未完成的具体 step
- **AND** 不得单独成为新的主上下文锚点
- **AND** 不得让某个 step 级实体反向重写当前 `Product` 主上下文

## MODIFIED Requirements

### Requirement: `Onboarding` 的阶段语义解释

`phase10-02` 修改了 `Onboarding` 在当前阶段的解释：它不再只被解释为“首轮登记入口”，而必须被解释为“首轮建链引导主线”。

#### Scenario: 判断 `Onboarding` 的新语义

- **WHEN** 后续 `/spec`、实现或验收描述 `Onboarding`
- **THEN** 必须明确其目标是降低登记后跨多个 detail page 手工补链的摩擦
- **AND** 必须明确它承接的是首轮建链引导，而不是单纯字段收集

### Requirement: `phase06` 首轮完成语义在 `phase10` 中的解释

`phase10-02` 修改了对 `phase06` 首轮完成语义的消费方式：后续不再把“四类对象都已创建”单独视为 onboarding 的充分完成解释，而必须进一步承接“主上下文已冻结、该步应承接的最小关系已解释、未完成关系已有单值 handoff”。

#### Scenario: 判断 `Onboarding` 完成是否只基于四类实体存在

- **WHEN** 后续 `/spec`、实现或验收判断当前 `Onboarding` 是否达成 `phase10` 语义
- **THEN** 不得只依据 `Product / Repository / Module / Decision` 四类实体是否存在
- **AND** 还必须判断每一步应承接的最小正式关系是否已被解释为直承接或单值 handoff

## REMOVED Requirements

### Requirement: 将 `Onboarding` 解释为孤立的六段式登记流程

**Reason**: 这种解释会把 `Onboarding` 留在 `phase06` 的“先创建四类对象”阶段，无法支撑 `phase10` 要求的首轮建链引导、canonical handoff 与后续 reread 语义，也会把补链摩擦继续推给多个 detail page。
**Migration**: 后续 `/spec` 与实现必须改为：保留六段式主线不变，但逐步冻结每一步的 canonical owner、最小动作、默认关系、延后 handoff、默认下一步与恢复语义；不能再把 `Onboarding` 仅作为字段录入顺序页面。

### Requirement: 通过“全局最新实体”猜测当前 onboarding 主线

**Reason**: 这会在多 `Product / Repository / Module` 并存时制造错误恢复，直接破坏 `phase10` 所要求的单值主上下文锚点与恢复语义。
**Migration**: 后续恢复逻辑必须统一回到 `current_product_id` 主锚点，并仅将 `current_repository_id / current_module_id / current_decision_id` 作为 step 级辅助恢复线索。
