# Phase Index

## 1. 职责

本目录承载项目推进主线。

任何新的功能推进、规格推进、结构性建设，都必须从这里开始，而不是直接创建孤立文档。

## 2. 当前状态

- 根级当前状态：`phase09_template_reuse_derived_intelligence_foundation` 已进入 `/plan`；`phase08_operating_review_loop_foundation` 继续保留为最近完成的正式业务 phase
- `phase01_*` 三件套已完成规划侧收口
- `phase02_*` 三件套已完成并已收口
- `phase03_*` 三件套已完成并已收口
- `phase04_*` 三件套已完成并已收口
- `phase05_*` 三件套保留为最近完成阶段的规划与冻结记录
- `phase06_*` 三件套已完成并保留为最近完成阶段的规划与冻结记录
- `phase07_*` 三件套已完成并继续保留为最近完成前置基础阶段的规划与冻结记录
- `phase08_*` 三件套已完成并保留为最近完成正式业务 phase 的规划与冻结记录
- `phase09_*` 三件套已建立并作为当前正式 `/plan` 入口
- `phase08-11` acceptance_report 已冻结为 `phase08` 的正式验收与收口入口
- `phase09` 复核通过后才允许继续进入 `/spec`
- 后续正式 phase 入口只允许在直接承接 `phase09` 验收结论、`phase08` 根级收口结论与当前根级真相源后建立（不得预设新的 phase 名称）
- 当前阶段直接上游已冻结为：`PSCO-mvp03-summarize-feedback.md`、`.trae/specs/phase07_07_formal_transport_mainline_cutover_spec/transport_mainline_cutover_spec_v0.1.md`、`.trae/specs/phase07_11_validate_phase01_phase06_regression_retirement_acceptance/acceptance_report.md`、`.trae/specs/phase08_11_validate_review_loop_integration_browser_regression_acceptance/acceptance_report.md`、`.trae/specs/phase05_10_dashboard_feedback_formal_spec/dashboard_feedback_spec_v0.1.md`、`.trae/specs/phase03_10_decision_center_formal_spec/decision_center_spec_v0.1.md`
- 当前项目技术路线：`Durable System Track`

## 2.1 当前规划中的正式 phase

- [phase09_template_reuse_derived_intelligence_foundation_architecture_plan.md](file:///home/dell/Projects/personal-software-company-os/docs/phase/phase09_template_reuse_derived_intelligence_foundation_architecture_plan.md)
- [phase09_template_reuse_derived_intelligence_foundation_dev_plan.md](file:///home/dell/Projects/personal-software-company-os/docs/phase/phase09_template_reuse_derived_intelligence_foundation_dev_plan.md)
- [phase09_template_reuse_derived_intelligence_foundation_shared_baseline.md](file:///home/dell/Projects/personal-software-company-os/docs/phase/phase09_template_reuse_derived_intelligence_foundation_shared_baseline.md)

## 2.2 最近完成正式业务 phase 的规划记录

- [phase08_operating_review_loop_foundation_architecture_plan.md](file:///home/dell/Projects/personal-software-company-os/docs/phase/phase08_operating_review_loop_foundation_architecture_plan.md)
- [phase08_operating_review_loop_foundation_dev_plan.md](file:///home/dell/Projects/personal-software-company-os/docs/phase/phase08_operating_review_loop_foundation_dev_plan.md)
- [phase08_operating_review_loop_foundation_shared_baseline.md](file:///home/dell/Projects/personal-software-company-os/docs/phase/phase08_operating_review_loop_foundation_shared_baseline.md)
- [phase08-11 正式验收与收口入口](file:///home/dell/Projects/personal-software-company-os/.trae/specs/phase08_11_validate_review_loop_integration_browser_regression_acceptance/acceptance_report.md)

## 2.3 最近完成前置基础阶段的规划记录

- [phase07_transport_contract_mainline_migration_architecture_plan.md](file:///home/dell/Projects/personal-software-company-os/docs/phase/phase07_transport_contract_mainline_migration_architecture_plan.md)
- [phase07_transport_contract_mainline_migration_dev_plan.md](file:///home/dell/Projects/personal-software-company-os/docs/phase/phase07_transport_contract_mainline_migration_dev_plan.md)
- [phase07_transport_contract_mainline_migration_shared_baseline.md](file:///home/dell/Projects/personal-software-company-os/docs/phase/phase07_transport_contract_mainline_migration_shared_baseline.md)
- [phase07-07 正式规格入口](file:///home/dell/Projects/personal-software-company-os/.trae/specs/phase07_07_formal_transport_mainline_cutover_spec/transport_mainline_cutover_spec_v0.1.md)
- [phase07-11 正式验收与收口入口](file:///home/dell/Projects/personal-software-company-os/.trae/specs/phase07_11_validate_phase01_phase06_regression_retirement_acceptance/acceptance_report.md)

## 3. 规则

- 每个 `phase` 至少包含：
  - `phase*_architecture_plan.md`
  - `phase*_dev_plan.md`
  - `phase*_shared_baseline.md`
- `/plan` 通过后，才能继续 `/spec` 与实现
- 已完成收口的 `phase`，后续实现应从正式 `/spec` 进入，而不是把三件套继续当作并列执行入口
- 后续 `dry-run` 或其他阶段在正式入口建立前，只允许以进入条件形式出现，不得提前命名未建立的新阶段
