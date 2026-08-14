# phase10-11 完成 `Asset-Action Closure` 联调、浏览器验收与反回归验证 Spec

## Why

`phase10-08 / 09 / 10` 已分别把 `Onboarding` 首轮建链、`Decision` 生命周期闭环、以及关键 detail pages 的下一步动作承接矩阵落到了实现层，但当前仍缺少一套被冻结为单值场景的机械验收规格。若不把工具链、浏览器链路、反回归路径、固定前置数据与明确不做边界一次性冻结，后续独立验收者仍会因为测试数据不同、入口不同、理解不同而得出不一致结论。

`phase10-11` 的目标不是继续扩写业务能力，而是把 `Asset-Action Closure` 当前阶段的联调、浏览器验收、反回归验证与边界证据收成一套可机械复现的正式验收规格。

## What Changes

- 冻结 `phase10` 当前阶段的单值机械验收前置数据与场景定义
- 冻结 `buf / go test / frontend build` 的正式工具链验收顺序与通过标准
- 冻结 `Onboarding / Dashboard / Daily Review / Decision Detail / Product Detail / Module Detail / Repository Detail` 的浏览器验收矩阵
- 冻结 `Current Focus / pending signals` 的反回归验证矩阵
- 冻结当前阶段明确不做 `Agent Consumption Layer / 新实体回归 / 第五态状态机` 的边界证据
- 将上述内容整理为后续独立验收者可直接执行的 spec 三件套

## Impact

- Affected specs:
  - `phase10_08_land_onboarding_first_run_chain_guidance_canonical_handoff`
  - `phase10_09_land_decision_lifecycle_detail_cta_pending_reread_unification`
  - `phase10_10_land_key_detail_pages_next_step_cta_handoff_matrix`
  - `docs/phase/phase10_asset_action_closure_foundation_dev_plan.md`
- Affected code:
  - `backend/` 中与 `buf / go test` 验收相关的 proto、service、query、command、candidate 承接位
  - `frontend/` 中与 `Onboarding / Dashboard / Review / Decision Detail / Product Detail / Module Detail / Repository Detail` 浏览器链路相关的 route、page、application、data 承接位
  - 浏览器级验收过程中实际使用到的固定前置数据、测试样本与最小 canonical 缺口组合

## Frozen Acceptance Baseline

### Isolated Acceptance Environment

- isolated database: `psco_phase10_11_e2e`
- isolated backend: `http://127.0.0.1:8082`
- isolated frontend: `http://127.0.0.1:4173`
- frontend API base URL: `VITE_API_BASE_URL=http://127.0.0.1:8082`
- 验收期间允许保留历史数据作为背景噪音，但正式入口、正式 pending 样本与 detail 浏览器链路只允许使用下列单值实体

### Frozen Single-Value Samples

- `Product Detail` 正式实体：
  - `phase10-11-product`
  - `product_id = 658ba43c-70de-42ec-8409-a8dccdbcdd0d`
- `Repository Detail` 正式实体：
  - `phase10-11-repo`
  - `repository_id = 4c094d74-0641-4c83-a787-c01670a0785b`
- `Module Detail` 正式实体：
  - `phase10-module-0814-1`
  - `module_id = f5a38fb9-1820-42a7-82d4-e89db5332ed7`
- `Decision pending` 正式样本：
  - `phase10 decision 0814-1`
  - `decision_id = 2778f553-e70d-4444-8337-631618f06212`
  - 机械验收起点状态：`proposed`
  - 验收完成后状态：`active`

### Frozen Toolchain Sequence

正式工具链顺序冻结为：

1. `proto/`: `make build && make lint`
2. `backend/`: `go test ./...`
3. `frontend/`: `npm run build`

通过标准冻结为：

- `buf build / buf lint` 必须通过，不允许以“浏览器可用”跳过合同失败
- `go test ./...` 必须通过，不允许以后端局部手测替代
- `npm run build` 必须通过，不允许以 dev server 运行替代正式构建
- warning 若不阻断命令退出，可单独记录，但不得篡改通过/失败归类

### Frozen Browser Acceptance Record

- `Onboarding` 正式入口冻结为 `Dashboard -> Onboarding`，并已在隔离环境中按 `welcome -> product -> repository -> module -> decision -> complete` 完整走通，最终返回 `Dashboard`
- `Decision pending` 正式入口冻结为 `Dashboard / Current Focus`；同一 `decision_id = 2778f553-e70d-4444-8337-631618f06212` 已完成 `proposed -> active` 推进，返回 `Dashboard` 后 pending 已消失，`Daily Review` 同步回读为无 pending 的空态解释
- `Weekly Review` 已作为 detail pages 的统一来源返回样本完成真实浏览器回归：
  - `Weekly Review -> Product Detail -> return`
  - `Weekly Review -> Module Detail -> Decision Detail -> return`
  - `Weekly Review -> Repository Detail -> Decision Detail -> return`
- `Repository Detail` 已验证真实动作承接位，而不是壳层 CTA：
  - 首轮映射模块后出现成功提示
  - reread 后页面级“下一步”切换为“查看相关决策”
  - 继续进入 `Decision Detail` 时保留 `fromReview / reviewKind / reviewReturnTo`
- `Current Focus / pending signals` 反回归冻结为：
  - 关键动作完成后必须回看 `Dashboard Current Focus`
  - pending 样本处理完成后不得继续残留到 `Dashboard / Daily Review`
  - 若仍指向非 detail canonical path，则直接判定 `phase10-11` 未通过

### Frozen Non-Goals

- `Agent Consumption Layer` 不属于本轮机械验收范围
- 新实体回归不属于本轮机械验收范围
- 第五态状态机不属于本轮机械验收范围
- `phase10-11` 只负责验证 `phase10-08 / 09 / 10` 已落地能力的联调闭环，不借验收过程扩写新业务能力

## ADDED Requirements

### Requirement: `phase10-11` 必须冻结单值验收前置数据，而不是临场造数

系统 SHALL 为 `phase10-11` 冻结一组单值机械验收前置数据，并将其作为所有联调、浏览器验收与反回归验证的唯一正式样本来源。后续独立验收者不得再自行拼装“差不多可测”的数据集替代。

#### Scenario: 冻结“空态或近空态”样本

- **WHEN** 定义 `Onboarding` 首轮建链浏览器验收前置数据
- **THEN** spec 必须单值写明：
  - 最小 Product 数量
  - 最小 Repository 数量
  - 最小 Module 数量
  - 最小 Decision 数量
  - 允许保留的历史脏数据范围
- **AND** 必须明确这是“空态”还是“近空态”
- **AND** 不得只写“准备一套基础数据”这类不可机械执行的表述

#### Scenario: 冻结“明确结构缺口”样本

- **WHEN** 定义 `Product / Module / Repository Detail` 的动作链浏览器验收
- **THEN** spec 必须单值写明至少三类 canonical 结构缺口组合：
  - `Product` 缺仓库、缺模块、或两者同时缺失
  - `Module` 缺 Product 绑定、缺 Repository 映射、或两者之一缺失
  - `Repository` 缺 Product 绑定、缺 Module 映射、或两者之一缺失
- **AND** 每类缺口都必须指向具体最小样本，而不是泛指“找一条有问题的数据”

#### Scenario: 冻结 `Decision pending` 样本

- **WHEN** 定义 `Decision` 生命周期与 pending reread 验收样本
- **THEN** spec 必须单值写明：
  - 至少一条 `status = proposed` 的最小样本
  - 该样本的 canonical 入口来源
  - 该样本在 `Dashboard / Daily Review / Current Focus` 中应出现的位置
- **AND** 不得允许同一轮验收临场更换为别的 `Decision`

### Requirement: 工具链验收必须冻结为单值顺序

系统 SHALL 将 `phase10-11` 的工具链验收顺序冻结为单值执行链，而不是“谁想到哪个就先跑哪个”。

#### Scenario: 工具链执行顺序

- **WHEN** 执行 `phase10-11` 工具链验收
- **THEN** 正式顺序必须至少包含：
  1. `buf` 相关合同验证
  2. `go test` 相关后端验证
  3. `frontend build` 前端构建验证
- **AND** 每一步都必须定义“通过 / 失败 / 可接受 warning”的判断口径
- **AND** 不得把某一步失败后“先跳过，等会儿再看”写成正式验收流程

#### Scenario: 工具链失败的处理语义

- **WHEN** 任一步工具链验证失败
- **THEN** 必须判定 `phase10-11` 尚未通过
- **AND** 必须明确记录失败归属是合同、后端、前端还是环境问题
- **AND** 不得在 spec 中把构建失败解释为“浏览器能开起来就先算通过”

### Requirement: `Onboarding` 首轮建链浏览器验收必须机械化

系统 SHALL 将 `Onboarding` 首轮建链浏览器验收冻结为单值、机械可复现的步骤序列，覆盖冷启动、逐步推进、成功 handoff、返回恢复与完成态。

#### Scenario: `Dashboard -> Onboarding` 首轮建链

- **WHEN** 用户从已冻结的“空态或近空态”场景进入 `Dashboard`
- **THEN** spec 必须明确从 `Dashboard` 进入 `Onboarding` 的唯一正式入口
- **AND** 必须逐步验证：
  - `welcome -> product -> repository -> module -> decision -> complete`
  - 每一步完成后默认下一步动作是否符合冻结矩阵
- **AND** 不得把“直接访问 `/onboarding`”与“从 `Dashboard` 进入 `Onboarding`”混为同一验证项

#### Scenario: `Onboarding` 每一步默认下一步动作

- **WHEN** 浏览器验收执行到 `product / repository / module / decision` 任一步
- **THEN** spec 必须明确当前页面应展示的默认下一步动作
- **AND** 必须写清是 stay-on-page、进入 canonical handoff，还是进入下一正式 step
- **AND** 后续验收者不得再自己判断“这一步看起来下一步应该去哪”

#### Scenario: `Onboarding` 完成后的回流理解

- **WHEN** 用户完成首轮建链并回到 canonical owner 或完成态
- **THEN** 页面必须能让用户看懂刚刚完成了什么
- **AND** spec 必须要求验证成功回流文案、来源返回与 reread 结果

### Requirement: `Decision` 生命周期与 pending reread 必须有真实页面机械验证

系统 SHALL 将 `Decision` 从 `proposed` 推进后，`Dashboard / Daily Review / Current Focus` 的 reread 行为冻结为机械验收矩阵，而不是仅依赖源码推断。

#### Scenario: `Decision Detail` 推进后 reread

- **WHEN** 用户从冻结的 pending 样本进入 `Decision Detail`
- **AND** 将该 `Decision` 从 `proposed` 推进到目标状态
- **THEN** spec 必须要求逐项验证：
  - 详情页 CTA 是否切换
  - 返回来源页后该 pending 是否消失
  - `Dashboard / Daily Review / Current Focus` 是否同步 reread

#### Scenario: pending 反回归

- **WHEN** 完成 `Decision` 生命周期浏览器验收
- **THEN** spec 必须要求额外回看：
  - `Dashboard` 的 `Current Focus`
  - `Daily Review` 的 pending 区块
  - `Current Focus / pending signals` 是否仍错误残留
- **AND** 不得只验证一处入口就宣布 pending reread 通过

### Requirement: 关键 detail pages 的动作链浏览器验收必须单值化

系统 SHALL 将 `Product Detail / Module Detail / Repository Detail` 的动作链浏览器验收冻结为单值矩阵，验证“下一步做什么”“是否进入 canonical path”“返回原入口后是否能看懂刚刚发生了什么”。

#### Scenario: `Product Detail`

- **WHEN** 浏览器验收进入冻结的 `Product Detail` 样本
- **THEN** 必须逐项验证：
  - 页面级主 CTA 是否明确
  - CTA 是否优先指向 canonical path
  - 完成动作后返回原入口时，来源页是否 reread 并能解释刚刚完成的结构闭合

#### Scenario: `Module Detail`

- **WHEN** 浏览器验收进入冻结的 `Module Detail` 样本
- **THEN** 必须逐项验证：
  - 页面级主 CTA 是否明确
  - handoff 是否进入 canonical `Product Detail / Repository Detail / Decision Detail`
  - 返回原入口时是否保留来源语义并 reread

#### Scenario: `Repository Detail`

- **WHEN** 浏览器验收进入冻结的 `Repository Detail` 样本
- **THEN** 必须逐项验证：
  - 页面级主 CTA 是否明确
  - CTA 是否真正进入动作承接位，而不只是滚动或显示壳层
  - 与 `Decision Detail` 的链路是否继续保留来源返回语义

### Requirement: `Current Focus / pending signals` 必须纳入正式反回归矩阵

系统 SHALL 将 `Current Focus / pending signals` 的反回归验证纳入 `phase10-11` 的正式验收矩阵，防止局部修复破坏整体解释。

#### Scenario: Current Focus 反回归

- **WHEN** 任一关键动作完成并返回 `Dashboard / Daily Review`
- **THEN** spec 必须要求回看 `Current Focus` 是否切换到新的 canonical 下一步
- **AND** 不得继续把已经闭合的结构缺口或已处理的 pending decision 当作当前主动作

#### Scenario: pending signals 反回归

- **WHEN** 完成 `Decision` 或 detail page 动作链验证
- **THEN** spec 必须要求回看 pending signal 列表、信号文案与跳转目标
- **AND** 确认没有因为回归导致重新指向非 canonical path

### Requirement: 非目标边界证据必须被正式留档

系统 SHALL 在 `phase10-11` 规格中正式留档当前阶段明确不做的边界证据，防止验收范围被不断追加。

#### Scenario: 明确不做的边界

- **WHEN** 完成 `phase10-11` spec
- **THEN** 必须明确写出当前阶段不做：
  - `Agent Consumption Layer`
  - 新实体回归
  - 第五态状态机
- **AND** 必须写明这些事项为何不属于本阶段机械验收范围
- **AND** 不得把“以后可能做”写成“本次验收顺手带一下”

## MODIFIED Requirements

### Requirement: `phase10` 当前阶段的“通过”解释必须从“局部子任务通过”升级为“联调与浏览器总体验收通过”

系统 SHALL 将 `phase10` 当前阶段的通过口径修改为：不仅要求 `phase10-08 / 09 / 10` 各自局部实现成立，还要求它们在单值前置数据下通过工具链、浏览器动作链与关键反回归验证。

#### Scenario: 判断 `phase10` 当前阶段是否可收口

- **WHEN** 后续验收者判断当前阶段是否可以收口
- **THEN** 不能只凭某个子任务构建通过或某条单链浏览器通过
- **AND** 必须同时参考 `phase10-11` 冻结的工具链、浏览器矩阵、反回归矩阵与单值样本

### Requirement: 验收步骤必须足够细到无需独立验收者补造主测试路径

系统 SHALL 将当前阶段机械验收步骤修改为：每一条关键链路都必须有明确入口、明确样本、明确动作、明确预期结果与明确回看点。

#### Scenario: 独立验收者执行 spec

- **WHEN** 一个不了解隐性上下文的独立验收者只阅读 `phase10-11` 的 spec 三件套
- **THEN** 其应能机械执行主要验收路径，而无需再自己推导：
  - 用哪条数据
  - 从哪个页面进入
  - 点哪一个按钮
  - 返回后看哪里

## REMOVED Requirements

### Requirement: 允许验收者临场拼装“相似场景”代替正式样本

**Reason**: 这会让 `Onboarding`、`Decision`、detail pages 与 pending reread 的结论建立在不同测试集之上，导致“同一子任务、不同人、不同结论”的不稳定状态。
**Migration**: 改为冻结单值前置数据、结构缺口样本与 pending decision 样本，所有后续验收统一复用。

### Requirement: 将“源码阅读 + 局部构建成功”视为本阶段充分验收

**Reason**: `Asset-Action Closure` 当前阶段的核心风险已经从静态设计转移到跨页面动作链、来源返回与 reread 行为，单纯静态阅读或局部构建无法证明闭环成立。
**Migration**: 改为工具链 + 浏览器动作链 + 关键反回归三层共同组成正式验收口径。
