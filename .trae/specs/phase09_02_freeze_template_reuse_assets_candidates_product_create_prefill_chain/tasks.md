# Tasks

- [x] Task 1: 对齐 `phase09-02` 的上游边界与模板复用最小语义
  - [x] SubTask 1.1: 对齐 `phase09-01`、`phase09` 三件套中关于 `Template Reuse` 的共同边界
  - [x] SubTask 1.2: 冻结模板级复用资产为 `Module` 组合快照 + `Product Create` 预填辅助
  - [x] SubTask 1.3: 明确本任务不越权冻结 `phase09-03` 的提示矩阵与 `phase09-04` 的合同/owner 细节

- [x] Task 2: 冻结模板候选来源、去重与空态规则
  - [x] SubTask 2.1: 明确模板候选只从 `product_modules` 已持久化事实派生
  - [x] SubTask 2.2: 明确原始候选输入、去重键、排序规则、空态与低质量候选规则
  - [x] SubTask 2.3: 明确 `Review` 只承担消费作用域与返回链上下文，不直接生成模板候选

- [x] Task 3: 冻结 `Weekly Review` 中模板候选的消费语义
  - [x] SubTask 3.1: 明确候选列表采用单选模型
  - [x] SubTask 3.2: 明确默认 active candidate、切换规则与无候选空态
  - [x] SubTask 3.3: 明确依赖模板上下文的后续 handoff 只能基于当前 active candidate

- [x] Task 4: 冻结 `Product Create` 模板 handoff 与预填读取入口
  - [x] SubTask 4.1: 明确 `/products/new` 的模板来源参数矩阵
  - [x] SubTask 4.2: 明确 `fromTemplateReuse` 与 `fromList / fromModuleDetail / fromDashboard` 的优先级、共存与互斥规则
  - [x] SubTask 4.3: 明确 `templateCandidateId` 是唯一正式预填读取入口，前端按 opaque string 消费

- [x] Task 5: 冻结模板预填成功会话与创建成功后的承接链
  - [x] SubTask 5.1: 明确"预填成立"的最小字段覆盖范围
  - [x] SubTask 5.2: 明确 `Product Create` 不自动写入 `product_modules`
  - [x] SubTask 5.3: 明确创建成功后继续携带 `fromTemplateReuse + templateCandidateId + templateSource`，并在 `Product Detail` 中复读模板来源与 canonical `Product <-> Module Binding` CTA

- [x] Task 6: 完成 `phase09-02` 规格自检与一致性校验
  - [x] SubTask 6.1: 校验 `spec.md` 已覆盖候选来源、active candidate、handoff、预填字段与成功回流链
  - [x] SubTask 6.2: 校验任务拆分与 `phase09` 三件套、`phase09-01` 上游边界一致
  - [x] SubTask 6.3: 校验本任务未越权冻结提示类型、合同字段名或读写 owner 细节

# Task Dependencies

- Task 2 depends on Task 1
- Task 3 depends on Task 2
- Task 4 depends on Task 1, Task 2
- Task 5 depends on Task 3, Task 4
- Task 6 depends on Task 1, Task 2, Task 3, Task 4, Task 5
