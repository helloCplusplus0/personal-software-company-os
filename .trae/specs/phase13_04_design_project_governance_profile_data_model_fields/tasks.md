# Tasks

- [x] Task 1: 冻结项目治理画像第一版核心字段集合
  - [x] SubTask 1.1: 明确 `project_profile_version / track_type / template_source / docs_workflow_layout` 的正式字段身份
  - [x] SubTask 1.2: 明确 `canonical_root_files / global_constraint_refs / current_phase_*` 属于同一字段模型
  - [x] SubTask 1.3: 复核第一版不再需要临场猜“还缺哪些同级核心字段”

- [x] Task 2: 冻结核心字段语义与嵌套字段矩阵
  - [x] SubTask 2.1: 明确每个核心字段承接的治理语义
  - [x] SubTask 2.2: 明确核心标量字段的结构形状、可空规则与受控取值
  - [x] SubTask 2.3: 明确 `canonical_root_files[]` 的最小子字段集合
  - [x] SubTask 2.4: 明确 `global_constraint_refs[]` 的最小子字段集合
  - [x] SubTask 2.5: 复核第一版不再临时补入未冻结语义的并列子字段

- [x] Task 3: 冻结字段分类矩阵
  - [x] SubTask 3.1: 明确哪些字段属于 `required`
  - [x] SubTask 3.2: 明确哪些字段属于 `optional`
  - [x] SubTask 3.3: 明确哪些字段属于具体 `read-only`
  - [x] SubTask 3.4: 明确哪些内容只属于 `future-auto-verifiable`
  - [x] SubTask 3.5: 复核第一版仍以手工维护合同为先，而不是自动推断为先

- [x] Task 4: 冻结当前项目范式 v1 的字段映射方式
  - [x] SubTask 4.1: 明确目录结构如何映射到治理画像字段模型
  - [x] SubTask 4.2: 明确 canonical 根级文件集合如何映射到 `canonical_root_files[]`
  - [x] SubTask 4.3: 复核 `docs_workflow_layout` 能承接当前 docs workflow 布局

- [x] Task 5: 完成 spec 包与 `phase13` 上游文档的一致性校验
  - [x] SubTask 5.1: 校验本 spec 包与 `phase13` 三件套在最小字段矩阵上保持单值一致
  - [x] SubTask 5.2: 校验本 spec 包没有引入与 `AGENTS.md / plan.md / project_rules.md` 相冲突的新解释
  - [x] SubTask 5.3: 校验后续 `phase13-05 ~ phase13-10` 可直接承接本 spec，而不再需要重新判断“第一版到底先做哪些字段”

# Task Dependencies

- `Task 2` depends on `Task 1`
- `Task 3` depends on `Task 1`, `Task 2`
- `Task 4` depends on `Task 1`, `Task 2`
- `Task 5` depends on `Task 1`, `Task 2`, `Task 3`, and `Task 4`

# 执行记录

- SubTask 1.1 结论：4 个标量字段身份已冻结——`project_profile_version`（画像正式版本）、`track_type`（正式技术路线，须与当前项目冻结的 `Durable System Track` 单值一致）、`template_source`（手工维护的范式/模板来源信息）、`docs_workflow_layout`（docs workflow 结构布局）；4 项与 `shared_baseline` §3.3 最小字段矩阵逐项一致（§3.3 同样列有 `docs_workflow_layout`，且其补充冻结已明确 `template_source` 属 A 层手工字段、不得解释为 C 层自动同步结果）
- SubTask 1.2 结论：`canonical_root_files[] / global_constraint_refs[] / current_phase_name / current_phase_ref / current_phase_status` 与上述 4 个标量字段同属一个字段模型，`shared_baseline` §3.3 与 spec 均为 9 项单值集合、逐项一致；spec Requirement 1 禁止临时新增“视情况再补”的并列第二批字段
- SubTask 1.3 结论：spec Requirement 1 Scenario 已将“再补一个同级核心字段试试”显式判定为越界；9 类字段覆盖治理画像版本、技术路线、模板来源、docs 布局、根级文件、规范引用与阶段状态，无缺口
- SubTask 2.1 结论：spec Requirement 2 逐字段冻结语义——尤其 `template_source` 为手工字段不得偷换为自动同步结果（对应 phase13-02 C 层后置）、`current_phase_*` 必须回到正式 phase 主线而不是散装阶段快照（对应 `AGENTS.md` 阶段状态以 `plan.md` 为准）
- SubTask 2.2 结论：核心标量字段现已冻结到“字段形状 + 可空规则 + 受控取值”层级：`project_profile_version` 为非空 `string` 且首个正式取值固定为 `project_governance_profile_v1`；`track_type` 为受控 `enum`，允许值只包括 `Product Track / Durable System Track`，当前项目固定 `Durable System Track`；`template_source` 为 `string | null`，允许为空；`docs_workflow_layout` 为非空受控 `string`，当前项目固定为 `phase/fix/audit/review`；`current_phase_name` 为非空 `string`，`current_phase_ref` 为指向正式 phase 入口的非空 `string`，`current_phase_status` 为受控 `enum`，允许值固定为 `planned / in_progress / completed / blocked`
- SubTask 2.3 结论：`canonical_root_files[]` 最小子字段冻结为 `file_name / role / required`，与 `shared_baseline` §3.3 逐字一致
- SubTask 2.4 结论：`global_constraint_refs[]` 最小子字段冻结为 `name / kind / entry_ref`，与 `shared_baseline` §3.3 逐字一致；口径关系说明：该数组承接的是全局规范资产的**入口引用与分类关系**（§3.3 口径），`shared_baseline` §3.5 的 `entry_ref + role + structured summary` 逐项承接策略由后续 `phase13-05` 后端合同设计承接，两者为不同矩阵、不冲突
- SubTask 2.5 结论：spec Requirement 3 补充冻结明确禁止第一版临时补入“全文内容 / 自动摘要结果 / 最近修改时间”等并列子字段；此类扩展只能作为后续 phase 显式进入项
- SubTask 3.1 结论：`required` 冻结为 8 项——`project_profile_version / track_type / docs_workflow_layout / canonical_root_files[] / global_constraint_refs[] / current_phase_name / current_phase_ref / current_phase_status`
- SubTask 3.2 结论：`optional` 冻结为 `template_source` 单项——手工维护的来源信息允许为空，不阻断画像成立
- SubTask 3.3 结论：`read-only` 现已落实到具体字段集合——`track_type / current_phase_name / current_phase_ref / current_phase_status`；这些字段都由根级正式上游冻结，只允许展示 / 回读，不允许在治理画像维护入口中自由改写；同时 spec 已显式冻结“字段分类是可叠加维度”，因此这些字段允许同时属于 `required` 与 `read-only`
- SubTask 3.4 结论：`future-auto-verifiable` 冻结为 3 类存在性校验（`canonical_root_files[].required` / `global_constraint_refs[].entry_ref` / `current_phase_ref`），均只属后续受控进入项，与 `shared_baseline` §3.3“存在性校验、自动建议与同步增强属于后续受控进入项”一致
- SubTask 3.5 结论：spec Requirement 4 补充冻结“第一版优先解决字段合同是否成立，不先解决字段能否被自动推断”，与 `shared_baseline` §3.3“第一版优先手工维护，不默认自动推断；优先管理合同，不是先管理扫描结果”单值一致；`read-only` 规则已从原则补齐为具体字段集合，不再停留在抽象约束
- SubTask 4.1 结论：目录结构承接方式分层说明——`docs/` 的 `phase / fix / audit / review` 四子目录布局由 `docs_workflow_layout` 明确承接（spec Requirement 2 第 4 项）；`backend/ / database/ / frontend/ / proto/` 四个顶层目录在 9 项字段中无专属承接字段，按 `architecture_plan` §4.5 定位为“治理画像的直接输入”而非必入字段，其是否需要显式字段承接留待 `phase13-05` 合同设计澄清（不得在第一版临时补入未冻结语义的字段）；与 `shared_baseline` §3.4 目录结构清单一致
- SubTask 4.2 结论：canonical 根级文件集合（`.env / AGENTS.md / architecture_map.md / plan.md / global_skills.md / project_skills.md / project_rules.md / README.md / TECH_STACK_BASELINE.md` 共 9 项）映射到 `canonical_root_files[]`，与 `shared_baseline` §3.4、`architecture_plan` §4.5 逐项一致；spec 补充冻结“结构化承接的治理事实”而非强制模板规则，与 §3.4“不要求立刻扩写成模板仓库”一致
- SubTask 4.3 结论：`docs_workflow_layout` 语义明确要求能承接 `docs/phase / fix / audit / review` 正式布局语义，覆盖 `shared_baseline` §3.4 的 docs 子目录结构与“`docs/README.md` 继续作为 docs workflow 总入口”
- SubTask 5.1 结论：spec 与三件套在字段模型上单值一致——`shared_baseline` §3.3 与 spec 的 9 项核心字段逐项完全一致（此前执行记录误称 §3.3 为 8 项，经独立复核与文件核实已修正）；spec 作为字段模型冻结点成为 05~07 的直接上游；已知非阻断性遗留项：四个顶层目录（backend/database/frontend/proto）的字段承接方式留待 phase13-05 澄清（见 SubTask 4.1）
- SubTask 5.2 结论：spec 未引入与 `AGENTS.md / plan.md / project_rules.md` 冲突的新解释；`track_type` 与 `TECH_STACK_BASELINE.md` 冻结的 `Durable System Track` 一致；字段集合未偷渡目录扫描、自动建议与同步投影字段（phase13-01 非目标 / phase13-02 C 层后置口径）
- SubTask 5.3 结论：9 类核心字段、嵌套子字段矩阵、`required / optional / read-only / future-auto-verifiable` 分类与项目范式 v1 映射均已单值冻结并配套机械判定 Scenario；`phase13-05 ~ phase13-10` 可直接按本 spec 字段模型进入合同、入口与 brief 设计，与 `dev_plan` §5 依赖关系（phase13-04 是 05/06/07 的共享字段前提）衔接一致
