# Tasks

- [x] Task 1: 冻结第一版前端正式承接位与入口层级
  - [x] SubTask 1.1: 明确 `Repository detail` 是第一版唯一正式承接位
  - [x] SubTask 1.2: 明确 `repository_id` 是唯一正式前端读取锚点
  - [x] SubTask 1.3: 复核第一版不会新增独立一级页面或第二入口

- [x] Task 2: 冻结 Repository detail 的信息架构分层
  - [x] SubTask 2.1: 明确治理画像概览层的展示范围
  - [x] SubTask 2.2: 明确结构化维护层的承接对象
  - [x] SubTask 2.3: 明确摘要回看层的承接对象
  - [x] SubTask 2.4: 复核第一版不会重演 `phase12` 的大块常驻解释 UI，且旧“项目上下文”设计会系统性退出前端

- [x] Task 3: 冻结前端手工维护与只读展示边界
  - [x] SubTask 3.1: 明确第一版允许手工维护的字段集合
  - [x] SubTask 3.2: 明确第一版只读展示的字段集合
  - [x] SubTask 3.3: 复核前端可编辑范围与 `phase13-05` 写边界保持单值一致

- [x] Task 4: 冻结人类维护、agent 消费与摘要回看的边界
  - [x] SubTask 4.1: 明确哪些内容供人类维护
  - [x] SubTask 4.2: 明确哪些内容只供 agent 消费
  - [x] SubTask 4.3: 明确哪些内容只保留摘要回看
  - [x] SubTask 4.4: 复核前端不会承担 agent brief 主消费职责

- [x] Task 5: 冻结 phase12 遗留“项目上下文”设计的退出规则
  - [x] SubTask 5.1: 明确 `Repository detail` 现有“项目上下文”区必须移除
  - [x] SubTask 5.2: 明确 `Decision detail` 现有“共享项目上下文入口”卡片必须移除
  - [x] SubTask 5.3: 明确不得以并区、换名或轻量保留方式延续 `phase12 project-context` 叙事

- [x] Task 6: 冻结目录真实路径与 `entry_ref` 的展示约束
  - [x] SubTask 6.1: 明确目录真实路径不是面向普通用户的主内容
  - [x] SubTask 6.2: 明确 `entry_ref` 的正式角色是定位元数据而不是正文主体
  - [x] SubTask 6.3: 复核结构化摘要优先于真实路径成为主阅读内容

- [x] Task 7: 完成 spec 包与上游冻结文档的一致性校验
  - [x] SubTask 7.1: 校验本 spec 包与 `phase13-03` 的前端承接位口径保持单值一致
  - [x] SubTask 7.2: 校验本 spec 包与 `phase13-05` 的后端读写边界保持单值一致
  - [x] SubTask 7.3: 校验本 spec 包足以支撑后续页面层级、表单结构与 UI 实现而不再需要临场判断

# Task Dependencies

- `Task 2` depends on `Task 1`
- `Task 3` depends on `Task 1`, `Task 2`
- `Task 4` depends on `Task 1`, `Task 2`, `Task 3`
- `Task 5` depends on `Task 1`, `Task 2`, `Task 4`
- `Task 6` depends on `Task 2`, `Task 4`, `Task 5`
- `Task 7` depends on `Task 1`, `Task 2`, `Task 3`, `Task 4`, `Task 5`, and `Task 6`

# 执行记录

- SubTask 1.1 结论：`Repository detail` 冻结为第一版唯一正式前端承接位，与 `dev_plan` phase13-06 范围（“第一版前端正式承接位冻结为：`Repository detail`”）、`shared_baseline` §3.7、phase13-03 Requirement 3 三处逐句一致
- SubTask 1.2 结论：`repository_id` 为唯一正式前端读取锚点，与 phase13-03 Requirement 3 第 1 条、phase13-05 承接位锚点口径单值一致；Web 读取与 agent 消费共享同一锚点驱动的治理事实
- SubTask 1.3 结论：spec Requirement 1 补充冻结四条——禁独立一级页面 / 禁并列第二入口 / 禁在 `Dashboard / Daily Review / Weekly Review / Product detail / Module detail / Decision detail` 长出并列主承接区 / 独立入口仅 `phase13` 收口后作为下一阶段进入条件讨论——前三条与 REMOVED Requirement（其他页面仅保留业务主语义）闭合，第四条与 `dev_plan` phase13-06 范围、phase13-03 Requirement 3 补充冻结逐句一致
- SubTask 2.1 结论：概览层承接治理画像版本、技术路线、docs workflow 布局等高层概览（对应 `project_profile_version / track_type / docs_workflow_layout` 只读展示），并冻结“保持轻量、不得成为大块解释型常驻 UI”
- SubTask 2.2 结论：结构化维护层承接第一版允许人类手工维护的结构化字段与全局规范资产承接结果（对应 Requirement 3 可编辑集合），以表单 / 面板 / 列表等结构化方式呈现
- SubTask 2.3 结论：摘要回看层承接全局规范资产的 `role / structured_summary / entry_ref` 与当前阶段入口状态回看，突出“结构化摘要”而非正文全文——与 phase13-05 读边界第 2/3 条（返回 entry_ref/role/structured_summary 承接结果、暴露正文可回源能力）衔接一致
- SubTask 2.4 结论：spec Requirement 2 补充冻结“第一版不得重演 `phase12` 中把解释性或验收性信息做成大块常驻 UI 的问题”，与 `dev_plan` phase13-06 DoD 第 1 条语义对应（spec 表述“解释性或验收性信息”略宽于 dev_plan 的“验收层信息”，方向一致且更完整）；三层同属一个正式承接区、不得拆成多个并列主内容区，与 `shared_baseline` §3.7“不应在四实体详情页内重复长成第二主内容区”一致；本轮已进一步收紧为：`phase12` 遗留“项目上下文”设计需系统性退出前端，而不是并区保留
- SubTask 3.1 结论：可编辑集合冻结为 `template_source`（04 分类中唯一 optional 字段）、`canonical_root_files[]`（file_name/role/required）、`global_asset_bindings[]`（entry_ref/role/structured_summary，摘要仅适用前 5 份资产）；补充冻结禁止前端编辑 markdown 正文、顶层目录矩阵、`track_type / current_phase_*` 只读字段
- SubTask 3.2 结论：只读集合冻结为 `project_profile_version / track_type / docs_workflow_layout / current_phase_name / current_phase_ref / current_phase_status` + 顶层目录矩阵；与可编辑集合合并后对 phase13-04 的 9 类字段形成完整划分（可编辑 3 + 只读 6 = 9），无遗漏、无重叠；本轮已将顶层目录矩阵的前端来源正式提升为 spec requirement：其第一版前端承接方式固定为“当前项目范式 v1 的前端只读基线表达”，不得反推为新增后端治理字段、不得依赖运行时目录扫描；`global_asset_bindings[]` 的 `name / kind` 子字段不在可编辑子字段集中（05 写边界第 2 条资产侧仅列 entry_ref/role/structured_summary 为可写项），按只读资产清单基线呈现（由初始化建立）
- SubTask 3.3 结论：前端可编辑集合是 phase13-05 写白名单（“结构化治理字段 + 资产 entry_ref/role/structured_summary 承接结果”）的子集——05 冻结后端写路径上界、06 冻结第一版前端暴露面，两层关系无冲突；05 禁止的写入项（markdown 全文 / 自动扫描 / 自动同步）在 06 同样被禁止
- SubTask 4.1 结论：供人类维护的内容仅限结构化字段维护与资产承接结果维护（Requirement 3 集合），与 `dev_plan` phase13-06 范围“明确哪些内容供人类维护”对应
- SubTask 4.2 结论：供 agent 消费的内容为完整 `project brief for agent`、数组摘要解析协议与目录全文读取能力边界说明，属于 agent 消费主线（phase13-07 承接），不需要在前端做成并列主内容区；spec 补充冻结禁止把 agent-only 协议、解析规则或 IDE 目录读取能力说明做成产品主界面常驻大块说明
- SubTask 4.3 结论：只保留摘要回看的内容为 `phase13` 治理画像正式承接的全局规范资产 `role / structured_summary / entry_ref` 与当前阶段入口状态结构化回看；不再默认承接 `phase12 project-context` 叙事内容本身；前端可提供人类可理解的摘要回看以帮助校对 agent 读取结果
- SubTask 4.4 结论：spec 明确“前端不得承担 agent brief 的主消费职责”，与 `shared_baseline` §3.6（agent 简报是结构化只读输入、与 IDE 目录读取是协作关系）、phase13-05 承接位“Web 与 agent 共享治理事实的唯一后端来源”口径一致
- SubTask 5.1 结论：`Repository detail` 现有 `phase12-09`“项目上下文”区必须移除，不再允许作为治理画像并列区或摘要层换壳保留
- SubTask 5.2 结论：`Decision detail` 现有“共享项目上下文入口”卡片必须移除，第一版不再把“去看项目上下文”作为用户侧正式交互目标
- SubTask 5.3 结论：`phase12 project-context` 设计的退出规则已正式冻结为“系统性退出前端”，不得通过并区、换标题、轻量入口或摘要层延续旧叙事
- SubTask 6.1 结论：目录真实路径不得被当作面向普通用户的主内容，与 `shared_baseline` §3.5“不把这些文件的真实路径直接做成面向普通用户的主内容”、§3.7“目录真实路径、全文规则与阶段快照不应误做成面向普通用户的大块说明 UI”逐句对应
- SubTask 6.2 结论：`entry_ref` 的正式角色冻结为定位与回源入口（允许查看、复制或跳转），以轻量 locator / secondary metadata 方式呈现，不是正文主体
- SubTask 6.3 结论：`structured_summary` 优先于真实路径成为主阅读内容；第一版不允许把路径列表做成“项目治理画像”的主价值展示——该约束同时保护 phase13-05“摘要可持久化、正文只回源”双层模式在前端的正确表达
- SubTask 7.1 结论：承接位（Repository detail）、锚点（repository_id）、独立入口后置条件与 phase13-03 Requirement 3 逐条一致；phase13-03“不得把治理层偷换为 Repository 业务字段扩张”由本 spec“四实体原有业务详情内容仍保持各自主语义”承接
- SubTask 7.2 结论：可编辑集合 ⊆ 05 写白名单、只读集合不进入写路径、markdown 正文不可编辑、顶层目录矩阵不可编辑且不入库，与 phase13-05 Requirement 4/5 及其 REMOVED（禁止全文入库）单值一致
- SubTask 7.3 结论：承接位、三层信息架构、可编辑/只读集合、三类消费边界、顶层目录矩阵前端来源、phase12 旧设计退出规则与 entry_ref 展示约束均已单值冻结并配套机械判定 Scenario；`phase13-07`（agent brief 设计）与 `phase13-09`（前端实现）可直接承接——“06 先于 07/09”顺序的正式依据是本 spec MODIFIED Requirement；组件树、文件级映射等实现粒度由 `phase13-09` 实现规格在“轻量、非常驻大块”约束内继续收敛，但不得再重开 `phase12 project-context` 是否保留的规格判断
