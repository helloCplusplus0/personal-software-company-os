# Phase02-02 Module 实体与版本主线 Spec

## Why

`phase02-01` 已经冻结了 `Module Registry` 的页面边界，但页面要真正可实现，仍然需要把 `Module` 在列表页和详情页到底展示什么、`Release` 如何登记和呈现、以及模块准入规则如何在页面上被用户理解和承接写成单值结论。否则后续 `/spec` 与实现仍会在字段、状态和版本主线之间继续漂移。

## What Changes

- 冻结 `Module` 在 `Module Registry / List` 与 `Module Detail` 中的最小展示字段
- 冻结 `Module` 的最小状态表达及其在页面中的展示方式
- 冻结 `Release` 的最小登记字段、最小展示字段与版本主线承接方式
- 冻结模块准入规则在页面侧的承接方式
- 明确当前阶段不引入超出正式 MVP 规格正文的新对象解释

## Impact

- Affected specs: `phase02_module_registry_foundation`
- Affected code: 后续 `frontend/` 中的模块列表卡片、模块详情信息区、版本列表区、版本登记表单与状态展示逻辑

## ADDED Requirements

### Requirement: Module 最小展示字段冻结

系统 SHALL 将 `Module` 在当前阶段的最小展示字段冻结为可同时服务列表读取与详情读取的一组单值字段。

#### Scenario: 列表页最小展示字段

- **WHEN** 用户在 `Module Registry / List` 查看模块列表
- **THEN** 每个模块项至少应展示 `name / description / status / latest_release / product_bind_count / repository_bind_count`
- **AND** 不得在当前阶段扩写为需要额外业务对象才能成立的复杂展示结构

#### Scenario: 详情页最小展示字段

- **WHEN** 用户进入 `Module Detail`
- **THEN** 页面至少应展示模块核心对象字段
- **AND** 必须展示版本列表、产品绑定、仓库映射与相关 `Decision` 入口
- **AND** 不得要求用户跳转其他独立主线才能理解该模块当前状态

### Requirement: Module 状态表达冻结

系统 SHALL 将当前阶段 `Module` 的最小状态表达冻结为页面可直接理解的状态集合，并与上游正式规格正文保持一致。

#### Scenario: Module 状态范围

- **WHEN** 后续 `/spec` 或实现讨论 `Module` 状态
- **THEN** 当前阶段只允许使用可直接支持注册与版本主线的最小状态集合
- **AND** 推荐最小状态集合为 `active / archived`
- **AND** 不得在当前阶段额外引入新的复杂生命周期体系作为实现前提

#### Scenario: Module 状态展示

- **WHEN** 用户在列表页或详情页查看某个模块
- **THEN** 页面必须明确展示该模块当前 `status`
- **AND** 状态表达必须同时适用于桌面与移动浏览器

### Requirement: Release 最小登记主线冻结

系统 SHALL 将 `Release` 冻结为从 `Module Detail` 进入、服务于版本主线的最小登记对象。

#### Scenario: Release 最小登记字段

- **WHEN** 用户执行 `CreateRelease`
- **THEN** 当前阶段最小登记字段至少应包含 `version / status / released_at`
- **AND** `module_id` 必须由当前模块上下文隐式承接，而不是要求用户重新选择
- **AND** 不得把 `Release` 扩写为独立页面主线之外的复杂对象体系

### Requirement: Release 最小展示方式冻结

系统 SHALL 将 `Release` 的最小展示方式冻结为 `Module Detail` 内的版本主线区块，而不是独立的版本管理工作台。

#### Scenario: Release 展示判定

- **WHEN** 用户进入 `Module Detail`
- **THEN** 页面必须能够看到该模块的版本列表或等价版本区块
- **AND** 必须能够识别 `latest_release`
- **AND** 当前阶段不得将 `Release` 扩写为独立主线页面集合

### Requirement: 模块准入规则页面承接冻结

系统 SHALL 将正式 MVP 规格正文中的模块准入规则冻结为当前阶段页面可感知、可承接的最小规则。

#### Scenario: 准入规则承接

- **WHEN** 用户创建或查看某个模块
- **THEN** 页面必须能够承接以下规则：
- **AND** `name` 与 `description` 未定义时，不得视为进入注册主线
- **AND** `name` 必须保持唯一
- **AND** `status` 必须明确后才进入版本主线
- **AND** 存在至少一个可追踪 `Release` 后，才进入版本主线
- **AND** 未绑定 `Product` 或未映射 `Repository` 不阻断登记，但不得进入有效复用反馈统计

#### Scenario: 页面侧提示策略

- **WHEN** 模块已登记但尚未满足版本主线或复用反馈条件
- **THEN** 页面必须通过状态说明、提示文案或结构化分区表达当前所处阶段
- **AND** 不得把这些规则隐藏为仅后端可见逻辑

## MODIFIED Requirements

### Requirement: Module Detail 信息结构

`Module Detail` 在当前阶段 SHALL 不只展示静态字段，而应被解释为“核心对象字段 + 版本主线 + 绑定关系 + 决策入口”的最小信息结构容器。

#### Scenario: 详情页结构解释

- **WHEN** 后续 `/spec` 或实现讨论 `Module Detail`
- **THEN** 必须将其理解为页面主线中的综合承接页
- **AND** 不得把版本、绑定关系与状态说明散落为无归属入口

## REMOVED Requirements

### Requirement: Module / Release 复杂生命周期前置
**Reason**: 当前阶段只需要支持 `Module Registry` 最小可执行主线，不需要引入超出正式规格正文的复杂生命周期体系。
**Migration**: 若后续需要更细的状态机或版本治理规则，必须在后续 `phase / audit` 中单独冻结，不在 `phase02-02` 当前规格中处理。
