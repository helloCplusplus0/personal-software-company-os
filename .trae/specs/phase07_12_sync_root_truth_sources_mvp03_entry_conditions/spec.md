# Phase07-12 完成根级同步与 `mvp0.3` 业务阶段进入条件回写 Spec

## Why

`phase07-08` 到 `phase07-11` 已经完成传输合同产物主线、Go 后端业务传输主线、前端调用主线，以及 `phase01 ~ phase06` 的统一联调、回归与退场验收。现在项目已经具备进入后续 `mvp0.3` 业务阶段的基础，但根级真相源若不及时同步，仓库就会继续保留“phase07 仍在迁移中”与“phase07 已完成基础收口”两套状态。

`phase07-12` 的目标，是把 `phase07` 的前置基础地位、已完成结论、正式证据入口，以及后续进入 `mvp0.3` 业务阶段的条件，统一回写到根级入口文档中，同时严格避免提前猜测尚未建立的业务 phase 名称或路线编号。

## What Changes

- 冻结 `AGENTS.md`、`plan.md`、`docs/README.md`、`architecture_map.md` 的根级同步边界
- 冻结 `phase07` 在根级入口中的收口定位：基础主线迁移已完成、为后续业务阶段提供前置基础
- 冻结 `mvp0.3` 业务阶段的进入条件表达，只允许写条件与上游输入，不允许预设未建立的 phase 名称
- 冻结 `phase07-07` 正式规格、`phase07-08/09/10` 实现结论与 `phase07-11` 验收结论进入根级真相源的同步关系
- 冻结 `docs/README.md` 与 `architecture_map.md` 对 `phase07` 活动文档、正式规格、验收结论的稳定入口
- **BREAKING**：根级文档不得继续保留“phase07 仍待继续迁移/继续切换”的旧表达
- **BREAKING**：根级文档不得凭空写入任何未建立的 `mvp0.3` 业务阶段名称、阶段编号或路线故事

## Impact

- Affected specs:
  - `phase07_transport_contract_mainline_migration`
  - `phase07_07_formal_transport_mainline_cutover_spec`
  - `phase07_08_land_buf_generation_connect_contract_mainline`
  - `phase07_09_cut_go_backend_transport_mainline`
  - `phase07_10_cut_frontend_client_slice_mainline`
  - `phase07_11_validate_phase01_phase06_regression_retirement_acceptance`
- Affected code:
  - 当前无业务代码改动，影响 `AGENTS.md`
  - 影响 `plan.md`
  - 影响 `docs/README.md`
  - 影响 `architecture_map.md`
  - 影响 `docs/phase/phase07_transport_contract_mainline_migration_*`
  - 影响 `.trae/specs/phase07_07_formal_transport_mainline_cutover_spec/`
  - 影响 `.trae/specs/phase07_11_validate_phase01_phase06_regression_retirement_acceptance/acceptance_report.md`

## ADDED Requirements

### Requirement: Phase07 收口事实必须先经过根级同步审核

系统 SHALL 在 `phase07-12` 中先审核 `phase07` 是否已经完成“正式规格冻结 → 生成链主线 → Go 后端主线 → 前端调用主线 → 联调回归与退场验收”的全链路，再允许回写根级状态。

#### Scenario: 审核 phase07 收口事实

- **WHEN** 接手者执行 `phase07-12`
- **THEN** 必须检查 `phase07` 三件套、`phase07-07` 正式规格、`phase07-08 / 09 / 10` 实现结论与 `phase07-11` 正式验收结论之间的入口关系与状态表达
- **AND** 必须确认 `phase07-07` 仍是当前阶段正式规格正文入口
- **AND** 必须确认 `phase07-11` 已给出“可进入 `phase07-12` 根级收口”的明确结论
- **AND** 若任一链路仍停留在“待实现”“待验收”或“待退场”状态，则不得宣告 `phase07` 已收口

### Requirement: 根级真相源必须反映 phase07 的前置基础地位

系统 SHALL 在 `phase07-12` 中把 `phase07` 明确表述为当前项目进入后续 `mvp0.3` 业务阶段前的基础主线建设阶段，而不是继续把它写成当前正在执行的业务功能 phase。

#### Scenario: 回写根级入口

- **WHEN** `phase07` 审核通过并进入根级同步
- **THEN** 必须同步 `AGENTS.md`、`plan.md`、`docs/README.md` 与 `architecture_map.md`
- **AND** 回写内容只允许包含：
  - `phase07` 的完成状态
  - `phase07` 作为前置基础阶段的角色表达
  - `phase07` 的正式规格入口、验收入口与活动文档入口
  - 后续进入 `mvp0.3` 业务阶段的条件
- **AND** 不得把 34 条 RPC 矩阵、浏览器验收细节、问题复测日志等正文原样复制到根级文档

### Requirement: `mvp0.3` 业务阶段进入条件必须明确但不得预设阶段名称

系统 SHALL 在 `phase07-12` 中清楚表达后续进入 `mvp0.3` 业务阶段的条件，但只允许写“条件”和“上游输入”，不允许猜测下一阶段的正式名称、编号或拆分方案。

#### Scenario: 表达 `mvp0.3` 进入条件

- **WHEN** 接手者查询 `phase07` 收口后的下一步
- **THEN** 必须能看到清楚的 `mvp0.3` 业务阶段进入条件
- **AND** 条件必须直接承接 `phase07-07` 正式规格、`phase07-11` 验收结论与当前根级真相源
- **AND** 该表达只允许使用“后续 `mvp0.3` 业务阶段”“待建立正式 phase 入口后切换”或等价受控说法
- **AND** 不得在 `AGENTS.md`、`plan.md`、`docs/README.md`、`architecture_map.md` 中提前写入任何未建立的阶段名称

### Requirement: 根级文档必须保持单值一致且无孤岛

系统 SHALL 在 `phase07-12` 中完成根级真相源与活动文档目录入口的互链复核，确保后续接手者无需依赖对话记忆即可找到 `phase07` 的正式收口证据。

#### Scenario: 复核入口一致性

- **WHEN** 接手者从根级入口或 `docs/README.md` 查找 `phase07` 文档
- **THEN** 必须能找到 `phase07` 三件套、`phase07-07` 正式规格入口、`phase07-11` 验收结论入口与 `phase07` 活动文档入口
- **AND** `architecture_map.md` 必须明确这些入口在收口后的职责与目录落点
- **AND** 不得存在只能通过搜索结果、历史对话或偶然路径发现的活动文档

### Requirement: Plan 状态必须从迁移执行转为收口后角色表达

系统 SHALL 要求 `plan.md` 中关于 `phase07` 的状态、当前主目标与后续目标表达，更新为与真实收口结果一致的单值口径。

#### Scenario: 更新 `plan.md`

- **WHEN** `phase07-12` 完成根级同步
- **THEN** `plan.md` 中 `phase07` 的状态必须从“推进迁移子任务”切换为“基础主线迁移已完成”的收口表达
- **AND** 当前主目标不得继续写成 `phase07-08 ~ 12` 的进行中串行任务
- **AND** 后续目标必须写成 `mvp0.3` 业务阶段的进入条件，而不是预设新的 phase 名称

## MODIFIED Requirements

### Requirement: Phase07 收口后的合法入口

`phase07` 收口完成后，后续接手与 `mvp0.3` 业务阶段 SHALL 同时满足两个入口约束：根级状态以 `AGENTS.md` / `plan.md` / `docs/README.md` / `architecture_map.md` 为准，`phase07` 的执行层边界以 `phase07-07` 正式规格正文与 `phase07-11` 验收结论为准。

#### Scenario: 后续阶段读取入口

- **WHEN** 后续接手者在 `phase07` 收口后继续推进项目
- **THEN** 必须先从根级文档读取项目当前状态、`phase07` 的完成结论与 `mvp0.3` 进入条件
- **AND** 再从 `phase07-07` 与 `phase07-11` 读取当前基础主线的正式边界与验证结论
- **AND** 不得把 `phase07-08 / 09 / 10` 的实现细节或零散复测记录提升为新的长期主入口

## REMOVED Requirements

### Requirement: 保留 phase07 仍处于“继续做迁移子任务”的旧表达

**Reason**: `phase07` 已完成正式规格、主线切换与统一验收，继续保留迁移进行中表述会导致根级真相源与实际阶段状态冲突。

**Migration**: 改为在根级文档中明确 `phase07` 已完成基础主线迁移与统一验收，并以 `phase07-07` 与 `phase07-11` 作为执行层正式入口；后续只表达 `mvp0.3` 进入条件，不预设阶段名称。

### Requirement: 直接在根级文档中猜测 `mvp0.3` 后续阶段名称

**Reason**: 当前只需要为后续业务阶段建立进入条件，不需要也不允许在没有正式入口时自由命名未来阶段。

**Migration**: 改为使用“后续 `mvp0.3` 业务阶段”与“待正式 phase 入口建立后切换”的受控表达。
