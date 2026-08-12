# phase09-10 落实派生提示展示、动作 handoff 与解释性回流 Spec

## Why

`phase09-08` 已经把 `TemplateReuseService`、派生提示读取合同与前端 read owner 正式落地，`phase09-09` 又把“模板候选 -> Product Create 预填 -> Product Detail 来源摘要”这条模板主链推进成真实页面行为。但当前两类正式提示还没有被稳定接入页面、动作 handoff 仍未与既有 canonical owner 对齐、提示消费后的 reread 与解释性回流也未形成浏览器可验收的单值路径；如果此时直接实现，很容易把提示退化成纯文案卡片，或者在 `Weekly Review / Product Create` 长出第二套长期智能主线。

## What Changes

- 在 `Weekly Review` 正式落地 `reuse_opportunity_hint` 与 `capability_gap_hint` 的最小展示位、解释文案与 CTA
- 在 `Product Create` 正式落地与当前 `templateCandidateId` 绑定的 `capability_gap_hint` 解释区与补齐 handoff
- 将提示动作 handoff 收敛到单一 application owner，统一承接 `trigger / explanation / CTA / target owner` 四元组
- 实现提示消费后的返回链、必要 reread、空态与错误语义，避免提示消费后丢失 active candidate 与 create 会话
- 明确只实现有稳定 canonical CTA 的两类正式提示；没有稳定 CTA 的候选提示直接裁撤，不进入实现态
- **BREAKING**：`phase09-10` 覆盖范围内不再允许保留 generic focus fallback、纯统计提示卡片或页面级临时提示编排作为长期稳态

## Impact

- Affected specs:
  - `phase09_03_freeze_derived_insight_hint_set_explanation_cta_handoff`
  - `phase09_05_design_template_reuse_hint_page_flow_interaction_return_chain`
  - `phase09_06_design_frontend_read_write_owner_state_flow`
  - `phase09_08_land_template_reuse_contract_backend_frontend_read_enablement`
  - `phase09_09_land_template_candidate_selection_product_create_prefill_result_handoff`
- Affected code:
  - `frontend/src/features/review/data/use-weekly-review-read.ts`
  - `frontend/src/features/review/pages/weekly-review-page.tsx`
  - `frontend/src/features/review/application/use-review-action.ts`
  - `frontend/src/features/template-reuse/data/use-derived-insight-hints-read.ts`
  - `frontend/src/features/template-reuse/application/use-derived-hint-handoff.ts`
  - `frontend/src/features/product-registry/application/use-product-create-template-handoff.ts`
  - `frontend/src/features/product-registry/pages/product-create-page.tsx`
  - `frontend/src/routes/modules/index.tsx`
  - `frontend/src/routes/modules/$moduleId.tsx`
  - `frontend/src/features/module-registry/pages/module-list-page.tsx`
  - `frontend/src/features/module-registry/pages/module-detail-page.tsx`
  - `frontend/src/features/template-reuse/components/*`

## ADDED Requirements

### Requirement: phase09-10 必须作为提示实现任务直接消费既有 enablement 与 create 主链

系统 SHALL 将 `phase09-10` 冻结为“在 `phase09-08 / 09` 既有能力之上落实正式提示展示、动作 handoff 与解释性回流”的源码实现任务，而不是重新定义一轮提示合同、owner 边界或第二套智能主线。

#### Scenario: 判断 phase09-10 的任务身份

- **WHEN** 后续执行 `phase09-10`
- **THEN** 主要工作必须落在 `Weekly Review / Product Create / Module Registry / Module Detail` 的真实提示行为与回流链
- **AND** 不得再把 `GetDerivedInsightHints`、`template-reuse/data/*` 或 `TemplateReuseService` 重新定义成新的主工作内容
- **AND** 不得新增第二套提示读取 transport、第二套提示宿主页面或第二套 create / module mutation

### Requirement: 只有具备稳定 canonical CTA 的正式提示才允许进入 phase09-10 实现

系统 SHALL 将 `phase09-10` 的正式提示范围收敛为 `reuse_opportunity_hint` 与 `capability_gap_hint` 两类，并要求每类提示都具备完整的 `trigger / explanation / CTA / target owner` 四元组。

#### Scenario: 判断允许进入实现的提示集合

- **WHEN** 审查 `phase09-10` 的提示实现范围
- **THEN** 只允许正式实现：
  - `reuse_opportunity_hint`
  - `capability_gap_hint`
- **AND** 没有稳定 canonical CTA 的候选提示必须直接裁撤
- **AND** 不得保留“先展示解释文案，CTA 后面再补”的半成品提示

#### Scenario: 判断提示是否满足四元组

- **WHEN** 某条提示进入页面正式展示
- **THEN** 它必须同时具备：
  - 明确的触发事实
  - 明确的解释文案
  - 明确的 CTA
  - 明确的 target owner
- **AND** 若只剩统计数字、badge 或孤立文案，则该提示不得进入正式实现态

### Requirement: Weekly Review 必须成为两类正式提示的主消费宿主

系统 SHALL 将 `Weekly Review` 冻结为 `reuse_opportunity_hint` 与 `capability_gap_hint` 的唯一主消费宿主，并要求提示语义与当前 active template candidate 单值绑定。

#### Scenario: reuse_opportunity_hint 在 Weekly Review 的正式展示

- **WHEN** `Weekly Review` 存在 active template candidate 且满足模板复用机会触发条件
- **THEN** 页面必须在模板候选区附近展示 `reuse_opportunity_hint`
- **AND** 提示至少必须包含：
  - 模板机会标题
  - 为什么当前组合值得复用的解释文案
  - 进入 `/products/new` 的正式 CTA
- **AND** CTA 必须只通过既有 `fromTemplateReuse / templateCandidateId / templateSource=weekly-review` 主链进入 `Product Create`
- **AND** 不得把该提示退化为纯统计卡片或与模板创建 CTA 分裂成两套入口

#### Scenario: capability_gap_hint 在 Weekly Review 的正式展示

- **WHEN** 当前 active template candidate 存在能力缺口
- **THEN** 页面必须在模板候选区附近展示 `capability_gap_hint`
- **AND** 提示至少必须包含：
  - 缺口标题
  - 缺口解释
  - 进入既有 `Module Registry / Module Detail` 的正式 CTA
- **AND** CTA 必须绑定当前 active template candidate 与当前 review 作用域
- **AND** 不得把缺口提示渲染成新的并列主工作台

#### Scenario: active candidate 切换后的提示刷新

- **WHEN** 用户在 `Weekly Review` 切换 active template candidate
- **THEN** `reuse_opportunity_hint` 与 `capability_gap_hint` 的解释、CTA 与目标参数都必须同步刷新到新候选
- **AND** 不得继续引用切换前候选的提示文本或目标
- **AND** 不得出现两个候选共用同一条旧提示的分裂状态

#### Scenario: 无 active template candidate 时的成功空态

- **WHEN** 当前不存在 active template candidate
- **THEN** `reuse_opportunity_hint` 与依赖模板上下文的 `capability_gap_hint` 都必须退回成功空态
- **AND** 不得回退到未冻结的 generic focus fallback
- **AND** 不得把该场景解释成 `Weekly Review` page error

### Requirement: Product Create 只允许承接 capability_gap_hint 的解释性延续与补齐 handoff

系统 SHALL 将 `Product Create` 冻结为模板创建主线中的解释性提示延续页；它只允许承接与当前 `templateCandidateId` 绑定的 `capability_gap_hint`，不得重新成长为第二套提示诊断中心。

#### Scenario: Product Create 的 capability_gap_hint 展示边界

- **WHEN** 用户通过模板 handoff 进入 `Product Create`
- **THEN** 页面只允许展示与当前 `templateCandidateId` 对应模板相关的 `capability_gap_hint`
- **AND** 提示必须位于模板来源摘要或预填说明附近
- **AND** 页面不得重新展示 `reuse_opportunity_hint`
- **AND** 页面不得重新提供模板候选切换器或第二套提示筛选逻辑

#### Scenario: Product Create 中 capability_gap_hint 的 CTA

- **WHEN** 用户在 `Product Create` 点击 `capability_gap_hint` CTA
- **THEN** 页面必须进入既有 `Module Registry / Module Detail` canonical 路径
- **AND** 必须继续保留：
  - `fromTemplateReuse=true`
  - `templateCandidateId=<opaque-id>`
  - `templateSource=<weekly-review | dashboard | product-detail>`
  - 当前 create 会话所需的返回链上下文
- **AND** 不得因为进入模块补齐页而丢失当前 create 会话

#### Scenario: 从模块补齐页返回 Product Create

- **WHEN** 用户从 `Product Create` 的缺口提示进入 `Module Registry / Module Detail` 后选择返回
- **THEN** 返回目标必须是原 `Product Create`
- **AND** 返回时必须恢复：
  - 模板来源摘要
  - 当前 create 表单草稿
  - 当前提示解释与 CTA
- **AND** 不得把返回结果降级为重新从 `Weekly Review` 或空白 create 重新开始

### Requirement: 提示动作 handoff 必须收敛到单一 application owner

系统 SHALL 为 `phase09-10` 新增或完善单一提示 handoff application owner，统一承接提示点击后的 search 参数拼装、返回链恢复与 target owner 跳转，避免导航逻辑散落在多个页面和卡片组件里。

#### Scenario: 提示 handoff owner 的最小职责

- **WHEN** 实现 `phase09-10`
- **THEN** 必须存在单一提示 handoff application owner
- **AND** 它至少必须承接：
  - `reuse_opportunity_hint` 的 Product Create handoff 参数拼装
  - `capability_gap_hint` 的 Module Registry / Module Detail handoff 参数拼装
  - `Weekly Review` active candidate 返回恢复
  - `Product Create` 会话返回恢复
  - 非法提示参数回退与错误归一化
- **AND** 页面组件不得各自重复拼装 `templateCandidateId / templateSource / capabilityKey / reviewScopeKey`

#### Scenario: target owner 只能是既有 canonical path

- **WHEN** 提示 handoff owner 计算下一跳
- **THEN** 目标只允许是既有 canonical path：
  - `reuse_opportunity_hint` -> `Product Create`
  - `capability_gap_hint` -> `Module Registry / Module Detail`
- **AND** 不得把提示直接跳到第二套 `Decision Workspace / AI Hint Center / Generic Focus Page`
- **AND** `Product Detail` 只允许作为 `reuse_opportunity_hint` 经 `Product Create` 成功创建后的结果回流页，不得作为独立提示目标页

### Requirement: 提示消费后的解释性回流与 reread 必须可机械验收

系统 SHALL 将提示消费后的回流语义冻结为“回到原页面时仍能恢复正确解释上下文与必要 reread”的浏览器可验收路径，而不是一次性跳转后语义丢失。

#### Scenario: Weekly Review 的 capability_gap_hint 消费回流

- **WHEN** 用户从 `Weekly Review` 进入 `Module Registry / Module Detail` 后返回
- **THEN** 页面必须恢复到原 `Weekly Review`
- **AND** 原 active template candidate 必须继续保持为 active
- **AND** 对应的 `capability_gap_hint` 必须重新读取并展示最新解释
- **AND** 不得退回默认候选或丢失当前提示上下文

#### Scenario: reuse_opportunity_hint 的结果回流

- **WHEN** 用户从 `reuse_opportunity_hint` 进入 `Product Create` 并成功创建 Product
- **THEN** 用户必须继续回流到既有 `Product Detail`
- **AND** `Product Detail` 必须继续展示 `phase09-09` 已落地的模板来源摘要与 canonical binding CTA
- **AND** 该页面承担的是提示消费后的解释性结果回流，而不是新的提示宿主

#### Scenario: Prompt reread 的最小刷新语义

- **WHEN** 提示消费路径完成一次返回或成功创建
- **THEN** 必须只触发与原页面相关的正式 reread
- **AND** 不得通过页面级临时 `window.location.reload()`、全局硬刷新或第二套 local store 模拟“已更新”
- **AND** 提示的成功空态、unavailable 与 error 语义必须继续由既有 read owner 派生

### Requirement: 提示空态、错误态与裁撤语义必须保持单值

系统 SHALL 将两类正式提示的空态、错误态与裁撤语义冻结为可恢复的局部状态，避免提示实现重新制造第二套错误主线。

#### Scenario: reuse_opportunity_hint 成功空态

- **WHEN** 当前没有满足触发条件的模板复用机会
- **THEN** `reuse_opportunity_hint` 必须展示成功空态或不展示提示位
- **AND** 不得以空态替代模板候选主 CTA
- **AND** 不得将该场景解释为接口错误

#### Scenario: capability_gap_hint 成功空态

- **WHEN** 当前 active template candidate 不存在能力缺口
- **THEN** `capability_gap_hint` 必须展示成功空态或不展示提示位
- **AND** 不得回退到 generic focus fallback
- **AND** 不得因为缺口为空而阻断模板创建主链

#### Scenario: 提示读取失败的局部错误

- **WHEN** `GetDerivedInsightHints` 发生真实读取失败
- **THEN** 错误必须停留在提示局部区域
- **AND** 页面主链仍然必须继续可用
- **AND** 提示区必须提供重试或等价恢复手段
- **AND** 不得因为提示失败把整页 `Weekly Review` 或 `Product Create` 打成 page error

## MODIFIED Requirements

### Requirement: phase09 的派生智能成立标准

自 `phase09-10` 起，系统 SHALL 不再把“页面上出现一段提示解释文案”视为 `Derived Intelligence Deepening` 已成立，而必须进一步满足“提示可重复消费、可进入正式动作链、可完成解释性回流与 reread”的实现标准。

#### Scenario: 判断派生智能是否真正成立

- **WHEN** 审查 `phase09-10` 是否完成
- **THEN** 必须同时验证：
  - `reuse_opportunity_hint` 已在 `Weekly Review` 正式可见且可进入 `Product Create`
  - `capability_gap_hint` 已在 `Weekly Review` 与 `Product Create` 按边界正式可见
  - 提示 handoff 已对齐既有 canonical owner
  - 无 active template candidate 与无缺口场景都表现为成功空态
  - 提示消费后的返回链与 reread 可由浏览器路径直接验证
- **AND** 不得再把“只有解释文案或统计摘要出现”视为通过

### Requirement: phase09-09 的模板主链在 phase09-10 中必须继续保持单值

自 `phase09-10` 起，系统 SHALL 将 `phase09-09` 已落地的模板候选选择、Product Create 预填与 Product Detail 来源摘要视为唯一模板主链，并要求提示实现只作为该主链的解释与动作增强层，不得长出并列长期主线。

#### Scenario: 审查提示是否破坏模板主链单值性

- **WHEN** 审查 `phase09-10` 的 spec、实现或验收
- **THEN** 必须看到提示只是在既有模板主链和模块补齐支链上增强解释与 handoff
- **AND** 不得把 `Weekly Review`、`Product Create` 或 `Product Detail` 改造成第二套智能中心

## REMOVED Requirements

### Requirement: 允许 generic focus fallback、无稳定 CTA 的提示或页面级散装提示编排继续进入正式实现
**Reason**: 这些模式会直接削弱 `phase09-03 / 05` 已冻结的单值提示矩阵，并把 `phase09-10` 重新拉回“先上几张提示卡、后面再决定怎么消费”的漂移状态。
**Migration**: 正式提示只保留 `reuse_opportunity_hint` 与 `capability_gap_hint`；提示导航、回流与错误语义统一收敛到单一 application owner 与既有 canonical path；不具备稳定 CTA 的候选提示保持裁撤，不进入当前实现范围。
