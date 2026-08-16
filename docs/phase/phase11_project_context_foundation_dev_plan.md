# phase11_project_context_foundation_dev_plan

## 1. 文档定位

本文档定义 `phase11_project_context_foundation` 的执行顺序、子任务范围、DoD 与明确不做。

`phase11` 是 `mvp0.5` 的首个正式 phase。它直接承接 `PSCO-mvp05-summarize-feedback.md` 的最终仲裁结论，不再继续扩大“方向是否清楚”的讨论，而是把已收敛的判断落成一轮可验证交付。

## 2. 本阶段目标

在 `phase10` 已完成 `Asset-Action Closure`、并且 `mvp0.5` 已明确“PSCO 是上下文系统、agent 当前先消费后维护”的前提下，交付：

- 根级上下文真相源治理
- 最小只读项目上下文导出
- AGENTS 风格上下文导出
- 以 PSCO 仓库自身为第一消费场景的 dogfooding 验证

使新接手 agent 不再需要阅读大量根级文档与历史评审，才能恢复当前项目核心上下文。

补充边界：

- `phase11` 同时承接“PSCO 自身仓库治理”与“面向未来项目的通用只读上下文能力”，但两层合同必须显式分离；
- `PSCO` 当前仓库真实存在的根级文件，只用于本仓库治理与第一轮 dogfooding，不得直接外推为未来所有项目的统一目录要求；
- “最佳实践项目模板”当前尚未进入正式讨论与冻结，不得在本阶段偷渡为默认前提。

## 3. 子任务清单

### 第一组：边界收敛类子任务

### phase11-01 冻结 `Project Context Foundation` 的范围边界、成功标准与非目标

范围：

- 冻结本阶段单一主交付能力为 `Project Context Foundation`
- 冻结本阶段与 MCP / CLI / agent 写回 / 前端对话式入口的边界
- 冻结本阶段与四实体结构重构、重型 GitHub 集成、知识图谱的边界
- 冻结本阶段成功标准、DoD 与阶段收口口径

DoD：

- 本阶段主交付能力与非目标单值化
- 不把后续能力偷渡到本阶段
- 进入 `/spec` 前，后续执行者不再需要猜“本阶段到底做什么”

### phase11-02 冻结 PSCO 作为“上下文系统”的正式定位与 web / agent 分工

范围：

- 冻结 PSCO 不是开发流程控制器的正式口径
- 冻结 web 继续作为全局查看、回顾、校对与最终确认渠道
- 冻结 agent 当前只承接现场上下文消费
- 冻结 web 与 agent 共享 Go backend canonical core 的约束

DoD：

- 后续执行者不再把 PSCO 理解成 IDE 现场流程编排器
- web 与 agent 的分工边界单值化
- 不会再出现第二套语义与第二套流程的设计冲动

### phase11-03 冻结四实体语义确认口径

范围：

- 冻结 `Product / Repository / Module / Decision` 的正式语义说明：
  - `Product`：经营目标与交付容器
  - `Repository`：代码仓库身份对象与项目锚点
  - `Module`：可复用能力资产，允许后置提炼
  - `Decision`：规则、约束、选择与依据的索引对象
- 冻结四实体当前只做语义澄清，不做结构重构、实体拆并或关系主线重写
- 冻结 `Module` 与 `Decision` 的当前阶段解释：
  - `Module` 当前代表可复用能力资产，允许在后续真实复用沉淀中继续提炼，当前阶段不要求重写其 schema、层级或注册主线
  - `Decision` 当前代表规则、约束、选择与依据的索引对象，用于支撑项目上下文恢复与只读导出，不在本子任务内扩写为审批流、流程引擎或结构重构入口

DoD：

- 四实体语义已在 `architecture_plan`、`dev_plan` 与 `shared_baseline` 中以单值口径冻结
- 后续 `/spec` 与实现不会再把语义确认误解为 schema 重构
- `Module` 与 `Decision` 的当前阶段解释已在三件套中可直接引用，不需要后续执行者二次解释

### 第二组：实现设计类子任务

### phase11-04 产出根级上下文真相源治理设计

范围：

- 冻结仅针对 `PSCO` 自身仓库的根级治理对象集合：`README.md / plan.md / AGENTS.md / architecture_map.md / docs/README.md / docs/phase/README.md / project_rules.md / global_skills.md`
- 冻结根级入口单一写者规则：
  - `plan.md`：阶段状态与推进路线的唯一正式承接位
  - `architecture_map.md`：目录结构、文档分类、迁移落点的唯一正式承接位
  - `TECH_STACK_BASELINE.md`：技术栈正文的唯一正式承接位
  - `README.md`：项目总览入口与受控跳转，不重复承载当前 phase 正文或最终共识正文
  - `AGENTS.md`：入口摘要，不重复承载完整 phase 状态正文
  - `global_skills.md`：项目内通用方法映射说明，不承接当前 phase 状态、最终共识或目录落点正文
  - `docs/README.md`：workflow 总入口，不重复承载完整目录落点正文
  - 其他入口只保留摘要式引用或受控跳转
- 冻结 `PSCO-mvp05-summarize-feedback.md` 作为根级最终共识入口的统一改写策略，相关根级入口统一指向该文档
- 冻结不再允许出现的重复表达模式：同一主结论在多个根级文档中重复承载正文
- 冻结治理清单不外推为未来所有项目的固定目录模板

正式产物至少包括：

- 根级入口治理矩阵（单一写者规则表）
- 重复承载清单与目标落点清单
- 悬空引用清理清单
- 收口后的单一写者规则表

DoD：

- 根级治理策略已在 `architecture_plan`、`dev_plan` 与 `shared_baseline` 中以单值口径冻结
- 治理矩阵、清理清单与单一写者规则可直接指导根级同步实现
- 当前阶段已明确冻结为"一次性校准"策略，不需要继续开放争论"是全量派生还是一次性校准"

### phase11-05 产出最小只读项目上下文聚合导出设计

范围：

- 产出项目上下文聚合只读读取的输入锚点设计
- 产出聚合内容范围与边界设计
- 产出与现有 canonical 数据对齐的投影设计
- 产出 AGENTS 风格 Markdown 导出的职责边界
- 明确通用能力与 `PSCO` 当前仓库文件结构之间的边界

至少明确：

- 以什么锚点读取当前项目上下文
- 聚合哪些 Product / Repository / Module / Decision / 规则信息
- 哪些信息属于结构化只读输出
- 哪些信息属于 Markdown 导出渲染
- 不做哪些协议与写路径

必须进一步冻结：

- `repository_id` 为当前阶段唯一正式结构化输入锚点
- 当前阶段不以本地路径、Git remote URL、`product_id` 或工作区扫描作为并列主锚点
- 当前阶段只承接“已完成 Repository Binding”的仓库上下文读取
- 未绑定仓库的失败语义
- `Decision` 聚合口径冻结为：以当前 `Repository` 为根，只合并三类直接 canonical 关系命中的 `Decision`：
  - 直接链接到当前 `Repository` 的 `Decision`
  - 直接链接到“当前 `Repository` 已绑定 `Product`”的 `Decision`
  - 直接链接到“当前 `Repository` 已映射 `Module`”的 `Decision`
- 当前阶段不得继续沿 `Product -> Module -> 其他 Repository` 做递归扩张；若超出上述三类命中范围，视为当前阶段不纳入
- 同一 `Decision` 若同时命中多类关系，必须以 `decision_id` 去重，并保留命中来源摘要；执行者不得临场决定保留哪一条
- 当前阶段结构化只读主列表只承接非 `archived` 的 `Decision`；`archived` 不进入主导出列表，不允许执行者自行决定是否混入
- 结构化只读读取继续落在 Go backend 的 `.proto + ConnectRPC` 正式主线
- Markdown 导出从同一结构化读取结果单向派生，不形成第二套事实源
- 当前阶段不把消费侧项目目录中的固定文件名/固定目录结构当作必要输入合同
- 当前阶段不把“统一项目模板”作为前置依赖

正式产物至少包括：

- 输入锚点与失败语义说明
- 结构化只读输出字段边界
- Markdown 导出字段边界
- 结构化读取 -> Markdown 导出的单向派生关系说明
- “哪些依赖 PSCO canonical 关系，哪些不依赖消费侧项目目录结构”的边界说明

DoD：

- 已形成后续 `/spec` 可直接承接的设计输入
- 聚合导出能力的输入、输出与职责边界单值化
- 不需要执行者临场再猜“上下文到底聚合到什么程度”

### 第三组：源代码实现类子任务

### phase11-06 落实根级上下文真相源治理

范围：

- 回收根级文档之间的重复 phase 状态、重复目录落点与重复共识入口
- 修复指向不存在文件 `PSCO-summarize-feedback.md` 的引用
- 让 `PSCO-mvp05-summarize-feedback.md` 成为根级最终共识的单值入口
- 同步根级治理矩阵中已冻结的活动入口，包括 `README.md` 与 `global_skills.md`
- 本子任务必须逐项审计并按需同步以下文件：
  - `README.md`
  - `AGENTS.md`
  - `plan.md`
  - `architecture_map.md`
  - `docs/README.md`
  - `docs/phase/README.md`
  - `project_rules.md`
  - `global_skills.md`
  - `PSCO-mvp05-summarize-feedback.md`

补充说明：

- 本子任务只治理 `PSCO` 自身仓库入口，不定义未来消费侧项目必须具备哪些同名文件
- 上述文件允许出现“无需改动”的结果，但必须在治理清单中逐项记录：`已审计 / 是否修改 / 不修改原因`

DoD：

- 根级入口不再互相复制主结论
- 悬空引用清零
- 新接手 agent 从根级入口读到的上下文单值一致
- 治理矩阵中的目标文件已全部完成逐项审计，不允许只改部分显眼入口后直接判定完成

### phase11-07 落实最小只读项目上下文聚合读取能力

范围：

- 落实一个最小只读“项目上下文聚合导出”正式承接位
- 保持其为聚合投影，不引入第二套业务事实源
- 复用既有 `.proto + ConnectRPC` 主线，不偷渡新协议层
- 以 `repository_id` 作为唯一正式结构化输入锚点
- 明确未绑定仓库的失败态与返回语义
- 不要求消费侧项目目录与 `PSCO` 当前仓库拥有相同结构

正式产物至少包括：

- 一个最小结构化只读读取承接位
- 对应的输入合同、输出边界与失败语义
- 与既有 canonical 数据的一致性说明

DoD：

- 已存在可供 agent 消费的最小只读上下文能力
- 只读边界清晰
- 不引入 agent 写回或第二套 canonical API
- 执行者不需要再临场决定“current project”如何绑定
- 执行者不需要假设未来每个项目都遵守 `PSCO` 当前仓库的文件布局

### phase11-08 落实 AGENTS 风格上下文导出

范围：

- 基于聚合上下文能力，提供 AGENTS 风格或等价 Markdown 风格导出
- 保证输出对新接手 agent 可直接消费
- 不把导出能力扩写为主动注入或仓库写入
- 保证导出内容从 phase11-07 的结构化只读结果单向派生

DoD：

- 已存在最小文档导出能力
- 可直接服务 PSCO 仓库自身 dogfooding
- 不主动写入外部仓库
- 不形成独立于结构化读取之外的第二套事实源

### 第四组：验证验收类子任务

### phase11-09 完成 `Project Context Foundation` 的联调、dogfooding 与反回归验证

范围：

- 完成最小工具链验证
- 完成 PSCO 仓库自身的 dogfooding 验证
- 验证新接手 agent 的上下文恢复成本是否下降
- 验证根级文档重复承载与悬空引用是否被收敛
- 留档本阶段明确不做 MCP / CLI / agent 写回 / 对话入口的边界证据
- 留档“PSCO 自身治理样本”与“跨项目通用能力合同”已分离的边界证据

验收协议至少冻结为：

- 采用同一份 dogfooding 剧本执行一次“旧路径基线”与一次“新路径目标”
- 旧路径基线入口集合固定为：
  - `AGENTS.md`
  - `plan.md`
  - `project_rules.md`
  - `architecture_map.md`
  - `docs/README.md`
  - `PSCO-mvp05-summarize-feedback.md`
- 旧路径基线执行时，不允许额外读取 `phase11` 新增的结构化导出结果或 AGENTS 风格项目上下文导出
- 新路径固定入口集合冻结为：
  - `AGENTS.md`
  - `PSCO-mvp05-summarize-feedback.md`
  - 基于同一 `repository_id` 生成的 AGENTS 风格项目上下文导出
- 新路径不允许临时增加第 `4` 个入口来补齐答案
- 验收提问固定为 `5` 问：当前 phase、直接上游、单一主交付、明确不做、当前项目关联的 `Repository / Product / Module / Decision` 摘要入口
- 新路径必须能够完整回答上述 `5` 问，且不再依赖散落的根级多文档补读
- 验收记录必须留档入口清单、回答结果、失败点与是否达标
- 验收结论必须额外确认：当前 dogfooding 只证明 `PSCO` 自身仓库样本可消费，不等于已经冻结未来所有项目模板

DoD：

- 新接手 agent 可通过不超过 `3` 个固定入口恢复项目核心上下文
- 根级治理与最小导出能力均可被真实验证
- 验收证据足以说明本阶段不是抽象设计停留
- dogfooding 结果可被客观复验，而不是只停留在主观体感

### 第五组：根级同步类子任务

### phase11-10 完成根级同步、阶段状态回写与下一阶段进入条件留档

范围：

- 回写本阶段正式验收与收口完成后的阶段状态、验收入口与下一阶段进入条件
- 留档本阶段正式验收与收口入口
- 明确下一阶段只允许在 `phase11` 正式收口后，再讨论更重的消费通道或受控维护能力

补充说明：

- `phase11-06` 负责治理实现与活动入口内容对齐
- `phase11-10` 只负责验收后的正式状态回写与进入条件留档，不重复承接治理实现本身

DoD：

- 根级状态、docs 入口与阶段记录在验收后同步完成
- 不长出新的孤岛文档
- 下一阶段进入条件单值化

## 4. 明确不做

本阶段明确不做：

1. MCP 协议层正式实现
2. CLI 工具正式实现
3. agent 自动写回、Draft 接口、审批流
4. 前端对话式 agent 入口
5. 知识图谱、自动扫描、重型 GitHub / Gitea 集成
6. 四实体结构重构或大规模 schema 扩张
7. 主动向外部仓库注入文件
8. 静态文件从 backend 全量派生

## 5. 子任务依赖关系

为避免后续执行时顺序错乱，当前阶段依赖关系冻结如下：

1. `phase11-01` 是全阶段边界前提，后续所有子任务都直接依赖它
2. `phase11-02` 与 `phase11-03` 是共享语义前提，`phase11-04 ~ 05` 必须直接承接这两项结论
3. `phase11-04 ~ 05` 属于实现设计层，必须先于 `phase11-06 ~ 08`
4. `phase11-06` 只依赖 `phase11-01 / 02 / 04`
5. `phase11-07` 与 `phase11-08` 依赖 `phase11-01 / 02 / 03 / 05`
6. `phase11-09` 依赖 `phase11-06 ~ 08`
7. `phase11-10` 依赖 `phase11-09`
