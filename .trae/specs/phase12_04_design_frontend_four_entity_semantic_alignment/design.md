# phase12-04 前端四实体语义一致性设计

> 本设计产出已足以进入后续实现链条（先承接 `phase12-05 / 07 / 06`，再进入 `phase12-08`），所有对象均按 `must-change / follow-regression / no-change` 逐项分类并留档。

---

## 一、影响对象全量清单与分类矩阵

### 1.1 Route 层（8 个）

Route 只负责稳定路由承接与交接，不作为语义事实源。全部 `no-change`。

| # | 文件 | 分类 | 理由 |
|---|---|---|---|
| R1 | `frontend/src/routes/dashboard.tsx` | no-change | 路由定义，不承载四实体语义 |
| R2 | `frontend/src/routes/onboarding.tsx` | no-change | 路由定义，不承载四实体语义 |
| R3 | `frontend/src/routes/reviews/daily.tsx` | no-change | 路由定义，不承载四实体语义 |
| R4 | `frontend/src/routes/reviews/weekly.tsx` | no-change | 路由定义，不承载四实体语义 |
| R5 | `frontend/src/routes/products/$productId.tsx` | no-change | 路由定义，不承载四实体语义 |
| R6 | `frontend/src/routes/repositories/$repositoryId.tsx` | no-change | 路由定义，不承载四实体语义 |
| R7 | `frontend/src/routes/modules/$moduleId.tsx` | no-change | 路由定义，不承载四实体语义 |
| R8 | `frontend/src/routes/decisions/$decisionId.tsx` | no-change | 路由定义，不承载四实体语义 |

### 1.2 Page 层（8 个）

> 注：P1-P4 在 phase12-02 中已冻结为 **Primary owner 页面**。当前设计不再把它们降级为被动容器，而是明确要求页面与其内含的 primary owner summary-card 一起显式承接冻结语义；summary-card 是最核心承接位，但不是唯一承接位。

| # | 文件 | primary owner? | 改动分类 | 理由 |
|---|---|---|---|---|
| P1 | `product-detail-page.tsx` | 是（phase12-02） | must-change | Product 详情页作为 primary owner 页面，需与 ProductSummaryCard 一起显式承接"经营目标与交付容器"语义 |
| P2 | `repository-binding-detail-page.tsx` | 是（phase12-02） | must-change | Repository 详情页作为 primary owner 页面，需与 RepositorySummaryCard 一起显式承接"代码仓库身份对象与项目锚点"语义 |
| P3 | `module-detail-page.tsx` | 是（phase12-02） | must-change | Module 详情页作为 primary owner 页面，需与 ModuleSummaryCard 和 ModuleNextActionBar 一起显式承接"可复用能力资产"语义 |
| P4 | `decision-detail-page.tsx` | 是（phase12-02） | must-change | Decision 详情页作为 primary owner 页面，需与 DecisionDetailSummaryCard 一起显式承接"规则、约束、选择与依据的索引对象"语义 |
| P5 | `dashboard-home-page.tsx` | 否（跟随回归） | follow-regression | 包含 CurrentFocusSection / AssetFeedbackSection / RecentActivitySection，均跟随回归 |
| P6 | `onboarding-page.tsx` | 否（跟随回归） | must-change | WelcomeStep 将四实体描述为 CRUD 步骤，需改为冻结语义 |
| P7 | `daily-review-page.tsx` | 否（跟随回归） | follow-regression | 显示 current focus / pending decisions，跟随 primary owner 语义 |
| P8 | `weekly-review-page.tsx` | 否（跟随回归） | follow-regression | 显示 current focus / pending decisions / template candidates，跟随 primary owner 语义 |

### 1.3 Component 层（12 个）

| # | 文件 | 分类 | 理由 |
|---|---|---|---|
| C1 | `product-summary-card.tsx` | must-change | 当前仅展示 name/status/description/created_at，需增加"经营目标与交付容器"语义表达 |
| C2 | `repository-summary-card.tsx` | must-change | 当前仅展示 name/status/url/provider/created_at，需增加"代码仓库身份对象与项目锚点"语义表达 |
| C3 | `module-summary-card.tsx` | must-change | **核心差距**：当前仅展示 name/status/description/created_at，需改为"可复用能力资产"语义，配合复用摘要展示 |
| C4 | `decision-detail-summary-card.tsx` | must-change | 已有结构化字段但缺少"规则/约束/选择与依据的索引对象"语义框架 |
| C5 | `current-focus-section.tsx` | follow-regression | 展示 feedback signals，跟随 signal 来源的实体语义 |
| C6 | `asset-feedback-section.tsx` | follow-regression | 展示 representative signals / reuse snapshot，跟随 Module 与反馈语义回归验证 |
| C7 | `recent-activity-section.tsx` | follow-regression | 展示 recent activity items，跟随 activity 来源的实体语义 |
| C8 | `review-page-shell.tsx` | follow-regression | Review 页面壳层本身不定义实体语义，但必须跟随 Daily/Weekly Review 的语义改动完成回归验证 |
| C9 | `module-next-action-bar.tsx` | must-change | 该组件直接承接 Module 相关下一步动作说明，必须显式回到"可复用能力资产"与 Repository 完整冻结语义 |
| C10 | `dashboard-primary-action-panel.tsx` | no-change | 主 CTA 面板属于 review 入口，不独立定义四实体语义 |
| C11 | `onboarding-cta-button.tsx` | no-change | onboarding 入口按钮，不独立定义四实体语义 |
| C12 | `review-action-footer.tsx` | no-change | review 完成动作区，不独立定义四实体语义 |

### 1.4 默认非 primary owner 对象（10 个）

| # | 对象 | 分类 | 理由 |
|---|---|---|---|
| N1 | `product-list-page.tsx` | must-change | Product 列表页仍不是 primary owner，但文件内直接承接 Product 空态与来源提示，必须把列表文件本身改回冻结语义 |
| N2 | `repository-binding-list-page.tsx` | must-change | Repository 列表页仍不是 primary owner，但文件内直接承接 Repository 空态与来源提示，必须把列表文件本身改回冻结语义 |
| N3 | `module-list-page.tsx` | must-change | Module 列表页仍不是 primary owner，但文件内直接承接 Module 空态与来源提示，必须把列表文件本身改回冻结语义 |
| N4 | `decision-list-page.tsx` | must-change | Decision 列表页仍不是 primary owner，但文件内直接承接 Decision 空态与来源提示，必须把列表文件本身改回冻结语义 |
| N5 | `product-list-toolbar.tsx` | follow-regression | 产品列表 toolbar 不独立定义实体语义，但需跟随 Product 列表语义回归验证 |
| N6 | `repository-binding-list-toolbar.tsx` | follow-regression | Repository 列表 toolbar 不独立定义实体语义，但需跟随 Repository 列表语义回归验证 |
| N7 | `module-list-toolbar.tsx` | follow-regression | Module 列表 toolbar 不独立定义实体语义，但需跟随 Module 列表语义回归验证 |
| N8 | `decision-list-toolbar.tsx` | follow-regression | Decision 列表 toolbar 不独立定义实体语义，但需跟随 Decision 列表语义回归验证 |
| N9 | `module-list-search-store.ts` | follow-regression | 搜索状态不承载实体解释，但需保持与 Module 列表过滤语义一致，不得长出第二套实体解释 |
| N10 | `decision-list-search-store.ts` | follow-regression | 搜索状态不承载实体解释，但需保持与 Decision 列表过滤语义一致，不得长出第二套实体解释 |

> 注：N1-N4 的文件级分类改为 `must-change`，仅表示这些非 primary owner 列表页中存在必须落盘的语义 surface；**不表示** 列表页 owner 身份被提升为 primary owner。

---

## 二、四实体 Surface 承接矩阵

### 2.1 摘要卡片（Summary Cards）

| Surface | 分类 | 当前承接 | 需改为 | 共享呈现/切片内保留 |
|---|---|---|---|---|
| ProductSummaryCard | must-change | 产品登记条目（name/status/description/created_at） | 在标题区或副标题区增加"经营目标与交付容器"语义标签，Description 区域可保留但需与"交付目标"语义对齐 | 切片内保留（Product feature 专属） |
| RepositorySummaryCard | must-change | 仓库登记条目（name/status/url/provider/created_at） | 在标题区增加"代码仓库身份对象与项目锚点"语义标签，URL/provider 保持为身份属性 | 切片内保留（Repository feature 专属） |
| ModuleSummaryCard | must-change | 模块登记条目（name/status/description/created_at） | 标题区增加"可复用能力资产"语义标签，description 需与"可复用能力"语义对齐，可配合 ReuseSummaryInline 展示复用事实 | 切片内保留（Module feature 专属），复用摘要已在 ReuseSummaryInline 中展示 |
| DecisionDetailSummaryCard | must-change | 结构化模板字段（context/problem/alternatives/choice/reason/impact） | 在卡片顶部增加一行语义框架说明（如"本条决策记录了以下规则、约束、选择与依据"），保持现有字段不变 | 切片内保留（Decision feature 专属） |

### 2.1A Primary owner Detail Page 页面级承接位

> 该小节用于把 P1-P4 的“页面级承接位”冻结成单值，避免再次退回到“只有 summary-card 在改、页面本身只是被动容器”的含糊状态。

| Page | 页面级承接位 | 分类 | 当前状态 | 需改为 | 最终承接位 |
|---|---|---|---|---|---|
| Product Detail Page | 页面标题区与 SummaryCard 之间的语义导语 | must-change | 当前页面顶部仅承接名称、来源模板与 summary-card，缺少 Product 页面级语义锚点 | 在标题区或 summary-card 前增加一句导语，明确“该 Product 代表一个经营目标与交付容器” | 切片内保留（`product-detail-page.tsx`），短语引用共享语义来源 |
| Repository Detail Page | 页面标题区与 SummaryCard 之间的语义导语 | must-change | 当前页面顶部仅承接名称与来源上下文，缺少 Repository 页面级语义锚点 | 在标题区或 summary-card 前增加一句导语，明确“该 Repository 代表一个代码仓库身份对象与项目锚点” | 切片内保留（`repository-binding-detail-page.tsx`），短语引用共享语义来源 |
| Module Detail Page | 页面标题区与 SummaryCard/NextActionBar 之间的语义导语 | must-change | 当前页面顶部仅承接名称与结构块，缺少 Module 页面级语义锚点 | 在标题区或 summary-card 前增加一句导语，明确“该 Module 代表一个可复用能力资产”，并与 NextActionBar 一致引用同一短语 | 切片内保留（`module-detail-page.tsx`），短语引用共享语义来源 |
| Decision Detail Page | 页面标题区与 SummaryCard 之间的语义导语 | must-change | 当前页面顶部仅承接标题与结构化内容入口，缺少 Decision 页面级语义锚点 | 在标题区或 summary-card 前增加一句导语，明确“该 Decision 用于索引规则、约束、选择与依据” | 切片内保留（`decision-detail-page.tsx`），短语引用共享语义来源 |

### 2.2 空态（Empty States）

| Surface | 位置 | 分类 | 当前语义 | 需改为 | 最终承接位 |
|---|---|---|---|---|---|
| ProductListPage 空态 | `product-registry/pages/` | must-change | "暂无数据"或等效无语义占位 | "暂无经营目标与交付容器" | 切片内保留（`product-list-page.tsx`） |
| RepositoryListPage 空态 | `repository-binding/pages/` | must-change | "暂无数据"或等效无语义占位 | "暂无代码仓库身份对象与项目锚点" | 切片内保留（`repository-binding-list-page.tsx`） |
| ModuleListPage 空态 | `module-registry/pages/` | must-change | "暂无数据"或等效无语义占位 | "暂无已登记的可复用能力资产" | 切片内保留（`module-list-page.tsx`） |
| DecisionListPage 空态 | `decision-center/pages/` | must-change | "暂无数据"或等效无语义占位 | "暂无规则与决策记录" | 切片内保留（`decision-list-page.tsx`） |
| Dashboard Current Focus 空态 | `current-focus-section.tsx` | follow-regression | "暂无待处理反馈信号" | 保持现状；该空态属于 review / signal 会话语义，不独立定义四实体语义 | 切片内保留（`current-focus-section.tsx`） |
| Dashboard Asset Feedback 空态 | `asset-feedback-section.tsx` | follow-regression | "暂无代表性缺口项" | 保持现状；该空态属于反馈缺口语义，不独立定义四实体语义 | 切片内保留（`asset-feedback-section.tsx`） |
| Dashboard Recent Activity 空态 | `recent-activity-section.tsx` | follow-regression | "暂无最近活动" | 保持现状；该空态属于活动流语义，不独立定义四实体语义 | 切片内保留（`recent-activity-section.tsx`） |
| Daily Review 焦点信号空态 | `daily-review-page.tsx` | follow-regression | "暂无待处理焦点信号" | 保持现状；跟随 review 会话语义回归验证 | 切片内保留（`daily-review-page.tsx`） |
| Daily Review 待处理决策空态 | `daily-review-page.tsx` | follow-regression | "暂无待处理决策" | 保持现状；跟随 decision primary owner 语义回归验证 | 切片内保留（`daily-review-page.tsx`） |
| Daily Review 代表性反馈空态 | `daily-review-page.tsx` | follow-regression | "暂无代表性反馈信号" | 保持现状；跟随 feedback/review 语义回归验证 | 切片内保留（`daily-review-page.tsx`） |
| Weekly Review 最近活动空态 | `weekly-review-page.tsx` | follow-regression | "暂无最近活动" | 保持现状；跟随 recent activity 语义回归验证 | 切片内保留（`weekly-review-page.tsx`） |
| Weekly Review 代表性反馈空态 | `weekly-review-page.tsx` | follow-regression | "暂无代表性反馈信号" | 保持现状；跟随反馈语义回归验证 | 切片内保留（`weekly-review-page.tsx`） |
| Weekly Review 模板候选空态 | `weekly-review-page.tsx` | follow-regression | "当前没有可复用模板候选" | 保持现状；该空态属于模板候选语义，不独立定义四实体语义 | 切片内保留（`weekly-review-page.tsx`） |
| Onboarding Welcome Step 空态 | `onboarding-page.tsx` | follow-regression | 当前无独立空态，直接展示 Welcome 引导卡 | 跟随 WelcomeStep 实体介绍语义改动回归验证 | 切片内保留（`onboarding-page.tsx`） |
| Onboarding Product Step 空态 | `onboarding-page.tsx` | follow-regression | 当前无独立空态，由表单或 DraftSummaryCard 驱动 | 跟随 Product 语义改动回归验证 | 切片内保留（`onboarding-page.tsx`） |
| Onboarding Repository Step 空态 | `onboarding-page.tsx` | follow-regression | 当前无独立空态，由表单或 DraftSummaryCard 驱动 | 跟随 Repository 语义改动回归验证 | 切片内保留（`onboarding-page.tsx`） |
| Onboarding Module Step 空态 | `onboarding-page.tsx` | follow-regression | 当前无独立空态，由表单或 DraftSummaryCard 驱动 | 跟随 Module 语义改动回归验证 | 切片内保留（`onboarding-page.tsx`） |
| Onboarding Decision Step 空态 | `onboarding-page.tsx` | follow-regression | 当前无独立空态，由表单或 DraftSummaryCard 驱动 | 跟随 Decision 语义改动回归验证 | 切片内保留（`onboarding-page.tsx`） |
| Onboarding Complete Step 空态 | `onboarding-page.tsx` | follow-regression | 当前无独立空态，展示完成提示 | 跟随链路完成语义回归验证 | 切片内保留（`onboarding-page.tsx`） |

### 2.3 说明文案（Hint / Description Copy）

| Surface | 位置 | 分类 | 当前语义 | 需改为 | 最终承接位 |
|---|---|---|---|---|---|
| Onboarding WelcomeStep 顶部引导说明 | `onboarding-page.tsx:L276-289` | must-change | "帮助你登记软件资产、记录决策并追踪复用反馈" / "每一步只需填写最小必填字段，其余可在后续补充" | 保留 onboarding 引导语气，但把"软件资产"改为与四实体冻结语义一致的表达，不再把四实体解释成 CRUD 步骤 | 切片内保留（`onboarding-page.tsx`），文案引用共享语义候选 |
| Onboarding WelcomeStep 实体介绍 | `onboarding-page.tsx:L281-287` | must-change | "创建一个产品（Product）" / "创建一个仓库（Repository）" / "创建一个模块（Module）" / "记录一条决策（Decision）" | "登记一个经营目标与交付容器（Product）" / "登记一个代码仓库身份对象与项目锚点（Repository）" / "登记一个可复用能力资产（Module）" / "记录一条规则、约束、选择与依据（Decision）" | 切片内保留（`onboarding-page.tsx`），文案引用共享语义候选 |
| Dashboard Current Focus 标题 | `current-focus-section.tsx` | no-change | "Current Focus" | 区块标题不直接承载实体语义，内容跟随 primary owner | 切片内保留（`current-focus-section.tsx`） |
| Dashboard Asset Feedback 标题 | `asset-feedback-section.tsx` | no-change | "Asset Feedback" | 区块标题不直接承载实体语义，内容跟随 primary owner | 切片内保留（`asset-feedback-section.tsx`） |
| Dashboard Recent Activity 标题 | `recent-activity-section.tsx` | no-change | "Recent Activity" | 区块标题不直接承载实体语义，内容跟随 primary owner | 切片内保留（`recent-activity-section.tsx`） |
| Dashboard Current Focus 说明文案 | `current-focus-section.tsx` | no-change | 当前无独立说明文案 | 保持 no-change；若后续新增说明文案，必须跟随 current focus / review 冻结语义 | 切片内保留（`current-focus-section.tsx`） |
| Dashboard Asset Feedback 说明文案 | `asset-feedback-section.tsx` | no-change | 当前无独立说明文案 | 保持 no-change；若后续新增说明文案，必须跟随 feedback / module 冻结语义 | 切片内保留（`asset-feedback-section.tsx`） |
| Dashboard Recent Activity 说明文案 | `recent-activity-section.tsx` | no-change | 当前无独立说明文案 | 保持 no-change；若后续新增说明文案，必须跟随 activity / decision 冻结语义 | 切片内保留（`recent-activity-section.tsx`） |
| Daily Review 页面头部说明 | `daily-review-page.tsx` | no-change | "今日要处理的焦点与决策" | review 会话语义，非实体定义语义 | 切片内保留（`daily-review-page.tsx`） |
| Weekly Review 页面头部说明 | `weekly-review-page.tsx` | no-change | "本周资产盘点与整理" | review 会话语义，非实体定义语义 | 切片内保留（`weekly-review-page.tsx`） |
| Product Detail 页面级语义导语 | `product-detail-page.tsx` | must-change | 当前无统一页面级语义导语 | 增加一句与 ProductSummaryCard 配套的页面级导语，明确“经营目标与交付容器” | 切片内保留（`product-detail-page.tsx`），短语引用共享语义来源 |
| Repository Detail 页面级语义导语 | `repository-binding-detail-page.tsx` | must-change | 当前无统一页面级语义导语 | 增加一句与 RepositorySummaryCard 配套的页面级导语，明确“代码仓库身份对象与项目锚点” | 切片内保留（`repository-binding-detail-page.tsx`），短语引用共享语义来源 |
| Module Detail 页面级语义导语 | `module-detail-page.tsx` | must-change | 当前无统一页面级语义导语 | 增加一句与 ModuleSummaryCard / ModuleNextActionBar 配套的页面级导语，明确“可复用能力资产” | 切片内保留（`module-detail-page.tsx`），短语引用共享语义来源 |
| Decision Detail 页面级语义导语 | `decision-detail-page.tsx` | must-change | 当前无统一页面级语义导语 | 增加一句与 DecisionDetailSummaryCard 配套的页面级导语，明确“规则、约束、选择与依据的索引对象” | 切片内保留（`decision-detail-page.tsx`），短语引用共享语义来源 |
| Product Detail 描述性标题 | `product-detail-page.tsx` | no-change | 标题由产品名称直接承接 | 标题来自数据，不独立定义 Product 语义 | 切片内保留（`product-detail-page.tsx`） |
| Repository Detail 描述性标题 | `repository-binding-detail-page.tsx` | no-change | 标题由仓库名称直接承接 | 标题来自数据，不独立定义 Repository 语义 | 切片内保留（`repository-binding-detail-page.tsx`） |
| Module Detail 描述性标题 | `module-detail-page.tsx` | no-change | 标题由模块名称直接承接 | 标题来自数据，不独立定义 Module 语义 | 切片内保留（`module-detail-page.tsx`） |
| Decision Detail 描述性标题 | `decision-detail-page.tsx` | no-change | 标题由决策标题直接承接 | 标题来自数据，不独立定义 Decision 语义 | 切片内保留（`decision-detail-page.tsx`） |
| Product Detail 引导文案 | `product-detail-page.tsx` | no-change | 来源模板块使用 "来源：..." 和 `templateDescription` 解释来源上下文 | 维持 no-change；其语义属于来源模板上下文，不替代 Product 冻结语义 | 切片内保留（`product-detail-page.tsx`） |
| Repository Detail 引导文案 | `repository-binding-detail-page.tsx` | no-change | "来源产品：" / "来源模块：" 承接来源上下文 | 维持 no-change；其语义属于来源上下文，不替代 Repository 冻结语义 | 切片内保留（`repository-binding-detail-page.tsx`） |
| Module Detail 引导文案 | `module-detail-page.tsx` | no-change | 当前无独立引导文案 | 保持 no-change；若后续新增引导文案，必须复用 ModuleSummaryCard 已冻结语义 | 切片内保留（`module-detail-page.tsx`） |
| Decision Detail 引导文案 | `decision-detail-page.tsx` | no-change | 当前无独立引导文案 | 保持 no-change；若后续新增引导文案，必须复用 DecisionDetailSummaryCard 已冻结语义 | 切片内保留（`decision-detail-page.tsx`） |

### 2.4 下一步动作说明（Next-Action Descriptions）

| Surface | 位置 | 分类 | 当前语义 | 需改为 | 最终承接位 |
|---|---|---|---|---|---|
| ModuleNextActionBar — 绑定产品 | `module-next-action-bar.tsx:L39-40` | must-change | "将模块绑定到所属产品，建立模块归属关系" | "将可复用能力资产绑定到所属经营目标（Product），建立资产归属关系" | 切片内保留（`module-next-action-bar.tsx`），文案引用共享语义候选 |
| ModuleNextActionBar — 映射仓库 | `module-next-action-bar.tsx:L63-64` | must-change | "将模块映射到代码仓库，建立可追溯关联" | "将可复用能力资产映射到代码仓库身份对象与项目锚点（Repository），建立可追溯关联" | 切片内保留（`module-next-action-bar.tsx`），文案引用共享语义候选 |
| ModuleNextActionBar — 查看决策 | `module-next-action-bar.tsx:L88-89` | must-change | "继续处理与模块相关的决策记录" | "继续处理与该可复用能力资产相关的规则与决策记录" | 切片内保留（`module-next-action-bar.tsx`），文案引用共享语义候选 |
| ModuleNextActionBar — 兜底完成 | `module-next-action-bar.tsx:L111` | must-change | "模块结构已完整" | "该可复用能力资产的关系已完整" | 切片内保留（`module-next-action-bar.tsx`），文案引用共享语义候选 |
| DecisionDetailSummaryCard — 状态推进 CTA | `decision-detail-summary-card.tsx` 状态推进区 | no-change | 状态标签如 "推进到 xx" | 状态推进是业务动作，非实体定义语义 | 切片内保留（`decision-detail-summary-card.tsx`） |
| DashboardPrimaryActionPanel | `dashboard-primary-action-panel.tsx` | no-change | "Daily Review" / "Weekly Review" | review 入口，非实体定义语义 | 切片内保留（`dashboard-primary-action-panel.tsx`） |
| OnboardingCtaButton | `onboarding-cta-button.tsx` | no-change | "开始首轮录入" | onboarding 入口，非实体定义语义 | 切片内保留（`onboarding-cta-button.tsx`） |
| ReviewActionFooter | `review-action-footer.tsx` | no-change | "完成 Review" | review 完成动作，非实体定义语义 | 切片内保留（`review-action-footer.tsx`） |

---

## 三、共享语义来源 vs 切片内渲染设计

### 3.1 共享语义来源候选

> 这里冻结的是**共享语义来源**，不是共享渲染组件。也就是说，短语与单行定义未来统一收口到 `frontend/src/features/project-context/`，但具体页面怎么排版、放在什么位置，仍由各切片渲染。

| 共享语义来源内容 | 复用面 | 到达条件 | 唯一共享语义来源 | 渲染承接方式 |
|---|---|---|---|---|
| 四实体冻结语义单行定义 | 4 个 Detail Page + 4 个 Summary Card + Onboarding + Dashboard 均需引用 | `3+` 页面已稳定复用 | `frontend/src/features/project-context/` | 各切片页面/组件本地渲染，共享的只是单行定义本身 |
| Module 的"可复用能力资产"短语 | ModuleDetailPage + ModuleSummaryCard + ModuleNextActionBar + Onboarding + Dashboard 均需 | 5 个位置复用 | `frontend/src/features/project-context/` | 各切片继续在自己的 page/component 中渲染，但引用同一短语来源 |
| Decision 的"规则、约束、选择与依据"短语 | DecisionDetailPage + DecisionDetailSummaryCard + Onboarding + Dashboard + Review 均需 | 5 个位置复用 | `frontend/src/features/project-context/` | 各切片继续在自己的 page/component 中渲染，但引用同一短语来源 |
| Repository 的"代码仓库身份对象与项目锚点"短语 | RepositoryDetailPage + RepositorySummaryCard + Onboarding + RepositoryListPage 空态 均需 | 4 个位置复用 | `frontend/src/features/project-context/` | 各切片继续在自己的 page/component 中渲染，但引用同一短语来源 |

**边界约束**：
- 当前阶段已冻结未来唯一共享承接位为 `frontend/src/features/project-context/`，但不要求 `phase12-04` 立即创建具体共享语义模块（该工作属于 `phase12-06` 读路径 owner 设计）；
- `phase12-04` 可在设计文本中引用冻结语义文案，但不得把共享呈现承接位重新打开成“常量或 adapter”等多选项；
- 不得在 `phase12-04` 中越权定义新的跨切片只读数据 owner。

### 3.2 切片内保留的渲染与结构

| 切片内保留对象 | 保留在切片内 | 理由 | 与共享语义来源的关系 |
|---|---|---|---|
| ProductSummaryCard 的布局、字段顺序与局部标签 | `product-registry/components/` | Product 卡片结构是切片专属 UI | 可引用共享短语，但渲染结构不共享 |
| RepositorySummaryCard 的布局、字段顺序与局部标签 | `repository-binding/components/` | Repository 卡片结构是切片专属 UI | 可引用共享短语，但渲染结构不共享 |
| ModuleSummaryCard 与 ReuseSummaryInline 的组合结构 | `module-registry/components/` | Module 详情与复用摘要组合依赖切片语境 | 可引用共享短语，但渲染结构不共享 |
| DecisionDetailSummaryCard 的结构化字段区块 | `decision-center/components/` | Decision 详情字段结构是切片专属 UI | 可引用共享短语，但结构化布局不共享 |
| 四个 Detail Page 的页面级导语插入位置 | 各 feature 的 `pages/` | 页面顶部布局与来源块结构各不相同 | 导语短语引用共享语义来源，但插入位置由页面本地决定 |
| 各 List 页空态与来源提示的渲染位置 | 各 feature 的 `pages/` | 列表页虽非 primary owner，但空态/提示位点由页面本身承接 | 空态短语可引用共享语义来源，但页面仍本地渲染 |
| Dashboard 区块内容 | `dashboard/` | Dashboard 为跟随回归页，内容跟随 primary owner | 若引用四实体短语，只能消费共享语义来源，不在 dashboard 内自造 |

---

## 四、Before / After 表达样例

### 4.1 页面级样例：ModuleSummaryCard

**Before**（当前 `module-summary-card.tsx`）：
```tsx
<CardContent className="space-y-2">
  <div>
    <p className="text-sm text-muted-foreground">描述</p>
    <p className="text-sm">{module.description}</p>
  </div>
  <div>
    <p className="text-sm text-muted-foreground">创建时间</p>
    <p className="text-sm">{new Date(module.created_at).toLocaleDateString()}</p>
  </div>
</CardContent>
```

**After**：
```tsx
<CardContent className="space-y-2">
  <div className="flex items-center gap-2">
    <span className="text-[10px] font-medium text-muted-foreground uppercase tracking-wide">
      可复用能力资产
    </span>
  </div>
  <div>
    <p className="text-sm text-muted-foreground">能力描述</p>
    <p className="text-sm">{module.description}</p>
  </div>
  <div>
    <p className="text-sm text-muted-foreground">登记时间</p>
    <p className="text-sm">{new Date(module.created_at).toLocaleDateString()}</p>
  </div>
</CardContent>
```

### 4.2 页面级样例：DecisionDetailSummaryCard

**Before**（当前顶部）：
```tsx
<CardHeader>
  <CardTitle className="text-lg">{decision.title}</CardTitle>
  <CardDescription>
    <Badge>{decision.status}</Badge>
  </CardDescription>
</CardHeader>
```

**After**（在 CardHeader 与 CardContent 之间增加语义框架行）：
```tsx
<CardHeader>
  <CardTitle className="text-lg">{decision.title}</CardTitle>
  <CardDescription>
    <Badge>{decision.status}</Badge>
  </CardDescription>
</CardHeader>
<div className="px-6 pb-1">
  <p className="text-[10px] text-muted-foreground uppercase tracking-wide">
    规则、约束、选择与依据的索引对象
  </p>
</div>
<CardContent className="space-y-4">
  {/* 现有结构化字段保持不变 */}
</CardContent>
```

### 4.3 Surface 级样例：Onboarding WelcomeStep 实体介绍

**Before**（`onboarding-page.tsx:L281-287`）：
```tsx
<ol className="list-decimal list-inside space-y-1 text-muted-foreground">
  <li>创建一个产品（Product）</li>
  <li>创建一个仓库（Repository）</li>
  <li>创建一个模块（Module）</li>
  <li>记录一条决策（Decision）</li>
</ol>
```

**After**：
```tsx
<ol className="list-decimal list-inside space-y-1 text-muted-foreground">
  <li>登记一个经营目标与交付容器（Product）</li>
    <li>登记一个代码仓库身份对象与项目锚点（Repository）</li>
  <li>登记一个可复用能力资产（Module）</li>
  <li>记录一条规则、约束、选择与依据（Decision）</li>
</ol>
```

### 4.4 Surface 级样例：ModuleNextActionBar 下一步动作说明

**Before**（`module-next-action-bar.tsx:L39-40`）：
```tsx
<p className="text-xs text-muted-foreground truncate">
  将 "{moduleName}" 绑定到所属产品，建立模块归属关系
</p>
```

**After**：
```tsx
<p className="text-xs text-muted-foreground truncate">
  将可复用能力资产 "{moduleName}" 绑定到所属经营目标（Product），建立资产归属关系
</p>
```

---

## 五、明确不做清单

| # | 不做事项 | 理由 |
|---|---|---|
| 1 | 不重写读路径 owner（`use-*-read.ts`） | 属于 phase12-06 范围 |
| 2 | 不新增后端合同、共享只读服务或 `.proto` 变更 | 属于 phase12-07 范围 |
| 3 | 不创建 `frontend/src/features/project-context/` 的具体共享语义承接模块 | 属于 phase12-06 范围；当前阶段只冻结未来唯一共享承接位，不提前实现模块细节 |
| 4 | 不进行页面重组或结构性路由改造 | 不属于 phase12 范围 |
| 5 | 不新增写回通道、MCP、CLI 或 agent 写回 | 超出 phase12 范围 |
| 6 | 不修改实体 schema 或关系主线 | 超出 phase12 范围 |
| 7 | 不新增第二套 canonical API 或影子状态表 | 超出 phase12 范围 |
| 8 | 不改写 List pages 的 owner 身份与页面主结构（列表页继续作为非 primary owner） | List pages 仍不承担 primary owner 职责，但允许按 `§1.4` / `§2.2` 调整空态与来源提示等页面内 surface，使其回到冻结语义 |

---

## 六、分类汇总

### 6.1 对象级分类

| 分类 | 数量 | 对象 |
|---|---|---|
| must-change | 14 | P1-P4（四个 primary owner detail page）、P6（OnboardingPage）、C1-C4（四个 SummaryCard）、C9（ModuleNextActionBar）、N1-N4（四个列表页文件） |
| follow-regression | 13 | P5、P7-P8（3 个页面）、C5-C8（4 个组件）、N5-N10（6 个默认非 primary owner 对象） |
| no-change | 11 | R1-R8（8 个 Route）、C10-C12（3 个动作/入口组件） |

### 6.2 Surface 级分类汇总

| Surface 类型 | must-change | follow-regression | no-change |
|---|---|---|---|
| 摘要卡片 | 4（Product/Repository/Module/Decision） | 0 | 0 |
| 空态 | 4（四实体 List 页空态） | 15（Dashboard 3 项 / Daily Review 3 项 / Weekly Review 3 项 / Onboarding 6 项） | 0 |
| 说明文案 | 6（WelcomeStep 顶部引导说明 / WelcomeStep 实体介绍 / 四个 Detail 页页面级语义导语） | 0 | 16（Dashboard 6 项 / Review 2 项 / Detail 标题 4 项 / Detail 既有引导文案 4 项） |
| 下一步动作说明 | 4（ModuleNextActionBar 四个场景） | 0 | 4（Decision CTA/DashboardPrimaryAction/OnboardingCta/ReviewActionFooter） |
