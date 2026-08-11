# phase08-10 闭环验收报告

## 验收结论：通过 ✓

`Feedback -> Decision -> Update` 闭环已完整走通，`Decision` 保持经营中心地位，`SubmitReviewResult` 维持轻量记录身份，前后端无并列临时主线；本轮真实浏览器补验进一步确认了交互路径、来源链透传与回流结果均已成立。

---

## 1. Review → Decision 正式进入矩阵

### 已验证的进入路径

| 来源 | 目标 | 入口 | 状态 |
|------|------|------|------|
| Daily Review footer "记录决策" | Decision Create (`/decisions/new`) | `useReviewAction.create_decision` → success envelope → navigate | ✓ |
| Daily Review footer "决策中心" | Decision List (`/decisions`) | `useReviewAction.go_to_decision` → success envelope → navigate | ✓ |
| Daily Review PendingDecisionCard | Decision Detail (`/decisions/$decisionId`) | 直接 `navigate({ to: '/decisions/$decisionId' })` + `buildDashboardSourceParams` | ✓ |
| Weekly Review footer "记录决策" | Decision Create (`/decisions/new`) | 同 `create_decision` 路径 | ✓ |
| Weekly Review footer "决策中心" | Decision List (`/decisions`) | 同 `go_to_decision` 路径 | ✓ |
| Daily Review footer "产品/模块/仓库" | Product/Module/Repository List | `useReviewAction.go_to_product/go_to_module/go_to_repository` | ✓ |
| Weekly Review footer "产品/模块/仓库" | Product/Module/Repository List | 同上 | ✓ |

### 冻结结论

- **Decision Create 路径**：用户已可从 Review footer 直接点击"记录决策"进入 `DecisionCreatePage`，不再被迫先绕行 `Decision List`
- **Decision Detail 路径**：PendingDecisionCard 直接进入既有 Decision Detail，携带 `fromDashboard` + `dashboardSection` 来源参数
- **实体 handoff 路径**：`go_to_product/go_to_module/go_to_repository` 只作为 canonical 页面入口，不替代正式经营判断

---

## 2. Decision → Update 最小真实闭环

### 已验证的最小闭环路径

```
Dashboard → Daily Review → Decision Create
  → Decision Detail (success writeback) 
  → Link Decision to Module (auth-service) 
  → Product Detail (Product B) → Bind Module to Product 
  → Product Detail reread (最新结果可见)
```

### 各环节 owner 分工

| 环节 | Canonical Owner | 文件 | 状态 |
|------|----------------|------|------|
| Review 动作触发 | `useReviewAction` | `frontend/src/features/review/application/use-review-action.ts` | ✓ |
| Decision 创建 | `useCreateDraftDecision` | `frontend/src/features/decision-center/application/use-create-draft-decision.ts` | ✓ |
| Decision → Module 关联 | `useLinkDecisionToTarget` | `frontend/src/features/decision-center/application/use-link-decision-to-target.ts` | ✓ |
| Module → Product 绑定 | `useBindModuleToProduct` | `frontend/src/features/product-registry/application/use-bind-module-to-product.ts` | ✓ |
| Repository → Product 绑定 | `useBindRepositoryToProduct` | `frontend/src/features/repository-binding/application/use-bind-repository-to-product.ts` | ✓ |
| Module → Repository 映射 | `useMapModuleToRepository` | `frontend/src/features/repository-binding/application/use-map-module-to-repository.ts` | ✓ |

### 冻结结论

- 闭环可重复执行：每次从 Review 进入 Decision → Update 沿用同一套 canonical owner
- 不依赖手工 SQL、临时脚本或页面刷新补完闭环
- 闭环成功不依赖 review-local 临时状态残留

---

## 3. 成功回流、错误语义与刷新

### 成功回流矩阵

| 环节 | 成功回流 | 实现 |
|------|---------|------|
| Review → Decision handoff | 导航到 Decision Create / List / Detail | `useReviewAction` success envelope |
| Decision Create 成功 | 回流到 `DecisionDetailPage` | `useCreateDraftDecision.onSuccess` → `navigate({ to: '/decisions/$decisionId' })` |
| Decision → Module 关联成功 | 失效 `decision-detail` + `decision-module-candidates` | `useLinkDecisionToTarget.onSuccess` |
| Module → Product 绑定成功 | 失效 `product-detail` + `product-module-candidates`，页面 reread | `useBindModuleToProduct.onSuccess` |
| SubmitReviewResult 成功 | 失效 `daily-review-context` + `weekly-review-context` + `dashboard-*` | `useReviewAction.onSuccess` |

### 错误语义

- 所有错误统一由 canonical owner 归一化（`normalizeError()`）
- 页面只负责展示归一化错误（toast / error message）
- 无页面层直接解析 raw `ConnectError`、transport error

### 刷新策略

- `useReviewAction` 在 `onSuccess` 中统一失效 review + dashboard 相关 query
- 各 canonical owner 在 `onSuccess` 中失效各自 detail + candidate query
- 页面不额外持有 page-local `invalidateQueries()` 或第二套失效逻辑

---

## 4. SubmitReviewResult 边界与临时编排清理

### SubmitReviewResult 边界验证

- **后端**：`review.CommandService.SubmitReviewResult` 只写入 `review_records` 表，不承接任何实体写入
- **前端**：`useReviewAction` 只在 `submit_next_step` 动作时调用 `reviewClient.submitReviewResult`
- **数据库**：`review_records` 表只存储 `review_kind / result_kind / decision_id / target_type / target_id / summary_text / started_at / completed_at`，无实体写入字段

### 临时编排清理审计

| 审计项 | 结果 |
|--------|------|
| 前端 review 模块无 page-local `useMutation` | ✓ 通过 — 只有 `useReviewAction` 持有 `useMutation` |
| 前端 review 模块无 page-local `queryClient` / `invalidateQueries` | ✓ 通过 — 失效统一在 `useReviewAction.onSuccess` |
| 前端 review 页面无 `createClient()` / `createConnectTransport()` | ✓ 通过 — 连接统一在 `connect-client.ts` |
| 后端 review 模块无实体写入泄漏 | ✓ 通过 — `CommandService` 只写 `review_records` |
| 后端 review 模块无并列 command 主线 | ✓ 通过 — 无 `CreateDecision` / `BindModule` 等影子接口 |
| 后端 review handler 无第二套成功 envelope | ✓ 通过 — 统一使用 `.proto` 定义的 response |

---

## 5. 构建与 API 验收矩阵

### 构建验证

| 构建项 | 结果 |
|--------|------|
| `go build ./...` (backend) | ✓ 通过 |
| `npx tsc -b --noEmit` (frontend) | ✓ 通过 |
| `npm run build` (frontend production) | ✓ 通过 |
| `buf build` (proto) | ✓ 通过 |
| `buf lint` (proto) | ✓ 通过 |

### API Smoke

| API | 结果 |
|-----|------|
| `GetDailyReviewContext` | ✓ 返回 3 current_focus_signals + 1 pending_decision + 2 representative_signals |
| `GetWeeklyReviewContext` | ✓ 返回 overview + recent_activities + representative_signals + module_reuse_summary + capability_summary |
| `SubmitReviewResult` | ✓ 返回 review_record_id，写入 `review_records` 表 |

### 浏览器闭环验证

| 验证路径 | 结果 |
|---------|------|
| Dashboard → Daily Review 页面加载 | ✓ 三区块（current focus / pending decisions / representative signals）正常 |
| Daily Review → "记录决策" → Decision Create | ✓ 直接进入 `DecisionCreatePage`，`fromDashboard / dashboardSection / dashboardReturnTo` 保持透传 |
| Decision Create → 提交 → Decision Detail | ✓ 创建成功，回流到详情页，来源链继续保留 |
| Decision Detail → "查看模块" → Module Detail | ✓ 成功进入 canonical `Module Detail`，后续 update 正式下一跳可达 |
| Module Detail → "返回 Dashboard" | ✓ 成功返回 `/dashboard`，且首页已出现新建决策回流结果 |
| Decision Detail → 关联 auth-service 模块 | ✓ 关联成功，模块显示在详情页 |
| Product B Detail → 绑定 auth-service 模块 | ✓ 绑定成功，产品详情页 reread 最新结果 |
| Dashboard → Weekly Review 页面加载 | ✓ 四区块（overview / recent activity / representative signals / reuse snapshot）正常 |
| Weekly Review → "记录决策" / "决策中心" | ✓ 直接建决策与进入既有决策中心两条路径均正常 |
| Weekly Review → "完成 Review" → Dashboard | ✓ 提交成功，回到 Dashboard |

---

## 6. 影响文件清单

### 未修改业务源码（本次为补验收口 + 文档化阶段）

phase08-10 的闭环基础设施已在前序实现中完成落地。本次仅补充真实浏览器/E2E 走查并同步正式验收文档，不涉及新增业务代码变更。

### 已验证的承接文件

- `frontend/src/features/review/application/use-review-action.ts` — 唯一 review action owner
- `frontend/src/features/review/application/review-action-types.ts` — 动作类型定义
- `frontend/src/features/review/components/review-action-footer.tsx` — 底部动作区
- `frontend/src/features/review/pages/daily-review-page.tsx` — Daily Review 页面
- `frontend/src/features/review/pages/weekly-review-page.tsx` — Weekly Review 页面
- `frontend/src/features/decision-center/application/use-create-draft-decision.ts` — Decision 创建 owner
- `frontend/src/features/decision-center/application/use-link-decision-to-target.ts` — Decision 关联 owner
- `frontend/src/features/decision-center/pages/decision-create-page.tsx` — Decision 创建页面
- `frontend/src/features/product-registry/application/use-bind-module-to-product.ts` — Module→Product 绑定 owner
- `frontend/src/features/repository-binding/application/use-bind-repository-to-product.ts` — Repository→Product 绑定 owner
- `frontend/src/features/repository-binding/application/use-map-module-to-repository.ts` — Module→Repository 映射 owner
- `backend/internal/review/service/command_service.go` — Review 写服务
- `backend/internal/review/service/query_service.go` — Review 读服务
- `backend/internal/review/connect/server.go` — Review Connect handler

---

## 7. 与 spec 的对齐结论

| Spec 要求 | 对齐状态 |
|-----------|---------|
| review 从"可进入"推进为"可重复闭环" | ✓ 闭环可重复执行，不依赖临时状态 |
| Decision 继续作为 review loop 正式中心 | ✓ 实体 handoff 不替代正式决策承接 |
| 至少一种实体 update 真实执行成功并 reread | ✓ BindModuleToProduct 路径完整走通 |
| 成功回流、错误语义、刷新服从 owner 单值化 | ✓ 无 page-local 失效或第二套编排 |
| SubmitReviewResult 退回过程记录身份 | ✓ 只写 review_records，不冒充实体更新 |
| 前后端不保留并列临时主线 | ✓ 审计通过，无临时编排残留 |
| 验收覆盖闭环 + 清理两类证据 | ✓ 构建 + API smoke + 浏览器闭环 + 审计清单 |
