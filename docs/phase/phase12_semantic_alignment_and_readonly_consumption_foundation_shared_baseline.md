# phase12_semantic_alignment_and_readonly_consumption_foundation_shared_baseline

## 1. 文档定位

本文档用于集中冻结 `phase12` 的共享基线，避免相同结论在 `architecture_plan`、`dev_plan`、后续 `/spec` 与根级入口中重复发散。

`phase12` 当前处于 `/plan` 阶段。本文档只承接当前阶段的单值基线、语义矩阵、只读消费矩阵与验收前提，不替代后续 `/spec`、实现或根级状态正文。

## 2. 当前单值基线

### 2.1 项目路线

- 当前项目：`PSCO`
- 当前 phase：`phase12_semantic_alignment_and_readonly_consumption_foundation`
- 当前技术路线：`Durable System Track`
- 当前根级阶段状态：`phase11_project_context_foundation` 已完成正式收口，`phase12` 当前作为下一阶段正式 `/plan` 入口建立
- 当前阶段规划上游统一以 `PSCO-mvp05-summarize-feedback.md` 与 `audit_002` 为准

### 2.2 当前阶段唯一直接执行层上游

- 直接执行层上游：
  - `PSCO-mvp05-summarize-feedback.md`
  - `docs/audit/audit_002_phase11_post_closeout_direction_issue.md`
  - `docs/audit/audit_002_phase11_post_closeout_direction_analysis.md`
  - `AGENTS.md`
  - `plan.md`
  - `project_rules.md`
  - `TECH_STACK_BASELINE.md`
  - `architecture_map.md`
  - `docs/README.md`
  - `docs/phase/README.md`
  - `docs/phase/phase11_project_context_foundation_*`
  - `.trae/specs/phase11_09_validate_project_context_foundation_dogfooding_regression/acceptance_report.md`
  - `.trae/specs/phase03_10_decision_center_formal_spec/decision_center_spec_v0.1.md`
  - `.trae/specs/phase04_10_product_repository_binding_formal_spec/product_repository_binding_spec_v0.1.md`
  - `.trae/specs/phase05_10_dashboard_feedback_formal_spec/dashboard_feedback_spec_v0.1.md`
  - `.trae/specs/phase06_12_onboarding_sovereignty_reuse_formal_spec/spec.md`
- 当前阶段只承接：
  - 四实体语义一致性收口
  - 只读消费深化
- 当前阶段不反向重写 `phase03 ~ phase11` 已完成的正式交付

### 2.3 当前阶段正式技术主线

- Web：`React + Vite + TypeScript`
- Frontend Delivery：单一 `React Web`
- Router：`TanStack Router`
- Data Fetching：`TanStack Query`
- UI：`Tailwind CSS + shadcn/ui`
- Backend：`Go + chi + net/http + ConnectRPC`
- Database：`PostgreSQL`
- Contract：`Protocol Buffers`
- 运行方式：`Single Server First`

### 2.4 当前阶段特别约束

- `.proto` 是唯一长期合同源
- `ConnectRPC` 是业务接口正式传输层
- `chi` 只承担 router shell、middleware 与非业务端点承接职责
- 当前阶段新增或演进后的前端业务动作必须继续遵守：
  - 写路径唯一 `application` 入口
  - `query` 层纯只读
  - mutation 固定承接位
- 当前阶段新增或演进后的只读能力必须继续保持只读
- 当前阶段不允许：
  - 引入 agent 专属一级业务对象
  - 引入第二套 canonical API
  - 把 PSCO 做成 IDE 现场流程控制器
  - 把 Web 做成对话式 agent 工作台
  - 把 MCP / CLI / agent 写回偷渡为当前阶段主交付
  - 把四实体语义澄清偷渡为 schema 重构

### 2.5 当前阶段交付模式

- `phase12` 是交付型 phase，不是纯文档冻结阶段
- 当前 `/plan` 只负责建立阶段上游、任务拆分与共享基线
- 当前阶段后续必须继续进入 `/spec`、源代码实现、验证验收与根级同步
- 当前阶段结束时必须新增可运行、可验收的“语义一致性收口 + 只读消费深化”主线代码

## 3. 当前阶段能力矩阵

### 3.1 Semantic Alignment & Read-Only Consumption Foundation 单值定义

`Semantic Alignment & Read-Only Consumption Foundation` 在当前阶段冻结为：

- Web 端四实体语义一致性收口
- Web / agent 共享的只读语义深化
- 规则、约束、文档入口与四实体摘要解释的共享读路径收敛
- 基于固定入口的语义一致性与只读消费 dogfooding 验证

当前阶段不把以下内容解释为该主交付：

- MCP 协议层
- CLI 工具
- agent 自动写回
- Draft / 审批流
- 前端对话式 agent 工作台
- 四实体结构重构

### 3.2 四实体语义矩阵

- `Product`：经营目标与交付容器
- `Repository`：代码仓库身份对象与项目锚点
- `Module`：可复用能力资产，允许后置提炼
- `Decision`：规则、约束、选择与依据的索引对象

补充冻结：

- 当前阶段只确认并落实表达层与消费层的一致性，不重写结构
- `Module` 在 Web 端必须被解释为可复用能力资产，而不是普通模块登记对象
- `Decision` 在 Web 端必须被解释为规则、约束、选择与依据的索引对象，而不是孤立文本卡片
- 后续 `/spec` 与实现不得把语义对齐偷渡为实体重构

### 3.3 Web / agent 共享语义矩阵

- Web：
  - 全局查看
  - 关系校对
  - review 与回顾
  - 历史查阅
  - 人工修正
  - 最终确认
  - 显式承接四实体冻结语义

- agent：
  - 在项目现场消费当前项目上下文
  - 读取与当前项目直接相关的规则、决策与文档入口
  - 复用同一套四实体最小摘要与只读解释

补充冻结：

- Web 不退化
- agent 不对称并行进入
- 二者共享同一套 Go backend canonical core
- 不允许二者各自长出第二套四实体解释

### 3.4 只读消费深化矩阵

当前阶段的只读消费深化至少应覆盖：

- 唯一结构化输入锚点：`repository_id`
- 当前 `Repository` 身份与最小上下文摘要
- 关联 `Product / Module / Decision` 的最小可读解释
- 与当前项目直接相关的规则、约束与文档入口
- 可被 Web / agent 共享消费的受控只读摘要或导出结果

补充冻结：

- 当前阶段不引入 `product_id`、本地路径、Git remote URL 或工作区扫描作为并列主锚点
- 当前阶段继续只承接“已完成 Repository Binding”的仓库上下文读取
- 结构化只读读取继续落在 Go backend 的 `.proto + ConnectRPC` 正式主线
- 共享只读结果必须从同一结构化只读事实源单向派生
- 当前阶段不把消费侧项目目录中的固定文件名作为必要输入合同

进一步冻结的 owner 与消费边界：

- 结构化 canonical owner：`proto/psco/project_context/v1/project_context.proto` 与 `backend/internal/projectcontext/*`
- agent-facing Markdown owner：`ProjectContextService.ExportProjectContext`
- Web 跨切片共享只读 owner：`frontend/src/features/project-context/`
- 切片内展示 owner：各 feature 的 `pages/`、`components/` 与 `data/`
- 直接 repository-scoped 页面只允许是 `repositories/$repositoryId`
- 间接 repository-scoped 页面包括 `products/$productId`、`modules/$moduleId`、`decisions/$decisionId`
- 衍生消费页包括 `dashboard`、`onboarding`、`reviews/daily`、`reviews/weekly`
- 衍生消费页不得自行升格为新的结构化主锚点入口

### 3.5 当前阶段前端承接矩阵

当前阶段至少应直接承接：

- `Dashboard`
- `Review`
- `Onboarding`
- `Product Detail / List`
- `Repository Detail / List / Binding`
- `Module Detail / List`
- `Decision Detail / List`

补充冻结：

- 页面层不得继续各自复制四实体解释逻辑
- 若存在共享语义摘要，应优先收敛到稳定承接位，而不是在每页内联拼文案
- 本阶段只演进读路径与呈现语义，不新增第二套写路径

进一步冻结的 primary owner 页面与组件：

- 主详情页：
  - `frontend/src/features/product-registry/pages/product-detail-page.tsx`
  - `frontend/src/features/repository-binding/pages/repository-binding-detail-page.tsx`
  - `frontend/src/features/module-registry/pages/module-detail-page.tsx`
  - `frontend/src/features/decision-center/pages/decision-detail-page.tsx`
- 跟随回归页：
  - `frontend/src/features/dashboard/pages/dashboard-home-page.tsx`
  - `frontend/src/features/onboarding/pages/onboarding-page.tsx`
  - `frontend/src/features/review/pages/daily-review-page.tsx`
  - `frontend/src/features/review/pages/weekly-review-page.tsx`
- 主摘要组件：
  - `frontend/src/features/product-registry/components/product-summary-card.tsx`
  - `frontend/src/features/repository-binding/components/repository-summary-card.tsx`
  - `frontend/src/features/module-registry/components/module-summary-card.tsx`
  - `frontend/src/features/decision-center/components/decision-detail-summary-card.tsx`
- 跟随摘要 / 容器组件：
  - `frontend/src/features/dashboard/components/current-focus-section.tsx`
    - `frontend/src/features/dashboard/components/asset-feedback-section.tsx`
  - `frontend/src/features/dashboard/components/recent-activity-section.tsx`
  - `frontend/src/features/review/components/review-page-shell.tsx`
    - `frontend/src/features/module-registry/components/module-next-action-bar.tsx`
    - `frontend/src/features/dashboard/components/dashboard-primary-action-panel.tsx`
    - `frontend/src/features/dashboard/components/onboarding-cta-button.tsx`
    - `frontend/src/features/review/components/review-action-footer.tsx`
  - 默认非 primary owner 但进入 `phase12-04` surface 审计面的文件：
    - `frontend/src/features/product-registry/pages/product-list-page.tsx`
    - `frontend/src/features/repository-binding/pages/repository-binding-list-page.tsx`
    - `frontend/src/features/module-registry/pages/module-list-page.tsx`
    - `frontend/src/features/decision-center/pages/decision-list-page.tsx`

### 3.5A 设计产物最小模板

为保证 `phase12-04 ~ 07` 的设计结果可以被不同执行者机械承接，当前阶段统一冻结以下最小设计产物模板：

1. 影响对象清单：
   - 逐项记录页面 / 组件 / data owner / 后端合同或服务对象
   - 每项必须标注 `must-change / follow-regression / no-change`
2. 结论矩阵：
   - 当前承接什么
   - 需要改成什么
   - 为什么要改
   - 若不改，为什么仍满足当前阶段冻结口径
3. 承接位矩阵：
   - 最终落到切片页面、切片组件、切片 data、`frontend/src/features/project-context/`、现有 `ProjectContextService` 或受控派生读取中的哪一处
4. 共享语义来源 vs 切片内渲染矩阵：
   - 显式区分哪些高频语义短语或共享解释应收敛到唯一共享语义来源
   - 显式区分哪些页面布局、组件结构、空态插入位或导语插入位继续保留在切片内渲染
5. Before / After 样例：
   - 至少覆盖一组文案、摘要字段或解释性结果的 before / after
6. 明确不做清单：
   - 本子任务没有扩入的结构重构、写回通道、协议扩张或页面重组事项

补充冻结：

- 不符合上述模板的设计结果，不得直接进入 `/spec`
- `/spec` 与实现只能承接已写清 owner、结论与不做项的设计结果

## 4. 当前阶段验收前提

当前阶段验收必须至少回答：

1. Web 端是否已能用固定入口准确解释 `Product / Repository / Module / Decision`
2. 同一固定入口集合下，人类用户与新接手 agent 是否能回到同一组规则、约束与文档入口
3. `Module` 是否仍被主要理解为普通模块登记对象
4. `Decision` 是否仍被主要理解为孤立文本记录
5. 只读消费深化是否仍完全保持只读，且唯一结构化输入锚点是否保持为 `repository_id`
6. 当前阶段是否严格没有引入 agent 写入路径、第二套语义、第二套事实源或更重协议层

补充冻结的固定样本：

- `Repository`：`personal-software-company-os`
- `repository_id`：`ca261521-8daf-4248-8f12-43525326e759`
- `Product`：`PSCO`
- `Module`：`project-context-foundation`
- `Decision`：`phase11 Project Context Foundation dogfooding 验收决策`

补充冻结的 dogfooding 入口集合：

- Web 侧固定入口集合至少应包括：
  - `AGENTS.md`
  - `PSCO-mvp05-summarize-feedback.md`
  - `/repositories/$repositoryId`
  - `/products/$productId`
  - `/modules/$moduleId`
  - `/decisions/$decisionId`
  - `/dashboard`
  - `/onboarding`
  - `/reviews/daily`
  - `/reviews/weekly`
- agent 侧固定入口集合至少应包括：
  - `AGENTS.md`
  - `PSCO-mvp05-summarize-feedback.md`
  - 基于同一 `repository_id` 生成或读取的共享只读结果
- 任一侧均不允许临时增加额外入口来补齐四实体解释
- 样本解析协议进一步冻结为：
  - `repository_id` 只允许使用固定样本中冻结的正式值
  - `product_id / module_id / decision_id` 只允许从同一 `repository_id` 驱动的结构化只读结果或其受控派生视图解析
  - 验收记录必须显式留档：解析入口、解析出的 id、是否一次成功
  - 若名称解析失败、结果不唯一或无法回到同一 `repository_id`，则本轮验收直接失败，不允许引入额外入口补救

补充冻结的固定验收问题：

1. 当前 `Product` 的正式定位是什么，在哪个固定入口可直接确认
2. 当前 `Repository` 的正式定位是什么，在哪个固定入口可直接确认
3. 当前 `Module` 为什么不是普通模块登记对象，在哪个固定入口可直接确认“可复用能力资产”语义
4. 当前 `Decision` 为什么不是孤立文本卡片，在哪个固定入口可直接确认“规则 / 约束 / 选择与依据索引对象”语义
5. 当前项目共享的规则、约束与文档入口从哪里查看，Web 与 agent 是否回到同一组入口
6. 当前只读消费是否仍由同一 `repository_id` 锚点驱动，且没有长出第二套事实源

## 5. 当前阶段完成条件

`phase12` 完成时，至少必须满足：

1. Web 端四实体语义表达已与 `phase11` 冻结口径单值一致
2. 在 `/repositories/$repositoryId`、`/products/$productId`、`/modules/$moduleId`、`/decisions/$decisionId` 四类详情页中，`Module` 与 `Decision` 的当前阶段解释可被直接复验
3. `dashboard`、`onboarding`、`reviews/daily`、`reviews/weekly` 已能通过共享语义摘要或固定入口解释当前四实体角色，而不是继续依赖切片内旧文案各自解释
4. `phase11` 的最小只读项目上下文能力已深化为更稳定、可定位、可共享的消费入口，且跨切片共享只读 owner 已单值化
5. 当前阶段未引入 MCP / CLI / agent 写回 / 对话入口 / schema 重写 / 第二套 canonical API
6. 已按固定样本、固定入口与固定 `6` 问完成 Web / agent 双侧语义一致性与只读消费验证
7. 更重消费通道或受控维护能力的进入条件已继续保持为后续条件，而不是当前事实
