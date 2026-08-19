# phase15_project_progress_timeline_foundation_dev_plan

## 1. 文档定位

- 本文档是 `phase15_project_progress_timeline_foundation` 的任务分解与推进计划，冻结该阶段全部子任务的范围、DoD 与依赖关系。
- 本文档的直接上游是同阶段 `architecture_plan` 与 `shared_baseline`；三者不一致时以 `shared_baseline` 的单值基线为准并回改。

## 2. 本阶段目标

- 单一主交付：项目推进时间轴最小主线（合同 → 存储 → 后端 → 前端 → agent 消费），无退役任务。
- 十一项裁决结论（①排序与能力定位 / ②repository 锚定 / ③append-only 事件流与派生当前值 / ④三轨 workflow / ⑤任务项级颗粒度 / ⑥web 手动维护于 Repository detail / ⑦evidence_ref 导航零托管 / ⑧source 预留仅 manual / ⑨3 RPC 无 Update / ⑩brief 摘要与完整流分层 / ⑪三重边界分离）是全部子任务的强制边界。

## 3. 子任务清单

### 第一组：边界收敛类子任务

### phase15-01 冻结 `Project Progress Timeline Foundation` 的范围边界、成功标准与非目标

- 范围：把十一项裁决结论与 `phase14-11` 进入条件收敛为 `phase15` 的单一主交付边界声明；产出 `.trae/specs/phase15_01_freeze_progress_timeline_scope_success_non_goals/` 三件套。
- DoD：spec 与三件套、根级文档单值一致；裁决矩阵与候选池顺延项留档；独立复核通过。

### phase15-02 冻结事件模型语义边界与 brief 演进对照

- 范围：三轨 × event_kind 合法矩阵（shared_baseline §3.3 前置确认）、`task_key` 格式规则冻结、派生规则冻结（当前 phase / 最新任务 / recent_events N=10）、brief 前后对照表（`progress = 9`、`BriefProgress` 字段清单、槽位 2/3/4 保持 reserved）、三重边界分离声明（plan.md / phase11 `PhaseEntry` / Decision）。
- DoD：合法矩阵与校验规则单值；brief 前后对照表留档；三重边界各有显式断言语句；独立复核通过。

### 第二组：实现设计类子任务

### phase15-03 产出数据模型与校验派生设计

- 范围：`progress_events` DDL 级设计（字段 / 枚举承载已冻结为 `TEXT + CHECK(IN ...)`〔沿 `0011`，不再选型〕/ 索引与三键读取全序 `(occurred_at DESC, created_at DESC, id DESC)` / 幂等 DDL，落 `0013`）；shared_baseline §3.3 的 9 条校验规则形式化（每条含判定逻辑 + 稳定错误码，可直接写单元测试）；§3.4 派生算法细节（SQL 或 Go 侧实现序，落实三键 tiebreak 链）。
- DoD：DDL 可直接进入迁移实现；校验规则每条有判定逻辑与错误码；派生算法有精确到排序键的实现描述。

### phase15-04 产出后端合同与读写边界设计

- 范围：`psco.progress.v1` 合同（`ProgressEvent` 消息 + 枚举 + 3 RPC 的请求/响应 envelope 与错误语义：`ListProgressEvents` 不分页 + `workflow_type` 过滤、`CreateProgressEvent` 校验规则执行位与 `repository_id` 不存在的 `invalid_argument` 语义、`DeleteProgressEvent` 误录修正语义、无 Update 的显式声明）；Go 模块分层（`backend/internal/progress/`：connect / service / repository）；`ProgressReader` 跨模块读取接口签名（输入 repository_id、输出事件集 + 派生摘要，沿 `StandardReader` 模式）；brief `progress = 9` 装配演进点。
- DoD：3 RPC 逐个有请求/响应/错误三要素定义；`BriefProgress` 内联消息与跨包导入关系单值；跨模块读取经独立 Read 接口；与 `.proto` 主线、ConnectRPC、canonical owner 单值化对齐。

### phase15-05 产出前端信息架构与维护入口设计

- 范围：Repository detail 内嵌"项目进度"区设计（时间轴倒序列表 + 当前 phase 派生卡 + 录入表单 + 误录删除确认）；录入表单最小摩擦交互（`occurred_at` 默认 now、`event_kind` 记住上次选择、`workflow_type × event_kind` 非法组合联动禁用）；`frontend/src/features/progress/` 切片结构（application / components / data / types.ts / index.ts）；与 Repository detail 既有布局（Standard 只读摘要区等）的分区关系；移动端适配对齐基线。
- DoD：组件树、文件级映射、交互规格（含非法组合禁用态与删除确认文案）达到"足以直接进入实现"的 DoD 标准；mutation 收敛切片固定承接位（project_rules §2.5）；无独立路由与导航项（裁决⑥）。

### 第三组：源代码实现类子任务

### phase15-06 落实后端主线

- 范围：`proto/psco/progress/v1/progress.proto` + buf 生成；`0013_phase15_progress_timeline.sql`（单表 + 枚举 `TEXT+CHECK` + 索引，幂等 DDL；落入 `database/migrations/` 即被 `RunMigrations` 自动登记——phase14-07 OBS-01 修复后机制，无需手工登记）；`backend/internal/progress/`（connect / service / repository 分层 + validate / errors 支撑文件单值化；3 RPC 全量实现 + 9 条校验规则 + 派生计算）；`projectcontext` 侧 `ProgressReader` 接口接入与 `GetProjectBrief` 装配 `progress = 9`（空态恒构造块）。
- DoD：`buf lint / build / breaking` 通过；`go build / vet / test` 通过；校验规则有单元测试覆盖（含非法用例矩阵）；brief 集成测试含 `progress` round-trip 与派生断言（含 phase_completed 后当前 phase 为空）。

### phase15-07 落实前端主线

- 范围：`frontend/src/features/progress/` 切片（query 纯只读 + mutation 固定承接位）；Repository detail 内嵌进度区（时间轴 + 当前卡 + 录入表单 + 删除）；与 Standard 只读摘要区的布局协排。
- DoD：`tsc --noEmit` 零错误；浏览器可完成"录入 phase_started → 录入多条 task_completed（含补录历史 occurred_at）→ 录入 note → 派生当前卡正确 → 删除误录 → phase_completed 后当前卡转完结态 → audit/fix 轨录入与过滤"完整会话。

### 第四组：验证验收类子任务

### phase15-08 完成进度时间轴的联调、dogfooding 与反回归验证

- 范围：固定样本录入 dogfooding 固定录入集（PSCO 自身 phase14 历史回放：1 条 phase_started + 11 条 task_completed + 1 条 phase_completed + 穿插 note；逐条明细〔task_key / title / occurred_at，取 git log 真实提交时间〕为本 spec 必备附件，保证 rerun 可复现）；固定问题取证（agent 经 brief 直答背景+进度 / ListProgressEvents 全量与三轨过滤 / append-only 断言 / 派生正确性 / 十一项裁决逐条验证 per shared_baseline §4）；工具链门禁；浏览器反回归矩阵（Repository detail 新增进度区 + Standard 摘要区协排回归 + 既有页面抽查）。
- DoD：固定问题全达标；十一项裁决验收门禁全绿；反回归矩阵全绿；acceptance_report 冻结验收结论；独立复核通过。

### 第五组：根级同步类子任务

### phase15-09 完成根级同步、阶段收口与下一阶段进入条件回写

- 范围：回写 `AGENTS.md / plan.md / docs/README.md / architecture_map.md / docs/phase/README.md`；冻结下一阶段进入条件（预期含 `source` 预留项排序：git 自动采集 / agent 写回 / 模板仓库自动接入 / 自动同步 / `standard_bindings` 扩展等候选池）。
- DoD：根级五文档单值一致；无孤岛；单一真相源不破坏；变更待用户确认后提交。

## 4. 明确不做

- git 自动采集与 agent 写回（`source = git / agent` 仅预留枚举位，不提供写入路径）
- MCP / CLI 消费通道、模板仓库自动接入、自动同步
- `standard_bindings` 目标类型扩展
- 进度事件与 Decision 互链
- 第六 CRUD 主实体化（独立路由 / 全局导航 / Dashboard 主卡片）
- `UpdateProgressEvent`（append-only 无更新语义）
- phase11 `GetProjectContextResponse` 的 `PhaseEntry` / 规则投影改动
- plan.md 正文接管与自动同步

## 5. 子任务依赖关系

- `phase15-02` depends on `phase15-01`
- `phase15-03` `phase15-04` `phase15-05` depend on `phase15-02`
- `phase15-06` depends on `phase15-03` `phase15-04`
- `phase15-07` depends on `phase15-05`（切片实现）+ `phase15-06`（浏览器完整会话验证的运行时前提）
- `phase15-08` depends on `phase15-06` `phase15-07`
- `phase15-09` depends on `phase15-08`
