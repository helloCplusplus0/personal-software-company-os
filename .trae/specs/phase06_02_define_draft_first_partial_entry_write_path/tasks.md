# Tasks

- [x] Task 1: 冻结四类核心对象的最小必填字段。把 `Product / Repository / Module / Decision` 在首轮草稿创建时真正需要用户输入的字段压缩成单值结论。
  - [x] SubTask 1.1: 明确 `CreateDraftProduct` 的最小必填字段与系统默认补位字段（`description = ''`，`status = active`）
  - [x] SubTask 1.2: 明确 `CreateDraftRepository` 的最小必填字段与系统默认补位字段（`provider = manual`，`status = active`）
  - [x] SubTask 1.3: 明确 `CreateDraftModule` 的最小必填字段与系统默认补位字段（`description = ''`，`status = active`）
  - [x] SubTask 1.4: 明确 `CreateDraftDecision` 的最小必填字段与系统默认补位字段（`context = ''`，`problem = ''`，`impact = ''`，`alternatives = []`，`status = proposed`）

- [x] Task 2: 冻结 `draft created`、partial-entry 与首轮成功会话之间的边界。明确单对象草稿创建成功、四对象最小闭环完成以及绑定关系后补三者的区别。
  - [x] SubTask 2.1: 明确单对象 `draft created` 的成立条件
  - [x] SubTask 2.2: 明确四类对象都已 `draft created` 后才允许成立首轮成功会话
  - [x] SubTask 2.3: 明确哪些字段允许后补
  - [x] SubTask 2.4: 明确哪些绑定关系允许后补

- [x] Task 3: 冻结四类对象唯一 `application` 写入承接位与 `query` 只读边界。把 mutation 放回切片固定承接位，避免 `Onboarding` 与既有页面并存两套写语义。
  - [x] SubTask 3.1: 明确四类对象各自只有一个正式 `application` 写入承接位
  - [x] SubTask 3.2: 明确 `Onboarding` 与 canonical create 页面共享同一套写入语义
  - [x] SubTask 3.3: 明确 `query` / read adapter 只承接只读逻辑
  - [x] SubTask 3.4: 明确新增 phase06 写路径不得在 read adapter、query hook 或展示组件中继续散落 `useMutation`

- [x] Task 4: 冻结 phase06 必须回收的既有 create 页面 / 组件级 mutation 范围。把当前仓库中会与 `Onboarding` 冲突的 page-level mutation 范围写成明确清单。
  - [x] SubTask 4.1: 明确 `ProductCreatePage` 必须回收 page-level mutation
  - [x] SubTask 4.2: 明确 `RepositoryCreatePage` 必须回收 page-level mutation
  - [x] SubTask 4.3: 明确 `ModuleCreatePage` 必须回收 page-level mutation
  - [x] SubTask 4.4: 明确 `DecisionCreatePage` 必须回收 page-level mutation
  - [x] SubTask 4.5: 明确表单组件只保留字段收集、提交事件与表单内错误展示职责

- [x] Task 5: 完成 `phase06-02` 一致性校验。确认本次规格与 `phase06-01`、当前 shared baseline、既有 create 页面现状和 `phase06` 子任务目标保持一致。
  - [x] SubTask 5.1: 验证最小必填字段与 `phase06` 低摩擦首轮录入目标一致
  - [x] SubTask 5.2: 验证 `draft created` / `first-run completed` 边界与 `phase06-01` 首轮成功会话定义一致
  - [x] SubTask 5.3: 验证 `application / query / mutation` 分层与 `phase06` shared baseline 一致
  - [x] SubTask 5.4: 验证 create 页面回收清单与当前代码实际分布一致

# Task Dependencies

- `Task 2` depends on `Task 1`
- `Task 3` depends on `Task 1`, `Task 2`
- `Task 4` depends on `Task 3`
- `Task 5` depends on `Task 1`, `Task 2`, `Task 3`, and `Task 4`
