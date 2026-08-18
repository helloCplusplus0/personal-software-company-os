# phase14-03 Standard 数据模型与目录树设计 Spec

## Why

phase14-01/02 已冻结范围与关系边界，但 `standards` / `standard_revisions` / `standard_bindings` 三表仍停留在字段矩阵层面，`DirectoryTreeNode` 的 jsonb 精确序列化规范、8 条树校验规则的形式化判定逻辑、以及旧画像两表数据到树节点的字段级迁移映射（含 `required` / `kind` 字段去向与同名冲突合并）尚未定义。本子任务产出可直接进入 `0011` 迁移与 proto 实现的完整设计，实现类子任务（phase14-07/09）不再做任何设计决策。本子任务为实现设计类，交付设计文档，不写实现代码、不创建迁移/proto 文件。

## What Changes

- 产出本 spec 三件套（`phase14_03_design_standard_data_model_directory_tree`）
- 三表 DDL 级设计草案（对齐仓库既有迁移风格：`gen_random_uuid()` / `TEXT+CHECK` 受控枚举 / `TIMESTAMPTZ` / 内联 UNIQUE / `idx_` 前缀索引），作为 `0011_phase14_standard_entity.sql` 的直接输入
- `DirectoryTreeNode` jsonb 精确序列化规范（Go 结构 + omitempty 语义 + nil/[] 等价规则）与 proto 字段对齐表（phase14-04 直接引用）
- 8 条树校验规则形式化（每条：判定逻辑 + 错误码 + 节点路径定位要求），补充字段长度细节（role 1-32、summary ≤2000）
- 旧画像两表 → 树节点的字段级迁移映射冻结：`required` 显式退役（结构性吸收）、`kind` 并入 summary 前缀格式、`entry_ref` → `ref` 规范化规则（裸名→`/` 前缀绝对路径、多级路径拆树、URL 原值保留）、同名冲突合并规则
- 根节点表示决策冻结：单根 `name="."`（仓库根语义）

## Impact

- Affected specs:
  - `docs/phase/phase14_standard_entity_foundation_dev_plan.md`（上游，phase14-03 定义 L29-32）
  - `docs/phase/phase14_standard_entity_foundation_shared_baseline.md`（§3.2 字段矩阵 / §3.3 节点矩阵与 8 条校验规则 / §3.4 绑定矩阵——本 spec 细化不修改）
  - `.trae/specs/phase14_02_freeze_standard_entity_relation_boundaries/spec.md`（绑定矩阵与零丢失映射的上游约束）
- Affected code: 无（零代码改动；设计对象为未来的 `database/migrations/0011_phase14_standard_entity.sql`、`proto/psco/standard/v1/standard.proto`、`backend/internal/standard/`，均由 phase14-07/09 落地）
- 设计参照（本轮实际读取）：`database/migrations/0010_phase13_governance_profile.sql`（DDL 风格 + 旧两表结构：`file_name/role/required` 与 `name/kind/entry_ref/role/structured_summary`）

## ADDED Requirements

### Requirement: 三表 DDL 级设计草案必须冻结

`0011_phase14_standard_entity.sql` 的表结构 SHALL 按以下草案实现（风格对齐 0010：注释块头 + 字段级行注释 + `TEXT+CHECK` 枚举 + `idx_` 前缀索引）：

```sql
-- standards：全局规范主表（全局作用域，无 repository 锚点——裁决⑤）
CREATE TABLE standards (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    -- 规范名（全局唯一，如 "Durable System Track 项目范式"）
    name            TEXT NOT NULL UNIQUE,
    -- 一段定位说明（可空）
    description     TEXT NULL,
    -- 生命周期状态（受控枚举）
    status          TEXT NOT NULL CHECK (status IN ('draft', 'active', 'retired')) DEFAULT 'draft',
    -- 目录树整树原子承载（单根结构，规范见 DirectoryTreeNode jsonb 规范）
    directory_tree  JSONB NOT NULL,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- standard_revisions：演进留痕（CON-06 最小承接；只追加，不更新）
CREATE TABLE standard_revisions (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    standard_id     UUID NOT NULL REFERENCES standards(id) ON DELETE CASCADE,
    -- 人工一句话说明本次调整（必填）
    change_summary  TEXT NOT NULL,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
CREATE INDEX idx_standard_revisions_standard_id ON standard_revisions (standard_id);

-- standard_bindings：多态绑定关系表（CON-02 不设限；target_id 无 DB 外键为多态必然，存在性由应用层校验）
CREATE TABLE standard_bindings (
    id           UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    standard_id  UUID NOT NULL REFERENCES standards(id) ON DELETE CASCADE,
    -- 绑定目标类型（受控枚举，可扩展：扩 enum + 登记 phase14-02 八格矩阵）
    target_type  TEXT NOT NULL CHECK (target_type IN ('repository', 'product', 'decision', 'module')),
    -- 绑定目标 id（多态：按 target_type 应用层查对应实体表校验存在性）
    target_id    UUID NOT NULL,
    -- 绑定角色（受控枚举，可扩展；组合合法性按 phase14-02 八格矩阵：template_source 仅 repository）
    role         TEXT NOT NULL CHECK (role IN ('template_source', 'adopts')),
    -- 绑定备注（可空）
    note         TEXT NULL,
    created_at   TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    -- 唯一约束：同一 (standard, target, role) 组合拒绝重复绑定
    UNIQUE (standard_id, target_type, target_id, role)
);
-- brief 反查索引：GetProjectBrief 按 (target_type='repository', target_id) 反查关联 Standard
CREATE INDEX idx_standard_bindings_target ON standard_bindings (target_type, target_id);
CREATE INDEX idx_standard_bindings_standard_id ON standard_bindings (standard_id);
```

设计决策说明：

- `directory_tree JSONB NOT NULL`：树永远是单根结构，创建时无内容则写入单根空树（见根节点表示决策），不存在无树状态
- `standard_bindings.target_id` 不设 DB 外键是多态绑定的必然结果（无法指向多张表）；存在性与组合合法性由 service 层校验（phase14-07 落地，错误语义 `invalid_argument`）
- `updated_at` 维护责任在 service 层（整树替换事务内同步更新）

#### Scenario: 迁移实现判定

- **WHEN** phase14-07 编写 `0011_phase14_standard_entity.sql`
- **THEN** 表名、列名、类型、约束、索引与本草案逐字一致；注释风格对齐 0010

### Requirement: DirectoryTreeNode jsonb 序列化规范必须冻结

`directory_tree` 列的 jsonb 内容 SHALL 满足以下精确规范：

Go 侧结构（`backend/internal/standard/` 的序列化目标，phase14-07 实现）：

```go
// DirectoryTreeNode 目录树节点（directory_tree jsonb 与 proto 的单值映射结构）
type DirectoryTreeNode struct {
    Name     string               `json:"name"`
    NodeType string               `json:"node_type"`          // "directory" | "file"
    Role     string               `json:"role,omitempty"`     // file 必填（校验层保证）；directory 可空
    Summary  string               `json:"summary,omitempty"`  // 结构化摘要（裁决①承接位）
    Ref      string               `json:"ref,omitempty"`      // 定位引用：/ 开头树内路径或 https:// URL
    Children []*DirectoryTreeNode `json:"children"`
}
```

序列化与等价规则：

1. `omitempty`：`role / summary / ref` 空字符串时序列化省略；反序列化缺失视为空字符串
2. `children`：`nil` 与 `[]` 等价（file 节点允许 `"children":[]` 或省略；directory 无子节点同样允许两者）；校验只要求 file 节点无非空 children
3. 根节点表示决策：**单根结构，根 `name` 固定为 `"."`（仓库根语义，Unix 风格）**；根 `node_type` 必须为 `directory`；根的 `role` 必须为空、`summary` 可选（承载整份范式的一句话说明）；根不参与同层唯一性判定（无同层）
4. proto 对齐表（phase14-04 落 proto 时逐字使用）：jsonb `name/node_type/role/summary/ref/children` ↔ proto `DirectoryTreeNode` 的 `string name / NodeType node_type（NODE_TYPE_DIRECTORY=1, NODE_TYPE_FILE=2）/ string role / string summary / string ref / repeated DirectoryTreeNode children`；字段名由 buf 按 snake_case→camelCase 惯例处理，Go jsonb 结构与 proto 生成结构字段名一一对应

示例（树形态与序列化样式演示；注意：phase13-11 固定样本的两表 `entry_ref` 实际均为根级路径——验收报告实证 8 项资产"entry_ref 均为资产自身根级路径"——故 phase14-09 真实迁移产物全部挂根下，不含拆树节点；`docs/phase/README.md` 拆树分支为通用规则演示，前向适用）：

```json
{
  "name": ".",
  "node_type": "directory",
  "summary": "Durable System Track 项目范式 v1",
  "children": [
    { "name": "AGENTS.md", "node_type": "file", "role": "entry", "summary": "[agents] agent 全局上下文入口", "ref": "/AGENTS.md" },
    { "name": "docs", "node_type": "directory", "children": [
        { "name": "phase", "node_type": "directory", "children": [
            { "name": "README.md", "node_type": "file", "role": "docs", "summary": "[summary] phase workflow 索引", "ref": "/docs/phase/README.md" }
        ]}
    ]}
  ]
}
```

#### Scenario: 序列化实现判定

- **WHEN** phase14-07 实现 Go 结构与 jsonb 读写
- **THEN** 结构定义、json tag、omitempty 语义与本规范逐字一致；nil/[] 等价规则在反序列化测试中覆盖

### Requirement: 8 条树校验规则必须形式化冻结

校验入口：`CreateStandard` / `UpdateStandard`（含整树替换与状态变更），service 层 `validate`（phase14-07 落地）。每条规则 SHALL 满足：判定逻辑单值、错误码稳定、错误信息含节点路径定位（自根起 `/` 连接，如 `/docs/phase`）：

| # | 规则 | 判定逻辑 | 错误码 |
|---|---|---|---|
| R1 | 单根 directory 且根 name="." | 树对象即根（无兄弟）；`node_type=="directory"`；`name=="."`；根 `role` 为空 | `INVALID_TREE_ROOT` |
| R2 | 空树仅 draft 合法 | 树中 `file` 节点计数 ≥1，除非 `status=="draft"`；在创建、整树替换、状态变更（draft→active/retired）三种时机检查 | `EMPTY_TREE_NOT_ALLOWED` |
| R3 | 同层 name 唯一 | DFS 每层建立 name 集合查重（大小写敏感）；根豁免（无同层） | `DUPLICATE_SIBLING_NAME` |
| R4 | name 字符集与长度 | 非根节点：`^[A-Za-z0-9._-]{1,64}$`；根固定 `"."` 天然满足 | `INVALID_NODE_NAME` |
| R5 | 深度上限 6 | 根=第 1 层，后代最深第 6 层；第 6 层只允许 `file`（第 6 层 directory 必无 children，等价于不允许） | `TREE_TOO_DEEP` |
| R6 | 序列化上限 64KB | `len(json.Marshal(tree)) <= 65536` 字节 | `TREE_TOO_LARGE` |
| R7 | file 节点完整性 | `file` 节点：`children` 为空、`role` 非空且长度 1-32；`summary` ≤2000 字符（directory 的 summary 同限） | `INVALID_FILE_NODE` / `INVALID_SUMMARY_LENGTH` |
| R8 | ref 格式 | 空（省略）或以 `/` 开头（树内绝对路径）或以 `https://` 开头（外部 URL）；ref 为树内路径时应与节点在树中的实际路径一致（路径自洽性，见迁移映射的 ref 规范化） | `INVALID_REF` |

补充细节（字段级细化，不修改 baseline §3.3 主干）：

- `role` 约定值域 `entry / plan / rules / baseline / spec / summary` 为软约定（校验不硬卡，允许自定义短标签；硬校验仅长度 1-32），延续 baseline"允许自定义"冻结
- `name` 中 `.` 与 `..` 作为普通字符串合法（在字符集内），不赋予特殊路径语义（树位置即路径，name 仅是段名）

#### Scenario: 校验实现判定

- **WHEN** phase14-07 实现 validate 层
- **THEN** 8 条规则逐条有单元测试（含每个错误码的非法用例与合法边界用例）；错误信息含节点路径

### Requirement: 旧画像数据到树节点的迁移映射必须冻结（字段级）

phase14-09 迁移执行的字段级映射 SHALL 冻结如下（数据源为 0010 两表，目标为单条全局 Standard 的 `directory_tree`，裁决⑧）。时间戳列口径：两源表的 `created_at` 按 phase14-02 已冻结的显式不迁移清单处理（`created_at + updated_at` 不迁移；目标 Standard 的 `created_at/updated_at` 由迁移写入时刻生成，首条 revision 承接合并来源留痕）。

**源 1：`governance_canonical_root_file_bindings`（`file_name / role / required`）→ 根下 file 节点**

| 旧字段 | 映射 | 说明 |
|---|---|---|
| `file_name` | `name` | 原值 |
| `role` | `role` | 原值保留不归一化（baseline §3.3 冻结） |
| `required` | **显式退役** | 全局规范语义下无项目级"当前项目是否必需"概念（规范对所有采用者一致）；"入树即规范声明文件"结构性吸收必需语义；不进入零丢失验收范围（同 `markdown_resolvable` 处理模式） |
| （无） | `summary` = `""` | 旧表无摘要 |
| （派生） | `ref` = `"/" + file_name` | 根级文件规范化为 `/` 前缀绝对路径 |
| （固定） | `node_type` = `"file"`、`children` 空 | |

**源 2：`governance_global_asset_bindings`（`name / kind / entry_ref / role / structured_summary`）→ file 节点（含路径拆树）**

| 旧字段 | 映射 | 说明 |
|---|---|---|
| `name` | `name` | 原值（节点名=最后路径段的文件名，见 entry_ref 拆树） |
| `role` | `role` | 原值保留 |
| `structured_summary` + `kind` | `summary`：`structured_summary` 非空 → `"[" + kind + "] " + structured_summary`；为空 → `"[" + kind + "]"` | phase14-02 冻结"kind 并入摘要语义"的单值化格式 |
| `entry_ref` | `ref` + 节点位置（拆树） | 规范化规则见下 |

**`entry_ref` 规范化与拆树规则（单值）**：

1. `entry_ref` 以 `https://` 开头 → `ref` 原值保留，节点挂根下（外部资源无树内路径）
2. `entry_ref` 为裸文件名（无 `/`，如 `project_rules.md`）→ 节点挂根下，`ref = "/" + entry_ref`
3. `entry_ref` 含路径（如 `docs/phase/README.md`，可能带前导 `/`）→ 去前导 `/` 后按 `/` 拆段：前面各段生成（或复用已存在的）`directory` 节点（`name`=段名、`role` 空、`summary` 空），最后一段为 `file` 节点；`ref` = 规范化全路径（`/` 前缀）
4. 拆树合并语义：同路径段复用既有 directory（多行数据共享目录树）；`external_url` 语义说明：0010 表无 `external_url` 列（本轮已实证），存量仅 `entry_ref` 形态；ref 的 URL 能力（规则 8）为前向能力，承接讨论中 `external_url` 的语义定位

**同名冲突合并规则（两源迁移到同一棵树的冲突处理）**：

- 迁移顺序固定：先源 1（canonical，全在根下），后源 2（global assets，含拆树）
- 冲突判定：同一父节点下出现同名 `file` 节点（如 `AGENTS.md` 同时存在于两表——phase13-11 样本即如此）
- 合并规则：保留先建节点的位置与 `ref`；源 2 行**增强**既有节点——`role` / `summary` 以源 2 非空值覆盖（源 2 信息更丰富：含 kind 前缀摘要与更细 role）；源 2 值为空则保留源 1 值
- 裁决⑧多画像合并：当前仅 1 个画像实例无此场景；若未来多画像不一致，字段取最新 `updated_at` 画像的值，差异记入首条 revision（沿用已冻结口径）

#### Scenario: 迁移实现判定

- **WHEN** phase14-06 设计迁移脚本 / phase14-09 执行迁移
- **THEN** 字段映射、拆树规则、冲突合并与本节逐字一致；迁移输出树通过 8 条校验（迁移产物天然合法）

#### Scenario: 零丢失验收范围核对

- **WHEN** phase14-10 执行零丢失对照验收
- **THEN** 验收范围 = phase14-02 冻结的 5 字段映射（`name/kind/entry_ref/role/structured_summary`）+ 本节源 1 的 `file_name/role`；`required` 与 `markdown_resolvable` 为显式退役项，不在验收范围（理由已留档）

## MODIFIED Requirements

无（本 spec 为新增设计规格；对 shared_baseline §3.2/§3.3/§3.4 仅为字段级细化，全部只收紧不放宽，不构成对主干规则的修改。细化清单：role 长度 1-32、summary ≤2000、根 name="."、children nil/[] 等价、R2 按 file 节点计数判定（覆盖"仅 directory 无 file"情形，对齐 §3.3 规则 2 第二句"active 必须有至少一个 file 节点"）、R5 第 6 层只允许 file（"第 6 层 directory 必无 children"与深度上限等价）、R8 ref 路径自洽性（树内路径 ref 必须与节点实际位置一致，服务 agent 无转译直读，格式约束之上的强化）。

## REMOVED Requirements

无。
