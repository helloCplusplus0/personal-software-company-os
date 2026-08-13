# phase10_asset_action_closure_foundation_dev_plan

## 1. 文档定位

本文档定义 `phase10_asset_action_closure_foundation` 的执行顺序、子任务范围、DoD 与明确不做。

`phase10` 是 `mvp0.4` 的第一个正式业务 phase。它直接承接 `mvp0.3 dry-run` 阻断项修复与收口，不再处理“能不能用”的 P0 问题，而是正式处理“用起来是否形成真实动作闭环”的结构性问题。

## 2. 本阶段目标

在 `phase06` 已交付 `Onboarding`、`phase08` 已交付 `Operating Review Loop`、`phase09` 已交付 `Template Reuse + Derived Intelligence`，且 `fix_001 ~ fix_003` 已完成收口的前提下，交付：

- `Onboarding` 首轮建链引导
- `Decision` 最小但真实的生命周期闭环
- Dashboard / Review / Detail pages 的下一步动作承接矩阵
- `Current Focus / pending signals` 的真实经营语义回归
- 页面、合同、服务与 reread 的统一回流路径

使用户不再需要靠跨多个详情页的手工补链与语义猜测推进项目，而能在 PSCO 中直接理解并执行“下一步动作”。

## 3. 子任务清单

### 第一组：边界收敛类子任务

### phase10-01 冻结 `Asset-Action Closure` 的范围边界、成功标准与非目标

范围：

- 冻结本阶段单一主交付能力为 `Asset-Action Closure`
- 冻结本阶段与 `Agent Consumption Layer / Cross-Project Convention Asset` 的边界
- 冻结本阶段与 `Onboarding` 工作流引擎化、`Decision Intelligence`、新实体回归的边界
- 冻结成功标准、DoD 与阶段收口口径

DoD：

- 本阶段主交付能力与非目标单值化
- 不把后续 phase 的内容偷渡到本阶段
- 进入 `/spec` 前，后续执行者不再需要猜“本阶段到底做什么”

### phase10-02 冻结 `Onboarding` 从首轮登记到首轮建链引导的语义与动作矩阵

范围：

- 冻结 `Onboarding` 六段式既有主线的保留边界
- 冻结哪些建链动作应在创建过程中直接承接，哪些保留为成功后 handoff
- 冻结从 `product -> repository -> module -> decision` 的最小动作与关系建立矩阵
- 冻结 `Onboarding` 与 canonical `Product / Repository / Module / Decision` 写路径的承接关系

DoD：

- `Onboarding` 不再只是六段登记说明页
- 建链动作、成功后回流与后续 handoff 语义单值化
- 不新增第二套草稿或工作流引擎语义

### phase10-03 冻结 `Decision` 生命周期、detail CTA 与 pending 语义的统一矩阵

范围：

- 冻结既有四态下的最小生命周期矩阵
- 冻结 `Decision Detail` 在各状态下允许展示的 CTA
- 冻结 `Dashboard / Daily Review / Current Focus` 与 `Decision.status` 的统一语义
- 冻结“退出 pending”的正式条件与 reread 规则

DoD：

- `Decision.status` 的解释口径单值化
- 详情页 CTA、review signal 与 dashboard signal 不再各讲一套语义
- 不引入第五态，不引入第二事实源

### 第二组：实现设计类子任务

### phase10-04 产出 `Onboarding` 建链流、返回链与空态/失败态设计

范围：

- 产出冷启动用户从 `/dashboard -> /onboarding` 的首轮建链流
- 产出创建成功后的默认下一步动作与返回链设计
- 产出空态、失败态、已存在实体场景与中途返回再进入的设计
- 产出移动浏览器下的最小降级策略

DoD：

- 页面流与交互流足以直接进入实现
- 创建、建链、返回与空态语义可被浏览器验收机械执行
- 不需要依赖页面临场发挥解释下一步

### phase10-05 产出 `Dashboard / Daily Review / Detail pages` 的下一步动作承接设计

范围：

- 产出 `Dashboard / Daily Review / Product Detail / Module Detail / Repository Detail / Decision Detail` 的 CTA 矩阵
- 产出每个页面“何时显示、显示什么、跳去哪里、返回后如何 reread”的设计
- 产出 `Current Focus` 与 detail CTA 之间的承接关系

DoD：

- 每个关键页面都有明确“下一步动作”语义
- CTA 不互相冲突、不重复拼装第二套业务主线
- 成功回流与 reread 规则足以指导后续实现

### phase10-06 产出前端读写 owner、route caller 与成功回流设计

范围：

- 产出 `Onboarding / Dashboard / Review / Detail pages` 的 caller 与 owner inventory
- 产出需要新增或回收的 read owner / application owner 设计
- 产出失效刷新、成功回流、错误归一化与返回链设计
- 识别必须回收的页面级散装 mutation / CTA 编排点

DoD：

- caller 与 owner 的一对一映射足以指导实现
- query 层与 application 层边界清晰
- 页面层不会继续内联第二套动作编排

### phase10-07 产出后端合同、服务与读模型承接设计

范围：

- 产出 `Onboarding` 建链引导所需的最小合同与服务设计
- 产出 `Decision` 生命周期推进与 pending reread 所需的最小合同与服务设计
- 产出 `Current Focus / pending signals / next-step CTA` 相关读模型与 query owner 设计
- 识别需要沿用与需要新增的 `.proto / Connect / service / store` 承接位

DoD：

- 合同、服务与数据承接位单值化
- 不复制既有 canonical facts
- 足以直接进入 `/spec` 与代码实现

### 第三组：源代码实现类子任务

### phase10-08 落实 `Onboarding` 首轮建链引导与 canonical handoff

范围：

- 实现 `Onboarding` 在既有六段主线下的建链引导
- 实现创建过程中的最小关系建立与成功后 handoff
- 实现相应的前后端承接位、读写路径与浏览器级交互

DoD：

- 冷启动用户能顺畅完成首轮建链
- 建链结果回到 canonical owner，而不是留在页面局部状态
- 后续补链摩擦显著下降

### phase10-09 落实 `Decision` 生命周期闭环、detail CTA 与 pending reread 统一

范围：

- 落实 `Decision Detail` 状态推进、CTA 矩阵与成功回流
- 落实 `Dashboard / Daily Review / Current Focus` 的统一 pending 语义
- 落实相关 reread、失效刷新与浏览器行为

DoD：

- `Decision` 已形成最小但真实生命周期
- Detail、Review、Dashboard 之间的语义一致
- 不再出现“有提示但无动作出口”或“已处理却仍误报”

### phase10-10 落实关键 detail pages 的下一步动作承接矩阵

范围：

- 落实 `Product Detail / Module Detail / Repository Detail` 的动作 CTA
- 落实从 detail page 返回 Dashboard / Review 的 reread 语义
- 回收当前页面中分散、歧义或缺失的动作承接位

DoD：

- 关键详情页都能回答“下一步做什么”
- CTA 指向 canonical path，而不是临时局部路径
- 详情页不再只是信息展示壳层

### 第四组：验证验收类子任务

### phase10-11 完成 `Asset-Action Closure` 联调、浏览器验收与反回归验证

范围：

- 完成 `buf / go test / frontend build` 工具链验证
- 完成 `Onboarding` 首轮建链流的浏览器验收
- 完成 `Dashboard / Daily Review / Decision Detail / Product Detail / Module Detail / Repository Detail` 的动作链浏览器验收
- 完成 `Current Focus / pending signals` 的反回归验证
- 留档本阶段明确不做 `Agent Consumption Layer / 新实体回归 / 第五态状态机` 的边界证据

DoD：

- 工具链、API、浏览器与关键反回归均通过
- 用户能从空态或接近空态场景完成首轮建链
- `Decision` 生命周期与 pending reread 行为可被真实页面机械验证
- Detail pages 的下一步动作承接证据已留档

### 第五组：根级同步类子任务

### phase10-12 完成根级同步、阶段收口与下一阶段进入条件回写

范围：

- 回写 `AGENTS.md / plan.md / docs/README.md / architecture_map.md / docs/phase/README.md`
- 留档本阶段正式验收与收口入口
- 明确下一阶段只允许在 `phase10` 收口后，再进入 `Agent Consumption Layer`

DoD：

- 根级状态、docs 入口与阶段记录同步完成
- 不长出新的孤岛文档
- 下一阶段进入条件单值化

## 4. 明确不做

本阶段明确不做：

1. 新增长期核心实体主线
2. 新增第五个 `DecisionStatus`
3. 把 `Onboarding` 做成通用工作流引擎
4. 把 `Dashboard / Review` 做成任务管理器
5. 提前实现 `Agent Consumption Layer`
6. 提前实现 `Cross-Project Convention Asset`
7. 引入 AI 工作台、聊天式主入口或 agent 写入主线
