# phase02_module_registry_foundation_dev_plan

## 1. 文档定位

本文档定义 `phase02_module_registry_foundation` 的执行顺序、子任务范围、DoD 与明确不做。

`phase02` 当前处于 `/plan` 阶段，因此本文档是后续 `/spec` 与收口的唯一任务来源。

## 2. 本阶段目标

在 `phase01-06` 正式 MVP 规格正文前提下，建立 `Module Registry` 最小可执行主线，作为 `v0.1` 首个直接进入实现的核心能力入口。

## 3. 子任务清单

### 第一组：冻结类子任务

### phase02-01 冻结 Module Registry 页面边界与信息结构

范围：

- 冻结模块列表页、创建入口、详情页的最小结构
- 冻结页面之间的跳转关系
- 冻结桌面与移动浏览器下的基础信息密度策略

DoD：

- `Module Registry` 页面职责单值化
- 无第二套移动端 UI 方案
- 页面边界与 `phase01-06` 正式规格正文一致

### phase02-02 冻结 Module 实体与版本主线

范围：

- 冻结 `Module` 的最小展示字段与状态表达
- 冻结 `Release` 的最小登记与展示方式
- 冻结模块准入规则在页面侧的承接方式

DoD：

- `Module` 与 `Release` 的最小主线明确
- 模块准入规则在页面侧可落地
- 不引入超出正式规格正文的新对象解释

### phase02-03 冻结模块创建与空状态路径

范围：

- 冻结首轮 `CreateModule` 的最小表单与录入路径
- 冻结空状态如何引导用户完成首个模块登记
- 冻结列表、详情、创建之间的最小闭环

DoD：

- 首轮用户能从零完成模块登记
- 空状态与录入路径一致
- 不把复杂导入、自动扫描或 AI 建议写入当前主线

### phase02-04 冻结版本登记、关联动作归属与最小读模型

范围：

- 冻结 `CreateRelease` 的最小交互
- 冻结 `BindModuleToProduct` 与 `MapModuleToRepository` 的页面承接归属
- 冻结到 `Decision` 的最小关联入口
- 冻结列表读取、详情读取、创建写入、版本写入的最小读写模型

DoD：

- 版本登记路径明确
- `BindModuleToProduct` 与 `MapModuleToRepository` 的动作拥有者明确
- 模块与 `Product / Repository / Decision` 的连接方式明确
- 列表页与详情页的最小读模型明确
- 不把 `phase03+` 的独立主线提前并入

### phase02-05 冻结数据读写范围与 API 承接前提

范围：

- 冻结 `Module Registry` 所需的数据读写范围
- 冻结最小接口承接前提
- 冻结写动作与读动作的最小接口分组
- 冻结 `Decision` 在当前阶段只读/跳转而不扩写为独立写主线的接口边界

DoD：

- 当前阶段所需的数据与 API 边界明确
- 与 `Contract First` 方向一致
- 不提前冻结完整查询矩阵或 Dashboard 聚合接口

### 第二组：源码实现设计产出类子任务

### phase02-06 产出前端页面、路由与组件分层设计

范围：

- 冻结 `Module List / Module Create / Module Detail / Release Create` 的页面分层
- 冻结页面之间的最小路由结构
- 冻结列表、表单、详情、关联面板的组件分层
- 冻结 `PC / 移动浏览器` 双场景下的布局降级策略

DoD：

- 前端页面与路由分层明确
- 页面级组件职责明确
- 无第二套移动端 UI 架构
- 设计结果足以直接进入 `/spec`

### phase02-07 产出前端状态模型与交互流设计

范围：

- 冻结列表页查询状态、筛选状态与空状态的最小模型
- 冻结创建表单、版本登记表单与关联动作面板的交互状态流转
- 冻结成功、失败、空状态与返回路径

DoD：

- 页面级状态模型明确
- 列表、详情、创建、版本登记之间的交互流明确
- 不把运行时实现细节提前写成当前既成事实

### phase02-08 产出后端模块边界与接口分组设计

范围：

- 冻结 `Module Registry` 在后端的模块边界
- 冻结列表读取、详情读取、创建写入、版本写入、关联写入的接口分组
- 冻结 `Module Registry` 与 `Product / Repository / Decision` 的服务侧连接边界

DoD：

- 后端模块边界明确
- 读写接口分组明确
- 不提前冻结 Go 数据访问层具体工具
- 设计结果足以直接进入 `/spec`

### phase02-09 产出首份 Module Registry 正式规格文档

范围：

- 基于前八个子任务产出 `phase02` 对应的 `/spec`
- 作为后续实现与 `phase03` 的直接上游规格来源

DoD：

- 文档完整覆盖页面、动作、数据读写、API、空状态、非目标、源码设计层结果与 Done 标准
- 与 `phase01-06` 正式 MVP 规格正文互链一致

### phase02-10 审核与根级同步

范围：

- 完成 `phase02` 文档互链复核
- 回写根级状态
- 确认下一阶段入口

DoD：

- 根级文档与 `phase02` 文档保持单值一致
- `plan.md` 中 `phase02` 状态更新正确
- `phase03` 的进入条件清楚

## 4. 明确不做

本阶段不做：

- 代码实现
- 运行时验证
- 数据库 migration
- Product / Decision / Repository 的独立阶段实现
- Dashboard 聚合反馈实现
- 独立 `React Native` 客户端
- 完整 `PWA`

## 5. 依赖关系

执行顺序固定为：

1. `phase02-01`
2. `phase02-02`
3. `phase02-03`
4. `phase02-04`
5. `phase02-05`
6. `phase02-06`
7. `phase02-07`
8. `phase02-08`
9. `phase02-09`
10. `phase02-10`

不允许跳过前置冻结，直接进入 `Module Registry` 正式规格正文或直接进入实现。
