-- seed_dashboard_fixture_pending_decisions.sql
-- phase05-12 验收 fixture 6：pending-decisions（有待决策信号，CTA 5）
--
-- 用途：
--   建立 Dashboard CTA 5 验收状态：pending_decision_signals → Decision Detail / List。
--   FeedbackSignalRead 产出 PENDING_DECISION 信号。
--
-- 数据状态（phase05-09 spec §"fixture 6 — pending-decisions"）：
--   - 至少 1 条 pending（proposed）状态的 decision
--   - 至少 1 条已绑定具体 decision_id 的 decision_link（验证单项信号跳转 Decision Detail）
--   - 允许 1 条未绑定单一 decision_id 的聚合决策信号（验证聚合信号跳转 Decision List）
--   - FeedbackSignalRead 产出 PENDING_DECISION 信号
--   - pending decision 状态判定沿用 phase03 proposed status 语义
--
-- 依赖：reset_dashboard_acceptance.sh --fixture pending-decisions 已执行清空操作
-- 幂等：使用 ON CONFLICT DO NOTHING + WHERE NOT EXISTS，可重复执行。
-- 使用方式：
--   ./database/scripts/reset_dashboard_acceptance.sh --fixture pending-decisions

BEGIN;

-- ============================================================================
-- 模块（1 条）：decision_links 依赖 modules
-- ============================================================================
INSERT INTO modules (name, description, status) VALUES
  ('auth-service', 'Dashboard fixture：认证服务模块', 'active')
ON CONFLICT (name) DO NOTHING;

-- ============================================================================
-- Decision 1：proposed 状态，将绑定 decision_link（单项信号 → Decision Detail）
-- ============================================================================
INSERT INTO decisions (title, context, problem, alternatives, choice, reason, impact, status)
SELECT 'Dashboard fixture：待决策记录 1',
       'Dashboard fixture 验证上下文',
       'Dashboard fixture 验证问题',
       '{}',
       'Dashboard fixture 验证选择',
       'Dashboard fixture 验证理由',
       '',
       'proposed'
WHERE NOT EXISTS (SELECT 1 FROM decisions WHERE title = 'Dashboard fixture：待决策记录 1');

-- ============================================================================
-- Decision 2：proposed 状态，不绑定 decision_link（聚合信号 → Decision List）
-- ============================================================================
INSERT INTO decisions (title, context, problem, alternatives, choice, reason, impact, status)
SELECT 'Dashboard fixture：待决策记录 2（无关联）',
       'Dashboard fixture 验证上下文',
       'Dashboard fixture 验证问题',
       '{}',
       'Dashboard fixture 验证选择',
       'Dashboard fixture 验证理由',
       '',
       'proposed'
WHERE NOT EXISTS (SELECT 1 FROM decisions WHERE title = 'Dashboard fixture：待决策记录 2（无关联）');

-- ============================================================================
-- decision_links（1 条）：Decision 1 → auth-service
-- 使 Decision 1 产生单项信号（target_type = DECISION_DETAIL）
-- ============================================================================
INSERT INTO decision_links (decision_id, module_id)
SELECT d.id, m.id
FROM decisions d
CROSS JOIN modules m
WHERE d.title = 'Dashboard fixture：待决策记录 1'
  AND m.name = 'auth-service'
ON CONFLICT (decision_id, module_id) DO NOTHING;

COMMIT;

-- ============================================================================
-- 校验
-- ============================================================================
SELECT 'modules' AS tbl, COUNT(*) AS cnt FROM modules
UNION ALL SELECT 'decisions', COUNT(*) FROM decisions
UNION ALL SELECT 'decision_links', COUNT(*) FROM decision_links;

-- pending decisions 与 decision_link 关联验证
SELECT d.title,
       d.status,
       EXISTS(SELECT 1 FROM decision_links dl WHERE dl.decision_id = d.id) AS has_link
FROM decisions d
WHERE d.status = 'proposed'
ORDER BY d.created_at DESC;
