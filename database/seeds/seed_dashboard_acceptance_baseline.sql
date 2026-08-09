-- seed_dashboard_acceptance_baseline.sql
-- phase05-12 验收辅助：Dashboard 验收基线补齐
--
-- 用途：
--   在 reset_dashboard_acceptance.sh --restore-only / 默认模式的恢复链最后一步执行，
--   补齐 Dashboard 验收所需的额外基线数据。
--
--   phase05-09 spec §"恢复范围"冻结的恢复顺序：
--     1. seed_readonly_prereqs.sql（products / repositories / decisions 占位）
--     2. reset_module_mainline.sh --restore-only（modules / module_releases / product_modules / module_repositories / decision_links）
--     3. reset_decision_mainline.sh --restore-only（decisions 正式基线 + decision_links）
--     4. reset_product_repository_mainline.sh --restore-only（products / repositories 正式基线 + 绑定关系）
--     5. seed_dashboard_acceptance_baseline.sql（本文件，补齐 Dashboard 验收所需额外数据）
--
--   经过前四步恢复后，基线状态已满足 Dashboard 最小有数据验收基线要求
--   （phase05-09 spec §"最小有数据验收基线建立方式"）：
--     - module_count >= 2（auth-service / integration-test-module）
--     - product_count >= 3（Product A / B / C）
--     - repository_count >= 2（main-repo / mirror-repo）
--     - decision_count >= 1（3 条 decisions）
--     - product_with_module_count = 1（Product A）
--     - product_with_repository_count = 1（Product A）
--     - 资产缺口：Product B / C 缺少双绑定（missing_both_bindings）
--     - 活动项：modules / module_releases / products / repositories / decisions / bindings 均有记录
--
--   本文件当前不新增额外数据，仅作为恢复链的收口校验点。
--   后续若 Dashboard 验收需要额外基线数据（如特定时间戳的活动项），
--   应在本文件中以幂等方式补充，不得发明第二套基线入口。
--
-- 幂等：可重复执行。本文件只包含 SELECT 校验查询，不修改数据，天然幂等。
--
-- 上游规格：phase05-09 spec §"恢复范围"
--           phase05-12 spec §"reset_dashboard_acceptance.sh 必须作为 Dashboard 验收统一入口"
-- 依赖：seed_readonly_prereqs.sql + 三个既有 reset 脚本的 --restore-only 已执行
--
-- 使用方式：
--   # 通过封装脚本（推荐）
--   ./database/scripts/reset_dashboard_acceptance.sh              # 清空 + 恢复（含本文件）
--   ./database/scripts/reset_dashboard_acceptance.sh --restore-only # 仅恢复（含本文件）
--
--   # 直接执行 SQL（需已执行前四步恢复）
--   podman exec -i rento-preview-postgres psql -U rento -d psco_development \
--     -v ON_ERROR_STOP=1 \
--     < database/seeds/seed_dashboard_acceptance_baseline.sql

-- ============================================================================
-- 校验：Dashboard 最小有数据基线验证（phase05-09 spec §"最小有数据基线验证点"）
-- ============================================================================
-- 以下查询为信息性输出，不阻断执行。若计数不满足要求，说明前四步恢复未完成。

SELECT 'modules' AS tbl, COUNT(*) AS cnt FROM modules
UNION ALL SELECT 'module_releases', COUNT(*) FROM module_releases
UNION ALL SELECT 'products', COUNT(*) FROM products
UNION ALL SELECT 'repositories', COUNT(*) FROM repositories
UNION ALL SELECT 'decisions', COUNT(*) FROM decisions
UNION ALL SELECT 'product_modules', COUNT(*) FROM product_modules
UNION ALL SELECT 'product_repositories', COUNT(*) FROM product_repositories
UNION ALL SELECT 'module_repositories', COUNT(*) FROM module_repositories
UNION ALL SELECT 'decision_links', COUNT(*) FROM decision_links;

-- DashboardOverview 关键计数验证
SELECT
  'dashboard_overview' AS check_name,
  (SELECT COUNT(*) FROM modules) AS module_count,
  (SELECT COUNT(*) FROM products) AS product_count,
  (SELECT COUNT(*) FROM repositories) AS repository_count,
  (SELECT COUNT(*) FROM decisions) AS decision_count,
  (SELECT COUNT(DISTINCT product_id) FROM product_modules) AS product_with_module_count,
  (SELECT COUNT(DISTINCT product_id) FROM product_repositories) AS product_with_repository_count;
