# phase10-05 产出 Dashboard / Daily Review / Detail pages 的下一步动作承接设计 Spec

## Why

`phase10-03` 已冻结 `Decision` 生命周期与 `Decision Detail` CTA 的 canonical 解释，`phase10-04` 也已冻结 `Onboarding` 建链流与返回合同。但如果不继续把 `Dashboard / Daily Review / Product Detail / Module Detail / Repository Detail / Decision Detail` 的“下一步动作”承接矩阵写成页面级单值规则，后续实现仍会在各页面临场判断显示哪个 CTA、先跳去哪里、成功后 reread 看什么。
因此，`phase10-05` 必须把关键页面的 CTA inventory、触发条件、默认跳转目标、回流 reread 规则与主 CTA / 次 CTA 优先级冻结下来，作为后续实现与浏览器验收的直接上游。

## What Changes

- 冻结 `Dashboard Home / Daily Review / Product Detail / Module Detail / Repository Detail / Decision Detail` 的页面级 CTA 矩阵
- 冻结每个页面“何时显示、显示什么、跳去哪里、返回后如何 reread”的单值规则
- 冻结 `Current Focus` 与各 detail CTA 之间的正式承接关系
- 冻结 `Product Detail / Module Detail / Repository Detail` 的 CTA inventory 与类型边界
- 冻结多个 CTA 同时成立时的主 CTA / 次 CTA 决策规则与展示优先级
- 冻结成功回流与 reread 规则，避免页面各自拼装第二套业务主线

## Impact

- Affected specs:
  - `phase10_01_freeze_asset_action_closure_scope_success_non_goals`
  - `phase10_03_freeze_decision_lifecycle_detail_cta_pending_semantics_matrix`
  - `phase10_04_design_onboarding_chain_return_empty_failure_flow`
  - `docs/phase/phase10_asset_action_closure_foundation_dev_plan.md`
  - `docs/phase/phase10_asset_action_closure_foundation_shared_baseline.md`
  - `docs/phase/phase10_asset_action_closure_foundation_architecture_plan.md`
  - 既有 `phase05_10_dashboard_feedback_formal_spec`
  - 既有 `phase08_02_freeze_dashboard_review_entry_page_route_handoff`
- Affected code:
  - 当前无直接源码改动
  - 后续会影响 `frontend/src/features/dashboard/*`
  - 后续会影响 `frontend/src/features/review/*`
  - 后续会影响 `frontend/src/features/product-registry/*`
  - 后续会影响 `frontend/src/features/module-registry/*`
  - 后续会影响 `frontend/src/features/repository-binding/*`
  - 后续会影响 `frontend/src/features/decision-center/*`
  - 后续会影响对应的 dashboard / review / detail 读模型与 application owner

## ADDED Requirements

### Requirement: 关键页面必须具备单值“下一步动作”语义

系统 SHALL 为 `Dashboard Home / Daily Review / Product Detail / Module Detail / Repository Detail / Decision Detail` 冻结单值“下一步动作”语义，使页面在任意时刻都能回答“当前最应该做什么”，而不是只堆叠一组无优先级按钮。

#### Scenario: 判断关键页面是否具备正式下一步语义

- **WHEN** 后续 `/spec`、实现或验收描述上述任一页面的 CTA
- **THEN** 必须至少明确：
  - 何时显示
  - 显示什么 CTA
  - 默认跳转目标
  - 成功后 reread 看什么
- **AND** 不得把这些判断分散到多个局部组件临场拼装

### Requirement: `Dashboard Home` 必须冻结首个主 CTA 与次 CTA 矩阵

系统 SHALL 为 `Dashboard Home` 冻结首个主 CTA 与次 CTA 的承接矩阵，使其继续作为经营入口，但不长出第二套业务编排主线。

#### Scenario: `Dashboard` 空态或近空态

- **WHEN** `Dashboard` 读取到空态或近空态，且当前系统尚未形成稳定经营闭环
- **THEN** 主 CTA 必须进入 `Onboarding`
- **AND** 不得把“查看列表页”或“返回 Review”作为空态下的主动作

#### Scenario: `Dashboard` 存在 pending decision

- **WHEN** `Dashboard` 读取到 canonical `Decision.status = proposed` 的 pending decision
- **AND** 当前不存在更高优先级的 canonical 结构缺口
- **THEN** 主 CTA 必须进入该决策的 `Decision Detail`
- **AND** CTA 文案必须明确这是正式状态推进入口
- **AND** 不得把局部 dismiss、隐藏卡片或列表筛选当作主动作

#### Scenario: `Dashboard` 同时存在结构缺口与 pending decision

- **WHEN** `Dashboard` 同时读取到明确结构缺口与 `Decision.status = proposed` 的 pending decision
- **THEN** 主 CTA 必须优先进入该结构缺口对应的 canonical owner
- **AND** pending decision 只能退为次 CTA
- **AND** 不得让执行者再自行判断两者谁应成为主动作

#### Scenario: `Dashboard` 存在明确结构缺口
- **WHEN** `Dashboard` 读取到明确结构缺口
- **THEN** 主 CTA 必须进入该结构缺口对应的单值目标页或 owner
- **AND** 次 CTA 才允许并列展示 reread 或次级导航

#### Scenario: `Dashboard` 主 CTA 成功回流

- **WHEN** 用户从 `Dashboard` 主 CTA 跳转并完成对应动作后返回
- **THEN** `Dashboard` reread 必须首先回答：
  - 原主动作是否已完成
  - 若已完成，新的主动作是否已切换
  - 若未完成，是否仍回到同一 canonical owner
- **AND** 不得仅靠 toast 或前端局部隐藏维持“已处理”的假象

### Requirement: `Daily Review` 必须冻结与 `Current Focus` 一致的动作承接矩阵

系统 SHALL 为 `Daily Review` 冻结与 `Current Focus` 一致的动作承接矩阵，使其只承担“今天要推进什么”的最小分诊，而不是重新发明一套局部 action hub。

#### Scenario: `Current Focus` 命中 pending decision

- **WHEN** `Daily Review` 的 `Current Focus` 命中 `Decision.status = proposed`
- **AND** 当前不存在更高优先级的 canonical 结构缺口
- **THEN** 主 CTA 必须进入该 `Decision Detail`
- **AND** 成功回流后 reread 必须优先判断该决策是否仍为 pending

#### Scenario: `Current Focus` 命中实体结构缺口

- **WHEN** `Current Focus` 命中 `Product / Module / Repository` 的结构缺口
- **THEN** 主 CTA 必须进入该结构缺口对应的单值目标页或 owner
- **AND** 不得先跳去无关列表页让用户自行寻找下一步

#### Scenario: `Daily Review` 与 `Dashboard` 的分工

- **WHEN** 同一条业务事实同时出现在 `Dashboard` 与 `Daily Review`
- **THEN** 两者必须共享同一 canonical owner 解释
- **AND** `Daily Review` 只负责更明确地指出“今天先推进哪一项”
- **AND** 不得与 `Dashboard` 形成两套相互竞争的主动作

### Requirement: `Current Focus` 必须正式 handoff 到 canonical detail CTA

系统 SHALL 冻结 `Current Focus` 与 detail CTA 的关系：`Current Focus` 只负责指出最优先事项，正式动作必须继续落到对应 canonical detail 或 canonical owner。

#### Scenario: `Current Focus` 命中 `Decision`

- **WHEN** `Current Focus` 命中某条 `Decision`
- **THEN** 正式 handoff 必须进入该 `Decision Detail`
- **AND** 后续状态推进或 reread 解释必须完全复用 `phase10-03` 的 canonical 规则

#### Scenario: `Current Focus` 命中结构缺口

- **WHEN** `Current Focus` 命中某个实体结构缺口
- **THEN** 正式 handoff 必须进入该结构缺口对应的单值目标页或 owner
- **AND** 不得在 `Current Focus` 卡片内联第二套长期写路径

### Requirement: 结构缺口必须映射到单值默认跳转目标

系统 SHALL 为 `Dashboard / Daily Review / Current Focus` 命中的结构缺口冻结单值默认跳转目标，避免“detail 或 binding owner”这种并列写法把目标选择重新留给实现层。

#### Scenario: 命中 `Product` 范围的结构缺口

- **WHEN** 结构缺口属于当前 `Product` 范围内仍可由 `Product Detail` 汇总并分发的下一步动作
- **THEN** 默认跳转目标必须是该 `Product Detail`
- **AND** 不得直接跳去无关列表页

#### Scenario: 命中 `Module` 范围的结构缺口

- **WHEN** 结构缺口属于当前 `Module` 范围内仍可由 `Module Detail` 汇总并分发的下一步动作
- **THEN** 默认跳转目标必须是该 `Module Detail`
- **AND** 不得再并列保留第二个默认目标

#### Scenario: 命中 `Repository` 范围的结构缺口

- **WHEN** 结构缺口属于当前 `Repository` 范围内仍可由 `Repository Detail` 汇总并分发的下一步动作
- **THEN** 默认跳转目标必须是该 `Repository Detail`
- **AND** 不得再并列保留第二个默认目标

#### Scenario: 命中纯关系闭合型结构缺口

- **WHEN** 结构缺口的唯一正式承接位已经明确为 canonical binding owner，且该关系不能在对应 detail 页内安全完成
- **THEN** 默认跳转目标必须直接进入该 canonical binding owner
- **AND** 不得同时把对应 detail 页继续保留为并列默认目标

### Requirement: `Product Detail` 必须冻结 CTA inventory 与主次优先级

系统 SHALL 为 `Product Detail` 冻结页面级 CTA inventory，至少承接“补仓库、补模块、承接关联决策、返回 reread”四类动作，并冻结其中的主 CTA / 次 CTA 优先级。

#### Scenario: `Product Detail` 的 CTA 类型边界

- **WHEN** 后续 `/spec`、实现或验收定义 `Product Detail` 的 CTA
- **THEN** 当前阶段只允许出现以下下一步动作类型：
  - 补齐 `Repository` 相关结构缺口
  - 补齐 `Module` 相关结构缺口
  - 进入相关 `Decision` 的正式承接页
  - 返回 `Dashboard / Daily Review` reread
- **AND** 不得长出与 `Product` 并列的局部旁路写路径

#### Scenario: `Product Detail` 存在结构缺口时的主 CTA

- **WHEN** `Product Detail` 同时存在多个潜在动作，但其中包含未闭合的 canonical 结构缺口
- **THEN** 结构缺口内部优先级必须冻结为：
  1. `Repository` 相关结构缺口
  2. `Module` 相关结构缺口
- **AND** 主 CTA 必须优先指向当前排序最高的结构缺口对应 owner
- **AND** `Decision` 相关 CTA 只能退为次 CTA

#### Scenario: `Product Detail` 结构 CTA 成功后的 reread

- **WHEN** 用户从 `Product Detail` 的结构 CTA 完成对应动作后返回
- **THEN** 页面必须先检查当前最高优先级结构缺口是否已经消失
- **AND** 若 `Repository` 缺口已闭合但 `Module` 缺口仍在，则主 CTA 必须切换为 `Module` 相关动作
- **AND** 若结构缺口均已闭合但仍存在 pending decision，则主 CTA 必须切换为相关 `Decision Detail`
- **AND** 只有在结构缺口与 pending decision 都已不存在时，reread 返回 CTA 才能升为兜底主动作

#### Scenario: `Product Detail` 无结构缺口但存在 pending decision

- **WHEN** `Product Detail` 已无更高优先级的结构缺口，但存在与当前产品相关的 pending decision
- **THEN** 主 CTA 必须进入对应 `Decision Detail`
- **AND** reread 后必须优先检查该决策是否已退出 pending

#### Scenario: `Product Detail` 只有 reread 型动作

- **WHEN** `Product Detail` 当前没有待补结构缺口，也没有需要优先推进的 pending decision
- **THEN** 页面允许仅保留返回 `Dashboard / Daily Review` reread 的 CTA
- **AND** 该 CTA 只能作为兜底主动作存在，不得掩盖仍未解释的结构缺口

### Requirement: `Module Detail` 必须冻结 CTA inventory 与主次优先级

系统 SHALL 为 `Module Detail` 冻结页面级 CTA inventory，至少承接“产品绑定、仓库映射、承接关联决策、返回 reread”四类动作，并冻结其中的主 CTA / 次 CTA 优先级。

#### Scenario: `Module Detail` 的 CTA 类型边界

- **WHEN** 后续 `/spec`、实现或验收定义 `Module Detail` 的 CTA
- **THEN** 当前阶段只允许出现以下下一步动作类型：
  - 进入 `Module -> Product` 的 canonical binding path
  - 进入 `Module -> Repository` 的 canonical mapping path
  - 进入相关 `Decision` 的正式承接页
  - 返回 `Dashboard / Daily Review` reread
- **AND** 不得把 `Module Detail` 扩写成局部工作台

#### Scenario: `Module Detail` 多个 CTA 同时成立

- **WHEN** `Module Detail` 同时存在产品绑定、仓库映射与决策承接三类候选 CTA
- **THEN** 结构缺口内部优先级必须冻结为：
  1. `Module -> Product` 绑定
  2. `Module -> Repository` 映射
- **AND** 主 CTA 必须优先指向当前排序最高的结构缺口
- **AND** `Decision` 状态推进只能排在结构闭合之后
- **AND** reread 返回动作只能排在最后

#### Scenario: `Module Detail` 结构 CTA 成功后的 reread

- **WHEN** 用户从 `Module Detail` 的结构 CTA 完成对应动作后返回
- **THEN** 页面必须先检查 `Module -> Product` 绑定是否仍缺失
- **AND** 若产品绑定已闭合但仓库映射仍缺失，则主 CTA 必须切换为 `Module -> Repository` 映射
- **AND** 若结构缺口均已闭合但仍存在相关 pending decision，则主 CTA 必须切换为相关 `Decision Detail`
- **AND** 只有在结构缺口与 pending decision 都已不存在时，reread 返回 CTA 才能升为兜底主动作

### Requirement: `Repository Detail` 必须冻结 CTA inventory 与主次优先级

系统 SHALL 为 `Repository Detail` 冻结页面级 CTA inventory，至少承接“产品绑定、模块映射、承接关联决策、返回 reread”四类动作，并冻结其中的主 CTA / 次 CTA 优先级。

#### Scenario: `Repository Detail` 的 CTA 类型边界

- **WHEN** 后续 `/spec`、实现或验收定义 `Repository Detail` 的 CTA
- **THEN** 当前阶段只允许出现以下下一步动作类型：
  - 进入 `Repository -> Product` 的 canonical binding path
  - 进入 `Repository -> Module` 的 canonical mapping path
  - 进入相关 `Decision` 的正式承接页
  - 返回 `Dashboard / Daily Review` reread
- **AND** 不得新增与 `Repository` 并列的局部写路径

#### Scenario: `Repository Detail` 多个 CTA 同时成立

- **WHEN** `Repository Detail` 同时存在产品绑定、模块映射与决策承接三类候选 CTA
- **THEN** 结构缺口内部优先级必须冻结为：
  1. `Repository -> Product` 绑定
  2. `Repository -> Module` 映射
- **AND** 主 CTA 必须优先指向当前排序最高的结构缺口
- **AND** 次 CTA 才允许保留其他结构动作或 `Decision` 导航
- **AND** 不得让多个结构 CTA 争夺同一主动作位置

#### Scenario: `Repository Detail` 结构 CTA 成功后的 reread

- **WHEN** 用户从 `Repository Detail` 的结构 CTA 完成对应动作后返回
- **THEN** 页面必须先检查 `Repository -> Product` 绑定是否仍缺失
- **AND** 若产品绑定已闭合但模块映射仍缺失，则主 CTA 必须切换为 `Repository -> Module` 映射
- **AND** 若结构缺口均已闭合但仍存在相关 pending decision，则主 CTA 必须切换为相关 `Decision Detail`
- **AND** 只有在结构缺口与 pending decision 都已不存在时，reread 返回 CTA 才能升为兜底主动作

### Requirement: `Decision Detail` 必须冻结为状态推进优先的页面级 CTA 宿主

系统 SHALL 将 `Decision Detail` 继续冻结为状态推进优先的页面级 CTA 宿主，使其页面级主 CTA / 次 CTA 继续服从 `phase10-03` 的生命周期矩阵，而不与其他 detail 页形成冲突语义。

#### Scenario: `Decision Detail` 为 `proposed`

- **WHEN** 当前 `Decision.status = proposed`
- **THEN** 主 CTA 必须是正式状态推进 CTA
- **AND** 其他 canonical 导航或返回 reread 只能作为次 CTA

#### Scenario: `Decision Detail` 为 `active`

- **WHEN** 当前 `Decision.status = active`
- **THEN** 主 CTA 必须继续围绕允许的正式状态推进或结果消费
- **AND** 不得回退为“返回 Dashboard”优先于正式业务动作

#### Scenario: `Decision Detail` 为终态

- **WHEN** 当前 `Decision.status = superseded` 或 `archived`
- **THEN** 页面不得再展示状态推进主 CTA
- **AND** 页面主动作只能退化为结果消费或返回 reread

### Requirement: 多个 CTA 同时成立时必须有单值主 CTA 决策规则

系统 SHALL 冻结一个跨页面统一规则：当多个 CTA 同时成立时，页面必须存在单值主 CTA，且执行者不需要再自行判断谁是主动作。

#### Scenario: 跨页面统一优先级

- **WHEN** `Dashboard / Daily Review / Product Detail / Module Detail / Repository Detail / Decision Detail` 中同时出现多个合法 CTA
- **THEN** 主 CTA 的统一优先级必须为：
  1. 先闭当前 canonical 结构缺口
  2. 再推进 `Decision` 的正式状态动作
  3. 最后才是返回 `Dashboard / Daily Review` reread
- **AND** 次 CTA 可以并列存在
- **AND** 次 CTA 不得与主 CTA 的目标 owner 相互冲突

#### Scenario: 页面级局部变体

- **WHEN** 某个页面需要呈现多个次 CTA
- **THEN** 只允许在不违反统一优先级的前提下调整次 CTA 排序
- **AND** 不得通过页面局部习惯重写主 CTA 优先级

### Requirement: CTA 成功回流后的 reread 规则必须页面级冻结

系统 SHALL 为关键页面冻结 CTA 成功回流后的 reread 规则，确保用户回到来源页后看到的是新的 canonical 结果，而不是旧 CTA 仍继续误报。

#### Scenario: 从 detail CTA 成功返回来源页

- **WHEN** 用户从任一 detail CTA 完成对应动作并返回来源页
- **THEN** 来源页 reread 必须先检查原主 CTA 对应的缺口或 pending 是否已消失
- **AND** 若已消失，页面必须切换到新的主 CTA
- **AND** 若未消失，页面必须继续保持同一主 CTA

#### Scenario: 返回 `Dashboard / Daily Review` 的 reread

- **WHEN** 用户从 `Product / Module / Repository / Decision Detail` 返回 `Dashboard / Daily Review`
- **THEN** reread 必须重新评估：
  - `Current Focus` 是否变化
  - 原主动作是否已完成
  - 是否出现新的更高优先级结构缺口或 pending decision
- **AND** 不得依赖来源页的本地临时状态保持一致性

## MODIFIED Requirements

### Requirement: `phase10` Shared Baseline 中的页面级 CTA inventory 解释

`phase10-05` 修改了 shared baseline 中“关键页面至少存在一个下一步动作 CTA”的解释：后续不再只要求页面上有可点击按钮，而必须将 CTA 落实为页面级 inventory、触发条件、默认目标、成功回流与主次优先级矩阵。

#### Scenario: 判断页面 CTA 是否只是形式满足

- **WHEN** 后续实现或验收只证明“页面上出现了一个按钮”
- **THEN** 不得视为满足 `phase10-05`
- **AND** 还必须证明该 CTA 的显示条件、跳转目标、reread 落点与优先级已被机械冻结

### Requirement: `Current Focus` 在 `phase10` 中的消费方式

`phase10-05` 修改了 `Current Focus` 的消费方式：它不再只作为提示信号，而必须正式把用户 handoff 到对应 canonical detail CTA。

#### Scenario: 判断 `Current Focus` 是否仍停留在提示层

- **WHEN** 后续实现或验收仍把 `Current Focus` 仅解释为提示文案或被动信息块
- **THEN** 必须判定为未满足 `phase10-05`
- **AND** `Current Focus` 必须继续回答“点完去哪、回来看什么”

## REMOVED Requirements

### Requirement: 由各页面自行判断哪个 CTA 应作为主动作

**Reason**: 这会让 `Dashboard / Daily Review / Detail pages` 各自长出不同的动作优先级，直接破坏 `phase10` 所要求的单值下一步动作语义。
**Migration**: 后续实现必须统一复用 `phase10-05` 冻结的主 CTA 优先级：先闭 canonical 结构缺口，再推进 `Decision` 正式状态动作，最后才是 reread 返回动作。

### Requirement: 用“返回列表页 / 返回 Dashboard”长期兜底替代正式下一步动作

**Reason**: 这会把当前阶段重新退回“用户自己猜下一步”的状态，无法满足 `Asset-Action Closure` 的交付目标。
**Migration**: 后续实现必须优先回答 canonical owner 的正式动作；只有在当前页面不存在未解释缺口和待推进 `Decision` 时，才允许把 reread 返回 CTA 升为兜底主动作。
