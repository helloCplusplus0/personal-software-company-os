// Package repository — decision_links 表数据访问层。
//
// 只负责持久化与读取，不承接业务校验或跨模块编排（phase03-10 §10.3 分层语义）。
//
// 文件落点：backend/internal/decisioncenter/repository/link_store.go
package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/psco/backend/internal/decisioncenter"
)

// LinkStore 承接 decision_links 表的读写。
type LinkStore struct {
	pool *pgxpool.Pool
}

// NewLinkStore 构造 LinkStore。
func NewLinkStore(pool *pgxpool.Pool) *LinkStore {
	return &LinkStore{pool: pool}
}

// Create 插入一条 Decision -> Module 关联记录。
// 重复关联（UNIQUE(decision_id, module_id) 冲突）返回 decisioncenter.ErrDuplicateLink。
//
// 调用方应先完成 Decision / Module 存在性校验；FK 约束提供兜底保护。
func (s *LinkStore) Create(ctx context.Context, decisionID, moduleID string) error {
	tag, err := s.pool.Exec(ctx, `
INSERT INTO decision_links (decision_id, module_id)
VALUES ($1, $2)
ON CONFLICT (decision_id, module_id) DO NOTHING`,
		decisionID, moduleID,
	)
	if err != nil {
		return fmt.Errorf("insert decision link: %w", err)
	}
	if tag.RowsAffected() == 0 {
		return decisioncenter.ErrDuplicateLink
	}
	return nil
}

// ListByDecisionID 读取指定 Decision 已关联的 Module 列表，附带 module_name
// （phase03-10 §5.8 详情读取中的 linked_modules 组装基础）。
//
// 结果按 module_name 升序排序，便于前端稳定展示。
func (s *LinkStore) ListByDecisionID(ctx context.Context, decisionID string) ([]decisioncenter.LinkedModule, error) {
	rows, err := s.pool.Query(ctx, `
SELECT dl.module_id, m.name
FROM decision_links dl
JOIN modules m ON m.id = dl.module_id
WHERE dl.decision_id = $1
ORDER BY m.name ASC`,
		decisionID,
	)
	if err != nil {
		return nil, fmt.Errorf("list decision links by decision: %w", err)
	}
	defer rows.Close()

	items := make([]decisioncenter.LinkedModule, 0)
	for rows.Next() {
		var lm decisioncenter.LinkedModule
		if err := rows.Scan(&lm.ModuleID, &lm.ModuleName); err != nil {
			return nil, fmt.Errorf("scan linked module row: %w", err)
		}
		items = append(items, lm)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iter linked module rows: %w", err)
	}
	return items, nil
}

// ExistsByDecisionAndModule 检测指定 Decision -> Module 关联是否已存在
// （重复关联检测，phase03-04 LinkDecisionToTarget 校验顺序第 4 步）。
func (s *LinkStore) ExistsByDecisionAndModule(ctx context.Context, decisionID, moduleID string) (bool, error) {
	var exists bool
	err := s.pool.QueryRow(ctx,
		`SELECT EXISTS(SELECT 1 FROM decision_links WHERE decision_id = $1 AND module_id = $2)`,
		decisionID, moduleID,
	).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("check decision link exists: %w", err)
	}
	return exists, nil
}

// ListLinkedModuleIDs 读取指定 Decision 已关联的 module_id 列表
// （候选读取排除已关联目标的前提，phase03-10 §5.10）。
func (s *LinkStore) ListLinkedModuleIDs(ctx context.Context, decisionID string) ([]string, error) {
	rows, err := s.pool.Query(ctx,
		`SELECT module_id FROM decision_links WHERE decision_id = $1`,
		decisionID,
	)
	if err != nil {
		return nil, fmt.Errorf("list linked module ids: %w", err)
	}
	defer rows.Close()

	ids := make([]string, 0)
	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			return nil, fmt.Errorf("scan linked module id: %w", err)
		}
		ids = append(ids, id)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iter linked module id rows: %w", err)
	}
	return ids, nil
}
