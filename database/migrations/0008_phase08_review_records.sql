-- 0008_phase08_review_records.sql
-- Phase08 Operating Review Loop 最小数据承接位
-- 直接承接表：review_records（新增）
-- 上游规格：phase08-04 冻结本阶段合同、读模型与记录模型的最小边界 Spec
--           phase08-07 产出后端服务、合同与最小数据承接设计 Spec
--           phase08-08 落实 review 相关合同、后端承接与前端 owner 收敛 Spec
--
-- 设计约束：
--   - review_records 是单表轻量过程记录，不升级为新的长期核心实体
--   - 不得复制完整实体快照、完整 review context 或 Decision/Product/Module/Repository 结构
--   - 只服务于 next-step result 或可选 review 过程留痕
--   - decision handoff / entity handoff 路径允许不写入此表

CREATE TABLE review_records (
    id                  UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at          TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    -- review 会话类型：daily / weekly
    review_kind         TEXT NOT NULL CHECK (review_kind IN ('daily', 'weekly')),
    -- review 结果类型：decision_handoff / entity_handoff / next_step_result
    result_kind         TEXT NOT NULL CHECK (result_kind IN ('decision_handoff', 'entity_handoff', 'next_step_result')),
    -- 可选关联的 decision_id
    decision_id         TEXT NULL,
    -- 可选关联的目标实体类型
    target_type         TEXT NULL,
    -- 可选关联的目标实体 id
    target_id           TEXT NULL,
    -- 最小摘要文本
    summary_text        TEXT NOT NULL DEFAULT '',
    -- review 开始时间
    started_at          TIMESTAMPTZ NOT NULL,
    -- review 完成时间
    completed_at        TIMESTAMPTZ NOT NULL
);

-- 按创建时间倒序读取最近 review 记录
CREATE INDEX idx_review_records_created_at ON review_records (created_at DESC);
-- 按 review 类型筛选
CREATE INDEX idx_review_records_review_kind ON review_records (review_kind);