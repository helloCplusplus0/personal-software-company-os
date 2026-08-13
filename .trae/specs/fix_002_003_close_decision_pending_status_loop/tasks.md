# Tasks

- [x] Task 1: 落地 `Decision.status` 的正式后端写链。
  - [x] SubTask 1.1: 在 `decision_center.proto` 中新增状态推进写接口与请求/响应消息，保持四态不变
  - [x] SubTask 1.2: 在 `backend/internal/decisioncenter/types.go` 中新增对应 domain request/response 类型
  - [x] SubTask 1.3: 在 `backend/internal/decisioncenter/repository/decision_store.go` 中新增单一 `UpdateStatus` 持久化入口
  - [x] SubTask 1.4: 在 `backend/internal/decisioncenter/service/command_service.go` 中新增状态推进编排与最小状态迁移校验
  - [x] SubTask 1.5: 在 `backend/internal/decisioncenter/connect/server.go` 中接入新的状态推进 handler 与错误映射

- [x] Task 2: 在 `Decision Detail` 中承接正式状态推进动作。
  - [x] SubTask 2.1: 新增前端固定 mutation owner，统一承接状态推进成功/失败与 query invalidation
  - [x] SubTask 2.2: 在 `Decision Detail` 概要区或等价固定区域展示当前状态与最小 CTA 组
  - [x] SubTask 2.3: 按 spec 落地最小状态推进矩阵，终态不再暴露 CTA
  - [x] SubTask 2.4: 保持 `Decision Detail` 仍为单页壳层，不拆新路由、不长第二工作台

- [x] Task 3: 统一 pending 语义与 reread 行为。
  - [x] SubTask 3.1: 保持 `Dashboard / Current Focus / Daily Review` 的 pending 判定继续锚定 canonical `Decision.status`
  - [x] SubTask 3.2: 禁止以 `decision_links` 或 `review_records` 代理 pending 退出语义
  - [x] SubTask 3.3: 状态推进成功后统一失效 `decision-detail / decision-list / review / dashboard` 相关读取

- [x] Task 4: 完成联动验证与回归。
  - [x] SubTask 4.1: 验证 `proposed` 决策可在 `Decision Detail` 中推进到 `active / superseded / archived`
  - [x] SubTask 4.2: 验证 `active` 决策可继续推进到 `superseded / archived`
  - [x] SubTask 4.3: 验证终态不再展示状态推进 CTA
  - [x] SubTask 4.4: 验证从 `Daily Review / Dashboard` 进入详情页后，状态推进成功能回流并消除 pending 误报
  - [x] SubTask 4.5: 验证 `SubmitReviewResult` 与 `LinkDecisionToTarget` 的职责边界未被破坏

# Task Dependencies
- Task 2 depends on Task 1
- Task 3 depends on Task 1
- Task 4 depends on Task 1
- Task 4 depends on Task 2
- Task 4 depends on Task 3
