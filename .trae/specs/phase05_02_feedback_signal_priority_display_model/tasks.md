# Tasks

- [x] Task 1: 冻结 `Feedback` 当前阶段的单值语义。把 `Feedback` 明确为派生信号层，而不是独立重实体、人工录入流或完整评分系统。
  - [x] SubTask 1.1: 明确 `Feedback` 来源于既有结构化对象与绑定关系
  - [x] SubTask 1.2: 明确 `Feedback` 至少回答“哪里不完整 / 现在去哪里补”
  - [x] SubTask 1.3: 明确当前阶段不引入独立 `Feedback` 表、人工录入或复杂评分模型

- [x] Task 2: 冻结 `pending_decision_signals` 与 `product_asset_coverage` 的 Dashboard 消费语义。把两类反馈源如何进入 `Current Focus / Next Action` 与 `Asset Feedback` 写成单值结论。
  - [x] SubTask 2.1: 明确 `pending_decision_signals` 先归一化为 `Feedback Signal Card`，再进入主队列
  - [x] SubTask 2.2: 明确 `product_asset_coverage` 的高优先级缺口项可进入主队列
  - [x] SubTask 2.3: 明确 `Asset Feedback` 只承接 `product_asset_coverage` 的补充摘要
  - [x] SubTask 2.4: 明确 `pending_decision_signals` 的空状态语义
  - [x] SubTask 2.5: 明确 `product_asset_coverage` 的空状态语义：无缺口时返回完整读模型结构、三类缺口计数归零、代表项为空列表、成功空态与局部读取失败显式区分

- [x] Task 3: 冻结 `product missing both bindings` 的独立读模型语义。把双缺口产品的计数、代表项和与单缺口类型的关系写成单值结论。
  - [x] SubTask 3.1: 明确双缺口产品必须单独归类为 `product missing both bindings`
  - [x] SubTask 3.2: 明确 `product_asset_coverage` 必须显式承接双缺口计数
  - [x] SubTask 3.3: 明确双缺口代表项允许以独立 `signal_code` 进入反馈卡片
  - [x] SubTask 3.4: 明确双缺口代表项不得重复计入两个单缺口代表项语义

- [x] Task 4: 冻结统一 `Feedback Signal Card` 的最小字段模板、优先级与排序前提。确保后续前端、后端与 `.proto` 使用同一套反馈卡片模型。
  - [x] SubTask 4.1: 明确反馈卡片最小字段集合
  - [x] SubTask 4.2: 明确优先级顺序固定为 `pending_decision_signals > missing both bindings > missing repository binding > missing module binding`
  - [x] SubTask 4.3: 明确同优先级内默认按“最近需要处理时间优先”，缺省回退为 `created_at DESC`

- [x] Task 5: 冻结 `Current Focus / Next Action` 与 `Asset Feedback` 的最大展示数量。把主队列与补充摘要的展示上限写成单值结论。
  - [x] SubTask 5.1: 明确 `Current Focus / Next Action` 最多展示 `5` 条主卡片
  - [x] SubTask 5.2: 明确 `Asset Feedback` 最多展示 `3` 条代表性缺口项
  - [x] SubTask 5.3: 明确超过上限的信号不得演变为第二条主队列

- [x] Task 6: 冻结 `recent_activity_feed` 与 `dashboard_overview` 的最小展示模型。把活动流与概览区块的最小字段和职责解释写成单值结论。
  - [x] SubTask 6.1: 明确 `recent_activity_feed` 的最小活动项字段集合
  - [x] SubTask 6.2: 明确 `recent_activity_feed` 最多展示 `10` 条活动项
  - [x] SubTask 6.3: 明确 `dashboard_overview` 的最小概览字段集合
  - [x] SubTask 6.4: 明确 `dashboard_overview` 只服务概览区块，不进入反馈优先级队列

- [x] Task 7: 冻结 `recent_activity_feed` 的显式排序字段与时间语义。避免后续实现依赖隐式时间字段或数据库默认顺序。
  - [x] SubTask 7.1: 明确活动流必须使用显式活动时间字段（如 `activity_at`）
  - [x] SubTask 7.2: 明确活动流默认按活动时间倒序排序
  - [x] SubTask 7.3: 明确不得依赖隐式 `created_at` 或默认数据库顺序推断最近活动

- [x] Task 8: 完成 `phase05-02` 规格一致性校验。确认本次规格与 `phase05` 三件套、`phase05-01` 页面边界以及 `phase01-06 / phase04-10` 的上游共识保持一致。
  - [x] SubTask 8.1: 验证本次规格与 `phase05_dashboard_feedback_foundation_dev_plan.md` 的范围与 DoD 一致
  - [x] SubTask 8.2: 验证本次规格与 `phase05_dashboard_feedback_foundation_shared_baseline.md` 的读模型、优先级与排序前提一致
  - [x] SubTask 8.3: 验证本次规格与 `phase05_dashboard_feedback_foundation_architecture_plan.md` 的四类读取归属与最小反馈主线一致
  - [x] SubTask 8.4: 验证本次规格未超出 `phase05-01` 已冻结的页面边界与 `phase01-06 / phase04-10` 的 MVP 范围

# Task Dependencies

- `Task 2` depends on `Task 1`
- `Task 3` depends on `Task 2`
- `Task 4` depends on `Task 1`, `Task 2`, `Task 3`
- `Task 5` depends on `Task 2`, `Task 4`
- `Task 6` depends on `Task 1`
- `Task 7` depends on `Task 6`
- `Task 8` depends on `Task 1`, `Task 2`, `Task 3`, `Task 4`, `Task 5`, `Task 6`, `Task 7`
