# phase08_operating_review_loop_foundation_dev_plan

## 1. 文档定位

本文档定义 `phase08_operating_review_loop_foundation` 的执行顺序、子任务范围、DoD 与明确不做。

`phase08` 是 `phase07` 收口后的首个正式业务 phase。它不是一次性承接整个 `mvp0.3`，而是先完成 **Operating Review Loop** 这一单一主交付能力，并把 `Template Reuse / Derived Intelligence Deepening / Real-Project Dry-Run` 留在后续 phase 或独立验收闸中承接。

## 2. 本阶段目标

在 `phase07` 已完成传输主线收口、`phase03 ~ phase06` 已交付业务能力可稳定消费的前提下，交付：

- Dashboard 正式 review 入口
- daily review 与 weekly review 两条独立可验收路径
- 当前焦点、代表性反馈与待处理决策的统一承接
- 对 `phase05` Feedback 与 `phase06` Reuse Awareness 的正式消费
- `Feedback -> Decision -> Update` 最小闭环
- review 结论回流既有实体的正式路径

使用户能够从 Dashboard 进入经营动作，而不是只停留在“查看状态”。

## 3. 子任务清单

### 第一组：边界冻结类子任务

### phase08-01 冻结 Operating Review Loop 的范围边界、成功标准与非目标

范围：

- 冻结当前 phase 的唯一中心主线为 `Operating Review Loop`
- 冻结当前阶段必须回答的经营动作问题与最小成功会话
- 冻结本阶段与 `Template Reuse / Derived Intelligence / Real-Project Dry-Run` 的边界
- 冻结 review loop 不得演化为通用任务管理器的约束

DoD：

- 当前 phase 的单一主交付能力明确
- review loop 的成功标准明确
- 后续支撑能力与独立验收闸不混入本阶段 DoD
- 不把未来阶段写成当前版本既成事实

### phase08-02 冻结 Dashboard review 入口、页面职责与路由承接位

范围：

- 冻结 daily / weekly review 的最小入口形态
- 冻结 daily / weekly review 的差异语义、输入数据范围与完成定义
- 冻结 daily / weekly review 各自的最小成功会话与独立验收口径
- 冻结 Dashboard 与 review 页面 / 面板 / 弹层之间的职责边界
- 冻结从 Dashboard 进入 review 的正式路由与回流路径
- 冻结 current focus、feedback signal、pending decision 的统一承接位

DoD：

- Dashboard 不再只是总览页
- review 入口与回流路径单值化
- daily / weekly review 不再是“同页双按钮”的模糊表达
- daily / weekly review 已形成各自单独可验证的完成定义
- 不在多个页面重复长出第二套 review 入口

### phase08-03 冻结 `Feedback -> Decision -> Update` 闭环的动作边界与 owner

范围：

- 冻结从反馈信号进入决策动作的正式路径
- 冻结 review 中 `Decision` 的承接位与回写语义
- 冻结决策结果回流 `Product / Module / Repository` 等既有实体的最小写路径与 action handoff 矩阵
- 冻结前端 mutation owner、后端 command owner 与错误归一化边界

DoD：

- `Decision` 继续保持经营中心地位
- review 结论不会停留在卡片或弹层内
- `Decision / Module / Product / Repository` 的允许动作与禁止动作明确
- 写路径不分散到页面级临时编排

### phase08-04 冻结本阶段合同、读模型与记录模型的最小边界

范围：

- 冻结 review 入口、review 上下文、review 动作与结果回流的最小正式合同
- 冻结当前 phase 必需的读模型与写模型边界
- 冻结 `phase05` Feedback 与 `phase06` Reuse Awareness 的正式消费边界
- 冻结是否需要 review 记录，以及其轻量化边界
- 冻结既有领域实体事实源与新读模型之间的关系
- 冻结当前真实 caller / route / query owner / application owner inventory

DoD：

- `.proto` 继续作为唯一长期合同源
- 若新增 review 记录，不升级为新的长期核心实体
- `phase05 / phase06` 既有数据的正式消费范围已冻结为当前 `/spec` 必答项
- 当前真实 caller / owner 清单足以直接进入 `/spec`
- 不长出第二套事实主线或并列状态体系

### 第二组：实现设计产出类子任务

### phase08-05 产出 Dashboard review 入口、页面流与交互流设计

范围：

- 产出 Dashboard 到 review 的页面流与交互流
- 产出 daily / weekly review 的最小信息组织方式与差异说明
- 产出 current focus、feedback signal、pending decision 的编排方式
- 产出 weekly review 对 `reuse snapshot / module_reuse_summary / capability_summary` 的消费方式
- 产出 review 完成后回流既有详情页或列表页的交互流

DoD：

- 页面流与交互流足以直接进入实现
- daily / weekly review 的 UI 骨架、数据来源与完成动作差异已明确
- daily / weekly review 各自的最小成功会话已能单独进入验收
- review 入口不与既有 Dashboard 信息密度冲突
- 已明确移动浏览器下的最小降级策略

### phase08-06 产出前端读写承接位与状态流设计

范围：

- 产出 review 相关 read layer 的切片落点
- 产出 `phase05` Feedback 与 `phase06` Reuse Awareness 在 review read layer 中的消费位置
- 产出 review action application owner、成功回流与 query 失效策略
- 产出 `Decision` 承接、实体更新回流与错误反馈设计
- 识别必须回收的页面级临时编排点
- 产出 caller 与 owner 的一对一映射表

DoD：

- `query` 与 `application` 边界明确
- `reuse snapshot / module_reuse_summary / capability_summary` 的消费路径明确
- caller 不会跨页面漂移成第二套 owner
- review 编排不散落在 route / 页面 / 展示组件
- 设计结果足以指导后续源码实现

### phase08-07 产出后端服务、合同与最小数据承接设计

范围：

- 产出 review 相关 proto 合同与 Connect 服务设计
- 产出 review 读模型、动作命令与结果回流的后端 owner
- 产出与既有 Dashboard / Decision / Product / Module / Repository 服务的协作边界
- 产出与 `phase06` 复用感知读模型协作的最小承接设计
- 产出 review 记录或结果落点的最小数据承接设计
- 产出 review 关键路径的工具链与 API 验收清单

DoD：

- 合同、服务与数据承接位单值化
- 不复制既有业务事实源
- `phase05 / phase06` 既有服务与读模型已被明确纳入 review 消费边界
- `buf / go build / API smoke` 验收口径已冻结
- 设计结果足以直接进入 `/spec`

### 第三组：源码实现类子任务

### phase08-08 落实 review 相关合同、后端承接与前端 owner 收敛

范围：

- 实现 review 相关 proto / Connect 合同
- 实现后端 owner、前端读写 owner 与必要数据承接
- 为后续 review 入口、动作流与回流路径提供正式使能位

DoD：

- review 相关合同与 owner 单值一致
- 前后端已具备进入 review 入口与动作实现的正式承接位
- 本阶段不以页面级临时编排作为长期稳态

当前结果（2026-08-11）：

- 已完成 `phase08-08` 源码落地、独立复核修复与联调验收
- `ReviewService`、`backend/internal/review/`、`frontend/src/features/review/`、`/reviews/daily`、`/reviews/weekly` 与 `DashboardPrimaryActionPanel` dual review launcher 已形成正式承接位
- 浏览器与 API 验收已确认 `Decision / Product / Module / Repository` canonical handoff 与 `next-step result` 提交路径可运行

### phase08-09 落实 Dashboard review 入口、双路径 review 会话与统一动作承接

范围：

- 实现 Dashboard review 入口
- 实现 daily review 与 weekly review 两条独立会话路径
- 实现 current focus、feedback signal、pending decision 的统一承接
- 实现 weekly review 对 `reuse snapshot / module_reuse_summary / capability_summary` 的最小消费
- 实现 review 入口与既有 Dashboard 模块的最小整合

DoD：

- 用户能从 Dashboard 分别进入 daily review 与 weekly review
- daily / weekly review 不以同一套数据装配和完成定义冒充双路径
- review 入口不再只是文案占位
- `phase05 / phase06` 已交付读模型已被正式消费进 review

### phase08-10 落实 `Feedback -> Decision -> Update` 闭环、结果回流与临时编排清理

范围：

- 实现从反馈到决策的动作流
- 实现决策结果回流既有实体的最小路径
- 实现错误语义、成功回流与必要刷新
- 回收本阶段新增过程中产生的临时散装编排点

DoD：

- 闭环可重复执行
- `Decision` 在 review 中保持中心地位
- 不引入第二套长期任务系统
- 前后端不保留并列临时主线

收口补记：

- `phase08-10` 已完成源码修复、构建验证、运行时健康检查与真实浏览器/E2E 走查
- 已验证 `Dashboard -> Daily Review -> Decision Create -> Decision Detail -> Module Detail -> Dashboard` 真实可交互闭环
- 当前可按通过验收并完结该子任务；后续仅需在 `phase08-11` 继续覆盖更完整的联调、反回归与双路径浏览器验收

### 第四组：验证验收类子任务

### phase08-11 完成 review loop 联调、浏览器验收与反回归验证

范围：

- 分别验证 `Dashboard -> Daily Review -> Decision -> Update` 与 `Dashboard -> Weekly Review -> Decision -> Update` 两条关键路径
- 验证 `Decision` 正式承接与至少一种实体回流 / action handoff 已落地
- 验证 weekly review 已正式消费 `reuse snapshot / module_reuse_summary / capability_summary` 最小读模型
- 完成 `buf build / lint / generate`
- 完成 `go build ./...`
- 完成 `npx tsc -b --noEmit`
- 完成 `frontend build`
- 完成 review 关键 Connect procedure 与 `/api` 访问 smoke
- 完成浏览器端关键交互验证
- 完成既有 `phase03 ~ phase06` 相关页面的最小反回归验证
- 记录不做 `Template Reuse / Derived Intelligence / dry-run` 的本阶段边界证据

DoD：

- 关键经营动作链通过验收
- daily / weekly review 两条路径均已独立通过验收
- 合同、前后端构建与 API smoke 全部通过
- 浏览器端不存在“API 成功但 UI 崩溃”的收口缺口
- 本阶段边界未漂移

### 第五组：根级同步类子任务

### phase08-12 完成根级同步与后续进入条件回写

范围：

- 回写 `AGENTS.md`
- 回写 `plan.md`
- 回写 `docs/README.md`
- 回写 `architecture_map.md`
- 回写 `docs/phase/README.md`

DoD：

- 根级入口已反映 `phase08` 为当前正式业务 phase
- `phase07` 已退回最近完成阶段的规划与冻结记录角色
- 后续支撑能力 phase 与 dry-run phase 只保留进入条件表达，不提前建立正式命名

## 4. 本阶段明确不做

- 在本阶段直接实现 `Template Reuse`
- 在本阶段直接实现 `Derived Intelligence Deepening`
- 在本阶段直接执行真实项目 `dry-run`
- 在本阶段引入新的长期核心实体
- 在本阶段把 review loop 演化为通用任务 / 项目管理系统
- 在本阶段回退 `phase07` 已完成的 ConnectRPC 传输主线

## 5. 本阶段 Done 标准

只有当以下条件同时满足时，`phase08` 才算完成：

1. Dashboard 已正式承接 daily / weekly review 双入口
2. daily review 与 weekly review 已形成各自独立的最小成功会话与验收路径
3. 当前焦点、代表性反馈与待处理决策已形成统一动作入口
4. `phase05` Feedback 与 `phase06` Reuse Awareness 已被正式消费进 review 主线
5. `Feedback -> Decision -> Update` 已形成最小闭环
6. review 结论已能回流既有实体或既有 canonical action handoff
7. review 相关合同、读模型、写路径与 owner 已单值收敛
8. 已完成 `buf / go build / tsc / frontend build / API smoke / 浏览器` 多层验收与最小反回归验证
9. 根级真相源已同步当前阶段状态与后续进入条件
