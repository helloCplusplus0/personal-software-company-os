# Tasks

- [x] Task 1: 冻结 MVP 核心实体范围。将 `v0.1` 主执行范围内的对象写成单值结论，避免后续规格回到泛化叙述。
  - [x] SubTask 1.1: 明确 `Product / Module / Release / Decision / Repository / Venture（可选）` 进入 `v0.1` 主执行范围
  - [x] SubTask 1.2: 明确 `Decision` 必须保留在 MVP
  - [x] SubTask 1.3: 验证核心实体范围与 `phase01` 上游文档一致

- [x] Task 2: 冻结派生层与后移层。明确哪些对象不进入 `v0.1` 主执行范围，但仍保留在长期理论模型中。
  - [x] SubTask 2.1: 明确 `Capability` 为派生层，不作为重实体
  - [x] SubTask 2.2: 明确 `Feature / Opportunity / Experiment` 为后移对象
  - [x] SubTask 2.3: 验证派生层与后移层边界无冲突

- [x] Task 3: 冻结 MVP 核心动作范围。将首轮必须承接的动作写成可执行清单。
  - [x] SubTask 3.1: 明确 `CreateProduct / CreateModule / CreateRelease / RecordDecision / LinkDecisionToTarget` 为核心动作
  - [x] SubTask 3.2: 验证核心动作与核心实体范围一致

- [x] Task 4: 冻结 Repository Binding 最小动作。将 `Product / Module / Repository` 之间的绑定关系细化到可执行动作。
  - [x] SubTask 4.1: 明确 `BindRepositoryToProduct`
  - [x] SubTask 4.2: 明确 `BindModuleToProduct`
  - [x] SubTask 4.3: 明确 `MapModuleToRepository`
  - [x] SubTask 4.4: 验证 `Repository Binding` 不再停留在泛化命名

- [x] Task 5: 完成规格校验。检查本次 `phase01-02` 规格是否具备进入后续子任务的条件。
  - [x] SubTask 5.1: 验证 `Decision` 明确保留
  - [x] SubTask 5.2: 验证 `Capability` 明确为派生层
  - [x] SubTask 5.3: 验证 `Feature / Opportunity / Experiment` 明确后移
  - [x] SubTask 5.4: 验证绑定动作至少细化到 `BindRepositoryToProduct / BindModuleToProduct / MapModuleToRepository`

# Task Dependencies

- `Task 2` depends on `Task 1`
- `Task 3` depends on `Task 1`
- `Task 4` depends on `Task 1` and `Task 3`
- `Task 5` depends on `Task 1`, `Task 2`, `Task 3`, and `Task 4`