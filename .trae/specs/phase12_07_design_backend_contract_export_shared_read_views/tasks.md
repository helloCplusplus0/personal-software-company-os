# Tasks

- [x] Task 1: 盘点 `phase12-07` 的直接上游冻结输入
  - [x] SubTask 1.1: 审阅 `dev_plan#L201-L219` 中 `phase12-07` 的范围与 DoD
  - [x] SubTask 1.2: 审阅 `phase12-05` 中已冻结的共享摘要、入口定位、三类页面 resolver 与候选判断题
  - [x] SubTask 1.3: 审阅 `architecture_plan §4.6B / §4.7` 与 `shared_baseline §3.5A` 的顺序、合同与设计模板冻结结论
  - [x] SubTask 1.4: 审阅当前 `project_context.proto` 与 `backend/internal/projectcontext/*` 的真实字段与服务边界（`design.md §1.1-1.2`）

- [x] Task 2: 产出后端影响对象清单与候选分组（写入 `design.md §2-3`）
  - [x] SubTask 2.1: 逐项列出真实 `.proto / Connect / service / renderer / frontend project-context adapter` 影响对象，并使用准确路径标记分类（`design.md §1.2` / `§2`）
  - [x] SubTask 2.2: 把共享摘要与入口定位需求拆成“继续复用 L1 真实字段 / 仅停留在 L3 或 renderer 单向映射 / 当前不做”三组（`design.md §3.1`）
  - [x] SubTask 2.3: 把 Product / Module / Decision 回锚 `repository_id` 的 resolver 候选逐项拆开，而不是合并成一句“既有 detail read 已覆盖”（`design.md §3.2`）
  - [x] SubTask 2.4: 对所有受控派生读取候选记录正式判断与回收收益说明（`design.md §3.3-§3.4`）

- [x] Task 3: 产出"复用 GetProjectContext vs 受控派生读取"的判断矩阵（写入 `design.md §3-4`）
  - [x] SubTask 3.1: 逐项判断哪些需求继续复用 `GetProjectContext` 真实字段（`design.md §3.1`：S1-S6）
  - [x] SubTask 3.2: 逐项判断哪些需求只允许停留在 L3 或 renderer 的单向映射（`design.md §3.1`：S7-S9）
  - [x] SubTask 3.3: 逐项判断哪些候选当前不做，并写清为何不足以进入 `ProjectContextService` 受控派生读取（`design.md §3.1-§3.3`）
  - [x] SubTask 3.4: 对 `Decision` resolver 候选单独写出正式评估，不再笼统归类为“已覆盖”（`design.md §3.2-§3.3`）

- [x] Task 4: 产出受控派生读取的最小合同与定位约束（写入 `design.md §4`）
  - [x] SubTask 4.1: 对每个候选写清最小输入锚点与其当前所属承接位（L1 / L3-renderer / 当前不做）（`design.md §3.1-§3.3` / `§4`）
  - [x] SubTask 4.2: 写清每个候选是否真的回收了前端 / agent 双侧重复解释逻辑（`design.md §3.3`）
  - [x] SubTask 4.3: 写清当前不做时如何继续保留 `repository_id` 与 `entry_ref / entry_kind` 的定位能力（`design.md §4.1-§4.2`）
  - [x] SubTask 4.4: 写清它与 `GetProjectContext`、`ExportProjectContext`、renderer 的关系边界（`design.md §4`）

- [x] Task 5: 产出导出结果与共享只读视图的关系设计（写入 `design.md §5`）
  - [x] SubTask 5.1: 冻结 `ExportProjectContext` 继续作为 agent-facing Markdown owner 的边界（`design.md §4.1-§4.2`）
  - [x] SubTask 5.2: 冻结 Web 共享只读结果与 renderer 的关系边界（`design.md §4.1-§4.2`）
  - [x] SubTask 5.3: 记录哪些结果共享事实源，哪些只是单向映射，不共享真相定义（`design.md §3.1` / `§4.2`）

- [x] Task 6: 完成与三件套及 `phase12-05` 的一致性校验
  - [x] SubTask 6.1: 校验 `phase12-05 -> 07 -> 06` 顺序与职责表达单值一致（`design.md §7`）
  - [x] SubTask 6.2: 校验不存在“把 L3 派生或 detail read 能力误写成后端真实字段”的结论（`design.md §3.1-§3.3` / `§6`）
  - [x] SubTask 6.3: 校验设计结果已满足统一最小模板与后端审计面（`design.md §1-§6`）
  - [x] SubTask 6.4: 校验本 spec 足以让后续执行者机械回答“这里是不是应该新建第二服务”（`design.md §3.3-§4` / `§6`）

# Task Dependencies

- Task 2 depends on Task 1
- Task 3 depends on Task 1
- Task 4 depends on Task 2
- Task 4 depends on Task 3
- Task 5 depends on Task 3
- Task 5 depends on Task 4
- Task 6 depends on Task 2
- Task 6 depends on Task 3
- Task 6 depends on Task 4
- Task 6 depends on Task 5
