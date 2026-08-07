# Phase04-03 三类绑定关系、候选范围与上下文入口 Spec

## Why

`phase04-01` 已冻结页面边界与动作 owner，`phase04-02` 已冻结 `Product / Repository` 模板、状态语义与最小展示模型，但要真正进入后续实现，还必须把三类绑定关系的关系语义、候选读取范围、重复绑定语义、上下文入口跳转参数与 `phase02` 临时承接点的迁移边界写成单值结论。否则后续前端、后端、`.proto` 与验收会继续在绑定关系口径、候选过滤策略、旧入口保留级别与迁移兼容性之间来回漂移。

## What Changes

- 冻结 `product_repositories / product_modules / module_repositories` 三类绑定关系的关系语义
- 冻结 `BindRepositoryToProduct / BindModuleToProduct / MapModuleToRepository` 三类绑定动作的 canonical owner 与 reread 承接页面
- 冻结三类绑定动作的候选范围、候选排序、已绑定排除规则与候选展示模型
- 冻结候选读取空状态语义与重复绑定语义
- 冻结 `Module Detail` 旧入口的兼容跳转参数与保留级别
- 冻结 `Product Detail` 进入 `Repository Binding Detail / Workspace` 的上下文跳转参数
- 冻结 `phase02` 临时承接点的迁移边界与历史数据兼容前提
- 明确当前阶段不把 `Decision Center`、`Module Registry` 重新扩写为并列绑定主线

## Impact

- Affected specs: `phase04_product_and_repository_binding_foundation`
- Affected code: 后续 `frontend/` 中的 `Product Detail` 绑定面板、`Repository Binding Detail / Workspace` 绑定工作台、`Module Detail` 兼容跳转入口，后续 `backend/` 中的 `Product Registry` 与 `Repository Binding` 模块候选读取、绑定写入与迁移承接，后续 `.proto` 中的 `Binding` 消息字段

## ADDED Requirements

### Requirement: 三类绑定关系关系语义冻结

系统 SHALL 将 `phase04` 的三类绑定关系冻结为多对多关系，并明确各自的关系语义与唯一性约束。

#### Scenario: 判断 product_repositories 关系语义

- **WHEN** 后续 `/spec`、前端展示、后端校验或 `.proto` 合同讨论 `product_repositories` 关系
- **THEN** `product_repositories` 必须表示 `Product` 与 `Repository` 之间的多对多绑定关系
- **AND** 该关系的语义为"`Repository R` 是 `Product P` 的实现锚点"
- **AND** 同一 `(product_id, repository_id)` 对只允许存在一条绑定记录
- **AND** 该关系由 `BindRepositoryToProduct` 建立

#### Scenario: 判断 product_modules 关系语义

- **WHEN** 后续 `/spec`、前端展示、后端校验或 `.proto` 合同讨论 `product_modules` 关系
- **THEN** `product_modules` 必须表示 `Product` 与 `Module` 之间的多对多绑定关系
- **AND** 该关系的语义为"`Module M` 被 `Product P` 使用"
- **AND** 同一 `(product_id, module_id)` 对只允许存在一条绑定记录
- **AND** 该关系由 `BindModuleToProduct` 建立

#### Scenario: 判断 module_repositories 关系语义

- **WHEN** 后续 `/spec`、前端展示、后端校验或 `.proto` 合同讨论 `module_repositories` 关系
- **THEN** `module_repositories` 必须表示 `Module` 与 `Repository` 之间的多对多映射关系
- **AND** 该关系的语义为"`Module M` 实现于 `Repository R`"
- **AND** 同一 `(module_id, repository_id)` 对只允许存在一条绑定记录
- **AND** 该关系由 `MapModuleToRepository` 建立

### Requirement: BindRepositoryToProduct 候选范围冻结

系统 SHALL 将 `BindRepositoryToProduct` 的候选读取范围冻结为当前已存在且未绑定到当前 `Repository` 的 `active` 状态 `Product`。

#### Scenario: 判断 BindRepositoryToProduct 候选范围

- **WHEN** 用户在 `Repository Binding Detail / Workspace` 发起 `BindRepositoryToProduct`
- **THEN** 候选列表必须只包含 `status = active` 的 `Product`
- **AND** 必须排除已经绑定到当前 `Repository` 的 `Product`
- **AND** 候选排序必须按 `created_at` 降序（ newest first ）
- **AND** 不得将 `archived` 状态的 `Product` 纳入候选列表

### Requirement: BindModuleToProduct 候选范围冻结

系统 SHALL 将 `BindModuleToProduct` 的候选读取范围冻结为当前已存在且未绑定到当前 `Product` 的 `active` 状态 `Module`。

#### Scenario: 判断 BindModuleToProduct 候选范围

- **WHEN** 用户在 `Product Detail` 发起 `BindModuleToProduct`
- **THEN** 候选列表必须只包含 `status = active` 的 `Module`
- **AND** 必须排除已经绑定到当前 `Product` 的 `Module`
- **AND** 候选排序必须按 `created_at` 降序
- **AND** 不得将 `archived` 状态的 `Module` 纳入候选列表

### Requirement: MapModuleToRepository 候选范围冻结

系统 SHALL 将 `MapModuleToRepository` 的候选读取范围冻结为当前已存在且未映射到当前 `Repository` 的 `active` 状态 `Module`。

#### Scenario: 判断 MapModuleToRepository 候选范围

- **WHEN** 用户在 `Repository Binding Detail / Workspace` 发起 `MapModuleToRepository`
- **THEN** 候选列表必须只包含 `status = active` 的 `Module`
- **AND** 必须排除已经映射到当前 `Repository` 的 `Module`
- **AND** 候选排序必须按 `created_at` 降序
- **AND** 不得将 `archived` 状态的 `Module` 纳入候选列表

### Requirement: 候选读取空状态语义冻结

系统 SHALL 将三类绑定动作的候选读取空状态语义冻结为返回空列表，不得映射为资源不存在或接口错误。

#### Scenario: 判断候选读取空状态

- **WHEN** 任一绑定动作的候选读取返回零条候选记录
- **THEN** 系统必须返回空列表语义
- **AND** 不得将空结果映射为资源不存在
- **AND** 不得将空结果映射为接口错误
- **AND** 页面必须展示明确的无可绑定候选空状态提示

### Requirement: 候选读取展示模型冻结

系统 SHALL 将三类绑定动作的候选读取展示模型冻结为与已绑定展示模型一致的最小字段集合。

#### Scenario: 判断候选 Product 展示模型

- **WHEN** `BindRepositoryToProduct` 的候选读取返回候选 `Product` 列表
- **THEN** 每条候选至少必须承接 `product_id / product_name / product_status`
- **AND** 不得在当前阶段扩写候选 `Product` 的展示字段

#### Scenario: 判断候选 Module 展示模型

- **WHEN** `BindModuleToProduct` 或 `MapModuleToRepository` 的候选读取返回候选 `Module` 列表
- **THEN** 每条候选至少必须承接 `module_id / module_name / module_status`
- **AND** 不得在当前阶段扩写候选 `Module` 的展示字段

### Requirement: 重复绑定语义冻结

系统 SHALL 将三类绑定动作的重复绑定语义冻结为返回明确的重复冲突语义。

#### Scenario: 判断重复绑定语义

- **WHEN** 用户提交的绑定关系在对应绑定表中已存在相同记录
- **THEN** 系统必须返回明确的重复冲突语义
- **AND** 不得降级为静默成功
- **AND** 不得降级为模糊通用错误
- **AND** 不得通过 `ON CONFLICT DO NOTHING` 隐式吞掉重复冲突

### Requirement: 三类绑定动作 canonical owner 冻结

系统 SHALL 将三类绑定动作的 canonical owner 冻结为单值结论，不得在后续 `/spec` 或实现中出现并行 owner。

#### Scenario: 判断 canonical owner

- **WHEN** 接手者判断三类绑定动作的页面归属
- **THEN** 必须得到以下单值结论：
- **AND** `BindModuleToProduct` → `Product Detail`（ `Product Registry` 模块 ）
- **AND** `BindRepositoryToProduct` → `Repository Binding Detail / Workspace`（ `Repository Binding` 模块 ）
- **AND** `MapModuleToRepository` → `Repository Binding Detail / Workspace`（ `Repository Binding` 模块 ）
- **AND** 不得由 `Module Detail`、`Repository Binding` 或其他页面并行拥有第二套主写入流程

### Requirement: 绑定成功后 reread 承接页面冻结

系统 SHALL 将三类绑定动作成功后的 reread 承接页面冻结为对应 canonical owner 页面。

#### Scenario: 判断 reread 承接页面

- **WHEN** 任一绑定动作写入成功
- **THEN** 系统必须回到对应 canonical owner 页面完成 reread
- **AND** `BindModuleToProduct` 成功后必须回到当前 `Product Detail` 重新读取已绑定模块列表
- **AND** `BindRepositoryToProduct` 成功后必须回到当前 `Repository Binding Detail / Workspace` 重新读取已绑定产品列表
- **AND** `MapModuleToRepository` 成功后必须回到当前 `Repository Binding Detail / Workspace` 重新读取已映射模块列表
- **AND** 不得只靠 `toast` 或局部通知作为成功依据

### Requirement: Module Detail 旧入口兼容跳转与参数冻结

系统 SHALL 将 `Module Detail` 旧入口的保留级别冻结为仅允许兼容跳转，并冻结跳转参数。

#### Scenario: 判断 Module Detail 旧入口保留级别

- **WHEN** 后续 `/spec` 或实现讨论 `Module Detail` 的绑定能力
- **THEN** `Module Detail` 只允许保留进入正式主入口的兼容跳转
- **AND** 不得在 `Module Detail` 内直接提交 `BindModuleToProduct`、`BindRepositoryToProduct` 或 `MapModuleToRepository`
- **AND** 不得把 `Module Detail` 扩写为第二个绑定工作台

#### Scenario: 判断 Module Detail 兼容跳转参数

- **WHEN** 用户从 `Module Detail` 发起与 `Product / Repository` 绑定相关的后续动作
- **THEN** 必须进入对应绑定动作的正式主入口，并携带 `moduleId / moduleName / fromModuleDetail` 作为上下文参数
- **AND** 上下文参数只表示来源模块身份与来源页面标记，不表示目标实体身份
- **AND** 若目标实体尚未确定，必须先进入对应列表页（ `Product Registry / List` 或 `Repository Binding / List` ）选择目标实体
- **AND** 若目标实体已确定，必须额外携带目标页身份参数 `productId` 或 `repositoryId`，与上下文参数拆开传递
- **AND** 接收方页面必须能基于上下文参数预填对应绑定面板的候选 `Module` 选择

### Requirement: Product Detail 上下文入口跳转参数冻结

系统 SHALL 将 `Product Detail` 发起 `Repository` 绑定相关动作时的上下文入口跳转参数冻结为单值结论。

#### Scenario: 判断 Product Detail 上下文跳转参数

- **WHEN** 用户从 `Product Detail` 发起与 `Repository` 绑定相关的后续动作
- **THEN** 必须进入 `BindRepositoryToProduct` 的正式主入口，并携带 `productId / productName / fromProductDetail` 作为上下文参数
- **AND** 上下文参数只表示来源产品身份与来源页面标记，不表示目标 `Repository` 身份
- **AND** 若目标 `Repository` 尚未确定，必须先进入 `Repository Binding / List` 选择目标 `Repository`
- **AND** 若目标 `Repository` 已确定，必须额外携带目标页身份参数 `repositoryId`，与上下文参数拆开传递
- **AND** 接收方页面必须能基于上下文参数预填 `BindRepositoryToProduct` 面板的候选 `Product` 选择
- **AND** `Product Detail` 自身不得承接第二套仓库绑定写入流程

### Requirement: phase02 临时承接迁移边界冻结

系统 SHALL 将 `phase02` 中由 `Module Registry` 临时承接的绑定相关接口与数据访问的迁移边界冻结为单值结论。

#### Scenario: 判断 ProductBindingCandidateRead 迁移边界

- **WHEN** `phase04` 实现 `BindRepositoryToProduct` 的候选 `Product` 读取
- **THEN** `phase02` 中的 `ProductBindingCandidateRead`（ 读取 `products` 表 ）必须从 `Module Registry` 的 `candidate/` 子包迁移到 `Repository Binding` 模块的 `candidate/` 子包
- **AND** 迁移后的接口契约由 `Repository Binding` 模块拥有
- **AND** 迁移后的候选读取必须满足本规格冻结的候选范围（ `active` 状态、已绑定排除、`created_at` 降序 ）

#### Scenario: 判断 RepositoryBindingCandidateRead 迁移边界

- **WHEN** `phase04` 实现 `MapModuleToRepository` 的候选 `Module` 读取
- **THEN** `phase02` 中的 `RepositoryBindingCandidateRead`（ 读取 `repositories` 表 ）必须标记为废弃
- **AND** 废弃原因：`phase04` 中 `MapModuleToRepository` 的候选为 `Module` 而非 `Repository`，该接口不再有消费方
- **AND** 废弃后不得在 `phase04` 中保留该接口的并行实现

#### Scenario: 判断 ModuleBindingWrite 迁移边界

- **WHEN** `phase04` 实现三类绑定动作的写入
- **THEN** `phase02` 中的 `ModuleBindingWrite` 必须拆分迁移
- **AND** `BindModuleToProduct` 写入迁移到 `Product Registry` 模块
- **AND** `MapModuleToRepository` 写入迁移到 `Repository Binding` 模块
- **AND** `BindRepositoryToProduct` 写入作为 `Repository Binding` 模块的新增能力实现

#### Scenario: 判断 binding_store 迁移边界

- **WHEN** `phase04` 实现绑定数据访问
- **THEN** `phase02` 中的 `binding_store.go`（ 管理 `product_modules` + `module_repositories` ）必须拆分迁移
- **AND** `product_modules` 数据访问迁移到 `Product Registry` 模块
- **AND** `module_repositories` 数据访问迁移到 `Repository Binding` 模块
- **AND** `product_repositories` 数据访问作为 `Repository Binding` 模块的新增能力实现

#### Scenario: 判断历史绑定数据兼容前提

- **WHEN** `phase04` 完成迁移后
- **THEN** `phase02` 中已存在的 `product_modules` 与 `module_repositories` 历史绑定数据必须保持可读
- **AND** 不得通过重建影子表、第二套绑定表或临时双写绕过迁移
- **AND** 既有前端入口升级后必须仍可继续读取历史绑定结果

### Requirement: 候选读取接口归属冻结

系统 SHALL 将三类绑定动作的候选读取接口归属冻结为消费方模块拥有，遵循 `phase02` 已建立的跨模块候选读取模式。

#### Scenario: 判断 Product Registry 模块候选读取归属

- **WHEN** `Product Detail` 需要 `BindModuleToProduct` 的候选 `Module` 读取
- **THEN** 候选读取接口必须由 `Product Registry` 模块的 `candidate/` 子包定义和拥有
- **AND** 必须通过独立 `Read` 接口隔离
- **AND** `Product Registry` 的 `service` 不得直接写跨模块 `SQL` 读取 `modules` 表

#### Scenario: 判断 Repository Binding 模块候选读取归属

- **WHEN** `Repository Binding Detail / Workspace` 需要 `BindRepositoryToProduct` 的候选 `Product` 读取或 `MapModuleToRepository` 的候选 `Module` 读取
- **THEN** 两个候选读取接口必须由 `Repository Binding` 模块的 `candidate/` 子包定义和拥有
- **AND** 必须通过独立 `Read` 接口隔离
- **AND** `Repository Binding` 的 `service` 不得直接写跨模块 `SQL` 读取 `products` 或 `modules` 表

### Requirement: 非目标冻结

系统 SHALL 明确当前阶段不把 `Decision Center`、`Module Registry` 重新扩写为并列绑定主线。

#### Scenario: 判断非目标边界

- **WHEN** 后续 `/spec` 或实现讨论 `phase04` 的绑定主线
- **THEN** 不得把 `Decision Center` 扩写为 `Product / Repository` 绑定写入主线
- **AND** 不得把 `Module Registry` 重新扩写为绑定写入主线
- **AND** `Module Detail` 只允许保留兼容跳转，不扩写为并列绑定工作台
- **AND** 不得引入 `Decision -> Product / Repository` 正式关联写入作为当前阶段范围

## MODIFIED Requirements

### Requirement: Module Detail 绑定承接方式解释

`Module Detail` 在 `phase04` 中 SHALL 从 `phase02` 的绑定写入承接位回落为摘要展示与兼容跳转入口，并且兼容跳转的上下文参数已冻结为 `moduleId / moduleName / fromModuleDetail`。

#### Scenario: Module Detail 绑定承接解释

- **WHEN** 后续 `/spec` 或实现讨论 `Module Detail` 的绑定能力
- **THEN** 必须将其理解为摘要展示与正式主入口跳转位
- **AND** 跳转参数必须使用 `moduleId / moduleName / fromModuleDetail`
- **AND** 不得继续把 `Module Detail` 实现为 `BindModuleToProduct`、`BindRepositoryToProduct`、`MapModuleToRepository` 的并列写入 owner

### Requirement: phase02 候选读取接口归属解释

`phase02` 中由 `Module Registry` 临时承接的 `ProductBindingCandidateRead` 与 `RepositoryBindingCandidateRead` SHALL 不再被解释为"长期由 `Module Registry` 承接的临时接口"，而必须被解释为已有明确迁移边界的历史接口。

#### Scenario: phase02 候选读取接口迁移解释

- **WHEN** 后续 `/spec` 或实现讨论 `phase02` 候选读取接口的归属
- **THEN** `ProductBindingCandidateRead` 必须理解为迁移到 `Repository Binding` 模块的 `candidate/` 子包
- **AND** `RepositoryBindingCandidateRead` 必须理解为已废弃
- **AND** 不得继续把 `Module Registry` 解释为这两个接口的长期 owner

## REMOVED Requirements

### Requirement: Module Detail 作为绑定写入主入口

**Reason**: `phase04` 已将三类绑定动作的 canonical owner 冻结为 `Product Detail` 与 `Repository Binding Detail / Workspace`，`Module Detail` 不再作为绑定写入主入口。
**Migration**: `Module Detail` 既有绑定入口在迁移后只允许保留为兼容跳转，可以带上 `moduleId / moduleName / fromModuleDetail` 上下文进入正式主入口，但不得继续停留在本页直接提交绑定写入。

### Requirement: 基于 archived 记录建立新绑定

**Reason**: `phase04-03` 已将三类绑定动作的候选范围冻结为仅包含 `active` 状态记录，`archived` 记录不进入候选列表。`archived` 语义为"已归档保留的历史事实"，`active` 语义为"可继续绑定的有效状态"。
**Migration**: 若后续确需允许 `archived` 记录作为绑定候选，必须进入新的冻结任务重新单值化，不在 `phase04-03` 当前规格中处理。
