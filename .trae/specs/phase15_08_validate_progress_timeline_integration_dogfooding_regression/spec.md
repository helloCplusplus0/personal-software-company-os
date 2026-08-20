# phase15-08 完成进度时间轴的联调、dogfooding 与反回归验证 Spec

## Why

`phase15-06`（后端主线）与 `phase15-07`（前端主线，含三轮 UI 返工）已交付可运行的进度时间轴最小主线，但尚未在冻结协议下完成整体验收：dogfooding 固定录入集（PSCO 自身 phase14 历史回放）尚未录入、固定问题取证与十一项裁决验收门禁（shared_baseline §4）未逐条验证、浏览器反回归矩阵未跑、验收结论未冻结。本子任务是 phase15 的验证验收类收口：零产品代码改动，产出 acceptance_report 与独立复核记录，作为 phase15-09 根级收口的唯一直接上游。

## What Changes

- 环境准备：执行 `database/scripts/restore_phase11_phase12_dogfooding_sample.sh` 恢复固定样本仓库（`ca261521-8daf-4248-8f12-43525326e759`，当前开发库缺失）；清理 `main-repo` 上 phase15-07 用户手录的 5 条 phase01 测试事件（API 删除，留档清理记录——phase15-07 tasks.md 返工记录三的数据说明承诺）
- dogfooding 固定录入集录入：在本 spec §附件 A 冻结的 16 条事件（1 phase_started + 11 task_completed + 1 phase_completed + 3 note）经 **web 表单路径**（真实浏览器驱动，沿 phase15-07 驱动模式）录入到固定样本仓库——PSCO 用自己的进度时间轴讲完自己的 phase14 故事
- 固定问题取证：agent 经 brief 直答"背景 + 进度"（一次调用含 repository/standards[] + progress）；`ListProgressEvents` 全量与三轨过滤；append-only 断言（无 Update / 历史零丢失 / 派生不落库）；派生正确性（phase14 完结态 / latest=phase14-11 / recent N=10）
- 十一项裁决验收门禁矩阵逐条验证（shared_baseline §4 冻结清单 ①~⑪）
- 工具链门禁：`proto/` 目录 `make lint && make build && make breaking`；`backend/` 目录 `go build ./... && go vet ./... && go test ./...`；`frontend/` 目录 `tsc --noEmit`
- 浏览器反回归矩阵：Repository detail（进度区三轮 UI 返工后最终形态 + Standard 摘要区协排 + 工作台四卡片布局）+ 既有页面抽查（沿 phase14-10 16 页矩阵抽样）
- 产出：`acceptance_report.md`（沿袭 phase14-10 十三节协议：固定样本 / 固定入口 / 恢复方式 / 双侧 dogfooding 取证 / 固定问题留档 / 工具链 / 反回归矩阵 / 裁决门禁矩阵 / 边界证据 / 失败恢复 rerun / 是否达标 / 独立复核 / rerun 指引）+ 独立复核子代理记录
- 不做（边界冻结）：零产品代码改动（发现阻断性缺陷 → 记录并按 fix 流程或回 phase15-07 返工，不在本子任务内修）；零根级文档回写（phase15-09）；零 schema_migrations / 种子脚本新增（复用既有恢复脚本）

## Impact

- Affected specs:
  - `docs/phase/phase15_project_progress_timeline_foundation_dev_plan.md` L58-61（本子任务范围与 DoD 唯一定义）
  - `docs/phase/phase15_project_progress_timeline_foundation_shared_baseline.md` §4（验收前提与十一项裁决门禁清单——本 spec 直接执行对象）
  - `phase15-06` / `phase15-07`（被验收对象：后端 3 RPC + 前端切片 + 三轮 UI 返工成果）
  - `phase14-10` acceptance_report（验收协议格式先例——沿袭其十三节结构）
  - `phase15-09`（下游：以本 spec acceptance_report 为唯一直接上游做根级收口）
- Affected code: **零源代码改动**。涉及数据：恢复固定样本仓库（seed 复用）；删除 main-repo 上 5 条测试事件；在 ca261521 录入 16 条 dogfooding 事件（验收后保留为 dogfooding 冻结数据，rerun 指引给出重放方法）
- 运行时约束：用户已重启前后端（后端 :8081 / vite :5173 运行中）；本子任务禁止重启任何服务器；浏览器驱动沿用 phase15-07 模式（playwright-core + 本地 chromium 二进制，/tmp/psco-browser-test）

## ADDED Requirements

### Requirement: 验收环境必须先恢复固定样本并清理测试残留

本子任务 SHALL 在取证前完成两项环境操作并留档：

1. **恢复固定样本仓库**：执行 `./database/scripts/restore_phase11_phase12_dogfooding_sample.sh`（幂等 seed，不清空现有开发数据），确认 `ca261521-8daf-4248-8f12-43525326e759`（`psco-fixed-sample`）仓库及其 canonical 关系（product/module/standard 绑定）就位——shared_baseline §4 验收环境前提
2. **清理测试残留**：经 `DeleteProgressEvent` API 删除 `main-repo`（1a863cb1）上 phase15-07 用户手录的 5 条 phase01 测试事件，确认 `main-repo` progress_events 归零、全库 progress_events 仅剩即将录入的 dogfooding 数据；清理记录（删除的 5 条 id 与 title）写入 acceptance_report §1

#### Scenario: 环境就绪判定

- **WHEN** 两项操作完成
- **THEN** `ca261521` 可被 GetProjectBrief 正常应答（repository + standards[] 非空——背景侧就绪）；`main-repo` ListProgressEvents 返回空；progress_events 全库 0 行（录入前）

### Requirement: dogfooding 固定录入集必须按附件 A 逐字录入且经 web 表单路径

本子任务 SHALL 按本 spec **§附件 A（dogfooding 固定录入集明细，16 条）** 经真实浏览器 web 表单路径（phase15-07 驱动模式）录入全部事件到 `ca261521`：

- **录入路径 = web 表单**（人类维护路径取证）：真实浏览器驱动 ProgressEventForm 逐条提交（workflow_type/event_kind/task_key/title/detail/evidence_ref/occurred_at 按附件 A 逐字；occurred_at 为补录历史——datetime-local 填附件 A 的本地时刻，DP-3 转换）；不使用 API 直插
- **录入顺序 = 附件 A 行序**（narrative 顺序：phase_started → 01 → note → …→ 11 → phase_completed → note）；同 occurred_at 相邻事件的展示序由 created_at DESC tiebreak 保证（录入顺序即展示顺序）
- **title 取 git commit message 子任务名**；note 标题与 detail 按附件 A；evidence_ref 按附件 A（`/` 仓库内路径纯文本标注——第三轮 UI 返工后语义）
- 录入完成后取证 web 侧：时间轴 16 条倒序、当前卡"暂无进行中 phase"（完结态）+ latest 行 phase14-11、三轨过滤

#### Scenario: dogfooding 录入完成判定

- **WHEN** 16 条全部录入成功
- **THEN** ListProgressEvents(ca261521) 返回 16 条；web 时间轴倒序首行 = 最后一条 note；当前卡空态文案 + latest_task_completed = phase14-11；brief `progress.currentPhaseKey` 空串（phase14 完结）、`recentEvents` 恰 10 条

### Requirement: 固定问题取证必须逐项留档

本子任务 SHALL 完成以下取证并逐项写入 acceptance_report（每项含 answer / 取证入口 / 是否达标）：

1. **agent 经 brief 直答背景 + 进度**：一次 `GetProjectBrief(ca261521)` 调用同时返回背景侧（repository / products / modules / decisions / standards[]）与进度侧（progress.currentPhaseKey 空串 / currentPhaseLabel 空串 / latestTaskCompleted=phase14-11 事件 / recentEvents 10 条）——裁决①"接管进度说明段，与 standards[] 合成一次调用"的落地证据
2. **ListProgressEvents 全量与三轨过滤**：全量 16 条（occurred_at DESC, created_at DESC, id DESC 三键链序）；phase 轨过滤 = 14 条（1 started + 11 task + 1 completed + 1 note）；audit / fix 轨 = 各 1 条 note（附件 A 的 note 均录在 phase 轨——若附件 A 将 note 分布在三轨则按附件 A 断言）
3. **append-only 断言**：无 `Update` RPC（proto 与 handler 双侧取证）；录入 16 条后前 15 条零变形零丢失（重取 List 比对）；"当前"为派生值不落库（progress_events 表无 phase/status 列取证 + 派生仅存在于 brief/List 响应）
4. **派生正确性**：current_phase 三态中"全部完结"态命中（最新 phase_started phase14 之后存在同 key phase_completed）；latest_task_completed = phase14-11（19:33:08，最新 task_completed）；recent_events = 最近 10 条（第 11~16 条 narrative 序，附件 A 可精确推算）

#### Scenario: 固定问题达标

- **WHEN** 四项取证完成
- **THEN** 每项 answer 与 shared_baseline §3.4 派生规则矩阵逐条一致；全部标记达标

### Requirement: 十一项裁决验收门禁矩阵必须逐条全绿

本子任务 SHALL 按 shared_baseline §4 验收门禁清单逐条验证并输出矩阵（裁决项 / 验证方法 / 证据 / 结论）：

| # | 裁决 | 验证方法（冻结） |
|---|---|---|
| ① | 排序与能力定位：brief 一次调用含背景+进度 | 固定问题取证 1 的应答体（standards[] 与 progress 同体） |
| ② | repository 锚定 | 16 条事件 repository_id 全等于 ca261521（SQL 取证）；main-repo List 为空（隔离） |
| ③ | append-only + 派生不落库 | 固定问题取证 3 |
| ④ | 三轨可录可滤 | dogfooding 含 phase 轨 16 条中占 14 + 补录 audit/fix 轨各 1 条事件（若附件 A 未覆盖则此处经表单补录 2 条验证取证后删除，恢复附件 A 16 条冻结态——**见附件 A 注 2**）+ 三轨过滤取证 |
| ⑤ | 任务项颗粒 + phase 边界标记 + audit/fix 无边界事件 | 附件 A task_key 格式逐条（phase14 / phase14-NN / audit_NNN / fix_NNN）；audit/fix 轨仅 task_completed 与 note |
| ⑥ | 维护入口仅在 Repository detail | 全站 grep 无 /progress 路由与导航（phase15-07 已验，复验一次）+ 浏览器矩阵确认唯一入口 |
| ⑦ | evidence_ref 导航且正文零托管 | 附件 A 的 evidence_ref 均为 / 或 https:// 前缀；web 渲染纯文本/外链不解析内容；DB 无 spec 正文副本列 |
| ⑧ | source 仅 manual 可写 | 16 条事件 source 全为 manual（SQL）；Create 请求不含 source 字段（前端代码 + 合同） |
| ⑨ | 无 Update RPC | proto service 方法列表 + `progressv1connect` 生成物方法集 |
| ⑩ | brief 摘要与完整流分离 | brief.recentEvents=10 截断 vs List 全量 16——分层证据 |
| ⑪ | 与 PhaseEntry / Decision / plan.md 三重边界分离 | phase11 GetProjectContextResponse.phases 未改动（grep 本 phase 零触碰）；进度事件与 decision_links 零关联（SQL）；plan.md 正文零复制（DB 无正文列） |

#### Scenario: 门禁全绿

- **WHEN** 11 行矩阵完成
- **THEN** 每行结论为 PASS 且证据可复查（命令 / SQL / 应答体摘录留档 acceptance_report §8）

### Requirement: 工具链门禁与浏览器反回归矩阵必须全绿

**工具链门禁**（单值顺序完整执行，退出码留档）：`proto/` `make lint && make build && make breaking`；`backend/` `go build ./... && go vet ./... && go test ./...`（集成测试连开发库执行非 skip）；`frontend/` `npx tsc -b`（tsc --noEmit 等价零错误）。

**浏览器反回归矩阵**（真实浏览器，沿 phase14-10 矩阵抽样 + phase15 新增页点）：
1. Repository detail（ca261521）：工作台右列四卡片布局（已绑定产品 / 已映射模块 / 相关决策 / **项目进度**）+ Standard 摘要底部注释位——三轮 UI 返工最终形态回归
2. Repository detail 进度区功能点：16 条时间轴倒序 + 限高滚动（max-h-80）+ evidence_ref 纯文本/外链双形态 + 三轨过滤四向 + 当前卡完结态
3. Standard 摘要区协排回归：ca261521 恢复样本的 standards[] 展示正常（与进度区互不影响）
4. 既有页面抽查（≥6 页）：Dashboard / Standards list / Standards detail(ca261521 样本标准) / Decisions list / Products list / Onboarding——零白屏零控制台错误、导航无 progress 项
5. main-repo detail：进度区空态文案正确（清理后回归）

#### Scenario: 门禁与矩阵全绿

- **WHEN** 全部执行完成
- **THEN** 七条工具链命令退出码全 0；矩阵每页 PASS（截图留档 acceptance_report screenshots/ 或 /tmp 留档路径引用）

### Requirement: acceptance_report 与独立复核必须完成收口

- `acceptance_report.md` 沿袭 phase14-10 十三节协议撰写，冻结：固定样本与 id / 固定入口 / 恢复方式（restore 脚本 + 附件 A 重放）/ 双侧 dogfooding 取证（web 维护路径 + brief agent 读取）/ 固定问题逐项留档 / 工具链逐条结果 / 反回归矩阵 / 十一项裁决门禁矩阵 / 边界证据（非目标）/ 失败恢复与 rerun 记录 / 是否达标结论 / 独立复核记录 / rerun 指引
- 独立复核子代理（非执行者）：对附件 A 与 git log 逐条比对、门禁矩阵证据复查、矩阵截图抽查、无偷渡（零产品代码改动 / 零根级文档 / dogfooding 16 条保持）——输出 PASS/FAIL
- 收口：tasks.md / checklist.md 全勾附执行记录；变更保持未提交（仅 spec 三件套目录新增 + acceptance_report），待用户确认后手动提交

#### Scenario: 验收收口

- **WHEN** acceptance_report 冻结"达标"结论且独立复核 PASS
- **THEN** phase15-08 收口，phase15-09 具备唯一直接上游

---

## 附件 A：dogfooding 固定录入集明细（16 条，rerun 可复现单值源）

**目标仓库**：`ca261521-8daf-4248-8f12-43525326e759`（恢复脚本重建）
**occurred_at 来源**：仓库 git log phase14 各提交真实时间（+0800 本地时区，录入 datetime-local 原值；同刻相邻事件按行序录入，展示序由 created_at DESC 保证）
**录入顺序 = 行序**；全部 source=manual（表单默认，不设置 source 字段）

| # | workflow | event_kind | task_key | title | occurred_at (+0800) | detail | evidence_ref |
|---|---|---|---|---|---|---|---|
| 1 | phase | phase_started | phase14 | phase14-01 冻结 Standard Entity Foundation 的范围边界、成功标准与非目标 | 2026-08-18 08:13:43 | phase14 开启：三件套与首个子任务同提交落地 | /docs/phase/phase14_standard_entity_foundation_dev_plan.md |
| 2 | phase | task_completed | phase14-01 | phase14-01 冻结范围边界、成功标准与非目标 | 2026-08-18 08:13:43 | | /docs/phase/phase14_standard_entity_foundation_shared_baseline.md |
| 3 | phase | task_completed | phase14-02 | phase14-02 冻结 Standard 与四实体主线、画像退役的关系边界 | 2026-08-18 08:22:56 | | |
| 4 | phase | task_completed | phase14-03 | phase14-03 产出 Standard 数据模型与目录树设计 | 2026-08-18 08:39:05 | | |
| 5 | phase | task_completed | phase14-04 | phase14-04 产出后端合同、存储与读写边界设计 | 2026-08-18 08:51:05 | | |
| 6 | phase | task_completed | phase14-05 | phase14-05 产出前端信息架构与维护入口设计 | 2026-08-18 09:01:38 | | |
| 7 | phase | task_completed | phase14-06 | phase14-06 产出画像退役与数据迁移设计 | 2026-08-18 09:42:32 | | |
| 8 | phase | note | | 后端主线就绪，进入前端承接 | 2026-08-18 10:36:46 | phase14-07 完成时刻记录 | |
| 9 | phase | task_completed | phase14-07 | phase14-07 落实 Standard 后端主线 | 2026-08-18 10:36:46 | | |
| 10 | phase | task_completed | phase14-08 | phase14-08 落实 Standard 前端主线 | 2026-08-18 11:36:35 | | |
| 11 | phase | task_completed | phase14-09 | phase14-09 落实画像系统性退役与 brief 切换 | 2026-08-18 15:01:53 | | |
| 12 | phase | task_completed | phase14-10 | phase14-10 联调、dogfooding 与反回归验证 | 2026-08-18 18:21:02 | | |
| 13 | phase | note | | 联调验收通过，Standard 主线全绿 | 2026-08-18 18:21:02 | 固定 6 问 6/6 + 八项裁决门禁全绿 | |
| 14 | phase | task_completed | phase14-11 | phase14-11 完成根级同步、阶段收口与下一阶段进入条件回写 | 2026-08-18 19:33:08 | | |
| 15 | phase | phase_completed | phase14 | phase14 完成根级同步、阶段收口与下一阶段进入条件回写 | 2026-08-18 19:33:08 | phase14 收口，冻结 phase15 进入条件 | /plan.md |
| 16 | phase | note | | phase14 收口，进入 phase15 进度时间轴 | 2026-08-18 19:33:08 | phase15 排序裁决次日完成 | |

**注 1（title 语义）**：行 1 phase_started 的 title 沿用 phase14-01 提交语义（同提交落地，无独立 /plan 提交——git 取证：三件套首次提交即 84f25c9）；行 15 phase_completed 沿用 phase14-11 提交语义。
**注 2（裁决④ audit/fix 轨取证）**：附件 A 16 条均为 phase 轨（phase14 回放的事实如此——phase14 期间无 audit/fix 子任务）。裁决④"三轨可录可滤"的 audit/fix 录入取证采用**临时取证对**：经表单补录 `audit_001`（task_completed）与 `fix_001`（note）各 1 条（occurred_at 用取证时刻），验证三轨过滤 17→14/2/1 分布与 audit/fix 无边界事件后**删除这 2 条**，恢复附件 A 16 条冻结态（删除留档 acceptance_report）。
**注 3（recent_events 推算）**：16 条中最近 10 条 = 行 7~16（narrative 序）；brief.recentEvents 断言以此为准。
**注 4（rerun）**：重放前删除 ca261521 全部 progress_events，按行序重录；恢复脚本本身幂等可先跑。

## 与上游单值一致声明

- 验收环境 / 固定录入集构成 / 十一项门禁清单：逐字源 = shared_baseline §4 + dev_plan L58-61
- 16 条明细 occurred_at：逐字源 = 本仓 git log（phase14 提交真实时间，验收时独立复核须重新执行 git log 比对）
- 协议格式：沿袭 phase14-10 acceptance_report 十三节结构
- 本 spec 不新增任何产品设计决策；发现阻断性产品缺陷时停止验收、记录、按既定流程回流（fix* 或 phase15-07 返工），不在本子任务内修复
