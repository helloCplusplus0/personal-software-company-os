# phase09-05 产出模板复用与派生提示的页面流、交互流与返回链设计 Spec

## Why

`phase09-02 ~ 04` 已经冻结了模板 handoff、提示矩阵、合同归属与 owner 边界，但“合同冻结”还不等于“页面可直接实现”。如果不继续把 `Weekly Review -> Product Create -> Product Detail` 的用户可见流转、空态、失败态、取消返回与移动端降级写成机械场景，后续实现仍会在页面宿主、返回链和提示呈现上重新长出第二套口径。

## What Changes

- 冻结 `Weekly Review -> Template / Hint -> Product Create -> Product Detail` 的正式页面流
- 冻结模板候选默认选中、单选切换、解释展示、CTA 与取消返回的交互流
- 冻结 `reuse_opportunity_hint / capability_gap_hint` 的最小展示方式、解释层级与动作承接
- 冻结 `Product Create` 中模板预填展示、失效降级、失败回退与继续编辑语义
- 冻结 `Product Detail` 中模板来源摘要与 canonical `Product <-> Module Binding` CTA 的可见行为
- 冻结 `fromDashboard` 返回链与模板 handoff 并存时的用户可见行为
- 冻结移动浏览器的最小降级策略，以及模板/提示/预填的空态与失败态

## Impact

- Affected specs:
  - `phase09_02_freeze_template_reuse_assets_candidates_product_create_prefill_chain`
  - `phase09_03_freeze_derived_insight_hint_set_explanation_cta_handoff`
  - `phase09_04_freeze_contract_read_model_owner_candidate_source_boundary`
  - `phase08_05_design_dashboard_review_entry_page_interaction_flow`
  - `phase04_06_frontend_state_interaction_flow`
- Affected code:
  - `frontend/src/routes/reviews/weekly.tsx`
  - `frontend/src/features/review/pages/weekly-review-page.tsx`
  - `frontend/src/features/review/data/use-weekly-review-read.ts`
  - `frontend/src/features/review/application/use-review-action.ts`
  - 后续新增 `frontend/src/features/template-reuse/*`
  - `frontend/src/routes/products/new.tsx`
  - `frontend/src/features/product-registry/pages/product-create-page.tsx`
  - `frontend/src/features/product-registry/components/product-create-form.tsx`
  - `frontend/src/features/product-registry/application/use-create-draft-product.ts`
  - `frontend/src/routes/products/$productId.tsx`
  - `frontend/src/features/product-registry/pages/product-detail-page.tsx`

## ADDED Requirements

### Requirement: `phase09` 的正式页面流必须保持单一主链与受控支链

系统 SHALL 将当前阶段的模板复用与派生提示页面流冻结为“单一模板创建主链 + 单一缺口补齐支链”。

#### Scenario: 判断 phase09 的正式主页面流

- **WHEN** 用户从 `Weekly Review` 消费模板候选或提示
- **THEN** 正式主链必须是：
  1. 在 `Weekly Review` 查看模板候选与提示
  2. 通过模板 CTA 或 `reuse_opportunity_hint` CTA 进入 `Product Create`
  3. 在 `Product Create` 查看模板预填与解释性提示并完成创建
  4. 成功进入 `Product Detail` 查看模板来源摘要并继续 canonical binding
- **AND** 不得把 `Dashboard`、`Product Detail` 或 `Product Create` 扩写成新的并列模板宿主
- **AND** 不得在主链外再发明第二套 “Template Center / Hint Center / AI Workspace”

#### Scenario: 判断 capability_gap_hint 的正式支链

- **WHEN** 用户消费 `capability_gap_hint`
- **THEN** 正式支链必须是：
  1. 在 `Weekly Review` 或 `Product Create` 查看当前模板上下文下的能力缺口提示
  2. 通过提示 CTA 进入既有 `Module Registry / Module Detail` canonical 路径
  3. 完成查看或补齐判断后，返回原提示发起页
- **AND** `capability_gap_hint` 不得进入 `Product Create` 主创建链作为并列 create 入口
- **AND** 不得把缺口补齐支链扩写成新的独立提示工作台

### Requirement: Weekly Review 必须成为模板候选与提示的唯一主消费宿主

系统 SHALL 冻结 `Weekly Review` 为模板候选与两类正式提示的唯一主消费宿主，`Product Create` 只承接模板解释性延续，不承担候选选择主线。

#### Scenario: 判断 Weekly Review 的模板区块位置与优先级

- **WHEN** 用户进入 `Weekly Review`
- **THEN** 模板候选与提示必须作为 `Weekly Review` 页面内稳定可见的正式消费位出现
- **AND** 模板候选区必须位于页面中上部的主要动作区域
- **AND** `reuse_opportunity_hint` 与 `capability_gap_hint` 必须围绕当前 active candidate 呈现
- **AND** 不得把模板候选藏入二级抽屉、弹窗或需额外展开的次级工作台

#### Scenario: 判断 Weekly Review 与 Product Create 的职责边界

- **WHEN** 用户从 `Weekly Review` 进入 `Product Create`
- **THEN** `Weekly Review` 继续负责：
  - 模板候选列表
  - 默认 active candidate
  - 单选切换
  - 模板复用机会提示
  - 能力缺口主诊断
- **AND** `Product Create` 只负责：
  - 读取 handoff 进入的模板预填
  - 展示与当前模板相关的解释性信息
  - 允许用户编辑并提交 create
- **AND** `Product Create` 不得重新提供模板候选列表或候选切换器

### Requirement: 模板候选的默认选中与单选切换交互必须可机械验收

系统 SHALL 将模板候选的默认选中、切换反馈与 CTA 更新语义冻结为浏览器可逐步核对的交互流。

#### Scenario: 默认 active candidate 自动生效

- **WHEN** `Weekly Review` 成功读取到至少一个模板候选
- **THEN** 排名第一的候选必须默认进入 active 状态
- **AND** 页面必须在该候选上展示清晰的选中态
- **AND** 模板解释文案、`reuse_opportunity_hint` 与进入 `Product Create` 的 CTA 必须默认跟随该 active candidate

#### Scenario: 用户切换 active candidate

- **WHEN** 用户在 `Weekly Review` 点击另一条模板候选
- **THEN** 新候选必须成为唯一 active candidate
- **AND** 原 active candidate 必须立即退回非选中态
- **AND** 与模板相关的解释内容、提示文案与 CTA 目标必须同步刷新到新候选
- **AND** 不得出现两个候选同时高亮或解释区与 CTA 指向不同候选的分裂状态

#### Scenario: active candidate 切换后的 CTA 语义

- **WHEN** 用户完成模板候选切换
- **THEN** “以该模板创建 Product”或等价 CTA 必须明确指向当前 active candidate
- **AND** `capability_gap_hint` 的解释与 CTA 也必须基于当前 active candidate 重算
- **AND** 不得继续引用切换前候选的解释摘要

### Requirement: 派生提示的展示层级与 CTA 承接必须单值

系统 SHALL 冻结两类正式提示的最小展示层级，保证它们既可见、可解释，又不会长成第二套工作台。

#### Scenario: 判断 reuse_opportunity_hint 的展示层级

- **WHEN** `reuse_opportunity_hint` 在 `Weekly Review` 中出现
- **THEN** 它必须与当前 active candidate 紧邻呈现，作为模板创建 CTA 的解释增强
- **AND** 提示至少必须包含：
  - 标题
  - 一段为何值得复用的解释
  - 进入 `Product Create` 的 CTA
- **AND** 不得把该提示展示成纯统计 badge 或无动作文案

#### Scenario: 判断 capability_gap_hint 的展示层级

- **WHEN** `capability_gap_hint` 在 `Weekly Review` 中出现
- **THEN** 它必须作为当前 active candidate 的补充提示出现
- **AND** 提示至少必须包含：
  - 缺口标题
  - 缺口解释
  - 进入既有 `Module Registry / Module Detail` canonical 路径的 CTA
- **AND** 不得把缺口提示与模板候选主入口并列渲染为第二主操作区

#### Scenario: 判断 Weekly Review 中 capability_gap_hint 的 CTA 页面流

- **WHEN** 用户在 `Weekly Review` 点击 `capability_gap_hint` CTA
- **THEN** 页面必须进入既有 `Module Registry` 或目标 `Module Detail`
- **AND** 当前 `Weekly Review` 的模板候选选择状态只允许作为返回链上下文保留，不得在目标页重新展开为第二套模板宿主
- **AND** 用户从目标页返回时，必须回到原 `Weekly Review`
- **AND** 返回后必须继续恢复进入前的 active candidate

#### Scenario: 判断 Product Create 中提示的展示边界

- **WHEN** `Product Create` 消费模板相关提示
- **THEN** 页面只允许展示与当前 `templateCandidateId` 对应模板有关的 `capability_gap_hint`
- **AND** 该提示必须位于预填摘要附近，作为“继续补齐”的解释性辅助
- **AND** `Product Create` 不得再次展示 `reuse_opportunity_hint`

#### Scenario: 判断 Product Create 中 capability_gap_hint 的 CTA 页面流

- **WHEN** 用户在 `Product Create` 点击 `capability_gap_hint` CTA
- **THEN** 页面必须进入既有 `Module Registry` 或目标 `Module Detail`
- **AND** 进入目标页时必须继续保留以下 create 会话上下文：
  - `fromTemplateReuse=true`
  - `templateCandidateId=<opaque-id>`
  - `templateSource=<weekly-review | dashboard | product-detail>`
- **AND** 若当前 create 页还携带 `fromDashboard / dashboardSection / dashboardReturnTo` 元数据，则这些元数据也只允许作为返回链辅助上下文一并透传
- **AND** 不得因为进入补齐页而丢失当前 create 会话

### Requirement: Product Create 的模板预填展示、编辑确认与取消返回必须单值

系统 SHALL 冻结 `Product Create` 的模板预填展示方式、用户编辑语义与取消返回链，确保模板 handoff 不会把 create 页变成第二套模板宿主。

#### Scenario: 判断 Product Create 的模板预填展示

- **WHEN** 用户通过 `fromTemplateReuse=true` 进入 `/products/new`
- **THEN** 页面必须在 create 表单上方或等价显著位置展示模板来源摘要
- **AND** 模板来源摘要至少必须包含：
  - 模板标题或说明
  - 模块组合摘要
  - 模板来源标记
- **AND** 表单字段必须以预填值进入可编辑状态
- **AND** 不得将预填内容做成只读锁定表单

#### Scenario: 判断用户在预填基础上继续编辑

- **WHEN** 用户修改预填后的 Product 名称、描述或其他 create 字段
- **THEN** 页面必须将这些修改视为当前 create 草稿的一部分
- **AND** 模板来源摘要仍必须保留可见
- **AND** 不得因为用户修改表单而丢失模板来源上下文或重置为 direct-entry

#### Scenario: 判断 Product Create 的取消返回

- **WHEN** 用户在 `Product Create` 主动取消、返回或关闭模板预填创建
- **THEN** 返回路径必须按 `templateSource` 单值决定：
  - `templateSource = weekly-review`：返回 `Weekly Review`
  - `templateSource = dashboard`：返回 `Dashboard`
  - `templateSource = product-detail`：返回原 `Product Detail`
- **AND** 若同时存在 `fromDashboard / dashboardSection / dashboardReturnTo` 元数据，则只作为恢复 Dashboard 焦点的返回链元数据，不覆盖 `templateSource` 的主返回语义
- **AND** 不得统一退回浏览器历史、`/products` 或根路由

#### Scenario: 判断从 Module 补齐页返回 Product Create

- **WHEN** 用户从 `Product Create` 通过 `capability_gap_hint` 进入 `Module Registry / Module Detail` 后选择返回
- **THEN** 返回目标必须是原 `Product Create`
- **AND** 返回时必须继续保留：
  - `fromTemplateReuse=true`
  - `templateCandidateId=<opaque-id>`
  - `templateSource=<weekly-review | dashboard | product-detail>`
- **AND** create 页必须恢复模板摘要、当前表单草稿与提示上下文
- **AND** 不得把返回结果降级为重新从 `Weekly Review`、`Dashboard` 或空白 create 重新开始

### Requirement: `fromDashboard` 返回链与模板 handoff 并存时的用户可见行为必须单值

系统 SHALL 冻结 Dashboard 元数据与模板来源参数并存时的用户可见行为，避免“页面看起来来自 A，返回却跳到 B”。

#### Scenario: Dashboard -> Weekly Review -> Product Create 的可见来源语义

- **WHEN** 用户从 `Dashboard` 进入 `Weekly Review`，再通过模板进入 `Product Create`
- **THEN** create 页的主来源语义必须显示为“来自 Weekly Review 的模板创建”
- **AND** `fromDashboard` 相关参数只允许作为返回 `Dashboard` 时的辅助元数据保留
- **AND** 不得在 create 页同时把主来源文案渲染为“来自 Dashboard”和“来自 Weekly Review”

#### Scenario: 从 Dashboard 直接触发模板 handoff 的返回链

- **WHEN** 用户未来通过 `Dashboard` 直接进入模板 handoff，且 `templateSource = dashboard`
- **THEN** create 页取消返回必须回到 `Dashboard`
- **AND** 若存在 `dashboardSection / dashboardReturnTo`，页面返回时必须尽量恢复用户原来的 Dashboard 关注位置
- **AND** 不得因为存在模板 handoff 就丢弃 Dashboard 返回链

### Requirement: 模板空态、提示空态、预填失败态与回退策略必须浏览器可验收

系统 SHALL 将模板、提示与预填的空态/失败态设计为局部可恢复状态，而不是整页级灾难错误。

#### Scenario: Weekly Review 中模板候选空态

- **WHEN** `Weekly Review` 当前不存在任何模板候选
- **THEN** 模板候选区必须展示成功空态
- **AND** 空态文案必须明确表达“当前没有可复用模板候选”
- **AND** `reuse_opportunity_hint` 与依赖模板上下文的 `capability_gap_hint` 必须一并进入成功空态
- **AND** 页面其余 `Weekly Review` 区块必须继续正常可用

#### Scenario: Weekly Review 中 capability_gap_hint 的成功空态

- **WHEN** 当前 active candidate 不存在能力缺口
- **THEN** `capability_gap_hint` 必须展示成功空态或不展示提示位
- **AND** 不得以空提示替代模板创建 CTA
- **AND** 不得把“无缺口”解释成模板区块失败

#### Scenario: Weekly Review 中模板读失败态

- **WHEN** 模板候选或模板相关提示读取发生请求失败
- **THEN** 失败必须停留在模板区块局部
- **AND** 模板区块必须提供局部重试或等价恢复手段
- **AND** 不得把整页 `Weekly Review` 回退为 page error

#### Scenario: Product Create 中模板预填 unavailable 成功态

- **WHEN** `templateCandidateId` 因读时派生漂移而不可解析
- **THEN** `Product Create` 必须展示“模板来源已失效，但仍可继续创建”的可恢复提示
- **AND** 模板摘要区必须进入 unavailable 空态
- **AND** create 表单必须退化为空白但仍可编辑
- **AND** 用户必须仍可提交普通 create

#### Scenario: Product Create 中模板预填请求失败态

- **WHEN** 模板预填请求发生网络或服务失败，而非 `UNAVAILABLE` 成功态
- **THEN** 页面必须展示局部失败提示
- **AND** create 表单仍必须可编辑
- **AND** 页面必须至少提供以下两种恢复路径：
  - 重试读取模板预填
  - 放弃模板预填并继续普通 create
- **AND** 不得因为模板预填失败阻断整个 `Product Create` 页面

#### Scenario: Product Create 中 capability_gap_hint 的成功空态与失败态

- **WHEN** 当前模板预填上下文不存在能力缺口
- **THEN** `capability_gap_hint` 必须展示成功空态或不展示提示位
- **AND** 不得影响 create 表单继续编辑
- **WHEN** `capability_gap_hint` 读取失败
- **THEN** 失败必须停留在提示区局部
- **AND** create 表单与模板摘要仍必须继续可用
- **AND** 页面必须允许用户忽略该提示继续完成 create

### Requirement: Product Detail 必须承接模板来源摘要与 canonical 下一步动作

系统 SHALL 冻结 `Product Detail` 对模板来源摘要的承接方式，确保创建成功后的模板语义不会在详情页消失。

#### Scenario: 判断 Product Detail 的模板来源摘要

- **WHEN** 用户通过模板预填成功创建 Product 并进入 `Product Detail`
- **THEN** 页面必须展示模板来源摘要区
- **AND** 该区至少必须包含：
  - 模板标题或说明
  - 模块组合摘要
  - 模板来源标记
- **AND** 模板来源摘要区不得升级为新的写操作工作台

#### Scenario: 判断 Product Detail 的 canonical 下一步动作

- **WHEN** `Product Detail` 展示模板来源摘要
- **THEN** 页面必须同时提供进入既有 `Product <-> Module Binding` canonical 路径的显式 CTA
- **AND** 该 CTA 必须作为“让模板来源继续落成真实绑定事实”的正式下一步动作
- **AND** 不得在模板摘要区内联自动绑定或第二套绑定写路径

#### Scenario: Product Detail 中模板来源失效时的展示

- **WHEN** `GetTemplateSourceSummary` 返回 `UNAVAILABLE`
- **THEN** 页面必须将模板来源摘要区降级为可恢复空态
- **AND** 空态必须说明“模板来源已不可复读”
- **AND** canonical `Product <-> Module Binding` CTA 仍必须可见

### Requirement: 移动浏览器下必须采用单页降级而不是第二套页面体系

系统 SHALL 在移动浏览器下继续沿用单一 `React Web` 页面主线，并通过布局降级而不是额外页面重写来承接模板与提示。

#### Scenario: Weekly Review 的移动端降级

- **WHEN** `Weekly Review` 在窄屏环境展示模板候选与提示
- **THEN** 页面必须改为单列纵向布局
- **AND** 模板候选必须按卡片纵向堆叠展示
- **AND** active candidate 的解释区与 CTA 必须跟随当前候选在同一阅读流内出现
- **AND** 不得引入左右双栏、横向对比表或移动端专用模板页面

#### Scenario: Product Create 的移动端降级

- **WHEN** `Product Create` 在窄屏环境展示模板来源摘要、提示与表单
- **THEN** 模板来源摘要必须位于表单上方
- **AND** 模板提示与 CTA 必须采用全宽按钮或等价紧凑布局
- **AND** 不得因移动端降级而隐藏模板来源、返回动作或 canonical binding 线索

## MODIFIED Requirements

### Requirement: `phase09` 成功会话中的 Template Reuse 页面解释

`phase09-02` 已冻结模板候选、模板 handoff 与成功回流链。

自 `phase09-05` 起，系统必须把该 requirement 修改为：

- 模板复用成功会话不再只定义“参数如何传”
- 还必须定义：
  - `Weekly Review` 中模板候选的正式消费位
  - active candidate 的切换反馈
  - `Product Create` 的模板摘要与取消返回
  - `Product Detail` 的模板来源摘要与 canonical 下一步动作

#### Scenario: 判断 Template Reuse 页面闭环是否成立

- **WHEN** 后续实现或浏览器验收判断 `Template Reuse` 是否形成正式闭环
- **THEN** 必须同时核对：
  - 模板候选可见
  - 默认 active candidate 成立
  - 切换后 CTA 与解释同步更新
  - create 页预填摘要可见且可编辑
  - 创建成功后 detail 页继续显示模板来源摘要
- **AND** 不得只凭“存在 handoff 参数”判定闭环成立

### Requirement: `phase09` 成功会话中的 Derived Intelligence 页面解释

`phase09-03` 已冻结 `reuse_opportunity_hint / capability_gap_hint` 的 trigger、解释与 CTA。

自 `phase09-05` 起，系统必须把该 requirement 修改为：

- 两类提示不仅要有合同与 CTA
- 还必须在 `Weekly Review` 与 `Product Create` 中拥有稳定的展示层级、页面流、返回链、空态和失败态

#### Scenario: 判断提示页面闭环是否成立

- **WHEN** 后续实现或浏览器验收判断提示链路是否成立
- **THEN** 必须同时核对：
  - 提示出现在正确页面
  - 提示解释紧邻正确的模板或预填上下文
  - CTA 指向既有 canonical 路径
  - 从提示进入补齐页后可返回正确原页面
  - 无候选、无缺口、读取失败时都进入预期局部状态
- **AND** 不得只凭接口返回了提示对象就判定提示闭环成立

## REMOVED Requirements

### Requirement: 允许在多个页面并列承接模板候选选择、提示主诊断与返回链判定

**Reason**: 这会让 `Weekly Review`、`Product Create`、`Product Detail` 各自长出半套模板工作流，直接破坏 `phase09-02 ~ 04` 已冻结的 handoff、owner 与 canonical return chain。

**Migration**: 模板候选选择与主诊断统一收敛到 `Weekly Review`；`Product Create` 只承接预填与解释延续；`Product Detail` 只承接模板来源摘要与 canonical binding CTA。
