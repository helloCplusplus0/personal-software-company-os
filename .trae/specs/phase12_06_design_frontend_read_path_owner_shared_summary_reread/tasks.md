# Tasks

- [x] Task 1: 盘点 `phase12-06` 的上游冻结输入与现有前端 read owner
  - [x] SubTask 1.1: 审阅 `dev_plan#L175-L199` 中 `phase12-06` 的范围与 DoD
  - [x] SubTask 1.2: 审阅 `phase12-05` 中已冻结的共享 owner、三类页面承接矩阵与 resolver 输入规则
  - [x] SubTask 1.3: 审阅 `phase12-07` 中已冻结的 L1 真实字段、L3 单向映射边界与当前不做候选
  - [x] SubTask 1.4: 审阅 8 个现有 `use-*-read.ts` 的输入锚点、缓存键、输出 shape 与页面职责（`design.md §1.1`：R1-R8 全部为纯数据映射，无语义解释，全部分类为 `no-change`）
  - [x] SubTask 1.5: 确认当前 `frontend/src/features/project-context/` 不存在，记录需新增 `data/*`（`design.md §1.2`：3 个新文件）

- [x] Task 2: 产出前端 read owner 影响对象清单与分类矩阵（写入 `design.md §1`）
  - [x] SubTask 2.1: 逐项列出 8 个既有 read owner 与 3 个新增 `project-context/data/*`（`design.md §1.1-1.2`）
  - [x] SubTask 2.2: 对每个对象标记 `must-change / follow-regression / no-change`（`design.md §1.3`：must-change 3 个、no-change 8 个）
  - [x] SubTask 2.3: 对每个 read owner 记录缓存键、输入锚点、输出 shape、是否消费共享只读（`design.md §1.1` 表）
  - [x] SubTask 2.4: 对 detail 页、dashboard、onboarding、daily review、weekly review 补充页面级消费对象（`design.md §3.1-3.3`）

- [x] Task 3: 产出 `query` 层、L3 共享只读 owner 与页面层的分工矩阵（写入 `design.md §2`）
  - [x] SubTask 3.1: 冻结哪些逻辑继续留在切片 `use-*-read.ts`（`design.md §2.2`：8 个切片 read owner 全部保留原有职责）
  - [x] SubTask 3.2: 冻结哪些共享摘要、入口定位视图与高频语义来源应进入 `project-context/data/*`（`design.md §2.2`：共享语义标签 → L3-1、GetProjectContext query → L3-2、入口定位裁剪 → L3-3）
  - [x] SubTask 3.3: 冻结哪些页面专属字段与转换明确不进入共享只读（`design.md §2.3`：8 项页面专属字段）
  - [x] SubTask 3.4: 冻结 `project-context/data/*` 的最小文件边界与 owner 身份（`design.md §5`：3 文件完整设计）

- [x] Task 4: 产出页面读取、缓存、成功回流与 reread 关系设计（写入 `design.md §3`）
  - [x] SubTask 4.1: 逐项设计四个 detail 页接入共享摘要时的读取先后关系（`design.md §3.1-3.2`：直接/间接页面流程图）
  - [x] SubTask 4.2: 逐项设计 dashboard / onboarding / daily review / weekly review 的共享只读消费方式（`design.md §3.3`：衍生页只消费 L3-1 静态常量）
  - [x] SubTask 4.3: 写清 mutation 成功后哪些失效动作仍归切片 owner，哪些共享只读 query 需要 reread（`design.md §3.4`：按“是否拿到唯一 repositoryId”冻结精确失效规则）
  - [x] SubTask 4.4: 区分初次加载、局部重试、整页重试与成功回流后的 reread 范围（`design.md §3.5`：四种场景的行为表）

- [x] Task 5: 产出散装解释逻辑回收清单（写入 `design.md §4`）
  - [x] SubTask 5.1: 冻结 detail 页共享上下文区未来入口定位 adapter 的唯一承接位（`design.md §4.2`：不再把 Dashboard / Review 误列为现成 consumer）
  - [x] SubTask 5.2: 盘点 dashboard / onboarding / review 与 summary card 内重复的四实体解释，并区分哪些只回收语义标签、哪些继续留在切片内 UI（`design.md §4.1` / `§4.3`）
  - [x] SubTask 5.3: 判断哪些重复逻辑应进入共享只读 data owner，哪些继续留在切片内渲染（`design.md §4.3`：5 处不回收 + 理由）
  - [x] SubTask 5.4: 校验该回收清单不反向改写 `phase12-04` 的 primary owner 页面职责（`design.md §4.4`：与 phase12-04 一致性声明）

- [x] Task 6: 完成模板、一致性与不做项校验
  - [x] SubTask 6.1: 校验设计结果满足统一最小模板：影响对象清单、结论矩阵、承接位矩阵、共享语义来源 vs 切片内渲染矩阵、Before/After、明确不做清单（`design.md §1-8`）
  - [x] SubTask 6.2: 校验 `phase12-06` 没有反向改写 `phase12-05 / 07`（`design.md §7` 不做清单 #8-10 + `§8` 一致性声明）
  - [x] SubTask 6.3: 校验不会把跨切片共享只读逻辑继续散落在 `dashboard / onboarding / review / detail page` 各自的数据层（`design.md §2.1` 三层职责边界 + `§4` 回收清单）
  - [x] SubTask 6.4: 校验本设计足以指导后续实现收口，而不需要页面层临场补第二套 query contract 或刷新策略（`design.md §5` 文件边界 + `§3.4` reread 策略 + `§6` Before/After 样例）

# Task Dependencies

- Task 2 depends on Task 1
- Task 3 depends on Task 1
- Task 4 depends on Task 2
- Task 4 depends on Task 3
- Task 5 depends on Task 2
- Task 5 depends on Task 3
- Task 6 depends on Task 2
- Task 6 depends on Task 3
- Task 6 depends on Task 4
- Task 6 depends on Task 5
