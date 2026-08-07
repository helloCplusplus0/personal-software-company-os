# Tasks

- [x] Task 1: 冻结 `Product / Repository / Binding` 正式规格正文的文件落点与覆盖范围。
  - [x] SubTask 1.1: 明确正式规格正文目标文件为 `product_repository_binding_spec_v0.1.md`
  - [x] SubTask 1.2: 明确正文至少覆盖页面、动作、模板、绑定关系、数据读写、API、合同、验收基线、迁移边界、非目标、实现设计层结果与 Done 标准
  - [x] SubTask 1.3: 明确正文不能只停留在 `phase04-01 ~ 09` 的零散冻结结论集合

- [x] Task 2: 冻结正式规格正文与上游规格的继承关系。
  - [x] SubTask 2.1: 明确 `mvp_spec_v0.1.md` 是当前阶段唯一执行层上游
  - [x] SubTask 2.2: 明确 `module_registry_spec_v0.1.md` 与 `phase02-12` 验收结论是当前阶段已交付的直接承接结果
  - [x] SubTask 2.3: 明确 `decision_center_spec_v0.1.md` 与 `phase03-14` 验收结论是当前阶段已交付的直接承接结果
  - [x] SubTask 2.4: 明确 `phase04-01 ~ 09` 的冻结结论必须被正式规格正文统一承接
  - [x] SubTask 2.5: 验证正式规格正文不得重新定义第二套阶段边界

- [x] Task 3: 冻结正式规格正文中的页面、动作与模板章节要求。
  - [x] SubTask 3.1: 明确六类页面主线与 `Module Detail` 兼容入口关系必须进入正式规格正文
  - [x] SubTask 3.2: 明确 `CreateProduct / CreateRepository / BindModuleToProduct / BindRepositoryToProduct / MapModuleToRepository` 的动作归属必须进入正式规格正文
  - [x] SubTask 3.3: 明确 `Product / Repository` 模板字段、required / optional、`status` 枚举与最小展示模型必须进入正式规格正文
  - [x] SubTask 3.4: 明确三类绑定关系、候选范围、上下文入口与 reread owner 必须进入正式规格正文

- [x] Task 4: 冻结正式规格正文中的数据、API、合同与迁移章节要求。
  - [x] SubTask 4.1: 明确直接承接数据、候选读取前提、最小读写接口分组与错误语义必须进入正式规格正文
  - [x] SubTask 4.2: 明确 `.proto` 合同边界、服务矩阵、字段编号、共享枚举与 RPC -> HTTP 映射必须进入正式规格正文
  - [x] SubTask 4.3: 明确 `0006_product_repository_binding_mainline.sql`、重置脚本、基线 seed 与兼容迁移边界必须进入正式规格正文

- [x] Task 5: 冻结正式规格正文中的前端与后端实现设计章节要求。
  - [x] SubTask 5.1: 明确页面集合、最小路由结构、组件职责、状态模型与多入口回流规则必须进入正式规格正文
  - [x] SubTask 5.2: 明确后端模块边界、接口分组、关系摘要读取链路、候选读取链路与兼容适配边界必须进入正式规格正文
  - [x] SubTask 5.3: 明确单一 `React Web` 与单一后端主线下的实现设计结果足以直接进入实现

- [x] Task 6: 冻结正式规格正文中的验收基线、非目标、Done 标准与直接上游定位。
  - [x] SubTask 6.1: 明确冷启动路径、异常路径、旧入口兼容与多入口回流验收矩阵必须进入正式规格正文
  - [x] SubTask 6.2: 明确当前阶段非目标矩阵必须进入正式规格正文
  - [x] SubTask 6.3: 明确正式 Done 标准必须进入正式规格正文
  - [x] SubTask 6.4: 明确 `product_repository_binding_spec_v0.1.md` 是后续实现与 `phase05` 的直接上游规格入口
  - [x] SubTask 6.5: 明确正文必须与根级真相源、`phase04` 三件套、`phase02 / phase03` 已交付结果互链一致

- [x] Task 7: 完成规格校验。
  - [x] SubTask 7.1: 验证正式规格正文覆盖范围已完整冻结
  - [x] SubTask 7.2: 验证正式规格正文的继承关系已与 `phase01-06`、`phase02` 已交付结果、`phase03` 已交付结果、`phase04-01 ~ 09` 单值一致
  - [x] SubTask 7.3: 验证正式规格正文的前端、后端、合同、迁移与验收基线承接要求已明确
  - [x] SubTask 7.4: 验证正式规格正文的直接上游定位与真相源关系已明确

- [x] Task 8: 产出 `Product / Repository / Binding` 正式规格正文。
  - [x] SubTask 8.1: 在当前 spec 目录下产出 `product_repository_binding_spec_v0.1.md`
  - [x] SubTask 8.2: 将 `phase04-01 ~ 09` 的冻结结论收敛为完整、单值、可互链的正式规格正文
  - [x] SubTask 8.3: 保证正式规格正文可直接作为实现与 `phase05` 的直接规格入口

# Task Dependencies

- `Task 2` depends on `Task 1`
- `Task 3` depends on `Task 1` and `Task 2`
- `Task 4` depends on `Task 1` and `Task 2`
- `Task 5` depends on `Task 1` and `Task 2`
- `Task 6` depends on `Task 1` and `Task 2`
- `Task 7` depends on `Task 2`, `Task 3`, `Task 4`, `Task 5`, and `Task 6`
- `Task 8` depends on `Task 7`
