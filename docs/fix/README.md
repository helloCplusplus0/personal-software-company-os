# Fix Index

## 1. 职责

本目录承载 bug 修复与局部问题处理。

## 2. 规则

- 每个 `fix` 先记录 issue
- 再产出 analysis
- 再进入 `/spec`、实现与验收

## 3. 当前状态

- 当前已创建：
  - `fix_001_onboarding_cold_start_state_issue.md`
  - `fix_001_onboarding_cold_start_state_analysis.md`
  - `fix_002_decision_pending_signal_semantics_issue.md`
  - `fix_002_decision_pending_signal_semantics_analysis.md`
  - `fix_003_decision_detail_status_advance_issue.md`
  - `fix_003_decision_detail_status_advance_analysis.md`
  - `fix_001_003_mvp03_dry_run_rerun_closure.md`
- 当前结论：
  - `fix_001` 已完成 analysis、`/spec`、实现与独立复核
  - `fix_002` 已完成 analysis，并与 `fix_003` 完成联动 `/spec`、实现、浏览器级验收与独立复核
  - `fix_003` 已完成 analysis，并与 `fix_002` 完成联动 `/spec`、实现、浏览器级验收与独立复核
  - `fix_001 ~ fix_003` 已完成本轮 `mvp0.3 dry-run` 阻断项聚焦 rerun 与收口
- 下一步：
  - 以 `PSCO-mvp04-summarize-feedback.md` 为上游基线，进入“候选阶段二：Asset-Action Closure 主线”的正式入口建立
