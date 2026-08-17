# phase13-12 完成根级同步、阶段收口与下一阶段进入条件回写 Spec

## Why

`phase13-11` 已按固定样本、固定入口、固定 `6` 问与固定 rerun 协议完成正式验收（`6/6` 达标、工具链通过、浏览器矩阵 `12/12`、独立复核无阻断问题），`Project Governance Profile Foundation` 的阶段主线交付已完成。但验收后的 dogfooding 级反思发现：当前"项目治理画像"的承接形态相对 PSCO 的初心使命存在结构性缺口，用户已明确判定 `phase12~13` 在该承接位上的设计方向需要修正。若不在阶段收口时把这份反思以正式缺口记录留档，`phase14 /plan` 将缺少单值上游，容易再次滑向"在画像上继续加字段"的错误路径。

同时，根级文档（`AGENTS.md / plan.md / docs/README.md / architecture_map.md / docs/phase/README.md`）仍停留在 `phase13 in_progress` 状态，需按 `dev_plan` L254-266 完成回写收口；`phase13-11` 独立复核留档的 `plan.md` 状态同步建议也需在本阶段承接。

`phase13-12` 的目标不是继续扩写功能，而是：完成根级同步与阶段收口，并把本次反思结论冻结为正式缺口记录，作为 `phase14 /plan` 的直接输入。

## What Changes

- 新增本阶段唯一正式缺口记录：完整留档用户对治理画像的反思结论、缺口清单（GAP-01 ~ GAP-07）与方向结论（CON-01 ~ CON-09）
- 冻结 `phase14` 进入条件，其中首次要求 `phase14 /plan` 必须优先裁决五项：全局规范实体（`Standard`）的数据模型与 pg 承载方式、信息维护颗粒度、`Standard` 信息与模板仓库的内容边界、治理画像重叠承接位的退役计划、第五实体正式地位
- 回写根级五文档：`AGENTS.md / plan.md / docs/README.md / architecture_map.md / docs/phase/README.md`
- 留档 `phase13` 正式验收与收口入口（`phase13_11` acceptance_report）
- 承接 `phase13-11` 复核非阻断建议 1（`plan.md` `phase13` 状态位同步）
- 明确本阶段不做任何 `Standard` 实体实现、不做画像代码改动（缺口只留档，不实施）

## Impact

- Affected specs:
  - `.trae/specs/phase13_11_validate_project_governance_profile_integration_dogfooding_regression/`（验收结论作为本阶段收口依据）
  - `docs/phase/phase13_project_governance_profile_foundation_dev_plan.md`（L254-266 根级同步范围；§4 非目标第 8 条"第五个业务主实体"的时效性由本记录重新界定）
  - `docs/phase/phase13_project_governance_profile_foundation_shared_baseline.md`（能力矩阵中画像字段矩阵成为 GAP-02/GAP-04 的定位对象）
  - `docs/phase/phase13_project_governance_profile_foundation_architecture_plan.md`（画像承接位设计成为 GAP-01/GAP-05 的定位对象）
- Affected docs（仅回写状态与入口，不改写正文结论）:
  - `AGENTS.md`、`plan.md`、`docs/README.md`、`architecture_map.md`、`docs/phase/README.md`
- Affected code: 无（本阶段不改任何前后端与数据库代码）
- 验收产物：本目录 `tasks.md / checklist.md`（根级回写完成后收口）

## ADDED Requirements

### Requirement: phase13 正式缺口记录必须完整留档本次反思结论

系统 SHALL 在本 spec 中冻结 `phase13-11` 验收后用户反思形成的完整缺口记录，作为 `phase14 /plan` 的直接与唯一上游输入。缺口记录 SHALL 覆盖：缺口背景、缺口清单（GAP-01 ~ GAP-07）、方向结论（CON-01 ~ CON-09）、与既有规则文档的衔接声明。

#### Scenario: phase14 规划者读取缺口记录

- **WHEN** `phase14 /plan` 的规划者需要确定下一阶段主交付方向
- **THEN** 能从本记录直接获得完整的缺口判定、用户已澄清的方向边界与必须优先裁决的开放问题
- **AND** 不需要重新回放本轮对话即可恢复全部关键结论

### Requirement: 缺口清单 GAP-01 ~ GAP-07 必须逐条冻结

以下缺口 SHALL 被逐条留档，每条均注明其在当前实现中的定位证据：

- **GAP-01 目录结构表示不成立**：治理画像以 `canonical_root_files[]`（`path / role / is_required` 平铺数组）承接"理想项目目录结构 + 必须全局文件"，但该形态不是目录结构——无树形层级、无目录节点，人类阅读需要转译，agent 理解精确度无保证，与"接管理想项目目录结构"的初心直接偏差。
- **GAP-02 双清单高度重复**：`canonical_root_files[]` 与 `global_asset_bindings[]` 语义大面积重叠——根级必须文件本就是理想目录结构的组成部分，`structured_summary` 本就是对这些文件的说明解释；拆成两份清单不合理，导致同一信息双份维护。
- **GAP-03 版本演进无基础**：画像提供 `project_profile_version`（`project_governance_profile_v1`）等版本字段，但无演进历史、无变更入口、无依赖版本语义的消费方，属无演进基础的超前设计。
- **GAP-04 阶段字段错位**：`current_phase_name / current_phase_ref / current_phase_status` 三字段属开发推进跟踪信息，被硬塞进治理画像；用户定位此类信息应独立以时间轴形式专门记录与展示（后端持久化、web 展示、agent 可读），而非画像字段。
- **GAP-05 全局性矛盾**：画像以 `repository_id` 为唯一锚点、寄生 `Repository detail`、逐仓库手动维护，与"全局性、跨项目、长效范式"定位直接矛盾——同一套理想范式不应要求用户在每个仓库重复编辑一次。
- **GAP-06 agent 原生友好度未达成**：结构化字段面（数组 + 枚举 + 布尔）对 agent 非原生友好，agent 需转译才能理解目录与文档语义，存在幻觉风险。用户真实 agent IDE 提示词用例证明：agent 实际需要的是"背景文档清单导航（`README / PSCO_0~4 / plan.md / TECH_STACK_BASELINE / PSCO-mvp01~05 / project_skills / project_rules / global_skills / architecture_map / AGENTS.md`）+ 进度说明 + 可回源内容"，而非当前画像字段面。
- **GAP-07 总体判定**：治理画像相对 phase13 使命是深度复杂的耦合设计，而非良好承接；`web / agent` 共享方向本身正确，错在承接形态。

#### Scenario: 缺口逐条可定位

- **WHEN** 后续执行者需要核对任一缺口在当前实现中的证据
- **THEN** 每条 GAP 均能映射到 `governance_profile.proto` 字段面、`00010_phase13_governance_profile.sql` 存储结构或 Repository detail 承接位
- **AND** 不存在无法定位证据的抽象抱怨

### Requirement: 方向结论 CON-01 ~ CON-09 必须冻结为 phase14 输入

以下方向结论 SHALL 冻结为 `phase14 /plan` 的强制上游：

- **CON-01 全局规范实体方向**：在 `Product / Repository / Module / Decision` 四实体之外新增全局规范实体（`Standard`），承接：身份信息、理想目录结构声明（真正的树形结构）、必须全局文档清单（文件 + 角色 + 内容或摘要）、与其他实体的绑定关系；独立于单一 `repository_id` 锚点。
- **CON-02 绑定关系不设限**：`Standard` 通过关系表关联其他实体——关联 `repository` 表明该规范存在实际模板仓库维护、关联 `product` 表明该产品遵守此规范、关联 `decision` 等亦允许；关联对象类型当前不封顶。
- **CON-03 模板仓库复用 Repository**：示范模板仓库本身作为 `Repository` 实例登记（新项目 `git clone / pull` 即用），不新造第二套仓库实体。
- **CON-04 取代不并存**：`Standard` 主线落地时 MUST 同步退役治理画像中与之重叠的承接位（`canonical_root_files[] / global_asset_bindings[]` 等），禁止"新开只读 RPC（如 `psco.standard.v1/GetProductStandard`）而旧画像原样保留"的只加不改倾向——用户已明确判定遗留过时设计不可接受。退役范围、先后顺序与存量数据迁移方式属于 `phase14` 优先裁决范围（见 phase14 进入条件第 3 条），本记录不越权给出实现细节。
- **CON-05 颗粒度必须首先裁决**：全局文档维护颗粒度存在三选项（对应用户原话）——① 仅文件名 / 入口导航（PSCO 只登记"有哪些必须文档及指向"，内容留在源仓库）；② PSCO 维护全面文本内容（文档正文由 PSCO 数据库承接维护）；③ 文本内容直接托管到模板仓库（PSCO 只维护结构与导航，正文以真实仓库为唯一事实源）。`phase14 /plan` MUST 首先在三选项间裁决（可组合），再进入数据模型设计。用户提示词实例同时证明全局文档不止技术栈一类（含 skills / rules / 架构文档 / agent 全局入口）。
- **CON-06 动态演进是本质属性**：理想目录结构与全局文档内容是长期实践总结、跟随探索不断调整，`phase14` 设计 MUST 把"持续沉淀、随时演进"作为一等需求，不允许一次性写死。
- **CON-07 效率底线约束**：PSCO 方案的综合成本 MUST 低于"本地目录 / GitHub 仓库手动维护 + 复制粘贴"基线；任何理论上复杂但实际增加维护负担的设计均不可接受（反假大空约束）。
- **CON-08 阶段推进跟踪独立承接**：开发推进跟踪（当前阶段 / 历史阶段）SHALL 独立成时间轴型承接——后端持久化、web 端时间轴展示、agent 可直接读取；从治理画像字段中剥离。具体承接形态（独立模块还是既有实体的派生视图）由 `phase14 /plan` 裁决，本记录只冻结"必须剥离"的判定。
- **CON-09 agent 交互终态与顺序约束**：终态为 web 端可读写 + agent 经 go 后端可读写；但"先消费、后维护"顺序不变，agent 写回仍不进入 `phase14` 首轮，Git 推进跟踪 / 模板仓库自动接入 / 自动同步继续按既有条件约束。

#### Scenario: 方向结论约束 phase14 规划

- **WHEN** `phase14 /plan` 拟定主交付与数据模型
- **THEN** CON-01 ~ CON-09 全部作为强制边界生效，尤其 CON-04（退役计划）与 CON-05（颗粒度裁决）不得跳过
- **AND** 与任一 CON 冲突的方案需显式说明理由并经用户裁决，不得默认绕过

### Requirement: 与既有规则文档的衔接声明必须明确

本记录 SHALL 显式声明以下衔接关系，避免规则冲突：

1. `dev_plan` §4 非目标第 8 条"第五个业务主实体"是 `phase13` 范围内的非目标，随 `phase13` 收口到期；`phase14 /plan` MUST 重新仲裁 `Standard` 的正式实体地位（业务主实体 vs 全局规范资产实体）。
2. `shared_baseline` 冻结的"治理层禁止做成第五主实体"约束在 `phase13` 范围内继续有效（画像实现确实未长成第五主实体）；`phase14` 引入 `Standard` 属于新阶段新裁决，不构成对 `phase13` 冻结结论的违反。
3. `phase13` 已交付的治理画像实现是已验收交付物，本阶段不实施任何改动；其重叠承接位的改造 / 收缩 / 退役由 `phase14 /plan` 依据 CON-04 裁决。
4. 本缺口记录不改变 `PSCO-mvp05-summarize-feedback.md` 的 mvp0.5 共识地位；`Standard` 方向是对其中"全局规范资产"承接形态的阶段级修正输入。

#### Scenario: 规则冲突预防

- **WHEN** 后续执行者在 `phase13` 旧约束与 `phase14` 新方向之间产生疑问
- **THEN** 能从本衔接声明直接判定各约束的时效范围
- **AND** 不需要重新解释规则优先级

### Requirement: 根级回写与阶段收口必须完成

系统 SHALL 完成以下根级同步，且不留孤岛文档：

- `AGENTS.md`：当前阶段状态更新为 `phase13` 已完成正式收口，登记缺口记录为 `phase14` 直接上游
- `plan.md`：`phase13` 状态位更新（承接 `phase13-11` 复核建议 1），登记下一阶段进入条件
- `docs/README.md`、`architecture_map.md`、`docs/phase/README.md`：登记本 spec 目录入口
- 留档 `phase13` 正式验收与收口入口指向 `phase13_11` acceptance_report

#### Scenario: 收口后无孤岛

- **WHEN** 收口完成后的任意执行者从根级入口出发
- **THEN** 能经 `AGENTS.md / plan.md / docs/README.md` 之一找到本缺口记录与 `phase13_11` 验收报告
- **AND** 本 spec 目录不是孤岛文档

### Requirement: phase14 进入条件必须单值化

`phase14` 的进入条件 SHALL 冻结为：

1. 本缺口记录（GAP-01 ~ GAP-07 + CON-01 ~ CON-09）已冻结
2. `phase13-12` 根级回写与阶段收口完成
3. `phase14 /plan` MUST 优先完成五项裁决后再进入实现拆分：
   - 信息维护颗粒度裁决（CON-05 三选项，可组合）
   - `Standard` 数据模型与 pg 承载方式（含目录树形结构的表示方案、agent 可良好读取理解的合同形态）
   - `Standard` 信息与模板仓库的内容边界（哪些信息由 PSCO 承接、哪些以真实仓库为唯一事实源）
   - 治理画像重叠承接位的退役计划（CON-04，含存量数据迁移）
   - `Standard` 正式实体地位仲裁（业务主实体 vs 全局规范资产实体）
4. 既有顺序约束继续生效：先消费后维护；Git 推进跟踪 / 模板仓库自动接入 / 自动同步 / agent 写回不因 `Standard` 引入自动解锁

#### Scenario: 进入条件可机械判定

- **WHEN** 用户或执行者评估是否可以开启 `phase14 /plan`
- **THEN** 以上 4 条可逐条机械判定，无需主观解释
