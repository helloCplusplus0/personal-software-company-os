# Phase07-08 落实 buf 生成链与 ConnectRPC 正式合同产物主线 Spec

> **执行产出**：代码实现已完成。`buf.gen.yaml` 从 2 插件升级为 3 插件正式矩阵（`protocolbuffers/go + connectrpc/gosimple + bufbuild/es`，含 `simple` 选项）；`Makefile` clean target 扩展为三目录清理；`frontend/package.json` 新增 `@connectrpc/connect@^2.1.2` + `@connectrpc/connect-web@^2.1.2`；Go 后端新增 `connectrpc.com/connect@v1.20.0`；`proto/README.md` 更新为 3 插件矩阵表并回写 compat 候选路由仍属阶段性保留。验证结果：10 个 Go protobuf + 9 个 Go Connect（simple 模式）+ 10 个 TS 产物完整；`go build ./...` / `npx tsc -b --noEmit` / `make build` / `make lint` 全部通过。
> **执行日期**：2026-08-11
> **状态**：✅ 已完成，tasks.md 与 checklist.md 全部勾选

## Why

`phase07-07` 已经把 `buf` 三插件矩阵、Go Connect 产物落点、前端 Connect client 运行时依赖、`proto/Makefile` 单一入口，以及本地/验收/CI 等价调用口径冻结为正式规格。但当前仓库真实状态仍停留在“2 插件 + 无 Connect-Go 产物 + 前端缺少 Connect runtime 依赖”的过渡阶段，后续 `phase07` 实现如果继续在这个基础上推进，就会再次出现“合同已冻结、生成链未落地、实现阶段手工补链”的断层。

因此，`phase07-08` 的目标不是重新设计一版工具链，而是把 formal spec 已冻结的生成链、脚本入口、产物目录与调用口径真正落实为仓库中的正式合同产物主线，为后续 Connect handler / frontend client 实现提供稳定上游。

## What Changes

- 将 `proto/buf.gen.yaml` 从当前 2 插件扩展为正式的 3 插件矩阵：`protocolbuffers/go + connectrpc/gosimple + bufbuild/es`
- 冻结 `proto/Makefile`、本地生成入口、验收入口与当前 CI 缺口的单值调用口径
- 冻结 Go Connect 生成产物落点 `backend/internal/gen/connect/**` 与 TypeScript 合同产物落点 `frontend/src/gen/proto/**`
- 冻结前端 `@connectrpc/connect`、`@connectrpc/connect-web` 的运行时依赖补齐要求
- 冻结 `.gitignore`、启动链、验收脚本与文档对新生成链的承接边界
- 明确本阶段只落实生成链与正式合同产物，不在本 spec 内并行展开后端 handler / 前端 owner 实现
- **BREAKING**：后续 `phase07` 的实现、验收与收口，不得再以“手写 DTO / 手工 client / 未生成 Connect 产物”作为默认前提；正式合同产物必须由统一生成链提供

## Impact

- Affected specs:
  - `phase07_02_freeze_chi_connectrpc_buf_formal_composition`
  - `phase07_06_design_business_api_migration_matrix_regression_acceptance`
  - `phase07_07_formal_transport_mainline_cutover_spec`
- Affected code:
  - `proto/buf.gen.yaml`
  - `proto/Makefile`
  - `proto/README.md`
  - `frontend/package.json`
  - `frontend/package-lock.json` 或等价 lockfile
  - `backend/internal/gen/connect/`
  - `backend/internal/gen/proto/`
  - `frontend/src/gen/proto/`
  - `.gitignore`
  - `scripts/`、`database/scripts/`、验收脚本与启动文档中所有引用生成链的入口

## ADDED Requirements

### Requirement: buf 生成链必须切换到 3 插件正式矩阵

系统 SHALL 将 `phase07` 的正式代码生成链切换到已冻结的 3 插件矩阵，而不是继续维持“只生成 Go protobuf + TS descriptor、不生成 Go Connect”的过渡状态。

#### Scenario: `buf.gen.yaml` 目标形态落地

- **WHEN** 团队执行 `phase07-08`
- **THEN** `proto/buf.gen.yaml` 必须同时包含：
  - `buf.build/protocolbuffers/go`
  - `buf.build/connectrpc/gosimple`
  - `buf.build/bufbuild/es`
- **AND** Go Connect 插件必须输出到 `../backend/internal/gen/connect`
- **AND** 必须显式保持 `simple` 模式
- **AND** 不得新增第二个 `buf.gen.yaml`、第二个 proto workspace、第二个 Go 生成根或第二个 TS 生成根

### Requirement: 生成入口必须继续复用 `proto/Makefile`

系统 SHALL 继续复用 `proto/Makefile` 作为唯一受控生成入口，并确保 `build / gen / lint / breaking / clean` 在新生成链下仍能稳定运行。

#### Scenario: `make gen` 与 `buf generate` 行为一致

- **WHEN** 实现者在 `proto/` 目录执行 `make gen` 或 `buf generate`
- **THEN** 两者都必须生成同一套 Go protobuf、Go Connect 与 TS 合同产物
- **AND** `make build`、`make lint`、`make breaking` 的调用路径与基准规则必须保持不变
- **AND** `make clean` 必须同时清理 `backend/internal/gen/proto`、`backend/internal/gen/connect` 与 `frontend/src/gen/proto`

### Requirement: Go Connect 合同产物必须进入正式目录主线

系统 SHALL 要求 9 个业务模块的 Connect-Go 产物进入 `backend/internal/gen/connect/**` 正式目录主线，并保持与现有 `backend/internal/gen/proto/**` 的同构关系。

#### Scenario: Go Connect 产物目录完整

- **WHEN** 团队执行 `buf generate`
- **THEN** 必须在 `backend/internal/gen/connect/psco/**` 下生成 9 个业务模块的 `*.connect.go`
- **AND** 目录层级必须对齐对应 proto 包版本
- **AND** 不得把 Connect 产物混写回 `backend/internal/gen/proto/**`

### Requirement: TypeScript 合同产物必须继续单值输出

系统 SHALL 保持 TypeScript 合同产物继续单值输出到 `frontend/src/gen/proto/**`，并由同一生成链直接提供 service descriptor 给前端 Connect client 使用。

#### Scenario: TS 生成目录不分叉

- **WHEN** 团队执行 `buf generate`
- **THEN** `frontend/src/gen/proto/**` 必须继续承接 `*_pb.ts` 与 service descriptor
- **AND** 不得额外新增 `frontend/src/gen/connect/`、`frontend/src/sdk/` 或其他并列生成根
- **AND** 前端后续 `createClient()` 必须以该目录为唯一正式 descriptor 来源

### Requirement: 前端运行时依赖必须补齐到 Connect 主线

系统 SHALL 在生成链正式落地时同步补齐前端 Connect runtime 依赖，使生成产物可以被后续前端实现直接消费。

#### Scenario: `frontend/package.json` 依赖补齐

- **WHEN** 团队落实 `phase07-08`
- **THEN** `frontend/package.json` 必须新增：
  - `@connectrpc/connect`
  - `@connectrpc/connect-web`
- **AND** 必须继续保留 `@bufbuild/protobuf`
- **AND** lockfile 必须与正式依赖保持一致

### Requirement: 本地开发、验收与当前 CI 缺口必须共享同一生成链口径

系统 SHALL 要求本地开发、验收脚本与当前 CI 缺口替代证据都通过同一生成链入口承接，不得长出“本地一套、验收一套、文档再写一套”的并列规则。

#### Scenario: 单值调用口径

- **WHEN** 文档、脚本或验收步骤需要描述生成链调用方式
- **THEN** 必须统一为：
  - `(cd proto && make build && make gen && make lint)`
  - 必要时补充 `(cd backend && go build ./...)`
  - 必要时补充 `(cd frontend && tsc -b --noEmit)` / `(cd frontend && npm run build)`
- **AND** 当前仓库无 `.github/workflows/` 的事实必须显式保留为缺口建账
- **AND** 不得把仓库根目录直接执行 `make build` 或 `go build ./...` 写回正式口径

### Requirement: 启动链、验收脚本与文档必须承接新生成链

系统 SHALL 要求所有引用 proto 产物的启动链、验收脚本与文档同步承接新的生成主线，避免出现“生成链已切换，但脚本/文档仍按旧产物结构假定”的分叉。

#### Scenario: 脚本与文档同步

- **WHEN** 仓库中的启动说明、验收说明、本地 reset 后联调说明或 proto 文档引用生成产物
- **THEN** 必须显式对齐 `backend/internal/gen/connect/**` 与 `frontend/src/gen/proto/**` 的最新落点
- **AND** 必须明确 reset 脚本只负责 DB / fixture 恢复，不直接承担 Connect 生成链调用
- **AND** 不得把生成链责任错误归到 `database/scripts/reset_*.sh`

### Requirement: 本阶段边界必须保持最小化

系统 SHALL 将 `phase07-08` 的范围限制在生成链、正式合同产物与调用入口落地，不在本 spec 内并行展开 Connect handler 业务实现或前端 owner 迁移实现。

#### Scenario: 与后续实现阶段的边界

- **WHEN** 团队执行 `phase07-08`
- **THEN** 当前阶段必须完成生成链落地、产物生成、依赖补齐与工具链入口对齐
- **AND** 当前阶段不得把“9 模块 Connect handler 完整实现”作为本 spec 的完成条件
- **AND** 后续 handler / frontend client 实现应继续按 `phase07-07 formal spec` 的正文承接到下一实现子任务

## MODIFIED Requirements

### Requirement: `phase07` 的下一实现前置条件

后续 `phase07` 实现阶段 SHALL 以“正式生成链已可稳定产出 Go Connect + TS 合同产物”为前置条件，而不是在 handler / client 实现过程中再临时补齐生成链。

#### Scenario: 进入后续实现前

- **WHEN** 团队准备进入 Connect handler 或前端 generated client 的实现
- **THEN** 必须先确认 `phase07-08` 已落地 3 插件生成链、依赖补齐与统一调用入口
- **AND** 不得一边写业务实现、一边临时修改生成链口径

### Requirement: `phase07` 的证据包工具链项

`phase07` 最终证据包中的 `toolchain_migration.csv` 与 `build_verification.txt` SHALL 以 `phase07-08` 落地后的正式生成链为基准记录验证结果。

#### Scenario: 收口证据引用 phase07-08 结果

- **WHEN** `phase07-11` 收口核销工具链项
- **THEN** 必须检查 `buf.gen.yaml`、`proto/Makefile`、前端依赖、生成产物目录与本地等价验证命令是否都对齐 `phase07-08`
- **AND** 不得继续以旧 2 插件口径或缺失 Connect runtime 的状态判定通过

## REMOVED Requirements

### Requirement: `phase07` 继续以 2 插件生成链作为默认基线

**Reason**: `phase07-02` 和 `phase07-07` 已明确冻结 3 插件正式矩阵，继续保留 2 插件默认基线会让正式合同产物主线与实现主线脱节。

**Migration**: 将现有 `proto/buf.gen.yaml` 扩展为 3 插件矩阵，保留 `proto/Makefile` 作为唯一入口，并由同一生成链生成 Go protobuf、Go Connect 与 TS 合同产物。

### Requirement: 前端通过手写类型或手工 client 承接 Connect 正式主线

**Reason**: `phase07` 正式口径要求前端 Connect client 以 `frontend/src/gen/proto/**` descriptor 和统一 transport 为基础，继续依赖手工 client 会形成并列主线。

**Migration**: 在 `phase07-08` 先补齐 `@connectrpc/connect` 与 `@connectrpc/connect-web`，后续前端实现统一从生成产物与共享 transport 承接。
