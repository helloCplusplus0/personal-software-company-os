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

- 明确 `Product / Repository / Module / Decision` 的正式语义说明
- 明确四实体当前只做语义澄清，不做结构重构
- 明确 `Module` 与 `Decision` 的当前阶段解释

DoD：

- 四实体语义可被写入 shared baseline
- 后续 `/spec` 与实现不会再把语义确认误解为 schema 重构

### 第二组：实现设计类子任务

### phase11-04 产出根级上下文真相源治理设计

范围：

- 盘点 `plan.md / AGENTS.md / architecture_map.md / docs/README.md / docs/phase/README.md / project_rules.md` 的重复承载与漂移点
- 明确“谁是单一写者、谁只保留摘要式引用”
- 明确 `PSCO-mvp05-summarize-feedback.md` 作为根级最终共识入口的统一改写策略
- 明确不再允许出现的重复表达模式

DoD：

- 根级治理策略单值化
- 直接可指导根级同步实现
- 不需要继续开放争论“是全量派生还是一次性校准”

### phase11-05 产出最小只读项目上下文聚合导出设计

范围：

- 产出项目上下文聚合只读读取的输入锚点设计
- 产出聚合内容范围与边界设计
- 产出与现有 canonical 数据对齐的投影设计
- 产出 AGENTS 风格 Markdown 导出的职责边界

至少明确：

- 以什么锚点读取当前项目上下文
- 聚合哪些 Product / Repository / Module / Decision / 规则信息
- 哪些信息属于结构化只读输出
- 哪些信息属于 Markdown 导出渲染
- 不做哪些协议与写路径

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
- 同步必要的根级入口摘要与导航

DoD：

- 根级入口不再互相复制主结论
- 悬空引用清零
- 新接手 agent 从根级入口读到的上下文单值一致

### phase11-07 落实最小只读项目上下文聚合读取能力

范围：

- 落实一个最小只读“项目上下文聚合导出”正式承接位
- 保持其为聚合投影，不引入第二套业务事实源
- 复用既有 `.proto + ConnectRPC` 主线，不偷渡新协议层

DoD：

- 已存在可供 agent 消费的最小只读上下文能力
- 只读边界清晰
- 不引入 agent 写回或第二套 canonical API

### phase11-08 落实 AGENTS 风格上下文导出

范围：

- 基于聚合上下文能力，提供 AGENTS 风格或等价 Markdown 风格导出
- 保证输出对新接手 agent 可直接消费
- 不把导出能力扩写为主动注入或仓库写入

DoD：

- 已存在最小文档导出能力
- 可直接服务 PSCO 仓库自身 dogfooding
- 不主动写入外部仓库

### 第四组：验证验收类子任务

### phase11-09 完成 `Project Context Foundation` 的联调、dogfooding 与反回归验证

范围：

- 完成最小工具链验证
- 完成 PSCO 仓库自身的 dogfooding 验证
- 验证新接手 agent 的上下文恢复成本是否下降
- 验证根级文档重复承载与悬空引用是否被收敛
- 留档本阶段明确不做 MCP / CLI / agent 写回 / 对话入口的边界证据

DoD：

- 新接手 agent 可通过少量固定入口恢复项目核心上下文
- 根级治理与最小导出能力均可被真实验证
- 验收证据足以说明本阶段不是抽象设计停留

### 第五组：根级同步类子任务

### phase11-10 完成根级同步、阶段状态回写与下一阶段进入条件留档

范围：

- 回写 `AGENTS.md / plan.md / architecture_map.md / docs/README.md / docs/phase/README.md`
- 留档本阶段正式验收与收口入口
- 明确下一阶段只允许在 `phase11` 正式收口后，再讨论更重的消费通道或受控维护能力

DoD：

- 根级状态、docs 入口与阶段记录同步完成
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
