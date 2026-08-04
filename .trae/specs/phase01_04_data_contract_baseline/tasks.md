# Tasks

- [x] Task 1: 冻结数据库主线。将 `v0.1` 当前数据库主线写成单值结论，并与 `Durable System Track` 和最终共识保持一致。
  - [x] SubTask 1.1: 明确 `PostgreSQL` 是当前唯一数据库主线
  - [x] SubTask 1.2: 明确 `Local First` 当前解释不等于切换到 `SQLite`
  - [x] SubTask 1.3: 验证当前规格没有引入第二套数据库解释

- [x] Task 2: 冻结最小数据模型方向。将 `v0.1` 数据结构压缩为核心表、关系表和派生视图三层方向。
  - [x] SubTask 2.1: 明确核心表至少包含 `ventures（可选） / products / repositories / modules / module_releases / decisions`
  - [x] SubTask 2.2: 明确关系表至少包含 `product_modules / product_repositories / module_repositories / decision_links`
  - [x] SubTask 2.3: 明确派生视图只承接 `Capability` 和反馈聚合结果

- [x] Task 3: 冻结 `Repository Binding` 结构要求。将产品、模块和仓库之间的关系落到可执行的数据结构方向。
  - [x] SubTask 3.1: 明确产品与仓库的绑定关系需要显式承载
  - [x] SubTask 3.2: 明确模块与产品的绑定关系需要显式承载
  - [x] SubTask 3.3: 明确模块与仓库的实现映射通过 `module_repositories` 显式承载
  - [x] SubTask 3.4: 验证 `Repository Binding` 没有退回为页面层泛化说明

- [x] Task 4: 冻结 `Decision Record` 结构要求。将决策记录压缩为可检索、可关联的最小结构。
  - [x] SubTask 4.1: 明确 `Decision` 最小字段至少包含 `title / context / problem / alternatives / choice / reason / impact / status`
  - [x] SubTask 4.2: 明确 `Decision` 必须通过结构化关系链接到目标对象
  - [x] SubTask 4.3: 验证 `Decision` 没有被退化为散装备注

- [x] Task 5: 冻结 `Contract First` 基线。明确当前项目 API 与跨语言边界的最小合同方向。
  - [x] SubTask 5.1: 明确后续 API 规格必须以结构化合同优先，而不是由前端猜字段
  - [x] SubTask 5.2: 明确当前长期合同方向继续冻结为 `Protocol Buffers`
  - [x] SubTask 5.3: 明确当前阶段不要求完整 `proto` 工具链
  - [x] SubTask 5.4: 验证未引入与 `Protocol Buffers` 冲突的第二套跨语言合同主线

- [x] Task 6: 完成规格校验。检查本次 `phase01-04` 规格是否具备进入后续子任务的条件。
  - [x] SubTask 6.1: 验证数据主线与 `Durable System Track` 一致
  - [x] SubTask 6.2: 验证没有第二套数据库解释
  - [x] SubTask 6.3: 验证没有与 `Protocol Buffers` 长期方向冲突的跨语言路线
  - [x] SubTask 6.4: 验证本次规格与 `phase01` 三件套和最终共识保持一致

# Task Dependencies

- `Task 2` depends on `Task 1`
- `Task 3` depends on `Task 2`
- `Task 4` depends on `Task 2`
- `Task 5` depends on `Task 1` and `Task 2`
- `Task 6` depends on `Task 1`, `Task 2`, `Task 3`, `Task 4`, and `Task 5`