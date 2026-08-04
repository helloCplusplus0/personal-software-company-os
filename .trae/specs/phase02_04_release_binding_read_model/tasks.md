# Tasks

- [x] Task 1: 冻结 `CreateRelease` 的最小交互路径。将版本登记入口、完成回流与上下文归属写成单值结论。
  - [x] SubTask 1.1: 明确 `CreateRelease` 只能从 `Module Detail` 进入
  - [x] SubTask 1.2: 明确版本登记完成后回流到当前模块详情上下文
  - [x] SubTask 1.3: 验证当前阶段未引入独立版本管理主线

- [x] Task 2: 冻结 `BindModuleToProduct` 与 `MapModuleToRepository` 的动作拥有者。确保绑定动作不再漂移到其他主线。
  - [x] SubTask 2.1: 明确 `BindModuleToProduct` 由 `Module Detail` 直接承接
  - [x] SubTask 2.2: 明确 `MapModuleToRepository` 由 `Module Detail` 直接承接
  - [x] SubTask 2.3: 验证当前阶段不是只跳转、不写入

- [x] Task 3: 冻结 `Decision` 的最小关联入口。将 `Decision` 控制在只读展示或跳转范围内。
  - [x] SubTask 3.1: 明确 `Decision` 在模块详情中只作为只读展示或跳转入口
  - [x] SubTask 3.2: 验证未把 `Decision Center` 全量主线提前并入当前阶段

- [x] Task 4: 冻结列表页与详情页的最小读模型。把当前阶段真正需要读取的内容写清。
  - [x] SubTask 4.1: 明确列表页最小读模型
  - [x] SubTask 4.2: 明确详情页最小读模型
  - [x] SubTask 4.3: 验证最小读模型只服务当前阶段主线，不扩写为复杂分析工作台

- [x] Task 5: 完成规格校验。检查本次 `phase02-04` 规格是否满足进入后续子任务的条件。
  - [x] SubTask 5.1: 验证版本登记路径已经明确
  - [x] SubTask 5.2: 验证 `BindModuleToProduct` 与 `MapModuleToRepository` 的动作拥有者已经明确
  - [x] SubTask 5.3: 验证模块与 `Product / Repository / Decision` 的连接方式已经明确
  - [x] SubTask 5.4: 验证列表页与详情页的最小读模型已经明确
  - [x] SubTask 5.5: 验证未把 `phase03+` 的独立主线提前并入

# Task Dependencies

- `Task 2` depends on `Task 1`
- `Task 3` depends on `Task 1`
- `Task 4` depends on `Task 1`
- `Task 5` depends on `Task 1`, `Task 2`, `Task 3`, and `Task 4`