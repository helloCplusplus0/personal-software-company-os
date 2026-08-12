# architecture_map.md

# Personal Software Company OS Architecture Map

## 1. 根级文件职责

### 1.1 项目公共文档

- `AGENTS.md`：OpenAI agent 默认全局上下文入口
- `README.md`：项目总览入口
- `plan.md`：全局开发预览文档，只展示 phase 计划、目标与进度
- `TECH_STACK_BASELINE.md`：项目统一技术栈基线与长期技术约束
- `architecture_map.md`：目录结构、文档分类、迁移落点
- `PSCO-summarize-feedback.md`：当前最终共识文档

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
- `PSCO-summarize-feedback.md`

当前保留在根目录的历史参考文档：

> 说明：`AGENTS-OLD.md` 已不存在于根目录，且不再作为当前项目技术栈来源；历史技术口径统一以 `TECH_STACK_BASELINE.md` 为准。

- `PSCO-mvp01-summarize-feedback.md`：`mvp0.1` 阶段最终共识与执行基线参考
- `PSCO-mvp02-summarize-feedback.md`：基于两轮 `mvp0.2` 评审与交叉汇总形成的下一阶段最终仲裁与 `/plan` 上游基线
- `PSCO-mvp03-summarize-feedback.md`：基于两轮 `mvp0.3` 评审与交叉汇总形成的最终仲裁、范围基线与推进顺序约束，作为 `phase07` 与后续 `mvp0.3` 业务阶段正式 `/plan` 的上游判断基线

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
  - `docs/phase/phase08_operating_review_loop_foundation_architecture_plan.md`（最近完成正式业务 phase 的架构规划与冻结记录）
  - `docs/phase/phase08_operating_review_loop_foundation_dev_plan.md`（最近完成正式业务 phase 的任务与 DoD 冻结记录）
  - `docs/phase/phase08_operating_review_loop_foundation_shared_baseline.md`（最近完成正式业务 phase 的共享基线冻结记录）
  - `.trae/specs/phase08_11_validate_review_loop_integration_browser_regression_acceptance/acceptance_report.md`（`phase08` 正式验收与收口结论入口）
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
- `phase08` 三件套在收口后只承担最近完成正式业务 phase 的规划与冻结记录角色；正式验收与收口入口收敛到 `phase08-11` acceptance_report
- `phase08` 只承接 `Operating Review Loop`，不把 `Template Reuse / Derived Intelligence / dry-run` 混写为当前并列主交付
- `phase09` 已直接承接 `PSCO-mvp03-summarize-feedback.md`、`phase08-11` 验收结论、`phase08` 三件套规划记录与 `phase06` 复用摘要主线，并已建立正式 `/plan` 入口
- `phase09` 三件套承担当前支撑能力 phase 的规划与冻结记录角色；当前只承接 `Template Reuse + Derived Intelligence Deepening` 的最小支撑能力
- 后续 `dry-run` 只允许在直接承接 `phase09` 验收结论、`phase08` 根级收口结论与当前根级真相源后建立正式入口；根级文档只写进入条件，不得提前猜测任何未建立阶段名称

### 4.2 fix

- 当前尚未创建 `fix*` 文档

### 4.3 audit

本目录只承接内部审计工作流（`audit_issue` / `audit_analysis`），当前已建立模板：

- `audit_issue_template.md`
- `audit_analysis_template.md`
- `audit_001_transport_contract_mainline_issue.md`
- `audit_001_transport_contract_mainline_analysis.md`

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

另保留迁移跳转文件：

- `docs/review/Personal Software Company OS v2.0.md` -> `TECH_STACK_BASELINE.md`

## 5. 迁移规则

- 已迁移文档在根目录保留受控跳转文件，避免旧引用直接失效
- 活动文档必须能从 `docs/README.md` 进入
- 新文档创建时，必须先判断它属于 `phase / fix / audit / review / archive` 哪一类
- 不再新增含义模糊、无法直接对应 workflow 的目录
