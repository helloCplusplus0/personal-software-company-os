# AGENTS.md

> 职责：PSCO 项目的全局上下文入口。
> 这是 OpenAI agent 默认读取的项目级入口文档，目标是让接手者快速恢复正确上下文，而不是在这里展开重复规则或阶段细节。

## 1. 项目定位

- 项目名称：`Personal Software Company OS`
- 当前阶段：`phase01_mvp_spec_convergence (/plan)`
- 当前主目标：把 PSCO 从“完整理念体系”压缩成“可在 4-8 周内落地的第一版执行规格”
- 当前下一阶段入口：`phase02_module_registry_foundation`
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
- `docs/` 当前只服务 `phase / fix / audit / archive` workflow

## 4. 当前状态

- 原始方案文档 `PSCO_0.md ~ PSCO_4.md` 已完成第一轮共识回正
- 专家评审与交叉汇总文档已归类到 `docs/audit/`
- 根目录保留真相源与主入口，不再作为散装文档堆放区
- `phase01_*` 三件套已创建，当前处于 `/plan` 审核前状态

## 5. 推荐阅读顺序

1. `AGENTS.md`
2. `plan.md`
3. `TECH_STACK_BASELINE.md`
4. `project_rules.md`
5. `architecture_map.md`
6. `PSCO-summarize-feedback.md`
7. `docs/README.md`
8. 当前目标对应的 `phase / fix / audit` 文档

## 6. 接手提醒

- 不要在 `AGENTS.md` 重复写实现细节、task 清单或目录细节
- `plan.md` 只看 phase 级推进计划与进度，不看 task
- 需要技术栈时先看 `TECH_STACK_BASELINE.md`
- 需要规则时看 `project_rules.md`
- 需要 Trae 内部上下文补充时看 `project_skills.md` 与 `global_skills.md`
- 需要找当前活动文档时看 `docs/README.md`
