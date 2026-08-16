# Tasks

- [x] Task 1: 盘点 `phase12-04` 的上游冻结输入
  - [x] SubTask 1.1: 审阅 `phase12-02` 中四实体 Web 端语义承接矩阵、surface 审计面与 `no-change` 规则
  - [x] SubTask 1.2: 审阅 `phase12-03` 中共享只读 owner、共享呈现边界与越权禁止项
  - [x] SubTask 1.3: 审阅 `dev_plan`、`architecture_plan` 与 `shared_baseline` 中 `phase12-04` 的范围、DoD 与最小设计产物模板

- [x] Task 2: 产出前端影响对象清单与分类矩阵（写入 `design.md §1`）
  - [x] SubTask 2.1: 逐项列出所有 route 对象并标记分类（`design.md §1.1`：8 个 Route 全部 `no-change`）
  - [x] SubTask 2.2: 逐项列出所有 page 对象并标记分类（`design.md §1.2`：P1-P4 与 P6 `must-change`，P5/P7/P8 `follow-regression`）
  - [x] SubTask 2.3: 逐项列出所有 component 对象并标记分类（`design.md §1.3`：C1-C12，C1-C4 与 C9 `must-change`，C5-C8 `follow-regression`，C10-C12 `no-change`）
  - [x] SubTask 2.4: 对四个 list page 文件、四个 list toolbar 与两个搜索 store 记录分类（`design.md §1.4`：N1-N4 `must-change` 但仍非 primary owner，N5-N10 `follow-regression`）

- [x] Task 3: 产出四实体 surface 承接矩阵（写入 `design.md §2`）
  - [x] SubTask 3.1: 对摘要卡片产出逐项承接矩阵（`design.md §2.1`：ProductSummaryCard / RepositorySummaryCard / ModuleSummaryCard / DecisionDetailSummaryCard 全部 `must-change`）
  - [x] SubTask 3.1A: 对四个 primary owner detail page 产出页面级语义承接位矩阵（`design.md §2.1A`：P1-P4 页面标题区与 SummaryCard 之间的语义导语全部 `must-change`）
  - [x] SubTask 3.2: 对空态产出逐项承接矩阵（`design.md §2.2`：四实体 List 页空态 / Dashboard Current Focus、Asset Feedback、Recent Activity 空态 / Daily Review 三类空态 / Weekly Review 三类空态 / Onboarding Welcome、Product、Repository、Module、Decision、Complete 六个 step 空态）
  - [x] SubTask 3.3: 对说明文案产出逐项承接矩阵（`design.md §2.3`：Onboarding WelcomeStep 顶部引导说明与实体介绍 `must-change` / 四个 Detail 页页面级语义导语 `must-change` / Dashboard 三个区块标题与三个区块说明文案 `no-change` / Daily 与 Weekly Review 页面头部说明 `no-change` / 四个 Detail 页标题与四个 Detail 页既有引导文案 `no-change`）
  - [x] SubTask 3.4: 对下一步动作说明产出逐项承接矩阵（`design.md §2.4`：ModuleNextActionBar `must-change` / Decision CTA `no-change` / DashboardPrimaryActionPanel `no-change` / OnboardingCtaButton `no-change` / ReviewActionFooter `no-change`）

- [x] Task 4: 产出共享呈现与切片内保留设计（写入 `design.md §3`）
  - [x] SubTask 4.1: 标记共享语义来源候选并冻结唯一共享承接位（`design.md §3.1`：四实体冻结语义单行定义与 Module / Decision / Repository 高频短语统一回收到 `frontend/src/features/project-context/`）
  - [x] SubTask 4.2: 对每项共享语义来源候选说明不越权改写 `phase12-03` 的共享只读 owner（`design.md §3.1` 边界约束）
  - [x] SubTask 4.3: 对每项切片内保留的渲染与结构说明理由，并显式区分“共享语义来源”与“切片本地渲染”（`design.md §3.2`：7 类保留对象）

- [x] Task 5: 产出可直接进入实现的表达样例（写入 `design.md §4`）
  - [x] SubTask 5.1: 至少给出一组详情页或摘要卡片的 before / after 样例（`design.md §4.1`：ModuleSummaryCard / `design.md §4.2`：DecisionDetailSummaryCard）
  - [x] SubTask 5.2: 至少给出一组空态、说明文案或下一步动作说明的 before / after 样例（`design.md §4.3`：Onboarding WelcomeStep / `design.md §4.4`：ModuleNextActionBar）
  - [x] SubTask 5.3: 为 `Module` 与 `Decision` 补出最容易漂移语义的表达样例（`design.md §4.1`：ModuleSummaryCard / `design.md §4.2`：DecisionDetailSummaryCard）

- [x] Task 6: 输出明确不做清单（写入 `design.md §5`）
  - [x] SubTask 6.1: 明确本任务不重写读路径 owner（`design.md §5` #1）
  - [x] SubTask 6.2: 明确本任务不新增后端合同、共享只读服务或写回通道（`design.md §5` #2-3）
  - [x] SubTask 6.3: 明确本任务不进行页面重组或结构性路由改造（`design.md §5` #4-8）

- [x] Task 7: 完成一致性与可执行性校验
  - [x] SubTask 7.1: 校验所有进入范围的对象都已被分类（`design.md §6`：38 个对象全部分类）
  - [x] SubTask 7.2: 校验设计结果足以回答"哪些页面必须改文案、改摘要或改说明"（`design.md §1-2` 逐对象/逐 surface 结论）
  - [x] SubTask 7.3: 校验设计结果没有引入第二套页面事实源（`design.md §3` 边界约束）
  - [x] SubTask 7.4: 校验 `no-change` 对象与 surface 都已逐项留档（`design.md §1.1/R1-R8`、`§2.3` 与 `§2.4` 中所有 `no-change` surface 全部含理由）

# Task Dependencies

- Task 2 depends on Task 1
- Task 3 depends on Task 1
- Task 4 depends on Task 1
- Task 5 depends on Task 2
- Task 5 depends on Task 3
- Task 6 depends on Task 1
- Task 7 depends on Task 2
- Task 7 depends on Task 3
- Task 7 depends on Task 4
- Task 7 depends on Task 5
- Task 7 depends on Task 6
