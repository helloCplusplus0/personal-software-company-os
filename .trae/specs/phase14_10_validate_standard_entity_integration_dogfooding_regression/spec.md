# phase14-10 `Standard Entity Foundation` 联调、dogfooding 与反回归验证 Spec

## Why

phase14-07 / 08 / 09 已交付 Standard 后端主线、前端主线与画像系统性退役，但三者的整合效果（agent 直读树形结构、单规范多实体复用、退役零丢失、八项裁决落地实况）尚未在同一固定样本上经统一工具链、双侧 dogfooding 与浏览器反回归矩阵验证。本子任务是 phase14 的验证验收类收口任务：按 phase13-11 冻结的验收协议（固定样本 / 固定入口 / 固定问题 / 固定工具链顺序 / 固定 rerun 记录格式）完成取证并冻结 acceptance_report。

## What Changes

- 产出本 spec 三件套（`phase14_10_validate_standard_entity_integration_dogfooding_regression`）+ `acceptance_report.md`（冻结验收结论，格式沿袭 phase13-11）
- 验收库数据变更（均经正式入口，非种子后门）：
  - 四实体基座样本确认（缺失时以既有 `database/scripts/restore_phase11_phase12_dogfooding_sample.sh` 兜底恢复）
  - Standard 固定样本正式化：盘点既有 dogfooding Standard（含 phase14-10 前置 UI 反馈轮经 web 创建的 `899ed9f3-614a-4f19-9c63-9da7a47c3db4`），满足本 spec 固定样本语义集合则正式化并补齐缺口，否则经 web 维护会话新建；补建四类绑定（覆盖八格矩阵全部合法格）与至少一条 revision；固定 `standard_id` 留档
- 验证执行（只读取证为主）：
  - agent 读取路径 dogfooding（`GetProjectBrief.standards[]` 树形直读）+ 固定 6 问逐题留档
  - 画像退役六触点（T1-T6）复验 + 画像 RPC 404 实测 + 裁决⑧两库形态复核（验收库无源数据分支 + phase14-09 一次性验证库证据引用）
  - 八项裁决验收门禁（shared_baseline §4）逐条取证
  - 工具链四步门禁（buf build / buf lint / go test / npm run build）+ `buf breaking` 豁免留痕专项步
  - 浏览器反回归矩阵 16 页（`/standards` 四页 + Repository detail 让位后回归 + 四实体列表/详情抽查 + 既有 8 页基线）
  - web 维护路径 dogfooding 完整会话（含 phase14-10 前置 UI 反馈轮两处修复的回归验证：树编辑器输入焦点稳定、详情页一致性布局）
- **裁决触发（2026-08-18 用户指令）——brief 画像残余解耦（新增 T7 触点，实现型变更）**：
  - 判定留档：`GetProjectBrief` 在 phase14 有专门承接（phase14-02 字段演进表 / phase14-07 `standards[]` 装配 / Repository detail 只读摘要同源消费），不属于画像附带过时设计，按"让 brief 承接 phase14 能力"分支处置
  - 阻断实证：画像写路径退役后，画像主记录行无任何正式入口可重建 → brief 按冻结语义 404，同时阻断 agent 直读与 Web 端 Repository detail Standard 摘要（`use-repository-standards-read` 同源消费 brief）
  - 处置：brief 移除 `governance_profile` 与 `current_phase` 两个画像残余块（字段号+名 reserved）；删除 `BriefGovernanceProfile` / `BriefCurrentPhase` / `BriefTrackType` / `BriefPhaseStatus`；删除 `backend/internal/governanceprofile/` 模块；新增 `0012` 迁移 drop `governance_profiles` 主表；`template_source` 语义由 Standard 绑定（role=template_source）承接；`track_type` 与时间轴无 phase14 数据源不保留（phase15 如进入时间轴另行新建正规字段，不复活画像派生形态）
- 若验证发现阻断性缺陷：最小修复 + 同端口重启 + 受影响门禁完整重跑，失败点/恢复路径/rerun 结果留档（沿袭 phase13-11 §9 格式）
- 不引入新功能设计、不做结构性代码变更（除用户裁决触发的 brief 画像残余解耦〔见上〕与缺陷最小修复）；变更保持未提交，待用户确认后手动提交

## Impact

- Affected specs:
  - `docs/phase/phase14_standard_entity_foundation_dev_plan.md`（上游，phase14-10 定义 L68-71）
  - `docs/phase/phase14_standard_entity_foundation_shared_baseline.md`（§3.5 退役矩阵 / §3.3 校验规则 / §4 八项裁决验收门禁——本 spec 验收门禁的直接依据）
  - `.trae/specs/phase14_07_land_standard_backend_mainline/spec.md`、`phase14_08_land_standard_frontend_mainline/spec.md`、`phase14_09_retire_governance_profile_switch_brief/spec.md`（验收对象）
  - `.trae/specs/phase13_11_validate_project_governance_profile_integration_dogfooding_regression/acceptance_report.md`（验收协议上游：固定样本 / 固定入口 / 固定问题 / 工具链顺序 / rerun 格式逐项沿袭）
- Affected code:
  - 只读验证覆盖：`proto/`（含 `buf.yaml` 豁免配置）、`backend/internal/standard/`、`backend/internal/governanceprofile/`、`backend/internal/projectcontext/`、`frontend/src/features/standard/`、`frontend/src/features/repository-binding/`（detail 让位区）
  - 潜在缺陷修复点（条件性）：上述范围内最小修复；禁止越出 phase14 交付面
  - 验收库（psco_development）：Standard 固定样本与绑定写入

## ADDED Requirements

### Requirement: 固定样本与验收环境协议必须沿袭 phase13-11 并补建 Standard 层

验收环境：Ubuntu 24，共享 PostgreSQL（psco_development），后端 8081（ConnectRPC），前端 5173（vite dev）。

**四实体基座（沿袭 phase13-11 冻结值）**：

| 实体 | 固定值 |
| --- | --- |
| Repository | `personal-software-company-os`（`ca261521-8daf-4248-8f12-43525326e759`） |
| Product | `PSCO`（`f0d034cc-3235-4d03-879d-6a3111b95b6b`） |
| Module | `project-context-foundation`（`9b02e0ca-3175-4b5a-bcf4-23cecbb06f72`） |
| Decision | `phase11 Project Context Foundation dogfooding 验收决策`（`aa8ee5ad-b224-4ea1-b393-7e4c30e42212`） |

**Standard 固定样本（本任务补建/正式化）**：

- 固定 Web 验证页面：`/standards/<固定 standard_id>`（详情页）+ `/standards`（列表页）
- 固定 agent 读取入口：`POST /api/psco.project_context.v1.ProjectContextService/GetProjectBrief`（同一 `repository_id` 锚点）的 `standards[]`；辅证 `ListStandards` / `GetStandard` / `ListStandardRevisions`
- 固定样本语义集合（树内容冻结为语义约束而非逐字节点，dogfooding 会话按真实 PSCO 项目范式填写）：
  1. 树为单根 directory（`name="."`），含至少一层嵌套 directory（如 `docs/`），验证多级结构展示与消费
  2. 含至少 5 个 `file` 节点，覆盖根级关键全局文档（至少 `AGENTS.md` / `plan.md` / `project_rules.md` / `TECH_STACK_BASELINE.md` 四项），各节点 `role` 必填（约定值域 `entry / plan / rules / baseline / spec / summary` 或自定义短标签）
  3. 至少 4 个 `file` 节点带非空 `summary`（混合式颗粒度承接位，裁决①）与 `/` 开头 `ref`（仓库内路径）；至少 1 个节点带 `https://` 开头 `ref`（外部链接形态覆盖）
  4. `status = active`（active 必须含至少一个 file 节点，校验规则 R2 联动）
- 固定绑定样本（单规范多消费，覆盖八格矩阵全部 5 个合法格）：
  - repository `ca261521...` → `template_source`（PSCO 仓库即结构导航源）
  - product `f0d034cc...` / module `9b02e0ca...` / decision `aa8ee5ad...` → 各一条 `adopts`
- 固定 revision 样本：至少一次经 web 编辑（`UpdateStandard`，`change_summary` 必填人工一句话）产生 revision 留痕（创建不记 revision，phase14-07 合同注释）
- 样本恢复方式：四实体基座缺失时执行既有 `restore_phase11_phase12_dogfooding_sample.sh`（不使用手工 SQL 插基座）；Standard 与绑定 100% 经 `/standards` web 维护会话补建（沿袭 phase13-11"画像维护 100% 经表单"纪律，非种子 SQL、非后门脚本）

#### Scenario: 固定样本判定

- **WHEN** Task 1 完成
- **THEN** 四实体基座在库且 `GetProjectContext` 对固定 `repository_id` 一次解析成功；固定 Standard 满足语义集合 1-4；`standard_bindings` 含 5 行固定绑定（4 target_type × role 分布如上）；`ListStandardRevisions` 返回 ≥1 条；固定 `standard_id` 留档于 acceptance_report §1

### Requirement: Standard 固定样本正式化协议必须单值

- 盘点验收库既有 dogfooding Standard（含 UI 反馈轮创建的 `899ed9f3-614a-4f19-9c63-9da7a47c3db4`）：
  - 若既有 Standard 满足语义集合 1-4 → 直接正式化为固定样本，经 web 会话补齐绑定/revision 缺口
  - 若不满足 → 经 web 维护会话新建固定样本（既有条目不删除、不干扰，仅在报告中留档说明）
- 全库迁移产物唯一性核对：`SELECT count(*) FROM standards WHERE name = '默认项目范式（迁移自治理画像）'` 在验收库 = 0（phase14-09 已冻结：验收库两源表 0 行，无画像数据分支不产生迁移产物）

#### Scenario: 正式化判定

- **WHEN** 盘点完成
- **THEN** 有且仅有一条被指定为固定样本的 Standard，其 id 冻结留档；报告中说明正式化路径（既有正式化 / 新建）与依据

### Requirement: agent 读取路径 dogfooding 必须实证树形结构与文档导航直读

对固定 `repository_id` 实时调用 `GetProjectBrief`（HTTP 200）并取证：

- `standards[]` 含固定 Standard（经 repository 绑定关联），每项含全量 `directory_tree`（嵌套 directory / file 结构、file 节点 `role / summary / ref` 完整）——agent 无转译直读树形结构
- 无画像残余块（裁决触发 T7）：`governance_profile` 与 `current_phase` 已整体移除（字段号 2/4 reserved）；`template_source` 语义经 Standard 绑定（role=template_source）随 `standards[]` 消费
- brief 字段面 = 5 顶层块（编号槽位 2/3/4 reserved）：repository / products / modules / decisions / standards[]；无 `global_assets` / 无 `governance_profile` / 无 `current_phase`、无目录扫描结果、无 Git 状态、无第二套事实源投影、无正文内容字段（裁决③正文零托管）
- 同源性：`GetStandard(固定 id)` 的 `standard.directory_tree` 与 brief 内该 Standard 树逐字段一致

### Requirement: 固定 6 问逐题留档必须全达标

沿袭 phase13-11 格式（answer / direct entry refs / 是否达标），覆盖 dev_plan L70 五类取证：

- **Q1（agent 直答树形结构与文档导航）**：当前项目的全局规范目录结构是什么？关键全局文档（入口 / 规划 / 规则 / 技术基线）各在什么角色与入口可直答？——取证 `GetProjectBrief.standards[].directory_tree` 与 file 节点 `role / ref`；Web 详情页树形展示同源
- **Q2（画像退役六触点零丢失）**：画像旧两清单信息（canonical 根文件 / 全局资产）现由什么承接，是否有丢失？——取证 T1-T7 复验结果 + brief 无旧字段 + phase14-09 `migration_evidence.md` §5 零丢失对照表（一次性验证库证据）
- **Q3（单规范多产品复用）**：一条 Standard 能否同时被 repository / product / module / decision 消费，绑定语义如何区分？——取证固定样本 5 行绑定（`template_source` 仅 repository、`adopts` 全开）+ brief（repository 锚点）standards[] 消费实况
- **Q4（revision 留痕回看）**：Standard 每次演进是否留痕，如何回看？——取证 `ListStandardRevisions` ≥1 条 + `change_summary` 人工一句话 + Web 详情页 Revision 历史区
- **Q5（维护唯一入口与先消费后维护）**：规范维护的唯一入口在哪，agent 是否仍无写回通道？——取证 `/standards` 四页唯一写路径 + `StandardService` 8 RPC 中写 RPC 仅 web 调用（前端切片 owner 唯一）+ 无 MCP / CLI / agent 写回承接位
- **Q6（画像残余彻底退役与时间轴未偷渡）**：brief 与库内是否还有画像残余？时间轴能力是否未被偷渡承接？——取证 brief 无 `governance_profile` / `current_phase` 块（字段号 2/4 reserved）+ `governance_profiles` 主表 `to_regclass` 为 NULL + `backend/internal/governanceprofile` 不存在 + 0011/0012 之后无时间轴类新表列 / 新 RPC（CON-08 归 phase15，届时新建正规字段而非复活画像派生形态）

### Requirement: 画像退役触点复验必须全绿（T1-T7）

逐条执行 phase14-09 checklist 同源机械检查并留痕（本轮为复验 + 汇总入 acceptance_report）：

- T1 proto：`proto/psco/governance_profile/v1/` 不存在；`buf.yaml` 豁免（`ignore` 单文件 + `ignore_only` 规则级单文件）留痕在且无扩大化；三端生成产物无画像包
- T2 存储：三张画像表（主表 + 两张 bindings 表）`to_regclass` 均 NULL（主表经 0012 裁决触发 drop）
- T3 后端模块：`backend/internal/governanceprofile` 目录整体不存在（T7 整体删除）；`grep -r "governanceprofile" backend/internal` 零命中（无残留 import）；`platform` 层无画像 RPC 挂载
- T4 RPC：`/api/psco.governance_profile.v1/GovernanceProfileService/GetGovernanceProfile` Connect 请求实测 404（对照 `ListStandards` 200 / `healthz` 200，排除整体路由故障）
- T5 前端：`frontend/src/features/governance-profile/` 不存在；`grep -r "governance-profile" frontend/src` 零命中（排除 phase14-10 前置修复的合法注释命中需逐条说明）
- T6 brief：字段面为 5 顶层块（槽位 2/3/4 reserved）；两清单信息唯一来源 `standards[]`；无 `BriefGovernanceProfile` / `BriefCurrentPhase` 等画像派生消息（消息定义已随 T7 删除）
- T7 画像残余（裁决触发 2026-08-18）：proto 字段号 2/4 + 字段名 reserved 且画像派生消息删除；`0012` 迁移已应用（主表 `to_regclass` NULL + `schema_migrations` 含 0012）；buf breaking 豁免最小扩展（仅 `project_context.proto` 单文件规则级）留痕且无目录前缀扩大化；Web 端 Repository detail Standard 摘要经 brief 正常加载（无画像前提不再 404）
- 裁决⑧两库形态：验收库 0011 重放幂等无产物无报错（`psql -f` 重放一次）+ `默认项目范式（迁移自治理画像）` count=0（无源数据分支）；一次性验证库证据（phase14-09 `migration_evidence.md` §3 恰 1 条 + revision 含 N/M 与源 repository + §5 零丢失对照）引用复核通过

### Requirement: 八项裁决验收门禁必须逐条取证（shared_baseline §4）

| 裁决 | 取证断言 |
| --- | --- |
| ① 混合式颗粒度 | Web 详情页 directory 与 file 节点 `summary` 可见；brief 树 JSON 含 summary；树内与 brief 响应均无正文内容字段（正文以模板仓库为唯一事实源） |
| ② 主表 + jsonb 树 + 多态绑定 | `standards.directory_tree` 为 jsonb（`\d standards`）；`standard_bindings` 含 4 类 target_type 实例行 + 四元组唯一约束在 |
| ③ 正文零托管 | file 节点仅含 `ref` 定位引用；`grep` 树 JSON / brief 响应无正文段落承载字段；代码无仓库正文抓取逻辑 |
| ④ 画像系统性退役 | T1-T7 全绿（上项 Requirement，含裁决触发 T7 画像残余彻底退役） |
| ⑤ 无第五主实体扩散 | Dashboard 无 Standard 主卡片挂载（grep `features/dashboard` 无 standard 引用）；四实体页面无 Standard CRUD 侵入（仅 Repository detail 只读摘要入口）；全局导航并列但 `/standards` 独立成组 |
| ⑥ 树编辑器无拖拽 | `grep -rn "draggable\|onDragStart\|dnd" frontend/src/features/standard` 零命中；dogfooding 编辑会话全程经 增删 / 上移下移 / 添加子节点 完成整树编辑 |
| ⑦ 绑定仅 Standard 详情页发起 | bind/unbind mutation caller 唯一（`StandardBindingPanel`）；module / product / repository / decision 四类页面无 Standard 绑定发起入口（grep + 浏览器抽查） |
| ⑧ 存量合并迁移 | 两库形态复核（上项 Requirement）；验收库无第二条迁移产物 |

### Requirement: 工具链四步门禁 + breaking 豁免专项必须单值顺序完整执行

沿袭 phase13-11 四步框架，本阶段增加 breaking 专项步（dev_plan L70"画像包 breaking 豁免留痕"）：

| 步骤 | 目录 | 命令 | 本阶段附加验证 |
| --- | --- | --- | --- |
| 1 | `proto/` | `buf build` | 退出码 0，无 warning |
| 2 | `proto/` | `buf lint` | 退出码 0，无 warning |
| 2b（专项） | `proto/` | `buf breaking`（对仓库 `.git` 基准） | 豁免留痕生效：画像包 `ignore` 单文件 + `ignore_only: FIELD_WIRE_JSON_COMPATIBLE_TYPE` 规则级单文件 + T7 触发的 brief 字段/消息移除所需最小规则级豁免（仅 `project_context.proto` 单文件），无目录前缀扩大化；退出码 0 |
| 3 | `backend/` | `go build ./... && go vet ./... && go test ./...` | 全绿；集成测试连验收库，重跑前 pg_dump 备份、重跑后复验固定样本（沿袭 phase13-11 保护协议） |
| 4 | `frontend/` | `npx tsc -b && npm run build` | 零错误（含 phase14-10 前置 UI 反馈轮两处修复的编译回归） |

### Requirement: 浏览器反回归矩阵必须逐页通过（16 页）

| # | 页面 | 通用检查点 | 专属检查点 |
| --- | --- | --- | --- |
| 1 | `/standards` | 加载 / 无 not_found / 无画像残留 | 列表含固定样本行；标题行/空态/紧凑基线对齐 Dashboard 基线 |
| 2 | `/standards/<固定id>` | 同上 | 树形展示（嵌套 + file 节点 role/summary/ref）+ 绑定管理区 5 行 + Revision 历史区；UI 反馈轮一致性布局在位（返回行 / 语义导语 / 左摘要卡右内容 grid） |
| 3 | `/standards/new` | 同上 | 创建表单 + 树编辑器可用 |
| 4 | `/standards/<固定id>/edit` | 同上 | 树编辑器整树编辑；**输入焦点稳定（UI 反馈修复回归：连续输入不跳焦）**；无拖拽交互 |
| 5 | `/repositories/ca261521-...` | 同上 | 画像区已让位：无"维护治理信息"入口；该仓库绑定 Standard 的只读摘要入口在且树 compact 模式正确；**解耦回归（T7）：Standard 摘要经 brief 正常加载（画像主表已 drop 后无 404）** |
| 6 | `/dashboard` | 同上 | 概览 / Current Focus / Asset Feedback / Recent Activity 正常，无 Standard 主卡片 |
| 7-8 | `/modules` + `/modules/9b02e0ca-...?fromList=true&statusFilter=all` | 同上 | 列表/详情正常 |
| 9-10 | `/decisions` + `/decisions/aa8ee5ad-...?fromList=true&statusFilter=all` | 同上 | 列表/详情正常 |
| 11-12 | `/products` + `/products/f0d034cc-...?fromList=true&statusFilter=all` | 同上 | 列表/详情正常 |
| 13 | `/repositories` | 同上 | 列表正常 |
| 14 | `/onboarding` | 同上 | 引导流程正常 |
| 15-16 | `/reviews/daily` + `/reviews/weekly` | 同上 | Review 双路径正常 |

### Requirement: web 维护路径 dogfooding 必须完成完整会话

经 `/standards` 真实操作（非 SQL 后门）完成：创建（或正式化核对）→ 结构化树编辑（节点增删 / 上移下移 / 添加子节点 / file 节点 role+summary+ref 填写）→ 保存 → 详情页绑定 repository / product / module / decision（含 `template_source` 联动禁用验证：target_type ≠ repository 时该选项禁用）→ 编辑更新产生 revision → Revision 回看。全程无控制台错误；关键步骤截图留档（沿袭 phase13-11 step-N 命名）。

### Requirement: acceptance_report 必须冻结验收结论并经独立复核

- 报告结构沿袭 phase13-11：固定样本与 id → 固定入口 → 样本恢复与维护方式 → 双侧 dogfooding 取证 → 固定 6 问逐题留档 → 工具链逐步结果 → 浏览器矩阵逐页结果 → 八项裁决门禁矩阵 → 边界证据（四类非目标）→ 失败点/恢复/rerun → 是否达标 → Rerun 指引
- 子代理独立复核（固定问题：样本协议一致性 / 六触点证据真实性 / 八项裁决取证完整性 / 工具链顺序与结果 / 矩阵覆盖完整性 / 范围外改动检查）；复核发现的阻断问题修复后方可收口
- DoD（dev_plan L71）：固定问题全达标；八项裁决验收门禁全绿；反回归矩阵全绿；acceptance_report 冻结验收结论；独立复核通过

## MODIFIED Requirements

### Requirement: brief 画像残余保留口径（phase14-02"内联三字段 + current_phase 三字段保留"）被 2026-08-18 用户裁决取代

**Reason**: phase14-10 验证实证——画像写路径退役后，画像主记录行无任何正式入口可重建，brief 对所有无存量画像行仓库（含全新部署的全部新建仓库）永久 404，并同源阻断 Web 端 Repository detail Standard 摘要；用户裁决画像为错误设计，残余依赖必须系统性移除、brief 改为承接 phase14 自有能力。

**Migration**: brief `governance_profile`（field 2）/ `current_phase`（field 4）移除并 reserved（号+名）；`BriefGovernanceProfile` / `BriefCurrentPhase` / `BriefTrackType` / `BriefPhaseStatus` 消息删除；`governanceprofile` 模块与 `governance_profiles` 主表（0012 迁移）整体退役；`template_source` 语义由 Standard 绑定（role=template_source）承接；`track_type` 与时间轴不保留，phase15 如进入时间轴另行新建正规字段。本 spec 正文（ADDED Requirements 新口径）为验收唯一依据，phase14-02 对照表"后"列中受影响两行以本节裁决为准。

## REMOVED Requirements

无。
