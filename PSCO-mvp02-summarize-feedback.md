# Personal Software Company OS

# MVP0.2 Final Consensus & Planning Baseline

**Author:** GPT54  
**Date:** 2026-08-09  
**Purpose:** 基于两轮共十份 `mvp0.2` 专家评审与交叉汇总文档，形成 PSCO 在 `phase05` 收口之后、进入下一阶段正式 `/plan` 之前的最终仲裁结论、范围基线与推进顺序。

---

## 1. 文档定位

本文不是第十一份评审意见，也不是直接替代下一阶段正式 `phase` `/plan` 的执行文档。

本文要解决的是四件事：

1. 明确哪些判断已经形成稳定共识；
2. 对仍有分歧的部分给出最终仲裁；
3. 冻结 `mvp0.2` 的正式方向、范围与非目标；
4. 为下一阶段正式 `/plan` 提供唯一上游判断基线。

本文的职责是为下一阶段正式 phase 入口建立提供上游依据，而不是在这里越权冻结正式 `phase` 名称、`.trae/specs/phaseXX_*` 路径或具体实现细节。

---

## 2. 证据来源

## 2.1 第一轮方向评审

- `docs/review/PSCO-mvp02-deepseekv4flash.md`
- `docs/review/PSCO-mvp02-deepseekv4pro.md`
- `docs/review/PSCO-mvp02-GLM52.md`
- `docs/review/PSCO-mvp02-GPT54.md`
- `docs/review/PSCO-mvp02-qwen37pro.md`

## 2.2 第二轮交叉汇总

- `docs/review/PSCO-mvp02-summarize-feedback-deepseekv4flash.md`
- `docs/review/PSCO-mvp02-summarize-feedback-deepseekv4pro.md`
- `docs/review/PSCO-mvp02-summarize-feedback-GLM52.md`
- `docs/review/PSCO-mvp02-summarize-feedback-GPT54.md`
- `docs/review/PSCO-mvp02-summarize-feedback-qwen37pro.md`

## 2.3 本文采用的仲裁原则

1. **高共识优先于单点新颖性。** 五方独立重复出现的结论，优先视为稳定输入。
2. **使用频率优先于概念扩张。** 下一阶段先证明用户会持续回来，再扩对象宽度。
3. **底座负债优先闭合。** 已在 `mvp_spec_v0.1.md` 冻结、但仍未实现的要求，不得继续当成普通“后续功能”处理。
4. **执行收缩优先于理论回退。** 可以收缩阶段范围，但不重写 PSCO 已建立的核心语言。
5. **评审不越权冻结正式 phase。** 正式阶段命名、spec 路径、技术落点留给后续 `/plan` 和 `/spec` 决定。

---

## 3. 最终共识

经过两轮评审，以下内容已经形成稳定共识，可直接视为 `mvp0.2` 的正式前提。

## 3.1 关于项目方向的共识

1. `phase01 ~ phase05` 已证明 PSCO 的最小资产主线成立。
2. 当前系统已经能够完成资产登记、决策留痕、绑定关系与 Dashboard 最小反馈闭环。
3. 当前系统仍偏“登记与查看”，还没有真正进入“日常经营操作”。
4. 下一阶段不应继续优先扩展 `Opportunity / Feature / Experiment` 等长期对象。

## 3.2 关于范围边界的共识

1. 不新增核心实体主线。
2. `Capability` 继续作为派生层，不进入重实体 CRUD。
3. `Decision` 继续保持核心地位，不能被弱化出主线。
4. 不做 GitHub OAuth 自动导入、不做 AI 一级工作台、不做 Rust Intelligence Layer、不做自动扫描 / 知识图谱。
5. 技术栈继续冻结，保持现有单服务、单合同源、单导航事实源。

## 3.3 关于必须补齐事项的共识

1. 冷启动 / Onboarding 摩擦必须在下一阶段被正面处理。
2. 导出 / 备份必须进入 `mvp0.2`，且不能再后移为尾部补丁。
3. `module_reuse_summary` 与 `capability_summary` 必须进入下一阶段范围。
4. Dashboard 必须继续从“总览页”推进为“经营动作起点”。
5. 真实项目 `dry-run` 必须进入后续验收视野，而不是只停留在 fixture 验收。

---

## 4. 最终仲裁结论

## 4.1 `mvp0.2` 的总主题

**最终结论：**

> `mvp0.2` 的总主题确定为：  
> **从 Asset Registry 走向 Operating System。**

它的正式主轴是：

1. `Onboarding Foundation`
2. `Operating Review Loop`
3. `Derived Asset Intelligence`

**理由：**

- 这条主轴最完整地覆盖了“进入 -> 使用 -> 复利”的因果链；
- 它保留了 PSCO 的差异化核心：`Module + Decision + Binding + Feedback`；
- 它能承接复用感知、模板级复用与数据主权，而不把 PSCO 收窄成单纯的“模块复利工具”。

## 4.2 对“经营闭环优先”与“复利感知优先”的仲裁

**最终结论：**

> 两条路线不是二选一。  
> 最终采用“经营闭环主轴 + 复利感知前置接入”的融合方案。

具体含义如下：

- 我**不采纳**“只做复利感知，再谈 operating loop”的排序；
- 我也**不采纳**“把复用感知整体后置到最后才进入”的排序；
- 我采纳的是：**以 Onboarding / Operating Loop 作为总叙事，同时把复用感知基础提前纳入下一阶段早段范围。**

这意味着：

1. `module_reuse_summary`
2. `capability_summary`
3. 导出 / 备份
4. 低摩擦 Decision capture

都不应被拖到“最后有空再做”。

## 4.3 导出 / 备份的定位

**最终结论：**

> 导出 / 备份不是普通增强项，而是 `mvp0.1` 遗留的规格合规负债；  
> 它必须在 `mvp0.2` 的早期阶段或贯穿性工作项中闭合。

**理由：**

- `mvp_spec_v0.1.md §7.3 / §7.4` 已明确要求；
- 这直接对应 PSCO 的“数据所有权优先”承诺；
- 如果继续后移，新阶段会叠加在未闭合的数据主权底座上。

## 4.4 复用感知与模板级复用的定位

**最终结论：**

> 复用感知是 `mvp0.2` 的核心能力之一，但不是唯一总主题；  
> 模板级复用是重要方向，但不单独取代 Onboarding 与 Review Loop。

正式采纳的能力包括：

- `module_reuse_summary`
- `capability_summary`
- `module_reuse_rate`
- `asset_import_duration`
- 模板级复用最小版

其中模板级复用的边界冻结为：

> **“Module 组合快照 + 预填辅助”，不是新一级业务实体，不引入完整模板系统。**

## 4.5 Decision 的下一阶段定位

**最终结论：**

> 可以后移的是 `Decision` 的高级复用机制，不能后移的是 `Decision` 在 operating loop 里的中心地位。

`mvp0.2` 必须继续推进：

- 更低摩擦的 `Decision capture`
- `Feedback -> Decision -> Update` 闭环
- review 中对 `Decision` 的回流承接

后移到 `mvp0.3+` 的才是：

- 相似决策匹配
- 历史决策引用推荐
- 更重的 AI 决策上下文增强

## 4.6 真实项目 dry-run 的定位

**最终结论：**

> 真实项目 `dry-run` 不是可有可无的补充说明，而应作为后段验收中的独立交付要求。

它不能替代 fixture 验收，但必须补上“真实使用摩擦”和“资产复利收益”是否成立的证据。

## 4.7 当前不允许提前冻结的内容

以下内容在本文件中明确**不冻结**，留给后续正式 `/plan` 与 `/spec`：

1. 下一阶段正式 `phase` 名称；
2. `.trae/specs/phaseXX_*` 路径；
3. 导出接口名、脚本名与目录落点；
4. 模板的具体存储方式；
5. `dry-run` 使用的最终真实项目对象。

---

## 5. `mvp0.2` 正式范围基线

## 5.1 必做范围

下一阶段正式 `/plan` 必须覆盖以下五组范围：

1. **Onboarding Foundation**
   - first-run 引导
   - 最小字段 / draft-first / partial-entry
   - 低摩擦 `Product / Repository / Module / Decision` 初始录入

2. **数据主权闭合**
   - 资产导出
   - 基础备份
   - 不依赖第三方平台作为唯一前提

3. **Reuse Awareness Foundation**
   - `module_reuse_summary`
   - `capability_summary`
   - Module / Dashboard 上对复用与能力增长的最小可见反馈

4. **Operating Review Loop**
   - daily / weekly review
   - `Feedback -> Decision -> Update`
   - 从 Dashboard 进入动作、并把结果回流到既有实体

5. **模板级复用与后段验证**
   - `Module` 组合快照
   - 新建 `Product` 时基于模板预填
   - 真实项目 `dry-run`

## 5.2 明确不做范围

`mvp0.2` 明确不做：

- `Opportunity / Feature / Experiment` 流程化
- `Capability` 重实体 CRUD
- GitHub OAuth / 自动导入
- AI 一级工作台
- 自动扫描 / 知识图谱
- Rust Intelligence Layer
- 通用项目管理系统化
- `Decision` 高级复用引擎

---

## 6. 对后续 `/plan` 的正式输入

## 6.1 推荐的候选阶段结构

本文不冻结正式 `phase` 命名，但冻结如下候选推进顺序：

### 阶段一：Onboarding + 数据主权 + 复用感知基础

目标：

- 让用户低摩擦进入系统；
- 让用户相信数据可带走；
- 让用户在早期就能看见最小复利反馈。

至少应承接：

- first-run onboarding
- 低摩擦 Decision capture
- 导出 / 备份闭合
- `module_reuse_summary`
- `capability_summary`

### 阶段二：Operating Review Loop + 模板级复用 + 派生智能深化

目标：

- 让 Dashboard 成为经营动作起点；
- 让复利反馈从“可见”走向“可行动”；
- 用真实项目验证飞轮是否成立。

至少应承接：

- daily / weekly review
- `Feedback -> Decision -> Update`
- 模板级复用最小版
- 基础度量
- 真实项目 `dry-run`

## 6.2 推荐的先后关系

后续 `/plan` 必须体现以下依赖关系：

1. 没有 Onboarding，真实数据进不来；
2. 没有导出 / 备份，数据主权承诺不成立；
3. 没有复用感知，用户看不到 PSCO 的差异化价值；
4. 没有 Review Loop，Dashboard 无法真正成为 operating console；
5. 没有 `dry-run`，就无法证明这套闭环在真实项目上成立。

---

## 7. 最终度量与验收口径

下一阶段正式验收至少应回答以下问题：

1. 新用户是否能在一次会话内完成首个 `Product + Repository + Module + Decision` 录入？
2. 用户是否能导出核心资产、并完成基础备份？
3. 用户是否能在 Dashboard 或相关详情页看见可解释的复用 / 能力派生反馈？
4. 用户是否能完成一次完整的 weekly review，并把结果回流到既有实体？
5. 模板级复用是否完成“保存组合 -> 新建预填 -> 继续编辑 -> 完成创建”的闭环？
6. 至少一个真实项目 `dry-run` 是否走通，并形成独立验收记录？

---

## 8. 最终结论

1. `mvp0.1` 已经证明 PSCO 不是空想模型，而是可运行的最小资产系统。
2. `mvp0.2` 不应继续优先扩实体，也不应把 PSCO 收窄为单一的“模块复利看板”。
3. 下一阶段的正式方向已经确定为：**`Onboarding Foundation + Operating Review Loop + Derived Asset Intelligence`**。
4. 复用感知、模板级复用、导出 / 备份、真实项目 `dry-run` 均进入 `mvp0.2`，但它们必须服从“从 Asset Registry 走向 Operating System”的总叙事。
5. 后续正式 `/plan` 应按“阶段一：Onboarding + 数据主权 + 复用感知基础；阶段二：Review Loop + 模板级复用 + 派生智能深化 + dry-run”的候选结构继续收敛。
6. 在正式 phase 入口建立前，不得越权预设新的 phase 名称或 spec 路径。

> **`mvp0.2` 的任务，不是给 PSCO 增加更多概念，而是把它从一个已经可登记资产的系统，推进成一个用户愿意反复回来使用、能看见复利、并且真正拥有自己数据的个人软件公司 operating system。**
