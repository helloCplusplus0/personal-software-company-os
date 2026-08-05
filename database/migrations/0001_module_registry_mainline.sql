-- 0001_module_registry_mainline.sql
-- Module Registry 主线核心表
-- 直接承接表：modules, module_releases
-- 上游规格：phase02-09 module_registry_spec_v0.1.md §5.1, §5.4, §5.5, §5.6

-- 迁移追踪表（由迁移运行器管理，幂等创建）
-- 注：migrate.go 也会创建此表，因此使用 IF NOT EXISTS 防止冲突；
--     业务表不使用 IF NOT EXISTS，因 schema_migrations 跟踪机制保证每个迁移只执行一次。
CREATE TABLE IF NOT EXISTS schema_migrations (
    version   TEXT PRIMARY KEY,
    applied_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- modules：模块登记核心实体
-- §5.4 对象最小字段：id / name / description / status / created_at
-- §5.5 状态集合冻结为 active / archived
-- §5.6 名称唯一约束
CREATE TABLE modules (
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name        TEXT NOT NULL UNIQUE,
    description TEXT NOT NULL,
    status      TEXT NOT NULL CHECK (status IN ('active', 'archived')),
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- module_releases：版本登记，依附 modules
-- §5.4 对象最小字段：id / module_id / version / status / released_at
-- 同一模块下版本号唯一
CREATE TABLE module_releases (
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    module_id   UUID NOT NULL REFERENCES modules(id) ON DELETE CASCADE,
    version     TEXT NOT NULL,
    status      TEXT NOT NULL CHECK (status IN ('active', 'archived')),
    released_at TIMESTAMPTZ NOT NULL,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE (module_id, version)
);

-- 列表读取性能：按状态过滤 + 创建时间倒序
CREATE INDEX idx_modules_status_created_at ON modules (status, created_at DESC);
-- 详情读取性能：按 module_id 取版本列表
CREATE INDEX idx_module_releases_module_id ON module_releases (module_id, released_at DESC);
