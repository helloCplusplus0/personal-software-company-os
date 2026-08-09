# Phase Index

## 1. 职责

本目录承载项目推进主线。

任何新的功能推进、规格推进、结构性建设，都必须从这里开始，而不是直接创建孤立文档。

## 2. 当前状态

- 根级当前状态：`phase05_dashboard_feedback_foundation` 已完成 `/plan -> /spec -> 实现 -> 验收 -> 收口`
- `phase01_*` 三件套已完成规划侧收口
- `phase02_*` 三件套已完成并已收口
- `phase03_*` 三件套已完成并已收口
- `phase04_*` 三件套已完成并已收口
- `phase05_*` 三件套保留为 `phase05` 的规划与冻结记录，不再覆盖根级当前状态
- 下一阶段正式 phase 入口：待建立后切换（不得预设新的 phase 名称）
- 下一阶段直接上游已冻结为：`.trae/specs/phase05_10_dashboard_feedback_formal_spec/dashboard_feedback_spec_v0.1.md`、`.trae/specs/phase05_11_dashboard_feedback_proto_mainline/`、`.trae/specs/phase05_14_dashboard_feedback_integration_validation_acceptance/acceptance_report.md`
- 当前项目技术路线：`Durable System Track`

## 2.1 最近完成阶段的规划记录

- [phase05_dashboard_feedback_foundation_architecture_plan.md](file:///home/dell/Projects/personal-software-company-os/docs/phase/phase05_dashboard_feedback_foundation_architecture_plan.md)
- [phase05_dashboard_feedback_foundation_dev_plan.md](file:///home/dell/Projects/personal-software-company-os/docs/phase/phase05_dashboard_feedback_foundation_dev_plan.md)
- [phase05_dashboard_feedback_foundation_shared_baseline.md](file:///home/dell/Projects/personal-software-company-os/docs/phase/phase05_dashboard_feedback_foundation_shared_baseline.md)

## 3. 规则

- 每个 `phase` 至少包含：
  - `phase*_architecture_plan.md`
  - `phase*_dev_plan.md`
  - `phase*_shared_baseline.md`
- `/plan` 通过后，才能继续 `/spec` 与实现
- 已完成收口的 `phase`，后续实现应从正式 `/spec` 进入，而不是把三件套继续当作并列执行入口
