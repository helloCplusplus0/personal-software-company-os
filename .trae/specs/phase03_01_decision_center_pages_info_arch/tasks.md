# Tasks

- [x] Task 1: 冻结 `Decision Center` 页面边界。将 `Decision Center / List`、`Decision Create`、`Decision Detail` 收敛为当前阶段唯一页面主线，并写成单值结论。
  - [x] SubTask 1.1: 明确列表页承接决策读取、筛选入口、创建入口与进入详情入口
  - [x] SubTask 1.2: 明确创建页只承接 `RecordDecision`
  - [x] SubTask 1.3: 明确详情页承接详情读取、已关联目标展示、`Decision -> Module` 候选读取与 `LinkDecisionToTarget`
  - [x] SubTask 1.4: 明确不得把独立 `AI Assistant` 一级导航纳入 `phase03` 页面主线

- [x] Task 2: 冻结页面跳转关系。把列表、创建、详情与 `Module Detail` 入口之间的最小跳转关系写清，避免后续页面主线漂移。
  - [x] SubTask 2.1: 明确 `Decision Center / List -> Decision Create`
  - [x] SubTask 2.2: 明确 `Decision Center / List -> Decision Detail`
  - [x] SubTask 2.3: 明确 `Decision Detail -> Decision Center / List`
  - [x] SubTask 2.4: 明确 `Module Detail` 的两个单值 `Decision` 入口触点，不保留"或"式双路线
  - [x] SubTask 2.5: 明确记录决策触点单值指向 `带上下文的 Decision Create`
  - [x] SubTask 2.6: 明确查看相关决策触点单值指向 `Decision Center / List`
  - [x] SubTask 2.7: 明确 `Module Detail` 当前阶段只作为轻量跳转或预填入口，不扩写为第二个 `Decision` 工作台

- [x] Task 3: 冻结 `Decision Detail` 中目标候选读取的页面归属。将 `Decision -> Module` 候选读取写成详情页的附属读取能力，避免实现期各自发明工作台边界。
  - [x] SubTask 3.1: 明确候选读取当前阶段只面向 `Module`
  - [x] SubTask 3.2: 明确候选读取属于 `Decision Detail` 附属能力
  - [x] SubTask 3.3: 明确候选读取只服务 `LinkDecisionToTarget`，不扩写为独立浏览工作台

- [x] Task 4: 冻结 `PC / 移动浏览器` 信息密度策略。保持单一 `React Web` 语义，同时明确桌面与窄屏下的信息组织方式。
  - [x] SubTask 4.1: 明确桌面端优先承接较高信息密度
  - [x] SubTask 4.2: 明确移动浏览器采用信息裁剪、垂直重排与分层展示
  - [x] SubTask 4.3: 明确当前阶段不引入第二套移动端 UI、独立 `React Native` 客户端或完整 `PWA`

- [x] Task 5: 冻结 `Decision List / Create / Detail` 最小页面级信息区块。为每个页面写出最小应包含的信息区块组成，使后续前端设计可直接进入实现。
  - [x] SubTask 5.1: 明确 `Decision Center / List` 至少由列表工具栏区、列表内容区与空状态区组成
  - [x] SubTask 5.2: 明确 `Decision Create` 至少由结构化表单区、来源上下文区与提交取消操作区组成
  - [x] SubTask 5.3: 明确 `Decision Detail` 至少由核心字段区、已关联目标区与候选读取及目标关联区组成

- [x] Task 6: 完成规格校验。检查本次 `phase03-01` 规格是否满足进入后续子任务的条件。
  - [x] SubTask 6.1: 验证 `Decision Center` 页面职责已经单值化
  - [x] SubTask 6.2: 验证 `Module Detail` 与 `Decision Center` 的入口归属已经明确
  - [x] SubTask 6.3: 验证无第二套移动端 UI 方案
  - [x] SubTask 6.4: 验证页面边界与 `phase01-06` 正式规格正文、`phase03` 三件套一致

# Task Dependencies

- `Task 2` depends on `Task 1`
- `Task 3` depends on `Task 1`
- `Task 4` depends on `Task 1`
- `Task 5` depends on `Task 1`
- `Task 6` depends on `Task 1`, `Task 2`, `Task 3`, `Task 4`, and `Task 5`
