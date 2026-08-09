# plan.md

# Personal Software Company OS Plan

## 1. 当前状态

- 当前阶段：`phase05_dashboard_feedback_foundation`
- 当前状态：`phase05` 已完成 `/plan -> /spec -> 实现 -> 验收 -> 收口`
- 当前目标：保持根级真相源与 `phase05` 已交付边界一致，并等待下一阶段正式 phase 入口建立后切换
- 当前下一阶段入口：待正式建立后切换（直接承接 `phase05-10 / 11 / 14`，不得预设新的 phase 名称）

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
- `phase03_decision_center_foundation` 已完成 `/plan -> /spec -> 实现 -> 验收 -> 收口`
- `Decision Center` 最小主线已形成可运行、可验收交付物
- `.trae/specs/phase03_10_decision_center_formal_spec/decision_center_spec_v0.1.md` 已冻结为 `Decision Center` 当前阶段唯一规格收敛入口
- `.trae/specs/phase03_11_decision_center_proto_mainline/` 已将 `Decision Center` 最小 `.proto` 合同落地为仓库主线
- `.trae/specs/phase03_14_decision_center_integration_validation_acceptance/` 已形成 `phase03` 联调验收与收口结论入口
- `phase04_product_and_repository_binding_foundation` 已完成 `/plan -> /spec -> 实现 -> 验收 -> 收口`
- `Product Registry + Repository Binding` 最小主线已形成可运行、可验收交付物
- `.trae/specs/phase04_10_product_repository_binding_formal_spec/product_repository_binding_spec_v0.1.md` 已冻结为 `Product / Repository / Binding` 当前阶段唯一规格收敛入口
- `.trae/specs/phase04_11_product_repository_binding_proto_mainline/` 已将 `Product / Repository / Binding` 最小 `.proto` 合同落地为仓库主线
- `.trae/specs/phase04_14_product_repository_binding_integration_validation_acceptance/acceptance_report.md` 已形成 `phase04` 联调验收与收口结论入口
- `phase05_dashboard_feedback_foundation` 已完成 `/plan -> /spec -> 实现 -> 验收 -> 收口`
- `Dashboard + Feedback` 最小主线已形成可运行、可验收交付物
- `.trae/specs/phase05_10_dashboard_feedback_formal_spec/dashboard_feedback_spec_v0.1.md` 已冻结为 `Dashboard + Feedback` 当前阶段唯一规格收敛入口
- `.trae/specs/phase05_11_dashboard_feedback_proto_mainline/` 已将 `Dashboard + Feedback` 最小 `.proto` 合同落地为仓库主线
- `.trae/specs/phase05_14_dashboard_feedback_integration_validation_acceptance/acceptance_report.md` 已形成 `phase05` 联调验收与收口结论入口
- `phase05` 三件套保留为该阶段 `/plan` 的规划与冻结记录，不再覆盖根级当前状态
- 下一阶段正式 phase 入口尚未在仓库中建立；根级阶段切换保持“待建立后切换”显式状态

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
- 状态：`completed`
- 当前收口结果：已交付 `Decision Center` 正式规格正文、最小 `.proto` 合同主线、后端与数据主线、前端主线与联调验收收口结果

### phase04：`phase04_product_and_repository_binding_foundation`

- 目标：交付 Product / Repository / Module Binding 主线
- 进入条件：直接承接 `.trae/specs/phase03_10_decision_center_formal_spec/decision_center_spec_v0.1.md`、`.trae/specs/phase03_11_decision_center_proto_mainline/` 与 `.trae/specs/phase03_14_decision_center_integration_validation_acceptance/`，以 `phase03` 已完成的 `Decision Center` 交付物作为唯一上游输入
- 范围约束：不得回退重做 `Decision Center`，而是以前一阶段已冻结并验收的能力为前提，推进 `Product / Repository / Module Binding` 最小主线
- 交付要求：作为交付型 phase 推进，不得只停留在规格冻结
- 状态：`completed`
- 当前收口结果：已交付 `Product / Repository / Binding` 正式规格正文、最小 `.proto` 合同主线、后端与数据主线、前端主线与联调验收收口结果

### phase05：`phase05_dashboard_feedback_foundation`

- 目标：交付 Dashboard + Feedback 最小闭环
- 进入条件：直接承接 `.trae/specs/phase04_10_product_repository_binding_formal_spec/product_repository_binding_spec_v0.1.md`、`.trae/specs/phase04_11_product_repository_binding_proto_mainline/` 与 `.trae/specs/phase04_14_product_repository_binding_integration_validation_acceptance/acceptance_report.md`，以 `phase04` 已完成的 `Product Registry + Repository Binding` 交付物作为唯一上游输入
- 范围约束：不得重复实现 `Product Registry`、`Repository Binding` 与三类绑定动作主线，而是以 `phase04` 已冻结页面、动作、数据、合同与联调结论为前提推进 `Dashboard + Feedback`
- 交付要求：作为交付型 phase 推进，不得只停留在规格冻结
- 状态：`completed`
- 当前收口结果：已交付 `Dashboard + Feedback` 正式规格正文、最小 `.proto` 合同主线、后端与数据主线、前端主线与联调验收收口结果

## 4. 下一阶段切换条件

当以下条件同时满足时，根级入口才允许从 `phase05` 已完成收口状态切换到下一阶段正式 phase 入口：

1. `phase05` 已完成 `/plan -> /spec -> 实现 -> 验收 -> 收口` 全链路
2. `.trae/specs/phase05_10_dashboard_feedback_formal_spec/dashboard_feedback_spec_v0.1.md` 已冻结为 `Dashboard + Feedback` 当前阶段唯一正式规格入口
3. `.trae/specs/phase05_11_dashboard_feedback_proto_mainline/` 已将 `Dashboard + Feedback` 最小 `.proto` 合同落地为仓库合同主线
4. `.trae/specs/phase05_14_dashboard_feedback_integration_validation_acceptance/acceptance_report.md` 已形成联调验收与问题收口结论
5. 下一阶段正式 `phase` 入口文档已在仓库中建立，并明确直接承接 `phase05-10 / 11 / 14`
6. 根级真相源已完成状态切换回写，且未在根级文档中凭空猜测未建立的 phase 名称

当前结论：`phase05` 已完成收口，下一阶段的直接上游输入已冻结为 `phase05-10` 正式规格正文、`phase05-11` 合同主线与 `phase05-14` 验收结论；由于下一阶段正式 phase 入口尚未建立，根级入口当前保持“待建立后切换”状态。

## 5. 说明

- 本文档只承载全局开发预览、phase 计划与进度
- 不展示 task 级拆分
- task 级内容只进入对应 `docs/phase/phase*_*_dev_plan.md`
- `phase01` 是规格收敛特例；`phase02+` 默认按交付型 phase 推进
