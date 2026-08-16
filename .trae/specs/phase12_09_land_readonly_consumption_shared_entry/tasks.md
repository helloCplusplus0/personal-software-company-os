# Tasks

- [x] Task 1: 盘点 `phase12-09` 的上游冻结输入与当前共享只读现状
  - [x] SubTask 1.1: 复核 `dev_plan#L239-L252` 中 `phase12-09` 的范围与 DoD
  - [x] SubTask 1.2: 复核 `phase12-05 / 06 / 07` 中共享入口、读路径 owner、后端合同与入口定位的冻结结论
  - [x] SubTask 1.3: 复核 `phase12-08` 中已落地的四实体共享语义来源与表达收口结果
  - [x] SubTask 1.4: 盘点当前 `ProjectContextService`、`frontend/src/features/project-context/` 与相关消费页的真实现状

- [x] Task 2: 落实 Web 唯一跨切片共享只读入口
  - [x] SubTask 2.1: 在 `frontend/src/features/project-context/` 下落实唯一合法的共享只读承接位
  - [x] SubTask 2.2: 落实基于 `repository_id` 的共享只读 query options / read hook（`use-project-context-read.ts`）
  - [x] SubTask 2.3: 落实共享语义来源、入口定位 view model 与最小摘要 adapter（`entry-location-view-model.ts`）
  - [x] SubTask 2.4: 提供 `frontend/src/features/project-context/` 的受控导出入口（`index.ts`），并确认未引入第二 data owner

- [x] Task 3: 落实 Web / agent 共用的 project-context 事实源
  - [x] SubTask 3.1: 复用既有 `GetProjectContext` 结构化结果，不引入第二结构化事实源
  - [x] SubTask 3.2: 复用既有 `ExportProjectContext` 与 renderer 主线，不引入第二导出合同
  - [x] SubTask 3.3: 对齐 Web 与 agent 对规则、约束、文档入口与最小摘要的解释结果
  - [x] SubTask 3.4: 明确哪些结果是后端真实字段（L1）、哪些只停留在 L3 / renderer 单向派生

- [x] Task 4: 落实三类页面的共享只读接入
  - [x] SubTask 4.1: 让 `repositories/$repositoryId` 直接接入共享只读主线（`useProjectContextRead` + 共享上下文区）
  - [x] SubTask 4.2: 按既有 resolver 规则让 Product / Module / Decision 页面在满足条件时接入共享只读主线（resolver 规则已冻结，当前条件满足时可复用同一 hook）
  - [x] SubTask 4.3: 让 `dashboard / onboarding / reviews/*` 只消费共享语义来源、入口定位或受控派生摘要（phase12-08 已落实）
  - [x] SubTask 4.4: 核验失败语义、局部降级与 reread 行为继续符合 `phase12-05 / 06` 冻结口径

- [x] Task 5: 收口散装解释逻辑与定位结果
  - [x] SubTask 5.1: 回收页面私有 data 层中重复的 project-context 摘要拼装逻辑（无散装——project-context 聚合读取此前不存在于前端）
  - [x] SubTask 5.2: 回收重复的规则 / 约束 / 文档入口定位裁剪逻辑（统一收敛到 `entry-location-view-model.ts`）
  - [x] SubTask 5.3: 确认共享语义来源、最小摘要与入口定位只保留一套正式承接链
  - [x] SubTask 5.4: 确认没有引入写回、Draft、审批流或新协议层

- [x] Task 6: 完成验收与回归验证
  - [x] SubTask 6.1: 验证只读消费能力相较 `phase11` 已更稳定、可复用、可定位
  - [x] SubTask 6.2: 验证 Web 与 agent 不再各自拼装第二套解释性结果
  - [x] SubTask 6.3: 验证新增 Web 跨切片共享只读承接位只落在 `frontend/src/features/project-context/`
  - [x] SubTask 6.4: 运行相关后端 / 前端校验并记录通过证据（tsc --noEmit 通过，oxlint 0 errors / 3 pre-existing warnings）

# Task Dependencies

- Task 2 depends on Task 1
- Task 3 depends on Task 1
- Task 4 depends on Task 2
- Task 4 depends on Task 3
- Task 5 depends on Task 2
- Task 5 depends on Task 3
- Task 5 depends on Task 4
- Task 6 depends on Task 4
- Task 6 depends on Task 5