# phase13-11 完成 `Project Governance Profile Foundation` 的联调、dogfooding 与反回归验证 Spec

## Why

`phase13-08 / 09 / 10` 已分别完成治理画像后端写读主线、前端承接与手工维护入口、agent 项目简报读取主线，但当前仍缺少一套被冻结为单值协议的正式验收规格。若不把固定样本、固定入口、固定 `6` 问、固定 rerun 记录格式与失败判定一次性冻结，后续独立执行者仍会因为样本不同、入口不同、解释不同而得出不一致结论；且固定样本当前在共享开发库中因集成测试切库已丢失，验收环境恢复路径也必须一并冻结。

`phase13-11` 的目标不是继续扩写功能，而是把 `Project Governance Profile Foundation` 当前阶段的联调、dogfooding 与反回归验证收成一套可机械复跑的正式验收规格。

## What Changes

- 冻结验收环境与固定样本恢复协议：仅允许使用仓库内正式恢复脚本，禁止手工 SQL 插样本或改锚点
- 冻结人类维护路径 dogfooding 协议：治理画像数据必须经 `Repository detail` 表单手工维护创建，禁止种子化
- 冻结 agent 读取路径 dogfooding 协议：`GetProjectBrief` 实时取证 + `AGENTS.md` / `plan.md` 固定入口核对
- 冻结固定 `6` 问与逐题留档格式（answer / direct entry refs / 是否达标）
- 冻结 `buf build / buf lint / go test / npm run build` 的最小工具链验证顺序与通过标准
- 冻结 `GetProjectBrief` 自动化集成测试补齐要求（承接 `phase13-10` 复核留档建议 1）
- 冻结浏览器反回归矩阵：主验证页 `/repositories/$repositoryId` + 四实体与跟随回归页面
- 冻结本阶段明确不做 Git 推进跟踪、模板仓库接入、自动同步与 agent 写回的边界证据要求
- **BREAKING**：`phase13` 的通过结论不再接受"画像接口能返回""表单能打开"或"另找入口补齐答案"作为替代证据，必须经由本阶段统一联调、双侧 dogfooding 与反回归验证

## Impact

- Affected specs:
  - `phase13_08_land_project_governance_profile_backend_mainline`
  - `phase13_09_land_frontend_governance_profile_manual_maintenance_entry`
  - `phase13_10_land_agent_project_brief_reading_mainline`
  - `docs/phase/phase13_project_governance_profile_foundation_dev_plan.md`
  - `docs/phase/phase13_project_governance_profile_foundation_shared_baseline.md`
- Affected code:
  - `backend/internal/projectcontext/`（新增 `GetProjectBrief` 集成测试用例）
  - `database/scripts/restore_phase11_phase12_dogfooding_sample.sh`（仅执行，不修改）
  - `frontend/`（仅浏览器验证与构建，不修改源码）
  - 固定 Web 验证页面 `/repositories/$repositoryId`
- 验收产物：本目录 `acceptance_report.md`

## ADDED Requirements

### Requirement: `phase13-11` 必须冻结单值验收样本、单值入口与单值问题

系统 SHALL 为 `phase13-11` 冻结一组单值机械验收协议，并要求所有联调、dogfooding 与反回归验证都基于同一固定样本、同一固定入口集合与同一固定 `6` 问执行。

固定样本冻结为（继承 `phase11 / phase12` dogfooding 样本）：

- `Repository`：`personal-software-company-os`
- `repository_id`：`ca261521-8daf-4248-8f12-43525326e759`
- `Product`：`PSCO`

固定 Web 验证页面第一版冻结为：

- `/repositories/ca261521-8daf-4248-8f12-43525326e759`

固定 agent 读取入口第一版冻结为：

1. 基于同一 `repository_id` 返回的 `project brief for agent`（`GetProjectBrief`）
2. `AGENTS.md`
3. `plan.md`
4. 由治理画像承接的全局规范资产结构化结果（`GetProjectBrief.global_assets` 或 `GetGovernanceProfile.global_asset_bindings`）

固定验收提问冻结为（`dev_plan` L232-238）：

1. 当前项目治理画像版本与技术路线是什么，在哪个固定入口可确认
2. 当前 canonical 根级文件集合是否已被正式承接，在哪个固定入口可确认
3. 当前全局规范资产是否以结构化摘要 + 入口关系被正式承接，在哪个固定入口可确认
4. 当前 agent 项目简报是否由同一 `repository_id` 驱动，且没有伪造第二套目录扫描结果
5. 当前第一版前端正式承接位是否仍是 `Repository detail`，而没有长出并列第二入口
6. 当前阶段是否仍严格没有进入 Git 推进跟踪、模板仓库接入、自动同步与 agent 写回

#### Scenario: 冻结 phase13-11 的正式样本与入口

- **WHEN** 执行者准备开始 `phase13-11` 的联调与 dogfooding
- **THEN** 能直接拿到唯一正式样本、唯一正式入口集合与唯一正式 `6` 问
- **AND** 不需要再临场更换仓库、切换锚点或补造新入口

### Requirement: 验收环境与固定样本恢复必须走正式脚本路径

系统 SHALL 将 `phase13-11` 的验收环境恢复协议冻结为：

1. 固定样本当前在共享开发库中丢失（`projectcontext` 集成测试切库所致，`phase12-11` 已有同款先例），必须先执行仓库内正式恢复脚本 `database/scripts/restore_phase11_phase12_dogfooding_sample.sh` 补齐 dogfooding 样本
2. 样本恢复后必须显式验证固定 `repository_id` 一次解析成功（兼容层 `GetProjectContext` 或 `GetProjectBrief` 的 repository 块均可作为解析证据）
3. 恢复完成后，画像未创建状态下 `GetProjectBrief` 应返回 `not_found`（`governance profile not created`），这是预期行为而非失败
4. 禁止通过手工 SQL 插入样本、临场改 `repository_id` 或用其他仓库样本替代

#### Scenario: 固定样本不在库中

- **WHEN** 执行者发现固定 `repository_id` 解析失败
- **THEN** 必须先通过正式恢复脚本恢复样本
- **AND** 不得手工插库或更换样本锚点

### Requirement: 人类维护路径 dogfooding 必须经 Web 表单完成画像创建

系统 SHALL 将 `phase13-11` 人类维护路径 dogfooding 冻结为：固定样本的治理画像数据必须通过 `/repositories/$repositoryId` 的治理画像手工维护表单创建与维护，而不是通过种子 SQL 或后门脚本。

当前阶段至少必须完成：

1. 经表单创建画像并维护 `template_source`、`canonical_root_files[]`（对齐当前项目范式 v1 根级文件集合）、`global_asset_bindings[]`（8 项资产全覆盖，前 5 项摘要型资产填写 `structured_summary`，后 3 项允许为空）
2. 验证概览区只读字段与后端 `RootFrozen*` 冻结常量一致：
   - `track_type = durable_system`
   - `docs_workflow_layout = phase/fix/audit/review`
   - `current_phase_name = phase13_project_governance_profile_foundation`
   - `current_phase_ref = plan.md#phase13_project_governance_profile_foundation`
   - `current_phase_status = in_progress`
3. 验证保存成功回流（无整页刷新的精准失效刷新）与刷新后回看数据一致
4. 维护完成后画像数据成为后续 agent 读取路径 dogfooding 的同源事实

#### Scenario: 画像数据从种子导入

- **WHEN** 执行者试图用种子 SQL 或脚本直接向治理画像表插数据来跳过表单维护
- **THEN** 必须判定为违背"手工维护优先"dogfooding 语义
- **AND** 本轮验收不得通过

### Requirement: agent 读取路径 dogfooding 必须以 brief 实时结果取证

系统 SHALL 将 `phase13-11` agent 读取路径 dogfooding 冻结为：

1. 对固定 `repository_id` 实时调用 `GetProjectBrief`，取证 7 顶层块完整性：`repository / governance_profile / global_assets / current_phase / products[] / modules[] / decisions[]`
2. 验证 brief 与 `GetGovernanceProfile` 读取结果同源（治理画像字段一致）
3. 验证数组摘要语义：`products / modules / decisions` 保持数组形态，即使长度为 1
4. 验证 `current_phase` 从治理画像主记录三 read-only 字段单向派生
5. 验证 `AGENTS.md` 推荐阅读顺序与 `plan.md` 阶段路线仍能回答 `phase13` 当前状态
6. 验证 brief 中无目录扫描结果、无第二套事实源投影、无自然语言指导词字段

#### Scenario: brief 数据与画像不同源

- **WHEN** 执行者比对 `GetProjectBrief.governance_profile` 与 `GetGovernanceProfile` 结果出现字段不一致
- **THEN** 必须判定本轮验收失败
- **AND** 不得以"接口都能返回"作为通过证据

### Requirement: 固定 `6` 问必须逐题留档且以固定格式回答

系统 SHALL 要求 `phase13-11` 的固定 `6` 问逐题回答留档，回答格式冻结为：`answer / direct entry refs / 是否达标`。

至少必须满足：

1. 每题的 answer 必须能被固定入口（Web 页面、brief 实时结果、`AGENTS.md`、`plan.md`）直接复验
2. 每题必须给出直接入口引用（页面区块、brief 字段或文档行号）
3. 任一题不达标即判定本轮验收失败，不得以其余题目达标折算通过
4. `6 / 6` 达标是 `phase13-11` 通过的必要条件

#### Scenario: 某题无法由固定入口复验

- **WHEN** 固定 `6` 问中任何一题只能靠执行者主观解释作答
- **THEN** 该题判定为不达标
- **AND** 本轮 `phase13-11` 验收失败

### Requirement: 工具链验证必须冻结为单值顺序

系统 SHALL 将 `phase13-11` 的最小工具链验证顺序冻结为单值执行链：

1. `proto/` 下 `buf build` 与 `buf lint`
2. `backend/` 下 `go test ./...`（须包含本轮新增的 `GetProjectBrief` 集成测试）
3. `frontend/` 下 `npm run build`

通过标准至少必须明确：

- `buf` 失败时不得以后续测试或浏览器可用替代
- `go test` 失败时不得以后端局部手测替代
- `npm run build` 失败时不得以 dev server 可打开或 `tsc --noEmit` 通过替代
- warning 若不阻断命令退出，可单独记录，但不得篡改通过/失败归类

#### Scenario: 工具链失败

- **WHEN** 任一步工具链验证失败
- **THEN** 必须判定 `phase13-11` 尚未通过
- **AND** 必须记录失败归属是合同、后端、前端还是环境问题

### Requirement: 必须补齐 `GetProjectBrief` 自动化集成测试

系统 SHALL 要求 `phase13-11` 在 `backend/internal/projectcontext/connect` 既有集成测试中补齐 `GetProjectBrief` 用例，承接 `phase13-10` 复核留档建议 1。

至少必须覆盖：

1. 固定 `repository_id` 下画像已创建 → 7 顶层块完整返回、`current_phase` 从主记录派生、数组语义正确
2. 画像未创建 → `not_found`
3. 治理画像与 brief 的同源性由断言保证（关键字段值一致）

补充冻结：

- 测试内画像准备必须走 service 正式写路径（`UpdateGovernanceProfile`）或 store 正式入口，不得绕过 service 校验直接构造半套状态
- 测试不得依赖共享开发库中的 dogfooding 样本（须自建 fixture 或测试内创建），避免再次出现"集成测试切库导致样本丢失"问题
- 本任务不改变任何生产代码行为，只新增测试

#### Scenario: 新增测试暴露实现缺陷

- **WHEN** 新增集成测试运行失败
- **THEN** 必须先定位是实现缺陷还是测试缺陷并修复
- **AND** 修复后必须完整重跑工具链，不得只重跑单个用例

### Requirement: 浏览器反回归矩阵必须覆盖主验证页与跟随回归页

系统 SHALL 冻结 `phase13-11` Web 侧浏览器验证矩阵：

主验证页（治理画像唯一正式承接位）：

- `/repositories/ca261521-8daf-4248-8f12-43525326e759`
  - 治理画像区三部分齐全：概览卡片、手工维护表单区、摘要回看区
  - 概览区 read-only 字段值正确
  - 摘要回看区显示 `structured_summary + entry_ref`，真实路径不成为主视觉

跟随回归页（四实体主线与关键页面不回归）：

- `/dashboard`
- `/modules` 列表与固定模块详情
- `/decisions` 列表与固定决策详情
- `/products` 列表与固定产品详情
- `/repositories` 列表
- `/onboarding`
- `/reviews/daily`、`/reviews/weekly`

对每个页面至少记录：正常加载、无 `not_found`、无第二套治理画像入口、四实体语义与共享只读消费不回归。

#### Scenario: 治理画像区出现在其他页面

- **WHEN** 浏览器验证发现 `Product / Module / Decision / Dashboard / Review` 页面长出治理画像主内容区或并列第二入口
- **THEN** 直接判定 `phase13-11` 失败（违背 `phase13-09` 唯一承接位冻结）

### Requirement: 边界证据必须显式留档

系统 SHALL 要求 `phase13-11` 留档以下边界证据，证明本阶段严格没有进入非目标范围：

1. 不做 Git 推进跟踪主线
2. 不做模板仓库接入 / bootstrap / clone / pull
3. 不做自动同步 / 目录全文扫描入库
4. 不做 agent 写回 / MCP / CLI / Draft / 审批流

证据来源至少包括：

- 后端 RPC 清单中治理画像与 brief 均为只读或既有手工写入口，无新增自动化同步接口
- `GetProjectBrief` 无写语义、无目录扫描字段
- 数据库迁移无目录扫描 / Git 状态 / 模板版本类表或列
- 根级与阶段文档的明确不做清单引用

#### Scenario: 边界证据缺失

- **WHEN** 验收记录无法给出四类非目标的直接证据
- **THEN** 本轮验收不得通过

### Requirement: 验收记录必须足以让不同执行者按同一协议 rerun

系统 SHALL 要求 `phase13-11` 产出单一正式验收记录 `acceptance_report.md`，rerun 记录格式至少留档：

1. 使用了哪个固定样本与 `repository_id`
2. 使用了哪个固定 Web 页面与 agent 入口
3. 样本恢复方式与画像维护方式（脚本名 + 表单路径）
4. 固定 `6` 问每题回答结果（answer / direct entry refs / 是否达标）
5. 工具链逐步结果（含命令、目录、结果、非阻断输出归类）
6. 浏览器矩阵逐页结果
7. 边界证据清单
8. 失败点与是否达标（含过程中出现的样本漂移、恢复路径与最终 rerun 结果）

#### Scenario: 第二执行者复跑 phase13-11

- **WHEN** 第二位执行者只拿到 `phase13-11` 的 spec 包与验收记录
- **THEN** 能按相同样本、相同入口、相同问题、相同工具链顺序 rerun
- **AND** 不需要再发明额外入口、额外问题或额外解释模板

## MODIFIED Requirements

### Requirement: `phase13` 的通过条件在 phase13-11 上收口为统一联调、双侧 dogfooding 与反回归验证

`phase13` 的通过条件 SHALL 从"`phase13-08 / 09 / 10` 已分别落地"推进为"`Project Governance Profile Foundation` 已在同一 `repository_id` 锚点下通过统一工具链、Web / agent 双侧 dogfooding、固定 `6` 问取证与浏览器反回归验证"。

至少必须同时满足：

- Web 与 agent 都能回到同一组项目治理画像与规范资产
- 当前项目范式 v1 能被 PSCO 稳定承接（canonical 根级文件集合 + 8 项全局规范资产经表单正式维护）
- 人类维护路径与 agent 读取路径均以真实操作取证，而非接口存在性取证
- 同一份验收记录足以被不同执行者复跑

#### Scenario: 判断 phase13 是否可正式收口

- **WHEN** 读者评估 `phase13` 当前是否可进入根级收口（`phase13-12`）
- **THEN** 必须同时参考 `phase13-11` 冻结的工具链、双侧 dogfooding 协议、固定 `6` 问、反回归矩阵与边界证据
- **AND** 不得仅凭局部源码修改或单接口可用给出"通过"结论

## REMOVED Requirements

### Requirement: 允许通过种子 SQL、后门脚本或临时样本补齐验收前提

**Reason**: 这会破坏 `phase13-11` 的单值验收协议与 `phase13-08` 已冻结的"手工维护优先"边界；用种子直接灌画像会让人类维护路径 dogfooding 失真，且不同执行者无法按同一协议复跑。

**Migration**: 固定样本丢失时只能走仓库内正式恢复脚本；治理画像数据只能经 `/repositories/$repositoryId` 表单维护创建；任一环节失败直接记录失败并回到实现修复。
