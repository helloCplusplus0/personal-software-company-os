# Phase13-04 产出项目治理画像数据模型与字段设计 Spec

## Why

`phase13-02` 已冻结三层信息边界，`phase13-03` 已冻结项目级治理层与四实体主线的关系，但如果不继续把项目治理画像的字段模型压成单值设计，后续执行者仍会在“第一版到底有哪些字段、哪些必须填、哪些只读、哪些先不自动校验”上重新猜测。

本次 `/spec` 的目标，是把项目治理画像的核心字段、嵌套字段矩阵、字段分类与第一版承接边界冻结成可直接进入后端合同、前端承接位与 agent brief 的正式规格。

## What Changes

- 冻结项目治理画像第一版的核心字段集合、字段语义与字段形状
- 冻结 `canonical_root_files / global_constraint_refs / current_phase_*` 的嵌套字段矩阵
- 冻结字段的 `required / optional / read-only / future-auto-verifiable` 分类
- 冻结当前项目范式 v1 如何被结构化承接到字段模型中

## Impact

- Affected specs:
  - `phase13_project_governance_profile_foundation`
  - 后续 `phase13-05 ~ phase13-10` 的后端合同、前端维护入口与 agent brief 设计
- Affected code:
  - 无直接源代码改动
  - 直接影响后续文档与实现边界：
    - `docs/phase/phase13_project_governance_profile_foundation_architecture_plan.md`
    - `docs/phase/phase13_project_governance_profile_foundation_dev_plan.md`
    - `docs/phase/phase13_project_governance_profile_foundation_shared_baseline.md`

## ADDED Requirements

### Requirement: 冻结项目治理画像第一版核心字段集合

系统 SHALL 将项目治理画像第一版核心字段冻结为以下单值集合：

1. `project_profile_version`
2. `track_type`
3. `template_source`
4. `docs_workflow_layout`
5. `canonical_root_files[]`
6. `global_constraint_refs[]`
7. `current_phase_name`
8. `current_phase_ref`
9. `current_phase_status`

补充冻结：

- 第一版字段集合必须足以承接当前项目范式 v1、全局规范资产入口关系与当前阶段入口状态
- 第一版不得临时新增与上述集合并列的第二批“视情况再补”的正式字段
- 第一版不得把目录全文扫描结果、自动建议结果或同步投影字段偷渡进该字段集合

#### Scenario: 执行者判断第一版字段是否已冻结

- **WHEN** 执行者评估项目治理画像第一版需要哪些正式字段
- **THEN** 必须以上述 9 类字段作为正式起点
- **AND** 不得重新临场决定“再补一个同级核心字段试试”

### Requirement: 冻结核心字段的正式语义

系统 SHALL 将第一版核心字段语义冻结如下：

1. `project_profile_version`
   - 用于标识当前项目治理画像遵循的正式版本
   - 字段形状固定为 `string`
   - 第一版必须为非空字符串，且首个正式取值固定为 `project_governance_profile_v1`
   - 第一版必须作为必填治理字段存在
2. `track_type`
   - 用于标识当前项目采用的正式技术路线 / 范式轨道
   - 字段形状固定为受控 `enum`
   - 允许值只包括 `Product Track` 与 `Durable System Track`
   - 第一版必须与当前项目冻结技术路线保持单值一致
3. `template_source`
   - 用于标识当前项目范式或模板来源的手工维护来源信息
   - 字段形状固定为 `string | null`
   - 第一版允许为空；为空时表示当前项目尚未声明显式模板来源
   - 第一版属于治理画像手工字段，不得偷换为自动同步结果
4. `docs_workflow_layout`
   - 用于标识当前项目采用的 docs workflow 结构布局
   - 字段形状固定为受控 `string`
   - 第一版必须为非空字符串，且当前项目的正式取值固定为 `phase/fix/audit/review`
   - 第一版必须能承接 `docs/phase / fix / audit / review` 这类正式布局语义
5. `current_phase_name / current_phase_ref / current_phase_status`
   - 用于承接当前阶段名称、正式入口引用与当前状态
   - `current_phase_name` 字段形状固定为非空 `string`
   - `current_phase_ref` 字段形状固定为非空 `string`，且语义上必须指向正式 phase 入口引用，而不是自由文本备注
   - `current_phase_status` 字段形状固定为受控 `enum`，第一版允许值只包括 `planned / in_progress / completed / blocked`
   - 第一版必须能回到当前正式 phase 主线，而不是散装阶段快照

#### Scenario: 执行者解释核心字段含义

- **WHEN** 执行者需要解释某个核心字段为什么存在
- **THEN** 必须能将其解释为“治理画像的正式合同字段”
- **AND** 必须能回答该字段的形状、是否可空、以及是否受控取值
- **AND** 不得将 `template_source` 解释为自动同步结果
- **AND** 不得将 `current_phase_*` 解释为随手记录的阶段笔记

### Requirement: 冻结嵌套字段矩阵与最小子字段集合

系统 SHALL 将以下嵌套字段矩阵冻结为第一版最小子字段集合：

1. `canonical_root_files[]`
   - `file_name`
   - `role`
   - `required`
2. `global_constraint_refs[]`
   - `name`
   - `kind`
   - `entry_ref`

补充冻结：

- 第一版 `canonical_root_files[]` 必须足以承接当前项目冻结的根级 canonical 文件集合
- 第一版 `global_constraint_refs[]` 必须足以承接全局规范资产的入口引用与分类关系
- 第一版不得在未冻结语义前，额外补入“全文内容”“自动摘要结果”“最近修改时间”等并列子字段

#### Scenario: 执行者设计嵌套字段

- **WHEN** 执行者设计 `canonical_root_files[]` 或 `global_constraint_refs[]`
- **THEN** 必须先满足上述最小子字段集合
- **AND** 若需要更多子字段，必须在后续 phase 明确进入，而不是在第一版临时扩写

### Requirement: 冻结字段分类矩阵

系统 SHALL 将第一版字段分类冻结为以下规则：

1. `required`
   - `project_profile_version`
   - `track_type`
   - `docs_workflow_layout`
   - `canonical_root_files[]`
   - `global_constraint_refs[]`
   - `current_phase_name`
   - `current_phase_ref`
   - `current_phase_status`
2. `optional`
   - `template_source`
3. `read-only`
   - `track_type`
   - `current_phase_name`
   - `current_phase_ref`
   - `current_phase_status`
   - 上述字段在第一版只允许来自根级正式上游的冻结结果回读，不允许在治理画像维护入口中被自由改写
4. `future-auto-verifiable`
   - `canonical_root_files[].required` 的存在性校验
   - `global_constraint_refs[].entry_ref` 的存在性校验
   - `current_phase_ref` 的存在性校验
   - 这些都只属于后续受控进入项，不在第一版字段设计中直接实现为自动同步

补充冻结：

- 字段分类是可叠加维度；同一字段允许同时属于 `required` 与 `read-only`
- 第一版优先解决“字段合同是否成立”，不是先解决“字段能否被自动推断”
- 第一版不得因为未来可自动校验，就把字段从手工维护合同偷换成自动派生结果

#### Scenario: 执行者判断字段是必填还是后置增强

- **WHEN** 执行者判断某字段应该属于必填、可选还是后续自动校验
- **THEN** 必须遵守上述分类矩阵
- **AND** 不得把未来自动校验能力误读成第一版必做实现

### Requirement: 冻结当前项目范式 v1 的结构化承接方式

系统 SHALL 要求第一版字段模型能够结构化承接当前项目范式 v1，至少包括：

1. 当前目录结构：
   - `backend/`
   - `database/`
   - `frontend/`
   - `docs/`
   - `proto/`
2. 当前根级 canonical 文件集合：
   - `.env`
   - `AGENTS.md`
   - `architecture_map.md`
   - `plan.md`
   - `global_skills.md`
   - `project_skills.md`
   - `project_rules.md`
   - `README.md`
   - `TECH_STACK_BASELINE.md`

补充冻结：

- 当前项目范式 v1 进入字段模型时，应表现为“被结构化承接的治理事实”
- 不得被误写成“要求所有未来项目复制当前目录结构”的强制模板规则
- 不得为了第一版字段模型，把当前项目范式 v1 扩写成模板仓库自动同步方案

#### Scenario: 执行者验证字段模型是否足以承接当前项目范式

- **WHEN** 执行者需要验证字段模型是否足以承接当前项目范式 v1
- **THEN** 必须能回答目录结构、canonical 根级文件集合与 docs workflow 布局分别映射到哪些正式字段
- **AND** 不得再需要靠脑内补充“还有哪些字段应该存在”

## MODIFIED Requirements

### Requirement: Project Governance Profile Foundation 的字段设计前提

`phase13_project_governance_profile_foundation` MUST 先完成项目治理画像核心字段集合、嵌套字段矩阵与字段分类的冻结，再进入 `phase13-05` 后端合同、`phase13-06` 前端承接位与 `phase13-07` agent brief 的具体设计；若字段模型仍未冻结，则后续任何合同或入口设计都不得视为稳定。

## REMOVED Requirements

### Requirement: 允许执行者在实现设计阶段继续临场补字段

**Reason**: 这种解释会让 `phase13-05 ~ phase13-07` 再次回到“第一版到底先做哪些字段”的反复讨论，破坏 `phase13-04` 作为字段模型冻结点的职责。

**Migration**: 将第一版核心字段、嵌套字段矩阵与字段分类统一冻结在本 spec；后续若需新增字段，只能作为下一阶段或后续子任务的显式进入项讨论。
