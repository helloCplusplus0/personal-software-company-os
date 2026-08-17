# Phase13-05 产出后端合同、存储与读写边界设计 Spec

## Why

`phase13-03` 已冻结项目级治理层不是第五个业务主实体，`phase13-04` 已冻结第一版字段模型，但后续执行仍然可能在“后端到底由谁承接”“哪些内容入库存结构化字段、哪些只允许正文回源”“新增接口是否再长出第二套 JSON API”这些问题上重新猜测。

本次 `/spec` 的目标，是把项目治理画像的后端承接位、`.proto` 合同边界、存储分层、读写边界，以及 8 份全局规范资产的逐项承接策略冻结为单值规格，供后续实现与验收直接遵循。

## What Changes

- 冻结项目治理画像第一版后端唯一正式承接位与唯一项目锚点
- 冻结 `.proto -> ConnectRPC -> Go application/service` 的正式合同链路
- 冻结结构化字段持久化、markdown 正文回源与全文禁止入库的存储边界
- 冻结全局规范资产逐项承接矩阵与读写权限边界

## Impact

- Affected specs:
  - `phase13_project_governance_profile_foundation`
  - 后续 `phase13-06` 前端承接位与 `phase13-07` agent brief 设计
- Affected code:
  - 无直接源代码改动
  - 直接影响后续实现边界：
    - `proto/`
    - `backend/`
    - `database/`

## ADDED Requirements

### Requirement: 冻结项目治理画像第一版后端唯一正式承接位

系统 SHALL 将项目治理画像第一版后端承接位冻结为：**由同一 `repository_id` 锚定的 repository-scoped governance aggregate**。

该承接位必须满足：

1. 以 `repository_id` 作为唯一正式项目锚点
2. 承接 `phase13-04` 已冻结的 9 类治理画像字段
3. 承接全局规范资产的结构化承接结果
4. 作为后续 Web 与 agent 共享治理事实的唯一后端来源

补充冻结：

- 该承接位是项目级治理层的后端合同位，不是第五个业务主实体
- 该承接位不得被实现为 `Product / Repository / Module / Decision` 任一实体字段的临时拼装结果
- 第一版不得再并列长出第二个“项目治理画像后端入口”或第二个“治理摘要表述协议”

#### Scenario: 执行者判断治理画像后端应挂在哪

- **WHEN** 执行者设计项目治理画像的后端正式承接位
- **THEN** 必须回到同一 `repository_id` 驱动的 repository-scoped governance aggregate
- **AND** 不得把治理画像拆散塞进四实体既有后端合同

### Requirement: 冻结第一版业务合同链路

系统 SHALL 将项目治理画像第一版业务合同链路冻结为：

1. `.proto` 是唯一长期合同源
2. `ConnectRPC` 是默认正式业务传输层
3. `Go application/service` 承接治理画像读写编排
4. `chi + net/http` 只承担基础设施端点或兼容适配，不得形成并列 canonical API

补充冻结：

- 对外字段、枚举、响应 envelope 与错误语义必须以 `.proto` 为准
- 第一版不得新增手写 `chi + JSON HTTP` 业务接口作为项目治理画像的正式合同
- 若后续出现兼容 HTTP handler，其 request / response struct 也必须从 `.proto` 单向派生或显式对齐映射

#### Scenario: 执行者设计治理画像对外接口

- **WHEN** 执行者设计项目治理画像的新增业务接口
- **THEN** 必须先定义 `.proto`
- **AND** 必须默认通过 `ConnectRPC` 暴露
- **AND** 不得额外长出与 `.proto` 并列的第二套 JSON 字段语义

### Requirement: 冻结第一版存储分层

系统 SHALL 将项目治理画像第一版存储分层冻结为以下三层：

1. `governance_profile_record`
   - 按 `repository_id` 持久化 `phase13-04` 已冻结的 9 类治理画像字段
2. `canonical_root_file_bindings[]`
   - 按 `repository_id` 持久化 `canonical_root_files[]`
   - 最小字段保持为 `file_name / role / required`
3. `global_asset_bindings[]`
   - 按 `repository_id` 持久化全局规范资产逐项承接结果
   - 至少承接 `name / kind / entry_ref / role`
   - 若该资产属于“需要结构化摘要”类别，可额外持久化 `structured_summary`

补充冻结：

- `global_asset_bindings[]` 是 `phase13-04` 中 `global_constraint_refs[]` 的后端承接扩展位，用于表达后端逐项承接策略；它不是新增核心字段集合
- 第一版只持久化治理画像结构化字段、资产入口关系、角色与必要摘要
- `docs/` 目录结构通过 `docs_workflow_layout` 进入正式持久化字段
- `backend / database / frontend / proto` 作为当前项目范式 v1 的固定顶层目录集合，第一版后端承接结论冻结为：只作为当前项目范式基线的只读 direct input 保留，不进入 repository-scoped 可写持久化层
- markdown 正文不得作为 canonical 存储副本写入数据库
- 第一版不得为 markdown 正文建立全文索引、全文副本表或并列缓存真相源

#### Scenario: 执行者判断某项数据应持久化还是回源

- **WHEN** 执行者评估某项治理相关数据是否应该直接入库
- **THEN** 只有结构化字段、入口关系、角色和必要摘要允许持久化
- **AND** markdown 正文只能保留回源能力，不得作为 canonical 副本入库

### Requirement: 冻结当前项目范式顶层目录的后端承接结论

系统 SHALL 将当前项目范式 v1 的顶层目录矩阵在第一版后端合同中的承接结论冻结为：

1. `docs/`
   - 通过 `docs_workflow_layout` 进入正式治理画像字段
2. `backend / database / frontend / proto`
   - 不新增专属 repository-scoped 持久化字段
   - 不进入治理画像写路径的可编辑字段集合
   - 只作为当前项目范式 v1 的只读基线输入保留

补充冻结：

- 上述 4 个顶层目录在第一版中不得继续留为“待后续澄清”
- 它们的存在不允许被误解释为需要新增第 10 个治理画像核心字段
- 若后续确需更细粒度目录结构合同，只能作为后续 phase 的显式进入项

#### Scenario: 执行者判断顶层目录该如何进入后端合同

- **WHEN** 执行者判断 `backend / database / frontend / proto` 在第一版后端合同中的承接方式
- **THEN** 必须将其视为当前项目范式 v1 的只读基线输入
- **AND** 不得为其新增 repository-scoped 可写存储字段
- **AND** 不得继续把它们保留为未冻结事项

### Requirement: 冻结第一版读边界

系统 SHALL 将项目治理画像第一版读边界冻结为：

1. 读接口可以返回治理画像结构化字段
2. 读接口可以返回全局规范资产的 `entry_ref / role / structured_summary` 承接结果
3. 读接口可以暴露“正文可回源”的能力或状态
4. 读接口不得把 markdown 全文当作数据库主内容直接返回为 canonical stored field

补充冻结：

- 正文回源属于 read-time resolution，不改变数据库中的正式结构化存储边界
- 第一版允许“结构化摘要 + markdown 正文回源”的双层模式
- 若正文回源失败，失败语义必须独立于治理画像结构化字段读取成功与否

#### Scenario: 执行者设计治理画像读取结果

- **WHEN** 执行者设计治理画像读接口的响应结构
- **THEN** 必须把结构化字段与正文回源能力区分开
- **AND** 不得把“可回源正文”误设计成“必须入库的全文字段”

### Requirement: 冻结第一版写边界

系统 SHALL 将项目治理画像第一版写边界冻结为：

1. 写路径只允许修改治理画像结构化字段
2. 写路径只允许修改全局规范资产的 `entry_ref / role / structured_summary` 等结构化承接结果
3. 写路径不得写入 markdown 全文
4. 写路径不得在保存时顺带自动扫描整个仓库、自动同步模板信息、或自动回填目录全文内容

补充冻结：

- 第一版必须遵守“手工维护优先、自动同步后置”
- `track_type / current_phase_name / current_phase_ref / current_phase_status` 虽属于治理画像结构化字段，但因 `phase13-04` 已冻结为 `read-only`，第一版不得由治理画像维护写路径改写，只允许来自根级正式上游的冻结结果回读
- 若未来进入自动存在性校验或模板接入同步，必须作为后续受控进入项单独设计
- 第一版应收敛为单一写入承接位，不得让多个后端 handler / service 各自维护一套治理画像写语义

#### Scenario: 执行者设计治理画像写入逻辑

- **WHEN** 执行者设计治理画像的保存接口或 application 写路径
- **THEN** 只允许写入结构化治理字段与结构化资产承接结果
- **AND** 必须将 `track_type / current_phase_*` 排除在可写集合之外
- **AND** 不得把仓库 markdown 全文一并存入数据库
- **AND** 不得在保存时偷渡目录扫描或自动同步逻辑

### Requirement: 冻结与四实体既有合同、读路径、写路径的关系

系统 SHALL 将项目治理画像与四实体既有后端主线的关系冻结为：

1. 四实体既有合同继续承接业务事实与业务关系
2. 项目治理画像后端承接位只承接治理画像、全局规范资产与当前阶段状态
3. 四实体读取治理画像时，只能通过同一 `repository_id` 驱动的治理承接位获得项目级治理事实
4. 项目治理画像不得反向复制四实体完整事实为第二套存储副本

补充冻结：

- 第一版允许在读接口中组合返回与 `repository_id` 相关的四实体摘要引用，但不得把这些摘要升级为治理画像自己的 canonical 业务事实
- 第一版不得把 `Decision` 现有“规则”语义误扩写为全局规范资产的正式存储位

#### Scenario: 执行者试图把四实体事实复制进治理画像

- **WHEN** 执行者尝试把 `Product / Repository / Module / Decision` 的业务事实完整复制进治理画像存储
- **THEN** 应判定为复制第二套四实体事实源
- **AND** 必须回收到引用、关联或 read-time aggregation 边界

### Requirement: 冻结全局规范资产逐项承接矩阵

系统 SHALL 将以下全局规范资产的第一版逐项承接策略冻结为单值矩阵：

1. `project_rules.md`
   - 持久化：`entry_ref + role + structured_summary`
   - 允许 markdown 正文回源
   - 第一版禁止全文入库
2. `TECH_STACK_BASELINE.md`
   - 持久化：`entry_ref + role + structured_summary`
   - 允许 markdown 正文回源
   - 第一版禁止全文入库
3. `AGENTS.md`
   - 持久化：`entry_ref + role + structured_summary`
   - 允许 markdown 正文回源
   - 第一版禁止全文入库
4. `architecture_map.md`
   - 持久化：`entry_ref + role + structured_summary`
   - 允许 markdown 正文回源
   - 第一版禁止全文入库
5. `plan.md`
   - 持久化：`entry_ref + role + structured_summary`
   - 允许 markdown 正文回源
   - 第一版禁止全文入库
6. `README.md`
   - 持久化：`entry_ref + role`
   - 允许 markdown 正文回源
   - 第一版禁止全文入库
7. `global_skills.md`
   - 持久化：`entry_ref + role`
   - 允许 markdown 正文回源
   - 第一版禁止全文入库
8. `project_skills.md`
   - 持久化：`entry_ref + role`
   - 允许 markdown 正文回源
   - 第一版禁止全文入库

补充冻结：

- 以上 8 项资产的逐项承接策略不得在实现阶段重新猜测
- `structured_summary` 只出现在前 5 项，不得在 `README.md / global_skills.md / project_skills.md` 第一版中偷偷提升为必做
- “允许 markdown 正文回源”不代表允许全文入库

#### Scenario: 执行者判断某份规范文件需要摘要还是只做引用

- **WHEN** 执行者判断某份全局规范资产需要持久化哪些结构化结果
- **THEN** 必须严格按上述 8 项矩阵执行
- **AND** 不得在实现阶段临时把只需 `entry_ref + role` 的文件提升为摘要必填项

## MODIFIED Requirements

### Requirement: Project Governance Profile Foundation 的后端设计前提

`phase13_project_governance_profile_foundation` MUST 先完成项目治理画像的后端唯一承接位、`.proto` 合同链路、存储分层、读写边界与全局规范资产逐项承接矩阵冻结，再进入 `phase13-06` 前端承接位设计与 `phase13-07` agent brief 设计；若这些后端边界仍未冻结，则后续任何前端入口或 agent 输入设计都不得视为稳定。

## REMOVED Requirements

### Requirement: 允许第一版把 markdown 全文入库作为默认做法

**Reason**: 这种解释会直接把“结构化摘要 + 正文回源”的双层模式打散，导致治理画像后端承接位退化为全文副本仓库，并形成与仓库正文并列的第二套真相源。

**Migration**: 第一版只持久化结构化治理字段、资产入口关系、角色与必要摘要；markdown 正文继续保留回源能力，但不得入库为 canonical 副本。
