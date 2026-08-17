# phase13_project_governance_profile_foundation_dev_plan

## 1. 文档定位

本文档定义 `phase13_project_governance_profile_foundation` 的执行顺序、子任务范围、DoD 与明确不做。

`phase13` 直接承接 `phase12` 正式收口后的新共识：PSCO 下一步优先要管理的是四实体信息、全局规范资产与项目级治理画像，并将其组织成 agent 在 IDE 场景可直接消费的项目简报输入。

## 2. 本阶段目标

在 `phase12` 已完成四实体语义与共享只读消费收口的前提下，交付：

- 项目级治理画像的正式边界
- 当前项目范式 v1 与 canonical 根级文件集合
- 全局规范资产的正式承接方式
- agent 项目简报输入的最小正式结构

使 PSCO 可以先稳定管理“长期有效、跨项目不应丢失的治理信息”，而不是一开始就扩写为目录扫描器、模板平台或 Git 推进跟踪平台。

补充边界：

- `phase13` 同时承接“人类在 Web 端手工维护项目治理信息”与“agent 在 IDE 场景读取该信息”的双侧需求，但二者共享同一套 `PSCO-native facts`；
- `phase13` 允许新建项目级治理承接层，但不得新增第五个业务主实体；
- Git 推进跟踪、模板仓库接入、自动同步与更重受控维护能力当前仍不进入本阶段正式实现。

## 3. 子任务清单

### 第一组：边界收敛类子任务

### phase13-01 冻结 `Project Governance Profile Foundation` 的范围边界、成功标准与非目标

范围：

- 冻结本阶段单一主交付能力为 `Project Governance Profile Foundation`
- 冻结本阶段与 Git 推进跟踪、模板仓库接入、自动同步、MCP / CLI / agent 写回的边界
- 冻结本阶段成功标准、DoD 与阶段收口口径

DoD：

- 本阶段主交付能力与非目标单值化
- 不把后续自动化与模板化能力偷渡到本阶段
- 进入 `/spec` 前，后续执行者不再需要猜“本阶段到底做什么”

### phase13-02 冻结 `PSCO-native facts / IDE-accessible context / controlled synced projection` 三层边界

范围：

- 冻结哪些信息属于 PSCO 正式管理对象
- 冻结哪些信息继续留给 IDE / agent 现场读取
- 冻结哪些信息属于后续才允许进入的受控同步投影

DoD：

- PSCO 不再被误设计为目录扫描器
- agent 项目输入不再默认依赖目录全文被 PSCO 接管
- 后续执行者能机械回答“什么该先做，什么后做”

### phase13-03 冻结项目级治理层与四实体主线的关系

范围：

- 冻结项目级治理层的正式职责
- 冻结它与 `Product / Repository / Module / Decision` 的关系
- 冻结“它不是第五个业务主实体”的口径

DoD：

- 不再把目录结构、全局规范与当前阶段信息硬塞进 `Decision`
- 不再让执行者临场判断“新承接层到底是不是新实体”
- 项目级治理层承接对象与边界单值化

### 第二组：实现设计类子任务

### phase13-04 产出项目治理画像数据模型与字段设计

范围：

- 产出 `project_profile_version / track_type / template_source / docs_workflow_layout` 等核心字段设计
- 产出 `canonical_root_files / global_constraint_refs / current_phase_*` 等最小字段矩阵
- 产出必填、可选、只读、后续可自动校验字段的分类

DoD：

- 字段模型足以直接进入 `/spec`
- 执行者不需要再猜“第一版到底先做哪些字段”
- 当前项目范式 v1 能被结构化承接

### phase13-05 产出后端合同、存储与读写边界设计

范围：

- 产出项目治理画像的后端承接位设计
- 明确它与四实体既有合同、读路径、写路径的关系
- 冻结第一版“手工维护优先、自动同步后置”的后端边界
- 产出全局规范资产逐项承接矩阵，至少覆盖：
  - `project_rules.md`
  - `TECH_STACK_BASELINE.md`
  - `AGENTS.md`
  - `architecture_map.md`
  - `plan.md`
  - `README.md`
  - `global_skills.md`
  - `project_skills.md`
- 对每项全局规范资产明确以下承接策略：
  - 只存 `entry_ref + role`
  - 存 `entry_ref + role + structured summary`
  - 允许 markdown 正文回源但不入库
  - 第一版禁止全文入库

DoD：

- 后端承接位单值化
- 不复制第二套四实体事实源
- 不把 markdown 全文入库误当成第一版必要条件
- 全局规范资产逐项承接策略单值化，后续执行者不再需要猜“哪份文件只做引用、哪份文件需要结构化摘要”

### phase13-06 产出前端信息架构与手工维护设计

范围：

- 产出项目治理画像在 Web 端的正式承接位设计
- 冻结它在信息架构中的层级、入口与展示边界
- 明确哪些内容供人类维护，哪些只供 agent 消费，哪些只保留摘要回看
- 第一版前端正式承接位冻结为：`Repository detail`
- 当前阶段不得同时新增独立“项目治理画像”一级页面或第二入口；若未来需要提升为独立入口，只能在 `phase13` 收口后作为新进入条件讨论

DoD：

- 前端不再重演 `phase12` 的“把验收层信息做成大块常驻 UI”的问题
- 页面层级、表单结构与读写边界足以直接进入 `/spec`
- 目录真实路径不被误当作面向普通用户的主内容
- 后续执行者不再需要临场决定“第一版到底挂在 Repository detail 还是另起页面”

### phase13-07 产出 agent 项目简报输入与读取设计

范围：

- 产出供 agent 在 IDE 场景读取的最小项目简报结构
- 冻结四实体信息、治理画像、全局规范资产与当前阶段入口如何组合
- 冻结该 brief 与 IDE 自带目录读取能力的边界
- 冻结 brief 的最小 schema，至少明确：
  - `repository`
  - `governance_profile`
  - `global_assets`
  - `current_phase`
  - `products[]`
  - `modules[]`
  - `decisions[]`
- 冻结 brief 的解析协议：
  1. `repository_id` 是唯一正式锚点
  2. `products / modules / decisions` 只能从同一 `repository_id` 驱动的 PSCO 结构化关系中解析
  3. 若存在多个 `Module / Decision`，必须返回数组摘要，不得伪造单一“当前 module / decision”
  4. agent 若需目录全文，继续由 IDE / agent 现场能力读取，不通过 PSCO brief 补做第二套目录扫描

DoD：

- agent 项目简报输入单值化
- 不再需要执行者临场决定“给 agent 发什么”
- 不让 PSCO 与 IDE 长出两套并列上下文机制
- 后端、前端与 agent 消费侧共享同一份 brief schema，不再各做一版“我理解的项目简报”

### 第三组：源代码实现类子任务

### phase13-08 落实项目治理画像后端主线

范围：

- 落实项目治理画像的后端合同、服务、存储与读取主线
- 保持四实体既有主线不变
- 保持“手工维护优先”的第一版策略

DoD：

- 后端已能正式保存并读取项目治理画像
- 不引入自动扫描、自动同步或全文入库
- 不产生第二套项目事实源

### phase13-09 落实前端治理画像承接与手工维护入口

范围：

- 落实 Web 端的项目治理画像承接位
- 落实当前项目范式 v1、canonical 根级文件、全局规范资产与当前阶段信息的维护入口
- 保持四实体主线与详情页结构稳定

DoD：

- 人类用户已能在 PSCO 中手工维护这层治理信息
- UI 层级与产品定位一致
- 不把目录真实文件路径误做成大块重复说明

### phase13-10 落实 agent 项目简报读取主线

范围：

- 落实供 agent 读取的项目简报输入
- 打通四实体信息、治理画像与全局规范资产的最小读取主线
- 保持只读消费优先，不新增写回

DoD：

- agent 已能获取“全局信息 + 当前项目 PSCO 管理信息”
- 不需要通过额外目录扫描才能恢复 PSCO 侧正式上下文
- 不引入第二套并列读取协议

### 第四组：验证验收类子任务

### phase13-11 完成 `Project Governance Profile Foundation` 的联调、dogfooding 与反回归验证

范围：

- 完成后端、前端与关键读取主线的最小工具链验证
- 完成人类维护路径与 agent 读取路径的双侧 dogfooding
- 验证当前项目范式 v1、全局规范资产与当前阶段入口是否能被真实复验
- 留档本阶段明确不做 Git 推进跟踪、模板仓库接入、自动同步与 agent 写回的边界证据
- 冻结固定验收协议，至少覆盖：
  - 固定样本
  - 固定页面 / 固定入口
  - 固定问题
  - 固定 rerun 记录格式
- 固定样本继续使用 `phase11 / phase12` dogfooding 样本：
  - `Repository`：`personal-software-company-os`
  - `repository_id`：`ca261521-8daf-4248-8f12-43525326e759`
  - `Product`：`PSCO`
- 固定 Web 验证页面第一版冻结为：
  - `/repositories/$repositoryId`
- 固定 agent 读取入口第一版冻结为：
  - 基于同一 `repository_id` 返回的 `project brief for agent`
  - `AGENTS.md`
  - `plan.md`
  - 由治理画像承接的全局规范资产结构化结果
- 固定验收提问至少覆盖：
  1. 当前项目治理画像版本与技术路线是什么，在哪个固定入口可确认
  2. 当前 canonical 根级文件集合是否已被正式承接，在哪个固定入口可确认
  3. 当前全局规范资产是否以结构化摘要 + 入口关系被正式承接，在哪个固定入口可确认
  4. 当前 agent 项目简报是否由同一 `repository_id` 驱动，且没有伪造第二套目录扫描结果
  5. 当前第一版前端正式承接位是否仍是 `Repository detail`，而没有长出并列第二入口
  6. 当前阶段是否仍严格没有进入 Git 推进跟踪、模板仓库接入、自动同步与 agent 写回
- 固定 rerun 协议至少要求记录：
  - 使用了哪个固定样本与 `repository_id`
  - 使用了哪个固定 Web 页面与 agent 入口
  - 每个问题的回答结果
  - 失败点与是否达标

DoD：

- Web 与 agent 都能回到同一组项目治理画像与规范资产
- 当前项目范式 v1 能被 PSCO 稳定承接
- 本阶段不是停留在抽象讨论，而是形成可维护、可读取、可复验的正式主线
- 不同执行者能够按同一固定样本、同一固定入口与同一固定问题 rerun

### 第五组：根级同步类子任务

### phase13-12 完成根级同步、阶段收口与下一阶段进入条件回写

范围：

- 回写 `AGENTS.md / plan.md / docs/README.md / architecture_map.md / docs/phase/README.md`
- 留档本阶段正式验收与收口入口
- 明确 Git 推进跟踪、模板仓库接入、自动同步与更重受控维护能力只允许在 `phase13` 正式收口后，再依据新条件讨论或进入

DoD：

- 根级状态、docs 入口与阶段记录同步完成
- 不长出新的孤岛文档
- 下一阶段进入条件单值化

## 4. 明确不做

本阶段明确不做：

1. Git 推进跟踪主线
2. 模板仓库自动 bootstrap
3. 目录全文扫描入库
4. MCP 协议层正式实现
5. CLI 工具正式实现
6. agent 自动写回、Draft 接口、审批流
7. IDE 插件或流程控制台
8. 第五个业务主实体
9. 第二套与四实体并列的事实源
10. 继续放大 `phase12` 的共享只读 UI 表达

## 5. 子任务依赖关系

为避免后续执行时再次出现“先自动化、后收口治理边界”的顺序错乱，当前阶段依赖关系冻结如下：

1. `phase13-01` 是全阶段边界前提，后续所有子任务都直接依赖它
2. `phase13-02` 与 `phase13-03` 是信息分层与治理层边界前提，`phase13-04 ~ 07` 必须直接承接这两项结论
3. `phase13-04 ~ 07` 属于实现设计层，必须先于 `phase13-08 ~ 10`
4. `phase13-04` 先冻结字段与项目范式 v1，再进入后端、前端与 agent brief 设计
5. `phase13-05 / 06 / 07` 必须共享同一套字段模型与边界，不得各自长出第二套语义
6. `phase13-08` 只依赖 `phase13-04 / 05`
7. `phase13-09` 只依赖 `phase13-04 / 06`
8. `phase13-10` 只依赖 `phase13-04 / 05 / 07`
9. `phase13-11` 依赖 `phase13-08 ~ 10`
10. `phase13-12` 依赖 `phase13-11`
