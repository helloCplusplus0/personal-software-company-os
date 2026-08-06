// Package candidate — Module 候选读取（由 Decision Center 拥有）。
//
// phase03-10 §10.5 ModuleCandidateRead 接口拥有者与接线：
//   - ModuleCandidateRead 的接口与实现必须由 Decision Center 的 candidate/ 子包自己定义和拥有
//   - Module Registry 不为 Decision Center 暴露专门的服务契约或服务方法
//   - candidate/ 子包通过独立 Read 接口隔离，service/ 层不得直接写跨模块 SQL
//   - 具体接线（构造与注入）必须在应用装配点（platform/router.go）完成
//
// 文件落点：backend/internal/decisioncenter/candidate/module_candidate_read.go
package candidate

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/psco/backend/internal/decisioncenter"
	"github.com/psco/backend/internal/moduleregistry"
)

// ModuleCandidateRead 承接 Decision -> Module 的候选读取。
//
// 只服务 LinkDecisionToTarget 的候选 Module 选择，不承接 Module 主线写入。
// 候选范围、排序与排除规则对齐 phase03-10 §5.10。
type ModuleCandidateRead struct {
	pool *pgxpool.Pool
}

// NewModuleCandidateRead 构造 ModuleCandidateRead。
func NewModuleCandidateRead(pool *pgxpool.Pool) *ModuleCandidateRead {
	return &ModuleCandidateRead{pool: pool}
}

// List 读取指定 Decision 可关联的 Module 候选列表。
//
// 候选范围与排序（phase03-10 §5.10）：
//   - 候选来源为当前已存在的 modules
//   - 候选范围同时覆盖 active 与 archived 的 Module，避免历史决策无法关联历史模块
//   - 排序采用 status(active 优先) -> module_name 升序
//   - 已建立 Decision -> Module 关联的目标不得再次出现在候选中
//   - status 复用 moduleregistry.ModuleStatus 类型
//
// 无可关联候选时返回空列表（[]），不返回 null，不得将空结果误报为接口错误
// （phase03-10 §5.10 空候选结果语义）。
func (r *ModuleCandidateRead) List(ctx context.Context, decisionID string) ([]decisioncenter.DecisionModuleCandidate, error) {
	rows, err := r.pool.Query(ctx, `
SELECT m.id, m.name, m.status
FROM modules m
WHERE m.id NOT IN (
    SELECT dl.module_id FROM decision_links dl WHERE dl.decision_id = $1
)
ORDER BY
    CASE m.status WHEN 'active' THEN 0 ELSE 1 END,
    m.name ASC`,
		decisionID,
	)
	if err != nil {
		return nil, fmt.Errorf("list module candidates: %w", err)
	}
	defer rows.Close()

	items := make([]decisioncenter.DecisionModuleCandidate, 0)
	for rows.Next() {
		var c decisioncenter.DecisionModuleCandidate
		if err := rows.Scan(&c.ModuleID, &c.ModuleName, &c.Status); err != nil {
			return nil, fmt.Errorf("scan module candidate row: %w", err)
		}
		items = append(items, c)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iter module candidate rows: %w", err)
	}
	return items, nil
}

// ModuleExists 校验指定 Module 是否存在（跨模块只读校验）。
//
// 用于 LinkDecisionToTarget 的目标 Module 存在性校验前提
// （phase03-04 校验顺序第 3 步）。
//
// 架构约束（phase03-10 §10.5）：
//   - service/ 层不得直接写跨模块 SQL
//   - 跨模块只读校验由 candidate/ 子包承接
//   - 本方法只读取 modules 表的存在性，不承接 Module 主线写入
func (r *ModuleCandidateRead) ModuleExists(ctx context.Context, moduleID string) (bool, error) {
	var exists bool
	err := r.pool.QueryRow(ctx,
		`SELECT EXISTS(SELECT 1 FROM modules WHERE id = $1)`,
		moduleID,
	).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("check module exists: %w", err)
	}
	return exists, nil
}

// 编译期断言：确保 DecisionModuleCandidate.Status 字段类型为 moduleregistry.ModuleStatus
// （phase03-10 §7.6 跨包依赖策略：不重定义本地等价枚举）。
var _ moduleregistry.ModuleStatus = decisioncenter.DecisionModuleCandidate{}.Status
