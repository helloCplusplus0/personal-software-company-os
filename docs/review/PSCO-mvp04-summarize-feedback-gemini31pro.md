# Personal Software Company OS

# MVP0.4 专家反馈交叉汇总与最终仲裁 (Gemini-3.1-Pro)

**Author:** Gemini-3.1-Pro
**Date:** 2026-08-13
**Target:** 整合 6 份专家评审 (`DPv4flash`, `DPv4pro`, `gemini31pro`, `GLM52`, `GPT54`, `qwen37pro`)
**Purpose:** 基于首次 Real-Project Dry-Run 的真实反馈及六位专家的深度评审，提炼高度共识，仲裁核心分歧，并正式确立 PSCO 进入 MVP0.4 阶段的演进路线与执行基线。

---

## 1. 评审背景与文档定位

在 `phase01 ~ phase09` 正式收口后，用户通过四轮真实的 `Real-Project Dry-Run` 产生了一份极具穿透力的使用反馈。与之前基于文档的推演不同，这次反馈暴露了 PSCO 在“实体登记完备”与“经营动作闭环”之间存在的真实语义断层。

六位专家对此进行了独立的诊断与方向预判。本文作为**最终的交叉汇总与仲裁文档**，不只是罗列观点，而是要：
1. 明确必须立刻执行的**绝对共识（Fix 范畴）**。
2. 剥离并仲裁各方在**MVP0.4 主线方向**上的分歧。
3. 冻结 MVP0.4 的**正式演进路线与边界约束**，作为后续 `phase10+` `/plan` 的直接上游。

---

## 2. 高度共识：必须立刻闭合的现实阻断项

六位专家在第一层面上达成了绝对一致：**Dry-run 反馈不是失败，而是揭示了现有架构的最后一公里断层。在进入任何新宏图之前，必须通过 `fix` 工作流闭合阻断项。**

### 2.1 共识修复清单 (P0 优先级)
1. **Onboarding 冷启动断裂与逻辑隔离**
   - **问题**: 空态首访无响应；向导只是“并列新增孤立实体”，未能在用户心智中形成“首轮建链”。
   - **共识行动**: 修复按钮响应逻辑；重构 Onboarding 为“首轮建链引导”（主动建立 Product-Repo-Module 绑定，而非事后手动关联）。
2. **Decision 状态机未闭环 (假 Pending)**
   - **问题**: Dashboard 的 Current Focus 持续提示已关联的决策，Decision Detail 缺乏状态推进 CTA。
   - **共识行动**: 为 Decision 引入最小生命周期（`draft -> proposed -> acknowledged -> active`），修复 Dashboard 的误报。
3. **Review Loop 业务闭环断裂**
   - **问题**: Daily Review 发现问题后，跳转到详情页却无法完成闭环。
   - **共识行动**: 补齐实体详情页的 Cross-Entity CTA 矩阵，让每次跳转都有明确的“下一步动作”承接，并正确回流至 Review 结束状态。

> **仲裁决定：** 这三个阻断项必须立刻启动一个短平快的 `fix` 阶段解决。在它们修复并通过新一轮 Dry-run 验证前，不允许开启任何 MVP0.4 的新实体建设。

---

## 3. 核心分歧与最终仲裁：MVP0.4 的主轴究竟是什么？

专家们在解决 Bug 后，对 MVP0.4 的战略走向产生了分歧，主要集中在如何回应用户关于“未触及真实代码”、“期待 Agent 协同”以及“精益创业 workflow”的深刻反思。

### 3.1 分歧点梳理
*   **阵营 A (DPv4flash, GLM52, GPT54)**: 坚持稳健路线，主张 MVP0.4 的中心主线只能是 **Asset-Action Closure (资产动作闭环)**，坚决反对扩张实体（如 Venture/Experiment）或过度开发 AI 工作台，认为 Agent 协同只能作为受控的只读支撑层。
*   **阵营 B (DPv4pro, gemini31pro)**: 认为痛点在于“文本描述者陷阱”。主张 MVP0.4 必须是 **Connected OS**，通过 GitHub API 读时获取仓库元数据实现“真实连接”；同时将 **Agent-Driven Workflow**（读写倒置，Agent 写入 PSCO）作为核心演进。
*   **阵营 C (qwen37pro)**: 认为最根本的痛点是战略错位。主张 MVP0.4 必须开始最小化承接 **“科学实验 Workflow” (想法->假设->验证)**，引入 Opportunity/Experiment 实体，否则 PSCO 无法支撑“个人软件公司”的市场探索。

### 3.2 最终仲裁决定

结合用户的真实痛点与系统工程的可控性，我做出以下仲裁：

1. **驳回“科学实验 Workflow”作为 MVP0.4 主线 (针对阵营 C)**
   *   **理由**: PSCO 连基础的代码协作闭环和 Agent 集成都没打通，此时引入 `Opportunity / Experiment` 等重战略实体，只会创造出更多“需要人工填写的孤立表单”，加剧摩擦。该方向**继续后移至 MVP0.5+**。
2. **采纳“真实连接 (Real Connection)”作为破局点 (针对阵营 B)**
   *   **理由**: 用户不想只做“文本描述者”。MVP0.4 必须让 PSCO 触及真实世界。
   *   **约束**: 严格限制在“读时获取 (Read-time Fetch)”，如通过 GitHub API 显示仓库最后更新、模块版本。绝不将 PSCO 做成全自动代码分析器。
3. **确立“Agent-Native Foundation”为核心跨越 (整合阵营 A/B)**
   *   **理由**: “Agent 维护为主，Web 展示为辅”是 AI 研发时代不可逆的趋势。
   *   **约束**: 不是做 AI Chat 界面，而是提供标准的 **Agent Context Query API (MCP / REST)**，让 Trae/Cursor 能够读取 PSCO 的资产上下文（如 `project_rules`、已登记 Module），并能通过 API 写入新的 Decision。
4. **提升“跨项目全局上下文”为核心资产**
   *   **理由**: 解决用户新项目重复建规范的切肤之痛。

---

## 4. MVP0.4 正式演进路线 (The Roadmap)

基于上述仲裁，MVP0.4 的正式推进路线被确立为**“三步走”**战略。这三个阶段是线性的，不可倒置。

### Phase 10: Asset-Action Closure (资产动作闭环与修复)
**目标**: 清理技术债务，让已有实体“真正活起来”。
*   **Fix-01**: Onboarding 重构为“首轮建链向导”。
*   **Fix-02**: Decision 状态机细化，彻底解决假 Pending 与 Dashboard 误报。
*   **Fix-03**: 实体详情页下沉 Cross-Entity CTA，闭合 Review Loop 业务流。
*   *验收闸*: Dry-run 能够顺畅无阻断地完成一次日常推进。

### Phase 11: Real Connection (真实工程连接)
**目标**: 摆脱纯文本记录，读时派生真实开发元数据。
*   **Repository 真实化**: 在仓库详情页，通过 GitHub API 获取并展示星标、最近 Commit 活跃度、主要语言。
*   **Module 真实化**: 展示该模块在关联仓库中的 Git Tag/Release 映射。
*   *约束*: 数据不落库，失败优雅降级。

### Phase 12: Agent-Native Foundation & Global Context
**目标**: 建立面向 Agent 的读写接口，实现多项目规范下发。
*   **Agent Context API**: 建立标准化接口，Agent 可查询当前项目的 PSCO 资产树。
*   **Cross-Project Convention Asset**: 将 `.trae/rules`、`TECH_STACK` 等全局规范资产化。新建立 Product 时，PSCO 可自动为目标目录注入这些基线文件，确保 Agent 进场即上下文对齐。

---

## 5. 明确的非目标 (What NOT to do in MVP0.4)

为了防止概念膨胀，以下事项被**严格禁止**在 MVP0.4 中出现：

1. ❌ 引入 `Venture`, `Experiment`, `Opportunity` 实体（后移至 MVP0.5+）。
2. ❌ 开发内置的 AI 工作台、AI 聊天窗口或自动写代码引擎。
3. ❌ 开发完整的 Decision Intelligence（如 AI 历史决策相似度匹配）。
4. ❌ 进行深度本地文件系统扫描、AST 解析或代码质量分析。
5. ❌ 开发可配置的复杂 Onboarding 工作流引擎。

---

## 6. 结论

`phase01 ~ phase09` 证明了 PSCO 具备优秀的骨架。这次 `Dry-run` 的伟大之处在于，它用刺痛的反馈告诉我们：**在 AI 编程的纪元，一个纯靠人工维护关联关系的“文本台账”，是不可能成为 OS 的。**

MVP0.4 将是 PSCO 的一次范式跃迁：
从 **“用户手动记录资产”** 走向 **“Agent 协同维护、连接真实工程数据的智能中枢”**。

路线已锁定，请立即启动 Phase 10 的 Fix 流程，修复阻断项，打响 MVP0.4 的第一枪。
