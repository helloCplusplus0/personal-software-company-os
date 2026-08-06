# Tasks

- [x] Task 1: 收敛 `Decision List` 的查询承接与页面状态模型。
  - [x] SubTask 1.1: 明确 `queryText / statusFilter` 必须归属于路由搜索参数
  - [x] SubTask 1.2: 明确列表搜索参数的默认来源与刷新恢复规则
  - [x] SubTask 1.3: 明确读取状态与派生视图状态的分层关系
  - [x] SubTask 1.4: 明确列表错误反馈位置、空状态主动作与返回列表后的上下文恢复规则
  - [x] SubTask 1.5: 明确 URL 中无 `queryText / statusFilter` 时的默认值（空字符串 / 全部状态）

- [x] Task 2: 收敛 `Decision Create` 的来源上下文、草稿状态、提交状态与成功回流设计。
  - [x] SubTask 2.1: 明确来源 `Module` 上下文在创建页中的承接方式
  - [x] SubTask 2.2: 明确 `title / context / problem / alternatives / choice / reason / impact / status` 的最小草稿状态与 `idle / dirty` 语义
  - [x] SubTask 2.3: 明确 `submitting / submit-success / submit-error` 的最小提交状态
  - [x] SubTask 2.4: 明确提交失败时保留输入与来源上下文，并在当前表单上下文展示错误
  - [x] SubTask 2.5: 明确提交成功后进入新建 `Decision` 对应的 `DecisionDetailPage`

- [x] Task 3: 收敛 `Decision Detail` 的读取、待关联目标、候选读取与关联动作交互流。
  - [x] SubTask 3.1: 明确 `DecisionDetailPage` 的最小读取状态
  - [x] SubTask 3.2: 明确入口上下文形成的待关联 `Module` 持续承接规则
  - [x] SubTask 3.3: 明确候选读取的最小状态 `pending / ready / empty / error`
  - [x] SubTask 3.4: 明确 `LinkDecisionToTarget` 的最小状态 `idle / submitting / submit-success / submit-error`
  - [x] SubTask 3.5: 明确关联成功后停留详情页、刷新已关联目标结果并清除待关联状态
  - [x] SubTask 3.6: 明确关联失败后保留当前候选选择与错误反馈位置

- [x] Task 4: 收敛页面返回路径与统一列表上下文恢复规则。
  - [x] SubTask 4.1: 明确从 `DecisionCreatePage` 主动返回时回到 `DecisionListPage`，按来源列表上下文存在性决定恢复原参数或落到默认参数
  - [x] SubTask 4.2: 明确从 `DecisionDetailPage` 主动返回时回到 `DecisionListPage`，按来源列表上下文存在性决定恢复原参数或落到默认参数
  - [x] SubTask 4.3: 明确从创建成功后的 `DecisionDetailPage` 返回列表时恢复原有 `queryText / statusFilter`，或不存在来源上下文时落到默认参数
  - [x] SubTask 4.4: 明确不得保留第二套默认回流路径
  - [x] SubTask 4.5: 明确 `DecisionCreatePage` 与 `DecisionDetailPage` 必须持有"来源列表上下文存在/不存在"的最小页面状态，并冻结进入路径与上下文继承规则

- [x] Task 5: 收敛局部 UI 状态归属边界。
  - [x] SubTask 5.1: 明确表单草稿、提交错误与来源上下文状态优先归属于创建页
  - [x] SubTask 5.2: 明确待关联目标、候选读取状态与关联动作状态优先归属于详情页上下文
  - [x] SubTask 5.3: 明确不得默认升级为跨路由全局状态

- [x] Task 6: 完成规格校验。
  - [x] SubTask 6.1: 验证页面级状态模型、状态归属与错误呈现位置已经明确
  - [x] SubTask 6.2: 验证列表、创建、详情、关联之间的交互流已经明确
  - [x] SubTask 6.3: 验证 `queryText / statusFilter` 的默认值、默认来源、刷新恢复与返回恢复规则已经明确
  - [x] SubTask 6.4: 验证设计结果足以直接进入实现
  - [x] SubTask 6.5: 验证未把 hook 命名、缓存细节或 store API 提前写成既成事实
  - [x] SubTask 6.6: 验证跨页面列表上下文承接策略已单值化，不会在实现期出现路由层与页面瞬时状态两套分叉

# Task Dependencies

- `Task 2` depends on `Task 1`
- `Task 3` depends on `Task 2`
- `Task 4` depends on `Task 1`, `Task 2`, and `Task 3`
- `Task 5` depends on `Task 2` and `Task 3`
- `Task 6` depends on `Task 1`, `Task 2`, `Task 3`, `Task 4`, and `Task 5`
