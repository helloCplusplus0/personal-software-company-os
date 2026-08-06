-- 0005_decision_source_context.sql
-- Decision Center 源上下文承接：decisions 表添加 source_module_id 字段
--
-- 上游规格：
--   - phase03-10 decision_center_spec_v0.1.md §5.11 入口上下文与正式关联结果边界
--   - phase03-03 spec §"Decision Create 入口上下文承接冻结"
--
-- 设计原则：
--   - phase03-10 §5.11 要求"该入口上下文中的 Module 必须带入 Decision Detail 作为
--     显式待关联目标继续承接，持续到用户完成正式 LinkDecisionToTarget"（当前阶段不提供主动放弃关联出口）
--   - "持续到完成"要求跨页面/跨刷新持久化，后端必须承接来源上下文存储（当前阶段不提供主动放弃出口）
--   - source_module_id 为可选字段（无来源时为 NULL），不参与 §5.5 结构化模板字段冻结
--   - ON DELETE SET NULL：Module 被删除时来源标识自动置空，不级联删除 Decision
--
-- 幂等性：IF NOT EXISTS 保证可重复执行

ALTER TABLE decisions ADD COLUMN IF NOT EXISTS source_module_id UUID REFERENCES modules(id) ON DELETE SET NULL;

CREATE INDEX IF NOT EXISTS idx_decisions_source_module_id ON decisions (source_module_id);
