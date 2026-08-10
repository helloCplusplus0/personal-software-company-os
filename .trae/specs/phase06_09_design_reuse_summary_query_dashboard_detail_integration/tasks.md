# Tasks

- [x] Task 1: 冻结 `ReuseSummaryRead` 的 query owner 与页面作用域。把 Dashboard / Module Detail / Product Detail 的最小查询作用域与 query key 规则写成单值结论。
  - [x] SubTask 1.1: 明确单数 `ReuseSummaryRead` 作为唯一 query owner
  - [x] SubTask 1.2: 明确 Dashboard / Module Detail / Product Detail 三种最小查询作用域
  - [x] SubTask 1.3: 明确 query key 以前缀 `['reuse-summary']` 承接局部失效
  - [x] SubTask 1.4: 明确每页只新增一个页面级 `ReuseSummaryRead` `useQuery`，其成功结果同时承接两类 summary

- [x] Task 2: 冻结 Dashboard 中复用感知的最小集成方式。确保不破坏 `phase05` 既有整页状态模型与 `Asset Feedback` 区块职责。
  - [x] SubTask 2.1: 明确 Dashboard 继续保留 `overview / feedback / recent-activity` 三路查询
  - [x] SubTask 2.2: 明确额外新增 `ReuseSummaryRead` 只服务 `Asset Feedback` 内的复用快照子区域
  - [x] SubTask 2.3: 明确 `Representative Feedback Signals` 与 `Reuse Snapshot` 两个子区域的独立状态模型
  - [x] SubTask 2.4: 明确 `ReuseSummaryRead` 失败不得让 Dashboard 整页回退为 `page-error`

- [x] Task 3: 冻结 Module Detail 与 Product Detail 的最小挂接设计。把 query 编排、挂接位和局部重试语义写清楚。
  - [x] SubTask 3.1: 明确 Module Detail 保留 canonical detail query，并追加当前 `moduleId` 作用域的 `ReuseSummaryRead`
  - [x] SubTask 3.2: 明确 Module Summary 邻近区域为 Module Detail 正式挂接位
  - [x] SubTask 3.3: 明确 Product Detail 保留 canonical detail query，并追加当前 `productId` 作用域的 `ReuseSummaryRead`
  - [x] SubTask 3.4: 明确 ProductModuleBindingPanel 邻近区域为 Product Detail 正式挂接位

- [x] Task 4: 冻结复用统计字段、时间字段、新鲜度与刷新语义。确保复用感知读模型在实现层不再漂移。
  - [x] SubTask 4.1: 明确 `module_reuse_summary / capability_summary` 的最小统计字段与时间字段
  - [x] SubTask 4.2: 明确新鲜度口径固定为“读取时反映最新已提交状态”
  - [x] SubTask 4.3: 明确 Dashboard / Detail 的局部重试只重读当前作用域 `ReuseSummaryRead`
  - [x] SubTask 4.4: 明确允许使用 `placeholderData` 稳定复用快照子区域 UI

- [x] Task 5: 冻结复用反馈与既有反馈信号的边界关系。避免 Dashboard 中出现第二条模糊主线。
  - [x] SubTask 5.1: 明确 `FeedbackSignalRead` 继续只服务 `Current Focus` 与 representative feedback signals
  - [x] SubTask 5.2: 明确 `ReuseSummaryRead` 只服务复用快照子区域
  - [x] SubTask 5.3: 明确复用快照不得升级为新的主 CTA 或一级行动队列

- [x] Task 6: 冻结空状态、零复用状态、有复用状态与 capability 缺省策略。让 UI 差异具有单值口径。
  - [x] SubTask 6.1: 明确 Dashboard 成功空态文案与读取失败态边界
  - [x] SubTask 6.2: 明确 `reuse_product_count <= 1` 的零复用解释语义
  - [x] SubTask 6.3: 明确 `capability_summary` 的系统内置 label 映射与 `empty_state_text` 缺省字段
  - [x] SubTask 6.4: 明确未填写 `capability_key` 的 Module 不参与 capability 聚合，但不算错误态

- [x] Task 7: 冻结当前子任务的非冻结边界。确保 `phase06-09` 只负责页面级 query owner 与挂接设计，不提前替后续合同/后端子任务做决定。
  - [x] SubTask 7.1: 明确当前子任务不冻结后端模块目录、candidate reader 文件名与 platform 装配函数
  - [x] SubTask 7.2: 明确当前子任务不冻结 `.proto` 文件路径、HTTP 路由数量与 query 参数细节
  - [x] SubTask 7.3: 明确后续合同设计只需满足“每页一个页面级 `ReuseSummaryRead` query 同时承接两类 summary”

- [x] Task 8: 完成规格一致性校验。验证本次设计与 `phase06-04 / 05`、shared baseline 和 `phase05` Dashboard 主线保持一致。
  - [x] SubTask 8.1: 验证 `ReuseSummaryRead` 与 `phase06-04` 已冻结读模型口径一致
  - [x] SubTask 8.2: 验证 `query` 只读边界与 `phase06-05` 一致
  - [x] SubTask 8.3: 验证 Dashboard 集成方式不破坏 `phase05` 整页状态与 `Asset Feedback` 区块职责
  - [x] SubTask 8.4: 验证本次规格足以直接指导前端 query、页面挂接与后续 `phase06-10` 合同设计

# Task Dependencies

- `Task 2` depends on `Task 1`
- `Task 3` depends on `Task 1`
- `Task 4` depends on `Task 1`, `Task 2`, and `Task 3`
- `Task 5` depends on `Task 2`
- `Task 6` depends on `Task 1`, `Task 3`, and `Task 4`
- `Task 7` depends on `Task 1`, `Task 4`, and `Task 6`
- `Task 8` depends on `Task 1`, `Task 2`, `Task 3`, `Task 4`, `Task 5`, `Task 6`, and `Task 7`
