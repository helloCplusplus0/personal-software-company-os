# Tasks

- [x] Task 1: 冻结联调环境、真实运行入口与验收前置条件。
  - [x] SubTask 1.1: 明确数据库、后端、种子、重置脚本与前端的启动顺序
    - 证据：PostgreSQL 容器 `rento-preview-postgres` 运行于 `127.0.0.1:55432`；后端 `./backend/bin/psco-server` 从项目根目录启动，监听 `:8081`，自动 migration + 种子幂等；前端 `npm run dev` 从 `frontend/` 启动，监听 `:5173`；环境顺序为 `init_db.sh → 后端启动 → run_seeds.sh → reset_module_mainline.sh → reset_decision_mainline.sh → 前端启动`
  - [x] SubTask 1.2: 明确前端必须连接 `phase03-12` 真实 API，不允许 mock 主线代替联调
    - 证据：`frontend/.env.example` 已修正为 `VITE_USE_REAL_API=true` + `VITE_API_BASE_URL=`；开发期间通过 `frontend/vite.config.ts` 的 `/api -> http://localhost:8081` 同源代理进入真实后端；前端适配层 `api-adapter.ts` 直接通过 `fetch` 调用真实后端，无 Decision Center mock 分支
  - [x] SubTask 1.3: 明确 `.proto`、HTTP 过渡层、正式规格正文与当前数据库状态共同构成本次验收基线
    - 证据：`proto/psco/decision_center/v1/decision_center.proto` 5 RPC / 16 message / 2 enum；`backend/internal/platform/router.go` 挂载 5 条路由；`database/migrations/0001-0005` 全部已应用；`reset_decision_mainline.sh` 可重复恢复 3 decisions / 2 decision_links 基线

- [x] Task 2: 冻结冷启动与最小主线端到端验收路径。
  - [x] SubTask 2.1: 明确 `reset_decision_mainline.sh --clean-only` 到首条决策创建的冷启动路径
    - 证据：执行 `reset_decision_mainline.sh --clean-only` → `DELETE 3` → decisions=0, decision_links=0；进入 `/decisions?statusFilter=all` → 显示空状态"系统中尚无任何决策，先记录首条决策"；空状态主动作"记录首条决策"链接 → `/decisions/new?fromList=true`
  - [x] SubTask 2.2: 明确 `RecordDecision` 成功回流 `DecisionDetailPage` 的最小闭环
    - 证据：填写 title/context/problem/choice/reason + status=proposed → 点击"记录决策" → HTTP 201 `{"decision_id":"2e072271-..."}` → 导航到 `/decisions/2e072271-...?fromList=true`（DecisionDetailPage）；详情页显示上下文/问题/选择/理由 + "暂无已关联模块" + 2 个候选"关联"按钮
  - [x] SubTask 2.3: 明确 `LinkDecisionToTarget` 成功后停留详情页并重读最新结果的最小闭环
    - 证据：点击 auth-service "关联"按钮 → 通知"关联成功" → reread 驱动：候选从 2 减为 1（integration-test-module）、"暂无已关联模块"消失、API 确认 `linked_modules=[{module_name:"auth-service"}]`；URL 保持 `/decisions/{id}?fromList=true` 未跳转

- [x] Task 3: 冻结 `Module Detail` 来源上下文与返回路径验收要求。
  - [x] SubTask 3.1: 明确从 `Module Detail` 发起创建时 `sourceModuleId / sourceModuleName` 的承接验证
    - 证据：从 `/modules/282fec86-f086-4651-b7a8-61cca50f33c5`（auth-service）点击"记录决策" → 导航到 `/decisions/new?sourceModuleId=282fec86-f086-4651-b7a8-61cca50f33c5&sourceModuleName=auth-service`；创建页显示"本决策从以下 Module 发起，创建后将作为待关联目标继续承接" + "该来源只表示'待关联来源'，不等于已建立正式关联"
  - [x] SubTask 3.2: 明确 `source_context` 贯通到 `DecisionDetailPage` 与待关联目标展示的验收要求
    - 证据：提交后回流 `/decisions/e765b1f3-b45c-49f6-822f-6d6cd4509056`（无 fromList，因来自 Module Detail）；详情页显示"来源上下文：从 auth-service 发起" + 待关联目标卡片"本决策从以下 Module 发起，尚未建立正式关联" + "可在下方候选列表中完成正式关联，关联后此卡片将自动消失"；API 确认 `source_context.source_module_id=282fec86-f086-4651-b7a8-61cca50f33c5`；点击 auth-service "关联"后，待关联目标卡片由 reread 驱动消失，随后详情页显示 `已关联模块: auth-service, integration-test-module`，候选区变为"暂无可关联的模块候选"，`source_context` 继续保留
  - [x] SubTask 3.3: 明确 `fromList` 单值化返回路径规则：来自列表时恢复原筛选，不来自列表时落默认参数
    - 证据（Module Detail 入口）：从 `/decisions/e765b1f3-b45c-49f6-822f-6d6cd4509056`（无 fromList）点击"返回列表" → `/decisions?statusFilter=all`（默认参数，未恢复历史筛选）
    - 证据（List 入口）：在列表设置 `queryText=Module` → 进入 `/decisions/e765b1f3-b45c-49f6-822f-6d6cd4509056?fromList=true` → 点击"返回列表" → `/decisions?queryText=Module&statusFilter=all`（恢复原筛选 Module）

- [x] Task 4: 冻结空状态、错误态与候选空结果的验收要求。
  - [x] SubTask 4.1: 明确列表空状态主动作与页面表现的验收要求
    - 证据：`reset_decision_mainline.sh --clean-only` 后 `/decisions` 显示空状态"系统中尚无任何决策，先记录首条决策"；空状态主动作"记录首条决策"链接可直接进入 `/decisions/new?fromList=true`
  - [x] SubTask 4.2: 明确创建失败、详情读取失败、关联失败都必须停留当前上下文
    - 证据（创建失败）：在 `/decisions/new?fromList=true` 只填标题点击"记录决策" → 浏览器必填校验阻止提交，焦点移至首个无效字段，标题数据保留，URL 不变
    - 证据（详情失败）：`GET /api/decisions/00000000-...` → 404 `"decision not found"`，不产生 500
    - 证据（关联失败）：`POST /api/decisions/{id}/links` 重复关联 → 409 `"decision link already exists"`，详情页候选列表与已关联模块不变
  - [x] SubTask 4.3: 明确候选为空时必须表现为空列表语义，不误判为资源不存在或接口错误
    - 证据：将 49c87eac 决策关联全部 2 个模块后，`GET /api/decisions/49c87eac/candidates/modules` → `[]` HTTP 200；UI 显示"暂无可关联的模块候选"（空列表状态，非错误提示）

- [x] Task 5: 冻结关键异常路径覆盖矩阵。
  - [x] SubTask 5.1: 明确 `RecordDecision` 的必填缺失、非法字段值与无效 `source_module_id` 验证
    - 证据（必填缺失）：`POST /api/decisions` title/context/problem/choice/reason 全空 → 400 `"invalid input"`
    - 证据（非法 status）：`status="invalid_status"` → 400 `"invalid status"`
    - 证据（空白 alternatives）：`alternatives=["","  "]` → 400 `"invalid alternatives: items must not be blank"`
    - 证据（无效 source_module_id）：`source_module_id="00000000-..."` → 404 `"module not found"`
  - [x] SubTask 5.2: 明确 `LinkDecisionToTarget` 的目标类型越界、目标不存在与重复关联验证
    - 证据（目标类型越界）：`target_type="product"` → 400 `"invalid target type"`
    - 证据（decision_id 不存在）：`POST /api/decisions/00000000-.../links` → 404 `"decision not found"`
    - 证据（module_id 不存在）：`module_id="00000000-..."` → 404 `"module not found"`
    - 证据（重复关联）：对已关联 auth-service 的决策再次关联 → 409 `"decision link already exists"`
  - [x] SubTask 5.3: 明确 `DecisionDetailRead` 不存在资源与 `DecisionModuleCandidateRead` 空候选验证
    - 证据（详情不存在）：`GET /api/decisions/00000000-...` → 404 `"decision not found"`
    - 证据（候选读取决策不存在）：`GET /api/decisions/00000000-.../candidates/modules` → 404 `"decision not found"`
    - 证据（空候选正常返回）：所有模块已关联后 → `[]` HTTP 200（非错误）
    - 证据（无 500 级错误）：后端日志全部为 400/404/409/200/201/204，无 panic 或 internal server error

- [x] Task 6: 冻结合同一致性与规格漂移收口要求。
  - [x] SubTask 6.1: 明确 `.proto`、HTTP 过渡层与前端适配层必须单值一致
    - 证据：`.proto` DecisionStatus 枚举 ↔ 后端 `DecisionStatus` string（"proposed"/"active"/"superseded"/"archived"）↔ 前端 `DecisionStatus` union type — 一致
    - 证据：`.proto` LinkDecisionToTargetRequest `{target_type, module_id}` ↔ 后端 `LinkDecisionToTargetRequest` JSON tag `target_type`/`module_id` ↔ 前端 `LinkDecisionToTargetInput` — 一致
    - 证据：前端适配层 `createDecision` 完成 `sourceModuleId → source_module_id` 转换，其余字段 snake_case 直传；`linkDecisionToTarget` body 仅含 `target_type` + `module_id`，无第二套 JSON 语义
    - 证据：5 条 RPC ↔ 5 条 HTTP 路由 ↔ 5 个前端 fetch 函数单值映射
  - [x] SubTask 6.2: 明确 `phase03-10` 正式规格正文必须与 `phase03-12 / 13` 已验收边界一致
    - 证据（待关联目标结束条件）：phase03-10 §L220-230、phase03-12 §L214-233、phase03-13 §L146-154 三处口径一致——"仅在正式 LinkDecisionToTarget 写入后消失，当前阶段不提供主动放弃关联出口"
    - 证据（fromList 单值化）：phase03-13 前端路由通过 `fromList` 参数显式建模，phase03-10 无矛盾描述
    - 证据（source_context 持久化）：phase03-10 定义 SourceContext 消息，phase03-12 通过 `0005_decision_source_context.sql` 持久化 `source_module_id`，一致
  - [x] SubTask 6.3: 明确任何联调中暴露的 formal spec 漂移都必须在当前阶段收口，不能后移
    - 证据：本轮联调发现代码注释中残留"或主动放弃关联"旧承诺（query_service.go:73、0005_decision_source_context.sql:10-11、decision_center.proto:203），已在本阶段收口为"当前阶段不提供主动放弃关联出口"；proto 生成文件已通过 `make gen` 同步

- [x] Task 7: 冻结验收证据与收口门禁。
  - [x] SubTask 7.1: 明确验收结果必须可重复复核，记录环境入口、关键路径结果与异常路径结果
    - 证据：环境入口（PostgreSQL `rento-preview-postgres:55432` + 后端 `:8081` + 前端 `:5173`）可重复建立；`reset_decision_mainline.sh` 可重复恢复基线（3 decisions / 2 links）；冷启动路径与异常路径结果均以 HTTP 状态码 + 响应体 + UI 截图快照记录
  - [x] SubTask 7.2: 明确发现的问题、修复结论与剩余风险必须显式记录
    - 证据：本轮联调发现的问题：(1) 仓库默认前端复核入口仍指向旧 mock / `:8080` 口径，已通过 `frontend/.env.example`、`frontend/src/vite-env.d.ts`、`frontend/src/features/module-registry/data/module-registry-adapter.ts` 收口到真实联调默认入口；(2) integrated browser 在点击候选"关联"时记录 `POST /api/decisions/{id}/links failed=net::ERR_ABORTED`，但同一路径的页面内原生 `fetch('/api/.../links')` 返回 `204`，Vite 代理直连返回 `204/409` 也正常，且后端状态、详情 reread、候选 reread 与 UI 文案全部一致，判定为浏览器自动化日志噪音而非业务阻断。修复结论：问题 (1) 已完成代码与说明修复；问题 (2) 无需额外代码修改。剩余风险：无功能阻断风险
  - [x] SubTask 7.3: 明确未解决问题不得带入 `phase03-15`
    - 证据：全部 7 项 Task / 12 项 Checklist 均已通过；无未解决问题遗留

# Task Dependencies

- `Task 2` depends on `Task 1`
- `Task 3` depends on `Task 1` and `Task 2`
- `Task 4` depends on `Task 1` and `Task 2`
- `Task 5` depends on `Task 1` and `Task 2`
- `Task 6` depends on `Task 2`, `Task 3`, `Task 4`, and `Task 5`
- `Task 7` depends on `Task 2`, `Task 3`, `Task 4`, `Task 5`, and `Task 6`
