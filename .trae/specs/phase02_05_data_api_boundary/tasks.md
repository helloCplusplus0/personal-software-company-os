# Tasks

- [x] Task 1: 冻结 `Module Registry` 当前阶段的数据读写范围。将当前主线真正需要的读写对象与动作写成单值结论。
  - [x] SubTask 1.1: 明确当前阶段写入范围只承接 `CreateModule`、`CreateRelease`、`BindModuleToProduct`、`MapModuleToRepository`
  - [x] SubTask 1.2: 明确当前阶段读取范围只承接列表读取、详情读取与相关只读关联入口
  - [x] SubTask 1.3: 验证当前阶段未引入 `Module Registry` 之外的独立实现主线

- [x] Task 2: 冻结最小接口承接前提。将当前页面主线与接口承接关系写清。
  - [x] SubTask 2.1: 明确 `Module Registry / List` 承接列表读取接口
  - [x] SubTask 2.2: 明确 `Module Create` 承接创建写入接口
  - [x] SubTask 2.3: 明确 `Module Detail` 承接详情读取、版本写入与关联写入接口

- [x] Task 3: 冻结读动作与写动作的最小接口分组。将接口分组控制在当前阶段最小可执行范围内。
  - [x] SubTask 3.1: 明确 `ModuleListRead` 与 `ModuleDetailRead` 的最小读接口分组
  - [x] SubTask 3.2: 明确 `ModuleCreateWrite`、`ModuleReleaseWrite`、`ModuleBindingWrite` 的最小写接口分组
  - [x] SubTask 3.3: 验证未提前冻结完整查询矩阵
  - [x] SubTask 3.4: 明确绑定动作候选目标读取前提（`ProductBindingCandidateRead` 与 `RepositoryBindingCandidateRead`）只服务当前绑定动作，不扩写为独立主线

- [x] Task 4: 冻结 `Decision` 在当前阶段的接口边界。将其明确限制为只读展示或跳转前提。
  - [x] SubTask 4.1: 明确 `Decision` 当前阶段只承接只读读取或跳转前提
  - [x] SubTask 4.2: 明确 `Decision` 入口作为 `ModuleDetailRead` 的附属读取承接，不设独立读接口组
  - [x] SubTask 4.3: 验证未把独立 `RecordDecision` 写接口主线提前并入
  - [x] SubTask 4.4: 验证未把 `Decision Center` 全量接口提前并入

- [x] Task 5: 完成规格校验。检查本次 `phase02-05` 规格是否满足进入后续源码实现设计类任务的条件。
  - [x] SubTask 5.1: 验证当前阶段所需的数据与 API 边界已经明确
  - [x] SubTask 5.2: 验证当前规格与 `Contract First` 方向一致
  - [x] SubTask 5.3: 验证未提前冻结完整查询矩阵或 `Dashboard` 聚合接口

# Task Dependencies

- `Task 2` depends on `Task 1`
- `Task 3` depends on `Task 1` and `Task 2`
- `Task 4` depends on `Task 1` and `Task 2`
- `Task 5` depends on `Task 1`, `Task 2`, `Task 3`, and `Task 4`