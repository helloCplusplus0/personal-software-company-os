# Phase13-08 落实项目治理画像后端主线 Spec

## Why

`phase13-05` 已冻结项目治理画像的后端合同、存储与读写边界，`phase13-07` 已冻结 agent brief 与现有 `ProjectContextService` 的受控演进关系，但如果不继续把后端实现主线本身压成单值规格，执行阶段仍会重新猜“新增哪个 `.proto` service`”“数据落在哪几张表”“现有只读聚合服务怎么演进”“写路径由谁正式承接”。

本次 `/spec` 的目标，是把项目治理画像第一版后端实现主线冻结为：同一 `repository_id` 驱动的 governance profile 持久化与读取主线，保持四实体既有主线不变，保持“手工维护优先”，同时不长出第二套项目事实源或第二套并列读取协议。

## What Changes

- 冻结项目治理画像第一版后端实现的 `.proto / ConnectRPC / Go service / repository / storage` 主线
- 冻结治理画像写路径与读取路径的唯一正式承接位
- 冻结现有 `ProjectContextService` 与治理画像后端实现的关系及演进边界
- 冻结数据库迁移与存储布局的第一版落地约束

## Impact

- Affected specs:
  - `phase13_project_governance_profile_foundation`
  - `phase13-05`
  - `phase13-07`
  - 后续 `phase13-09` 前端治理画像承接
  - 后续 `phase13-10` agent brief 读取主线
- Affected code:
  - `proto/`
  - `backend/`
  - `database/`

## ADDED Requirements

### Requirement: 冻结项目治理画像后端实现的唯一正式主线

系统 SHALL 将项目治理画像第一版后端实现主线冻结为：

1. 同一 `repository_id` 锚定的 repository-scoped governance profile 持久化主线
2. `.proto -> ConnectRPC -> Go application/service -> repository/store -> PostgreSQL` 的单一实现链路
3. 结构化治理字段、canonical 根级文件绑定、全局规范资产绑定的统一读写承接位

补充冻结：

- 第一版不得再并列长出第二个“治理画像后端实现主线”
- 该主线是项目治理层后端实现位，不是第五个业务主实体
- 四实体既有 service / repository / proto 主线保持不变，不因治理画像后端实现而迁移或重写

#### Scenario: 执行者判断治理画像后端实现应挂在哪条主线上

- **WHEN** 执行者设计项目治理画像的后端正式实现主线
- **THEN** 必须回到同一 `repository_id` 驱动的 repository-scoped governance profile 持久化与读取主线
- **AND** 不得把治理画像写读逻辑散落到四实体既有 service 中各自维护

### Requirement: 冻结第一版 `.proto` 与 ConnectRPC 落点

系统 SHALL 将项目治理画像第一版业务合同落点冻结为：

1. 使用现有 `proto/` workspace
2. 沿用 `.proto` 作为唯一长期合同源
3. 沿用 `ConnectRPC` 作为正式业务传输层
4. 治理画像写读 RPC 必须以 `repository_id` 作为唯一正式结构化输入锚点

补充冻结：

- 第一版不得新增并列 proto 根目录、并列 buf workspace 或第二套代码生成链
- 第一版不得新增手写 `chi + JSON HTTP` 业务接口作为治理画像正式合同
- 治理画像读写 message 字段必须直接对齐 `phase13-04 / phase13-05` 已冻结字段矩阵与后端边界
- 若需要响应 envelope、错误语义或枚举，必须在 `.proto` 中正式声明

#### Scenario: 执行者设计治理画像新增 RPC

- **WHEN** 执行者设计治理画像的新增 RPC
- **THEN** 必须先在现有 `proto/` workspace 中定义 `.proto`
- **AND** 必须通过 `ConnectRPC` 暴露
- **AND** 不得绕过 `.proto` 先写临时 JSON handler

### Requirement: 冻结第一版 Go 分层与模块落点

系统 SHALL 将项目治理画像第一版 Go 实现分层冻结为：

1. `connect/`
   - 只承接 Connect transport、请求解包、错误映射与 response 组装
2. `service/`
   - 承接治理画像读写编排、字段分类约束、兼容规则与事务边界
3. `repository/`
   - 承接 PostgreSQL 持久化读写
4. `types.go` / `errors.go`
   - 承接领域结构与业务错误定义

补充冻结：

- 治理画像写入校验、`read-only` 字段排除与资产矩阵分类约束必须收敛在 `service/` 主线，不得散落在多个 transport handler 中
- 数据库 SQL 与扫描逻辑必须收敛在 `repository/`，不得让 `connect/` 直接操作数据库
- 第一版不得把治理画像实现塞进 `projectcontext/candidate` 这类只读 candidate reader 层
- router 装配仍应沿用现有 `backend/internal/platform/router.go` 模式接入

#### Scenario: 执行者安排治理画像后端代码结构

- **WHEN** 执行者安排治理画像后端的目录与职责
- **THEN** 必须按 `connect / service / repository / domain types` 分层
- **AND** 不得让 transport 层或 candidate 层承接正式写入语义

### Requirement: 冻结第一版 PostgreSQL 持久化结构

系统 SHALL 将项目治理画像第一版 PostgreSQL 持久化结构冻结为：

1. 一条按 `repository_id` 唯一锚定的治理画像主记录
2. 一组 `canonical_root_file_bindings`
3. 一组 `global_asset_bindings`

其正式持久化要求如下：

1. 主记录承接 `phase13-04` 已冻结的治理画像结构化字段
2. `canonical_root_file_bindings` 承接 `file_name / role / required`
3. `global_asset_bindings` 承接 `name / kind / entry_ref / role / structured_summary`
4. 顶层目录矩阵 `backend / database / frontend / proto` 不新增 repository-scoped 可写持久化字段

补充冻结：

- 第一版数据库迁移必须显式创建治理画像主记录与两组 bindings 的正式表结构
- `repository_id` 必须作为治理画像主记录唯一业务锚点，并在 bindings 层形成明确外键或等价关联约束
- `structured_summary` 在前 5 份摘要型资产中为必填，在 `README.md / global_skills.md / project_skills.md` 中第一版允许为空
- markdown 正文不得进入任何治理画像表的 canonical 存储列

#### Scenario: 执行者设计治理画像数据库迁移

- **WHEN** 执行者设计治理画像第一版数据库迁移
- **THEN** 必须创建治理画像主记录与两组 bindings 的正式结构
- **AND** 必须让 `repository_id` 成为唯一项目锚点
- **AND** 不得把 markdown 正文设计成持久化列

### Requirement: 冻结第一版写路径主线

系统 SHALL 将项目治理画像第一版写路径主线冻结为：

1. 后端存在一个正式治理画像保存入口
2. 该入口只允许写入 `phase13-04 / phase13-05` 已冻结的可写结构化字段
3. 该入口必须显式排除 `track_type / current_phase_name / current_phase_ref / current_phase_status`
4. 该入口必须在同一事务边界内处理主记录、canonical 根级文件 bindings 与全局规范资产 bindings 的保存

补充冻结：

- 第一版写路径只承接手工维护优先的结构化输入
- 第一版不得在保存时触发目录扫描、模板自动同步、正文拉取入库或自动状态建议
- 第一版不得让多个 RPC、多个 handler 或多个 service 各自维护一套治理画像写语义
- 若出现部分资产 summary 更新失败，应按同一写事务整体失败，而不是写入半套治理画像状态

#### Scenario: 执行者设计治理画像保存逻辑

- **WHEN** 执行者设计治理画像第一版保存逻辑
- **THEN** 必须只允许写入可写结构化字段
- **AND** 必须显式排除 `track_type / current_phase_*`
- **AND** 必须保持单一正式写入承接位与单一事务边界

### Requirement: 冻结第一版读取主线

系统 SHALL 将项目治理画像第一版读取主线冻结为：

1. 后端存在一个正式治理画像结构化读取入口
2. 读取结果必须返回治理画像结构化字段、canonical 根级文件绑定与全局规范资产结构化承接结果
3. 读取结果可以返回“markdown 正文可回源”的能力状态
4. 读取结果不得把 markdown 全文作为 canonical stored field 返回

补充冻结：

- 正文回源属于 read-time resolution，不改变数据库正式存储边界
- 第一版读取失败语义必须区分“结构化读取失败”与“正文回源失败”
- 第一版治理画像读取主线是 `phase13-09` 前端回看与维护入口的同源事实来源
- 第一版读取主线与后续 `phase13-10` agent brief 主线是上下游关系：治理画像读取提供结构化数据源，但不在 `phase13-08` 直接承担完整 brief 组装

#### Scenario: 执行者设计治理画像读取结果

- **WHEN** 执行者设计治理画像第一版读取结果
- **THEN** 必须把结构化字段、结构化摘要与正文回源能力区分开
- **AND** 不得把“可回源正文”误设计为“必须入库的全文字段”
- **AND** 不得在 `phase13-08` 直接把治理画像读取入口扩写成完整 agent brief 入口

### Requirement: 冻结与现有 ProjectContextService 的演进边界

系统 SHALL 将项目治理画像后端主线与现有 `ProjectContextService` 的关系冻结为：**同属 repository-scoped 读取主线体系，但职责分层不同、演进节奏不同，不得并列形成第二套项目事实源。**

其正式约束如下：

1. `ProjectContextService` 继续承接 phase12 既有最小只读项目上下文聚合读取与 Markdown 导出
2. 项目治理画像后端主线承接治理画像结构化写读与资产承接结果持久化
3. 第一版不得为了治理画像写读而重写或替换四实体既有读取主线
4. 第一版不得把治理画像读取接口直接伪装成完整 agent brief 主线
5. 后续 `phase13-10` 若要将 brief 合同收口到 `ProjectContextService` 主线，必须基于 `phase13-07` 已冻结的受控演进规则处理

补充冻结：

- `phase13-08` 可以新增治理画像专属读写 service / RPC 主线，但其职责必须限定为治理画像结构化写读，不得与 `ProjectContextService` 长期并列提供两套“当前项目 PSCO 正式上下文读取”
- 若治理画像读取结果被后续 brief 组装复用，复用方式应是 read-time composition，而不是复制出第二套持久化事实源
- `ProjectContextService` 的 `GetProjectContext / ExportProjectContext` 在 `phase13-08` 不强制立即变更，但不得被重新定义为治理画像正式写读承接位

#### Scenario: 执行者判断现有 ProjectContextService 是否应承担治理画像写读

- **WHEN** 执行者判断现有 `ProjectContextService` 是否直接承担治理画像第一版写读
- **THEN** 应判定其仍以 phase12 既有最小只读聚合职责为主
- **AND** 项目治理画像应沿自己的结构化写读主线落地
- **AND** 不得让两者长期并列为两套项目正式事实源

### Requirement: 冻结第一版非目标与禁止事项

系统 SHALL 将以下事项冻结为 `phase13-08` 的明确非目标：

1. 自动扫描仓库目录并自动填充治理画像
2. 自动同步模板接入状态、模板版本或存在性校验结果
3. markdown 正文全文入库
4. 重写四实体既有 service / repository 主线
5. 在 `phase13-08` 直接实现完整 agent brief 组装与读取主线
6. 新增第二套项目事实源或第二套并列项目上下文读取协议

#### Scenario: 执行者试图在 phase13-08 偷渡更多自动化能力

- **WHEN** 执行者试图在 `phase13-08` 中加入目录自动扫描、模板自动同步、全文入库或完整 brief 主线
- **THEN** 应判定为越出当前阶段边界
- **AND** 必须回到治理画像结构化写读主线的最小实现定位

## MODIFIED Requirements

### Requirement: Project Governance Profile Foundation 的后端实现前提

`phase13_project_governance_profile_foundation` MUST 在 `phase13-08` 中完成项目治理画像的正式后端实现主线，包括 `.proto / ConnectRPC / Go service / repository / PostgreSQL` 的单一承接位、治理画像结构化写读主线、数据库迁移与读取主线；若这些后端实现位仍未落地，则 `phase13-09` 前端承接入口与 `phase13-10` agent brief 读取主线都不得视为可联调状态。

## REMOVED Requirements

### Requirement: 允许 phase13-08 直接以自动扫描或全文入库换取“更快落地”

**Reason**: 这种解释会直接破坏 `phase13-05` 已冻结的“手工维护优先、结构化字段入库、markdown 正文只回源”的边界，并且会让治理画像后端主线退化成目录扫描与全文副本系统，违背“不得产生第二套项目事实源”的阶段目标。

**Migration**: `phase13-08` 只实现治理画像结构化写读主线、数据库正式迁移与受控读取；自动扫描、自动同步、存在性校验和全文能力全部继续留在后续受控进入项。 
