# phase06_onboarding_sovereignty_reuse_foundation_dev_plan

## 1. 文档定位

本文档定义 `phase06_onboarding_sovereignty_reuse_foundation` 的执行顺序、子任务范围、DoD 与明确不做。

> 收口说明：`phase06` 已完成 `/plan -> /spec -> 实现 -> 验收 -> 收口`。本文保留为该阶段 `/plan` 的执行规划与冻结记录，不承担根级当前阶段状态说明；当前阶段状态以 `AGENTS.md`、`plan.md` 与 `docs/README.md` 为准。

`phase06` 是 `mvp0.2` 的首个交付型 phase。它不是只冻结“下一阶段应该做什么”，而是要完成 `/plan -> /spec -> 实现 -> 验收 -> 收口` 全链路，并最终交付 `Onboarding + Data Sovereignty + Reuse Awareness` 的最小可执行主线。

相较于 `phase05`，本阶段在任务拆分上显式吸取以下经验：

- first-run 与写路径摩擦必须前置冻结，而不是交互实现时临场收缩
- `.proto` 合同、HTTP 映射与前端 mutation 承接位必须在早期任务中明确
- 导出 / 备份属于底座负债，不能后移为尾部补丁
- 复用感知必须尽早进入页面可见反馈，而不是只留在后端内部派生

## 2. 本阶段目标

在 `phase05` 已交付最小 Dashboard 主线的前提下，交付：

- first-run onboarding 与最小录入入口
- 核心资产导出与基础备份路径
- `module_reuse_summary / capability_summary` 的最小可见反馈

使用户能够完成一次“进入系统 -> 录入核心资产 -> 确认资产可带走 -> 看见最小复用反馈”的首轮闭环。

## 3. 子任务清单

### 第一组：冻结类子任务

### phase06-01 冻结 first-run onboarding 边界、页面入口与最小完成条件

范围：

- 冻结 first-run onboarding 是否独立页面、是否与既有 Dashboard 并存
- 冻结首轮最小完成条件与首个成功会话定义
- 冻结冷启动引导与回访用户入口的区别
- 冻结 Onboarding 与既有导航的关系
- 冻结 Onboarding 的正式业务入口路由与回访入口
- 冻结首轮成功会话的推荐执行顺序
- 冻结 `first_run_state` 的最小状态集合、状态跃迁与根级入口判定承接位
- 冻结 cold-start 用户自动导向与 `in_progress` 回访用户 `Continue Onboarding` 的区别

DoD：

- first-run onboarding 职责单值化
- 首轮成功会话完成条件明确
- `Onboarding` 的正式业务入口路由单值化为 `/onboarding`
- 冷启动主 CTA 与回访继续入口单值化
- `Product -> Repository -> Module -> Decision` 的推荐执行顺序单值化
- 不形成第二套根级宿主或复杂向导系统
- `first_run_state` 最小状态机单值化为 `not_started / in_progress / completed`
- 应用入口判定固定由前端根级路由入口守卫承接，不分散到页面级 `useEffect`
- 已明确 `not_started` 与 `in_progress` 的进入条件、默认落点与是否强制跳转

### phase06-02 冻结 `Product / Repository / Module / Decision` 的 draft-first / partial-entry 写路径

范围：

- 冻结四类核心对象的最小必填字段
- 冻结草稿优先、后补完整信息的规则
- 冻结页面、表单、面板之间的唯一 `application` 写入承接位
- 冻结新增前端 mutation 的固定承接位原则
- 冻结四类对象在首轮成功会话中的最小完成状态
- 冻结当前阶段允许的“未绑定也可完成首轮会话”边界
- 冻结本阶段必须回收的既有 create 页面 / 组件级 mutation 范围，避免 Onboarding 与 canonical owner 并存两套写语义

DoD：

- 最小必填字段明确
- `query` 层与写路径边界明确
- 四类对象的“草稿创建成功”与“首轮会话完成”边界明确
- 已明确哪些绑定关系允许后补，哪些对象必须先落一条已持久化记录
- 页面级散装写路径不再作为新增模式继续扩散
- 已明确 phase06 结束时哪些 create 页面必须回收到切片固定承接位，不允许保留第二套成功回流 / 失效刷新语义

### phase06-03 冻结导出、基础备份与恢复前提的正式语义

范围：

- 冻结导出与备份的职责边界
- 冻结最小覆盖资产范围
- 冻结基础恢复前提与错误语义
- 冻结不依赖第三方平台的前提
- 冻结导出 / 备份的正式用户入口位
- 冻结 `Export` 与 `Backup` 的最小覆盖矩阵
- 冻结当前阶段 `backup verified` 的最小验证动作，避免把“文件写出成功”误判为恢复前提已验证

DoD：

- 导出 / 备份不是同义词
- 核心资产覆盖范围明确
- `Export` 最小覆盖 `products / modules / releases / repositories / decisions / decision_links / product_modules / product_repositories / module_repositories`
- `Backup` 的最小覆盖范围不小于 `Export`
- `Export` 正式执行路由单值化为 `/dashboard/export`
- `Backup` 正式执行路由单值化为 `/dashboard/backup`
- 不把自动同步、复杂灾备、多端同步写入前置范围
- `backup verified` 的最小成立条件明确到可执行动作级别：产物生成、manifest 可读、覆盖矩阵可核对、schema / 版本前提可校验

### phase06-04 冻结 `module_reuse_summary / capability_summary` 的最小读模型与页面挂接位

范围：

- 冻结两类派生读模型的最小字段与展示语义
- 冻结 Dashboard、Module Detail、Product Detail 中的最小挂接位
- 冻结最小解释文案与空状态语义
- 冻结两类派生读模型的最小计算口径
- 冻结复用感知的刷新 / 新鲜度语义
- 冻结 `capability_summary` 的最小事实来源、未声明 capability 的 Module 处理方式与空聚合语义

DoD：

- 复用反馈变成正式页面能力
- `module_reuse_summary` 的最小统计口径单值化为“一个 Module 当前被多少 Product 直接复用”
- `capability_summary` 的最小聚合单位与最小字段集合明确
- 复用感知的新鲜度口径单值化为“读取时反映最新已提交状态”
- `Capability` 仍保持派生层，不进入重实体 CRUD
- 不引入新的一级“复用中心”
- `capability_summary` 的最小事实来源单值化，避免实现阶段再临时猜测派生口径
- 已明确未填写 capability 归属的 Module 是否参与聚合，以及不参与时的页面空状态语义

### phase06-05 冻结当前阶段合同、传输与源码约束的执行前提

范围：

- 冻结 `.proto` 作为唯一合同源的新增接口要求
- 冻结 `chi + HTTP JSON` 的承接位与映射策略
- 冻结前端 `application / query / mutation / shared` 四条约束在当前阶段的执行口径
- 冻结 `buf` 工具链与 breaking check 前提
- 冻结复用感知读模型与导出 / 备份 DTO 的字段一致性要求

DoD：

- 新增接口不再长出第二套 JSON contract
- 前端新增写路径必须遵守固定承接位
- `Export / Backup / Reuse Summary` 的 HTTP DTO 必须与 `.proto` 字段保持单值一致
- 合同与源码约束进入阶段验收口径

### 第二组：实现设计产出类子任务

### phase06-06 产出 Onboarding 前端页面、路由与交互流设计

范围：

- 产出 first-run onboarding 页面与步骤分层
- 产出首次录入、草稿保存、继续补全与完成回流的交互流
- 产出移动浏览器下的布局降级策略
- 产出根级默认进入路径、路由入口守卫与 `/onboarding` 之间的跳转关系

DoD：

- 页面、步骤与交互流明确
- 设计结果足以直接进入实现
- 已明确 cold-start / in-progress / completed 三类状态分别由哪个入口承接与如何回流

### phase06-07 产出前端写路径收敛与 mutation 承接位设计

范围：

- 产出四类核心对象新增 / 更新写路径的切片落点
- 产出统一错误归一化、成功回流与 query 失效策略
- 识别需要从页面 / 组件中回收到切片承接位的旧模式
- 列出本阶段必须回收的既有 create 页面清单与对应 application owner

DoD：

- `application` 承接位明确
- mutation 与 query 分层明确
- 设计结果足以指导现有漂移点后续回收
- 已给出旧 create 页面回收清单、迁移顺序与 phase06 收口标准

### phase06-08 产出导出、备份与恢复前提的后端模块边界设计

范围：

- 产出导出 / 备份相关后端模块边界与接口分组
- 产出数据装配、归档与恢复校验前提
- 产出与既有服务、脚本、数据库的边界关系
- 产出最小覆盖矩阵对应的数据装配责任边界
- 产出 `backup verified` 的最小后端校验链与 owner 归属

DoD：

- 后端模块边界明确
- 导出与备份接口 owner 明确
- 最小覆盖矩阵与装配 owner 一一对应
- 不提前冻结脚本名与最终目录细节
- 已明确“备份产物生成成功”和“恢复前提已校验”分别由哪些接口或校验步骤承接

### phase06-09 产出复用感知派生读、Dashboard 挂接与详情挂接设计

范围：

- 产出 `module_reuse_summary / capability_summary` 的 query owner
- 产出 Dashboard 与详情页的最小集成方式
- 产出复用反馈与既有反馈信号的边界关系
- 产出最小统计字段、时间字段与刷新语义
- 产出空状态、零复用状态与有复用状态的展示差异
- 产出 `capability_summary` 的最小事实来源、字段映射与缺省值策略

DoD：

- 复用感知与既有 Dashboard 主线不冲突
- 两类派生读模型的计算口径、字段集合与新鲜度语义明确
- 设计结果足以直接进入实现
- `capability_summary` 的聚合事实来源、显示映射与空状态差异明确

### phase06-10 产出当前阶段最小 Protocol Buffers 合同设计

范围：

- 产出 Onboarding、Export、Backup、Reuse Summary 的最小 `.proto` 合同设计
- 明确消息结构、服务接口、包名与版本语义
- 明确 `.proto -> DTO / HTTP JSON` 的映射策略

DoD：

- `.proto` 合同不晚于正式规格正文进入主线
- 新增字段、枚举与错误语义单值化
- breaking 演进前提明确

### phase06-11 产出联调验收环境、fixture 与恢复基线设计

范围：

- 产出冷启动、部分录入、已完成录入、已导出、已备份、可见复用反馈等最小 fixture 组合
- 产出恢复验证、失败路径与重复执行入口
- 产出首轮成功会话与阶段完成验收矩阵
- 产出“缺少绑定关系仍完成首轮会话”与“绑定补全后再次验证”两类 fixture
- 产出复用感知最新状态验证 fixture
- 产出 cold-start / in-progress / completed 三类 `first_run_state` fixture
- 产出 `backup verified` 基于 manifest 与覆盖矩阵校验的 fixture

DoD：

- 验收环境可重复建立
- 关键状态与失败路径都有单值基线
- 首轮成功会话、数据主权闭合与复用感知最新状态均有独立 fixture 证据
- 不依赖手工补数据才能完成验收
- 入口判定、回访继续、备份校验与复用感知最新状态均可通过 fixture 重复验证

### 第三组：规格、实现与验收子任务

### phase06-12 产出首份 Onboarding + Sovereignty + Reuse 正式规格文档

范围：

- 基于前置冻结与设计产出 `phase06` 对应的 `/spec`
- 作为后续实现与下一阶段的直接上游规格来源

DoD：

- 文档完整覆盖页面、交互、写路径、导出 / 备份、复用读模型、合同、验收基线、非目标与 Done 标准

### phase06-13 落实当前阶段最小 Protocol Buffers 合同主线

范围：

- 将 `phase06-10` 已冻结的 `.proto` 合同正式落地为仓库唯一合同源
- 落地 `buf build / lint / generate / breaking` 的最小工具链入口
- 明确当前阶段 HTTP DTO 与 `.proto` 的语义映射落点

DoD：

- 当前阶段最小 `.proto` 合同已落地为单一合同源
- 过渡传输层不得形成第二套合同源

### phase06-14 实现后端、数据与脚本主线

范围：

- 实现 Onboarding、导出、备份与复用感知所需的最小后端主线
- 实现数据装配、派生读与联调所需脚本 / fixture
- 实现最小恢复验证前提

DoD：

- 后端与数据主线可运行
- 当前阶段关键路径可重复执行
- 不依赖手工补票完成数据主权验证

### phase06-15 实现前端主线

范围：

- 实现 first-run onboarding 前端主线
- 实现新增写路径承接位与既有页面回流
- 实现 Dashboard / Detail 上的最小复用反馈
- 实现导出 / 备份的用户入口

DoD：

- 首轮录入可走通
- 新增写路径遵守固定 mutation 承接位
- 复用反馈可见且可解释
- phase06 覆盖范围内的既有 create 页面不再保留与 Onboarding 冲突的第二套写入语义

### phase06-16 联调、验证与验收

范围：

- 验证首轮录入、部分补全、导出、备份、恢复前提与复用反馈路径
- 验证失败语义、局部错误与重复执行前提
- 验证与 `phase05` 已交付 Dashboard / Feedback 主线的兼容性
- 验证首轮成功会话严格满足四类对象都已落持久化记录
- 验证导出包含绑定 / 关联关系而不是仅主实体
- 验证 `module_reuse_summary / capability_summary` 反映最新已提交状态
- 验证 cold-start 与 `in_progress` 回访入口行为符合冻结口径
- 验证 `backup verified` 依赖 manifest / 覆盖矩阵 / schema 前提校验，而不是仅以文件生成成功代替

DoD：

- 首轮成功会话可重复走通
- 数据主权路径可验证
- 复用反馈路径可验证
- 首轮成功会话、导出覆盖矩阵与复用感知最新状态均通过验收
- 无破坏 `phase05` 已交付边界的回归
- 根级入口行为、回访继续入口与备份校验语义均通过验收

### phase06-17 审核、根级同步与下一阶段进入条件回写

范围：

- 审核 `phase06` 是否完成 `/plan -> /spec -> 实现 -> 验收 -> 收口`
- 回写根级真相源、入口与下一阶段切换条件
- 归档当前阶段规划、规格与验收结论

DoD：

- 根级状态与阶段实际收口结论一致
- 活动文档无孤岛
- 不在根级文档中越权预设 `phase07` 名称

## 4. 本阶段明确不做

- `Opportunity / Feature / Experiment` 的流程化主线
- `Capability` 重实体 CRUD
- GitHub OAuth / 自动导入
- AI 一级工作台
- 自动扫描 / 知识图谱
- Rust Intelligence Layer
- 模板级复用正式落地
- 真实项目 `dry-run` 作为本阶段主交付
