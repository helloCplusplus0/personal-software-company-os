# phase05-14 Dashboard + Feedback 联调、验证与验收报告

## 1. 验收环境

### 1.1 运行时入口与前置条件

| 项目 | 实际结果 |
|------|---------|
| 前端验收入口 | `http://127.0.0.1:5173/dashboard` |
| 后端健康检查 | `http://127.0.0.1:8081/healthz -> 200 {"status":"ok"}` |
| 前端真实 API 模式 | `frontend/.env` 已配置 `VITE_USE_REAL_API=true`、`VITE_API_BASE_URL=` |
| 开发期 API 代理 | Vite proxy 已将 `/api` 代理到 `localhost:8081` |
| 统一重置入口 | `database/scripts/reset_dashboard_acceptance.sh` |
| fixture 白名单 | 已由 `reset_dashboard_acceptance.sh` 统一承接，不使用手工 SQL 替代 |

### 1.2 合同与单值链路核对

本轮已确认 `.proto -> HTTP -> 前端` 单值链路成立，未出现并列 DTO 或 page-only contract：

- `GetDashboardOverviewResponse { overview }`
- `GetFeedbackSignalsResponse { current_focus_signals, asset_feedback_summary }`
- `GetRecentActivitiesResponse { activities }`
- 后端 handler 与前端 adapter 已按上述字段语义对齐

## 2. 最小主线端到端验收

### 2.1 Dashboard 聚合读最小闭环

#### 空系统基线

- `empty-system` 下：
  - `overview -> 200`，返回全零 envelope
  - `feedback-signals -> 200`，`current_focus_signals=[]`，`asset_feedback_summary` 全零
  - `recent-activities -> 200`，`activities=[]`

#### 默认恢复基线

- 默认恢复后，`overview` 返回非零概览：`2 / 3 / 2 / 3 / 1 / 1`

#### `recent-activities` fixture 基线

- `overview = 1 / 1 / 1 / 1 / 1 / 1`
- `feedback-signals -> 200`，为空成功态
- `recent-activities -> 200`，覆盖 8 类活动类型

#### 局部错误模拟

- 已确认错误模拟入口存在并实测通过：
  - `DASHBOARD_SIMULATE_FEEDBACK_ERROR`
  - `DASHBOARD_SIMULATE_RECENT_ACTIVITY_ERROR`
- 在 `8082 / 8083` 端口实测确认：附属聚合可独立返回 `500`，`overview` 仍保持 `200`

**结论**：`/dashboard -> overview / feedback-signals / recent-activities` 的最小聚合读链路已在真实前后端与真实数据库上跑通。`overview` 不会被附属聚合错误拖垮。✅

### 2.2 浏览器联调结果

#### 空状态

- 访问 `/dashboard` 后展示全零概览
- 主 CTA 文案为 `登记首个模块`
- 点击后跳转到：
  - `/modules/new?fromDashboard=true&dashboardSection=empty-state&dashboardReturnTo=%2Fdashboard`
- 关键空态文案均正确：
  - `暂无待处理反馈信号`
  - `暂无代表性缺口项`
  - `暂无最近活动`
- 截图：
  - `/tmp/trae/screenshots/dashboard-empty-state-verification.png`

#### 有数据状态

- 页面可见非零概览
- 反馈区块可表现为真实反馈信号或中性成功态
- `Recent Activity` 正常显示
- 最近活动区块已按限高 + 内部滚动实现，页面不会被无限拉长

#### 局部错误状态

- 使用 `5174` 前端 + `8082` 后端验证 feedback error
- 页面保持 `ready`，未升级为整页 `page-error`
- `overview / stat bar` 正常
- `Current Focus` 显示：`反馈信号读取失败：internal server error`
- `Asset Feedback` 显示：`资产缺口读取失败：internal server error`
- `Recent Activity` 保持正常显示
- 点击局部 `重试` 后仍保持局部错误，不升级为整页错误
- 截图：
  - `/tmp/trae/screenshots/dashboard-5174-local-error-check.png`

**结论**：Dashboard 空状态、有数据状态与局部错误状态均已完成真实浏览器复测，页面级与区块级语义符合正式规格。✅

## 3. 状态矩阵验收结果

| 状态场景 | 建立方式 | 实际结果 |
|----------|----------|----------|
| 空系统 | `reset_dashboard_acceptance.sh --fixture empty-system` | 概览全零成功态、主 CTA 正确、三块空态文案正确，已通过 |
| 默认恢复 | 默认恢复 | `overview` 非零 `2 / 3 / 2 / 3 / 1 / 1`，页面进入真实有数据状态，已通过 |
| `recent-activities` fixture | `reset_dashboard_acceptance.sh --fixture recent-activities` | `overview = 1 / 1 / 1 / 1 / 1 / 1`，Recent Activity 覆盖 8 类活动类型，已通过 |
| Feedback Signals 局部错误 | `DASHBOARD_SIMULATE_FEEDBACK_ERROR=true` | 反馈区块局部失败，`overview` 与页面 ready 语义保持稳定，已通过 |
| Recent Activities 局部错误 | `DASHBOARD_SIMULATE_RECENT_ACTIVITY_ERROR=true` | 已确认附属聚合可独立 `500` 且不拖垮 `overview`，局部错误隔离成立，已通过 |

## 4. 聚合读取与前端消费闭环

### 4.1 Overview

| 检查项 | 实际结果 |
|--------|---------|
| HTTP 响应结构 | 返回 `{"overview": ...}` envelope |
| 零值成功态 | `empty-system` 下返回全零 envelope，前端按成功空态消费 |
| 非零成功态 | 默认恢复返回 `2 / 3 / 2 / 3 / 1 / 1`，`recent-activities` fixture 返回 `1 / 1 / 1 / 1 / 1 / 1`，前端可稳定展示 |

### 4.2 Feedback Signals

| 检查项 | 实际结果 |
|--------|---------|
| HTTP 响应结构 | 同次返回 `current_focus_signals` 与 `asset_feedback_summary` |
| 空成功态 | `empty-system` 与 `recent-activities` fixture 下均返回 `200` 成功态，不伪造错误 |
| 局部错误隔离 | `DASHBOARD_SIMULATE_FEEDBACK_ERROR=true` 时，反馈区块局部失败但不升级为整页错误 |
| 前端共享消费 | `Current Focus` 与 `Asset Feedback` 共享同一次读取结果，handler 与 adapter 已对齐 |

### 4.3 Recent Activities

| 检查项 | 实际结果 |
|--------|---------|
| HTTP 响应结构 | 返回 `{"activities": [...]}` envelope |
| 空数组成功态 | `empty-system` 下返回 `activities=[]`，前端按成功空态展示 |
| 活动类型覆盖 | `recent-activities` fixture 已覆盖 8 类活动类型 |
| 排序稳定性 | `seed_dashboard_fixture_recent_activities.sql` 已补齐 8 类活动的显式可排序时间证据，reset 后通过 `/api/dashboard/recent-activities` 复测顺序正确 |

## 5. 跳转与返回路径验收

### 5.1 直接跳转返回

- `Dashboard -> Decision Detail -> Dashboard` 通过
- `Dashboard -> Product Detail -> Dashboard` 通过
- `Dashboard -> Repository Detail -> Dashboard` 通过
- `Dashboard -> Module Detail -> Dashboard` 通过

### 5.2 多跳返回链

- `Dashboard -> Modules List -> Module Detail -> List -> Dashboard` 通过
- URL 同时保留：
  - `fromList=true`
  - `fromDashboard=true&dashboardSection=overview&dashboardReturnTo=/dashboard`

### 5.3 来源区块恢复

- `Current Focus` 恢复通过：
  - 返回后 `Current Focus` 区块位于视口内
  - `document.activeElement` 落在该区块内容容器
- `overview` 恢复通过：
  - 返回后 overview 区块位于视口内
  - `document.activeElement` 落在统计条容器
- 截图：
  - `/tmp/trae/screenshots/dashboard-current-focus-restored.png`
  - `/tmp/trae/screenshots/dashboard-overview-restored.png`

**结论**：Dashboard 到四类 canonical owner 的直接返回、多跳返回与来源区块恢复均已通过真实复测；`fromList` 与 `fromDashboard` 可以并存恢复。✅

## 6. 合同与正式规格一致性核对

- `.proto` 合同主线字段与 HTTP envelope 保持单值一致：
  - `overview`
  - `current_focus_signals`
  - `asset_feedback_summary`
  - `activities`
- 前端运行于真实 API 模式，开发期通过 Vite proxy 访问真实 `/api/dashboard/*`
- 后端 handler 与前端 adapter 已按同一字段语义消费，不存在额外 page-only contract
- Dashboard 主链保持 `真实 API -> HTTP envelope -> 前端区块渲染` 的单一消费路径

## 7. 联调中发现的问题与修复

| # | 问题 | 级别 | 处理结果 | 复测结论 |
|---|------|------|----------|----------|
| 1 | 返回 Dashboard 后只渲染 `sr-only` 提示，未真正恢复来源区块视图上下文 | P1 | 已在 `frontend/src/features/dashboard/pages/dashboard-home-page.tsx` 修复为按 `dashboardSection` 执行 `scrollIntoView + focus`，并清理一次性 state | 已复测通过 |
| 2 | `seed_dashboard_fixture_recent_activities.sql` 仅 release 带显式时间，其余活动时间证据不足 | P1 | 已在 `database/seeds/seed_dashboard_fixture_recent_activities.sql` 为 8 类活动补齐显式可排序时间证据 | 已通过 reset + `/api/dashboard/recent-activities` 复测顺序正确 |
| 3 | 从 `backend/` 目录启动时 `.env` 中 migrations / seeds 相对路径失效 | P1 | 已在 `backend/cmd/server/main.go`、`backend/internal/platform/config.go`、`backend/.env.example` 修复为显式尝试加载 `.env / backend/.env / ../backend/.env`，并将路径解析到仓库真实位置 | 已用 `HTTP_PORT=18081 go run ./cmd/server` 复测通过 |

## 8. 阶段收口结论

- `Dashboard + Feedback` 最小主线已在真实前端、真实后端、真实数据库上完整走通 ✅
- 空状态、有数据状态、局部错误状态与 Recent Activity 展示已形成可重复复核的验收矩阵 ✅
- `overview / feedback-signals / recent-activities` 三类聚合读取与前端消费闭环已形成单一事实链路 ✅
- Dashboard 到 `Product / Repository / Module / Decision` 的跳转、返回与来源区块恢复已完成真实复测 ✅
- 本轮发现的问题已在当前阶段修复并复测通过，不遗留隐性阻断 ✅

补充说明：

- 浏览器在切换 fixture、重置环境或服务刚启动后的首次观察中，可能出现短暂旧数据或空态瞬时现象；本轮结论以重置后稳定复测结果为准。
- 基于本轮真实结果，`phase05-14` 通过。

## 9. DoD 达成情况

| DoD 项 | 达成情况 |
|--------|---------|
| `Dashboard + Feedback` 最小主线可完整走通 | ✅ |
| 最小反馈信号、聚合读取与跳转返回路径可被重复复核 | ✅ |
| 验收结果可重复复核，并可明确证明当前阶段已形成可运行交付物 | ✅ |
| 发现的问题已收口到当前阶段，不遗留隐性阻断 | ✅ |
