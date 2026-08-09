# Tasks

- [x] Task 1: 冻结 Dashboard 当前阶段最小聚合读范围。把 `DashboardOverviewRead / FeedbackSignalRead / RecentActivityRead` 作为唯一聚合读集合写成单值结论。
  - [x] SubTask 1.1: 明确当前阶段只承接三类聚合读取，不新增趋势分析、通知中心、导出或写入接口
  - [x] SubTask 1.2: 明确 `DashboardOverviewRead` 负责概览与系统状态判定，`FeedbackSignalRead` 负责反馈主队列与补充摘要，`RecentActivityRead` 负责独立活动流
  - [x] SubTask 1.3: 明确 Dashboard 当前阶段不承接新的业务写入接口

- [x] Task 2: 冻结三类聚合读取的最小接口边界。把输出字段范围、输入参数边界与最大返回数量写成单值结论。
  - [x] SubTask 2.1: 明确 `DashboardOverviewRead` 至少承接六个概览计数，并禁止混入反馈信号与最近活动结果
  - [x] SubTask 2.2: 明确 `FeedbackSignalRead` 只消费 `pending_decision_signals + product_asset_coverage`，并承接主队列最多 `5` 条、补充摘要最多 `3` 条
  - [x] SubTask 2.3: 明确 `RecentActivityRead` 至少承接活动类型、显式活动时间与目标跳转信息，并最多返回 `10` 条
  - [x] SubTask 2.4: 明确三类读取当前阶段都不引入业务筛选、分页、时间范围或排序切换参数

- [x] Task 3: 冻结空系统、无信号、无活动、局部失败与整页失败的错误语义前提。把成功空态与真实失败区分清楚。
  - [x] SubTask 3.1: 明确冷启动空系统与“已有结构化资产但仍无 Module”是两类不同的成功状态
  - [x] SubTask 3.2: 明确无信号、`Asset Feedback` 成功空态与无活动都返回成功语义，不映射为资源不存在或接口失败
  - [x] SubTask 3.3: 明确 `DashboardOverviewRead` 失败触发整页失败
  - [x] SubTask 3.4: 明确 `FeedbackSignalRead` 或 `RecentActivityRead` 失败只触发局部失败，不强制整页失败
  - [x] SubTask 3.5: 明确超出冻结边界的非法参数必须返回校验失败语义
  - [x] SubTask 3.6: 明确“无缺口且无活动”属于非空成功态，不与冷启动空系统混同，也不回退到创建导向 CTA
  - [x] SubTask 3.7: 明确 `FeedbackSignalRead` 局部失败且 1-4 条 overview 条件未命中时，必须抑制强制主 CTA 并降级为区块内重试入口

- [x] Task 4: 冻结 Dashboard 主 CTA 优先级矩阵与单主 CTA 约束。把冷启动、缺口态、待决策与系统已就绪状态的命中顺序写成单值结论。
  - [x] SubTask 4.1: 明确冷启动空系统、非空但无 Module、缺 Product、缺 Repository 四类状态优先于反馈信号命中
  - [x] SubTask 4.2: 明确 `pending_decision_signals` 与三类产品资产缺口的主 CTA 命中顺序
  - [x] SubTask 4.3: 明确“无缺口且有活动数据”进入中性状态，不再展示强制主 CTA
  - [x] SubTask 4.4: 明确同一时刻只允许存在一个主 CTA，其他动作降级为区块内次级入口
  - [x] SubTask 4.5: 明确“无缺口且无活动”同样进入非空中性状态，不展示强制主 CTA

- [x] Task 5: 冻结 `phase05-04` 的非目标边界。确认本规格不提前抢 `phase05-06 / 07 / 08` 的职责。
  - [x] SubTask 5.1: 明确不提前冻结前端页面级状态机、交互细节与路由恢复策略，由 `phase05-06` 承接
  - [x] SubTask 5.2: 明确不提前冻结后端模块边界、接口分组、query service owner 与 DTO 映射策略，由 `phase05-07` 承接
  - [x] SubTask 5.3: 明确不提前冻结 `.proto` 服务命名、包名版本、消息编号与过渡传输层映射，由 `phase05-08` 承接
  - [x] SubTask 5.4: 明确外部埋点、趋势分析、通知中心与超出当前阶段的聚合接口仍属非目标

- [x] Task 6: 完成 `phase05-04` 规格一致性校验。确认本次冻结与 `phase05` 三件套、`phase05-02 / 03` 及后续任务分工保持一致。
  - [x] SubTask 6.1: 验证三类聚合读取与 `shared_baseline.md`、`architecture_plan.md` 的读取归属保持一致
  - [x] SubTask 6.2: 验证空系统、非空但无 Module 与主 CTA 优先级矩阵和共享基线一致
  - [x] SubTask 6.3: 验证 Dashboard 仍只承接读与跳转，不形成新的业务写入接口
  - [x] SubTask 6.4: 验证本规格未越界冻结 `phase05-06 / 07 / 08` 的职责

# Task Dependencies

- `Task 2` depends on `Task 1`
- `Task 3` depends on `Task 1`, `Task 2`
- `Task 4` depends on `Task 1`, `Task 2`, `Task 3`
- `Task 5` depends on `Task 1`, `Task 2`
- `Task 6` depends on `Task 1`, `Task 2`, `Task 3`, `Task 4`, `Task 5`
