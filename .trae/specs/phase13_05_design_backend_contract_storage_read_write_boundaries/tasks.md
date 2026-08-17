# Tasks

- [x] Task 1: 冻结项目治理画像第一版后端唯一承接位
  - [x] SubTask 1.1: 明确后端唯一正式承接位为同一 `repository_id` 驱动的 repository-scoped governance aggregate
  - [x] SubTask 1.2: 明确该承接位不是第五个业务主实体，也不是四实体字段拼装
  - [x] SubTask 1.3: 复核第一版不会并列长出第二个治理画像后端入口

- [x] Task 2: 冻结第一版业务合同链路与传输边界
  - [x] SubTask 2.1: 明确 `.proto` 是治理画像新增业务接口的唯一长期合同源
  - [x] SubTask 2.2: 明确 `ConnectRPC` 是默认正式业务传输层
  - [x] SubTask 2.3: 明确 `chi + net/http` 不得形成并列 canonical API

- [x] Task 3: 冻结存储分层与读写边界
  - [x] SubTask 3.1: 明确 `governance_profile_record / canonical_root_file_bindings[] / global_asset_bindings[]` 的第一版分层
  - [x] SubTask 3.2: 明确哪些结构化数据允许持久化，以及顶层目录矩阵的后端承接结论
  - [x] SubTask 3.3: 明确 markdown 正文只允许回源、不允许全文入库
  - [x] SubTask 3.4: 明确 `read-only` 治理字段不会进入写白名单
  - [x] SubTask 3.5: 复核第一版写路径仍坚持“手工维护优先、自动同步后置”

- [x] Task 4: 冻结与四实体主线的后端关系
  - [x] SubTask 4.1: 明确治理画像与四实体既有合同、读路径、写路径的关系
  - [x] SubTask 4.2: 明确治理画像不得复制第二套四实体事实源
  - [x] SubTask 4.3: 复核 `Decision` 不会被误用为全局规范资产正式承接位

- [x] Task 5: 冻结全局规范资产逐项承接矩阵
  - [x] SubTask 5.1: 明确前 5 份文件的承接策略为 `entry_ref + role + structured_summary`
  - [x] SubTask 5.2: 明确后 3 份文件的承接策略为 `entry_ref + role`
  - [x] SubTask 5.3: 明确 8 份文件统一允许 markdown 正文回源但第一版禁止全文入库
  - [x] SubTask 5.4: 复核后续执行者不再需要猜“哪份文件只做引用、哪份文件需要结构化摘要”

- [x] Task 6: 完成 spec 包与上游冻结文档的一致性校验
  - [x] SubTask 6.1: 校验本 spec 包与 `phase13-03 / phase13-04` 在锚点、字段与分层边界上保持单值一致
  - [x] SubTask 6.2: 校验本 spec 包没有引入与 `project_rules.md` 的 `.proto / ConnectRPC` 约束相冲突的新解释
  - [x] SubTask 6.3: 校验本 spec 包为 `phase13-06 / phase13-07` 提供了稳定后端前提

# Task Dependencies

- `Task 2` depends on `Task 1`
- `Task 3` depends on `Task 1`, `Task 2`
- `Task 4` depends on `Task 1`, `Task 2`, `Task 3`
- `Task 5` depends on `Task 1`, `Task 3`
- `Task 6` depends on `Task 1`, `Task 2`, `Task 3`, `Task 4`, and `Task 5`

# 执行记录

- SubTask 1.1 结论：承接位冻结为同一 `repository_id` 驱动的 repository-scoped governance aggregate，与 phase13-03 冻结的“`repository_id` 是第一版唯一正式锚点”及 `shared_baseline` §3.6 单值一致；该承接位同时承接 04 冻结的 9 类字段、全局规范资产结构化承接结果，并作为 Web 与 agent 共享治理事实的唯一后端来源
- SubTask 1.2 结论：spec Requirement 1 补充冻结“不是第五个业务主实体、不得实现为四实体任一实体字段的临时拼装结果”，与 phase13-03 Requirement 1/`architecture_plan` §4.4 逐句对应
- SubTask 1.3 结论：spec Requirement 1 禁止并列长出第二个“项目治理画像后端入口”或第二个“治理摘要表述协议”；Scenario 将“把治理画像拆散塞进四实体既有后端合同”显式判定为越界
- SubTask 2.1 结论：`.proto` 为唯一长期合同源，对外字段、枚举、响应 envelope 与错误语义以 `.proto` 为准——与 `project_rules.md` §2.6 第 1 条逐字一致
- SubTask 2.2 结论：`ConnectRPC` 为默认正式业务传输层——与 `project_rules.md` §2.6 第 3 条“业务相关接口默认优先采用 ConnectRPC”一致；spec 明确第一版不得新增手写 `chi + JSON HTTP` 业务接口作为正式合同
- SubTask 2.3 结论：`chi + net/http` 只承担基础设施端点或兼容适配、不得形成并列 canonical API——与 `project_rules.md` §2.6 第 4/5 条（非业务端点允许 chi 承接、存量业务接口只能是兼容适配层）一致；兼容 handler 的 struct 必须从 `.proto` 单向派生或显式对齐映射（对应 §2.6 第 6 条）
- SubTask 3.1 结论：存储分层冻结为三层——`governance_profile_record`（9 类字段）、`canonical_root_file_bindings[]`（file_name/role/required）、`global_asset_bindings[]`（name/kind/entry_ref/role + 可选 structured_summary）；`global_asset_bindings[]` 显式定位为 phase13-04 `global_constraint_refs[]` 的后端承接扩展位、非新增核心字段集合，与 04 REMOVED Requirement（禁止临场补字段）一致
- SubTask 3.2 结论：持久化白名单冻结为“治理画像结构化字段、资产入口关系、角色与必要摘要”；并已在 spec 正文显式冻结顶层目录矩阵的后端承接结论：`docs/` 通过 `docs_workflow_layout` 进入正式字段，`backend / database / frontend / proto` 作为当前项目范式 v1 的只读基线输入保留，不新增 repository-scoped 可写持久化字段，也不再保留为待澄清事项
- SubTask 3.3 结论：markdown 正文不得作为 canonical 存储副本写入数据库，第一版不得建立全文索引、全文副本表或并列缓存真相源；正文回源属 read-time resolution，失败语义独立于结构化字段读取成功与否——与 `shared_baseline` §3.5“markdown 正文继续保留在仓库内作为正文载体”、“统一允许回源但不全文入库”及 `architecture_plan` §4.6“结构化摘要 + markdown 正文回源”双层模式一致
- SubTask 3.4 结论：`track_type / current_phase_name / current_phase_ref / current_phase_status` 虽属于治理画像结构化字段，但因 `phase13-04` 已冻结为 `read-only`，本 spec 已将其显式排除出治理画像维护写路径的可写集合，只允许来自根级正式上游的冻结结果回读
- SubTask 3.5 结论：写路径白名单冻结为可写的结构化治理字段与结构化资产承接结果；禁止保存时顺带自动扫描仓库、自动同步模板、回填目录全文——“手工维护优先、自动同步后置”与 `shared_baseline` §3.3 补充冻结、phase13-01 非目标、phase13-02 C 层后置三处口径一致；写路径收敛为单一写入承接位
- SubTask 4.1 结论：四实体既有合同继续承接业务事实与业务关系；治理画像承接位只承接治理画像、全局规范资产与当前阶段状态；四实体读取治理事实只能通过同一 `repository_id` 驱动的治理承接位——与 phase13-03 关系矩阵逐项对应
- SubTask 4.2 结论：禁止反向复制四实体完整事实为第二套存储副本；读接口允许组合返回 `repository_id` 相关四实体摘要引用，但不得升级为治理画像 canonical 业务事实（read-time aggregation 边界）——与 phase13-03 Requirement 4“同一概念不得两套正式语义”一致
- SubTask 4.3 结论：第一版不得把 `Decision` 现有“规则”语义误扩写为全局规范资产的正式存储位——与 phase13-03 Requirement 2 第 4 项、`dev_plan` phase13-03 DoD“不再把目录结构、全局规范与当前阶段信息硬塞进 Decision”一致
- SubTask 5.1 结论：前 5 份文件（project_rules.md / TECH_STACK_BASELINE.md / AGENTS.md / architecture_map.md / plan.md）承接策略冻结为 `entry_ref + role + structured_summary`，与 `shared_baseline` §3.5（当前版）逐项一致
- SubTask 5.2 结论：后 3 份文件（README.md / global_skills.md / project_skills.md）承接策略冻结为 `entry_ref + role`，与 `shared_baseline` §3.5 逐项一致；spec 补充冻结禁止把后 3 项偷偷提升为摘要必做
- SubTask 5.3 结论：8 份文件统一允许 markdown 正文回源、第一版禁止全文入库，与 `shared_baseline` §3.5“第一版统一允许 markdown 正文回源，但不把上述文件全文入库”单值一致
- SubTask 5.4 结论：spec Requirement 7 Scenario 已将“临时把只需 entry_ref + role 的文件提升为摘要必填项”显式判定为越界；8 项矩阵覆盖 `dev_plan` phase13-05 要求的全部四类策略维度（只存引用 / 摘要 / 回源不入库 / 禁全文入库）
- SubTask 6.1 结论：锚点（repository_id）承接 03 冻结、字段（9 类 + 嵌套子字段）承接 04 冻结、存储分层与读写边界为本 spec 新增单值设计且不与 03/04 冲突；`canonical_root_file_bindings[]` 子字段与 04 冻结的 file_name/role/required 逐字一致
- SubTask 6.1 补充：本轮修复后，05 已显式接住 04 的两项关键约束——`read-only` 字段不进入写白名单，以及 `backend / database / frontend / proto` 的第一版后端承接结论不再只留在执行记录中
- SubTask 6.2 结论：spec 合同链路四条与 `project_rules.md` §2.6 约束逐条对齐（.proto 唯一合同源对应第 1 条 / ConnectRPC 默认对应第 3 条 / chi 仅基础设施与兼容适配对应第 4/5 条 / handler struct 单向派生对应第 6 条；§2.6 共 7 条，第 7 条“未来演进传输协议”约束与 05 spec 无冲突亦无对应需求），无冲突新解释
- SubTask 6.3 结论：后端承接位、合同链路、存储分层、读写边界与 8 项资产矩阵均已单值冻结并配套机械判定 Scenario；`phase13-06`（前端承接位）与 `phase13-07`（agent brief）可直接以本 spec 为稳定后端前提——“05 先于 06/07”顺序的正式依据是本 spec MODIFIED Requirement（“MUST 先完成……再进入 phase13-06 / phase13-07”），`dev_plan` §5 相关条文（第 4 条“04 先于 05/06/07”、第 5 条“05/06/07 共享同一套字段模型”）与之兼容
