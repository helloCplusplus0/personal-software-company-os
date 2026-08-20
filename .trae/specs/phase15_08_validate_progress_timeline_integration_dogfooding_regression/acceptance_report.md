# phase15-08 进度时间轴联调、dogfooding 与反回归验收报告

> 验收对象：phase15-06（后端主线）+ phase15-07（前端主线，含三轮 UI 返工）交付的"项目推进时间轴最小主线"
> 协议格式：沿袭 phase14-10 十三节结构
> 验收日期：2026-08-20

## 1. 固定样本与 id

- **dogfooding 样本仓库**：`ca261521-8daf-4248-8f12-43525326e759`（`personal-software-company-os`，phase11/12 冻结固定样本），经 `./database/scripts/restore_phase11_phase12_dogfooding_sample.sh` 恢复
- **dogfooding 固定录入集**：本 spec §附件 A 的 16 条（1 phase_started + 11 task_completed + 1 phase_completed + 3 note），occurred_at 逐条取仓库 git log phase14 十一提交真实时间（+0800，datetime-local 分钟截断录入，DP-3 转换 UTC——DB 侧 phase14-11 occurred_at = `2026-08-18T11:33:00Z`）
- **测试残留清理记录**（phase15-07 承诺兑现）：main-repo（`1a863cb1`）上用户手录 5 条已删——`56f454ff`(phase01-06) / `c55a9687`(phase01-02) / `35a8a0c3`(fix_001:NNBVN) / `961639b6`(phase01-01:XXCC) / `0922bd7a`(phase01_01:XXXX)，经 DeleteProgressEvent API；main-repo List 归零
- **最终冻结态**：progress_events 全库恰 16 条（SQL 复核：total=16 / anchored(ca261521)=16 / source=manual=16；main-repo 随后端集成测试 fixture 清库移除，非验收数据）

## 2. 固定 Web 页面与 agent 入口

- Web 维护与展示唯一入口：Repository detail `/repositories/ca261521-...` 工作台右列第四卡片"项目进度"（裁决⑥：无独立路由/导航/Dashboard 入口，全站 grep 与导航 hrefs 复验）
- agent 读取入口：`GetProjectBrief`（progress = 9 摘要块）+ `psco.progress.v1.ProgressService.ListProgressEvents`（完整流）

## 3. 样本恢复方式与进度维护方式

- 恢复：幂等 restore 脚本（不清空开发数据）+ 附件 A 16 条 web 表单重放（rerun 指引见 §13）
- 维护：仅 web 手动录入（source 归一 manual）；误录唯一修正 = window.confirm 确认后整条删除（append-only，裁决⑨）

## 4. 双侧 dogfooding 取证

### 4.1 人类维护路径（Web 表单）

真实浏览器（playwright-core + 本地 chromium，沿 phase15-07 驱动模式）经 ProgressEventForm 按附件 A 行序录入 16 条（occurred_at 补录历史逐条填入；每条提交后表单重置语义实测：title 清空 + event_kind 保持）。两轮完整录入（首录 + fixture 清库后重放），全部成功；过程留证脚本与截图：`/tmp/psco-browser-test/`（p08-drive.mjs / p08-replay.mjs / p08-dogfooding-16.png / p08-final-frozen.png）。

### 4.2 agent 读取路径（brief 实时取证）

`GetProjectBrief(ca261521)` 一次应答同体返回：背景侧 repository(name=personal-software-company-os) + standards[1] + products[1] + modules[1] + decisions[1]，进度侧 progress（currentPhaseKey="" / latestTaskCompleted.taskKey="phase14-11" / recentEvents 恰 10 条）——裁决①"背景+进度一次调用"落地实证。

## 5. 固定问题逐项留档

| # | 问题 | answer（取证值） | 入口 | 达标 |
|---|---|---|---|---|
| 1 | agent 经 brief 直答背景+进度 | 五背景块 + progress 同体；recent=10 | GetProjectBrief 应答体 | ✅ |
| 2 | ListProgressEvents 全量与三轨过滤 | 全量 16 且三键链序 = 附件 A 逆行序；phase=16 / audit=0 / fix=0（冻结态）；临时对下 18/16/1/1 | List API + SQL | ✅ |
| 3 | append-only 断言 | proto 与 connect 生成物均仅 List/Create/Delete（无 Update）；两次 List 逐字段零变形（id/title/taskKey/occurredAt 比对一致，len=1713）；表 11 列无派生/状态列 | proto L172-186 + 生成物 Procedure 常量 + API 双取 + information_schema | ✅ |
| 4 | 派生正确性 | current_phase 命中"全部完结"态（brief currentPhaseKey 空串）；latest=phase14-11@11:33:00Z；recent 首条=行16 note、末条=行7（附件 A 注 3 一致） | brief 应答 | ✅ |

## 6. 工具链逐步结果（单值顺序完整执行）

| 命令 | 结果 |
|---|---|
| proto/ `make lint` / `make build` / `make breaking` | 0 / 0 / 0 |
| backend/ `go build ./...` / `go vet ./...` | 0 / 0 |
| backend/ `go test ./...` | **0**（11 包 ok，projectcontext/connect 真跑 18.658s；首轮 FAIL 见 §10，修复后全绿） |
| frontend/ `npx tsc -b` | 0（零错误） |

## 7. 浏览器反回归矩阵

| 组 | 页面/点位 | 结果 |
|---|---|---|
| 6.1 | Repository detail 工作台右列四卡片（已绑定产品→已映射模块→相关决策→项目进度）+ Standard 摘要底部注释位 | PASS（截图 p08-layout.png / p08-final-frozen.png） |
| 6.2 | 进度区功能点：16 条倒序=附件 A 逆行序 / max-h-80 限高滚动 / `/` 前缀 evidence_ref 纯文本 3 条 / 三轨过滤 16-0-0 + 空态文案"该轨暂无事件。" / 当前卡完结态 | 全 PASS |
| 6.3 | Standard 摘要协排：样本 standards[] 展示正常，与进度区互不影响 | PASS |
| 6.4 | 既有页面 6 页（Dashboard / Standards list / Standards detail / Decisions list / Products list / Onboarding）零白屏零 console error + 导航 hrefs（/dashboard /modules /decisions /products /repositories /standards）无 progress 项 | 7/7 PASS |
| 6.5 | 空态回归：当前卡"暂无进行中 phase + 暂无任务完成记录"组合 + 时间轴"暂无推进事件，从上方表单录入第一条。"（0 事件窗口实测，截图 p08-empty-state.png） | PASS |

## 8. 十一项裁决验收门禁矩阵（shared_baseline §4）

| # | 裁决 | 证据（可复查） | 结论 |
|---|---|---|---|
| ① | brief 一次调用含背景+进度 | §5-1 应答体（standards[] 与 progress 同体） | PASS |
| ② | repository 锚定 | SQL：anchored=16/16 全等 ca261521 | PASS |
| ③ | append-only + 派生不落库 | §5-3（无 Update 双侧 + 零变形 + 11 列无派生列） | PASS |
| ④ | 三轨可录可滤 | 临时取证对（audit_001 task_completed + fix_001 note，表单路径录入）：分布 18/16/1/1 → 删除恢复 16（删除 id：c878a650 / e286c634 留档） | PASS |
| ⑤ | 任务项颗粒 + audit/fix 无边界事件 | SQL 分组：phase14/phase14-01..11 各 1 + note 3；audit/fix 轨边界事件计数=0 | PASS |
| ⑥ | 维护入口仅 Repository detail | 全站 grep 无 /progress 路由（routes 内 "progress" 仅 first_run_state 枚举）；§7-6.4 导航 hrefs 复验 | PASS |
| ⑦ | evidence_ref 导航且零托管 | SQL：3 条非空全 `/` 前缀，违规计数=0；web 纯文本渲染 3 条；表无正文列 | PASS |
| ⑧ | source 仅 manual | SQL：16/16 manual；前端 owner 代码注释与请求体不含 source（grep L8/L60） | PASS |
| ⑨ | 无 Update RPC | proto L172/182/186 三方法 + 生成物三 Procedure 常量 | PASS |
| ⑩ | brief 摘要与完整流分离 | recent=10 截断 vs List=16（同库同派生两通道分层） | PASS |
| ⑪ | 三重边界分离 | phase11 PhaseEntry 零触碰（本 phase 零后端代码改动，git 工作区仅 spec 目录+fixture 脚本）；progress 模块零 decision 引用（grep）；表无 plan.md 正文列 | PASS |

## 9. 边界证据清单（非目标）

- git 自动采集 / agent 写回：source 枚举含 git/agent 但 16 条全 manual，无任何 git/agent 写入口（grep 零命中）
- MCP / CLI / 模板仓库 / 自动同步 / standard_bindings 扩展 / Decision 互链 / UpdateProgressEvent / phase11 PhaseEntry 改动 / plan.md 接管：全部零实现零触碰
- 产品代码改动：**零**（git 工作区改动仅本 spec 目录 + 两处验收基建脚本，见 §10）

## 10. 失败点、恢复路径与最终 rerun 结果

| 失败点 | 根因 | 恢复 | rerun |
|---|---|---|---|
| `go test ./...` 首轮 FAIL：projectcontext 集成测试 3 场景 fixture 重置报 FK `progress_events_repository_id_fkey` 违规 | phase15-06 的 0013 引入 FK RESTRICT 后，phase04 时代 fixture 清空链（reset_product_repository_mainline.sh 与其 baseline seed）未同步先清 progress_events——验收基建滞后，非产品缺陷 | 验收基建最小修复（2 行）：两文件清空序在 `DELETE FROM repositories` 前补 `DELETE FROM progress_events` | `go test ./...` 全绿（11 ok）；副作用为 fixture 清库抹掉首录 dogfooding 与 main-repo → 按 §3 恢复脚本 + p08-replay 重放 16 条并全量复验（§5/§7 结论均基于重放后冻结态） |
| 断言脚本 3 处定位误差（6.1 标题取值 / title 行选择器 / 过滤按钮嵌套定位）与 1 处假阳性（3.3 首版 note 行 taskKey KeyError 致空比对） | 驱动脚本问题，非产品缺陷 | 脚本修正（evaluateAll 首子节点 / font-medium 精确匹配 / .get 兜底） | 全部 PASS |

**修复涉及的验收基建文件**（保持未提交，随本 spec 一并待确认）：
- `database/scripts/reset_product_repository_mainline.sh`（清空序 +1 行）
- `database/seeds/seed_product_repository_mainline_baseline.sql`（同 +1 行）

## 11. 是否达标

**达标。** 固定问题 4/4 全达标；十一项裁决门禁 11/11 全绿；工具链七命令全零退出码；浏览器反回归矩阵 5 组全 PASS；dogfooding 16 条冻结态在库（rerun 可复现）。

## 12. 独立复核记录

见同目录 `independent_review.md`（独立复核子代理输出：附件 A 与 git log 逐条比对 / 门禁证据复查 / 无偷渡核查 / 最终结论）。

## 13. Rerun 指引（面向不同执行者）

1. 前置：后端(:8081) + vite(:5173) 运行中；共享 PostgreSQL 容器在运行
2. 恢复样本：`./database/scripts/restore_phase11_phase12_dogfooding_sample.sh`
3. 清旧进度数据（如重放）：`DELETE FROM progress_events;`（容器 psql）或逐条 API 删除
4. 重放 16 条：`node /tmp/psco-browser-test/p08-replay.mjs`（含空态取证与核心断言；如 /tmp 已清理，按本 spec 附件 A 行序经 Web 表单录入）
5. 复跑门禁与矩阵：§6 七命令 + `/tmp/psco-browser-test/p08-sweep.mjs`（或等价浏览器抽查）
6. 断言口径：全部以附件 A 与 shared_baseline §3.4/§4 为单值源
