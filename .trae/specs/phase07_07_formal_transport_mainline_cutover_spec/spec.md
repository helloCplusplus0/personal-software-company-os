# Phase07-07 产出首份传输主线完全切换正式规格文档 Spec

> **执行产出**：`transport_mainline_cutover_spec_v0.1.md` — phase07 传输主线完全切换的**唯一正式规格正文**，将 phase07-01~06 的全部冻结与设计结论收口为单一文档。覆盖 8 个章节：文档定位与上游继承、迁移范围与边界（9 模块 / 34 RPC / 4 legacy / infra keep list）、合同与工具链（.proto 唯一合同源 / 单一 /api / buf 3 插件 / 前端依赖 / CI 缺口）、后端正式主线（chi 职责 / Connect 挂载 / 分层保持 / 错误映射 / compat 过渡组）、前端正式主线（shared transport / slice-local client / query 约束 / 11 mutation owner / 13 项 adapter 资产回收）、验收退场与收口（34 RPC 核销 / 9 跨模块回归 / 4 legacy 退场证据 / 双门槛 / 8 份证据包 / 8 条阻断）、执行顺序、一致性声明（10 项上游对齐）。
> **执行日期**：2026-08-11
> **状态**：✅ 已完成，tasks.md 与 checklist.md 全部勾选

## Why

`phase07-01 ~ 06` 已分别冻结了迁移范围、`chi + ConnectRPC + buf` 正式组合、compat 退场标准、后端 Connect 接线、前端 generated client / query / application 迁移设计，以及 34 条 RPC 的迁移矩阵与回归验收设计。但这些结论仍分散在多个冻结或设计文档中，后续 `phase07-08 ~ 11` 若继续并列消费这些文档，仍会回到“实现前临场拼装、多入口解释、验收与退场标准分散”的状态。

因此，当前必须产出 `phase07` 首份正式规格正文，把 `phase07-01 ~ 06` 的单值结论统一收口为后续实现、回归、退场与收口的唯一直接上游规格来源。

## What Changes

- 冻结 `phase07` 正式规格正文的角色、文件落点与唯一直接上游入口职责
- 冻结该 formal spec 对 `phase07-01 ~ 06` 的继承关系，不允许再形成并列执行入口
- 冻结正式规格正文必须完整覆盖的章节：迁移范围、生成链、后端、前端、验收与退场标准
- 冻结 9 个业务模块、34 条 canonical RPC、4 条 legacy / compat endpoint、11 项前端正式写动作 owner 在 formal spec 中的统一承接要求
- 冻结 `phase07-08 ~ 11` 只允许以 formal spec 为直接上游规格来源推进实现与核销
- **BREAKING**：`phase07-01 ~ 06` 在 formal spec 生效后退为冻结来源与证据链，不再承担并列直接执行层入口职责

## Impact

- Affected specs:
  - `phase07_transport_contract_mainline_migration`
  - `phase07_01_freeze_transport_migration_scope_canonical_boundary`
  - `phase07_02_freeze_chi_connectrpc_buf_formal_composition`
  - `phase07_03_freeze_compat_migration_exit_criteria`
  - `phase07_04_design_go_connect_handler_service_chi_mount`
  - `phase07_05_design_frontend_generated_client_query_application_migration`
  - `phase07_06_design_business_api_migration_matrix_regression_acceptance`
- Affected code:
  - 当前无代码改动；影响后续 `backend/internal/platform/`、`backend/internal/*/connect/`、`frontend/src/shared/rpc/`、`frontend/src/features/*/{data,application}/`、`proto/buf.gen.yaml`、`proto/Makefile`、`database/scripts/reset_*.sh` 与 phase07 收口证据包的实现与验收入口

## ADDED Requirements

### Requirement: `phase07` 正式规格正文必须收口到单一文档

系统 SHALL 将 `phase07` 的传输主线完全切换规格收口到单一正文文档，而不是继续以 `phase07-01 ~ 06` 多份冻结 / 设计文档作为长期并列入口。

#### Scenario: 正式规格正文文件落点

- **WHEN** 后续实现或验收需要引用 `phase07` 的正式规格
- **THEN** 必须存在单一正文文件 `phase07_07_formal_transport_mainline_cutover_spec/transport_mainline_cutover_spec_v0.1.md`
- **AND** 该正文文件必须成为 `phase07-08 ~ 11` 的唯一直接上游规格入口
- **AND** 不得要求接手者先手工拼装 `phase07-01 ~ 06` 后才能理解当前阶段正式口径

### Requirement: `phase07-01 ~ 06` 结论必须统一继承

系统 SHALL 要求 `phase07` 正式规格正文完整继承 `phase07-01 ~ 06` 已冻结的单值结论，而不是重新定义第二套迁移边界、第二套生成链、第二套后端装配、第二套前端承接位或第二套验收门槛。

#### Scenario: formal spec 继承前置冻结

- **WHEN** 正式规格涉及迁移范围、`/api` 基址、buf 生成链、Connect handler 挂载方式、前端 generated client 组合、compat 退场窗口或回归验收矩阵
- **THEN** 必须分别完整承接 `phase07-01 ~ 06` 的已冻结结论
- **AND** 不得绕开这些上游文档扩范围、改口径或回退到旧数据
- **AND** 不得把已经在上游修正过的口径重新写回过时版本

### Requirement: 正式规格正文必须完整覆盖迁移范围

系统 SHALL 在 formal spec 的同一正文中完整覆盖本阶段的迁移范围与当前阶段 Done 边界。

#### Scenario: 迁移范围章节是否完整

- **WHEN** 接手者编写或审阅 `phase07` 正式规格正文
- **THEN** 文档必须明确 9 个 canonical 业务模块、34 条 canonical RPC、4 条 legacy / compat endpoint 与非业务 infra keep list
- **AND** 必须明确 `.proto` 是唯一长期合同源，业务 transport 正式切换到 `ConnectRPC`
- **AND** 不得把 hand-written JSON business handler 继续写成长期主线

### Requirement: 正式规格正文必须完整覆盖生成链与工具链承接位

系统 SHALL 在 formal spec 中完整覆盖 `buf` 生成链、Go / TS 产物目录、前端运行时依赖、Vite `/api` proxy、本地验证入口与当前 CI 缺口。

#### Scenario: 生成链与工具链章节是否完整

- **WHEN** 接手者审阅 `phase07` 正式规格正文的工具链章节
- **THEN** 必须明确 `buf.build/protocolbuffers/go + buf.build/connectrpc/gosimple + buf.build/bufbuild/es` 的正式插件矩阵
- **AND** 必须明确 `proto/Makefile` 与 `buf generate` 的唯一承接位
- **AND** 必须明确前端通过 `@connectrpc/connect + @connectrpc/connect-web` 运行时组合访问 `/api`
- **AND** 必须明确当前仓库没有现成 `.github/workflows/*` 的事实属于显式缺口，而不是假定已有 CI

### Requirement: 正式规格正文必须完整覆盖后端主线切换

系统 SHALL 在 formal spec 中完整覆盖后端从手写 JSON handler 主线切换到 `chi + Connect handler + existing service` 主线的正式形态。

#### Scenario: 后端章节是否完整

- **WHEN** 接手者审阅 formal spec 的后端章节
- **THEN** 必须明确：
  - `chi` 只承担 middleware 外壳、infra keep list 与 `/api` 子路由承接
  - Connect handler 必须消费 generated `path`，正式挂载方式对齐 `phase07-02 / 04`
  - `service` 层继续保持 Query / Command 分层
  - compat 路由只允许作为过渡资产存在，并按 `phase07-03` 最晚时点退场
- **AND** 不得重新长出第二套 canonical API 或第二套 transport 主线

### Requirement: 正式规格正文必须完整覆盖前端主线切换

系统 SHALL 在 formal spec 中完整覆盖前端 generated client、read owner、application owner、route caller 与 adapter 退场规则的正式承接方式。

#### Scenario: 前端章节是否完整

- **WHEN** 接手者审阅 formal spec 的前端章节
- **THEN** 必须明确：
  - `shared/rpc/connect-transport.ts` 是唯一共享 transport 承接位
  - `features/<slice>/data/connect-client.ts` 是 slice-local generated client 承接位
  - query 层纯只读，mutation 收敛到切片内固定 application owner
  - 11 项正式写动作 owner 与 4 组 candidate read / mutation 联合验收必须进入正文
  - route / page / component 不得再直连旧 `api-adapter.ts`
- **AND** 不得遗漏 `Onboarding` route caller、`SovereigntyPanel` 过渡位与 compat adapter 的退场窗口

### Requirement: 正式规格正文必须完整覆盖验收、退场与 phase 收口标准

系统 SHALL 在 formal spec 中完整覆盖跨模块回归、legacy endpoint retirement、前端 owner 收口、工具链核销与最终 evidence package。

#### Scenario: 验收与退场章节是否完整

- **WHEN** 接手者审阅 formal spec 的验收与收口章节
- **THEN** 必须明确：
  - 34 条 canonical RPC 的迁移核销方式
  - 4 条 legacy / compat endpoint 的 endpoint 级删除证据与替代 Connect 回归证据
  - 11 项 mutation owner 的验收映射
  - `/api` 单一基址在 dev / 验收 / 部署链路中的承接要求
  - `phase07-11` 的最终证据包结构与阻断条件
- **AND** 不得把“Connect 主线已接通”或“单接口 smoke test 通过”视为 phase07 收口充分条件

### Requirement: 正式规格正文必须成为 `phase07-08 ~ 11` 的唯一直接上游

系统 SHALL 将 `phase07` 正式规格正文定位为后续实现、验收、退场与收口核销的唯一直接上游规格来源。

#### Scenario: 后续阶段引用 formal spec

- **WHEN** `phase07-08 ~ 11` 需要引用当前阶段正式边界
- **THEN** 必须优先引用 `transport_mainline_cutover_spec_v0.1.md`
- **AND** 不得继续并列引用 `phase07-01 ~ 06` 作为长期直接执行入口
- **AND** 不得绕开正式规格正文直接以对话结论替代规格入口

### Requirement: 正式规格正文必须与根级真相源和当前阶段文档保持单值一致

系统 SHALL 要求 `phase07` 正式规格正文与根级真相源、当前阶段 phase 文档及项目规则保持单值一致。

#### Scenario: 真相源一致性判定

- **WHEN** 正式规格正文引用项目定位、技术基线、目录规则、transport 合同规则或当前阶段完成条件
- **THEN** 必须与 `TECH_STACK_BASELINE.md`、`project_rules.md`、`plan.md`、`architecture_map.md`、`docs/README.md`、`phase07` 的 architecture / baseline / dev plan 保持一致
- **AND** 必须与 `phase07-01 ~ 06` 的修正后结论保持一致
- **AND** 不得在正式规格正文中重写一套冲突的主结论

## MODIFIED Requirements

### Requirement: phase07 的直接执行层规格入口

后续 `phase07-08 ~ 11` 的实现、验收、退场与收口核销 SHALL 以本次产出的 `transport_mainline_cutover_spec_v0.1.md` 作为唯一直接上游规格入口；而该正文自身继续继承 `phase07-01 ~ 06` 的冻结与设计结论。

#### Scenario: 后续阶段引用正式规格

- **WHEN** 后续阶段需要引用 `phase07` 当前阶段正式规格
- **THEN** 必须从 `transport_mainline_cutover_spec_v0.1.md` 进入
- **AND** 该正文必须显式继承 `phase07-01 ~ 06`
- **AND** 不得把 `phase07-01 ~ 06` 长期保留为并列直接执行入口

### Requirement: phase07 的收口判定口径

`phase07` 的收口判定 SHALL 从“分散在 `phase07-01 ~ 06` 的冻结与设计结论”升级为“formal spec 中单值化的收口标准与证据模型”。

#### Scenario: 收口判定入口

- **WHEN** 团队检查 `phase07` 是否具备收口条件
- **THEN** 必须从 formal spec 中读取双门槛与最终 evidence package 口径
- **AND** 不得要求先翻阅多个冻结文档后自行拼装收口标准

## REMOVED Requirements

### Requirement: 后续实现继续并列拼装 `phase07-01 ~ 06`

**Reason**: `phase07-07` 的目标就是把 `phase07-01 ~ 06` 的冻结与设计结论收口为单一正式规格正文，避免 `phase07-08 ~ 11` 再次回到多文档临场拼装。

**Migration**: 后续实现、退场与验收若需要引用 `phase07` 当前阶段正式边界，应统一从 `transport_mainline_cutover_spec_v0.1.md` 进入；`phase07-01 ~ 06` 仅保留为该正文的冻结来源与追踪依据。

### Requirement: 以局部设计文档替代 formal spec 作为后续实现入口

**Reason**: `phase07-04 ~ 06` 解决的是后端设计、前端设计与迁移矩阵设计，不承担后续整个阶段的唯一直接规格入口职责。

**Migration**: 将这些设计文档的关键结论按章节收束进 formal spec，后续实现统一按 formal spec 承接，设计文档作为追溯证据保留。
