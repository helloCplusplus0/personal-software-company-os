# Tasks

- [x] Task 1: 冻结项目级治理层的正式身份与职责边界
  - [x] SubTask 1.1: 复核 `architecture_plan / dev_plan / shared_baseline` 是否都将项目级治理层定义为“项目级承接层 / 治理层 / 合同层”
  - [x] SubTask 1.2: 显式冻结“它不是第五个业务主实体、不是新的业务 CRUD 主线”
  - [x] SubTask 1.3: 复核后续执行者不再需要猜“项目级治理层到底算不算新实体”

- [x] Task 2: 冻结项目级治理层与四实体的关系矩阵
  - [x] SubTask 2.1: 明确它与 `Repository` 的锚点关系与非字段扩张边界
  - [x] SubTask 2.2: 明确它与 `Product / Module` 的治理背景关系，而非业务事实替代关系
  - [x] SubTask 2.3: 明确 `Decision` 不再兼容承接目录结构、全局规范与当前阶段信息
  - [x] SubTask 2.4: 复核四实体各自继续承接什么、项目级治理层正式承接什么，已形成单值矩阵

- [x] Task 3: 冻结第一版项目锚点、前端承接位与双侧读取口径
  - [x] SubTask 3.1: 明确第一版唯一项目锚点为 `repository_id`
  - [x] SubTask 3.2: 明确第一版前端正式承接位为 `Repository detail`
  - [x] SubTask 3.3: 复核 Web 与 agent 都读取同一 `repository_id` 驱动的项目级治理事实，而不是各自解释一套

- [x] Task 4: 冻结项目级治理层与四实体事实的禁止越界规则
  - [x] SubTask 4.1: 明确项目治理画像、全局规范资产与当前阶段入口状态统一由项目级治理层承接
  - [x] SubTask 4.2: 明确四实体继续承接业务事实、业务关系与业务动作语义
  - [x] SubTask 4.3: 复核同一概念不会在项目级治理层与四实体业务事实中形成两套正式语义

- [x] Task 5: 完成 spec 包与 `phase13` 上游文档的一致性校验
  - [x] SubTask 5.1: 校验本 spec 包与 `phase13` 三件套在项目级治理层定位上保持单值一致
  - [x] SubTask 5.2: 校验本 spec 包没有引入与 `AGENTS.md / plan.md / project_rules.md` 相冲突的新解释
  - [x] SubTask 5.3: 校验后续 `phase13-04 ~ phase13-10` 可直接承接本 spec，而不再需要重新判断“治理层与四实体到底是什么关系”

# Task Dependencies

- `Task 2` depends on `Task 1`
- `Task 3` depends on `Task 1`, `Task 2`
- `Task 4` depends on `Task 1`, `Task 2`
- `Task 5` depends on `Task 1`, `Task 2`, `Task 3`, and `Task 4`

# 执行记录

- SubTask 1.1 结论：`architecture_plan` §4.4 明确定义为“四实体之外的**项目级承接层 / 治理层 / 合同层**”；`dev_plan` §2 与 `shared_baseline` §3.7 使用同一“项目级治理层/项目级承接层”口径；无漂移
- SubTask 1.2 结论：spec Requirement 1 “正式职责不包括：充当新的资产登记主线 / 充当第五个业务 CRUD 主实体”与 `architecture_plan` §4.2 补充冻结“不允许因此长出第五个业务主实体”、`shared_baseline` §2.4“当前阶段新增承接层不得演化为第五个业务主实体”一致
- SubTask 1.3 结论：spec Requirement 1 Scenario 将“把项目级治理层设计为与四实体并列的新业务对象”显式判定为越界并要求回到正式定位；执行者无需临场判断
- SubTask 2.1 结论：spec Requirement 2 第 1 项与 `architecture_plan` §4.4 一致：`Repository` 是代码仓库身份对象与项目锚点、第一版以同一 `repository_id` 为唯一正式锚点、前端承接位为 `Repository detail`、不得偷换为 `Repository` 业务字段扩张；spec Requirement 3 补充冻结“锚点只表示定位与读取对齐方式，不表示治理层变成 Repository 的业务子类型”为该结论的细化表述
- SubTask 2.2 结论：spec Requirement 2 第 2/3 项中 `Product`“经营目标与交付容器”、`Module`“可复用能力资产”与 phase12-02 冻结的前端 `shared-semantic-constants.ts` 四实体语义常量逐字一致（语义标签的正式冻结源为 phase12-02，三件套仅部分承载该表述，如 `Repository`“代码仓库身份对象与项目锚点”见 `architecture_plan` §4.4，三件套中所有实际出现的表述与 spec 一致无冲突）；治理层只提供治理背景、不承接业务事实
- SubTask 2.3 结论：spec Requirement 2 第 4 项与 `dev_plan` phase13-03 DoD“不再把目录结构、全局规范与当前阶段信息硬塞进 `Decision`”逐句对应；`Decision`“规则、约束、选择与依据的索引对象”语义保持不变
- SubTask 2.4 结论：关系矩阵已单值化——四实体承接业务事实/关系/动作语义，治理层承接治理画像/全局规范资产/当前阶段入口状态，职责互不重叠
- SubTask 3.1 结论：`repository_id` 为第一版唯一项目锚点，与 `architecture_plan` §4.7 brief 解析协议第 1 条、`shared_baseline` §3.6 一致
- SubTask 3.2 结论：`Repository detail` 为第一版前端正式承接位，与 `architecture_plan` §4.4、`shared_baseline` §3.7、`dev_plan` phase13-06 一致；“不新增独立页面或并列第二入口”与 spec Requirement 3 补充冻结一致
- SubTask 3.3 结论：spec Requirement 3 第 3 条“Web 与 agent 读取必须都回到同一 repository_id 驱动的正式事实”与 `shared_baseline` §4 验收问题 6、§2.4“不允许让 Web 与 agent 各自长出第二套项目治理解释”一致
- SubTask 4.1 结论：spec Requirement 4 第 2 条与主交付单值定义（治理画像 + 全局规范资产 + 当前阶段入口状态）一致；REMOVED Requirement 的 Migration 将这三类信息统一回收到项目级治理层
- SubTask 4.2 结论：spec Requirement 4 第 1 条保持四实体业务主线不变，与 `shared_baseline` §2.4“四实体业务主线继续保持”一致
- SubTask 4.3 结论：spec Requirement 4 第 3/4 条“同一概念不得两套正式语义、治理背景不得回写为四实体主语义”与 `architecture_plan` §4.9 第 7 项“第二套与四实体并列的事实源”非目标一致
- SubTask 5.1 结论：spec 与三件套在治理层定位、关系矩阵、锚点、承接位与越界规则上单值一致；spec 对三件套结论做了矩阵化细化（逐实体关系 + Scenario 判定规则），未引入冲突
- SubTask 5.2 结论：spec 未引入与 `AGENTS.md / plan.md / project_rules.md` 冲突的新解释；`AGENTS.md` §3“当前项目目录结构与根级 canonical 文件集合，已进入 phase13 的直接规划范围”与治理层承接范围一致；`plan.md` phase13 条目范围约束一致
- SubTask 5.3 结论：治理层身份、四实体关系矩阵、唯一锚点、承接位与越界规则均已单值冻结并配套机械判定 Scenario；`phase13-04 ~ phase13-10` 可直接承接，与 `dev_plan` §5 依赖关系第 2 条（phase13-02/03 是 04~07 的边界前提）衔接一致
