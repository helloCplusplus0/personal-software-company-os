# Tasks

- [x] Task 1: 冻结 review 到 `Decision` 的正式进入矩阵
  - [x] SubTask 1.1: 明确 `Daily Review / Weekly Review` 中哪些动作必须先进入 `Decision Create` 或 `Decision Detail`
  - [x] SubTask 1.2: 明确 review 来源参数、成功回流与 `DecisionDetailPage` 的最小承接关系
  - [x] SubTask 1.3: 明确"直接实体 handoff 只是 canonical 入口，不等于经营判断已正式承接"的边界

- [x] Task 2: 冻结 `Decision -> Update` 的最小真实闭环
  - [x] SubTask 2.1: 在 `BindModuleToProduct / BindRepositoryToProduct / MapModuleToRepository` 中选择至少一条作为 `phase08-10` 必做闭环
  - [x] SubTask 2.2: 明确该闭环必须继续复用哪些 canonical application owner、canonical page 与 reread 结果
  - [x] SubTask 2.3: 明确 route handoff、Decision 正式承接与实体 update 成功三者的先后关系

- [x] Task 3: 冻结成功回流、错误语义与必要刷新
  - [x] SubTask 3.1: 明确 `useReviewAction`、`useCreateDraftDecision` 与所选实体 canonical owner 的失效矩阵分工
  - [x] SubTask 3.2: 明确成功后页面如何回流到 `DecisionDetailPage` 或实体 canonical 页面，并以 reread 结果作为完成依据
  - [x] SubTask 3.3: 明确错误语义仍由 canonical owner / review owner 统一归一化，不允许页面层直连 raw transport error

- [x] Task 4: 冻结 `SubmitReviewResult` 与临时散装编排清理边界
  - [x] SubTask 4.1: 明确 `SubmitReviewResult` 继续只承接轻量流程记录，不升级为并列实体写入主线
  - [x] SubTask 4.2: 列出前端必须回收的临时编排点，避免保留第二套 review-local mutation / invalidation 主线
  - [x] SubTask 4.3: 列出后端必须回收的临时编排点，避免保留 review-local 并列 command / handler 语义

- [x] Task 5: 冻结 `phase08-10` 的闭环验收矩阵
  - [x] SubTask 5.1: 明确至少一条 `Review -> Decision -> canonical update -> reread` 的最小验收路径
  - [x] SubTask 5.2: 明确 `SubmitReviewResult` 的辅助验收位置，避免其冒充实体 update
  - [x] SubTask 5.3: 明确前后端"无并列临时主线"的审查清单
  - [x] SubTask 5.4: 明确构建、API smoke、浏览器闭环验证三层证据各自承担什么结论

# Task Dependencies

- `Task 2` depends on `Task 1`
- `Task 3` depends on `Task 1`
- `Task 3` depends on `Task 2`
- `Task 4` depends on `Task 1`
- `Task 4` depends on `Task 2`
- `Task 5` depends on `Task 1`
- `Task 5` depends on `Task 2`
- `Task 5` depends on `Task 3`
- `Task 5` depends on `Task 4`

## 修复记录

- `frontend/src/features/review/application/review-action-types.ts` / `use-review-action.ts` / `review-action-footer.tsx`：恢复 `create_decision` 正式动作，将 Review 主按钮从“进入决策列表”收敛为“记录决策 -> /decisions/new”，同时保留 `go_to_decision` 作为进入既有 `Decision Center` 的辅路径
- `frontend/src/routes/decisions/new.tsx` / `decision-list-page.tsx` / `decision-create-page.tsx`：让 `Decision List -> Decision Create -> Decision Detail` 全链路继续透传 `fromDashboard / dashboardSection / dashboardReturnTo`，不再在新建决策分支丢失 review/dashboard 来源链
- `frontend/src/features/decision-center/pages/decision-detail-page.tsx` / `decision-linked-targets-section.tsx` / `decision-pending-link-target-card.tsx` / `decision-module-candidate-panel.tsx`：补齐 `Decision Detail -> Module Detail` 正式下一跳，使 `Decision` 正式承接后能继续进入既有 `Module Detail -> Product/Repository canonical update` 主线，而不是停留在只读 badge

## 验收补记

- 前端构建：`frontend/npm run build` 通过
- 后端回归构建：`backend/go build ./...` 通过
- 最小运行验证：本地后端 `go run ./cmd/server` 成功启动并通过 `GET /healthz`
- 真实浏览器/E2E 走查：已在用户开启的前后端服务上完成 `Dashboard -> Daily Review -> Decision Create -> Decision Detail -> Module Detail -> Dashboard` 实际点击验收；`fromDashboard / dashboardSection / dashboardReturnTo` 全链路透传保持成立，且 Dashboard 已出现新建决策的回流结果
- 当前结论：`phase08-10` 已完成源码修复、构建验证、运行时健康检查与真实浏览器交互验收，可按通过验收并收口该子任务
