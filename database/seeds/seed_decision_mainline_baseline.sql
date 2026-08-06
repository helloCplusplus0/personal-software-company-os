-- seed_decision_mainline_baseline.sql
-- phase03-12 验收辅助：清空 Decision 主线 + 恢复基线数据
--
-- 用途：
--   为 phase03-14 空状态验收与基线复验提供可重复执行的"清空 -> 恢复基线"入口。
--   清空仅影响 Decision 主线（decisions 及其级联表 decision_links），
--   不影响只读前提数据（products / repositories）与模块主线（modules）。
--
-- 幂等：可重复执行。
--   - 清空阶段：DELETE FROM decisions 依赖 ON DELETE CASCADE 级联删除 decision_links
--   - 恢复阶段：
--     * Decision 1（复用 phase02 标题）使用"先 UPDATE 收口已有 placeholder → 再 INSERT 补缺"，
--       确保 --restore-only 模式下也能把 readonly placeholder 收口为正式基线内容
--     * Decision 2/3（新标题）使用 WHERE NOT EXISTS 守卫
--     * decision_links 使用 ON CONFLICT (decision_id, module_id) DO NOTHING
--
-- 上游规格：phase03-09 spec §"Decision Center 基线种子数据范围必须冻结"
--           phase03-10 decision_center_spec_v0.1.md §11.3 基线种子数据
-- 依赖：reset_module_mainline.sh 已执行（提供 modules 基线，decision_links 依赖 modules）
--
-- 使用方式：
--   # 通过封装脚本（推荐）
--   ./database/scripts/reset_decision_mainline.sh              # 清空 + 恢复
--   ./database/scripts/reset_decision_mainline.sh --clean-only # 仅清空（验证空状态）
--   ./database/scripts/reset_decision_mainline.sh --restore-only # 仅恢复
--
--   # 直接执行 SQL（需已运行 reset_module_mainline.sh）
--   podman exec -i rento-preview-postgres psql -U rento -d psco_development \
--     -v ON_ERROR_STOP=1 \
--     < database/seeds/seed_decision_mainline_baseline.sql

BEGIN;

-- ============================================================================
-- 清空阶段：删除 Decision 主线所有数据
-- ============================================================================
-- 依赖 schema 定义的 ON DELETE CASCADE：
--   - decision_links.decision_id REFERENCES decisions(id) ON DELETE CASCADE
-- 因此只需 DELETE FROM decisions 即可级联清空 decision_links。
DELETE FROM decisions;

-- ============================================================================
-- 恢复阶段：重建 phase03-14 验收基线数据
-- ============================================================================
-- 设计原则：通过 module_name 与 decision title 查找 ID，避免硬编码 UUID
-- 覆盖维度（phase03-10 §11.3）：
--   - 3 条结构化 Decision：1 proposed + 1 active + 1 archived
--   - 1 条保留 phase02 原有 title（关于 auth-service 技术选型的决策）
--   - 至少 1 条包含 alternatives 数组（2 个以上条目）
--   - 至少 1 条 alternatives 为空数组
--   - 至少 1 条 impact 为空字符串
--   - 2 条 decision_links，关联到 modules 基线
--   - 1 条 Decision 无 decision_links（验证 link_count=0 与 linked_module_summary=''）
--   - 1 条 decision_links 复用 phase02 原有 auth-service 关联

-- Decision 1：保留 phase02 原有 title，proposed 状态，空 alternatives，空 impact
-- （兼容 phase02 decision_links 基线）
-- 带来源上下文：source_module_id 设为 auth-service（模拟从 Module Detail 带上下文创建）
-- 幂等说明：decisions.title 无 UNIQUE 约束，先 UPDATE 收口已有 placeholder（含
--   seed_readonly_prereqs.sql 插入的占位内容）为正式基线内容，再 INSERT 补缺。
--   这保证 --restore-only 模式下也能恢复到正式基线，而不是保留 readonly placeholder。
UPDATE decisions
SET context     = '团队需要为微服务架构选择认证服务实现方案',
    problem     = '在自研与开源认证服务之间做出选择',
    alternatives = '{}',
    choice      = '采用自研 auth-service',
    reason      = '自研方案更贴合现有架构且可控性更高',
    impact      = '',
    status      = 'proposed',
    source_module_id = (SELECT id FROM modules WHERE name = 'auth-service')
WHERE title = '关于 auth-service 技术选型的决策'
  AND (context IS NULL
       OR context != '团队需要为微服务架构选择认证服务实现方案');

INSERT INTO decisions (title, context, problem, alternatives, choice, reason, impact, status, source_module_id)
SELECT '关于 auth-service 技术选型的决策',
       '团队需要为微服务架构选择认证服务实现方案',
       '在自研与开源认证服务之间做出选择',
       '{}',
       '采用自研 auth-service',
       '自研方案更贴合现有架构且可控性更高',
       '',
       'proposed',
       (SELECT id FROM modules WHERE name = 'auth-service')
WHERE NOT EXISTS (SELECT 1 FROM decisions WHERE title = '关于 auth-service 技术选型的决策');

-- Decision 2：active 状态，含 alternatives 数组（2 个条目），有 impact
INSERT INTO decisions (title, context, problem, alternatives, choice, reason, impact, status)
SELECT '统一日志采集方案决策',
       '微服务数量增加后需要统一日志采集与查询',
       '在 ELK 与 Loki 之间做出选择',
       '{"ELK 全家桶", "Loki + Grafana"}',
       '选择 Loki + Grafana',
       'Loki 资源占用更低且与现有 Grafana 集成',
       '需迁移现有日志管道',
       'active'
WHERE NOT EXISTS (SELECT 1 FROM decisions WHERE title = '统一日志采集方案决策');

-- Decision 3：archived 状态，空 alternatives，空 impact
-- （此条 Decision 无 decision_links，验证 link_count=0 与 linked_module_summary=''）
INSERT INTO decisions (title, context, problem, alternatives, choice, reason, impact, status)
SELECT 'API 网关选型决策',
       '需要为对外 API 引入网关层',
       '在 Kong 与 Traefik 之间做出选择',
       '{}',
       '选择 Traefik',
       'Traefik 与 Caddy 部署栈更一致',
       '',
       'archived'
WHERE NOT EXISTS (SELECT 1 FROM decisions WHERE title = 'API 网关选型决策');

-- ============================================================================
-- decision_links 基线（2 条）
-- ============================================================================
-- 通过 module_name 与 decision title 查找 ID，不硬编码 UUID
-- 其中 1 条复用 phase02 原有 auth-service 关联

-- Link 1：关于 auth-service 技术选型的决策 -> auth-service（复用 phase02 关联）
INSERT INTO decision_links (decision_id, module_id)
SELECT d.id, m.id
FROM decisions d
CROSS JOIN modules m
WHERE d.title = '关于 auth-service 技术选型的决策'
  AND m.name = 'auth-service'
ON CONFLICT (decision_id, module_id) DO NOTHING;

-- Link 2：统一日志采集方案决策 -> auth-service
INSERT INTO decision_links (decision_id, module_id)
SELECT d.id, m.id
FROM decisions d
CROSS JOIN modules m
WHERE d.title = '统一日志采集方案决策'
  AND m.name = 'auth-service'
ON CONFLICT (decision_id, module_id) DO NOTHING;

-- 注：API 网关选型决策 不建立 decision_links，验证 link_count=0 与 linked_module_summary=''

COMMIT;

-- ============================================================================
-- 验证恢复结果（信息性输出，不阻断）
-- ============================================================================
SELECT 'decisions' AS tbl, COUNT(*) FROM decisions
UNION ALL SELECT 'decision_links', COUNT(*) FROM decision_links;
