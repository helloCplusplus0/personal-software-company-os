# plan.md

# Personal Software Company OS Plan

## 1. 当前状态

- 当前阶段：`phase01_mvp_spec_convergence`
- 当前状态：`completed`
- 当前目标：完成 `phase01` 收口，并以正式 MVP 规格正文作为 `phase02` 的唯一执行层上游
- 当前下一阶段入口：`phase02_module_registry_foundation`

## 2. 当前进度概览

- 原始方案文档 `PSCO_0.md ~ PSCO_4.md` 已完成第一轮共识回正
- 根级真相源职责已完成第一轮去重
- `docs/` 已收口到 `phase / fix / audit / review / archive`
- 评审与交叉汇总文档已归类到 `docs/review/`
- `phase01_*` 三件套已完成规划侧收口
- `.trae/specs/phase01_06_formal_mvp_spec/mvp_spec_v0.1.md` 已冻结为执行层唯一规格入口
- 当前项目技术路线已明确为 `Durable System Track`

## 3. Phase 路线预览

### phase01：`phase01_mvp_spec_convergence`

- 目标：完成 MVP 规格收敛
- 状态：`completed`
- 产出：`architecture_plan / dev_plan / shared_baseline + phase01-01 ~ phase01-07 + formal_mvp_spec`

### phase02：`phase02_module_registry_foundation`

- 目标：建立 Module Registry 最小可执行主线
- 进入条件：以 `.trae/specs/phase01_06_formal_mvp_spec/mvp_spec_v0.1.md` 为唯一执行层上游，只承接已冻结的 `v0.1` 边界
- 范围约束：不得重新引入 `Feature / Opportunity / Experiment`、独立 `AI Assistant`、独立 `React Native` 客户端或完整 `PWA` 能力作为前置范围
- 状态：`draft`

### phase03：`phase03_decision_center_foundation`

- 目标：建立 Decision Center 最小闭环
- 状态：`draft`

### phase04：`phase04_product_and_repository_binding_foundation`

- 目标：建立 Product / Repository / Module Binding 主线
- 状态：`draft`

### phase05：`phase05_dashboard_feedback_foundation`

- 目标：建立 Dashboard 最小反馈闭环
- 状态：`draft`

## 4. 当前阶段完成标志

当以下条件同时满足时，当前阶段结束，并进入 `phase02_module_registry_foundation`：

1. 根级文档职责稳定
2. `plan.md` 已只承载 phase 级预览
3. `AGENTS.md` 已能作为稳定上下文入口
4. `docs/` 结构与 workflow 完全一致
5. `.trae/specs/phase01_06_formal_mvp_spec/mvp_spec_v0.1.md` 已冻结为执行层唯一规格入口
6. `phase02` 已明确只承接 `v0.1` 冻结范围，不重新解释前端端策略或后移对象边界

## 5. 说明

- 本文档只承载全局开发预览、phase 计划与进度
- 不展示 task 级拆分
- task 级内容只进入对应 `docs/phase/phase*_*_dev_plan.md`
