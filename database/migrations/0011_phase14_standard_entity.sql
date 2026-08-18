-- 0011_phase14_standard_entity.sql — Standard 全局规范实体三表建表迁移（phase14-07）
--
-- 文档定位：phase14 Standard Entity Foundation 存储层第一段建表。
-- 上游规格：
--   - phase14-03 Standard 数据模型与目录树设计 Spec（三表 DDL 级草案，逐字源）
--   - phase14-06 画像退役与数据迁移设计 Spec（0011 三段结构设计）
--
-- 结构预留：phase14-09 将按 phase14-06 冻结设计在本文件追加
--   第二段（存量画像两表数据迁移 DO 块）与第三段（drop 两张画像 bindings 表）。
--   因此本文件全部语句使用 IF NOT EXISTS 幂等形态，保证追加段后整文件可安全重放。
--
-- 本文件不触碰 phase13 画像三表（governance_profiles 主表与两张 bindings 表）。

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
