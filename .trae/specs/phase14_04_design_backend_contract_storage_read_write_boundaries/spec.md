# phase14-04 Standard 后端合同、存储与读写边界设计 Spec

## Why

phase14-03 已冻结三表 DDL 与目录树规范，但 `psco.standard.v1` 合同（消息 + 8 个 RPC 的请求/响应 envelope 与错误语义）、Go 模块分层、写路径归属（web 写 / agent 只读）、`StandardReader` 跨模块读取接口签名与 brief 装配演进点尚未定义。本子任务产出可直接进入 phase14-07 实现的完整后端设计，实现类子任务不再做任何设计决策。本子任务为实现设计类，交付设计文档，不写实现代码、不创建 proto/Go/迁移文件。

## What Changes

- 产出本 spec 三件套（`phase14_04_design_backend_contract_storage_read_write_boundaries`）
- `psco.standard.v1` 合同包结构冻结：文件头注释 / package / go_package / 4 枚举 / 5 核心消息（`Standard` / `DirectoryTreeNode` / `StandardBinding` / `StandardRevision` / 请求响应 envelope），风格逐字对齐 `governance_profile.proto` 既有合同
- 8 个 RPC 三要素逐个冻结：`CreateStandard` / `ListStandards`（不分页）/ `GetStandard` / `UpdateStandard`（整树原子替换 + `change_summary` 必填）/ `DeleteStandard`（active 拒绝防误删）/ `BindStandard` / `UnbindStandard` / `ListStandardRevisions`，每个含请求字段 / 响应字段 / 错误语义
- Go 模块分层冻结（`backend/internal/standard/` 8 文件清单与职责单值映射，沿袭工程约定"支撑文件按职责单值化映射到唯一文件"）
- 写路径归属冻结：5 写 RPC（web 发起）+ 3 读 RPC（web/agent 共享）矩阵；无 agent 写回承接位（CON-09）
- `StandardReader` 跨模块读取接口签名冻结（沿袭 phase13-10 `GovernanceProfileReader` 消费方拥有模式，落点 `projectcontext/candidate/context_readers.go`）
- brief 装配演进点冻结：phase14-07（standards[] 新增接入）与 phase14-09（画像退役切换）的边界分工
- 单值化决策留档：brief `standards = 8` 类型直接引用 `psco.standard.v1.Standard`（不另造 StandardSummary 消息），对 phase14-02 对照表中"StandardSummary"占位名的实现级收敛

## Impact

- Affected specs:
  - `docs/phase/phase14_standard_entity_foundation_dev_plan.md`（上游，phase14-04 定义 L34-37）
  - `.trae/specs/phase14_03_design_standard_data_model_directory_tree/spec.md`（直接上游：DDL 草案 / jsonb 序列化规范 / proto 对齐表 / 8 条校验规则）
  - `.trae/specs/phase14_02_freeze_standard_entity_relation_boundaries/spec.md`（八格矩阵 / brief 对照表 / T6 装配切换上游约束）
  - `docs/phase/phase14_standard_entity_foundation_shared_baseline.md`（§2.3 技术主线 / §3.2-3.4 能力矩阵）
- Affected code: 无（零代码改动；设计对象为未来的 `proto/psco/standard/v1/standard.proto`、`backend/internal/standard/`、`backend/internal/projectcontext/candidate/context_readers.go` 追加、`backend/internal/platform/router.go` 装配点，均由 phase14-07 落地）
- 设计参照（本轮实际读取）：`proto/psco/governance_profile/v1/governance_profile.proto`（合同风格 / 错误语义注释模式）、`backend/internal/governanceprofile/`（模块分层现状）、`backend/internal/governanceprofile/candidate/repository_reader.go`（candidate 模式）、`backend/internal/governanceprofile/errors.go`（哨兵错误模式）、`backend/internal/projectcontext/candidate/context_readers.go`（消费方拥有 reader 接口模式）

## ADDED Requirements

### Requirement: psco.standard.v1 合同包结构必须冻结

`proto/psco/standard/v1/standard.proto` SHALL 按以下结构实现（文件头注释风格对齐 `governance_profile.proto`：文档定位 / 上游规格 / 合同约束 / 生成入口四段）：

- `syntax = "proto3"` / `package psco.standard.v1` / `option go_package = "github.com/psco/backend/internal/gen/proto/psco/standard/v1;standardv1"`
- `import "google/protobuf/timestamp.proto"`（created_at / updated_at 承载）
- 枚举 4 个（全部带 `_UNSPECIFIED = 0`）：

```proto
// StandardStatus 规范生命周期（对齐 DDL status CHECK 枚举）
enum StandardStatus {
  STANDARD_STATUS_UNSPECIFIED = 0;
  STANDARD_STATUS_DRAFT = 1;
  STANDARD_STATUS_ACTIVE = 2;
  STANDARD_STATUS_RETIRED = 3;
}

// NodeType 目录树节点类型（值序冻结于 phase14-03 proto 对齐表）
enum NodeType {
  NODE_TYPE_UNSPECIFIED = 0;
  NODE_TYPE_DIRECTORY = 1;
  NODE_TYPE_FILE = 2;
}

// BindingTargetType 绑定目标类型（可扩展：扩 enum + 登记 phase14-02 八格矩阵）
enum BindingTargetType {
  BINDING_TARGET_TYPE_UNSPECIFIED = 0;
  BINDING_TARGET_TYPE_REPOSITORY = 1;
  BINDING_TARGET_TYPE_PRODUCT = 2;
  BINDING_TARGET_TYPE_DECISION = 3;
  BINDING_TARGET_TYPE_MODULE = 4;
}

// BindingRole 绑定角色（可扩展；组合合法性按 phase14-02 八格矩阵）
enum BindingRole {
  BINDING_ROLE_UNSPECIFIED = 0;
  BINDING_ROLE_TEMPLATE_SOURCE = 1;
  BINDING_ROLE_ADOPTS = 2;
}
```

- 核心消息 4 个（不含请求响应 envelope）：

```proto
// DirectoryTreeNode 目录树节点（与 directory_tree jsonb 单值映射，phase14-03 冻结结构）
message DirectoryTreeNode {
  string name = 1;                        // 根固定 "."；非根 ^[A-Za-z0-9._-]{1,64}$
  NodeType node_type = 2;
  string role = 3;                        // file 必填（校验层保证）；directory 可空
  string summary = 4;                     // 结构化摘要 ≤2000
  string ref = 5;                         // "/..." 树内绝对路径或 https:// URL
  repeated DirectoryTreeNode children = 6; // file 节点必须为空
}

// Standard 全局规范（主表投影 + 整树；List/Get/brief 统一使用本消息——不另造 StandardSummary，见设计决策）
message Standard {
  string id = 1;
  string name = 2;                        // 全局唯一
  string description = 3;                 // 可空
  StandardStatus status = 4;
  DirectoryTreeNode directory_tree = 5;   // 整树原子承载（单根）
  google.protobuf.Timestamp created_at = 6;
  google.protobuf.Timestamp updated_at = 7;
}

// StandardBinding 多态绑定关系（standard_bindings 表投影）
message StandardBinding {
  string id = 1;
  string standard_id = 2;
  BindingTargetType target_type = 3;
  string target_id = 4;                   // UUID；存在性由应用层校验
  BindingRole role = 5;
  string note = 6;                        // 可空
  google.protobuf.Timestamp created_at = 7;
}

// StandardRevision 演进留痕（standard_revisions 表投影，只追加）
message StandardRevision {
  string id = 1;
  string standard_id = 2;
  string change_summary = 3;              // 必填，人工一句话
  google.protobuf.Timestamp created_at = 4;
}
```

设计决策（单一 Standard 消息，无 StandardSummary）：

- phase14-02 brief 对照表中 `repeated StandardSummary`（含 directory_tree 全树）为语义占位名；本 spec 单值化为直接引用 `Standard` 消息
- 理由：`Standard` 本身即轻量全量投影（无正文类重型字段，树 ≤64KB 由 R6 保证）；单用户系统规范数量级为个位数到十位数，List 场景带树传输总量可控（反假大空，不预制分页/摘要双消息）；避免同一树结构在两个消息中重复定义形成第二套字段语义
- 若未来规范数量增长到列表传输成为负担，再增加 Summary 消息（YAGNI，当前不预制）

#### Scenario: proto 实现判定

- **WHEN** phase14-07 编写 `standard.proto`
- **THEN** package / go_package / 4 枚举值序 / 4 消息字段与字段号与本节逐字一致；文件头注释四段齐全

### Requirement: 8 个 RPC 三要素必须逐个冻结

`StandardService` SHALL 承接 8 个 RPC，每个 RPC 的请求 / 响应 / 错误语义按下表单值冻结（错误 code 三类沿袭仓库既有模式：NotFound / InvalidArgument / Internal，不引入新分类）：

| # | RPC | 请求 | 响应 | 错误语义 |
|---|---|---|---|---|
| 1 | `CreateStandard` | `name`(必填) / `description`(可选) / `directory_tree`(必填，允许单根空树) / `status`(可选，默认 DRAFT；不允许 RETIRED) | `Standard` | name 重复 / 树校验 R1-R8 失败 / `status=RETIRED` / `status=ACTIVE` 且空树(R2) → InvalidArgument；创建失败 → Internal |
| 2 | `ListStandards` | 空请求（无参数） | `repeated Standard`（含树，按 `updated_at` DESC） | 读取失败 → Internal |
| 3 | `GetStandard` | `standard_id` | `Standard` + `repeated StandardBinding bindings`（绑定管理区直接消费） | id 非法 UUID → InvalidArgument；不存在 → NotFound；读取失败 → Internal |
| 4 | `UpdateStandard` | `standard_id` / `name`(optional，未设置不变更) / `description`(optional) / `status`(optional) / `directory_tree`(**必带**，整树原子替换) / `change_summary`(**必填**) | `Standard`（更新后） | 不存在 → NotFound；`change_summary` 空 / 树校验失败 / name 重复 / 目标状态非 DRAFT 且树空(R2) → InvalidArgument；事务失败 → Internal（整树替换 + updated_at + revision 追加同一事务，不写半套状态） |
| 5 | `DeleteStandard` | `standard_id` | 空 | 不存在 → NotFound；`status=ACTIVE` → InvalidArgument（防误删，先经 Update 置 RETIRED）；删除失败 → Internal；CASCADE 连带删除 bindings 与 revisions（DDL 冻结） |
| 6 | `BindStandard` | `standard_id` / `target_type` / `target_id` / `role` / `note`(可选) | `StandardBinding` | standard 不存在 → NotFound；枚举非法 / 八格矩阵非法（template_source 携非 repository） / target 不存在（按 target_type 查实体表，phase14-02 冻结） / 四元组重复绑定 → InvalidArgument；写入失败 → Internal |
| 7 | `UnbindStandard` | `standard_id` / `target_type` / `target_id` / `role`（四元组定位，note 不参与） | 空 | 绑定不存在 → NotFound；删除失败 → Internal |
| 8 | `ListStandardRevisions` | `standard_id` | `repeated StandardRevision`（按 `created_at` DESC，不分页） | 不存在 → NotFound；读取失败 → Internal |

语义要点（单值化说明）：

- **CreateStandard 不记 revision**：创建时间由 `Standard.created_at` 承接；revision 从第一次 Update 开始留痕。裁决⑧迁移场景的"首条 revision 记录合并来源"由迁移脚本直接 INSERT 承接，不经本 RPC
- **UpdateStandard 无部分更新**：`directory_tree` 必带（整树替换语义，baseline §2.4 冻结"编辑保存 = 整树校验 + 替换 + 追加 revision"）；仅改 status（如 draft→active 激活）时前端同样读后写全量，不设增量协议
- **CreateStandard 的 status 约束**：允许 DRAFT（默认）/ ACTIVE（创建即激活，树必须满足 R2）；拒绝 RETIRED（不能创建已退役规范，防呆）
- **ListStandards / ListStandardRevisions 不分页**：单用户系统数量级（规范个位数、revision 单规范几十条内），分页为过度设计（CON-07 反假大空）
- **重复绑定错误归类 InvalidArgument**：错误信息含 "already bound" 区分语义；不引入 AlreadyExists 第四类 code（沿袭哨兵错误三类归一化模式）

#### Scenario: RPC 实现判定

- **WHEN** phase14-07 实现 connect handler 与 service
- **THEN** 8 个 RPC 的请求/响应字段、错误映射与本表逐字一致；UpdateStandard 单事务边界（树替换 + updated_at + revision 追加）有集成测试断言

#### Scenario: 绑定校验链判定

- **WHEN** phase14-07 实现 `BindStandard`
- **THEN** 校验顺序固定：standard 存在 → 枚举合法 → 八格矩阵 → target 存在 → 唯一约束；每步失败返回对应错误，不跳步

### Requirement: Go 模块分层必须冻结（文件级）

`backend/internal/standard/` SHALL 按以下 8 文件清单实现（沿袭 governanceprofile 模块分层 + 工程约定"支撑文件按职责单值化映射到唯一文件"）：

| 文件 | 职责 |
|---|---|
| `connect/server.go` | ConnectRPC handler：8 RPC 装配；调用 `connecterrors.MapToConnectError` 统一错误映射；不含业务逻辑 |
| `service/command_service.go` | 写路径 owner：Create / Update / Delete / Bind / Unbind；UpdateStandard 单事务边界；调用 validate 与 candidate |
| `service/query_service.go` | 读路径 owner：List / Get / ListRevisions；`StandardReader` 接口实现位（ListStandardsByRepository） |
| `repository/standard_store.go` | 三表 SQL（standards / standard_revisions / standard_bindings）；jsonb 编解码；不含业务判断 |
| `candidate/target_reader.go` | 跨模块 target 存在性校验（4 实体表 EXISTS 查询；沿袭 governanceprofile/candidate 模式，service 不直接写跨模块 SQL） |
| `errors.go` | 哨兵错误（三类：ErrStandardNotFound / ErrBindingNotFound → NotFound；ErrInvalidInput → InvalidArgument；ErrStandardReadFailed / ErrStandardSaveFailed → Internal） |
| `types.go` | `DirectoryTreeNode` Go 结构（phase14-03 冻结 json tag）+ `StandardReadResult` 等读写模型 |
| `validate.go` | 8 条树校验规则 R1-R8（判定逻辑 / 错误码 / 节点路径定位，phase14-03 冻结） |

分层约束：

- connect 层不写 SQL、repository 层不做业务判断、service 层不直接写跨模块 SQL（跨模块读取一律经 candidate 子包）
- `platform/router.go` 装配点：构造 store → command/query service → connect server → 路由注册；`StandardReader` 实现注入 `projectcontext` 装配（见下一 Requirement）

#### Scenario: 模块实现判定

- **WHEN** phase14-07 落地 `backend/internal/standard/`
- **THEN** 文件清单与本表一一对应，无多余文件、无职责漂移；validate.go 的 8 条规则有单元测试（phase14-03 Scenario 要求）

### Requirement: 写路径归属必须冻结（web 写 / agent 只读）

8 个 RPC 的消费方归属 SHALL 冻结为：

| RPC | web（维护路径） | agent（消费路径） |
|---|---|---|
| CreateStandard / UpdateStandard / DeleteStandard / BindStandard / UnbindStandard | ✓（唯一写入口，前端 `/standards` 承接，裁决⑦） | ✗（无 agent 写回承接位，CON-09） |
| ListStandards / GetStandard / ListStandardRevisions | ✓ | ✓（agent 经 Connect 直读，或经 GetProjectBrief 聚合消费） |

约束说明：

- 当前无鉴权体系，本边界为**语义边界**（合同注释 + 前端切片归属表达），非技术强制；未来引入鉴权时按本矩阵落权
- agent 的规范消费主路径是 `GetProjectBrief.standards[]`（按 repository 聚合）；`ListStandards` / `GetStandard` 为全量直读补充（如跨仓库规范检索场景）

#### Scenario: 归属边界判定

- **WHEN** phase14-08 实现前端切片或 phase14-10 验收 dogfooding
- **THEN** 前端 mutation 仅出现在 `/standards` 切片（裁决⑦维护动作唯一入口）；agent 侧取证仅出现读调用

### Requirement: StandardReader 跨模块读取接口必须冻结

brief 对 Standard 的读取 SHALL 经消费方拥有的 candidate 接口承接（沿袭 phase13-10 `GovernanceProfileReader` 冻结模式），落点 `backend/internal/projectcontext/candidate/context_readers.go` 追加：

```go
// StandardReader 全局规范读取接口（消费方拥有的 candidate 接口）。
//
// phase14-04 冻结：brief 对 Standard 的读取必须通过本接口承接，
// 由 platform 装配点注入 standard/service.QueryService 作为实现；
// projectcontext 不得直接书写 standard 表 SQL 或复制其存储读取逻辑。
type StandardReader interface {
	// ListStandardsByRepository 经 standard_bindings（任意 role）反查
	// 该仓库关联的全部 Standard（含 directory_tree 全树）。
	// 失败语义：读取失败 → standard.ErrStandardReadFailed；
	//           仓库无关联 Standard → 返回空列表（非错误）。
	ListStandardsByRepository(ctx context.Context, repositoryID string) ([]standard.StandardReadResult, error)
}
```

装配约束：

- 实现位：`standard/service/query_service.go`（读路径 owner）
- 注入位：`platform/router.go` 装配点（与 `GovernanceProfileReader` 注入并列）
- 反查 SQL 归属：实现方（standard 模块）拥有 `standard_bindings` 表的读取；projectcontext 经接口隔离

#### Scenario: 接口实现判定

- **WHEN** phase14-07 实现 StandardReader 并接入 brief 装配
- **THEN** 接口签名与本节逐字一致；`projectcontext` 无 standard 表直接 SQL；`GetProjectBrief` 集成测试含 `standards[]` round-trip 断言（dev_plan phase14-07 DoD）

### Requirement: brief 装配演进点必须冻结（phase14-07 与 phase14-09 分工）

`GetProjectBrief` 的演进 SHALL 按以下边界分两步落地（对齐 dev_plan 子任务划分与 phase14-02 T6）：

| 步骤 | 子任务 | 动作 | 不做 |
|---|---|---|---|
| 步骤 1（新增接入） | phase14-07 | `project_context.proto` 增 `import "psco/standard/v1/standard.proto"` 与 `standards = 8`（`repeated psco.standard.v1.Standard`）；`query_service.go` 装配 `StandardReader` 调用；集成测试断言 `standards[]` | 不动 `global_assets = 3` 与 `governance_profile = 2` 现状（此时新旧并存，phase14-09 前的过渡态） |
| 步骤 2（退役切换） | phase14-09 | 移除 `global_assets = 3`（reserved 3）；`governance_profile = 2` 切内联 `BriefGovernanceProfile`（phase14-02 字段清单）；解除 governance_profile 包 import | 不新增第 9 顶层块；不改动 `repository / current_phase / products / modules / decisions` |

跨包引用说明（设计决策）：

- `project_context.proto` 引用 `psco.standard.v1.Standard` 构成新的跨包依赖——这是**对活跃合同源的引用**，与 phase14-02 T6 要解除的"对退役包的跨包依赖"性质不同；phase13 现状（project_context 引用当时活跃的 governance_profile 包）即此模式
- 理由：树结构是 agent 消费的核心物，brief 必须携带全树；跨包引用保证单一消息定义源（不在 project_context.proto 重复定义树结构形成第二套字段语义）
- phase14-09 完成后 brief 的最终字段面 = 8 顶层块（phase14-02 对照表），其中两清单信息唯一来源 `standards[]`

#### Scenario: 分步落地判定

- **WHEN** phase14-07 完成（步骤 1）
- **THEN** brief 同时含 `global_assets[]`（旧）与 `standards[]`（新）——过渡态字段面有集成测试快照；`governance_profile` 仍为旧跨包类型
- **WHEN** phase14-09 完成（步骤 2）
- **THEN** brief 字段面与 phase14-02 对照表"后"列逐字一致；无 governance_profile 包 import；buf breaking 对画像包删除豁免留痕

## MODIFIED Requirements

无（本 spec 为新增设计规格；对 phase14-02 brief 对照表中 `StandardSummary` 占位名的单值化收敛已在该 Requirement 内显式声明理由，不构成对上游冻结口径的修改——"含 directory_tree 全树与导航摘要"语义由 `Standard` 消息完整承接）。

## REMOVED Requirements

无。
