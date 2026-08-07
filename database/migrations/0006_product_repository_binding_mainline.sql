-- 0006_product_repository_binding_mainline.sql
-- Product / Repository / Binding 主线核心表
-- 直接承接表：products (升级), repositories (升级), product_repositories (新增)
-- 上游规格：phase04-10 product_repository_binding_spec_v0.1.md §数据模型
--           phase04-12 spec §"0006 migration 必须将 products / repositories 升级为 phase04 正式主线"
--
-- 升级策略：
--   - products 原位新增 description / status（不创建 products_v2 或影子表）
--   - repositories 原位新增 url / provider / status（不创建 repositories_v2 或影子表）
--   - 新增 product_repositories 关系表
--   - 历史数据通过 ALTER TABLE ADD COLUMN ... DEFAULT ... 原位回填
--   - 回填后 DROP DEFAULT，确保后续 INSERT 必须显式提供字段值
--
-- 兼容性：
--   - 历史 product_modules / module_repositories 数据保持可读
--   - 历史 products / repositories 记录获得默认回填值，不阻断既有读取

-- ============================================================================
-- products 原位升级：新增 description / status
-- ============================================================================

-- 新增 description（回填历史记录）
ALTER TABLE products
  ADD COLUMN description TEXT NOT NULL DEFAULT '（历史产品，phase04 升级前无描述）';

-- 新增 status（回填历史记录为 active，附 CHECK 约束）
ALTER TABLE products
  ADD COLUMN status TEXT NOT NULL DEFAULT 'active' CHECK (status IN ('active', 'archived'));

-- 移除 DEFAULT，后续 INSERT 必须显式提供 description / status
ALTER TABLE products
  ALTER COLUMN description DROP DEFAULT,
  ALTER COLUMN status DROP DEFAULT;

-- 列表读取索引：按状态过滤 + 创建时间倒序
CREATE INDEX idx_products_status_created_at ON products (status, created_at DESC);

-- ============================================================================
-- repositories 原位升级：新增 url / provider / status
-- ============================================================================

-- 新增 url（回填历史记录）
ALTER TABLE repositories
  ADD COLUMN url TEXT NOT NULL DEFAULT 'https://example.com/legacy';

-- 新增 provider（回填历史记录）
ALTER TABLE repositories
  ADD COLUMN provider TEXT NOT NULL DEFAULT 'legacy';

-- 新增 status（回填历史记录为 active，附 CHECK 约束）
ALTER TABLE repositories
  ADD COLUMN status TEXT NOT NULL DEFAULT 'active' CHECK (status IN ('active', 'archived'));

-- 移除 DEFAULT，后续 INSERT 必须显式提供 url / provider / status
ALTER TABLE repositories
  ALTER COLUMN url DROP DEFAULT,
  ALTER COLUMN provider DROP DEFAULT,
  ALTER COLUMN status DROP DEFAULT;

-- 列表读取索引：按状态过滤 + 创建时间倒序
CREATE INDEX idx_repositories_status_created_at ON repositories (status, created_at DESC);

-- ============================================================================
-- product_repositories：Product 与 Repository 的绑定关系（新增）
-- ============================================================================
-- phase04-10 §数据模型：BindRepositoryToProduct 写入载体
-- 当前阶段不引入额外绑定属性字段

CREATE TABLE product_repositories (
    id            UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    product_id    UUID NOT NULL REFERENCES products(id) ON DELETE RESTRICT,
    repository_id UUID NOT NULL REFERENCES repositories(id) ON DELETE RESTRICT,
    created_at    TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE (product_id, repository_id)
);

-- 按产品读取已绑定仓库列表
CREATE INDEX idx_product_repositories_product_id ON product_repositories (product_id);
-- 按仓库读取已绑定产品列表
CREATE INDEX idx_product_repositories_repository_id ON product_repositories (repository_id);
