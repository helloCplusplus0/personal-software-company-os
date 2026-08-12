# Tasks

- [x] Task 1: 对齐 `phase09-04` 的上游边界、真实代码入口与既有 owner 现状
  - [x] SubTask 1.1: 对齐 `phase09-01 ~ 03` 与 `phase09` 三件套中关于合同、提示、模板 handoff 的共同约束
  - [x] SubTask 1.2: 复核当前真实前端入口：`/reviews/weekly`、`/products/new`、`/products/$productId`
  - [x] SubTask 1.3: 复核当前真实后端 owner：`review.QueryService`、`reusesummary.QueryService`、`productregistry.CommandService`

- [x] Task 2: 冻结 `phase09` 的 `.proto` 合同归属与最小消息边界
  - [x] SubTask 2.1: 明确 `review.proto`、`reuse_summary.proto`、`product_registry.proto` 的既有 canonical 职责不被改写
  - [x] SubTask 2.2: 明确新增 `template_reuse.proto` 作为模板候选、组合快照、派生提示与 create 预填的唯一定义源
  - [x] SubTask 2.3: 冻结 `TemplateCandidateSummary / TemplateModuleRef / TemplateCandidatePrefill / DerivedInsightHint` 的最小正式字段边界
  - [x] SubTask 2.4: 拍板 `TemplateReuseService` 作为模板读能力的唯一 canonical transport owner，并明确 `ReviewService` 只承担页面级组合读取

- [x] Task 3: 冻结模板候选的数据源、派生规则与快照决策
  - [x] SubTask 3.1: 明确模板候选继续只从 `product_modules` 已持久化事实读时派生
  - [x] SubTask 3.2: 明确 `templateCandidateId` 继续由 normalized module-set key 单向派生并由后端 owner 负责解析
  - [x] SubTask 3.3: 最终拍板当前阶段不允许引入轻量快照持久化或第二套长期稳态
  - [x] SubTask 3.4: 冻结 `templateCandidateId` 在读时派生下失效时的可恢复 unavailable 语义，覆盖 `Product Create` 与 `Product Detail`

- [x] Task 4: 冻结后端 query/command owner 与前端 read/application owner 矩阵
  - [x] SubTask 4.1: 明确 `review / reusesummary / templatereuse / productregistry` 的后端 query owner 分工
  - [x] SubTask 4.2: 明确 `review.CommandService` 与 `productregistry.CommandService` 的写路径边界不被 phase09 侵入
  - [x] SubTask 4.3: 明确 `useWeeklyReviewRead / useReviewAction / useCreateDraftProduct` 等前端 owner 的继续职责与禁止项

- [x] Task 5: 冻结当前真实 caller inventory 与模板来源回流链
  - [x] SubTask 5.1: 列出当前真实 route / page / query owner / application owner inventory，作为后续 `/spec` 强制输入
  - [x] SubTask 5.2: 明确 `fromTemplateReuse + templateCandidateId + templateSource` 在 `Product Create -> Product Detail` 成功回流链中的保留方式
  - [x] SubTask 5.3: 明确 `Product Detail` 只能通过模板来源参数复读摘要，并继续导向 canonical binding 路径

- [x] Task 6: 完成 `phase09-04` 规格自检与一致性校验
  - [x] SubTask 6.1: 校验本规格已对 `.proto` 单一合同源、读写 owner 矩阵与 caller inventory 给出单值结论
  - [x] SubTask 6.2: 校验本规格与 `phase09-02 / 03` 的模板 handoff、提示矩阵与成功回流链一致
  - [x] SubTask 6.3: 校验本规格已最终拍板“只允许读时派生，不允许轻量快照持久化”
  - [x] SubTask 6.4: 校验本规格已消除模板读合同 transport owner 歧义，并补齐 `templateCandidateId` 漂移时的机械验收语义

# Task Dependencies

- Task 2 depends on Task 1
- Task 3 depends on Task 1, Task 2
- Task 4 depends on Task 1, Task 2, Task 3
- Task 5 depends on Task 2, Task 4
- Task 6 depends on Task 1, Task 2, Task 3, Task 4, Task 5
