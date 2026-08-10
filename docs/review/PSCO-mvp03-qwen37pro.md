# Personal Software Company OS

# MVP0.3 推进方向计划

**Author:** Qwen3.7-Pro  
**Date:** 2026-08-10  
**Purpose:** 基于 PSCO_0~4 原始愿景与 mvp0.1/mvp0.2 已验收交付基础，系统分析下一阶段（mvp0.3）的推进方向、范围基线与 phase 拆分，作为后续正式 `/plan` 的参考评价文档。  
**参照标准:** `PSCO-mvp02-summarize-feedback.md`（GPT54 签署的 mvp0.2 最终共识）+ `PSCO_0.md ~ PSCO_4.md`（Document 00-08 长期基线）

---

## 0. 文档定位与独立声明

- 本文是 **参考评价文档**，归入 `docs/review/`，不直接承担当前阶段正式规则。
- 本文不预设正式 phase 名称，不绕过 `phase*` workflow；正式推进仍需按 `project_rules.md §4.1` 走 `/plan -> /spec -> 实现 -> 验收 -> 收口`。
- 本文价值在于：在 mvp0.2 收口与下一阶段正式 phase 入口建立之间，提供一份系统化的方向分析与范围基线。
- 凡与 `PSCO-mvp02-summarize-feedback.md` 仲裁结论冲突之处，以共识文档为准。

---

## 1. 认知基础

### 1.1 PSCO 的核心价值与差异化定位（来自 Document 00-08）

PSCO 不是代码管理工具，不是 AI Chat 产品，不是自动扫描系统。它是 **个人软件公司的经营与资产系统**。

**长期愿景的最终形态**（Document 00 §7）：
- 五年目标：几十个成熟模块 + 多个商业产品 + 完整产品方法论 + 稳定软件生产流程
- 十年目标：形成个人软件公司基础设施（产品部门 + 技术部门 + 知识部门 + 战略部门）

**差异化核心**长期稳定为四件事：
```
Module（能力单元）+ Decision（决策留痕）+ Binding（资产锚定）+ Feedback（复利反馈）
```

**三条理念**：`Build / Accumulate / Compound`  
**长期飞轮**：`更多机会 → 更多产品 → 更多模块 → 更强能力 → 更快产品 → 更多机会`

**关键判断**：PSCO 的护城河不是"能登记多少资产"，而是"能否让资产产生复利"。如果没有复利反馈，PSCO 退化为模块台账工具。

### 1.2 mvp0.1 已交付基础（phase01~05 验收完结）

| Phase | 交付主线 | 验收结论 | 核心交付物 |
|-------|---------|---------|-----------|
| phase01 | MVP 规格收敛 | completed | `mvp_spec_v0.1.md` 冻结 |
| phase02 | Module Registry | phase02-12 验收通过 | 前后端最小主线 + .proto 合同 |
| phase03 | Decision Center | phase03-14 验收通过 | 正式规格 + .proto 合同 + 联调验收 |
| phase04 | Product Registry + Repository Binding | phase04-14 验收通过 | 三类绑定 + .proto 合同 + 联调验收 |
| phase05 | Dashboard + Feedback | phase05-14 验收通过 | 聚合读 + 反馈信号 + CTA + 局部错误隔离 |

**mvp0.1 已验证的核心假设**：
1. 资产登记闭环成立（Product / Module / Release / Repository / Decision）
2. 绑定动作闭环成立（三类绑定可执行）
3. 决策留痕闭环成立（Decision 可关联到任意目标）
4. 基础反馈闭环成立（Dashboard 聚合读 + 反馈信号 + CTA 矩阵）

### 1.3 mvp0.2 已交付基础（phase06 验收完结）

**mvp0.2 主题**：从 Asset Registry 走向 Operating System

**phase06 交付内容**（Onboarding + Data Sovereignty + Reuse Awareness）：
1. **Onboarding Foundation**
   - First-run 引导流程
   - 最小字段 / draft-first / partial-entry
   - 低摩擦 Product / Repository / Module / Decision 初始录入

2. **Data Sovereignty 闭合**
   - 资产导出（ExportAssets）
   - 基础备份（pg_dump 脚本）
   - 数据所有权承诺兑现

3. **Reuse Awareness Foundation**
   - `module_reuse_summary` 派生视图
   - `capability_summary` 派生视图
   - Module Detail / List 复用度展示
   - Dashboard Capability Growth 区块
   - 模块复用率度量指标

**mvp0.2 验收结论**：phase06-16 验收通过，已交付正式规格正文、最小 .proto 合同主线、后端与数据主线、前端主线与联调验收收口结果。

### 1.4 mvp0.2 遗留项（根据 PSCO-mvp02-summarize-feedback.md §6.1）

mvp0.2 的完整范围包括两个阶段，当前只完成了阶段一（phase06），阶段二尚未启动：

**阶段二待完成内容**：
1. **Operating Review Loop**
   - Daily / weekly review
   - Feedback -> Decision -> Update 闭环
   - 从 Dashboard 进入动作，并把结果回流到既有实体

2. **模板级复用最小版**
   - Module 组合快照
   - 新建 Product 时基于模板预填

3. **基础度量补齐**
   - 资产导入耗时
   - 录入摩擦感知

4. **真实项目 dry-run**
   - 优先以 Rento-miniX 或等价真实项目走完整流程
   - 形成独立验收记录

---

## 2. 下一阶段方向分析

### 2.1 候选方向

基于 mvp0.2 遗留项与长期愿景，mvp0.3 存在三个候选方向：

#### 方向 A：Operating Review Loop 深化

- Daily / weekly review 流程化
- Feedback -> Decision -> Update 闭环
- Review 结论回流到既有实体
- Action handoff（最小后续动作记录）
- **核心价值**：让 Dashboard 真正成为经营动作起点，兑现"Operating System"承诺

#### 方向 B：Decision 高级复用机制

- 相似决策匹配
- 历史决策引用推荐
- Decision 上下文增强
- **核心价值**：强化 Decision 护城河，为后续 AI 增强准备上下文前提

#### 方向 C：AI Context-Aware Enhancement 最小版本

- 页面内低干扰提示（Module / Product / Decision 页面）
- 结构化上下文准备（Product / Module / Decision 可结构化检索）
- 不实现完整 AI 工作台，只做最小增强层
- **核心价值**：为后续 AI Composition Assistant 准备上下文前提，兑现 Document 07 §9 的长期方向

### 2.2 方向仲裁

**最终结论：mvp0.3 以方向 A 为主体，方向 B 和 C 作为可选深化项，视阶段进展决定是否进入。**

理由：

1. **共识对齐**：
   - `PSCO-mvp02-summarize-feedback.md §6.1` 明确将 Operating Review Loop 列为阶段二核心内容
   - 这是 mvp0.2"从 Asset Registry 走向 Operating System"总叙事的自然延续
   - 方向 A 直接承接 mvp0.2 遗留项，不需要重新仲裁

2. **差异化护城河**：
   - mvp0.1 验证了"资产能登记"
   - mvp0.2 验证了"资产能复利可见"（复用感知）
   - mvp0.3 应该验证"资产能驱动日常经营"（Operating Loop）
   - 这是 PSCO 从"工具"走向"系统"的关键一步

3. **执行收缩原则**：
   - 方向 A 不引入新的重集成（OAuth / LLM 调用链）
   - 不引入新核心实体，只在已有数据之上做流程化
   - 符合 `PSCO-mvp02-summarize-feedback.md §4.2`"执行收缩而不是理论回退"

4. **上下文基础设施**：
   - 方向 A 产出的 Review Loop 是后续方向 B（Decision 复用）和方向 C（AI 增强）的使用场景前提
   - 先做 A，才能验证 B 和 C 的真实价值

5. **Decision 复用后移的合理性**：
   - `PSCO-mvp02-summarize-feedback.md §4.5` 明确"Decision 高级复用机制"后移到 mvp0.3+
   - mvp0.2 的 Decision 留痕已能完成基础闭环
   - Decision 复用需要更复杂的上下文匹配逻辑，适合在 Operating Loop 成熟后再推进

6. **AI 增强后移的合理性**：
   - `PSCO-mvp02-summarize-feedback.md §5.2` 明确 AI 一级工作台不进入 mvp0.2
   - AI context-aware enhancement 需要 Operating Loop 作为使用场景
   - 在没有真实 review 数据前，AI 增强的价值有限

### 2.3 非方向（明确不做）

以下方向在 mvp0.3 明确不进入，避免范围蔓延：

- Opportunity / Experiment / Feature 流程化（`PSCO-mvp02-summarize-feedback.md §5.2`）
- AI Assistant 一级工作台（`PSCO-mvp02-summarize-feedback.md §5.2`）
- Rust Intelligence Layer（`PSCO-mvp02-summarize-feedback.md §5.2`）
- 自动扫描 / 知识图谱（`PSCO-mvp02-summarize-feedback.md §5.2`）
- 完整 PMM / PCP 正式标准（`PSCO-mvp02-summarize-feedback.md §5.2`）
- Venture 强制实现（`PSCO-mvp01-summarize-feedback.md §4.5` 保留可选）
- GitHub OAuth 自动导入（`PSCO-mvp01-summarize-feedback.md §4.7` 后移 P1/P2）
- Capability 重实体 CRUD（`PSCO-mvp01-summarize-feedback.md §4.4` 继续派生层）

---

## 3. mvp0.3 范围基线

### 3.1 实体范围

**不新增核心实体。** mvp0.3 在 mvp0.1/mvp0.2 已有实体（Product / Module / Release / Decision / Repository）之上深化，不引入 Venture 强制、不引入 Capability 重实体、不引入 Feature/Opportunity/Experiment。

`Capability` 继续作为派生结果层，不建核心表。

### 3.2 流程化能力（核心新增）

| 能力 | 数据来源 | 用途 |
|------|---------|------|
| Daily Review | 聚合读（当前焦点、待处理决策、近期变化） | 每日经营起点 |
| Weekly Review | 汇总新增 Decision、Module / Release 变化、缺口/停滞/复用进展 | 周复盘与动作规划 |
| Review Conclusion | 用户输入的 review 结论（最小文本块） | 回流到既有实体 |
| Action Handoff | 从 review 结论转成的最小后续动作记录 | 锚定到既有实体的动作跟踪 |

实现方式：优先用现有 Dashboard 聚合读 + 新增 review 流程页面，不引入新的重基础设施。

### 3.3 页面增强

| 页面 | mvp0.3 增强 |
|------|------------|
| Dashboard | 新增 Daily Review 入口（当前焦点、待处理决策、近期变化） |
| Dashboard | 新增 Weekly Review 入口（汇总、缺口、动作规划） |
| Review Detail | 新增 review 结论录入与回流机制 |
| Decision Detail | 新增"相关历史决策"提示（可选，视进展决定） |

### 3.4 动作（核心新增）

mvp0.1/mvp0.2 已有动作保持不变。mvp0.3 新增：

- `StartDailyReview`：进入 daily review 流程
- `StartWeeklyReview`：进入 weekly review 流程
- `RecordReviewConclusion`：记录 review 结论并回流到既有实体
- `CreateActionFromReview`：从 review 结论创建后续动作

不新增核心实体的 CRUD。

### 3.5 度量指标（补齐 mvp0.2 遗留）

| 指标 | 实现方式 | mvp0.2 状态 | mvp0.3 目标 |
|------|---------|-----------|-----------|
| 资产导入耗时 | 从 CreateProduct 到完成首轮 Module 绑定的时间差 | 未实现 | **必须落地** |
| 录入摩擦感知 | 每个核心对象最小登记字段数 | 未实现 | 可观测不强制 |
| Review 完成率 | 完成 weekly review 的频率 | 未实现 | 可观测不强制 |
| Action 回流率 | 从 review 创建的动作被执行的比率 | 未实现 | 可观测不强制 |

mvp0.3 至少落地 **资产导入耗时**，其余可观测不强制。

---

## 4. phase 拆分与推进计划

### 4.1 拆分原则

- 每个 phase 仍是交付型 phase（`/plan -> /spec -> 实现 -> 验收 -> 收口`）
- 每个 phase 直接承接上一 phase 已冻结交付物，不回退重做
- mvp0.3 控制在 2 个 phase 内完成，保持紧凑

### 4.2 phase07：Operating Review Loop 基础

**目标**：让 Dashboard 成为经营动作起点，兑现"Operating System"承诺。

**范围**：
- Daily Review 流程（当前焦点、待处理决策、近期变化）
- Weekly Review 流程（汇总、缺口、动作规划）
- Review Conclusion 录入与回流机制
- Action Handoff 最小版本

**进入条件**：直接承接 phase06-12 / 13 / 16 已冻结交付物。  
**非目标**：不做 Decision 高级复用（留 phase08 可选），不做 AI 增强（留 mvp0.4+），不改 mvp0.1/mvp0.2 已冻结合同。

**关键交付物**：
- 正式规格正文（`.trae/specs/phase07_10_operating_review_loop_formal_spec/`）
- .proto 合同主线（`.trae/specs/phase07_11_operating_review_loop_proto_mainline/`）
- 后端与数据主线（`.trae/specs/phase07_12_operating_review_loop_backend_data_mainline/`）
- 前端主线（`.trae/specs/phase07_13_operating_review_loop_frontend_mainline/`）
- 联调验收与收口（`.trae/specs/phase07_14_operating_review_loop_integration_validation_acceptance/`）

### 4.3 phase08：度量收口、模板复用与真实项目 dry-run（可选）

**目标**：补齐 mvp0.2 遗留的度量指标，兑现模板级复用最小版，用真实项目 dry-run 验证飞轮。

**范围**：
- 资产导入耗时度量落地
- 模板级复用最小版（SaveModuleCompositionAsTemplate + ApplyModuleCompositionTemplate）
- Product Create 增强：基于模板预填
- 真实项目 dry-run（优先以 Rento-miniX 或等价真实项目走完整流程）

**进入条件**：直接承接 phase07 已冻结交付物。  
**非目标**：不做 GitHub OAuth（留 mvp0.4），不做 AI 增强（留 mvp0.4+），不做 Decision 高级复用（除非 phase07 进展顺利且有余力）。

**关键交付物**：
- 正式规格正文（`.trae/specs/phase08_10_metrics_template_dryrun_formal_spec/`）
- .proto 合同主线（`.trae/specs/phase08_11_metrics_template_dryrun_proto_mainline/`）
- 后端与数据主线（`.trae/specs/phase08_12_metrics_template_dryrun_backend_data_mainline/`）
- 前端主线（`.trae/specs/phase08_13_metrics_template_dryrun_frontend_mainline/`）
- 联调验收与收口（`.trae/specs/phase08_14_metrics_template_dryrun_integration_validation_acceptance/`）
- 真实项目 dry-run 报告（`.trae/specs/phase08_15_real_project_dry_run/`）

**可选深化项**（视 phase07 进展决定）：
- Decision 高级复用最小版本（相似决策匹配、历史决策引用推荐）
- AI context-aware enhancement 最小版本（页面内低干扰提示）

### 4.4 推进预览

```
phase07 Operating Review Loop 基础
  └─ Daily/Weekly Review + Review Conclusion + Action Handoff
        ↓
phase08 度量收口 + 模板复用 + 真实项目 dry-run（可选）
  └─ 资产导入耗时 + 能力模板最小版 + dry-run
        ↓
mvp0.3 收口 → 下一阶段（mvp0.4）候选：GitHub 集成 / AI context-aware 增强 / Decision 高级复用
```

---

## 5. mvp0.3 非目标

明确不做，避免范围蔓延：

1. ❌ Opportunity / Experiment / Feature 流程化
2. ❌ Venture 强制实现（继续可选不强制，不建表）
3. ❌ Capability 主动 CRUD（继续派生层）
4. ❌ GitHub OAuth + Repository 自动导入
5. ❌ AI Assistant 一级工作台 / context-aware AI 增强（除非 phase08 可选深化）
6. ❌ Rust Intelligence Layer / 自动扫描 / 知识图谱
7. ❌ 完整 PMM / PCP 正式标准
8. ❌ 第二套 UI 框架 / 第二套路由 / 第二套 ORM（遵守 TECH_STACK_BASELINE）
9. ❌ Decision 高级复用机制（除非 phase08 可选深化）

---

## 6. 度量与 Done 标准

### 6.1 mvp0.3 Done 标准

只有当以下条件同时成立，才可认为 mvp0.3 规格成立并可进入稳定实现：

1. Daily Review 流程可完成，用户能看到当前焦点、待处理决策、近期变化
2. Weekly Review 流程可完成，用户能汇总变化、记录结论、创建动作
3. Review Conclusion 可回流到既有实体（Product / Module / Decision）
4. Action Handoff 可创建并锚定到既有实体
5. 资产导入耗时度量可在 Dashboard 观测
6. 至少一个真实项目 dry-run 走通（如果 phase08 执行）
7. 未引入超出 TECH_STACK_BASELINE 的技术选择
8. 未回退重做 mvp0.1/mvp0.2 已冻结交付物

### 6.2 度量与验收原则

- 验收环境继续通过 `reset_*_acceptance.sh` + fixture 建立，保持"他人可稳定复验"
- Review 流程必须有 fixture 覆盖（含空状态、有数据、多轮 review 三类场景）
- 真实项目 dry-run 不替代 fixture 验收，二者并存
- Action Handoff 必须验证"创建 -> 锚定 -> 执行"闭环

---

## 7. 风险与依赖

### 7.1 风险

| 风险 | 级别 | 缓解措施 |
|------|------|---------|
| Review Loop 过度工程化为完整任务管理系统 | 中 | 严格限定为"最小后续动作记录"，不做 Kanban / Sprint / 通用任务平台 |
| Review Conclusion 回流机制复杂度过高 | 中 | 第一版只做文本块 + 实体关联，不做复杂工作流 |
| 真实项目 dry-run 依赖外部项目可用性 | 低 | 优先 Rento-miniX；若不可用，允许用等价真实项目，但必须在验收报告中说明 |
| 模板级复用边界失控（过度工程为完整模板系统） | 中 | 严格限定为"Module 组合快照 + 预填"，不做模板版本管理、不做参数化 |

### 7.2 依赖

- mvp0.3 直接依赖 mvp0.1/mvp0.2 已冻结的 .proto 合同主线、数据库 schema、前后端模块结构
- 不依赖任何外部服务（GitHub / LLM），保持 mvp0.3 自闭可验收
- 不依赖 Venture 实现（继续可选）
- 不依赖 Decision 高级复用机制（除非 phase08 可选深化）

---

## 8. 与长期愿景的对齐度

| 长期愿景要素 | mvp0.3 推进情况 | 对齐度 |
|-------------|---------------|--------|
| Build（持续创造） | 不直接推进（Opportunity/Product 创建仍为 mvp0.1 水平） | 低 |
| Accumulate（持续积累） | 间接推进（Review Loop 让积累可见可回顾） | 中 |
| Compound（持续复利） | 直接推进（Operating Loop 让复利可行动） | **高** |
| Module 差异化核心 | 间接推进（Review 中可见模块复用） | 中 |
| Decision 护城河 | 直接推进（Review 中 Decision 回流承接） | **高** |
| AI 增强层 | 不推进（留 mvp0.4+） | 低 |
| Local First | 已兑现（mvp0.2 已交付导出/备份） | **高** |
| 五年目标（几十个成熟模块 + 多个商业产品） | 间接推进（Operating Loop 加速产品迭代） | 中 |

**对齐度判断**：mvp0.3 聚焦在 `Compound / Decision 护城河 / Local First` 三个长期要素上，是 mvp0.2 验证"能复利可见"之后验证"能驱动日常经营"的自然下一步，不贪多，不跳跃。

---

## 9. 与 mvp0.2 共识的交叉参照

`PSCO-mvp02-summarize-feedback.md` 提供了 mvp0.2 的最终仲裁与规划基线。本文与其的主要一致性：

1. **主题一致**：均以"从 Asset Registry 走向 Operating System"为总叙事
2. **范围一致**：均将 Operating Review Loop 列为阶段二核心内容
3. **非目标一致**：均明确不做 GitHub OAuth / AI 一级工作台 / Venture 强制
4. **Done 标准一致**：均要求 Review Loop 落地 + 模板闭环 + 真实项目 dry-run

**本文的补充视角**：

1. **更强调 phase 拆分的紧凑性**：本文将 mvp0.3 控制在 2 个 phase 内（phase07 + phase08），保持紧凑
2. **更强调 phase08 的可选性**：本文将模板级复用、度量收口、真实项目 dry-run 作为 phase08 的可选深化项，视 phase07 进展决定
3. **更强调 Decision 高级复用的后移理由**：本文认为 Decision 高级复用需要 Operating Loop 作为使用场景，在 Review Loop 成熟前推进会增加复杂度
4. **更强调 AI context-aware enhancement 的后移理由**：本文认为 AI 增强需要 Operating Loop 作为使用场景，在没有真实 review 数据前，AI 增强的价值有限

---

## 10. 最终结论

mvp0.1 验证了 PSCO 的"资产登记 + 决策留痕 + 基础复用反馈"闭环成立。  
mvp0.2 验证了 PSCO 的"低摩擦进入 + 数据主权 + 复用感知"闭环成立。  
下一阶段的核心命题是 **让已登记的资产驱动日常经营**。

因此 mvp0.3 的推进方向应聚焦于：

> **兑现 Operating Review Loop，让 Dashboard 成为经营动作起点，用真实项目 dry-run 验证飞轮可转。**

一句话收口：

> **不扩实体，先做流程；不引集成，先做闭环；不碰 AI，先用真实项目验证经营收益。**

mvp0.3 完成后，PSCO 将从"资产复利系统"升级为"资产经营系统"，此时再进入 mvp0.4 的 GitHub 集成、AI context-aware 增强与 Decision 高级复用才有意义——因为 Operating Loop 是 AI 增强与 Decision 复用的使用场景前提，没有它，智能推荐无从谈起。

---

## 11. 给后续 `/plan` 的执行建议

1. **正式 phase 入口建立**：建议按 `project_rules.md §4.1` 先建立 phase07 的 `/plan` 三件套（architecture_plan / dev_plan / shared_baseline），再进入 `/spec -> 实现 -> 验收 -> 收口`。

2. **根级真相源切换**：phase07 完成后，需同步更新 `AGENTS.md`、`plan.md`、`architecture_map.md`、`docs/README.md`，将当前阶段从 phase06 切换到 phase07。

3. **规格正文互链**：phase07/08 的规格正文必须明确直接承接 phase06-12 / 13 / 16，不回退重做 mvp0.1/mvp0.2 已冻结交付物。

4. **验收环境**：继续通过 `reset_*_acceptance.sh` + fixture 建立可重复复核的验收环境，Review 流程必须有 fixture 覆盖。

5. **真实项目 dry-run**：建议在 phase08 结束时，优先以 Rento-miniX 或等价真实项目走完整 dry-run，兑现 `PSCO-mvp02-summarize-feedback.md §6.4 Phase 3` 的承诺。

---

**Reference Document**  
**Author: Qwen3.7-Pro（以 GPT54 共识标准为参照）**  
**Date: 2026-08-10**
