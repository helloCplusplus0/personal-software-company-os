# phase14-07 落实 Standard 后端主线 Spec

## Why

phase14-03/04 已冻结 Standard 数据模型、目录树规范、`psco.standard.v1` 合同与 Go 模块分层设计，phase14-06 已冻结 0011 迁移结构与幂等守卫，但均停留在设计文档。本子任务是 phase14 首个源代码实现类子任务：把 Standard 后端主线（proto 合同 → 数据库迁移 → Go 模块 → brief 装配）落地为可运行、可测试的交付物。本子任务为源代码实现类，设计决策一律以上游冻结为准，不在实现中引入新设计。

## What Changes

- 产出本 spec 三件套（`phase14_07_land_standard_backend_mainline`）
- 新增 `proto/psco/standard/v1/standard.proto`（4 枚举 + 4 核心消息 + 8 RPC envelope + `StandardService`，逐字落地 phase14-04 冻结合同）
- 修改 `proto/psco/project_context/v1/project_context.proto`：追加 `import "psco/standard/v1/standard.proto"` 与 `GetProjectBriefResponse.standards = 8`（`repeated psco.standard.v1.Standard`，过渡态新增接入——`global_assets = 3` 与 `governance_profile = 2` 现状不动，phase14-09 才退役切换）；同步更新 brief 顶层字段注释（phase13 的"7 顶层 / 不新增第 8 块"表述随 `standards = 8` 演进更新，避免注释与字段自相矛盾）
- 执行 `make gen`：三端生成产物（Go protobuf / Go Connect / TS）
- 新增 `database/migrations/0011_phase14_standard_entity.sql`：**仅第一段建表**（三表 + 唯一约束 + 索引，`CREATE TABLE IF NOT EXISTS` 幂等形态；数据迁移 DO 块与 drop 两表段由 phase14-09 按 phase14-06 设计追加，dev_plan L53 字面范围"三表 + 枚举 + 唯一约束"）
- 新增 `backend/internal/standard/` 模块（8 文件清单逐字 phase14-04：`connect/server.go` / `service/command_service.go` / `service/query_service.go` / `repository/standard_store.go` / `candidate/target_reader.go` / `errors.go` / `types.go` / `validate.go`）+ 测试资产（`validate_test.go` R1-R8 单测含非法树用例；store/service 集成测试沿袭 governanceprofile 既有 integration test 模式）
- 修改 `backend/internal/projectcontext/`：`candidate/context_readers.go` 追加 `StandardReader` 接口（签名逐字 phase14-04）；`service/query_service.go` 注入 standardReader 并在 `GetProjectBrief` 装配 `standards[]`；`types.go` 的 `ProjectBriefReadResult` 追加 `Standards` 字段；`connect/server.go` brief 响应组装 standards；集成测试追加 `standards[]` round-trip 断言
- 修改 `backend/internal/connecterrors/connect_errors.go`：注册 standard 哨兵映射（`ErrStandardNotFound` / `ErrBindingNotFound` → NotFound；`ErrInvalidInput` → InvalidArgument；`ErrStandardReadFailed` / `ErrStandardSaveFailed` → Internal 段显式留痕）
- 修改 `backend/internal/platform/router.go` + `server.go`：`buildStandard` / `mountStandardConnect` 装配；`buildProjectContext` 签名扩展双 Reader 注入（phase14-06 已预留该动作归属本子任务）

## Impact

- Affected specs:
  - `docs/phase/phase14_standard_entity_foundation_dev_plan.md`（上游，phase14-07 定义 L51-54）
  - `.trae/specs/phase14_03_design_standard_data_model_directory_tree/spec.md`（直接上游：三表 DDL / DirectoryTreeNode jsonb 规范 / R1-R8 形式化）
  - `.trae/specs/phase14_04_design_backend_contract_storage_read_write_boundaries/spec.md`（直接上游：proto 合同逐字源 / 8 RPC 三要素 / 模块分层 / StandardReader 签名 / brief 装配步骤 1）
  - `.trae/specs/phase14_06_design_profile_retirement_data_migration/spec.md`（0011 三段结构上游：本子任务只落地第一段；phase14-09 追加第二、三段）
  - `.trae/specs/phase14_02_freeze_standard_entity_relation_boundaries/spec.md`（八格矩阵 / brief 对照表上游约束）
- Affected code:
  - 新增：`proto/psco/standard/v1/standard.proto`、`database/migrations/0011_phase14_standard_entity.sql`、`backend/internal/standard/`（8 文件 + 测试）、三端生成产物
  - 修改：`proto/psco/project_context/v1/project_context.proto`、`backend/internal/projectcontext/`（candidate / service / types / connect / 集成测试）、`backend/internal/connecterrors/connect_errors.go`、`backend/internal/platform/router.go`、`backend/internal/platform/server.go`
  - 不动：`governanceprofile` 模块、画像两表、`global_assets = 3` / `governance_profile = 2` 字段（phase14-09 退役范围）

## ADDED Requirements

### Requirement: psco.standard.v1 合同必须逐字落地

`proto/psco/standard/v1/standard.proto` SHALL 按 phase14-04 §ADDED-1 冻结内容实现：

- 文件头注释四段（文档定位 / 上游规格 / 合同约束 / 生成入口），风格对齐 `governance_profile.proto`
- `syntax = "proto3"` / `package psco.standard.v1` / `option go_package = "github.com/psco/backend/internal/gen/proto/psco/standard/v1;standardv1"` / `import "google/protobuf/timestamp.proto"`
- 4 枚举逐字：`StandardStatus`（UNSPECIFIED/DRAFT/ACTIVE/RETIRED = 0-3）、`NodeType`（UNSPECIFIED/DIRECTORY/FILE = 0-2）、`BindingTargetType`（UNSPECIFIED/REPOSITORY/PRODUCT/DECISION/MODULE = 0-4）、`BindingRole`（UNSPECIFIED/TEMPLATE_SOURCE/ADOPTS = 0-2）
- 4 核心消息逐字（字段名与字段号）：`DirectoryTreeNode`（name/node_type/role/summary/ref/children=1-6）、`Standard`（id/name/description/status/directory_tree/created_at/updated_at=1-7）、`StandardBinding`（id/standard_id/target_type/target_id/role/note/created_at=1-7）、`StandardRevision`（id/standard_id/change_summary/created_at=1-4）
- 8 RPC 的请求/响应 envelope 与 `service StandardService` 按 phase14-04 §ADDED-2 三要素表逐字实现（错误语义以注释承载，对齐既有合同风格）
- 生成入口：`proto/` 目录 `make gen`；产物落点 Go `backend/internal/gen/proto/psco/standard/v1/`、Go Connect `backend/internal/gen/connect/psco/standard/v1/`、TS `frontend/src/gen/proto/psco/standard/v1/`

#### Scenario: 合同落地判定

- **WHEN** 本子任务完成 proto 编写并执行 `make gen`
- **THEN** `make lint` / `make build` / `make breaking` 三门禁全绿（新增包为向后兼容变更）；生成产物三端齐备

### Requirement: 0011 迁移第一段建表必须落地

`database/migrations/0011_phase14_standard_entity.sql` SHALL 仅承载第一段建表（phase14-06 冻结三段结构的第一段）：

- 三表 DDL 逐字 phase14-03 §ADDED-1 草案：`standards`（name 全局唯一 / status CHECK / directory_tree JSONB NOT NULL）/ `standard_revisions`（CASCADE + `idx_standard_revisions_standard_id`）/ `standard_bindings`（target_type CHECK / role CHECK / 四元组 UNIQUE + 两个 `idx_` 索引）
- 幂等形态：`CREATE TABLE IF NOT EXISTS` + `CREATE INDEX IF NOT EXISTS`（为 phase14-09 追加数据迁移段后的整文件重放预留，phase14-06 §ADDED-1 第一段冻结）
- 头注释声明文件定位与"phase14-09 将追加第二段数据迁移与第三段 drop 两表（phase14-06 设计）"的结构预留
- 本子任务**不写**数据迁移 DO 块、**不 drop** 画像两表（phase14-09 范围）

#### Scenario: 建表落地判定

- **WHEN** 后端启动执行迁移（或验收环境 apply）
- **THEN** 三表建成、schema_migrations 记录 `0011_phase14_standard_entity`；画像两表与 `governance_profiles` 主表原样保留

### Requirement: backend/internal/standard 模块必须按 8 文件分层实现

模块 SHALL 按 phase14-04 §ADDED-3 冻结的 8 文件清单实现（职责单值映射，无多余文件、无职责漂移）：

| 文件 | 实现要点（上游冻结） |
|---|---|
| `errors.go` | 5 哨兵：`ErrStandardNotFound` / `ErrBindingNotFound`（NotFound）、`ErrInvalidInput`（InvalidArgument）、`ErrStandardReadFailed` / `ErrStandardSaveFailed`（Internal）；注释风格对齐 governanceprofile/errors.go |
| `types.go` | `DirectoryTreeNode` Go 结构逐字 phase14-03 §ADDED-2（json tag / omitempty 语义 / Children `[]*DirectoryTreeNode`）；`StandardReadResult` 等读写模型（含 status / node_type / target_type / role 的受控字符串与枚举转换） |
| `validate.go` | R1-R8 逐条实现 phase14-03 §ADDED-3 形式化表（判定逻辑 / 稳定错误码 `INVALID_TREE_ROOT` 等 8 类 / 错误信息含自根起 `/` 连接的节点路径）；R2 检查时机含创建 / 整树替换 / 状态变更（draft→active/retired）；R6 序列化上限 65536 字节 |
| `repository/standard_store.go` | 三表 SQL + jsonb 编解码；`UpdateStandard` 单事务边界（整树替换 + updated_at + revision 追加同一事务，不写半套状态）；不含业务判断 |
| `candidate/target_reader.go` | 4 实体表 EXISTS 存在性校验（repositories / products / decisions / modules），沿袭 governanceprofile/candidate 模式 |
| `service/command_service.go` | Create / Update / Delete / Bind / Unbind 写路径 owner；BindStandard 校验顺序固定 phase14-04 Scenario：standard 存在 → 枚举合法 → 八格矩阵（template_source 仅 repository）→ target 存在 → 唯一约束；CreateStandard 不记 revision（phase14-04 语义要点） |
| `service/query_service.go` | List（updated_at DESC）/ Get（含 bindings）/ ListRevisions（created_at DESC）读路径 owner；`StandardReader` 接口实现位 `ListStandardsByRepository`（经 standard_bindings 任意 role 反查，含全树；无关联返回空列表非错误） |
| `connect/server.go` | 8 RPC Connect handler：proto 解包 → service 调用 → 响应组装；错误统一 `connecterrors.MapToConnectError`；无业务逻辑 |

RPC 错误语义单值对齐 phase14-04 §ADDED-2 三要素表（全部归一 NotFound / InvalidArgument / Internal 三类；重复绑定 → InvalidArgument 且错误信息含 "already bound"）。

#### Scenario: 模块实现判定

- **WHEN** 模块代码完成
- **THEN** 文件清单与上表一一对应；`go build ./...` / `go vet ./...` 零错误

### Requirement: 树校验必须有单元测试覆盖（含非法树用例）

`validate_test.go` SHALL 对 R1-R8 逐条覆盖：

- 每条规则至少 1 个非法用例（命中该规则稳定错误码）+ 1 个合法边界用例
- R1-R8 非法用例合计覆盖全部 8 个错误码；错误信息断言含节点路径定位（如 `/docs/phase`）
- phase14-03 §ADDED-2 序列化等价规则（children nil/[] 等价、omitempty 反序列化缺失视为空）在反序列化用例中覆盖
- 合法完整树样例（含 directory 嵌套 + file 节点 + ref 两种形态）全规则通过

#### Scenario: 校验测试判定

- **WHEN** `go test ./internal/standard/...`
- **THEN** validate 单测全绿；非法树用例逐错误码可追溯

### Requirement: UpdateStandard 事务与 StandardReader 必须有集成测试

集成测试 SHALL 沿袭仓库既有 integration test 模式（真实 PostgreSQL + fixture，环境检测跳过机制与 governanceprofile / projectcontext 既有测试一致）：

- `UpdateStandard`：整树替换 + `updated_at` 变更 + revision 追加在同一事务内完成的 round-trip 断言（更新后 GetStandard 返回新树；`ListStandardRevisions` 多出一条且 `change_summary` 匹配）
- `StandardReader.ListStandardsByRepository`：绑定（含 adopts / template_source 两 role）后反查返回含全树 Standard；无绑定 repository 返回空列表非错误
- 写路径关键错误语义抽查：CreateStandard name 重复 / RETIRED 拒绝 / active 空树 R2 拒绝 / DeleteStandard active 拒绝 / BindStandard 八格矩阵非法（template_source 携非 repository）与 target 不存在

#### Scenario: 集成测试判定

- **WHEN** 验收环境数据库可用并运行集成测试
- **THEN** 上述断言全绿；数据库不可用时测试按既有模式跳过（不 fail）

### Requirement: projectcontext 必须接入 StandardReader 并装配 standards[]

- `candidate/context_readers.go` 追加 `StandardReader` 接口，签名逐字 phase14-04 §ADDED-5（`ListStandardsByRepository(ctx, repositoryID) ([]standard.StandardReadResult, error)`；失败语义 `standard.ErrStandardReadFailed`；空列表非错误）
- `service/query_service.go`：`QueryService` 构造签名追加 `standardReader candidate.StandardReader`；`GetProjectBrief` 编排追加步骤——经 `standardReader` 读取 `standards[]`（nil 归一为空切片）；既有 7 步编排与失败语义不变
- `types.go`：`ProjectBriefReadResult` 追加 `Standards []standard.StandardReadResult`（json tag `standards`）
- `connect/server.go`：brief proto 响应组装 standards（domain → proto 转换，树结构递归组装复用模块内转换逻辑）
- **过渡态冻结**：`global_assets = 3` 与 `governance_profile = 2` 装配现状不动（phase14-04 步骤 1 "不做"清单）；brief 集成测试在既有 round-trip 用例上追加 `standards[]` 断言（绑定 Standard 后 brief 返回含全树；未绑定时空数组）

#### Scenario: brief 装配判定

- **WHEN** 绑定 Standard（任意 role）的 repository 调用 `GetProjectBrief`
- **THEN** `standards[]` 含该 Standard 全树 round-trip（节点字段零丢失）；`global_assets[]` 行为与 phase13 现状一致（过渡态并存）

### Requirement: 装配与错误映射必须闭合

- `connecterrors/connect_errors.go`：standard 5 哨兵注册（NotFound 分支 2 项 / InvalidArgument 分支 1 项 / Internal 段显式引用留痕 2 项，对齐 governanceprofile 注册模式）
- `platform/router.go`：追加 `buildStandard(pool)`（candidate → store → query/command service）与 `mountStandardConnect(r, querySvc, commandSvc)`；`buildProjectContext` 签名扩展追加 `standardReader` 参数（双 Reader 注入，phase14-06 行级清单 L407-409 归属本子任务的时序冻结）
- `platform/server.go`：`/api` 路由装配追加 standard 挂载；`buildProjectContext` 调用点传入 standard 模块 QueryService 作为 Reader 实现
- 分层约束复核：connect 层无 SQL / repository 层无业务判断 / service 层无跨模块直写 SQL（跨模块一律经 candidate）

#### Scenario: 装配判定

- **WHEN** 后端启动
- **THEN** `/api/psco.standard.v1/StandardService/...` 8 个 RPC 可达；`GetProjectBrief` 正常返回含 `standards[]`

### Requirement: DoD 验证门禁必须全绿

- `proto/`：`make lint` && `make build` && `make breaking` 通过
- `backend/`：`go build ./...` && `go vet ./...` && `go test ./...` 通过（含新模块全部测试）
- brief 集成测试含 `standards[]` round-trip 断言（验收环境 DB 可用时执行并留痕）
- 树校验单元测试覆盖含非法树用例（全部 8 个错误码可追溯）

## MODIFIED Requirements

无（本 spec 为实现类规格；全部实现语义引用 phase14-03/04/06 冻结口径，不修改上游设计）。

## REMOVED Requirements

无。
