# phase12_semantic_alignment_and_readonly_consumption_foundation_architecture_plan

## 1. 文档定位

本文档是 `phase12_semantic_alignment_and_readonly_consumption_foundation` 的架构规划文档。

`phase11_project_context_foundation` 已完成正式 `/plan -> /spec -> 实现 -> 验收 -> 收口`，并把 `PSCO` 正式推进为可被 agent 稳定只读消费的项目上下文系统。与此同时，`audit_002` 已明确：`phase11` 对 `Product / Repository / Module / Decision` 的冻结，当前主要尚未在 Web 端形成系统化语义收口；下一步也不应直接跳到更重的 `MCP / CLI / agent 写回 / 更重消费通道`。

因此，`phase12` 的职责不是：

- 重写四实体 schema 或关系主线；
- 把 Web 做成对话式 agent 工作台；
- 直接建设 MCP / CLI / agent 写回；
- 让 Web 与 agent 各自长出第二套四实体解释。

`phase12` 的职责是交付单一主交付能力：

> **Semantic Alignment & Read-Only Consumption Foundation**

也就是：

1. 让四实体冻结语义在 Web 端形成可见、可复用、可回看的单值表达；
2. 在保持只读边界不变的前提下，继续深化 agent / Web 共享的只读消费能力；
3. 为更重消费通道或受控维护能力建立明确进入条件，而不是提前实现。

## 2. 上游输入

本阶段直接上游输入如下：

1. `AGENTS.md`
2. `plan.md`
3. `project_rules.md`
4. `TECH_STACK_BASELINE.md`
5. `architecture_map.md`
6. `docs/README.md`
7. `docs/phase/README.md`
8. `PSCO-mvp05-summarize-feedback.md`
9. `docs/audit/audit_002_phase11_post_closeout_direction_issue.md`
10. `docs/audit/audit_002_phase11_post_closeout_direction_analysis.md`
11. `docs/phase/phase11_project_context_foundation_architecture_plan.md`
12. `docs/phase/phase11_project_context_foundation_dev_plan.md`
13. `docs/phase/phase11_project_context_foundation_shared_baseline.md`
14. `.trae/specs/phase11_09_validate_project_context_foundation_dogfooding_regression/acceptance_report.md`
15. `.trae/specs/phase03_10_decision_center_formal_spec/decision_center_spec_v0.1.md`
16. `.trae/specs/phase04_10_product_repository_binding_formal_spec/product_repository_binding_spec_v0.1.md`
17. `.trae/specs/phase05_10_dashboard_feedback_formal_spec/dashboard_feedback_spec_v0.1.md`
18. `.trae/specs/phase06_12_onboarding_sovereignty_reuse_formal_spec/spec.md`

补充说明：

- `PSCO-mvp05-summarize-feedback.md` 继续作为 `mvp0.5` 的唯一直接共识上游；
- `audit_002` 负责给出 `phase11` 收口后的去向仲裁，不替代本阶段 `/plan`；
- `phase03 ~ phase06` 提供的是四实体与前端既有主线的正式规格来源，不允许在本阶段临时虚构第二套语义来源。

## 3. 本阶段目标

`phase12` 的目标是：

> 在不重写四实体结构、不突破只读边界的前提下，让 Web 与 agent 对 `Product / Repository / Module / Decision` 的理解重新收敛到同一套冻结语义，并把最小只读消费从“能读”推进到“更容易读、更稳定读、可在 Web/agent 两侧复用地读”。

本阶段需要回答的核心问题：

1. 四实体冻结语义在 Web 端的正式承接位应该是什么；
2. 哪些页面、摘要卡片、空态、说明文案与下一步动作表达必须回收语义漂移；
3. `phase11` 已交付的只读项目上下文能力，哪些部分需要继续深化到更稳定的消费入口；
4. Web 与 agent 如何共享同一套只读上下文解释，而不是在前端另写一套“页面语义”；
5. 更重的 agent 通道在什么前提下才允许进入下一阶段讨论。

## 4. 架构冻结结论

### 4.1 当前阶段唯一直接执行层上游

`phase12` 必须直接承接：

- `PSCO-mvp05-summarize-feedback.md`
- `audit_002` issue / analysis
- `phase11` 三件套与 `phase11-09` 正式验收入口
- `phase03 / phase04 / phase05 / phase06` 已冻结的前端与实体主线规格

不允许在本阶段重新解释：

- `.proto` 作为唯一长期合同源；
- `ConnectRPC` 作为正式业务传输主线；
- `phase11` 已完成“根级真相源治理 + 最小只读项目上下文导出”的事实；
- `phase11` 对四实体的冻结是“语义澄清，不做结构重构”；
- `mvp0.5` 当前仍遵守“先消费、后维护”的推进顺序。

### 4.2 当前阶段单一主交付能力

`phase12` 的单一主交付能力冻结为：

> **让四实体冻结语义在 Web 与 agent 的只读消费面上形成单值一致。**

它由两部分组成：

1. **四实体语义一致性收口**
2. **只读消费深化**

补充冻结：

- 语义一致性收口优先落在页面表达、摘要呈现、入口说明与共享读模型解释，不等于结构重构；
- 只读消费深化优先复用 `phase11` 已交付的项目上下文主线，不得借机打开写回或新协议层；
- 两部分必须放在同一交付型 phase 中推进，而不是拆成“纯前端文案整理”与“纯 agent 能力探索”两条散线。

### 4.3 四实体语义承接的正式边界

本阶段继续冻结四实体语义如下：

- `Product`：经营目标与交付容器；
- `Repository`：代码仓库身份对象与项目锚点；
- `Module`：可复用能力资产，允许后置提炼；
- `Decision`：规则、约束、选择与依据的索引对象。

本阶段正式结论：

- 四实体当前只承接**表达与消费层**的对齐，不承接 schema、关系主线与 canonical owner 重写；
- Web 端的摘要卡片、详情页说明、空态、引导文案、Review / Dashboard / Onboarding 中的解释性语言必须回收到上述单值语义；
- 若现有类型、页面命名或字段解释与上述语义存在张力，优先通过显示语义、共享摘要与只读读模型解释解决，而不是先动结构。

本阶段明确不做：

- 四实体改名重构；
- schema 重写；
- 把 `Module` 升格为新的重型能力平台对象；
- 把 `Decision` 扩写为审批流、流程引擎或 agent 写入入口。

### 4.4 只读消费深化的正式边界

`Read-Only Consumption` 在当前阶段冻结为：

- 保持 `phase11` 的 `.proto + ConnectRPC + Project Context` 正式主线；
- 优先让 Web 与 agent 共享同一套规则、约束、文档入口与四实体最小摘要解释；
- 优先提升“可定位性、可回看性、可复用性”，而不是新增长协议出口。

当前阶段进一步冻结：

- `repository_id` 继续是项目上下文聚合读取的唯一正式结构化输入锚点；
- 新增或演进后的只读能力必须继续落在 Go backend 正式只读业务接口层；
- Markdown / 结构化输出仍必须从同一只读事实源单向派生；
- 若 Web 端需要复用 `phase11` 导出的解释性结果，应优先走共享只读主线，而不是前端本地再拼一套“页面版真相源”。

当前阶段明确不做：

- MCP / CLI；
- agent 自动写回；
- 前端对话式 agent 入口；
- 第二套 canonical API；
- 以消费便利为由新增影子状态表或第二套语义字段。

### 4.4A 共享只读 owner 与承接矩阵

为避免 `phase12` 在实现时再次长出“后端一套 Project Context、前端一套页面语义摘要”的并列事实源，当前阶段额外冻结以下唯一 owner：

1. **结构化 canonical owner**
   - 唯一 owner：`proto/psco/project_context/v1/project_context.proto` 与 `backend/internal/projectcontext/*`
   - 承接职责：`Repository / Product / Module / Decision` 的项目上下文聚合只读结果、规则/约束/文档入口定位、AGENTS 风格导出
   - 禁止事项：前端不得并列定义第二套跨切片 canonical 摘要合同
2. **Web 侧跨切片共享只读 owner**
   - 唯一允许的新承接位：`frontend/src/features/project-context/`
   - 启用条件：当 `3+` 个页面/切片需要复用同一份 repository-scoped 语义摘要、规则入口或文档入口时，才允许晋升为该跨切片 owner
   - 承接职责：只读 `connect-client`、query options、受控只读 adapter、跨切片共享语义摘要与入口定位视图
   - 禁止事项：不得在 `project-context` 中承接写路径、页面私有状态或与后端 canonical 并列的新字段语义
3. **切片内展示 owner**
   - 唯一 owner：各 feature 自己的 `pages/`、`components/` 与 `data/` 目录
   - 承接职责：把共享只读结果映射为页面文案、摘要卡片、空态与说明文案
   - 禁止事项：不得在切片内重新拼装跨切片共享语义摘要；若 `3+` 页面复用同一段解释，必须回收到 `frontend/src/features/project-context/`

进一步冻结：

- `ProjectContextService.GetProjectContext` 是 Web / agent 共享语义摘要的唯一结构化事实源；
- `ProjectContextService.ExportProjectContext` 继续作为 agent-facing Markdown 导出承接位，不承担 Web 页面专属真相源职责；
- Web 端若只需要局部解释性结果，可以在 `frontend/src/features/project-context/` 做受控裁剪视图，但其字段语义必须单向派生自 `GetProjectContext` 结构化结果。

### 4.5 Web / agent 共享语义的正式分工

当前阶段继续冻结：

- PSCO 是上下文系统，不是开发流程控制器；
- Web 继续承接全局查看、关系校对、回顾、人工修正与最终确认；
- agent 当前继续承接项目现场只读消费；
- 二者共享同一套 Go backend canonical core 与同一套四实体冻结语义。

补充冻结：

- Web 不退化为聊天工作台；
- agent 不提前进入维护主线；
- 不允许在 Web 与 agent 两侧分别定义“Module / Decision 到底是什么”的不同版本。

### 4.6 当前阶段前端承接策略

当前阶段前端继续统一遵守：

- 单一 `React Web`
- `TanStack Router + TanStack Query`
- 业务写路径唯一 `application` 入口
- `query` 层纯只读
- mutation 固定承接位

当前阶段重点是：

- 回收前端四实体页面与摘要组件中的语义漂移；
- 把四实体的解释性语言尽可能收敛到稳定承接位，而不是散落在多个页面中各讲一套；
- 在不推翻既有切片边界的前提下，让 `Dashboard / Review / Onboarding / Detail pages` 的四实体解释重新对齐。

当前阶段前端正式影响面进一步冻结为以下 route / page / component owner：

- 路由层：
  - `frontend/src/routes/dashboard.tsx`
  - `frontend/src/routes/onboarding.tsx`
  - `frontend/src/routes/reviews/daily.tsx`
  - `frontend/src/routes/reviews/weekly.tsx`
  - `frontend/src/routes/products/$productId.tsx`
  - `frontend/src/routes/repositories/$repositoryId.tsx`
  - `frontend/src/routes/modules/$moduleId.tsx`
  - `frontend/src/routes/decisions/$decisionId.tsx`
- 页面层：
  - `frontend/src/features/dashboard/pages/dashboard-home-page.tsx`
  - `frontend/src/features/onboarding/pages/onboarding-page.tsx`
  - `frontend/src/features/review/pages/daily-review-page.tsx`
  - `frontend/src/features/review/pages/weekly-review-page.tsx`
  - `frontend/src/features/product-registry/pages/product-detail-page.tsx`
  - `frontend/src/features/repository-binding/pages/repository-binding-detail-page.tsx`
  - `frontend/src/features/module-registry/pages/module-detail-page.tsx`
  - `frontend/src/features/decision-center/pages/decision-detail-page.tsx`
- 摘要/说明组件层：
  - `frontend/src/features/product-registry/components/product-summary-card.tsx`
  - `frontend/src/features/repository-binding/components/repository-summary-card.tsx`
  - `frontend/src/features/module-registry/components/module-summary-card.tsx`
  - `frontend/src/features/decision-center/components/decision-detail-summary-card.tsx`
  - `frontend/src/features/dashboard/components/current-focus-section.tsx`
  - `frontend/src/features/dashboard/components/recent-activity-section.tsx`
  - `frontend/src/features/review/components/review-page-shell.tsx`
- 读路径 owner：
  - `frontend/src/features/*/data/use-*-read.ts`
  - 若产生跨切片共享只读语义，则唯一 owner 为 `frontend/src/features/project-context/`

补充冻结：

- `List pages`、toolbar 与搜索 store 默认只做跟随回归，不作为语义冻结的 primary owner；
- 当前阶段必须至少显式审计以上 route / page / component 清单，允许有“无需改动”的结果，但必须在后续 `/spec` 与验收记录中逐项留档。

### 4.6A repository-scoped 消费边界矩阵

为避免 `repository_id` 作为唯一结构化输入锚点时，被后续实现错误外推到所有页面，当前阶段冻结以下消费边界：

1. **直接 repository-scoped 页面**
   - `frontend/src/routes/repositories/$repositoryId.tsx`
   - 允许直接消费 `GetProjectContext(repository_id)`
2. **间接 repository-scoped 详情页**
   - `products/$productId`
   - `modules/$moduleId`
   - `decisions/$decisionId`
   - 允许通过既有 detail read 先解析关联 `repository_id`，再复用共享只读语义摘要；不得直接在页面内临时猜测 repository 锚点
3. **衍生消费页**
   - `dashboard`
   - `onboarding`
   - `reviews/daily`
   - `reviews/weekly`
   - 只允许消费受控共享摘要、固定入口链接或派生只读解释；不得把这些页面伪装成新的结构化输入锚点入口

补充冻结：

- 若某页面无法稳定解析回同一 `repository_id` 或其派生共享摘要，则该页面不得被纳入本阶段的语义一致性 primary owner；
- 当前阶段任何新增 Web 读路径都必须先回答：它是“直接 repository-scoped”“间接 repository-scoped”还是“衍生消费页”，否则不得进入实现。
- 对 `products/$productId`、`modules/$moduleId`、`decisions/$decisionId` 的 id 解析，唯一允许的上游是同一 `repository_id` 驱动的结构化只读结果或其受控派生视图；不得通过页面内临场搜索、额外脚本或手工数据库查询补齐。

### 4.6B 设计推进顺序冻结

为避免 `phase12` 在 `/spec` 前再次出现“共享 owner 先定还是字段 shape 先定”的往返修改，当前阶段额外冻结以下实现设计顺序：

1. 先由 `phase12-05` 冻结：
   - Web / agent 共享只读 owner
   - 页面承接分类
   - 最小共享摘要需求
   - repository-scoped resolver 需求
2. 再由 `phase12-07` 冻结：
   - 后端合同与共享只读视图是否继续复用 `GetProjectContext`
   - 若需新增受控派生读取，其最小字段、定位能力与回收的重复解释逻辑
3. 最后由 `phase12-06` 冻结：
   - 前端读路径 owner
   - `frontend/src/features/project-context/` 的 adapter / query options / 共享摘要裁剪视图
   - 各衍生消费页的 reread 与回流关系

补充冻结：

- `phase12-06` 不得反向改写 `phase12-05` 已冻结的 owner 分类；
- `phase12-06` 也不得反向要求 `phase12-07` 在无明确重复解释回收收益的情况下新增第二服务或第二事实源；
- 若 `phase12-07` 判定现有 `GetProjectContext` 已足够承接，则 `phase12-06` 必须直接以前者结果为准，而不是在前端再长一套影子摘要合同。

### 4.7 当前阶段后端、合同与读模型承接策略

当前阶段继续统一遵守：

- `.proto` 是唯一长期合同源
- `ConnectRPC` 是业务接口正式传输层
- `chi` 继续只承接 router shell、middleware 与非业务端点
- `phase11` 的项目上下文聚合读取与导出是当前只读消费深化的正式基础

当前阶段重点是：

- 识别哪些只读上下文字段、入口定位或导出结果可以被 Web / agent 同时稳定消费；
- 优先演进共享读模型、共享解释字段或受控导出视图，而不是扩张新协议；
- 任何新增只读承接位都必须显式回答“它回收了哪段前端 / agent 重复解释逻辑”，否则不得进入实现。

当前阶段合同与读模型冻结补充为：

- 若后端需要为 Web 端提供受控共享摘要，优先在现有 `GetProjectContext` 结构化结果内追加最小字段，而不是新增第二服务；
- 若确需新增只读视图，其唯一合法身份是 `ProjectContextService` 下的受控派生读取，不得长成并列业务域；
- `RuleEntry / PhaseEntry / BoundaryEntry` 与四实体最小摘要继续作为 Web / agent 共享解释的优先来源；
- 任何面向 Web 的共享只读 adapter 都必须保留到 `entry_ref / entry_kind / repository_id` 或其单向派生定位关系，避免前端丢失可定位性。
- 任何样本 id 解析或固定入口定位方案，都必须能被验收记录逐项复现；若无法在验收记录中写清解析入口、解析结果与失败条件，则不得进入实现。

### 4.8 当前阶段业务边界原则

为避免 `phase12` 再次滑向“大重构 + 大通道 + 大而泛 agent 平台”，当前阶段先冻结以下边界：

- 本阶段只交付 `Semantic Alignment & Read-Only Consumption Foundation`
- 更重的 `MCP / CLI / agent 写回 / 更重消费通道 / 受控维护能力` 只允许作为后续进入条件表达
- 当前阶段不重开一轮广义专家原则争论；若需要新增外部评审，只允许围绕已压缩候选方向展开
- 当前阶段不回头推翻 `phase03 ~ phase06` 已完成的实体、页面与写路径主线

## 5. 当前阶段完成条件

`phase12` 完成时，至少必须满足：

1. 四实体冻结语义已在 Web 端形成稳定、可回看的正式表达；
2. 至少在 `repositories/$repositoryId`、`products/$productId`、`modules/$moduleId`、`decisions/$decisionId` 四类详情页中，可直接验证 `Module = 可复用能力资产`、`Decision = 规则/约束/选择与依据索引对象` 的正式表达；
3. `dashboard`、`onboarding`、`reviews/daily`、`reviews/weekly` 已能通过共享语义摘要或固定入口解释当前四实体角色，而不是继续依赖各自切片内旧文案；
4. `phase11` 已交付的只读项目上下文能力已被深化为更稳定、可复用、可定位的消费入口，且 Web 侧跨切片共享读模型 owner 单值冻结；
5. Web 与 agent 对四实体、规则、约束与文档入口的理解已能通过固定入口集合回答同一组验收问题；
6. 本阶段未引入 MCP / CLI / agent 写回 / 前端对话入口 / schema 重写 / 第二套 canonical API；
7. 根级入口已回写当前正式 phase 为 `phase12_semantic_alignment_and_readonly_consumption_foundation`。
