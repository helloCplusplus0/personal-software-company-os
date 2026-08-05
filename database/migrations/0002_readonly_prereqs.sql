-- 0002_readonly_prereqs.sql
-- 当前阶段只读前提表
-- 上游规格：phase02-09 module_registry_spec_v0.1.md §5.2, §5.3
-- 约束：这些表在 phase02 中只服务读取，不扩写为新的写入主线

-- products：Product 候选读取前提（只读）
-- §6.2 ProductBindingCandidateRead 的数据来源
CREATE TABLE products (
    id         UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name       TEXT NOT NULL UNIQUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- repositories：Repository 候选读取前提（只读）
-- §6.2 RepositoryBindingCandidateRead 的数据来源
CREATE TABLE repositories (
    id         UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name       TEXT NOT NULL UNIQUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- decisions：Decision 入口只读前提
-- §6.3 Decision 作为 ModuleDetailRead 的内嵌附属读取承接，不设独立读接口组
CREATE TABLE decisions (
    id         UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title      TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- decision_links：Decision 与 Module 的关联（只读前提）
-- §6.3 不得为 Decision 单独设立独立 Read 接口、独立 handler 或独立 service 文件
-- 此表数据由种子数据初始化，phase02 不实现 RecordDecision / LinkDecisionToTarget 写入
CREATE TABLE decision_links (
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    decision_id UUID NOT NULL REFERENCES decisions(id) ON DELETE CASCADE,
    module_id   UUID NOT NULL REFERENCES modules(id) ON DELETE CASCADE,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE (decision_id, module_id)
);

CREATE INDEX idx_decision_links_module_id ON decision_links (module_id);
