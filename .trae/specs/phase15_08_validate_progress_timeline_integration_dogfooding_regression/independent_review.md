# phase15-08 独立复核记录（independent_review.md)

> 复核者：独立复核子代理（非执行者，审计视角）
> 复核日期：2026-08-20
> 复核对象：本 spec 的 acceptance_report.md 全部验收证据
> 复核方法：**全部证据亲自重新取证**（git log / psql SQL / ConnectRPC curl / 编译工具链 / git 工作区 / 进程只读抽查），不采信报告文字

---

## 维度 1：附件 A 与 git log / DB 逐条比对（单值源校验）— **PASS**

### 1.1 git log 比对（11 提交 ↔ 11 条 task_completed）

亲自执行 `git log --format="%h %ad %s" --date=format:"%Y-%m-%d %H:%M:%S %z" --grep="phase14" -i`，得 11 个提交，与附件 A 逐条比对全部一致：

| 附件 A 行 | task_key | 附件 A occurred_at | git 提交 | 提交时间 | 一致 |
|---|---|---|---|---|---|
| 2 | phase14-01 | 2026-08-18 08:13:43 | 84f25c9 | 08:13:43 +0800 | ✅ |
| 3 | phase14-02 | 08:22:56 | 13f7a08 | 08:22:56 | ✅ |
| 4 | phase14-03 | 08:39:05 | 6bc916f | 08:39:05 | ✅ |
| 5 | phase14-04 | 08:51:05 | 45ee128 | 08:51:05 | ✅ |
| 6 | phase14-05 | 09:01:38 | f8e076a | 09:01:38 | ✅ |
| 7 | phase14-06 | 09:42:32 | f836a89 | 09:42:32 | ✅ |
| 9 | phase14-07 | 10:36:46 | 5d8d42c | 10:36:46 | ✅ |
| 10 | phase14-08 | 11:36:35 | 8b7efdd | 11:36:35 | ✅ |
| 11 | phase14-09 | 15:01:53 | caa5eb4 | 15:01:53 | ✅ |
| 12 | phase14-10 | 18:21:02 | 66c348c | 18:21:02 | ✅ |
| 14 | phase14-11 | 19:33:08 | 31ae2cb | 19:33:08 | ✅ |

非 task_completed 行的语义核验（按附件 A 注 1）：

- 行 1 `phase_started(phase14)`：occurred_at=08:13:43 = 84f25c9（phase14-01 三件套同提交落地），title 沿用该提交语义——与注 1 一致 ✅
- 行 15 `phase_completed(phase14)`：occurred_at=19:33:08 = 31ae2cb（phase14-11），title 沿用该提交语义——与注 1 一致 ✅
- 3 条 note 的 occurred_at 分别锚定 phase14-07（5d8d42c 10:36:46）/ phase14-10（66c348c 18:21:02）/ phase14-11（31ae2cb 19:33:08）提交时刻，且行 8 detail「phase14-07 完成时刻记录」在附件 A 内自洽 ✅

结论：11 提交与 phase14-01~11 一一对应，秒级时间戳零偏差。

### 1.2 DB 数据逐行比对（16 条）

亲自执行任务指定 SQL（`SELECT task_key, title, to_char(occurred_at AT TIME ZONE 'Asia/Shanghai','YYYY-MM-DD HH24:MI') ... ORDER BY occurred_at, created_at`），返回 16 行，与附件 A 逐行比对：

- 16 行的 task_key / title / 本地时间（分钟截断）与附件 A 行 1~16 **全部逐行一致**
- 3 条 note 行（行 8/13/16）task_key 为 NULL/空 ✅
- 全部 source=manual ✅
- 分钟截断语义一致（git 秒级 08:13:43 → DB 08:13，datetime-local 分钟粒度）
- 行序 = 附件 A 行序（同刻相邻事件 created_at 递增：08:13 对、10:36 对、18:21 对、19:33 三连）——录入顺序即行序 ✅

## 维度 2：验收证据复查 — **PASS**

### 2.1 brief 断言（亲自 curl，http=200）

`POST http://localhost:8081/api/psco.project_context.v1.ProjectContextService/GetProjectBrief` `{repository_id: ca261521-...}`：

| 断言 | 取证值 | 结论 |
|---|---|---|
| repository.name | `personal-software-company-os` | ✅ |
| standards 数 | 1 | ✅ |
| progress.currentPhaseKey 空串 | proto 零值空串（protojson 默认省略零值字段，应答 progress 块仅含 `latestTaskCompleted` + `recentEvents` 两键——**语义为空串，完结态成立**，见备注 3） | ✅ |
| latestTaskCompleted.taskKey | `phase14-11`（occurredAt=`2026-08-18T11:33:00Z` = 本地 19:33，分钟截断） | ✅ |
| recentEvents 条数 | 恰 10；首条=行 16 note、末条=行 7 phase14-06（= 附件 A 注 3 推算"最近 10 条 = 行 7~16"） | ✅ |

### 2.2 List 断言（亲自 curl，http=200）

`POST .../api/psco.progress.v1.ProgressService/ListProgressEvents`：

- 全量恰 **16** 条 ✅
- 首条 title=「phase14 收口，进入 phase15 进度时间轴」（行 16 note）✅
- 末条 title=「phase14-01 冻结 Standard Entity Foundation 的范围边界、成功标准与非目标」（行 1 phase_started）✅
- eventKind 分布：PHASE_STARTED 1 / TASK_COMPLETED 11 / PHASE_COMPLETED 1 / NOTE 3 ✅
- 倒序 taskKey 序列（—, phase14, phase14-11, —, phase14-10, …, phase14-01, phase14）= 附件 A 逆行序 ✅

### 2.3 SQL 权衡（亲自执行）

| 权衡项 | 取证值 | 结论 |
|---|---|---|
| total | 16 | ✅ |
| anchored(ca261521) | 16（⇒ 其它仓库含 main-repo 名下 0 条，隔离成立） | ✅ |
| source=manual | 16/16 | ✅ |
| audit / fix 轨事件（冻结态） | 0 / 0 | ✅ |
| evidence_ref 非空条数 | 3（附件 A 行 1/2/15），全部 `/` 前缀，违规 0 | ✅ |
| 分组 | phase14(started)1 + phase14(completed)1 + phase14-01..11 各 1 + note 3 | ✅ |

### 2.4 表结构（亲自查 information_schema）

`progress_events` 恰 **11 列**：id / repository_id / workflow_type / event_kind / task_key / title / detail / evidence_ref / source / occurred_at / created_at——**无派生列（phase/status）、无 decision 关联列、无正文列** ✅

### 2.5 无 Update（亲自 grep proto）

`proto/psco/progress/v1/progress.proto` 的 rpc 行恰 3 个：`ListProgressEvents`(L172) / `CreateProgressEvent`(L182) / `DeleteProgressEvent`(L186)——**无 Update** ✅（与后端 `mountProgressConnect` 注释「无 Update——误录修正 = Delete + 重新 Create」互证）

## 维度 3：工具链抽查 — **PASS**

| 命令 | 亲自执行结果 |
|---|---|
| `frontend/ npx tsc -b` | 退出码 **0** |
| `backend/ go build ./...` | 退出码 **0** |
| `backend/ go vet ./...` | 退出码 **0** |
| `backend/ go test ./internal/progress/ -count=1` | **ok**（0.008s） |

**go test 未全量重跑的原因**：全量 `go test ./...` 含连库集成测试，其 fixture 会清空 `progress_events`（acceptance_report §10 已记录该行为——首轮即因该 fixture 清库抹掉 dogfooding 首录数据）。为保护 dogfooding 16 条冻结态，本复核仅抽查非集成纯单元包 `internal/progress`（validate/derive；集成测试位于 `internal/progress/service` 子包未触碰），并以「抽查 ok + 复核全程冻结态保持 16 条」作为独立信号。执行者侧全量结果（11 包 ok，projectcontext/connect 真跑 18.658s）以报告 §6/§10 留档为准。

## 维度 4：无偷渡核查 — **PASS**

### 4.1 工作区改动

亲自执行 `git status --porcelain`：

```
 M database/scripts/reset_product_repository_mainline.sh
 M database/seeds/seed_product_repository_mainline_baseline.sql
?? .trae/specs/phase15_08_validate_progress_timeline_integration_dogfooding_regression/
```

= 本 spec 目录（新增，spec/tasks/checklist/acceptance_report）+ 2 个验收基建文件，**零后端/前端/proto/迁移产品代码改动、零根级文档改动** ✅

### 4.2 两文件 diff 逐行核验

`git diff` 确认两文件**各恰 +1 行**（`2 files changed, 2 insertions(+)`，无删改）：均为在 `DELETE FROM repositories;` 之前补 `DELETE FROM progress_events;`——与报告 §10 的验收基建修复描述逐字一致，合法性成立（0013 FK RESTRICT 引入后 fixture 清空序滞后，非产品缺陷）✅

### 4.3 服务器进程（只读抽查）

`ps` / `ss`：psco server（`main`，监听 `*:8081`，pid 2103052）与 vite（`127.0.0.1:5173`，pid 2103305）均在运行；本复核未执行任何重启 ✅

### 4.4 dogfooding 冻结态保持

复核收尾时复查 `frozen_total = 16`——本复核全程只读取证 + 非集成测试，冻结态未被破坏 ✅

## 维度 5：报告一致性 — **PASS**

- **十三节齐全**：§1 固定样本与 id / §2 固定入口 / §3 恢复与维护方式 / §4 双侧 dogfooding / §5 固定问题 / §6 工具链 / §7 反回归矩阵 / §8 十一项裁决门禁 / §9 边界证据 / §10 失败恢复 rerun / §11 是否达标 / §12 独立复核 / §13 rerun 指引 ✅
- **§11 结论**：「达标。」✅
- **§8 矩阵**：①~⑪ 共 11 行，全部 PASS ✅
- **§10 记录完整**：首轮 `go test` FAIL（projectcontext 集成测试 fixture 报 FK `progress_events_repository_id_fkey` 违规）→ 根因（0013 引入 FK RESTRICT 后 fixture 清空序未同步）→ 2 行基建修复 → rerun 全绿（11 ok）+ 重放 16 条恢复冻结态——链条完整 ✅
- 报告关键断言（§1 冻结态数字 / §5-2 三轨分布 / §8-②⑦⑧ SQL 口径）与本复核独立取证**全部吻合**

---

## 阻断项清单

**0 项。**

## 非阻断备注（供收口知悉，不影响结论）

1. **tasks.md / checklist.md checkbox 当前均未勾选**：按 spec SubTask 7.3，「全部勾选附执行记录」排在独立复核（7.2）之后执行，属流程内收口尾巴。本复核输出后应完成勾选，方构成完整收口。
2. **报告 §8-⑤ 措辞**：「phase14/phase14-01..11 各 1 + note 3」中 phase14 实为 2 条（phase_started 1 + phase_completed 1，同 task_key 不同 event_kind）。本复核 SQL 分组证实实际数据正确（14 分组行合计 16），属简写措辞不精确，非数据错误。
3. **currentPhaseKey 的 JSON 编码语义**：protojson 默认省略零值字段，故 brief 应答体中该字段不可见（progress 块仅两键），proto 层面语义=空串，与报告 §5-1「currentPhaseKey=''」口径在此编码行为下成立，完结态断言不受影响。
4. **main-repo 当前 id 为 71f5f1c4**（原 1a863cb1 已被集成测试 fixture 清库移除）——报告 §1 已如实披露并给出理由（非验收数据）；「main-repo 0 条事件」经 total=16 ∧ anchored=16 推出成立。

## 最终结论

# **PASS**

附件 A 与 git log / DB 双侧逐条一致（单值源校验通过）；brief / List / SQL / 表结构 / 无 Update 断言全部亲自复验成立；工具链抽查全零退出码；工作区改动恰为 spec 目录 + 2 处已披露的验收基建最小修复（各 +1 行），零产品代码偷渡、零根级文档回写；dogfooding 16 条冻结态在库保持；acceptance_report 十三节齐全、结论「达标」有据。phase15-08 验收证据真实、完整，具备进入 phase15-09 根级收口的条件。
