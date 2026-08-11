# Tasks

- [x] Task 1: 落实 review proto 合同与生成主线
  - [x] SubTask 1.1: 新增 `proto/psco/review/v1/review.proto`，只定义 `GetDailyReviewContext / GetWeeklyReviewContext / SubmitReviewResult`
  - [x] SubTask 1.2: 对齐 `buf` 生成主线，确保后端 Go proto / Go Connect 与前端 TS 产物都生成到既有落点
  - [x] SubTask 1.3: 确认 review proto 复用既有 canonical message，不重复抄字段

- [x] Task 2: 落实 review 后端承接位与数据承接
  - [x] SubTask 2.1: 新增 `backend/internal/review/service/query_service.go`，组合 `dashboard / decisioncenter / reusesummary` 既有 query service
  - [x] SubTask 2.2: 新增 `backend/internal/review/service/command_service.go` 与 `repository/review_record_store.go`，只承接 `next-step result` 的最小持久化
  - [x] SubTask 2.3: 新增 review record migration，并对齐 `review_records` 单表最小字段集合
  - [x] SubTask 2.4: 新增 `backend/internal/review/connect/server.go` 并接入 `platform/router.go`

- [x] Task 3: 落实 review 前端切片与 owner 收敛
  - [x] SubTask 3.1: 新增 `frontend/src/features/review/data/connect-client.ts` 与 `review-query-options.ts`
  - [x] SubTask 3.2: 新增 `use-daily-review-read.ts` 与 `use-weekly-review-read.ts`，输出稳定的 page / section 状态模型
  - [x] SubTask 3.3: 新增 `use-review-action.ts` 与 `review-action-types.ts`，统一承接 success envelope、query invalidation 与错误归一化

- [x] Task 4: 落实 review route、page shell 与最小页面
  - [x] SubTask 4.1: 新增 `/reviews/daily` 与 `/reviews/weekly` 文件路由，并接入 `dashboardSourceSearchSchema`
  - [x] SubTask 4.2: 新增 `ReviewPageShell`、`ReviewActionFooter`、`DailyReviewPage`、`WeeklyReviewPage`
  - [x] SubTask 4.3: 确保页面只消费 review slice owner，不直接拼底层 query / mutation

- [x] Task 5: 收敛 Dashboard review 入口 caller
  - [x] SubTask 5.1: 将 `DashboardPrimaryActionPanel` 从旧单 CTA 命中器切换为 dual review launcher
  - [x] SubTask 5.2: 在 launcher 中正式导航到 `/reviews/daily` 与 `/reviews/weekly`，并透传 `buildDashboardSourceParams('empty-state')`
  - [x] SubTask 5.3: 对齐 `DashboardHomePageShell` 的响应式布局责任，避免把 review 编排下沉到 panel

- [x] Task 6: 完成构建与关键路径验收
  - [x] SubTask 6.1: 执行 `(cd proto && make build && make gen && make lint)` — ✅ 全部通过
  - [x] SubTask 6.2: 执行 `(cd backend && go build ./...)` — ✅ 通过
  - [x] SubTask 6.3: 执行 `(cd frontend && npm run build)` — ✅ 通过
  - [x] SubTask 6.4: 覆盖 review API smoke 与 `/dashboard -> /reviews/*` 路由 smoke — ✅ 全部通过

# 实现总结

## 修复项
- `backend/internal/review/repository/review_record_store.go`：修复 NULL 扫描错误——`decision_id`/`target_type`/`target_id` 为数据库 NULL 可空列，但 `ReviewRecord` 对应字段为 `string`（非指针），pgx 无法将 NULL 直接扫描到 `string` 字段地址。使用中间 `*string` 变量承接后再解引用赋值。
- `frontend/src/features/review/application/use-review-action.ts`：修复 `submit_next_step` 硬编码为 `ReviewKind.DAILY`——在 `ReviewActionInput` 新增 `reviewKind` 字段，由 `ReviewActionFooter` 从页面透传当前 review 类型，`useReviewAction` 根据 `reviewKind` 选择 `ReviewKind.DAILY` 或 `ReviewKind.WEEKLY`。
- `frontend/src/features/review/application/use-review-action.ts` / `review-action-types.ts` / `review-action-footer.tsx`：独立复核后收敛 review 正式动作面——移除尚未冻结为正式能力的 `create_decision` 路径，避免把 Dashboard 来源参数送入 `/decisions/new` 后被路由 schema 静默丢失；同时补齐 `go_to_repository` 的用户可达入口，使 `Repository` handoff 不再停留在类型与 owner 预留分支。

## 验证结果
- proto: `buf build` ✅ | `buf gen` ✅ | `buf lint` ✅
- backend: `go build ./...` ✅
- frontend: `npm run build` ✅
- API smoke: `GetDailyReviewContext` ✅ | `GetWeeklyReviewContext` ✅ | `SubmitReviewResult`（含 NULL 和非 NULL 路径）✅
- 数据库: `review_records` 记录正确写入 ✅
- 前端路由: `/reviews/daily` ✅ | `/reviews/weekly` ✅
- Dashboard launcher: `DashboardPrimaryActionPanel` 已收敛为 dual review launcher ✅
- 页面 owner 收敛: `DailyReviewPage` 只消费 `useDailyReviewRead / useReviewAction` ✅ | `WeeklyReviewPage` 只消费 `useWeeklyReviewRead / useReviewAction` ✅
- Read owner transport: `review-query-options.ts` 通过 `reviewClient`（ReviewService generated client）消费 ✅
- 浏览器 smoke: `/dashboard -> /reviews/daily` ✅ | `/dashboard -> /reviews/weekly` ✅ | review 页面进入 `Decision / Product / Module / Repository` canonical path 时继续透传 `fromDashboard / dashboardSection / dashboardReturnTo` ✅
- 动作闭环 smoke: `Weekly Review -> SubmitReviewResult -> /dashboard` ✅，后端日志确认 `ReviewService/SubmitReviewResult` 实际命中并返回 `200`

# Task Dependencies

- `Task 2` depends on `Task 1`
- `Task 3` depends on `Task 1`
- `Task 4` depends on `Task 3`
- `Task 5` depends on `Task 4`
- `Task 6` depends on `Task 1`
- `Task 6` depends on `Task 2`
- `Task 6` depends on `Task 4`
- `Task 6` depends on `Task 5`
