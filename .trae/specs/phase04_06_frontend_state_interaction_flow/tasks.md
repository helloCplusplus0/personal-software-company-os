# Tasks

- [x] Task 1: 冻结两个列表页的查询条件承接策略、查询状态、筛选状态与刷新恢复规则
  - [x] SubTask 1.1: 明确列表查询条件冻结到路由搜索参数层（`queryText / statusFilter`），默认值 `queryText=空 / statusFilter=all`
    - 证据：`spec.md` §ADDED Requirements「列表查询条件承接策略冻结」Scenario: 判断列表查询条件的路由搜索参数承接 L35-44
  - [x] SubTask 1.2: 明确列表查询条件的唯一事实源为路由搜索参数，不引入 `sessionStorage` 等持久化层作为缺参回退源
    - 证据：`spec.md` §ADDED Requirements「列表查询条件承接策略冻结」Scenario: 判断列表查询条件的路由搜索参数承接 L44 + 判断列表刷新行为 L56-62
  - [x] SubTask 1.3: 明确 `fromList` 来源标记建模"来源列表上下文存在/不存在"，存在时返回导航携带原参数，不存在时落默认参数
    - 证据：`spec.md` §ADDED Requirements「列表查询条件承接策略冻结」Scenario: 判断 fromList 来源标记与返回列表的参数保留 L46-54
  - [x] SubTask 1.4: 明确列表刷新行为（路由搜索参数缺失时使用默认参数，不得回退持久化层，无参 URL 稳定表现默认筛选）
    - 证据：`spec.md` §ADDED Requirements「列表查询条件承接策略冻结」Scenario: 判断列表刷新行为 L56-62
  - [x] SubTask 1.5: 明确两个列表页的读取状态（`pending/success/error`）、派生视图状态（`initial-loading/ready/empty/error`）、空状态主动作与错误呈现位置
    - 证据：`spec.md` §ADDED Requirements「列表查询状态与空状态模型冻结」Scenario: 判断列表读取状态 L68-74 + 判断 Product List 空状态 L76-81 + 判断 Repository Binding / List 空状态 L83-88 + 判断列表错误呈现位置 L90-94
  - [x] SubTask 1.6: 单值化列表筛选维度（`Product List` 与 `Repository Binding / List` 均只冻结 `queryText / statusFilter`，不引入 `providerFilter`）
    - 证据：`spec.md` §ADDED Requirements「列表筛选维度冻结」Scenario: 判断 Product List 筛选维度 L100-105 + 判断 Repository Binding / List 筛选维度 L107-113

- [x] Task 2: 冻结 `Product Create` 的来源上下文、草稿状态、提交状态与成功回流（含来源标记继承）
  - [x] SubTask 2.1: 明确 `ProductCreatePage` 来源上下文由路由搜索参数派生（`fromList` / `fromModuleDetail` / `direct-entry`），不发明第二套来源标记命名
    - 证据：`spec.md` §ADDED Requirements「Product Create 交互状态流转冻结」Scenario: 判断 Product Create 来源上下文 L119-127
  - [x] SubTask 2.2: 明确 `ProductCreatePage` 最小草稿状态（`idle/dirty`，`status` 默认预填 `active`）与提交状态（`submitting/submit-success/submit-error`）
    - 证据：`spec.md` §ADDED Requirements「Product Create 交互状态流转冻结」Scenario: 判断 Product Create 草稿与提交状态 L129-135
  - [x] SubTask 2.3: 明确创建失败后停留当前页、保留草稿与来源上下文、错误显示在表单上下文
    - 证据：`spec.md` §ADDED Requirements「Product Create 交互状态流转冻结」Scenario: 判断 Product Create 提交失败处理 L137-144
  - [x] SubTask 2.4: 明确创建成功回流到 `ProductDetailPage` 时必须继续携带来源标记与上下文参数（`fromList` 带原参数 / `fromModuleDetail` 带原参数 / `direct-entry` 不带），不得丢失导致退化为 `direct-entry`
    - 证据：`spec.md` §ADDED Requirements「Product Create 交互状态流转冻结」Scenario: 判断 Product Create 提交成功回流 L146-157
  - [x] SubTask 2.5: 明确取消返回按真实来源决定（`fromList` 回列表携带原参数 / `fromModuleDetail` 回 ModuleDetail / `direct-entry` 回列表默认参数），不统一伪造成回列表
    - 证据：`spec.md` §ADDED Requirements「Product Create 交互状态流转冻结」Scenario: 判断 Product Create 取消返回 L159-167

- [x] Task 3: 冻结 `Repository Create` 的来源上下文、草稿状态、提交状态与成功回流（含来源标记继承）
  - [x] SubTask 3.1: 明确 `RepositoryCreatePage` 来源上下文由路由搜索参数派生（`fromList` / `fromProductDetail` / `fromModuleDetail` / `direct-entry`）
    - 证据：`spec.md` §ADDED Requirements「Repository Create 交互状态流转冻结」Scenario: 判断 Repository Create 来源上下文 L173-182
  - [x] SubTask 3.2: 明确 `RepositoryCreatePage` 最小草稿状态与提交状态（`name/url/provider/status`，`status` 默认预填 `active`）
    - 证据：`spec.md` §ADDED Requirements「Repository Create 交互状态流转冻结」Scenario: 判断 Repository Create 草稿与提交状态 L184-190
  - [x] SubTask 3.3: 明确创建失败后停留当前页、保留草稿与来源上下文、错误显示在表单上下文
    - 证据：`spec.md` §ADDED Requirements「Repository Create 交互状态流转冻结」Scenario: 判断 Repository Create 提交失败处理 L192-199
  - [x] SubTask 3.4: 明确创建成功回流到 `RepositoryBindingDetailPage` 时必须继续携带来源标记与上下文参数（`fromList` / `fromProductDetail` / `fromModuleDetail` 带原参数 / `direct-entry` 不带），不得丢失导致退化为 `direct-entry`
    - 证据：`spec.md` §ADDED Requirements「Repository Create 交互状态流转冻结」Scenario: 判断 Repository Create 提交成功回流 L201-213
  - [x] SubTask 3.5: 明确取消返回按真实来源决定（`fromList` 回列表携带原参数 / `fromProductDetail` 回 ProductDetail / `fromModuleDetail` 回 ModuleDetail / `direct-entry` 回列表默认参数），不统一伪造成回列表
    - 证据：`spec.md` §ADDED Requirements「Repository Create 交互状态流转冻结」Scenario: 判断 Repository Create 取消返回 L215-224

- [x] Task 4: 冻结 `Product Detail` 的详情读取状态、来源上下文（含从 Create 继承）与 `BindModuleToProduct` 绑定交互流
  - [x] SubTask 4.1: 明确 `ProductDetailPage` 读取状态（`pending/success/error`，资源不存在派生 `not-found`）与来源上下文（含从 `ProductCreatePage` 成功创建后继承来源上下文）
    - 证据：`spec.md` §ADDED Requirements「Product Detail 交互状态流转冻结」Scenario: 判断 Product Detail 读取状态 L230-235 + 判断 Product Detail 来源上下文 L237-245
  - [x] SubTask 4.2: 明确 `ProductModuleBindingPanel` 候选读取状态（`closed/pending/ready/empty/error`）与写入状态（`idle/submitting/submit-success/submit-error`），候选读取独立于详情读取
    - 证据：`spec.md` §ADDED Requirements「Product Detail 交互状态流转冻结」Scenario: 判断 ProductModuleBindingPanel 候选读取状态 L247-252 + 判断 ProductModuleBindingPanel 写入状态 L254-258
  - [x] SubTask 4.3: 明确绑定成功后停留当前 `ProductDetailPage`、重新读取已绑定模块列表完成 reread、面板回到 `closed`，承接 `phase04-03`
    - 证据：`spec.md` §ADDED Requirements「Product Detail 交互状态流转冻结」Scenario: 判断 BindModuleToProduct 提交成功 reread L260-266
  - [x] SubTask 4.4: 明确候选为空状态（空列表语义，不误报为接口错误）与绑定失败后状态保持（保留选择、错误停留面板上下文、重复绑定不静默成功）
    - 证据：`spec.md` §ADDED Requirements「Product Detail 交互状态流转冻结」Scenario: 判断 BindModuleToProduct 候选为空状态 L268-272 + 判断 BindModuleToProduct 提交失败处理 L274-280

- [x] Task 5: 冻结 `Repository Binding Detail / Workspace` 的详情读取状态、来源上下文（含从 Create 继承）与两类绑定交互流
  - [x] SubTask 5.1: 明确 `RepositoryBindingDetailPage` 读取状态（`pending/success/error`，资源不存在派生 `not-found`）与来源上下文（含从 `RepositoryCreatePage` 成功创建后继承来源上下文）
    - 证据：`spec.md` §ADDED Requirements「Repository Binding Detail / Workspace 绑定工作台交互状态流转冻结」Scenario: 判断 Repository Binding Detail 读取状态 L286-291 + 判断 Repository Binding Detail 来源上下文 L293-302
  - [x] SubTask 5.2: 明确 `RepositoryProductBindingPanel` 与 `RepositoryModuleMappingPanel` 的候选读取状态与写入状态
    - 证据：`spec.md` §ADDED Requirements「Repository Binding Detail / Workspace 绑定工作台交互状态流转冻结」Scenario: 判断两类绑定面板状态与互斥展开 L304-310
  - [x] SubTask 5.3: 明确单活动面板规则（同一时刻只允许一个绑定面板打开，互斥展开）
    - 证据：`spec.md` §ADDED Requirements「Repository Binding Detail / Workspace 绑定工作台交互状态流转冻结」Scenario: 判断两类绑定面板状态与互斥展开 L304-310
  - [x] SubTask 5.4: 明确两类绑定成功后停留当前页、重新读取对应已绑定列表完成 reread、面板回到 `closed`，承接 `phase04-03`
    - 证据：`spec.md` §ADDED Requirements「Repository Binding Detail / Workspace 绑定工作台交互状态流转冻结」Scenario: 判断 BindRepositoryToProduct 提交成功 reread L319-325 + 判断 MapModuleToRepository 提交成功 reread L327-333
  - [x] SubTask 5.5: 明确候选为空状态与绑定失败后状态保持（保留选择、错误停留面板上下文、重复绑定不静默成功）
    - 证据：`spec.md` §ADDED Requirements「Repository Binding Detail / Workspace 绑定工作台交互状态流转冻结」Scenario: 判断两类绑定候选为空状态 L312-317 + 判断两类绑定提交失败处理 L335-341

- [x] Task 6: 冻结多入口来源上下文、返回路径、创建成功链路来源继承与刷新恢复规则
  - [x] SubTask 6.1: 明确从 `Product List` / `Repository Binding / List` 进入 Detail 后返回列表时携带 `fromList` 标记并按路由搜索参数保留原上下文
    - 证据：`spec.md` §ADDED Requirements「多入口返回路径与上下文恢复规则冻结」Scenario: 判断从 Product List 进入 Detail 后返回 List 的上下文恢复 L367-372 + 判断从 Repository Binding / List 进入 Detail 后返回 List 的上下文恢复 L374-379
  - [x] SubTask 6.2: 明确从 `Module Detail` 兼容入口进入后的来源标记（`fromModuleDetail`）、绑定成功 reread 停留 canonical owner、用户主动返回默认回到 `ModuleDetailPage`
    - 证据：`spec.md` §ADDED Requirements「多入口返回路径与上下文恢复规则冻结」Scenario: 判断从 Module Detail 兼容入口进入后的来源标记与 reread L349-356
  - [x] SubTask 6.3: 明确从 `Product Detail` 上下文入口进入后的来源标记（`fromProductDetail`）、绑定成功 reread 停留 `Repository Binding Detail / Workspace`、用户主动返回默认回到 `ProductDetailPage`
    - 证据：`spec.md` §ADDED Requirements「多入口返回路径与上下文恢复规则冻结」Scenario: 判断从 Product Detail 上下文入口进入后的来源标记与 reread L358-365
  - [x] SubTask 6.4: 明确无来源列表上下文时返回列表落到默认参数，不伪造上一份筛选条件、不回退持久化层
    - 证据：`spec.md` §ADDED Requirements「多入口返回路径与上下文恢复规则冻结」Scenario: 判断无来源列表上下文时返回列表的默认行为 L392-397
  - [x] SubTask 6.5: 明确来源上下文刷新恢复（`fromModuleDetail` / `fromProductDetail` 刷新后不丢失）
    - 证据：`spec.md` §ADDED Requirements「多入口返回路径与上下文恢复规则冻结」Scenario: 判断来源上下文刷新恢复 L399-403
  - [x] SubTask 6.6: 明确 reread 与返回路径的区分（绑定成功 reread 停留 canonical owner，与用户主动返回是两个独立动作）
    - 证据：`spec.md` §ADDED Requirements「多入口返回路径与上下文恢复规则冻结」关键区分说明 L347 + Scenario: 判断从 Module Detail 兼容入口进入后的来源标记与 reread L349-356
  - [x] SubTask 6.7: 明确从创建成功后的 Detail 页返回来源页时基于继承的来源标记按真实来源返回，承接 `phase03-06` 创建成功后来源上下文继承模式
    - 证据：`spec.md` §ADDED Requirements「多入口返回路径与上下文恢复规则冻结」Scenario: 判断从创建成功后的 Detail 页返回来源页 L381-390

- [x] Task 7: 冻结页面级 UI 状态局部归属原则与运行时实现细节不冻结
  - [x] SubTask 7.1: 明确列表页、创建页与详情页的局部 UI 状态优先归属当前页面或详情页上下文，不默认升级为跨路由全局状态
    - 证据：`spec.md` §ADDED Requirements「页面级 UI 状态局部归属原则冻结」Scenario: 判断列表与创建页局部状态归属 L409-413 + 判断详情页局部状态归属 L415-419
  - [x] SubTask 7.2: 明确当前阶段不冻结 hook 命名、Query key、store API、缓存时间、optimistic update 方案，且不引入 `sessionStorage` 作为事实源
    - 证据：`spec.md` §ADDED Requirements「运行时实现细节不冻结」Scenario: 判断当前阶段允许冻结的内容 L425-428 + 判断当前阶段不得冻结的内容 L430-434

- [x] Task 8: 完成规格校验，确认与上游冻结结论一致且足以直接进入实现
  - [x] SubTask 8.1: 验证列表筛选维度、默认值与 `providerFilter` 结论和 `phase04-02 / 04` 一致
    - 证据：`spec.md` §ADDED Requirements「列表筛选维度冻结」L96-113 + §REMOVED Requirements「providerFilter 作为当前阶段筛选维度」L462-466
  - [x] SubTask 8.2: 验证来源上下文参数与 `phase04-03 / 05` 的路由参数冻结保持一致（`fromModuleDetail` / `fromProductDetail` / `fromList` 无第二套命名）
    - 证据：`spec.md` §MODIFIED Requirements「phase04-05 上下文承接路由参数的状态语义解释」L436-448 + §ADDED Requirements 各 Create/Detail 来源上下文 Scenario
  - [x] SubTask 8.3: 验证 reread 行为与 `phase04-03` canonical owner 与 reread 承接页面冻结结论一致
    - 证据：`spec.md` §MODIFIED Requirements「phase04-03 canonical owner 与 reread 的前端状态承接解释」L450-458
  - [x] SubTask 8.4: 验证列表查询条件唯一事实源为路由搜索参数（不引入 `sessionStorage`）与 `fromList` 来源标记承接 `phase02-09` §7.4/§8.4 与 `phase03-06` 既有模式
    - 证据：`spec.md` §ADDED Requirements「列表查询条件承接策略冻结」Scenario: 判断列表查询条件的路由搜索参数承接 L35-44 + 判断 fromList 来源标记与返回列表的参数保留 L46-54 + 判断列表刷新行为 L56-62
  - [x] SubTask 8.5: 验证未越界冻结 `phase04-07 / 08` 的后端模块边界与 `.proto` 合同设计
    - 证据：`spec.md` Why 部分阶段分工约束 + §ADDED Requirements「运行时实现细节不冻结」L421-434
  - [x] SubTask 8.6: 验证 Create 成功回流来源标记继承与 `phase03-06` 创建成功后来源上下文继承模式一致
    - 证据：`spec.md` §ADDED Requirements「Product Create 交互状态流转冻结」Scenario: 判断 Product Create 提交成功回流 L146-157 + 「Repository Create 交互状态流转冻结」Scenario: 判断 Repository Create 提交成功回流 L201-213 + 「多入口返回路径与上下文恢复规则冻结」Scenario: 判断从创建成功后的 Detail 页返回来源页 L381-390
  - [x] SubTask 8.7: 验证设计结果足以直接进入实现
    - 证据：`spec.md` 全文状态模型、交互流、查询条件承接策略、返回路径与来源标记继承均单值化，运行时实现细节显式不冻结

# Task Dependencies

- `Task 2` depends on `Task 1`
- `Task 3` depends on `Task 1`
- `Task 4` depends on `Task 1`
- `Task 5` depends on `Task 1`
- `Task 6` depends on `Task 1`, `Task 2`, `Task 3`, `Task 4`, `Task 5`
- `Task 7` depends on `Task 2`, `Task 3`, `Task 4`, `Task 5`
- `Task 8` depends on `Task 1` through `Task 7`
