# Tasks

- [x] Task 1: 冻结 `Module` 的最小展示字段。将列表页与详情页需要承接的字段写成单值结论，避免后续页面结构继续漂移。
  - [x] SubTask 1.1: 明确列表页最小展示字段为 `name / description / status / latest_release / product_bind_count / repository_bind_count`
  - [x] SubTask 1.2: 明确详情页最小展示字段承接核心对象字段、版本列表、产品绑定、仓库映射与相关 `Decision` 入口

- [x] Task 2: 冻结 `Module` 的最小状态表达。将状态范围与页面展示方式写成可直接实现的最小规则。
  - [x] SubTask 2.1: 明确当前阶段 `Module` 的推荐最小状态集合
  - [x] SubTask 2.2: 明确列表页与详情页都必须展示模块状态
  - [x] SubTask 2.3: 验证状态表达不引入超出正式规格正文的新生命周期体系

- [x] Task 3: 冻结 `Release` 的最小登记与展示主线。把版本登记字段、版本列表展示与 `latest_release` 承接方式写清。
  - [x] SubTask 3.1: 明确 `CreateRelease` 的最小登记字段
  - [x] SubTask 3.2: 明确 `module_id` 由 `Module Detail` 上下文承接
  - [x] SubTask 3.3: 明确 `Release` 在 `Module Detail` 中以版本主线区块展示
  - [x] SubTask 3.4: 明确必须能够识别 `latest_release`

- [x] Task 4: 冻结模块准入规则在页面侧的承接方式。确保正式 MVP 规格正文中的准入规则可以在页面中直接表达。
  - [x] SubTask 4.1: 明确 `name / description / status / Release` 对注册与版本主线的约束
  - [x] SubTask 4.2: 明确未绑定 `Product / Repository` 不阻断登记但影响复用反馈统计
  - [x] SubTask 4.3: 明确页面必须通过状态说明、提示文案或结构化分区表达当前阶段

- [x] Task 5: 完成规格校验。检查本次 `phase02-02` 规格是否满足进入后续子任务的条件。
  - [x] SubTask 5.1: 验证 `Module` 与 `Release` 的最小主线已经明确
  - [x] SubTask 5.2: 验证模块准入规则在页面侧可落地
  - [x] SubTask 5.3: 验证未引入超出正式规格正文的新对象解释

# Task Dependencies

- `Task 2` depends on `Task 1`
- `Task 3` depends on `Task 1`
- `Task 4` depends on `Task 1`, `Task 2`, and `Task 3`
- `Task 5` depends on `Task 1`, `Task 2`, `Task 3`, and `Task 4`