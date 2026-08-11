# plan.md

# Personal Software Company OS Plan

## 1. 当前状态

- 当前阶段：`phase08_operating_review_loop_foundation`
- 当前状态：`phase08` 已建立正式 `/plan` 入口，当前作为 `mvp0.3` 的首个正式业务 phase 推进
- 当前目标：让 Dashboard 正式承担经营动作入口，并闭合 `Feedback -> Decision -> Update` 最小经营回路
- 当前下一阶段入口：待 `phase08` 正式收口后，后续 `mvp0.3` 支撑能力 phase 才允许建立正式入口；当前只保留进入条件表达

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
- `PSCO-mvp02-summarize-feedback.md` 已冻结为 `mvp0.2` 下一阶段 `/plan` 的最终仲裁与规划基线
- `phase06_onboarding_sovereignty_reuse_foundation` 已完成 `/plan -> /spec -> 实现 -> 验收 -> 收口`
- `Onboarding + Data Sovereignty + Reuse Awareness` 最小主线已形成可运行、可验收交付物
- `phase07-08` 已完成 `buf + ConnectRPC` 正式合同产物主线切换，并通过 `proto / backend / frontend` 生成链验证
- `phase07-09` 已完成 Go 后端业务传输主线切换、源码优先独立复核、阻断修复与运行时验收
- `phase07-10` 已完成前端 generated client / owner 主线切换、旧 adapter 与 compat facade 退场，并通过前端构建与静态核销验收
- `.trae/specs/phase07_07_formal_transport_mainline_cutover_spec/transport_mainline_cutover_spec_v0.1.md` 已冻结为 `phase07` 当前阶段唯一正式规格入口
- `.trae/specs/phase07_11_validate_phase01_phase06_regression_retirement_acceptance/acceptance_report.md` 已形成 `phase07` 联调、退场验收与收口结论入口
- `.trae/specs/phase06_12_onboarding_sovereignty_reuse_formal_spec/spec.md` 已冻结为 `phase06` 当前阶段唯一规格收敛入口
- `.trae/specs/phase06_13_land_minimal_proto_contract_mainline/` 已将 `phase06` 最小 `.proto` 合同落地为仓库主线
- `.trae/specs/phase06_16_integration_validation_acceptance/acceptance_report.md` 已形成 `phase06` 联调验收与收口结论入口
- `phase06` 三件套保留为该阶段 `/plan` 的规划与冻结记录，不再覆盖根级当前状态
- `PSCO-mvp03-summarize-feedback.md` 已冻结为 `phase08` 与后续 `mvp0.3` 业务阶段的最终仲裁与完整路线规划基线
- `docs/audit/audit_001_transport_contract_mainline_issue.md` 与 `docs/audit/audit_001_transport_contract_mainline_analysis.md` 已冻结为当前传输主线收敛议题的正式审计入口
- `phase07_transport_contract_mainline_migration` 已完成 `/plan -> /spec -> 实现 -> 验收 -> 收口`，当前作为 `mvp0.3` 业务阶段之前的前置基础结论保留
- `phase07` 的正式执行层入口已收敛到 `phase07-07` 正式规格与 `phase07-11` 验收报告；三件套只保留为规划与冻结记录
- `phase08_operating_review_loop_foundation` 已建立正式 `/plan` 入口，当前作为 `mvp0.3` 的首个正式业务 phase 推进
- `phase08` 当前只承接 `Operating Review Loop`，不把 `Template Reuse / Derived Intelligence / Real-Project Dry-Run` 混写为并列主交付

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

### phase06：`phase06_onboarding_sovereignty_reuse_foundation`

- 目标：交付 `Onboarding + Data Sovereignty + Reuse Awareness` 最小主线
- 进入条件：直接承接 `PSCO-mvp02-summarize-feedback.md` 与 `phase05` 已冻结的正式规格、合同主线和联调验收结果，以 `phase05` 已完成的 Dashboard / Feedback 主线作为唯一交付上游
- 范围约束：不得重新扩写 `Opportunity / Feature / Experiment`，不得把 `Capability` 升级为重实体 CRUD，且不得把导出 / 备份后移为尾部补丁
- 交付要求：作为交付型 phase 推进，必须完成 `/plan -> /spec -> 实现 -> 验收 -> 收口`
- 状态：`completed`
- 当前收口结果：已交付 `Onboarding + Data Sovereignty + Reuse Awareness` 正式规格正文、最小 `.proto` 合同主线、后端与数据主线、前端主线与联调验收收口结果

### phase07：`phase07_transport_contract_mainline_migration`

- 目标：交付 `phase01 ~ phase06` canonical 业务接口向 `.proto + ConnectRPC` 的正式传输主线切换
- 进入条件：直接承接 `PSCO-mvp03-summarize-feedback.md`、`audit_001` 审计结论与 `phase06` 已冻结的正式规格、合同主线和联调验收结果，以 `phase06` 已完成的业务能力作为唯一迁移上游
- 范围约束：不得把本 phase 收窄为“新增接口默认走 ConnectRPC”的试点；必须完成既有 canonical 业务接口迁移，同时保留 `chi` 作为装配层与 `healthz / readyz / metrics / debug` 等非业务端点承载层
- 交付要求：作为交付型 phase 推进，必须完成 `/plan -> /spec -> 实现 -> 验收 -> 收口`
- 状态：`completed`
- 当前收口结果：已完成 `phase01 ~ phase06` canonical 业务接口向 `.proto + ConnectRPC` 的正式传输主线切换；`phase07-07` 为正式规格入口，`phase07-11` 为联调、退场验收与收口结论入口；三件套保留为规划与冻结记录

### phase08：`phase08_operating_review_loop_foundation`

- 目标：交付 `Operating Review Loop` 最小经营回路
- 进入条件：直接承接 `PSCO-mvp03-summarize-feedback.md`、`phase07-07` 正式规格、`phase07-11` 验收结论，以及 `phase03 ~ phase06` 已冻结的业务规格与验收结果
- 范围约束：不得把 `Template Reuse / Derived Intelligence Deepening / Real-Project Dry-Run` 混写为本 phase 并列主交付；不得把 review loop 演化为通用任务管理器
- 交付要求：作为交付型 phase 推进，必须完成 `/plan -> /spec -> 实现 -> 验收 -> 收口`
- 状态：`current`
- 当前阶段结果：已建立 `phase08` 三件套，当前待进入 `/spec` 收敛 `Dashboard -> Review -> Decision -> Update` 的正式执行边界

## 4. 下一阶段切换条件

当以下条件同时满足时，根级入口才允许从 `phase08` 当前阶段切换到后续 `mvp0.3` 支撑能力 phase 的正式入口：

1. `phase08` 已完成 `/plan -> /spec -> 实现 -> 验收 -> 收口` 全链路
2. Dashboard 已正式承担 review 入口职责，而不再只是总览页
3. `Feedback -> Decision -> Update` 已形成可重复执行的最小闭环
4. `phase08` 对应正式规格、验收结论与根级同步入口已形成正式证据
5. 下一阶段正式 `phase` 入口文档已在仓库中建立，并明确直接承接 `phase08` 已冻结的边界与收口结论
6. 根级真相源已完成状态切换回写，且未在根级文档中凭空猜测未建立的后续 phase 名称

当前结论：`phase08` 已建立正式 `/plan` 入口，当前直接上游已冻结为 `PSCO-mvp03-summarize-feedback.md`、`phase07-07` 正式规格、`phase07-11` 验收结论与相关已交付业务规格；在 `phase08` 正式收口前，后续支撑能力 phase 与 dry-run phase 只保留进入条件表达，不预设新的正式阶段名称。

## 5. 说明

- 本文档只承载全局开发预览、phase 计划与进度
- 不展示 task 级拆分
- task 级内容只进入对应 `docs/phase/phase*_*_dev_plan.md`
- `phase01` 是规格收敛特例；`phase02+` 默认按交付型 phase 推进
