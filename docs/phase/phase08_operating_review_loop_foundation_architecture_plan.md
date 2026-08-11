# phase08_operating_review_loop_foundation_architecture_plan

## 1. 文档定位

本文档是 `phase08_operating_review_loop_foundation` 的架构规划文档。

`phase07` 已完成 `/plan -> /spec -> 实现 -> 验收 -> 收口`，并正式交付了 `phase01 ~ phase06` canonical 业务接口向 `.proto + ConnectRPC` 主线的切换结果。当前仓库已经具备进入后续 `mvp0.3` 业务阶段的前置基础，因此需要建立第一个正式业务 phase 入口。

根据 `PSCO-mvp03-summarize-feedback.md` 已冻结的仲裁结论，`mvp0.3` 的正式方向不是把 `Operating Review Loop / Template Reuse / Derived Intelligence Deepening / Real-Project Dry-Run` 作为完全并列主线推进，而是以 **`Operating Review Loop` 为当前唯一中心主线**，其余能力作为后续支撑 phase 或独立验收闸承接。

因此，`phase08` 的职责不是一次性做完 `mvp0.3` 全部主题，而是先完成“让 Dashboard 真正成为经营动作起点，并闭合 `Feedback -> Decision -> Update`”这一单一主交付能力。

## 2. 上游输入

本阶段直接上游输入如下：

1. `AGENTS.md`
2. `plan.md`
3. `TECH_STACK_BASELINE.md`
4. `project_rules.md`
5. `architecture_map.md`
6. `docs/README.md`
7. `PSCO-mvp03-summarize-feedback.md`
8. `.trae/specs/phase07_07_formal_transport_mainline_cutover_spec/transport_mainline_cutover_spec_v0.1.md`
9. `.trae/specs/phase07_11_validate_phase01_phase06_regression_retirement_acceptance/acceptance_report.md`
10. `.trae/specs/phase05_10_dashboard_feedback_formal_spec/dashboard_feedback_spec_v0.1.md`
11. `.trae/specs/phase05_14_dashboard_feedback_integration_validation_acceptance/acceptance_report.md`
12. `.trae/specs/phase03_10_decision_center_formal_spec/decision_center_spec_v0.1.md`
13. `.trae/specs/phase04_10_product_repository_binding_formal_spec/product_repository_binding_spec_v0.1.md`
14. `.trae/specs/phase06_12_onboarding_sovereignty_reuse_formal_spec/spec.md`

补充说明：

- `PSCO-mvp03-summarize-feedback.md` 继续冻结 `mvp0.3` 的业务方向排序与范围边界，但不直接替代本阶段 `/plan`
- `phase07-07` 与 `phase07-11` 是当前阶段唯一直接执行层迁移上游，不允许重新长出第二套业务传输主线
- `phase03 / 04 / 05 / 06` 的正式规格与验收结果继续作为本阶段可消费的既有业务资产上游

## 3. 本阶段目标

`phase08` 的目标是：

> 在不扩写长期业务实体、不把系统演化成通用任务管理器的前提下，让 Dashboard 正式承担 daily / weekly operating cycle 的入口职责，并首次稳定闭合 `Feedback -> Decision -> Update` 经营动作链。

本阶段需要回答的核心问题：

1. Dashboard 如何从“总览页”切换为正式 review 入口，而不是继续增加摘要卡片
2. daily / weekly review 的差异语义、输入数据与完成定义应如何冻结，避免退化成“同页双按钮”
3. 当前焦点、代表性反馈与待处理决策应如何在 review 中形成单一动作入口
4. `Decision` 如何继续保持经营中心地位，而不是退化为 review 附属备注
5. review 结论如何低摩擦回流到既有实体，而不长出第二套长期任务系统
6. 当前真实 caller、route、query owner 与 application owner 应如何编目，避免 `/spec` 遗漏真实调用方
7. 当前 phase 的合同、读模型、写路径与前端 owner 应如何收敛，避免把 review 编排散落到 route、页面与组件中
8. 当前 phase 完成时，如何证明“用户能够从 Dashboard 进入经营动作”，而不仅是“页面多了一个 review 区块”

## 4. 架构冻结结论

### 4.1 当前阶段唯一直接执行层上游

`phase08` 必须直接承接：

- `PSCO-mvp03-summarize-feedback.md`
- `.trae/specs/phase07_07_formal_transport_mainline_cutover_spec/transport_mainline_cutover_spec_v0.1.md`
- `.trae/specs/phase07_11_validate_phase01_phase06_regression_retirement_acceptance/acceptance_report.md`
- `.trae/specs/phase05_10_dashboard_feedback_formal_spec/dashboard_feedback_spec_v0.1.md`
- `.trae/specs/phase05_14_dashboard_feedback_integration_validation_acceptance/acceptance_report.md`
- `.trae/specs/phase03_10_decision_center_formal_spec/decision_center_spec_v0.1.md`
- `.trae/specs/phase04_10_product_repository_binding_formal_spec/product_repository_binding_spec_v0.1.md`
- `.trae/specs/phase06_12_onboarding_sovereignty_reuse_formal_spec/spec.md`

不允许在本阶段重新解释：

- `Decision` 在 MVP 中的中心地位
- `.proto` 作为唯一长期合同源
- `phase07` 已完成的 ConnectRPC 业务主线切换结论
- `mvp0.3` 的方向排序与“单一中心主线 + 后续支撑能力 + 独立验收闸”的结构

### 4.2 当前阶段主交付对象

`phase08` 的主交付对象是：

- `Operating Review Loop Entry`
- `Dashboard Review Flow`
- `Feedback -> Decision -> Update Action Chain`
- `Review Result Writeback`

其最小主线必须优先承接：

- `Dashboard Home`
- `Feedback Signals`
- `Decision Center`
- `Product / Module / Repository` 既有详情页与更新动作

当前阶段不把以下内容作为主交付对象：

- `Template Reuse`
- `Derived Intelligence Deepening`
- `Real-Project Dry-Run`
- 新的长期核心实体
- 通用任务系统 / backlog 系统

### 4.3 当前阶段前端交付策略

前端继续统一遵守：

- 单一 `React Web`
- `TanStack Router + TanStack Query`
- 业务写路径唯一 `application` 入口
- `query` 层纯只读
- mutation 固定承接位

当前阶段重点是：

- 让 Dashboard 正式承接 review 入口与经营动作编排
- 把 review 相关读取、回流与错误处理收敛到固定切片 owner
- 保持页面层只消费已经收敛的 read / application 承接位
- 不在 route、卡片组件或弹层中长期内联第二套 review 编排语义

### 4.4 当前阶段后端、合同与数据承接策略

当前阶段继续统一遵守：

- `.proto` 是唯一长期合同源
- `ConnectRPC` 是业务接口正式传输层
- `chi` 继续只承接 router shell、middleware 与非业务端点
- 既有 `phase07` 业务接口主线不回退为手写 JSON contract

当前阶段重点是：

- 为 review 入口、review 动作承接与结论回流补齐最小正式合同
- 保持 `Feedback / Decision / Product / Module / Repository` 的事实主线继续由既有领域实体承接
- 若需要新增 review 记录或派生读模型，必须保持轻量，不得升级为新的长期核心业务主线
- review 结论应回流既有实体，而不是复制一套并列状态体系

### 4.5 当前阶段业务边界原则

为了避免 `phase08` 演化为“经营 review + 模板复用 + 指标系统 + 真实 dry-run”大杂烩，当前阶段先冻结以下边界：

- 本阶段只交付 `Operating Review Loop`
- `Template Reuse` 与 `Derived Intelligence Deepening` 只允许作为当前 `/plan` 的后续依赖表达，不进入本阶段实现承诺
- `Real-Project Dry-Run` 只允许作为后续独立验收闸，不在本阶段混写为实现子任务
- review loop 不能演化为通用任务管理器；动作必须围绕 `Feedback / Decision / Update` 闭环
- `Decision` 必须保持经营中心地位，不允许被弱化为 review 辅助备注

### 4.6 当前阶段与后续阶段的依赖关系

`phase08` 是 `mvp0.3` 的第一个正式业务 phase，但不是全部。

后续依赖关系冻结如下：

1. 没有 `phase08`，Dashboard 仍只是总览页，`Decision` 也还不是经营枢纽
2. 只有在 `phase08` 先闭合 `Feedback -> Decision -> Update` 后，`Template Reuse` 与 `Derived Intelligence Deepening` 才有稳定消费场景
3. `Real-Project Dry-Run` 必须等待 review loop 与支撑能力进入可运行态后，再作为独立验收闸推进
4. 后续支撑能力 phase 与 dry-run phase 的正式命名，留给下一次 `/plan` 建立时再冻结

## 5. 当前阶段完成条件

`phase08` 完成时，至少必须满足：

1. Dashboard 已成为正式 daily / weekly review 入口
2. 当前焦点、代表性反馈与待处理决策已形成统一动作承接位
3. `Feedback -> Decision -> Update` 已形成可重复执行的最小闭环
4. review 结论能够回流到既有实体，而不是停留在页面展示层
5. 前端与后端的 review 相关合同、读取与写路径已收敛到单值 owner
6. 本阶段未把 `Template Reuse / Derived Intelligence / Real-Project Dry-Run` 偷渡为并列主交付
7. 根级入口已回写当前正式业务 phase 为 `phase08_operating_review_loop_foundation`
