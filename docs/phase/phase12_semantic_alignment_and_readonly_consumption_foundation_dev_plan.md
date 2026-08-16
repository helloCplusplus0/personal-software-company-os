# phase12_semantic_alignment_and_readonly_consumption_foundation_dev_plan

## 1. 文档定位

本文档定义 `phase12_semantic_alignment_and_readonly_consumption_foundation` 的执行顺序、子任务范围、 DoD 与明确不做。

`phase12` 直接承接 `phase11` 正式收口与 `audit_002` 的仲裁结论，不再继续停留在“要不要再开一轮大范围方向讨论”的状态，而是把已收敛的判断落成一轮新的交付型 phase。

## 2. 本阶段目标

在 `phase11` 已完成最小只读项目上下文导出、并且 `audit_002` 已明确“下一步不做结构重构，也不直接进入更重 agent 通道”的前提下，交付：

- 四实体在 Web 端的语义一致性收口
- Web / agent 共享的只读消费语义深化
- 关键读模型、摘要表达与入口说明的单值化
- 基于真实固定入口的语义一致性与只读消费 dogfooding 验证

使用户与新接手 agent 不再分别依赖两套不同的四实体解释来理解当前项目。

补充边界：

- `phase12` 同时承接“人类用户在 Web 端的语义对齐”与“agent 侧只读消费深化”，但两者共享同一套 canonical 只读语义，不得分裂为两套事实源；
- `phase12` 允许演进只读读模型、摘要视图与文案表达，但不得借机重写四实体结构、关系主线或写路径 owner；
- “更重消费通道 / 受控维护能力”当前仍不进入本阶段正式实现。

## 3. 子任务清单

### 第一组：边界收敛类子任务

### phase12-01 冻结 `Semantic Alignment & Read-Only Consumption Foundation` 的范围边界、成功标准与非目标

范围：

- 冻结本阶段单一主交付能力为 `Semantic Alignment & Read-Only Consumption Foundation`
- 冻结本阶段与四实体结构重构、schema 重写、agent 写回、MCP / CLI / 前端对话入口的边界
- 冻结本阶段成功标准、DoD 与阶段收口口径

DoD：

- 本阶段主交付能力与非目标单值化
- 不把后续更重能力偷渡到本阶段
- 进入 `/spec` 前，后续执行者不再需要猜“本阶段到底做什么”

### phase12-02 冻结四实体在 Web 端的正式语义承接矩阵

范围：

- 冻结 `Product / Repository / Module / Decision` 在 Web 端的正式解释口径
- 冻结哪些页面、摘要卡片、空态、提示文案与下一步动作说明必须显式承接冻结语义
- 冻结四实体解释的共享承接位与禁止散落重复解释的范围
- 冻结 primary owner 页面清单、跟随回归页面清单与允许“无需改动”的记录口径

DoD：

- Web 端不再需要临场解释四实体“到底意味着什么”
- `Module` 与 `Decision` 的当前阶段解释单值化
- 后续执行者能机械回答“哪些页面必须展示什么语义”
- 后续执行者能机械回答“哪些页面是 primary owner，哪些只是跟随回归”

### phase12-03 冻结只读消费深化边界与更重通道进入条件

范围：

- 冻结 `phase11` 已交付的只读项目上下文能力在本阶段可以如何深化
- 冻结 Web / agent 共享只读语义的承接边界
- 冻结共享只读 owner：后端 canonical owner、Web 跨切片 owner 与切片内展示 owner
- 冻结 `repository_id` 在“直接 repository-scoped / 间接 repository-scoped / 衍生消费页”三类页面中的承接边界
- 冻结更重消费通道、受控维护能力与额外专家讨论的进入条件

DoD：

- 不再把“只读消费深化”误读为“直接进入完整 agent 平台”
- 更重能力进入条件单值化
- 当前阶段仍遵守“先消费、后维护”
- 后续执行者不需要临场决定“这里该新建前端共享读模型，还是直接拼页面解释”

### 第二组：实现设计类子任务

### 实现设计类子任务统一产物模板

为避免 `phase12-04 ~ 07` 在进入 `/spec` 前再次出现“设计名义上完成，但执行者仍需临场补判断”的情况，当前阶段统一冻结以下最小产物模板。

`phase12-04 ~ 07` 每项设计子任务至少必须显式产出：

1. **影响对象清单**
   - 逐项列出本子任务覆盖的页面 / 组件 / data owner / 后端合同或服务对象
   - 每项必须标注：`must-change / follow-regression / no-change`
2. **结论矩阵**
   - 逐项回答：本对象当前承接什么、需要改成什么、为什么需要改
   - 若为 `no-change`，必须记录“不改仍满足当前阶段冻结口径”的理由
3. **承接位矩阵**
   - 显式说明该结论最终落在哪个 owner：切片页面、切片组件、切片 data、`frontend/src/features/project-context/`、现有 `ProjectContextService` 或其受控派生读取
4. **共享语义来源 vs 切片内渲染矩阵**
   - 显式区分哪些高频语义短语或共享解释应收敛到唯一共享语义来源
   - 显式区分哪些页面布局、组件结构、空态插入位或导语插入位继续保留在切片内渲染
5. **Before / After 表达样例**
   - 对涉及文案、摘要、解释性语言、共享摘要字段或导出结果的对象，至少给出一组 before / after 样例
6. **明确不做清单**
   - 显式写出本子任务没有扩入的结构重构、写回通道、协议扩张或页面重组事项

补充冻结：

- 若某项设计产物无法按上述模板回答，则该子任务不得视为完成；
- `/spec` 只能承接符合该模板的设计结论，不再接受“实现时再决定”的空位。

### phase12-04 产出前端四实体语义一致性设计

范围：

- 产出 `Product / Repository / Module / Decision` 相关页面的语义对齐设计
- 产出摘要卡片、空态、说明文案、下一步动作说明的承接矩阵
- 产出哪些语义应收敛为共享呈现、哪些保留在切片内的设计
- 产出页面/路由/组件级影响清单，至少覆盖：
  - `frontend/src/routes/dashboard.tsx`
  - `frontend/src/routes/onboarding.tsx`
  - `frontend/src/routes/reviews/daily.tsx`
  - `frontend/src/routes/reviews/weekly.tsx`
  - `frontend/src/routes/products/$productId.tsx`
  - `frontend/src/routes/repositories/$repositoryId.tsx`
  - `frontend/src/routes/modules/$moduleId.tsx`
  - `frontend/src/routes/decisions/$decisionId.tsx`
  - `frontend/src/features/dashboard/pages/dashboard-home-page.tsx`
  - `frontend/src/features/onboarding/pages/onboarding-page.tsx`
  - `frontend/src/features/review/pages/daily-review-page.tsx`
  - `frontend/src/features/review/pages/weekly-review-page.tsx`
  - `frontend/src/features/product-registry/pages/product-detail-page.tsx`
  - `frontend/src/features/repository-binding/pages/repository-binding-detail-page.tsx`
  - `frontend/src/features/module-registry/pages/module-detail-page.tsx`
  - `frontend/src/features/decision-center/pages/decision-detail-page.tsx`
  - `frontend/src/features/product-registry/components/product-summary-card.tsx`
  - `frontend/src/features/repository-binding/components/repository-summary-card.tsx`
  - `frontend/src/features/module-registry/components/module-summary-card.tsx`
  - `frontend/src/features/decision-center/components/decision-detail-summary-card.tsx`
  - `frontend/src/features/dashboard/components/current-focus-section.tsx`
    - `frontend/src/features/dashboard/components/asset-feedback-section.tsx`
  - `frontend/src/features/dashboard/components/recent-activity-section.tsx`
  - `frontend/src/features/review/components/review-page-shell.tsx`
    - `frontend/src/features/module-registry/components/module-next-action-bar.tsx`
    - `frontend/src/features/dashboard/components/dashboard-primary-action-panel.tsx`
    - `frontend/src/features/dashboard/components/onboarding-cta-button.tsx`
    - `frontend/src/features/review/components/review-action-footer.tsx`
    - `frontend/src/features/product-registry/pages/product-list-page.tsx`
    - `frontend/src/features/repository-binding/pages/repository-binding-list-page.tsx`
    - `frontend/src/features/module-registry/pages/module-list-page.tsx`
    - `frontend/src/features/decision-center/pages/decision-list-page.tsx`

DoD：

- 页面级语义对齐结果足以直接进入实现
- 不需要执行者临场再猜哪些页面必须改文案、改摘要或改说明
- 不引入第二套页面事实源
- 允许“无需改动”的对象被逐项记录，而不是被整体跳过

### phase12-05 产出只读消费深化与共享入口设计

范围：

- 产出 `phase11` 项目上下文能力在本阶段的深化方向设计
- 产出 Web / agent 共享规则、约束、文档入口与四实体最小摘要的承接设计
- 产出哪些继续复用既有只读合同、哪些允许新增最小只读承接位的判定规则
- 产出共享只读 owner 矩阵，至少明确：
  - `ProjectContextService.GetProjectContext` 是唯一结构化 canonical owner
  - `ProjectContextService.ExportProjectContext` 是 agent-facing Markdown 导出 owner
  - `frontend/src/features/project-context/` 是唯一允许的新 Web 跨切片共享只读 owner
- 产出“直接 repository-scoped / 间接 repository-scoped / 衍生消费页”三类页面的承接规则
- 显式输出供 `phase12-07` 继续承接的共享摘要字段需求、入口定位需求与最小 resolver 需求

DoD：

- 只读消费深化边界单值化
- 共享消费入口足以直接进入 `/spec`
- 不需要执行者再猜“这里是复用 phase11，还是该新增最小只读承接位”
- 不需要执行者再猜“这里该直接用 repository_id，还是先解析回 repository_id 再复用共享摘要”

### phase12-06 产出前端读路径 owner、共享摘要与回流设计

范围：

- 产出四实体语义对齐所需的读路径 owner 与共享摘要承接位设计
- 产出页面读取、缓存、成功回流与 reread 关系设计
- 识别需要回收的散装页面解释逻辑
- 只承接已经由 `phase12-05` 冻结的共享 owner / 入口矩阵与由 `phase12-07` 冻结的结构化字段结果，不再反向改写二者
- 产出以下读路径 owner 审计清单：
  - `frontend/src/features/product-registry/data/use-product-detail-read.ts`
  - `frontend/src/features/repository-binding/data/use-repository-detail-read.ts`
  - `frontend/src/features/module-registry/data/use-module-detail-read.ts`
  - `frontend/src/features/decision-center/data/use-decision-detail-read.ts`
  - `frontend/src/features/dashboard/data/use-dashboard-overview-read.ts`
  - `frontend/src/features/onboarding/data/use-onboarding-read.ts`
  - `frontend/src/features/review/data/use-daily-review-read.ts`
  - `frontend/src/features/review/data/use-weekly-review-read.ts`
  - 若存在跨切片共享只读，则新增 `frontend/src/features/project-context/data/*`

DoD：

- `query` 层与共享读摘要边界清晰
- 页面层不再各自复制一套四实体解释逻辑
- 足以指导后续实现收口
- 不会把跨切片共享只读逻辑继续散落在 `dashboard / onboarding / review / detail page` 各自的数据层

### phase12-07 产出后端合同、导出结果与共享只读视图设计

范围：

- 产出本阶段若需演进 `.proto / Connect / service / renderer` 的最小设计
- 识别哪些字段、入口定位与摘要结果可被 Web / agent 同时消费
- 冻结“复用 vs 新增”规则，明确任何新增只读承接位必须回收哪段重复解释逻辑
- 若新增字段或视图，必须显式说明：
  - 是否继续复用 `GetProjectContext`
  - 若不复用，为什么仍属于 `ProjectContextService` 下的受控派生读取
  - 它具体回收了哪段前端 / agent 重复解释逻辑
  - 它如何保留 `repository_id` 与 `entry_ref / entry_kind` 的定位能力

DoD：

- 合同、服务与共享只读视图承接位单值化
- 不复制既有 canonical facts
- 足以直接进入 `/spec` 与代码实现
- 后续执行者不需要再猜“这里是不是应该新建第二服务”

### 第三组：源代码实现类子任务

### phase12-08 落实前端四实体语义一致性收口

范围：

- 落实 `Product / Repository / Module / Decision` 相关页面、摘要卡片、空态与说明文案的语义对齐
- 回收关键页面中与冻结语义不一致或不够清晰的表达
- 保持既有切片边界与写路径 owner 不变

DoD：

- Web 端已能稳定表达四实体冻结语义
- 用户不再主要把 `Module` 理解为普通模块登记对象
- 用户不再主要把 `Decision` 理解为孤立文本卡片
- 本子任务只负责表达层与只读呈现层收口，不负责结构重构
- 至少完成 `repositories/$repositoryId`、`products/$productId`、`modules/$moduleId`、`decisions/$decisionId` 四类详情页的 primary owner 收口

### phase12-09 落实只读消费深化与共享只读入口

范围：

- 落实 `phase11` 项目上下文能力在本阶段的共享只读深化
- 落实 Web / agent 共用的规则、约束、文档入口或最小摘要承接结果
- 落实必要的后端合同、服务、导出或共享只读视图

DoD：

- 只读消费比 `phase11` 更稳定、可复用、可定位
- Web 与 agent 不再各自拼装第二套解释性结果
- 不引入写回、Draft、审批流或新协议层
- 若新增 Web 侧跨切片共享只读承接位，其唯一合法路径为 `frontend/src/features/project-context/`

### phase12-10 落实关键页面与共享读路径的联动回流

范围：

- 落实 `Dashboard / Review / Onboarding / Detail pages` 对共享语义摘要的接入
- 落实只读入口更新后的 reread、失效刷新与返回链
- 回收当前页面中语义重复、漂移或入口不稳定的读路径

DoD：

- 关键页面能用同一套共享语义摘要与解释回看当前项目
- 页面返回链与 reread 不再放大语义漂移
- 本子任务只负责读路径与呈现回流，不新增并列写路径
- `dashboard / onboarding / reviews/daily / reviews/weekly` 只能作为衍生消费页接入共享摘要或受控入口，不得自长新的结构化主锚点
- `dashboard / onboarding / reviews/daily / reviews/weekly` 必须能通过共享语义摘要或固定入口解释当前四实体角色；若仍依赖切片内旧文案单独解释，则不得视为完成

### 第四组：验证验收类子任务

### phase12-11 完成 `Semantic Alignment & Read-Only Consumption Foundation` 的联调、dogfooding 与反回归验证

范围：

- 完成 `buf / go test / frontend build` 最小工具链验证
- 完成 Web 端四实体语义一致性的浏览器与读路径验证
- 完成 Web / agent 共享只读消费入口的 dogfooding 验证
- 留档本阶段明确不做 schema 重写、MCP / CLI / agent 写回 / 对话入口的边界证据
- 冻结固定验收问题，至少覆盖：
  - Web 端当前如何解释 `Product / Repository / Module / Decision`
  - 新接手 agent 与人类用户是否可回到同一组规则、约束与文档入口
  - 当前 `Module / Decision` 是否仍被错误理解
  - 当前只读消费是否仍保持单一结构化输入锚点与只读边界
  - 共享只读 owner 是否已经单值化

验收协议至少冻结为：

- 固定样本继续使用 `phase11` 已补齐的 `PSCO` dogfooding 样本：
  - `Repository`：`personal-software-company-os`
  - `repository_id`：`ca261521-8daf-4248-8f12-43525326e759`
  - `Product`：`PSCO`
  - `Module`：`project-context-foundation`
  - `Decision`：`phase11 Project Context Foundation dogfooding 验收决策`
- agent 侧固定入口集合冻结为：
  - `AGENTS.md`
  - `PSCO-mvp05-summarize-feedback.md`
  - 基于同一 `repository_id` 读取的共享只读结果（结构化读取或其受控派生视图）
- Web 侧 primary owner 固定验证页面冻结为：
  - `/repositories/$repositoryId`
  - `/products/$productId`
  - `/modules/$moduleId`
  - `/decisions/$decisionId`
- Web 侧跟随回归页面冻结为：
  - `/dashboard`
  - `/onboarding`
  - `/reviews/daily`
  - `/reviews/weekly`
- Web 侧详情页允许先通过稳定名称解析 `product_id / module_id / decision_id`，但不允许在页面内临场猜测 `repository_id`
- 样本解析协议进一步冻结为：
  1. `repository_id` 只允许使用固定样本中冻结的正式值，不允许临场改样本或改锚点
  2. `product_id / module_id / decision_id` 只允许从同一 `repository_id` 驱动的结构化只读结果或其受控派生视图解析，不允许绕到数据库手工查询、浏览器临场搜索或额外脚本补齐
  3. 解析链路必须在验收记录中显式留档：使用了哪个固定入口、解析出了什么 id、是否一次成功
  4. 若名称解析失败、结果不唯一或无法回到同一 `repository_id`，则本轮验收直接判定失败，不允许通过新增第 `7` 个入口补救
- 固定验收提问冻结为 `6` 问：
  1. 当前 `Product` 的正式定位是什么，在哪个固定入口可直接确认
  2. 当前 `Repository` 的正式定位是什么，在哪个固定入口可直接确认
  3. 当前 `Module` 为什么不是普通模块登记对象，在哪个固定入口可直接确认“可复用能力资产”语义
  4. 当前 `Decision` 为什么不是孤立文本卡片，在哪个固定入口可直接确认“规则 / 约束 / 选择与依据索引对象”语义
  5. 当前项目共享的规则、约束与文档入口从哪里查看，Web 与 agent 是否回到同一组入口
  6. 当前只读消费是否仍由同一 `repository_id` 锚点驱动，且没有长出第二套事实源
- 验收记录必须留档：
  - 样本解析方式
  - 固定页面 / 固定入口清单
  - 每个问题的回答结果
  - 失败点与是否达标
- 不允许通过额外第 `7` 个临时入口补齐回答

DoD：

- 工具链、API、浏览器与关键反回归均通过
- Web 与 agent 对四实体语义的回答可被真实固定入口复验
- 只读消费深化证据足以说明本阶段不是停留在抽象描述
- 同一份验收记录足以让不同执行者按相同问题、相同样本与相同入口 rerun

### 第五组：根级同步类子任务

### phase12-12 完成根级同步、阶段收口与下一阶段进入条件回写

范围：

- 回写 `AGENTS.md / plan.md / docs/README.md / architecture_map.md / docs/phase/README.md`
- 留档本阶段正式验收与收口入口
- 明确更重消费通道或受控维护能力只允许在 `phase12` 正式收口后，再依据新条件讨论或进入

DoD：

- 根级状态、docs 入口与阶段记录同步完成
- 不长出新的孤岛文档
- 下一阶段进入条件单值化

## 4. 明确不做

本阶段明确不做：

1. 四实体 schema 重写或关系主线重构
2. MCP 协议层正式实现
3. CLI 工具正式实现
4. agent 自动写回、Draft 接口、审批流
5. 前端对话式 agent 入口
6. 把 Web 做成 agent 工作台
7. 第二套 canonical API 或影子状态表

## 5. 子任务依赖关系

为避免后续执行时再次出现“先大通道扩张、后补语义”的顺序错乱，当前阶段依赖关系冻结如下：

1. `phase12-01` 是全阶段边界前提，后续所有子任务都直接依赖它
2. `phase12-02` 与 `phase12-03` 是共享语义与共享只读边界前提，`phase12-04 ~ 07` 必须直接承接这两项结论
3. `phase12-04 ~ 07` 属于实现设计层，必须先于 `phase12-08 ~ 10`
4. `phase12-05` 先冻结共享只读 owner、页面承接分类与最小共享摘要需求
5. `phase12-07` 再基于 `phase12-05` 冻结的共享摘要需求，决定是否复用 `GetProjectContext` 或增加受控派生读取
6. `phase12-06` 最后基于 `phase12-05 / 07` 已冻结的 owner、字段与入口结果，收敛前端读路径 owner、adapter、query options 与 reread 关系
7. `phase12-08` 只依赖 `phase12-02 / 04 / 06`
8. `phase12-09` 只依赖 `phase12-03 / 05 / 07`
9. `phase12-10` 依赖 `phase12-04 ~ 07`，且不得回头改写 `phase12-08` 已冻结的四实体语义表达
10. `phase12-11` 依赖 `phase12-08 ~ 10`
11. `phase12-12` 依赖 `phase12-11`
