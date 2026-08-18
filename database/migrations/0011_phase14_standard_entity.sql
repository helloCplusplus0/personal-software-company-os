-- 0011_phase14_standard_entity.sql — Standard 全局规范实体三表建表迁移（phase14-07）
--
-- 文档定位：phase14 Standard Entity Foundation 存储层第一段建表。
-- 上游规格：
--   - phase14-03 Standard 数据模型与目录树设计 Spec（三表 DDL 级草案，逐字源）
--   - phase14-06 画像退役与数据迁移设计 Spec（0011 三段结构设计）
--
-- 结构状态：phase14-09 已按 phase14-06 冻结设计补全三段结构——
--   第一段：Standard 三表建表（phase14-07 落地，IF NOT EXISTS 幂等）
--   第二段：存量画像两表 → 单条全局 Standard 数据迁移（单个 DO 块，含可重放守卫）
--   第三段：drop 两张画像 bindings 表（governance_profiles 主表保留——T2）
--   整文件可安全重放：建表段 IF NOT EXISTS 空过、DO 块守卫真跳过、DROP IF EXISTS 空过。
--
-- 本文件读取 governance_profiles 主表（repository_id / updated_at，不修改），
-- 读取并最终 drop 两张画像 bindings 表；不触碰主表数据。

-- standards：全局规范主表（全局作用域，无 repository 锚点——裁决⑤）
CREATE TABLE IF NOT EXISTS standards (
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
CREATE TABLE IF NOT EXISTS standard_revisions (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    standard_id     UUID NOT NULL REFERENCES standards(id) ON DELETE CASCADE,
    -- 人工一句话说明本次调整（必填）
    change_summary  TEXT NOT NULL,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
CREATE INDEX IF NOT EXISTS idx_standard_revisions_standard_id ON standard_revisions (standard_id);

-- standard_bindings：多态绑定关系表（CON-02 不设限；target_id 无 DB 外键为多态必然，存在性由应用层校验）
CREATE TABLE IF NOT EXISTS standard_bindings (
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
CREATE INDEX IF NOT EXISTS idx_standard_bindings_target ON standard_bindings (target_type, target_id);
CREATE INDEX IF NOT EXISTS idx_standard_bindings_standard_id ON standard_bindings (standard_id);

-- ============================================================================
-- 第二段：数据迁移（phase14-09；phase14-06 冻结算法的 SQL 落地）
-- ============================================================================
-- 单个 DO 块承接全部数据迁移（块首 to_regclass 守卫保证两表已 drop 的重放场景真跳过）。
-- 算法两段式：节点行物化（源 1 / 源 2 展开 + 同名冲突聚合合并）→ 固定 6 轮非递归组树
-- （弃 WITH RECURSIVE：递归项内禁聚合）→ 幂等产物写入。
-- 字段映射逐字对齐 phase14-03 迁移映射节（entry_ref 三态 / 冲突合并 / summary 合成格式）。
DO $$
DECLARE
    n_root_files    integer;
    n_global_assets integer;
    n_profiles      integer;
    v_repository_id uuid;
    v_standard_id   uuid;
    v_root_subtree  jsonb;
BEGIN
    -- 可重放守卫：两张源表已 drop 时整块真跳过（不执行任何块内语句，含临时表创建）
    IF to_regclass('governance_canonical_root_file_bindings') IS NULL THEN
        RETURN;
    END IF;

    -- 无画像数据守卫：两源表零行或主表零行 → 不产生任何迁移产物
    -- （不生成空树 Standard；用户后续手工创建），第三段 drop 仍在本块外执行
    SELECT count(*) INTO n_root_files    FROM governance_canonical_root_file_bindings;
    SELECT count(*) INTO n_global_assets FROM governance_global_asset_bindings;
    SELECT count(*) INTO n_profiles      FROM governance_profiles;
    IF (n_root_files = 0 AND n_global_assets = 0) OR n_profiles = 0 THEN
        RETURN;
    END IF;

    -- 绑定目标：主表最新 updated_at 行的 repository_id（裁决⑧"多画像取最新"落地）
    SELECT repository_id
      INTO v_repository_id
      FROM governance_profiles
     ORDER BY updated_at DESC
     LIMIT 1;

    -- ===== 步骤 1a：节点行物化（临时表生命周期 = 本 DO 块事务，不跨语句泄漏）=====
    CREATE TEMP TABLE standard_migration_nodes (
        path        text,
        parent_path text,
        name        text,
        node_type   text,    -- 'directory' | 'file'
        role        text,
        summary     text,
        ref         text,
        priority    integer, -- 0=根节点 1=源1(canonical 根文件) 2=源2(全局资产)
        subtree     jsonb
    ) ON COMMIT DROP;

    -- 根节点：单根结构，name='.'（仓库根语义）；根 role/ref 必空（R1 / R8）
    INSERT INTO standard_migration_nodes (path, parent_path, name, node_type, role, summary, ref, priority)
    VALUES ('/', NULL, '.', 'directory', '', '', '', 0);

    -- 源 1 展开（governance_canonical_root_file_bindings）：每行 → 根下 file 节点，
    -- ref 规范化为 '/' + file_name；required 显式退役（phase14-03 冻结，不入树）
    INSERT INTO standard_migration_nodes (path, parent_path, name, node_type, role, summary, ref, priority)
    SELECT '/' || b.file_name,
           '/',
           b.file_name,
           'file',
           b.role,
           '',
           '/' || b.file_name,
           1
      FROM governance_canonical_root_file_bindings AS b;

    -- 源 2 展开（governance_global_asset_bindings）：entry_ref 三态
    --   a) https:// 前缀 → 根下 file 节点（ref 原值保留，外部资源无树内路径）
    INSERT INTO standard_migration_nodes (path, parent_path, name, node_type, role, summary, ref, priority)
    SELECT '/' || b.name,
           '/',
           b.name,
           'file',
           b.role,
           COALESCE('[' || b.kind || '] ' || NULLIF(b.structured_summary, ''), '[' || b.kind || ']'),
           b.entry_ref,
           2
      FROM governance_global_asset_bindings AS b
     WHERE b.entry_ref LIKE 'https://%';

    --   b) 裸文件名（不含 /）→ 根下 file 节点（ref = '/' + entry_ref）
    INSERT INTO standard_migration_nodes (path, parent_path, name, node_type, role, summary, ref, priority)
    SELECT '/' || b.entry_ref,
           '/',
           b.name,
           'file',
           b.role,
           COALESCE('[' || b.kind || '] ' || NULLIF(b.structured_summary, ''), '[' || b.kind || ']'),
           '/' || b.entry_ref,
           2
      FROM governance_global_asset_bindings AS b
     WHERE b.entry_ref NOT LIKE 'https://%'
       AND position('/' IN b.entry_ref) = 0;

    --   c) 含路径 → 去前导 / 按 / 拆段：中间段逐层物化 directory 节点（role/summary 空），
    --      末段为 file 节点（name = 最后路径段，OBS-7 消歧；ref = 规范化全路径）
    WITH asset_paths AS (
        SELECT b.role  AS src_role,
               COALESCE('[' || b.kind || '] ' || NULLIF(b.structured_summary, ''), '[' || b.kind || ']') AS src_summary,
               regexp_replace(b.entry_ref, '^/+', '') AS norm
          FROM governance_global_asset_bindings AS b
         WHERE b.entry_ref NOT LIKE 'https://%'
           AND position('/' IN b.entry_ref) > 0
    ),
    asset_segs AS (
        SELECT ap.src_role,
               ap.src_summary,
               ap.norm,
               string_to_array(ap.norm, '/') AS segs
          FROM asset_paths AS ap
    )
    INSERT INTO standard_migration_nodes (path, parent_path, name, node_type, role, summary, ref, priority)
    SELECT '/' || array_to_string(s.segs[1:g.i], '/'),
           CASE WHEN g.i = 1
                THEN '/'
                ELSE '/' || array_to_string(s.segs[1:g.i - 1], '/')
           END,
           s.segs[g.i],
           CASE WHEN g.i = cardinality(s.segs) THEN 'file'      ELSE 'directory' END,
           CASE WHEN g.i = cardinality(s.segs) THEN s.src_role  ELSE '' END,
           CASE WHEN g.i = cardinality(s.segs) THEN s.src_summary ELSE '' END,
           CASE WHEN g.i = cardinality(s.segs) THEN '/' || s.norm ELSE '' END,
           2
      FROM asset_segs AS s
      CROSS JOIN LATERAL generate_series(1, cardinality(s.segs)) AS g(i);

    -- ===== 步骤 1b：同名冲突合并（聚合完成，非过程式；GROUP BY path）=====
    -- role/summary：源 2 非空值优先覆盖、源 2 空保留源 1（phase14-03 冻结合并规则）；
    -- name/node_type/ref：保留先建节点的位置与 ref（源 1 优先，先建=源 1），
    -- 兼顾中间 directory 节点与 file 节点同名同 path 的合并口径；根节点（priority=0）
    -- 各 CASE 均不命中，稳定回落 MAX 聚合值。
    CREATE TEMP TABLE standard_migration_merged ON COMMIT DROP AS
    SELECT path,
           MAX(parent_path) AS parent_path,
           COALESCE(MAX(CASE WHEN priority = 1 THEN name END),
                    MAX(CASE WHEN priority = 2 THEN name END),
                    MAX(name)) AS name,
           COALESCE(MAX(CASE WHEN priority = 1 THEN node_type END),
                    MAX(CASE WHEN priority = 2 THEN node_type END),
                    MAX(node_type)) AS node_type,
           COALESCE(MAX(CASE WHEN priority = 2 THEN role END),
                    MAX(CASE WHEN priority = 1 THEN role END),
                    '') AS role,
           COALESCE(MAX(CASE WHEN priority = 2 THEN summary END),
                    MAX(CASE WHEN priority = 1 THEN summary END),
                    '') AS summary,
           COALESCE(MAX(CASE WHEN priority = 1 THEN ref END),
                    MAX(CASE WHEN priority = 2 THEN ref END),
                    '') AS ref,
           MIN(priority) AS priority
      FROM standard_migration_nodes
     GROUP BY path;

    DELETE FROM standard_migration_nodes;
    INSERT INTO standard_migration_nodes (path, parent_path, name, node_type, role, summary, ref, priority, subtree)
    SELECT path, parent_path, name, node_type, role, summary, ref, priority, NULL
      FROM standard_migration_merged;
    DROP TABLE standard_migration_merged;

    -- ===== 步骤 1c：subtree 初值（自身节点对象 + children=[]）=====
    -- 空 role/summary/ref 以 NULLIF 置 NULL 后经 jsonb_strip_nulls 省略键，
    -- 对齐 phase14-03 DirectoryTreeNode 的 omitempty 序列化语义
    UPDATE standard_migration_nodes
       SET subtree = jsonb_strip_nulls(jsonb_build_object(
               'name',     name,
               'node_type', node_type,
               'role',     NULLIF(role, ''),
               'summary',  NULLIF(summary, ''),
               'ref',      NULLIF(ref, ''),
               'children', '[]'::jsonb
           ));

    -- ===== 步骤 2：自底向上固定 6 轮组树（R5 冻结深度 ≤6；弃 WITH RECURSIVE）=====
    -- 每轮一条非递归 UPDATE，相关子查询按 UPDATE 快照语义每轮恰好上卷一层，
    -- 重复轮次幂等；子节点按 name 升序保证产物字节稳定
    FOR i IN 1..6 LOOP
        UPDATE standard_migration_nodes AS t
           SET subtree = jsonb_set(
                   t.subtree,
                   '{children}',
                   COALESCE((
                       SELECT jsonb_agg(c.subtree ORDER BY c.name)
                         FROM standard_migration_nodes AS c
                        WHERE c.parent_path = t.path
                   ), '[]'::jsonb)
               );
    END LOOP;

    -- ===== 步骤 3：产物写入（幂等守卫）=====
    SELECT subtree
      INTO v_root_subtree
      FROM standard_migration_nodes
     WHERE path = '/';

    -- 整树单行原子写入：固定名 NOT EXISTS 幂等守卫
    INSERT INTO standards (name, description, status, directory_tree)
    SELECT '默认项目范式（迁移自治理画像）',
           '由 phase14-09 迁移自动创建：合并项目治理画像的 canonical 根文件与全局资产两清单',
           'active',
           v_root_subtree
     WHERE NOT EXISTS (
         SELECT 1 FROM standards WHERE name = '默认项目范式（迁移自治理画像）'
     )
    RETURNING id INTO v_standard_id;

    -- 首条 revision（合并来源留痕，裁决⑧）与 adopts 绑定仅当 Standard 为本次新插入时写入
    IF v_standard_id IS NOT NULL THEN
        INSERT INTO standard_revisions (standard_id, change_summary)
        VALUES (v_standard_id,
                '迁移自项目治理画像：' || n_root_files::text || ' 项 canonical 根文件 + '
                || n_global_assets::text || ' 项全局资产合并（源 repository '
                || v_repository_id::text || '）');

        INSERT INTO standard_bindings (standard_id, target_type, target_id, role)
        VALUES (v_standard_id, 'repository', v_repository_id, 'adopts')
        ON CONFLICT DO NOTHING;
    END IF;
END $$;

-- ============================================================================
-- 第三段：drop 两张画像 bindings 表（phase14-09；governance_profiles 主表保留——T2）
-- ============================================================================
DROP TABLE IF EXISTS governance_canonical_root_file_bindings;
DROP TABLE IF EXISTS governance_global_asset_bindings;
