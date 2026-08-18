# phase14-09 落实画像系统性退役与 brief 切换 Spec

## Why

phase14-06 已冻结画像退役的完整执行设计（0011 迁移三段结构与两段式算法 / 后端模块收缩文件级清单 / 前端切片移除清单 / brief 装配切换 / buf breaking 豁免 / 六触点断言矩阵），phase14-07/08 已交付 Standard 后端与前端承接位（Repository detail 让位已完成），画像退役的全部前置条件就绪。本子任务把退役设计执行为代码与数据变更，完成六触点退役与 brief 切换。本子任务为源代码实现类，执行语义一律以 phase14-06/02 冻结为准，不引入新设计。

## What Changes

- 产出本 spec 三件套（`phase14_09_retire_governance_profile_switch_brief`）
- **BREAKING** `database/migrations/0011_phase14_standard_entity.sql` 补第二、三段：数据迁移 DO 块（两段式算法：节点物化 TEMP TABLE → 同名合并 → 自底向上固定 6 轮组树 → 产物幂等写入）+ drop 两张 bindings 表（`governance_profiles` 主表保留）
- **BREAKING** 删除 `proto/psco/governance_profile/v1/` 整个包；`proto/buf.yaml` breaking 段追加单文件路径豁免留痕；`make gen` 三端产物清理
- **BREAKING** `project_context.proto` brief 切换：`governance_profile = 2` 类型改内联 `BriefGovernanceProfile`（phase14-02 逐字草案）；`global_assets = 3` 移除并 `reserved 3;`；`current_phase = 4` 的 status 改内联 `BriefPhaseStatus`；删 `psco/governance_profile` import
- 后端 `governanceprofile` 模块收缩为纯读：删 `connect/server.go` / `connect/server_test.go` / `service/command_service.go` / `candidate/repository_reader.go` 4 文件；收缩 `service/query_service.go`（删 `GetGovernanceProfile` 编排，新增 `ReadProfileCore`）/ `repository/profile_store.go`（删 `SaveProfile` 与两表读取，`ReadProfile` 只读主表）/ `types.go`（删写模型，新增 `GovernanceProfileCoreReadResult`）/ `errors.go`（删 3 哨兵保留 2）；收缩联动 `platform/server.go`（单值接收 + 删 mount 调用）与 `connecterrors/connect_errors.go`（删 3 映射行）
- `projectcontext/candidate/context_readers.go` `GovernanceProfileReader` 接口签名收缩为 `ReadProfileCore`；`projectcontext` service/types/集成测试联动切换（删 `GlobalAssetBinding` 装配与旧类型引用）
- 删除 `frontend/src/features/governance-profile/` 整目录（挂载点已由 phase14-08 让位，无引用）；验证 TS 生成产物消失
- 在验收 DB 执行补全后的 0011（全语句幂等，可手工 psql -f 重放），产出一条合并全局 Standard（裁决⑧）

## Impact

- Affected specs:
  - `docs/phase/phase14_standard_entity_foundation_dev_plan.md`（上游，phase14-09 定义 L61-64）
  - `.trae/specs/phase14_06_design_profile_retirement_data_migration/spec.md`（直接上游：本 spec 的执行对象，迁移算法/收缩清单/豁免配置/断言矩阵逐字源）
  - `.trae/specs/phase14_02_freeze_standard_entity_relation_boundaries/spec.md`（直接上游：brief 对照表 / `BriefGovernanceProfile` 内联消息逐字草案 L133-146 / 六触点边界）
  - `.trae/specs/phase14_03_design_standard_data_model_directory_tree/spec.md`（迁移字段级映射与拆树规则——DO 块算法落地依据）
  - `.trae/specs/phase14_07_land_standard_backend_mainline/spec.md`（standard 后端承接位现状）
  - `.trae/specs/phase14_08_land_standard_frontend_mainline/spec.md`（前端让位现状——挂载点已替换）
- Affected code:
  - `database/migrations/0011_phase14_standard_entity.sql`（追加第二、三段）
  - `proto/psco/governance_profile/v1/`（整包删除）、`proto/psco/project_context/v1/project_context.proto`、`proto/buf.yaml`
  - `backend/internal/gen/`、`frontend/src/gen/`（`make gen` 重生成）
  - `backend/internal/governanceprofile/`（删 4 收缩 4）、`backend/internal/platform/server.go`、`backend/internal/platform/router.go`、`backend/internal/connecterrors/connect_errors.go`
  - `backend/internal/projectcontext/`（candidate 接口 / service / types / connect / 集成测试）
  - `frontend/src/features/governance-profile/`（整目录删除）

## ADDED Requirements

### Requirement: 0011 迁移第二、三段必须按 phase14-06 冻结算法逐字落地

`database/migrations/0011_phase14_standard_entity.sql` SHALL 在现有第一段建表后追加（逐字执行 phase14-06 §ADDED-1，不引入新设计）：

**第二段：数据迁移 DO 块**

- 整段收敛为单个 `DO $$ ... $$` PL/pgSQL 块，块首守卫 `IF to_regclass('governance_canonical_root_file_bindings') IS NULL THEN RETURN; END IF;`（两表已 drop 的重放真跳过）
- 算法步骤：
  1. 节点行物化：`CREATE TEMP TABLE ... ON COMMIT DROP` 承载 `(path, parent_path, name, node_type, role, summary, ref, priority, subtree)`；源 1（`governance_canonical_root_file_bindings`）与源 2（`governance_global_asset_bindings`）按 phase14-03 字段映射与 `entry_ref` 三态规则展开；`GROUP BY path` 同名合并（源 2 非空值优先，`COALESCE(MAX(CASE WHEN priority=2 ...), MAX(CASE WHEN priority=1 ...), '')`）
  2. 自底向上固定轮次组树：**不用 `WITH RECURSIVE`**；节点行含 `subtree` 列（初值自身节点对象、`children='[]'`）；FOR 循环固定 6 轮，每轮一条非递归 UPDATE 重算 `children = (SELECT jsonb_agg(子行 subtree ORDER BY name) ...)`；根节点（`path='/'`、`name='.'`、`node_type='directory'`）末轮取 `subtree` 即单根 `directory_tree`
  3. 产物幂等写入：`INSERT INTO standards ... WHERE NOT EXISTS (name = 固定名)`；同批次首条 revision（含 N/M 计数与源 repository）；`standard_bindings ... ON CONFLICT DO NOTHING`（`adopts` 角色）
  4. 无画像数据场景：两源表零行 → 不产生迁移产物
- 迁移产物固定值（phase14-06 冻结表逐字）：`name = '默认项目范式（迁移自治理画像）'`、`status = 'active'`、`description = '由 phase14-09 迁移自动创建：合并项目治理画像的 canonical 根文件与全局资产两清单'`、首条 revision `change_summary = '迁移自项目治理画像：N 项 canonical 根文件 + M 项全局资产合并（源 repository <repository_id>）'`、binding `target_type='repository'` / `target_id=旧主表最新行 repository_id` / `role='adopts'`

**第三段：drop 两表**

- `DROP TABLE IF EXISTS governance_canonical_root_file_bindings;` + `DROP TABLE IF EXISTS governance_global_asset_bindings;`；`governance_profiles` 主表保留不触碰

#### Scenario: 迁移执行判定

- **WHEN** 在验收 DB（phase13-11 固定样本）手工 `psql -f` 执行补全后的 0011
- **THEN** 产出一条合并全局 Standard（树含全部源节点 / 首条 revision 留痕 / adopts 绑定正确）；重放一次无报错无重复产物（幂等实证）；迁移前导出两表全量基线，迁移后按 5 字段映射集逐项零丢失对照

### Requirement: proto 画像包删除与 brief 内联切换必须按冻结清单执行

- 删除 `proto/psco/governance_profile/v1/` 整包；`buf.yaml` breaking 段追加 `ignore: psco/governance_profile/v1/governance_profile.proto`（单文件路径，禁止目录前缀扩大化）+ 三行注释留痕（phase14-06 冻结 YAML 逐字）
- **执行期缺口修正留痕（T1）**：phase14-06"单文件豁免后 breaking 全绿"的断言与实测不符——brief 内联切换本身在 `project_context.proto` 上必然触发 `FIELD_WIRE_JSON_COMPATIBLE_TYPE`（governance_profile 字段与 BriefCurrentPhase.status 两处类型切换）与 `FIELD_NO_DELETE_UNLESS_NAME_RESERVED`（global_assets 需补字段名 reserved）。最小修正：① proto 侧 `reserved 3;` 增强为 `reserved 3;` + `reserved "global_assets";`（对 phase14-02 冻结口径纯增强，protobuf 语法要求编号与名称分列两条语句）；② buf.yaml 追加 `ignore_only: FIELD_WIRE_JSON_COMPATIBLE_TYPE: [psco/project_context/v1/project_context.proto]`（规则级 + 单文件豁免，比整文件 ignore 更窄；合入 main 重置基准后成为无实际作用的留痕）。修正后 `make breaking` 对 main 基准全绿（实证）。
- `project_context.proto` 按 phase14-02 逐字草案切换：`BriefTrackType` / `BriefPhaseStatus` / `BriefGovernanceProfile`（`repository_id = 1` / `track_type = 2` / `optional string template_source = 3`）内联；`governance_profile = 2` 类型替换；`global_assets = 3` 删除并 `reserved 3;`；`current_phase` 的 status 改 `BriefPhaseStatus`；删 `psco/governance_profile` import
- `make gen` 重生成后三端画像产物消失；`make lint && make build && make breaking` 全绿

#### Scenario: proto 退役判定

- **WHEN** Task 1 完成
- **THEN** `find proto backend/internal/gen frontend/src/gen -path '*governance_profile*'` 零输出（project_context 内联消息名不含该前缀）；brief 字段面 = phase14-02 对照表"后"列（8 顶层块）

### Requirement: 后端模块收缩与 brief 装配切换必须按文件级清单执行

按 phase14-06 §ADDED-2/3/4 逐字执行（清单与行级变更见该 spec，符号定位优先）：

- 删 4 文件：`governanceprofile/connect/server.go` / `connect/server_test.go` / `service/command_service.go` / `candidate/repository_reader.go`
- 收缩 4 文件：`service/query_service.go`（增 `ReadProfileCore`，构造去 `repositoryReader` 依赖）/ `repository/profile_store.go`（删 `SaveProfile` 与两表读取，`ReadProfile` 只读主表）/ `types.go`（增 `GovernanceProfileCoreReadResult`）/ `errors.go`（删 `ErrInvalidInput` / `ErrGovernanceProfileSaveFailed` / `ErrRepositoryNotFound`，保留 `ErrGovernanceProfileNotFound` / `ErrGovernanceProfileReadFailed`）
- 联动收缩：`platform/server.go`（L92 单值接收 / L93 删 mount 调用 / L97 调用对齐）、`platform/router.go`（删 `mountGovernanceProfileConnect` 函数与 3 处 import；`buildGovernanceProfile` 收缩单返回值）、`connecterrors/connect_errors.go`（删 3 映射行）
- `projectcontext` 切换：candidate `GovernanceProfileReader` 接口签名收缩为 `ReadProfileCore`；service 画像装配改新接口 + 删 `GlobalAssetBinding` 装配（两清单信息唯一来自 `StandardReader`，phase14-07 已接入 `standards = 8` 不回退）；types domain 收缩；connect 响应组装切换；集成测试切换（内联类型断言 + `BriefGovernanceProfile` 三字段 + `standards[]` 保持）
- 前端移除：删 `frontend/src/features/governance-profile/` 整目录；验证 `frontend/src/gen/proto/psco/governance_profile/` 不存在；`grep -r "governance-profile" frontend/src` 零命中

#### Scenario: 收缩判定

- **WHEN** Task 2-4 完成
- **THEN** `go build / vet / test ./...` 全绿；`tsc --noEmit` 零错误；模块内 `grep -r "CommandService\|SaveProfile\|connect/" backend/internal/governanceprofile` 零命中；`grep -rn "GovernanceProfileService\|governance_profile.v1" backend/internal/platform` 零命中

### Requirement: 退役六触点断言矩阵必须全量执行并留痕

phase14-06 §ADDED-5 六触点机械检查（T1-T6）逐条执行留痕；画像 RPC 404 实测（`/api/psco.governance_profile.v1/...` Connect 请求返回 404）；DoD 三项达成：

1. 六触点断言全绿（proto 目录无画像包 / 旧表已删无残留 / 模块无写方法 / 前端切片无残留 / 无画像 RPC 调用 / brief 语义不缺失）
2. brief 对照验收：旧清单信息经 `standards[]` 零丢失，`ref` 承接 `entry_ref / external_url`
3. 全库仅存在一条合并全局 Standard 且首条 revision 记录合并来源（裁决⑧）

#### Scenario: DoD 门禁判定

- **WHEN** Task 5-6 完成
- **THEN** checklist 全部勾选附证据；`SELECT count(*) FROM standards WHERE name = '默认项目范式（迁移自治理画像）'` = 1；两张 bindings 表 `to_regclass` 均 NULL；`governance_profiles` 主表行数与迁移前一致

## MODIFIED Requirements

无（执行类子任务；对 phase14-06 冻结设计逐字落地，不修改上游口径）。

## REMOVED Requirements

### Requirement: 治理画像写路径与跨包 brief 引用

**Reason**: phase14-02/06 冻结的系统性退役（六触点）；画像两组 bindings 信息已由 Standard 实体承接（phase14-07/08 交付）。
**Migration**: 存量数据经 0011 第二段 DO 块合并迁至一条全局 Standard（裁决⑧映射，零丢失对照验收）；brief `governance_profile` 切换为 `project_context.proto` 内联轻量消息；buf breaking 豁免留痕。
