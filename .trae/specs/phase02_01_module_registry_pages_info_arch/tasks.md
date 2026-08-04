# Tasks

- [x] Task 1: 冻结 `Module Registry` 页面边界。将 `Module Registry / List`、`Module Create`、`Module Detail`、`Release Create` 收敛为当前阶段唯一页面主线，并写成单值结论。
  - [x] SubTask 1.1: 明确列表页承接读取、筛选入口、创建入口与进入详情入口
  - [x] SubTask 1.2: 明确创建页只承接 `CreateModule`
  - [x] SubTask 1.3: 明确详情页承接详情读取、`CreateRelease`、`BindModuleToProduct` 与 `MapModuleToRepository`
  - [x] SubTask 1.4: 明确 `Release Create` 作为版本登记的页面级入口
  - [x] SubTask 1.5: 明确不得把独立 `AI Assistant` 一级导航纳入 `phase02` 页面主线

- [x] Task 2: 冻结页面跳转关系。把列表、创建、详情、版本登记与外部轻量入口之间的最小跳转关系写清，避免后续页面主线漂移。
  - [x] SubTask 2.1: 明确 `Module Registry / List -> Module Create`
  - [x] SubTask 2.2: 明确 `Module Registry / List -> Module Detail`
  - [x] SubTask 2.3: 明确 `Module Detail -> Release Create`
  - [x] SubTask 2.4: 明确 `Module Detail -> Product Registry / Repository Binding / Decision Center` 只作为轻量跳转或关联入口

- [x] Task 3: 冻结 `PC / 移动浏览器` 信息密度策略。保持单一 `React Web` 语义，同时明确桌面与窄屏下的信息组织方式。
  - [x] SubTask 3.1: 明确桌面端优先承接较高信息密度
  - [x] SubTask 3.2: 明确移动浏览器采用信息裁剪、垂直重排与分层展示
  - [x] SubTask 3.3: 明确当前阶段不引入第二套移动端 UI、独立 `React Native` 客户端或完整 `PWA`

- [x] Task 4: 完成规格校验。检查本次 `phase02-01` 规格是否满足进入后续子任务的条件。
  - [x] SubTask 4.1: 验证页面职责已经单值化
  - [x] SubTask 4.2: 验证无第二套移动端 UI 方案
  - [x] SubTask 4.3: 验证页面边界与 `phase01-06` 正式规格正文一致

# Task Dependencies

- `Task 2` depends on `Task 1`
- `Task 3` depends on `Task 1`
- `Task 4` depends on `Task 1`, `Task 2`, and `Task 3`