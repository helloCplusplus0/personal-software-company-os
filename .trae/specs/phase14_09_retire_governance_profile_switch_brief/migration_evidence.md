# phase14-09 Task 5 迁移证据：0011 第二、三段补全与验收 DB 执行

> 执行日期：2026-08-18（Asia/Shanghai）
> 执行对象：`database/migrations/0011_phase14_standard_entity.sql`（第二段数据迁移 DO 块 + 第三段 drop 两表）
> 验收 DB：`psco_development`（容器 `rento-preview-postgres`，PostgreSQL 16-alpine，`127.0.0.1:55432`；宿主机 psql 不可用，按仓库既有 `docker exec -i` 模式执行）
> 算法逐字源：`.trae/specs/phase14_06_design_profile_retirement_data_migration/spec.md` §"0011 迁移文件必须按三段结构 + 两段式算法设计"；字段映射逐字源：`.trae/specs/phase14_03_design_standard_data_model_directory_tree/spec.md`

---

## 1. 迁移前基线导出（验收 DB psco_development）

### 1.1 源 1：governance_canonical_root_file_bindings（全量行）

```sql
SELECT * FROM governance_canonical_root_file_bindings ORDER BY file_name;
```

```
 id | governance_profile_id | file_name | role | required | created_at
----+-----------------------+-----------+------+----------+------------
(0 rows)
```

**行数 = 0。**

### 1.2 源 2：governance_global_asset_bindings（全量行）

```sql
SELECT * FROM governance_global_asset_bindings ORDER BY entry_ref;
```

```
 id | governance_profile_id | name | kind | entry_ref | role | structured_summary | created_at
----+-----------------------+------+------+-----------+------+--------------------+------------
(0 rows)
```

**行数 = 0。**

### 1.3 主表：governance_profiles（repository_id + updated_at，按 updated_at 倒序）

```sql
SELECT repository_id, updated_at FROM governance_profiles ORDER BY updated_at DESC;
```

```
            repository_id             |          updated_at
--------------------------------------+-------------------------------
 be2d7b13-ed93-4bbc-90f9-1317a75597e5 | 2026-08-18 04:12:43.246584+00
(1 row)
```

### 1.4 主表行数

```sql
SELECT count(*) AS governance_profiles_count FROM governance_profiles;
```

```
 governance_profiles_count
---------------------------
                         1
(1 row)
```

### 1.5 基线结论：无画像数据场景（冻结决策路径）

两源表均为 **0 行**（主表 1 行）→ 命中 phase14-06 冻结决策第 4 条与任务边界预设场景：

> 两源表存在但零行（或主表零行）→ 不产生迁移产物（standards 无该固定名行、无 revision、无 binding）——用户后续手工创建；空数据不生成空树 Standard；**第三段 drop 仍执行**。

按任务预授权路径继续：验收 DB 走"无产物 + drop"路径；同时因验收 DB 无源数据无法触发算法路径，另建**一次性验证库**（`psco_phase14_09_migcheck`，验证后已删除）以覆盖全算法分支的样本执行 0011 整文件，证明第二段 DO 块算法端到端正确（见 §3）。

---

## 2. 0011 追加段落地要点

文件：`database/migrations/0011_phase14_standard_entity.sql`（第一段建表未改动；头部"结构预留"注释已更新为"结构状态：三段已补全"）

**第二段：数据迁移（单个 `DO $$ ... $$` 块，逐字对齐 phase14-06 冻结算法）**

1. 块首可重放守卫：`IF to_regclass('governance_canonical_root_file_bindings') IS NULL THEN RETURN; END IF;`（两表已 drop 的重放场景整块真跳过，含临时表创建）
2. 无画像数据守卫：两源表计数均 0，或主表计数 0 → RETURN（不产生任何迁移产物）
3. 绑定目标：`repository_id` 取主表最新 `updated_at` 行（裁决⑧"多画像取最新"落地）
4. 节点物化：`CREATE TEMP TABLE standard_migration_nodes (...) ON COMMIT DROP`，列 `(path, parent_path, name, node_type, role, summary, ref, priority, subtree)`；根节点行（path='/'、parent_path=NULL、name='.'、node_type='directory'、priority=0）
5. 源 1 展开：每行 → parent_path='/'、path='/'+file_name、node_type='file'、role=原值、summary=''、ref='/'+file_name、priority=1（`required` 显式退役，不入树）
6. 源 2 展开（entry_ref 三态，priority=2，summary 合成 `COALESCE('[kind] ' || NULLIF(structured_summary,''), '[kind]')`）：
   - `https://` 前缀 → 根下 file 节点（ref=原值）
   - 裸文件名（不含 /）→ 根下 file 节点（name=源 name，ref='/'+entry_ref）
   - 含路径 → 去前导 `/` 按 `/` 拆段（`regexp_replace` + `string_to_array` + `generate_series` lateral 展开）：中间段逐层 directory 节点（role/summary/ref 空），末段 file 节点（**name=最后路径段**，OBS-7 消歧；ref='/'+规范化全路径）
7. 同名冲突合并（聚合，`GROUP BY path`）：
   - role/summary：`COALESCE(MAX(CASE WHEN priority=2 THEN 列 END), MAX(CASE WHEN priority=1 THEN 列 END), '')`——源 2 非空值优先、源 2 空保留源 1
   - name/node_type/ref：`COALESCE(MAX(CASE WHEN priority=1 THEN 列 END), MAX(CASE WHEN priority=2 THEN 列 END), ...)`——保留先建节点（源 1）的位置与 ref（phase14-03 冻结合并口径），同时覆盖中间 directory 与 file 同名同 path 场景
8. subtree 初值：`jsonb_strip_nulls(jsonb_build_object('name',…, 'node_type',…, 'role', NULLIF(role,''), 'summary', NULLIF(summary,''), 'ref', NULLIF(ref,''), 'children','[]'::jsonb))`——空 role/summary/ref 键省略，对齐 phase14-03 DirectoryTreeNode `omitempty` 序列化语义（规范冻结："空字符串时序列化省略"）
9. 组树：`FOR i IN 1..6 LOOP UPDATE ... SET subtree = jsonb_set(subtree, '{children}', COALESCE((SELECT jsonb_agg(c.subtree ORDER BY c.name) FROM ... WHERE c.parent_path = t.path), '[]'::jsonb)) END LOOP;`——固定 6 轮非递归 UPDATE，每轮恰好上卷一层，子节点按 name 升序
10. 产物写入（幂等）：`INSERT INTO standards (name, description, status, directory_tree) SELECT '默认项目范式（迁移自治理画像）', '由 phase14-09 迁移自动创建：合并项目治理画像的 canonical 根文件与全局资产两清单', 'active', 根subtree WHERE NOT EXISTS (同名行) RETURNING id`；仅当本次新插入时写首条 revision（`迁移自项目治理画像：N 项 canonical 根文件 + M 项全局资产合并（源 repository <repository_id>）`）与 binding（`repository` / 主表最新行 repository_id / `adopts`，`ON CONFLICT DO NOTHING`）

**第三段：drop 两表**

```sql
DROP TABLE IF EXISTS governance_canonical_root_file_bindings;
DROP TABLE IF EXISTS governance_global_asset_bindings;
```

---

## 3. 一次性算法验证库执行（psco_phase14_09_migcheck，已删除）

### 3.1 样本设计（源 1 = 5 行、源 2 = 8 行，覆盖全部算法分支）

- 源 1（5 行，phase13-11 形态根级文件）：`AGENTS.md/entry`、`README.md/entry`、`plan.md/plan`、`project_rules.md/rules`、`TECH_STACK_BASELINE.md/baseline`
- 源 2（8 行）覆盖：与源 1 同名冲突 ×3（AGENTS.md / project_rules.md / TECH_STACK_BASELINE.md，其中 TECH 行 `structured_summary=NULL` 验证空摘要格式）、无冲突裸文件名 ×2（architecture_map.md；`global_skills`——name≠entry_ref，验证 name=源 name、ref=entry_ref 规范化）、多级路径拆树 ×2（`docs/phase/README.md`——源 name=`docs_phase_readme`≠末段，验证 OBS-7 末段 name=最后路径段；`/docs/phase/phase14_spec.md`——带前导 / + 空摘要）、`https://` 前缀 ×1（ref 原值保留）
- 主表 1 行：repository_id=`be2d7b13-ed93-4bbc-90f9-1317a75597e5`（与验收库真实值一致）

### 3.2 首次执行输出

```
CREATE TABLE / CREATE TABLE / CREATE INDEX / CREATE TABLE / CREATE INDEX / CREATE INDEX
DO
DROP TABLE / DROP TABLE
=== 0011 APPLY EXIT 0 ===
```

### 3.3 产物断言（全部通过）

**固定名 Standard 恰 1 条，固定值逐字一致：**

```
 fixed_name_standard_count
---------------------------
                         1

              name              | status |                                   description
--------------------------------+--------+---------------------------------------------------------------------------------
 默认项目范式（迁移自治理画像） | active | 由 phase14-09 迁移自动创建：合并项目治理画像的 canonical 根文件与全局资产两清单
```

**directory_tree（jsonb_pretty，键序为 jsonb 内部序，无语义影响）：**

```json
{
    "name": ".", "node_type": "directory",
    "children": [
        { "ref": "/AGENTS.md", "name": "AGENTS.md", "role": "agent-context", "summary": "[agents] agent 全局上下文入口", "children": [], "node_type": "file" },
        { "ref": "/README.md", "name": "README.md", "role": "entry", "children": [], "node_type": "file" },
        { "ref": "/TECH_STACK_BASELINE.md", "name": "TECH_STACK_BASELINE.md", "role": "baseline", "summary": "[baseline]", "children": [], "node_type": "file" },
        { "ref": "/architecture_map.md", "name": "architecture_map.md", "role": "summary", "summary": "[summary] 目录结构真相源", "children": [], "node_type": "file" },
        { "name": "docs", "children": [
            { "name": "phase", "children": [
                { "ref": "/docs/phase/README.md", "name": "README.md", "role": "summary", "summary": "[summary] phase workflow 索引", "children": [], "node_type": "file" },
                { "ref": "/docs/phase/phase14_spec.md", "name": "phase14_spec.md", "role": "spec", "summary": "[spec]", "children": [], "node_type": "file" }
            ], "node_type": "directory" }
        ], "node_type": "directory" },
        { "ref": "https://example.com/std.md", "name": "external_std.md", "role": "external", "summary": "[external] 外部规范引用", "children": [], "node_type": "file" },
        { "ref": "/global_skills.md", "name": "global_skills", "role": "skills", "summary": "[skills] 全局技能", "children": [], "node_type": "file" },
        { "ref": "/plan.md", "name": "plan.md", "role": "plan", "children": [], "node_type": "file" },
        { "ref": "/project_rules.md", "name": "project_rules.md", "role": "rules", "summary": "[rules] 项目级正式规则", "children": [], "node_type": "file" }
    ]
}
```

逐节点核对结论：根 name='.' 且 role/summary/ref 键省略（R1 + omitempty）；根 children 按 name 升序；`AGENTS.md` 节点 role/summary 为源 2 非空值覆盖、ref 保留源 1 位置（同名合并三规则同时命中）；`README.md` / `plan.md` 无 summary 键（空省略）；`[baseline]` / `[spec]` 空摘要格式正确；`docs/phase` 两层中间 directory 的 role/summary/ref 键全部省略；`/docs/phase/README.md` 末段 name=`README.md`（OBS-7，源 name=`docs_phase_readme` 未采用）；`https://` ref 原值保留；`global_skills` 节点 name=源 name、ref=`/global_skills.md`（entry_ref 规范化）。

**首条 revision（N/M 计数 + 源 repository id 逐字正确）：**

```
迁移自项目治理画像：5 项 canonical 根文件 + 8 项全局资产合并（源 repository be2d7b13-ed93-4bbc-90f9-1317a75597e5）
```

**binding：**

```
 target_type |              target_id               |  role
-------------+--------------------------------------+--------
 repository  | be2d7b13-ed93-4bbc-90f9-1317a75597e5 | adopts
```

**树结构近似校验（递归 CTE 机械断言）：**

```
 total_nodes | max_depth | file_nodes | file_missing_ref | root_nodes
-------------+-----------+------------+------------------+------------
          13 |         4 |         10 |                0 |         1
```

单根 ✓；深度 4 ≤ 6（R5）✓；全部 10 个 file 节点均有 ref（R7/R8）✓；`octet_length(tree)=1603` ≤ 64KB（R6）✓。

**两源表 drop：**

```
 src1_after | src2_after
------------+------------
            |
(两列均为 NULL)
```

### 3.4 重放（同一文件第二次 psql -f）

```
NOTICE: relation "standards" already exists, skipping        ×3 表 + ×3 索引
DO                                                            ← DO 块守卫真跳过，零报错
NOTICE: table "governance_canonical_root_file_bindings" does not exist, skipping
NOTICE: table "governance_global_asset_bindings" does not exist, skipping
=== REPLAY EXIT 0 ===

 standards_rows | revision_rows | binding_rows | fixed_name_rows
----------------+---------------+--------------+-----------------
              1 |             1 |              1 |               1
```

重放零报错；产物无重复（Standard/revision/binding 均不新增）。

---

## 4. 验收 DB（psco_development）执行与断言

### 4.1 首次执行输出（无画像数据路径）

```
NOTICE: relation "standards" already exists, skipping          （第一段 IF NOT EXISTS 空过，×3 表 ×3 索引）
DO                                                              ← 无画像数据守卫 RETURN（两源表 0 行），零报错
DROP TABLE                                                      ← 第三段真 drop（两表此时存在）
DROP TABLE
=== ACCEPTANCE DB APPLY EXIT 0 ===
```

### 4.2 断言结果（无画像数据场景冻结预期形态）

| 断言 | 结果 | 冻结预期 | 判定 |
|---|---|---|---|
| `SELECT count(*) FROM standards WHERE name='默认项目范式（迁移自治理画像）'` | 0 | 无画像数据不产生产物 | ✓ |
| standards / standard_revisions / standard_bindings 总行数 | 0 / 0 / 0 | 无产物 | ✓ |
| `SELECT to_regclass('governance_canonical_root_file_bindings'), to_regclass('governance_global_asset_bindings')` | NULL / NULL | 两表已 drop（T2） | ✓ |
| `SELECT count(*) FROM governance_profiles` | 1（迁移前同为 1） | 主表保留、数据未动 | ✓ |
| `SELECT repository_id, updated_at FROM governance_profiles ORDER BY updated_at DESC` | `be2d7b13-… / 2026-08-18 04:12:43.246584+00`（与基线一致） | 0011 不触碰主表 | ✓ |

树对照 / revision / binding 断言在验收 DB 为"无产物"形态（源数据为空集，无从迁移）；算法正确性由 §3 一次性验证库全量覆盖。

### 4.3 重放（验收 DB，同一文件第二次 psql -f）

```
NOTICE: relation "..." already exists, skipping                 （建表段全部空过）
DO                                                              ← 块首 to_regclass 守卫真跳过，零报错
NOTICE: table "governance_canonical_root_file_bindings" does not exist, skipping
NOTICE: table "governance_global_asset_bindings" does not exist, skipping
=== ACCEPTANCE DB REPLAY EXIT 0 ===

 standards_rows | revision_rows | binding_rows | profile_rows
----------------+---------------+--------------+-------------
              0 |             0 |            0 |           1

 src1_final | src2_final
------------+------------
   NULL     |   NULL
```

重放零报错、状态零变化。

---

## 5. 零丢失对照表（5 字段映射集）

> 验收 DB 两源表为空集（0+0 行），空集对照成立（无源行需要承接，无丢失可能）；本表按一次性验证库样本（源 1 = 5 行、源 2 = 8 行）逐行落证，映射规则与 phase14-03 冻结字段级映射逐字一致。

### 5.1 源 1（file_name→name / role→role / 派生 ref='/'+file_name）

| # | 源 1 file_name | 源 1 role | → 树节点 name | → 树节点 role | → 树节点 ref | 打勾 |
|---|---|---|---|---|---|---|
| 1 | AGENTS.md | entry | AGENTS.md | agent-context（源 2 非空增强） | /AGENTS.md（源 1 位置保留） | ✓ |
| 2 | README.md | entry | README.md | entry | /README.md | ✓ |
| 3 | plan.md | plan | plan.md | plan | /plan.md | ✓ |
| 4 | project_rules.md | rules | project_rules.md | rules（源 2 同值） | /project_rules.md | ✓ |
| 5 | TECH_STACK_BASELINE.md | baseline | TECH_STACK_BASELINE.md | baseline | /TECH_STACK_BASELINE.md | ✓ |

`required` 显式退役（phase14-03 冻结：结构性吸收，不入零丢失验收范围）。

### 5.2 源 2（name→name / kind+structured_summary→summary / entry_ref→ref / role→role）

| # | 源 2 name | kind | entry_ref | role | structured_summary | → 树位置 | → name | → role | → summary | → ref | 打勾 |
|---|---|---|---|---|---|---|---|---|---|---|---|
| 1 | AGENTS.md | agents | AGENTS.md | agent-context | agent 全局上下文入口 | /AGENTS.md（合并入源 1 节点） | AGENTS.md | agent-context | [agents] agent 全局上下文入口 | /AGENTS.md | ✓ |
| 2 | project_rules.md | rules | project_rules.md | rules | 项目级正式规则 | /project_rules.md | project_rules.md | rules | [rules] 项目级正式规则 | /project_rules.md | ✓ |
| 3 | TECH_STACK_BASELINE.md | baseline | TECH_STACK_BASELINE.md | baseline | NULL | /TECH_STACK_BASELINE.md | TECH_STACK_BASELINE.md | baseline | [baseline]（空摘要格式） | /TECH_STACK_BASELINE.md | ✓ |
| 4 | architecture_map.md | summary | architecture_map.md | summary | 目录结构真相源 | /architecture_map.md | architecture_map.md | summary | [summary] 目录结构真相源 | /architecture_map.md | ✓ |
| 5 | global_skills | skills | global_skills.md | skills | 全局技能 | /global_skills.md（裸文件名） | global_skills（源 name） | skills | [skills] 全局技能 | /global_skills.md | ✓ |
| 6 | docs_phase_readme | summary | docs/phase/README.md | summary | phase workflow 索引 | /docs/phase/README.md（多级拆树） | README.md（末段，OBS-7） | summary | [summary] phase workflow 索引 | /docs/phase/README.md | ✓ |
| 7 | phase14_spec | spec | /docs/phase/phase14_spec.md | spec | NULL | /docs/phase/phase14_spec.md（前导 / 规范化） | phase14_spec.md（末段） | spec | [spec] | /docs/phase/phase14_spec.md | ✓ |
| 8 | external_std.md | external | https://example.com/std.md | external | 外部规范引用 | 根下（URL 无树内路径） | external_std.md | external | [external] 外部规范引用 | https://example.com/std.md（原值） | ✓ |

时间戳列（created_at / updated_at）按 phase14-02 冻结为显式不迁移项，不在验收范围。

---

## 6. 结论

1. `0011_phase14_standard_entity.sql` 已完成三段结构补全（第一段未改动），头部注释同步更新；SQL 固定值中文串（`默认项目范式（迁移自治理画像）`、description、`迁移自项目治理画像：…（源 repository …）`）与冻结值逐字一致（含全角括号）。
2. 算法端到端正确性已在一次性验证库证明：三态展开、同名冲突合并（源 2 非空 role/summary 覆盖 + 源 1 位置 ref 保留）、多级拆树（OBS-7 末段命名）、空摘要 `[kind]` 格式、omitempty 键省略、children name 升序、固定 6 轮组树、幂等产物写入与重放全部通过。
3. 验收 DB `psco_development` 为**无画像数据场景**（两源表 0 行、主表 1 行）：按冻结决策不产生迁移产物，两张 bindings 表已 drop（T2 断言 NULL），`governance_profiles` 主表 1 行数据与结构未动，整文件重放零报错零变化。
4. 后续提示：`Standards` 固定名产物在验收 DB 中不存在属冻结预期（空数据不生成空树）；若需在验收 DB 出现该产物，需先有画像两表数据再执行迁移（当前已 drop，数据不可再生，需经 Standard 维护路径手工创建）。
