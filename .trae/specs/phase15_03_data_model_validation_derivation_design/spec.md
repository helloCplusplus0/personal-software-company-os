# phase15-03 产出数据模型与校验派生设计 Spec

## Why

`phase15-02` 已冻结事件模型语义基线（合法矩阵、K-1~K-5 正则、派生规则含空值语义），但三处实现设计仍未落地到"可直接进入实现"粒度：`progress_events` 的 DDL 级设计（落 `0013` 迁移）、9 条校验规则的形式化（每条判定逻辑 + 稳定错误码，可直接写单元测试）、派生算法实现序（三键 tiebreak 链的精确落地）。`phase15-06`（后端实现）以本 spec 为直接设计上游。

`phase15-03` 是实现设计类子任务：纯设计文档冻结，不写任何代码、不建迁移文件、不改 proto；DDL 以注释版 SQL 草案承载（实现归 phase15-06）。

## What Changes

- 冻结 `progress_events` DDL 级设计：逐字段定义（含类型 / 约束 / 注释）、枚举承载 `TEXT + CHECK(IN ...)`、索引 `(repository_id, occurred_at DESC, created_at DESC)`、FK `ON DELETE RESTRICT`、幂等 DDL（`IF NOT EXISTS`）、迁移落点 `database/migrations/0013_phase15_progress_timeline.sql` 与自动登记机制
- 形式化 9 条校验规则：每条判定逻辑（伪码级）+ 稳定错误码（沿 standard `SCREAMING_SNAKE_CASE` 模式）+ 执行序冻结 + 报第一个错误策略 + 2 个 envelope 前置校验（非 9 条业务规则，显式区分）
- 冻结派生算法实现序：repository 层单一查询（SQL `ORDER BY` 三键链为唯一排序执行位）+ service 层纯函数派生（零 I/O 可单元测试）；三派生项精确算法（含 DESC 切片索引语义）；List / brief / web 当前卡共用同一查询与派生函数
- 冻结 task_key / title 的 TrimSpace 规范化边界（标识符 trim 后持久化 vs 自由文本原值持久化）

## Impact

- Affected specs:
  - `phase15-06`（本 spec 为其存储层 + 校验 + 派生的直接设计上游；DDL 草案逐字转写为 `0013` 迁移文件）
  - `phase15-04`（错误码清单为其合同错误语义设计的输入；`repository_id` 存在性校验承接位 = DP-2 归其裁决，本 spec 只声明语义边界不裁决承接位）
  - `phase15-02` spec（本 spec 为其形式化下游：矩阵 / K 正则 / 派生语义 1:1 转译，零语义新增）
- Affected code: 无（设计冻结；`database/migrations/` / `backend/internal/progress/` 零改动）
- 验收产物：本目录 `tasks.md / checklist.md`

## ADDED Requirements

### Requirement: progress_events DDL 级设计必须可直接进入迁移实现

本 spec SHALL 冻结以下 DDL 设计（phase15-06 逐字转写为 `0013_phase15_progress_timeline.sql`，字段 / 约束 / 索引 / 注释零再决策）：

```sql
-- 0013_phase15_progress_timeline.sql — Project Progress Timeline 事件流建表迁移（phase15-06）
--
-- 上游规格：phase15-03 数据模型与校验派生设计 Spec（本 DDL 逐字源）
-- 幂等：CREATE TABLE / INDEX IF NOT EXISTS，整文件可安全重放（沿 0011 第一段模式）
-- 登记：落入 database/migrations/ 即被 RunMigrations 按文件名升序自动登记执行
--       （phase14-07 OBS-01 修复后机制，无需手工登记）

-- progress_events：repository 锚定的三轨 append-only 推进事件流（裁决②③）
-- 无 Update 语义（裁决⑨）；"当前 phase"等为读取侧派生值，不落库（裁决③）
CREATE TABLE IF NOT EXISTS progress_events (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    -- 锚点：进度事实唯一归属仓库；RESTRICT 保护 append-only 历史
    -- （仓库被事件引用时不可物理删除，沿 0006 product_repositories FK 惯例）
    repository_id   UUID NOT NULL REFERENCES repositories(id) ON DELETE RESTRICT,
    -- 三轨 workflow（对齐 docs/ 三目录与三推进链，裁决④）
    workflow_type   TEXT NOT NULL CHECK (workflow_type IN ('phase', 'audit', 'fix')),
    -- 事件类型（裁决⑤：audit/fix 轨禁止 phase 边界标记——应用层规则 7 承接，DB 不建组合约束）
    event_kind      TEXT NOT NULL CHECK (event_kind IN ('phase_started', 'phase_completed', 'task_completed', 'note')),
    -- 任务项标识（可空；格式随 workflow×kind 矩阵变化，应用层 K-1~K-5 承接）
    task_key        TEXT NULL,
    -- 一句话标题（非空上限 200 字符，应用层承接）
    title           TEXT NOT NULL,
    -- 展开说明（可空上限 2000 字符，应用层承接）
    detail          TEXT NULL,
    -- 证据导航引用（/ 或 https:// 前缀，应用层承接；正文零托管，裁决⑦）
    evidence_ref    TEXT NULL,
    -- 来源（预留 manual/git/agent 三值；本阶段创建入口仅 manual，裁决⑧）
    source          TEXT NOT NULL CHECK (source IN ('manual', 'git', 'agent')) DEFAULT 'manual',
    -- 用户声明发生时间（允许补录历史，与 created_at 分离）
    occurred_at     TIMESTAMPTZ NOT NULL,
    -- 系统录入时间
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- 时间轴回看与派生计算统一读取序索引（三键链前两键 + repository 过滤；
-- 最终 tiebreak id DESC 由 ORDER BY 补齐，不重复入索引——沿 shared_baseline §3.2 冻结形态）
CREATE INDEX IF NOT EXISTS idx_progress_events_repository_sort
    ON progress_events (repository_id, occurred_at DESC, created_at DESC);
```

设计决策冻结（DDL 草案中注释承载 + 本节展开理由）：

1. **FK `ON DELETE RESTRICT`**：沿 `0006` `product_repositories` FK 惯例；语义 = 仓库被进度事件引用时不可物理删除，保护 append-only 历史完整性（裁决③"历史永不丢弃"的存储层落地）。
2. **DB 只承接单列枚举完整性**：`workflow_type` / `event_kind` / `source` 三列各建 `CHECK(IN ...)`；workflow×kind 组合法性与 task_key 格式**不建 DB 约束**，全部由应用层校验承接（沿 standard 模式：树校验 R1-R8 全在应用层，`0011` 只做单列 CHECK）。
3. **索引不含 `id DESC`**：索引保持 shared_baseline §3.2 冻结形态 `(repository_id, occurred_at DESC, created_at DESC)`；读取全序第三键 `id DESC` 由查询 `ORDER BY` 补齐（个人规模无索引全覆盖的性能诉求，不为 tiebreak 扩索引）。
4. **无 `updated_at` 列**：append-only 无更新语义（裁决⑨），不存在行级更新时间。
5. **无数据迁移段**：纯新建表（T7 裁决后时间轴零承接位），迁移文件只有建表段，无 DO 块迁移段、无 DROP 段（与 `0011` 三段结构对照显式排除）。

#### Scenario: DDL 可逐字转写

- **WHEN** phase15-06 实现迁移文件
- **THEN** 以上 SQL 草案逐字转写为 `0013_phase15_progress_timeline.sql`（仅可调整：无），字段 / 约束 / 索引 / 注释 / 文件头格式零再决策
- **AND** 迁移文件落入 `database/migrations/` 后由 `RunMigrations` 按文件名升序自动登记执行，无需任何手工登记改动

### Requirement: 9 条校验规则必须形式化为判定逻辑 + 稳定错误码

本 spec SHALL 将 phase15-02 冻结的 9 条规则形式化（编号 V1-V9 与规则 1-9 一一对应），每条含判定逻辑与稳定错误码；错误统一包装 `ErrInvalidInput`（映射 Connect `CodeInvalidArgument`），错误信息格式沿 standard 模式 `%w: [CODE] message`。

**稳定错误码总表**（业务码 9 个 + envelope 前置码 2 个）：

| 错误码 | 对应规则 | 判定逻辑（伪码） |
|---|---|---|
| `INVALID_WORKFLOW_TYPE` | V1a | `workflow_type ∈ {"phase","audit","fix"}`（proto 枚举 UNSPECIFIED / 越界字符串 → 失败） |
| `INVALID_EVENT_KIND` | V1b | `event_kind ∈ {"phase_started","phase_completed","task_completed","note"}`（同上） |
| `EVENT_KIND_NOT_ALLOWED` | V7 | `NOT (workflow_type ∈ {"audit","fix"} AND event_kind ∈ {"phase_started","phase_completed"})`（枚举值各自合法但组合被矩阵禁止） |
| `TASK_KEY_REQUIRED` | V2,V3,V4,V5,V6 必填 | 按 workflow×kind 矩阵：`task_key` 必填格 `TrimSpace(task_key) ≠ ""`（V2 通用 task_completed + V3-V6 各矩阵格的必填前置） |
| `TASK_KEY_FORMAT_INVALID` | V3,V4,V5,V6 格式 | 必填通过后按矩阵格匹配 K-1 / K-2 / K-3 / K-4；错误信息携带期望格式说明（如 `expected ^phase[0-9]{2,}$`） |
| `INVALID_TITLE` | V9a | `TrimSpace(title) ≠ "" AND runeLen(title) ≤ 200`（空与超长同码，信息区分） |
| `INVALID_DETAIL` | V9b | `detail = "" OR runeLen(detail) ≤ 2000`（可空字段仅校验上限） |
| `INVALID_EVIDENCE_REF` | V9c | `evidence_ref = "" OR HasPrefix(evidence_ref, "/") OR HasPrefix(evidence_ref, "https://")` |
| `INVALID_SOURCE` | V9d | `source = "manual"`（现值约束：`git` / `agent` 为预留枚举位，本阶段创建入口拒绝） |

**V8（note 轨 task_key 可空）无独立错误码**：V8 是"允许"规则（合法路径不产生错误），其语义已由 V2-V6 的矩阵分支覆盖（note 格不走必填与格式分支）。

**envelope 前置校验（非 9 条业务规则，proto 解包层输入完整性，显式区分防规则清单 1:1 破坏）**：

| 错误码 | 判定逻辑 |
|---|---|
| `INVALID_REPOSITORY_ID` | `repository_id` 为合法 UUID 格式（非 UUID 字符串 → 失败；存在性校验承接位归 phase15-04 DP-2 裁决，本 spec 只冻结格式层） |
| `INVALID_OCCURRED_AT` | `occurred_at` 已设置（proto Timestamp nil / 零值 → 失败；`occurred_at` 是用户声明发生时间，无合法零值） |

**执行序冻结**（报第一个错误，沿 standard DFS 报第一个错误模式）：

1. envelope 前置：`repository_id` UUID 格式 → `occurred_at` 已设置
2. V1a `workflow_type` 枚举 → V1b `event_kind` 枚举
3. V7 矩阵组合判定（先于 task_key 判定：组合非法时 task_key 无从校验）
4. task_key 矩阵分支（V2-V6）：先必填（`TASK_KEY_REQUIRED`）后格式（`TASK_KEY_FORMAT_INVALID`）
5. V9 文本字段顺序：title → detail → evidence_ref → source
6. repository 存在性校验（语义 `invalid_argument`；承接位 = phase15-04 DP-2 裁决；DB FK RESTRICT 为存储层兜底，非校验承接位）

**TrimSpace 规范化边界冻结**：

- `task_key`（结构化标识符，用于匹配 / 过滤 / 派生）：服务端 `TrimSpace` 后判定**并持久化 trim 后值**——标识符字段 trim 是无害规范化
- `title` / `detail` / `evidence_ref`（自由文本 / 引用）：`TrimSpace` 仅用于必填判定（title），**原值持久化**（沿 standard `name` 处理模式：TrimSpace 非空校验 + 原值入库）
- 长度计量单位：Unicode 字符数（Go rune 计数，沿 standard `maxSummaryRunes` 模式；中文标题按字符而非字节计）

#### Scenario: 校验规则可直接写单元测试

- **WHEN** phase15-06 编写校验单元测试
- **THEN** 每条规则的判定逻辑可直接转译为表驱动用例：合法矩阵 12 格正例 + 逐规则反例（每码至少 1 个触发用例），执行序可经构造多错输入断言首个报错
- **AND** 测试断言按错误码（非错误信息全文）定位，错误信息格式变化不破坏测试

### Requirement: 派生算法实现序必须精确到排序键

本 spec SHALL 冻结派生算法实现序为"repository 单一查询 + service 纯函数"方案：

**实现序设计**：

- **排序唯一执行位 = SQL `ORDER BY`**：repository 层提供单一查询（List RPC / brief 装配 / 派生计算共用同一查询语句）；Go 侧**不重排**，直接消费 SQL 返回序
- **repository 层查询**（唯一读取语句形态）：

```sql
SELECT id, repository_id, workflow_type, event_kind, task_key, title,
       detail, evidence_ref, source, occurred_at, created_at
  FROM progress_events
 WHERE repository_id = $1
   [AND workflow_type = $2]  -- 可选过滤（ListProgressEvents 的 workflow_type 参数；派生计算不过滤）
 ORDER BY occurred_at DESC, created_at DESC, id DESC
```

- **service 层纯函数派生**：`DeriveProgressSummary(events []ProgressEventRecord) ProgressSummary`——零 I/O、零外部依赖、零时间函数（输入即全部信息），可纯单元测试覆盖全部边界（含空值双情形）
- **共用实现约束落地**：List RPC 响应、brief `progress` 装配（经 ProgressReader）、web 当前卡数据通道（DP-1 裁决后）全部消费同一 repository 查询 + 同一派生函数（shared_baseline §3.4"web 当前卡与 brief 摘要共用同一实现"的实现形态）

**三派生项精确算法**（输入 `events` 为三键链 DESC 全序切片；**索引 i 越小 = 事件越晚**）：

1. **recent_events**：`events[0 : min(10, len(events))]`；`len(events) = 0` → 空数组（N=10 冻结）
2. **latest_task_completed**：自 `i = 0` 起找第一个 `event_kind = "task_completed"` 的元素；不存在 → 零值（proto 侧不设置）
3. **current_phase_key / current_phase_label**：
   - 找 `latestStartedIdx` = 第一个 `event_kind = "phase_started"` 的索引
   - 无 `phase_started` → 空值（**从未开始**）
   - 有：检查是否存在 `j < latestStartedIdx` 使 `events[j].event_kind = "phase_completed"` 且 `events[j].task_key = events[latestStartedIdx].task_key`（`j` 更小 = 序更晚；同 key 匹配）→ 存在 → 空值（**全部完结**）
   - 否则：`current_phase_key = events[latestStartedIdx].task_key`；`current_phase_label = events[latestStartedIdx].title`
   - 两种空值情形输出同型（空字符串零值，沿 phase15-02 冻结；不引入第二套状态字段）

**算法边界说明**：

- 个人规模假设：全量加载单仓库事件集（百级）无压力；不分页已由裁决⑨冻结；万级以上优化属未来进入条件（本阶段不设计）
- `ORDER BY` 第三键 `id DESC` 的确定性：PostgreSQL UUID 排序为字节序，与 Go `uuid.UUID` 字节比较一致；同 `(occurred_at, created_at)` 碰撞时结果确定（补录同刻与未来批量插入两类场景，沿 shared_baseline §3.2 tiebreak 论证）

#### Scenario: 派生算法可直接进入实现与测试

- **WHEN** phase15-06 实现 repository 查询与 service 派生函数
- **THEN** SQL 语句逐字采用上述形态（仅绑定参数与可选过滤子句按调用方拼接）；派生函数按三算法逐行实现，无第二套排序或派生路径
- **AND** 单元测试可纯内存构造 events 切片覆盖：空集 / 仅 note / 从未开始 / 进行中 / 全部完结 / 补录同刻 tiebreak / recent 截断（>10 条）全部边界，无需数据库

### Requirement: 本 spec 与 phase15-02 语义上游单值一致且不偷渡后续合同与前端设计

本 spec 的全部形式化内容 SHALL 与 `phase15-02` spec 冻结语义 1:1 转译（零语义新增、零收窄、零放宽）：

- DDL 字段与 `shared_baseline` §3.2 字段矩阵逐字段一致（11 列 + 类型 + 约束）；枚举三列值域与 §3.2 / §3.3 一致；索引与 §3.2 冻结索引一致
- V1-V9 判定逻辑与 phase15-02 合法矩阵 12 格、K-1~K-5 正则、9 条规则清单逐条对应（V8 为允许规则无码属结构性处理，非语义改动）
- 派生三算法与 phase15-02 派生规则 4 项 + 三键链 + 空值双情形同型语义逐项一致
- 本 spec 不承载：RPC 请求 / 响应 envelope 定义与字段号、`ProgressReader` 接口签名、`BriefProgress` 字段号分配、`repository_id` 存在性校验承接位裁决（DP-2）、前端组件树与交互规格——以上归 `phase15-04` / `phase15-05`
- 本 spec 不承载：Go 代码、迁移文件实体、proto 文件实体（实现归 `phase15-06`）

#### Scenario: 一致性可校验

- **WHEN** 独立复核执行
- **THEN** DDL 草案 11 列与 §3.2 字段矩阵逐列比对一致；V1-V9 与 phase15-02 规则清单逐条比对一致；三派生算法与 phase15-02 派生冻结逐项一致
- **AND** git 工作区中本 spec 仅为目录新增，零代码 / 零迁移文件 / 零 proto / 零根级文档改动
