# phase09_template_reuse_derived_intelligence_foundation_dev_plan

## 1. 文档定位

本文档定义 `phase09_template_reuse_derived_intelligence_foundation` 的执行顺序、子任务范围、DoD 与明确不做。

`phase09` 是 `phase08` 收口后的后续支撑能力 phase。它不是另起一条业务主线，而是承接 `phase08` 已完成的 `Operating Review Loop`，把 `Template Reuse` 与 `Derived Intelligence Deepening` 从“已有读模型和方向共识”推进为“可直接支撑下一次创造”的正式交付能力。

## 2. 本阶段目标

在 `phase08` 已完成 `Feedback -> Decision -> Update` 最小经营回路、`phase06` 已交付 `reuse_summary` 基础读模型的前提下，交付：

- `Module` 组合快照的最小正式资产定义
- 从模板候选进入 `Product Create` 预填的正式动作链
- `capability gap / reuse opportunity` 的最小解释性提示
- `Weekly Review` 与 `Product Create` 的正式消费路径
- 支撑能力结果回流到 `Dashboard / Review / Product Detail` 的最小 reread 路径

使用户能够更快开始下一次创造，而不是只停留在“看见复用”与“看见摘要”。

## 3. 子任务清单

### 第一组：边界冻结类子任务

### phase09-01 冻结 `Template Reuse + Derived Intelligence` 的范围边界、成功标准与非目标

范围：

- 冻结当前 phase 的单一主交付能力为“下一次创造的加速支撑层”
- 冻结 `Template Reuse` 与 `Derived Intelligence` 的正式职责与相互关系
- 冻结本阶段与 `Operating Review Loop`、`Real-Project Dry-Run`、`AI Context Enhancement` 的边界
- 冻结不新增长期核心实体、不建设模板平台与 AI 工作台的约束

DoD：

- 当前 phase 的单一主交付能力明确
- `Template Reuse / Derived Intelligence` 的最小成功定义明确
- `dry-run` 与 `mvp0.4+` 候选方向不混入本阶段 DoD
- 不把未来阶段写成当前版本既成事实

### phase09-02 冻结模板级复用资产、候选来源与 `Product Create` 预填动作链

范围：

- 冻结 `Module` 组合快照的最小语义与候选来源
- 冻结模板候选如何只从既有 `Product / Module / Binding` 已持久化事实派生
- 冻结模板候选进入 `Product Create` 的正式 handoff 与返回链
- 冻结预填后继续编辑、提交创建与结果回流的最小闭环
- 冻结 `/products/new` 的模板来源参数矩阵、优先级与互斥规则
- 冻结 `templateCandidateId` 作为唯一模板预填读取入口的语义
- 冻结 `Weekly Review` 中模板候选的默认 active candidate、单选切换与无候选空态规则
- 冻结模板预填至少覆盖哪些字段，才能算“预填闭环成立”
- 冻结创建成功后模板来源语义如何继续承接到 `Product Detail` 与 canonical `Product <-> Module Binding` 路径

DoD：

- `Template Reuse` 已明确是“预填辅助”而不是独立 CRUD 主线
- `Product Create` 仍是唯一正式创建承接位
- 模板候选来源与动作链已单值化
- 不在多个页面复制第二套产品创建路径
- 模板 handoff 已明确冻结为：`fromTemplateReuse + templateCandidateId + templateSource`
- 模板 handoff 与 `fromList / fromModuleDetail / fromDashboard` 的共存与互斥关系已单值化
- `Weekly Review` 的 active candidate 语义已单值化，不再允许“未选模板时退回另一套 generic focus”并列口径
- “预填成立”的最小字段级判定已明确，后续浏览器验收可机械执行
- 创建成功后如何继续承接模板组合已明确，不允许退化为“预填时看见、成功后丢失”

### phase09-03 冻结派生智能提示集、解释口径与动作承接

范围：

- 冻结 `capability gap / reuse opportunity` 的最小提示集合
- 冻结解释性指标、文案与动作指向的最小口径
- 冻结 `Weekly Review` 与 `Product Create` 各自承接的提示粒度
- 冻结提示如何与既有 `Decision / Product / Module / Review` 动作链对接
- 为每类提示冻结：`trigger -> explanation -> CTA -> target owner` 的单值矩阵
- 冻结 `capability_gap_hint` 对 active template candidate 的依赖关系与无 active candidate 时的成功空态语义
- 冻结没有稳定 canonical CTA 的提示不得进入 `phase09` 正式范围

DoD：

- 提示不是纯统计展示，也不是泛化 AI 输出
- 每类提示都能指向一个正式下一步动作或可解释空态
- `Weekly Review` 与 `Product Create` 的提示职责边界明确
- 不长出独立智能中心或第二套任务系统
- `reuse_opportunity_hint` 与 `capability_gap_hint` 的事实来源、触发条件与 CTA 已分别冻结
- `capability_gap_hint` 不再保留“实现时再决定是否回退到 review focus”的灰区
- 当前阶段不再保留“后面实现时再决定提示类型”的灰区

### phase09-04 冻结本阶段合同、读模型、owner 与候选数据源边界

范围：

- 冻结模板候选、组合快照、派生提示与 create 预填的最小正式合同
- 冻结当前 phase 必需的读模型与写模型边界
- 冻结 `phase06 reuse_summary`、`phase08 review`、`phase04 product create` 的正式消费边界
- 冻结前端 read owner / application owner、后端 query / command owner 的承接矩阵
- 冻结当前真实 caller / route / query owner / application owner inventory
- 冻结模板候选的 canonical 来源、去重键、排序规则与空态规则
- 冻结模板来源参数在 `Product Create -> Product Detail` 成功回流链中的保留方式
- 在本子任务中最终拍板：当前阶段是否允许引入轻量快照记录；未拍板前不得继续进入实现设计类子任务

DoD：

- `.proto` 继续作为唯一长期合同源
- 读写 owner 与正式 caller 清单足以直接进入 `/spec`
- 模板与提示不复制既有事实主线
- 当前真实 caller / owner 清单足以避免后续 `/spec` 漏掉正式消费面
- 模板候选的生成口径已单值化，不再允许后续 `/spec` 与实现各猜一套
- 创建成功后的模板来源复读链与 canonical binding 承接位已单值化
- 已明确冻结“默认读时派生”或“受控轻量快照记录”其一，且不能保留两套长期稳态

### 第二组：实现设计产出类子任务

### phase09-05 产出模板复用与派生提示的页面流、交互流与返回链设计

范围：

- 产出 `Weekly Review -> Template / Hint -> Product Create -> Product Detail` 的页面流
- 产出模板候选默认选中、单选切换、预填展示、编辑确认与取消返回的交互流
- 产出派生提示的最小展示方式、解释层级与动作 CTA
- 产出移动浏览器下的最小降级策略
- 产出模板空态、提示空态、预填失败态与回退策略
- 产出 `fromDashboard` 返回链与模板 handoff 并存时的用户可见行为
- 产出创建成功后 `Product Detail` 中模板来源摘要与 canonical binding CTA 的可见行为

DoD：

- 页面流与交互流足以直接进入实现
- 模板与提示的正式消费位、返回链与空态语义已明确
- `Weekly Review` 不与 `Product Create` 互相覆盖职责
- 不在多个页面重复长出第二套入口
- 模板选择、无候选空态与创建后承接行为可被浏览器验收逐步核对
- 空态、失败态与回退路径可直接被浏览器验收用例消费

### phase09-06 产出前端读写承接位与状态流设计

范围：

- 产出模板候选与派生提示的 read layer 切片落点
- 产出 `Product Create` 预填 handoff 的 application owner 设计
- 产出成功回流、query 失效与错误反馈策略
- 识别必须回收的页面级临时编排点
- 产出 caller 与 owner 的一对一映射表
- 产出 `Product Create` 页面如何只消费预填 read owner，而不侵入既有 canonical mutation owner
- 产出 `Weekly Review` 页面新增消费位与既有 `reuseSnapshot / representativeSignals` 的边界矩阵
- 产出 `Product Detail` 如何消费模板来源复读上下文，并继续导向既有 canonical binding path

DoD：

- `query` 与 `application` 边界明确
- `Product Create` canonical mutation owner 不被模板逻辑侵入替换
- caller 不会跨页面漂移成第二套 owner
- 设计结果足以指导后续源码实现
- `Product Detail` 的模板来源复读不会长出第二套详情写路径
- 前端不会因为模板预填而新增第二套 create form state 主线

### phase09-07 产出后端服务、合同与最小数据承接设计

范围：

- 产出模板候选 / 派生提示相关 proto 合同与 Connect 服务设计
- 产出后端 query owner、必要 command owner 与结果回流设计
- 产出与既有 `Review / Product / Decision / ReuseSummary` 服务的协作边界
- 产出是否需要轻量快照记录的评估与受控边界
- 产出本阶段关键路径的工具链、API 与浏览器验收清单
- 产出模板候选读取、模板预填详情读取、派生提示读取三类接口的正式责任边界
- 明确 `Review` 不是模板候选 canonical 事实源，只提供消费作用域与返回链元数据
- 若引入轻量快照记录，必须同时给出“为何纯读时派生不足”的明确证据；否则默认关闭该方案

DoD：

- 合同、服务与数据承接位单值化
- 不复制既有业务事实源
- 是否需要轻量快照记录已被明确为“受控支撑资产”或“无需新增记录”
- `buf / go build / frontend type-check / browser acceptance` 口径已冻结
- 模板候选、active candidate 与创建成功后模板来源复读链都已有单值合同语义
- 三类读取接口的职责边界足以直接驱动 API smoke 与前端 owner 实现

### 第三组：源码实现类子任务

### phase09-08 落实支撑能力相关合同、后端承接与前端 read owner

范围：

- 实现模板候选 / 派生提示相关 proto / Connect 合同
- 实现后端 query owner、前端 read owner 与必要数据承接
- 为后续模板 handoff、提示展示与 create 预填提供正式使能位

DoD：

- 支撑能力相关合同与 owner 单值一致
- 前后端已具备进入模板候选与提示消费的正式承接位
- 本阶段不以页面级临时拼装作为长期稳态

### phase09-09 落实模板候选选择、`Product Create` 预填与结果回流

范围：

- 实现模板候选列表与正式选择路径
- 实现从候选进入 `Product Create` 的预填 handoff
- 实现预填后编辑、创建成功回流与 reread
- 清理临时并列创建入口与散装导航逻辑
- 实现模板来源参数校验、非法组合参数回退与空候选回退
- 实现创建成功后 `Product Detail` 的模板来源摘要与 canonical binding CTA
- 实现预填字段的可见标记或说明，使浏览器验收可以明确判断“预填已生效”

DoD：

- 用户能够从既有复用事实进入正式预填创建路径
- `Product Create` 仍是唯一 canonical create 主线
- 创建成功后结果能回流到既有正式详情页与读模型
- 不保留并列临时创建主线
- 创建成功后模板来源语义不会丢失，且能继续导向既有 canonical `Product <-> Module Binding` 路径
- 非法模板参数、无效 `templateCandidateId` 与空候选场景不会把页面打成不可恢复错误

### phase09-10 落实派生提示展示、动作 handoff 与解释性回流

范围：

- 实现 `capability gap / reuse opportunity` 等最小提示
- 实现提示到 `Product Create / Product Detail / Decision / Module` 等正式动作链的 handoff
- 实现提示消费后的必要 reread、空态与错误语义
- 回收本阶段新增过程中产生的临时散装提示编排点
- 实现提示矩阵中的 `trigger / explanation / CTA / target owner` 四元组，不允许只落解释文案
- 对无 active template candidate 场景只返回成功空态，不再落未冻结的 generic focus fallback
- 对没有稳定 CTA 的候选提示直接裁撤，不进入实现态

DoD：

- 提示可重复消费并能指向正式下一步动作
- 提示不会退化为纯统计卡片或孤立文案
- 前后端不保留第二套长期智能主线
- `capability_gap_hint` 的 active candidate 依赖与空态语义已能被浏览器路径直接验证
- `reuse_opportunity_hint` 与 `capability_gap_hint` 都能被浏览器路径直接验证

### 第四组：验证验收类子任务

### phase09-11 完成模板复用与派生提示联调、浏览器验收与反回归验证

范围：

- 完成 `buf / frontend type-check / build / backend build` 工具链验证
- 完成模板候选、预填创建与提示 handoff 的 API smoke
- 完成 `Weekly Review -> Product Create 预填 -> Product Detail` 浏览器验收
- 完成 `Dashboard / Review / Product Detail / ReuseSummary` 的最小反回归验证
- 记录本阶段明确不做 `dry-run / AI Context Enhancement / Venture` 的边界证据
- 固定最小浏览器闭环步骤、字段级预填判定与成功/空态/失败态判定标准
- 固定最小 API smoke 矩阵：模板候选读取、模板预填详情读取、派生提示读取
- 固定反回归页面清单：`Weekly Review / Product Create / Product Detail / Dashboard / ReuseSummary`
- 固定最小 reread 观察断言：`Product Detail` 模板来源摘要、候选 `Module` 组合摘要、canonical binding CTA，以及 `Dashboard / Review / ReuseSummary` 成功返回语义

DoD：

- 工具链、API、浏览器与关键反回归均通过
- 模板预填真实缩短创建路径的证据已留档
- 派生提示真实支撑下一步动作的证据已留档
- 不把 `dry-run` 偷渡为本阶段既成事实
- 浏览器验收已能机械回答：
  - 是否出现模板候选或成功空态
  - 是否存在单值 active candidate 或成功空态
  - 是否出现派生提示或成功空态
  - 是否通过 `templateCandidateId` 进入了可编辑预填创建页
  - 是否创建成功并在 `Product Detail` 中看到模板来源摘要、候选 `Module` 组合摘要与 canonical binding CTA
  - 是否完成 `Dashboard / Review / ReuseSummary` 成功 reread，且未把“无统计变化”误判为失败

### 第五组：根级同步类子任务

### phase09-12 完成根级同步与后续 `dry-run` 进入条件回写

范围：

- 回写 `AGENTS.md`、`plan.md`、`docs/README.md`、`architecture_map.md`、`docs/phase/README.md`
- 同步 `phase09` 作为当前正式支撑能力 phase 的状态、入口与冻结结论
- 同步 `phase08` 退回最近完成正式业务 phase 的角色表达
- 只保留后续 `dry-run` 的进入条件表达，不提前扩大 `mvp0.4+` 范围

DoD：

- 根级真相源与 docs 入口单值一致
- `phase09` 的活动文档与正式验收入口可稳定找到
- `dry-run` 仍保持后续独立验收闸定位
- 不引入第二套根级状态口径

## 4. 本阶段明确不做

- 在本阶段直接执行真实项目 `dry-run`
- 在本阶段引入 `Venture / Decision Intelligence / AI Context Enhancement`
- 在本阶段建设完整模板平台或模板版本系统
- 在本阶段把 `Product Create` 改造成第二套模板宿主
- 在本阶段长出独立智能工作台、自动扫描或知识图谱
