-- 0003_module_registry_bindings.sql
-- Module Registry 关联写入表
-- 直接承接表：product_modules, module_repositories
-- 上游规格：phase02-09 module_registry_spec_v0.1.md §4.1, §5.1
-- §4.1 BindModuleToProduct / MapModuleToRepository 归属 Module Registry 后端模块

-- product_modules：模块与产品的绑定关系
-- §4.1 BindModuleToProduct 写入载体
CREATE TABLE product_modules (
    id         UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    module_id  UUID NOT NULL REFERENCES modules(id) ON DELETE CASCADE,
    product_id UUID NOT NULL REFERENCES products(id) ON DELETE RESTRICT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE (module_id, product_id)
);

-- module_repositories：模块与仓库的映射关系
-- §4.1 MapModuleToRepository 写入载体
CREATE TABLE module_repositories (
    id            UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    module_id     UUID NOT NULL REFERENCES modules(id) ON DELETE CASCADE,
    repository_id UUID NOT NULL REFERENCES repositories(id) ON DELETE RESTRICT,
    created_at    TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE (module_id, repository_id)
);

CREATE INDEX idx_product_modules_module_id ON product_modules (module_id);
CREATE INDEX idx_module_repositories_module_id ON module_repositories (module_id);
