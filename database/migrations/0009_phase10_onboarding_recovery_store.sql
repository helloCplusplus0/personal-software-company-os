-- 0009_phase10_onboarding_recovery_store.sql
--
-- phase10-08：为 Onboarding 建链恢复引入最小 current_product_id 锚点表。
-- 该表只承接单值恢复事实源，不复制页面草稿或派生状态。

CREATE TABLE IF NOT EXISTS onboarding_recovery_store (
    key TEXT PRIMARY KEY,
    value TEXT NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
