# Phase13-02 冻结 PSCO-native facts / IDE-accessible context / controlled synced projection 三层边界 Spec

## Why

`phase13` 已明确 PSCO 下一步优先承接“项目级治理画像 + 全局规范资产 + agent 项目简报输入”，但如果不先把三层信息边界冻结清楚，后续执行仍会把 IDE 已可读取的目录上下文、Git 推进信息或自动同步候选误带入本阶段，重新把 PSCO 拉回目录扫描器或更重自动化平台。

本次 `/spec` 的目标，是把 `PSCO-native facts / IDE-accessible context / controlled synced projection` 三层模型压成可机械执行的正式规格，让后续执行者能稳定判断“什么该先做，什么后做，什么本阶段明确不做”。

## What Changes

- 冻结三层信息模型的正式定义、承接对象与执行顺序
- 冻结哪些信息属于 `phase13` 当前正式实现范围，哪些继续留给 IDE / agent 现场读取
- 冻结哪些能力只作为后续受控进入项，不得在 `phase13` 起点直接实现
- 冻结后续 `/spec`、实现与验收使用同一套三层边界口径

## Impact

- Affected specs:
  - `phase13_project_governance_profile_foundation`
  - 后续 `phase13-03 ~ phase13-12` 子任务的边界判断
- Affected code:
  - 无直接源代码改动
  - 直接影响后续文档与实现边界：
    - `docs/phase/phase13_project_governance_profile_foundation_architecture_plan.md`
    - `docs/phase/phase13_project_governance_profile_foundation_dev_plan.md`
    - `docs/phase/phase13_project_governance_profile_foundation_shared_baseline.md`

## ADDED Requirements

### Requirement: 冻结三层信息模型的正式定义

系统 SHALL 将 `phase13` 的三层信息模型冻结为以下单值定义：

1. `PSCO-native facts`
2. `IDE-accessible context`
3. `Controlled synced projection`

且后续所有 `phase13` 子任务都必须以该三层模型作为正式边界前提，不得再自由重命名、重新拆层或新增并列第四层。

#### Scenario: 执行者判断某类信息归属

- **WHEN** 执行者评估某类信息应由 PSCO 正式管理、继续留给 IDE / agent 读取，还是下沉为后续能力
- **THEN** 必须先将该信息归入上述三层之一
- **AND** 不得跳过三层判断直接进入实现设计或源码落地

### Requirement: 冻结 PSCO-native facts 的正式承接范围

系统 SHALL 将以下内容冻结为 `PSCO-native facts`，并作为 `phase13` 当前正式承接对象：

1. 四实体信息与关系：`Product / Repository / Module / Decision`
2. 全局规范资产：长期有效的 workflow、技术基线、协作约束、目录与文档职责
3. 项目级治理画像：项目范式版本、`docs workflow` 布局、canonical 根级文件集合、当前阶段入口与状态

补充冻结：

- 同一概念不得跨 `PSCO-native facts / IDE-accessible context / Controlled synced projection` 多层重复承接
- `template_source` 作为第一版手工维护的治理画像字段，属于 `PSCO-native facts`
- 模板仓库接入状态与模板版本自动同步属于 `Controlled synced projection`，不得与 `template_source` 混写为同一层语义

#### Scenario: 判断某信息是否属于本阶段优先实现对象

- **WHEN** 执行者评估某信息是否应进入 `phase13` 当前正式实现范围
- **THEN** 若该信息属于四实体关系、全局规范资产或项目级治理画像
- **THEN** 应将其视为 `PSCO-native facts`
- **AND** 应优先进入 `phase13` 的 `/spec`、实现与验收主线

### Requirement: 冻结 IDE-accessible context 的保留边界

系统 SHALL 将以下内容冻结为 `IDE-accessible context`，并明确其默认继续由 IDE / agent 现场读取，不默认上升为 PSCO 正式事实：

1. 当前工作区文件全文
2. 临时目录漂移
3. 局部实现细节
4. agent 在 IDE 现场即时读取的项目文件内容

#### Scenario: 执行者尝试把目录即时上下文纳入 PSCO

- **WHEN** 执行者提出把工作区文件全文、目录即时状态或局部实现细节直接纳入 `phase13` 的 PSCO 正式事实层
- **THEN** 本 spec 必须将其判定为越出 `PSCO-native facts` 边界
- **AND** 默认要求该信息继续由 IDE / agent 现场能力消费，而不是由 PSCO 接管

### Requirement: 冻结 Controlled synced projection 的后置进入条件

系统 SHALL 将以下内容冻结为 `Controlled synced projection`，并明确它们只作为后续受控进入项，不在 `phase13` 起点直接实现：

1. Git 推进摘要
2. 模板仓库接入状态与模板版本自动同步
3. 自动存在性校验
4. 自动状态建议

#### Scenario: 执行者尝试把同步增强直接纳入 phase13 起点

- **WHEN** 执行者提出在 `phase13` 起点直接实现 Git 推进跟踪、模板自动接入、自动存在性校验或自动状态建议
- **THEN** 本 spec 必须将其判定为 `Controlled synced projection`
- **AND** 要求将其下沉为 `phase13` 正式收口后的下一阶段进入条件

### Requirement: 冻结三层边界的执行顺序

系统 SHALL 将 `phase13` 的执行顺序冻结为：

1. 先正式承接 `PSCO-native facts`
2. 保持 `IDE-accessible context` 继续由 IDE / agent 现场读取
3. 将 `Controlled synced projection` 下沉为后续进入条件

#### Scenario: 判断什么该先做、什么后做

- **WHEN** 执行者需要决定 `phase13` 的先后顺序
- **THEN** 必须先做 `PSCO-native facts`
- **AND** 不得为了图省事先做 Git 同步、模板接入或自动推断
- **AND** 必须能机械回答“先承接正式事实，再保留现场读取，最后才讨论受控同步投影”

### Requirement: 冻结后续规格与实现的统一口径

系统 SHALL 要求后续 `phase13-03 ~ phase13-12` 的 `/spec`、实现与验收共享同一套三层边界口径，不得：

1. 把 `IDE-accessible context` 偷换成 `PSCO-native facts`
2. 把 `Controlled synced projection` 偷渡为本阶段默认实现内容
3. 让 Web 与 agent 分别长出两套不同的信息分层解释
4. 让同一概念在 A / B / C 多层同时承接

#### Scenario: 复核后续子任务是否遵守三层边界

- **WHEN** 执行者复核后续子任务设计或实现
- **THEN** 必须能用同一套三层模型解释其信息来源与边界
- **AND** 若出现跨层偷换或双重解释，则应判定为不通过

## MODIFIED Requirements

### Requirement: Project Governance Profile Foundation 的阶段推进顺序

`phase13_project_governance_profile_foundation` MUST 先完成三层边界冻结，再进入字段模型、后端承接位、前端承接位与 agent brief 设计；若三层边界仍未冻结，则后续任何实现设计都不得视为稳定。

## REMOVED Requirements

### Requirement: 让 phase13 起点同时承接正式事实与自动同步增强

**Reason**: 这种解释会模糊本阶段“先承接正式事实、后讨论受控同步”的顺序，重新把 PSCO 拉回目录扫描器或更重自动化平台。

**Migration**: 将 Git 推进摘要、模板仓库接入状态与模板版本自动同步、自动存在性校验与自动状态建议统一保留为 `Controlled synced projection`，只在 `phase13` 正式收口后，再作为下一阶段进入条件讨论。
