# phase02_module_registry_foundation_dev_plan

## 1. 文档定位

本文档定义 `phase02_module_registry_foundation` 的执行顺序、子任务范围、DoD 与明确不做。

`phase02` 虽然当前处于 `/plan` 阶段，但它本身不是纯文档冻结阶段，而是一个交付型 phase。因此本文档既要定义后续 `/spec` 的任务来源，也要定义后续实现、验收与收口的推进顺序。

## 2. 本阶段目标

在 `phase01-06` 正式 MVP 规格正文前提下，交付 `Module Registry` 最小可执行主线，作为 `v0.1` 首个直接进入实现并完成验收的核心能力入口。

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

### 第二组：实现设计产出类子任务

### phase02-06 产出前端页面、路由与组件分层设计

范围：

- 产出 `Module List / Module Create / Module Detail / Release Create` 的页面分层
- 产出页面之间的最小路由结构
- 产出列表、表单、详情、关联面板的组件分层
- 产出 `PC / 移动浏览器` 双场景下的布局降级策略

DoD：

- 前端页面与路由分层明确
- 页面级组件职责明确
- 无第二套移动端 UI 架构
- 设计结果足以直接进入实现

### phase02-07 产出前端状态模型与交互流设计

范围：

- 产出列表页查询状态、筛选状态与空状态的最小模型
- 产出创建表单、版本登记表单与关联动作面板的交互状态流转
- 产出成功、失败、空状态与返回路径

DoD：

- 页面级状态模型明确
- 列表、详情、创建、版本登记之间的交互流明确
- 设计结果足以直接进入实现
- 不把运行时实现细节提前写成当前既成事实

### phase02-08 产出后端模块边界与接口分组设计

范围：

- 产出 `Module Registry` 在后端的模块边界
- 产出列表读取、详情读取、创建写入、版本写入、关联写入的接口分组
- 产出 `Module Registry` 与 `Product / Repository / Decision` 的服务侧连接边界

DoD：

- 后端模块边界明确
- 读写接口分组明确
- 不提前冻结 Go 数据访问层具体工具
- 设计结果足以直接进入实现

### 第三组：规格、实现与验收子任务

### phase02-09 产出首份 Module Registry 正式规格文档

范围：

- 基于前八个子任务产出 `phase02` 对应的 `/spec`
- 作为后续实现与 `phase03` 的直接上游规格来源

DoD：

- 文档完整覆盖页面、动作、数据读写、API、空状态、非目标、实现设计层结果与 Done 标准
- 与 `phase01-06` 正式 MVP 规格正文互链一致

### phase02-10 实现前端 Module Registry 主线

范围：

- 实现 `Module List / Module Create / Module Detail / Release Create` 前端主线
- 实现列表、创建、详情、版本登记与关联面板的最小交互
- 落实单一 `React Web` 下的 `PC / 移动浏览器` 双场景布局策略

DoD：

- 前端主线可运行
- 列表、创建、详情、版本登记与关联动作在前端可走通
- 无第二套移动端 UI 架构

### phase02-11 实现后端与数据主线

范围：

- 实现 `Module Registry` 所需的最小后端读写接口
- 实现 `modules / module_releases / product_modules / module_repositories` 的数据主线
- 实现 `Decision` 在当前阶段所需的只读/跳转边界

DoD：

- 后端读写接口可运行
- 数据主线与已冻结边界一致
- 不引入超出当前阶段的新对象解释或第二套数据主线

### phase02-12 联调、验证与验收

范围：

- 完成前后端联调
- 完成 `CreateModule / CreateRelease / BindModuleToProduct / MapModuleToRepository` 的最小验收路径
- 完成空状态、错误态与返回路径验证

DoD：

- `Module Registry` 最小主线可完整走通
- 验收结果可明确证明当前阶段已形成可运行交付物
- 发现的问题已收口到当前阶段，不遗留隐性阻断

### phase02-13 审核与根级同步

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

- `Feature / Opportunity / Experiment` 的重新引入
- Product / Decision / Repository 的独立阶段实现
- Dashboard 聚合反馈实现
- 自动扫描代码
- 自动知识图谱
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
11. `phase02-11`
12. `phase02-12`
13. `phase02-13`

不允许跳过前置收敛与设计，直接进入 `Module Registry` 实现；也不允许只完成文档冻结而不完成本阶段的代码交付、验收与收口。
