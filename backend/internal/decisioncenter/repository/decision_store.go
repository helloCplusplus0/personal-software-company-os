// Package repository — decisions 表数据访问层。
//
// 只负责持久化与读取，不承接业务校验或跨模块编排（phase03-10 §10.3 分层语义）。
//
// 文件落点：backend/internal/decisioncenter/repository/decision_store.go
package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/psco/backend/internal/decisioncenter"
)

// DecisionStore 承接 decisions 表的读写。
type DecisionStore struct {
	pool *pgxpool.Pool
}

// NewDecisionStore 构造 DecisionStore。
func NewDecisionStore(pool *pgxpool.Pool) *DecisionStore {
	return &DecisionStore{pool: pool}
}

// DecisionWithSource 包装 Decision 核心对象与其来源上下文信息。
//
// SourceModuleID / SourceModuleName 来自 decisions.source_module_id 与
// modules 表的 LEFT JOIN，用于 service 层组装 DecisionDetail.SourceContext
// （phase03-10 §5.11 入口上下文与正式关联结果边界）。
// SourceModuleID 为空字符串表示无来源上下文（source_module_id IS NULL）。
type DecisionWithSource struct {
	Decision         decisioncenter.Decision
	SourceModuleID   string
	SourceModuleName string
}

// Create 插入一条结构化决策记录，返回带 id / created_at 的完整对象。
//
// alternatives 按输入顺序保留；空数组写入 '{}'，不写入 NULL
// （phase03-10 §5.5 / §5.7 alternatives 写入语义）。
// source_module_id 为可选来源上下文（§5.11），空字符串写入 NULL。
func (s *DecisionStore) Create(ctx context.Context, req decisioncenter.CreateDecisionRequest) (*decisioncenter.Decision, error) {
	// 空切片统一为 '{}'，避免 pgx 把 nil 切片写入 NULL
	alts := req.Alternatives
	if alts == nil {
		alts = []string{}
	}

	// source_module_id 空字符串转为 NULL
	var sourceModuleID any
	if req.SourceModuleID != "" {
		sourceModuleID = req.SourceModuleID
	}

	d := &decisioncenter.Decision{}
	err := s.pool.QueryRow(ctx, `
INSERT INTO decisions (title, context, problem, alternatives, choice, reason, impact, status, source_module_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING id, title, context, problem, alternatives, choice, reason, impact, status, created_at`,
		req.Title, req.Context, req.Problem, alts,
		req.Choice, req.Reason, req.Impact, string(req.Status), sourceModuleID,
	).Scan(
		&d.ID, &d.Title, &d.Context, &d.Problem, &d.Alternatives,
		&d.Choice, &d.Reason, &d.Impact, &d.Status, &d.CreatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("insert decision: %w", err)
	}
	return d, nil
}

// GetByID 按 id 读取单个决策及其来源上下文。未找到时返回 decisioncenter.ErrDecisionNotFound。
//
// 通过 LEFT JOIN modules 获取 source_module_name，用于组装 SourceContext
// （phase03-10 §5.11）。source_module_id 为 NULL 时 SourceModuleID / SourceModuleName
// 返回空字符串。
func (s *DecisionStore) GetByID(ctx context.Context, id string) (*DecisionWithSource, error) {
	result := &DecisionWithSource{}
	d := &result.Decision
	err := s.pool.QueryRow(ctx, `
SELECT d.id, d.title, d.context, d.problem, d.alternatives, d.choice, d.reason, d.impact, d.status, d.created_at,
       COALESCE(d.source_module_id::text, ''), COALESCE(m.name, '')
FROM decisions d
LEFT JOIN modules m ON m.id = d.source_module_id
WHERE d.id = $1`,
		id,
	).Scan(
		&d.ID, &d.Title, &d.Context, &d.Problem, &d.Alternatives,
		&d.Choice, &d.Reason, &d.Impact, &d.Status, &d.CreatedAt,
		&result.SourceModuleID, &result.SourceModuleName,
	)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, decisioncenter.ErrDecisionNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("get decision by id: %w", err)
	}
	return result, nil
}

// List 按筛选条件读取决策列表，返回列表项（含 link_count 与 linked_module_summary）。
//
// 实现说明（phase03-10 §5.9 计算口径）：
//   - link_count：通过子查询统计 decision_links 中已建立的 Decision -> Module 有效关联数
//   - linked_module_summary：通过聚合已关联 Module 名称（按 module_name 升序取前 3 + +N）
//     使用 array_agg + array_to_string 在 SQL 层完成摘要生成，避免 N+1 查询
//   - 无关联时 link_count = 0，linked_module_summary = ''（不返回 null）
//
// statusFilter 为空字符串时不过滤状态。queryText 走 ILIKE 模糊匹配 title。
// 结果按 created_at DESC 排序（对齐 idx_decisions_status_created_at 索引）。
func (s *DecisionStore) List(ctx context.Context, queryText, statusFilter string) ([]decisioncenter.DecisionListItem, error) {
	rows, err := s.pool.Query(ctx, `
SELECT
    d.id,
    d.title,
    d.status,
    d.created_at,
    COUNT(dl.id)::int AS link_count,
    COALESCE(
        (SELECT string_agg(m.name, ', ' ORDER BY m.name)
         FROM (
             SELECT m.name
             FROM decision_links dl2
             JOIN modules m ON m.id = dl2.module_id
             WHERE dl2.decision_id = d.id
             ORDER BY m.name ASC
             LIMIT 3
         ) m
        ) || CASE
            WHEN (SELECT COUNT(*) FROM decision_links dl3 WHERE dl3.decision_id = d.id) > 3
            THEN ' +' || ((SELECT COUNT(*) FROM decision_links dl4 WHERE dl4.decision_id = d.id) - 3)::text
            ELSE ''
        END,
        ''
    ) AS linked_module_summary
FROM decisions d
LEFT JOIN decision_links dl ON dl.decision_id = d.id
WHERE ($1 = '' OR d.status = $1)
  AND ($2 = '' OR d.title ILIKE '%' || $2 || '%')
GROUP BY d.id, d.title, d.status, d.created_at
ORDER BY d.created_at DESC`,
		statusFilter, queryText,
	)
	if err != nil {
		return nil, fmt.Errorf("list decisions: %w", err)
	}
	defer rows.Close()

	items := make([]decisioncenter.DecisionListItem, 0)
	for rows.Next() {
		var item decisioncenter.DecisionListItem
		if err := rows.Scan(
			&item.ID, &item.Title, &item.Status, &item.CreatedAt,
			&item.LinkCount, &item.LinkedModuleSummary,
		); err != nil {
			return nil, fmt.Errorf("scan decision list item: %w", err)
		}
		items = append(items, item)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iter decision list rows: %w", err)
	}
	return items, nil
}

// Exists 校验 decision 是否存在（关联写入前提校验）。
func (s *DecisionStore) Exists(ctx context.Context, decisionID string) (bool, error) {
	var exists bool
	err := s.pool.QueryRow(ctx,
		`SELECT EXISTS(SELECT 1 FROM decisions WHERE id = $1)`,
		decisionID,
	).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("check decision exists: %w", err)
	}
	return exists, nil
}

// UpdateStatus 更新决策状态并返回更新后的完整对象。
//
// 只做持久化，不做业务校验（状态迁移合法性由 service 层承接）。
// 未找到决策时返回 decisioncenter.ErrDecisionNotFound。
func (s *DecisionStore) UpdateStatus(ctx context.Context, decisionID string, status decisioncenter.DecisionStatus) (*decisioncenter.Decision, error) {
	d := &decisioncenter.Decision{}
	err := s.pool.QueryRow(ctx, `
	UPDATE decisions
	SET status = $1
	WHERE id = $2
	RETURNING id, title, context, problem, alternatives, choice, reason, impact, status, created_at`,
		string(status), decisionID,
	).Scan(
		&d.ID, &d.Title, &d.Context, &d.Problem, &d.Alternatives,
		&d.Choice, &d.Reason, &d.Impact, &d.Status, &d.CreatedAt,
	)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, decisioncenter.ErrDecisionNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("update decision status: %w", err)
	}
	return d, nil
}
