# Phase06-02 Draft-First / Partial-Entry 写路径冻结 Spec

## Why

`phase06-01` 已经把 first-run onboarding 的页面边界、入口判定和首轮成功会话条件冻结成单值结论，接下来必须把 `Product / Repository / Module / Decision` 四类核心对象的 draft-first / partial-entry 写路径收住。否则 `Onboarding` 会与既有 create 页面并存两套写语义，前端 mutation 继续散落在 page 组件内，后续 `.proto`、后端 owner 和验收 fixture 都会继续漂移。

## What Changes

- 冻结四类核心对象在 `phase06` 首轮成功会话中的最小必填字段
- 冻结四类核心对象的 `draft created` 与 `first-run completed` 语义边界
- 冻结“草稿优先、后补完整信息”的写入模型，不要求先完成完整表单再落第一条持久化记录
- 冻结每类对象唯一 `application` 写入承接位，禁止页面、表单、面板各自长出第二套 mutation 语义
- 冻结 `query` 层纯只读边界与 mutation 固定承接位原则
- 冻结 phase06 必须回收的既有 create 页面 / 组件级 mutation 范围
- **BREAKING** 现有四个 create 页面中的 page-level `useMutation` 与页面自带成功回流 / 失效刷新语义不再能继续作为 phase06 新增写路径的 canonical 模式

## Impact

- Affected specs: `phase06_onboarding_sovereignty_reuse_foundation`
- Affected code:
  - `frontend/src/features/product-registry/pages/product-create-page.tsx`
  - `frontend/src/features/repository-binding/pages/repository-create-page.tsx`
  - `frontend/src/features/module-registry/pages/module-create-page.tsx`
  - `frontend/src/features/decision-center/pages/decision-create-page.tsx`
  - `frontend/src/features/product-registry/components/product-create-form.tsx`
  - `frontend/src/features/repository-binding/components/repository-create-form.tsx`
  - `frontend/src/features/module-registry/components/module-create-form.tsx`
  - `frontend/src/features/decision-center/components/decision-create-form.tsx`
  - 后续各 feature slice 的 `application` 写入 owner
  - 后续 `OnboardingWrite` owner 与前端 `Onboarding` 页面流程

## ADDED Requirements

### Requirement: Draft-First 写模型冻结

系统 SHALL 将 `phase06` 首轮录入的四类对象创建语义冻结为“先创建最小可持久化草稿，再允许后补完整信息”，而不是要求用户先完成完整表单再写入第一条记录。

#### Scenario: draft-first 创建成立

- **WHEN** 用户在 `Onboarding` 或既有 canonical create 页面发起 `CreateDraftProduct / CreateDraftRepository / CreateDraftModule / CreateDraftDecision`
- **THEN** 当前阶段必须允许基于最小必填字段直接创建已持久化记录
- **AND** 该记录一旦创建成功，即视为 `draft created`
- **AND** 后续字段补全可以在 `Onboarding` 后续步骤或既有 canonical owner 页面继续完成
- **AND** 当前阶段不得要求用户先完成完整详情字段后才允许落第一条持久化记录

### Requirement: Product Draft 最小必填字段冻结

系统 SHALL 将 `CreateDraftProduct` 的最小用户必填字段冻结为 `name`。

#### Scenario: Product draft 创建

- **WHEN** 用户创建首轮 `Product` 草稿
- **THEN** 用户最少只需要提供 `name`
- **AND** `description` 在当前阶段允许后补
- **AND** 当首轮未填写 `description` 时，草稿创建阶段必须由系统固定填充空字符串 `''` 作为持久化占位值
- **AND** `status` 在当前阶段必须由系统固定填充为 `active`，不要求用户在首轮录入时手动选择
- **AND** 只要记录已持久化，就算当前轮 `Product` 草稿创建成功

### Requirement: Repository Draft 最小必填字段冻结

系统 SHALL 将 `CreateDraftRepository` 的最小用户必填字段冻结为 `name + url`。

#### Scenario: Repository draft 创建

- **WHEN** 用户创建首轮 `Repository` 草稿
- **THEN** 用户最少只需要提供 `name` 与 `url`
- **AND** `provider` 在当前阶段不得再作为首轮人工必填字段
- **AND** `provider` 在草稿创建时必须由系统固定填充为 `manual` 作为单值占位，用于承接当前阶段“手动登记 / 手动绑定”的仓库来源语义
- **AND** 用户后续可以在 canonical owner 页面把 `provider` 从 `manual` 改为真实提供商值
- **AND** `status` 在当前阶段必须由系统固定填充为 `active`，不要求用户在首轮录入时手动选择
- **AND** 只要记录已持久化，就算当前轮 `Repository` 草稿创建成功

### Requirement: Module Draft 最小必填字段冻结

系统 SHALL 将 `CreateDraftModule` 的最小用户必填字段冻结为 `name`。

#### Scenario: Module draft 创建

- **WHEN** 用户创建首轮 `Module` 草稿
- **THEN** 用户最少只需要提供 `name`
- **AND** `description` 在当前阶段允许后补
- **AND** 当首轮未填写 `description` 时，草稿创建阶段必须由系统固定填充空字符串 `''` 作为持久化占位值
- **AND** `status` 在当前阶段必须由系统固定填充为 `active`，不要求用户在首轮录入时手动选择
- **AND** `capability_key` 不属于当前子任务的首轮强制输入项，不得阻断本轮 `Module` 草稿创建成功
- **AND** 只要记录已持久化，就算当前轮 `Module` 草稿创建成功

### Requirement: Decision Draft 最小必填字段冻结

系统 SHALL 将 `CreateDraftDecision` 的最小用户必填字段冻结为 `title + choice + reason`。

#### Scenario: Decision draft 创建

- **WHEN** 用户创建首轮 `Decision` 草稿
- **THEN** 用户最少只需要提供 `title`、`choice` 与 `reason`
- **AND** `context`、`problem`、`alternatives`、`impact` 在当前阶段允许后补
- **AND** 当首轮未填写 `context`、`problem`、`impact` 时，草稿创建阶段必须分别由系统固定填充空字符串 `''` 作为持久化占位值
- **AND** 当首轮未填写 `alternatives` 时，草稿创建阶段必须由系统固定填充空数组 `[]` 作为持久化占位值
- **AND** `status` 在当前阶段必须由系统固定填充为 `proposed`，不要求用户在首轮录入时手动选择
- **AND** 若存在来源对象上下文（如后续承接 `source_module_id`），该上下文允许由系统自动带入，但不得作为额外人工必填项
- **AND** 只要记录已持久化，就算当前轮 `Decision` 草稿创建成功

### Requirement: Draft Created 与首轮成功会话边界冻结

系统 SHALL 明确四类对象的 `draft created` 与 `first-run completed` 不是同一个判定层级。

#### Scenario: 单对象草稿创建成功

- **WHEN** 任一对象满足其最小必填字段并成功写入持久化记录
- **THEN** 系统只能判定该对象 `draft created`
- **AND** 不得把单对象草稿创建成功直接等价为整个首轮成功会话完成

#### Scenario: 首轮成功会话成立

- **WHEN** 同一个 first-run onboarding run 中，四类对象都已分别满足各自 `draft created`
- **THEN** 当前阶段才允许把首轮成功会话判定为成立
- **AND** 此时仍不要求 `Product / Repository / Module / Decision` 之间的绑定关系全部完成

### Requirement: Partial-Entry 后补规则冻结

系统 SHALL 将 partial-entry 的后补规则冻结为“可以补完整字段，但不要求补齐绑定关系才保留草稿记录有效性”。

#### Scenario: 后补完整字段

- **WHEN** 用户后续回到任一已创建的草稿对象
- **THEN** 系统必须允许继续补全非首轮必填字段
- **AND** 草稿对象在后补完整字段前后都应保持同一条持久化记录身份
- **AND** 当前阶段用于 draft 持久化占位的空字符串 `''`、空数组 `[]` 与 `provider = manual` 只承担首轮占位语义，不得被解释为业务上“真实已补全”

#### Scenario: 绑定关系后补

- **WHEN** 用户已经完成四类对象的最小草稿创建
- **THEN** 当前阶段允许把以下关系后补，而不阻断首轮成功会话：
  - `Repository -> Product` 绑定
  - `Module -> Product` 绑定
  - `Module -> Repository` 映射
  - `Decision -> target` 链接

### Requirement: 唯一 Application 写入承接位冻结

系统 SHALL 为四类对象各自冻结一个唯一的 feature-slice `application` 写入承接位，供 `Onboarding` 与既有 canonical create 页面共享。

#### Scenario: 写入承接位判定

- **WHEN** 前端发起任一 `CreateDraft*` 写动作
- **THEN** 同一业务动作必须只有一个正式 `application` 承接位
- **AND** `Onboarding` 页面不得单独拼装第二套 mutation、成功回流与失效刷新语义
- **AND** 既有 canonical create 页面也不得继续各自保留独立的 phase06 canonical 写入语义

### Requirement: Query 层只读边界冻结

系统 SHALL 冻结 `query` 层纯只读边界，禁止在读适配层继续承接 phase06 新增写路径。

#### Scenario: Query 与 mutation 分层

- **WHEN** 接手者判断四类对象在 phase06 的前端分层
- **THEN** `query` 或读适配层只允许承接读取、缓存键与只读解包
- **AND** `create / update / bind / link` 一类写动作必须收敛到固定 `application` 承接位
- **AND** 当前阶段不得在 read adapter、query hook 或展示组件中继续散落 `useMutation`

### Requirement: 既有 Create 页面回收范围冻结

系统 SHALL 明确 phase06 必须回收的 page-level mutation 范围，避免 `Onboarding` 与既有 create 页面并存两套写语义。

#### Scenario: 必须回收的页面范围

- **WHEN** phase06 实现前端新增写路径
- **THEN** 以下页面中的 page-level mutation 必须在 phase06 范围内回收到各自 feature-slice 固定 `application` 承接位：
  - `ProductCreatePage`
  - `RepositoryCreatePage`
  - `ModuleCreatePage`
  - `DecisionCreatePage`
- **AND** 表单组件只承接字段收集、提交事件与表单内错误展示
- **AND** 页面组件不得继续长期保留第二套成功回流 / 失效刷新语义

### Requirement: 当前阶段不要求新增绑定前置字段

系统 SHALL 明确首轮最小草稿创建不要求用户额外填写绑定前置字段或关系型补充字段。

#### Scenario: 非前置字段判定

- **WHEN** 用户在首轮 `Onboarding` 中完成四类对象最小草稿创建
- **THEN** 当前阶段不得额外要求用户先填写以下信息才允许草稿创建成功：
  - `Product` 的绑定目标
  - `Repository` 的 Product / Module 关系
  - `Module` 的 Repository / Product 关系
  - `Decision` 的 target link

## MODIFIED Requirements

### Requirement: 既有 Create 页面前端写入模式

现有 `Product / Repository / Module / Decision` create 页面在 `phase06-02` 中 SHALL 继续承担 canonical 进入页与表单宿主职责，但不再长期拥有各自独立的 page-level mutation、成功回流与失效刷新语义。

#### Scenario: 既有 create 页面职责收敛

- **WHEN** 接手者改造既有 create 页面以承接 `phase06`
- **THEN** 页面可以继续承接来源上下文、返回路径与表单宿主布局
- **AND** 但同类 `CreateDraft*` 写动作的实际承接必须回收到 feature-slice `application` owner
- **AND** 页面与 `Onboarding` 必须共享同一套写入语义

## REMOVED Requirements

### Requirement: 由页面组件内联 `useMutation` 直接定义 canonical 写入语义

**Reason**: 这种模式会让 `Onboarding` 与既有 create 页面各自维护一套 mutation、成功回流与 query 失效策略，直接违反 phase06 的“唯一 application 入口 / mutation 固定承接位”约束。

**Migration**: phase06 后续实现统一改为：四类对象的 `CreateDraft*` 写动作收敛到各 feature slice 的固定 `application` owner；页面与表单只负责字段收集、提交事件与展示语义。
