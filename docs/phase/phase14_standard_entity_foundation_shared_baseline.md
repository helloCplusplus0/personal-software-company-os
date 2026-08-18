# phase14_standard_entity_foundation_shared_baseline

## 1. 文档定位

- 本文档集中冻结 `phase14_standard_entity_foundation` 的单值基线与能力矩阵，是该阶段全部 `/spec` 子任务的共享参照。
- 三件套不一致时以本文档为准并回改；阶段收口后本文档保留为该阶段冻结记录。

## 2. 当前单值基线

### 2.1 项目路线

- 当前阶段：`phase14_standard_entity_foundation`（`/plan` 三件套已建立，进入子任务执行）
- 直接规划上游：`phase13-12` 正式缺口记录（唯一）
- 单一主交付：`Standard` 全局规范实体最小主线 + 治理画像重叠承接位完整退役
- 下一阶段进入条件在 `phase14-11` 收口时冻结（预期含 `CON-08` 时间轴承接）

### 2.2 八项裁决结论（本阶段最高优先级基线）

前五项为 `/plan` 主裁决（2026-08-17 拍板），后三项为可执行性补裁（同日反思轮拍板）：

| 裁决项 | 结论 | 对应 GAP/CON |
|---|---|---|
| ① 信息维护颗粒度 | 混合式：PSCO 承接结构 + 导航 + 结构化摘要；全文以模板仓库为唯一事实源 | CON-05 / CON-07 |
| ② 数据模型与 pg 承载 | 主表 + jsonb 树形目录结构 + 多态绑定关系表 | GAP-01 / GAP-02 |
| ③ 与模板仓库内容边界 | PSCO=结构导航与演进记录；仓库=正文事实源 | CON-03 / CON-07 |
| ④ 画像退役计划 | phase14 内完整退役；`current_phase` 三字段随画像主记录保留在 brief，时间轴 phase15 承接 | CON-04 / CON-08 |
| ⑤ Standard 实体地位 | 全局规范资产实体：治理层地位、全局作用域、独立维护入口，不做第五 CRUD 主实体 | GAP-05 / GAP-07 |
| ⑥ 整树编辑交互形态 | 结构化树编辑器：树形缩进列表，每节点一行表单化编辑（name / node_type / role / summary / ref），支持节点增删 / 上移下移 / 添加子节点，无拖拽 | 可执行性 |
| ⑦ 绑定 UI 承接位 | Standard 详情页内发起绑定（选择 target_type + 检索目标 + role）；维护动作唯一入口 `/standards` | 可执行性 |
| ⑧ 存量迁移合并策略 | 全部存量画像数据合并为一条全局 Standard；不一致字段以最新 `updated_at` 为准；多源差异记入首条 revision 留痕 | GAP-05 / 裁决⑤ |

### 2.3 当前阶段正式技术主线

- 合同源：`proto/psco/standard/v1/standard.proto`（package `psco.standard.v1`），`.proto` 唯一长期合同
- 传输：ConnectRPC 正式传输；`chi` 仅基础设施端点（project_rules §2.6）
- 后端：Go 单体 `backend/internal/standard/`（connect / service / repository 分层），跨模块读取经 `projectcontext` 侧 `StandardReader` 独立 Read 接口
- 前端：`frontend/src/features/standard/` 切片 + `/standards` 路由；query 纯只读、mutation 收敛切片固定承接位（project_rules §2.5）
- 存储：`0011_phase14_standard_entity.sql`（`standards` + `standard_revisions` + `standard_bindings`）

### 2.4 当前阶段特别约束

- 双清单合一：必须文档即 `directory_tree` 的 `file` 节点（`role` + `summary` + `ref` 承载角色、说明与定位引用），禁止再现第二份清单（GAP-02 修正）
- 树整体原子替换：编辑保存 = 整树校验 + 替换 + 追加 revision；不做节点级增量协议
- 正文零托管：`file` 节点仅含 `ref` 定位引用（仓库内路径或 URL），不含正文内容；PSCO 不复制 / 缓存仓库正文（裁决③；`ref` 是导航引用而非正文，不违反本约束）
- 绑定不设限：`standard_bindings.target_type` 枚举可扩展（repository / product / decision / module 起步），不为目标类型建分表（CON-02）
- 画像系统性移除：`phase13` 画像设计在 phase14 内系统性退役——`proto/psco/governance_profile/v1` 合同包整体退役、`GetGovernanceProfile / UpdateGovernanceProfile` RPC 退役、前端 `features/governance-profile` 切片整体移除、后端模块收缩为纯读（仅保留 `GovernanceProfileReader` 服务 brief）；画像主记录三字段（`track_type / template_source / current_phase`）数据保留并以轻量内联消息存在于 brief（用户强化要求 + 裁决④）
- 退役完整性：退役触点（proto 包 / 存储表 / 后端模块 / RPC / 前端切片与编辑位 / brief 装配）缺一不可（CON-04）
- 先消费后维护：agent 写回不进入本阶段（CON-09）
- 反假大空：任何子任务设计若使综合成本高于"本地目录手动维护"基线，即违反 CON-07，须回退

### 2.5 当前阶段交付模式

- 沿袭交付型 phase 模式：每个子任务产出 `.trae/specs/phase14_XX_*/` 三件套（spec / tasks / checklist），实现类子任务附代码与验证证据，验收类子任务附 acceptance_report，全部经独立复核后收口。

## 3. 当前阶段能力矩阵

### 3.1 `Standard Entity Foundation` 单值定义

- 一句话定义：跨项目的全局规范资产实体——以树形目录结构 + 必须文档导航（含摘要）承接长期实践沉淀的最佳实践，经绑定关系被任意产品 / 仓库 / 决策消费，web 可维护、agent 可直读。
- 解决的初心问题：理想项目目录结构 + 必须全局文件不再停留脑袋或散落各仓库，一次维护、跨项目复用、agent 开发环境直接消费。

### 3.2 Standard 实体字段矩阵

| 字段 | 类型 | 承载 | 说明 |
|---|---|---|---|
| `id` | uuid PK | 主表 | |
| `name` | text 唯一 | 主表 | 规范名（如 `Durable System Track 项目范式`） |
| `description` | text | 主表 | 一段定位说明 |
| `status` | enum | 主表 | `draft / active / retired` |
| `directory_tree` | jsonb | 主表 | 递归 `DirectoryTreeNode` 整树 |
| `created_at / updated_at` | timestamp | 主表 | |
| `change_summary` | text | revisions 表 | 每次更新追加一行，一句话演进留痕 |

### 3.3 目录树节点矩阵（DirectoryTreeNode）

| 节点字段 | 值域 | 说明 |
|---|---|---|
| `name` | 文本 | 目录名或文件名；字符集限制 `[A-Za-z0-9._-]`，禁止路径分隔符 `/` 与空白 |
| `node_type` | `directory / file` | 仅二值 |
| `role` | 短标签 | file 必填、directory 可空；约定值域 `entry / plan / rules / baseline / spec / summary`，允许自定义短标签；迁移时旧画像 role 值原值保留不归一化 |
| `summary` | 一段结构化摘要 | 混合式颗粒度的承接位（裁决①） |
| `ref` | 定位引用，可选 | 仓库内路径或外部 URL；承接旧 `entry_ref / external_url` 语义；是导航引用而非正文（裁决③边界内） |
| `children` | 节点数组 | `file` 节点必须为空数组 |

校验规则（完整版，phase14-03 实现依据）：

1. 树为单根结构：根节点有且仅有一个，且 `node_type` 必须为 `directory`
2. 空树（根无 children）仅在 `status=draft` 时合法；`active` 状态必须有至少一个 `file` 节点
3. 同层 `name` 唯一（大小写敏感）
4. `name` 字符集 `[A-Za-z0-9._-]`，非空，长度上限 64
5. 树深度上限 6（根为第 1 层）
6. 整树序列化大小上限 64KB
7. `file` 节点 `children` 必须为空数组；`role` 必填
8. `ref` 若填写必须为以 `/` 开头的仓库内相对路径或 `https://` 开头的 URL

### 3.4 绑定关系矩阵（standard_bindings）

| 字段 | 值域 | 说明 |
|---|---|---|
| `standard_id` | uuid FK | |
| `target_type` | `repository / product / decision / module`（可扩展） | CON-02 不设限 |
| `target_id` | uuid | 应用层校验目标存在 |
| `role` | `template_source / adopts`（可扩展） | `template_source`：该规范有实际模板仓库维护；`adopts`：该目标遵守此规范 |
| `note` | text 可选 | 绑定备注 |
| 唯一约束 | `(standard_id, target_type, target_id, role)` | |

### 3.5 画像系统性退役矩阵（裁决④ + 用户强化要求执行基线）

| 触点 | 退役动作 | 验收断言 |
|---|---|---|
| proto 包 | `proto/psco/governance_profile/v1` 整体退役（含 `GetGovernanceProfile / UpdateGovernanceProfile` RPC 与全部消息） | proto 目录无该包残留；buf breaking 受控通过（豁免留痕） |
| 存储 | drop 两张 bindings 表（先迁数据）；`governance_profile_record` 主表保留（brief 三字段数据源） | 旧 bindings 表无数据残留且已删除 |
| 后端模块 | `backend/internal/governanceprofile` 收缩为纯读：删除 connect 层与写路径，保留 repository + service 中服务 `GovernanceProfileReader` 的最小读取（brief 三字段） | 模块无 connect handler / 无写方法；`StandardReader` 成为 brief 两清单信息唯一来源 |
| brief 装配 | brief 的 `governance_profile` 字段类型改为 `project_context.proto` 内联轻量消息（`track_type / template_source / current_phase` 三组字段），解除对 governance_profile.proto 的依赖；两清单信息唯一来源 `standards[]` | brief 编译无跨包依赖；旧清单信息经 `standards[]` 零丢失 |
| 前端切片 | `frontend/src/features/governance-profile/` 整体移除；Repository detail 画像区（含"维护治理信息"入口与回看区）整体让位，由"该仓库绑定 Standard 的只读摘要入口"替代 | 切片目录无残留；Repository detail 无画像编辑 / 维护入口 |
| RPC 消费 | 前端与 agent 均不再调用画像独立 RPC；规范维护唯一入口 `/standards`（裁决⑦） | 网络请求无画像 RPC 调用 |

保留项：画像主记录三字段（`track_type` / `template_source` / `current_phase`）数据保留于 `governance_profile_record` 主表，经收缩后模块只读装配进 brief；`current_phase` 时间轴承接为 phase15 进入条件。

### 3.6 agent 简报演进矩阵（GetProjectBrief）

| 字段 | 演进 | 说明 |
|---|---|---|
| `standards[]` | 新增 | repository 经绑定关联的 Standard 列表（含全树 + 文档导航摘要 + ref） |
| `governance_profile` | 收缩为内联轻量消息 | `track_type / template_source / current_phase` 三组字段；类型定义内联于 `project_context.proto` |
| 两清单信息 | 迁移 | 唯一来源 `standards[]`，agent 无转译直读树形结构 |

### 3.7 前端承接矩阵

| 承接位 | 形态 | 说明 |
|---|---|---|
| `/standards` 列表 / 详情 / 创建 / 编辑 | 新建 | 树形展示 + 结构化树编辑器（裁决⑥）+ revision 回看 |
| 绑定发起位 | Standard 详情页内（裁决⑦） | target_type 选择 + 目标检索 + role；不在目标实体页发起 |
| 全局导航项 | 新增 `Standards` | 与四实体导航并列（但非第五 CRUD 主实体扩散） |
| Repository detail | 画像区整体让位 + 只读回看 | 画像切片移除；新增该仓库绑定 Standard 的只读摘要入口 |
| 移动端适配 | 对齐既有基线 | 响应式单列布局 |

### 3.8 非目标矩阵

`CON-08` 时间轴 / agent 写回 / MCP / CLI / Git 推进跟踪 / 模板仓库自动接入 / 自动同步 / Standard 版本分支与全文快照 / 第五 CRUD 主实体化 / 四实体无关重构——全部为 phase14 非目标，进入条件在 `phase14-11` 冻结。

## 4. 当前阶段验收前提

- 验收环境：phase13-11 固定 dogfooding 样本（`repository_id: ca261521-8daf-4248-8f12-43525326e759`）上补建 Standard 与绑定；恢复脚本可重复执行。
- 验收协议：固定样本 / 固定入口 / 固定问题 / 固定 rerun 记录格式，沿袭 phase13-11 协议。
- 验收门禁：八项裁决结论逐条可验证（①摘要可见且正文不在 PSCO / ②树形 jsonb + 多态绑定 / ③正文零托管 / ④画像系统性退役六触点断言 / ⑤无第五主实体扩散迹象 / ⑥树编辑器无拖拽可完成整树编辑 / ⑦绑定仅从 Standard 详情页发起 / ⑧迁移后仅存在一条合并全局 Standard 且首条 revision 记录合并来源）。
