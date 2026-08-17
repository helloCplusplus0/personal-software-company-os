-- 0010_phase13_governance_profile.sql
-- Phase13 Project Governance Profile Foundation 最小数据承接位
-- 直接承接表：governance_profiles / governance_canonical_root_file_bindings /
--             governance_global_asset_bindings（新增三张）
-- 上游规格：phase13-04 项目治理画像数据模型与字段设计 Spec（9 类核心字段与分类矩阵）
--           phase13-05 后端合同、存储与读写边界设计 Spec（存储三层与 8 项资产矩阵）
--           phase13-08 落实项目治理画像后端主线 Spec（唯一后端实现主线）
--
-- 设计约束：
--   - repository_id 是治理画像主记录唯一业务锚点（UNIQUE + 外键）
--   - bindings 层通过 governance_profile_id 形成明确外键关联（ON DELETE CASCADE）
--   - read-only 字段（track_type / current_phase_*）持久化在主记录中，
--     但只允许来自根级正式上游冻结结果的受控初始化，不由维护写路径改写
--   - markdown 正文不进入任何 canonical 存储列；structured_summary 只承接结构化摘要
--   - backend / database / frontend / proto 顶层目录矩阵不进入本迁移
--     （只作为项目范式 v1 的只读基线输入保留）

-- ============================================================================
-- governance_profiles：治理画像主记录（按 repository_id 唯一锚定）
-- ============================================================================
CREATE TABLE governance_profiles (
    id                     UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    -- 唯一项目锚点
    repository_id          UUID NOT NULL UNIQUE REFERENCES repositories(id) ON DELETE CASCADE,
    -- 治理画像正式版本（第一版固定 project_governance_profile_v1，由服务端写入）
    project_profile_version TEXT NOT NULL,
    -- 技术路线（read-only，受控枚举）
    track_type             TEXT NOT NULL CHECK (track_type IN ('product', 'durable_system')),
    -- 模板来源（optional，允许为空）
    template_source        TEXT NULL,
    -- docs workflow 结构布局（非空受控字符串，当前项目为 phase/fix/audit/review）
    docs_workflow_layout   TEXT NOT NULL,
    -- 当前阶段名称（read-only）
    current_phase_name     TEXT NOT NULL,
    -- 当前阶段正式入口引用（read-only）
    current_phase_ref      TEXT NOT NULL,
    -- 当前阶段状态（read-only，受控枚举）
    current_phase_status   TEXT NOT NULL CHECK (current_phase_status IN ('planned', 'in_progress', 'completed', 'blocked')),
    created_at             TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at             TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- ============================================================================
-- governance_canonical_root_file_bindings：canonical 根级文件绑定
-- ============================================================================
-- 最小字段集合（phase13-04 嵌套矩阵）：file_name / role / required
CREATE TABLE governance_canonical_root_file_bindings (
    id                     UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    governance_profile_id  UUID NOT NULL REFERENCES governance_profiles(id) ON DELETE CASCADE,
    -- 根级 canonical 文件名（如 project_rules.md）
    file_name              TEXT NOT NULL,
    -- 该文件在项目治理中的正式角色
    role                   TEXT NOT NULL,
    -- 是否为当前项目范式必需文件
    required               BOOLEAN NOT NULL DEFAULT FALSE,
    created_at             TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE (governance_profile_id, file_name)
);

CREATE INDEX idx_governance_root_file_bindings_profile_id
    ON governance_canonical_root_file_bindings (governance_profile_id);

-- ============================================================================
-- governance_global_asset_bindings：全局规范资产绑定
-- ============================================================================
-- 最小字段集合（phase13-05 矩阵）：name / kind / entry_ref / role / structured_summary
-- structured_summary 前 5 项摘要型资产必填（由 service 层校验），
-- README.md / global_skills.md / project_skills.md 第一版允许为空。
CREATE TABLE governance_global_asset_bindings (
    id                     UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    governance_profile_id  UUID NOT NULL REFERENCES governance_profiles(id) ON DELETE CASCADE,
    -- 资产名（必须属于 8 项冻结矩阵，由 service 层校验）
    name                   TEXT NOT NULL,
    -- 资产分类
    kind                   TEXT NOT NULL,
    -- 资产正式入口引用（定位与回源入口，不是正文）
    entry_ref              TEXT NOT NULL,
    -- 该资产在项目治理中的正式角色
    role                   TEXT NOT NULL,
    -- 结构化摘要（可空列；必填性由 service 层按 8 项矩阵校验）
    structured_summary     TEXT NULL,
    created_at             TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE (governance_profile_id, name)
);

CREATE INDEX idx_governance_global_asset_bindings_profile_id
    ON governance_global_asset_bindings (governance_profile_id);
