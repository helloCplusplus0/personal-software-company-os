# phase01_mvp_spec_convergence_dev_plan

## 1. 文档定位

本文档定义 `phase01_mvp_spec_convergence` 的执行顺序、子任务范围、DoD 与明确不做。

本文档在 `phase01` 的 `/plan` 阶段定义后续 `/spec` 与收口的唯一任务来源。

> 状态说明：上述表述描述的是 `phase01` 的规划时语境。当前 `phase01` 已完成收口；执行层唯一规格入口为 `.trae/specs/phase01_06_formal_mvp_spec/mvp_spec_v0.1.md`，本文档仅保留为规划任务上游。

## 2. 本阶段目标

在 `Durable System Track` 前提下，完成 PSCO `v0.1` MVP 的规格收敛，并产出可审阅、可冻结、可承接后续实现的单值规格文档。

## 3. 子任务清单

### phase01-01 冻结技术路线与系统边界

范围：

- 明确 PSCO 当前项目走 `Durable System Track`
- 明确前端、后端、数据库、部署与跨语言合同方向
- 明确 `Rust` 在 `v0.1` 中的排除边界

DoD：

- 技术路线单值化
- 无 `Product Track` / `Durable System Track` 混写
- 无 `AGENTS-OLD.md` 旧技术栈残留解释

### phase01-02 冻结 MVP 对象与动作范围

范围：

- 冻结核心实体
- 冻结派生层与后移层
- 冻结 MVP 核心动作清单
- 冻结 `Product / Module / Repository` 之间的最小绑定动作

DoD：

- `Decision` 明确保留
- `Capability` 明确为派生层
- `Feature / Opportunity / Experiment` 明确后移
- `Repository Binding` 不再停留在泛化表述，至少细化到 `BindRepositoryToProduct / BindModuleToProduct / MapModuleToRepository`

### phase01-03 冻结页面与输入路径

范围：

- 冻结 MVP 页面范围
- 冻结每个页面的最小职责
- 冻结录入路径与低摩擦原则

DoD：

- 页面范围与动作范围一致
- 不新增独立 AI 工作台
- 不把未来页面提前写进 `v0.1`

### phase01-04 冻结数据与合同基线

范围：

- 冻结数据库主线与最小数据模型方向
- 冻结 API / contract-first 基线
- 冻结 repository binding 与 decision record 的结构要求

DoD：

- 数据主线与 `Durable System Track` 一致
- 不引入第二套数据库解释
- 不引入与 `Protocol Buffers` 长期方向冲突的跨语言路线

### phase01-05 冻结冷启动、导入与导出要求

范围：

- 冻结 MVP 冷启动路径
- 冻结导入路径与手动录入边界
- 冻结最小导出 / 备份要求

DoD：

- 已明确首轮用户如何从零开始建立 `Product / Module / Decision / Repository` 基础资产
- 已明确当前阶段是否支持导入、支持什么导入、哪些仍采用手动录入
- 已明确 `Local First` 语义下的最小导出 / 备份要求

### phase01-06 产出首份正式 MVP 规格文档

范围：

- 基于前五个子任务产出 `phase01` 对应的 `/spec`
- 作为后续实现与 phase02 的唯一上游规格来源

DoD：

- 文档完整覆盖对象、动作、页面、数据、API、非目标、Done 标准
- 与根级真相源互链一致

### phase01-07 审核与根级同步

范围：

- 完成 phase01 文档互链复核
- 回写根级状态
- 确认下一阶段入口

DoD：

- 根级文档与 phase 文档保持单值一致
- `plan.md` 中 phase01 状态更新正确
- phase02 的进入条件清楚

## 4. 明确不做

本阶段不做：

- 代码实现
- 运行时验证
- 数据库 migration
- 前端 UI 细节打磨
- Go 数据访问层工具的最终实现选型落地
- Rust 计算引擎接入

## 5. 依赖关系

执行顺序固定为：

1. `phase01-01`
2. `phase01-02`
3. `phase01-03`
4. `phase01-04`
5. `phase01-05`
6. `phase01-06`
7. `phase01-07`

不允许跳过前置冻结，直接写规格正文或直接进入实现。
