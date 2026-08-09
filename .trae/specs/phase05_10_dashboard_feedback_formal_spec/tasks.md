# Tasks

- [x] Task 1: 对齐 `phase05-10` 的直接上游与正式正文样式基线，明确这次任务要产出的是“正式规格正文入口”，不是第十份并列子规格。
  - [x] SubTask 1.1: 对齐 `docs/phase/phase05_dashboard_feedback_foundation_dev_plan.md` 中 `phase05-10` 的范围与 DoD
  - [x] SubTask 1.2: 对齐 `phase01-06`、`phase04-10` 的正式正文写法与章节风格
  - [x] SubTask 1.3: 对齐 `phase04-14` 的验收结论，明确 `phase05` 必须承接既有 canonical owner、reread 与返回路径前提

- [x] Task 2: 冻结 `phase05` 正式规格正文的唯一落点、命名与文档定位。
  - [x] SubTask 2.1: 冻结目录落点为 `.trae/specs/phase05_10_dashboard_feedback_formal_spec/`
  - [x] SubTask 2.2: 冻结正式正文文件名为 `dashboard_feedback_spec_v0.1.md`
  - [x] SubTask 2.3: 冻结正文开头必须写清“文档定位 / 上游收敛 / 互链前提 / 当前阶段状态约束”

- [x] Task 3: 冻结正式正文必须完整收敛 `phase05-01 ~ 09` 的全部结果。
  - [x] SubTask 3.1: 冻结页面、区块、路由与点击入口来自 `phase05-01 / 05`
  - [x] SubTask 3.2: 冻结反馈信号、优先级、CTA 与错误语义来自 `phase05-02 / 04 / 06`
  - [x] SubTask 3.3: 冻结跳转、返回与来源上下文来自 `phase05-03 / 06`
  - [x] SubTask 3.4: 冻结后端模块边界、接口分组与合同边界来自 `phase05-07 / 08`
  - [x] SubTask 3.5: 冻结验收环境、fixture、局部错误与返回矩阵来自 `phase05-09`

- [x] Task 4: 冻结正式正文必须覆盖的章节骨架与正文范围。
  - [x] SubTask 4.1: 冻结技术路线、对象范围、页面矩阵、区块矩阵、动作矩阵、数据模型章节
  - [x] SubTask 4.2: 冻结聚合读与反馈信号、跳转与返回上下文、前端页面/状态模型章节
  - [x] SubTask 4.3: 冻结后端模块边界、API 边界与合同矩阵章节
  - [x] SubTask 4.4: 冻结冷启动、fixture 与验收基线、非目标、Done 标准章节
  - [x] SubTask 4.5: 冻结正文不得把上述范围拆到第二入口文档

- [x] Task 5: 冻结正式正文与 `phase01-06`、`phase04-10 / 14` 的互链规则。
  - [x] SubTask 5.1: 冻结 `mvp_spec_v0.1.md` 是 `phase05` 的唯一执行层总上游
  - [x] SubTask 5.2: 冻结 `product_repository_binding_spec_v0.1.md` 是 `Dashboard` canonical owner 与写入边界的直接上游
  - [x] SubTask 5.3: 冻结 `phase04-14` 验收结论是 `phase05` 运行前提与回流语义前提
  - [x] SubTask 5.4: 冻结 `phase05` 不得回退或改写 `phase04` 已验收通过的 reread / 返回路径 / owner 结论

- [x] Task 6: 冻结正式正文的非目标与 Done 标准矩阵。
  - [x] SubTask 6.1: 冻结 Dashboard 当前阶段不得扩写为第二套写入主线、复杂驾驶舱、通知中心或外部导入系统
  - [x] SubTask 6.2: 冻结 Done 标准必须覆盖页面、区块、聚合读、反馈信号、合同、验收基线与下一阶段引用前提
  - [x] SubTask 6.3: 冻结 Done 标准足以支撑后续 `phase05-11 ~ 14`

- [x] Task 7: 产出正式规格正文 `dashboard_feedback_spec_v0.1.md`。
  - [x] SubTask 7.1: 产出文档头部（文档定位 / 上游收敛 / 互链前提 / 状态约束）
  - [x] SubTask 7.2: 产出 §1 技术路线、§2 对象范围、§3 页面矩阵、§4 区块矩阵
  - [x] SubTask 7.3: 产出 §5 动作矩阵、§6 数据模型、§7 聚合读与反馈信号
  - [x] SubTask 7.4: 产出 §8 跳转与返回上下文、§9 前端页面与状态模型
  - [x] SubTask 7.5: 产出 §10 后端模块边界与接口分组、§11 API 边界与合同矩阵
  - [x] SubTask 7.6: 产出 §12 冷启动/fixture/验收基线、§13 非目标、§14 Done 标准、§15 与根级真相源互链、§16 追溯来源矩阵
  - [x] SubTask 7.7: 验证正式正文完整覆盖 dev plan DoD 中"页面、区块、聚合读、反馈信号、API、合同、验收基线、非目标与 Done 标准"

- [x] Task 8: 完成 `phase05-10` 规格一致性校验与入口收口。
  - [x] SubTask 8.1: 验证本 spec 没有把 `phase05-10` 错写成新的并列子规格
  - [x] SubTask 8.2: 验证本 spec 已明确 `dashboard_feedback_spec_v0.1.md` 为唯一正式正文入口
  - [x] SubTask 8.3: 验证本 spec 已覆盖 dev plan DoD 中"页面、区块、聚合读、反馈信号、API、合同、验收基线、非目标与 Done 标准"
  - [x] SubTask 8.4: 验证本 spec 已写清与 `phase01-06`、`phase04-10`、`phase04-14` 的互链一致性
  - [x] SubTask 8.5: 验证 `phase05-01 ~ 09` 被定义为追溯来源而非并列直接执行层入口
  - [x] SubTask 8.6: 验证正式正文 `dashboard_feedback_spec_v0.1.md` 已实际产出并落点在正确目录

# Task Dependencies

- `Task 2` depends on `Task 1`
- `Task 3` depends on `Task 1`, `Task 2`
- `Task 4` depends on `Task 1`, `Task 3`
- `Task 5` depends on `Task 1`, `Task 3`, `Task 4`
- `Task 6` depends on `Task 3`, `Task 4`, `Task 5`
- `Task 7` depends on `Task 1` through `Task 6`
- `Task 8` depends on `Task 1` through `Task 7`
