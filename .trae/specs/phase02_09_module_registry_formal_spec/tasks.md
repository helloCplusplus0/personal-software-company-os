# Tasks

- [x] Task 1: 冻结 `Module Registry` 正式规格正文的文件落点与覆盖范围。
  - [x] SubTask 1.1: 明确正式规格正文目标文件为 `module_registry_spec_v0.1.md`
  - [x] SubTask 1.2: 明确正文至少覆盖页面、动作、数据读写、API、空状态、非目标、实现设计层结果与 Done 标准
  - [x] SubTask 1.3: 明确正文不能只停留在 `phase02-01 ~ 08` 的零散冻结结论集合

- [x] Task 2: 冻结正式规格正文与上游规格的继承关系。
  - [x] SubTask 2.1: 明确 `mvp_spec_v0.1.md` 是当前阶段唯一执行层上游
  - [x] SubTask 2.2: 明确 `phase02-01 ~ 08` 的冻结结论必须被正式规格正文统一承接
  - [x] SubTask 2.3: 验证正式规格正文不得重新定义第二套阶段边界

- [x] Task 3: 冻结正式规格正文中的页面、动作与空状态章节要求。
  - [x] SubTask 3.1: 明确 `Module Registry / List`、`Module Create`、`Module Detail`、`Release Create` 的页面主线必须进入正式规格正文
  - [x] SubTask 3.2: 明确 `CreateModule`、`CreateRelease`、`BindModuleToProduct`、`MapModuleToRepository` 与 `Decision` 只读入口的动作归属必须进入正式规格正文
  - [x] SubTask 3.3: 明确空状态、冷启动路径、默认回流路径与错误呈现原则必须进入正式规格正文

- [x] Task 4: 冻结正式规格正文中的数据、API 与后端边界章节要求。
  - [x] SubTask 4.1: 明确直接承接数据、最小读取前提与候选读取前提必须进入正式规格正文
  - [x] SubTask 4.2: 明确 `ModuleListRead / ModuleDetailRead / ModuleCreateWrite / ModuleReleaseWrite / ModuleBindingWrite` 的接口分组必须进入正式规格正文
  - [x] SubTask 4.3: 明确 `Product / Repository` 候选读取边界与 `Decision` 附属读取边界必须进入正式规格正文
  - [x] SubTask 4.4: 明确后端 `handler / service / repository / candidate` 分层与文件落点必须进入正式规格正文

- [x] Task 5: 冻结正式规格正文中的前端实现设计章节要求。
  - [x] SubTask 5.1: 明确页面集合、最小路由结构、URL 语义与组件职责必须进入正式规格正文
  - [x] SubTask 5.2: 明确列表、创建、详情、版本登记与绑定动作的状态模型必须进入正式规格正文
  - [x] SubTask 5.3: 明确单一 `React Web` 下的 `PC / 移动浏览器` 布局降级策略必须进入正式规格正文

- [x] Task 6: 冻结正式规格正文中的非目标、Done 标准与上游定位。
  - [x] SubTask 6.1: 明确当前阶段非目标矩阵必须进入正式规格正文
  - [x] SubTask 6.2: 明确正式 Done 标准必须进入正式规格正文
  - [x] SubTask 6.3: 明确 `module_registry_spec_v0.1.md` 是后续实现、验收与 `phase03` 的直接上游规格入口
  - [x] SubTask 6.4: 明确正文必须与根级真相源和 `phase02` 阶段文档互链一致

- [x] Task 7: 完成规格校验。
  - [x] SubTask 7.1: 验证正式规格正文覆盖范围已完整冻结
  - [x] SubTask 7.2: 验证正式规格正文的继承关系已与 `phase01-06`、`phase02-01 ~ 08` 单值一致
  - [x] SubTask 7.3: 验证正式规格正文的前端与后端实现设计承接要求已明确
  - [x] SubTask 7.4: 验证正式规格正文的直接上游定位与真相源关系已明确

- [x] Task 8: 产出 `Module Registry` 正式规格正文。
  - [x] SubTask 8.1: 在当前 spec 目录下产出 `module_registry_spec_v0.1.md`
  - [x] SubTask 8.2: 将 `phase02-01 ~ 08` 的冻结结论收敛为完整、单值、可互链的正式规格正文
  - [x] SubTask 8.3: 保证正式规格正文可直接作为 `phase02-10 / 11 / 12 / 13` 与 `phase03` 的直接规格入口

# Task Dependencies

- `Task 2` depends on `Task 1`
- `Task 3` depends on `Task 1` and `Task 2`
- `Task 4` depends on `Task 1` and `Task 2`
- `Task 5` depends on `Task 1` and `Task 2`
- `Task 6` depends on `Task 1` and `Task 2`
- `Task 7` depends on `Task 2`, `Task 3`, `Task 4`, `Task 5`, and `Task 6`
- `Task 8` depends on `Task 7`
