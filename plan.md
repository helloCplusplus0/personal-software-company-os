# plan.md

# Personal Software Company OS Plan

## 1. 当前状态

- 当前阶段：`phase07_transport_contract_mainline_migration`
- 当前状态：`phase07` 已进入实现与验收推进段；`phase07-08` 生成链主线与 `phase07-09` Go 后端传输主线切换已完成，当前待继续推进 `phase07-10 ~ phase07-12`
- 当前目标：完成 `phase01 ~ phase06` canonical 业务接口向 `.proto + ConnectRPC` 正式主线的一次性迁移，并为后续 `mvp0.3` 业务主线清空传输层历史负担
- 当前下一阶段入口：待 `phase07` 正式收口后，切换到 `mvp0.3` 业务主线对应的正式 phase 入口（不得预设 `phase08` 名称）

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
- `.trae/specs/phase06_12_onboarding_sovereignty_reuse_formal_spec/spec.md` 已冻结为 `phase06` 当前阶段唯一规格收敛入口
- `.trae/specs/phase06_13_land_minimal_proto_contract_mainline/` 已将 `phase06` 最小 `.proto` 合同落地为仓库主线
- `.trae/specs/phase06_16_integration_validation_acceptance/acceptance_report.md` 已形成 `phase06` 联调验收与收口结论入口
- `phase06` 三件套保留为该阶段 `/plan` 的规划与冻结记录，不再覆盖根级当前状态
- `PSCO-mvp03-summarize-feedback.md` 已冻结为 `phase07` 与后续 `mvp0.3` 业务阶段的最终仲裁与完整路线规划基线
- `docs/audit/audit_001_transport_contract_mainline_issue.md` 与 `docs/audit/audit_001_transport_contract_mainline_analysis.md` 已冻结为当前传输主线收敛议题的正式审计入口
- `phase07_transport_contract_mainline_migration` 已建立正式 `/plan` 入口，当前作为 `mvp0.3` 业务阶段之前的前置基础 phase 推进
- `phase07` 的目标不是新增业务能力，而是完成 `phase01 ~ phase06` canonical 业务接口向 `.proto + ConnectRPC` 的一次性正式迁移

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
- 状态：`current`
- 当前阶段结果：`phase07-01 ~ phase07-09` 已完成冻结、规格、生成链与 Go 后端主线切换；下一步继续推进 `phase07-10` 前端调用切换、`phase07-11` 联调回归验收与 `phase07-12` 根级收口

### `mvp0.3` 业务候选阶段结构（承接 `phase07` 之后）

> 说明：以下内容用于一次性表达 `phase07` 收口之后的完整路线预览，来源于 `PSCO-mvp03-summarize-feedback.md`。
> 它们是后续正式 `/plan` 应承接的候选阶段结构，不等于已建立的正式 `phase08 / phase09 / phase10`，不得在根级文档中写成既成事实。

### 候选阶段一：Operating Review Loop 主线建立

- 目标：让 Dashboard 真正接管经营动作起点，并首次闭合 `Feedback -> Decision -> Update`
- 直接上游：`PSCO-mvp03-summarize-feedback.md` 与 `phase07` 已冻结的传输主线切换结果
- 范围约束：不得把 review loop 演化为通用任务管理器；必须以既有 `Decision / Product / Module / Feedback / Reuse` 为事实基础
- 至少应承接：
  - daily / weekly review 入口
  - 当前焦点、代表性反馈与待处理决策的统一承接
  - review 结论记录
  - action handoff 回流既有实体
- 目标验收：证明用户能够从 Dashboard 进入经营动作，且 `Decision` 真正成为 review 中的经营中心

### 候选阶段二：Template Reuse + Derived Intelligence Deepening

- 目标：让复用从“可见”推进到“可用、可行动”
- 直接上游：候选阶段一已建立的 review loop 主线，以及 `phase06` 已交付的 `module_reuse_summary / capability_summary`
- 范围约束：模板边界继续冻结为“Module 组合快照 + 新建预填辅助”；不得演化为完整模板平台，不得把派生智能扩成 AI 主线
- 至少应承接：
  - `Module` 组合快照
  - 新建 `Product` 时基于组合快照预填
  - capability gap / reuse opportunity 提示
  - 最小解释性指标落地
- 目标验收：证明 review 后的下一步创造成本被真实降低，且复用感知不再停留于展示层

### 候选阶段三：Real-Project Dry-Run 与收口验证

- 目标：用真实项目证明 `mvp0.3` 核心命题成立，并形成下一阶段仲裁输入
- 直接上游：候选阶段一与候选阶段二已形成可运行闭环
- 范围约束：不得用纯 fixture 结论替代真实项目证据；必须保留独立验收记录，而不是混写进普通联调说明
- 至少应承接：
  - 至少一个真实项目完整走通
  - 明确记录摩擦点、收益点与修正项
  - 与 fixture 验收并列留档
- 目标验收：证明用户会围绕真实项目持续 review、持续决策、持续复用，并形成下一次创造加速

### 后续候选方向：`mvp0.4+`

- `Venture`
- `Decision Intelligence`
- `AI Context Enhancement`

说明：以上方向已在 `PSCO-mvp03-summarize-feedback.md` 中明确后移，当前仅保留为后续候选范围，不进入当前后续 phase 主线规划

## 4. 下一阶段切换条件

当以下条件同时满足时，根级入口才允许从 `phase07` 当前进行中状态切换到下一阶段正式业务 phase 入口：

1. `phase07` 已完成 `/plan -> /spec -> 实现 -> 验收 -> 收口` 全链路
2. `phase07` 对应正式规格正文已冻结为当前阶段唯一正式规格入口
3. `phase07` 对应 `buf` 生成链、Go Connect 产物与前端客户端主线已稳定落地
4. `phase07` 对应联调回归验收与旧 JSON 业务主线退场结论已形成正式入口
5. 下一阶段正式 `phase` 入口文档已在仓库中建立，并明确直接承接 `phase07` 已冻结的规格、迁移结果与验收结论
6. 根级真相源已完成状态切换回写，且未在根级文档中凭空猜测未建立的 phase 名称

当前结论：`phase07` 已建立正式 `/plan` 入口，当前直接上游已冻结为 `PSCO-mvp03-summarize-feedback.md`、`audit_001` 审计结论与 `phase06-12 / 13 / 16` 三个正式入口；在 `phase07` 正式收口前，后续 `mvp0.3` 业务阶段只保留为候选路线预览，不预设 `phase08` 名称。

## 5. 说明

- 本文档只承载全局开发预览、phase 计划与进度
- 不展示 task 级拆分
- task 级内容只进入对应 `docs/phase/phase*_*_dev_plan.md`
- `phase01` 是规格收敛特例；`phase02+` 默认按交付型 phase 推进
