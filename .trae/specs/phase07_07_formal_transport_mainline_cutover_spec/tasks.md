# Tasks

- [x] Task 1: 对齐 phase07-07 的 formal spec 角色与直接上游边界，明确本次任务是"收敛单一正式规格正文"，不是再次产出并列设计文档。
  - [x] SubTask 1.1: 继承 phase07_transport_contract_mainline_migration_dev_plan.md#L157-166 的范围与 DoD，确认 formal spec 必须覆盖迁移范围、生成链、后端、前端、验收与退场标准 → `transport_mainline_cutover_spec_v0.1.md` §1-6 全部覆盖
  - [x] SubTask 1.2: 对齐 phase07-01 ~ 06 的职责边界，明确它们在 formal spec 生效后退为冻结来源与证据链 → `transport_mainline_cutover_spec_v0.1.md` §1.2 继承链图 + §1.3 禁止项
  - [x] SubTask 1.3: 冻结正式正文文件落点与作为 phase07-08 ~ 11 唯一直接上游的角色 → `transport_mainline_cutover_spec_v0.1.md` 文档头 + §1.1

- [x] Task 2: 收敛 formal spec 的迁移范围、合同与工具链章节。
  - [x] SubTask 2.1: 将 9 个 canonical 业务模块、34 条 canonical RPC、4 条 legacy / compat endpoint 与 infra keep list 收入统一正文 → `transport_mainline_cutover_spec_v0.1.md` §2.1-2.4
  - [x] SubTask 2.2: 将 .proto 唯一合同源、单一 /api 前缀、ConnectRPC 正式传输主线与 compat 只作过渡资产的规则收入正文 → `transport_mainline_cutover_spec_v0.1.md` §3.1-3.2
  - [x] SubTask 2.3: 将 buf 三插件矩阵、proto/Makefile、Go / TS 产物落点、前端 Connect runtime 依赖与当前 CI 缺口收入正文 → `transport_mainline_cutover_spec_v0.1.md` §3.3-3.5

- [x] Task 3: 收敛 formal spec 的后端正式主线章节。
  - [x] SubTask 3.1: 将 chi 的唯一正式职责、Connect handler 消费 generated path 的挂载方式与 compat route 的退场规则收入正文 → `transport_mainline_cutover_spec_v0.1.md` §4.1-4.2 + §4.6
  - [x] SubTask 3.2: 将 service 分层保持、backend/internal/*/connect/ 落点、domain error → Connect error 映射收入正文 → `transport_mainline_cutover_spec_v0.1.md` §4.3-4.5
  - [x] SubTask 3.3: 明确 formal spec 不允许重新长出第二套 canonical API 或第二套路由/传输主线 → `transport_mainline_cutover_spec_v0.1.md` §4.4 禁止项

- [x] Task 4: 收敛 formal spec 的前端正式主线章节。
  - [x] SubTask 4.1: 将 shared/rpc/connect-transport.ts、slice-local connect-client.ts、query owner、application owner 的固定承接位收入正文 → `transport_mainline_cutover_spec_v0.1.md` §5.1-5.4
  - [x] SubTask 4.2: 将 11 项 mutation owner、Onboarding route caller、candidate read 与 SovereigntyPanel 过渡位收入正文 → `transport_mainline_cutover_spec_v0.1.md` §5.5
  - [x] SubTask 4.3: 明确旧 api-adapter.ts、route/page/component 直连 transport/client 的禁止事项与退场窗口 → `transport_mainline_cutover_spec_v0.1.md` §5.6-5.7

- [x] Task 5: 收敛 formal spec 的验收、退场与收口章节。
  - [x] SubTask 5.1: 将 phase07-06 的 34 条 RPC 迁移矩阵、跨模块回归清单与 route 级回归要求收入正文 → `transport_mainline_cutover_spec_v0.1.md` §6.1-6.2
  - [x] SubTask 5.2: 将 4 条 legacy / compat endpoint 的 endpoint 级删除证据、替代 Connect 回归证据与最晚退场时点收入正文 → `transport_mainline_cutover_spec_v0.1.md` §6.3
  - [x] SubTask 5.3: 将 11 项 frontend mutation owner 验收映射、最终 evidence package 与 phase 收口阻断条件收入正文 → `transport_mainline_cutover_spec_v0.1.md` §6.4-6.6

- [x] Task 6: 完成 formal spec 一致性校验，确保后续阶段只能从单一入口推进。
  - [x] SubTask 6.1: 校验 formal spec 与 phase07-01 ~ 06 的修正后结论一致 → `transport_mainline_cutover_spec_v0.1.md` §8 一致性声明
  - [x] SubTask 6.2: 校验 formal spec 与根级真相源、phase07 architecture/shared baseline/dev plan 以及 project_rules.md 一致 → `transport_mainline_cutover_spec_v0.1.md` §8 一致性声明
  - [x] SubTask 6.3: 校验 formal spec 已明确 phase07-08 ~ 11 不得再并列引用 phase07-01 ~ 06 作为长期直接执行入口 → `transport_mainline_cutover_spec_v0.1.md` §1.3 禁止项

# Task Dependencies

- Task 2 depends on Task 1 ✅
- Task 3 depends on Task 1, Task 2 ✅
- Task 4 depends on Task 1, Task 2 ✅
- Task 5 depends on Task 2, Task 3, Task 4 ✅
- Task 6 depends on Task 2, Task 3, Task 4, Task 5 ✅