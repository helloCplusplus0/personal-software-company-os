-- seed_decision_link_fixture.sql
-- phase02-12 联调用 fixture：为已存在的首个模块建立一条 Decision 关联
--
-- 使用方式（在创建首个模块后执行）：
--   podman exec -i rento-preview-postgres psql -U rento -d psco_development \
--     < database/seeds/seed_decision_link_fixture.sql
--
-- 本脚本通过模块名查找目标模块，若模块不存在则跳过（不报错），
-- 以便在 phase02-12 联调流程中按需手动启用 Decision 入口展示。

DO $$
DECLARE
    target_module_id UUID;
    target_decision_id UUID;
BEGIN
    -- 取首个 active 模块作为关联目标
    SELECT id INTO target_module_id FROM modules WHERE status = 'active' ORDER BY created_at LIMIT 1;
    -- 取首条 decision 作为关联来源
    SELECT id INTO target_decision_id FROM decisions ORDER BY created_at LIMIT 1;

    IF target_module_id IS NOT NULL AND target_decision_id IS NOT NULL THEN
        INSERT INTO decision_links (decision_id, module_id)
        VALUES (target_decision_id, target_module_id)
        ON CONFLICT (decision_id, module_id) DO NOTHING;
        RAISE NOTICE 'Linked decision % to module %', target_decision_id, target_module_id;
    ELSE
        RAISE NOTICE 'Skipped: no active module or no decision found';
    END IF;
END $$;
