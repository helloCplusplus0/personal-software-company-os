# phase14-05 Standard 前端信息架构与维护入口设计 Spec

## Why

phase14-04 已冻结后端合同（8 RPC）与 StandardReader 装配，但 `/standards` 前端承接（页面映射 / URL / 组件树 / 结构化树编辑器交互 / 绑定管理区）、切片结构、Repository detail 画像区让位方案尚未定义。本子任务产出"足以直接进入实现"的前端设计，phase14-08 不再做任何设计决策。本子任务为实现设计类，交付设计文档，不写实现代码。

## What Changes

- 产出本 spec 三件套（`phase14_05_design_frontend_info_arch_maintenance_entry`）
- URL 语义与路由文件级映射冻结（4 条路由 + 全局导航项接入 `__root.tsx` NAV_ITEMS）
- `frontend/src/features/standard/` 切片结构冻结（22 文件清单：data 5 / application 5 / pages 4 / components 6 / types / index；query 纯只读 + mutation 固定承接位，project_rules §2.5）
- 四页面组件树冻结（列表 / 详情 / 创建 / 编辑）
- 结构化树编辑器交互规格冻结（裁决⑥：节点行 5 字段、操作清单、禁用态规则表、node_type 切换规则、校验反馈双层模型、无拖拽）
- 绑定管理区交互规格冻结（裁决⑦：仅 Standard 详情页；target_type 选择 → 目标检索 → role 联动禁用 → note；数据源复用四实体 owning 切片 list query owner）
- Repository detail 画像区让位方案与 Standard 只读摘要入口冻结（数据源 `GetProjectBrief.standards[]`；让位时序：phase14-08 同位替换挂载 / phase14-09 删除切片目录）
- 移动端适配基线冻结（对齐 Dashboard 基线与既有紧凑化规范）
- 单值化决策留档：摘要数据源走 brief（不新增第 9 RPC）、目标检索复用既有 list owner（不新建检索 RPC）、前端轻量校验 + 后端权威 R1-R8 双层模型、树编辑器本地 draft state 整树提交

## Impact

- Affected specs:
  - `docs/phase/phase14_standard_entity_foundation_dev_plan.md`（上游，phase14-05 定义 L39-42）
  - `.trae/specs/phase14_04_design_backend_contract_storage_read_write_boundaries/spec.md`（直接上游：8 RPC 三要素 / 消息结构——前端消费模型与 mutation 的唯一合同依据）
  - `.trae/specs/phase14_03_design_standard_data_model_directory_tree/spec.md`（树节点 6 字段 / R1-R8 校验规则——编辑器禁用态与校验反馈的依据）
  - `.trae/specs/phase14_02_freeze_standard_entity_relation_boundaries/spec.md`（八格矩阵——绑定表单 role 联动禁用依据）
  - `docs/phase/phase14_standard_entity_foundation_shared_baseline.md`（§2.2 裁决⑥⑦ / §3.7 前端承接矩阵）
- Affected code: 无（零代码改动；设计对象为未来的 `frontend/src/features/standard/`、`frontend/src/routes/standards/`、`__root.tsx` NAV_ITEMS 追加、`repository-binding-detail-page.tsx` L27/L305-306 挂载替换，均由 phase14-08 落地）
- 设计参照（本轮实际读取）：`routes/` 扁平文件路由现状、`__root.tsx` L25-31 NAV_ITEMS、`module-registry` 切片（query owner `use-module-list-read.ts` / mutation owner `use-create-draft-module.ts` 模式）、四实体 data 层 list owner 实证（`use-module-list-read` / `use-product-list-read` / `use-repository-list-read` / `use-decision-list-read`）、`repository-binding-detail-page.tsx` 画像区挂载 L27/L305-306

## ADDED Requirements

### Requirement: URL 语义与路由映射必须冻结

路由 SHALL 按 TanStack Router 既有扁平文件模式新增 4 文件，URL 语义单值如下：

| 路由文件 | URL | 页面组件 | 语义 |
|---|---|---|---|
| `routes/standards/index.tsx` | `/standards` | `StandardListPage` | 规范列表（无搜索参数：ListStandards RPC 无参数不分页，第一版不做筛选） |
| `routes/standards/new.tsx` | `/standards/new` | `StandardCreatePage` | 创建规范 |
| `routes/standards/$standardId.tsx` | `/standards/:id` | `StandardDetailPage` | 规范详情（树展示 + 绑定管理 + revision 回看） |
| `routes/standards/$standardId.edit.tsx` | `/standards/:id/edit` | `StandardEditPage` | 编辑规范（整树替换 + change_summary） |

导航接入：`__root.tsx` NAV_ITEMS 追加 `{ to: '/standards', label: 'Standards' }`，位置在 Repository Binding 之后（与四实体导航并列，baseline §3.7）。

#### Scenario: 路由实现判定

- **WHEN** phase14-08 落地路由与导航
- **THEN** 4 路由文件与 NAV_ITEMS 追加与本节逐字一致；导航 PC 常驻与移动端菜单同步出现 Standards 项（沿袭 `__root.tsx` 既有双端渲染）

### Requirement: 切片结构必须冻结（文件级）

`frontend/src/features/standard/` SHALL 按以下 22 文件清单实现（结构沿袭 module-registry 切片模式；types 对齐后端 snake_case）：

**data/（5 文件，全部纯只读 query owner）**

| 文件 | query key | 承接 |
|---|---|---|
| `connect-client.ts` | — | 导出 `standardClient`（StandardService）与 `projectContextClient`（brief 读取用） |
| `use-standard-list-read.ts` | `['standard-list']` | ListStandards → `Standard[]` |
| `use-standard-detail-read.ts` | `['standard-detail', id]` | GetStandard → `{ standard, bindings }` |
| `use-standard-revisions-read.ts` | `['standard-revisions', id]` | ListStandardRevisions → `StandardRevision[]` |
| `use-repository-standards-read.ts` | `['repository-standards', repositoryId]` | GetProjectBrief → 投影 `standards[]`（Repository detail 摘要数据源，见让位方案 Requirement） |

**application/（5 文件，固定 mutation 承接位）**

| 文件 | RPC | 失效矩阵 |
|---|---|---|
| `use-create-standard.ts` | CreateStandard | `['standard-list']` |
| `use-update-standard.ts` | UpdateStandard | `['standard-list']`、`['standard-detail', id]` |
| `use-delete-standard.ts` | DeleteStandard | `['standard-list']` |
| `use-bind-standard.ts` | BindStandard | `['standard-detail', id]`、`['repository-standards']` 前缀 |
| `use-unbind-standard.ts` | UnbindStandard | `['standard-detail', id]`、`['repository-standards']` 前缀 |

**pages/**（4 文件）：`standard-list-page.tsx` / `standard-detail-page.tsx` / `standard-create-page.tsx` / `standard-edit-page.tsx`

**components/**（6 文件）：`standard-tree-view.tsx`（只读树）/ `standard-tree-editor.tsx`（编辑器外壳与递归渲染）/ `tree-node-editor-row.tsx`（单节点行）/ `standard-binding-panel.tsx` / `standard-revision-list.tsx` / `standard-readonly-summary.tsx`（Repository detail 挂载）

**根文件**（2 文件）：`types.ts`（Standard / DirectoryTreeNode / StandardBinding / StandardRevision / 枚举 string union / 绑定表单模型）、`index.ts`（barrel，仅导出页面与 `standard-readonly-summary`）

切片纪律（project_rules §2.5）：

- query 层零写动作；5 写 RPC 全部收敛 application 5 owner，页面与组件不内联 `useMutation`
- `standard-tree-view` / `tree-node-editor-row` 留在本切片，不晋升 `shared`（切片优先，延迟晋升）

#### Scenario: 切片实现判定

- **WHEN** phase14-08 落地切片
- **THEN** 文件清单与本表一一对应；`grep useMutation` 仅命中 application 5 文件；失效矩阵逐字实现

### Requirement: 四页面组件树必须冻结

**StandardListPage**：页壳（标题行 `flex-col gap-2 sm:flex-row sm:items-center sm:justify-between` + 新建 CTA `h-9`）→ 错误区（`p-3 text-xs`）→ 加载态 → 空态（`p-4 text-xs` + 引导新建文案）→ 摘要卡列表（每卡 `p-3 space-y-2 hover:bg-muted/50`：name + status Badge + description `truncate` + updated_at；整卡 Link 进详情）。

**StandardDetailPage**：头部（name + status Badge + description + updated_at + 操作组：编辑→`/standards/:id/edit`、删除→确认弹窗后 `use-delete-standard`，ACTIVE 态删除按钮禁用并提示先 Retire）→ `StandardTreeView`（树形只读）→ `StandardBindingPanel`（绑定管理区）→ `StandardRevisionList`（`border-t pt-2` 分隔，轻量列表）。

**StandardCreatePage**：基本信息表单（name 必填 / description 可选 / status select 默认 draft，不含 retired 选项——对齐 CreateStandard 约束）→ `StandardTreeEditor`（初始树 = 单根 directory `name="."`）→ 提交（`use-create-standard`，成功 → 详情页）/ 取消（→ `/standards`）。

**StandardEditPage**：预填基本信息（`use-standard-detail-read`）+ `change_summary` 必填输入 → `StandardTreeEditor`（预填全树）→ 保存（`use-update-standard`：整树 + change_summary + optional name/description/status，成功 → 详情页）/ 取消（→ 详情页）。

四页面共享页面壳层模式：主标题 `text-xl`、语义导语 `text-xs text-muted-foreground`、子区域 `border-t` 分隔（避免三层卡片堆叠）、容器型元素仅 `:focus-visible`。

#### Scenario: 页面实现判定

- **WHEN** phase14-08 落地四页面
- **THEN** 组件树与本节逐层一致；无 Card 重型嵌套；mutation 仅经 application owner 调用

### Requirement: 结构化树编辑器交互规格必须冻结（裁决⑥）

`StandardTreeEditor` SHALL 满足以下规格（无拖拽；树形缩进列表，每节点一行表单化编辑）：

**状态承载**：编辑器持有整树本地 draft state（React state），作为所属页面表单 state 的组成部分；提交时整树随 Create/Update 请求发出（对齐"整树原子替换"，无节点级增量协议）。

**节点行（`tree-node-editor-row.tsx`）**：按层级缩进（每层 `pl-4`）；一行内 5 个输入：`name`(text) / `node_type`(select：directory/file) / `role`(text) / `summary`(text) / `ref`(text) + 操作按钮组（添加子节点 / 删除 / 上移 / 下移，`h-7 px-2 text-xs variant=outline`）。移动端节点行折行为字段网格（`grid-cols-1 sm:grid-cols-2 lg:grid-cols-[...]`），操作组 `flex-wrap` 独立成行。

**节点操作清单**：添加子节点（插入为该层末尾，默认 `node_type=directory`、字段全空）/ 删除 / 上移（同层内交换）/ 下移（同层内交换）/ 添加根级节点（编辑器工具行：根 children 追加，等同对根"添加子节点"）。

**禁用态规则表（单值判据）**：

| 场景 | 规则 | 依据 |
|---|---|---|
| 根节点 name / node_type / role | 只读或禁用（name 固定 `"."`、node_type 固定 directory、role 必空）；summary 可编辑；ref 不提供输入位 | R1 / phase14-03 根规范 |
| 根节点 删除 / 上移 / 下移 | 禁用（无父层语义） | R1 |
| 根节点 添加子节点 | 允许 | — |
| file 节点 添加子节点 | 禁用 | R7 |
| directory（有 children）node_type 切换为 file | 禁用（须先清空子节点） | R7 |
| 第 5 层节点 添加子节点 | 允许（产生第 6 层节点） | R5 |
| 第 6 层节点 添加子节点 / node_type=directory 选项 | 禁用（第 6 层只允许 file） | R5 |
| 同层首节点 上移 / 末节点 下移 | 禁用 | 交互完整性 |
| 删除 directory 含 children | 允许但确认弹窗（提示连带删除全部后代） | 防误操作 |
| name 含 `/` 或空白即时输入 | 行内警告提示（字符集 `[A-Za-z0-9._-]`） | R4（前端轻量） |

**node_type 切换规则**：file→directory 直接允许（children 置 `[]`）；directory→file 仅当无 children。

**校验反馈双层模型（单值决策）**：

- 前端轻量层：即时行内提示（name 非空/字符集/长度 64、file 的 role 必填/长度 32、summary 长度 2000、ref 格式 `/` 或 `https://` 前缀）——仅 UX 提示，不阻断输入
- 后端权威层：R1-R8 完整判定在 Create/Update 响应错误中返回（phase14-04 错误语义）；前端将错误信息中的节点路径映射回对应节点行高亮显示
- 编辑器工具行显示节点计数（轻量规模提示；不做字节计数，避免过度设计）

#### Scenario: 编辑器实现判定

- **WHEN** phase14-08 落地树编辑器
- **THEN** 禁用态规则表逐行可验证（每个禁用场景有对应交互）；无拖拽交互；draft state 不发起任何网络请求；提交为整树单次调用

### Requirement: 绑定管理区交互规格必须冻结（裁决⑦）

`StandardBindingPanel` SHALL 仅存在于 StandardDetailPage（绑定发起位唯一；Repository detail 摘要为纯只读无操作）：

**现有绑定列表**：每行（target_type 标签 / target 名称 / role 标签 / note / created_at / 解绑按钮）。target 名称解析：消费对应实体 owning 切片 list query owner 的缓存数据（id→name 映射；未命中缓存显示 id 短版）。

**发起绑定表单（inline 非弹窗）**：

1. `target_type` select（repository / product / decision / module）
2. 目标检索：按所选 target_type 消费对应 owning 切片 list query owner（`use-repository-list-read` / `use-product-list-read` / `use-module-list-read` / `use-decision-list-read`），全量加载 + 下拉前端过滤（单用户数量级，不新建检索 RPC——单值决策）
3. `role` select 联动禁用：target_type ≠ repository 时 `template_source` 选项禁用（八格矩阵 UI 投影）；`adopts` 全开
4. `note` 可选输入
5. 提交（`use-bind-standard`）；`invalid_argument`（含重复绑定 "already bound"）行内错误回显

**解绑**：确认后 `use-unbind-standard`（四元组；note 不参与）。

#### Scenario: 绑定 UI 归属判定

- **WHEN** phase14-08 落地或 phase14-10 验收
- **THEN** 全站绑定发起控件仅出现在 StandardDetailPage；四实体页面与 Repository detail 摘要无任何绑定操作入口

### Requirement: Repository detail 画像区让位方案与 Standard 只读摘要入口必须冻结

**现状挂载**：`repository-binding-detail-page.tsx` L27（import `GovernanceProfileSection`）+ L305-306（画像区渲染）。

**让位方案（单值时序）**：

- phase14-08：同位替换——移除 L27 import 与 L305-306 画像区，原位置挂载 `<StandardReadonlySummary repositoryId={...} />`（swap 即让位，不保留并存的双治理区）
- phase14-09：删除 `features/governance-profile/` 切片目录（此时已无挂载点，删除不破坏编译）+ 后端退役六触点

**`StandardReadonlySummary` 规格**：

- 数据源：`use-repository-standards-read`（query key `['repository-standards', repositoryId]`，调用 `projectContextClient.getProjectBrief`，仅投影 `standards[]`）——与 agent 消费主路径同源（GetProjectBrief.standards[]），不新增第 9 RPC（单值决策）
- 展示：每条 Standard（name + status Badge + 链接 `/standards/:id`）+ 紧凑树形只读（复用 `standard-tree-view` 紧凑模式：`text-xs`、`divide-y`、file 节点带 role/summary/ref）
- 空态：`p-4 text-xs`"该仓库尚未关联 Standard"
- 过渡态说明：phase14-08 落地至 phase14-09 迁移完成之间，旧画像数据尚未迁入 Standard，摘要可能显示空态——为设计内过渡态，phase14-09 迁移后消失

#### Scenario: 让位实现判定

- **WHEN** phase14-08 完成
- **THEN** detail 页无 `governance` 引用（grep 零命中）；摘要入口可见且经 brief 读取；`tsc --noEmit` 零错误

### Requirement: 移动端适配基线必须冻结

四页面与编辑器 SHALL 对齐既有基线：

- 页头：`flex-col gap-2 sm:flex-row sm:items-center sm:justify-between`；主行动按钮移动端 `w-full sm:w-auto`
- 卡片/列表密度：`p-3 space-y-2`；hover 仅可交互元素（`hover:bg-muted/50` 限于 Link/Button 行）
- 树形结构：缩进 `pl-4`/层不收缩；节点行移动端折行（字段网格 + 操作组独立行）；`min-w-0 truncate` 防横向溢出
- 绑定面板：表单字段移动端单列（`grid-cols-1 sm:grid-cols-2`）；绑定列表行 `flex-wrap`
- 字段标签 `text-xs`；子区域 `border-t pt-2`；容器型元素仅 `:focus-visible`

## MODIFIED Requirements

无（本 spec 为新增设计规格；Repository detail 画像区的移除动作归属 phase14-08/09 执行，其退役边界已由 phase14-02 T5 冻结，本 spec 仅设计前端承接位与让位时序，不修改上游冻结口径）。

## REMOVED Requirements

无。
