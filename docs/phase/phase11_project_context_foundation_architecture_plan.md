# phase11_project_context_foundation_architecture_plan

## 1. 文档定位

本文档是 `phase11_project_context_foundation` 的架构规划文档。

`phase10_asset_action_closure_foundation` 已完成正式 `/plan -> /spec -> 实现 -> 验收 -> 收口`，`PSCO-mvp05-summarize-feedback.md` 已进一步冻结：`mvp0.5` 的中心任务不再是继续闭动作链，而是让 PSCO 以 Go 后端为中心，先成为一个能被 agent 稳定读取、稳定理解、低成本接手的项目上下文系统。

因此，`phase11` 的职责不是：

- 重做四实体结构；
- 直接建设 MCP / CLI / agent 写回；
- 把 PSCO 做成 IDE 现场的流程控制器；
- 把 web 做成对话式 agent 工作台。

`phase11` 的职责是交付单一主交付能力：

> **Project Context Foundation**

也就是：

1. 先治好根级上下文真相源的重复承载、引用漂移与悬空入口；
2. 再提供最小只读“项目上下文聚合导出”能力；
3. 让 agent 可以通过单一稳定入口读取当前项目的完整核心上下文。

补充边界澄清：

- 本阶段同时包含两层对象，但二者不可混写：
  1. `PSCO` 自身仓库的根级上下文真相源治理；
  2. `PSCO` 作为能力基座面向未来不同项目提供的通用只读项目上下文能力。
- 第一层允许也必须基于 `PSCO` 当前仓库真实存在的根级入口文件落地；
- 第二层不得把 `PSCO` 当前仓库的目录结构、固定文件名或入口清单外推为所有未来项目的必备前提。

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
9. `docs/review/PSCO-mvp045-GPT54.md`
10. `docs/review/PSCO-mvp045-DPv4pro.md`
11. `docs/review/PSCO-mvp045-gemini31pro.md`
12. `docs/review/PSCO-mvp045-GLM52.md`
13. `docs/review/PSCO-mvp045-qwen38max.md`
14. `docs/phase/phase10_asset_action_closure_foundation_architecture_plan.md`
15. `docs/phase/phase10_asset_action_closure_foundation_dev_plan.md`
16. `docs/phase/phase10_asset_action_closure_foundation_shared_baseline.md`
17. `.trae/specs/phase10_11_complete_asset_action_closure_integration_browser_regression_validation/spec.md`

补充说明：

- `PSCO-mvp05-summarize-feedback.md` 是本阶段唯一直接共识上游；
- 五份 `mvp0.45` 评审不再作为并列正式规则源，而作为 `mvp0.5` 共识的证据上游；
- `phase10` 提供的是最近完成正式业务 phase 的规划记录与验收结论，不允许在本阶段回头重写 `Asset-Action Closure` 主线。

## 3. 本阶段目标

`phase11` 的目标是：

> 在不引入第二套 canonical API、不建设 agent 写入路径、不启动重型协议层的前提下，把 PSCO 从“已闭合的经营动作系统”推进为“可被 agent 稳定读取的项目上下文系统”。

本阶段需要回答的核心问题：

1. 根级静态文档中，哪些结论必须只保留一个正式承接位；
2. 当前根级入口如何统一直接指向 `PSCO-mvp05-summarize-feedback.md`；
3. agent 当前真正需要的“项目上下文”最小集合是什么；
4. 当前 12 个既有 `.proto + ConnectRPC` 合同之上，最小的聚合只读投影应如何承接；
5. AGENTS 风格 Markdown 导出与结构化只读读取的边界应如何划分；
6. 如何用 PSCO 仓库自身作为第一 dogfooding 场景，验证上下文恢复成本是否真实下降。

## 4. 架构冻结结论

### 4.1 当前阶段唯一直接执行层上游

`phase11` 必须直接承接：

- `PSCO-mvp05-summarize-feedback.md`
- `project_rules.md`
- `TECH_STACK_BASELINE.md`
- `plan.md`
- `AGENTS.md`
- `architecture_map.md`
- `docs/README.md`
- `docs/phase/README.md`
- `phase10` 三件套与 `phase10-11` 验收入口

不允许在本阶段重新解释：

- `.proto` 作为唯一长期合同源；
- `ConnectRPC` 作为正式业务传输主线；
- PSCO 的技术栈与 Durable System Track 冻结选择；
- `phase10` 已完成正式收口这一事实；
- `mvp0.5` 已冻结的方向：根级真相源治理优先、最小只读项目上下文导出为中心主交付。

### 4.2 当前阶段单一主交付能力

`phase11` 的单一主交付能力冻结为：

> **让 agent 通过单一稳定入口读取当前项目的完整核心上下文。**

它由两部分组成：

1. **根级上下文真相源治理**
2. **最小只读项目上下文导出**

补充冻结：

- 上下文导出是聚合投影，不是新实体主线；
- 根级治理是前置能力，不是附属清理项；
- 二者必须放在同一交付型 phase 内，而不是拆成纯讨论 phase + 交付 phase。

### 4.3 根级上下文真相源治理的正式边界

本阶段的根级治理冻结为：

- `plan.md` 继续作为阶段状态与路线的唯一正式承接位；
- `architecture_map.md` 继续作为目录结构与文档落点的唯一正式承接位；
- `TECH_STACK_BASELINE.md` 继续作为技术栈正文的唯一正式承接位；
- `README.md` 继续作为项目总览入口，只保留总览与受控跳转，不重复承载当前 phase 正文或最终共识正文；
- `AGENTS.md` 继续作为入口摘要，而不是重复承载完整 phase 状态正文；
- `project_rules.md` 继续作为项目级协作规则与单一真相源约束承接位，不重复承载当前 phase 状态、最终共识或目录落点正文；
- `global_skills.md` 继续作为项目内通用方法映射说明，不承接当前 phase 状态、最终共识或目录落点正文；
- `docs/README.md` 继续作为 workflow 总入口，而不是重复承载完整目录落点正文。
- `docs/phase/README.md` 继续作为 phase 文档入口索引，而不是重复承载根级治理策略、最终共识或目录落点正文。

补充冻结：

- 本节治理对象仅针对 `PSCO` 自身仓库的根级/入口文档；
- 本节列出的文件清单不构成未来被 `PSCO` 承载项目的通用目录模板；
- 不允许把 `PSCO` 当前仓库的根级文件结构，偷渡成“所有项目必须长这样”的隐含合同。
- 本子任务的正式设计产物至少应包括：
  - 根级入口治理矩阵；
  - 重复承载清单与目标落点清单；
  - 悬空引用清理清单；
  - 收口后的单一写者规则表。

本阶段明确不做：

- 静态文件从 backend 全量派生；
- 在根级继续保留同一主结论的重复承载；
- 继续保留指向不存在文件的引用；
- 通过新增新文件绕开既有单一真相源规则。

当前阶段根级共识入口冻结为：

- **直接以 `PSCO-mvp05-summarize-feedback.md` 作为根级最终共识文档**
- 当前阶段的治理路线冻结为：先完成根级入口的一次性校准与正文收敛，不在本阶段引入静态文件全量派生机制。

### 4.4 项目上下文导出的正式边界

`Project Context` 在当前阶段冻结为：

- 基于当前 canonical 数据的聚合只读结果；
- 围绕“当前项目/仓库”构造；
- 面向 agent 读取，不面向页面展示分页；
- 输出“当前项目最小必需上下文”，而不是全库扫描或知识图谱。

当前阶段进一步冻结以下输入与承接合同：

- 结构化只读读取的唯一输入锚点为既有 `Repository` canonical 身份；
- `repository_id` 是当前阶段唯一正式结构化输入，不引入 `product_id`、本地路径、Git remote URL 或工作区扫描作为并列主锚点；
- 当前阶段只承接“已在 PSCO 中完成登记/绑定的仓库”上下文读取；
- 若当前仓库尚未完成 `Repository Binding`，其身份属于当前阶段明确失败态，而不是由执行者临场猜测补锚；
- 最小结构化读取的正式承接位应落在 Go backend 只读业务接口，并继续遵守 `.proto + ConnectRPC` 主线；
- AGENTS 风格 Markdown 导出必须从同一结构化只读结果单向派生，不得形成第二套事实源。

补充冻结：

- 当前阶段的通用项目上下文能力，以 `PSCO` 中已登记的 canonical 实体关系为准，而不是以消费侧项目目录中是否存在特定文件名为准；
- 除 `repository_id` 与其在 `PSCO` 中已绑定的实体关系外，当前阶段不把消费侧项目目录结构当作必要输入合同；
- 若未来存在“最佳实践项目模板”或“项目约定资产”，其身份只能是增强型 convention/profile，不是当前 phase11 的前置前提；
- 当前阶段不要求未来所有项目都必须通过统一模板 `git clone / pull` 才能被 `PSCO` 消费。

当前阶段至少应覆盖的上下文维度：

- 当前项目对应的 `Repository`；
- 关联的 `Product` 摘要；
- 关联的 `Module` 摘要与状态；
- 关联的 `Decision` 摘要与状态（过滤或弱化已归档终态）；
- 关键规则、约束与文档入口；
- 与当前 phase 直接相关的 spec / baseline / 根级入口。

`Decision` 聚合边界进一步冻结为：

- 以当前 `Repository` 为根，只合并三类直接 canonical 关系命中的 `Decision`：
  - 直接链接到当前 `Repository` 的 `Decision`
  - 直接链接到“当前 `Repository` 已绑定 `Product`”的 `Decision`
  - 直接链接到“当前 `Repository` 已映射 `Module`”的 `Decision`
- 当前阶段不得继续沿 `Product -> Module -> 其他 Repository` 做递归扩张
- 同一 `Decision` 若同时命中多类关系，必须以 `decision_id` 去重，并保留命中来源摘要
- 当前阶段结构化只读主列表只承接非 `archived` 的 `Decision`

当前阶段明确不做：

- agent 写入建议的正式提交路径；
- MCP / CLI 协议层；
- 前端对话式入口；
- 重型 GitHub / Gitea 集成；
- 仓库外主动注入文件；
- 任何 agent 专属一级业务对象。

### 4.5 Web 与 agent 的正式分工

当前阶段继续冻结：

- PSCO 是上下文系统，不是开发流程控制器；
- IDE / agent 负责项目内微观推进；
- web 继续承接全局查看、关系校对、回顾、人工修正与最终确认；
- agent 当前只承接现场上下文消费；
- 二者共享同一套 Go backend canonical core。

补充冻结：

- web 不退化；
- agent 不对称并行进入；
- 不允许 web 与 agent 各自长出第二套领域语义或第二套流程。

### 4.6 四实体语义冻结

当前阶段不重写 `Product / Repository / Module / Decision` 结构，只冻结语义口径：

- `Product`：经营目标与交付容器；
- `Repository`：代码仓库身份对象与项目锚点；
- `Module`：可复用能力资产，允许后置提炼；
- `Decision`：规则、约束、选择与依据的索引对象。

补充冻结：

- `Module` 当前代表可复用能力资产，允许在后续真实复用沉淀中继续提炼，当前阶段不要求重写其 schema、层级或注册主线；
- `Decision` 当前代表规则、约束、选择与依据的索引对象，用于支撑项目上下文恢复与只读导出，不在本阶段扩写为审批流、流程引擎或结构重构入口；
- 这组语义用于后续 `/spec` 与实现设计边界说明；
- 不允许把上述语义口径偷渡成 schema 重构或实体关系重写。

## 5. 当前阶段完成条件

`phase11` 完成时，至少必须满足：

1. 根级入口文档不再重复承载 phase 状态、目录落点或技术栈正文；
2. 根级入口中不再存在指向不存在文件 `PSCO-summarize-feedback.md` 的引用；
3. `PSCO-mvp05-summarize-feedback.md` 已成为根级最终共识的单值入口；
4. 已存在一个最小只读“项目上下文聚合导出”正式承接方案；
5. 已存在一个 AGENTS 风格或等价 Markdown 风格的正式导出承接方案；
6. `repository_id` 已被冻结为当前阶段唯一正式结构化输入锚点，未绑定仓库失败语义已冻结；
7. 本阶段未引入 agent 写入路径、MCP、CLI、第二套 canonical API 或 agent 专属一级业务对象；
8. 已明确区分“PSCO 自身仓库治理”与“面向未来项目的通用上下文能力”，未把当前仓库目录外推为所有项目模板；
9. 以 PSCO 仓库自身为第一 dogfooding 场景的验收路径与测量协议已冻结。
