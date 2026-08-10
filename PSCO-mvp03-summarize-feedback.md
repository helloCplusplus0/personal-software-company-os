# Personal Software Company OS

# MVP0.3 Final Consensus & Planning Baseline

**Author:** GPT54
**Date:** 2026-08-10
**Purpose:** 基于两轮 `mvp0.3` 专家评审与交叉汇总文档，形成 PSCO 在 `phase06` 收口之后、进入下一阶段正式 `/plan` 之前的最终仲裁结论、范围基线、推进顺序与明确非目标。

---

## 1. 文档定位

本文不是第九份评审意见，也不是直接替代下一阶段正式 `phase` `/plan` 的执行文档。

本文要解决的是五件事：

1. 明确哪些判断已经形成稳定共识；
2. 对存在分歧的方向给出最终仲裁；
3. 冻结 `mvp0.3` 的正式方向、范围与非目标；
4. 约束下一阶段正式 `/plan` 的输入边界与推进顺序；
5. 为后续 `/plan -> /spec -> 实现 -> 验收 -> 收口` 提供唯一上游判断基线。

本文的职责，是为下一阶段正式 phase 入口建立提供上游依据，而不是在这里越权冻结正式 `phase` 名称、`.trae/specs/phaseXX_*` 路径、接口名或实现细节。

## 1.1 本文与 `PSCO-mvp02-summarize-feedback.md` 的关系

本文不是把 `mvp0.2` 重新写一遍，也不是把 `mvp0.2` 尚未完成的阶段二机械复制成另一份文档。

本文与 `PSCO-mvp02-summarize-feedback.md` 的关系，明确冻结为：

1. `PSCO-mvp02-summarize-feedback.md` 已经完成了从 Asset Registry 走向 Operating System 的总方向仲裁，并给出了“阶段一 / 阶段二”的候选结构；
2. `phase06` 已经实际完成了其中的阶段一：`Onboarding + Data Sovereignty + Reuse Awareness`；
3. 本文要解决的，不是再次仲裁 `phase06` 做得对不对，而是把 `mvp0.2` 尚未落地的阶段二，正式收敛为 `mvp0.3` 的执行主题、范围基线与完整路线表达；
4. 因此，本文既有承接 `mvp0.2` 的部分，也有新增冻结的部分；新增冻结的重点，不在于发明更多长期对象，而在于把“Review Loop / Template Reuse / Derived Intelligence / Dry-Run”从候选方向升级为下一步必须完成的正式主线。

换句话说：

> **`mvp0.2` 解决的是“总方向与阶段一 / 阶段二如何划分”；`mvp0.3` 解决的是“在阶段一已经完成之后，阶段二应如何被完整地、一次性地规划为下一段正式路线”。**

---

## 2. 证据来源

## 2.1 第一轮方向评审

- `docs/review/PSCO-mvp03-DPv4flash.md`
- `docs/review/PSCO-mvp03-DPv4pro.md`
- `docs/review/PSCO-mvp03-GLM52.md`
- `docs/review/PSCO-mvp03-GPT54.md`
- `docs/review/PSCO-mvp03-qwen37pro.md`

## 2.2 第二轮交叉汇总

- `docs/review/PSCO-mvp03-summarize-feedback-DPv4flash.md`
- `docs/review/PSCO-mvp03-summarize-feedback-GLM52.md`
- `docs/review/PSCO-mvp03-summarize-feedback-GPT54.md`

## 2.3 本文采用的仲裁原则

1. **高共识优先于单点激进扩张。** 四方以上独立收敛的判断，优先视为稳定输入。
2. **时机优先于愿景。** 某个长期方向即使正确，只要当前前提未成立，就不能提前升格为下一阶段主线。
3. **先消费既有交付，再扩新对象宽度。** `phase05 / phase06` 已交付能力必须先进入真实经营动作，否则会沦为展示型孤儿能力。
4. **使用频率优先于概念完整。** 下一步先证明用户会持续回来使用，再讨论更大的长期模型。
5. **评审不越权冻结正式 phase。** 正式阶段命名、spec 路径与具体落点留给后续 `/plan` 与 `/spec` 决定。

---

## 3. 最终共识

经过两轮评审，以下内容已经形成稳定共识，可直接视为 `mvp0.3` 的正式前提。

## 3.1 关于当前项目状态的共识

1. `phase01 ~ phase06` 已经证明 PSCO 的最小资产主线成立。
2. 当前系统已经具备：资产登记、决策留痕、绑定关系、Dashboard 反馈、Onboarding、导出 / 备份与基础复用感知。
3. 当前系统仍偏“登记 + 展示 + 首轮引导”，还没有真正证明“用户会围绕真实项目持续 review、持续决策、持续复用”。
4. 下一步的关键问题，已经不是“还能不能再加功能”，而是“已交付能力能不能被持续消费并形成经营节奏”。

## 3.2 关于下一步总方向的共识

1. **Operating Review Loop** 必须成为下一步中心任务。
2. Dashboard 必须从“总览页”继续推进为“经营动作起点”。
3. `Feedback -> Decision -> Update` 必须进入下一步正式主线。
4. 模板级复用必须进入下一步，但边界继续保持最小版。
5. 真实项目 `dry-run` 必须成为独立验收要求，而不是顺手补充说明。
6. AI 仍然不应成为下一步主线。

## 3.3 关于范围边界的共识

1. 不新增长期核心实体主线。
2. `Capability` 继续作为派生层，不进入重实体 CRUD。
3. `Decision` 必须继续保持经营中心地位，不能被弱化成普通备注数据。
4. 不做 GitHub OAuth / 自动导入、不做 AI 一级工作台、不做自动扫描 / 知识图谱、不做 Rust Intelligence Layer。
5. 技术栈、`.proto` 单一合同源、前端 query / application 边界等既有工程约束继续冻结。
6. 在正式进入 `mvp0.3` 业务主线之前，必须先解决 `audit_001` 指出的 Go 业务传输主线问题；该问题不得继续留在“边做业务边顺手收敛”的状态。

---

## 4. 最终仲裁结论

## 4.1 `mvp0.3` 的版本语义

**最终结论：**

> `mvp0.3` 的实质任务，不是跳过 `mvp0.2` 尚未完成的阶段二去开启一套更大的长期模型；
> 而是把 `mvp0.2` 尚未落地的“Operating Review Loop + 模板级复用 + 派生智能深化 + 真实项目 dry-run”真正做完，并将其作为 PSCO 从“登记系统”走向“经营系统”的正式跨越。

这意味着我**不采纳**把 `mvp0.3` 建立在“先默认 `phase07` 已经完成”这一前提上的推进方式。

原因很简单：

1. `phase06` 只完成了 `mvp0.2` 的阶段一；
2. 阶段二仍是仓库现实里尚未闭合的核心能力；
3. 在这一步未完成前，直接上跳到 `Venture / Decision Intelligence / AI Context Enhancement`，会把系统带回“概念扩张先于使用闭环”的旧问题。

进一步冻结为：

> **`mvp0.3` 不是对 `mvp0.2` 的否定，也不是平行新路线；它是 `mvp0.2` 阶段二在 `phase06` 完成之后的正式执行化、完整化与收口化。**

## 4.2 `mvp0.3` 的总主题

**最终结论：**

> `mvp0.3` 的总主题确定为：
> **从 Reuse-Aware Registry 走向 Operating Review System。**

它的正式主轴不是三条并列新世界，也不是提前切到战略与智能主线，而是：

1. **Operating Review Loop**（中心主线）
2. **Template Reuse**（复用执行支撑）
3. **Derived Intelligence Deepening**（复利行动支撑）
4. **Real-Project Dry-Run**（独立验收闸）

## 4.3 对“三并列主线”与“一主两翼一闸”的仲裁

**最终结论：**

> 我采纳 **“Operating Review Loop 为中心主线，Template Reuse 与 Derived Intelligence Deepening 为支撑，Real-Project Dry-Run 为验收闸”** 的结构。

原因如下：

1. Review Loop 是消费 `phase05` Feedback、`phase06` Reuse Awareness 与既有 `Decision` 主线的唯一经营枢纽；
2. Template Reuse 脱离 review，容易退化为孤立模板功能；
3. 派生智能深化脱离 review，容易退化为“看起来更丰富的统计展示”；
4. `dry-run` 的价值不在证明功能能跑，而在证明 review loop 是否真的改变了真实项目中的下一步动作。

因此我**不采纳**把所有主题完全并列推进的结构，也**不采纳**把模板 / 派生智能 / `dry-run` 降为可有可无的尾部事项。

## 4.4 Dashboard 的下一阶段定位

**最终结论：**

> Dashboard 不是继续增加摘要卡片，而是必须成为 daily / weekly operating cycle 的起点。

下一阶段的 Dashboard 必须至少承接：

1. daily review 入口；
2. weekly review 入口；
3. 当前焦点、代表性反馈与待处理决策的统一承接；
4. 从 review 进入 `Decision`、再回流更新到既有实体的动作链。

换句话说，Dashboard 必须从“看状态的地方”变成“推动下一步动作的地方”。

## 4.5 `Decision` 的下一阶段定位

**最终结论：**

> 可以后移的是 `Decision` 的高级智能复用；不能后移的是 `Decision` 在经营回路中的中心地位。

`mvp0.3` 必须推进的是：

- review 中对 `Decision` 的承接；
- `Feedback -> Decision -> Update` 闭环；
- 低摩擦的 decision action handoff 与实体回流。

后移到 `mvp0.4+` 的才是：

- 相似决策匹配；
- 历史决策引用推荐；
- 更重的 decision intelligence；
- AI 驱动的上下文增强。

## 4.6 模板级复用的下一阶段定位

**最终结论：**

> 模板级复用必须进入 `mvp0.3`，但其边界继续冻结为：
> **“Module 组合快照 + 新建预填辅助”，不是新一级核心实体，不做完整模板平台。**

模板的职责是：

1. 把已有复用事实转化成下一次创造的低摩擦起点；
2. 让复利从“可见”推进到“可用”；
3. 服务于 review 之后的下一步动作，而不是独立长成另一套系统。

## 4.7 派生智能深化的下一阶段定位

**最终结论：**

> 派生智能深化必须进入 `mvp0.3`，但它不是 AI 主线，而是 review loop 的行动支撑层。

它至少应承接：

- 能力缺口提示；
- 复用机会提示；
- 能力演化反馈；
- 与 review 动作相连的最小解释性指标。

它的目标不是“看起来更聪明”，而是让用户在 review 时更容易判断“下一步应该复用什么、补齐什么、推进什么”。

## 4.8 真实项目 `dry-run` 的定位

**最终结论：**

> 真实项目 `dry-run` 不是可有可无的附录，而应作为 `mvp0.3` 后段验收中的独立交付要求。

它不能替代 fixture 验收，但必须补上三件事：

1. 真实使用摩擦是否可接受；
2. review loop 是否真的改变下一步动作；
3. 模板复用是否真的缩短了下一次创造路径。

## 4.9 当前明确后移的方向

以下方向在本文件中明确**后移到 `mvp0.4+` 候选范围**：

1. `Venture` 由“可选概念”升格为下一阶段正式主线；
2. `Decision Intelligence` 作为独立主线进入当前版本；
3. `AI Context Enhancement` 进入当前版本主线；
4. `Opportunity / Feature / Experiment` 流程化；
5. 完整模板平台、参数化模板体系或模板版本管理；
6. GitHub OAuth / 自动导入；
7. 自动扫描 / 知识图谱 / 更重的 AI 工作台。

这里的判断不是否定这些方向的长期价值，而是冻结它们的**时机**：当前前提还未成熟。

## 4.10 当前不允许提前冻结的内容

以下内容在本文件中明确**不冻结**，留给后续正式 `/plan` 与 `/spec`：

1. 下一阶段正式 `phase` 名称；
2. `.trae/specs/phaseXX_*` 路径；
3. review、模板与指标的最终接口名；
4. 模板的具体存储方式；
5. `dry-run` 最终选用的真实项目对象；
6. 是否将若干子能力拆成一个还是多个交付型 phase。

---

## 5. `mvp0.3` 正式范围基线

## 5.1 必做范围

下一阶段正式 `/plan` 必须覆盖以下四组范围：

1. **Operating Review Loop**
   - daily / weekly review 最小入口
   - 当前焦点、代表性信号、待处理决策的统一承接
   - `Feedback -> Decision -> Update` 闭环
   - review 结论回流既有实体

2. **Template Reuse**
   - `Module` 组合快照
   - 新建 `Product` 时基于组合快照预填
   - 预填后继续编辑并完成创建

3. **Derived Intelligence Deepening**
   - capability gap / reuse opportunity 的最小提示
   - `module_reuse_summary / capability_summary` 的行动化消费
   - 解释性指标最小落地

4. **Real-Project Dry-Run**
   - 至少一个真实项目走通完整闭环
   - 独立留存验收记录
   - 明确记录摩擦点、收益点与后续修正项

## 5.2 明确不做范围

`mvp0.3` 明确不做：

- `Venture` 正式主线化
- `Opportunity / Feature / Experiment` 流程化
- `Capability` 重实体 CRUD
- GitHub OAuth / 自动导入
- AI 一级工作台
- 自动扫描 / 知识图谱
- Rust Intelligence Layer
- 完整模板平台
- `Decision` 高级智能复用机制

---

## 6. 对后续 `/plan` 的正式输入

## 6.1 推荐的完整候选阶段结构

本文不冻结正式 `phase` 命名，但冻结如下**完整后续路线预览**。

这里的“阶段一 / 二 / 三”是根级 `plan.md` 应表达的候选阶段结构，不等同于已经建立的正式 `phase07 / phase08 / phase09`。它们的职责，是一次性把后续路线说完整，而不是提前把未建立阶段写成既成事实。

在进入下面三段 `mvp0.3` 业务路线之前，后续正式 `/plan` 必须先承接 `audit_001` 给出的前置结论：

- 先建立一个独立交付型 phase，完成 Go 业务传输主线从 `chi + JSON HTTP` 过渡实现向 `.proto + ConnectRPC` 正式主线的切换
- 该前置 phase 的完成标准，不是“以后新增业务接口默认走 ConnectRPC”，而是 `phase01 ~ phase06` 已落地 canonical 业务接口已经完成正式切换
- `healthz / readyz / metrics / debug` 等非业务端点继续保留在 `chi + net/http`
- 旧 hand-written JSON 业务主线只允许在迁移过程短时存在，不允许作为该前置 phase 收口后的长期稳态

### 候选阶段一：Operating Review Loop 主线建立

目标：

- 让 Dashboard 真正接管经营动作起点；
- 让 `Feedback -> Decision -> Update` 首次形成稳定回路；
- 让 review 从展示层进入动作层。

至少应承接：

- daily / weekly review 入口
- review 结论记录
- 回流到既有实体的 action handoff
- 对 phase05 / phase06 既有数据的正式消费

这一阶段的核心验收问题不是“有没有新增页面”，而是：

- 用户是否真的从 Dashboard 进入经营动作；
- `Decision` 是否被 review loop 真正消费；
- `Feedback -> Decision -> Update` 是否第一次稳定闭合。

### 候选阶段二：Template Reuse + Derived Intelligence Deepening

目标：

- 让复用从“可见”走向“可用”；
- 让派生智能从“统计展示”走向“动作支撑”。

至少应承接：

- `Module` 组合快照
- 新建 `Product` 预填
- capability gap / reuse opportunity 提示
- 解释性指标落地

这一阶段的核心验收问题不是“是不是多了更多统计”，而是：

- review 后是否真的更容易开始下一次创造；
- 复用提示是否真正支撑了动作；
- 模板预填是否带来了可感知的进入成本下降。

### 候选阶段三：Real-Project Dry-Run 与收口验证

目标：

- 用真实项目证明 `mvp0.3` 核心命题成立；
- 留存独立验收记录，形成下一阶段仲裁输入。

至少应承接：

- 至少一个真实项目完整走通
- 明确记录摩擦、收益与修正项
- 与 fixture 验收并列留档

这一阶段的核心验收问题不是“系统能不能跑”，而是：

- 真实项目里，用户是否愿意反复回来使用；
- review loop 是否真的改变了下一步动作；
- 模板与派生智能是否真的带来了复利加速。

## 6.2 推荐的先后关系

后续正式 `/plan` 必须体现以下依赖关系：

1. 没有候选阶段一，Dashboard 仍只是总览页，`Decision` 也还不是经营枢纽；
2. 没有候选阶段二，`phase06` 的复用感知仍停留在“可见”而非“可用”；
3. 没有候选阶段三，就无法判断这套经营闭环在真实项目上是否成立；
4. 只有三段都成立，`mvp0.3` 才算真正完成了从“可见复利”到“可运行经营系统”的跨越。

## 6.3 对正式 `/spec` 的硬约束

后续 `/spec` 与实现必须继续遵守：

1. `.proto` 仍是唯一长期合同源；
2. query 层继续保持纯只读；
3. 前端写路径继续收敛到切片固定承接位；
4. 不引入第二套路由、第二套状态管理、第二套 UI 事实源；
5. review loop 不能演化为通用任务管理器；
6. 模板不能演化为独立平台型系统。

## 6.4 对后续验收的硬要求

下一阶段正式验收至少必须回答以下问题：

1. 用户是否能从 Dashboard 进入 review 并完成回流？
2. `Decision` 是否真正成为 review 中的经营中心？
3. 模板预填是否真的缩短了新建路径？
4. 派生智能提示是否真正帮助了下一步动作？
5. 真实项目中，用户是否愿意围绕 PSCO 持续回来使用？

---

## 7. 最终结论

`phase06` 的意义，不只是又完成了一个阶段，而是把 PSCO 从“资产登记 + 复用可见”推进到了“可以开始承接经营动作”的门口。

因此，`mvp0.3` 的最终仲裁结论是：

> **下一步不应继续扩长期对象宽度，也不应提前切到 Venture / Decision Intelligence / AI Context Enhancement；而应优先完成 `mvp0.2` 尚未落地的阶段二，把 PSCO 从“已能进入、已能带走、已能看见复用”推进到“会围绕真实项目持续 review、持续决策、持续复用”的 Operating Review System。**

正式收敛方向如下：

1. **Operating Review Loop** 为中心主线；
2. **Template Reuse** 与 **Derived Intelligence Deepening** 为支撑能力；
3. **Real-Project Dry-Run** 为独立验收闸；
4. `Venture / Decision Intelligence / AI Context Enhancement` 后移到 `mvp0.4+` 候选范围。

如果用一句话收口：

> **`phase01 ~ phase06` 证明了 PSCO “能登记、能带走、能看见复利”；`mvp0.3` 必须继续证明它“会被周期性经营，而且经营会加速下一次创造”。**
