# plan.md

# Personal Software Company OS Plan

## 1. 当前状态

- 当前阶段：`phase15` 进入条件已冻结，等待用户裁决开启 `phase15 /plan`
- 当前状态：`phase14_standard_entity_foundation` 已完成正式 `/plan -> /spec -> 实现 -> 验收 -> 收口`；`phase13` 作为上一完成正式业务 phase 保留，`phase09` 继续作为最近完成正式支撑能力 phase
- 当前最近完成正式验收与收口入口：`.trae/specs/phase14_10_validate_standard_entity_integration_dogfooding_regression/acceptance_report.md`
- 当前目标：`phase15 /plan` 前须完成后续项排序裁决——候选池与边界冻结于 `phase14-11` spec（`CON-08` 时间轴〔T7 裁决后口径：新建正规承接，不复活画像派生形态〕/ `standard_bindings` 目标类型扩展 / agent 写回〔不自动解锁〕/ Git 推进跟踪 / 模板仓库自动接入 / 自动同步）
- 当前下一阶段进入条件：已冻结于 `.trae/specs/phase14_11_sync_root_level_closeout_freeze_phase15_entry_conditions/spec.md`

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
- `.trae/specs/phase07_07_formal_transport_mainline_cutover_spec/transport_mainline_cutover_spec_v0.1.md` 已冻结为 `phase07` 最近完成前置基础阶段的正式规格入口
- `.trae/specs/phase07_11_validate_phase01_phase06_regression_retirement_acceptance/acceptance_report.md` 已形成 `phase07` 联调、退场验收与前置基础阶段收口结论入口
- `.trae/specs/phase06_12_onboarding_sovereignty_reuse_formal_spec/spec.md` 已冻结为 `phase06` 当前阶段唯一规格收敛入口
- `.trae/specs/phase06_13_land_minimal_proto_contract_mainline/` 已将 `phase06` 最小 `.proto` 合同落地为仓库主线
- `.trae/specs/phase06_16_integration_validation_acceptance/acceptance_report.md` 已形成 `phase06` 联调验收与收口结论入口
- `phase06` 三件套保留为该阶段 `/plan` 的规划与冻结记录，不再覆盖根级当前状态
- `PSCO-mvp03-summarize-feedback.md` 已冻结为 `phase08`、`phase09` 与后续 `dry-run` 进入条件的最终仲裁与规划基线
- `docs/audit/audit_001_transport_contract_mainline_issue.md` 与 `docs/audit/audit_001_transport_contract_mainline_analysis.md` 已冻结为当前传输主线收敛议题的正式审计入口
- `phase07_transport_contract_mainline_migration` 已完成 `/plan -> /spec -> 实现 -> 验收 -> 收口`，当前退回最近完成前置基础阶段的规划与冻结记录角色
- `phase07` 的正式执行层入口已收敛到 `phase07-07` 正式规格与 `phase07-11` 验收报告；三件套只保留为规划与冻结记录，不再覆盖当前业务状态
- `phase08_operating_review_loop_foundation` 已完成 `/plan -> /spec -> 实现 -> 验收 -> 收口`，当前作为已完成正式业务 phase 保留
- `phase08_operating_review_loop_foundation` 三件套继续保留为已完成正式业务 phase 的规划与冻结记录，不再承接最近完成正式业务 phase 的根级角色
- `.trae/specs/phase08_11_validate_review_loop_integration_browser_regression_acceptance/acceptance_report.md` 已形成 `phase08` 统一联调、浏览器验收、反回归验证与正式收口结论入口
- `phase08` 已完成 `Dashboard` review 入口、Daily / Weekly Review 双路径会话承接，以及 `Feedback -> Decision -> Update` 最小经营回路，并通过统一联调与浏览器验收收口
- `phase08` 当前只承接 `Operating Review Loop`，不把 `Template Reuse / Derived Intelligence / Real-Project Dry-Run` 混写为并列主交付
- `phase09_template_reuse_derived_intelligence_foundation` 已完成 `/plan -> /spec -> 实现 -> 验收 -> 收口`，当前作为最近完成正式支撑能力 phase 保留
- `docs/phase/phase09_template_reuse_derived_intelligence_foundation_architecture_plan.md`、`docs/phase/phase09_template_reuse_derived_intelligence_foundation_dev_plan.md` 与 `docs/phase/phase09_template_reuse_derived_intelligence_foundation_shared_baseline.md` 已保留为 `phase09` 的规划与冻结记录入口
- `.trae/specs/phase09_11_validate_template_reuse_derived_hint_integration_browser_regression_acceptance/acceptance_report.md` 已形成 `phase09` 统一联调、浏览器验收、反回归验证与正式收口结论入口
- `phase09` 已完成模板候选、`Product Create` 预填回流、派生提示展示与 handoff 的正式交付，并通过统一联调、浏览器验收与反回归验证收口
- `phase09` 当前只承接 `Template Reuse + Derived Intelligence Deepening` 的最小支撑能力，不把 `Real-Project Dry-Run`、`Venture` 或 `AI Context Enhancement` 偷渡为并列主交付
- `fix_001 ~ fix_003` 已完成修复、独立复核、聚焦 rerun 与 `mvp0.3` 收口，`docs/fix/fix_001_003_mvp03_dry_run_rerun_closure.md` 已形成进入 `mvp0.4` 首个正式 phase 的直接结论入口
- `phase10_asset_action_closure_foundation` 三件套已完成 `/plan -> /spec -> 实现 -> 验收 -> 收口`，保留为历史完成正式业务 phase 的规划与冻结记录
- `.trae/specs/phase10_11_complete_asset_action_closure_integration_browser_regression_validation/spec.md` 已冻结为 `phase10-11` 的正式验收入口；同目录 `tasks.md / checklist.md` 共同构成 `phase10` 的正式验收与收口证据
- `phase10` 已完成 `Onboarding` 首轮建链、`Decision` 生命周期闭环、关键 detail pages 动作承接矩阵，以及 `Current Focus / pending signals` 反回归验证；更重能力方向已统一顺延为 `phase11` 正式收口后的下一阶段进入条件
- `PSCO-mvp05-summarize-feedback.md` 已冻结为 `mvp0.5` 的最终仲裁与 `phase11 / phase12` 的直接共识上游
- `.trae/specs/phase11_09_validate_project_context_foundation_dogfooding_regression/acceptance_report.md` 已冻结为 `phase11` 的正式验收与收口入口
- `phase11_project_context_foundation` 已完成正式 `/plan -> /spec -> 实现 -> 验收 -> 收口`，并完成“根级上下文真相源治理 + 最小只读项目上下文导出”交付；下一阶段只允许在本阶段正式收口后，再讨论或进入更重的 `MCP / CLI / agent 写回 / 更重消费通道 / 受控维护能力`
- `docs/audit/audit_002_phase11_post_closeout_direction_analysis.md` 已冻结为 `phase11` 收口后下一步方向的正式审计结论入口
- `.trae/specs/phase12_11_validate_semantic_alignment_readonly_consumption_foundation/acceptance_report.md` 已冻结为 `phase12` 的正式验收与收口入口
- `phase12_semantic_alignment_and_readonly_consumption_foundation` 已完成正式 `/plan -> /spec -> 实现 -> 验收 -> 收口`，并完成“前端四实体语义一致性收口 + Web / agent 共享只读消费深化”交付；`Dashboard / Review / Onboarding / Detail pages` 的语义表达与共享只读消费已完成阶段收口
- `phase13_project_governance_profile_foundation` 已完成正式 `/plan -> /spec -> 实现 -> 验收 -> 收口`，并完成“项目级治理画像 + 全局规范资产 + agent 项目简报输入”交付
- `.trae/specs/phase13_11_validate_project_governance_profile_integration_dogfooding_regression/acceptance_report.md` 已冻结为 `phase13` 的正式验收与收口入口
- `.trae/specs/phase13_12_sync_root_level_closeout_next_phase_entry_conditions/spec.md` 已冻结为 `phase13` 正式缺口记录（GAP-01 ~ GAP-07 + CON-01 ~ CON-09）与 `phase14 /plan` 的唯一直接上游
- `phase14_standard_entity_foundation` 三件套已建立并完成独立复核：五项裁决结论、`Standard` 实体设计（主表+jsonb树+多态绑定）、画像退役映射与 11 个子任务已冻结
- `phase14_standard_entity_foundation` 已完成正式 `/plan -> /spec -> 实现 -> 验收 -> 收口`，并完成“`Standard` 全局规范实体最小主线 + 治理画像系统性退役 + brief 画像残余解耦（T7 用户裁决）”交付；`phase14` 现为最近完成正式业务 phase
- `.trae/specs/phase14_10_validate_standard_entity_integration_dogfooding_regression/acceptance_report.md` 已冻结为 `phase14` 的正式验收与收口入口
- `.trae/specs/phase14_11_sync_root_level_closeout_freeze_phase15_entry_conditions/spec.md` 已冻结为 `phase15` 进入条件（T7 裁决后口径）与 `phase15 /plan` 的唯一直接上游；`phase13-12` 缺口记录退位为历史输入

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
- 状态：`completed`
- 当前收口结果：已完成 `Dashboard` review 入口、Daily / Weekly Review 双路径会话承接与 `Feedback -> Decision -> Update` 最小经营回路；`phase08-11` 为统一联调、浏览器验收、反回归验证与正式收口结论入口，三件套保留为规划与冻结记录

### phase09：`phase09_template_reuse_derived_intelligence_foundation`

- 目标：交付 `Template Reuse + Derived Intelligence Deepening` 最小支撑能力
- 进入条件：直接承接 `PSCO-mvp03-summarize-feedback.md`、`phase08-11` 正式验收结论、`phase08` 三件套规划记录、`phase06` 复用摘要主线，以及 `phase03 ~ phase05` 已冻结的 `Decision / Product / Dashboard` 正式规格与验收结果
- 范围约束：不得把本 phase 扩写为模板平台、AI 工作台或 `dry-run`；`Product Create` 必须继续保持唯一 canonical 创建承接位；支撑能力必须直接服务“下一次创造”，而不是另起第二条业务主线
- 交付要求：作为交付型 phase 推进，必须完成 `/plan -> /spec -> 实现 -> 验收 -> 收口`
- 状态：`completed`
- 当前收口结果：已完成模板候选、`Product Create` 预填回流、派生提示展示与 handoff 的正式交付；`phase09-11` 为统一联调、浏览器验收、反回归验证与正式收口结论入口，三件套保留为规划与冻结记录；当前作为最近完成正式支撑能力 phase 保留

### phase10：`phase10_asset_action_closure_foundation`

- 目标：交付 `Asset-Action Closure` 主线
- 进入条件：直接承接 `PSCO-mvp04-summarize-feedback.md`、`docs/fix/fix_001_003_mvp03_dry_run_rerun_closure.md`、`fix_001 ~ fix_003` 三份 analysis、以及 `phase06 / phase08 / phase09` 已冻结的正式规格与验收结果
- 范围约束：不得提前混入 `Agent Consumption Layer`、`Cross-Project Convention Asset`、新实体主线、第五态 `DecisionStatus`、AI 工作台或真实连接重型集成
- 交付要求：作为交付型 phase 推进，必须完成 `/plan -> /spec -> 实现 -> 验收 -> 收口`
- 状态：`completed`
- 当前收口结果：已完成 `Onboarding` 首轮建链引导、`Decision` 最小真实生命周期、Dashboard / Review / Detail pages 下一步动作承接矩阵，以及 `Current Focus / pending signals` 真实经营语义回归；`.trae/specs/phase10_11_complete_asset_action_closure_integration_browser_regression_validation/spec.md` 为正式验收入口，三件套保留为历史完成正式业务 phase 的规划与冻结记录；更重能力方向已统一顺延为 `phase11` 正式收口后的下一阶段进入条件

### phase11：`phase11_project_context_foundation`

- 目标：交付 `Project Context Foundation`
- 进入条件：直接承接 `PSCO-mvp05-summarize-feedback.md`、`phase10` 三件套、`phase10-11` 正式验收入口，以及当前根级真相源文档的已知漂移治理需求
- 范围约束：不得提前混入 MCP / CLI / agent 写回 / 前端对话式入口 / 四实体结构重构 / 第二套 canonical API；只允许推进“根级上下文真相源治理 + 最小只读项目上下文导出”
- 交付要求：作为交付型 phase 推进，必须先完成 `/plan` 三件套复核，再按 `dev_plan` 子任务顺序进入 `/spec`、实现、验收与收口
- 状态：`completed`
- 当前收口结果：已完成“根级上下文真相源治理 + 最小只读项目上下文导出”正式交付；`.trae/specs/phase11_09_validate_project_context_foundation_dogfooding_regression/acceptance_report.md` 为正式验收与收口入口；更重的 `MCP / CLI / agent 写回 / 更重消费通道 / 受控维护能力` 只允许在 `phase11` 正式收口后，作为下一阶段进入条件讨论或进入

### phase12：`phase12_semantic_alignment_and_readonly_consumption_foundation`

- 目标：交付 `Semantic Alignment & Read-Only Consumption Foundation`
- 进入条件：直接承接 `PSCO-mvp05-summarize-feedback.md`、`audit_002` 审计结论、`phase11` 三件套与 `phase11-09` 正式验收入口，以及当前前端四实体语义表达与只读消费入口尚未完全收口的已知差距
- 范围约束：不得提前混入 MCP / CLI / agent 写回 / 前端对话式入口 / 四实体结构重构 / 第二套 canonical API；只允许推进“前端四实体语义一致性收口 + Web / agent 共享只读消费深化”
- 交付要求：作为交付型 phase 推进，必须先完成 `/plan` 三件套复核，再按 `dev_plan` 子任务顺序进入 `/spec`、实现、验收与收口
- 状态：`completed`
- 当前收口结果：已完成“前端四实体语义一致性收口 + Web / agent 共享只读消费深化”正式交付；`.trae/specs/phase12_11_validate_semantic_alignment_readonly_consumption_foundation/acceptance_report.md` 为正式验收与收口入口；更重的项目推进跟踪、模板仓库接入、自动同步与受控维护能力只允许在 `phase13` 正式收口后，再作为下一阶段进入条件讨论或进入

### phase13：`phase13_project_governance_profile_foundation`

- 目标：交付 `Project Governance Profile Foundation`
- 进入条件：直接承接 `phase12` 三件套、`phase12-11` 正式验收与收口入口、当前根级真相源文档，以及已形成的新共识：PSCO 优先管理四实体信息、全局规范资产与项目级治理画像，而不把 IDE 目录即时上下文默认上升为正式事实源
- 范围约束：不得把本 phase 扩写为 Git 推进跟踪平台、模板仓库自动化、全量目录扫描器、MCP / CLI / agent 写回或 IDE 插件；只允许推进“项目级治理画像 + 全局规范资产 + agent 项目简报输入”的最小正式主线
- 交付要求：作为交付型 phase 推进，必须先完成 `/plan` 三件套复核，再按 `dev_plan` 子任务顺序进入 `/spec`、实现、验收与收口
- 状态：`completed`
- 当前收口结果：已完成“项目级治理画像 + 全局规范资产 + agent 项目简报输入”正式交付；`phase13-11` 为正式验收与收口入口，`phase13-12` 为正式缺口记录与 `phase14` 进入条件冻结入口；`phase14 /plan` 必须以缺口记录为唯一直接上游，优先完成五项裁决后再进入实现拆分

### phase14：`phase14_standard_entity_foundation`

- 目标：交付 `Standard Entity Foundation`
- 进入条件：`phase13-12` 正式缺口记录（唯一直接上游）+ 五项裁决已完成（2026-08-17 用户结构化逐项拍板）
- 范围约束：单一主交付 = `Standard` 全局规范实体最小主线（合同 → 存储 → 后端 → 前端 → agent 消费）+ 治理画像重叠承接位完整退役；`CON-08` 时间轴、agent 写回、MCP / CLI、Git 推进跟踪、模板仓库自动接入、自动同步均为非目标
- 交付要求：作为交付型 phase 推进，按 `dev_plan` 11 个子任务（phase14-01 ~ phase14-11）顺序进入 `/spec`、实现、验收与收口
- 状态：`completed`
- 当前收口结果：已完成 `Standard` 全局规范实体最小主线（`standards` + `standard_revisions` + `standard_bindings` 三表与 jsonb 树形目录、`/standards` 前端四页、`StandardService` 8 RPC、brief `standards[]` agent 直读）+ 治理画像七触点系统性退役；期间按用户裁决完成 brief 画像残余解耦（brief 收敛为 5 顶层块，槽位 2/3/4 reserved）；`phase14-10` 为正式验收与收口入口（固定 6 问 6/6 + 八项裁决门禁全绿 + 16 页反回归矩阵 + 独立复核 PASS），`phase14-11` 为根级收口与 `phase15` 进入条件冻结入口；三件套保留为该阶段 `/plan` 的规划与冻结记录

## 4. 说明

- 本文档只承载全局开发预览、phase 计划与进度
- 不展示 task 级拆分
- task 级内容只进入对应 `docs/phase/phase*_*_dev_plan.md`
- `phase01` 是规格收敛特例；`phase02+` 默认按交付型 phase 推进
