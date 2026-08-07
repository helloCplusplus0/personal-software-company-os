# Phase04-04 数据读写范围、接口边界与错误语义前提 Spec

## Why

`phase04-01` 已冻结页面边界与动作 owner，`phase04-02` 已冻结 `Product / Repository` 模板、状态语义与最小展示模型，`phase04-03` 已冻结三类绑定关系语义、候选范围、上下文入口、canonical owner、reread 承接页面、候选读取接口归属与 `phase02` 迁移边界。但要真正进入后续实现设计（ `phase04-05 ~ phase04-09` ），还必须把 `Product / Repository / Binding` 的最小数据读写范围、最小接口边界与错误语义前提写成单值结论。否则后续前端状态模型、后端模块边界、`.proto` 合同与验收异常路径会继续在数据范围、错误语义与空状态口径之间来回漂移，且关键异常路径会被推迟到联调时临时补写。

> 阶段分工约束：本规格只冻结 `phase04-04` 范围内的数据读写范围、最小接口边界与错误语义前提。接口分组、方向级 API 矩阵、canonical owner 兼容策略、候选读取接线位置、合同演进约束与实现工具选型分别由 `phase04-07`（后端模块边界与接口分组设计）与 `phase04-08`（`.proto` 合同设计）承接，不在本规格中提前冻结。

## What Changes

- 冻结 `Product Registry` 与 `Repository Binding` 两个模块的最小数据读写范围（列表读取、详情读取、创建写入、绑定写入、候选读取）
- 冻结列表读取、详情读取、创建写入、绑定写入与候选读取的最小接口边界（输入参数、输出数据范围、筛选与排序）
- 冻结创建失败、目标不存在、重复绑定、候选空结果与列表空结果的错误语义前提
- 冻结三类绑定写入的校验失败类型
- 明确当前阶段不提前冻结 Dashboard 聚合接口

## Impact

- Affected specs: `phase04_product_and_repository_binding_foundation`
- Affected code: 后续 `frontend/` 中的列表、详情、创建、绑定工作台的数据消费模型与错误处理，后续 `backend/` 中的 `Product Registry` 与 `Repository Binding` 模块数据范围校验与错误哨兵值，后续 `.proto` 中的消息字段范围前提

## ADDED Requirements

### Requirement: Product Registry 最小数据读写范围冻结

系统 SHALL 将 `Product Registry` 模块的最小数据读写范围冻结为单值结论，覆盖列表读取、详情读取、创建写入、绑定写入与候选读取。

> 本规格只冻结数据读写范围（字段、筛选、排序）。接口分组、方向级 API 矩阵与接线位置由 `phase04-07` 承接。本规格中使用的接口名（如 `ProductListRead`、`ProductDetailRead` 等）仅为引用数据范围的目的，不冻结接口分组或接口命名决策。

#### Scenario: 判断 Product List 读取数据范围

- **WHEN** 后续 `/spec`、前端展示、后端校验或 `.proto` 合同讨论 `ProductListRead` 的数据范围
- **THEN** 列表读取至少必须承接 `name / description / status / created_at / module_bind_count / repository_bind_count`
- **AND** 列表读取的筛选参数只允许 `queryText` 与 `statusFilter`
- **AND** `queryText` 只匹配 `name` 字段（模糊匹配）
- **AND** `statusFilter` 只允许 `all / active / archived`，且 `all` 只存在于 UI/路由层，不得写入数据库、HTTP 持久化 DTO 或 `.proto` 持久化字段
- **AND** 列表排序必须按 `created_at` 降序（ newest first ）
- **AND** 当前阶段不引入分页参数

#### Scenario: 判断 Product Detail 读取数据范围

- **WHEN** 后续 `/spec`、前端展示、后端校验或 `.proto` 合同讨论 `ProductDetailRead` 的数据范围
- **THEN** 详情读取至少必须承接核心对象字段 `id / name / description / status / created_at`
- **AND** 必须承接已绑定模块列表，每条至少包含 `module_id / module_name / module_status`
- **AND** 必须承接已绑定仓库列表，每条至少包含 `repository_id / repository_name / provider / repository_status`
- **AND** 不得在当前阶段扩写详情读取的展示字段

#### Scenario: 判断 CreateProduct 写入数据范围

- **WHEN** 后续 `/spec`、前端表单、后端 DTO 或 `.proto` 合同讨论 `ProductCreateWrite` 的数据范围
- **THEN** 创建写入必须承接 `name / description / status`
- **AND** `name / description / status` 必须为必填字段
- **AND** 必填字段在去首尾空白后不得为空字符串
- **AND** `status` 必须属于 `active / archived`，默认预填并显式提交 `active`
- **AND** 不得在当前阶段额外引入 `customer / value_proposition / business_model / metrics / remote_import_source` 等超出 `v0.1` 的字段

#### Scenario: 判断 BindModuleToProduct 写入数据范围

- **WHEN** 后续 `/spec`、前端表单、后端 DTO 或 `.proto` 合同讨论 `ProductModuleBindingWrite` 的数据范围
- **THEN** 绑定写入必须承接 `product_id / module_id`
- **AND** `product_id / module_id` 必须为必填字段
- **AND** `product_id` 由当前 `Product Detail` 上下文隐式承接
- **AND** `module_id` 由用户从候选列表中选择
- **AND** 不得在当前阶段引入额外的绑定属性字段（如绑定权重、绑定备注）

#### Scenario: 判断 Product Module 候选读取数据范围

- **WHEN** 后续 `/spec`、前端展示、后端校验或 `.proto` 合同讨论 `ProductModuleCandidateRead` 的数据范围
- **THEN** 候选读取至少必须承接每条候选的 `module_id / module_name / module_status`
- **AND** 候选范围必须只包含 `status = active` 的 `Module`
- **AND** 必须排除已经绑定到当前 `Product` 的 `Module`
- **AND** 候选排序必须按 `created_at` 降序
- **AND** 不得将 `archived` 状态的 `Module` 纳入候选列表

### Requirement: Repository Binding 最小数据读写范围冻结

系统 SHALL 将 `Repository Binding` 模块的最小数据读写范围冻结为单值结论，覆盖列表读取、详情读取、创建写入、绑定写入与候选读取。

> 本规格只冻结数据读写范围（字段、筛选、排序）。接口分组、方向级 API 矩阵与接线位置由 `phase04-07` 承接。本规格中使用的接口名仅为引用数据范围的目的，不冻结接口分组或接口命名决策。

#### Scenario: 判断 Repository List 读取数据范围

- **WHEN** 后续 `/spec`、前端展示、后端校验或 `.proto` 合同讨论 `RepositoryListRead` 的数据范围
- **THEN** 列表读取至少必须承接 `name / url / provider / status / created_at / product_bind_count / module_bind_count`
- **AND** 列表读取的筛选参数只允许 `queryText` 与 `statusFilter`
- **AND** `queryText` 只匹配 `name` 字段（模糊匹配）
- **AND** `statusFilter` 只允许 `all / active / archived`，且 `all` 只存在于 UI/路由层
- **AND** 列表排序必须按 `created_at` 降序
- **AND** 当前阶段不引入分页参数
- **AND** 当前阶段不引入 `providerFilter`

#### Scenario: 判断 Repository Detail 读取数据范围

- **WHEN** 后续 `/spec`、前端展示、后端校验或 `.proto` 合同讨论 `RepositoryDetailRead` 的数据范围
- **THEN** 详情读取至少必须承接核心对象字段 `id / name / url / provider / status / created_at`
- **AND** 必须承接已绑定产品列表，每条至少包含 `product_id / product_name / product_status`
- **AND** 必须承接已映射模块列表，每条至少包含 `module_id / module_name / module_status`
- **AND** 不得在当前阶段扩写详情读取的展示字段

#### Scenario: 判断 CreateRepository 写入数据范围

- **WHEN** 后续 `/spec`、前端表单、后端 DTO 或 `.proto` 合同讨论 `RepositoryCreateWrite` 的数据范围
- **THEN** 创建写入必须承接 `name / url / provider / status`
- **AND** `name / url / provider / status` 必须为必填字段
- **AND** 必填字段在去首尾空白后不得为空字符串
- **AND** `status` 必须属于 `active / archived`，默认预填并显式提交 `active`
- **AND** `provider` 必须为必填字符串字段，不采用受控枚举
- **AND** 不得在当前阶段额外引入 `oauth_binding / remote_import_status / sync_cursor / scanned_commit` 等自动化集成字段

#### Scenario: 判断 BindRepositoryToProduct 写入数据范围

- **WHEN** 后续 `/spec`、前端表单、后端 DTO 或 `.proto` 合同讨论 `RepositoryProductBindingWrite` 的数据范围
- **THEN** 绑定写入必须承接 `repository_id / product_id`
- **AND** `repository_id / product_id` 必须为必填字段
- **AND** `repository_id` 由当前 `Repository Binding Detail / Workspace` 上下文隐式承接
- **AND** `product_id` 由用户从候选列表中选择
- **AND** 不得在当前阶段引入额外的绑定属性字段

#### Scenario: 判断 MapModuleToRepository 写入数据范围

- **WHEN** 后续 `/spec`、前端表单、后端 DTO 或 `.proto` 合同讨论 `RepositoryModuleMappingWrite` 的数据范围
- **THEN** 映射写入必须承接 `repository_id / module_id`
- **AND** `repository_id / module_id` 必须为必填字段
- **AND** `repository_id` 由当前 `Repository Binding Detail / Workspace` 上下文隐式承接
- **AND** `module_id` 由用户从候选列表中选择
- **AND** 不得在当前阶段引入额外的映射属性字段

#### Scenario: 判断 Repository Product 候选读取数据范围

- **WHEN** 后续 `/spec`、前端展示、后端校验或 `.proto` 合同讨论 `ProductBindingCandidateRead` 的数据范围
- **THEN** 候选读取至少必须承接每条候选的 `product_id / product_name / product_status`
- **AND** 候选范围必须只包含 `status = active` 的 `Product`
- **AND** 必须排除已经绑定到当前 `Repository` 的 `Product`
- **AND** 候选排序必须按 `created_at` 降序
- **AND** 不得将 `archived` 状态的 `Product` 纳入候选列表

#### Scenario: 判断 Repository Module 候选读取数据范围

- **WHEN** 后续 `/spec`、前端展示、后端校验或 `.proto` 合同讨论 `RepositoryModuleCandidateRead` 的数据范围
- **THEN** 候选读取至少必须承接每条候选的 `module_id / module_name / module_status`
- **AND** 候选范围必须只包含 `status = active` 的 `Module`
- **AND** 必须排除已经映射到当前 `Repository` 的 `Module`
- **AND** 候选排序必须按 `created_at` 降序
- **AND** 不得将 `archived` 状态的 `Module` 纳入候选列表

### Requirement: 详情读取与候选读取边界冻结

系统 SHALL 将详情读取与候选读取的接口边界冻结为独立承接，不得合并或拆散。

#### Scenario: 判断详情读取与候选读取独立性

- **WHEN** 后续 `/spec`、后端实现或 `.proto` 合同讨论详情读取与候选读取的关系
- **THEN** `ProductDetailRead` 只承接详情本体、已绑定模块列表与已绑定仓库列表
- **AND** `ProductModuleCandidateRead` 必须通过独立 request / response 承接候选读取
- **AND** `RepositoryDetailRead` 只承接详情本体、已绑定产品列表与已映射模块列表
- **AND** `ProductBindingCandidateRead` 与 `RepositoryModuleCandidateRead` 必须通过独立 request / response 承接候选读取
- **AND** 不得把候选读取结果直接并入 `ProductDetailRead` 或 `RepositoryDetailRead` 的最小数据边界
- **AND** 不得把候选读取拆成需要前端自行拼装的多个独立业务入口

### Requirement: 错误语义前提冻结

系统 SHALL 将 `Product Registry` 与 `Repository Binding` 两个模块的错误语义前提冻结为单值结论，覆盖创建失败、目标不存在、重复绑定、候选空结果与列表空结果。

#### Scenario: 判断创建失败错误语义

- **WHEN** 用户提交缺少必填字段或非法 `status` 的 `CreateProduct`
- **THEN** 系统必须返回明确的校验失败语义
- **AND** 归属接口为 `ProductCreateWrite`
- **AND** 不得降级为模糊通用错误
- **AND** 不得出现 `500` 级未收口错误替代业务错误
- **WHEN** 用户提交缺少必填字段或非法 `status` 的 `CreateRepository`
- **THEN** 系统必须返回明确的校验失败语义
- **AND** 归属接口为 `RepositoryCreateWrite`
- **AND** 不得降级为模糊通用错误
- **AND** 不得出现 `500` 级未收口错误替代业务错误

#### Scenario: 判断目标不存在错误语义

- **WHEN** `ProductDetailRead` 接收到不存在的 `product_id`
- **THEN** 系统必须返回资源不存在语义
- **AND** 归属接口为 `ProductDetailRead`
- **WHEN** `RepositoryDetailRead` 接收到不存在的 `repository_id`
- **THEN** 系统必须返回资源不存在语义
- **AND** 归属接口为 `RepositoryDetailRead`
- **WHEN** `BindModuleToProduct` 的目标 `Product` 或 `Module` 不存在
- **THEN** 系统必须返回资源不存在语义
- **AND** 归属接口为 `ProductModuleBindingWrite`
- **WHEN** `BindRepositoryToProduct` 的目标 `Product` 或 `Repository` 不存在
- **THEN** 系统必须返回资源不存在语义
- **AND** 归属接口为 `RepositoryProductBindingWrite`
- **WHEN** `MapModuleToRepository` 的目标 `Module` 或 `Repository` 不存在
- **THEN** 系统必须返回资源不存在语义
- **AND** 归属接口为 `RepositoryModuleMappingWrite`
- **WHEN** `ProductModuleCandidateRead` 接收到不存在的 `product_id`
- **THEN** 系统必须返回资源不存在语义
- **AND** 归属接口为 `ProductModuleCandidateRead`
- **WHEN** `ProductBindingCandidateRead` 或 `RepositoryModuleCandidateRead` 接收到不存在的 `repository_id`
- **THEN** 系统必须返回资源不存在语义
- **AND** 归属接口为对应候选读取接口

#### Scenario: 判断重复绑定错误语义

- **WHEN** 用户提交的绑定关系在对应绑定表中已存在相同记录
- **THEN** 系统必须返回明确的重复冲突语义
- **AND** `BindModuleToProduct` 重复绑定的归属接口为 `ProductModuleBindingWrite`
- **AND** `BindRepositoryToProduct` 重复绑定的归属接口为 `RepositoryProductBindingWrite`
- **AND** `MapModuleToRepository` 重复映射的归属接口为 `RepositoryModuleMappingWrite`
- **AND** 不得降级为静默成功
- **AND** 不得降级为模糊通用错误
- **AND** 不得通过 `ON CONFLICT DO NOTHING` 隐式吞掉重复冲突

#### Scenario: 判断候选读取空结果语义

- **WHEN** 任一候选读取接口返回零条候选记录
- **THEN** 系统必须返回空列表语义
- **AND** 不得将空结果映射为资源不存在
- **AND** 不得将空结果映射为接口错误
- **AND** 不得把空候选结果误报为接口失败
- **AND** 页面必须展示明确的无可绑定候选空状态提示

#### Scenario: 判断列表读取空结果语义

- **WHEN** `ProductListRead` 或 `RepositoryListRead` 返回零条记录
- **THEN** 系统必须返回空列表语义
- **AND** 不得将空结果映射为资源不存在或接口错误
- **AND** 页面必须展示明确的空状态引导

#### Scenario: 判断三类绑定写入的校验失败类型

- **WHEN** 后续 `/spec`、后端实现或验收讨论三类绑定写入的校验失败类型
- **THEN** 必须得到以下单值结论：
- **AND** 必填字段缺失（`product_id / module_id / repository_id` 缺失）→ 校验失败
- **AND** 目标实体不存在 → 资源不存在
- **AND** 目标实体非 `active` 状态 → 校验失败
- **AND** 重复绑定 → 重复冲突
- **AND** 不得出现 `500` 级未收口错误替代上述业务错误

### Requirement: 非目标冻结

系统 SHALL 明确当前阶段不提前冻结超出 `v0.1` 的 Dashboard 聚合接口与跨模块聚合读取接口。

#### Scenario: 判断非目标边界

- **WHEN** 后续 `/spec` 或实现讨论 `phase04-04` 的接口范围
- **THEN** 不得提前冻结 Dashboard 聚合反馈接口
- **AND** 不得提前冻结 `product_asset_coverage` 或类似跨模块聚合读取接口
- **AND** 不得提前冻结 `Decision -> Product / Repository` 正式关联写入接口
- **AND** 不得提前冻结超出当前阶段的分页、复杂检索或批量操作接口

## MODIFIED Requirements

### Requirement: phase02 临时承接接口迁移后的迁移边界解释

`phase02` 中由 `Module Registry` 临时承接的 `ProductBindingCandidateRead`、`RepositoryBindingCandidateRead` 与 `ModuleBindingWrite` SHALL 不再被解释为"长期由 `Module Registry` 承接的临时接口"，而必须被解释为已有明确迁移边界的历史接口。

> 本规格只解释迁移边界（迁移目标模块与废弃状态），不冻结接口名是否沿用、拆分后的具体接口名或接口分组。接口名沿用、接口分组与方向级 API 矩阵由 `phase04-07` 统一收敛。

#### Scenario: phase02 临时承接接口迁移后的迁移边界解释

- **WHEN** 后续 `/spec` 或实现讨论 `phase02` 临时承接接口的归属与分组
- **THEN** `ProductBindingCandidateRead` 必须理解为迁移到 `Repository Binding` 模块，由该模块拥有
- **AND** `RepositoryBindingCandidateRead` 必须理解为已废弃，不在 `phase04` 中保留并行实现
- **AND** `ModuleBindingWrite` 必须理解为已拆分迁移到 `Product Registry` 模块与 `Repository Binding` 模块
- **AND** 不得继续把 `Module Registry` 解释为这些接口的长期 owner

## REMOVED Requirements

### Requirement: Dashboard 聚合接口

**Reason**: `phase04-04` 只冻结 `Product / Repository / Binding` 的最小数据读写范围与接口边界，不提前冻结超出当前阶段的 Dashboard 聚合接口。Dashboard 聚合反馈属于 `shared_baseline` §8 非目标矩阵，且 `architecture_plan` §5.3 已明确"Dashboard 最小反馈闭环实现"为本阶段不做项。
**Migration**: 若后续确需引入 Dashboard 聚合反馈接口，必须进入新的冻结任务重新单值化，不在 `phase04-04` 当前规格中处理。

### Requirement: 超出当前阶段的分页与复杂检索接口

**Reason**: `phase04-04` 冻结的列表读取接口只承接当前阶段最小主线所需的 `queryText + statusFilter` 筛选与 `created_at` 降序排序，不引入分页参数。分页与复杂检索属于后续阶段优化项，提前冻结会扩大当前阶段接口边界。
**Migration**: 若后续确需引入分页或复杂检索接口，必须进入新的冻结任务重新单值化，不在 `phase04-04` 当前规格中处理。
