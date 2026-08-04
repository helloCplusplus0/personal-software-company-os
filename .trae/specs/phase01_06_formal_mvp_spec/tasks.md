# Tasks

- [x] Task 1: 冻结正式 MVP 规格正文的覆盖范围。将 `phase01-06` 要产出的正式规格正文写成明确的章节清单和覆盖要求。
  - [x] SubTask 1.1: 明确正式规格正文至少覆盖对象、动作、页面、数据、API、非目标与 Done 标准
  - [x] SubTask 1.2: 明确正式规格正文不能只停留在零散冻结结论，而必须形成统一正文

- [x] Task 2: 冻结前五个子任务与正式规格正文的继承关系。确保 `phase01-01` 到 `phase01-05` 的结论被统一承接。
  - [x] SubTask 2.1: 明确技术路线、对象动作、页面输入、数据合同、冷启动导入导出要求都必须被正式规格正文继承
  - [x] SubTask 2.2: 验证正式规格正文不得重新定义第二套边界

- [x] Task 3: 冻结正式规格正文中的核心章节要求。将对象、动作、页面、数据与 API 的正式承接要求写成可校验前提。
  - [x] SubTask 3.1: 明确对象与动作章节必须完整承接核心实体、派生层、后移对象与动作矩阵
  - [x] SubTask 3.2: 明确页面与空状态章节必须完整承接页面范围、页面职责与冷启动路径
  - [x] SubTask 3.3: 明确数据与 API 章节必须完整承接表结构方向、关系结构与合同边界

- [x] Task 4: 冻结正式规格正文中的补充章节要求。将导入导出、基础度量与 Done 标准写成必备章节。
  - [x] SubTask 4.1: 明确导入与导出章节必须承接手动录入优先、最小导出与最小备份要求
  - [x] SubTask 4.2: 明确基础度量指标必须进入正式规格正文
  - [x] SubTask 4.3: 明确正式 Done 标准必须进入正式规格正文

- [x] Task 5: 冻结正式规格正文与根级真相源的关系。确保它成为后续唯一上游，而不是第二套真相源。
  - [x] SubTask 5.1: 明确正式规格正文是后续实现与 `phase02` 的唯一上游规格来源
  - [x] SubTask 5.2: 明确正式规格正文必须与根级真相源互链一致
  - [x] SubTask 5.3: 验证后续阶段不得继续把前置子规格当作长期并列主入口

- [x] Task 6: 完成规格校验。检查本次 `phase01-06` 规格是否具备进入后续子任务的条件。
  - [x] SubTask 6.1: 验证正式规格正文覆盖对象、动作、页面、数据、API、非目标、Done 标准
  - [x] SubTask 6.2: 验证正式规格正文要求与 `phase01-01` 到 `phase01-05` 保持一致
  - [x] SubTask 6.3: 验证正式规格正文要求与根级真相源互链一致
  - [x] SubTask 6.4: 验证该正式规格正文被明确定位为后续实现与 `phase02` 的唯一上游规格来源

- [x] Task 7: 产出正式 MVP 规格正文。将 `phase01-01` 到 `phase01-05` 的冻结结论收敛成一份完整、单值、可互链的正式规格正文文档。
  - [x] SubTask 7.1: 产出覆盖对象、动作、页面、数据、API、非目标、Done 标准的正式规格正文
  - [x] SubTask 7.2: 保证正式规格正文完整继承 `phase01-01` 到 `phase01-05` 的单值结论，不引入第二套边界
  - [x] SubTask 7.3: 保证正式规格正文与根级真相源互链一致，不形成第二套真相源
  - [x] SubTask 7.4: 将正式规格正文与 `phase01-06` spec 目录关联，使其可被追踪而非孤立

- [x] Task 8: 补齐最小 API 矩阵。在正式规格正文中列出核心动作到最小接口的映射。
  - [x] SubTask 8.1: 将 9 个核心动作映射到对应最小 API
  - [x] SubTask 8.2: 将可选动作映射到对应最小 API
  - [x] SubTask 8.3: 明确该 API 矩阵为方向级映射，最终以 `Contract First` 合同为准

- [x] Task 9: 补齐对象最小字段。在正式规格正文中定义各核心对象的最小字段集。
  - [x] SubTask 9.1: 定义 `Product / Module / Release / Repository / Venture / Decision` 的最小字段
  - [x] SubTask 9.2: 明确对象最小字段服务于登记与检索闭环，派生视图与 `Capability` 不参与

- [x] Task 10: 补齐模块准入规则。在正式规格正文中定义 `Module` 进入注册与版本主线的条件。
  - [x] SubTask 10.1: 定义模块登记、命名唯一、状态明确、版本准入与未绑定不阻断的准入条件
  - [x] SubTask 10.2: 明确准入规则不扩大 `v0.1` 范围，未达标模块仅保留登记

# Task Dependencies

- `Task 2` depends on `Task 1`
- `Task 3` depends on `Task 1` and `Task 2`
- `Task 4` depends on `Task 1` and `Task 2`
- `Task 5` depends on `Task 1` and `Task 2`
- `Task 6` depends on `Task 1`, `Task 2`, `Task 3`, `Task 4`, and `Task 5`
- `Task 7` depends on `Task 6`
- `Task 8` depends on `Task 7`
- `Task 9` depends on `Task 7`
- `Task 10` depends on `Task 7`