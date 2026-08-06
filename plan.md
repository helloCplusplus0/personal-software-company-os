# plan.md

# Personal Software Company OS Plan

## 1. 当前状态

- 当前阶段：`phase03_decision_center_foundation`
- 当前状态：`/plan`
- 当前目标：完成 `Decision Center` 三件套规划，作为后续 `/spec`、实现、验收与收口的直接上游
- 当前下一阶段入口：`phase03_decision_center_foundation`

## 2. 当前进度概览

- 原始方案文档 `PSCO_0.md ~ PSCO_4.md` 已完成第一轮共识回正
- 根级真相源职责已完成第一轮去重
- `docs/` 已收口到 `phase / fix / audit / review / archive`
- 评审与交叉汇总文档已归类到 `docs/review/`
- `phase01_*` 三件套已完成规划侧收口
- `.trae/specs/phase01_06_formal_mvp_spec/mvp_spec_v0.1.md` 已冻结为执行层唯一规格入口
- `phase02_*` 三件套已建立
- 当前项目技术路线已明确为 `Durable System Track`
- `phase02_module_registry_foundation` 已完成 `/plan -> /spec -> 实现 -> 验收 -> 收口`
- `Module Registry` 最小主线已形成可运行、可验收交付物
- `.trae/specs/phase02_09_module_registry_formal_spec/module_registry_spec_v0.1.md` 已冻结为 `Module Registry` 当前阶段唯一规格收敛入口
- `.trae/specs/phase02_12_module_registry_integration_validation_acceptance/acceptance_report.md` 已给出通过验收结论

## 3. Phase 路线预览

### phase01：`phase01_mvp_spec_convergence`

- 目标：完成 MVP 规格收敛
- 状态：`completed`
- 产出：`architecture_plan / dev_plan / shared_baseline + phase01-01 ~ phase01-07 + formal_mvp_spec`

### phase02：`phase02_module_registry_foundation`

- 目标：交付 Module Registry 最小可执行主线
- 进入条件：以 `.trae/specs/phase01_06_formal_mvp_spec/mvp_spec_v0.1.md` 为唯一执行层上游，只承接已冻结的 `v0.1` 边界
- 范围约束：不得重新引入 `Feature / Opportunity / Experiment`、独立 `AI Assistant`、独立 `React Native` 客户端或完整 `PWA` 能力作为前置范围
- 交付要求：本 phase 必须完成 `/plan -> /spec -> 实现 -> 验收 -> 收口`
- 状态：`completed`
- 当前收口结果：已交付 `Module Registry` 前后端最小主线、数据主线、最小 `.proto` 合同源与联调验收结果

### phase03：`phase03_decision_center_foundation`

- 目标：交付 Decision Center 最小闭环
- 进入条件：直接承接 `phase02` 已交付的 `Module Registry` 主线，以 `.trae/specs/phase02_09_module_registry_formal_spec/module_registry_spec_v0.1.md`、`.trae/specs/phase02_11a_module_registry_proto_contract/` 与 `.trae/specs/phase02_12_module_registry_integration_validation_acceptance/acceptance_report.md` 为直接上游输入
- 交付要求：作为交付型 phase 推进，不得只停留在规格冻结
- 状态：`/plan`

### phase04：`phase04_product_and_repository_binding_foundation`

- 目标：交付 Product / Repository / Module Binding 主线
- 交付要求：作为交付型 phase 推进，不得只停留在规格冻结
- 状态：`draft`

### phase05：`phase05_dashboard_feedback_foundation`

- 目标：交付 Dashboard 最小反馈闭环
- 交付要求：作为交付型 phase 推进，不得只停留在规格冻结
- 状态：`draft`

## 4. 当前阶段完成标志

当以下条件同时满足时，`phase03_decision_center_foundation` 结束，并进入 `phase04_product_and_repository_binding_foundation`：

1. `phase03` 已完成 `/plan -> /spec -> 实现 -> 验收 -> 收口` 全链路
2. `Decision Center` 已具备可运行、可验收的前后端最小主线
3. `Decision` 的页面、动作、数据、API、合同与验收基线已单值化
4. `phase03` 已明确只承接已冻结的 `v0.1` 范围，不重新解释后移对象边界
5. 单一 `React Web` 同时覆盖 `PC` 与移动浏览器的交付策略已写清并落实到实现
6. `Decision -> Module` 的最小闭环已完整打通
7. `phase04` 的进入条件已清楚

当前结论：`phase03` 当前处于 `/plan`，完成标志如上，后续进入 `/spec`、实现、验收与收口后再切换到 `phase04`。

## 5. 说明

- 本文档只承载全局开发预览、phase 计划与进度
- 不展示 task 级拆分
- task 级内容只进入对应 `docs/phase/phase*_*_dev_plan.md`
- `phase01` 是规格收敛特例；`phase02+` 默认按交付型 phase 推进
