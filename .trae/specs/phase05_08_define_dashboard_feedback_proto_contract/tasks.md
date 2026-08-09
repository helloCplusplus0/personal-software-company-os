# Tasks

- [x] Task 1: 对齐 `phase05-08` 直接上游与现有 proto workspace 基线，确定本次合同源只在单一 `proto/` 主线内扩展。
  - [x] SubTask 1.1: 对齐 `phase05-02 / 03 / 04 / 07` 已冻结的反馈字段模板、导航语义、读组边界与后端 owner
  - [x] SubTask 1.2: 对齐 `phase04-11` 已落地的 proto 主线模式、`proto/README.md`、`buf.yaml`、`buf.gen.yaml` 与 `proto/Makefile`
  - [x] SubTask 1.3: 冻结 `Dashboard + Feedback` 只新增 `proto/psco/dashboard/v1/dashboard.proto`，不新增第二套 proto workspace

- [x] Task 2: 冻结 `DashboardService` 的最小 RPC 矩阵、包名版本与文件落点。
  - [x] SubTask 2.1: 冻结包名为 `psco.dashboard.v1`，文件落点为 `proto/psco/dashboard/v1/dashboard.proto`
  - [x] SubTask 2.2: 冻结单一 `DashboardService`，承接 `GetDashboardOverview / GetFeedbackSignals / GetRecentActivities`
  - [x] SubTask 2.3: 冻结三个 request 的最小边界为无筛选、无分页、无排序切换
  - [x] SubTask 2.4: 冻结 Go / TypeScript 生成产物落点与现有 proto 主线保持同构

- [x] Task 3: 冻结 `DashboardOverviewRead` 的消息结构与字段编号方案。
  - [x] SubTask 3.1: 冻结 `DashboardOverview` 至少承接六个概览计数字段
  - [x] SubTask 3.2: 冻结 `GetDashboardOverviewResponse` 的最小返回边界，不混入反馈或活动字段
  - [x] SubTask 3.3: 冻结 `DashboardOverview` 与相关 request / response 的具体字段编号表，确保 `.proto` 落地可机械映射

- [x] Task 4: 冻结 `FeedbackSignalRead` 的统一卡片模型、资产缺口摘要与枚举体系。
  - [x] SubTask 4.1: 冻结 `FeedbackSignal` 最小字段模板，保持 `signal_family / signal_code / priority / title / summary / action_label / target_type / target_id / target_label`
  - [x] SubTask 4.2: 冻结 `FeedbackSignalFamily / FeedbackSignalCode / FeedbackSignalPriority / DashboardTargetType` 的最小枚举语义与具体编号表
  - [x] SubTask 4.3: 冻结 `GetFeedbackSignalsResponse` 的两层结构：`current_focus_signals` + `asset_feedback_summary`
  - [x] SubTask 4.4: 冻结 `ProductAssetCoverageSummary` 的最小字段、代表项上限与具体字段编号表
  - [x] SubTask 4.5: 冻结 `missing both bindings` 必须同时拥有独立 `signal_code` 与独立计数字段，不得退回隐式组合表达

- [x] Task 5: 冻结 `RecentActivityRead` 的活动项模型、类型枚举、显式时间字段与导航解释。
  - [x] SubTask 5.1: 冻结 `RecentActivityItem` 最小字段模板
  - [x] SubTask 5.2: 冻结 `RecentActivityType` 的最小枚举语义与具体编号表，去除笼统 `binding`
  - [x] SubTask 5.3: 冻结 `activity_at` 为唯一显式活动时间字段与排序锚点
  - [x] SubTask 5.4: 冻结 `target_type / target_id` 对 `Module / Product / Repository / Decision` canonical 落点的解释
  - [x] SubTask 5.5: 冻结 `GetRecentActivitiesRequest / Response` 的具体字段编号表，并明确最多返回 `10` 条活动项

- [x] Task 6: 冻结 `.proto` 与 `chi + JSON HTTP` 的过渡映射策略。
  - [x] SubTask 6.1: 冻结 RPC → HTTP 映射矩阵：`/api/dashboard/overview`、`/api/dashboard/feedback-signals`、`/api/dashboard/recent-activities`
  - [x] SubTask 6.2: 冻结当前阶段错误语义由传输层映射承接，不在 `.proto` 中发明第二套错误包络
  - [x] SubTask 6.3: 冻结手写 JSON DTO / 前端 adapter 只能从 `.proto` 单向承接，不得形成第二套合同源

- [x] Task 7: 冻结合同演进、`reserved` 与 breaking check 前提。
  - [x] SubTask 7.1: 冻结当前版本字段编号表、枚举编号表与后续新增字段必须递增的规则
  - [x] SubTask 7.2: 冻结删除字段、删除枚举值后的 `reserved` 约束
  - [x] SubTask 7.3: 冻结 `buf build / lint / generate / breaking` 为最小校验链
  - [x] SubTask 7.4: 冻结 `buf breaking --against '../.git#branch=main,subdir=proto'` 为唯一 breaking 基准

- [x] Task 8: 完成 `phase05-08` 规格一致性校验。
  - [x] SubTask 8.1: 验证 `.proto` 合同设计不晚于正式规格正文进入阶段主线
  - [x] SubTask 8.2: 验证字段语义、字段编号与 `phase05-02 / 03 / 04` 单值一致
  - [x] SubTask 8.3: 验证服务接口矩阵与 `phase05-07` 后端读组边界一致
  - [x] SubTask 8.4: 验证 `FeedbackSignal` 最小字段、优先级枚举与 `RecentActivityItem` 类型枚举已单值化
  - [x] SubTask 8.5: 验证 `activity_at`、`missing both bindings`、HTTP 映射与 `reserved / breaking` 前提已单值化

# Task Dependencies

- `Task 2` depends on `Task 1`
- `Task 3` depends on `Task 1` and `Task 2`
- `Task 4` depends on `Task 1` and `Task 2`
- `Task 5` depends on `Task 1` and `Task 2`
- `Task 6` depends on `Task 2`, `Task 4`, and `Task 5`
- `Task 7` depends on `Task 1`, `Task 2`, `Task 3`, `Task 4`, `Task 5`, and `Task 6`
- `Task 8` depends on `Task 1` through `Task 7`
