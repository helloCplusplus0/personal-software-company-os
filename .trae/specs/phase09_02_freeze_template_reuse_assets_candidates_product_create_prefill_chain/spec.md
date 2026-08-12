# phase09-02 冻结模板级复用资产、候选来源与 Product Create 预填动作链 Spec

## Why

`phase09-01` 已冻结当前阶段的单一主交付能力、成功标准与非目标，但 `Template Reuse` 要真正进入后续实现，还必须把模板资产的最小语义、候选来源、`Weekly Review` 中的候选消费方式，以及进入 `Product Create` 的 handoff 与成功回流链继续收敛成单值口径。否则，后续 `/spec -> 实现 -> 验收` 仍会在“候选从哪里来、预填怎么读、成功后是否丢失模板语义”这些关键点上反复仲裁。

## What Changes

- 冻结模板级复用资产的最小正式语义为 `Module` 组合快照 + `Product Create` 预填辅助
- 冻结模板候选只从既有 `Product / Module / Binding` 已持久化事实派生
- 冻结 `Weekly Review` 中模板候选的默认 active candidate、单选切换与无候选空态规则
- 冻结 `/products/new` 的模板 handoff 参数矩阵、优先级与互斥规则
- 冻结 `templateCandidateId` 作为唯一模板预填读取入口的语义
- 冻结模板预填成功会话的最小字段覆盖范围
- 冻结创建成功后模板来源语义如何承接到 `Product Detail` 与 canonical `Product <-> Module Binding` 路径

## Impact

- Affected specs:
  - `phase09_01_freeze_template_reuse_derived_intelligence_scope_success_non_goals`
  - `docs/phase/phase09_template_reuse_derived_intelligence_foundation_dev_plan.md`
  - `docs/phase/phase09_template_reuse_derived_intelligence_foundation_architecture_plan.md`
  - `docs/phase/phase09_template_reuse_derived_intelligence_foundation_shared_baseline.md`
  - 后续 `phase09-03 ~ phase09-07` 的提示、合同、owner 与交互设计规格
- Affected code:
  - 后续会影响 `Weekly Review`、`Product Create`、`Product Detail`、`Product <-> Module Binding` 相关前后端承接位
  - 当前无直接源码改动

## ADDED Requirements

### Requirement: 模板级复用资产必须保持最小正式语义

系统 SHALL 将当前阶段的模板级复用资产冻结为“`Module` 组合快照 + 面向 `Product Create` 的预填辅助”，并明确其职责是服务下一次创建，而不是成长为独立模板平台。

#### Scenario: 判断模板资产的正式语义

- **WHEN** 后续 `/spec`、实现或验收描述 `Template Reuse`
- **THEN** 必须只把模板资产解释为已存在 `Module` 组合的快照化表达
- **AND** 必须要求该资产直接服务 `Product Create` 的预填与继续编辑
- **AND** 不得把模板资产扩写为独立模板 CRUD、模板版本管理或参数化模板系统

### Requirement: 模板候选必须只从已持久化绑定事实派生

系统 SHALL 冻结模板候选的 canonical 来源为已持久化的 `product_modules` 绑定事实，并明确 `Review` 不直接生成模板候选。

#### Scenario: 判断模板候选是否来自正确事实源

- **WHEN** 后续 `/spec`、实现或验收生成模板候选
- **THEN** 候选只能从一个 Product 当前绑定的全部 Module 集合派生
- **AND** 候选去重键必须基于去重并升序排序后的 `module_id` 集合
- **AND** 候选排序必须固定为：
  1. `source_product_count DESC`
  2. `total_reuse_product_count DESC`
  3. `latest_source_product_updated_at DESC`
- **AND** `Review` 只允许提供消费作用域与返回链上下文
- **AND** 不得从未持久化草稿、页面本地状态或临时 review 输入直接生成模板候选

#### Scenario: 判断模板候选的空态与低质量规则

- **WHEN** 某个 Product 没有任何绑定 Module，或当前不存在满足条件的组合
- **THEN** 系统必须返回成功空态
- **AND** 不得把“无候选”视为页面错误
- **AND** 空 Module 集合不得形成模板候选

### Requirement: Weekly Review 模板候选消费语义必须单值

系统 SHALL 冻结 `Weekly Review` 中模板候选的消费语义为单选模型，以避免前后端各自实现一套“当前候选”判断逻辑。

#### Scenario: 判断默认 active candidate 规则

- **WHEN** `Weekly Review` 存在模板候选
- **THEN** 系统必须按既定排序结果取第一名作为默认 active candidate
- **AND** 当前 review 会话中同一时刻只允许一个 active candidate 生效

#### Scenario: 判断用户切换 active candidate 的规则

- **WHEN** 用户在 `Weekly Review` 中切换模板候选
- **THEN** 系统必须只保留一个新的 active candidate
- **AND** 依赖模板上下文的解释文案、CTA 与后续 handoff 都必须只基于当前 active candidate 计算

#### Scenario: 判断无候选时的行为

- **WHEN** `Weekly Review` 不存在任何模板候选
- **THEN** 模板相关消费位必须返回成功空态
- **AND** 不得回退到另一套未冻结的 generic focus 口径

### Requirement: Product Create 模板 handoff 参数矩阵必须单值

系统 SHALL 冻结 `/products/new` 的模板 handoff 方式为 route search 参数承接，并明确新增模板来源参数与既有来源参数的优先级、共存与互斥关系。

#### Scenario: 判断模板 handoff 的正式参数

- **WHEN** 用户从模板候选进入 `Product Create`
- **THEN** 系统必须只使用以下正式来源参数：
  - `fromTemplateReuse=true`
  - `templateCandidateId=<opaque-id>`
  - `templateSource=weekly-review | dashboard | product-detail`
- **AND** 不得在 search 参数中直接携带完整模板 payload

#### Scenario: 判断模板 handoff 与既有来源参数的关系

- **WHEN** `fromTemplateReuse=true`
- **THEN** 模板来源必须成为本次 create 的唯一业务来源语义
- **AND** `fromTemplateReuse` 必须与 `fromList / fromModuleDetail` 互斥
- **AND** `fromDashboard / dashboardSection / dashboardReturnTo` 只允许作为返回链与来源链元数据与模板来源参数共存

### Requirement: templateCandidateId 必须是唯一预填读取入口

系统 SHALL 冻结 `templateCandidateId` 为 `Product Create` 页面读取模板预填内容的唯一正式入口，并要求前端按 opaque string 消费。

#### Scenario: 判断 templateCandidateId 的消费方式

- **WHEN** `Product Create` 页面根据模板来源预填创建表单
- **THEN** 页面必须只读取 `templateCandidateId` 对应的正式预填读模型
- **AND** 前端不得自行编码或解码业务语义
- **AND** 前端不得通过页面本地拼装的 `Module` 快照替代正式读取入口

### Requirement: 模板预填成功会话必须覆盖最小字段集合

系统 SHALL 冻结“预填闭环成立”的最小字段级判定，避免后续把“只是跳转到创建页”误判为模板复用已成立。

#### Scenario: 判断预填是否真实成立

- **WHEN** 用户通过模板候选进入 `Product Create`
- **THEN** 创建页必须出现可继续编辑的预填内容
- **AND** 预填内容至少必须覆盖：
  - 模板名称或说明
  - 候选 `Module` 组合摘要
  - 待创建 Product 的初始模块上下文
- **AND** 用户必须能够在预填基础上修改并继续完成创建

### Requirement: 创建成功后模板来源语义必须继续承接

系统 SHALL 冻结模板创建成功后的承接链，确保模板组合不会在“预填时看见、成功后丢失”。

#### Scenario: 判断创建成功后的正式回流链

- **WHEN** 用户通过模板预填成功创建 Product
- **THEN** 既有 `ProductCreate` canonical mutation owner 只负责创建 Product 本身
- **AND** 当前阶段不得在同一 mutation 内自动写入 `product_modules`
- **AND** 成功跳转到 `Product Detail` 时必须继续携带以下模板来源上下文：
  - `fromTemplateReuse=true`
  - `templateCandidateId=<opaque-id>`
  - `templateSource=<weekly-review | dashboard | product-detail>`

#### Scenario: 判断 Product Detail 的模板来源复读

- **WHEN** 新建 Product 进入 `Product Detail`
- **THEN** 页面必须能够读取并展示模板来源摘要与候选 `Module` 组合摘要
- **AND** 页面必须提供进入既有 canonical `Product <-> Module Binding` 路径的正式 CTA
- **AND** 不得通过新增长期事实表或第二套写路径承接模板组合

## MODIFIED Requirements

### Requirement: phase09 成功会话中的 Template Reuse 解释

`phase09-02` 修改了 `phase09-01` 中对最小成功会话的解释：成功会话中的 `Template Reuse` 不再只是“看到候选并跳转创建”，而是必须包含模板候选来源、单选消费语义、`Product Create` 预填读取入口、以及 `Product Detail` 中的模板来源复读与 canonical 下一步动作。

#### Scenario: 判断 phase09 成功会话是否具备模板承接闭环

- **WHEN** 后续 `/spec`、实现或验收判断 `Template Reuse` 是否成立
- **THEN** 必须同时验证候选来源、active candidate、模板 handoff、预填字段覆盖与成功回流链
- **AND** 不得只用“出现候选卡片”或“出现跳转”代替模板闭环成立

## REMOVED Requirements

### Requirement: 允许模板候选来自未持久化输入、并在创建成功后自动沉淀为新 Product 的已绑定事实

**Reason**: 当前阶段已冻结模板候选只来自 `product_modules` 已持久化事实，且 `Product Create` canonical mutation owner 不应在 `phase09` 中被模板逻辑侵入或扩写为第二套绑定写路径。
**Migration**: 若后续需要让模板组合真正落成新 Product 的绑定事实，只能在 `Product Detail` 之后继续进入既有 canonical `Product <-> Module Binding` 路径完成，而不是在 `phase09-02` 里新增自动绑定主线。
