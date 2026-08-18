# phase14-01 冻结 `Standard Entity Foundation` 范围边界、成功标准与非目标 Spec

## Why

`phase14` 三件套已冻结八项裁决结论与整体架构方向，但 `phase14` 共 11 个子任务在进入各自 `/spec` 前，需要一个单值的执行层边界声明入口：把八项裁决 + `phase13-12` 缺口记录（GAP-01~07 / CON-01~09）收敛为"做什么、成功长什么样、不做什么"的冻结结论，防止后续子任务在执行中出现范围漂移或对裁决口径的二次解释。本子任务为边界收敛类，不写任何实现代码。

## What Changes

- 产出本 spec 三件套（`phase14_01_freeze_standard_entity_scope_success_non_goals`），作为 `phase14` 执行层第一个正式规格入口
- 冻结单一主交付边界：`Standard` 全局规范实体最小主线（合同 → 存储 → 后端 → 前端 → agent 消费）+ 治理画像设计系统性退役（六触点）
- 冻结成功标准（4 条）与非目标（7 条）
- 冻结八项裁决的边界声明（每项裁决的"允许 / 禁止"二值化表述）
- 留档 GAP-01~07 / CON-01~09 覆盖矩阵（每条在本 phase 的承接落点）
- 不修改 phase14 三件套正文、不修改根级文档、不写实现代码

## Impact

- Affected specs:
  - `docs/phase/phase14_standard_entity_foundation_architecture_plan.md`（上游，本 spec 单值引用其 §2/§3/§4 冻结结论）
  - `docs/phase/phase14_standard_entity_foundation_dev_plan.md`（上游，phase14-01 定义 L17-20）
  - `docs/phase/phase14_standard_entity_foundation_shared_baseline.md`（上游，§2.2 八项裁决 / §2.4 特别约束为最高优先级基线）
  - `.trae/specs/phase13_12_sync_root_level_closeout_next_phase_entry_conditions/spec.md`（缺口记录上游）
- Affected code: 无（边界收敛类子任务，无代码改动）

## ADDED Requirements

### Requirement: 单一主交付边界必须冻结

`phase14` 的交付范围 SHALL 冻结为一条主线两个组成部分，不并列第二主交付：

1. **`Standard` 全局规范实体最小主线**：合同（`proto/psco/standard/v1/standard.proto`，8 RPC）→ 存储（`0011` 迁移三表：`standards` / `standard_revisions` / `standard_bindings`）→ 后端（`backend/internal/standard/` 分层）→ 前端（`frontend/src/features/standard/` 切片 + `/standards` 路由）→ agent 消费（`GetProjectBrief` 装配 `standards[]`）
2. **治理画像设计系统性退役**（六触点）：proto 包整体退役 / 两张 bindings 表 drop（先迁数据）/ 后端模块收缩纯读 / 前端切片整体移除 / 无画像 RPC 调用 / brief 改内联轻量消息

#### Scenario: 后续子任务范围判定

- **WHEN** phase14-02 ~ phase14-10 任一子任务需要判定某项工作是否属于 phase14
- **THEN** 以本 Requirement 的两个组成部分为单值判据
- **AND** 超出两者的诉求一律落入非目标或 phase15 进入条件讨论

### Requirement: 成功标准必须冻结为可验证的 4 条

`phase14` 的成功标准 SHALL 冻结为（与 architecture_plan §3 一致）：

1. 用户可在 web 端创建并维护一份 `Standard`（结构化树编辑器维护目录树 + 必须文档节点 + 摘要），并从 Standard 详情页绑定模板仓库 `Repository` 与遵守该规范的 `Product`
2. agent 经 `GetProjectBrief` 可直接读到该 `Standard` 的树形目录结构与文档导航（含摘要与 `ref`），无需转译、无需回源即可定位每份文档的用途
3. 治理画像设计六触点已系统性退役，存量数据按裁决⑧合并迁移至一条全局 Standard，旧信息经 `standards[]` 零丢失
4. 同一套规范只需维护一次，任意数量产品 / 仓库通过绑定消费（全局性达成）

#### Scenario: phase14-10 验收入口

- **WHEN** phase14-10 联调验收执行
- **THEN** 4 条成功标准逐条有取证证据，任一条不可演示即验收不通过

### Requirement: 非目标必须冻结为 7 条且不可偷渡

以下 SHALL 冻结为 `phase14` 非目标（与 dev_plan §4 一致）：

1. `CON-08` 阶段推进时间轴（`current_phase` 三字段仅以 brief 内联轻量消息保留，不新增承接位、不改动数据；时间轴归 phase15 进入条件）
2. agent 写回、MCP / CLI 消费通道
3. Git 推进跟踪、模板仓库自动接入、自动同步
4. `Standard` 的版本分支、全文快照、正文托管
5. 第五 CRUD 主实体化（Dashboard 主卡片、四实体关系主链接入）
6. 四实体既有页面的无关重构（Repository detail 仅触达画像让位局部）
7. 画像主记录三字段（`track_type / template_source / current_phase`）的进一步演进（保留即止）

#### Scenario: 非目标偷渡拦截

- **WHEN** 任一子任务的 spec 或实现出现非目标范围内的交付物
- **THEN** 该交付物构成范围违约，必须移除或显式升级为新的 phase 级裁决后才允许存在

### Requirement: 八项裁决边界声明必须二值化冻结

八项裁决（shared_baseline §2.2）SHALL 以"允许 / 禁止"二值化表述冻结为执行层边界：

| 裁决 | 允许 | 禁止 |
|---|---|---|
| ① 颗粒度 | PSCO 承接结构、导航、结构化摘要（`summary`）与定位引用（`ref`） | PSCO 存储 / 复制 / 缓存任何文档正文；正文以模板仓库为唯一事实源 |
| ② 数据模型 | 主表 + jsonb 整树 + 多态绑定表（一张 `standard_bindings`） | 全范式化树节点表；单表全 jsonb；按目标类型分表 |
| ③ 内容边界 | PSCO=结构导航与演进记录；仓库=正文事实源 | 双份维护正文；仓库承载 PSCO 结构定义文件 |
| ④ 画像退役 | phase14 内六触点系统性退役；三字段以内联消息保留于 brief | 画像承接位以 deprecated 双轨保留；本阶段给 `current_phase` 新增承接位 |
| ⑤ 实体地位 | 全局规范资产实体：全局作用域、独立 `/standards` 入口、绑定关联四实体 | `repository_id` 锚点；第五 CRUD 主实体；Dashboard 主卡片 |
| ⑥ 树编辑形态 | 结构化树编辑器（缩进列表 + 节点表单化编辑 + 增删 / 上移下移 / 添加子节点） | 拖拽交互；裸 JSON 文本域编辑；节点级增量保存协议 |
| ⑦ 绑定 UI | 仅 Standard 详情页内发起（target_type 选择 + 目标检索 + role） | 目标实体页发起绑定；双侧入口 |
| ⑧ 迁移合并 | 全部存量合并为一条全局 Standard；不一致取最新 `updated_at`；差异记入首条 revision | 按 repository 拆分 N 条 Standard；静默丢弃不一致数据 |

#### Scenario: 裁决口径争议仲裁

- **WHEN** 任一后续子任务对裁决含义产生解释分歧
- **THEN** 以本二值化表 + shared_baseline §2.2 为准，不得二次解释或放宽
- **AND** 确需放宽时必须回到用户显式补裁，不得默认绕过

### Requirement: GAP/CON 覆盖矩阵必须留档

本 spec SHALL 留档 GAP-01~07 / CON-01~09 到 phase14 承接落点的完整映射矩阵：

**GAP 修正矩阵**：

| GAP | phase14 承接落点 |
|---|---|
| GAP-01 目录结构表示不成立 | `DirectoryTreeNode` 递归 jsonb 树（单根 directory、8 条校验规则）；proto 递归消息单值映射（裁决②，arch §4.5） |
| GAP-02 双清单高度重复 | 双清单合一：必须文档即树的 `file` 节点（`role + summary + ref`），禁止第二份清单（shared_baseline §2.4） |
| GAP-03 版本演进无基础 | `standard_revisions` 轻量留痕（`change_summary` 一句话 + 时间戳，无全文快照）（裁决②，arch §4.4） |
| GAP-04 阶段字段错位 | `current_phase` 三字段保留于 brief 内联消息不扩展；时间轴承接为 phase15 进入条件（裁决④） |
| GAP-05 全局性矛盾 | 全局作用域无 repository 锚点 + 多态绑定关系 + 迁移合并为一条全局 Standard（裁决⑤⑧） |
| GAP-06 agent 原生友好度 | 树形 proto 消息 + brief `standards[]`（全树 + 导航摘要 + `ref`）无转译直读（arch §4.8） |
| GAP-07 总体耦合判定 | 画像设计六触点系统性退役（裁决④ + 用户强化要求，arch §4.7） |

**CON 承接矩阵**：

| CON | phase14 承接方式 |
|---|---|
| CON-01 全局规范实体方向 | 直接实现（`Standard` 实体主线） |
| CON-02 绑定关系不设限 | `standard_bindings` 多态绑定（target_type 枚举可扩展，起步 4 类） |
| CON-03 模板仓库复用 Repository | 模板仓库作为普通 Repository 实例 + `role=template_source` 绑定 |
| CON-04 取代不并存 | 裁决④：phase14 内六触点完整退役，无双轨兼容期 |
| CON-05 颗粒度裁决 | 裁决①：混合式（结构 + 导航 + 摘要在 PSCO，正文在仓库） |
| CON-06 动态演进一等需求 | 整树原子替换 + revision 每次更新追加留痕 |
| CON-07 反假大空效率底线 | 全程约束（shared_baseline §2.4）；设计方案成本高于手动维护基线即回退 |
| CON-08 阶段跟踪独立时间轴 | **非目标**，冻结为 phase15 进入条件（phase14-11 收口时正式回写） |
| CON-09 先消费后维护 | agent 写回为非目标；web 写 + agent 只读 |

#### Scenario: 覆盖完整性核对

- **WHEN** 独立复核执行
- **THEN** 16 条（7 GAP + 9 CON）逐条有落点，无"悬空"条目（既无承接又未显式列入非目标）

## MODIFIED Requirements

无（本 spec 为新增边界收敛规格，不修改既有规格）。

## REMOVED Requirements

无。
