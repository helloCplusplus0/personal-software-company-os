# Phase03-07 后端模块边界与接口分组设计 Spec

## Why

`phase03-04` 已经冻结了 `Decision Center` 的最小数据读写范围、接口边界与错误语义前提，`phase03-05 / 06` 又把前端页面结构与交互流收敛到了可直接进入实现的粒度。当前仍缺少一份后端实现前设计结果，用来回答“哪些能力归 `Decision Center` 后端模块自己负责、哪些能力只通过与 `Module Registry` 的服务侧连接边界承接、读写接口应该如何分组、这些接口落到 `backend/` 的哪些文件上”，否则后续实现仍会在模块归属、接口入口、跨模块调用与代码组织上发生漂移。

## What Changes

- 冻结 `Decision Center` 在后端的最小模块边界
- 冻结列表读取、详情读取、候选读取、创建写入、关联写入的接口分组
- 冻结 `Decision Center` 与 `Module Registry` 的服务侧连接边界
- 冻结读接口、写接口与候选读取接口的职责拆分
- 冻结后端文件/目录落点到可直接创建文件的层级
- 明确当前阶段不提前冻结 Go HTTP/RPC 框架、数据访问层工具或完整合同文件结构

## Impact

- Affected specs: `phase03_decision_center_foundation`
- Affected code: 后续 `backend/` 中 `Decision Center` 的 handler/service/repository/candidate 边界、读写接口组织与跨模块连接方式
- Affected code: 预期会直接映射到 `backend/internal/decisioncenter/` 下的 `handler`、`service`、`repository` 与 `candidate` 文件落点

## ADDED Requirements

### Requirement: Decision Center 后端边界必须冻结为单一业务模块

系统 SHALL 将 `Decision Center` 在 `phase03` 的后端承接方式冻结为单一业务模块，由其统一承接当前阶段属于决策主线的读取、写入、候选读取与关系写入能力。

#### Scenario: 当前阶段模块归属

- **WHEN** 后续实现定义 `Decision Center` 的后端模块边界
- **THEN** `RecordDecision`、`LinkDecisionToTarget`、`DecisionListRead`、`DecisionDetailRead`、`DecisionModuleCandidateRead` 必须归属于 `Decision Center` 后端模块
- **AND** 不得把这些能力拆散到 `Module Registry`、`Product` 或 `Repository` 的后端模块中

#### Scenario: 当前阶段非归属能力

- **WHEN** 后续实现讨论 `CreateModule`、`CreateProduct`、`CreateRepository`
- **THEN** 必须判定为当前阶段 `Decision Center` 后端模块的非归属能力
- **AND** `Decision Center` 只能通过最小连接边界读取或校验 `Module`，不得吸收其他对象的主线写入

### Requirement: 后端接口必须按读组与写组拆分

系统 SHALL 将 `Decision Center` 的后端接口冻结为读组与写组两大类，并在读组内部保留候选读取子组，避免把读取语义与写入语义混入同一入口。本 spec 的接口命名与分组沿用 `phase03-04` 已冻结的 `*Read` / `*Write` 命名体系，不引入第二套命名。

#### Scenario: 最小读接口分组

- **WHEN** 后续实现设计后端读取入口
- **THEN** 最小读接口分组至少应包含 `DecisionListRead`、`DecisionDetailRead` 与 `DecisionModuleCandidateRead`
- **AND** `DecisionListRead` 只承接列表展示、筛选入口与进入详情所需最小读取
- **AND** `DecisionDetailRead` 只承接详情展示、已关联目标展示与最小来源上下文读取
- **AND** `DecisionModuleCandidateRead` 只承接 `Decision -> Module` 候选读取

#### Scenario: 最小写接口分组

- **WHEN** 后续实现设计后端写入入口
- **THEN** 最小写接口分组至少应包含 `DecisionWrite` 与 `DecisionLinkWrite`
- **AND** `DecisionWrite` 只承接 `RecordDecision`
- **AND** `DecisionLinkWrite` 只承接 `LinkDecisionToTarget`
- **AND** 不得把创建写入与目标关联写入混成单个无边界的大接口

### Requirement: Decision List 读组必须只承接列表语义

系统 SHALL 将列表读组冻结为只服务 `Decision Center / List` 的读取接口，而不是提前膨胀成通用搜索中心。

#### Scenario: 列表读取职责

- **WHEN** 后续实现 `DecisionListRead`
- **THEN** 其职责只承接列表最小字段、筛选前提与进入详情所需标识
- **AND** 最小读取字段必须与 `phase03-02` 与 `shared_baseline §5.1` 已冻结的列表读模型保持一致
- **AND** 当前阶段不提前扩写分页统计、Dashboard 聚合或跨对象分析读取

### Requirement: Decision Detail 读组必须作为后端读模型宿主

系统 SHALL 将 `DecisionDetailRead` 冻结为 `Decision Detail` 页面的统一后端读模型宿主，集中承接详情页当前阶段所有必要读取。

#### Scenario: 详情读取承接范围

- **WHEN** 后续实现 `DecisionDetailRead`
- **THEN** 它必须统一承接决策核心字段、结构化模板字段、最小来源上下文与已关联 `Module` 列表读取
- **AND** 不得把这些详情读取拆成需要前端自行拼装的多个独立业务入口

#### Scenario: 详情读模型与候选读取的边界

- **WHEN** `DecisionDetailRead` 需要与 `Decision -> Module` 候选读取协作
- **THEN** `DecisionDetailRead` 只承接详情本体与已建立关联结果
- **AND** 候选读取必须由 `DecisionModuleCandidateRead` 独立承接
- **AND** 不得把候选读取结果直接并入 `DecisionDetailRead` 的最小合同边界

### Requirement: 创建写组必须只负责 Decision 自身登记

系统 SHALL 将 `DecisionWrite` 冻结为只承接 `RecordDecision` 的写入入口，不混入目标关联写入。

#### Scenario: 创建写入职责

- **WHEN** 后续实现 `RecordDecision`
- **THEN** `DecisionWrite` 只负责 `Decision` 自身最小结构化字段写入
- **AND** 成功结果必须能返回新建 `Decision` 标识，以支持前端回流到 `DecisionDetailPage`

### Requirement: 关联写组必须统一承接目标关联写入

系统 SHALL 将 `LinkDecisionToTarget` 冻结为统一的关系写入组，由 `DecisionLinkWrite` 承接。

#### Scenario: 关系写入职责

- **WHEN** 后续实现 `LinkDecisionToTarget`
- **THEN** `DecisionLinkWrite` 必须统一承接当前阶段允许的目标关联写入
- **AND** 当前阶段唯一必交付目标类型是 `Module`
- **AND** 不得把目标关联写入责任转交给 `Module Registry`

#### Scenario: 关系写入后的读取语义

- **WHEN** `DecisionLinkWrite` 提交成功
- **THEN** 当前阶段必须以“详情页重新读取当前已关联目标结果”为默认后端语义
- **AND** 不得额外设计一套脱离 `DecisionDetailRead` 的回流读取路径

### Requirement: Module Registry 连接边界必须冻结为候选读取与目标校验

系统 SHALL 将 `Decision Center` 与 `Module Registry` 的后端连接边界冻结为最小候选读取与目标校验，不承接 `Module` 主线写入。

#### Scenario: Module 最小连接边界

- **WHEN** `Decision Center` 需要支持 `Decision -> Module` 候选读取与关联写入
- **THEN** 它只允许依赖 `DecisionModuleCandidateRead` 一类的最小候选读取接口
- **AND** 只允许依赖目标存在性、可展示名称与最小状态语义所需的读取/校验
- **AND** 不得在当前阶段调用或吸收 `CreateModule`、`CreateRelease` 或其他 `Module Registry` 主线写入

#### Scenario: Module 连接边界不反向改写 Module Registry

- **WHEN** `Decision Center` 与 `Module Registry` 的服务侧边界被实现
- **THEN** `Module Registry` 仍然只作为被连接方提供最小只读协作前提
- **AND** 不得为了 `Decision Center` 反向重写 `Module Registry` 已冻结的主线接口与职责

#### Scenario: 候选读取接口拥有者与接线位置

- **WHEN** 后续实现 `DecisionModuleCandidateRead` 的接口定义与依赖注入
- **THEN** `ModuleCandidateRead` 的接口与实现必须由 `Decision Center` 的 `candidate/` 子包自己定义和拥有
- **AND** `Module Registry` 不为 `Decision Center` 暴露专门的服务契约或服务方法
- **AND** `candidate/` 子包通过独立 Read 接口隔离，`service/` 层不得直接写跨模块 SQL
- **AND** 具体接线（构造与注入）必须在应用装配点完成，不得在 `service/` 或 `handler/` 内部自行构造
- **AND** 后续若 `Module Registry` 提供更稳定的读模型协作实现，可迁移 `candidate/` 内的具体实现，但接口契约与拥有者保持不变

### Requirement: Product 与 Repository 必须保持后移边界

系统 SHALL 明确当前阶段不为 `Decision Center` 设计 `Product / Repository` 的正式后端连接主线。

#### Scenario: 当前阶段后移对象

- **WHEN** 后续实现讨论 `Decision -> Product` 或 `Decision -> Repository`
- **THEN** 只能保留为未来合同保留位或后续 phase 的连接前提
- **AND** 不得在 `phase03-07` 中设计其独立读接口、写接口或候选读取实现

### Requirement: 后端文件落点必须冻结到实现结构层

系统 SHALL 将 `Decision Center` 的后端设计冻结到“包 + 文件落点”层级，确保后续实现可直接按既定结构创建文件，而不是由实现者临场决定代码组织方式。本 Requirement 冻结职责到文件的映射，但不冻结文件内部的具体 Go HTTP 框架、RPC 框架、ORM 或 SQL Builder 选型。

#### Scenario: 后端模块根包落点

- **WHEN** 后续实现开始创建 `Decision Center` 后端文件
- **THEN** 模块根包必须落在 `backend/internal/decisioncenter/`
- **AND** 模块内部必须按 `handler/`、`service/`、`repository/`、`candidate/` 四个子包组织
- **AND** 不得把 `Decision Center` 的代码散落到 `backend/` 根目录或其他模块目录中

#### Scenario: 读组文件落点

- **WHEN** 后续实现 `DecisionListRead`、`DecisionDetailRead` 与 `DecisionModuleCandidateRead`
- **THEN** 入口层必须落在 `backend/internal/decisioncenter/handler/query_handler.go`
- **AND** 业务编排层必须落在 `backend/internal/decisioncenter/service/query_service.go`
- **AND** `query_service.go` 必须统一承接列表读取、详情读取与候选读取编排
- **AND** 不得把读组拆成多个并列 service 文件，保持读组单文件编排

#### Scenario: 写组文件落点

- **WHEN** 后续实现 `DecisionWrite` 与 `DecisionLinkWrite`
- **THEN** 入口层必须落在 `backend/internal/decisioncenter/handler/command_handler.go`
- **AND** 业务编排层必须落在 `backend/internal/decisioncenter/service/command_service.go`
- **AND** `command_service.go` 必须统一承接创建写入与目标关联写入编排
- **AND** 不得把写组拆成两个独立 service 文件，保持写组单文件编排

#### Scenario: 数据访问层文件落点

- **WHEN** 后续实现 `Decision Center` 的数据访问层
- **THEN** `decisions` 表访问必须落在 `backend/internal/decisioncenter/repository/decision_store.go`
- **AND** `decision_links` 表访问必须落在 `backend/internal/decisioncenter/repository/link_store.go`
- **AND** 数据访问层只负责持久化与读取，不承接业务校验或跨模块编排

#### Scenario: 跨模块候选读取文件落点

- **WHEN** 后续实现 `DecisionModuleCandidateRead`
- **THEN** 候选读取必须落在 `backend/internal/decisioncenter/candidate/module_candidate_read.go`
- **AND** `candidate/` 子包必须通过独立 Read 接口定义隔离，不得在 `service/` 中直接写跨模块 SQL
- **AND** 后续若 `Module Registry` 提供更稳定的读模型协作实现，可迁移具体实现，但接口契约保持不变

#### Scenario: 支撑文件落点

- **WHEN** 后续实现 `Decision Center` 的错误定义、共享类型、输入校验与 HTTP 协议工具
- **THEN** 业务错误哨兵值必须落在 `backend/internal/decisioncenter/errors.go`
- **AND** 跨层共享 API 消息结构（DTO、枚举、请求/响应类型、列表查询参数）必须落在 `backend/internal/decisioncenter/types.go`
- **AND** 输入校验辅助必须落在 `backend/internal/decisioncenter/validate.go`
- **AND** HTTP 协议层共享工具（JSON 编解码、错误到状态码映射）必须落在 `backend/internal/decisioncenter/handler/response.go`
- **AND** 支撑文件的组织方式必须与现有 `moduleregistry` 模块的 `errors.go / types.go / validate.go / handler/response.go` 保持同构
- **AND** 不得把错误定义、共享类型或校验逻辑散落到 `handler/` 或 `service/` 内部

#### Scenario: 当前阶段不冻结的实现工具

- **WHEN** 当前 spec 讨论后端文件落点
- **THEN** 不得提前冻结 Go HTTP 框架、RPC 框架、ORM、SQL Builder、Repository 模板或目录生成器
- **AND** 文件内部的函数签名、中间件选型、数据库访问方式不作为当前阶段冻结内容
- **AND** 只冻结职责分工、接口归属与文件落点，不冻结实现手段

### Requirement: 后端模块内部应保持分层而不提前钉死工具

系统 SHALL 将 `Decision Center` 的后端实现前结构冻结为“入口层（handler）-> 业务编排层（service）-> 数据访问/外部连接层（repository + candidate）”的分层语义，但不提前冻结具体框架与工具。

#### Scenario: 后端分层语义

- **WHEN** 后续实现开始定义 `Decision Center` 后端内部结构
- **THEN** 入口层（`handler/`）只负责承接请求与返回结果
- **AND** 业务编排层（`service/`）只负责动作语义、校验顺序与跨连接口编排
- **AND** 数据访问层（`repository/`）与外部连接层（`candidate/`）只负责持久化与依赖调用

#### Scenario: 当前阶段不得冻结的实现工具

- **WHEN** 当前 spec 讨论后端边界与接口分组
- **THEN** 不得提前冻结 Go HTTP 框架、RPC 框架、ORM、SQL Builder、Repository 模板或目录生成器
- **AND** 只冻结职责分工与接口归属，不冻结实现手段

### Requirement: 合同边界必须与存储模型解耦

系统 SHALL 将当前阶段的接口分组解释为业务合同边界，而不是数据库表结构的直接暴露。

#### Scenario: 合同与存储解耦

- **WHEN** 后续实现开始定义接口请求与响应结构
- **THEN** 不得直接将 `decisions`、`decision_links` 的存储模型原样暴露为外部合同
- **AND** 必须允许后续在 `Contract First` 路线下独立演进接口消息结构

#### Scenario: 合同演进前提

- **WHEN** 后续阶段需要调整接口消息结构
- **THEN** 不得复用已删除字段编号或字段语义
- **AND** 必须保持与 `Protocol Buffers` 长期方向兼容

## MODIFIED Requirements

### Requirement: Decision Detail 后端承接方式

`Decision Detail` 在当前阶段 SHALL 不仅是前端复合详情页，也必须在后端对应为一个统一读取宿主与动作回流宿主。

#### Scenario: Detail 后端收口

- **WHEN** 后续实现讨论 `Decision Detail` 的后端承接方式
- **THEN** `DecisionDetailRead` 必须作为统一详情读取宿主
- **AND** `DecisionLinkWrite` 成功后必须回到 `DecisionDetailRead` 的读取语义
- **AND** 不得为详情页再设计第二套并列回流读取入口

## REMOVED Requirements

### Requirement: 预先冻结 Go 数据访问层具体工具
**Reason**: 当前阶段需要先收敛模块边界、接口分组、分层语义与文件落点，避免实现期漂移；但具体采用何种 ORM、SQL Builder 或 Repository 模板，属于实现层决策，不应在此时写成既成事实。
**Migration**: 后续实现时必须继续遵守本 spec 已冻结的职责分层、接口归属与文件落点；具体工具选型若需偏离项目既有模式，应在实现评审或后续 `audit` 中单独说明。
