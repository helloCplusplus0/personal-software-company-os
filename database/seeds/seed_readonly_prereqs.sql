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
-- phase03-12 升级：从 title-only 升级为结构化字段插入
-- 必须保持原有 title（关于 auth-service 技术选型的决策）以兼容 phase02 decision_links
-- 必须补全 context / problem / choice / reason / status 必填字段
-- alternatives 设为 '{}'，impact 设为 ''（对齐 phase03-10 §5.5 创建可选语义）
-- 上游规格：phase03-09 spec §"seed_readonly_prereqs.sql decisions seed 更新"
INSERT INTO decisions (title, context, problem, alternatives, choice, reason, impact, status)
SELECT '关于 auth-service 技术选型的决策',
       'phase02 只读前提占位上下文',
       'phase02 只读前提占位问题',
       '{}',
       'phase02 只读前提占位选择',
       'phase02 只读前提占位理由',
       '',
       'proposed'
WHERE NOT EXISTS (SELECT 1 FROM decisions WHERE title = '关于 auth-service 技术选型的决策');
