# phase14-08 落实 Standard 前端主线 Spec

## Why

phase14-05 已冻结 `/standards` 前端承接的完整设计（URL / 22 文件切片清单 / 四页面组件树 / 结构化树编辑器交互 / 绑定管理区 / Repository detail 画像区让位 / 移动端基线），phase14-07 已交付后端 8 RPC 与 brief `standards[]`，但均停留在设计与后端层面。本子任务把 Standard 前端主线落地为可运行、可浏览器验收的交付物。本子任务为源代码实现类，设计决策一律以 phase14-05 冻结为准，不在实现中引入新设计。

## What Changes

- 产出本 spec 三件套（`phase14_08_land_standard_frontend_mainline`）
- 新增 `frontend/src/features/standard/` 切片，按 phase14-05 §ADDED-2 冻结的 22 文件清单实现（data 5 / application 5 / pages 4 / components 6 / types / index；query 纯只读 + mutation 固定承接位，project_rules §2.5）
- 新增 4 条路由：`routes/standards/index.tsx`（`/standards`）/ `routes/standards/new.tsx`（`/standards/new`）/ `routes/standards/$standardId.tsx`（`/standards/:id`）/ `routes/standards/$standardId.edit.tsx`（`/standards/:id/edit`）
- 修改 `frontend/src/routes/__root.tsx`：NAV_ITEMS 追加 `{ to: '/standards', label: 'Standards' }`（Repository Binding 之后）
- 修改 `frontend/src/routes/repository-binding-detail-page.tsx`：画像区同位替换——移除 `GovernanceProfileSection` import（L27）与挂载（L305-306），原位挂载 `<StandardReadonlySummary repositoryId={...} />`（swap 即让位，不保留双治理区）
- 不动：`features/governance-profile/` 切片目录（phase14-09 删除）、四实体既有页面（Repository detail 仅触达画像让位局部）

## Impact

- Affected specs:
  - `docs/phase/phase14_standard_entity_foundation_dev_plan.md`（上游，phase14-08 定义 L56-59）
  - `.trae/specs/phase14_05_design_frontend_info_arch_maintenance_entry/spec.md`（唯一直接上游：全部前端设计决策的冻结源）
  - `.trae/specs/phase14_04_design_backend_contract_storage_read_write_boundaries/spec.md`（8 RPC 三要素——前端消费合同依据）
  - `.trae/specs/phase14_03_design_standard_data_model_directory_tree/spec.md`（树节点 6 字段 / R1-R8——编辑器禁用态与校验反馈依据）
  - `.trae/specs/phase14_02_freeze_standard_entity_relation_boundaries/spec.md`（八格矩阵——绑定 role 联动禁用依据）
- Affected code:
  - 新增：`frontend/src/features/standard/`（22 文件）、`frontend/src/routes/standards/`（4 路由文件）
  - 修改：`frontend/src/routes/__root.tsx`（NAV_ITEMS）、`frontend/src/routes/repository-binding-detail-page.tsx`（画像区让位）
  - 消费（不修改）：`frontend/src/gen/proto/psco/standard/v1/`（TS 生成产物，phase14-07 已就位）、四实体 owning 切片 list query owner、`project-context` brief 读取
- 后端依赖：phase14-07 已交付的 8 RPC（`/api/psco.standard.v1.StandardService/...`）与 `GetProjectBrief.standards[]`

## ADDED Requirements

### Requirement: 路由与导航必须按冻结映射落地

- 4 路由文件与 URL 语义逐字 phase14-05 §ADDED-1 表：`/standards`（列表）/ `/standards/new`（创建）/ `/standards/:id`（详情）/ `/standards/:id/edit`（编辑）
- `__root.tsx` NAV_ITEMS 追加 Standards 项，位置在 Repository Binding 之后；PC 常驻与移动端菜单同步出现（沿袭既有双端渲染）
- 路由文件仅做页面组件接线（page component + route param 传递），不含业务逻辑

#### Scenario: 路由落地判定

- **WHEN** 访问 4 条 URL
- **THEN** 分别渲染对应页面组件；导航出现 Standards 项；`tsc --noEmit` 零错误

### Requirement: 切片必须按 22 文件清单实现且切片纪律成立

`frontend/src/features/standard/` SHALL 逐字落地 phase14-05 §ADDED-2 的 22 文件清单：

- **data/（5 文件，纯只读）**：`connect-client.ts`（standardClient + projectContextClient）/ `use-standard-list-read.ts`（key `['standard-list']`）/ `use-standard-detail-read.ts`（key `['standard-detail', id]`）/ `use-standard-revisions-read.ts`（key `['standard-revisions', id]`）/ `use-repository-standards-read.ts`（key `['repository-standards', repositoryId]`，brief 投影 `standards[]`）
- **application/（5 文件，mutation 固定承接位）**：`use-create-standard.ts` / `use-update-standard.ts` / `use-delete-standard.ts` / `use-bind-standard.ts` / `use-unbind-standard.ts`，失效矩阵逐字 phase14-05 §ADDED-2 表
- **pages/（4 文件）**：`standard-list-page.tsx` / `standard-detail-page.tsx` / `standard-create-page.tsx` / `standard-edit-page.tsx`
- **components/（6 文件）**：`standard-tree-view.tsx` / `standard-tree-editor.tsx` / `tree-node-editor-row.tsx` / `standard-binding-panel.tsx` / `standard-revision-list.tsx` / `standard-readonly-summary.tsx`
- **根文件（2 文件）**：`types.ts`（Standard / DirectoryTreeNode / StandardBinding / StandardRevision / 枚举 string union / 绑定表单模型，对齐后端 snake_case）/ `index.ts`（barrel，仅导出页面与 `standard-readonly-summary`）

切片纪律（验收断言）：

- `grep -r useMutation frontend/src/features/standard/` 仅命中 application 5 文件；query 层零写动作；页面与组件不内联 `useMutation`
- `standard-tree-view` / `tree-node-editor-row` 留在本切片不晋升 `shared`

#### Scenario: 切片落地判定

- **WHEN** 切片文件完成
- **THEN** 文件清单与冻结清单一一对应（无多余文件无缺失）；失效矩阵逐字实现；useMutation grep 断言通过

### Requirement: 四页面组件树必须按冻结结构落地

逐字 phase14-05 §ADDED-3：

- **StandardListPage**：页壳（标题行 + 新建 CTA `h-9`）→ 错误区 → 加载态 → 空态（引导新建）→ 摘要卡列表（`p-3 space-y-2 hover:bg-muted/50`：name + status Badge + description truncate + updated_at；整卡 Link）
- **StandardDetailPage**：头部（name + status Badge + description + updated_at + 操作组：编辑 / 删除（确认弹窗，ACTIVE 态禁用并提示先 Retire））→ `StandardTreeView` → `StandardBindingPanel` → `StandardRevisionList`（`border-t pt-2`）
- **StandardCreatePage**：基本信息表单（name 必填 / description 可选 / status select 默认 draft 不含 retired）→ `StandardTreeEditor`（初始单根 directory `name="."`）→ 提交（成功 → 详情页）/ 取消（→ `/standards`）
- **StandardEditPage**：预填（`use-standard-detail-read`）+ `change_summary` 必填 → `StandardTreeEditor` 预填全树 → 保存（整树 + change_summary + optional 字段，成功 → 详情页）/ 取消（→ 详情页）
- 共享壳层模式：主标题 `text-xl`、导语 `text-xs text-muted-foreground`、子区域 `border-t` 分隔、容器型元素仅 `:focus-visible`；无 Card 重型嵌套；mutation 仅经 application owner

#### Scenario: 页面落地判定

- **WHEN** 浏览器访问四页面
- **THEN** 组件树与冻结结构逐层一致；错误/加载/空态齐备；mutation 全部经 application owner 调用

### Requirement: 结构化树编辑器必须按裁决⑥规格落地

逐字 phase14-05 §ADDED-4：

- 编辑器持有整树本地 draft state（页面表单 state 组成部分）；提交时整树随 Create/Update 单次发出；draft state 不发起网络请求
- 节点行：层级缩进（每层 `pl-4`）；一行 5 输入（name / node_type select / role / summary / ref）+ 操作组（添加子节点 / 删除 / 上移 / 下移，`h-7 px-2 text-xs variant=outline`）；移动端折行为字段网格 + 操作组独立行
- 节点操作：添加子节点（插入层末尾，默认 directory 全空字段）/ 删除 / 上移 / 下移 / 编辑器工具行添加根级节点 + 节点计数
- 禁用态规则表逐行落地（根只读与操作禁用 / file 添加子节点禁用 / directory 有 children 时切 file 禁用 / 第 6 层添加子节点与 directory 选项禁用 / 同层首末节点移动禁用 / 删 directory 含 children 确认弹窗 / name 非法字符行内警告）
- node_type 切换：file→directory 直接允许（children 置 `[]`）；directory→file 仅当无 children
- 校验反馈双层：前端轻量层即时行内提示（不阻断输入）；后端权威层 R1-R8 错误信息中的节点路径映射回对应节点行高亮
- 无拖拽交互（验收断言：实现中无任何 drag 事件与 dnd 依赖）

#### Scenario: 编辑器落地判定

- **WHEN** 在创建/编辑页操作树编辑器
- **THEN** 禁用态规则表每个场景有对应交互；无拖拽；提交为整树单次调用；后端校验错误能定位到节点行

### Requirement: 绑定管理区必须按裁决⑦规格落地

逐字 phase14-05 §ADDED-5：

- `StandardBindingPanel` 仅存在于 StandardDetailPage（全站唯一绑定发起位）
- 现有绑定列表：每行 target_type 标签 / target 名称 / role 标签 / note / created_at / 解绑按钮；target 名称经对应实体 owning 切片 list query owner 缓存解析（id→name；未命中显示 id 短版）
- 发起绑定 inline 表单：target_type select → 目标检索（消费四实体 owning 切片 list owner，全量加载 + 下拉前端过滤，不新建检索 RPC）→ role select 联动禁用（target_type ≠ repository 时 template_source 选项禁用）→ note 可选 → 提交（`invalid_argument` 含 "already bound" 行内回显）
- 解绑：确认后四元组调用（note 不参与）

#### Scenario: 绑定 UI 归属判定

- **WHEN** phase14-08 完成
- **THEN** 全站绑定发起控件仅出现在 StandardDetailPage；四实体页面与 Repository detail 摘要无任何绑定操作入口

### Requirement: Repository detail 画像区同位替换必须落地

逐字 phase14-05 §ADDED-6 让位方案：

- 移除 `repository-binding-detail-page.tsx` 的 `GovernanceProfileSection` import 与挂载，原位挂载 `<StandardReadonlySummary repositoryId={...} />`（不保留双治理区并存）
- `StandardReadonlySummary`：数据源 `use-repository-standards-read`（brief `standards[]` 投影，与 agent 消费主路径同源）；每条 Standard（name + status Badge + 链接 `/standards/:id`）+ 紧凑树形只读（复用 `standard-tree-view` 紧凑模式）；空态"该仓库尚未关联 Standard"
- 过渡态说明：phase14-09 迁移前摘要可能空态，为设计内过渡态

#### Scenario: 让位落地判定

- **WHEN** phase14-08 完成
- **THEN** `grep -r governance frontend/src/routes/repository-binding-detail-page.tsx` 零命中；摘要入口经 brief 读取；`tsc --noEmit` 零错误

### Requirement: 移动端适配必须对齐既有基线

逐字 phase14-05 §ADDED-7：页头响应式折行；主行动按钮移动端 `w-full sm:w-auto`；列表密度 `p-3 space-y-2`；hover 限可交互元素；树缩进不收缩 + 节点行折行 + `min-w-0 truncate`；绑定面板字段 `grid-cols-1 sm:grid-cols-2` + 列表行 `flex-wrap`；字段标签 `text-xs`；子区域 `border-t pt-2`；容器型元素仅 `:focus-visible`。

#### Scenario: 移动端判定

- **WHEN** 窄视口（< 640px）访问四页面与 Repository detail
- **THEN** 无横向溢出；操作组折行；主按钮全宽

### Requirement: DoD 验证门禁必须全绿

- `frontend/`：`tsc --noEmit` 零错误
- 浏览器完整会话（DoD 原文）：创建 Standard → 结构化树编辑 → 从详情页绑定 repository 与 product（各至少 1 条）→ revision 留痕（编辑触发）→ 回看（ListStandardRevisions 展示）全链可完成；无拖拽交互
- 验证环境：后端 dev server（phase14-07 交付）+ 前端 dev server；验收后清理会话产生的测试数据

## MODIFIED Requirements

无（本 spec 为实现类规格；全部实现语义引用 phase14-05 冻结口径，不修改上游设计）。

## REMOVED Requirements

无。
