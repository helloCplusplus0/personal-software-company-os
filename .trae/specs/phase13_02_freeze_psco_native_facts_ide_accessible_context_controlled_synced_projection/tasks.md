# Tasks

- [x] Task 1: 冻结三层信息模型的正式定义与命名边界
  - [x] SubTask 1.1: 复核 `architecture_plan / dev_plan / shared_baseline` 是否都使用同一套 `PSCO-native facts / IDE-accessible context / Controlled synced projection` 命名
  - [x] SubTask 1.2: 复核三层模型没有被自由重命名、重拆层或扩写为并列第四层
  - [x] SubTask 1.3: 若三件套间仍存在命名或层级漂移，回写为单值结论

- [x] Task 2: 冻结 `PSCO-native facts` 的正式承接范围
  - [x] SubTask 2.1: 明确四实体信息与关系属于 `PSCO-native facts`
  - [x] SubTask 2.2: 明确全局规范资产属于 `PSCO-native facts`
  - [x] SubTask 2.3: 明确项目级治理画像属于 `PSCO-native facts`
  - [x] SubTask 2.4: 复核后续执行者不再需要猜“哪些信息该先进入 phase13 主线”

- [x] Task 3: 冻结 `IDE-accessible context` 的保留边界
  - [x] SubTask 3.1: 明确工作区文件全文、目录漂移与局部实现细节继续留给 IDE / agent 现场读取
  - [x] SubTask 3.2: 复核 spec 中已明确“这些信息不默认上升为 PSCO 正式事实”
  - [x] SubTask 3.3: 复核后续执行者无法再把目录即时上下文误读成 `phase13` 当前正式承接对象

- [x] Task 4: 冻结 `Controlled synced projection` 的后置进入条件
  - [x] SubTask 4.1: 明确 Git 推进摘要、模板仓库接入状态与模板版本自动同步、自动存在性校验与自动状态建议属于后续受控进入项
  - [x] SubTask 4.2: 复核这些能力未被写成 `phase13` 起点默认实现内容
  - [x] SubTask 4.3: 复核执行顺序已冻结为“先正式事实，后现场读取，最后才讨论同步增强”

- [x] Task 5: 完成 spec 包与 phase13 三件套的一致性校验
  - [x] SubTask 5.1: 校验本 spec 包与 `phase13` 三件套在三层边界定义上保持单值一致
  - [x] SubTask 5.2: 校验本 spec 包没有引入与 `AGENTS.md / plan.md / project_rules.md` 相冲突的新解释
  - [x] SubTask 5.3: 校验后续 `phase13-03 ~ phase13-12` 能直接承接本 spec，而不再需要重新判断“什么该先做，什么后做”
  - [x] SubTask 5.4: 校验同一概念不会在 A / B / C 多层重复承接

# Task Dependencies

- `Task 2` depends on `Task 1`
- `Task 3` depends on `Task 1`
- `Task 4` depends on `Task 1`
- `Task 5` depends on `Task 1`, `Task 2`, `Task 3`, and `Task 4`

# 执行记录

- SubTask 1.1 结论：`architecture_plan` §4.3 与 `shared_baseline` §3.2 使用完全一致的 `PSCO-native facts / IDE-accessible context / Controlled synced projection` 三层命名；`dev_plan` phase13-02 标题为小写 `controlled synced projection`，仅大小写差异，语义等价，无漂移
- SubTask 1.2 结论：三件套均保持三层结构，无重命名、无重拆层、无并列第四层；spec Requirement 1 已显式禁止后续自由重命名、重新拆层或新增第四层
- SubTask 1.3 结论：未发现命名或层级漂移，无需回写修改
- SubTask 2.1 结论：spec Requirement 2 第 1 项与 `architecture_plan` §4.3 第 1 层、`shared_baseline` §3.2 A 层第 1 项一致，四实体（Product / Repository / Module / Decision）信息与关系明确归属 `PSCO-native facts`
- SubTask 2.2 结论：全局规范资产（workflow、技术基线、协作约束、目录与文档职责）在三处均归属 A 层，与 spec Requirement 2 第 2 项一致
- SubTask 2.3 结论：项目级治理画像在三处均归属 A 层；本轮已回写 `docs_workflow_layout` 进入第一版正式字段矩阵，`spec / architecture_plan / shared_baseline / dev_plan` 现已一致承认“项目范式版本 + docs workflow 布局 + canonical 根级文件集合 + 当前阶段入口与状态”属于治理画像范围
- SubTask 2.4 结论：spec Requirement 2 Scenario 已给出机械判定规则（属于四实体关系/全局规范资产/治理画像 → A 层 → 优先进入主线），执行者无需猜测
- SubTask 3.1 结论：spec Requirement 3 的 4 项（工作区文件全文/临时目录漂移/局部实现细节/agent 现场即时读取）与 `architecture_plan` §4.3 第 2 层逐项对应，继续留给 IDE / agent 现场读取
- SubTask 3.2 结论：spec Requirement 3 明确“不默认上升为 PSCO 正式事实”；`shared_baseline` §3.2 补充冻结同口径（B 层默认继续由 IDE / agent 现场读取）
- SubTask 3.3 结论：spec Requirement 3 Scenario 将“把工作区文件全文、目录即时状态或局部实现细节纳入正式事实层”显式判定为越界，执行者无法误读
- SubTask 4.1 结论：spec Requirement 4 的 4 项（Git 推进摘要/模板仓库接入状态与模板版本自动同步/自动存在性校验/自动状态建议）与 `architecture_plan` §4.3、`shared_baseline` §3.2 C 层现已逐项一致，均为后续受控进入项
- SubTask 4.2 结论：三件套均未把 C 层能力写为本阶段默认实现内容；`architecture_plan` §4.9 与 `dev_plan` §4 的非目标清单进一步覆盖该约束
- SubTask 4.3 结论：spec Requirement 5 冻结执行顺序为“先正式承接 A 层 → 保持 B 层现场读取 → C 层下沉为后续进入条件”，与 `shared_baseline` §3.2 补充冻结的顺序口径一致；与 `dev_plan` §5 依赖关系（phase13-02/03 是 04~07 的边界前提）衔接一致
- SubTask 5.1 结论：spec 与三件套在三层边界定义、A/B/C 各层内容与执行顺序上单值一致；当前仅保留非阻断性粒度差异：A 层里 `shared_baseline` 继续将“当前阶段入口与状态”单列，而 spec/`architecture_plan` 仍将其写在治理画像描述内，但都明确位于 A 层；B 层措辞“临时/局部”与“目录局部/即时”逐项等价
- SubTask 5.2 结论：spec 未引入与 `AGENTS.md / plan.md / project_rules.md` 冲突的新解释；`AGENTS.md` §3“PSCO 优先管理 PSCO-native facts…不把 IDE 目录即时上下文默认上升为正式事实源”与 `plan.md` phase13 条目“不提前混入 Git 推进跟踪、模板仓库接入、自动同步或 agent 写回”均被 spec 正确承接
- SubTask 5.3 结论：三层模型、各层内容、执行顺序与统一口径（Requirement 6 禁止跨层偷换与双重解释）均已单值冻结，`phase13-03 ~ phase13-12` 可直接承接本 spec 做边界判断
- SubTask 5.4 结论：已新增“同一概念不得跨 A / B / C 多层重复承接”的显式校验；`template_source` 现仅作为 A 层治理画像第一版手工字段，模板仓库接入状态与模板版本自动同步仅保留在 C 层，不再跨层冲突
