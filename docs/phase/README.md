# Phase Index

## 1. 职责

本目录承载项目推进主线。

任何新的功能推进、规格推进、结构性建设，都必须从这里开始，而不是直接创建孤立文档。

## 2. 当前状态

- 根级当前状态：`phase07_transport_contract_mainline_migration` 已建立正式 `/plan` 入口
- `phase01_*` 三件套已完成规划侧收口
- `phase02_*` 三件套已完成并已收口
- `phase03_*` 三件套已完成并已收口
- `phase04_*` 三件套已完成并已收口
- `phase05_*` 三件套保留为最近完成阶段的规划与冻结记录
- `phase06_*` 三件套已完成并保留为最近完成阶段的规划与冻结记录
- `phase07_*` 三件套已建立，并作为当前阶段正式 `/plan` 入口
- 下一阶段正式 phase 入口：待 `phase07` 正式收口后切换（不得预设新的 phase 名称）
- 当前阶段直接上游已冻结为：`PSCO-mvp03-summarize-feedback.md`、`docs/audit/audit_001_transport_contract_mainline_analysis.md`、`.trae/specs/phase06_12_onboarding_sovereignty_reuse_formal_spec/spec.md`、`.trae/specs/phase06_13_land_minimal_proto_contract_mainline/`、`.trae/specs/phase06_16_integration_validation_acceptance/acceptance_report.md`
- 当前项目技术路线：`Durable System Track`

## 2.1 当前活动阶段的规划记录

- [phase07_transport_contract_mainline_migration_architecture_plan.md](file:///home/dell/Projects/personal-software-company-os/docs/phase/phase07_transport_contract_mainline_migration_architecture_plan.md)
- [phase07_transport_contract_mainline_migration_dev_plan.md](file:///home/dell/Projects/personal-software-company-os/docs/phase/phase07_transport_contract_mainline_migration_dev_plan.md)
- [phase07_transport_contract_mainline_migration_shared_baseline.md](file:///home/dell/Projects/personal-software-company-os/docs/phase/phase07_transport_contract_mainline_migration_shared_baseline.md)

## 2.2 最近完成阶段的规划记录

- [phase06_onboarding_sovereignty_reuse_foundation_architecture_plan.md](file:///home/dell/Projects/personal-software-company-os/docs/phase/phase06_onboarding_sovereignty_reuse_foundation_architecture_plan.md)
- [phase06_onboarding_sovereignty_reuse_foundation_dev_plan.md](file:///home/dell/Projects/personal-software-company-os/docs/phase/phase06_onboarding_sovereignty_reuse_foundation_dev_plan.md)
- [phase06_onboarding_sovereignty_reuse_foundation_shared_baseline.md](file:///home/dell/Projects/personal-software-company-os/docs/phase/phase06_onboarding_sovereignty_reuse_foundation_shared_baseline.md)

## 3. 规则

- 每个 `phase` 至少包含：
  - `phase*_architecture_plan.md`
  - `phase*_dev_plan.md`
  - `phase*_shared_baseline.md`
- `/plan` 通过后，才能继续 `/spec` 与实现
- 已完成收口的 `phase`，后续实现应从正式 `/spec` 进入，而不是把三件套继续当作并列执行入口
