# phase04_product_and_repository_binding_foundation_dev_plan

## 1. 文档定位

本文档定义 `phase04_product_and_repository_binding_foundation` 的执行顺序、子任务范围、DoD 与明确不做。

当前根级真相源已完成 `phase04` 收口并切换到 `phase05_dashboard_feedback_foundation`。本文档保留为 `phase04` 的执行规划与任务拆分记录；文中“当前阶段”均指 `phase04` 当时上下文，不覆盖项目当前已进入 `phase05` 的根级状态。

`phase04` 继续遵守交付型 phase 原则：不是只把 `Product` 与 `Repository` 写成两份规格，而是要完成 `/plan -> /spec -> 实现 -> 验收 -> 收口` 全链路，并最终交付 `Product Registry + Repository Binding` 最小可执行闭环。

相较于 `phase02` 与 `phase03`，本阶段在任务拆分上显式吸取以下经验：

- 绑定动作归属不能继续长期临时承接，必须在规划期单值化
- `.proto` 合同与验收环境必须前置进入任务主线
- 主交付对象、只读连接对象与兼容迁移对象必须分开写清
- 正式规格正文要在前端、后端、合同、验收设计都冻结后再进入，减少后续返工

## 2. 本阶段目标

在 `phase02` 已交付 `Module Registry`、`phase03` 已交付 `Decision Center` 的前提下，交付 `Product Registry + Repository Binding` 最小可执行闭环，使 `Product / Repository / Binding` 从“存在于正式规格与临时入口”进入“可结构化登记、可稳定绑定、可独立验收”的正式执行主线。

## 3. 子任务清单

### 第一组：冻结类子任务

### phase04-01 冻结 Product Registry 与 Repository Binding 页面边界和信息结构

范围：

- 冻结 `Product List / Product Create / Product Detail`
- 冻结 `Repository Binding / List / Create / Detail(or Workspace)` 的最小页面结构
- 冻结 `Product Registry`、`Repository Binding`、`Module Detail` 之间的入口关系
- 冻结五个核心动作与六个页面之间的单值动作归属矩阵
- 冻结桌面与移动浏览器下的基础信息密度策略

DoD：

- `Product Registry` 与 `Repository Binding` 页面职责单值化
- `CreateProduct / CreateRepository / BindRepositoryToProduct / BindModuleToProduct / MapModuleToRepository` 的页面拥有者单值化
- `Module Detail` 不再被长期保留为第二绑定工作台
- 不形成第二套移动端 UI 方案
- 页面边界与 `phase01-06` 正式规格正文一致

### phase04-02 冻结 Product / Repository 模板、状态语义与最小展示模型

范围：

- 冻结 `Product` 最小结构化模板
- 冻结 `Repository` 最小结构化模板
- 冻结 `status` 的最小状态语义
- 冻结列表展示、详情展示与绑定展示所需的最小字段集合

DoD：

- `Product` 与 `Repository` 记录模板明确
- 最小字段集合与字段级 required / optional 已明确
- 列表与详情的最小读模型明确
- `status` 的持久化枚举、默认值与列表 `statusFilter` 取值已明确
- `status` 必须进一步单值化为：当前阶段创建写入是否要求前端/DTO/`.proto` 显式提交；若继续要求创建必填，则“默认 `active`”只表示预填并显式提交 `active`，不得再保留“由服务端隐式补默认值”的并行解释
- `provider` 是否采用受控枚举、以及仓库列表是否引入 `providerFilter` 已明确
- 不引入超出 `v0.1` 的复杂生命周期、自动扫描字段或远程导入字段

### phase04-03 冻结三类绑定关系、候选范围与上下文入口

范围：

- 冻结 `BindRepositoryToProduct / BindModuleToProduct / MapModuleToRepository` 的当前阶段关系范围
- 冻结三类绑定动作的页面归属与上下文入口
- 冻结 `phase02` 临时绑定承接点的迁移边界
- 冻结迁移完成后的 canonical owner、旧入口保留级别与兼容跳转参数

DoD：

- `product_repositories / product_modules / module_repositories` 的关系语义明确
- `Product Detail`、`Repository Binding`、`Module Detail` 的交互归属明确
- 候选读取、已绑定排除、无候选空状态与重复绑定前提明确
- `BindModuleToProduct` 的主归属是否留在 `Product Detail` 已单值化
- 旧 `Module Detail` 入口只保留为兼容跳转还是仍允许本页直写已单值化
- 不把 `Decision Center`、`Module Registry` 重新扩写为并列绑定主线

### phase04-04 冻结数据读写范围、接口边界与错误语义前提

范围：

- 冻结 `Product / Repository / Binding` 所需的最小数据读写范围
- 冻结列表读取、详情读取、创建写入、绑定写入与候选读取的最小接口边界
- 冻结创建失败、目标不存在、重复绑定等错误语义前提

DoD：

- 当前阶段数据与接口边界明确
- 三类绑定写入的校验失败类型已进入阶段规划
- 候选读取与空结果语义已明确
- 关键异常路径已进入阶段规划，而不是联调时临时补写
- 不提前冻结超出当前阶段的 Dashboard 聚合接口

### 第二组：实现设计产出类子任务

### phase04-05 产出前端页面、路由与组件分层设计

范围：

- 产出 `Product List / Create / Detail` 的页面分层
- 产出 `Repository Binding / List / Create / Detail(or Workspace)` 的页面分层
- 产出与 `Module Detail` 之间的最小路由与上下文承接方式
- 产出列表、详情、绑定工作台的组件分层
- 产出 `PC / 移动浏览器` 双场景下的布局降级策略

DoD：

- 前端页面与路由分层明确
- 页面级组件职责明确
- 无第二套移动端 UI 架构
- 设计结果足以直接进入实现

### phase04-06 产出前端状态模型与交互流设计

范围：

- 产出产品列表、仓库列表的查询状态、筛选状态与空状态模型
- 产出产品创建、仓库创建、绑定工作台与详情页的交互状态流转
- 产出成功、失败、空状态、候选为空与返回路径
- 产出列表查询条件在路由搜索参数与页面状态之间的承接策略
- 产出从 `Module Detail`、`Product Detail`、`Repository Binding` 多入口进入时的返回路径与上下文恢复规则

DoD：

- 页面级状态模型明确
- 列表、创建、详情、绑定之间的交互流明确
- 列表筛选维度已明确，至少冻结 `queryText` 与 `statusFilter`
- 仓库列表若增加 `providerFilter`，必须在此阶段单值化
- 查询条件的默认来源、路由持久化策略与刷新恢复策略已明确
- 从创建页或详情页返回列表时，必须恢复原有搜索上下文
- 从 `Module Detail` 兼容入口跳入正式绑定主入口后，返回路径、来源标记与默认回流页面已明确
- 设计结果足以直接进入实现

### phase04-07 产出后端模块边界与接口分组设计

范围：

- 产出 `Product Registry` 与 `Repository Binding` 在后端的模块边界
- 产出列表读取、详情读取、创建写入、绑定写入与候选读取的接口分组
- 产出与 `Module Registry`、`Decision Center` 的服务侧连接边界
- 产出迁移完成后每个动作的 canonical owner 与旧接口兼容策略

DoD：

- 后端模块边界明确
- 读写接口分组明确
- 临时候选读取迁移策略明确
- `BindModuleToProduct`、`BindRepositoryToProduct`、`MapModuleToRepository` 的写组 owner 与 reread owner 已明确
- 不提前冻结 Go 数据访问层具体工具
- 设计结果足以直接进入实现

### phase04-08 产出 Product / Repository / Binding 最小 Protocol Buffers 合同设计

范围：

- 基于前置冻结结果产出 `Product / Repository / Binding` 最小 `.proto` 合同设计
- 明确 `ProductListRead / ProductDetailRead / ProductWrite / RepositoryListRead / RepositoryDetailRead / RepositoryWrite / BindingWrite / CandidateRead` 的消息结构、服务接口、包名与版本语义
- 明确当前阶段 `.proto` 与 `chi + JSON HTTP` 过渡传输层的显式映射策略

DoD：

- `.proto` 合同不晚于正式规格正文进入阶段主线
- 合同字段语义、字段编号与页面动作单值一致
- 三类绑定写入与候选读取的消息边界明确
- 合同演进规则明确，包括 `reserved` 与 breaking check 前提

### phase04-09 产出联调验收环境、重置基线与兼容迁移设计

范围：

- 产出 `Product / Repository / Binding` 联调所需的数据库重置、基线种子与 fixture 设计
- 产出绑定关系迁移、旧入口兼容与最小验收基线
- 产出从空状态到首个产品、首个仓库与首轮绑定的冷启动验收路径
- 产出三类绑定成功后的 reread 页面、兼容跳转路径与多入口回流验收矩阵

DoD：

- 验收环境建立方式明确
- 重置脚本、基线数据、旧绑定兼容与异常路径验证要求已进入阶段任务
- 冷启动主路径、兼容入口主路径与三类绑定写入后的 reread 验收路径已明确
- 不再依赖临时手工 SQL 才能完成验收

### 第三组：规格、实现与验收子任务

### phase04-10 产出首份 Product / Repository / Binding 正式规格文档

范围：

- 基于前九个子任务产出 `phase04` 对应的 `/spec`
- 作为后续实现与 `phase05` 的直接上游规格来源

DoD：

- 文档完整覆盖页面、动作、数据读写、API、合同、验收基线、非目标、迁移边界与 Done 标准
- 与 `phase01-06` 正式 MVP 规格正文、`phase02` 正式规格与 `phase03` 正式规格/验收结论互链一致

### phase04-11 落实 Product / Repository / Binding 最小 Protocol Buffers 合同主线

范围：

- 将 `phase04-08` 已冻结的 `.proto` 合同正式落地为仓库内唯一合同源
- 落地 `buf build / lint / generate / breaking` 的最小工具链入口
- 明确当前阶段 DTO/HTTP 适配层与 `.proto` 的语义映射落点

DoD：

- `Product / Repository / Binding` 最小 `.proto` 合同已落地为单一合同源
- `buf` 校验链可运行，breaking check 基准路径正确
- 过渡传输层不得形成与 `.proto` 并列的第二套合同源

### phase04-12 实现后端与数据主线

范围：

- 实现 `Product Registry` 所需的最小后端读写接口
- 实现 `Repository Binding` 所需的最小后端读写接口
- 实现 `products / repositories / product_modules / product_repositories / module_repositories` 与候选读取的数据主线
- 实现联调所需的重置脚本、基线 seed 与最小 fixture

DoD：

- 后端读写接口可运行
- 数据主线与已冻结边界一致
- 历史绑定与临时承接迁移逻辑可重复执行
- 联调环境可重复建立，不依赖手工补数据

### phase04-13 实现前端 Product Registry 与 Repository Binding 主线

范围：

- 实现 `Product List / Create / Detail` 前端主线
- 实现 `Repository Binding / List / Create / Detail(or Workspace)` 前端主线
- 实现绑定关系面板、候选选择与从 `Module Detail / Product Detail` 发起的最小上下文入口
- 落实单一 `React Web` 下的 `PC / 移动浏览器` 双场景布局策略

DoD：

- 前端主线可运行
- 产品创建、仓库创建、三类绑定关系在前端可走通
- `Module Detail` 中旧绑定入口已收敛为兼容入口或轻量跳转，不再形成第二主工作台
- 无第二套移动端 UI 架构

### phase04-14 联调、验证与验收

范围：

- 完成前后端联调
- 完成 `CreateProduct / CreateRepository / BindRepositoryToProduct / BindModuleToProduct / MapModuleToRepository` 的最小验收路径
- 完成空状态、错误态、候选为空、目标不存在、重复绑定与返回路径验证
- 完成三类绑定成功后的 reread、旧入口兼容跳转与多入口返回路径验证

DoD：

- `Product Registry + Repository Binding` 最小主线可完整走通
- 三类绑定动作的 canonical owner 页面、成功写入后的 reread 页面与返回路径可被重复复核
- 验收结果可重复复核，并可明确证明当前阶段已形成可运行交付物
- 发现的问题已收口到当前阶段，不遗留隐性阻断

### phase04-15 审核与根级同步

范围：

- 完成 `phase04` 文档互链复核
- 回写根级状态
- 确认下一阶段入口

DoD：

- 根级文档与 `phase04` 文档保持单值一致
- `plan.md` 中 `phase04` 状态更新正确
- `phase05` 的进入条件清楚

## 4. 明确不做

本阶段不做：

- Dashboard 聚合反馈实现
- `Decision -> Product / Repository` 正式关联写入
- `Feature / Opportunity / Experiment` 的重新引入
- GitHub OAuth / 自动导入
- 自动扫描代码
