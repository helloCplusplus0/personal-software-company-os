# Tasks

- [x] Task 1: 冻结 MVP 页面范围。将 `v0.1` 页面级主范围写成单值结论，避免后续页面体系漂移。
  - [x] SubTask 1.1: 明确 `Dashboard / Module Registry / Product Registry / Decision Center / Repository Binding` 为 MVP 页面范围
  - [x] SubTask 1.2: 验证页面范围与 `phase01` 上游文档一致

- [x] Task 2: 冻结每个页面的最小职责。将页面职责压缩到 MVP 必需承载的动作和信息范围。
  - [x] SubTask 2.1: 明确 `Dashboard` 只承担最小聚合反馈
  - [x] SubTask 2.2: 明确 `Module Registry` 承担模块资产登记、`CreateRelease` 与查看入口
  - [x] SubTask 2.3: 明确 `Product Registry` 承担产品资产登记与绑定查看入口
  - [x] SubTask 2.4: 明确 `Decision Center` 承担决策记录与关联入口
  - [x] SubTask 2.5: 明确 `Repository Binding` 承担仓库创建与绑定关系入口

- [x] Task 3: 冻结页面与动作范围的一致性。确保页面范围可以映射到已冻结的动作范围。
  - [x] SubTask 3.1: 验证每个页面至少对应一组已冻结的核心动作（含 `CreateRelease` 由 `Module Registry` 承载）
  - [x] SubTask 3.2: 验证不存在无动作承载的空页面

- [x] Task 4: 冻结冷启动录入路径与低摩擦原则。将首轮用户从零建资产的最小路径写成可执行前提。
  - [x] SubTask 4.1: 明确首轮冷启动路径至少覆盖 `Product / Repository / Module / Decision` 基础资产建立
  - [x] SubTask 4.2: 明确空状态必须服务于首轮录入路径
  - [x] SubTask 4.3: 明确首轮以低摩擦手动录入为主，不依赖自动扫描或复杂导入

- [x] Task 5: 冻结非目标页面与入口边界。防止未来页面和独立工作台提前进入 `v0.1`。
  - [x] SubTask 5.1: 明确独立 `AI Assistant` 工作台不进入 `v0.1`
  - [x] SubTask 5.2: 明确 `Feature / Opportunity / Experiment` 页面不进入 `v0.1`
  - [x] SubTask 5.3: 验证没有第二套导航结构或第二套路径解释

- [x] Task 6: 完成规格校验。检查本次 `phase01-03` 规格是否具备进入后续子任务的条件。
  - [x] SubTask 6.1: 验证页面范围与动作范围一致
  - [x] SubTask 6.2: 验证未新增独立 AI 工作台
  - [x] SubTask 6.3: 验证未把未来页面提前写进 `v0.1`

# Task Dependencies

- `Task 2` depends on `Task 1`
- `Task 3` depends on `Task 1` and `Task 2`
- `Task 4` depends on `Task 1` and `Task 2`
- `Task 5` depends on `Task 1`
- `Task 6` depends on `Task 1`, `Task 2`, `Task 3`, `Task 4`, and `Task 5`