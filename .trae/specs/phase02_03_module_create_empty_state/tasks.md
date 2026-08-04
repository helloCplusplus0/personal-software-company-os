# Tasks

- [x] Task 1: 冻结 `CreateModule` 的最小表单。将首轮模块登记必须填写的字段写成单值结论，避免创建入口继续扩范围。
  - [x] SubTask 1.1: 明确最小表单字段为 `name / description / status`
  - [x] SubTask 1.2: 明确创建时不要求同步完成 `Product` 绑定、`Repository` 映射或 `Release` 登记
  - [x] SubTask 1.3: 验证未引入超出正式规格正文的复杂字段

- [x] Task 2: 冻结空状态引导路径。把“首次进入列表时如何完成首个模块登记”写成清晰闭环。
  - [x] SubTask 2.1: 明确空状态必须提供进入 `Module Create` 的主入口
  - [x] SubTask 2.2: 明确空状态文案围绕首个模块登记展开
  - [x] SubTask 2.3: 明确非空列表状态下仍保留创建入口但不再展示首轮空状态主提示

- [x] Task 3: 冻结列表、创建、详情之间的最小闭环。确保用户可以从读取进入创建，并从创建进入详情或返回列表。
  - [x] SubTask 3.1: 明确 `Module Registry / List -> Module Create`
  - [x] SubTask 3.2: 明确 `Module Create -> Module Detail` 或等价创建成功承接路径
  - [x] SubTask 3.3: 明确取消创建或创建失败时返回列表的路径

- [x] Task 4: 冻结当前主线的非目标边界。明确复杂导入、自动扫描与 AI 建议不进入当前创建主线。
  - [x] SubTask 4.1: 明确复杂导入不作为当前阶段创建入口
  - [x] SubTask 4.2: 明确自动扫描代码不作为当前阶段创建前提
  - [x] SubTask 4.3: 明确 AI 建议不作为首轮必需入口

- [x] Task 5: 完成规格校验。检查本次 `phase02-03` 规格是否满足进入后续子任务的条件。
  - [x] SubTask 5.1: 验证首轮用户能从零完成模块登记
  - [x] SubTask 5.2: 验证空状态与录入路径一致
  - [x] SubTask 5.3: 验证未把复杂导入、自动扫描或 AI 建议写入当前主线

# Task Dependencies

- `Task 2` depends on `Task 1`
- `Task 3` depends on `Task 1` and `Task 2`
- `Task 4` depends on `Task 1`
- `Task 5` depends on `Task 1`, `Task 2`, `Task 3`, and `Task 4`