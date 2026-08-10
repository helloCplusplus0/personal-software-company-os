-- seed_phase06_acceptance_baseline.sql
-- phase06-14 验收辅助：Phase06 验收基线补齐
--
-- 用途：
--   在 reset_phase06_acceptance.sh --restore-only / 默认模式的恢复链最后一步执行，
--   补齐 phase06 验收所需的额外基线数据（capability_key 与复用绑定）。
--
--   phase06-14 spec §"reset_phase06_acceptance.sh 必须把 11 个 fixture 变成可重复执行的正式主线"：
--     1. reset_dashboard_acceptance.sh --restore-only（恢复 phase05 Dashboard / Feedback 验收基线）
--     2. seed_phase06_acceptance_baseline.sql（本文件，补齐 phase06 验收所需额外数据）
--
--   经过前一步恢复后，基线状态已满足 phase05 最小有数据验收基线要求。
--   本文件在此基础上补齐 phase06 特有数据：
--     - 为部分 modules 设置 capability_key（支撑 capability_summary 聚合）
--     - 补充 product_modules 绑定（支撑 module_reuse_summary 跨 Product 复用统计）
--
-- 幂等：使用 UPDATE WHERE + ON CONFLICT DO NOTHING，可重复执行。
--
-- 上游规格：phase06-11 spec §"reset_phase06_acceptance.sh 必须作为统一验收入口"
--           phase06-14 spec §"Fixture 白名单文件落点"
-- 依赖：reset_dashboard_acceptance.sh --restore-only 已执行
--
-- 使用方式：
--   ./database/scripts/reset_phase06_acceptance.sh              # 清空 + 恢复（含本文件）
--   ./database/scripts/reset_phase06_acceptance.sh --restore-only # 仅恢复（含本文件）

BEGIN;

-- ============================================================================
-- 为 phase05 基线 modules 补充 capability_key
-- ============================================================================
-- phase06-04 / phase06-12 §"capability_summary 事实来源、映射与缺省策略冻结"：
--   必须以 Module.capability_key 作为唯一聚合主键来源

UPDATE modules SET capability_key = 'auth' WHERE name = 'auth-service' AND capability_key IS NULL;
UPDATE modules SET capability_key = 'web_frontend' WHERE name = 'integration-test-module' AND capability_key IS NULL;

-- ============================================================================
-- 补充复用绑定：让 auth-service 被 2 个 Product 复用（形成跨 Product 复用）
-- ============================================================================
-- phase06-14 spec §"module_reuse_summary 聚合口径"：
--   reuse_product_count 必须表示"当前被多少 Product 直接复用"

-- auth-service 已绑定 Product A（phase05 基线），再绑定 Product B
INSERT INTO product_modules (module_id, product_id)
SELECT m.id, p.id
FROM modules m
CROSS JOIN products p
WHERE m.name = 'auth-service'
  AND p.name = 'Product B'
ON CONFLICT (module_id, product_id) DO NOTHING;

COMMIT;

-- ============================================================================
-- 校验
-- ============================================================================
SELECT 'modules_with_capability' AS tbl, COUNT(*) AS cnt FROM modules WHERE capability_key IS NOT NULL
UNION ALL SELECT 'product_modules_total', COUNT(*) FROM product_modules
UNION ALL SELECT 'auth_service_reuse', COUNT(DISTINCT pm.product_id)
FROM product_modules pm
JOIN modules m ON m.id = pm.module_id
WHERE m.name = 'auth-service';
