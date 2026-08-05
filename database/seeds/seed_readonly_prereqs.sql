-- seed_readonly_prereqs.sql
-- phase02-12 联调验收最小候选数据与示例数据
-- 上游规格：phase02-11 spec §"迁移与最小种子数据必须可支撑 phase02-12 验收"
--
-- 本文件幂等可重复执行（ON CONFLICT DO NOTHING），只服务当前阶段联调与验收，
-- 不扩写为第二套业务主线，不视为 phase02 写入主线。
--
-- 说明：decision_links 表因 FK 引用 modules，无法在用户创建模块前预置，
-- 故不在此处 seed。phase02-12 联调时可在创建首个模块后，
-- 通过 scripts/seed_decision_link_fixture.sql 手动建立一条示例关联。

-- products 候选读取前提（对齐前端 mock 数据）
INSERT INTO products (name) VALUES ('Product A') ON CONFLICT (name) DO NOTHING;
INSERT INTO products (name) VALUES ('Product B') ON CONFLICT (name) DO NOTHING;
INSERT INTO products (name) VALUES ('Product C') ON CONFLICT (name) DO NOTHING;

-- repositories 候选读取前提（对齐前端 mock 数据）
INSERT INTO repositories (name) VALUES ('main-repo') ON CONFLICT (name) DO NOTHING;
INSERT INTO repositories (name) VALUES ('mirror-repo') ON CONFLICT (name) DO NOTHING;

-- decisions 只读入口前提（对齐前端 mock 数据）
-- title 唯一性通过 ON CONFLICT 在插入重复时跳过（title 无 UNIQUE 约束，此处靠应用层幂等；
-- 若需严格幂等，可在 decisions.name 上加 UNIQUE 约束，但 phase02 不强求）
INSERT INTO decisions (title)
SELECT '关于 auth-service 技术选型的决策'
WHERE NOT EXISTS (SELECT 1 FROM decisions WHERE title = '关于 auth-service 技术选型的决策');
