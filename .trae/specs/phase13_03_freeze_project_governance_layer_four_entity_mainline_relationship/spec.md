# Phase13-03 冻结项目级治理层与四实体主线关系 Spec

## Why

`phase13-01` 已冻结本阶段主交付与非目标，`phase13-02` 已冻结三层信息边界，但“项目级治理层到底如何与四实体并存”仍可能被后续执行者误读为：把治理信息塞进 `Decision`、把治理层做成 `Repository` 字段扩张，或再长出第五个业务主实体。

本次 `/spec` 的目标，是把项目级治理层与 `Product / Repository / Module / Decision` 的职责、锚点、依赖方向与禁止越界压成单值规格，让后续字段设计、后端承接位、前端承接位与 agent brief 都建立在同一关系模型上。

## What Changes

- 冻结项目级治理层的正式身份、职责范围与非实体属性
- 冻结它与 `Product / Repository / Module / Decision` 的关系矩阵与禁止越界规则
- 冻结第一版项目治理层以 `repository_id` 为唯一项目锚点、以 `Repository detail` 为前端正式承接位的关系解释
- 冻结后续 `phase13-04 ~ phase13-10` 必须遵守的承接顺序与边界口径

## Impact

- Affected specs:
  - `phase13_project_governance_profile_foundation`
  - 后续 `phase13-04 ~ phase13-10` 的字段、合同、前端承接位与 agent brief 设计
- Affected code:
  - 无直接源代码改动
  - 直接影响后续文档与实现边界：
    - `docs/phase/phase13_project_governance_profile_foundation_architecture_plan.md`
    - `docs/phase/phase13_project_governance_profile_foundation_dev_plan.md`
    - `docs/phase/phase13_project_governance_profile_foundation_shared_baseline.md`

## ADDED Requirements

### Requirement: 冻结项目级治理层的正式身份

系统 SHALL 将项目级治理层冻结为四实体之外的项目级承接层 / 治理层 / 合同层，而不是新的业务主实体。

其正式职责仅包括：

1. 承接项目治理画像
2. 承接全局规范资产的结构化摘要、职责与入口关系
3. 承接当前阶段入口与状态
4. 为 Web 与 agent 提供同一套项目级治理事实

其正式职责不包括：

1. 取代 `Product / Repository / Module / Decision` 的业务含义
2. 充当新的资产登记主线
3. 充当第五个业务 CRUD 主实体

#### Scenario: 执行者判断项目级治理层是否属于新实体

- **WHEN** 执行者尝试把项目级治理层设计为与四实体并列的新业务对象
- **THEN** 本 spec 必须将其判定为越出当前阶段边界
- **AND** 必须回到“项目级承接层 / 治理层 / 合同层”的正式定位

### Requirement: 冻结项目级治理层与四实体的关系矩阵

系统 SHALL 将项目级治理层与四实体的关系冻结为以下单值矩阵：

1. 与 `Repository` 的关系：
   - `Repository` 是当前项目的代码仓库身份对象与项目锚点
   - 第一版项目级治理层以同一 `repository_id` 作为唯一正式锚点
   - 第一版前端正式承接位冻结为 `Repository detail`
   - 但项目级治理层不得被偷换成 `Repository` 业务字段扩张
2. 与 `Product` 的关系：
   - `Product` 继续承接经营目标与交付容器语义
   - 项目级治理层可为同一项目提供治理约束与规范背景
   - 但不得把治理画像字段直接塞进 `Product` 业务事实
3. 与 `Module` 的关系：
   - `Module` 继续承接可复用能力资产语义
   - 项目级治理层只为其提供项目级治理背景，不承接模块业务状态本身
4. 与 `Decision` 的关系：
   - `Decision` 继续承接规则、约束、选择与依据的索引对象语义
   - 项目级治理层不得把全局规范资产、目录结构或当前阶段信息硬塞进 `Decision`
   - `Decision` 可引用或受治理规则约束，但不承担项目治理层的正式承接职责

#### Scenario: 执行者判断某类项目信息应该挂在哪一层

- **WHEN** 执行者评估目录结构、canonical 根级文件、全局规范资产或当前阶段信息的正式承接位置
- **THEN** 应将其归入项目级治理层
- **AND** 不得因为 `Decision` 具备“规则”语义就把这些信息并入 `Decision`
- **AND** 不得因为 `Repository` 是项目锚点就把治理层偷换为 `Repository` 业务字段扩张

### Requirement: 冻结项目级治理层的唯一项目锚点与消费锚点

系统 SHALL 冻结：

1. `repository_id` 是第一版项目级治理层的唯一正式项目锚点
2. `Repository detail` 是第一版前端正式承接位
3. Web 与 agent 对项目级治理层的读取，必须都回到同一 `repository_id` 驱动的正式事实

补充冻结：

- “以 `Repository` 为锚点”只表示项目定位与读取对齐方式，不表示治理层变成 `Repository` 的业务子类型
- 第一版不得新增独立项目治理画像一级页面或并列第二入口
- 若未来需要脱离 `Repository detail` 提升为独立入口，只能在 `phase13` 正式收口后作为下一阶段进入条件讨论

#### Scenario: 执行者设计前端或 agent 的治理层入口

- **WHEN** 执行者设计项目级治理层的读取入口
- **THEN** 必须以同一 `repository_id` 作为正式锚点
- **AND** 第一版人类侧正式入口必须保持在 `Repository detail`
- **AND** 不得让 Web 与 agent 各自长出不同的治理层主入口协议

### Requirement: 冻结项目级治理层与四实体事实的承接边界

系统 SHALL 将以下边界冻结为正式规则：

1. 四实体继续承接业务事实、业务关系与业务动作语义
2. 项目级治理层继续承接项目治理画像、全局规范资产与当前阶段入口状态
3. 同一概念不得同时在项目级治理层与四实体业务事实中承接为两套正式语义
4. 项目级治理层可以为四实体提供治理背景，但不得回写为四实体主语义的一部分

#### Scenario: 执行者尝试把同一事实写成两套正式语义

- **WHEN** 执行者同时在项目级治理层与四实体业务事实中定义同一治理概念
- **THEN** 本 spec 必须将其判定为第二套并列事实源或跨层重复承接
- **AND** 要求回收为单一正式承接位

### Requirement: 冻结后续子任务必须遵守的关系前提

系统 SHALL 要求后续 `phase13-04 ~ phase13-10` 的字段设计、后端合同、前端承接位与 agent brief 设计，都必须建立在以下前提上：

1. 先承认项目级治理层不是第五个业务主实体
2. 再承认它与四实体是“项目级治理背景 -> 业务事实主线”的关系，而不是替代关系
3. 再以 `repository_id` 作为唯一项目锚点展开字段、合同与入口设计

#### Scenario: 复核后续设计是否遵守 phase13-03

- **WHEN** 执行者复核 `phase13-04 ~ phase13-10` 的设计或实现
- **THEN** 必须能回答项目级治理层的正式职责、四实体各自继续承接什么、以及为什么第一版锚点是 `repository_id`
- **AND** 若仍需临场判断“治理层到底算不算新实体”或“治理信息到底挂到哪个实体”，则应判定为不通过

## MODIFIED Requirements

### Requirement: Project Governance Profile Foundation 的字段与入口设计前提

`phase13_project_governance_profile_foundation` MUST 先完成“项目级治理层不是第五个业务主实体、且以 `repository_id` 为唯一项目锚点”的关系冻结，再进入 `phase13-04` 字段设计、`phase13-05` 后端承接位设计、`phase13-06` 前端承接位设计与 `phase13-07` agent brief 设计；若该关系仍未冻结，则后续任何字段或入口设计都不得视为稳定。

## REMOVED Requirements

### Requirement: 允许以四实体中任一实体兼容承接项目治理信息

**Reason**: 这种解释会让后续执行者继续把目录结构、全局规范资产或当前阶段信息零散塞进 `Decision`、`Repository` 或其他业务实体，破坏项目级治理层的单值边界。

**Migration**: 将项目治理画像、全局规范资产与当前阶段入口状态统一回收到项目级治理层；四实体只保留各自业务事实与业务关系语义。
