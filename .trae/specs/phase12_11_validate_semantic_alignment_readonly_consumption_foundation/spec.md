# phase12-11 完成 `Semantic Alignment & Read-Only Consumption Foundation` 的联调、dogfooding 与反回归验证 Spec

## Why

`phase12-08 / 09 / 10` 已分别完成四实体语义收口、共享只读入口落地与关键页面联动回流，但当前仍缺少一套被冻结为单值协议的正式验收规格。若不把工具链、浏览器验证、Web / agent dogfooding、固定样本、固定入口、固定 `6` 问与失败判定一次性冻结，后续独立执行者仍会因为样本不同、入口不同、解释不同而得出不一致结论。

`phase12-11` 的目标不是继续扩写功能，而是把 `Semantic Alignment & Read-Only Consumption Foundation` 当前阶段的联调、dogfooding 与反回归验证收成一套可机械复跑的正式验收规格。

## What Changes

- 冻结 `buf / go test / frontend build` 的最小工具链验证顺序与通过标准
- 冻结 Web 侧 primary owner 页面与跟随回归页面的浏览器验证矩阵
- 冻结 Web / agent 共享只读消费入口的 dogfooding 协议、固定样本、固定入口与固定 `6` 问
- 冻结基于同一 `repository_id` 的样本解析协议，禁止数据库手工查询、浏览器临场搜索、额外脚本或新增第 `7` 个入口补救
- 冻结本阶段明确不做 schema 重写、MCP / CLI / agent 写回 / 对话入口的边界证据要求
- **BREAKING**：`phase12` 的通过结论不再接受“局部页面看起来对了”“某个入口能回答出来”或“另找入口补齐答案”作为替代证据，必须经由本阶段统一联调、dogfooding 与反回归验证

## Impact

- Affected specs:
  - `phase12_01_freeze_semantic_alignment_scope_success_non_goals`
  - `phase12_05_design_readonly_consumption_shared_entry`
  - `phase12_06_design_frontend_read_path_owner_shared_summary_reread`
  - `phase12_08_land_frontend_four_entity_semantic_alignment`
  - `phase12_09_land_readonly_consumption_shared_entry`
  - `phase12_10_land_key_pages_shared_read_linkage_reread`
  - `docs/phase/phase12_semantic_alignment_and_readonly_consumption_foundation_dev_plan.md`
- Affected code:
  - `proto/` 下与 `buf` 验证相关的合同源
  - `backend/` 下与 `go test`、`ProjectContext`、四实体读取与导出相关的承接位
  - `frontend/` 下与 `/repositories/$repositoryId`、`/products/$productId`、`/modules/$moduleId`、`/decisions/$decisionId`、`/dashboard`、`/onboarding`、`/reviews/daily`、`/reviews/weekly` 相关的 route、page、application、data 承接位
  - `AGENTS.md`
  - `PSCO-mvp05-summarize-feedback.md`

## ADDED Requirements

### Requirement: `phase12-11` 必须冻结单值验收样本、单值入口与单值问题
系统 SHALL 为 `phase12-11` 冻结一组单值机械验收协议，并要求所有联调、dogfooding 与反回归验证都基于同一固定样本、同一固定入口集合与同一固定 `6` 问执行。

固定样本至少冻结为：

- `Repository`：`personal-software-company-os`
- `repository_id`：`ca261521-8daf-4248-8f12-43525326e759`
- `Product`：`PSCO`
- `Module`：`project-context-foundation`
- `Decision`：`phase11 Project Context Foundation dogfooding 验收决策`

固定入口至少冻结为：

- agent 侧：`AGENTS.md`、`PSCO-mvp05-summarize-feedback.md`、基于同一 `repository_id` 读取的共享只读结果（结构化读取或其受控派生视图）
- Web 侧 primary owner：`/repositories/$repositoryId`、`/products/$productId`、`/modules/$moduleId`、`/decisions/$decisionId`
- Web 侧跟随回归页：`/dashboard`、`/onboarding`、`/reviews/daily`、`/reviews/weekly`

固定验收提问必须冻结为：

1. 当前 `Product` 的正式定位是什么，在哪个固定入口可直接确认
2. 当前 `Repository` 的正式定位是什么，在哪个固定入口可直接确认
3. 当前 `Module` 为什么不是普通模块登记对象，在哪个固定入口可直接确认“可复用能力资产”语义
4. 当前 `Decision` 为什么不是孤立文本卡片，在哪个固定入口可直接确认“规则 / 约束 / 选择与依据索引对象”语义
5. 当前项目共享的规则、约束与文档入口从哪里查看，Web 与 agent 是否回到同一组入口
6. 当前只读消费是否仍由同一 `repository_id` 锚点驱动，且没有长出第二套事实源

#### Scenario: 冻结 phase12-11 的正式样本与入口
- **WHEN** 执行者准备开始 `phase12-11` 的联调与 dogfooding
- **THEN** 能直接拿到唯一正式样本、唯一正式入口集合与唯一正式 `6` 问
- **AND** 不需要再临场更换仓库、切换锚点或补造新入口

### Requirement: 样本解析协议必须继续由同一 `repository_id` 锚点驱动
系统 SHALL 将 `phase12-11` 的样本解析协议冻结为：`repository_id` 只允许使用固定样本中的正式值，`product_id / module_id / decision_id` 只允许从同一 `repository_id` 驱动的结构化只读结果或其受控派生视图解析。

当前阶段至少必须明确：

1. `repository_id` 不允许临场改样本或改锚点；
2. `product_id / module_id / decision_id` 不允许绕到数据库手工查询、浏览器临场搜索或额外脚本补齐；
3. 解析链路必须显式留档：使用了哪个固定入口、解析出了什么 id、是否一次成功；
4. 若名称解析失败、结果不唯一或无法回到同一 `repository_id`，则本轮验收直接判定失败；
5. Web 侧详情页允许先通过稳定名称解析 `product_id / module_id / decision_id`，但不允许在页面内临场猜测 `repository_id`。

#### Scenario: 名称解析失败或不唯一
- **WHEN** 执行者无法从固定入口一次解析出唯一 `product_id / module_id / decision_id`
- **THEN** 必须直接判定 `phase12-11` 当前验收失败
- **AND** 不得通过新增第 `7` 个入口或额外脚本补救

### Requirement: 工具链验证必须冻结为单值顺序
系统 SHALL 将 `phase12-11` 的最小工具链验证顺序冻结为单值执行链，而不是“谁想到哪个就先跑哪个”。

正式顺序至少包含：

1. `buf` 相关合同验证
2. `go test` 相关后端验证
3. `frontend build` 前端构建验证

通过标准至少必须明确：

- `buf` 失败时不得以后续测试或浏览器可用替代；
- `go test` 失败时不得以后端局部手测替代；
- `frontend build` 失败时不得以 dev server 可打开替代；
- warning 若不阻断命令退出，可单独记录，但不得篡改通过/失败归类。

#### Scenario: 工具链失败
- **WHEN** 任一步工具链验证失败
- **THEN** 必须判定 `phase12-11` 尚未通过
- **AND** 必须记录失败归属是合同、后端、前端还是环境问题

### Requirement: Web 侧浏览器验证必须覆盖 primary owner 与跟随回归页面
系统 SHALL 在 `phase12-11` 中冻结 Web 侧浏览器验证矩阵，并要求 primary owner 页面与跟随回归页面都能用真实固定入口复验四实体语义、共享只读消费与返回链行为。

当前阶段至少必须覆盖：

- primary owner 页面：`/repositories/$repositoryId`、`/products/$productId`、`/modules/$moduleId`、`/decisions/$decisionId`
- 跟随回归页面：`/dashboard`、`/onboarding`、`/reviews/daily`、`/reviews/weekly`
- 对每个页面至少记录：
  - 使用了哪个固定入口进入；
  - 当前页面如何解释 `Product / Repository / Module / Decision`；
  - 当前页面是否仍保持只读消费边界；
  - 返回链与 reread 是否回到最新共享解释。

#### Scenario: 验证 Web 侧四实体语义
- **WHEN** 执行者按冻结页面清单逐页浏览器验证
- **THEN** 必须能用真实页面回答固定 `6` 问中的 Web 相关部分
- **AND** 若页面仍依赖切片内旧文案、第二套解释链或无法复验的临时入口，则直接判定失败

### Requirement: Web / agent dogfooding 必须回到同一组规则、约束与入口
系统 SHALL 要求 `phase12-11` 在同一轮验收中同时验证 Web 与 agent 两侧的共享只读消费入口，并证明两者能回到同一组规则、约束与文档入口。

当前阶段至少必须验证：

- agent 侧可从 `AGENTS.md`、`PSCO-mvp05-summarize-feedback.md` 与同一 `repository_id` 的共享只读结果回答固定 `6` 问；
- Web 侧可从冻结页面清单回答同一组问题；
- 两侧对 `Product / Repository / Module / Decision` 的正式解释不出现并列版本；
- 两侧对共享规则、约束、文档入口与只读边界的回答可以互相复验。

#### Scenario: 比对 Web 与 agent 的回答
- **WHEN** 执行者对 Web 与 agent 分别执行固定 `6` 问
- **THEN** 必须能指出两侧共用的固定入口与共享只读锚点
- **AND** 不会出现“一边靠页面旧文案，一边靠根级文档补救”的双轨答案

### Requirement: `Module / Decision` 的误读风险必须纳入固定验收
系统 SHALL 将 `Module / Decision` 的误读风险纳入 `phase12-11` 的固定验收问题，而不是只做通用页面巡检。

当前阶段至少必须显式验证：

- `Module` 不是普通模块登记对象，而是“可复用能力资产”；
- `Decision` 不是孤立文本卡片，而是“规则 / 约束 / 选择与依据的索引对象”；
- 两者都必须能在至少一个固定入口被直接确认，而不是依赖执行者主观解释。

#### Scenario: 复核 Module / Decision 的正式语义
- **WHEN** 执行者回答固定 `6` 问中的第 `3`、`4` 题
- **THEN** 必须指出对应的固定入口与可直接确认的正式语义
- **AND** 若仍出现“普通模块”“文本卡片”式误读，则本轮验收失败

### Requirement: 验收记录必须足以让不同执行者按同一协议 rerun
系统 SHALL 要求 `phase12-11` 产出一份单一正式验收记录，并保证不同执行者可以按相同样本、相同入口、相同问题与相同失败判定重新执行。

验收记录至少必须留档：

- 样本解析方式；
- 固定页面 / 固定入口清单；
- 每个问题的回答结果；
- 工具链结果；
- 浏览器与 dogfooding 结果；
- 失败点与是否达标；
- 本阶段明确不做 schema 重写、MCP / CLI / agent 写回 / 对话入口的边界证据。

#### Scenario: 第二执行者复跑 phase12-11
- **WHEN** 第二位执行者只拿到 `phase12-11` 的正式验收记录与 spec 包
- **THEN** 能按相同问题、相同样本与相同入口 rerun
- **AND** 不需要再发明额外入口、额外问题或额外解释模板

## MODIFIED Requirements

### Requirement: `phase12` 的通过条件在 phase12-11 上收口为统一联调、dogfooding 与反回归验证
`phase12` 的通过条件 SHALL 从“`phase12-08 / 09 / 10` 已分别落地”推进为“`Semantic Alignment & Read-Only Consumption Foundation` 已在同一 `repository_id` 锚点下通过统一工具链、浏览器、Web / agent dogfooding 与关键反回归验证”。

至少必须同时满足：

- Web 与 agent 对四实体语义的回答可被真实固定入口复验；
- 当前只读消费仍保持单一结构化输入锚点与只读边界；
- 共享只读 owner 已单值化；
- 同一份验收记录足以被不同执行者复跑。

#### Scenario: 判断 phase12 是否可正式收口
- **WHEN** 读者评估 `phase12` 当前是否可进入根级收口
- **THEN** 必须同时参考 `phase12-11` 冻结的工具链、浏览器矩阵、dogfooding 协议、固定 `6` 问与边界证据
- **AND** 不得仅凭局部源码修改或单页浏览器点击给出“通过”结论

## REMOVED Requirements

### Requirement: 允许通过临时新增入口、临时样本或临时问题补齐验收答案
**Reason**: 这会破坏 `phase12-11` 的单值验收协议，使不同执行者继续依赖临场判断，无法保证 Web / agent 与共享只读入口的结论可复验。
**Migration**: 改为只允许使用冻结样本、冻结入口、冻结 `6` 问与冻结失败判定执行验收；若任一环节失败，直接记录失败并回到实现修复。
