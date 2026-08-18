# phase14-10 独立复核报告

> 复核日期：2026-08-18。复核人：独立复核子代理（只读取证 + 重执行验证命令；除本报告外未修改任何文件）。
> 复核对象：本目录 spec.md / tasks.md / checklist.md / acceptance_report.md；上游 dev_plan L67-71、shared_baseline §4。
> 复核环境实测基线：PostgreSQL（`rento-preview-postgres` / `psco_development`）、后端 8081、前端产物重验，均在线实测。

## 复核结论：PASS（阻断问题 0 项 / 观察项 4 项）

R1-R7 全部 PASS。所有关键断言经命令重执行独立复证，未发现虚假证据、范围外改动或矩阵缺口。验收结论"达标"成立，可进入 Task 8.3 收口（勾选 checklist / tasks 收尾）。

---

## R1 样本协议一致性 — PASS

### 命令与结果

- `docker exec rento-preview-postgres psql -U rento -d psco_development -c "SELECT id, name, status FROM standards;"`
  → 恰 1 行：`85a5d8b7-f41a-44ed-8f6f-421e548b53ed` / `Durable System Track 项目范式` / `active` ✓
- `curl -s -X POST http://localhost:8081/api/psco.standard.v1.StandardService/GetStandard -d '{"standard_id":"85a5d8b7-..."}'`
  → `"status":"STANDARD_STATUS_ACTIVE"` ✓；树结构实测：
  - **语义集合 1**（单根 directory `.` + 嵌套 directory `docs`）✓
  - **语义集合 2**（file 节点 6 个 ≥5：AGENTS.md / plan.md / project_rules.md / TECH_STACK_BASELINE.md / PSCO-summarize-feedback.md / docs/README.md；四项根级关键文档全在；role 全必填：entry / plan / rules / baseline / summary / spec）✓
  - **语义集合 3**（summary + `/` ref 同时具备的 file 节点恰 4 个：AGENTS.md / plan.md / project_rules.md / TECH_STACK_BASELINE.md，均非空 summary + `/` 开头 ref，满足 ≥4；`https://` ref 1 个：PSCO-summarize-feedback.md → github.com 外链）✓
  - **语义集合 4**（status=active 且树含 file 节点）✓
- 绑定行数：`SELECT count(*) FROM standard_bindings WHERE standard_id='85a5d8b7...'` → **5** ✓；GetStandard 响应 bindings 数组逐行核对：repository→template_source 1 行 + repository / product / module / decision→adopts 各 1 行，target_id 与四实体冻结 id 全匹配（八格矩阵 5 合法格全覆盖）✓
- revisions：`SELECT count(*) FROM standard_revisions ...` → **3 条（≥1）** ✓；change_summary 依次为"phase14-10 dogfooding：补全关键文档 ref 与摘要，固化首条 revision 留痕" / "变更状态" / "恢复固定样本 active 状态（phase14-10 验收口径）"，均为人工一句话，与 report §1 / Q4 留档一致 ✓
- 固定 id 留档一致性：acceptance_report §1 = tasks.md Task 1 执行记录 = 实测库值 = `85a5d8b7-f41a-44ed-8f6f-421e548b53ed` ✓；正式化路径（web 新建、`899ed9f3` 因库重置不存在）两处留档一致 ✓

### 结论

report §1 对 spec 固定样本语义集合 1-4 的声明与库内/API 实测完全一致。

---

## R2 六触点证据真实性（T1-T7 + 裁决⑧）— PASS

### 命令与结果（全部重执行）

| 触点 | 复核命令 | 实测结果 | 判定 |
| --- | --- | --- | --- |
| T2 | `SELECT to_regclass('governance_profiles'), to_regclass('governance_profile_repository_bindings'), to_regclass('governance_profile_bindings')` | 三者均 NULL | ✓ |
| T4 | 画像 RPC curl + 对照 | `GetGovernanceProfile` **404**；`ListStandards` **200**；`healthz` **200** | ✓ |
| T7 proto | 读 `project_context.proto` L193-199 | `reserved 2, 3, 4;` + `reserved "governance_profile", "global_assets", "current_phase";` 号+名双 reserved 实际在位（L198-199） | ✓ |
| T7 迁移 | `SELECT version FROM schema_migrations ORDER BY version` | 含 `0012_phase14_brief_profile_decoupling`，且最高版本即 0012（无时间轴类新迁移，Q6 断言成立） | ✓ |
| T7 豁免 | 读 `proto/buf.yaml` | `breaking.ignore` 仅 governance_profile.proto 单文件；`ignore_only: FIELD_WIRE_JSON_COMPATIBLE_TYPE` 仅 project_context.proto 单文件；无目录前缀扩大化 | ✓ |
| 裁决⑧ | `SELECT count(*) FROM standards WHERE name='默认项目范式（迁移自治理画像）'` | **0** ✓；`migration_evidence.md` §3（L107"一次性算法验证库执行"）与 §5（L261"零丢失对照表"）章节实际存在 | ✓ |
| brief 字段面 | GetProjectBrief(ca261521) 实时调用 + JSON 解析 | HTTP 200；顶层 keys 恰 5 个：`decisions / modules / products / repository / standards`；`governanceProfile / currentPhase / globalAssets` 均 absent；standards[] 含固定样本全树 | ✓ |
| T1/T3/T5（附） | 目录 ls + grep | `proto/psco/governance_profile/`、`backend/internal/governanceprofile/`、`frontend/src/features/governance-profile/` 三目录均不存在；`grep -rn "governanceprofile" backend/internal` 与 `grep -rn "governance-profile" frontend/src` 均**零命中** | ✓ |

### 备注

T3 实测零命中，acceptance_report §8.1 T3 行写"仅 router.go:433 一行裁决留痕注释"——实测 router.go:433 注释现为中文措辞（"buildGovernanceProfile 已于 2026-08-18 phase14-10 T7 用户裁决删除……"），不含 `governanceprofile` 字面。该注释确实存在于 L433，报告描述的是 Task 3 复验时点快照；Task 7 注释改写后现状更干净。实质断言（无挂载 / 无残留 import）成立，列入观察项 2。

---

## R3 八项裁决取证完整性 — PASS

- acceptance_report §8.2 八行逐行审查：每行均含具体证据摘要（文件+行号 / 符号名 / 行数分布 / SQL 断言），非空泛表述 ✓
- ④（"T1-T7 全绿（§8.1）"）与 ⑧（"两库形态复核通过（§8.1 裁决⑧行）"）正确汇入 §8.1，无重复取证亦无缺失 ✓
- 裁决⑤抽查：`grep -rn "standard" frontend/src/features/dashboard/ -i` → **零命中**（无 Standard 主卡片挂载）✓
- 裁决⑦抽查：`grep -rn "useBindStandard|useUnbindStandard" frontend/src` → 命中文件全部位于 `frontend/src/features/standard/`（use-bind-standard.ts / use-unbind-standard.ts 两个固定承接位 + standard-binding-panel.tsx 唯一消费组件），切片外零命中 ✓
- 其余裁决（①②③⑥）证据摘要中的关键断言（树 JSON 字段面 / jsonb 列 / 四元组唯一约束 / 拖拽符号零命中）与本轮 R2 实测的 proto/库表/目录状态相互印证，无矛盾 ✓

---

## R4 工具链顺序与结果 — PASS

### 顺序合规性

report §6 步骤顺序：0 备份 → 1 `buf build` → 2 `buf lint` → 2b `make breaking` → 3 `go build && go vet && go test` → 3b 固定样本复验 → 4 `tsc -b && npm run build`，符合 spec"单值顺序"（buf build → lint → breaking → go → tsc/build）且附加保护步（0/3b）均沿袭 phase13-11 保护协议 ✓

### 本轮全量独立重验（绕过缓存）

| 步骤 | 命令 | 退出码 |
| --- | --- | --- |
| buf build | `cd proto && buf build` | **0** |
| buf lint | `buf lint` | **0** |
| buf breaking | `make breaking`（`--against '../.git#branch=main,subdir=proto'`） | **0** |
| go | `go build ./... && go vet ./...` | **0 / 0** |
| go test | `go test ./... -count=1`（**强制实跑、绕过全部缓存**） | **0**（projectcontext/connect 18.812s 集成实跑 ok、standard/service 0.178s ok，0 FAIL） |
| tsc | `npx tsc -b` | **0** |
| vite | `npm run build` | **0**（built in 1.51s） |

### 缓存命中说明如实性判定

report §6 步骤 3 声明"8 包缓存命中——代码自 Task 7 门禁实跑后未变更故缓存有效"。本轮以 `-count=1` 强制全量实跑同样全绿，证明缓存命中不构成失败掩盖；缓存有效性前提（代码未再变更）与 git 变更面（R6：Task 5 后无新增代码改动，仅工件文档更新至 17:51-17:55）相互印证。**说明如实。**

### 复核过程库保护记录（透明留档）

本轮 `go test -count=1` 后四实体基座被集成测试清空（repo/prod/mod/dec = 0/0/0/0），**standards 三表完好（std_ok=1 / bindings=5 / revisions=3）**——该行为与 tasks.md Task 7 失败点 2 记录的集成测试行为一致。已按验收协议执行官方恢复脚本 `restore_phase11_phase12_dogfooding_sample.sh` 重建基座（四 id 与冻结值一致），复验最终库形态：四基座 1/1/1/1 + std_ok=1 + bindings=5 + revisions=3，`GetProjectBrief` 复测 200。复核前已先行 pg_dump 备份（/tmp/psco_dev_backup_independent_review.sql，32595B——与 report §6 步骤 0 备份同字节数，佐证库形态自 Task 5 后未漂移）。

---

## R5 矩阵覆盖完整性 — PASS

- report §7 逐页核对 spec 16 页矩阵：#1-#16 页面 URL 与 spec 完全一致（`/standards` 四页 + Repository detail + dashboard + 四实体列表/详情 7 页 + onboarding + reviews 双路径），每页有专属检查点结果 + console error 列 + PASS 结论，无缺页 ✓
- 截图存在性：`ls screenshots/task6-p*.png` → **16 张全在**（task6-p01.png ~ task6-p16.png）✓；另 Task 1 web 维护会话 step*.png 32 张在档（含 step4-binding-role-disabled.png 联动禁用验证截图，与 report §4.1 引用一致）✓

---

## R6 范围外改动检查 — PASS

### `git status --porcelain` 全量清单与逐类判定（19 项）

| 变更 | 类别 | 判定 |
| --- | --- | --- |
| `M proto/psco/project_context/v1/project_context.proto` | T7 双 reserved + 消息删除 + 裁决注释 | phase14 交付面 ✓ |
| `D backend/internal/governanceprofile/`（5 文件：errors.go / profile_store.go / 集成测试 / query_service.go / types.go） | T7 整目录删除 | phase14 交付面 ✓ |
| `M backend/internal/connecterrors/connect_errors.go` | T7 删画像哨兵错误 | phase14 交付面 ✓ |
| `M backend/internal/platform/router.go` / `server.go` | T7 去装配（router.go:433 裁决留痕注释在位） | phase14 交付面 ✓ |
| `M backend/internal/projectcontext/`（types / candidate / connect / service / 集成测试 5 文件） | T7 brief 解耦收缩 | phase14 交付面 ✓ |
| `M backend/internal/standard/service/service_integration_test.go` | git diff 实查：仅 5 行注释措辞改写（governanceprofile 字面 → 中文描述），T7 grep 清理联动 | phase14 交付面 ✓ |
| `M frontend/src/features/standard/components/standard-tree-editor.tsx`（5 行） | UI 反馈轮焦点修复 | 前置已确认范围 ✓ |
| `M frontend/src/features/standard/pages/standard-detail-page.tsx`（210 行） | UI 反馈轮一致性布局 | 前置已确认范围 ✓ |
| `?? database/migrations/0012_phase14_brief_profile_decoupling.sql` | T7 迁移 | phase14 交付面 ✓ |
| `?? .trae/specs/phase14_10_validate_standard_entity_integration_dogfooding_regression/` | 本任务三件套 + acceptance_report + screenshots + 本报告 | phase14 交付面 ✓ |

- 未越出 phase14 范围的文件：**0**；phase14-07/08/09 交付物（standard 后端/前端主线、0011 迁移、proto standard 包）已在 main 基准内，当前未提交增量即 phase14-10 全部变更面，与 report 口径一致 ✓
- 散装文档/孤立 spec 检查：未跟踪文件仅上述两项（0012 + phase14_10 目录）；phase14_10 目录内仅 spec / tasks / checklist / acceptance_report / screenshots / 本报告，无意外文件；无根级散装文档 ✓

---

## R7 工件完整性 — PASS

- **tasks.md**：Task 1-7 全部勾选且每任务附实质执行记录（命令结果 / id / 时间线 / 证据文件路径）；Task 8 未勾——8.1（报告冻结）实质已完成，8.2（独立复核）即本报告，8.3（全量勾选收口）待复核通过后执行，时序合理 ✓
- **checklist.md**：检查项结构与 spec Requirements 一一对应（固定样本 7 项 / agent dogfooding 5 项 / 触点 9 项 / 八裁决 8 项 / 工具链 6 项 / 矩阵 7 项 / Task 7 3 项 / 边界 1 项 / DoD 7 项）；**当前全部未勾选**（如实记录——按任务约定不作为 FAIL 依据，属 Task 8.3 待办，见观察项 1）✓
- **acceptance_report.md** 章节结构对照 spec"必须冻结"条目：§1 固定样本与 id / §2 固定入口 / §3 样本恢复与维护方式 / §4 双侧 dogfooding / §5 固定 6 问（answer + direct entry refs + 是否达标，6/6）/ §6 工具链逐步结果 / §7 浏览器矩阵逐页 / §8 八项裁决门禁矩阵（+T1-T7 触点表）/ §9 边界证据（四类非目标 + CON-08 留痕）/ §10 失败点-恢复-rerun / §11 是否达标（DoD 对照）/ §12 独立复核记录 / §13 Rerun 指引——**全部条目覆盖，无缺章** ✓

---

## 阻断问题清单

无。

## 观察项清单（不阻断）

1. **checklist.md 尚未勾选**（Task 8.3 待办）：43 个检查框当前均未勾。建议收口时逐项勾选并（如需）附一行结果引用（指向 acceptance_report 对应章节），保持与 tasks.md 执行记录同等留痕强度。
2. **acceptance_report §8.1 T3 行与当前代码状态存在时点差**：报告写"仅 router.go:433 一行裁决留痕注释"命中；Task 7 注释措辞改写后当前实测 `grep governanceprofile` 零命中（注释在位但不含字面）。实质断言（无挂载/无残留 import）成立且现状更干净。建议收口时将该行口径微调为"裁决留痕注释在位（中文措辞，不含字面），grep 零命中"，非必须。
3. **裁决⑦ caller 计数口径**：report/tasks 写"4 文件"，本轮 grep 直接命中 3 文件（两个 mutation 承接位 + binding-panel；第 4 应为挂载点 standard-detail-page）。计数口径差异不影响"caller 全在切片内"的实质断言。建议无需改动，或统一口径说明。
4. **go test 全量重跑会清空四实体基座**（standards 三表不受影响）：本轮复核实测证实该行为与 Task 7 记录一致，恢复脚本兜底有效。report §13 Rerun 指引第 2 条已涵盖恢复操作；建议后续任何执行者重跑 `go test ./...` 后必须复验四基座并按协议恢复（本报告 R4 节留档了完整操作先例）。

---

## 复核方法与证据强度说明

- 本轮全部关键断言以命令重执行独立复证：psql 直查 5 组、curl 实测 4 个端点（含 404 对照）、grep/ls 目录级断言 6 组、proto/buf.yaml/migration_evidence 文件级核读、工具链七步全量重跑（含 `-count=1` 绕缓存）。
- 复核期间唯一写操作为本报告文件与验收库基座恢复（协议内兜底脚本，已透明留档于 R4 节）；未修改任何项目源码或验收工件。
