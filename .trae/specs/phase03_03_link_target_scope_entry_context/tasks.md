# Tasks

- [x] Task 1: 冻结 `LinkDecisionToTarget` 的当前阶段目标范围。将当前阶段允许正式写入的目标类型写成单值结论，避免后续接口与合同继续扩 scope。
  - [x] SubTask 1.1: 明确 `Module` 是当前阶段唯一必交付的可写入目标类型
  - [x] SubTask 1.2: 明确当前阶段不引入其他新目标类型

- [x] Task 2: 冻结 `Decision -> Module` 的直接闭环。把候选读取、目标选择、关联写入与读模型回流之间的关系写清。
  - [x] SubTask 2.1: 明确 `Decision Detail` 中必须完成 `Decision -> Module` 候选读取、目标选择与关联写入
  - [x] SubTask 2.2: 明确关联结果必须回流到 `Decision Detail` 已关联目标区
  - [x] SubTask 2.3: 明确关联结果必须能被 `Decision List` 的 `link_count / linked_module_summary` 消费

- [x] Task 3: 冻结 `Product / Repository` 的受控连接位解释。确保它们保留未来扩展位，但不提前扩写为当前阶段第二主线。
  - [x] SubTask 3.1: 明确 `Product / Repository` 当前只保留合同保留位、轻量候选读取前提或未来扩展语义
  - [x] SubTask 3.2: 明确当前阶段不实现 `Decision -> Product` 或 `Decision -> Repository` 正式写入闭环
  - [x] SubTask 3.3: 明确不为 `Product / Repository` 扩写新的页面主线、写入动作或验收主线

- [x] Task 4: 冻结 `Module Detail` 发起记录决策的入口上下文。把从 `Module Detail` 到 `Decision Create` 的最小上下文传递规则写成单值结论。
  - [x] SubTask 4.1: 明确 `Module Detail` 的记录决策触点单值指向带上下文的 `Decision Create`
  - [x] SubTask 4.2: 明确入口上下文至少携带当前 `Module` 的目标标识与可展示名称
  - [x] SubTask 4.3: 明确入口上下文只作为预填来源与候选建议依据，不等于正式关联结果

- [x] Task 5: 冻结 `Decision Create` 与 `Decision Detail` 的上下文承接边界。确保入口上下文与正式写入不会被混淆，且创建成功后入口上下文能无歧义地继续承接。
  - [x] SubTask 5.1: 明确从 `Module Detail` 带上下文进入时，`Decision Create` 必须展示来源 `Module` 信息
  - [x] SubTask 5.2: 明确从列表直接进入时，`Decision Create` 必须承接“无特定来源目标”的语义
  - [x] SubTask 5.3: 明确正式 `LinkDecisionToTarget` 写入只能在 `Decision Detail` 中完成
  - [x] SubTask 5.4: 明确预填来源在正式写入前不得计入 `Decision List` 的已关联统计
  - [x] SubTask 5.5: 明确从 `Module Detail` 带上下文创建成功后，必须默认进入新建 `Decision` 的 `Decision Detail`，不得回流到 `Decision Center / List`
  - [x] SubTask 5.6: 明确入口上下文中的 `Module` 必须在 `Decision Detail` 中作为显式待关联目标继续承接
  - [x] SubTask 5.7: 明确在候选读取面板中，该 `Module` 必须作为首选候选或显式待确认目标出现
  - [x] SubTask 5.8: 明确该显式待关联状态必须持续到用户完成正式 `LinkDecisionToTarget` 或主动放弃关联

- [x] Task 6: 完成规格校验。确认本次 `phase03-03` 规格可直接作为后续接口、合同、状态流与验收设计的上游。
  - [x] SubTask 6.1: 验证 `Decision -> Module` 关联路径已经明确
  - [x] SubTask 6.2: 验证 `Module Detail` 与 `Decision Center` 的交互归属已经明确
  - [x] SubTask 6.3: 验证 `Product / Repository` 只保留受控连接位，未扩写为第二主线
  - [x] SubTask 6.4: 验证入口上下文与正式关联结果的边界已经明确
  - [x] SubTask 6.5: 验证创建成功回流路径与入口上下文在 `Decision Detail` 中的继续承接已经单值化

# Task Dependencies

- `Task 2` depends on `Task 1`
- `Task 3` depends on `Task 1`
- `Task 4` depends on `Task 1`
- `Task 5` depends on `Task 2` and `Task 4`
- `Task 6` depends on `Task 1`, `Task 2`, `Task 3`, `Task 4`, and `Task 5`
