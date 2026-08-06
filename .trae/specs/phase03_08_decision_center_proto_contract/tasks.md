# Tasks

- [x] Task 1: 冻结 `Decision Center` Proto 合同源落点与版本语义。
  - [x] SubTask 1.1: 冻结 `Decision Center` Proto 文件目录、文件名、package 与版本号
  - [x] SubTask 1.2: 明确 `.proto` 作为当前阶段唯一合同源，手写 JSON 结构不得再并列扩张
  - [x] SubTask 1.3: 明确当前阶段只要求最小合同设计与生成入口冻结，不要求立即完成完整传输层迁移

- [x] Task 2: 冻结核心消息结构、枚举与字段编号。
  - [x] SubTask 2.1: 定义 `Decision`、`DecisionListItem`、`DecisionDetail`、`LinkedModule`、`SourceContext` 等核心消息
  - [x] SubTask 2.2: 定义 `DecisionStatus` 与 `DecisionLinkTargetType` 的最小枚举
  - [x] SubTask 2.3: 定义 `DecisionModuleCandidate` 与候选读取消息
  - [x] SubTask 2.4: 为字段分配稳定编号，并明确删除字段后的 `reserved` 兼容性规则
  - [x] SubTask 2.5: 冻结核心读模型消息字段语义来源（`Decision` / `DecisionListItem` / `DecisionDetail` / `LinkedModule` / `linked_module_summary`）
  - [x] SubTask 2.6: 冻结 `alternatives` 合同建模约束（`repeated string`，不得嵌套 `message`）
  - [x] SubTask 2.7: 冻结 `SourceContext` 字段结构（`source_module_id` + `source_module_name`，含无来源时空值语义）
  - [x] SubTask 2.8: 冻结候选读取消息字段语义（`module_id` + `module_name` + `status`，复用 `ModuleStatus`）
  - [x] SubTask 2.9: 冻结写组 request 字段语义（`CreateDecisionRequest` + `LinkDecisionToTargetRequest`）
  - [x] SubTask 2.10: 冻结字段语义与页面动作单值映射
  - [x] SubTask 2.11: 冻结枚举字段编号（`DecisionStatus` 0-4 / `DecisionLinkTargetType` 0-1）
  - [x] SubTask 2.12: 冻结核心对象字段编号（`Decision` / `DecisionListItem` / `LinkedModule` / `SourceContext` / `DecisionDetail` / `DecisionModuleCandidate`）
  - [x] SubTask 2.13: 冻结读组 request / response 字段编号
  - [x] SubTask 2.14: 冻结写组 request / response 字段编号与返回语义（`CreateDecisionResponse` 返回 `decision_id`，`LinkDecisionToTargetResponse` 空响应）
  - [x] SubTask 2.15: 冻结 `ModuleStatus` 跨包依赖策略（`import psco.module_registry.v1.ModuleStatus`，不本地重定义）

- [x] Task 3: 冻结服务接口与 request / response 语义。
  - [x] SubTask 3.1: 定义 `DecisionListRead` 对应的 `ListDecisions` request / response
  - [x] SubTask 3.2: 定义 `DecisionDetailRead` 对应的 `GetDecisionDetail` request / response
  - [x] SubTask 3.3: 定义 `DecisionWrite` 对应的 `CreateDecision` request / response
  - [x] SubTask 3.4: 定义 `DecisionLinkWrite` 对应的 `LinkDecisionToTarget` request / response
  - [x] SubTask 3.5: 定义 `DecisionModuleCandidateRead` 对应的 `ListDecisionModuleCandidates` request / response
  - [x] SubTask 3.6: 明确详情读取与候选读取不得合并为同一消息边界

- [x] Task 4: 冻结 `.proto` 与 `chi + JSON HTTP` 过渡传输层的显式映射策略。
  - [x] SubTask 4.1: 明确 `chi + JSON HTTP` 只作为过渡传输层，不能形成第二套合同源
  - [x] SubTask 4.2: 明确 RPC → HTTP 映射矩阵
  - [x] SubTask 4.3: 明确路径参数与 Proto request 字段之间的显式组装策略
  - [x] SubTask 4.4: 明确错误语义由合同承接、状态码由过渡层映射

- [x] Task 5: 冻结生成链、演进规则与 breaking check 前提。
  - [x] SubTask 5.1: 明确 `buf build / lint / generate / breaking` 为最小校验链
  - [x] SubTask 5.2: 明确 `buf breaking` 必须直接对照仓库主线 `.git` 基准且不得吞掉失败退出码
  - [x] SubTask 5.3: 明确后续新增接口必须优先修改 `.proto`，而不是先改 JSON 结构

- [x] Task 6: 完成规格校验。
  - [x] SubTask 6.1: 验证 `.proto` 合同不晚于正式规格正文进入阶段主线
  - [x] SubTask 6.2: 验证合同字段语义、字段编号与页面动作单值一致
  - [x] SubTask 6.3: 验证服务接口与 `phase03-04 / 07` 已冻结接口分组一致
  - [x] SubTask 6.4: 验证 `.proto` 与 `chi + JSON HTTP` 的显式映射策略已经明确
  - [x] SubTask 6.5: 验证合同演进规则、`reserved` 约束与 `breaking check` 前提已经明确
  - [x] SubTask 6.6: 验证核心消息字段语义已单值映射上游读模型（含 `SourceContext` 字段结构）
  - [x] SubTask 6.7: 验证字段语义与页面动作单值映射已明确
  - [x] SubTask 6.8: 验证具体字段编号方案已冻结（满足 DoD "字段编号单值一致"，未后移到 `phase03-11`）
  - [x] SubTask 6.9: 验证写组 response 返回语义已单值化（`CreateDecisionResponse` 返回 `decision_id`，`LinkDecisionToTargetResponse` 空响应）
  - [x] SubTask 6.10: 验证 `ModuleStatus` 跨包依赖策略已单值化（`import`，不本地重定义）

# Task Dependencies

- `Task 2` depends on `Task 1`
- `Task 3` depends on `Task 1` and `Task 2`
- `Task 4` depends on `Task 2` and `Task 3`
- `Task 5` depends on `Task 1`, `Task 2`, `Task 3`, and `Task 4`
- `Task 6` depends on `Task 1`, `Task 2`, `Task 3`, `Task 4`, and `Task 5`
