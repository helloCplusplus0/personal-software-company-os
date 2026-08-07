# AGENTS.md

> 职责：PSCO 项目的全局上下文入口。
> 这是 OpenAI agent 默认读取的项目级入口文档，目标是让接手者快速恢复正确上下文，而不是在这里展开重复规则或阶段细节。

## 1. 项目定位

- 项目名称：`Personal Software Company OS`
- 当前阶段：`phase04_product_and_repository_binding_foundation`
- 当前主目标：以 `phase03-10` 正式规格、`phase03-11` 合同主线与 `phase03-14` 验收结论为直接上游，完成 `Product / Repository / Module Binding` 最小主线的 `phase04` 三件套规划
- 当前下一阶段入口：`phase04_product_and_repository_binding_foundation`
- 当前定位：PSCO 是个人软件公司的经营与资产系统，不是代码管理工具，不是 AI Chat 产品，也不是自动扫描系统

## 2. 当前唯一上游

- 最终共识只以 `PSCO-summarize-feedback.md` 为准
- 全局推进预览只以 `plan.md` 为准
- 技术栈标准只以 `TECH_STACK_BASELINE.md` 为准
- workflow、协作门禁只以 `project_rules.md` 为准
- 目录结构、文档分类和迁移落点只以 `architecture_map.md` 为准

## 3. 当前关键共识

- `v0.1` 的正式目标是：软件资产登记、决策留痕与基础复用反馈
- `Decision` 必须进入 MVP
- `Capability` 在 `v0.1` 中只作为派生层
- `Venture` 保留，但作为可选实体，不强制创建
- `Feature / Opportunity / Experiment` 保留在长期理论模型中，但不进入 `v0.1` 主执行范围
- 当前项目必须遵守统一技术栈方案，禁止越出 `TECH_STACK_BASELINE.md` 自由发挥
- 当前项目技术路线已明确冻结为 `Durable System Track`
- `v0.1` 前端正式交付物为单一 `React Web`，同时考虑 `PC` 与移动浏览器 UI；当前不引入独立 `React Native` 客户端，`PWA` 仅作可兼容增强方向
- `docs/` 当前只服务 `phase / fix / audit / review / archive` workflow

## 4. 当前状态

- 原始方案文档 `PSCO_0.md ~ PSCO_4.md` 已完成第一轮共识回正
- 专家评审与交叉汇总文档已归类到 `docs/review/`
- 根目录保留真相源与主入口，不再作为散装文档堆放区
- `phase01_*` 三件套已完成规划收口，并作为 `phase01-06` 正式 MVP 规格正文的上游
- 执行层唯一规格入口已冻结为 `.trae/specs/phase01_06_formal_mvp_spec/mvp_spec_v0.1.md`
- `phase02_module_registry_foundation` 已完成 `/plan -> /spec -> 实现 -> 验收 -> 收口`
- `Module Registry` 最小主线已形成可运行交付物，并完成前后端、数据、Proto 合同与联调验收
- `phase02-09` 的 `module_registry_spec_v0.1.md` 已冻结为 `Module Registry` 当前阶段唯一规格收敛入口
- `phase02-11A` 已将 `Protocol Buffers` 落地为当前阶段最小 `.proto` 合同源
- `phase03_decision_center_foundation` 已完成 `/plan -> /spec -> 实现 -> 验收 -> 收口`
- `Decision Center` 最小主线已形成可运行交付物，并完成正式规格、`.proto` 合同、前后端实现与联调验收收口
- `.trae/specs/phase03_10_decision_center_formal_spec/decision_center_spec_v0.1.md` 已冻结为 `Decision Center` 当前阶段唯一规格收敛入口
- `.trae/specs/phase03_11_decision_center_proto_mainline/` 已将 `Decision Center` 最小 `.proto` 合同接入仓库既有 `proto/` 主线
- `.trae/specs/phase03_14_decision_center_integration_validation_acceptance/` 已给出 `phase03` 联调验收与收口结论入口
- 当前阶段入口已切换为 `phase04_product_and_repository_binding_foundation`
- `phase04_product_and_repository_binding_foundation` 已正式进入 `/plan`
- `phase04` 三件套已建立，用于冻结 `Product Registry + Repository Binding` 的最小主线边界、执行顺序与共享基线

## 5. 推荐阅读顺序

1. `AGENTS.md`
2. `plan.md`
3. `TECH_STACK_BASELINE.md`
4. `project_rules.md`
5. `architecture_map.md`
6. `PSCO-summarize-feedback.md`
7. `docs/README.md`
8. `.trae/specs/phase01_06_formal_mvp_spec/mvp_spec_v0.1.md`
9. `.trae/specs/phase02_09_module_registry_formal_spec/module_registry_spec_v0.1.md`
10. `.trae/specs/phase02_12_module_registry_integration_validation_acceptance/acceptance_report.md`
11. `.trae/specs/phase03_10_decision_center_formal_spec/decision_center_spec_v0.1.md`
12. `.trae/specs/phase03_11_decision_center_proto_mainline/`
13. `.trae/specs/phase03_14_decision_center_integration_validation_acceptance/`
14. `docs/phase/phase03_decision_center_foundation_architecture_plan.md`
15. `docs/phase/phase03_decision_center_foundation_dev_plan.md`
16. `docs/phase/phase03_decision_center_foundation_shared_baseline.md`
17. `docs/phase/phase04_product_and_repository_binding_foundation_architecture_plan.md`
18. `docs/phase/phase04_product_and_repository_binding_foundation_dev_plan.md`
19. `docs/phase/phase04_product_and_repository_binding_foundation_shared_baseline.md`
20. 当前目标对应的 `phase / fix / audit` 文档

## 6. 接手提醒

- 不要在 `AGENTS.md` 重复写实现细节、task 清单或目录细节
- `plan.md` 只看 phase 级推进计划与进度，不看 task
- 需要技术栈时先看 `TECH_STACK_BASELINE.md`
- 需要规则时看 `project_rules.md`
- 需要 Trae 内部上下文补充时看 `project_skills.md` 与 `global_skills.md`
- 需要找当前活动文档时看 `docs/README.md`
