# phase13-11 `Project Governance Profile Foundation` 联调、dogfooding 与反回归验收报告

> 验收日期：2026-08-17
> 验收依据：本目录 `spec.md` 冻结的单值验收协议（固定样本 / 固定入口 / 固定 6 问 / 固定工具链顺序 / 固定 rerun 记录格式）
> 验收环境：Ubuntu 24，共享 PostgreSQL（psco_development），后端 8081（ConnectRPC），前端 5173（vite dev）

---

## 1. 固定样本与 `repository_id`

| 项 | 值 |
| --- | --- |
| Repository | `personal-software-company-os` |
| `repository_id` | `ca261521-8daf-4248-8f12-43525326e759` |
| Product | `PSCO`（`f0d034cc-3235-4d03-879d-6a3111b95b6b`） |
| Module | `project-context-foundation`（`9b02e0ca-3175-4b5a-bcf4-23cecbb06f72`） |
| Decision | `phase11 Project Context Foundation dogfooding 验收决策`（`aa8ee5ad-b224-4ea1-b393-7e4c30e42212`） |

## 2. 固定 Web 页面与 agent 入口

- 固定 Web 验证页面（第一版冻结）：`/repositories/ca261521-8daf-4248-8f12-43525326e759`
- 固定 agent 读取入口（第一版冻结）：
  1. `POST /api/psco.project_context.v1.ProjectContextService/GetProjectBrief`（同一 `repository_id` 驱动）
  2. `AGENTS.md`
  3. `plan.md`
  4. `GetProjectBrief.global_assets`（与 `GetGovernanceProfile.global_asset_bindings` 同源）

## 3. 样本恢复方式与画像维护方式

- 样本恢复：固定样本曾因 `projectcontext` 集成测试切库在共享开发库中丢失，已执行仓库内正式恢复脚本 `database/scripts/restore_phase11_phase12_dogfooding_sample.sh` 补齐；未使用手工 SQL 插样本、未更换 `repository_id` 锚点。
- 恢复后验证：兼容层 `GetProjectContext` 对固定 `repository_id` 一次解析成功（repository 块完整返回）；画像未创建时 `GetProjectBrief` 返回 HTTP 404 `{"code":"not_found","message":"governanceprofile: get governance profile: governance profile not created for repository"}`，与 spec 冻结的预期行为一致。
- 画像维护：治理画像数据 100% 经 `/repositories/ca261521-8daf-4248-8f12-43525326e759` 治理画像表单手工创建（空态"初始化治理画像"入口），非种子 SQL、非后门脚本：
  - `template_source` = `manual://psco-phase13-dogfooding`
  - `canonical_root_files[]` = 表单预填的当前项目范式 v1 根级文件集合 9 项（.env/AGENTS.md/architecture_map.md/plan.md/project_rules.md/README.md/TECH_STACK_BASELINE.md/global_skills.md/project_skills.md，含 role/required）
  - `global_asset_bindings[]` = 8 项资产全承接（entry_ref 均为资产自身根级路径），前 5 项摘要型资产（project_rules.md / TECH_STACK_BASELINE.md / AGENTS.md / architecture_map.md / plan.md）填写 `structured_summary`，后 3 项允许为空

## 4. 双侧 dogfooding 取证

### 4.1 人类维护路径（Web 表单）

- 空态 → "初始化治理画像" → 表单（预填 9 项根级文件）→ 填写 template_source 与 8 项资产 → 保存成功回流回看态（无整页刷新的精准失效刷新）。
- 概览区只读字段与后端 `RootFrozen*` 冻结常量逐项一致：`track_type=Durable System Track`、`docs_workflow_layout=phase/fix/audit/review`、`current_phase_name=phase13_project_governance_profile_foundation`、`current_phase_ref=plan.md#phase13_project_governance_profile_foundation`、`current_phase_status=进行中（in_progress）`。
- 刷新（重新导航）后回看数据与保存时完全一致；摘要回看区以 `structured_summary + entry_ref` 为主视觉，真实路径未成为主内容。
- 全程无控制台错误；关键步骤截图留档（step1-initial-page / step2-form-loaded / step4-form-filled / step5-after-save / step6-overview / step7-summary）。

### 4.2 agent 读取路径（brief 实时取证）

`GetProjectBrief`（画像创建后）HTTP 200，7 顶层块完整：

| 顶层块 | 取证结果 |
| --- | --- |
| `repository` | `ca261521-...` / `personal-software-company-os` / github |
| `governanceProfile` | `projectProfileVersion=project_governance_profile_v1`、`trackType=TRACK_TYPE_DURABLE_SYSTEM`、`templateSource=manual://psco-phase13-dogfooding`、`docsWorkflowLayout=phase/fix/audit/review`、`canonicalRootFiles[9]`、`globalAssetBindings[8]`、`currentPhaseName/Ref/Status` |
| `globalAssets` | 8 项资产，与 `governanceProfile.globalAssetBindings` 逐项一致；前 5 项含 `structuredSummary` |
| `currentPhase` | `{name, entryRef, status}` 与主记录三 read-only 字段单向派生一致 |
| `products` | 数组形态（长度 1：PSCO），未退化为单对象 |
| `modules` | 数组形态（长度 1：project-context-foundation） |
| `decisions` | 数组形态（长度 1：phase11 dogfooding 决策） |

同源性：`GetGovernanceProfile.profile` 与 `GetProjectBrief.governanceProfile` 逐字段一致（repositoryId / version / trackType / templateSource / docsWorkflowLayout / canonicalRootFiles / globalAssetBindings / currentPhase* / timestamps）。

字段面边界：brief 响应无目录扫描结果、无 Git 状态、无第二套事实源投影、无自然语言指导词字段。

文档入口核对：`AGENTS.md` §1（当前阶段 phase13 + 主目标）与 §5（推荐阅读顺序）、`plan.md` L7/L10/L184-191（phase13 当前状态、目标与路线块）均可直接回答 phase13 当前阶段状态。

## 5. 固定 6 问逐题留档（answer / direct entry refs / 是否达标）

### Q1 当前项目治理画像版本与技术路线是什么，在哪个固定入口可确认

- **answer**：画像版本为 `project_governance_profile_v1`；技术路线冻结为 `TRACK_TYPE_DURABLE_SYSTEM`（Durable System Track：React + Go + PostgreSQL + .proto + ConnectRPC）。
- **direct entry refs**：`GetProjectBrief.governanceProfile.projectProfileVersion` / `.trackType`；`GetGovernanceProfile.profile` 同值；Web 概览区（主验证页"项目治理画像"区，截图 step6-overview）。
- **达标**：是。

### Q2 当前 canonical 根级文件集合是否已被正式承接，在哪个固定入口可确认

- **answer**：已正式承接，9 项根级文件（含 role/required）经表单维护写入画像并持久化。
- **direct entry refs**：`GetProjectBrief.governanceProfile.canonicalRootFiles[]`（9 项）；`GetGovernanceProfile.profile.canonicalRootFiles[]` 同源；Web 表单预填与摘要回看（截图 step2/step7）。
- **达标**：是。

### Q3 当前全局规范资产是否以结构化摘要 + 入口关系被正式承接，在哪个固定入口可确认

- **answer**：已承接，8 项资产全覆盖（name/kind/entry_ref/role），前 5 项摘要型资产填写 `structured_summary`，后 3 项按矩阵允许为空。
- **direct entry refs**：`GetProjectBrief.global_assets[]`（8 项，前 5 项含 structuredSummary）；`GetGovernanceProfile.profile.globalAssetBindings[]` 同源；Web 摘要回看区（截图 step7-summary）。
- **达标**：是。

### Q4 当前 agent 项目简报是否由同一 `repository_id` 驱动，且没有伪造第二套目录扫描结果

- **answer**：是。`GetProjectBrief` 请求唯一参数为 `repository_id`（与 `GetGovernanceProfile` / `GetProjectContext` 同锚），响应 7 顶层块全部由该锚点关联的 PSCO-native facts 派生；字段面无目录扫描、无 Git 状态、无第二套事实源投影。
- **direct entry refs**：实时调用请求体 `{"repository_id":"ca261521-8daf-4248-8f12-43525326e759"}` 与响应 JSON 字段面（本报告 §4.2 表格）；合同源 `proto/psco/project_context/v1/project_context.proto`。
- **达标**：是。

### Q5 当前第一版前端正式承接位是否仍是 `Repository detail`，而没有长出并列第二入口

- **answer**：是。治理画像主内容区唯一存在于 Repository detail；`GovernanceProfileSection` 唯一页面消费入口为 `repository-binding-detail-page.tsx`（L27 import / L306 挂载）；浏览器反回归矩阵 12 页中仅主验证页出现治理画像区。
- **direct entry refs**：`frontend/src/features/repository-binding/pages/repository-binding-detail-page.tsx` L306；`frontend/src/features/governance-profile/index.ts` 冻结注释；浏览器矩阵逐页结果（本报告 §7）。
- **达标**：是。

### Q6 当前阶段是否仍严格没有进入 Git 推进跟踪、模板仓库接入、自动同步与 agent 写回

- **answer**：是。四类非目标均未进入（边界证据见 §8）。
- **direct entry refs**：`proto/psco/governance_profile/v1/governance_profile.proto` RPC 清单（仅 Get/Update 手工维护）；`proto/psco/project_context/v1/project_context.proto` RPC 清单（纯读）；`database/migrations/0010_phase13_governance_profile.sql`（无 git/模板/扫描/同步类表列）；`plan.md` L188 范围约束；`AGENTS.md` §1/§4。
- **达标**：是。

**固定 6 问最终结论：6 / 6 达标。**

## 6. 工具链逐步结果（单值顺序，缺陷修复后完整重跑）

| 步骤 | 目录 | 命令 | 结果 | 非阻断输出归类 |
| --- | --- | --- | --- | --- |
| 1 | `proto/` | `buf build` | 通过（退出码 0） | 无输出、无 warning |
| 2 | `proto/` | `buf lint` | 通过（退出码 0） | 无输出、无 warning |
| 3 | `backend/` | `go test ./...` | 通过（全部 package ok） | 集成包 cache 命中属正常缓存语义，非失败隐藏；review 包（本轮修复所在，无独立测试文件）以 `go build ./...` + `go vet ./internal/review/...` + `GetWeeklyReviewContext` 定向实调 200 三重复核，独立复核代理已重跑 build+vet 确认 |
| 4 | `frontend/` | `npm run build` | 通过（built in 1.45s） | 产物体积正常，无 warning |

说明：首轮工具链（缺陷修复前）同样全部通过；§7 缺陷修复后按 spec 要求完整重跑一遍，结果一致。重跑 `go test` 前已对共享开发库做 pg_dump 备份，重跑后复验 `GetProjectBrief` 与 `GetWeeklyReviewContext` 均 200，画像与 dogfooding 样本未受影响。

## 7. 浏览器反回归矩阵（逐页结果）

| # | 页面 | 加载 | not_found | 治理画像区越界 | 备注 |
| --- | --- | --- | --- | --- | --- |
| 1 | `/repositories/ca261521-...`（主验证页） | 正常 | 无 | —（唯一合法承接位） | 三区齐全、维护按钮可进入编辑态、取消可退回 |
| 2 | `/dashboard` | 正常 | 无 | 无 | 概览 / Current Focus / Recent Activity 正常 |
| 3 | `/modules?statusFilter=all` | 正常 | 无 | 无 | 列表含 project-context-foundation |
| 4 | `/modules/9b02e0ca-...?fromList=true&statusFilter=all` | 正常 | 无 | 无 | 模块详情正常 |
| 5 | `/decisions?statusFilter=all` | 正常 | 无 | 无 | 列表含 phase11 dogfooding 决策 |
| 6 | `/decisions/aa8ee5ad-...?fromList=true&statusFilter=all` | 正常 | 无 | 无 | 决策详情正常 |
| 7 | `/products?statusFilter=all` | 正常 | 无 | 无 | 列表含 PSCO |
| 8 | `/products/f0d034cc-...?fromList=true&statusFilter=all` | 正常 | 无 | 无 | 产品详情正常 |
| 9 | `/repositories` | 正常 | 无 | 无 | 列表含 personal-software-company-os |
| 10 | `/onboarding` | 正常 | 无 | 无 | 引导流程表单正常 |
| 11 | `/reviews/daily` | 正常 | 无 | 无 | Daily Review 正常 |
| 12 | `/reviews/weekly` | 正常（修复后） | 无 | 无 | 首轮失败 → 缺陷修复 → 复验通过（见 §9） |

## 8. 边界证据清单（四类非目标）

### 8.1 不做 Git 推进跟踪主线

- `proto/psco/governance_profile/v1/governance_profile.proto` RPC 仅 `GetGovernanceProfile`（读）+ `UpdateGovernanceProfile`（手工写），无任何 commit/branch/PR/进度类 RPC。
- `proto/psco/project_context/v1/project_context.proto` RPC 均为纯读（`GetProjectBrief` / `GetProjectContext` / `ExportProjectContext`）。
- `database/migrations/0010_phase13_governance_profile.sql` 无 git/commit/branch 类表列（关键字检索零命中）。
- `plan.md` L188：范围约束明确"不得把本 phase 扩写为 Git 推进跟踪平台…"。

### 8.2 不做模板仓库接入 / bootstrap / clone / pull

- 迁移 0010 无 template_repo/bootstrap/clone/pull 类表列（关键字检索零命中）。
- 画像 `template_source` 是自由文本手工字段（本轮 dogfooding 取值 `manual://psco-phase13-dogfooding`），不触发任何仓库接入行为。
- `plan.md` L188 明确"不得…模板仓库自动化…"。

### 8.3 不做自动同步 / 目录全文扫描入库

- `GetProjectBrief` 响应字段面（§4.2）无目录扫描结果；顶层目录矩阵（backend/database/frontend/proto）是前端只读基线常量（`frontend/src/features/governance-profile/data/governance-profile-baseline.ts`），非运行时扫描产物。
- 后端 RPC 清单无 scan/sync 类接口；迁移 0010 无 scan/auto_sync 类表列（关键字检索零命中）。
- `AGENTS.md` §4："全文扫描、模板仓库接入与自动同步不作为本阶段起点"。

### 8.4 不做 agent 写回 / MCP / CLI / Draft / 审批流

- `GetProjectBrief` 为纯读 RPC，无写语义。
- 治理画像唯一写入口是 `UpdateGovernanceProfile`（Web 表单手工维护），无 draft/审批/MCP/CLI 通道。
- `plan.md` L188 明确"不得…MCP / CLI / agent 写回或 IDE 插件"；`AGENTS.md` §4 同口径。

## 9. 失败点、恢复路径与最终 rerun 结果

| # | 过程事件 | 处置 | 最终结果 |
| --- | --- | --- | --- |
| 1 | 固定样本曾因集成测试切库在共享开发库丢失（phase12-11 同款先例） | 执行正式恢复脚本 `restore_phase11_phase12_dogfooding_sample.sh` | 样本在库，`repository_id` 一次解析成功 |
| 2 | 浏览器矩阵首轮 `/reviews/weekly` 显示"Review 上下文加载失败"（`GetWeeklyReviewContext` HTTP 500 空 body） | 定位为 phase08 遗留 nil 解引用缺陷：`review/connect/server.go` 两个转换函数直接解引用 `*time.Time` 可空指针，库中存在"从未被复用模块"（auth-service）即 panic；对齐 `reusesummary/connect` 既有 nil 保护口径做最小修复；同端口 8081 重启后端 | `GetWeeklyReviewContext` 200；浏览器复验正常渲染；工具链三步完整重跑通过；该缺陷属本轮验收发现并修复的真实回归，非 phase13 新增代码引入 |
| 3 | 修复后重跑 `go test ./...` 存在集成测试重置共享开发库的风险 | 重跑前 pg_dump 备份开发库；重跑后复验画像与样本 | 画像与样本未受影响（cache 命中未触发重置），双 RPC 复验 200 |

## 10. 是否达标

- 固定 6 问：6 / 6 达标。
- 工具链：4 步（buf build / buf lint / go test / npm run build）全部通过。
- 人类维护路径 dogfooding：经真实表单操作完成画像创建并回流一致。
- agent 读取路径 dogfooding：brief 实时取证 7 顶层块完整、同源、数组形态、派生正确、无扫描字段。
- 浏览器反回归矩阵：12 / 12 页通过（含 1 页缺陷修复后复验）。
- 边界证据：四类非目标全部留档。
- 过程失败点：全部定位、修复或按协议恢复，无遗留。

**结论：phase13-11 达标，`Project Governance Profile Foundation` 已在同一 `repository_id` 锚点下通过统一工具链、Web / agent 双侧 dogfooding、固定 6 问取证与浏览器反回归验证。**

## 11. Rerun 指引（面向不同执行者）

1. 若固定样本缺失：执行 `database/scripts/restore_phase11_phase12_dogfooding_sample.sh`，再以 `GetProjectContext` 验证 `ca261521-8daf-4248-8f12-43525326e759` 解析。
2. 若画像已存在且需完整重走维护路径：在主验证页进入"维护治理信息"核对既有数据，或经后端确认字段与本报告 §3 一致。
3. agent 侧取证：依次调用 `GetProjectBrief` 与 `GetGovernanceProfile`（请求体仅 `repository_id`），按 §4.2 表格与 §5 六问逐项复验。
4. 工具链：按 §6 表格顺序在 `proto/` → `backend/` → `frontend/` 执行。
5. 浏览器矩阵：按 §7 表格 12 页逐页检查加载 / not_found / 治理画像区越界三点。
