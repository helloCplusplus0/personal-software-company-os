# phase12-11 acceptance report

## 1. 文档定位

本报告承接 `phase12-11` 的全量验收收口，覆盖：

1. 固定样本、固定入口与同一 `repository_id` 锚点下的样本解析；
2. 正式工具链重跑；
3. shared-readonly / agent 固定三入口与 fixed `6` 问取证；
4. Web 侧 primary owner / 跟随回归页面的浏览器验收；
5. 本阶段明确不做 `schema` 重写、`MCP / CLI / agent 写回 / 对话入口` 的边界证据；
6. 浏览器回归期间出现的样本漂移根因、恢复路径与最终 rerun 结果。

本报告不使用数据库手工查询、不引入额外第 `7` 个入口，样本恢复也通过仓库内正式脚本执行。

## 2. 固定入口、固定样本与解析协议

### 2.1 固定 agent 入口集合

本轮只使用以下固定入口：

1. [AGENTS.md](file:///home/dell/Projects/personal-software-company-os/AGENTS.md)
2. [PSCO-mvp05-summarize-feedback.md](file:///home/dell/Projects/personal-software-company-os/PSCO-mvp05-summarize-feedback.md)
3. 同一 `repository_id` 驱动的共享只读结果：
   - 结构化读取：`GetProjectContext`
   - 受控派生视图：`ExportProjectContext`

### 2.2 固定样本

- `Repository`: `personal-software-company-os`
- `repository_id`: `ca261521-8daf-4248-8f12-43525326e759`
- `Product`: `PSCO`
- `Module`: `project-context-foundation`
- `Decision`: `phase11 Project Context Foundation dogfooding 验收决策`

### 2.3 实际解析结果

本轮通过同一 `repository_id` 一次解析成功得到：

| 实体 | 名称 | 解析出的 ID | 入口 |
| --- | --- | --- | --- |
| Repository | `personal-software-company-os` | `ca261521-8daf-4248-8f12-43525326e759` | `GetProjectContext` / `ExportProjectContext` |
| Product | `PSCO` | `f0d034cc-3235-4d03-879d-6a3111b95b6b` | `GetProjectContext.product` / `ExportProjectContext > Product` |
| Module | `project-context-foundation` | `9b02e0ca-3175-4b5a-bcf4-23cecbb06f72` | `GetProjectContext.modules[0]` / `ExportProjectContext > Modules` |
| Decision | `phase11 Project Context Foundation dogfooding 验收决策` | `aa8ee5ad-b224-4ea1-b393-7e4c30e42212` | `GetProjectContext.decisions[0]` / `ExportProjectContext > Decisions` |

结论：本轮样本解析满足“同一 `repository_id` 锚点驱动、一次成功、未引入额外第 `7` 个入口”。

### 2.4 Toolchain / Sample Resolution

本轮按用户要求对 `phase12-11` 执行正式工具链重跑，不修改源码，也不改 `tasks.md / checklist.md`。执行顺序严格固定为：

1. `buf` 相关正式验证；
2. `backend/` 下 `go test ./...`；
3. `frontend/` 下 `npm run build`。

其中，`buf.yaml` 当前已显式配置 `lint.use: STANDARD`，因此若只跑 `buf build`，只能覆盖编译级合同校验，不能覆盖 lint 级 API 规则；本轮将 `buf lint` 一并纳入第 `1` 步，作为完整 protobuf 合同验证。

本轮实际执行与结果如下：

| 步骤 | 执行目录 | 命令 | 结果 | warning / 非阻断输出 | 失败归类 |
| --- | --- | --- | --- | --- | --- |
| 1a | `proto/` | `buf build` | 通过 | 无输出 | 无 |
| 1b | `proto/` | `buf lint` | 通过 | 无输出 | 无 |
| 2 | `backend/` | `go test ./...` | 通过 | 多个 package 输出 `[no test files]`；已有测试包出现 `(cached)`，均为 Go 常规非阻断信息 | 无 |
| 3 | `frontend/` | `npm run build` | 通过 | `vite build` 输出产物体积、gzip 体积与 chunk 列表，无 warning；构建完成于 `2.09s` | 无 |

结论：本轮 `phase12-11` 请求范围内的正式工具链验证已全通过，当前未发现合同、后端、前端或环境类阻断点。

## 3. shared-readonly 实际结果摘要

### 3.1 结构化读取摘要

本轮实时 `GetProjectContext` 返回以下关键事实：

- `repository`: `personal-software-company-os`
- `product`: `PSCO`
- `modules`: `project-context-foundation`
- `decisions`: `phase11 Project Context Foundation dogfooding 验收决策`
- `rules`:
  - `Product = 经营目标与交付容器`
  - `Repository = 代码仓库身份对象与项目锚点`
  - `project_rules.md`
  - `plan.md`
  - `architecture_map.md`
  - `docs/README.md`
  - `TECH_STACK_BASELINE.md`
- `phases`:
  - `phase12_semantic_alignment_and_readonly_consumption_foundation`
  - `phase11_project_context_foundation`
  - `phase10_asset_action_closure_foundation`
- `boundaries`:
  - 不做 `MCP / CLI / agent 自动写回`
  - 不把 PSCO 做成开发流程控制器
  - 不把 web 退化为对话式 agent 工作台
  - 不要求消费侧项目复制 PSCO 目录结构
  - 不形成第二套事实源

### 3.2 Markdown 导出摘要

本轮实时 `ExportProjectContext` 返回以下受控派生视图结构：

1. `Repository`
2. `Product`
3. `Modules`
4. `Decisions`
5. `Rules & Constraints`
6. `Current Phase`
7. `Boundaries (What This Project Does NOT Do)`

其中 `Rules & Constraints` 已显式带出：

- `Product = 经营目标与交付容器`
- `Repository = 代码仓库身份对象与项目锚点`
- `plan.md`
- `architecture_map.md`
- `docs/README.md`
- `project_rules.md`
- `TECH_STACK_BASELINE.md`

其中 `Current Phase` 已显式带出 `phase12` 当前入口。

结论：agent 侧固定三入口中的第三入口已经具备回答四实体样本、规则/约束入口、当前 phase 入口与边界摘要的最低信息面。

## 4. 固定 6 问逐题验证

回答格式冻结为：`answer / direct entry refs / 是否达标`。

| 题目 | answer | direct entry refs | 是否达标 |
| --- | --- | --- | --- |
| 1. 当前 `Product` 的正式定位是什么，在哪个固定入口可直接确认 | 当前 `Product` 的正式定位是“经营目标与交付容器”。固定三入口里，可由 shared-readonly 的 `Rules & Constraints` 直接确认该语义，再由 `GetProjectContext.product` / `ExportProjectContext > Product` 把该语义锚到当前样本 `PSCO`。 | 本报告 [2.3](file:///home/dell/Projects/personal-software-company-os/.trae/specs/phase12_11_validate_semantic_alignment_readonly_consumption_foundation/acceptance_report.md#L35-L46)、[3.1](file:///home/dell/Projects/personal-software-company-os/.trae/specs/phase12_11_validate_semantic_alignment_readonly_consumption_foundation/acceptance_report.md#L72-L98)、[3.2](file:///home/dell/Projects/personal-software-company-os/.trae/specs/phase12_11_validate_semantic_alignment_readonly_consumption_foundation/acceptance_report.md#L99-L123)、[AGENTS.md:L14-L18](file:///home/dell/Projects/personal-software-company-os/AGENTS.md#L14-L18) | 达标 |
| 2. 当前 `Repository` 的正式定位是什么，在哪个固定入口可直接确认 | 当前 `Repository` 的正式定位是“代码仓库身份对象与项目锚点”。固定三入口里，可由 shared-readonly 的 `Rules & Constraints` 直接确认该语义，再由 `GetProjectContext.repository` / `ExportProjectContext > Repository` 把该语义锚到当前样本 `personal-software-company-os`。 | 本报告 [2.3](file:///home/dell/Projects/personal-software-company-os/.trae/specs/phase12_11_validate_semantic_alignment_readonly_consumption_foundation/acceptance_report.md#L35-L46)、[3.1](file:///home/dell/Projects/personal-software-company-os/.trae/specs/phase12_11_validate_semantic_alignment_readonly_consumption_foundation/acceptance_report.md#L72-L98)、[3.2](file:///home/dell/Projects/personal-software-company-os/.trae/specs/phase12_11_validate_semantic_alignment_readonly_consumption_foundation/acceptance_report.md#L99-L123) | 达标 |
| 3. 当前 `Module` 为什么不是普通模块登记对象，在哪个固定入口可直接确认“可复用能力资产”语义 | `Module` 当前应被理解为“可复用能力资产”，不是普通模块登记对象。该语义可由固定共识上游直接确认，shared-readonly 再把该语义锚到当前样本 `project-context-foundation`。 | [PSCO-mvp05-summarize-feedback.md:L111-L117](file:///home/dell/Projects/personal-software-company-os/PSCO-mvp05-summarize-feedback.md#L111-L117)、本报告 [2.3](file:///home/dell/Projects/personal-software-company-os/.trae/specs/phase12_11_validate_semantic_alignment_readonly_consumption_foundation/acceptance_report.md#L35-L46)、[3.2](file:///home/dell/Projects/personal-software-company-os/.trae/specs/phase12_11_validate_semantic_alignment_readonly_consumption_foundation/acceptance_report.md#L99-L123) | 达标 |
| 4. 当前 `Decision` 为什么不是孤立文本卡片，在哪个固定入口可直接确认“规则 / 约束 / 选择与依据索引对象”语义 | `Decision` 当前应被理解为“规则 / 约束 / 选择与依据的索引对象”，不是孤立文本卡片。固定共识上游已明文给出该定位；shared-readonly 中 `hitSources` 同时命中 `bound_product_module` 与 `repository_module_mapping`，说明它也不是孤立脱链文本。 | [PSCO-mvp05-summarize-feedback.md:L111-L117](file:///home/dell/Projects/personal-software-company-os/PSCO-mvp05-summarize-feedback.md#L111-L117)、[AGENTS.md:L26-L27](file:///home/dell/Projects/personal-software-company-os/AGENTS.md#L26-L27)、本报告 [3.1](file:///home/dell/Projects/personal-software-company-os/.trae/specs/phase12_11_validate_semantic_alignment_readonly_consumption_foundation/acceptance_report.md#L72-L98) | 达标 |
| 5. 当前项目共享的规则、约束与文档入口从哪里查看，Web 与 agent 是否回到同一组入口 | 当前共享的规则、约束与文档入口可从 `AGENTS.md` 的“当前唯一上游”回到 `plan.md / architecture_map.md / project_rules.md / TECH_STACK_BASELINE.md`，并通过 shared-readonly 的 `Rules & Constraints / Current Phase / Boundaries` 直接看到 `plan.md / architecture_map.md / docs/README.md / project_rules.md / TECH_STACK_BASELINE.md` 以及 `phase12` 当前入口。基于固定三入口，当前已能判定 Web 与 agent 回到同一组 canonical 入口。 | [AGENTS.md:L16-L22](file:///home/dell/Projects/personal-software-company-os/AGENTS.md#L16-L22)、[PSCO-mvp05-summarize-feedback.md:L54-L72](file:///home/dell/Projects/personal-software-company-os/PSCO-mvp05-summarize-feedback.md#L54-L72)、本报告 [3.1](file:///home/dell/Projects/personal-software-company-os/.trae/specs/phase12_11_validate_semantic_alignment_readonly_consumption_foundation/acceptance_report.md#L72-L98)、[3.2](file:///home/dell/Projects/personal-software-company-os/.trae/specs/phase12_11_validate_semantic_alignment_readonly_consumption_foundation/acceptance_report.md#L99-L123) | 达标 |

汇总结论：固定 `6` 问当前结果为 `6 / 6` 达标。`Q1 / Q2 / Q5` 已由“不达标”转为“达标”，当前 fixed 三入口已能直接回答 `Product / Repository` 语义，并显式回到 `phase12` 当前规则、约束与文档入口集。

## 5. Module / Decision 误读复核

### 5.1 Module

结论：按固定三入口联合复核，`Module` 当前**不应再被判为“普通模块登记对象”**。

直接证据：

1. [PSCO-mvp05-summarize-feedback.md:L115-L116](file:///home/dell/Projects/personal-software-company-os/PSCO-mvp05-summarize-feedback.md#L115-L116) 已直接冻结 `Module` 更适合作为“可复用能力资产”理解；
2. 本报告 [2.3](file:///home/dell/Projects/personal-software-company-os/.trae/specs/phase12_11_validate_semantic_alignment_readonly_consumption_foundation/acceptance_report.md#L35-L46) 已把该语义锚到当前样本 `project-context-foundation`；
3. 若去掉 `PSCO-mvp05-summarize-feedback.md`，单靠 shared-readonly 当前仍只会显示通用 `Modules` 列表，语义纠偏能力会下降。

### 5.2 Decision

结论：按固定三入口联合复核，`Decision` 当前**不应再被判为“孤立文本卡片”**。

直接证据：

1. [PSCO-mvp05-summarize-feedback.md:L115-L116](file:///home/dell/Projects/personal-software-company-os/PSCO-mvp05-summarize-feedback.md#L115-L116) 已直接冻结 `Decision` 更适合作为“规则、约束、选择与依据的索引对象”；
2. [AGENTS.md:L26-L27](file:///home/dell/Projects/personal-software-company-os/AGENTS.md#L26-L27) 明确 `Decision` 必须进入 MVP；
3. 本报告 [3.1](file:///home/dell/Projects/personal-software-company-os/.trae/specs/phase12_11_validate_semantic_alignment_readonly_consumption_foundation/acceptance_report.md#L72-L98) 记录的实时 `hitSources` 同时命中 `bound_product_module` 与 `repository_module_mapping`，说明它挂在实体关系链上而不是孤立文本。

## 6. Boundary Evidence

### 6.1 不做 schema 重写

上游冻结证据：

1. [PSCO-mvp05-summarize-feedback.md:L111-L117](file:///home/dell/Projects/personal-software-company-os/PSCO-mvp05-summarize-feedback.md#L111-L117)：四实体需要语义澄清，不做结构重构；
2. [PSCO-mvp05-summarize-feedback.md:L196-L199](file:///home/dell/Projects/personal-software-company-os/PSCO-mvp05-summarize-feedback.md#L196-L199)：最小交付物是只读聚合投影，不是新一轮实体重构；
3. [PSCO-mvp05-summarize-feedback.md:L246-L255](file:///home/dell/Projects/personal-software-company-os/PSCO-mvp05-summarize-feedback.md#L246-L255)：当前阶段明确不做四实体结构重构或大规模 schema 扩张；
4. [AGENTS.md:L91-L93](file:///home/dell/Projects/personal-software-company-os/AGENTS.md#L91-L93)：`phase12` 当前只承接四实体语义一致性收口与共享只读消费深化，不提前混入 schema 重写。

### 6.2 不做 MCP / CLI / agent 写回 / 对话入口

上游冻结证据：

1. [AGENTS.md:L11-L12](file:///home/dell/Projects/personal-software-company-os/AGENTS.md#L11-L12) 与 [AGENTS.md:L91-L93](file:///home/dell/Projects/personal-software-company-os/AGENTS.md#L91-L93)；
2. [PSCO-mvp05-summarize-feedback.md:L96-L109](file:///home/dell/Projects/personal-software-company-os/PSCO-mvp05-summarize-feedback.md#L96-L109)；
3. [PSCO-mvp05-summarize-feedback.md:L246-L251](file:///home/dell/Projects/personal-software-company-os/PSCO-mvp05-summarize-feedback.md#L246-L251)；
4. 本报告 [3.1](file:///home/dell/Projects/personal-software-company-os/.trae/specs/phase12_11_validate_semantic_alignment_readonly_consumption_foundation/acceptance_report.md#L72-L98) 记录的实时 `boundaries`。

### 6.3 不允许额外第 7 个入口

本轮实际只使用了三类固定入口，未引入：

- 数据库查询；
- 额外脚本；
- 新增文档入口；
- 浏览器临场搜索。

因此本轮证据符合 [phase12-11 spec:L68-L82](file:///home/dell/Projects/personal-software-company-os/.trae/specs/phase12_11_validate_semantic_alignment_readonly_consumption_foundation/spec.md#L68-L82) 的失败判定约束。

## 7. Web / agent 同组入口证据点

以下证据说明 Web 与 agent 的正式方向应回到同一组规则、约束与文档入口，而不是各自维护双轨答案：

1. [phase12-11 spec:L123-L136](file:///home/dell/Projects/personal-software-company-os/.trae/specs/phase12_11_validate_semantic_alignment_readonly_consumption_foundation/spec.md#L123-L136) 明确要求 Web / agent dogfooding 回到同一组规则、约束与入口；
2. [phase12-09 spec:L78-L97](file:///home/dell/Projects/personal-software-company-os/.trae/specs/phase12_09_land_readonly_consumption_shared_entry/spec.md#L78-L97) 明确要求 Web 与 agent 共用同一 `ProjectContextService.GetProjectContext / ExportProjectContext` 事实源；
3. [phase12-09 spec:L57-L76](file:///home/dell/Projects/personal-software-company-os/.trae/specs/phase12_09_land_readonly_consumption_shared_entry/spec.md#L57-L76) 明确 Web 跨切片共享只读入口只能落在 `frontend/src/features/project-context/`；
4. 本报告 [3.1](file:///home/dell/Projects/personal-software-company-os/.trae/specs/phase12_11_validate_semantic_alignment_readonly_consumption_foundation/acceptance_report.md#L72-L98) 记录的 shared-readonly `rules` / `phases` / `boundaries` 已显式把规则、phase 与边界入口投影到同一只读结果中；
5. [AGENTS.md:L16-L22](file:///home/dell/Projects/personal-software-company-os/AGENTS.md#L16-L22) 继续给出根级唯一上游，避免 agent 侧另长一套入口。

## 8. Web 浏览器验收与样本恢复

### 8.1 浏览器回归中的环境漂移根因

首轮 focused browser regression 曾出现固定详情页 `not_found`。根因不是前端路由或 shared-readonly 改动本身，而是 `projectcontext` 相关集成测试把共享开发库切回了 `phase06 completed-bound` fixture，导致 `phase11 / phase12` 冻结样本短暂丢失。

本轮通过仓库内正式恢复路径重新补齐 dogfooding 样本：

- `database/scripts/restore_phase11_phase12_dogfooding_sample.sh`
- `database/seeds/seed_phase11_phase12_dogfooding_sample.sql`

随后将 live `8081` 后端进程切换到最新代码，再执行最终浏览器复验。

### 8.2 最终浏览器矩阵结果

最终浏览器验收覆盖以下固定页面：

- `/repositories/$repositoryId`
- `/products/$productId`
- `/modules/$moduleId`
- `/decisions/$decisionId`
- `/dashboard`
- `/onboarding`
- `/reviews/daily`
- `/reviews/weekly`

最终结果：

- 8 / 8 页面正常加载，不再出现 `not_found`
- `repository / product / module` 共享上下文区均已出现：
  - `Product = 经营目标与交付容器`
  - `Repository = 代码仓库身份对象与项目锚点`
  - `plan.md`
  - `architecture_map.md`
  - `docs/README.md`
  - `phase12` 当前入口
- `dashboard / onboarding / reviews/daily / reviews/weekly` 均能通过固定页面与共享语义摘要回看四实体角色
- 未观察到页面内临场猜测 `repository_id` 或第二套解释链

结论：`Task 4` 最终浏览器验收 **通过**。

## 9. Any Mismatch

当前仅保留一项非阻断观察：

1. **`Module / Decision` 语义显式化仍更依赖上游共识文档**：`Q3 / Q4` 当前已达标，但直接纠偏证据仍主要依赖 `PSCO-mvp05-summarize-feedback.md`；这不阻断 `phase12-11` 当前验收结论。

## 10. 当前结论

对 `phase12-11` 而言，本轮已完成：

1. 固定样本、固定入口与同一 `repository_id` 锚点下的样本解析；
2. shared-readonly / agent 固定三入口取证；
3. fixed `6` 问逐题回答与达标判定留档；
4. `Module / Decision` 误读风险复核；
5. `schema` 重写、`MCP / CLI / agent 写回 / 对话入口` 的边界证据留档；
6. 正式工具链重跑；
7. Web 侧 primary owner / 跟随回归页面浏览器验收；
8. 样本漂移根因、正式恢复路径与 rerun 结果留档。

最终判断：

- `Task 4`：**通过**
- `Task 5`：**通过**
- `Task 6`：**通过**
- 最终正式工具链重跑：**通过**
  - `buf build`：通过
  - `buf lint`：通过
  - `go test ./...`：通过
  - `npm run build`：通过
- fixed `6` 问：**6 / 6 达标**
- Web 侧固定页面矩阵：**通过**
- shared-readonly / agent 固定三入口阻断：**已消除**

因此，`phase12-11` 当前可判定为 **通过验收**。
