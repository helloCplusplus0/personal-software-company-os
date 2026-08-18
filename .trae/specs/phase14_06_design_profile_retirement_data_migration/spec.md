# phase14-06 画像退役与数据迁移设计 Spec

## Why

phase14-02 已冻结六触点退役边界与 brief 字段演进对照，phase14-03 已冻结两表 → 树节点的字段级迁移映射，但退役的执行层设计——`0011` 迁移的机械结构与幂等语义、后端模块收缩的文件级清单与 `GovernanceProfileReader` 接口签名收缩、前端切片移除清单、buf breaking 豁免配置——尚未单值化。本子任务把这些执行细节冻结为 phase14-09 的直接执行输入。本子任务为实现设计类，交付设计文档，不写实现代码、不创建 SQL/Go/TS 文件。

## What Changes

- 产出本 spec 三件套（`phase14_06_design_profile_retirement_data_migration`）
- `0011_phase14_standard_entity.sql` 迁移设计冻结：三段结构（建表 / 数据迁移 / drop 两表）、数据迁移两段式算法（节点行物化 → DO 块自底向上固定轮次组树，独立复核修正：弃 `WITH RECURSIVE` 形态）、幂等守卫矩阵（逐语句级别）、迁移产物固定值（name / status / description / 首条 revision / adopts 绑定决策）、无画像数据场景决策、可重放设计（DO 块 `to_regclass` 真跳过守卫）
- 后端模块收缩文件级清单冻结：删 4 文件、收缩 4 文件 + 收缩联动 2 文件（`server.go` / `connect_errors.go`，独立复核补齐的清单外硬编译断点）、`GovernanceProfileReader` 接口签名收缩设计（新轻量类型 `GovernanceProfileCoreReadResult` + 方法 `ReadProfileCore`）、`router.go` 行级变更清单
- 前端 `features/governance-profile` 切片移除清单冻结（10 文件 + 挂载点 + 时机）
- brief 装配切换执行设计冻结（双 Reader 注入 / 对照表引用 / TS 生成产物清理）
- buf breaking 对画像包删除的豁免配置方案冻结（`buf.yaml` `breaking.ignore` + 注释留痕）
- 退役六触点验收断言矩阵执行化（每触点一条机械可验证检查命令）

## Impact

- Affected specs:
  - `docs/phase/phase14_standard_entity_foundation_dev_plan.md`（上游，phase14-06 定义 L44-47）
  - `.trae/specs/phase14_02_freeze_standard_entity_relation_boundaries/spec.md`（直接上游：六触点 T1-T6 / brief 对照表 / `BriefGovernanceProfile` 内联消息）
  - `.trae/specs/phase14_03_design_standard_data_model_directory_tree/spec.md`（直接上游：三表 DDL / 迁移字段级映射 / 拆树规则 / 同名冲突合并——本 spec 的算法落地对象）
  - `.trae/specs/phase14_04_design_backend_contract_storage_read_write_boundaries/spec.md`（standard 模块结构与 StandardReader——与收缩后的画像模块并存形态）
- Affected code: 无（零代码改动；设计对象为未来的 `database/migrations/0011_phase14_standard_entity.sql`、`backend/internal/governanceprofile/` 收缩、`backend/internal/platform/router.go`、`proto/buf.yaml`、`frontend/src/features/governance-profile/` 移除，均由 phase14-09 落地）
- 设计参照（本轮实际读取）：`backend/internal/platform/migrate.go`（L103-124 单事务整文件执行机制——0011 原子性依据）、`database/migrations/0010_phase13_governance_profile.sql`（两源表实证）、`backend/internal/governanceprofile/` 全模块 10 文件（收缩对象现状）、`backend/internal/projectcontext/candidate/context_readers.go` L26-36（GovernanceProfileReader 接口现状）、`backend/internal/platform/router.go` L69-90/L305-313/L396-422（装配与挂载现状）、`proto/buf.yaml`（v2 breaking 配置现状）、`proto/Makefile`（breaking target 现状）

## ADDED Requirements

### Requirement: 0011 迁移文件必须按三段结构 + 两段式算法设计

`database/migrations/0011_phase14_standard_entity.sql` SHALL 按以下结构实现（三段顺序固定，单事务原子——`migrate.go` applyMigration 整文件单事务，失败全回滚不丢旧数据）：

**第一段：建表（DDL）**

- 按相位 phase14-03 冻结的三表 DDL 草案落地 `standards` / `standard_revisions` / `standard_bindings`（`CREATE TABLE IF NOT EXISTS` + 索引）
- 注意：phase14-07 后端主线落地时创建此文件；本 spec 冻结其完整设计，编号 `0011` 承 phase14-07 dev_plan 已修正的四位编号规范

**第二段：数据迁移（两段式算法，仅当画像数据存在时执行）**

执行守卫（可重放设计核心）：**数据迁移整段收敛为单个 `DO $$ ... $$` PL/pgSQL 块**，块首 `IF to_regclass('governance_canonical_root_file_bindings') IS NULL THEN RETURN; END IF;`——两表已 drop 的重放场景真跳过（块内后续语句不执行，不报错）。该形态同时自洽于两条执行路径：migrate.go 单事务整文件执行（DO 块在文件事务内原子执行）与手工 `psql -f` 重放（psql autocommit 下 DO 块为单语句单事务，守卫不满足时整块空过）。
> 独立复核 FAIL-3 修正留痕：原设计"`WHERE` 条件包裹 + 多条独立语句"在 psql autocommit 重放下，`TEMP TABLE ... ON COMMIT DROP` 语句级提交即被删除，后续引用语句必然报 `relation does not exist`；收敛为单 DO 块后中间表生命周期 = 块事务，两条路径均安全。

算法步骤（冻结为步骤级，phase14-09 落地为 SQL；字段映射逐字引用 phase14-03 迁移映射节，不重复定义）：

1. **节点行物化**：DO 块内 `CREATE TEMP TABLE ... ON COMMIT DROP` 承载中间节点行 `(path, parent_path, name, node_type, role, summary, ref, priority, subtree)`（临时表生命周期 = DO 块事务，不跨语句泄漏）
   - 源 1 展开（`governance_canonical_root_file_bindings`）：每行 → `parent_path='/'`、`path='/'+file_name`、`node_type='file'`、`role` 原值、`summary=''`、`ref='/'+file_name`、`priority=1`
   - 源 2 展开（`governance_global_asset_bindings`）：按 phase14-03 `entry_ref` 三态规则展开——`https://` 前缀 → 根下 file 节点（`ref` 原值）；裸文件名 → 根下 file 节点（`ref='/'+entry_ref`）；含路径 → 去前导 `/` 按 `/` 拆段，中间段逐层物化 `directory` 节点（`role`/`summary` 空），末段为 file 节点（末段节点 `name` = 最后路径段，按 phase14-03 迁移映射表格主口径执行——独立复核 OBS-7 消歧）；`summary` 按 `"[" + kind + "] " + structured_summary` 格式合成（空 structured_summary 时为 `"[" + kind + "]"`）；`priority=2`
   - 同名冲突合并（聚合完成，非过程式）：`GROUP BY path` 后 `role`/`summary` 取 `COALESCE(MAX(CASE WHEN priority=2 THEN ...), MAX(CASE WHEN priority=1 THEN ...), '')`——即源 2 非空值优先覆盖、源 2 空保留源 1（phase14-03 冻结合并规则的集合化落地）
2. **自底向上固定轮次组树**：**不使用 `WITH RECURSIVE`**（独立复核 FAIL-1 修正留痕：PostgreSQL 递归项内禁止聚合函数、子查询内禁止递归引用，"自顶向下每层 `jsonb_agg`"形态不可实现）。改为：节点行含 `subtree` 列（初值 = 自身节点对象、`children='[]'`）；DO 块内 FOR 循环固定 6 轮（R5 冻结深度 ≤6），每轮一条非递归 UPDATE 对全部节点重算 `children = (SELECT jsonb_agg(子行 subtree ORDER BY name) FROM 同临时表 WHERE parent_path = 本行 path)`——相关子查询作用于普通临时表，无递归 CTE 限制；按 UPDATE 快照语义每轮恰好上卷一层，重复轮次幂等（已正确子树重算不变）。根节点（`path='/'`、`name='.'`、`node_type='directory'`）亦物化于临时表，末轮取其 `subtree` 即单根 `directory_tree` jsonb。子节点排序默认 `name` 升序保证产物字节稳定（若 phase14-03 另有冻结排序从之）
3. **产物写入**（幂等守卫）：
   - `INSERT INTO standards (...) SELECT ... WHERE NOT EXISTS (SELECT 1 FROM standards WHERE name = $固定名)`——整树单行原子写入
   - 同批次 `INSERT INTO standard_revisions (standard_id, change_summary)` 引用刚插入行（首条 revision 记录合并来源，裁决⑧）
   - `INSERT INTO standard_bindings (standard_id, target_type, target_id, role) ... ON CONFLICT DO NOTHING`（唯一约束四元组兜底）——源 `repository_id` 取旧 `governance_profiles` 主表最新 `updated_at` 行（裁决⑧"多画像取最新"口径的落地；当前单实例）
4. **无画像数据场景决策**：两源表存在但零行（或主表零行）→ 不产生迁移产物（standards 无该固定名行、无 revision、无 binding）——用户后续手工创建；空数据不生成空树 Standard（避免无意义 draft 行）

**第三段：drop 两表**

- `DROP TABLE IF EXISTS governance_canonical_root_file_bindings;` + `DROP TABLE IF EXISTS governance_global_asset_bindings;`（先迁数据后 drop，dev_plan 冻结顺序；`governance_profiles` 主表保留——T2）

**迁移产物固定值（单值冻结）**：

| 字段 | 值 | 说明 |
|---|---|---|
| `name` | `默认项目范式（迁移自治理画像）` | 固定名（幂等守卫键）；中文自明，phase14-10 断言"仅存在一条合并全局 Standard"以此为准 |
| `status` | `'active'` | 迁移产物即正在使用的规范；树含 file 节点天然满足 R2 |
| `description` | `由 phase14-09 迁移自动创建：合并项目治理画像的 canonical 根文件与全局资产两清单` | 固定说明 |
| 首条 revision `change_summary` | `迁移自项目治理画像：N 项 canonical 根文件 + M 项全局资产合并（源 repository <repository_id>）` | N/M 为迁移时实际计数（动态）；合并来源留痕（裁决⑧验收断言依据） |
| binding | `target_type='repository'`、`target_id=旧主表最新行 repository_id`、`role='adopts'` | **决策留档**：迁移默认建 `adopts`（保守语义：该仓库遵守此规范）；`template_source` 是更强语义主张（规范有实际示范仓库），不由迁移自动断言，需用户显式维护 |

#### Scenario: 迁移可机械执行判定

- **WHEN** phase14-09 落地 `0011` 并在 phase13-11 固定样本库上执行
- **THEN** 三段顺序执行、单事务原子；产物树的每个节点可按 phase14-03 字段映射逐字段回溯到源行（零丢失对照验收）；`GetStandard` 读取该树通过 8 条校验（只读断言）

#### Scenario: 幂等与可重放判定

- **WHEN** `0011` 已应用后再次执行（手工 `psql -f` 重放，如 phase14-10 验收 runbook）
- **THEN** 全部语句幂等不报错：`CREATE TABLE IF NOT EXISTS` 跳过、数据迁移 DO 块守卫真跳过（两表已 drop，块内语句不执行——含临时表创建，无 `relation does not exist` 风险）、`DROP TABLE IF EXISTS` 跳过；不产生重复 Standard / revision / binding

### Requirement: 后端模块收缩必须按文件级清单执行

`backend/internal/governanceprofile/` SHALL 从现状 9 文件收缩为纯读模块（现状行数本轮实证；独立复核 OBS-1 修正：模块文件计数 9 非笔误性的 10）：

| 动作 | 文件 | 现状行数 | 说明 |
|---|---|---|---|
| 删除 | `connect/server.go` | — | RPC handler 层（T4） |
| 删除 | `connect/server_test.go` | — | 随 handler 删除 |
| 删除 | `service/command_service.go` | 134 | 写路径编排（T3） |
| 删除 | `candidate/repository_reader.go` | 42 | 画像写入的仓库存在性前提校验——纯写入候选服务（本轮实证：仅被 Query/Command 构造注入），随写路径退役（T3"candidate 若仅为画像写入候选服务则一并删除"确认适用） |
| 收缩 | `service/query_service.go` | 61 | 删 `GetGovernanceProfile` RPC 编排方法；新增 `ReadProfileCore` 方法（Reader 实现见下）；构造签名去 `repositoryReader` 依赖 |
| 收缩 | `repository/profile_store.go` | 274 | 删 `SaveProfile`（L49）写方法及两 bindings 表读取；`ReadProfile`（L124）改为只读主表（两表已 drop，JOIN 必然失败——收缩是编译期外的前置必要） |
| 收缩 | `types.go` | 167 | 删写输入模型与 bindings 相关模型；新增 `GovernanceProfileCoreReadResult`（轻量三字段组：`TrackType` / `TemplateSource` / `CurrentPhaseName/Ref/Status`） |
| 收缩 | `errors.go` | 31 | 删 `ErrInvalidInput` / `ErrGovernanceProfileSaveFailed` / `ErrRepositoryNotFound`（哨兵单值口径见下）；保留 `ErrGovernanceProfileNotFound` / `ErrGovernanceProfileReadFailed` |
| 保留 | `repository/profile_store_integration_test.go` | — | 收缩为只读断言（读主表路径保留覆盖） |

**`GovernanceProfileReader` 接口签名收缩设计**（`projectcontext/candidate/context_readers.go` L26-36 同步修改）：

```go
// GovernanceProfileReader 治理画像主记录轻量读取接口（消费方拥有的 candidate 接口）。
//
// phase14-06 冻结：接口随画像退役收缩为只读主表三组字段
// （track_type / template_source / current_phase 三字段，服务 brief 内联装配）；
// 两组 bindings 信息已迁移至 Standard（经 StandardReader 读取）。
// 实现仍由 platform 装配点注入 governanceprofile/service.QueryService。
type GovernanceProfileReader interface {
	// ReadProfileCore 读取画像主记录核心字段（不含已退役的两组 bindings）。
	// 失败语义：画像未创建 → ErrGovernanceProfileNotFound；
	//           其他读取失败 → ErrGovernanceProfileReadFailed。
	ReadProfileCore(ctx context.Context, repositoryID string) (*governanceprofile.GovernanceProfileCoreReadResult, error)
}
```

**`router.go` 行级变更清单**（现状本轮实证）：

| 行 | 变更 |
|---|---|
| L69（`governanceprofilecandidate` import） | 删 |
| L70（`governanceprofileconnect` import） | 删 |
| L90（`governanceprofilev1connect` gen import） | 删 |
| L305-315（`mountGovernanceProfileConnect` 函数定义） | 整体删除（T4：路由表无画像服务注册；本轮实证该段仅函数定义，实际调用点在 `server.go` L93，见下方联动清单） |
| L407-409（`buildProjectContext`） | 签名扩展：追加 `standardReader projectcontextcandidate.StandardReader` 参数（双 Reader 注入，对齐 phase14-04 装配设计；**该项由 phase14-07 过渡态先行落地，phase14-09 仅核对不再改**） |
| L412-422（`buildGovernanceProfile`） | 收缩：删 commandSvc 构造与 candidate 构造；返回值改为单 `*governanceprofileservice.QueryService`（Reader 实现注入 brief 与 Standard 装配并列） |

> 行号按 phase14-06 时点现状冻结（本轮 grep 实证）；phase14-07 将先行改动 router.go（standard 模块装配），phase14-09 执行时**以符号定位为准**，行号仅作时点参照（独立复核 OBS-2 补注）。

**收缩联动文件（独立复核 FAIL-2 补齐——清单外硬编译断点，缺此 `go build ./...` 零错误不成立）**：

| 文件 | 行 | 变更 |
|---|---|---|
| `backend/internal/platform/server.go` | L92 | `governanceProfileQuerySvc, governanceProfileCommandSvc := buildGovernanceProfile(pool)` 双赋值改单值接收（buildGovernanceProfile 收缩为单返回值） |
| `backend/internal/platform/server.go` | L93 | 删 `mountGovernanceProfileConnect(r, ...)` 调用（函数已删；实际调用点在此） |
| `backend/internal/platform/server.go` | L97 | `buildProjectContext(pool, governanceProfileQuerySvc)` 调用对齐收缩后签名（standardReader 参数已由 phase14-07 先行注入） |
| `backend/internal/connecterrors/connect_errors.go` | L67 | 删 `governanceprofile.ErrRepositoryNotFound` 映射行（哨兵删除口径见下） |
| `backend/internal/connecterrors/connect_errors.go` | L102 | 删 `governanceprofile.ErrInvalidInput` 映射行（哨兵在删除清单内） |
| `backend/internal/connecterrors/connect_errors.go` | L138 | 删 `governanceprofile.ErrGovernanceProfileSaveFailed` 映射行（哨兵在删除清单内） |

**errors.go 哨兵单值口径（独立复核 FAIL-2/OBS-3 补齐）**：保留 `ErrGovernanceProfileNotFound` / `ErrGovernanceProfileReadFailed`；删除 `ErrInvalidInput` / `ErrGovernanceProfileSaveFailed` / `ErrRepositoryNotFound`。其中 `ErrRepositoryNotFound` 现由 `candidate/repository_reader.go` L39 返回，随 candidate 删除后模块内无返回方（`projectcontext` 侧同名哨兵为独立定义，不受影响），故一并删除并联动 `connect_errors.go` L67。

#### Scenario: 收缩实现判定

- **WHEN** phase14-09 执行收缩
- **THEN** 模块内 `grep -r "SaveProfile\|CommandService\|connect"` 零命中；`go build ./...` 零错误；brief 集成测试经新接口通过（`BriefGovernanceProfile` 三字段断言）

### Requirement: 前端切片移除清单必须冻结

`frontend/src/features/governance-profile/` 整目录删除（10 文件，现状实证）：

- `application/use-update-governance-profile.ts`
- `data/connect-client.ts` / `data/governance-profile-baseline.ts` / `data/use-governance-profile-read.ts`
- `components/governance-profile-form.tsx` / `components/governance-profile-overview.tsx` / `components/governance-profile-readonly-summary.tsx` / `components/governance-profile-section.tsx`
- `types.ts` / `index.ts`
- 挂载点：`frontend/src/features/repository-binding/pages/repository-binding-detail-page.tsx` L27（import）与 L305-306（渲染）——**phase14-08 已同位替换为 StandardReadonlySummary**（phase14-05 让位设计），phase14-09 仅执行目录删除与残留引用清扫

时机与断言：phase14-08 完成让位后，phase14-09 删除目录；断言 = 切片目录不存在 + `grep -r "governance-profile" frontend/src` 零命中 + `tsc --noEmit` 零错误（T5）。

TS 生成产物清理：`frontend/src/gen/proto/psco/governance_profile/` 随 `make gen` 重生成自动消失（proto 源删除后），phase14-09 验证其不存在。

### Requirement: brief 装配切换与 buf breaking 豁免必须冻结

**brief 装配切换**（T6，执行序在 phase14-09）：

1. `project_context.proto`：按 phase14-02 对照表执行——`governance_profile = 2` 类型改内联 `BriefGovernanceProfile`（02 已冻结 proto 草案，逐字引用不重复）、`global_assets = 3` 移除并 `reserved 3;`、`current_phase = 4` 的 status 改内联 `BriefPhaseStatus`、移除 `psco/governance_profile` import（L39）
2. 后端装配（`projectcontext/service/query_service.go`）：画像三字段改经收缩后 `GovernanceProfileReader.ReadProfileCore`；两清单信息改经 `StandardReader.ListStandardsByRepository`（phase14-07 已在过渡态接入 `standards = 8`）；删除旧 `GlobalAssetBinding` 装配代码；**同步收缩同包联动文件**（独立复核 OBS-4 补齐）——`projectcontext/types.go` domain 字段（`GovernanceProfile *governanceprofile.GovernanceProfileReadResult` / `GlobalAssets []governanceprofile.GlobalAssetBinding` 随旧类型删除替换为内联轻量类型）与 `connect/server_integration_test.go` 测试数据（`governanceprofileservice.NewQueryService` 真实构造、`governanceprofile.GlobalAssetBinding` 字面量与 gen 包引用替换）
3. 前端：`gen/proto` 重生成；brief 消费侧（phase13 遗留只读展示，如有）切换内联类型
4. 断言：`project_context.proto` 无 `governance_profile` import；Go/TS 编译零错误；brief 字段面 = phase14-02 对照表"后"列（8 顶层块）

**buf breaking 豁免配置**（T1 断言配套）：

`proto/buf.yaml` 的 `breaking` 段追加 `ignore`（v2 配置，路径相对 module 根）：

```yaml
breaking:
  use:
    - WIRE_JSON
  ignore:
    # phase14-09（T1）：治理画像包整体退役，豁免其删除的 breaking 检测。
    # 决策留痕：phase14_06 spec（画像退役六触点 T1）；phase14-02 冻结退役边界。
    # 该项长期保留：豁免决策可追溯；变更合入 main 形成新基准后 ignore 无实际作用但无害。
    - psco/governance_profile/v1/governance_profile.proto
```

约束：豁免范围**仅限画像包单文件路径**（不得目录前缀扩大化，避免误豁免未来 `psco/governance_profile` 命名空间下的新包）；`make breaking` 在配置落地后对 main 基准全绿。

### Requirement: 退役六触点验收断言矩阵必须执行化

phase14-02 六触点断言 SHALL 逐项落为机械可执行检查（phase14-09 自验 + phase14-10 复验共用）：

| 触点 | 断言 | 机械检查 |
|---|---|---|
| T1 proto 包 | 源与三端产物无残留 | `find proto backend/internal/gen frontend/src/gen -path '*governance_profile*'` 零输出；`make breaking` 通过（豁免生效） |
| T2 存储表 | 两 bindings 表不存在、主表保留 | psql：`SELECT to_regclass('governance_canonical_root_file_bindings'), to_regclass('governance_global_asset_bindings')` 均 NULL；`SELECT count(*) FROM governance_profiles` 与迁移前基线一致（0011 不触碰主表，列结构不变由脚本内容保证——独立复核 OBS-5 修正：去掉非机械占位表述） |
| T3 后端模块 | 无 connect / 无写方法 / 仅 Reader | `grep -r "CommandService\|SaveProfile\|connect/" backend/internal/governanceprofile` 零命中；模块对外导出仅 Reader 实现 |
| T4 画像 RPC | 路由表无注册 | `grep -rn "GovernanceProfileService\|governance_profile.v1" backend/internal/platform` 零命中；对 `/api/psco.governance_profile.v1/...` 的 Connect 请求返回 404 |
| T5 前端切片 | 目录不存在 / 无引用 / 编译过 | `test ! -d frontend/src/features/governance-profile`；`grep -r "governance-profile" frontend/src` 零命中；`tsc --noEmit` 零错误 |
| T6 brief 装配 | 无跨包 import / 字段面对照 | `grep "governance_profile" proto/psco/project_context/v1/project_context.proto` 零命中（注释除外）；`buf build` 过；brief 集成测试字段面 = 02 对照表"后"列 |

迁移专属断言（phase14-10 验收门禁第 ⑧ 条配套）：

- `SELECT count(*) FROM standards WHERE name = '默认项目范式（迁移自治理画像）'` = 1（仅一条合并 Standard）
- 首条 revision 的 `change_summary` 含计数与源 repository id（合并来源留痕）
- 零丢失对照：迁移前导出的两表全量行（5 字段映射集）逐项在树中找到保真承接（phase14-02 零丢失 Scenario）

## MODIFIED Requirements

无（本 spec 为新增设计规格；对 phase14-02 六触点与 brief 对照表仅为执行层细化，对 phase14-03 迁移映射仅为算法落地设计，均不修改上游冻结口径）。

## REMOVED Requirements

无（画像能力的正式退役发生在 phase14-09 实现；本 spec 仅冻结执行设计）。
