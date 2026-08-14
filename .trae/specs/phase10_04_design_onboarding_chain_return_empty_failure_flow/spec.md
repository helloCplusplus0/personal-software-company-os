# phase10-04 产出 Onboarding 建链流、返回链与空态/失败态设计 Spec

## Why

`phase10-02` 已冻结 `Onboarding` 从“首轮登记”升级为“首轮建链引导”的语义与逐步建链矩阵，但如果不继续把页面流、成功后默认下一步、detail handoff 返回链、空态/失败态和中途中断恢复写成可机械执行的单值规则，后续实现仍会在页面层各自猜测“成功后去哪里”“失败后留在哪一步”“从 detail 返回时该回到哪个 step”。
因此，`phase10-04` 必须把 `Onboarding` 的页面版执行清单、返回合同与异常语义冻结下来，作为后续前端实现、后端恢复读模型与浏览器验收的直接上游。

## What Changes

- 冻结冷启动用户从 `Dashboard -> Onboarding` 的首轮建链页面流
- 冻结 `welcome / product / repository / module / decision / complete` 六段式页面版执行清单
- 冻结每一步创建成功后的默认下一步、默认返回链与 reread 落点
- 冻结 canonical detail handoff 的来源参数、返回优先级与失效回退规则
- 冻结空态、失败态、已存在实体场景与中途中断再进入的单值解释
- 冻结移动浏览器下的最小降级策略与浏览器验收前提

## Impact

- Affected specs:
  - `phase10_01_freeze_asset_action_closure_scope_success_non_goals`
  - `phase10_02_freeze_onboarding_first_run_chain_guidance_action_matrix`
  - `docs/phase/phase10_asset_action_closure_foundation_dev_plan.md`
  - `docs/phase/phase10_asset_action_closure_foundation_shared_baseline.md`
  - `docs/phase/phase10_asset_action_closure_foundation_architecture_plan.md`
  - 既有 `phase06_06_design_onboarding_frontend_route_interaction_flow`
- Affected code:
  - 当前无直接源码改动
  - 后续会影响 `frontend/src/features/onboarding/*`
  - 后续会影响 `frontend/src/routes/onboarding.tsx`
  - 后续会影响 `frontend/src/features/onboarding/lib/onboarding-return.ts`
  - 后续会影响四类 canonical detail 页面与其 `validateSearch`
  - 后续会影响 `backend/internal/onboarding/*` 的恢复读模型与最小错误语义

## ADDED Requirements

### Requirement: `Onboarding` 必须冻结为单一路由宿主的页面版建链流

系统 SHALL 将 `Onboarding` 冻结为单一路由宿主的页面版建链流：`/onboarding` 承接 `welcome / product / repository / module / decision / complete` 六段式主线，页面职责是解释当前 step、展示当前最小动作、给出单值下一步，而不是在多个页面间散落一套隐式编排。

#### Scenario: 冷启动用户从 `Dashboard` 正式进入 `Onboarding`

- **WHEN** 用户处于 `Asset-Action Closure` 的冷启动或近空态，并从 `Dashboard` 进入首轮建链
- **THEN** 系统必须进入单一路由宿主 `/onboarding`
- **AND** 页面必须先落到 `welcome` 或当前最近未完成 step
- **AND** 不得要求用户先去 `Product / Repository / Module / Decision` 列表页自行判断下一步

#### Scenario: `Onboarding` 页面版执行清单的单值解释

- **WHEN** 后续 `/spec`、实现或验收描述某个 `Onboarding` step 的页面行为
- **THEN** 必须同时回答以下问题：
  - 当前 step 的 canonical owner 是谁
  - 当前 step 的最小动作是什么
  - 成功后默认下一步去哪里
  - 若关系无法同页闭合，去哪个 canonical handoff
  - reread 后页面应停留在哪个 step
- **AND** 不得把这些判断拆散到多个页面组件临场决定

### Requirement: 六段式 step 必须存在页面级成功态与默认下一步规则

系统 SHALL 为 `welcome / product / repository / module / decision / complete` 六段式 step 冻结页面级成功态与默认下一步规则，使每一步在成功后都只有一个正式前进方向。

#### Scenario: `welcome` 的成功动作

- **WHEN** 用户在 `welcome` 点击正式开始首轮建链
- **THEN** 系统必须单值进入 `product`
- **AND** 不得在 `welcome` 预创建实体或并列提供第二条起始主线

#### Scenario: `product` 的成功动作

- **WHEN** 用户在 `product` 完成首个 `Product` 创建，且写入成功
- **THEN** 系统必须冻结该 `Product` 为当前 `current_product_id`
- **AND** 页面必须默认进入 `repository`
- **AND** reread 后不得仍停留在 `welcome`

#### Scenario: `repository` 的成功动作

- **WHEN** 用户在 `repository` 完成首个 `Repository` 创建，且当前承接位可同页闭合 `Repository -> Product`
- **THEN** 系统必须在本步完成最小正式绑定
- **AND** 页面必须默认进入 `module`

#### Scenario: `repository` 成功但需要 canonical handoff

- **WHEN** 用户在 `repository` 创建成功，但当前承接位无法同页完成 `Repository -> Product` 最小正式绑定
- **THEN** 系统必须展示单值 canonical handoff
- **AND** 用户完成 handoff 返回后，页面必须继续落回 `module`
- **AND** 不得在 handoff 完成后把用户重新送回 `repository` 让其再次猜测

#### Scenario: `module` 的成功动作

- **WHEN** 用户在 `module` 完成首个 `Module` 创建，且当前承接位可同页闭合 `Module -> Product`
- **THEN** 系统必须在本步完成最小正式关系
- **AND** 页面必须默认进入 `decision`

#### Scenario: `module` 成功但 `Repository` 映射仍未闭合

- **WHEN** 用户在 `module` 完成创建，但 `Repository` 映射尚未闭合
- **THEN** 系统只允许将该缺口解释为从属于 `module` 步骤的 detail CTA 或 canonical handoff
- **AND** 页面主线仍必须单值进入 `decision`
- **AND** 不得因此在 `module` 步骤并列展开第二条以 `Repository` 为中心的主线

#### Scenario: `decision` 的成功动作

- **WHEN** 用户在 `decision` 完成首个 `Decision` 创建，且当前承接位可形成与当前 `Product` 的最小正式承接
- **THEN** 页面必须默认进入 `complete`
- **AND** 该 `Decision` 必须能在后续 reread 中被解释

#### Scenario: `decision` 成功但需要 canonical handoff

- **WHEN** 用户在 `decision` 创建成功，但当前承接位无法同页完成最小正式承接
- **THEN** 系统必须展示单值 canonical handoff
- **AND** 用户完成 handoff 返回后，页面必须进入 `complete`
- **AND** 不得因为 handoff 存在而阻断进入 `complete`

#### Scenario: `complete` 的成功动作

- **WHEN** 用户进入 `complete`
- **THEN** 页面只允许承接完成确认、结果回看与返回 `Dashboard`
- **AND** 默认主动作必须返回 `Dashboard`
- **AND** 不得在 `complete` 再长出新的并列写路径

### Requirement: 六段式 step 必须存在页面级空态、失败态与 reread 落点规则

系统 SHALL 为 `welcome / product / repository / module / decision / complete` 六段式 step 继续冻结页面级空态、失败态与 reread 落点规则，确保每一步在“尚未开始”“中途中断”“刷新 reread”“写入失败”时都只有一个正式页面落点，而不是让页面组件临场解释下一步。

#### Scenario: `welcome` 的空态与 reread 落点

- **WHEN** 当前不存在 `current_product_id`，且首轮建链尚未开始或来源参数全部缺失
- **THEN** `/onboarding` 必须落到 `welcome`
- **AND** `welcome` 的正式空态只允许解释为“尚未建立首轮建链主上下文”
- **AND** reread 后不得跳过 `product`

#### Scenario: `product` 的空态、失败态与 reread 落点

- **WHEN** 当前 `current_product_id` 尚未存在，且系统已进入首轮建链主线
- **THEN** `/onboarding` 必须落到 `product`
- **AND** `product` 的正式空态必须解释为“需要创建首个 `Product` 以冻结主上下文”
- **AND** 若本步写入失败，页面必须继续停留在 `product`
- **AND** 若刷新或 reread 时仍未创建成功，页面必须继续停留在 `product`

#### Scenario: `repository` 的空态、失败态与 reread 落点

- **WHEN** `current_product_id` 已存在，但当前主上下文下的 `Repository` 尚未创建或其最小正式绑定仍未解释完成
- **THEN** `/onboarding` 必须落到 `repository`
- **AND** `repository` 的正式空态必须解释为“需要创建首个 `Repository` 并优先完成 `Repository -> Product` 承接”
- **AND** 若本步创建失败或同页绑定失败，页面必须继续停留在 `repository`
- **AND** 若刷新或 reread 时该步仍未满足完成条件，页面必须继续停留在 `repository`

#### Scenario: `module` 的空态、失败态与 reread 落点

- **WHEN** `current_product_id` 已存在，且 `repository` 已完成，但 `Module` 尚未创建或其最小正式关系仍未解释完成
- **THEN** `/onboarding` 必须落到 `module`
- **AND** `module` 的正式空态必须解释为“需要创建首个 `Module` 并优先完成 `Module -> Product` 承接”
- **AND** 若本步创建失败或同页承接失败，页面必须继续停留在 `module`
- **AND** 若刷新或 reread 时该步仍未满足完成条件，页面必须继续停留在 `module`

#### Scenario: `decision` 的空态、失败态与 reread 落点

- **WHEN** `current_product_id` 已存在，且 `module` 已完成，但 `Decision` 尚未创建或其最小正式承接仍未解释完成
- **THEN** `/onboarding` 必须落到 `decision`
- **AND** `decision` 的正式空态必须解释为“需要创建首个 `Decision` 并使其在 reread 中可被解释”
- **AND** 若本步创建失败或同页承接失败，页面必须继续停留在 `decision`
- **AND** 若刷新或 reread 时该步仍未满足完成条件，页面必须继续停留在 `decision`

#### Scenario: `complete` 的 reread 落点

- **WHEN** 用户已经进入 `complete`，且当前主线上不存在未解释的结构缺口
- **THEN** 刷新或 reread 后页面必须继续停留在 `complete`
- **AND** 不得因为历史实体并存或返回参数缺失而回退到中间 step
- **AND** 用户离开 `Onboarding` 后的默认去向才允许是 `Dashboard`

### Requirement: canonical detail handoff 必须有统一返回合同

系统 SHALL 为从 `Onboarding` 进入的 canonical detail handoff 冻结统一返回合同，使 `Product / Repository / Module / Decision` detail 页面都能在完成后回到同一条 `Onboarding` 主线，而不是各自回到列表页或 `Dashboard`。

#### Scenario: 从 `Onboarding` 进入 canonical detail handoff

- **WHEN** 用户从 `Onboarding` 某个 step 进入 canonical detail handoff
- **THEN** 跳转必须携带统一且参数级冻结的来源合同：
  - `fromOnboarding=true`
  - `onboardingProductId=<current_product_id>`
  - `onboardingStep=<repository|module|decision>`
- **AND** detail 页不得重新发明第二套来源标记
- **AND** 若后续还需要携带其他页面局部参数，它们只允许作为附属参数存在，不得替代上述三项正式返回合同

#### Scenario: detail handoff 成功后的返回优先级

- **WHEN** 用户在 canonical detail 页完成来自 `Onboarding` 的 handoff 后返回
- **THEN** 系统必须优先采用显式 step 返回线索
- **AND** 前提是该线索仍隶属于当前 `current_product_id` 主上下文
- **AND** 若显式线索已失效，才允许回退为恢复读模型推导出的最近未完成 step
- **AND** 不得跳回全局默认列表页或 `Dashboard`

#### Scenario: detail handoff 返回线索失效

- **WHEN** 用户从 canonical detail 返回 `Onboarding`，但携带的 step 返回线索已失效、越过前置步骤或不再隶属于当前主上下文
- **THEN** 系统必须回退为基于 `current_product_id` 的最近未完成 step
- **AND** 不得继续强行落到失效 step
- **AND** 不得按“最近访问页面”猜测返回落点

#### Scenario: 用户离开 canonical handoff 但未完成承接

- **WHEN** 用户从 `Onboarding` 进入 canonical detail handoff 后，没有完成该 handoff 要求的正式承接就直接返回、关闭或离开页面
- **THEN** 系统必须将当前 step 继续解释为未完成
- **AND** 返回 `Onboarding` 后必须落回原所属的未完成 step
- **AND** 不得错误推进到下一步
- **AND** 不得把“进入过 handoff 页面”误判为“该步已完成”

### Requirement: `Onboarding` 必须冻结空态与已存在实体场景的单值解释

系统 SHALL 为 `Onboarding` 冻结空态与已存在实体场景的单值解释，确保在“完全冷启动”“部分已存在”“多实体并存”时，页面都能给出可预测的正式入口与下一步。

#### Scenario: 完全冷启动空态

- **WHEN** 当前不存在 `current_product_id`，且首轮建链尚未开始
- **THEN** `Dashboard` 必须提供进入 `Onboarding` 的主入口
- **AND** `/onboarding` 默认进入 `welcome`
- **AND** `welcome` 后必须单值进入 `product`

#### Scenario: 已存在 `Product` 但未进入后续 step

- **WHEN** 当前 `current_product_id` 已存在，但 `Repository` 尚未闭合到当前主上下文
- **THEN** 系统必须将当前主线解释为 `repository`
- **AND** 不得把“已有 Product”解释为无需继续 onboarding

#### Scenario: 已存在多个实体并存

- **WHEN** 系统中已经存在多个 `Product / Repository / Module / Decision`
- **THEN** `Onboarding` 恢复与页面落点必须优先围绕 `current_product_id`
- **AND** 不得按全局最新创建时间选择当前主线
- **AND** 不得因为存在历史实体而把用户直接送去 `complete`

#### Scenario: 已存在实体但关系仍未解释

- **WHEN** 某一步对应实体已存在，但该步应承接的最小正式关系尚未解释为直承接或 canonical handoff
- **THEN** 当前 step 仍不得视为完成
- **AND** 页面必须把下一步解释为补齐该关系的 detail CTA 或 canonical handoff
- **AND** 不得把“实体已存在”误判为“本步已完成”

### Requirement: `Onboarding` 必须冻结失败态与恢复规则

系统 SHALL 为 `Onboarding` 冻结读失败、写失败、handoff 失败与返回失败四类最小失败态，确保失败后页面保持单值可恢复，不会偷偷前进或清空主上下文。

#### Scenario: step 内写失败

- **WHEN** 用户在任一步骤触发创建或绑定写动作失败
- **THEN** 页面必须停留在当前 step
- **AND** 当前已冻结的主上下文不得被清空
- **AND** 系统必须展示可重试的失败反馈
- **AND** 不得在写失败后自动前进到下一步

#### Scenario: `Onboarding` 读模型加载失败

- **WHEN** `/onboarding` 初始化或 reread 时读取恢复读模型失败
- **THEN** 页面必须展示统一失败态与重试动作
- **AND** 不得回退为凭前端局部缓存猜当前 step
- **AND** 不得静默重定向到 `Dashboard`

#### Scenario: canonical handoff 自身失败

- **WHEN** 用户从某个 step 进入 canonical handoff，但 handoff 页面上的绑定或承接动作失败
- **THEN** 失败处理应留在 handoff 页完成重试
- **AND** 成功返回前不得把当前 step 标记为已完成
- **AND** 返回 `Onboarding` 后仍必须落回原本所属主线的下一正式 step 或当前未完成 step

#### Scenario: 返回参数丢失或不可解释

- **WHEN** 用户从 canonical detail 返回 `Onboarding` 时，来源参数缺失或不可解释
- **THEN** 系统必须回退为基于 `current_product_id` 的最近未完成 step
- **AND** 若连主上下文也不存在，则回退为 `welcome`
- **AND** 不得停在无定义的中间态

### Requirement: 中途中断再进入必须复用单值恢复主线

系统 SHALL 将中途中断、从 `Dashboard` 再进入、从浏览器刷新后重载，统一解释为同一套恢复主线：优先尊重当前主上下文，其次定位最近未完成 step，不得另起第二套“继续草稿”解释。

#### Scenario: 从 `Dashboard` 继续首轮建链

- **WHEN** 用户已开始但未完成 `Onboarding`，并从 `Dashboard` 再次进入
- **THEN** 系统必须直接落到当前 `current_product_id` 所锚定的最近未完成 step
- **AND** 不得重新从 `welcome` 开始
- **AND** 不得按全局最新实体猜测当前主线

#### Scenario: 浏览器刷新后重新加载

- **WHEN** 用户在 `Onboarding` 任一步骤刷新页面
- **THEN** 系统必须依据恢复读模型重新解释当前 step
- **AND** 若刷新前存在合法的显式 step 线索，则仍应先验证其是否隶属于当前主上下文
- **AND** 不得仅凭本地页面状态决定恢复位置

### Requirement: 移动浏览器必须提供最小降级但不改变主线语义

系统 SHALL 为移动浏览器冻结最小降级策略：允许在布局、文案密度与 CTA 呈现方式上做压缩，但不得改变六段式主线、默认下一步、返回合同与失败恢复规则。

#### Scenario: 移动浏览器下的 step 呈现

- **WHEN** 用户在移动浏览器访问 `/onboarding`
- **THEN** 页面可将 stepper 降级为更紧凑的进度提示或当前步骤摘要
- **AND** 当前 step 的主 CTA 必须保持单值可见
- **AND** 不得因为屏幕受限而隐藏正式下一步或返回动作

#### Scenario: 移动浏览器下的 detail handoff 返回

- **WHEN** 用户在移动浏览器完成 canonical detail handoff 后返回
- **THEN** 返回语义必须与桌面浏览器完全一致
- **AND** 只允许在布局层做紧凑化处理
- **AND** 不得因为端上差异形成第二套 step 恢复规则

## MODIFIED Requirements

### Requirement: `phase06` Onboarding 页面与路由流在 `phase10` 中的解释

`phase10-04` 修改了既有 `phase06-06` 对 `Onboarding` 页面与路由流的解释：`/onboarding` 不再是围绕 draft-first 录入与草稿摘要的页面壳，而必须升级为围绕首轮建链、canonical handoff、返回链与 reread 恢复的单一路由宿主。

#### Scenario: 判断 `Onboarding` 页面流是否仍停留在 draft-first 语义

- **WHEN** 后续实现或验收仍把 `Onboarding` 解释为“创建草稿后展示摘要，再自行去 detail 编辑”
- **THEN** 必须判定为未满足 `phase10-04`
- **AND** 页面流必须继续回答“此步最小建链动作是什么、若未闭合去哪里补、补完后回来落在哪一步”

### Requirement: `Onboarding` 完成态在 `phase10` 中的解释

`phase10-04` 修改了 `complete` 的消费方式：进入 `complete` 不再只表示四类实体存在，而必须表示六段式主线的最小关系缺口都已被解释为直承接或 canonical handoff，且返回 `Dashboard` 的 reread 路径已经明确。

#### Scenario: 判断是否可以进入 `complete`

- **WHEN** 用户完成 `decision` 并准备进入 `complete`
- **THEN** 系统不得只以“四类实体存在”作为充分条件
- **AND** 还必须保证未即时完成的关系已经被单值解释为正式 handoff 或 detail CTA

## REMOVED Requirements

### Requirement: 将 `Onboarding` 返回链退化为 detail 页各自处理的局部逻辑

**Reason**: 这会让 `Product / Repository / Module / Decision` detail 页再次各自生长一套来源参数、返回优先级与失败回退规则，直接破坏 `phase10` 所要求的单一路由宿主与单值恢复主线。
**Migration**: 后续实现必须统一复用 `Onboarding` 返回合同，由 `fromOnboarding`、主上下文锚点和显式 step 返回线索共同解释返回落点；detail 页只承接合同，不再各自发明第二套规则。

### Requirement: 将“实体已创建”直接等同于“当前 step 已完成”

**Reason**: `phase10` 的目标是首轮建链引导而不是对象存在性登记，若只凭实体存在就判定 step 完成，会把最小正式关系缺口重新留给用户在多个 detail 页间猜测。
**Migration**: 后续实现必须改为同时检查实体存在与该步最小正式关系是否已被解释为直承接或 canonical handoff，缺一不可。
