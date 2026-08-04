# Personal Software Company OS

# Final Consensus & Execution Plan

**Author:** GPT54  
**Date:** 2026-08-04  
**Purpose:** 基于第一轮五份专家评审文档与第二轮三份交叉汇总文档，形成 PSCO 当前阶段的最终共识、仲裁结论与后续执行计划，作为后续 `PSCO_MVP_Spec_v0.1.md` 与工程推进的正式依据。

---

## 1. 文档定位

本文不是新增一轮评审意见，而是对两轮评审结果进行最终收口。

它要回答的不是“谁说得更好”，而是：

1. 哪些结论已经形成稳定共识；
2. 哪些分歧需要明确仲裁；
3. 仲裁后的最终方案应该如何推进；
4. 下一步应该产出什么规格文档、按什么顺序实现。

本文将作为 PSCO 从“理念设计阶段”进入“执行规格阶段”的桥接文档。

---

## 2. 证据来源

## 2.1 第一轮评审文档

- `PSCO_Evaluation-GPT54.md`
- `PSCO_Review_deepseek-v4-flash.md`
- `PSCO-Design-Review-GLM-52.md`
- `PSCO-evaluation-deepseek-v4-pro.md`
- `PSCO-review-qwen37-pro.md`

## 2.2 第二轮交叉汇总文档

- `PSCO-summarize-feedback-GPT54.md`
- `PSCO-summarize-feedback-dsv4flash.md`
- `PSCO-summarize-feedback-GLM52.md`

## 2.3 本文采用的仲裁原则

在两轮意见基础上，本文按以下标准做最终收口：

1. **优先保留高共识结论**  
   若同一结论在第一轮与第二轮均重复出现，视为稳定共识。

2. **优先保留低摩擦、高验证密度的方案**  
   若两个方案都成立，优先选择更能在 4-8 周内验证核心假设的方案。

3. **坚持“执行收缩”而不是“理论回退”**  
   可以缩减 v0.1 的执行对象，但不轻易重写整套领域语言。

4. **凡是会明显拖慢 MVP 的重规范、重集成、重自动化，默认后移**  
   只有当它直接影响核心验证闭环，才进入 P0。

---

## 3. 最终共识

经过两轮评审，以下内容已经可以视为 PSCO 当前阶段的最终共识。

## 3.1 战略共识

1. **PSCO 的方向成立。**
2. **AI 边界判断正确。**
3. **PSCO 不是代码管理工具，也不是 AI Chat 产品，而是个人软件公司的经营与资产系统。**
4. **当前最大问题不是理念不足，而是执行规格尚未收敛。**

## 3.2 MVP 共识

1. **v0.1 必须继续收敛。**
2. **v0.1 的核心不是“完整公司操作系统”，而是“资产沉淀闭环是否成立”。**
3. **`Decision` 必须保留在 MVP。**
4. **`Capability` 不应在 v0.1 作为重实体维护。**
5. **Rust Intelligence Layer 不应进入 v0.1 主范围。**
6. **冷启动、导入路径、录入摩擦、基础度量，是 MVP 的硬问题。**

## 3.3 工程共识

1. **先定义核心写入动作，再定义页面。**
2. **需要跨文档一致性收口。**
3. **需要最小模块准入规则。**
4. **需要最小 Decision 模板。**
5. **需要最小数据导出/备份策略。**

---

## 4. 最终仲裁结论

以下是基于两轮讨论后的正式仲裁结果。后续规格与实现应以此为准。

## 4.1 仲裁 1：v0.1 的产品目标

**最终结论：**

> v0.1 的正式目标定义为：  
> **一个面向个人开发者的软件资产登记、决策留痕与基础复用反馈系统。**

**理由：**

- 这保留了 PSCO 的差异化核心；
- 同时避免把 v0.1 做成“完整公司操作系统”；
- 比“只做模块目录工具”更完整；
- 比“全面覆盖 Opportunity / Experiment / AI / Composition”更克制。

## 4.2 仲裁 2：v0.1 核心实体范围

**最终结论：**

v0.1 保留以下核心实体：

1. `Product`
2. `Module`
3. `Release`
4. `Decision`
5. `Repository`
6. `Venture`（可选实体，不强制）

延后或弱化：

- `Opportunity`
- `Feature`
- `Experiment`
- `Capability`

**理由：**

- 该方案兼顾了 GPT54 / DS-Flash / GLM52 的稳健收敛；
- 避免了 DS-Pro 的过度压缩；
- 也避免了 Qwen 方案中过度削弱 PSCO 差异化核心。

## 4.3 仲裁 3：`Decision` 是否进入 MVP

**最终结论：进入。**

**理由：**

- `Decision` 是 PSCO 长期护城河的一部分；
- 它是 AI 上下文的重要原料；
- 没有它，PSCO 容易退化成模块台账工具；
- 录入成本问题应通过低摩擦模板解决，而不是把对象整体移出 MVP。

## 4.4 仲裁 4：`Capability` 的定位

**最终结论：**

> `Capability` 在 v0.1 中作为派生结果层，不单独建核心表，不要求用户主动 CRUD。

建议的实现方式：

- 先作为模块聚合结果展示；
- 可先保留轻量 `capability_tag` 或分类辅助字段；
- 基于复用次数、稳定状态、版本演进、关联决策密度等事实生成视图。

**理由：**

- 符合“能力来自事实沉淀，而不是手工填写”的共识；
- 与低摩擦 UX 原则一致；
- 保住了长期哲学，不增加短期输入负担。

## 4.5 仲裁 5：`Venture` 是否删除或合并

**最终结论：**

> `Venture` 保留为可选上层分组实体，不强制创建，不与 `Product` 合并，也不重写为新术语。

**理由：**

- 保留长期经营层语义；
- 避免重写整套领域语言；
- 通过“可选而非强制”解决 solo 场景负担；
- 比“直接合并”更稳，也更有延展性。

## 4.6 仲裁 6：数据库与 Local First

**最终结论：**

> 当前默认技术基线维持 `PostgreSQL`，不在这一阶段改成 `SQLite`。  
> 但必须补充“可导出、可备份、可迁移”的产品要求。

**理由：**

- 原始技术基线已经明确为 `React + Go + PostgreSQL`；
- 此时切换 SQLite 会引入额外分叉；
- Local First 的核心是数据所有权，而不是必须单机 SQLite；
- 同时保留后续提供“轻量部署模式”的开放性。

## 4.7 仲裁 7：GitHub 集成的优先级

**最终结论：**

> v0.1 必须支持 `Repository` 概念与手动绑定；  
> `GitHub OAuth + 自动导入` 不作为 P0 阻断项，后移到 P1/P2。

**理由：**

- 资产登记闭环先于自动化集成；
- 手动绑定足以完成第一轮验证；
- 避免 OAuth、token、API 限流等额外复杂度挤占 MVP 时间。

## 4.8 仲裁 8：PMM / PCP 的定位

**最终结论：**

> 需要最小版本的 PMM 与 PCP 思路，但不将“完整规范”作为当前阶段的阻断项。

更具体地说：

- **PMM**：v0.1 先做轻量模块契约字段，不做重型规范化；
- **PCP**：v0.1 先保证 `Product / Module / Decision` 可结构化检索，不要求完整协议完备。

**理由：**

- GLM52 对工程契约的提醒是正确的；
- 但完整协议前置会过度工程；
- 当前更需要“可演化的最小契约”，而不是“一步到位的正式标准”。

## 4.9 仲裁 9：模板级复用是否进入当前计划

**最终结论：**

> 模板级复用是重要方向，但不作为进入 MVP Spec 的 P0 阻断项。  
> 它应被列为 v0.1 后段验证项或 v0.2 的第一优先级。

**理由：**

- DS-Flash 关于“没有最小复用感知，飞轮难启动”的提醒是有价值的；
- 但当前更基础的闭环仍是“登记 - 绑定 - 决策 - 反馈”；
- 因此它不应阻断规格收敛，但必须被明确列入后续路线图，而不是丢失。

---

## 5. 最终 MVP 基线

## 5.1 MVP 页面范围

v0.1 建议保留以下页面或主功能块：

1. **Dashboard**
2. **Module Registry**
3. **Product Registry**
4. **Decision Center**
5. **Project/Repository Binding**

其中：

- Dashboard 只做最小聚合，不做复杂驾驶舱；
- AI Assistant 不作为一级独立导航；
- Opportunity / Experiment / Feature 页面不进入 v0.1 主体范围。

## 5.2 MVP 核心动作

后续规格文档应先定义以下动作：

- `CreateProduct`
- `CreateModule`
- `CreateRelease`
- `CreateRepository`
- `RecordDecision`
- `BindRepositoryToProduct`
- `BindModuleToProduct`
- `MapModuleToRepository`
- `LinkDecisionToTarget`

可选补充动作：

- `CreateVenture`
- `BindProductToVenture`

## 5.3 MVP 最小数据模型

**核心表：**

- `ventures`（可选）
- `products`
- `repositories`
- `modules`
- `module_releases`
- `decisions`

**关系表：**

- `product_modules`
- `product_repositories`
- `decision_links`

**派生视图：**

- `capability_summary`
- `module_reuse_summary`
- `product_asset_coverage`
- `pending_decision_signals`

## 5.4 MVP 非目标

v0.1 明确不做：

- 自动扫描全部代码
- 自动知识图谱
- Rust 智能层
- AI 自动判断方案
- 完整 Opportunity / Experiment 流程化
- GitHub OAuth 自动导入
- 完整 PMM / PCP 正式标准
- 独立 AI Assistant 工作台

---

## 6. 规格收敛计划

在进入代码实现前，建议按以下顺序推进。

## 6.1 Phase 0：一致性收口

目标：解决“文档语言已完整，但术语和边界仍有断层”的问题。

需要完成：

1. 校正文档中 `Knowledge` 等未定义概念；
2. 统一状态命名与术语；
3. 明确 `Venture / Product / Module / Decision / Capability` 的边界；
4. 明确哪些概念保留在理论层，哪些进入 v0.1 执行层。

**产出物：**

- 一份轻量术语表
- 一份 v0.1 实体边界说明

## 6.2 Phase 1：编写 `PSCO_MVP_Spec_v0.1.md`

这是下一份关键文档，必须回答：

1. v0.1 保留哪些对象；
2. 每个对象最小字段是什么；
3. 页面有哪些；
4. 动作有哪些；
5. 表结构有哪些；
6. API 有哪些；
7. 哪些内容明确不做；
8. 每个功能的 Done 标准是什么。

**额外必须包含：**

- 冷启动流程
- 空状态设计
- 导入路径说明
- 基础度量指标
- 导出/备份要求

## 6.3 Phase 2：执行顺序

建议的实现顺序如下：

1. **Module Registry**
2. **Decision Center**
3. **Product Registry**
4. **Repository / Product / Module Binding**
5. **Dashboard**

这个顺序的理由是：

- `Module` 是资产基础；
- `Decision` 是上下文基础；
- `Product + Binding` 把资产与现实项目连接起来；
- Dashboard 最后做聚合展示。

## 6.4 Phase 3：验证与复盘

建议使用一个真实项目做 dry-run，优先以 `Rento-miniX` 为样本。

最小验证流程：

1. 创建 Product；
2. 绑定 Repository；
3. 注册 2-3 个 Module；
4. 录入关键 Decision；
5. 观察 Dashboard 是否能形成第一版资产反馈。

若该流程跑不顺，优先修规格与 UX，不扩功能。

---

## 7. 最终度量与验收

## 7.1 建议的 MVP 先导指标

建议至少埋以下 4 个指标：

1. **模块复用率**  
   被 2+ Product 使用的 Module 占比

2. **决策复用率**  
   新建相似决策时引用历史 Decision 的比例

3. **资产导入耗时**  
   从创建 Product 到完成首轮 Module 绑定的时间

4. **录入摩擦感知**  
   每个核心对象完成最小登记所需步骤数/字段数

## 7.2 v0.1 Done 标准

只有当以下条件成立时，才可认为 v0.1 规格成立并可进入稳定实现：

1. 核心实体范围明确；
2. 冷启动路径明确；
3. 模块准入规则明确；
4. Decision 模板明确；
5. 页面与动作一一对应；
6. 表结构与关系表明确；
7. 非目标列表明确；
8. 导出/备份要求明确；
9. 至少有一个真实项目 dry-run 走通。

---

## 8. 最终结论

基于两轮评审与三份第二轮汇总文档，PSCO 当前阶段的最终结论可以明确为：

> **PSCO 的方向已经足够清楚。现在真正需要做的，不是继续扩展概念，而是把“资产、决策、绑定、反馈”压缩成一个最小但可运行的系统，并用真实项目验证它是否产生资产化收益。**

因此，后续推进应遵循以下总原则：

1. **不退回概念层重写理论；**
2. **不在 MVP 前引入重型自动化与重型规范；**
3. **保留 PSCO 的差异化核心：Module + Decision + Binding + Feedback；**
4. **优先解决冷启动、录入摩擦、导入路径和基础度量；**
5. **以 `PSCO_MVP_Spec_v0.1.md` 作为下一阶段唯一主目标。**

一句话收口：

> **Decision 不砍，Capability 不建表，Rust 不进 MVP，Venture 保留但不强制，GitHub 自动化后移，MVP 先把资产闭环跑通。**

---

**Final Record**  
**Signed by: GPT54**
