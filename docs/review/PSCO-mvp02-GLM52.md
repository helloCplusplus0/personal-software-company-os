# Personal Software Company OS

# mvp0.2 推进方向计划

**Author:** GLM-5.2
**Date:** 2026-08-09
**Purpose:** 基于 PSCO_0~4 原始愿景与 mvp0.1（phase01~05）已验收交付基础，系统分析下一阶段（mvp0.2）的推进方向、范围基线与 phase 拆分，作为 GPT54 主导正式 phase 规划时的参考评价文档。
**参照标准:** `PSCO-mvp0.1-summarize-feedback.md`（GPT54 签署的最终共识）+ `PSCO_0.md ~ PSCO_4.md`（Document 00-08 长期基线）

---

## 0. 文档定位与免责

- 本文是 **参考评价文档**，归入 `docs/review/`，不直接承担当前阶段正式规则。
- 本文不预设正式 phase 名称，不绕过 `phase*` workflow；正式推进仍需由 GPT54 按 `project_rules.md §4.1` 走 `/plan -> /spec -> 实现 -> 验收 -> 收口`。
- 本文价值在于：在 mvp0.1 收口与下一阶段正式 phase 入口建立之间，提供一份系统化的方向分析与范围基线，避免下一阶段凭直觉启动。
- 凡与 `PSCO-mvp0.1-summarize-feedback.md` 仲裁结论冲突之处，以共识文档为准。

---

## 1. 认知基础

### 1.1 PSCO 的核心价值（来自 Document 00-08）

PSCO 不是代码管理工具，不是 AI Chat 产品，不是自动扫描系统。它是 **个人软件公司的经营与资产系统**。

差异化核心长期稳定为四件事：

```
Module（能力单元）+ Decision（决策留痕）+ Binding（资产锚定）+ Feedback（复利反馈）
```

三条理念：`Build / Accumulate / Compound`。
五条原则：产品优先 / 能力优先 / 实践优先 / 简单优先 / AI 增强而非替代。
长期飞轮：`更多机会 → 更多产品 → 更多模块 → 更强能力 → 更快产品 → 更多机会`。

**关键判断**：PSCO 的护城河不是"能登记多少资产"，而是"能否让资产产生复利"。如果没有复利反馈，PSCO 退化为模块台账工具（summarize-feedback §4.3 的明确警告）。

### 1.2 mvp0.1 已交付基础（phase01~05 验收完结）

| Phase | 交付主线 | 验收结论 |
|-------|---------|---------|
| phase01 | MVP 规格收敛（`mvp_spec_v0.1.md` 冻结） | completed |
| phase02 | Module Registry（CRUD + Release + .proto 合同） | phase02-12 验收通过 |
| phase03 | Decision Center（CRUD + LinkDecisionToTarget + .proto） | phase03-14 验收通过 |
| phase04 | Product Registry + Repository Binding（三类绑定 + .proto） | phase04-14 验收通过 |
| phase05 | Dashboard + Feedback（聚合读 + 反馈信号 + CTA + 跳转返回 + 局部错误隔离） | phase05-14 验收通过 |

工程基础已稳定：
- 技术栈冻结为 `Durable System Track`（React + Go + PostgreSQL + Drizzle + TanStack + shadcn/ui）
- `.proto` 合同主线已落地（module / decision / product+binding / dashboard 四套）
- 数据库 reset + fixture 验收体系可重复复核
- 多跳返回 + 来源上下文恢复机制已成型

### 1.3 mvp0.1 已验证的核心假设

按 `PSCO-mvp0.1-summarize-feedback.md §6.4 Phase 3` 的 dry-run 精神，mvp0.1 在真实前后端 + 真实数据库上验证了：

1. **资产登记闭环成立**：Product / Module / Release / Repository / Decision 可完整登记
2. **绑定动作闭环成立**：三类绑定（Product-Module / Product-Repository / Module-Repository）可执行
3. **决策留痕闭环成立**：Decision 可关联到任意目标，source_context 可持久化承接
4. **基础反馈闭环成立**：Dashboard 聚合读 + 反馈信号（pending_decision + product_asset_coverage）+ CTA 矩阵可驱动下一步动作
5. **局部错误隔离成立**：附属聚合失败不拖垮主链路

### 1.4 mvp0.1 的真实缺口（实测确认）

通过代码与 schema 实测，mvp0.1 以下能力 **未实现**，均为共识文档明确延后或未列入范围：

| 缺口 | 共识定位 | 实测状态 |
|------|---------|---------|
| `capability_summary` 派生视图 | summarize-feedback §5.3 列为派生视图 | 未实现 |
| `module_reuse_summary` 派生视图 | summarize-feedback §5.3 列为派生视图 | 未实现 |
| `Venture` 可选实体 | summarize-feedback §4.5 保留为可选 | 未实现（符合"可选不强制"） |
| 数据导出 / 备份 | summarize-feedback §4.6 / §5.2 Local First 要求 | 未实现 |
| 4 个先导度量指标 | summarize-feedback §7.1 | 仅 product_asset_coverage / pending_decision_signals 间接覆盖，复用率/导入耗时未实现 |
| 真实项目 dry-run | summarize-feedback §6.4 Phase 3 | 验收用 fixture，未以真实项目（如 Rento-miniX）走完整 dry-run |
| 模板级复用 | summarize-feedback §4.9 明确为 v0.2 第一优先级 | 未实现 |
| GitHub OAuth 自动导入 | summarize-feedback §4.7 后移 P1/P2 | 未实现 |
| AI Assistant 工作台 | summarize-feedback §5.1 不作为一级主导航 | 未实现 |
| Rust Intelligence Layer | summarize-feedback §4.6 不进 v0.1 | 未实现 |

---

## 2. 下一阶段方向分析

### 2.1 候选方向

基于原始愿景与 mvp0.1 缺口，下一阶段存在三个候选方向：

#### 方向 A：复用感知深化

- 补齐 `module_reuse_summary` + `capability_summary` 派生视图
- Module Detail / List 增强：展示"被哪些 Product 使用"、"复用次数"
- Dashboard 增强：Capability Growth 区块（Document 08 §5.2 的派生反馈）
- 模板级复用最小版：从已有 Product 的 Module 组合导出为能力模板，新建 Product 时可预填绑定
- 度量指标补齐：模块复用率、资产导入耗时

#### 方向 B：导入与集成

- GitHub OAuth + Repository 自动导入（summarize-feedback §4.7 的 P1/P2）
- 数据导出 / 备份（Local First §4.6）
- 仓库元数据同步

#### 方向 C：决策深化与 AI 增强

- Decision 复用（新建相似决策时引用历史 Decision）
- 页面内 context-aware AI enhancement（Document 07 §13 / Document 08 §13）
- AI 辅助文档生成（Document 07 §7）

### 2.2 方向仲裁

**最终结论：mvp0.2 以方向 A 为主体，方向 B 的数据导出部分作为收尾，方向 C 延后 mvp0.3+。**

理由：

1. **共识对齐**：summarize-feedback §4.9 明确"模板级复用是 v0.2 第一优先级"，§5.3 将 `capability_summary` / `module_reuse_summary` 列为派生视图。方向 A 直接承接共识，不需要重新仲裁。

2. **差异化护城河**：mvp0.1 已验证"资产能登记"，但 **未验证"资产能复利"**。`module_reuse_summary` 与 `capability_summary` 缺失，意味着用户看不到"我拥有了什么能力"（Document 08 §9 的核心页面语义）。没有复用感知，飞轮转不起来（DS-Flash 在第二轮汇总中的明确提醒）。

3. **执行收缩原则**：方向 A 不引入新的重集成（OAuth / token / API 限流 / LLM 调用链），不引入新核心实体，只在已有数据之上做派生视图与最小模板功能，符合 summarize-feedback §2.3"执行收缩而不是理论回退"。

4. **上下文基础设施**：方向 A 产出的 `module_reuse_summary` 是后续方向 C（AI Composition Assistant，Document 07 §9）的上下文前提。先做 A，才能做 C。

5. **GitHub OAuth 后移的合理性**：summarize-feedback §4.7 已判定 GitHub 自动导入不作为 P0。mvp0.1 的手动绑定已能完成验证闭环，mvp0.2 仍可维持手动绑定，把集成复杂度留给 mvp0.3。

6. **数据导出提前到 mvp0.2 的理由**：Local First 是 PSCO 的核心原则（Document 05 §3.3），数据所有权/导出/备份是 summarize-feedback §4.6 的明确要求。mvp0.1 已积累真实资产数据，若 mvp0.2 仍不提供导出，用户的数据所有权承诺无法兑现。导出实现成本低（JSON 序列化），不应继续延后。

### 2.3 非方向（明确不做）

以下方向在 mvp0.2 明确不进入，避免范围蔓延：

- Opportunity / Experiment / Feature 流程化（summarize-feedback §4.2 延后）
- AI Assistant 一级工作台（summarize-feedback §5.1）
- Rust Intelligence Layer（summarize-feedback §4.6）
- 自动扫描 / 知识图谱（summarize-feedback §5.4）
- 完整 PMM / PCP 正式标准（summarize-feedback §4.8）
- Venture 强制实现（summarize-feedback §4.5 保留可选）
- GitHub OAuth 自动导入（后移 mvp0.3）

---

## 3. mvp0.2 范围基线

### 3.1 实体范围

**不新增核心实体。** mvp0.2 在 mvp0.1 已有实体（Product / Module / Release / Decision / Repository）之上深化，不引入 Venture 强制、不引入 Capability 重实体、不引入 Feature/Opportunity/Experiment。

`Capability` 继续作为派生结果层，不建核心表（与 summarize-feedback §4.4 一致）。

### 3.2 派生视图（核心新增）

| 派生视图 | 数据来源 | 用途 |
|---------|---------|------|
| `module_reuse_summary` | `product_modules` 聚合 | 每个 Module 被哪些 Product 使用、复用次数、跨产品复用标记 |
| `capability_summary` | Module 分类 + 复用 + 稳定状态聚合 | 按能力分类（Foundation/Application/Domain/AI/Data）聚合模块数、稳定模块数、复用模块数 |

实现方式：优先用 SQL 查询 / 视图，不引入物化视图除非性能证明需要。保持 Local First 的可导出性。

### 3.3 页面增强

| 页面 | mvp0.2 增强 |
|------|------------|
| Module Detail | 新增 "Used By Products" 区块（哪些 Product 使用此 Module） |
| Module List | 新增按复用度排序/筛选；展示复用次数 |
| Dashboard | 新增 "Capability Growth" 区块（Document 08 §5.2 派生反馈） |
| Dashboard | 新增度量指标区块（模块复用率、资产导入耗时） |
| Product Create | 新增"基于能力模板预填"入口（模板级复用最小版） |

### 3.4 动作（核心新增）

mvp0.1 已有动作保持不变。mvp0.2 新增：

- `ExportAssets`：导出全部资产（products / modules / releases / decisions / repositories / bindings）为 JSON
- `SaveModuleCompositionAsTemplate`：将某 Product 的 Module 组合保存为能力模板
- `ApplyModuleCompositionTemplate`：新建 Product 时基于模板预填 Module 绑定

不新增核心实体的 CRUD。

### 3.5 度量指标（补齐 summarize-feedback §7.1）

| 指标 | 实现方式 | mvp0.1 状态 |
|------|---------|-----------|
| 模块复用率 | `module_reuse_summary` 派生（被 2+ Product 使用的 Module 占比） | 未实现 |
| 决策复用率 | Decision 引用历史 Decision 的比例（需 Decision 复用最小支持，mvp0.2 可先观测不强制） | 未实现 |
| 资产导入耗时 | 从 CreateProduct 到完成首轮 Module 绑定的时间差（基于现有 timestamp） | 未实现 |
| 录入摩擦感知 | 每个核心对象最小登记字段数（静态可计算） | 未实现 |

mvp0.2 至少落地 **模块复用率** 与 **资产导入耗时** 两个，其余可观测不强制。

---

## 4. phase 拆分与推进计划

### 4.1 拆分原则

- 每个 phase 仍是交付型 phase（`/plan -> /spec -> 实现 -> 验收 -> 收口`）
- 每个 phase 直接承接上一 phase 已冻结交付物，不回退重做
- mvp0.2 控制在 2 个 phase 内完成，保持紧凑

### 4.2 phase06：复用感知与能力派生

**目标**：让"能力复利"可见，补齐 mvp0.1 缺失的两个派生视图与对应页面增强。

**范围**：
- `module_reuse_summary` 派生视图 + Module Detail "Used By" 区块 + Module List 复用度排序
- `capability_summary` 派生视图 + Dashboard "Capability Growth" 区块
- 模块复用率度量指标落地到 Dashboard

**进入条件**：直接承接 phase05-10 / 11 / 14 已冻结交付物。
**非目标**：不做模板级复用（留 phase07），不做数据导出（留 phase07），不改 mvp0.1 已冻结合同。

### 4.3 phase07：模板级复用、度量收口与数据导出

**目标**：兑现 summarize-feedback §4.9 的"v0.2 第一优先级"与 §4.6 的 Local First 导出要求。

**范围**：
- 能力模板最小版（SaveModuleCompositionAsTemplate + ApplyModuleCompositionTemplate）
- Product Create 增强：基于模板预填
- 资产导入耗时度量落地
- `ExportAssets` 动作 + JSON 导出 endpoint + 前端导出入口
- 真实项目 dry-run（优先以 Rento-miniX 或等价真实项目走完整流程，兑现 summarize-feedback §6.4 Phase 3）

**进入条件**：直接承接 phase06 已冻结交付物。
**非目标**：不做 GitHub OAuth（留 mvp0.3），不做 AI 增强（留 mvp0.3+）。

### 4.4 推进预览

```
phase06 复用感知与能力派生
  └─ module_reuse_summary + capability_summary + Dashboard Capability Growth
        ↓
phase07 模板级复用 + 度量收口 + 数据导出
  └─ 能力模板最小版 + ExportAssets + 真实项目 dry-run
        ↓
mvp0.2 收口 → 下一阶段（mvp0.3）候选：GitHub 集成 / AI context-aware 增强
```

---

## 5. mvp0.2 非目标

明确不做，避免范围蔓延：

1. ❌ Opportunity / Experiment / Feature 流程化
2. ❌ Venture 强制实现（继续可选不强制，不建表）
3. ❌ Capability 主动 CRUD（继续派生层）
4. ❌ GitHub OAuth + Repository 自动导入
5. ❌ AI Assistant 一级工作台 / context-aware AI 增强
6. ❌ Rust Intelligence Layer / 自动扫描 / 知识图谱
7. ❌ 完整 PMM / PCP 正式标准
8. ❌ 第二套 UI 框架 / 第二套路由 / 第二套 ORM（遵守 TECH_STACK_BASELINE）

---

## 6. 度量与 Done 标准

### 6.1 mvp0.2 Done 标准

只有当以下条件同时成立，才可认为 mvp0.2 规格成立并可进入稳定实现：

1. `module_reuse_summary` 与 `capability_summary` 派生视图已落地且可重复复核
2. Module Detail / List / Dashboard 对应增强已通过真实前后端联调
3. 能力模板最小版可完成"保存组合 → 新建 Product 时预填"闭环
4. `ExportAssets` 可导出完整资产 JSON，且可重新导入或离线阅读
5. 模块复用率与资产导入耗时度量可在 Dashboard 观测
6. 至少一个真实项目 dry-run 走通（创建 Product → 绑定 Repository → 注册 Module → 录入 Decision → 观察复用与能力反馈）
7. 未引入超出 TECH_STACK_BASELINE 的技术选择
8. 未回退重做 mvp0.1 已冻结交付物

### 6.2 度量与验收原则

- 验收环境继续通过 `reset_*_acceptance.sh` + fixture 建立，保持"他人可稳定复验"
- 派生视图必须有 fixture 覆盖（含零复用、单产品复用、跨产品复用三类场景）
- 导出必须验证"导出 → 离线可读 → 结构完整"三段式
- 真实项目 dry-run 不替代 fixture 验收，二者并存

---

## 7. 风险与依赖

### 7.1 风险

| 风险 | 级别 | 缓解措施 |
|------|------|---------|
| 模板级复用最小版边界失控（过度工程为完整模板系统） | 中 | 严格限定为"Module 组合快照 + 预填"，不做模板版本管理、不做参数化 |
| 派生视图性能（数据量增长后聚合慢） | 低 | mvp0.2 数据量为个人级，先用 SQL 查询；性能证明需要时再加物化视图/缓存，但不引入 Redis（违反基线） |
| 度量指标定义模糊导致无法验收 | 中 | 模块复用率/资产导入耗时必须在 phase06/07 spec 中给出精确计算公式与 fixture 期望值 |
| 真实项目 dry-run 依赖外部项目可用性 | 低 | 优先 Rento-miniX；若不可用，允许用等价真实项目，但必须在验收报告中说明 |

### 7.2 依赖

- mvp0.2 直接依赖 mvp0.1 已冻结的 `.proto` 合同主线、数据库 schema、前后端模块结构
- 不依赖任何外部服务（GitHub / LLM），保持 mvp0.2 自闭可验收
- 不依赖 Venture 实现（继续可选）

---

## 8. 与长期愿景的对齐度

| 长期愿景要素 | mvp0.2 推进情况 |
|-------------|---------------|
| Build（持续创造） | 不直接推进（Opportunity/Product 创建仍为 mvp0.1 水平） |
| Accumulate（持续积累） | 直接推进（复用视图 + 能力模板让积累可见可用） |
| Compound（持续复利） | 直接推进（复用感知是复利的前提出现） |
| Module 差异化核心 | 直接推进 |
| Decision 护城河 | 不直接推进（Decision 复用留 mvp0.3） |
| AI 增强层 | 不推进（留 mvp0.3+） |
| Local First | 直接推进（数据导出兑现所有权承诺） |

**对齐度判断**：mvp0.2 聚焦在 `Accumulate / Compound / Module 差异化 / Local First` 四个长期要素上，是 mvp0.1 验证"能登记"之后验证"能复利"的自然下一步，不贪多，不跳跃。

---

## 9. 最终结论

mvp0.1 验证了 PSCO 的"资产登记 + 决策留痕 + 基础复用反馈"闭环成立。下一阶段的核心命题不是"扩张实体范围"，而是 **让已登记的资产产生复利反馈**。

因此 mvp0.2 的推进方向应聚焦于：

> **补齐复用感知（module_reuse_summary + capability_summary），兑现模板级复用最小版，落地 Local First 数据导出，并用真实项目 dry-run 验证飞轮可转。**

一句话收口：

> **不扩实体，先做复利；不引集成，先做导出；不碰 AI，先用真实项目验证资产化收益。**

mvp0.2 完成后，PSCO 将从"资产登记系统"升级为"资产复利系统"，此时再进入 mvp0.3 的 GitHub 集成与 AI context-aware 增强才有意义——因为复用感知是 AI Composition 的上下文前提，没有它 AI 推荐无从谈起。

---

**Reference Document**
**Author: GLM-5.2（以 GPT54 共识标准为参照）**
**Date: 2026-08-09**
