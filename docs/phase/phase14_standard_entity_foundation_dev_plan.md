# phase14_standard_entity_foundation_dev_plan

## 1. 文档定位

- 本文档是 `phase14_standard_entity_foundation` 的任务分解与推进计划，冻结该阶段全部子任务的范围、DoD 与依赖关系。
- 本文档的直接上游是同阶段 `architecture_plan` 与 `shared_baseline`；三者不一致时以 `shared_baseline` 的单值基线为准并回改。

## 2. 本阶段目标

- 单一主交付：`Standard` 全局规范实体最小主线（合同 → 存储 → 后端 → 前端 → agent 消费）+ 治理画像设计系统性退役。
- 八项裁决结论（①混合式颗粒度 / ②主表+jsonb树+多态绑定表 / ③PSCO=结构导航仓库=正文 / ④phase14 内画像系统性退役 / ⑤全局规范资产实体 / ⑥结构化树编辑器 / ⑦绑定从 Standard 详情页发起 / ⑧存量合并为一条全局 Standard）是全部子任务的强制边界。

## 3. 子任务清单

### 第一组：边界收敛类子任务

### phase14-01 冻结 `Standard Entity Foundation` 的范围边界、成功标准与非目标

- 范围：把八项裁决结论与 `phase13-12` 缺口记录（GAP-01~07 / CON-01~09）收敛为 `phase14` 的单一主交付边界声明；产出 `.trae/specs/phase14_01_freeze_standard_entity_scope_success_non_goals/` 三件套。
- DoD：spec 与三件套、根级文档单值一致；GAP/CON 覆盖矩阵留档；独立复核通过。

### phase14-02 冻结 Standard 与四实体主线、画像退役的关系边界

- 范围：冻结多态绑定关系矩阵（`standard_bindings` 的 `target_type` / `role` 枚举与语义）、画像系统性退役映射表（architecture_plan §4.7 六触点：proto 包 / 存储表 / 后端模块收缩 / RPC / 前端切片 / brief 装配）的执行层确认、`GetProjectBrief` 字段演进前后对照（含 `governance_profile` 内联轻量消息的字段清单）。
- DoD：关系矩阵单值；退役映射覆盖六触点且每触点有验收断言；`current_phase` 三字段保留口径显式冻结（phase15 时间轴进入条件）；brief 前后对照表留档。

### 第二组：实现设计类子任务

### phase14-03 产出 Standard 数据模型与目录树设计

- 范围：`standards` / `standard_revisions` / `standard_bindings` 三表字段级设计；`DirectoryTreeNode` 递归结构（`name / node_type / role / summary / ref / children`）jsonb 规范；树校验规则完整落地（shared_baseline §3.3 的 8 条：单根 directory、空树仅 draft 合法、同层 name 唯一、name 字符集 `[A-Za-z0-9._-]` 上限 64、深度上限 6、序列化上限 64KB、file 无 children 且 role 必填、ref 格式约束）。
- DoD：字段矩阵与校验规则可直接进入迁移与 proto 实现；双清单合一（GAP-02 修正）落为树节点单一承载；`ref` 字段承接旧 `entry_ref / external_url` 语义无丢失。

### phase14-04 产出后端合同、存储与读写边界设计

- 范围：`psco.standard.v1` 合同（消息 + 8 个 RPC 的请求/响应 envelope 与错误语义：`ListStandards` 不分页、`UpdateStandard` 整树原子替换且 `change_summary` 必填、`BindStandard` target 不存在的错误语义、`ListStandardRevisions` 作为 revision 回看读取位）；Go 模块分层（`backend/internal/standard/`：connect / service / repository）；写路径归属（web 写、agent 只读）；`StandardReader` 跨模块读取接口签名（输入 repository_id、输出 Standard 列表含整树）；brief 装配演进点。
- DoD：合同与 `.proto` 主线、ConnectRPC、canonical owner 单值化对齐；跨模块读取经独立 Read 接口（沿袭 phase13 candidate 模式）；8 个 RPC 逐个有请求/响应/错误三要素定义。

### phase14-05 产出前端信息架构与维护入口设计

- 范围：`/standards` 列表页 + 详情页（树形展示 + 文档节点摘要 + 绑定管理区 + revision 回看）+ 创建/编辑（结构化树编辑器：树形缩进列表、每节点一行表单化编辑 name/node_type/role/summary/ref、节点增删/上移下移/添加子节点、无拖拽）；绑定发起位冻结为 Standard 详情页内（target_type 选择 + 目标检索 + role）；全局导航项接入；`frontend/src/features/standard/` 切片结构（data / api-adapter / application owner / pages）；Repository detail 画像区整体让位方案与 Standard 只读摘要入口；移动端适配对齐既有基线。
- DoD：页面文件级映射、URL 语义、组件树、树编辑器交互规格（含节点操作清单与禁用态规则）达到"足以直接进入实现"的 DoD 标准；mutation 收敛到切片内固定承接位（project_rules §2.5）；绑定 UI 仅在 Standard 详情页（裁决⑦）。

### phase14-06 产出画像退役与数据迁移设计

- 范围：存量 `canonical_root_files` / `global_asset_bindings` → 一条全局 Standard 的迁移映射（裁决⑧：合并策略、不一致字段取最新 `updated_at`、多源差异记入首条 revision；`role / summary` 保真，`entry_ref / external_url` → `ref`）；多级 `path` 拆树规则（如 `docs/phase/x.md` → 嵌套 directory 节点）；迁移脚本入口与幂等语义（`WHERE NOT EXISTS` / `ON CONFLICT` 级别明确）；`0011` 迁移中 drop 两表的顺序（先迁数据后 drop）；前端 `features/governance-profile` 切片整体移除清单；后端画像模块收缩为纯读的文件级清单（删 connect 层与写路径、保留 `GovernanceProfileReader`）；brief `governance_profile` 内联轻量消息设计；buf breaking 对画像包删除的豁免配置方案。
- DoD：迁移可机械执行且幂等；合并策略单值；退役六触点（proto 包 / 存储表 / 后端模块 / RPC / 前端切片与编辑位 / brief 装配）逐项有验收断言；brief 前后对照表留档。

### 第三组：源代码实现类子任务

### phase14-07 落实 Standard 后端主线

- 范围：`proto/psco/standard/v1/standard.proto` + buf 生成；`0011_phase14_standard_entity.sql`（三表 + 枚举 + 唯一约束）；`backend/internal/standard/`（connect / service / repository 分层 + validate / errors 支撑文件单值化；8 个 RPC 全量实现 + 8 条树校验规则）；`projectcontext` 侧 `StandardReader` 接口接入与 `GetProjectBrief` 装配 `standards[]`。
- DoD：`buf lint / build / breaking` 通过；`go build / vet / test` 通过；brief 集成测试含 `standards[]` round-trip 断言；树校验规则有单元测试覆盖（含非法树用例）。

### phase14-08 落实 Standard 前端主线

- 范围：`frontend/src/features/standard/` 切片（query 纯只读 + mutation 固定承接位）；`/standards` 路由（列表 / 详情 / 创建 / 编辑）；结构化树编辑器组件（裁决⑥规格）与树形只读展示；详情页绑定管理区（裁决⑦）；revision 回看；导航项接入；Repository detail 中该仓库绑定 Standard 的只读摘要入口。
- DoD：`tsc --noEmit` 零错误；浏览器可完成"创建 Standard → 结构化树编辑 → 从详情页绑定 repository/product → revision 留痕 → 回看"完整会话；无拖拽交互。

### phase14-09 落实画像系统性退役与 brief 切换

- 范围：执行 phase14-06 迁移设计（存量数据按裁决⑧合并迁至一条全局 Standard，drop 两张 bindings 表）；删除 `proto/psco/governance_profile/v1` 整个包并在 proto 配置显式豁免留痕；后端 `governanceprofile` 模块收缩为纯读（删 connect 层与写路径，保留 `GovernanceProfileReader` 最小读取）；brief `governance_profile` 字段切换为 `project_context.proto` 内联轻量消息；整体移除前端 `features/governance-profile` 切片与 Repository detail 画像区。
- DoD：退役六触点断言全绿（proto 目录无画像包 / 旧表已删无残留 / 模块无写方法 / 前端切片无残留 / 无画像 RPC 调用 / brief 语义不缺失）；brief 对照验收（旧清单信息经 `standards[]` 零丢失，`ref` 承接 `entry_ref / external_url`）；迁移后全库仅存在一条合并全局 Standard 且首条 revision 记录合并来源（裁决⑧）。

### 第四组：验证验收类子任务

### phase14-10 完成 `Standard Entity Foundation` 的联调、dogfooding 与反回归验证

- 范围：固定样本（phase13-11 样本上补建 Standard 与绑定）；固定问题取证（agent 直答树形结构与文档导航 / 画像系统性退役六触点零丢失 / 单规范多产品复用 / revision 留痕回看 / 八项裁决逐条验证 per shared_baseline §4）；工具链四步门禁（画像包 breaking 豁免留痕）；浏览器反回归矩阵（`/standards` 新页面 + Repository detail 画像区让位后回归 + 四实体列表/详情抽查）；双侧 dogfooding（web 维护路径 + agent 读取路径）。
- DoD：固定问题全达标；八项裁决验收门禁全绿；反回归矩阵全绿；acceptance_report 冻结验收结论；独立复核通过。

### 第五组：根级同步类子任务

### phase14-11 完成根级同步、阶段收口与下一阶段进入条件回写

- 范围：回写 `AGENTS.md / plan.md / docs/README.md / architecture_map.md / docs/phase/README.md`；冻结 `phase15` 进入条件（含 `CON-08` 时间轴承接 `current_phase` 三字段、`standard_bindings` 目标类型扩展、agent 写回等后续项的排序）。
- DoD：根级五文档单值一致；无孤岛；单一真相源不破坏；变更待用户确认后提交。

## 4. 明确不做

- `CON-08` 阶段推进时间轴（`current_phase` 三字段仅以 brief 内联轻量消息保留，不新增承接位、不改动数据）
- agent 写回、MCP / CLI、Git 推进跟踪、模板仓库自动接入、自动同步
- `Standard` 版本分支、全文快照、正文托管
- 第五 CRUD 主实体化（Dashboard 主卡片、四实体关系主链）
- 四实体既有页面的无关重构（Repository detail 仅触达画像让位局部）
- 画像主记录三字段（`track_type / template_source / current_phase`）的进一步演进（保留即止，时间轴归 phase15）

## 5. 子任务依赖关系

- `phase14-02` depends on `phase14-01`
- `phase14-03` `phase14-04` `phase14-05` `phase14-06` depend on `phase14-02`
- `phase14-07` depends on `phase14-03` `phase14-04`
- `phase14-08` depends on `phase14-05`（树形组件与切片设计）
- `phase14-09` depends on `phase14-06` `phase14-07`（迁移需 Standard 存储与读取就绪）
- `phase14-10` depends on `phase14-07` `phase14-08` `phase14-09`
- `phase14-11` depends on `phase14-10`
