# PSCO MVP0.4 专家评审意见

**Author:** Gemini-3.1-Pro
**Date:** 2026-08-13
**Target:** `PSCO-real-project-dry-run-user-manual-GPT54 feedback.md`
**Purpose:** 基于用户在 Real-Project Dry-Run 中的真实体感与深度战略反思，评估 MVP0.3 的实际交付质量，并为下一阶段（MVP0.4+）的系统演进提供标准化的架构与产品方向建议。

---

## 1. 核心洞察与总体评价

用户的真实反馈极具穿透力。MVP0.3 成功证明了 PSCO 在“资产关联、来源追踪、留痕管理”上的工程扎实度（并非粗糙的玩具），但也精准暴露了当前架构在**“用户心智模型”**与**“现代 AI 研发工作流”**上的严重脱节。

用户的第五至第八点反思，触及了 PSCO 的灵魂：
> **PSCO 究竟是一个“需要人工录入的文本管理后台”，还是一个“深度融合 Agent 工作流、驱动个人软件公司运转的现代 OS”？**

结论很明确：当前的 PSCO 停留在前者，而它的未来必须是后者。

---

## 2. 现有系统缺陷的修复定调 (Fix & Improvement)

在进入宏大的 MVP0.4 之前，Dry-Run 中暴露的工程断点必须通过 `fix` 流程迅速收口，否则会持续产生高摩擦。

### 2.1 Onboarding 流程的空态阻断与逻辑割裂
*   **缺陷本质**：底层实体的解耦（Module/Product/Repository 互相独立）被错误地暴露给了业务向导层。Onboarding 应该是“上下文继承”的组合器，而不是批量新增孤立数据的批处理脚本。
*   **修复建议**：
    1. 修复空态下 `/onboarding` 首步无响应的阻断 Bug。
    2. 重构 Onboarding 状态机：在向导中创建的 Product、Module、Repository，必须在向导结束时**自动建立绑定关系**，而不是让用户去详情页手动关联。

### 2.2 Decision 状态流转未闭环
*   **缺陷本质**：Decision 目前只有“被提出”和“被关联”，缺乏明确的“已解决/已执行 (Resolved/Actioned)” 状态。
*   **修复建议**：为 Decision 增加明确的生命周期状态。当用户在详情页完成动作后，允许标记该决策为“已处理”，从而使其从 Dashboard 和 Daily Review 的 `Current Focus` 中正确退出，完成经营回路的物理闭环。

---

## 3. MVP0.4 战略路线建议：向“智能经营 OS”跨越

基于用户的反思，MVP0.4 不应再继续堆砌传统的 CRUD 表单，而应围绕**“Agent 协同”**与**“精益经营”**进行架构范式转移。我建议 MVP0.4 聚焦以下三大主线：

### 3.1 主线一：Agent-Driven Workflow (读写倒置)
**问题背景**：用户感觉自己只是个“文本描述者”，真正的代码是在 IDE 里由 Agent (Trae/Cursor) 编写的。让人去手动维护 PSCO 状态是反人性的。
**演进方向**：
*   **Web 侧重读与宏观决策**：PSCO Web 端退后一步，成为大盘看板、趋势分析和高管（用户本人）下达指令的驾驶舱。
*   **Agent 侧重写与微观执行**：为 PSCO 提供标准化接口（如本地 REST API 或专门的 MCP Server）。
*   **场景示例**：当 Trae Agent 在本地新建了一个微服务代码目录，它应通过 MCP 自动调用 PSCO 接口，静默完成 Module 的登记与绑定；当代码中遇到技术难点，Agent 自动向 PSCO 写入一条 Pending Decision。

### 3.2 主线二：跨项目全局上下文资产化 (Context as an Asset)
**问题背景**：新开项目时，极具价值的全局规范、架构约束需要手动复制粘贴，易错且繁琐，不利于 Agent 保持执行一致性。
**演进方向**：
*   将 `project_rules`, `TECH_STACK_BASELINE`, `global_skills` 等提权为 PSCO 管理的**核心资产 (Global Contexts)**。
*   **下发机制**：在 PSCO 中执行“基于模板创建 Product”或“绑定 Repository”时，PSCO 后端自动向目标仓库注入 `.trae/rules` 文件夹或 `.cursorrules` 文件。
*   **价值**：PSCO 不仅“记录”资产，更“下发”资产，真正成为多项目 Agent 开发的**母舰 (Mothership)**。

### 3.3 主线三：融入精益创业模型 (Lean Business Loop)
**问题背景**：Personal Software Company 包含“公司”二字，不仅是写代码，还要验证市场、找用户。
**演进方向**：
*   解锁并升格 `Venture / Experiment / Opportunity` 实体。
*   参考《Testing Business Ideas》与《Value Proposition Design》，将这些实体的数据结构改造为：`假设 (Hypothesis) -> 实验设计 (Experiment) -> 数据指标 (Metrics) -> 验证结果 (Validation)`。
*   **业务数据挂载**：为 Product 实体预留外部业务指标（如 Stripe 支付数据、Google Analytics 流量）的接入点。让 Weekly Review 不仅 review 代码进度，更 review 商业指标。

---

## 4. 架构与设计的硬性约束建议

为了确保 MVP0.4 不会因为愿景宏大而失控，建议设定以下约束：
1. **Local-First 延续**：Agent 协同机制必须基于本地网络（如 `localhost:xxxx` API 或本地 MCP），不得强依赖公网网关，保护数据主权。
2. **渐进式精益**：精益模型的引入不要一上来就做极其复杂的画布，先做一个最简单的 `Experiment` 实体，只包含“假设”和“结果”两个字段，跑通 `Dashboard -> 实验验证 -> 记录结论` 的闭环。
3. **隔离修复与重构**：在正式进入 MVP0.4 的 Phase 之前，必须先用一个专门的 `fix` 工作流，把当前 Onboarding 和 Decision 未闭环的问题彻底清理干净。

---

## 5. 总结

MVP0.3 是一次极其成功的“压力测试”。它证明了系统底盘稳固，但也证明了**“纯人工维护的资产登记系统，在单人+AI的研发模式下是没有未来的”**。

PSCO 下一步的伟大之处，不在于它能提供多少个精美的表单，而在于它能**多大程度上与 Agent (Trae/Cursor) 融为一体**，以及它能**多大程度上用数据证明你的商业假设**。MVP0.4，应该是 PSCO 真正蜕变为 **“OS”** 的起点。
