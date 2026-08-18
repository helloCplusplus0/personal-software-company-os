// Package repository — standard 数据访问层。
//
// 本文件承接 Standard 三张表（standards / standard_revisions / standard_bindings）
// 的 PostgreSQL 持久化：
//   - CreateStandard / ListStandards / GetStandardByID / UpdateStandard / DeleteStandard：主表
//   - InsertBinding / DeleteBinding / ListBindingsByStandardID：绑定表
//   - ListRevisions：演进留痕表
//   - ListStandardsByRepository：brief 反查（按 repository 绑定关系）
//
// 编解码语义（对齐 0011 迁移列集）：
//   - directory_tree jsonb：领域树 ↔ []byte（json.Marshal / Unmarshal）
//   - status / target_type / role：受控小写字符串直存
//   - description / note 可空 TEXT：读出 NULL 归空串
//
// 本层只做 SQL 与编解码，无业务判断；业务校验收敛在 service 层。
//
// 文件落点：backend/internal/standard/repository/standard_store.go
package repository

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/psco/backend/internal/standard"
)

// isUniqueViolation 判断是否为 PostgreSQL 唯一约束冲突（SQLSTATE 23505）。
func isUniqueViolation(err error) bool {
	var pgErr *pgconn.PgError
	return errors.As(err, &pgErr) && pgErr.Code == "23505"
}

// StandardStore 承接 Standard 三张表的持久化操作。
type StandardStore struct {
	pool *pgxpool.Pool
}

// NewStandardStore 构造 StandardStore。
func NewStandardStore(pool *pgxpool.Pool) *StandardStore {
	return &StandardStore{pool: pool}
}

// CreateStandard 插入规范主记录并返回完整读取结果。
// UNIQUE(name) 冲突映射为 ErrInvalidInput（name already exists）。
func (s *StandardStore) CreateStandard(ctx context.Context, input standard.CreateStandardInput) (*standard.StandardReadResult, error) {
	treeBytes, err := json.Marshal(input.DirectoryTree)
	if err != nil {
		return nil, fmt.Errorf("%w: encode directory_tree: %w", standard.ErrStandardSaveFailed, err)
	}

	result, err := scanStandard(s.pool.QueryRow(ctx, `
		INSERT INTO standards (name, description, status, directory_tree)
		VALUES ($1, $2, $3, $4)
		RETURNING id::text, name, description, status, directory_tree, created_at, updated_at
	`, input.Name, input.Description, string(input.Status), treeBytes))
	if err != nil {
		if isUniqueViolation(err) {
			return nil, fmt.Errorf("%w: standard name %q already exists", standard.ErrInvalidInput, input.Name)
		}
		return nil, fmt.Errorf("%w: insert standard: %w", standard.ErrStandardSaveFailed, err)
	}
	return result, nil
}

// ListStandards 读取全量规范列表（按 updated_at DESC，不分页）。
// 空结果返回空切片非 nil。
func (s *StandardStore) ListStandards(ctx context.Context) ([]standard.StandardReadResult, error) {
	rows, err := s.pool.Query(ctx, `
		SELECT id::text, name, description, status, directory_tree, created_at, updated_at
		FROM standards
		ORDER BY updated_at DESC
	`)
	if err != nil {
		return nil, fmt.Errorf("%w: list standards: %w", standard.ErrStandardReadFailed, err)
	}
	defer rows.Close()

	results := []standard.StandardReadResult{}
	for rows.Next() {
		result, err := scanStandard(rows)
		if err != nil {
			return nil, fmt.Errorf("%w: scan standard: %w", standard.ErrStandardReadFailed, err)
		}
		results = append(results, *result)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("%w: iterate standards: %w", standard.ErrStandardReadFailed, err)
	}
	return results, nil
}

// GetStandardByID 按 id 读取单条规范。未找到时返回 standard.ErrStandardNotFound。
func (s *StandardStore) GetStandardByID(ctx context.Context, id string) (*standard.StandardReadResult, error) {
	result, err := scanStandard(s.pool.QueryRow(ctx, `
		SELECT id::text, name, description, status, directory_tree, created_at, updated_at
		FROM standards
		WHERE id = $1
	`, id))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, standard.ErrStandardNotFound
		}
		return nil, fmt.Errorf("%w: get standard: %w", standard.ErrStandardReadFailed, err)
	}
	return result, nil
}

// ListBindingsByStandardID 读取规范的绑定集合（按 created_at ASC，绑定管理区稳定回读）。
func (s *StandardStore) ListBindingsByStandardID(ctx context.Context, standardID string) ([]standard.StandardBindingReadResult, error) {
	rows, err := s.pool.Query(ctx, `
		SELECT id::text, standard_id::text, target_type, target_id::text, role, note, created_at
		FROM standard_bindings
		WHERE standard_id = $1
		ORDER BY created_at ASC
	`, standardID)
	if err != nil {
		return nil, fmt.Errorf("%w: list standard bindings: %w", standard.ErrStandardReadFailed, err)
	}
	defer rows.Close()

	bindings := []standard.StandardBindingReadResult{}
	for rows.Next() {
		binding, err := scanBinding(rows)
		if err != nil {
			return nil, fmt.Errorf("%w: scan standard binding: %w", standard.ErrStandardReadFailed, err)
		}
		bindings = append(bindings, *binding)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("%w: iterate standard bindings: %w", standard.ErrStandardReadFailed, err)
	}
	return bindings, nil
}

// UpdateStandard 在单一事务边界内执行整树替换 + revision 追加，返回更新后的完整记录。
//
// 事务流程：
//  1. UPDATE standards：name / description / status 为 COALESCE 部分变更
//     （nil 不变更，非 nil 含空串写入），directory_tree 整树必写
//  2. INSERT standard_revisions：追加人工一句话留痕
//
// 任一步失败即整体回滚，不写入半套状态。
func (s *StandardStore) UpdateStandard(ctx context.Context, input standard.UpdateStandardInput) (*standard.StandardReadResult, error) {
	treeBytes, err := json.Marshal(input.DirectoryTree)
	if err != nil {
		return nil, fmt.Errorf("%w: encode directory_tree: %w", standard.ErrStandardSaveFailed, err)
	}

	var statusPtr *string
	if input.Status != nil {
		v := string(*input.Status)
		statusPtr = &v
	}

	tx, err := s.pool.Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("%w: begin tx: %w", standard.ErrStandardSaveFailed, err)
	}
	defer tx.Rollback(ctx) // commit 后 rollback 为 no-op

	// 1. 主记录整树替换（COALESCE 承接 optional 部分变更语义）
	result, err := scanStandard(tx.QueryRow(ctx, `
		UPDATE standards SET
			name           = COALESCE($2, name),
			description    = COALESCE($3, description),
			status         = COALESCE($4, status),
			directory_tree = $5,
			updated_at     = NOW()
		WHERE id = $1
		RETURNING id::text, name, description, status, directory_tree, created_at, updated_at
	`, input.StandardID, input.Name, input.Description, statusPtr, treeBytes))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, standard.ErrStandardNotFound
		}
		if isUniqueViolation(err) {
			newName := ""
			if input.Name != nil {
				newName = *input.Name
			}
			return nil, fmt.Errorf("%w: standard name %q already exists", standard.ErrInvalidInput, newName)
		}
		return nil, fmt.Errorf("%w: update standard: %w", standard.ErrStandardSaveFailed, err)
	}

	// 2. revision 追加
	if _, err := tx.Exec(ctx, `
		INSERT INTO standard_revisions (standard_id, change_summary)
		VALUES ($1, $2)
	`, input.StandardID, input.ChangeSummary); err != nil {
		return nil, fmt.Errorf("%w: insert standard revision: %w", standard.ErrStandardSaveFailed, err)
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, fmt.Errorf("%w: commit tx: %w", standard.ErrStandardSaveFailed, err)
	}
	return result, nil
}

// DeleteStandard 删除规范（CASCADE 连带删除 bindings 与 revisions）。
// 未找到时返回 standard.ErrStandardNotFound。
func (s *StandardStore) DeleteStandard(ctx context.Context, id string) error {
	tag, err := s.pool.Exec(ctx, `
		DELETE FROM standards WHERE id = $1
	`, id)
	if err != nil {
		return fmt.Errorf("%w: delete standard: %w", standard.ErrStandardSaveFailed, err)
	}
	if tag.RowsAffected() == 0 {
		return standard.ErrStandardNotFound
	}
	return nil
}

// InsertBinding 插入绑定关系并返回完整读取结果。
// 四元组 UNIQUE 冲突映射为 ErrInvalidInput（错误信息含 "already bound"）。
func (s *StandardStore) InsertBinding(ctx context.Context, input standard.BindStandardInput) (*standard.StandardBindingReadResult, error) {
	binding, err := scanBinding(s.pool.QueryRow(ctx, `
		INSERT INTO standard_bindings (standard_id, target_type, target_id, role, note)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id::text, standard_id::text, target_type, target_id::text, role, note, created_at
	`, input.StandardID, string(input.TargetType), input.TargetID, string(input.Role), input.Note))
	if err != nil {
		if isUniqueViolation(err) {
			return nil, fmt.Errorf("%w: target already bound (standard=%s, target_type=%s, target_id=%s, role=%s)",
				standard.ErrInvalidInput, input.StandardID, input.TargetType, input.TargetID, input.Role)
		}
		return nil, fmt.Errorf("%w: insert standard binding: %w", standard.ErrStandardSaveFailed, err)
	}
	return binding, nil
}

// DeleteBinding 按四元组删除绑定。未找到时返回 standard.ErrBindingNotFound。
func (s *StandardStore) DeleteBinding(ctx context.Context, input standard.UnbindStandardInput) error {
	tag, err := s.pool.Exec(ctx, `
		DELETE FROM standard_bindings
		WHERE standard_id = $1 AND target_type = $2 AND target_id = $3 AND role = $4
	`, input.StandardID, string(input.TargetType), input.TargetID, string(input.Role))
	if err != nil {
		return fmt.Errorf("%w: delete standard binding: %w", standard.ErrStandardSaveFailed, err)
	}
	if tag.RowsAffected() == 0 {
		return standard.ErrBindingNotFound
	}
	return nil
}

// ListRevisions 读取规范演进留痕（按 created_at DESC，不分页）。
func (s *StandardStore) ListRevisions(ctx context.Context, standardID string) ([]standard.StandardRevisionReadResult, error) {
	rows, err := s.pool.Query(ctx, `
		SELECT id::text, standard_id::text, change_summary, created_at
		FROM standard_revisions
		WHERE standard_id = $1
		ORDER BY created_at DESC
	`, standardID)
	if err != nil {
		return nil, fmt.Errorf("%w: list standard revisions: %w", standard.ErrStandardReadFailed, err)
	}
	defer rows.Close()

	revisions := []standard.StandardRevisionReadResult{}
	for rows.Next() {
		var r standard.StandardRevisionReadResult
		if err := rows.Scan(&r.ID, &r.StandardID, &r.ChangeSummary, &r.CreatedAt); err != nil {
			return nil, fmt.Errorf("%w: scan standard revision: %w", standard.ErrStandardReadFailed, err)
		}
		revisions = append(revisions, r)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("%w: iterate standard revisions: %w", standard.ErrStandardReadFailed, err)
	}
	return revisions, nil
}

// ListStandardsByRepository 按 repository 绑定关系反查关联规范（按 updated_at DESC）。
// 无关联返回空切片非错误；失败 wrap ErrStandardReadFailed。
func (s *StandardStore) ListStandardsByRepository(ctx context.Context, repositoryID string) ([]standard.StandardReadResult, error) {
	rows, err := s.pool.Query(ctx, `
		SELECT DISTINCT s.id::text, s.name, s.description, s.status, s.directory_tree, s.created_at, s.updated_at
		FROM standards s
		JOIN standard_bindings b ON b.standard_id = s.id
		WHERE b.target_type = 'repository' AND b.target_id = $1
		ORDER BY s.updated_at DESC
	`, repositoryID)
	if err != nil {
		return nil, fmt.Errorf("%w: list standards by repository: %w", standard.ErrStandardReadFailed, err)
	}
	defer rows.Close()

	results := []standard.StandardReadResult{}
	for rows.Next() {
		result, err := scanStandard(rows)
		if err != nil {
			return nil, fmt.Errorf("%w: scan standard: %w", standard.ErrStandardReadFailed, err)
		}
		results = append(results, *result)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("%w: iterate standards: %w", standard.ErrStandardReadFailed, err)
	}
	return results, nil
}

// rowScanner 统一 pgx.Row / pgx.Rows 的 Scan 承接位。
type rowScanner interface {
	Scan(dest ...any) error
}

// scanStandard 将 standards 表单行解码为 StandardReadResult（jsonb 树反序列化 + 可空 TEXT 归空串）。
func scanStandard(row rowScanner) (*standard.StandardReadResult, error) {
	var result standard.StandardReadResult
	var description *string
	var treeBytes []byte
	if err := row.Scan(&result.ID, &result.Name, &description, &result.Status, &treeBytes, &result.CreatedAt, &result.UpdatedAt); err != nil {
		return nil, err
	}
	if description != nil {
		result.Description = *description
	}
	var tree standard.DirectoryTreeNode
	if err := json.Unmarshal(treeBytes, &tree); err != nil {
		return nil, fmt.Errorf("decode directory_tree: %w", err)
	}
	result.DirectoryTree = &tree
	return &result, nil
}

// scanBinding 将 standard_bindings 表单行解码为 StandardBindingReadResult（可空 note 归空串）。
func scanBinding(row rowScanner) (*standard.StandardBindingReadResult, error) {
	var result standard.StandardBindingReadResult
	var note *string
	if err := row.Scan(&result.ID, &result.StandardID, &result.TargetType, &result.TargetID, &result.Role, &note, &result.CreatedAt); err != nil {
		return nil, err
	}
	if note != nil {
		result.Note = *note
	}
	return &result, nil
}
