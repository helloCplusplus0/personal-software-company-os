# Tasks

- [x] Task 1: 对齐 `phase08-04` 的直接上游、合同基线与真实 inventory
  - [x] SubTask 1.1: 对齐 `phase08-01 / 02 / 03` 已冻结的范围、入口、动作 owner 与 writeback 边界
  - [x] SubTask 1.2: 消费 `phase08` shared baseline 中的合同、读写边界与 inventory 清单
  - [x] SubTask 1.3: 对齐当前 `Dashboard / Reuse Summary / Decision` 的真实 route、query owner 与 application owner 事实

- [x] Task 2: 冻结 review 最小正式合同边界
  - [x] SubTask 2.1: 明确 `.proto` 继续作为唯一长期合同源，review 继续走 ConnectRPC
  - [x] SubTask 2.2: 明确 review 合同只承接 review context、最小动作命令与可选过程记录
  - [x] SubTask 2.3: 明确不得扩写 `dashboard.proto` / `reuse_summary.proto` 去吞并 review session / review result

- [x] Task 3: 冻结 review read model / write model 的最小边界
  - [x] SubTask 3.1: 明确 daily review 的最小正式消费范围
  - [x] SubTask 3.2: 明确 weekly review 的最小正式消费范围
  - [x] SubTask 3.3: 明确 review read model 是对既有事实源的轻量组合层
  - [x] SubTask 3.4: 明确 review-local write model 不得复制既有实体写模型

- [x] Task 4: 冻结 `phase05` Feedback 与 `phase06` Reuse Awareness 的正式消费边界
  - [x] SubTask 4.1: 明确 `phase05` 当前正式消费边界至少包括 overview / current focus / representative signals / recent activities
  - [x] SubTask 4.2: 明确 `phase06` 当前正式消费边界至少包括 dashboard scope 下的 `module_reuse_summary / capability_summary`
  - [x] SubTask 4.3: 明确 review 不得合并或重写 `FeedbackSignalRead` 与 `ReuseSummaryRead` 的 canonical 职责

- [x] Task 5: 冻结 review 记录的可选性与轻量化边界
  - [x] SubTask 5.1: 明确 `phase08` 不以引入 review record 为前置阻断项
  - [x] SubTask 5.2: 明确无 review record 路径当前只覆盖 `decision handoff / entity handoff`
  - [x] SubTask 5.3: 明确 `next-step result` 若作为正式输出，必须落到轻量 review record
  - [x] SubTask 5.4: 明确若新增 review 记录，其身份只允许是经营回路过程记录
  - [x] SubTask 5.5: 明确若新增 review 记录，其最小字段边界与禁止复制的影子状态

- [x] Task 6: 冻结真实 caller / route / owner inventory 如何进入后续 `/spec`
  - [x] SubTask 6.1: 明确 `DashboardRoute / DashboardHomePage / DashboardPrimaryActionPanel / FeedbackSignalCard / RecentActivityItemCard / DashboardStatBar` 的消费要求
  - [x] SubTask 6.2: 明确 `CurrentFocusSection / AssetFeedbackSection / RecentActivitySection / OnboardingCtaButton / SovereigntyPanel` 的消费要求与负向约束
  - [x] SubTask 6.3: 明确 `useDashboardOverviewRead / useFeedbackSignalsRead / useRecentActivitiesRead / useReuseSummaryRead` 的消费要求
  - [x] SubTask 6.4: 明确 `Decision` 相关 route / page / read owner / write owner 与 `dashboard-source.ts / BackToDashboardButton` 的消费要求
  - [x] SubTask 6.5: 明确后续 `/spec` 必须逐项说明“复用 / 扩展 / 禁止升级”为哪一类

- [x] Task 7: 完成 `phase08-04` 规格自检与一致性校验
  - [x] SubTask 7.1: 校验 `spec.md` 已覆盖合同边界、读写模型边界、复用感知消费、记录模型边界与 inventory 要求
  - [x] SubTask 7.2: 校验规格与 `phase08` 三件套、`phase08-01 ~ 03` 规格及当前源码事实一致
  - [x] SubTask 7.3: 校验本任务未越权冻结 review proto 文件名、RPC 名称、query key 细节或实现目录结构

# Task Dependencies

- Task 2 depends on Task 1
- Task 3 depends on Task 1, Task 2
- Task 4 depends on Task 1, Task 3
- Task 5 depends on Task 2, Task 3
- Task 6 depends on Task 1, Task 3, Task 4
- Task 7 depends on Task 1, Task 2, Task 3, Task 4, Task 5, Task 6
