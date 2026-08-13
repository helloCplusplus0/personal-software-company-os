# Fix 001-003 Closure - MVP0.3 Dry-Run 阻断项聚焦 Rerun 收口记录

## 1. 文档定位

本文不是新的 `phase` 文档，也不是新的评审汇总文档。

本文只承接一件事：

- 为 `fix_001 ~ fix_003` 在真实 `dry-run` 阻断项上的修复、复验与收口提供一份极简但正式的结论记录；
- 作为 `mvp0.3` `Real-Project Dry-Run` 从“发现阻断”走到“阻断消失”的验收钉子；
- 为是否进入 `PSCO-mvp04-summarize-feedback.md` 中的“候选阶段二：Asset-Action Closure 主线”提供直接判断依据。

## 2. 本次收口覆盖范围

本次聚焦 rerun 只覆盖以下三项阻断项：

1. `fix_001`：`Onboarding` 冷启动欢迎页首次点击无响应；
2. `fix_002`：`Decision` 待处理信号语义错位；
3. `fix_003`：`Decision Detail` 缺少正式状态推进入口。

本文不重新评估 `mvp0.3` 全量体验，也不扩写 `mvp0.4` 的下一阶段实现设计。

## 3. 验收输入

本次收口以以下材料与验证结果为直接输入：

- [PSCO-mvp03-real-project-dry-run-user-manual-GPT54.md](file:///home/dell/Projects/personal-software-company-os/docs/review/PSCO-mvp03-real-project-dry-run-user-manual-GPT54.md)
- [PSCO-mvp03-real-project-dry-run-user-manual-GPT54 feedback.md](file:///home/dell/Projects/personal-software-company-os/docs/review/PSCO-mvp03-real-project-dry-run-user-manual-GPT54%20feedback.md)
- [fix_001_onboarding_cold_start_state_issue.md](file:///home/dell/Projects/personal-software-company-os/docs/fix/fix_001_onboarding_cold_start_state_issue.md)
- [fix_001_onboarding_cold_start_state_analysis.md](file:///home/dell/Projects/personal-software-company-os/docs/fix/fix_001_onboarding_cold_start_state_analysis.md)
- [fix_002_decision_pending_signal_semantics_issue.md](file:///home/dell/Projects/personal-software-company-os/docs/fix/fix_002_decision_pending_signal_semantics_issue.md)
- [fix_002_decision_pending_signal_semantics_analysis.md](file:///home/dell/Projects/personal-software-company-os/docs/fix/fix_002_decision_pending_signal_semantics_analysis.md)
- [fix_003_decision_detail_status_advance_issue.md](file:///home/dell/Projects/personal-software-company-os/docs/fix/fix_003_decision_detail_status_advance_issue.md)
- [fix_003_decision_detail_status_advance_analysis.md](file:///home/dell/Projects/personal-software-company-os/docs/fix/fix_003_decision_detail_status_advance_analysis.md)
- `fix_001` 的源码独立复核结论与前端构建验证；
- `fix_002 + fix_003` 的源码独立复核、浏览器端实际页面验证与后端只读结果确认；
- 真实使用者对 rerun 结果的最终确认：上述三项阻断均已消失。

## 4. 聚焦 Rerun 结论

### 4.1 `fix_001` 结论

- `Onboarding` 欢迎页“开始首轮录入”首次点击无响应问题已消失；
- 冷启动用户可以从 `/dashboard -> /onboarding -> 开始首轮录入` 正常进入首轮流程；
- 该阻断不再构成 `mvp0.3 dry-run` 的继续使用障碍。

### 4.2 `fix_002` 结论

- `Dashboard / Daily Review / Current Focus` 对 `Decision` pending signal 的误报问题已消失；
- `pending decision` 继续锚定 canonical `Decision.status`，未长出第二事实源；
- “我明明已经处理过，但系统仍在催促”的核心语义冲突已消失。

### 4.3 `fix_003` 结论

- `Decision Detail` 已具备正式状态推进 CTA；
- 用户可以通过正式写链推进 `Decision.status`，不再停留在“看得到待处理、却没有处理出口”的断裂状态；
- 状态推进后，`Decision Detail / Dashboard / Daily Review` 的 reread 与回流语义已经闭环。

## 5. 本次收口的正式判断

经过 `fix_001 ~ fix_003` 修复、独立复核、浏览器级验证与真实使用者 rerun 确认，本次 `mvp0.3` `Real-Project Dry-Run` 第一轮暴露的三项 P0 阻断均已消失。

正式判断冻结为：

1. `fix_001 ~ fix_003` 已完成本轮 fix workflow 的核心目标；
2. `mvp0.3` `Real-Project Dry-Run` 已从“发现阻断并暂停推进”进入“阻断消失，可正式收口”状态；
3. `PSCO-mvp04-summarize-feedback.md` 中“候选阶段一：dry-run 阻断项 fix 与收口”的前置目标已满足；
4. 项目可以正式结束本轮 `mvp0.3 dry-run` 阻断项处置，并进入下一步“候选阶段二：Asset-Action Closure 主线”的正式准备与建立入口。

## 6. 明确不在本文继续处理的内容

1. 不在本文中直接展开“候选阶段二”的 `/plan` 或 `/spec`；
2. 不在本文中追加新的 `dry-run` 长清单或泛化验收范围；
3. 不在本文中重写 `mvp0.4` 的阶段命名、接口名或实现方案；
4. 不把本轮收口扩写为新一轮专家评审。

## 7. 收口后下一步

收口完成后，下一步正式工作应转向：

- 以 [PSCO-mvp04-summarize-feedback.md](file:///home/dell/Projects/personal-software-company-os/PSCO-mvp04-summarize-feedback.md#L380-L387) 为上游基线；
- 围绕“候选阶段二：Asset-Action Closure 主线”建立新的正式入口；
- 聚焦推进：
  1. 让 `Onboarding` 变成真正的首轮建链引导；
  2. 让 `Decision` 形成最小但真实的生命周期；
  3. 让 Dashboard / Review / Detail pages 共同承接“下一步动作”；
  4. 让 `Current Focus` 与 pending signals 回到真实经营语义。
