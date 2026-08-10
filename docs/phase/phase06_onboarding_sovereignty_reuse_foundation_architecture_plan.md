# phase06_onboarding_sovereignty_reuse_foundation_architecture_plan

## 1. 文档定位

本文档是 `phase06_onboarding_sovereignty_reuse_foundation` 的架构规划文档。

> 收口说明：`phase06` 已完成 `/plan -> /spec -> 实现 -> 验收 -> 收口`。本文保留为该阶段 `/plan` 的架构规划与冻结记录，不承担根级当前阶段状态说明；当前阶段状态以 `AGENTS.md`、`plan.md` 与 `docs/README.md` 为准。

`phase05` 已完成 `/plan -> /spec -> 实现 -> 验收 -> 收口`，当前阶段作为 `mvp0.2` 的首个正式交付型 phase，负责把 `PSCO-mvp02-summarize-feedback.md` 中冻结的“阶段一：Onboarding + 数据主权 + 复用感知基础”转化为可继续进入 `/spec` 与实现的正式阶段入口。

当前阶段不是继续扩写新实体，而是优先闭合三个缺口：

- 冷启动与首轮录入摩擦过高
- 数据导出 / 基础备份仍未闭合
- 复用感知与能力派生仍未成为可见反馈

## 2. 上游输入

本阶段直接上游输入如下：

1. `AGENTS.md`
2. `plan.md`
3. `TECH_STACK_BASELINE.md`
4. `project_rules.md`
5. `architecture_map.md`
6. `PSCO-summarize-feedback.md`
7. `PSCO-mvp02-summarize-feedback.md`
8. `.trae/specs/phase05_10_dashboard_feedback_formal_spec/dashboard_feedback_spec_v0.1.md`
9. `.trae/specs/phase05_11_dashboard_feedback_proto_mainline/`
10. `.trae/specs/phase05_14_dashboard_feedback_integration_validation_acceptance/acceptance_report.md`

补充说明：

- 当前阶段直接承接 `phase05` 已交付的 `Dashboard + Feedback` 主线
- `phase02 ~ phase04` 的既有边界继续经由 `phase05` 正式规格、合同主线与验收结论继承
- `PSCO-mvp02-summarize-feedback.md` 只提供方向仲裁与候选顺序，不替代本阶段 `/plan`

## 3. 本阶段目标

`phase06` 的目标是：

> 在 `phase05` 已交付最小 Dashboard 闭环的前提下，交付 `Onboarding + Data Sovereignty + Reuse Awareness` 的首个可执行主线，使新用户能够低摩擦完成首轮资产录入、确认资产可带走，并在系统中看到最小复用 / 能力反馈。

本阶段需要回答的核心问题：

1. 首轮进入系统时，最小录入路径如何从“完整表单”收敛到“draft-first / partial-entry”
2. 导出与基础备份在当前阶段的正式语义、范围与承接位是什么
3. `module_reuse_summary` 与 `capability_summary` 应在哪些页面成为正式可见反馈
4. 前端新增写路径如何正式遵守 `application` 单入口、`query` 纯只读与固定 mutation 承接位
5. `.proto` 唯一合同源与 `chi + HTTP JSON` 传输适配如何在新阶段新增接口中被严格执行

## 4. 架构冻结结论

### 4.1 当前阶段唯一直接执行层上游

`phase06` 必须直接承接：

- `PSCO-mvp02-summarize-feedback.md`
- `.trae/specs/phase05_10_dashboard_feedback_formal_spec/dashboard_feedback_spec_v0.1.md`
- `.trae/specs/phase05_11_dashboard_feedback_proto_mainline/`
- `.trae/specs/phase05_14_dashboard_feedback_integration_validation_acceptance/acceptance_report.md`

不允许在本阶段重新解释：

- `phase05` 已冻结的 Dashboard 路由、反馈信号与返回路径主线
- `phase04` 已冻结的 `Product / Repository / Binding` canonical owner
- `Decision` 作为核心主线的地位
- 独立 `AI Assistant`、独立 `React Native` 客户端、完整 `PWA`

### 4.2 当前阶段主交付对象

`phase06` 的主交付对象是：

- `Onboarding Foundation`
- `Data Sovereignty Baseline`
- `Reuse Awareness Surface`

其最小主线必须优先承接：

- `Product / Repository / Module / Decision` 的低摩擦初始录入
- first-run 引导与草稿优先录入路径
- 核心资产导出
- 面向当前实例的基础备份
- `module_reuse_summary`
- `capability_summary`
- Dashboard / Detail 中的最小复用反馈入口

当前阶段不把以下对象提升为新一级重实体：

- `Capability`
- `Template`
- `Opportunity / Feature / Experiment`

### 4.3 当前阶段前端交付策略

前端继续统一遵守：

- 单一 `React Web`
- 同时考虑 `PC` 与移动浏览器 UI
- 业务写路径唯一 `application` 入口
- `query` 层纯只读
- mutation 固定承接位
- 切片优先与 `shared` 延迟晋升

当前阶段重点是：

- 把首轮录入从“高完整度表单”收敛为“最小可提交 + 后续补全”
- 避免在 Onboarding 页面、详情页、弹窗面板中长出多套并行写路径
- 让复用反馈进入既有 Dashboard / Detail，而不是再造第二套“智能中心”

### 4.4 当前阶段数据、合同与传输承接原则

当前阶段关于数据主权的冻结如下：

- 导出语义面向用户带走核心资产数据
- 基础备份语义面向当前实例保留与恢复
- 二者都不得依赖第三方平台作为唯一前提

当前阶段关于导出 / 备份的最小覆盖矩阵补充冻结如下：

- `Export` 最小必须覆盖：
  - `products`
  - `modules`
  - `releases`
  - `repositories`
  - `decisions`
  - `decision_links`
  - `product_modules`
  - `product_repositories`
  - `module_repositories`
- `Backup` 的最小覆盖范围不得小于 `Export`
- `Backup` 除上述核心资产外，至少还必须带出：
  - 当前备份清单或 manifest
  - 备份创建时间
  - 当前实例恢复所需的 schema / 版本前提
- 当前阶段不允许把“只导出主实体，不导出绑定 / 关联关系”解释为完成数据主权闭合

当前阶段关于合同与传输的冻结如下：

- `.proto` 是唯一长期合同源
- `buf build / lint / generate / breaking` 继续作为合同工具链基线
- `chi + HTTP JSON` 在当前阶段继续保留为传输适配层
- 新增 HTTP DTO 必须从 `.proto` 单向派生或显式映射，不得长出第二套 canonical contract

当前阶段关于复用感知读模型的计算与新鲜度补充冻结如下：

- `module_reuse_summary` 由当前 canonical 绑定关系读时聚合得到；当前阶段不引入独立统计表作为唯一事实源
- `module_reuse_summary` 的最小统计口径冻结为“一个 Module 当前被多少 Product 直接复用”
- `module_reuse_summary` 至少带出：
  - `module_id`
  - `reuse_product_count`
  - `latest_reuse_at`
  - 最小解释文案
- `capability_summary` 由当前 Module 派生信息读时聚合得到；当前阶段不引入独立 Capability 重实体
- `capability_summary` 至少带出：
  - `capability_key`
  - `capability_label`
  - `supporting_module_count`
  - `latest_capability_update_at`
  - 最小解释文案
- `capability_summary` 的最小事实来源冻结为 Module 写模型中的轻量 `capability_key`（可空）与系统内置 label 映射；未填写 `capability_key` 的 Module 不参与当前阶段 capability 聚合，但不阻断首轮成功会话
- 当前阶段复用感知的新鲜度口径冻结为“读取时反映最新已提交状态”，不引入异步离线刷新前提

当前阶段不在架构规划中提前冻结：

- 导出文件格式的最终单一实现
- 备份脚本名、systemd 落点与目录结构
- 模板的存储方式
- `module_reuse_rate` 的最终计算公式

### 4.5 当前阶段交互归属原则

为了避免 `phase06` 在后续 `/spec` 与实现阶段出现“Onboarding 既引导又变成第二套业务工作台”的歧义，当前阶段先冻结以下交互归属原则：

- Onboarding 页面只承接首轮进入、最小录入与后续跳转
- 正式详情写入仍回到既有 canonical owner 页面或当前阶段新建的单值写路径
- Dashboard 继续承接经营入口与反馈展示，不承接导出 / 备份的全部管理职责
- 导出 / 备份必须有明确入口，但不扩写为完整运维中心
- 复用反馈优先挂接在 Dashboard、Module Detail、Product Detail 等既有页面，而不是新建一级导航

当前阶段关于正式入口位补充冻结如下：

- `Onboarding` 的正式业务入口路由冻结为 `/onboarding`
- 冷启动空系统下，主 CTA 必须落到 `/onboarding`
- 首轮未完成用户在回访时，必须看到 `Continue Onboarding` 入口，并回到 `/onboarding`
- `Export` 的正式用户入口冻结为 `Dashboard` 动作区中的独立入口，正式执行路由冻结为 `/dashboard/export`
- `Backup` 的正式用户入口冻结为 `Dashboard` 动作区中的独立入口，正式执行路由冻结为 `/dashboard/backup`
- `Export / Backup` 不允许以内联方式直接塞回 `Dashboard Home` 主内容区完成全部操作
- 首次进入应用时的 cold-start 判定固定由前端根级路由入口守卫承接（`beforeLoad` 或等价根级 loader），不得分散到页面组件 `useEffect` 中各自判断
- `first_run_state = not_started` 时，根级默认进入路径必须回落到 `/onboarding`
- `first_run_state = in_progress` 时，不要求劫持所有 canonical detail 路由；根级默认进入路径与 `Dashboard` 必须提供 `Continue Onboarding` 入口

当前阶段关于首轮成功会话的单值定义补充冻结如下：

- 当前阶段首轮成功会话的推荐执行顺序冻结为：
  - `Product -> Repository -> Module -> Decision`
- 当前阶段“完成一次首轮成功会话”指：
  - 在一次连续会话中，至少各创建 `1` 条已持久化记录：
    - `Product`
    - `Repository`
    - `Module`
    - `Decision`
  - 允许这些记录在当前阶段先以 `draft-first / partial-entry` 形式存在
  - 会话结束时用户能够回到 `Dashboard` 或任一 canonical owner 页面继续补全
- 当前阶段首轮成功会话不强制要求：
  - `Product` 已完成全部绑定
  - `Repository` 已绑定到 `Product`
  - `Module` 已映射到 `Repository`
  - `Decision` 已完成对象链接
- 但 `Decision` 至少必须完成最小可持久化记录，不允许把“尚未打开 Decision 写路径”也算作 onboarding 已完成
- `first_run_state` 的最小状态跃迁冻结为：
  - 尚未开始任何首轮对象写入：`not_started`
  - 已至少创建 `1` 条首轮对象记录、但四类对象未全部持久化：`in_progress`
  - 四类对象均已持久化并满足首轮成功会话条件：`completed`

### 4.6 当前阶段源码设计层输出要求

`phase06` 虽然当前处于 `/plan`，但为了保证后续 `/spec` 可直接进入实现，本阶段必须把以下源码设计层结果纳入任务规划：

- first-run onboarding 页面、步骤与最小字段模型
- `Product / Repository / Module / Decision` 的 draft-first / partial-entry 写路径设计
- 导出、备份的后端模块边界、接口分组与恢复前提
- `module_reuse_summary / capability_summary` 的读模型、owner 与页面挂接位
- 前端新增 mutation 的固定承接位与失效刷新归一化策略
- `.proto` 合同、HTTP DTO 与前端消费模型的一致性校验策略
- 验收环境中的冷启动、部分录入、导出、备份与复用反馈 fixture
- `backup verified` 基于备份产物与 manifest 读取校验成立，而不是仅以“写出文件成功”代替

### 4.7 当前阶段规划吸取的历史经验

本阶段必须明确吸取前几阶段经验，避免重复补票：

- 低摩擦入口要在设计阶段先冻结，不能实现时再“顺手简化”
- `.proto` 合同与传输映射必须从阶段早期就进入任务主线，不再后补
- 写路径 owner 必须在前端阶段设计时单值化，不再允许页面级散装 `useMutation`
- 数据主权能力必须视为底座负债，不再等到阶段尾部作为“有空再做”
