# Phase Index

## 1. 职责

本目录承载项目推进主线。

任何新的功能推进、规格推进、结构性建设，都必须从这里开始，而不是直接创建孤立文档。

## 2. 当前状态

- 当前阶段：`phase05_dashboard_feedback_foundation (/plan)`
- `phase01_*` 三件套已完成规划侧收口
- `phase02_*` 三件套已完成并已收口
- `phase03_*` 三件套已完成并已收口
- `phase04_*` 三件套已完成并已收口
- `phase05_*` 三件套已建立，当前处于 `/plan`
- 当前阶段直接上游：`.trae/specs/phase04_10_product_repository_binding_formal_spec/product_repository_binding_spec_v0.1.md`、`.trae/specs/phase04_11_product_repository_binding_proto_mainline/`、`.trae/specs/phase04_14_product_repository_binding_integration_validation_acceptance/acceptance_report.md`
- 当前项目技术路线：`Durable System Track`

## 2.1 当前活动文档

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
