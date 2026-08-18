# phase14_standard_entity_foundation_architecture_plan

## 1. 文档定位

- 本文档是 `phase14_standard_entity_foundation` 的架构规划与冻结记录，承担该阶段 `/plan` 的架构结论冻结职责。
- 本文档在阶段收口后保留为规划与冻结记录，不覆盖根级当前状态；根级状态只以 `plan.md` 为准。
- 本文档的直接下游是同阶段 `dev_plan` 与 `shared_baseline`；三件套共同构成 `phase14` 全部 `/spec` 子任务的强制边界上游。

## 2. 上游输入

- 唯一直接规划上游：`.trae/specs/phase13_12_sync_root_level_closeout_next_phase_entry_conditions/spec.md`（`phase13` 正式缺口记录：GAP-01 ~ GAP-07 + CON-01 ~ CON-09 + 五项裁决进入条件）
- 八项裁决已由用户于 `2026-08-17` 结构化逐项拍板完成（前五项 `/plan` 主裁决 + 后三项可执行性补裁）：
  - 裁决①（信息维护颗粒度，对应 CON-05）：混合式——PSCO 承接结构 + 导航 + 结构化摘要，全文以模板仓库为唯一事实源
  - 裁决②（数据模型与 pg 承载）：主表 + jsonb 树形目录结构 + 多态绑定关系表
  - 裁决③（与模板仓库内容边界，对应 CON-05/CON-07）：PSCO=结构导航与演进记录，仓库=正文事实源
  - 裁决④（画像退役计划，对应 CON-04）：phase14 内完整退役重叠承接位；`current_phase` 三字段随画像主记录保留在 brief，时间轴由 `phase15` 承接；用户强化要求：`phase13` 画像设计系统性移除（proto 包 / RPC / 前端切片 / 后端写路径整体退役）
  - 裁决⑤（实体地位）：全局规范资产实体——治理层地位、全局作用域、独立维护入口，不做第五 CRUD 主实体
  - 裁决⑥（整树编辑交互形态）：结构化树编辑器——树形缩进列表 + 节点表单化编辑 + 增删 / 上移下移，无拖拽
  - 裁决⑦（绑定 UI 承接位）：Standard 详情页内发起绑定，维护动作唯一入口 `/standards`
  - 裁决⑧（存量迁移合并策略）：全部存量画像数据合并为一条全局 Standard，不一致字段取最新 `updated_at`，多源差异记入首条 revision
- 参照基线：`PSCO-mvp05-summarize-feedback.md`（mvp0.5 共识）、`TECH_STACK_BASELINE.md`（技术主线）、`project_rules.md` §2.6（传输合同约束）、四实体既有建模模式（`proto/psco/<entity>/v1` + `backend/internal/<module>` 分层 + `frontend/src/features/<slice>`）

## 3. 本阶段目标

- 单一主交付：建立 `Standard` 全局规范实体最小主线——跨项目的理想目录结构与必须全局文档以全局实体承接，web 端可维护、agent 经 go 后端可直读；同步完成治理画像重叠承接位在 phase14 内的完整退役。
- 成功标准：
  1. 用户可在 web 端创建并维护一份 `Standard`（目录树 + 必须文档清单 + 摘要），并将其绑定到模板仓库 `Repository` 与遵守该规范的 `Product`
  2. agent 经 `GetProjectBrief` 可直接读到该 `Standard` 的树形目录结构与文档导航（含摘要），无需转译、无需回源即可定位每份文档的用途
  3. 治理画像的 `canonical_root_files[]` / `global_asset_bindings[]` 承接位（proto 字段、存储表、前端编辑位）已完整退役，存量数据已迁移至 `Standard`
  4. 同一套规范只需维护一次，任意数量产品 / 仓库通过绑定消费（解决 GAP-05 全局性矛盾）

## 4. 架构冻结结论

### 4.1 当前阶段唯一直接执行层上游

- 规划上游：`phase13-12` 正式缺口记录（唯一直接上游）
- 执行层上游：本三件套冻结后，`phase14` 全部 `/spec` 子任务以三件套为强制边界上游
- 既有不可回退项继续有效：四实体语义、`.proto + ConnectRPC` 传输主线、canonical owner 单值化、`chi` 仅承担基础设施端点

### 4.2 当前阶段单一主交付能力

- `Standard Global Asset Foundation`：全局规范资产实体的最小完整主线（合同 → 存储 → 后端 → 前端 → agent 消费 → 画像退役），一能力贯通，不并列第二主交付。
- 交付边界内的四个组成部分：
  1. `Standard` 实体合同与存储（新建）
  2. web 端维护入口（新建 `/standards`）
  3. agent 消费链路（`GetProjectBrief` 演进改读 `Standard`）
  4. 治理画像重叠承接位退役（含存量数据迁移）

### 4.3 Standard 实体的正式定位（裁决⑤）

- 正式地位：全局规范资产实体，属治理层资产，不是第五业务主实体。
- 作用域：全局作用域，无 `repository_id` 锚点（解决 GAP-05）；通过多态绑定关系与四实体关联（CON-02 不设限）。
- 交互入口：独立列表 / 维护入口（`/standards` 路由 + 全局导航项），不进入四实体关系主链，不建 Dashboard 主卡片主线。
- 与模板仓库的关系（CON-03）：示范模板仓库本身作为 `Repository` 实例正常登记；`Standard` 经绑定关系（`role=template_source`）指向该 `Repository`，不新造第二套仓库实体。

### 4.4 Standard 数据模型与 pg 承载（裁决②）

存储分三层，落在一张编号迁移文件（`0011_phase14_standard_entity.sql`）：

1. **主表 `standards`**：身份与状态层
   - `id`（uuid PK）、`name`（唯一）、`description`、`status`（draft / active / retired）
   - `directory_tree`（jsonb，整树原子承载，见 4.5）
   - `created_at` / `updated_at`
2. **演进留痕表 `standard_revisions`**（CON-06 动态演进一等需求的最小承接）
   - `id`、`standard_id`（FK）、`change_summary`（text，人工一句话说明本次调整）、`created_at`
   - 每次主表更新（结构 / 文档清单 / 摘要变动）追加一行；不存全文快照（反假大空，正文与结构的当前值以主表为唯一事实源）
3. **多态绑定关系表 `standard_bindings`**（CON-02 关联不设限）
   - `id`、`standard_id`（FK）、`target_type`（enum：`repository` / `product` / `decision` / `module`，可扩展）、`target_id`（uuid，应用层校验目标存在性）、`role`（enum：`template_source` / `adopts`，可扩展）、`note`（可选）
   - 唯一约束：`(standard_id, target_type, target_id, role)`
   - 不为每类目标建分表；新增目标类型只扩 enum，不加表

### 4.5 目录树表示（裁决①+②+⑥，直接解决 GAP-01 / GAP-02）

- jsonb 树以递归节点承载，节点结构统一为：

```json
{
  "name": "docs",
  "node_type": "directory",
  "role": null,
  "summary": "项目推进 workflow 文档",
  "ref": null,
  "children": [
    { "name": "AGENTS.md", "node_type": "file", "role": "entry", "summary": "agent 全局上下文入口", "ref": "/AGENTS.md", "children": [] }
  ]
}
```

- 必须文档即树中的 `file` 节点，不再单列第二份清单（GAP-02 双清单重复的修正：目录结构与必须文档清单合一，`role` + `summary` + `ref` 就是文档的角色、说明与定位引用）。
- `node_type` 只有 `directory` / `file` 两值；`role` 为短标签（约定值域 `entry / plan / rules / baseline / spec / summary`，允许自定义）；`summary` 为一段结构化摘要（裁决①混合式颗粒度的承接位）；`ref` 为可选定位引用（以 `/` 开头的仓库内相对路径或 `https://` 开头的 URL，承接旧 `entry_ref / external_url` 语义——`ref` 是导航引用而非正文，不违反裁决③正文零托管）。
- proto 侧以递归消息 `DirectoryTreeNode` 单值映射，agent 原生可读树形结构，无转译成本。
- 树整体读写、原子替换（编辑保存 = 整树校验后替换 `directory_tree` + 追加 revision），不做节点级增量协议。
- 完整校验规则（8 条：单根 directory / 空树仅 draft 合法 / 同层 name 唯一 / name 字符集 `[A-Za-z0-9._-]` 上限 64 / 深度上限 6 / 序列化上限 64KB / file 无 children 且 role 必填 / ref 格式约束）以 `shared_baseline` §3.3 为单值来源。

### 4.6 内容边界（裁决③）

- PSCO 承接：规范身份、目录结构声明、必须文档清单与摘要、定位引用（`ref`）、绑定关系、演进留痕（revision）。
- 模板仓库承接：文档正文、可执行示例。仓库是正文的唯一事实源；PSCO 不复制、不缓存正文。
- `file` 节点的 `ref` 仅是导航引用（路径 / URL），不含正文内容；agent 需要正文时按 `ref` 到模板仓库读取。

### 4.7 治理画像系统性退役映射（裁决④ + 用户强化要求，六触点）

| 退役对象 | 当前位置 | 退役后去向 |
|---|---|---|
| `governance_profile` proto 包（全部消息 + `GetGovernanceProfile / UpdateGovernanceProfile` RPC） | `proto/psco/governance_profile/v1/` | 整体退役删除；brief 侧类型改内联（见 4.8） |
| `canonical_root_files[]` / `global_asset_bindings[]` 数据 | `0010` 迁移两张 bindings 表 | 迁移为一条全局 Standard 的 `directory_tree` 节点（裁决⑧合并；`role / summary` 保真，`entry_ref / external_url` → `ref`） |
| `governance_canonical_root_file_bindings` / `governance_global_asset_bindings` 表 | `0010` 迁移 | drop（先迁数据后 drop）；`governance_profile_record` 主表保留（brief 三字段数据源） |
| 前端画像切片 | `frontend/src/features/governance-profile/` | 整体移除；Repository detail 画像区整体让位，由 Standard 只读摘要入口替代 |
| 后端画像写路径与 connect 层 | `backend/internal/governanceprofile/` | 收缩为纯读模块（仅保留 `GovernanceProfileReader` 最小读取服务 brief 三字段） |
| `GetProjectBrief` 两清单装配 | `projectcontext/query_service.go` | 改读 `StandardReader`；`governance_profile` 字段改内联轻量消息 |

- 保留项：画像主记录三字段（`track_type` / `template_source` / `current_phase`）数据保留于主表，以 `project_context.proto` 内联轻量消息存在于 brief；`current_phase` 的时间轴承接是 `phase15` 进入条件（CON-08），phase14 不新增承接位。
- buf breaking 受控策略：`governance_profile` 包整体删除属破坏性变更，在迁移完成后的删除提交中执行，并在 proto 配置层面显式豁免该已退役包、于 phase14-09 spec 留痕；`standard` / `project_context` 包继续执行全量 breaking 检查。
- 退役验收门禁：proto 目录无画像包残留、旧 bindings 表无数据残留且已删除、前端切片无残留、brief 无画像 RPC 调用、brief 返回值语义不缺失（原信息经 `Standard` 完整可达）。

### 4.8 agent 消费链路（GetProjectBrief 演进）

- `GetProjectBriefResponse` 新增 `standards[]`：该 repository 经 `standard_bindings`（任意 role）关联的 `Standard` 列表（含 `directory_tree` 全树、文档导航摘要与 `ref`）。
- 原 `governance_profile` 字段收缩为 `project_context.proto` 内联轻量消息（`track_type / template_source / current_phase` 三组字段），解除对 `governance_profile.proto` 的跨包依赖；两清单信息唯一来源变为 `standards[]`。
- 读取路径：web 与 agent 同源（`GetProjectBrief` 单值装配），符合 phase12 建立的共享只读消费基线。
- agent 写回不在本阶段（CON-09 先消费后维护不变）。

### 4.9 合同与传输主线（project_rules §2.6 对齐）

- 新合同包：`proto/psco/standard/v1/standard.proto`，package `psco.standard.v1`。
- RPC 最小集（8 个）：`ListStandards`（不分页，个人规模，返回全量；显式冻结避免实现摇摆）/ `GetStandard`（含基础信息与整树）/ `CreateStandard` / `UpdateStandard`（整树原子替换，`change_summary` 必填）/ `BindStandard`（target 不存在 → `invalid_argument` 语义错误）/ `UnbindStandard` / `ListStandardBindings` / `ListStandardRevisions`（前端 revision 回看的读取承接位）；全部走 ConnectRPC。
- web 写路径与 agent 读路径同包不同 RPC，不出现第二套 canonical API。

### 4.10 当前阶段验收协议前提

- 固定样本：沿用 `phase13-11` 冻结的 dogfooding 样本（`repository_id: ca261521-8daf-4248-8f12-43525326e759`），在其上补建 `Standard` 与绑定。
- 固定问题：brief 中 `standards[]` 树形结构与文档导航的可直答性、画像系统性退役后旧信息零丢失、单套规范多产品复用可演示、revision 留痕可回看。
- 工具链门禁：`buf lint / build / breaking`（画像包豁免留痕）、`go build / vet / test`、前端 `tsc --noEmit`、浏览器反回归矩阵（含 `/standards` 新页面与 Repository detail 让位后的回归）。
- 八项裁决逐条可验证（验收门禁清单见 `shared_baseline` §4）。

### 4.11 当前阶段明确不做

- `CON-08` 阶段推进时间轴（phase15 承接，`current_phase` 三字段仅保留不扩展）
- agent 写回 / MCP / CLI 消费通道
- Git 推进跟踪、模板仓库自动接入、自动同步
- `Standard` 的版本分支 / 全文快照 / 正文托管
- 第五 CRUD 主实体化（Dashboard 主卡片、四实体关系主链接入）
- 四实体既有页面与合同的重构（除画像让位触达的 Repository detail 局部）
