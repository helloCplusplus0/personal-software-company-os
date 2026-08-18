# phase14-02 冻结 Standard 与四实体主线、画像退役的关系边界 Spec

## Why

phase14-01 已冻结整体范围，但 Standard 与四实体的绑定关系（target_type × role 组合语义）、画像系统性退役的六触点执行层定位（文件级）、`GetProjectBrief` 字段演进前后对照（含内联轻量消息字段清单）尚未单值化。本子任务把这三组关系边界冻结为执行层确认，作为 phase14-04/06/07/09 设计与实现的强制输入，防止实现时对绑定组合合法性、brief 字段去留、跨包引用内联化范围产生二次解释。本子任务为边界收敛类，不写任何实现代码。

## What Changes

- 产出本 spec 三件套（`phase14_02_freeze_standard_entity_relation_boundaries`）
- 冻结多态绑定关系矩阵：`target_type` 4 枚举 × `role` 2 枚举的组合语义与合法性（8 格矩阵），target 存在性校验规则，唯一约束
- 冻结画像退役六触点的执行层确认：每触点附文件级现状定位（基于本轮实际探查）、退役动作、验收断言
- 冻结 `GetProjectBrief` 字段演进前后对照表（7 字段 → 8 字段，含 1 处 **BREAKING** 移除与字段号 reserved 规则）
- 冻结 `BriefGovernanceProfile` 内联轻量消息字段清单与 `current_phase` 三字段保留口径（phase15 时间轴进入条件）
- 不修改 phase14 三件套正文、不修改根级文档、不写实现代码

## Impact

- Affected specs:
  - `docs/phase/phase14_standard_entity_foundation_dev_plan.md`（上游，phase14-02 定义 L22-25）
  - `docs/phase/phase14_standard_entity_foundation_architecture_plan.md`（§4.4 数据模型 / §4.7 退役映射 / §4.8 brief 演进）
  - `docs/phase/phase14_standard_entity_foundation_shared_baseline.md`（§3.4 绑定矩阵 / §3.5 退役矩阵 / §3.6 简报演进矩阵）
  - `.trae/specs/phase14_01_freeze_standard_entity_scope_success_non_goals/spec.md`（前一子任务，边界总纲）
- Affected code: 无（本子任务零代码改动；本 spec 的事实定位以下列现存文件为对象，供后续子任务执行）
  - `proto/psco/governance_profile/v1/governance_profile.proto`（退役对象 T1）
  - `database/migrations/0010_phase13_governance_profile.sql`（退役对象 T2）
  - `backend/internal/governanceprofile/`（收缩对象 T3）
  - `frontend/src/features/governance-profile/`（移除对象 T5）
  - `proto/psco/project_context/v1/project_context.proto`（演进对象 T6）
  - `backend/internal/projectcontext/`（装配切换对象 T6）

## ADDED Requirements

### Requirement: 多态绑定关系矩阵必须单值冻结

`standard_bindings` 的组合语义 SHALL 冻结为以下 8 格矩阵：

| target_type \ role | `template_source`（规范有实际仓库维护） | `adopts`（目标遵守此规范） |
|---|---|---|
| `repository` | 合法：该仓库是此规范的示范模板仓库（CON-03 复用 Repository） | 合法：该仓库遵守此规范 |
| `product` | **非法**：product 不是仓库，不能承载模板正文事实源 | 合法：该产品遵守此规范 |
| `decision` | **非法**：同上 | 合法：该决策在此规范背景下做出 / 关联 |
| `module` | **非法**：同上 | 合法：该模块遵守此规范 |

约束说明：

- `template_source` 仅对 `repository` 合法是 role 的固有语义约束（指向正文的仓库事实源），不违反 CON-02"关联不设限"（CON-02 约束的是 target 类型广度：`adopts` 对四类全部开放）
- `target_type` 与 `role` 枚举均设计为可扩展（新增值只需扩 enum 并在本矩阵登记组合合法性，不加表）
- target 存在性校验：`BindStandard` 在应用层按 `target_type` 查对应实体表（repositories / products / decisions / modules），目标不存在 → `invalid_argument`
- 唯一约束：`(standard_id, target_type, target_id, role)`，重复绑定拒绝
- brief 装配规则：`GetProjectBrief` 按 `repository_id` 经 `standard_bindings`（任意 role）反查关联 Standard

#### Scenario: 绑定组合合法性判定

- **WHEN** phase14-04 设计 `BindStandard` 校验规则或 phase14-07 实现
- **THEN** 8 格矩阵为单值判据；`template_source` 携带非 repository 目标一律 `invalid_argument`

#### Scenario: 绑定矩阵扩展

- **WHEN** 未来需要新增 target 类型（如 venture）或 role
- **THEN** 仅扩 enum 值 + 在本矩阵登记组合合法性，不加分表、不改唯一约束结构

### Requirement: 画像退役六触点必须执行层确认（文件级）

六触点 SHALL 以本轮实际探查的现状为基准冻结（现状定位精确到文件与行号，作为 phase14-09 的执行清单）：

**T1 proto 包整体退役**

- 现状：`proto/psco/governance_profile/v1/governance_profile.proto`（190 行：`TrackType` / `PhaseStatus` 枚举、`CanonicalRootFileBinding` / `GlobalAssetBinding` / `GovernanceProfile` 消息、`GetGovernanceProfile` / `UpdateGovernanceProfile` RPC、`GovernanceProfileService` 服务定义）
- 动作：删除整个 `proto/psco/governance_profile/` 目录；生成产物（`backend/internal/gen/proto|connect/psco/governance_profile/`、`frontend/src/gen/proto/psco/governance_profile/`）随之清理并重生成
- 断言：proto 源与三端生成产物均无该包残留；buf breaking 对该包删除显式豁免并留痕（phase14-06 设计豁免配置）

**T2 存储表退役（主表保留）**

- 现状：`database/migrations/0010_phase13_governance_profile.sql` 三表——`governance_profiles` 主表（L21，`repository_id` UNIQUE 锚点）、`governance_canonical_root_file_bindings`（L47）、`governance_global_asset_bindings`（L69）
- 动作：`0011` 迁移中先迁数据（两张 bindings 表数据 → 一条全局 Standard 的 `directory_tree`，裁决⑧）后 `DROP TABLE` 两张 bindings 表；**`governance_profiles` 主表保留**（brief 三字段数据源，列不增不减）
- 断言：两张 bindings 表不存在；`governance_profiles` 行数与列结构与迁移前一致

**T3 后端模块收缩为纯读**

- 现状：`backend/internal/governanceprofile/` 含 `candidate/`、`connect/`、`repository/` 子包与 `errors.go`、`types.go`
- 动作：删除 `connect/`（RPC handler 层）与全部写路径；保留最小读取能力（`GovernanceProfileReader`：读主表三组字段服务 brief）；`candidate/` 若仅为画像写入候选服务则一并删除
- 断言：模块内无 connect handler、无写方法；对外仅暴露 Reader 读取接口

**T4 画像 RPC 退役**

- 现状：`GovernanceProfileService` 的 `GetGovernanceProfile` / `UpdateGovernanceProfile`，经路由注册暴露于 `/api/psco.governance_profile.v1/...`
- 动作：随 T1/T3 执行移除路由注册与 handler
- 断言：路由表无画像服务注册；对退役路径的请求不再命中任何 handler

**T5 前端切片整体移除**

- 现状：`frontend/src/features/governance-profile/` 共 10 文件（`application/use-update-governance-profile.ts`；`data/connect-client.ts` / `data/governance-profile-baseline.ts` / `data/use-governance-profile-read.ts`；`components/` 4 组件；`types.ts`；`index.ts`）+ 挂载点 `repository-binding-detail-page.tsx` L27（import）与 L305（画像区渲染）
- 动作：删除整个切片目录；detail 页移除 import 与画像区渲染，让位给 Standard 只读摘要入口（phase14-08 承接）
- 断言：切片目录不存在；detail 页无 `governance` 引用；`tsc --noEmit` 零错误

**T6 brief 装配切换（跨包引用内联化）**

- 现状：`project_context.proto` 3 处跨包引用——L39（import 语句）、L190（`governance_profile = 2` 引用 `GovernanceProfile`）、L191（`global_assets = 3` 引用 `GlobalAssetBinding`）、L165（`BriefCurrentPhase.status` 引用 `PhaseStatus`）；后端装配在 `backend/internal/projectcontext/service/`
- 动作：内联轻量消息（见下一 Requirement）替代跨包引用；装配改读 `StandardReader`（两清单信息）+ `GovernanceProfileReader`（主表三字段）；移除 import
- 断言：`project_context.proto` 无 `governance_profile` 包 import；Go/TS 两端编译零错误；brief 语义对照验收通过（本 spec 对照表）

#### Scenario: 退役完整性验收

- **WHEN** phase14-09 执行完成
- **THEN** 六触点断言逐项全绿，任一触点残留即退役不完整、验收不通过

### Requirement: GetProjectBrief 字段演进前后对照表必须留档

`GetProjectBriefResponse` SHALL 按下表演进（前：phase13-10 冻结的 7 顶层字段；后：phase14 冻结的 8 顶层字段）：

| # | 字段 | 前（现状） | 后（phase14） | 说明 |
|---|---|---|---|---|
| 1 | `repository = 1` | `RepositorySummary` | 不变 | |
| 2 | `governance_profile = 2` | 跨包 `psco.governance_profile.v1.GovernanceProfile` | **内联 `BriefGovernanceProfile`**（轻量，字段清单见下一 Requirement） | 类型替换，语义收缩 |
| 3 | `global_assets = 3` | 跨包 `repeated GlobalAssetBinding` | **BREAKING 移除**；字段号 3 `reserved` | 信息迁入 `standards[]` 的树节点（`role/summary/ref` 保真映射） |
| 4 | `current_phase = 4` | `BriefCurrentPhase`（status 引用跨包 `PhaseStatus`） | 保留；status 改用**内联 `BriefPhaseStatus` 枚举**（值域不变） | 三字段数据源仍为主表 |
| 5-7 | `products = 5` / `modules = 6` / `decisions = 7` | 摘要数组 | 不变 | |
| 8 | `standards = 8` | 无 | **新增** `repeated StandardSummary`（含 `directory_tree` 全树与导航摘要） | 该 repository 经绑定关联的 Standard 列表 |

约束：

- 移除字段号必须 `reserved 3;`（防止字段号复用造成旧消费端语义错乱）
- 不新增与 8 字段并列的第 9 个顶层块（延续 phase13-07 约束）
- brief 装配信息源单值：两清单信息唯一来自 `StandardReader`；主表三字段唯一来自 `GovernanceProfileReader`

#### Scenario: 零丢失对照验收

- **WHEN** phase14-10 执行 brief 对照验收
- **THEN** 以 phase13-11 固定样本的 brief 前值为基准：`global_assets[]` 每一项（`name/kind/entry_ref/role/structured_summary`）在迁移后 `standards[].directory_tree` 中有保真承接（`name`→节点名、`entry_ref`→`ref`、`role`→`role`、`structured_summary`→`summary`、`kind` 并入摘要语义）；`canonical_root_files[]` 同理；无信息静默丢失

### Requirement: BriefGovernanceProfile 内联轻量消息与 current_phase 保留口径必须冻结

`project_context.proto` 内联消息 SHALL 冻结为：

```proto
// BriefTrackType 内联技术路线枚举（值域与退役的 TrackType 一致）
enum BriefTrackType { BRIEF_TRACK_TYPE_UNSPECIFIED = 0; BRIEF_TRACK_TYPE_PRODUCT = 1; BRIEF_TRACK_TYPE_DURABLE_SYSTEM = 2; }
// BriefPhaseStatus 内联阶段状态枚举（值域与退役的 PhaseStatus 一致）
enum BriefPhaseStatus { BRIEF_PHASE_STATUS_UNSPECIFIED = 0; BRIEF_PHASE_STATUS_PLANNED = 1; BRIEF_PHASE_STATUS_IN_PROGRESS = 2; BRIEF_PHASE_STATUS_COMPLETED = 3; BRIEF_PHASE_STATUS_BLOCKED = 4; }
// BriefGovernanceProfile 画像主记录轻量内联消息
message BriefGovernanceProfile {
  string repository_id = 1;            // 锚点保留
  BriefTrackType track_type = 2;       // read-only
  optional string template_source = 3; // 手工维护字段保留
}
```

显式不迁移字段（随 GovernanceProfile 消息退役）：

- `project_profile_version`（GAP-03：版本字段无演进基础，退役）
- `docs_workflow_layout`（phase13 已 deprecated 的死字段，退役）
- `canonical_root_files` / `global_asset_bindings`（T2/T6 迁移至 Standard）
- `GlobalAssetBinding.markdown_resolvable`（服务端按资产矩阵计算的只读能力状态，非用户维护信息；其"正文可否按 ref 回源"语义被裁决③结构性吸收——ref 一律可回源到仓库事实源，故显式退役，不进入零丢失验收范围）
- `current_phase_name/ref/status`（不在此消息内重复：顶层 `current_phase = 4` 块是唯一承载位，从主表直接装配，消除原消息与顶层块的重复）
- `created_at / updated_at`（brief 无消费场景，不迁移）

`current_phase` 三字段保留口径 SHALL 冻结为：

1. 数据层：`governance_profiles` 主表三列保留不动（T2）
2. 读取层：仅经 `GetProjectBrief.current_phase`（`BriefCurrentPhase`：`name / entry_ref / status`）只读装配，status 类型改内联 `BriefPhaseStatus`
3. 写入层：无任何写入承接（延续 phase13 read-only 冻结）
4. 演进层：phase14 不新增时间轴 / 历史承接位；`CON-08` 时间轴承接为 phase15 进入条件，在 phase14-11 收口时正式回写

#### Scenario: 内联消息实现判定

- **WHEN** phase14-07/09 修改 `project_context.proto`
- **THEN** 内联消息字段与本清单逐字一致；不迁移字段清单中的字段不得以任何形式复活

## MODIFIED Requirements

无（本 spec 为新增关系边界冻结规格，不修改既有规格正文）。

## REMOVED Requirements

无（画像相关能力的正式退役发生在 phase14-09 实现，本 spec 仅冻结其执行层映射，不在此执行移除）。
