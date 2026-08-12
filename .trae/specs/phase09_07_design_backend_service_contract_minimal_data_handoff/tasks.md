# Tasks

- [x] Task 1: 对齐 `phase09-07` 的上游冻结结论、真实 proto/service 主线与后端装配基线
  - [x] SubTask 1.1: 对齐 `phase09-04` 的合同、owner、候选来源与来源复读边界
  - [x] SubTask 1.2: 对齐 `phase09-05` 的页面流、返回链与 unavailable 成功态要求
  - [x] SubTask 1.3: 对齐 `phase09-06` 的前端 owner / 状态流承接口径
  - [x] SubTask 1.4: 复核当前真实后端入口：`review.proto`、`reuse_summary.proto`、`product_registry.proto`、`review.QueryService`、`reusesummary.QueryService`、`productregistry.QueryService/CommandService`、`platform/router.go`

- [x] Task 2: 冻结 `template_reuse.proto` 与 `TemplateReuseService` 的正式合同边界
  - [x] SubTask 2.1: 明确 `template_reuse.proto` 是模板候选、预填、提示与来源复读的唯一合同源
  - [x] SubTask 2.2: 明确 `TemplateReuseService` 的最小 RPC 矩阵与四类读取职责
  - [x] SubTask 2.3: 明确模板候选、active candidate、预填、提示与来源复读的最小字段/枚举边界
  - [x] SubTask 2.3.1: 明确四个读取 RPC 的最小 request 字段，尤其冻结 `GetDerivedInsightHintsRequest` 的 `consumer_surface + review_scope_key` 作用域输入
  - [x] SubTask 2.4: 明确三类核心读取接口与成功回流来源复读接口的职责不重叠

- [x] Task 3: 冻结后端 query / command owner 与既有服务协作边界
  - [x] SubTask 3.1: 明确 `templatereuse.QueryService` 是模板读能力的唯一后端 owner
  - [x] SubTask 3.2: 明确 `Review / ReuseSummary / Product / Decision` 的协作边界与禁止事项
  - [x] SubTask 3.3: 明确当前阶段不得新增 `TemplateReuseCommandService` 或第二套 create/template write RPC
  - [x] SubTask 3.4: 明确 Connect handler 物理落点、`/api` 挂载方式与错误映射口径

- [x] Task 4: 冻结最小数据承接设计与轻量快照评估结论
  - [x] SubTask 4.1: 明确模板候选继续只从 `product_modules` 读时派生
  - [x] SubTask 4.2: 写出“当前无需轻量快照记录”的证据链
  - [x] SubTask 4.3: 明确若未来提议引入轻量快照，必须满足的受控前提
  - [x] SubTask 4.4: 明确 `templateCandidateId` 漂移时的 unavailable 成功态由哪些接口承接
  - [x] SubTask 4.4.1: 明确 `GetDerivedInsightHints` 在候选漂移时必须返回 `UNAVAILABLE` 成功态，而不是与“无提示空态”混用

- [x] Task 5: 冻结工具链、API smoke 与浏览器验收口径
  - [x] SubTask 5.1: 明确 `buf / go build / frontend type-check` 的正式命令口径
  - [x] SubTask 5.2: 明确 `TemplateReuseService` 四个关键 RPC 的 API smoke 清单
  - [x] SubTask 5.3: 明确 `Weekly Review -> Product Create -> Product Detail` 的最小浏览器验收闭环
  - [x] SubTask 5.4: 明确 `UNAVAILABLE` 成功态必须进入 smoke 与浏览器验收

- [x] Task 6: 完成 `phase09-07` 规格自检与一致性校验
  - [x] SubTask 6.1: 校验合同、服务与数据承接位已单值化
  - [x] SubTask 6.2: 校验没有复制既有业务事实源
  - [x] SubTask 6.3: 校验“无需新增轻量快照记录”已被明确冻结
  - [x] SubTask 6.4: 校验 `buf / go build / frontend type-check / browser acceptance` 口径已冻结
  - [x] SubTask 6.5: 校验模板候选、active candidate 与来源复读链已有单值合同语义
  - [x] SubTask 6.6: 校验三类核心读取接口的职责边界足以直接驱动 API smoke 与前端 owner 实现
  - [x] SubTask 6.7: 校验 request 侧合同已经足以单值驱动 `Weekly Review` 与 `Product Create / Product Detail` 的 owner 接线
  - [x] SubTask 6.8: 校验 `GetDerivedInsightHints` 已可区分“无提示成功空态”与“候选漂移 unavailable 成功态”

# Task Dependencies

- Task 2 depends on Task 1
- Task 3 depends on Task 1, Task 2
- Task 4 depends on Task 1, Task 2, Task 3
- Task 5 depends on Task 1, Task 2, Task 3, Task 4
- Task 6 depends on Task 1, Task 2, Task 3, Task 4, Task 5
