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
- 当前已完成：
  - `.trae/specs/phase01_06_formal_mvp_spec/mvp_spec_v0.1.md`（`v0.1` 执行层唯一规格入口）
  - `.trae/specs/phase02_09_module_registry_formal_spec/module_registry_spec_v0.1.md`（`Module Registry` 当前阶段唯一规格入口）
  - `.trae/specs/phase02_12_module_registry_integration_validation_acceptance/acceptance_report.md`（`phase02` 联调验收通过结论）
  - `.trae/specs/phase03_10_decision_center_formal_spec/decision_center_spec_v0.1.md`（`Decision Center` 当前阶段唯一规格入口）
  - `.trae/specs/phase03_11_decision_center_proto_mainline/`（`Decision Center` 最小 `.proto` 合同主线入口）
  - `.trae/specs/phase03_14_decision_center_integration_validation_acceptance/`（`phase03` 联调验收与收口结论入口）
- 当前项目技术路线已冻结为 `Durable System Track`
- `phase02_module_registry_foundation` 已完成 `/plan -> /spec -> 实现 -> 验收 -> 收口`
- `phase03_decision_center_foundation` 已完成 `/plan -> /spec -> 实现 -> 验收 -> 收口`
- `phase04_product_and_repository_binding_foundation` 已正式进入 `/plan`
- `phase02` 直接承接正式 MVP 规格正文，不在根级文档重复正文内容
- `phase03` 直接承接 `phase02` 已交付的 `Module Registry` 主线与 `.proto` 合同结果，并已完成 `Decision Center` 正式规格、合同、实现与联调验收收口
- `phase04` 直接承接 `phase03` 已交付的 `Decision Center` 主线与联调验收结果，并已建立 `Product Registry + Repository Binding` 三件套
- 当前阶段入口已切换为 `phase04_product_and_repository_binding_foundation`
- `phase04` 将直接承接 `phase03-10` 正式规格、`phase03-11` 合同主线与 `phase03-14` 验收结论，不在根级文档重复正文内容

### 4.2 fix

- 当前尚未创建 `fix*` 文档

### 4.3 audit

本目录只承接内部审计工作流（`audit_issue` / `audit_analysis`），当前已建立模板：

- `audit_issue_template.md`
- `audit_analysis_template.md`

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

另保留迁移跳转文件：

- `docs/review/Personal Software Company OS v2.0.md` -> `TECH_STACK_BASELINE.md`

## 5. 迁移规则

- 已迁移文档在根目录保留受控跳转文件，避免旧引用直接失效
- 活动文档必须能从 `docs/README.md` 进入
- 新文档创建时，必须先判断它属于 `phase / fix / audit / review / archive` 哪一类
- 不再新增含义模糊、无法直接对应 workflow 的目录
