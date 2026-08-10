# Tasks

- [x] Task 1: 对齐 `phase06-12` 的直接上游与 formal spec 职责边界，明确这次任务是“收敛正式规格正文”，不是进入实现或合同主线落地。
  - [x] SubTask 1.1: 对齐 `docs/phase/phase06_onboarding_sovereignty_reuse_foundation_dev_plan.md` 中 `phase06-12` 的范围与 DoD
  - [x] SubTask 1.2: 对齐 `phase06` shared baseline / architecture plan 中的技术路线、页面矩阵、动作矩阵与非目标
  - [x] SubTask 1.3: 对齐 `phase06-01 ~ 11` 的冻结结论，确认不再形成并列直接执行层入口
  - [x] SubTask 1.4: 对齐 `phase01-06 / phase05-10` 的 formal spec 收敛模式，复用正文定位而不重造格式

- [x] Task 2: 冻结 `phase06` 正式规格正文的角色与上游收口关系。
  - [x] SubTask 2.1: 冻结 `phase06-12` 为 `phase06` 唯一直接执行层规格入口
  - [x] SubTask 2.2: 冻结 `phase06-01 ~ 11` 退为追溯来源与证据链
  - [x] SubTask 2.3: 冻结后续 `phase06-13` 及实现只能承接本文，不再并列消费子规格

- [x] Task 3: 冻结页面、路由、`first_run_state` 与 Onboarding 交互主线。
  - [x] SubTask 3.1: 冻结 `Onboarding / Dashboard / Export / Backup / Detail` 页面矩阵与正式路由
  - [x] SubTask 3.2: 冻结 `not_started / in_progress / completed` 与根级默认进入路径
  - [x] SubTask 3.3: 冻结 `Product -> Repository -> Module -> Decision` 的首轮录入主线与完成条件
  - [x] SubTask 3.4: 冻结 `Continue Onboarding`、回访继续与完成回流语义

- [x] Task 4: 冻结 draft-first 写模型、前端写路径 owner 与 canonical create 合同复用关系。
  - [x] SubTask 4.1: 冻结四类最小人工必填字段与 partial-entry 语义
  - [x] SubTask 4.2: 冻结继续复用 `CreateProduct / CreateRepository / CreateModule / CreateDecision`
  - [x] SubTask 4.3: 冻结 `application` 单入口、`query` 纯只读与 mutation 固定承接位
  - [x] SubTask 4.4: 冻结 `Onboarding` 与既有 create 页面共享同一套写入语义

- [x] Task 5: 冻结 Export / Backup / Reuse Summary 的正式业务边界。
  - [x] SubTask 5.1: 冻结 Export 的语义、覆盖矩阵、入口位与成功/失败语义
  - [x] SubTask 5.2: 冻结 Backup 的语义、恢复前提、`backup verified` 与错误边界
  - [x] SubTask 5.3: 冻结 `module_reuse_summary / capability_summary` 的事实来源、字段与页面挂接位
  - [x] SubTask 5.4: 冻结复用感知的新鲜度、空态与失败态边界

- [x] Task 6: 冻结 `phase06` 的合同、传输与演进基线。
  - [x] SubTask 6.1: 冻结 `.proto` 为唯一长期合同源
  - [x] SubTask 6.2: 冻结 `Onboarding / Export / Backup / Reuse Summary` 的 package、service 与 HTTP 映射矩阵
  - [x] SubTask 6.3: 冻结 DTO 单向派生与 `BackupWrite.read_verify` 读取侧合同一致性
  - [x] SubTask 6.4: 冻结 enum 零值、`reserved` 与 breaking 演进规则

- [x] Task 7: 冻结联调验收基线、非目标与 Done 标准。
  - [x] SubTask 7.1: 冻结 `reset_phase06_acceptance.sh` 与 fixture 白名单
  - [x] SubTask 7.2: 冻结首轮成功会话、阶段完成与合同一致性验收矩阵
  - [x] SubTask 7.3: 冻结不得依赖手工补数据的验收门禁
  - [x] SubTask 7.4: 冻结 `phase06` 非目标矩阵与 Done 标准

- [x] Task 8: 完成 formal spec 一致性校验与后续承接检查。
  - [x] SubTask 8.1: 验证正文完整覆盖页面、交互、写路径、导出 / 备份、复用读模型、合同、验收基线、非目标与 Done 标准
  - [x] SubTask 8.2: 验证正文与 `phase06-01 ~ 11` 单值一致，不反向改写上游结论
  - [x] SubTask 8.3: 验证正文足以直接作为 `phase06-13`、实现与下一阶段的唯一上游规格来源

# Task Dependencies

- `Task 2` depends on `Task 1`
- `Task 3` depends on `Task 1`, `Task 2`
- `Task 4` depends on `Task 1`, `Task 2`, `Task 3`
- `Task 5` depends on `Task 1`, `Task 2`, `Task 3`
- `Task 6` depends on `Task 1`, `Task 2`, `Task 4`, `Task 5`
- `Task 7` depends on `Task 1`, `Task 2`, `Task 3`, `Task 5`, `Task 6`
- `Task 8` depends on `Task 1` through `Task 7`
