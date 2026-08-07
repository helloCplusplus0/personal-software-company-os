// Package candidate — Module 候选读取（Product Registry 拥有）。
//
// 对齐 phase04-07 跨模块候选读取边界：
//   - ProductModuleCandidateRead 由 Product Registry 的 candidate 子包定义和拥有
//   - 通过独立 Read 接口隔离，service 层不直接写跨模块 SQL
//   - 读取 modules 表是为了获取候选 Module 列表（排除已绑定的 Module）
//
// 文件落点：backend/internal/productregistry/candidate/module_candidate_read.go
package candidate

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/psco/backend/internal/moduleregistry"
	"github.com/psco/backend/internal/productregistry"
)

// ModuleCandidateRead 承接 ProductModuleCandidateRead 接口。
//
// 只服务 BindModuleToProduct 的候选 Module 选择，不承接 Module 主线写入。
type ModuleCandidateRead struct {
	pool *pgxpool.Pool
}

// NewModuleCandidateRead 构造 ModuleCandidateRead。
func NewModuleCandidateRead(pool *pgxpool.Pool) *ModuleCandidateRead {
	return &ModuleCandidateRead{pool: pool}
}

// List 读取指定产品的可绑定 Module 候选，排除已绑定的 Module。
//
// 排序按 modules.created_at DESC（phase04-04 冻结：候选读取均按 created_at 降序）。
// 无可关联候选时返回空列表，不返回错误。
func (r *ModuleCandidateRead) List(ctx context.Context, productID string) ([]productregistry.ProductModuleCandidate, error) {
	rows, err := r.pool.Query(ctx, `
SELECT m.id, m.name, m.status
FROM modules m
WHERE m.status = 'active'
  AND m.id NOT IN (
    SELECT pm.module_id FROM product_modules pm WHERE pm.product_id = $1
  )
ORDER BY m.created_at DESC`,
		productID,
	)
	if err != nil {
		return nil, fmt.Errorf("list product module candidates: %w", err)
	}
	defer rows.Close()

	var items []productregistry.ProductModuleCandidate
	for rows.Next() {
		var c productregistry.ProductModuleCandidate
		if err := rows.Scan(&c.ModuleID, &c.ModuleName, &c.ModuleStatus); err != nil {
			return nil, fmt.Errorf("scan product module candidate row: %w", err)
		}
		items = append(items, c)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iter product module candidate rows: %w", err)
	}
	return items, nil
}

// CheckModuleExistsActive 校验 Module 存在性与 active 状态（绑定前提校验）。
//
// 返回三态（phase04-04 错误语义前提：不存在 → 404，存在但非 active → 400）：
//   - exists=false, active=false：Module 不存在（service 层映射 ErrModuleNotFound → 404）
//   - exists=true, active=false：Module 存在但状态非 active（service 层映射 ErrModuleNotActive → 400）
//   - exists=true, active=true：可用于绑定
//
// 单次 SQL 同时取 EXISTS(id) 与 EXISTS(id AND status='active')，避免两次往返。
func (r *ModuleCandidateRead) CheckModuleExistsActive(ctx context.Context, moduleID string) (exists, active bool, err error) {
	err = r.pool.QueryRow(ctx, `
SELECT
  EXISTS(SELECT 1 FROM modules WHERE id = $1),
  EXISTS(SELECT 1 FROM modules WHERE id = $1 AND status = 'active')`,
		moduleID,
	).Scan(&exists, &active)
	if err != nil {
		return false, false, fmt.Errorf("check module exists and active: %w", err)
	}
	return exists, active, nil
}

// 确保编译期 moduleregistry 包被引用（module_status 类型复用 moduleregistry.ModuleStatus）。
var _ moduleregistry.ModuleStatus
