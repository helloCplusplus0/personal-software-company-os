# phase14-10 `Standard Entity Foundation` 联调、dogfooding 与反回归验收报告

> 验收日期：2026-08-18。验收协议沿袭 phase13-11（固定样本 / 固定入口 / 固定问题 / 固定工具链顺序 / 固定 rerun 记录格式）。本报告为 phase14 验证验收类收口任务的冻结结论。

## 1. 固定样本与 id

| 实体 | 固定值 | 状态 |
| --- | --- | --- |
| Repository | `personal-software-company-os`（`ca261521-8daf-4248-8f12-43525326e759`） | 在库（沿袭 phase13-11 冻结值） |
| Product | `PSCO`（`f0d034cc-3235-4d03-879d-6a3111b95b6b`） | 在库（沿袭） |
| Module | `project-context-foundation`（`9b02e0ca-3175-4b5a-bcf4-23cecbb06f72`） | 在库（沿袭） |
| Decision | `phase11 Project Context Foundation dogfooding 验收决策`（`aa8ee5ad-b224-4ea1-b393-7e4c30e42212`） | 在库（沿袭） |
| **Standard（本阶段新增固定样本）** | `Durable System Track 项目范式`（`85a5d8b7-f41a-44ed-8f6f-421e548b53ed`） | 在库，status=active（Task 1 web 新建；Task 6 修复一次 status 偏离后经 web 会话恢复 active，见 §10） |

- Standard 固定样本树语义集合达标：单根 directory（`.`）+ 嵌套 directory（`docs`）+ 6 个 file 节点（AGENTS.md/entry、plan.md/plan、project_rules.md/rules、TECH_STACK_BASELINE.md/baseline、PSCO-summarize-feedback.md/summary、docs README/spec）；role 必填；5 个 `/` 开头 ref + 1 个 `https://` 开头 ref；≥4 节点非空 summary。
- 固定绑定 5 行（八格矩阵 5 合法格全覆盖）：repository→template_source 1 行 + repository/product/module/decision→adopts 各 1 行。
- Revision 3 条（创建后第一次 Update 留痕 / "变更状态" / "恢复固定样本 active 状态（phase14-10 验收口径）"）。
- 正式化路径：**web 新建**（Task 1 盘点时 standards 三表全空，phase14-10 前置 UI 反馈轮的 `899ed9f3` 已因库重置不存在）；迁移产物核对：`SELECT count(*) FROM standards WHERE name = '默认项目范式（迁移自治理画像）'` = 0（验收库无源数据分支，符合 phase14-09 冻结结论）。

## 2. 固定 Web 页面与 agent 入口

- 固定 Web 页面：`/standards`（列表）、`/standards/85a5d8b7-f41a-44ed-8f6f-421e548b53ed`（详情）、`/standards/new`（创建）、`/standards/<id>/edit`（编辑）、`/repositories/ca261521-...`（Repository detail Standard 只读摘要）
- 固定 agent 入口：`POST /api/psco.project_context.v1.ProjectContextService/GetProjectBrief`（repository_id=ca261521 锚点，`standards[]` 全树直读）；辅证 `ListStandards` / `GetStandard` / `ListStandardRevisions`

## 3. 样本恢复方式与 Standard 维护方式

- 四实体基座：缺失时执行既有 `database/scripts/restore_phase11_phase12_dogfooding_sample.sh` 兜底（Task 1 实际使用一次；Task 7 集成测试重置后按备份 ON CONFLICT 逐行恢复一次，见 §10）。
- Standard 与绑定：100% 经 `/standards` web 维护会话补建与正式化（无种子 SQL、无后门脚本）；status 偏离修复亦经 web 编辑表单完成（change_summary 留痕）。

## 4. 双侧 dogfooding 取证

### 4.1 人类维护路径（Web 表单，Task 1 + Task 6 会话）

创建（表单 + 树编辑器填写根级四文档）→ 一次 Update 补全 ref/summary 产生首条 revision → 绑定管理区补 5 行绑定（`template_source` 联动禁用验证：target_type ≠ repository 时该选项禁用，截图 step4-binding-role-disabled.png）→ Revision 回看 → status 偏离修复会话（draft→active，change_summary 留痕）。全程无控制台错误；关键步骤截图归档 `screenshots/step*.png`。

### 4.2 agent 读取路径（brief 实时取证，Task 2）

- `GetProjectBrief`(ca261521) HTTP **200**（两次独立调用）；顶层字段恰好 5 个：`repository / products / modules / decisions / standards`（槽位 2/3/4 reserved）；`governanceProfile / currentPhase / globalAssets` 全部 present=False。
- `standards[]` 含固定 Standard 全树（根 `.` + `docs` + 6 file 节点 role/summary/ref 完整，无转译直读）。
- 同源性：`GetStandard(85a5d8b7)` 的 `directoryTree` 与 brief 内该树规范化 JSON 全等（逐字段一致），6 个标量字段全部 equal。
- 证据文件：`/tmp/brief_task2.json`、`/tmp/getstandard_task2.json`。

## 5. 固定 6 问逐题留档（answer / direct entry refs / 是否达标）

### Q1 当前项目的全局规范目录结构是什么？关键全局文档各在什么角色与入口可直答？

- **answer**：brief `standards[].directoryTree` 直答——根目录四文档（AGENTS.md/entry `/AGENTS.md`、plan.md/plan `/plan.md`、project_rules.md/rules `/project_rules.md`、TECH_STACK_BASELINE.md/baseline `/TECH_STACK_BASELINE.md`）+ 外链共识文档（PSCO-summarize-feedback.md/summary `https://github.com/...`）+ docs/README.md/spec `/docs/README.md`；directory 与 file 节点均带 summary。
- **direct entry refs**：`GetProjectBrief.standards[0].directoryTree`（/tmp/brief_task2.json）；Web `/standards/85a5d8b7` 详情页树形展示同源。
- **达标**：是。

### Q2 画像旧两清单信息（canonical 根文件 / 全局资产）现由什么承接，是否有丢失？

- **answer**：由 `standards[]` 单一来源承接——canonical 根文件经树内根级 file 节点（entry/plan/rules/baseline 角色 + `/` ref）承接；全局规范资产经树形结构 + summary + 外链 ref 承接；phase14-09 一次性验证库证据（migration_evidence.md §3 恰 1 条合并产物 + §5 零丢失对照表）复核在位。
- **direct entry refs**：brief 树 JSON 节点清单；`.trae/specs/phase14_09_retire_governance_profile_switch_brief/migration_evidence.md` §3/§5；T1-T7 复验（§8 裁决④）。
- **达标**：是。

### Q3 一条 Standard 能否同时被 repository / product / module / decision 消费，绑定语义如何区分？

- **answer**：能。固定样本 5 行绑定实测：repository→template_source（仓库即结构导航源，仅 repository 合法）+ 四类目标各 1 行 adopts（全开）；brief 以 repository 锚点消费 standards[]。
- **direct entry refs**：`SELECT target_type, role FROM standard_bindings WHERE standard_id='85a5d8b7...'`（5 行，target_id 与固定四实体 id 全匹配）；详情页绑定管理区 5 行（截图 task6-p02.png）。
- **达标**：是。

### Q4 Standard 每次演进是否留痕，如何回看？

- **answer**：是。`ListStandardRevisions` 返回 3 条，change_summary 均为人工一句话；Web 详情页 Revision 历史区可回看。
- **direct entry refs**：ListStandardRevisions 响应（"变更状态" / "phase14-10 dogfooding：补全关键文档 ref 与摘要..." / "恢复固定样本 active 状态（phase14-10 验收口径）"）；截图 task6-p02.png。
- **达标**：是。

### Q5 规范维护的唯一入口在哪，agent 是否仍无写回通道？

- **answer**：唯一维护入口为 `/standards` 四页（前端切片 owner 唯一）；`StandardService` 写 RPC 调用 100% 位于 `frontend/src/features/standard/`；backend 无 MCP / CLI / agent 写回承接位。
- **direct entry refs**：grep 写 RPC 调用文件清单（application/ 5 个 mutation owner + binding-panel + 3 pages，均在切片内）；backend grep 零命中。
- **达标**：是。

### Q6 brief 与库内是否还有画像残余？时间轴能力是否未被偷渡承接？

- **answer**：无残余。brief 无 governance_profile/current_phase 块（字段号 2/4 + 名 reserved）；`governance_profiles` 主表 `to_regclass` NULL（0012 已应用）；`backend/internal/governanceprofile` 不存在；migrations 仅到 0012，无时间轴类新表列/新 RPC/活字段。
- **direct entry refs**：brief JSON present=False 取证；`schema_migrations` 含 0012；T6/T7 复验（§8）；proto 仅 1 处裁决留痕注释命中。
- **达标**：是。

**6/6 全达标。**

## 6. 工具链逐步结果（单值顺序完整执行）

| 步骤 | 命令 | 退出码 | 结果 |
| --- | --- | --- | --- |
| 0（保护） | `pg_dump` 备份验收库 | 0 | `/tmp/psco_dev_backup_phase14_10_task5.sql`（32595B，未动用） |
| 1 | `proto/ buf build` | 0 | 无 warning |
| 2 | `proto/ buf lint` | 0 | 无 warning |
| 2b（专项） | `proto/ make breaking`（`buf breaking --against '../.git#branch=main,subdir=proto'`） | 0 | 豁免留痕生效：`breaking.ignore` 单文件（governance_profile.proto）+ `breaking.ignore_only: FIELD_WIRE_JSON_COMPATIBLE_TYPE` 规则级单文件（project_context.proto），无目录前缀扩大化 |
| 3 | `backend/ go build ./... && go vet ./... && go test ./...` | 0 | 9 测试包 ok 0 FAIL（8 包缓存命中——代码自 Task 7 门禁实跑〔集成 18.6s 全绿〕后未变更，缓存结果有效；standard/service 实跑 0.194s） |
| 3b（复验） | 固定样本复验 SQL | — | std_ok=1 / bindings=5 / revisions=3 / 四基座各 1——全达标，未触发恢复 |
| 4 | `frontend/ npx tsc -b && npm run build` | 0 | 零错误；vite build 2462 modules 1.35s（含 standard 全部页面产物与 UI 反馈轮两处修复的编译回归） |

## 7. 浏览器反回归矩阵（16 页逐页结果）

| # | 页面 | 加载 | 专属检查点 | console error | 结论 |
| --- | --- | --- | --- | --- | --- |
| 1 | `/standards` | 正常 | 列表含固定样本行且状态 active | 无 | PASS |
| 2 | `/standards/85a5d8b7` | 正常 | 树形（根+docs+6 file role/summary/ref）/ 绑定 5 行 / Revision ≥3 条含恢复条目 / 一致性布局五要素（返回行/语义导语/左摘要右内容 grid）齐全 | 无 | PASS |
| 3 | `/standards/new` | 正常 | 创建表单 + 树编辑器可用 | 无 | PASS |
| 4 | `/standards/<id>/edit` | 正常 | 整树编辑在位；节点名称输入框连续输入不跳焦（UI 反馈修复回归通过）；无拖拽交互 | 无 | PASS |
| 5 | `/repositories/ca261521` | 正常 | 画像区已让位（无"维护治理信息"入口）；Standard 只读摘要正常加载，compact 树正确（T7 解耦回归：无 404） | 无 | PASS |
| 6 | `/dashboard` | 正常 | 概览/Current Focus/Asset Feedback/Recent Activity 正常；无 Standard 主卡片 | 无 | PASS |
| 7 | `/modules` | 正常 | 列表含 project-context-foundation | 无 | PASS |
| 8 | `/modules/9b02e0ca?fromList=true&statusFilter=all` | 正常 | 详情正常 | 无 | PASS |
| 9 | `/decisions` | 正常 | 列表正常 | 无 | PASS |
| 10 | `/decisions/aa8ee5ad?fromList=true&statusFilter=all` | 正常 | 详情正常 | 无 | PASS |
| 11 | `/products` | 正常 | 列表含 PSCO | 无 | PASS |
| 12 | `/products/f0d034cc?fromList=true&statusFilter=all` | 正常 | 详情正常 | 无 | PASS |
| 13 | `/repositories` | 正常 | 列表正常 | 无 | PASS |
| 14 | `/onboarding` | 正常 | 引导流程正常 | 无 | PASS |
| 15 | `/reviews/daily` | 正常 | Current Focus/Pending Decisions/Representative Signals 完整 | 无 | PASS |
| 16 | `/reviews/weekly` | 正常 | 系统概览/模板候选/派生智能提示/最近活动完整 | 无 | PASS |

全部页面无 not_found、无画像残留文案。截图归档：`screenshots/task6-p01.png ~ task6-p16.png`（16 张）+ Task 1 会话 `step*.png`。

## 8. 八项裁决验收门禁矩阵（shared_baseline §4）+ 画像退役触点复验（T1-T7）

### 8.1 画像退役触点（T1-T7 + 裁决⑧）

| 触点 | 断言与证据 | 结论 |
| --- | --- | --- |
| T1 proto | `proto/psco/governance_profile/v1/` 不存在；buf.yaml 两条豁免均精确单文件无扩大化；生成产物命中仅为 project_context 裁决注释与 reserved 名编码残留（protoc 正常编码） | PASS |
| T2 存储 | 三张画像表（主表+两 bindings）`to_regclass` 均 NULL（主表经 0012 drop） | PASS |
| T3 后端 | `backend/internal/governanceprofile` 不存在；grep 零命中；platform 层仅 router.go:433 一行裁决留痕注释，无挂载 | PASS |
| T4 RPC | 画像 Connect 请求 404（对照 ListStandards 200 / healthz 200） | PASS |
| T5 前端 | `features/governance-profile/` 不存在；grep 零命中 | PASS |
| T6 brief | 5 顶层块（槽位 2/3/4 reserved）；无画像派生消息（proto 仅 1 处裁决注释命中） | PASS |
| T7 画像残余（裁决触发） | proto L198-199 双 reserved（号+名）+ 4 画像派生消息删除；schema_migrations 含 0012；豁免无新增（号+名双 reserved 已满足 WIRE_JSON 规则）；Repository detail Standard 摘要经 brief 正常加载 | PASS |
| 裁决⑧验收库形态 | 0011 `psql -f` 重放 exit=0 全幂等 NOTICE；迁移种子 count=0；standards 总数 1 固定样本完好；migration_evidence.md §3/§5 章节存在 | PASS |

### 8.2 八项裁决矩阵

| 裁决 | 取证摘要 | 结论 |
| --- | --- | --- |
| ① 混合式颗粒度 | brief 树 JSON 仅六字段无正文承载；directory（docs）与 file 节点均带 summary；standard-tree-view.tsx L55-76 双侧渲染 | PASS |
| ② 主表+jsonb树+多态绑定 | `directory_tree jsonb NOT NULL`；四元组唯一约束 + target_type CHECK 四类；4 类实例行（1/1/1/2） | PASS |
| ③ 正文零托管 | DirectoryTreeNode 仅 name/node_type/role/summary/ref/children；backend/internal/standard 抓取类符号零命中 | PASS |
| ④ 画像系统性退役 | T1-T7 全绿（§8.1） | PASS |
| ⑤ 无第五主实体扩散 | dashboard 目录零命中；交叉引用仅 Repository detail 只读摘要；`/standards` 独立导航项 | PASS |
| ⑥ 树编辑器无拖拽 | 拖拽符号零命中；交互原语 handleAddRootChild/replaceChild/removeChild/swapChildren 按钮式；编辑会话全程经增删/上移下移完成 | PASS |
| ⑦ 绑定仅 Standard 详情页发起 | bind/unbind caller 4 文件全在 features/standard；挂载点唯一（standard-detail-page L182-184）；四实体页面零绑定入口 | PASS |
| ⑧ 存量合并迁移 | 两库形态复核通过（§8.1 裁决⑧行）；验收库无第二条迁移产物 | PASS |

## 9. 边界证据清单（非目标）

- **不做 Git 推进跟踪 / 模板仓库自动接入 / 自动同步**：backend 无 clone/pull/同步类代码（Q5/Q6 grep 取证零命中）；brief 无目录扫描与 Git 状态字段。
- **不做正文托管 / 全文快照 / 版本分支**：树节点无正文字段（裁决③取证）；ref 仅定位引用；revision 为演进留痕非全文快照。
- **不做第五 CRUD 主实体化**：Dashboard 无 Standard 主卡片；四实体页面无 Standard CRUD 侵入（裁决⑤取证）。
- **不做 agent 写回 / MCP / CLI**：写 RPC 仅 web 前端切片调用（Q5 取证）；backend 无写回承接位。
- **CON-08 时间轴口径变更留痕**：dev_plan §4 原表述"`current_phase` 三字段仅以 brief 内联轻量消息保留"已被 2026-08-18 用户裁决取代——brief 内联轻量消息（`BriefGovernanceProfile`/`BriefCurrentPhase` 等 4 消息）随 T7 整体删除并 reserved；phase15 如进入时间轴另行新建正规字段，不复活画像派生形态（本 spec MODIFIED Requirements 冻结）。

## 10. 失败点、恢复路径与最终 rerun 结果（沿袭 phase13-11 §9 格式）

| # | 失败点 | 根因 | 恢复路径 | 最终 rerun 结果 |
| --- | --- | --- | --- | --- |
| 1 | Task 2 首验 `GetProjectBrief` 404（画像主记录行丢失且无重建入口） | brief 残余依赖画像读取（phase12-13 遗留设计，phase14 裁决画像为错误设计） | **T7 裁决触发修复**：proto 双 reserved + 4 画像派生消息删除 + `governanceprofile` 模块整删 + `0012` drop 主表 + 三端再生成；同端口重启 | 门禁完整重跑全绿（buf build/lint/breaking 0 + go build/vet/test 0〔集成 18.6s〕+ tsc/build 0）；`GetProjectBrief` 200，brief 5 顶层块 |
| 2 | Task 5 前置（Task 7 期间）go test 集成重置清空四实体基座 | 集成测试库重置行为 | pg_dump 备份 → `INSERT ... ON CONFLICT DO NOTHING` 逐行恢复（COPY 整段重放因同名异 id fixture 冲突不可用） | 四基座 1/1/1/1 + standard 1 + bindings 5 + revisions 1 恢复；后续样本复验全达标 |
| 3 | Task 2 取证发现固定 Standard status=draft（违反 spec 语义集合 4） | Task 1 会话后一次"变更状态"编辑产生偏离 | web 编辑表单正式修复（change_summary="恢复固定样本 active 状态（phase14-10 验收口径）"，产生第 3 条 revision） | 详情页确认 active；Task 6 p01/p02 复验通过 |

## 11. 是否达标

**达标。** DoD 对照（dev_plan L71）：

- 固定问题全达标：Q1-Q6 = 6/6（§5）
- 八项裁决验收门禁全绿：①-⑧全 PASS（§8）
- 反回归矩阵全绿：16/16 页 PASS（§7）
- acceptance_report 冻结验收结论：本报告
- 独立复核通过：见 §12 复核记录

## 12. 独立复核记录

见本目录 `independent_review.md`（子代理独立复核：样本协议一致性 / 六触点证据真实性 / 八项裁决取证完整性 / 工具链顺序与结果 / 矩阵覆盖完整性 / 范围外改动检查）。

## 13. Rerun 指引（面向不同执行者）

1. 启动环境：PostgreSQL（docker `rento-preview-postgres`，库 `psco_development`）→ 后端 `cd backend && go run ./cmd/server`（8081）→ 前端 `cd frontend && npm run dev`（5173）。
2. 固定样本核验：`GetProjectContext` / `GetProjectBrief`（repository_id=ca261521-8daf-4248-8f12-43525326e759）应 200 且 standards[] 含 85a5d8b7 全树；四实体缺失时执行 `bash database/scripts/restore_phase11_phase12_dogfooding_sample.sh`。
3. 固定 6 问取证：按 §5 问题与入口复测（brief JSON + psql 绑定行 + revisions）。
4. 触点复验：按 §8.1 命令逐条执行（目录不存在 / to_regclass NULL / grep 零命中 / 画像 RPC 404）。
5. 工具链：按 §6 单值顺序（步骤 0 备份 → buf build → lint → breaking → go build/vet/test → 样本复验 → tsc/build）。
6. 浏览器矩阵：按 §7 逐页访问 16 URL 核对专属检查点。
7. 样本基线：standards 1 行（85a5d8b7/active）+ bindings 5 行 + revisions ≥3 条；偏离时经 `/standards` web 会话修正（非 SQL 后门）。
