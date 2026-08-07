# Tasks

- [x] Task 1: 冻结 `Product Registry` 页面边界。将 `Product Registry / List`、`Product Create`、`Product Detail` 收敛为当前阶段唯一 `Product` 页面主线，并写成单值结论。
  - [x] SubTask 1.1: 明确列表页承接产品读取、筛选入口、创建入口与进入详情入口
    - 证据：`spec.md` §ADDED Requirements「Product Registry 列表页职责冻结」Scenario: Product 列表页职责判定 — 明确承接 `Product` 列表读取、提供进入 `Product Create` 与 `Product Detail` 的明确入口、筛选只作为列表页入口能力
  - [x] SubTask 1.2: 明确创建页只承接 `CreateProduct`
    - 证据：`spec.md` §ADDED Requirements「Product Create 页面职责冻结」Scenario: Product 创建页职责判定 — 明确为 `CreateProduct` 的唯一页面级承接入口，不得分散到列表页内联复杂编辑流
  - [x] SubTask 1.3: 明确详情页承接详情读取、已绑定模块/仓库读取、`BindModuleToProduct` 与进入仓库绑定工作台的上下文入口
    - 证据：`spec.md` §ADDED Requirements「Product Detail 页面职责冻结」Scenario: Product 详情页职责判定 — 展示核心信息、已绑定 `Module` 列表、已绑定 `Repository` 列表，承接 `BindModuleToProduct` 最小写入触点与进入 `Repository Binding Detail / Workspace` 的上下文入口
  - [x] SubTask 1.4: 明确 `Product Detail` 不承接第二套仓库绑定写入流程
    - 证据：`spec.md` §ADDED Requirements「Product Detail 页面职责冻结」Scenario L70 — 「不得在当前阶段并行承接第二套仓库绑定写入流程」

- [x] Task 2: 冻结 `Repository Binding` 页面边界。将 `Repository Binding / List`、`Repository Create`、`Repository Binding Detail / Workspace` 收敛为当前阶段唯一 `Repository` 页面主线，并写成单值结论。
  - [x] SubTask 2.1: 明确列表页承接仓库读取、筛选入口、创建入口与进入工作台入口
    - 证据：`spec.md` §ADDED Requirements「Repository Binding 列表页职责冻结」Scenario: Repository 列表页职责判定 — 承接 `Repository` 列表读取、提供进入 `Repository Create` 与 `Repository Binding Detail / Workspace` 的明确入口
  - [x] SubTask 2.2: 明确创建页只承接 `CreateRepository`
    - 证据：`spec.md` §ADDED Requirements「Repository Create 页面职责冻结」Scenario: Repository 创建页职责判定 — 明确为 `CreateRepository` 的唯一页面级承接入口
  - [x] SubTask 2.3: 明确工作台页承接详情读取、候选读取、`BindRepositoryToProduct` 与 `MapModuleToRepository`
    - 证据：`spec.md` §ADDED Requirements「Repository Binding Detail / Workspace 页面职责冻结」Scenario: Repository 工作台职责判定 — 展示核心信息、已绑定 `Product` 列表、已映射 `Module` 列表，承接 `BindRepositoryToProduct` 与 `MapModuleToRepository` 的最小写入触点
  - [x] SubTask 2.4: 明确 `Repository Binding Detail / Workspace` 不承接 `BindModuleToProduct`
    - 证据：`spec.md` §ADDED Requirements「Repository Binding Detail / Workspace 页面职责冻结」Scenario L106 — 「不得把 `BindModuleToProduct` 迁入该页面形成并列主写入流程」

- [x] Task 3: 冻结五个核心动作与六个页面之间的单值 owner 矩阵。把页面 owner 写清，避免后续 `/spec` 和实现期出现双 owner。
  - [x] SubTask 3.1: 明确 `CreateProduct` → `Product Create`
    - 证据：`spec.md` §ADDED Requirements「五个核心动作 owner 冻结」Scenario: 动作 owner 判定 L116
  - [x] SubTask 3.2: 明确 `CreateRepository` → `Repository Create`
    - 证据：`spec.md` §ADDED Requirements「五个核心动作 owner 冻结」Scenario: 动作 owner 判定 L117
  - [x] SubTask 3.3: 明确 `BindModuleToProduct` → `Product Detail`
    - 证据：`spec.md` §ADDED Requirements「五个核心动作 owner 冻结」Scenario: 动作 owner 判定 L118
  - [x] SubTask 3.4: 明确 `BindRepositoryToProduct` → `Repository Binding Detail / Workspace`
    - 证据：`spec.md` §ADDED Requirements「五个核心动作 owner 冻结」Scenario: 动作 owner 判定 L119
  - [x] SubTask 3.5: 明确 `MapModuleToRepository` → `Repository Binding Detail / Workspace`
    - 证据：`spec.md` §ADDED Requirements「五个核心动作 owner 冻结」Scenario: 动作 owner 判定 L120

- [x] Task 4: 冻结 `Product Registry`、`Repository Binding`、`Module Detail` 之间的最小跳转关系与上下文入口。把正式主入口、兼容入口与返回路径前提写清，避免页面主线漂移。
  - [x] SubTask 4.1: 明确 `Product Registry / List -> Product Create`
    - 证据：`spec.md` §ADDED Requirements「页面跳转关系冻结」Scenario: Product 主线最小跳转关系判定 L129
  - [x] SubTask 4.2: 明确 `Product Registry / List -> Product Detail`
    - 证据：`spec.md` §ADDED Requirements「页面跳转关系冻结」Scenario: Product 主线最小跳转关系判定 L130
  - [x] SubTask 4.3: 明确 `Product Detail -> Repository Binding Detail / Workspace` 的带上下文跳转
    - 证据：`spec.md` §ADDED Requirements「页面跳转关系冻结」Scenario: Product 主线最小跳转关系判定 L132
  - [x] SubTask 4.4: 明确 `Repository Binding / List -> Repository Create`
    - 证据：`spec.md` §ADDED Requirements「页面跳转关系冻结」Scenario: Repository 主线最小跳转关系判定 L137
  - [x] SubTask 4.5: 明确 `Repository Binding / List -> Repository Binding Detail / Workspace`
    - 证据：`spec.md` §ADDED Requirements「页面跳转关系冻结」Scenario: Repository 主线最小跳转关系判定 L138
  - [x] SubTask 4.6: 明确 `Module Detail` 只保留绑定摘要与正式主入口跳转，不再保留本页直写
    - 证据：`spec.md` §ADDED Requirements「页面跳转关系冻结」Scenario: Module Detail 兼容入口判定 L143-146 + §MODIFIED Requirements「Module Detail 的绑定承接方式」— 只允许轻量跳转或带上下文入口进入正式主入口，不得在本页直接提交绑定写入

- [x] Task 5: 冻结六个页面的最小页面级信息区块组成。为后续前端页面/路由/组件设计提供直接上游。
  - [x] SubTask 5.1: 明确 `Product Registry / List` 至少由列表工具栏区、列表内容区与空状态区组成
    - 证据：`spec.md` §ADDED Requirements「Product 页面级信息区块冻结」Scenario: Product 列表页信息区块判定 L155-158
  - [x] SubTask 5.2: 明确 `Product Create` 至少由结构化表单区与提交取消操作区组成
    - 证据：`spec.md` §ADDED Requirements「Product 页面级信息区块冻结」Scenario: Product 创建页信息区块判定 L163-165
  - [x] SubTask 5.3: 明确 `Product Detail` 至少由核心信息区、已绑定模块区、已绑定仓库区与绑定入口区组成
    - 证据：`spec.md` §ADDED Requirements「Product 页面级信息区块冻结」Scenario: Product 详情页信息区块判定 L170-174
  - [x] SubTask 5.4: 明确 `Repository Binding / List` 至少由列表工具栏区、列表内容区与空状态区组成
    - 证据：`spec.md` §ADDED Requirements「Repository 页面级信息区块冻结」Scenario: Repository 列表页信息区块判定 L183-187 — 工具栏冻结为 `queryText` + `statusFilter` + 进入 `Repository Create` 入口三项；`providerFilter` 不在 `phase04-01` 冻结范围，留给 `phase04-02 / phase04-06` 单值化
  - [x] SubTask 5.5: 明确 `Repository Create` 至少由结构化表单区与提交取消操作区组成
    - 证据：`spec.md` §ADDED Requirements「Repository 页面级信息区块冻结」Scenario: Repository 创建页信息区块判定 L192-194
  - [x] SubTask 5.6: 明确 `Repository Binding Detail / Workspace` 至少由核心信息区、已绑定产品区、已映射模块区与绑定工作台区组成
    - 证据：`spec.md` §ADDED Requirements「Repository 页面级信息区块冻结」Scenario: Repository 工作台信息区块判定 L199-203

- [x] Task 6: 冻结 `PC / 移动浏览器` 信息密度策略。保持单一 `React Web` 语义，同时明确桌面与窄屏下的信息组织方式。
  - [x] SubTask 6.1: 明确桌面端优先承接较高信息密度
    - 证据：`spec.md` §ADDED Requirements「PC 与移动浏览器信息密度策略冻结」Scenario: 桌面端信息密度 L211-213
  - [x] SubTask 6.2: 明确移动浏览器采用信息裁剪、区块折叠、垂直重排与分层展示
    - 证据：`spec.md` §ADDED Requirements「PC 与移动浏览器信息密度策略冻结」Scenario: 移动浏览器信息密度 L217-219（WHEN + THEN + 信息裁剪 AND）
  - [x] SubTask 6.3: 明确当前阶段不引入第二套移动端 UI、独立 `React Native` 客户端或完整 `PWA`
    - 证据：`spec.md` §ADDED Requirements「PC 与移动浏览器信息密度策略冻结」Scenario L221 + §REMOVED Requirements「第二套移动端 UI 方案」Reason/Migration

- [x] Task 7: 完成规格校验。检查本次 `phase04-01` 规格是否满足进入后续子任务的条件。
  - [x] SubTask 7.1: 验证 `Product Registry` 与 `Repository Binding` 页面职责已经单值化
    - 证据：`spec.md` §ADDED Requirements 中「Product Registry 列表页职责冻结」「Product Create 页面职责冻结」「Product Detail 页面职责冻结」「Repository Binding 列表页职责冻结」「Repository Create 页面职责冻结」「Repository Binding Detail / Workspace 页面职责冻结」六个 Requirement 各自单值化，无并列 owner
  - [x] SubTask 7.2: 验证五个核心动作与六个页面的 owner 已经单值化
    - 证据：`spec.md` §ADDED Requirements「五个核心动作 owner 冻结」Scenario: 动作 owner 判定 — 五个动作各对应唯一页面 owner，无并行 owner
  - [x] SubTask 7.3: 验证 `Module Detail` 未被保留为第二绑定工作台
    - 证据：`spec.md` §MODIFIED Requirements「Module Detail 的绑定承接方式」+ §ADDED Requirements「页面跳转关系冻结」Scenario: Module Detail 兼容入口判定 — `Module Detail` 回落为摘要与上下文入口，不作为并列写入 owner
  - [x] SubTask 7.4: 验证无第二套移动端 UI 方案
    - 证据：`spec.md` §REMOVED Requirements「第二套移动端 UI 方案」+ §ADDED Requirements「PC 与移动浏览器信息密度策略冻结」— 当前阶段冻结为单一 `React Web`，不引入第二套移动端 UI / 独立 `React Native` / 完整 `PWA`
  - [x] SubTask 7.5: 验证页面边界与 `phase01-06` 正式规格正文、`phase04` 三件套一致
    - 证据（与 phase01-06 一致）：`mvp_spec_v0.1.md` §2.1 核心实体含 `Product` / `Repository`、§3 动作矩阵含五个核心动作、§4.1 页面范围含 `Product Registry` / `Repository Binding`、§4.2 明确不进入 `AI Assistant` 一级导航 — 与 `spec.md` §ADDED Requirements「Phase04 页面边界冻结」L32-34 一致
    - 证据（与 phase04 三件套一致）：`dev_plan` §3 phase04-01 范围五项 → `spec.md` 五组 Requirement 覆盖；`shared_baseline` §4 页面矩阵六类 + §4.1 交互归属矩阵 → `spec.md` 六个页面职责 Requirement + 动作 owner Requirement 一致；`architecture_plan` §4.5 交互归属原则 → `spec.md` 动作 owner 与跳转关系一致

# Task Dependencies

- `Task 2` depends on `Task 1`
- `Task 3` depends on `Task 1`, `Task 2`
- `Task 4` depends on `Task 1`, `Task 2`, `Task 3`
- `Task 5` depends on `Task 1`, `Task 2`, `Task 3`
- `Task 6` depends on `Task 1`, `Task 2`
- `Task 7` depends on `Task 1`, `Task 2`, `Task 3`, `Task 4`, `Task 5`, and `Task 6`
