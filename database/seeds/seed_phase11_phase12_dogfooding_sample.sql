-- seed_phase11_phase12_dogfooding_sample.sql
-- phase11 / phase12 固定 dogfooding 样本恢复
--
-- 用途：
--   恢复 phase11 / phase12 验收冻结的 PSCO 自身固定样本，
--   使 repository / product / module / decision detail 与 shared-readonly
--   能基于同一 repository_id 重放验证。
--
-- 范围：
--   仅补齐固定样本与其 canonical 关系，不清空现有开发数据；
--   若同名实体已被其他 id 占用，则显式失败，避免静默覆盖。
--
-- 固定样本：
--   Repository: personal-software-company-os
--   repository_id: ca261521-8daf-4248-8f12-43525326e759
--   Product: PSCO
--   product_id: f0d034cc-3235-4d03-879d-6a3111b95b6b
--   Module: project-context-foundation
--   module_id: 9b02e0ca-3175-4b5a-bcf4-23cecbb06f72
--   Decision: phase11 Project Context Foundation dogfooding 验收决策
--   decision_id: aa8ee5ad-b224-4ea1-b393-7e4c30e42212
--
-- 幂等：
--   - 实体以固定 id upsert
--   - 关系以唯一键 ON CONFLICT DO NOTHING
--   - 同名异 id 冲突时显式报错，要求先清理错误样本

BEGIN;

DO $$
BEGIN
  IF EXISTS (
    SELECT 1
    FROM repositories
    WHERE name = 'personal-software-company-os'
      AND id <> 'ca261521-8daf-4248-8f12-43525326e759'
  ) THEN
    RAISE EXCEPTION 'repository name personal-software-company-os already exists with different id';
  END IF;

  IF EXISTS (
    SELECT 1
    FROM products
    WHERE name = 'PSCO'
      AND id <> 'f0d034cc-3235-4d03-879d-6a3111b95b6b'
  ) THEN
    RAISE EXCEPTION 'product name PSCO already exists with different id';
  END IF;

  IF EXISTS (
    SELECT 1
    FROM modules
    WHERE name = 'project-context-foundation'
      AND id <> '9b02e0ca-3175-4b5a-bcf4-23cecbb06f72'
  ) THEN
    RAISE EXCEPTION 'module name project-context-foundation already exists with different id';
  END IF;

  IF EXISTS (
    SELECT 1
    FROM decisions
    WHERE title = 'phase11 Project Context Foundation dogfooding 验收决策'
      AND id <> 'aa8ee5ad-b224-4ea1-b393-7e4c30e42212'
  ) THEN
    RAISE EXCEPTION 'decision title phase11 Project Context Foundation dogfooding 验收决策 already exists with different id';
  END IF;
END $$;

INSERT INTO repositories (id, name, url, provider, status)
VALUES (
  'ca261521-8daf-4248-8f12-43525326e759',
  'personal-software-company-os',
  'https://github.com/psco/personal-software-company-os',
  'github',
  'active'
)
ON CONFLICT (id) DO UPDATE
SET
  name = EXCLUDED.name,
  url = EXCLUDED.url,
  provider = EXCLUDED.provider,
  status = EXCLUDED.status;

INSERT INTO products (id, name, description, status)
VALUES (
  'f0d034cc-3235-4d03-879d-6a3111b95b6b',
  'PSCO',
  'phase11/12 dogfooding 固定样本：项目上下文系统产品锚点',
  'active'
)
ON CONFLICT (id) DO UPDATE
SET
  name = EXCLUDED.name,
  description = EXCLUDED.description,
  status = EXCLUDED.status;

INSERT INTO modules (id, name, description, status, capability_key)
VALUES (
  '9b02e0ca-3175-4b5a-bcf4-23cecbb06f72',
  'project-context-foundation',
  'phase11/12 dogfooding 固定样本：共享只读项目上下文能力资产',
  'active',
  'project_context'
)
ON CONFLICT (id) DO UPDATE
SET
  name = EXCLUDED.name,
  description = EXCLUDED.description,
  status = EXCLUDED.status,
  capability_key = EXCLUDED.capability_key;

INSERT INTO decisions (
  id,
  title,
  context,
  problem,
  alternatives,
  choice,
  reason,
  impact,
  status,
  source_module_id
)
VALUES (
  'aa8ee5ad-b224-4ea1-b393-7e4c30e42212',
  'phase11 Project Context Foundation dogfooding 验收决策',
  'phase11/12 固定样本：以 PSCO 自身仓库验证共享只读与语义对齐。',
  '缺少固定样本时，browser regression 与 shared-readonly 无法使用同一 repository_id 重放。',
  ARRAY['继续依赖 phase06 fixture', '补齐 PSCO 自身固定样本'],
  '补齐 PSCO 自身固定样本，并通过同一 repository_id 统一 detail 与 shared-readonly 验证入口。',
  '该样本是 phase11/12 验收协议冻结的一部分，必须可重复恢复。',
  '',
  'active',
  '9b02e0ca-3175-4b5a-bcf4-23cecbb06f72'
)
ON CONFLICT (id) DO UPDATE
SET
  title = EXCLUDED.title,
  context = EXCLUDED.context,
  problem = EXCLUDED.problem,
  alternatives = EXCLUDED.alternatives,
  choice = EXCLUDED.choice,
  reason = EXCLUDED.reason,
  impact = EXCLUDED.impact,
  status = EXCLUDED.status,
  source_module_id = EXCLUDED.source_module_id;

INSERT INTO product_modules (module_id, product_id)
VALUES (
  '9b02e0ca-3175-4b5a-bcf4-23cecbb06f72',
  'f0d034cc-3235-4d03-879d-6a3111b95b6b'
)
ON CONFLICT (module_id, product_id) DO NOTHING;

INSERT INTO product_repositories (product_id, repository_id)
VALUES (
  'f0d034cc-3235-4d03-879d-6a3111b95b6b',
  'ca261521-8daf-4248-8f12-43525326e759'
)
ON CONFLICT (product_id, repository_id) DO NOTHING;

INSERT INTO module_repositories (module_id, repository_id)
VALUES (
  '9b02e0ca-3175-4b5a-bcf4-23cecbb06f72',
  'ca261521-8daf-4248-8f12-43525326e759'
)
ON CONFLICT (module_id, repository_id) DO NOTHING;

INSERT INTO decision_links (decision_id, module_id)
VALUES (
  'aa8ee5ad-b224-4ea1-b393-7e4c30e42212',
  '9b02e0ca-3175-4b5a-bcf4-23cecbb06f72'
)
ON CONFLICT (decision_id, module_id) DO NOTHING;

COMMIT;

SELECT 'repository_id', id::text FROM repositories WHERE id = 'ca261521-8daf-4248-8f12-43525326e759'
UNION ALL
SELECT 'product_id', id::text FROM products WHERE id = 'f0d034cc-3235-4d03-879d-6a3111b95b6b'
UNION ALL
SELECT 'module_id', id::text FROM modules WHERE id = '9b02e0ca-3175-4b5a-bcf4-23cecbb06f72'
UNION ALL
SELECT 'decision_id', id::text FROM decisions WHERE id = 'aa8ee5ad-b224-4ea1-b393-7e4c30e42212';
