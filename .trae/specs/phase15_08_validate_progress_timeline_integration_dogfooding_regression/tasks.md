# Tasks

- [x] Task 1: 验收环境准备（恢复固定样本 + 清理测试残留）
  - [x] SubTask 1.1: 执行 `./database/scripts/restore_phase11_phase12_dogfooding_sample.sh` 恢复固定样本仓库；验证 `ca261521-8daf-4248-8f12-43525326e759` 就位且 `GetProjectBrief(ca261521)` 应答 repository + standards[] 非空（背景侧就绪）
  - [x] SubTask 1.2: 经 `DeleteProgressEvent` API 删除 `main-repo`（1a863cb1）上 phase15-07 用户手录的 5 条 phase01 测试事件；确认 main-repo List 返回空、全库 progress_events 归零；5 条删除记录（id/title）留档（写入 acceptance_report §1 素材）

- [x] Task 2: dogfooding 固定录入集录入（附件 A 16 条，web 表单路径）
  - [x] SubTask 2.1: 编写真实浏览器驱动脚本（沿 phase15-07 /tmp/psco-browser-test 模式：playwright-core + 本地 chromium，HOME 重定向 /tmp）：导航至 `/repositories/ca261521-...`，经 ProgressEventForm 按附件 A 行序逐条录入 16 条（occurred_at 补录历史填附件 A 本地时刻；note 行不填 task_key；evidence_ref 按附件 A）
  - [x] SubTask 2.2: 录入过程断言：每条提交成功后表单重置（title 清空 + event_kind 保持）；16 条完成后 web 时间轴 16 条倒序（首行 = 行 16 note）+ 当前卡"暂无进行中 phase"完结态 + latest 行 phase14-11；三轨过滤 phase=16/audit=0/fix=0（临时取证对之前的附件 A 冻结态）
  - [x] SubTask 2.3: 录入留证：每条提交后截图或最终整卡截图（/tmp/psco-browser-test/ 留档，acceptance_report 引用）；SQL 复核 16 条 repository_id 全等于 ca261521、source 全为 manual、task_key 格式逐条符合附件 A

- [x] Task 3: 固定问题取证（四项逐项留档）
  - [x] SubTask 3.1: agent 经 brief 直答背景+进度——curl `GetProjectBrief(ca261521)`：一次应答含 repository/products/modules/decisions/standards[]（背景）+ progress（currentPhaseKey 空串 / latestTaskCompleted=phase14-11 / recentEvents 恰 10 条 = 附件 A 行 7~16）；应答体摘录留档
  - [x] SubTask 3.2: ListProgressEvents 全量与三轨过滤——全量 16 条三键链倒序断言（与附件 A 逆行序一致）；phase 过滤 16 条；audit/fix 过滤 0 条（临时取证对之前）
  - [x] SubTask 3.3: append-only 断言——proto service 方法集无 Update + connect 生成物方法集取证；重取 List 与录入后首取比对前 15 条零变形零丢失；progress_events 表结构取证（无 phase/status/派生列——派生仅存在响应侧）
  - [x] SubTask 3.4: 派生正确性——current_phase"全部完结"态（最新 phase_started phase14 后存在同 key phase_completed）取 brief 断言；latest_task_completed=phase14-11；recent_events 截断 10 条与附件 A 注 3 推算一致

- [x] Task 4: 十一项裁决验收门禁矩阵（shared_baseline §4 逐条）
  - [x] SubTask 4.1: 按 spec §"十一项裁决验收门禁矩阵必须逐条全绿"表格逐行取证（①brief 同体 / ②repository 锚定 SQL / ③append-only（复用 Task 3.3）/ ④三轨可录可滤——执行 spec 附件 A 注 2 临时取证对：表单补录 audit_001(task_completed)+fix_001(note) 各 1 条 → 三轨过滤 17/14/2/1 分布断言 + audit/fix 无边界事件验证 → 删除 2 条恢复 16 条冻结态（删除留档）/ ⑤task_key 格式 SQL 逐条 / ⑥全站 grep 无 /progress 路由 + 浏览器确认唯一入口 / ⑦evidence_ref 前缀 SQL + web 渲染双形态截图 + 零正文托管 / ⑧source 全 manual SQL + Create 无 source 字段 / ⑨无 Update RPC 双侧取证 / ⑩recentEvents 10 vs List 16 分层 / ⑪PhaseEntry 零触碰 grep + decision 零关联 SQL + plan.md 零复制）
  - [x] SubTask 4.2: 矩阵汇总（11 行 × 裁决/方法/证据/结论），全部 PASS 方可进入 Task 5

- [x] Task 5: 工具链门禁（单值顺序完整执行，退出码留档）
  - [x] SubTask 5.1: `proto/` 目录 `make lint && make build && make breaking` 三命令零退出码
  - [x] SubTask 5.2: `backend/` 目录 `go build ./... && go vet ./... && go test ./...` 零退出码（集成测试连开发库执行非 skip——dogfooding 数据存在下不破坏既有测试断言：测试均用独立 fixture）
  - [x] SubTask 5.3: `frontend/` 目录 `npx tsc -b` 零错误

- [x] Task 6: 浏览器反回归矩阵（真实浏览器，5 组）
  - [x] SubTask 6.1: Repository detail（ca261521）——工作台右列四卡片布局（已绑定产品/已映射模块/相关决策/项目进度）+ Standard 摘要底部注释位：三轮 UI 返工最终形态回归（截图留档）
  - [x] SubTask 6.2: 进度区功能点回归——16 条倒序 + max-h-80 限高滚动 + evidence_ref 纯文本（/ 前缀 3 条）/ 外链双形态 + 三轨过滤四向 + 当前卡完结态组合断言
  - [x] SubTask 6.3: Standard 摘要区协排回归——ca261521 样本 standards[] 展示正常（树形/链接），与进度区互不影响
  - [x] SubTask 6.4: 既有页面抽查 ≥6 页——Dashboard / Standards list / Standards detail / Decisions list / Products list / Onboarding：零白屏 + 控制台零错误（console 消息采集）+ 导航无 progress 项
  - [x] SubTask 6.5: main-repo detail——进度区空态文案正确（"暂无推进事件，从上方表单录入第一条。"）+ 当前卡空态组合（清理后回归）

- [x] Task 7: acceptance_report 撰写、独立复核与收口
  - [x] SubTask 7.1: 撰写 `acceptance_report.md`（沿袭 phase14-10 十三节协议：固定样本与 id / 固定 Web 页面与 agent 入口 / 样本恢复方式与进度维护方式 / 双侧 dogfooding 取证 / 固定问题逐项留档 / 工具链逐步结果 / 浏览器反回归矩阵 / 十一项裁决门禁矩阵 / 边界证据清单 / 失败点恢复路径与 rerun 结果 / 是否达标 / 独立复核记录 / rerun 指引）
  - [x] SubTask 7.2: 独立复核子代理（非执行者）：附件 A 与 git log 逐条重新比对（occurred_at/title 单值）/ 门禁矩阵证据复查 / 截图抽查 / 无偷渡（零产品代码改动〔git status 仅 spec 目录〕/ 零根级文档 / dogfooding 16 条保持 / main-repo 0 条）——输出 PASS/FAIL 与阻断项
  - [x] SubTask 7.3: 修复独立复核发现的阻断性问题（如有——产品缺陷回流既定流程不在本任务修；报告缺陷则验收暂停）并复验；tasks.md / checklist.md 全部勾选附执行记录；变更保持未提交（spec 三件套 + acceptance_report + independent_review），待用户最终确认后手动提交

# Task Dependencies

- Task 1 为全部后续前提（样本仓库与干净数据）
- Task 2 depends on Task 1；Task 3 depends on Task 2（取证对象为录入结果）
- Task 4 depends on Task 2 + Task 3（复用 ③ 证据；④ 的临时取证对在 16 条冻结态上执行并恢复）
- Task 5 与 Task 2~4 可并行（门禁不依赖 dogfooding 数据；但 go test 须在临时取证对删除后跑，避免数据干扰——实践上置于 Task 4 后串行）
- Task 6 depends on Task 2（进度区回归需 16 条数据）+ Task 4（临时取证对已删，16 条冻结态）
- Task 7 depends on Task 1 ~ Task 6
- 后续：phase15-09 depends on 本 spec acceptance_report（唯一直接上游）

# 执行记录（2026-08-20）

- Task 1：restore 脚本恢复 ca261521 + 关联 Standard；API 删除 main-repo 5 条手录残留（id 留档 acceptance_report §1）；全库归零。
- Task 2：真实浏览器 web 表单路径录入 16 条（两轮：首录 p08-drive + fixture 清库后重放 p08-replay）；每条重置语义实测；SQL 复核 16/16 锚定、16/16 manual、分组与附件 A 逐行一致。
- Task 3：brief 同体五背景块+进度（recent=10 首末条与注 3 一致）；List 16 条三键链序=附件 A 逆行序；append-only 双侧无 Update + 两次 List 逐字段零变形（首版比对因 note 行 proto3 省略 taskKey 出现 KeyError 假阳性，.get 修正版实证 len=1713 一致）；派生四断言全 PASS（currentPhaseKey 空串/latest=phase14-11@11:33Z/recent 截断）。
- Task 4：11 行门禁全 PASS。④ 临时取证对（表单路径 audit_001+fix_001）分布 18/16/1/1 → 删除恢复 16（id c878a650/e286c634 留档）；②⑤⑦⑧ SQL 权衡；①⑩ brief 应答；③⑨ proto+生成物+表结构；⑥ grep+导航；⑪ 零触碰+零关联+零正文。注：spec 表格"17→14/2/1"与"phase 轨 14 条"为撰写期算术笔误，执行按事实断言 18/16/1/1 与 16（附件 A 16 条全 phase 轨），独立复核认可。
- Task 5：七命令全零退出码。首轮 go test FAIL（projectcontext fixture 重置触发 progress_events FK RESTRICT——phase15-06 的 0013 引入 FK 后 phase04 时代 fixture 清空链未同步）：验收基建最小修复 2 行（reset_product_repository_mainline.sh + baseline seed 各在 DELETE repositories 前补 DELETE progress_events），重跑 11 包 ok（projectcontext 真跑 18.658s）；副作用清库 → 恢复样本 + p08-replay 重放并全量复验。
- Task 6：5 组全 PASS——四卡片布局/16 条倒序/max-h-80/纯文本 3 条/三轨过滤/完结态（6.1-6.2）；Standard 协排（6.3）；6 页零白屏零 console error + 导航无 progress 项（6.4）；空态组合文案在 0 事件窗口实测（6.5）。截图留档 /tmp/psco-browser-test/。
- Task 7：acceptance_report.md 十三节冻结"达标"；独立复核子代理五维度全 PASS 0 阻断（附件 A 与 git log 秒级零偏差 / 证据独立复取 / 工具链抽查 / 无偷渡——工作区恰 = spec 目录 + 2 个各 +1 行验收基建文件 / 报告一致性），4 条非阻断备注留档 independent_review.md；tasks/checklist 全勾；变更未提交待用户确认。
