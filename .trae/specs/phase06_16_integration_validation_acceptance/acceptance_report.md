# phase06-16 联调、验证与验收报告

## 1. 验收环境

### 1.1 运行时入口与前置条件

| 项目 | 实际结果 |
|------|---------|
| 前端验收入口 | `http://127.0.0.1:5173` |
| 后端健康检查 | `http://127.0.0.1:8081/healthz -> 200 {"status":"ok"}` |
| 前端真实 API 模式 | `frontend/.env` 已配置 `VITE_USE_REAL_API=true`、`VITE_API_BASE_URL=` |
| 后端运行方式 | `go run ./cmd/server`（从 `backend/` 目录启动） |
| 数据库 | `rento-preview-postgres` 容器运行中 |
| 统一重置入口 | `database/scripts/reset_phase06_acceptance.sh` |
| fixture 白名单 | 已按正式脚本加载，不使用手工 SQL 替代验收路径 |

### 1.2 合同与单值链路核对

本轮已确认以下正式合同均通过真实 HTTP 出口读取，且前后端消费语义保持单值一致：

- `GET /api/onboarding/state -> { first_run_state }`
- `GET /api/dashboard/export -> { snapshot }`
- `POST /api/dashboard/export -> { snapshot }`
- `GET /api/dashboard/backup -> { snapshot }`
- `POST /api/dashboard/backup -> { snapshot }`
- `GET /api/reuse-summary -> { module_reuse_summary, capability_summary }`

补充说明：

- 当前浏览器验收优先使用 `integrated_browser` 完成；`mcp_chrome-devtools` 在本机因缺少 Chrome stable 可执行文件未作为主验收通道。

## 2. 实际执行的 fixture 与验收入口

本轮实际使用的正式 fixture：

- `cold-start-empty`
- `in-progress-partial-entry`
- `completed-bound`
- `export-ready`
- `backup-verified`
- `backup-manifest-missing`
- `backup-coverage-incomplete`
- `backup-schema-mismatch`
- `reuse-latest-after-binding`

说明：

- `completed-bound` 覆盖了 `completed-unbound` 的正式成功态验收目标；
- `reuse-latest-after-binding` 覆盖了“读取最新已提交状态”的验收目标，因此未再单独保留 `reuse-latest` 作为最终证据。

## 3. 首轮录入与根级入口验收

### 3.1 cold-start 根级入口

使用 `reset_phase06_acceptance.sh --fixture cold-start-empty` 后：

- `GET /api/onboarding/state` 返回：
  - `status = not_started`
  - `current_step = welcome`
  - `completion_progress = 0`
- 浏览器访问 `/` 自动进入 `/onboarding`
- `/dashboard` 中展示 `开始首轮录入`
- Dashboard 同时保持：
  - Export 为 preview 成功态
  - Backup 为正式空态 `暂无备份记录。点击下方按钮触发首次备份。`
  - `overview / feedback-signals / recent-activities` 继续为成功空态

结论：cold-start 根级入口、首轮主 CTA 与空系统语义通过。✅

### 3.2 `in_progress` 回访继续

使用 `reset_phase06_acceptance.sh --fixture in-progress-partial-entry` 后：

- `GET /api/onboarding/state` 返回：
  - `status = in_progress`
  - `current_step = repository`
  - `completion_progress = 50`
- 浏览器访问 `/` 自动进入 `/dashboard`
- Dashboard 展示 `继续首轮录入`
- 浏览器访问 `/onboarding` 时，页面自动定位到 `创建仓库` 步骤，而不是回到 `welcome`

结论：`in_progress` 回访继续入口与步骤自动定位通过。✅

### 3.3 首轮成功会话真实跑通

在 `cold-start-empty` 基础上，直接通过正式 create API 完成一次真实首轮会话：

- `POST /api/products` 创建 `Acceptance Product`
- `POST /api/repositories` 创建 `acceptance-repo`
- `POST /api/modules` 创建 `acceptance-module`
- `POST /api/decisions` 创建 `Acceptance Decision`

实际返回的正式持久化 ID：

- Product: `20f0a64b-a4d4-4c60-85a5-0546fd5f4818`
- Repository: `dc089f41-6eb5-46ee-81a3-44b9e430d4eb`
- Module: `a732ad67-d7c5-489a-9aa7-b549627ce7fd`
- Decision: `5c6e9d56-4010-4191-8f07-e6ff0be20056`

随后回读：

- `GET /api/onboarding/state -> status = completed, current_step = complete, completion_progress = 100`
- 数据库真实计数：
  - `products = 1`
  - `repositories = 1`
  - `modules = 1`
  - `decisions = 1`

结论：首轮成功会话已通过真实写入链路走通，并验证四类对象均已真实持久化。✅

### 3.4 `fromOnboarding` 回流链

使用 `completed-bound` fixture，通过浏览器执行真实回流链：

1. 进入 `/onboarding?...productDraftId=...&productDraftLabel=Product A`
2. 页面展示 `继续编辑产品` 与 `返回完成页`
3. 点击 `继续编辑产品` 进入 Product Detail，URL 带有：
   - `fromOnboarding=true`
   - `onboardingStep=product`
4. 点击 Product Detail 中的 `返回首轮录入`
5. 成功返回 `/onboarding`，并恢复到产品步骤摘要，而不是掉回完成页

结论：`Onboarding -> canonical detail -> /onboarding` 回流链已通过真实浏览器复测。✅

## 4. Export / Backup 数据主权路径验收

### 4.1 Export 覆盖矩阵

使用 `reset_phase06_acceptance.sh --fixture export-ready` 后：

- `GET /api/dashboard/export` 返回 preview 快照
- `POST /api/dashboard/export` 返回正式导出结果
- 返回的 `asset_scope` 覆盖 9 类正式资产：
  - `products`
  - `modules`
  - `releases`
  - `repositories`
  - `decisions`
  - `decision_links`
  - `product_modules`
  - `product_repositories`
  - `module_repositories`

结论：导出正式覆盖矩阵通过，且包含绑定 / 关联关系，不是只导出主实体。✅

### 4.2 Backup Verified 与重复执行

使用 `reset_phase06_acceptance.sh --fixture backup-verified` 后，连续执行两轮：

- `POST /api/dashboard/backup`
- `GET /api/dashboard/backup`

实际结果：

- 每次 `POST` 均先返回 `verified_status = unverified`
- 随后 `GET` 均返回 `verified_status = verified`
- `manifest_summary`
- `asset_coverage`
- `schema_version_prerequisite`
  三组校验前提均存在且可读

结论：`backup verified` 的正式语义成立，且备份动作可重复执行。✅

### 4.3 Backup 三类失败语义

分别使用三类正式 fixture 复测：

| fixture | 实际读取结果 |
|--------|-------------|
| `backup-manifest-missing` | `verify_failed + manifest_missing` |
| `backup-coverage-incomplete` | `verify_failed + coverage_incomplete` |
| `backup-schema-mismatch` | `verify_failed + schema_mismatch` |

同时在 Dashboard 浏览器界面确认：

- `校验状态 -> 校验失败`
- `失败原因 -> manifest 缺失`

结论：三类失败语义已在 `.proto -> HTTP DTO -> 前端展示` 边界保持单值一致，没有被折叠成泛化错误。✅

## 5. 复用反馈与最新状态验收

### 5.1 Dashboard 复用快照

使用 `reset_phase06_acceptance.sh --fixture reuse-latest-after-binding` 后：

- `GET /api/reuse-summary?scope=dashboard` 返回：
  - `auth-service -> reuse_product_count = 3`
  - `integration-test-module -> reuse_product_count = 3`
  - `capability_summary` 包含：
    - `Web Frontend -> 1`
    - `Authentication -> 1`

Dashboard 浏览器文本同时展示：

- `模块「auth-service」当前被 3 个 Product 复用`
- `模块「integration-test-module」当前被 3 个 Product 复用`
- `能力分布`
- `Web Frontend`
- `Authentication`

结论：Dashboard 复用快照已反映最新已提交状态。✅

### 5.2 Detail 页复用反馈

在同一 fixture 下使用真实 UUID 复测：

- `GET /api/reuse-summary?scope=module_detail&module_id=2298610b-7227-4de8-8f99-188f38aa149f`
- `GET /api/reuse-summary?scope=product_detail&product_id=0a6a3e28-4e20-4f8e-b760-062316017124`

实际结果：

- Module Detail 返回单模块复用摘要 + 对应 capability 摘要
- Product Detail 返回当前 Product 绑定范围内的复用摘要与 capability 摘要
- 未观察到 detail 读取承接任何绑定写入或筛选写语义

结论：Detail 页复用反馈路径通过。✅

## 6. `phase05` 兼容性与局部边界

在 cold-start 与 `reuse-latest-after-binding` 两类状态下，均复测了：

- `GET /api/dashboard/overview`
- `GET /api/dashboard/feedback-signals`
- `GET /api/dashboard/recent-activities`

实际结果：

- cold-start 下三者均返回成功空态
- `reuse-latest-after-binding` 下三者继续返回正式数据，不受 `phase06` 新增块破坏
- Backup 失败状态停留在 `数据主权` 子区域，不会把整页拖垮
- `Asset Feedback` 与 `Reuse Snapshot` 可同时存在，未再出现 sibling loading 互相遮挡

结论：`phase05` 已交付的 Dashboard 主线兼容性通过，未发现回归。✅

## 7. 联调中发现的问题与修复

| # | 问题 | 级别 | 处理结果 | 复测结论 |
|---|------|------|----------|----------|
| 1 | `GET /api/dashboard/backup` 在无备份记录时返回空快照对象，前端被误渲染为异常时间，而不是正式空态 | P1 | 已修复 [backend/internal/backup/handler/query_handler.go](file:///home/dell/Projects/personal-software-company-os/backend/internal/backup/handler/query_handler.go)；空态改回 `snapshot: null` | 冷启动下 API 返回 `{\"snapshot\":null}`，Dashboard 正常展示 `暂无备份记录` |
| 2 | `completed` 状态会清掉从 detail 页回流的一次性 `focusedStep`，导致“返回首轮录入”掉回完成页而不是目标步骤 | P1 | 已修复 [frontend/src/features/onboarding/pages/onboarding-page.tsx](file:///home/dell/Projects/personal-software-company-os/frontend/src/features/onboarding/pages/onboarding-page.tsx)；保留 `focusedStep` 对 `completed` 默认完成页的优先级 | 真实浏览器链路已确认可回到产品步骤摘要，并可手动返回完成页 |

## 8. 额外静态验证

- `frontend`: `npm run lint`
  - 结果：0 error，4 条既有 warning（UI primitives / `__root.tsx`，非本轮新增）
- `frontend`: `npm run build`
  - 结果：通过

## 9. 阶段收口结论

- `phase06` 的首轮录入、数据主权与复用反馈最小主线已在真实前端、真实后端、真实数据库上完成联调验收 ✅
- 首轮成功会话已通过真实 create API 走通，并验证四类对象均已持久化 ✅
- `Export` 覆盖矩阵、`Backup` 校验语义与三类失败语义均已通过验收 ✅
- `module_reuse_summary / capability_summary` 已证明能反映最新已提交状态 ✅
- `phase05` Dashboard 主线兼容性保持稳定，未发现破坏性回归 ✅
- 本轮发现的两个真实阻断项已在当前阶段修复并复测通过，不遗留未收口问题 ✅

因此，`phase06-16` 通过。

## 10. DoD 达成情况

| DoD 项 | 达成情况 |
|--------|---------|
| 首轮成功会话可重复走通 | ✅ |
| 数据主权路径可验证 | ✅ |
| 复用反馈路径可验证 | ✅ |
| 首轮成功会话、导出覆盖矩阵与复用感知最新状态均通过验收 | ✅ |
| 无破坏 `phase05` 已交付边界的回归 | ✅ |
| 根级入口行为、回访继续入口与备份校验语义均通过验收 | ✅ |
