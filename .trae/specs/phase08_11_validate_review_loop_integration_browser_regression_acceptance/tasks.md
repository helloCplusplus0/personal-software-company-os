# Tasks

- [x] Task 1: 冻结 `phase08-11` 的统一验收输入边界与上游证据来源
  - [x] SubTask 1.1: 对齐 `docs/phase/phase08_operating_review_loop_foundation_dev_plan.md#L217-L239` 的范围、DoD 与非目标
  - [x] SubTask 1.2: 继承 `phase08-08 / 09 / 10` 的已交付能力，明确它们在本阶段只作为上游输入与待复验项
  - [x] SubTask 1.3: 继承 `phase03 ~ phase06` 的相关验收入口，明确 Decision、Product / Repository Binding、Dashboard / Feedback、Onboarding、Reuse Summary 是本阶段最小反回归范围

- [x] Task 2: 冻结 `phase08-11` 的环境准备、工具链验证与 API smoke 顺序
  - [x] SubTask 2.1: 明确前端、后端、数据库与 `/api` 主线的统一前置检查顺序
  - [x] SubTask 2.2: 明确 `buf build / gen / lint`、`go build ./...`、`npx tsc -b --noEmit`、`frontend build` 的执行顺序与通过标准
  - [x] SubTask 2.3: 明确 review 关键 Connect procedure 与 `/api` 访问 smoke 的最小验证矩阵

- [x] Task 3: 冻结 daily / weekly review 两条关键经营路径的浏览器验收矩阵
  - [x] SubTask 3.1: 为 `Dashboard -> Daily Review -> Decision -> Update` 定义最小浏览器步骤、成功回流与返回链检查
  - [x] SubTask 3.2: 为 `Dashboard -> Weekly Review -> Decision -> Update` 定义独立浏览器步骤、成功回流与返回链检查
  - [x] SubTask 3.3: 明确两条路径必须分别给出通过结论，禁止一条路径替代另一条

- [x] Task 4: 冻结 Weekly Review 对 `phase05 / phase06` 读模型的正式消费验证
  - [x] SubTask 4.1: 明确 `overview / recent activity / representative signals` 的验收口径
  - [x] SubTask 4.2: 明确 `reuse snapshot / module_reuse_summary / capability_summary` 的验收口径
  - [x] SubTask 4.3: 明确局部失败边界、错误语义与页面稳定性的检查方式

- [x] Task 5: 冻结 `phase03 ~ phase06` 相关页面的最小反回归矩阵
  - [x] SubTask 5.1: 明确 Decision、Module、Product、Repository 关键页面的最小反回归路径
  - [x] SubTask 5.2: 明确 Dashboard、Onboarding、Reuse Summary 的最小反回归路径
  - [x] SubTask 5.3: 明确“API 成功但 UI 崩溃”“来源链丢失”“owner 越界”三类阻断问题的判定规则

- [x] Task 6: 冻结本阶段边界证据与正式验收结论结构
  - [x] SubTask 6.1: 明确 `Template Reuse / Derived Intelligence / dry-run` 继续不做的记录方式
  - [x] SubTask 6.2: 明确正式验收结论必须包含环境、步骤、结果、问题、复测与 DoD 判定
  - [x] SubTask 6.3: 明确 `phase08-11` 通过后才能进入 `phase08-12` 根级同步

- [x] Task 7: 完成 `phase08-11` 规格一致性校验
  - [x] SubTask 7.1: 校验本规格与 `phase08-08 / 09 / 10` 的冻结口径一致
  - [x] SubTask 7.2: 校验本规格未把未来能力写成当前阶段交付
  - [x] SubTask 7.3: 校验本规格未提前混入 `phase08-12` 的根级同步正文，只保留进入条件

# Task Dependencies

- `Task 2` depends on `Task 1`
- `Task 3` depends on `Task 1`, `Task 2`
- `Task 4` depends on `Task 1`, `Task 3`
- `Task 5` depends on `Task 1`, `Task 2`
- `Task 6` depends on `Task 2`, `Task 3`, `Task 4`, `Task 5`
- `Task 7` depends on `Task 1` through `Task 6`

# 实现总结

## 修复项

- `frontend/src/features/review/data/review-query-options.ts`：将 TanStack Query 提供的 `signal` 透传给 `reviewClient.getDailyReviewContext()` 与 `getWeeklyReviewContext()`，让取消语义按正式 transport 链路传递
- `frontend/src/features/review/data/use-daily-review-read.ts`：新增 `isCanceledReadError()`，把 `Code.Canceled`、`AbortError`、`ERR_ABORTED`、`canceled/cancelled` 识别为取消类错误，并将整页失败条件收紧为“首次加载失败且不是取消类错误”

## 验证结果

- 工具链验证通过：
  - `(cd proto && make build && make gen && make lint)` ✅
  - `(cd backend && go build ./...)` ✅
  - `(cd frontend && npx tsc -b --noEmit)` ✅
  - `(cd frontend && npm run build)` ✅
- 环境与 API smoke 通过：
  - 复用用户已开启的 `http://127.0.0.1:5173` 与 `http://127.0.0.1:8081`
  - `GET /healthz` ✅
  - `GetDailyReviewContext / GetWeeklyReviewContext / SubmitReviewResult` ✅
  - Decision / Module / Product / Repository 关键读取与 `/api` 主线 smoke ✅
- 浏览器验收通过：
  - Daily 路径 ✅：`Dashboard -> Daily Review -> Decision -> Module -> Dashboard`
  - Weekly 路径 ✅：关键区块、`记录决策`、`决策中心`、`完成 Review` 与返回链均成立
- 最小反回归通过：
  - Decision / Module / Product / Repository / Onboarding / Reuse Summary 嵌入式消费 ✅
- 非阻断说明：
  - `GetDailyReviewContext` 与 `GetWeeklyReviewContext` 在浏览器中仍可见偶发 `net::ERR_ABORTED`，但 UI 已不再退化为整页失败，后续成功请求可恢复可用状态
