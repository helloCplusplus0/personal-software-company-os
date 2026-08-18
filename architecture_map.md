# architecture_map.md

# Personal Software Company OS Architecture Map

## 1. 根级文件职责

### 1.1 项目公共文档

- `AGENTS.md`：OpenAI agent 默认全局上下文入口
- `README.md`：项目总览入口
- `plan.md`：全局开发预览文档，只展示 phase 计划、目标与进度
- `TECH_STACK_BASELINE.md`：项目统一技术栈基线与长期技术约束
- `architecture_map.md`：目录结构、文档分类、迁移落点
- `PSCO-mvp05-summarize-feedback.md`：当前最终共识文档

### 1.2 Trae 内部 agent 上下文文档

- `project_rules.md`：技术基线、workflow、协作门禁
- `project_skills.md`：项目专属语义、边界与风险提醒
- `global_skills.md`：通用执行方法在本项目中的映射

## 2. 根目录保留策略

根目录只保留以下四类文件：

1. 项目公共入口文档
2. agent 上下文文档
3. 当前仍作为主入口的基础方案文档
4. 迁移后的受控跳转文件

当前继续保留在根目录的活动文档：

- `AGENTS.md`
- `README.md`
- `plan.md`
- `TECH_STACK_BASELINE.md`
- `architecture_map.md`
- `project_rules.md`
- `project_skills.md`
- `global_skills.md`
- `PSCO_0.md ~ PSCO_4.md`
- `PSCO-mvp05-summarize-feedback.md`

当前保留在根目录的历史参考文档：

> 说明：`AGENTS-OLD.md` 已不存在于根目录，且不再作为当前项目技术栈来源；历史技术口径统一以 `TECH_STACK_BASELINE.md` 为准。

- `PSCO-mvp01-summarize-feedback.md`：`mvp0.1` 阶段最终共识与执行基线参考
- `PSCO-mvp02-summarize-feedback.md`：基于两轮 `mvp0.2` 评审与交叉汇总形成的下一阶段最终仲裁与 `/plan` 上游基线
- `PSCO-mvp03-summarize-feedback.md`：基于两轮 `mvp0.3` 评审与交叉汇总形成的最终仲裁、范围基线与推进顺序约束，作为 `phase07` 与后续 `mvp0.3` 业务阶段正式 `/plan` 的上游判断基线
- `PSCO-mvp04-summarize-feedback.md`：基于 `mvp0.4` 评审与交叉汇总形成的最终仲裁与 `phase10` 的直接规划基线
- `docs/review/PSCO-real-project-dry-run-user-manual-GPT54.md`：`phase08` 与 `phase09` 收口后，为后续 `Real-Project Dry-Run` 准备的标准使用手册与真实使用前参考输入；不承担正式 `phase` 规则冻结职责

## 3. docs 目录结构

```text
docs/
├── README.md      # workflow 文档总入口
├── phase/         # 项目推进主线
├── fix/           # bug 修复与局部问题
├── audit/         # 跨模块复核、路线仲裁、结构审计（内部审计工作流）
├── review/        # 专家评审与交叉汇总文档（历史留档）
└── archive/       # 沉默文档、旧规范、历史资料
```

## 4. 当前文档落点

### 4.1 phase

- 当前已创建：
  - `phase01_mvp_spec_convergence_architecture_plan.md`
  - `phase01_mvp_spec_convergence_dev_plan.md`
  - `phase01_mvp_spec_convergence_shared_baseline.md`
- 当前已创建：
  - `phase02_module_registry_foundation_architecture_plan.md`
  - `phase02_module_registry_foundation_dev_plan.md`
  - `phase02_module_registry_foundation_shared_baseline.md`
- 当前已创建：
  - `phase03_decision_center_foundation_architecture_plan.md`
  - `phase03_decision_center_foundation_dev_plan.md`
  - `phase03_decision_center_foundation_shared_baseline.md`
- 当前已创建：
  - `phase04_product_and_repository_binding_foundation_architecture_plan.md`
  - `phase04_product_and_repository_binding_foundation_dev_plan.md`
  - `phase04_product_and_repository_binding_foundation_shared_baseline.md`
- 当前已创建：
  - `phase05_dashboard_feedback_foundation_architecture_plan.md`
  - `phase05_dashboard_feedback_foundation_dev_plan.md`
  - `phase05_dashboard_feedback_foundation_shared_baseline.md`
- 当前已创建：
  - `phase06_onboarding_sovereignty_reuse_foundation_architecture_plan.md`
  - `phase06_onboarding_sovereignty_reuse_foundation_dev_plan.md`
  - `phase06_onboarding_sovereignty_reuse_foundation_shared_baseline.md`
- 当前已创建：
  - `phase07_transport_contract_mainline_migration_architecture_plan.md`
  - `phase07_transport_contract_mainline_migration_dev_plan.md`
  - `phase07_transport_contract_mainline_migration_shared_baseline.md`
- 当前已创建：
  - `phase10_asset_action_closure_foundation_architecture_plan.md`
  - `phase10_asset_action_closure_foundation_dev_plan.md`
  - `phase10_asset_action_closure_foundation_shared_baseline.md`
- 当前已创建：
  - `phase11_project_context_foundation_architecture_plan.md`
  - `phase11_project_context_foundation_dev_plan.md`
  - `phase11_project_context_foundation_shared_baseline.md`
- 当前已创建：
  - `phase12_semantic_alignment_and_readonly_consumption_foundation_architecture_plan.md`
  - `phase12_semantic_alignment_and_readonly_consumption_foundation_dev_plan.md`
  - `phase12_semantic_alignment_and_readonly_consumption_foundation_shared_baseline.md`
- 当前已创建：
  - `phase13_project_governance_profile_foundation_architecture_plan.md`
  - `phase13_project_governance_profile_foundation_dev_plan.md`
  - `phase13_project_governance_profile_foundation_shared_baseline.md`
  - `phase14_standard_entity_foundation_architecture_plan.md`
  - `phase14_standard_entity_foundation_dev_plan.md`
  - `phase14_standard_entity_foundation_shared_baseline.md`
- 当前已创建：
  - `phase09_template_reuse_derived_intelligence_foundation_architecture_plan.md`
  - `phase09_template_reuse_derived_intelligence_foundation_dev_plan.md`
  - `phase09_template_reuse_derived_intelligence_foundation_shared_baseline.md`
- 当前已完成：
  - `.trae/specs/phase01_06_formal_mvp_spec/mvp_spec_v0.1.md`（`v0.1` 执行层唯一规格入口）
  - `.trae/specs/phase02_09_module_registry_formal_spec/module_registry_spec_v0.1.md`（`Module Registry` 当前阶段唯一规格入口）
  - `.trae/specs/phase02_12_module_registry_integration_validation_acceptance/acceptance_report.md`（`phase02` 联调验收通过结论）
  - `.trae/specs/phase03_10_decision_center_formal_spec/decision_center_spec_v0.1.md`（`Decision Center` 当前阶段唯一规格入口）
  - `.trae/specs/phase03_11_decision_center_proto_mainline/`（`Decision Center` 最小 `.proto` 合同主线入口）
  - `.trae/specs/phase03_14_decision_center_integration_validation_acceptance/`（`phase03` 联调验收与收口结论入口）
  - `.trae/specs/phase04_10_product_repository_binding_formal_spec/product_repository_binding_spec_v0.1.md`（`phase04` 正式规格收口入口，同时作为 `phase05` 直接上游）
  - `.trae/specs/phase04_11_product_repository_binding_proto_mainline/`（`phase04` 最小 `.proto` 合同收口入口，同时作为 `phase05` 直接上游）
  - `.trae/specs/phase04_14_product_repository_binding_integration_validation_acceptance/acceptance_report.md`（`phase04` 联调验收与收口结论入口，同时作为 `phase05` 直接上游）
  - `.trae/specs/phase05_10_dashboard_feedback_formal_spec/dashboard_feedback_spec_v0.1.md`（`phase05` 正式规格收口入口，同时作为下一阶段直接上游）
  - `.trae/specs/phase05_11_dashboard_feedback_proto_mainline/`（`phase05` 最小 `.proto` 合同收口入口，同时作为下一阶段直接上游）
  - `.trae/specs/phase05_14_dashboard_feedback_integration_validation_acceptance/acceptance_report.md`（`phase05` 联调验收与收口结论入口，同时作为下一阶段直接上游）
  - `.trae/specs/phase06_12_onboarding_sovereignty_reuse_formal_spec/spec.md`（`phase06` 正式规格收口入口）
  - `.trae/specs/phase06_13_land_minimal_proto_contract_mainline/`（`phase06` 最小 `.proto` 合同收口入口）
  - `.trae/specs/phase06_16_integration_validation_acceptance/acceptance_report.md`（`phase06` 联调验收与收口结论入口）
  - `.trae/specs/phase07_07_formal_transport_mainline_cutover_spec/transport_mainline_cutover_spec_v0.1.md`（`phase07` 正式规格收口入口）
  - `.trae/specs/phase07_11_validate_phase01_phase06_regression_retirement_acceptance/acceptance_report.md`（`phase07` 联调、退场验收与收口结论入口）
  - `docs/phase/phase08_operating_review_loop_foundation_architecture_plan.md`（已完成正式业务 phase 的架构规划与冻结记录）
  - `docs/phase/phase08_operating_review_loop_foundation_dev_plan.md`（已完成正式业务 phase 的任务与 DoD 冻结记录）
  - `docs/phase/phase08_operating_review_loop_foundation_shared_baseline.md`（已完成正式业务 phase 的共享基线冻结记录）
  - `.trae/specs/phase08_11_validate_review_loop_integration_browser_regression_acceptance/acceptance_report.md`（`phase08` 正式验收与收口结论入口）
  - `docs/phase/phase09_template_reuse_derived_intelligence_foundation_architecture_plan.md`（最近完成正式支撑能力 phase 的架构规划与冻结记录）
  - `docs/phase/phase09_template_reuse_derived_intelligence_foundation_dev_plan.md`（最近完成正式支撑能力 phase 的任务与 DoD 冻结记录）
  - `docs/phase/phase09_template_reuse_derived_intelligence_foundation_shared_baseline.md`（最近完成正式支撑能力 phase 的共享基线冻结记录）
  - `.trae/specs/phase09_11_validate_template_reuse_derived_hint_integration_browser_regression_acceptance/acceptance_report.md`（`phase09` 正式验收与收口结论入口）
  - `docs/phase/phase10_asset_action_closure_foundation_architecture_plan.md`（历史完成正式业务 phase 的架构规划与冻结记录）
  - `docs/phase/phase10_asset_action_closure_foundation_dev_plan.md`（历史完成正式业务 phase 的任务与 DoD 冻结记录）
  - `docs/phase/phase10_asset_action_closure_foundation_shared_baseline.md`（历史完成正式业务 phase 的共享基线冻结记录）
  - `.trae/specs/phase10_11_complete_asset_action_closure_integration_browser_regression_validation/spec.md`（`phase10-11` 正式验收入口；同目录 `tasks.md / checklist.md` 共同构成 `phase10` 正式验收与收口证据）
  - `.trae/specs/phase12_11_validate_semantic_alignment_readonly_consumption_foundation/acceptance_report.md`（`phase12` 正式验收与收口入口）
  - `.trae/specs/phase13_11_validate_project_governance_profile_integration_dogfooding_regression/acceptance_report.md`（`phase13` 正式验收与收口入口）
  - `.trae/specs/phase13_12_sync_root_level_closeout_next_phase_entry_conditions/`（`phase13` 正式缺口记录与 `phase14` 进入条件冻结入口）
  - `.trae/specs/phase14_10_validate_standard_entity_integration_dogfooding_regression/`（`phase14` 正式验收与收口入口）
  - `.trae/specs/phase14_11_sync_root_level_closeout_freeze_phase15_entry_conditions/`（`phase15` 进入条件冻结入口）
- 当前项目技术路线已冻结为 `Durable System Track`
- `phase02_module_registry_foundation` 已完成 `/plan -> /spec -> 实现 -> 验收 -> 收口`
- `phase03_decision_center_foundation` 已完成 `/plan -> /spec -> 实现 -> 验收 -> 收口`
- `phase04_product_and_repository_binding_foundation` 已完成 `/plan -> /spec -> 实现 -> 验收 -> 收口`
- `phase05_dashboard_feedback_foundation` 已完成 `/plan -> /spec -> 实现 -> 验收 -> 收口`
- `phase04` 三件套保留为该阶段 `/plan` 的规划与冻结记录，不再承担项目当前阶段状态说明
- `phase02` 直接承接正式 MVP 规格正文，不在根级文档重复正文内容
- `phase03` 直接承接 `phase02` 已交付的 `Module Registry` 主线与 `.proto` 合同结果，并已完成 `Decision Center` 正式规格、合同、实现与联调验收收口
- `phase04` 直接承接 `phase03` 已交付的 `Decision Center` 主线与联调验收结果，并已完成 `Product Registry + Repository Binding` 正式规格、合同、实现与联调验收收口
- `phase05` 已直接承接 `phase04-10` 正式规格、`phase04-11` 合同主线与 `phase04-14` 验收结论，并完成 `Dashboard + Feedback` 正式规格、合同、实现与联调验收收口
- `phase05` 三件套保留为该阶段 `/plan` 的规划与冻结记录，不再承担根级当前阶段状态说明
- `phase06` 已直接承接 `PSCO-mvp02-summarize-feedback.md` 与 `phase05-10 / 11 / 14`，并已完成正式规格、合同、实现与联调验收收口
- `phase06` 三件套保留为该阶段 `/plan` 的规划与冻结记录，不再承担根级当前阶段状态说明
- `phase07` 已直接承接 `PSCO-mvp03-summarize-feedback.md`、`audit_001` 审计结论与 `phase06-12 / 13 / 16`，并已完成正式规格、实现、验收与根级收口
- `phase07` 三件套在收口后只承担最近完成前置基础阶段的规划与冻结记录角色；执行层正式入口只收敛到 `phase07-07` 正式规格与 `phase07-11` 验收结论，不再覆盖当前业务状态
- `phase07-08 / 09 / 10` 只保留为生成链、后端与前端迁移实现结论记录，不提升为根级长期主入口
- `phase08` 已直接承接 `PSCO-mvp03-summarize-feedback.md`、`phase07-07` 正式规格与 `phase07-11` 验收结论，并已完成 `Operating Review Loop` 的正式交付、统一验收与根级收口
- `phase08` 三件套在收口后只承担已完成正式业务 phase 的规划与冻结记录角色；正式验收与收口入口继续收敛到 `phase08-11` acceptance_report
- `phase08` 只承接 `Operating Review Loop`，不把 `Template Reuse / Derived Intelligence / dry-run` 混写为当前并列主交付
- `phase09` 已直接承接 `PSCO-mvp03-summarize-feedback.md`、`phase08-11` 验收结论、`phase08` 三件套规划记录与 `phase06` 复用摘要主线，并已完成 `Template Reuse + Derived Intelligence Deepening` 的正式交付、统一验收与收口
- `phase09` 三件套在收口后只承担最近完成正式支撑能力 phase 的规划与冻结记录角色；正式验收与收口入口收敛到 `phase09-11` acceptance_report
- `phase09-08 / 09 / 10` 只保留为合同落地、模板预填回流与派生提示实现结论记录，不提升为根级长期主入口
- `phase10` 已直接承接 `PSCO-mvp04-summarize-feedback.md`、`fix_001_003` 收口记录、`fix_001 ~ fix_003` analysis 与 `phase06 / phase08 / phase09` 正式交付结论，并已完成 `Asset-Action Closure` 的正式 `/plan -> /spec -> 实现 -> 验收 -> 收口`
- `phase10` 三件套在收口后只承担历史完成正式业务 phase 的规划与冻结记录角色；正式验收与收口入口收敛到 `phase10-11` 规格目录
- `phase11` 已直接承接 `PSCO-mvp05-summarize-feedback.md` 与 `phase10` 已完成的正式交付结论，并已完成“根级上下文真相源治理 + 最小只读项目上下文导出”的正式 `/plan -> /spec -> 实现 -> 验收 -> 收口`
- `phase11` 三件套在收口后只承担历史完成阶段的规划与冻结记录角色；正式验收与收口入口收敛到 `phase11-09` acceptance_report
- `phase12` 已直接承接 `PSCO-mvp05-summarize-feedback.md`、`audit_002` 审计结论与 `phase11` 已完成的正式交付结论，并已完成“前端四实体语义一致性收口 + Web / agent 共享只读消费深化”的正式 `/plan -> /spec -> 实现 -> 验收 -> 收口`
- `phase12` 三件套在收口后只承担历史完成阶段的规划与冻结记录角色；正式验收与收口入口收敛到 `phase12-11` acceptance_report
- `phase13` 已直接承接 `phase12` 已完成的正式交付结论与当前根级真相源文档，并已完成“项目级治理画像 + 全局规范资产 + agent 项目简报输入”的正式 `/plan -> /spec -> 实现 -> 验收 -> 收口`
- `phase13` 三件套在收口后只承担该阶段 `/plan` 的规划与冻结记录角色；正式验收与收口入口收敛到 `phase13-11` acceptance_report，`phase13-12` 缺口记录已随 `phase14` 收口退位为历史输入
- `phase14_standard_entity_foundation` 已完成正式 `/plan -> /spec -> 实现 -> 验收 -> 收口`，三件套在收口后只承担该阶段 `/plan` 的规划与冻结记录角色；正式验收与收口入口收敛到 `phase14-10` acceptance_report，`phase14-11` 为 `phase15` 进入条件冻结入口

### 4.2 fix

- 当前已创建：
  - `docs/fix/fix_001_onboarding_cold_start_state_issue.md`
  - `docs/fix/fix_001_onboarding_cold_start_state_analysis.md`
  - `docs/fix/fix_002_decision_pending_signal_semantics_issue.md`
  - `docs/fix/fix_002_decision_pending_signal_semantics_analysis.md`
  - `docs/fix/fix_003_decision_detail_status_advance_issue.md`
  - `docs/fix/fix_003_decision_detail_status_advance_analysis.md`
  - `docs/fix/fix_001_003_mvp03_dry_run_rerun_closure.md`
- 当前 fix 主线含义：用于承接 `Real-Project Dry-Run` 暴露的第一轮阻断项，先完成 issue 与 analysis，再进入 `/spec`、实现、聚焦 rerun 与 `mvp0.3` 收口
- 当前 fix 状态：`fix_001 ~ fix_003` 已完成修复、独立复核、聚焦 rerun 与 `mvp0.3` 收口；后续正式工作应转向 `PSCO-mvp04-summarize-feedback.md` 所定义的“候选阶段二：Asset-Action Closure 主线”

### 4.3 audit

本目录只承接内部审计工作流（`audit_issue` / `audit_analysis`），当前已建立模板：

- `audit_issue_template.md`
- `audit_analysis_template.md`
- `audit_001_transport_contract_mainline_issue.md`
- `audit_001_transport_contract_mainline_analysis.md`
- `audit_002_phase11_post_closeout_direction_issue.md`
- `audit_002_phase11_post_closeout_direction_analysis.md`

> 专家评审与交叉汇总文档不归入 `docs/audit/`，统一归入 `docs/review/`。

### 4.4 archive

以下不直接服务当前 workflow 的文档，已下沉到 `docs/archive/`：

- `PSCO_Glossary.md`
- `PSCO_v0.1_entity_boundary.md`
- `agent_workflow_notes.md`
- `agent_session_checklist.md`

### 4.5 review

以下专家评审与交叉汇总文档统一归入 `docs/review/`：

- `PSCO_Evaluation-GPT54.md`
- `PSCO_Review_deepseek-v4-flash.md`
- `PSCO-Design-Review-GLM-52.md`
- `PSCO-evaluation-deepseek-v4-pro.md`
- `PSCO-review-qwen37-pro.md`
- `PSCO-summarize-feedback-dsv4flash.md`
- `PSCO-summarize-feedback-GLM52.md`
- `PSCO-summarize-feedback-GPT54.md`
- `PSCO-mvp02-GLM52.md`
- `PSCO-next-phase-mvp02-GPT54.md`
- `PSCO-mvp02-deepseekv4flash.md`
- `PSCO-mvp02-deepseekv4pro.md`
- `PSCO-mvp02-qwen37pro.md`
- `PSCO-mvp02-summarize-feedback-GPT54.md`
- `PSCO-mvp03-GPT54.md`
- `PSCO-mvp03-summarize-feedback-GPT54.md`
- `PSCO-mvp03-DPv4flash.md`
- `PSCO-mvp03-DPv4pro.md`
- `PSCO-mvp03-GLM52.md`
- `PSCO-mvp03-qwen37pro.md`
- `PSCO-mvp03-summarize-feedback-DPv4flash.md`
- `PSCO-mvp03-summarize-feedback-GLM52.md`
- `PSCO-real-project-dry-run-user-manual-GPT54.md`
- `PSCO-real-project-dry-run-user-manual-GPT54 feedback.md`
- `PSCO-mvp04-DPv4pro.md`
- `PSCO-mvp04-summarize-feedback-DPv4pro.md`
- `PSCO-mvp045-GPT54.md`
- `PSCO-mvp045-DPv4pro.md`
- `PSCO-mvp045-gemini31pro.md`
- `PSCO-mvp045-GLM52.md`
- `PSCO-mvp045-qwen38max.md`

另保留迁移跳转文件：

- `docs/review/Personal Software Company OS v2.0.md` -> `TECH_STACK_BASELINE.md`

## 5. 迁移规则

- 已迁移文档在根目录保留受控跳转文件，避免旧引用直接失效
- 活动文档必须能从 `docs/README.md` 进入
- 新文档创建时，必须先判断它属于 `phase / fix / audit / review / archive` 哪一类
- 不再新增含义模糊、无法直接对应 workflow 的目录
