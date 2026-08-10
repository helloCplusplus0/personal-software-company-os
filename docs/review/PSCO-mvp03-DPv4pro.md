# Personal Software Company OS

# MVP0.3 推进方向展望

**Author:** DeepSeek-V4-Pro
**Date:** 2026-08-10
**Role:** 基于原始方案（`PSCO_0.md ~ PSCO_4.md`）与 MVP0.1~MVP0.2（`phase01 ~ phase06`）已验收现实，系统输出 PSCO 下一阶段推进方向的标准计划文档，供后续正式 `/plan` 参考评价
**Document Type:** `review`
**Status:** 供后续 `/plan` 参考，不直接构成正式 phase 命名或执行指令

---

## 0. 文档定位与前置说明

### 0.1 本文与 MVP0.2 共识的关系

MVP0.2 的最终共识（`PSCO-mvp02-summarize-feedback.md`）将 MVP0.2 定义为两个阶段：

- **阶段一（已完成 = phase06）**：Onboarding + 数据主权闭合 + 复用感知基础
- **阶段二（尚未完成）**：Operating Review Loop + 模板级复用 + 派生智能深化 + 真实项目 dry-run

本文在制定 MVP0.3 展望时，**假定阶段二将作为 phase07 完成 MVP0.2 收口**。这意味着 MVP0.3 的推进基础是：

- phase01~05：MVP0.1 资产登记系统（已完成）
- phase06：MVP0.2 阶段一（Onboarding + 导出/备份 + 复用感知）（已完成）
- phase07：MVP0.2 阶段二（Review Loop + 模板复用 + 派生智能 + dry-run）（本文假定将完成）

### 0.2 本文不替代正式 `/plan`

本文是方向性展望文档，归入 `docs/review/`。正式推进仍需按 `project_rules.md §4.1` 走 `/plan -> /spec -> 实现 -> 验收 -> 收口`。本文不冻结正式 phase 名称、spec 路径或实现细节。

---

## 1. 第一项任务：PSCO 核心价值与最终形态再确认

### 1.1 一句话定位（来自 Document 00）

> 让一个个人开发者具备一家软件公司的组织能力，使其能够持续发现机会、验证需求、构建产品、积累软件能力，并形成长期复利资产。

### 1.2 三大核心理念

| 理念 | 含义 | 对应系统能力 |
| --- | --- | --- |
| **Build（持续创造）** | 像公司一样持续寻找机会、构建产品 | 当前以 Product / Module / Decision 承载 |
| **Accumulate（持续积累）** | 每次开发都留下资产，从项目归档升级为能力提取 | 当前以 Module Registry + Reuse Awareness 承载 |
| **Compound（持续复利）** | 长期价值来自能力复用，模块被未来产品反复复用 | 当前以 capability_summary + module_reuse_summary 承载 |

### 1.3 四条差异化支柱

```
Module（能力单元）+ Decision（决策留痕）+ Binding（资产锚定）+ Feedback（复利反馈）
```

这四条在 MVP0.1 和 MVP0.2 中已全部落地，构成了 PSCO 不可替代的差异化护城河。

### 1.4 长期飞轮（来自 Document 02/03）

```
更多机会 → 更多产品 → 更多模块 → 更强能力 → 更快产品 → 更多机会
```

这个飞轮在 MVP0.1 中建立了"资产登记"基础，在 MVP0.2 中建立了"复用感知"与"经营闭环"基础。MVP0.3 需要让这个飞轮**从"可观测"升级为"可加速"**。

### 1.5 最终呈现形态：四重视角的 Command Center

PSCO 的最终形态（Document 08）是一个具备四重视角的个人软件公司运行系统：

1. **经营视角**：知道自己在做什么、为什么做、当前最值得推进什么
2. **产品视角**：把产品、价值主张、模块、决策、发布串成可追溯的价值链
3. **工程资产视角**：看见哪些模块已沉淀、是否稳定、是否被复用、是否值得进一步抽象
4. **上下文增强视角**：AI 在上下文中做总结、提醒、候选与文档辅助，但不抢走判断权

**当前进展：**

| 视角 | MVP0.1 状态 | MVP0.2 状态（phase06 完成） | MVP0.2 状态（phase07 假定完成） |
| --- | --- | --- | --- |
| 经营视角 | Dashboard 基础聚合读 | + Onboarding 入口 + 导出/备份 | + Daily/Weekly Review |
| 产品视角 | Product / Module / Decision CRUD | + 复用感知 + 部分回写 | + 模板复用 + Feedback→Decision 闭环 |
| 工程资产视角 | Module Registry + Release | + module_reuse_summary + capability_summary | + 派生智能深化 |
| 上下文增强视角 | 未启动 | 未启动 | 未启动 |

**关键发现：** 即使 MVP0.2 全部完成，**第四重视角（上下文增强视角）仍然完全空白**。这是 MVP0.3 最自然的切入方向。

---

## 2. 第二项任务：当前推进基础（phase01~06 已验收现实）

### 2.1 已交付的六个 phase

| Phase | 交付主题 | 状态 | 核心验证 |
| --- | --- | --- | --- |
| phase01 | MVP 规格收敛 | ✅ | 执行层唯一规格入口冻结 |
| phase02 | Module Registry | ✅ | 模块 CRUD + Release + .proto |
| phase03 | Decision Center | ✅ | 决策 CRUD + LinkDecisionToTarget |
| phase04 | Product Registry + Repository Binding | ✅ | 产品 CRUD + 三类绑定 |
| phase05 | Dashboard + Feedback | ✅ | 聚合读 + 反馈信号 + CTA + 跳转返回 |
| phase06 | Onboarding + 数据主权 + 复用感知 | ✅ | 首轮录入 + 导出/备份 + 复用反馈 |

### 2.2 phase06 已验证的核心假设

根据 `phase06-16 acceptance_report.md`，以下事实已经成立：

1. **首轮成功会话可重复走通**：新用户可在一次会话内完成 Product + Repository + Module + Decision 四类对象的首轮录入
2. **数据主权路径可验证**：Export 覆盖 9 类资产（含绑定/关联关系），Backup 校验语义（含三类失败语义）完整
3. **复用反馈路径可验证**：module_reuse_summary 和 capability_summary 在 Dashboard 和 Detail 页均可见
4. **Onboarding 回流链完整**：fromOnboarding 标记在 Onboarding ↔ canonical detail 之间正确传递
5. **phase05 兼容性保持**：Dashboard 原有主线不受 phase06 新增块破坏

### 2.3 MVP0.2 阶段二（phase07 假定完成）的预期基础

基于 `PSCO-mvp02-summarize-feedback.md` §6.1，phase07 完成后系统将具备：

- Daily / Weekly Review：用户在 Dashboard 中可完成日常/周复盘
- Feedback → Decision → Update 闭环：从反馈信号到决策到实体更新
- 模板级复用最小版：Module 组合快照 + 新建 Product 时预填
- 派生智能深化：候选能力提示、跨产品复用标记、基础度量
- 真实项目 dry-run：至少一个真实项目走通完整流程

### 2.4 当前系统的真实边界

即使 MVP0.2 全部完成，以下能力仍然**不在系统范围内**：

| 缺口 | 状态 | 长期重要性 |
| --- | --- | --- |
| Venture（战略容器） | 可选，未实现 | 高 — 经营闭环建立后，"为什么做"成为自然需求 |
| Opportunity（机会管理） | 延后，未实现 | 中 — 飞轮起点，但当前阶段可继续手工管理 |
| 决策复用/相似匹配 | 延后，未实现 | 高 — 决策数量增长后，历史决策的价值需要被释放 |
| GitHub 集成 | 延后，未实现 | 中 — 数据主权闭合后，代码连接成为自然延伸 |
| AI 上下文增强 | 延后，未实现 | 高 — 原始方案中定义为"个人 AI 增强层"的长期目标 |
| 能力模板的参数化与版本管理 | 明确不做（模板最小版） | 低 — 长期方向，但当前阶段非必要 |

---

## 3. MVP0.3 的定位：从 Operating System 到 Intelligence-Enhanced Growth System

### 3.1 三个阶段的产品演进

```
MVP0.1: Asset Registry（资产登记系统）
  "我记录了什么资产"
  
MVP0.2: Operating System（经营系统）
  "我知道接下来该做什么，以及这些动作如何沉淀为资产"
  
MVP0.3: Intelligence-Enhanced Growth System（智能增强的增长系统）
  "我能更快地做出更好的决策，看见更远的增长路径"
```

### 3.2 MVP0.3 的核心命题

> **在经营闭环已经建立的基础上，如何让 PSCO 开始"加速"用户的决策质量和资产增长速度？**

这个命题的三个子问题：

1. **战略层**：当用户有了多个 Product 和 Module 后，如何帮用户回答"我应该优先投入什么方向"？（Venture）
2. **决策层**：当决策数量增长到几十上百条后，如何让历史决策产生复利而非只占据存储？（Decision Intelligence）
3. **增强层**：当系统中的 Product / Module / Decision / Review 数据足够丰富后，如何让 AI 在正确的时间提供正确的上下文增强？（AI Context Enhancement）

### 3.3 MVP0.3 的三条主线

我建议 MVP0.3 收敛为三条主线，它们之间有明确的依赖关系：

```
主线 A: Venture & Strategic Context
  让用户把多个 Product 放进一个战略容器，回答"为什么做"
  
主线 B: Decision Intelligence
  让历史决策从"留痕"升级为"可检索、可复用、可产生洞察"
  
主线 C: AI Context Enhancement
  让 AI 在正确上下文中提供辅助，但不替代人的判断
```

**依赖关系：**

```
主线 A（Venture）
    │  提供战略上下文
    ▼
主线 B（Decision Intelligence）
    │  提供结构化决策数据
    ▼
主线 C（AI Context Enhancement）
    │  提供智能增强能力
    ▼
长期飞轮加速
```

---

## 4. 主线 A：Venture & Strategic Context（战略上下文）

### 4.1 为什么现在是引入 Venture 的合适时机

Venture 在原始方案中被定义为"长期战略容器"，在 MVP0.1 和 MVP0.2 中一直保持"可选，不强制"的地位。现在是引入它的合适时机，原因如下：

1. **经营闭环已建立**：用户已经能够 daily/weekly review 自己的产品与模块。当用户看到多个 Product 时，自然的下一步是问"这些 Product 之间是什么关系？我为什么在做它们？"
2. **复用感知已存在**：module_reuse_summary 和 capability_summary 让用户看见跨产品复用。Venture 是解释"为什么这些产品共享这些模块"的自然容器。
3. **模板复用已建立**：当用户能够从一个 Product 的 Module 组合创建新 Product 时，Venture 为该组合提供了"为什么这样组合"的战略理由。

### 4.2 Venture 的最小引入方式

**原则：Venture 在 MVP0.3 中仍然保持"可选，不强制"，但不再是"不存在"。**

| 维度 | 最小定义 |
| --- | --- |
| 实体 | `ventures` 表，字段：`id / name / mission / status / created_at / updated_at` |
| 关系 | `Venture 1:N Product`（一个 Venture 可包含多个 Product） |
| 创建 | 可选创建，不强制成为 Product 的前置条件 |
| 展示 | Dashboard 中可选展示"Active Ventures"区块；Product Detail 中展示所属 Venture |
| 非目标 | 不做 Venture 的复杂生命周期管理、不做 Venture 级度量、不做 Venture 模板 |

### 4.3 Venture 对现有系统的增强

| 现有页面 | 增强内容 |
| --- | --- |
| Dashboard | 新增可选区块"Active Ventures"（若用户创建了 Venture） |
| Product Detail | 展示所属 Venture（若有），点击跳转 Venture 详情 |
| Product Create | 可选选择所属 Venture |
| Product List | 可选按 Venture 筛选 |
| Module Detail | 展示"被哪些 Venture 的产品使用"（间接复用上下文） |

### 4.4 Venture 的明确非目标

- ❌ 不强制要求 Product 必须属于 Venture
- ❌ 不做 Opportunity → Venture 的完整工作流
- ❌ 不做 Venture 级的度量仪表盘
- ❌ 不做 Venture 模板或 Venture 组合
- ❌ 不引入 Feature / Experiment 实体

---

## 5. 主线 B：Decision Intelligence（决策智能）

### 5.1 为什么决策智能是 MVP0.3 的核心

PSCO 的差异化核心之一是"决策留痕"（Decision）。在 MVP0.1 中，Decision 完成了 CRUD 和关联目标。在 MVP0.2 中，Decision 通过低摩擦 capture 和 Feedback → Decision 闭环变得更易用。

但当系统中的 Decision 数量增长到几十上百条后，一个自然的问题浮现：

> 我三年前做过类似的决策吗？当时是怎么选的？结果如何？

这是 PSCO 从"记录工具"变成"判断加速器"的关键一步。

### 5.2 Decision Intelligence 的最小能力集

| 能力 | 描述 | 优先级 |
| --- | --- | --- |
| **Decision 可检索** | 按标题、上下文、关联目标类型、时间范围检索历史决策 | P0 |
| **Decision 可引用** | 新建决策时可引用历史决策作为"前置决策" | P1 |
| **相似决策发现** | 基于标题和上下文的文本相似度，在新建决策时推荐相关历史决策 | P1 |
| **决策密度视图** | 在 Product / Module Detail 中展示"关联决策时间线" | P1 |
| **决策结果追踪** | 决策记录增加"结果"字段（outcome），支持后续复盘 | P2 |

### 5.3 关键设计约束

1. **Decision 引用不是复杂依赖图**：引用关系是"参考"而非"阻塞"，不引入决策依赖链的复杂状态管理
2. **相似度匹配是辅助，不是替代**：推荐结果以"你可能想参考"的方式呈现，不做自动关联
3. **结果追踪是可选字段**：不强制用户填写结果，但提供"后续复盘"的入口
4. **不引入 AI 驱动的决策自动生成**：AI 可以辅助总结和推荐，但决策本身始终由人撰写和确认

### 5.4 Decision Intelligence 对现有系统的增强

| 现有能力 | 增强内容 |
| --- | --- |
| Decision Create | 新建时展示"你可能想参考的历史决策"（相似度推荐） |
| Decision Create | 支持引用历史决策作为"前置决策" |
| Decision Detail | 展示"被哪些后续决策引用" |
| Decision List | 支持全文检索 |
| Product Detail | 展示"关联决策时间线" |
| Module Detail | 展示"关联决策时间线" |
| Dashboard Review | 周复盘时提示"近期关键决策" |

---

## 6. 主线 C：AI Context Enhancement（AI 上下文增强）

### 6.1 为什么现在是引入 AI 增强的合适时机

PSCO 原始方案中，AI 被定义为"增强层，不是决策层"（Document 07）。这个原则在 MVP0.1 和 MVP0.2 中一直得到严格遵守。

引入 AI 增强的时机判断标准不是"AI 技术是否成熟"，而是：

> **系统是否已经积累了足够丰富的结构化上下文，使 AI 的增强建议有实际价值？**

在 MVP0.2 完成后，系统将具备：

- 多个 Product 及其价值主张
- 多个 Module 及其分类、稳定状态、复用关系
- 多条 Decision 及其上下文、关联目标
- 多轮 Review 结论
- 能力派生摘要（capability_summary）

这些数据构成了 PSCO 原始方案中定义的"Personal Context Layer"——这是 AI 增强的前提。现在这个前提已经开始成立。

### 6.2 AI Context Enhancement 的最小能力集

**关键原则：AI 永远是增强层。不做 AI 一级导航，不做 AI 自动决策，不做 AI 自动扫描。**

| 能力 | 描述 | 触发场景 | 优先级 |
| --- | --- | --- | --- |
| **决策草稿辅助** | 基于当前 Product / Module / 历史决策上下文，AI 生成决策草稿（标题、背景、候选方案），由人确认和修改 | Decision Create 页面 | P0 |
| **模块复用提示** | 当用户创建新 Product 时，AI 基于已有模块和复用模式，推荐可能适合的模块组合 | Product Create 或 Product Detail 的模块绑定区 | P1 |
| **周报总结** | AI 基于本周的 Review 数据、新增 Decision、Module 变化，生成周报草稿 | Dashboard Weekly Review | P1 |
| **能力增长洞察** | AI 基于 capability_summary 的历史变化，生成"能力增长趋势"的自然语言描述 | Dashboard Capability Growth 区块 | P2 |
| **决策影响分析** | 在 Decision Detail 中，AI 总结该决策影响了哪些 Product / Module 的后续变化 | Decision Detail 页面 | P2 |

### 6.3 AI 增强的架构约束

1. **AI 调用必须显式触发**：不做后台自动分析，不做定时扫描。AI 增强只在用户主动进入相关页面时触发。
2. **AI 输出永远可编辑**：AI 生成的内容（决策草稿、周报总结等）永远以"草稿"形式呈现，用户必须确认或修改后才能保存。
3. **AI 不直接访问代码仓库**：AI 的上下文来源仅限于 PSCO 系统内部的结构化数据（Product / Module / Decision / Review / Capability），不扫描外部代码仓库。
4. **AI 失败不阻塞主流程**：AI 调用超时或失败时，页面回退到纯人工模式，不展示错误，不阻塞用户操作。
5. **AI 模型选择可配置**：支持通过环境变量配置 AI endpoint 和模型，不硬编码任何特定 AI 服务。

### 6.4 AI 增强的明确非目标

- ❌ 不做 AI 一级主导航
- ❌ 不做 AI Chat 界面
- ❌ 不做自动代码扫描
- ❌ 不做自动知识图谱
- ❌ 不做 AI 自动判断最佳方案
- ❌ 不做 AI 自动创建/修改实体
- ❌ 不做 Rust Intelligence Layer

---

## 7. 三条主线的整体推进建议

### 7.1 推荐的 phase 拆分

```
phase07（MVP0.2 收口）:
  Review Loop + 模板复用 + 派生智能深化 + 真实项目 dry-run
  
phase08（MVP0.3 启动 — 主线 A）:
  Venture & Strategic Context
  
phase09（MVP0.3 — 主线 B）:
  Decision Intelligence
  
phase10（MVP0.3 收口 — 主线 C）:
  AI Context Enhancement
```

### 7.2 依赖关系论证

```
phase07（Review Loop）
    │  提供：review 数据、Feedback→Decision 闭环
    │  提供：模板复用基础
    ▼
phase08（Venture）
    │  提供：战略上下文
    │  依赖：phase07 的 review 数据使"为什么做"成为可被 review 的问题
    ▼
phase09（Decision Intelligence）
    │  提供：决策可检索、可引用、相似匹配
    │  依赖：phase08 的 Venture 提供"为什么做"的战略上下文
    │  依赖：phase07 的 review 数据使决策密度可观测
    ▼
phase10（AI Context Enhancement）
    │  提供：AI 辅助决策草稿、模块复用提示、周报总结
    │  依赖：phase08 的 Venture 提供战略上下文
    │  依赖：phase09 的 Decision Intelligence 提供结构化决策数据
    │  依赖：phase07 的 review 数据提供时间维度的上下文
```

### 7.3 为什么这个顺序

1. **phase07 必须先完成**：它是 MVP0.2 的收口，是 MVP0.3 的地基
2. **Venture 先于 Decision Intelligence**：决策的"为什么"需要战略上下文；没有 Venture，"为什么做这个决策"只能停留在单个 Product 层面
3. **Decision Intelligence 先于 AI Enhancement**：AI 增强需要结构化决策数据作为上下文；如果决策只是散装文本，AI 的增强建议质量会很低
4. **AI 增强最后进入**：它是对前三层数据（Product/Module/Decision/Review/Venture）的智能增强，必须在数据层成熟后才能产生真正价值

---

## 8. MVP0.3 的非目标（明确不做）

### 8.1 继续不做的事（从 MVP0.1/MVP0.2 延续）

1. ❌ Opportunity / Feature / Experiment 流程化
2. ❌ Capability 重实体 CRUD（继续派生层）
3. ❌ GitHub OAuth / 全自动导入
4. ❌ 自动扫描 / 知识图谱
5. ❌ Rust Intelligence Layer
6. ❌ 通用项目管理系统化
7. ❌ 第二套技术栈（路由 / 合同 / ORM / UI 框架）
8. ❌ AI 一级主导航 / AI Chat 界面

### 8.2 MVP0.3 新增的非目标

9. ❌ Venture 强制实现（继续可选，但不再是"不存在"）
10. ❌ 完整的 Opportunity → Venture 工作流
11. ❌ AI 自动决策或自动创建实体
12. ❌ AI 后台自动分析或定时扫描
13. ❌ 决策依赖图或复杂决策链管理
14. ❌ 能力模板的参数化与版本管理

---

## 9. MVP0.3 的度量与验收标准

### 9.1 对 MVP0.3 整体的验收问题

当 MVP0.3 完成时，系统应能回答以下问题：

1. 用户是否能把多个 Product 组织到一个 Venture 下，并理解它们之间的战略关系？
2. 用户是否能在新建决策时，快速找到和引用相关的历史决策？
3. 用户是否能在 AI 辅助下，更快地完成决策草稿、周报总结和模块复用建议？
4. AI 增强是否做到"可触发、可编辑、可关闭、失败不阻塞"？
5. 系统是否继续遵守"AI 是增强层，不是决策层"的核心原则？

### 9.2 各主线的独立 DoD

#### 主线 A（Venture）DoD

- Venture 可创建、编辑、查看
- Product 可可选关联到 Venture
- Dashboard 展示"Active Ventures"区块（若用户创建了 Venture）
- Product Detail 展示所属 Venture
- Venture 不强制成为 Product 的前置条件

#### 主线 B（Decision Intelligence）DoD

- 决策可全文检索
- 新建决策时可引用历史决策
- 新建决策时展示相似历史决策推荐
- Product / Module Detail 展示关联决策时间线
- 决策引用不引入复杂依赖图管理

#### 主线 C（AI Context Enhancement）DoD

- Decision Create 页面提供 AI 辅助决策草稿
- Product Create / Detail 页面提供 AI 模块复用提示
- Dashboard Weekly Review 提供 AI 周报总结
- AI 输出永远可编辑、可丢弃
- AI 失败时页面回退到纯人工模式，不阻塞用户操作
- AI endpoint 与模型可通过环境变量配置

---

## 10. 与长期愿景的对齐度

| 长期愿景要素 | MVP0.1 状态 | MVP0.2 状态 | MVP0.3 展望 |
| --- | --- | --- | --- |
| Build（持续创造） | Product CRUD | + Onboarding + 模板复用 | + Venture 战略容器 |
| Accumulate（持续积累） | Module Registry | + 复用感知 | + 决策可检索/可引用 |
| Compound（持续复利） | 基础绑定 | + capability_summary | + AI 模块复用建议 |
| Module 差异化核心 | Module CRUD | + capability_summary | + 决策关联 + AI 复用提示 |
| Decision 护城河 | Decision CRUD | + 低摩擦 capture | + 决策智能（检索/引用/相似） |
| AI 增强层 | 未启动 | 未启动 | **首次进入**（决策草稿/周报/复用提示） |
| Local First | 未闭合 | + 导出/备份 | 继续维持 |
| 五年目标 | 资产登记 | 经营闭环 | 智能加速 |

**对齐度判断：** MVP0.3 是整个 PSCO 演进中**首次将 AI 增强层引入系统**的阶段，也是首次让"决策"从一个"记录对象"升级为"可复利的知识资产"的阶段。它不新建重实体，不偏离 Build / Accumulate / Compound 的核心理念，而是在已有数据基础上，让系统开始"加速"用户的决策质量和资产增长速度。

---

## 11. 风险与缓解

| 风险 | 级别 | 缓解措施 |
| --- | --- | --- |
| AI 增强过度工程化（演变为 AI Chat 或自动化系统） | 🔴 高 | 严格遵守"AI 是可触发、可编辑、可关闭的增强层"原则；AI 失败不阻塞主流程；不做 AI 一级导航 |
| Venture 过度工程化（演变为完整战略管理系统） | 🟡 中 | 严格限定为最小字段 + 可选关联；不做 Venture 生命周期、度量仪表盘、模板 |
| 决策引用演变为复杂依赖图 | 🟡 中 | 引用关系定义为"参考"而非"阻塞"；不引入依赖链状态管理 |
| AI 上下文质量不足导致增强建议无价值 | 🟡 中 | 先在 phase08/09 中建立高质量结构化数据，再在 phase10 中引入 AI；AI 增强以"草稿"形式呈现，用户始终可以丢弃 |
| 三条主线的 phase 拆分导致 MVP0.3 战线过长 | 🟡 中 | 每个 phase 独立交付、独立验收；phase08（Venture）本身即可独立产生价值，不依赖 phase09/10 |

---

## 12. 最终结论

### 12.1 一句话总结

> **MVP0.3 的任务，是在经营闭环已经建立的基础上，让 PSCO 从"记录和经营的系统"升级为"加速决策和增长的系统"。**

### 12.2 三条主线

1. **Venture & Strategic Context**：让用户回答"为什么做"，把多个 Product 放进战略容器
2. **Decision Intelligence**：让历史决策从"留痕"升级为"可检索、可引用、可产生洞察"
3. **AI Context Enhancement**：让 AI 在正确上下文中辅助决策草稿、模块复用提示和周报总结，但永远不替代人的判断

### 12.3 推进顺序

```
phase07（MVP0.2 收口）:
  Review Loop + 模板复用 + 派生智能深化 + 真实项目 dry-run

phase08: Venture & Strategic Context
phase09: Decision Intelligence
phase10: AI Context Enhancement
```

### 12.4 核心原则

1. **不扩实体，先做智能**：Venture 是唯一新增实体，且保持可选、不强制
2. **AI 永远是增强层**：可触发、可编辑、可关闭、失败不阻塞
3. **数据质量先于 AI 能力**：先建立结构化决策数据（phase08/09），再引入 AI 增强（phase10）
4. **继续坚持技术栈冻结**：不引入新基础设施、不引入第二套事实源
5. **继续坚持"资产登记 → 经营闭环 → 智能加速"的渐进式演进路径**

---

*End of PSCO-mvp03-DPv4pro.md*