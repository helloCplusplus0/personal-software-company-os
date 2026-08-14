# Tasks

- [x] Task 1: 收敛 `Decision` 生命周期推进的最小后端承接位
  - [x] SubTask 1.1: 复核并补齐 `UpdateDecisionStatus` 在 `.proto`、Connect、service、repository` 的最小承接，确保四态迁移矩阵与 `phase10-03` 一致
  - [x] SubTask 1.2: 校正 `backend/internal/dashboard/` 与 `backend/internal/review/` 的 pending 读取口径，只允许以 canonical `Decision.status` 解释 `proposed / active / superseded / archived`
  - [x] SubTask 1.3: 为状态推进成功、终态禁止推进与 pending reread 结果补齐最小后端测试或等价验证

- [x] Task 2: 落实 `Decision Detail` 的页面级 read owner、action owner 与 CTA 矩阵
  - [x] SubTask 2.1: 在 `frontend/src/features/decision-center/data/` 收敛单一 `Decision Detail` 页面读承接位，稳定产出状态、CTA 描述符与来源返回上下文
  - [x] SubTask 2.2: 在 `frontend/src/features/decision-center/application/` 收敛单一 `Decision Detail` 页面写承接位，统一承接状态推进、错误归一化、成功 envelope 与页面级失效刷新
  - [x] SubTask 2.3: 改造 `decision-detail-page.tsx`，落地 `proposed / active / superseded / archived` 四态下的真实 CTA 矩阵与来源返回入口

- [x] Task 3: 统一 `Dashboard / Daily Review / Current Focus` 的 pending handoff 主线
  - [x] SubTask 3.1: 让 `Dashboard` 的 feedback signals、`Current Focus` 与主 CTA 在命中 pending decision 时统一 handoff 到 `Decision Detail`
  - [x] SubTask 3.2: 让 `Daily Review` 的 pending decision 入口统一 handoff 到 `Decision Detail`，不再并行承接状态推进语义
  - [x] SubTask 3.3: 移除页面局部"已处理"假象，包括依赖 `decision_links`、`review_records` 或本地 dismiss 的代理退出模式

- [x] Task 4: 落实状态推进成功后的回流、失效刷新与 reread
  - [x] SubTask 4.1: 统一 `Decision Detail` 状态推进成功后的 detail reread 与来源页相关 query 失效范围
  - [x] SubTask 4.2: 确保返回 `Dashboard / Daily Review / Current Focus` 后，pending count、pending card 与主 CTA 都基于 reread 结果同步更新
  - [x] SubTask 4.3: 补齐失败态、重复点击、终态禁用与刷新浏览器后的最小行为约束

- [x] Task 5: 完成 `phase10-09` 浏览器级闭环验收
  - [x] SubTask 5.1: 验证从 `Dashboard / Current Focus` 进入 `Decision Detail` 并推进状态后，返回 `Dashboard` 不再误报 pending
  - [x] SubTask 5.2: 验证从 `Daily Review` 进入 `Decision Detail` 并推进状态后，返回 `Daily Review` 不再残留该 pending 项
  - [x] SubTask 5.3: 验证 `proposed / active / superseded / archived` 四态在 `Decision Detail` 上展示的 CTA 集合与规格一致
  - [x] SubTask 5.4: 验证本子任务未顺带改写 `Product / Module / Repository Detail` 的独立 CTA inventory

- [x] Task 6: 完成 `phase10-09` 交付自检与边界复核
  - [x] SubTask 6.1: 复核实现是否对齐 `phase10-03 / 05 / 06 / 07` 的单值边界，不回退到页面级第二套写路径
  - [x] SubTask 6.2: 复核 `Decision Detail + Dashboard / Daily Review / Current Focus` 是否已经共享同一套 canonical pending 解释
  - [x] SubTask 6.3: 记录仍属于非目标的事项，避免把 `Product / Module / Repository Detail` 的 CTA inventory 混入本子任务

# Task Dependencies

- `Task 2` depends on `Task 1`
- `Task 3` depends on `Task 1`
- `Task 4` depends on `Task 2` and `Task 3`
- `Task 5` depends on `Task 1` to `Task 4`
- `Task 6` depends on `Task 1` to `Task 5`