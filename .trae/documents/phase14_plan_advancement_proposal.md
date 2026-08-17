# phase14 /plan 推进计划

## Summary

phase14 定位为 **Standard 全局规范实体基础（单一主交付）+ 治理画像重叠承接位退役**。推进链路：五项结构化裁决（phase13-12 冻结的强制前提）→ 产出三件套 → 子代理复核 → 根级同步。CON-08 阶段推进时间轴列入 phase14 非目标（留待 phase15），`current_phase` 三字段去向并入裁决④处理。

## Current State Analysis

- **阶段状态**：phase13 已正式收口（plan.md §1 已更新）；`.trae/specs/phase13_12_sync_root_level_closeout_next_phase_entry_conditions/spec.md` 为 phase14 唯一规划上游，冻结 GAP-01~07 + CON-01~09 + 五项裁决进入条件
- **画像实现足迹（裁决④退役范围清单）**：
  - 合同：`proto/psco/governance_profile/v1/governance_profile.proto`（L64-122 核心消息含 `canonical_root_files[]` / `global_asset_bindings[]`）
  - 存储：`database/migrations/0010_phase13_governance_profile.sql`（`governance_canonical_root_file_bindings` L47-58、`governance_global_asset_bindings` L69-84）
  - 后端：`backend/internal/governanceprofile/`（connect/service/repository 分层）
  - 前端：`frontend/src/features/governance-profile/` 切片 + Repository detail 挂载位
  - 消费方：`project_context.proto` L174-196 `GetProjectBriefResponse` 引用画像完整消息；`backend/internal/projectcontext/query_service.go` 编排
- **四实体建模模式（Standard 参照基线）**：`proto/psco/<entity>/v1` 合同目录、`backend/internal/<module>` connect/service/repository 分层、`frontend/src/features/<slice>` 切片、编号迁移文件
- **三件套模板**：phase13 三件套（architecture_plan / dev_plan / shared_baseline），dev_plan 含五类子任务（边界收敛 / 实现设计 / 源代码实现 / 验证验收 / 根级同步）

## Proposed Changes

### Step 1: 五项裁决 Round 1（结构化问答，3 问）

每问给出选项 + 利弊 + 推荐方案，用户逐项拍板：

- **裁决① 信息维护颗粒度（CON-05）**：A) 混合式——PSCO 承接结构 + 导航 + 结构化摘要，全文以模板仓库为唯一事实源（推荐，符合 CON-07 效率底线）/ B) 仅文件名导航（最轻）/ C) PSCO 托管全文（最重）
- **裁决② Standard 数据模型与 pg 承载**：A) 主表 + jsonb 树形目录结构 + 绑定关系表，proto 用树形消息对 agent 原生友好（推荐）/ B) 全范式化树节点表 / C) 单表全 jsonb
- **裁决③ Standard 与模板仓库内容边界**：A) PSCO=结构与导航+演进记录，仓库=内容事实源（推荐，与裁决①A 联动）/ B) PSCO 含全文内容 / C) 仓库承载全部、PSCO 仅外链

### Step 2: 五项裁决 Round 2（结构化问答，2 问）

- **裁决④ 画像退役计划（CON-04）**：A) phase14 内完整退役——迁移数据至 Standard、删除 proto 重叠字段、brief 改读 Standard；`current_phase` 三字段去向一并裁决（暂留 brief 派生或随退役移除，时间轴 phase15 承接）（推荐）/ B) 兼容期保留 deprecated 至 phase15 删除 / C) 仅冻结写入不删
- **裁决⑤ Standard 实体地位**：A) 全局规范资产实体——治理层地位、全局作用域（非 repository 锚点）、独立维护入口但不做第五 CRUD 主实体（推荐，符合 phase13-12 衔接声明第 1 条的重新仲裁要求）/ B) 第五业务主实体（独立导航 CRUD 主线）

裁决结果全部冻结进 shared_baseline，标注与 GAP/CON 的对应关系。

### Step 3: 产出 phase14 三件套

- `docs/phase/phase14_standard_entity_foundation_architecture_plan.md`：定位与边界（信息分层模型更新）、Standard 实体设计（字段矩阵、目录树表示、绑定关系矩阵）、与画像退役的映射、agent 消费链路（GetProjectBrief 演进）
- `docs/phase/phase14_standard_entity_foundation_dev_plan.md`：五类子任务编号（phase14-01 边界收敛 → 02/03 实现设计 → 04+ 源码实现（proto/迁移/后端/前端/退役）→ 验证验收 → 根级同步），每任务附 DoD
- `docs/phase/phase14_standard_entity_foundation_shared_baseline.md`：五项裁决结论、能力矩阵、非目标（CON-08 时间轴 / agent 写回 / Git 推进跟踪 / 模板仓库自动同步）

### Step 4: 子代理独立复核三件套

复核清单：GAP-01~07 / CON-01~09 覆盖矩阵（逐条映射到三件套章节）、五项裁决结论与用户拍板结果一致、内部单值性（无第二事实源 / 无双写路径）、DoD 可执行性、与 project_rules §2.6 传输约束（.proto + ConnectRPC）对齐。

### Step 5: 根级同步（五文档）

- `plan.md`：§1 当前阶段更新为 phase14 planned、§2 登记三件套、§3 新增 phase14 块（状态 planned + 五项裁决结论摘要 + 进入条件）
- `AGENTS.md`：§1 当前阶段与主目标、§4 当前状态、§5 阅读顺序
- `docs/README.md` / `architecture_map.md` / `docs/phase/README.md`：登记三件套入口与阶段状态

## Assumptions & Decisions

- phase14 = Standard 单一主交付（用户已确认）；CON-08 时间轴为 phase14 非目标
- 裁决方式 = 结构化逐项（用户已确认），两轮共 5 问
- agent 写回、Git 推进跟踪、模板仓库自动接入、自动同步继续不进入 phase14（CON-09 + phase13-12 进入条件第 4 条）
- phase14 命名暂定 `phase14_standard_entity_foundation`，可在三件套产出时微调
- 三件套仅规划冻结，不含代码实现（实现按 dev_plan 子任务经 /spec 逐个推进）

## Verification Steps

1. 三件套复核报告 PASS（GAP/CON 覆盖矩阵无缺口、裁决结论无漂移）
2. 根级五文档同步后：无孤岛（三件套可从根级入口到达）、无过时表述（phase13"待裁决"类表述已更新）、单一真相源（phase 路线状态仅 plan.md 承载）
3. 五项裁决留痕完整（两轮问答结果 + shared_baseline 冻结条目一一对应）
