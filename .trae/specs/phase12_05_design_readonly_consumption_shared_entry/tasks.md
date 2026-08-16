# Tasks

- [x] Task 1: 盘点 `phase12-05` 的直接上游冻结输入
  - [x] SubTask 1.1: 审阅 `dev_plan#L154-L173` 中 `phase12-05` 的范围与 DoD
  - [x] SubTask 1.2: 审阅 `phase12-03` 中只读消费深化边界、共享只读 owner 与三类页面承接矩阵
  - [x] SubTask 1.3: 审阅 `phase11-07 / 08` 中既有结构化读取与 Markdown 导出能力边界
  - [x] SubTask 1.4: 审阅 `architecture_plan §4.4A/§4.6A/§4.6B` 与 `shared_baseline §3.4` 的冻结结论

- [x] Task 2: 产出共享只读 owner 与共享入口设计矩阵（写入 `design.md §2-§4`）
  - [x] SubTask 2.1: 固定 `GetProjectContext` 为 L1 结构化 canonical owner（`design.md §2.1` / `§3.1`）
  - [x] SubTask 2.2: 固定 `ExportProjectContext` 为 L2 agent-facing Markdown 导出 owner（`design.md §2.1` / `§3.1`）
  - [x] SubTask 2.3: 固定 `frontend/src/features/project-context/` 为 L3 Web 跨切片共享只读 owner（`design.md §2.1` / `§3.1`）
  - [x] SubTask 2.4: 固定各 feature 的 `pages/`、`components/`、`data/` 为 L4 切片内展示 owner（`design.md §2.1` / `§3.1`）
  - [x] SubTask 2.5: 明确 `project-context/` 的启用条件、共享语义来源边界与禁止事项（`design.md §4` / `§7`）

- [x] Task 3: 产出三类页面承接矩阵与 resolver 责任边界（写入 `design.md §3`）
  - [x] SubTask 3.1: 冻结 `repositories/$repositoryId` 为直接 repository-scoped 页面（`design.md §3.2`）
  - [x] SubTask 3.2: 分别冻结 `products/$productId`、`modules/$moduleId`、`decisions/$decisionId` 的间接 resolver 规则，不再合并成一句笼统表述（`design.md §2.2` / `§3.2-§3.3`）
  - [x] SubTask 3.3: 冻结 `dashboard`、`onboarding`、`reviews/daily`、`reviews/weekly` 为衍生消费页（`design.md §3.2`）
  - [x] SubTask 3.4: 为三类页面分别写出输入、解析、复用、唯一候选条件与失败语义责任边界（`design.md §3.2-§3.3`）
  - [x] SubTask 3.5: 明确禁止页面内搜索、手工查询、工作区扫描或静默选取多候选锚点（`design.md §3.2-§3.3` / `§7`）

- [x] Task 4: 产出"复用既有合同 vs 新增最小承接位"的判定规则（写入 `design.md §2` / `§6` / `§7`）
  - [x] SubTask 4.1: 明确哪些共享需求已可继续复用 `GetProjectContext` 真实字段（`design.md §2.3` / `§6.1`）
  - [x] SubTask 4.2: 明确哪些共享结果只能作为 `phase12-07` 候选评估，而不能在 `phase12-05` 预支为“已覆盖”（`design.md §6.3-§6.4`）
  - [x] SubTask 4.3: 明确哪些内容继续停留在 L3 单向派生或 L4 切片展示层（`design.md §4` / `§6.2`）
  - [x] SubTask 4.4: 明确任何新增承接位都必须写清输入锚点、回收逻辑与失败条件，且不得通过禁止事项绕过 `phase12-07`（`design.md §6.4` / `§7`）

- [x] Task 5: 输出供 `phase12-07` 继续承接的正式输入需求（写入 `design.md §6`）
  - [x] SubTask 5.1: 列出当前真实已覆盖的共享摘要字段与入口定位字段（`design.md §6.1`）
  - [x] SubTask 5.2: 列出可由 L3 单向派生、但不属于 L1 真实字段的共享结果（`design.md §6.2`）
  - [x] SubTask 5.3: 列出 Product / Module / Decision 的 resolver 候选补充项，并明确哪些只能进入 `phase12-07` 候选评估（`design.md §6.3`）
  - [x] SubTask 5.4: 显式写出 `phase12-07` 必须继续回答的判断题，而不是在 `phase12-05` 预支“无需新增”的结论（`design.md §6.4`）

- [x] Task 6: 完成与三件套及上游 spec 的一致性校验
  - [x] SubTask 6.1: 校验 `dev_plan`、`architecture_plan`、`shared_baseline` 与本 spec 对 `phase12-05` 职责表达单值一致（`design.md §8.2`）
  - [x] SubTask 6.2: 校验 owner 分层、三类页面矩阵、`repository_id` 锚点规则与复用判定规则单值一致（`design.md §2-§4`）
  - [x] SubTask 6.3: 校验 `phase12-05 -> 07 -> 06` 顺序与交接边界单值一致（`design.md §8.1`）
  - [x] SubTask 6.4: 校验本 spec 已足以让后续执行者机械回答"这里该复用 phase11，还是该新增最小只读承接位"与"这里该直接用 repository_id，还是先解析回 repository_id 再复用共享摘要"（`design.md §3` / `§6`）
  - [x] SubTask 6.5: 校验 `design.md` 已满足统一最小设计模板：影响对象清单、结论矩阵、承接位矩阵、共享语义来源 vs 切片内渲染矩阵、Before/After、明确不做清单（`design.md §1-§7`）

# Task Dependencies

- Task 2 depends on Task 1
- Task 3 depends on Task 1
- Task 4 depends on Task 1
- Task 5 depends on Task 2
- Task 5 depends on Task 3
- Task 5 depends on Task 4
- Task 6 depends on Task 2
- Task 6 depends on Task 3
- Task 6 depends on Task 4
- Task 6 depends on Task 5
