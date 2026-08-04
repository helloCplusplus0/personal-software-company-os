# Phase Index

## 1. 职责

本目录承载项目推进主线。

任何新的功能推进、规格推进、结构性建设，都必须从这里开始，而不是直接创建孤立文档。

## 2. 当前状态

- 当前阶段已完成：`phase01_mvp_spec_convergence`
- `phase01_*` 三件套已完成规划侧收口
- 当前执行层规格入口：`.trae/specs/phase01_06_formal_mvp_spec/mvp_spec_v0.1.md`
- 当前项目技术路线：`Durable System Track`

## 3. 规则

- 每个 `phase` 至少包含：
  - `phase*_architecture_plan.md`
  - `phase*_dev_plan.md`
  - `phase*_shared_baseline.md`
- `/plan` 通过后，才能继续 `/spec` 与实现
- 已完成收口的 `phase`，后续实现应从正式 `/spec` 进入，而不是把三件套继续当作并列执行入口
