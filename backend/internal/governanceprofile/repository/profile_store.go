// Package repository — governanceprofile 数据访问层。
//
// 本文件承接治理画像主记录与两组 bindings 的 PostgreSQL 持久化：
//   - SaveProfile：单一事务边界内保存主记录 + 两组 bindings（整体成功或整体失败）
//   - ReadProfile：聚合读取主记录 + 两组 bindings
//
// 事务语义对齐 phase13-08 写路径冻结：
//   - 主记录 UPSERT：INSERT 分支以根级上游冻结投影初始化 read-only 字段；
//     ON CONFLICT UPDATE 分支继续用根级上游冻结投影刷新 read-only 列，
//     但这些列仍不接受用户输入改写
//   - bindings 全量替换：DELETE + 批量 INSERT（同一事务内）
//
// 文件落点：backend/internal/governanceprofile/repository/profile_store.go
package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/psco/backend/internal/governanceprofile"
)

// ProfileStore 承接治理画像三张表的持久化操作。
type ProfileStore struct {
	pool *pgxpool.Pool
}

// NewProfileStore 构造 ProfileStore。
func NewProfileStore(pool *pgxpool.Pool) *ProfileStore {
	return &ProfileStore{pool: pool}
}

// SaveProfile 在单一事务边界内保存治理画像并返回保存后的完整聚合。
//
// 事务流程：
//  1. UPSERT governance_profiles
//     - INSERT：可写字段 + read-only 受控投影值 + project_profile_version 固定版本
//     - ON CONFLICT (repository_id) DO UPDATE：继续用根级受控投影刷新
//     track_type / current_phase_*，同时更新可写字段与 updated_at
//  2. DELETE + 批量 INSERT canonical_root_file_bindings（全量替换）
//  3. DELETE + 批量 INSERT global_asset_bindings（全量替换）
//  4. COMMIT 后聚合读取返回
//
// 任一步失败即整体回滚，不写入半套治理画像状态。
func (s *ProfileStore) SaveProfile(ctx context.Context, input governanceprofile.UpdateGovernanceProfileInput) (*governanceprofile.GovernanceProfileReadResult, error) {
	tx, err := s.pool.Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("%w: begin tx: %w", governanceprofile.ErrGovernanceProfileSaveFailed, err)
	}
	defer tx.Rollback(ctx) // commit 后 rollback 为 no-op

	// 1. 主记录 UPSERT。
	// INSERT 分支用根级冻结投影初始化 read-only 字段；
	// UPDATE 分支继续使用根级冻结投影刷新 read-only 字段；
	// 这些值不是来自维护写路径输入，而是来自服务端正式冻结常量。
	const upsertSQL = `
		INSERT INTO governance_profiles (
			repository_id, project_profile_version, track_type, template_source,
			docs_workflow_layout, current_phase_name, current_phase_ref, current_phase_status
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		ON CONFLICT (repository_id) DO UPDATE SET
			project_profile_version = EXCLUDED.project_profile_version,
                        track_type              = EXCLUDED.track_type,
			template_source         = EXCLUDED.template_source,
			docs_workflow_layout    = EXCLUDED.docs_workflow_layout,
                        current_phase_name      = EXCLUDED.current_phase_name,
                        current_phase_ref       = EXCLUDED.current_phase_ref,
                        current_phase_status    = EXCLUDED.current_phase_status,
			updated_at              = NOW()
		RETURNING repository_id::text, project_profile_version, track_type, template_source,
		          docs_workflow_layout, current_phase_name, current_phase_ref, current_phase_status,
		          created_at, updated_at
	`
	record := governanceprofile.GovernanceProfileRecord{}
	err = tx.QueryRow(ctx, upsertSQL,
		input.RepositoryID,
		governanceprofile.CurrentProfileVersion,
		string(governanceprofile.RootFrozenTrackType),
		input.TemplateSource,
		input.DocsWorkflowLayout,
		governanceprofile.RootFrozenCurrentPhaseName,
		governanceprofile.RootFrozenCurrentPhaseRef,
		string(governanceprofile.RootFrozenCurrentPhaseStatus),
	).Scan(
		&record.RepositoryID,
		&record.ProjectProfileVersion,
		&record.TrackType,
		&record.TemplateSource,
		&record.DocsWorkflowLayout,
		&record.CurrentPhaseName,
		&record.CurrentPhaseRef,
		&record.CurrentPhaseStatus,
		&record.CreatedAt,
		&record.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("%w: upsert governance profile: %w", governanceprofile.ErrGovernanceProfileSaveFailed, err)
	}

	// 2. canonical 根级文件 bindings 全量替换。
	if err := replaceCanonicalRootFileBindings(ctx, tx, input.RepositoryID, input.CanonicalRootFiles); err != nil {
		return nil, err
	}

	// 3. 全局规范资产 bindings 全量替换。
	if err := replaceGlobalAssetBindings(ctx, tx, input.RepositoryID, input.GlobalAssetBindings); err != nil {
		return nil, err
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, fmt.Errorf("%w: commit tx: %w", governanceprofile.ErrGovernanceProfileSaveFailed, err)
	}

	// 4. COMMIT 后聚合读取（复用同一读取主线，保证保存后回读一致）。
	return s.ReadProfile(ctx, input.RepositoryID)
}

// ReadProfile 聚合读取治理画像主记录与两组 bindings。
// 主记录不存在时返回 ErrGovernanceProfileNotFound。
func (s *ProfileStore) ReadProfile(ctx context.Context, repositoryID string) (*governanceprofile.GovernanceProfileReadResult, error) {
	result := &governanceprofile.GovernanceProfileReadResult{
		CanonicalRootFiles:  []governanceprofile.CanonicalRootFileBinding{},
		GlobalAssetBindings: []governanceprofile.GlobalAssetBinding{},
	}

	// 1. 主记录
	err := s.pool.QueryRow(ctx, `
		SELECT repository_id::text, project_profile_version, track_type, template_source,
		       docs_workflow_layout, current_phase_name, current_phase_ref, current_phase_status,
		       created_at, updated_at
		FROM governance_profiles
		WHERE repository_id = $1
	`, repositoryID).Scan(
		&result.Record.RepositoryID,
		&result.Record.ProjectProfileVersion,
		&result.Record.TrackType,
		&result.Record.TemplateSource,
		&result.Record.DocsWorkflowLayout,
		&result.Record.CurrentPhaseName,
		&result.Record.CurrentPhaseRef,
		&result.Record.CurrentPhaseStatus,
		&result.Record.CreatedAt,
		&result.Record.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, governanceprofile.ErrGovernanceProfileNotFound
		}
		return nil, fmt.Errorf("%w: read governance profile: %w", governanceprofile.ErrGovernanceProfileReadFailed, err)
	}

	// 2. canonical 根级文件 bindings
	rootFiles, err := s.readCanonicalRootFileBindings(ctx, repositoryID)
	if err != nil {
		return nil, err
	}
	result.CanonicalRootFiles = rootFiles

	// 3. 全局规范资产 bindings
	assetBindings, err := s.readGlobalAssetBindings(ctx, repositoryID)
	if err != nil {
		return nil, err
	}
	result.GlobalAssetBindings = assetBindings

	return result, nil
}

// readCanonicalRootFileBindings 读取 canonical 根级文件 bindings（按 file_name 排序，稳定回读）。
func (s *ProfileStore) readCanonicalRootFileBindings(ctx context.Context, repositoryID string) ([]governanceprofile.CanonicalRootFileBinding, error) {
	rows, err := s.pool.Query(ctx, `
		SELECT b.file_name, b.role, b.required
		FROM governance_canonical_root_file_bindings b
		JOIN governance_profiles p ON p.id = b.governance_profile_id
		WHERE p.repository_id = $1
		ORDER BY b.file_name ASC
	`, repositoryID)
	if err != nil {
		return nil, fmt.Errorf("%w: read canonical root file bindings: %w", governanceprofile.ErrGovernanceProfileReadFailed, err)
	}
	defer rows.Close()

	bindings := []governanceprofile.CanonicalRootFileBinding{}
	for rows.Next() {
		var b governanceprofile.CanonicalRootFileBinding
		if err := rows.Scan(&b.FileName, &b.Role, &b.Required); err != nil {
			return nil, fmt.Errorf("%w: scan canonical root file binding: %w", governanceprofile.ErrGovernanceProfileReadFailed, err)
		}
		bindings = append(bindings, b)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("%w: iterate canonical root file bindings: %w", governanceprofile.ErrGovernanceProfileReadFailed, err)
	}
	return bindings, nil
}

// readGlobalAssetBindings 读取全局规范资产 bindings（按 name 排序，稳定回读）。
func (s *ProfileStore) readGlobalAssetBindings(ctx context.Context, repositoryID string) ([]governanceprofile.GlobalAssetBinding, error) {
	rows, err := s.pool.Query(ctx, `
		SELECT b.name, b.kind, b.entry_ref, b.role, b.structured_summary
		FROM governance_global_asset_bindings b
		JOIN governance_profiles p ON p.id = b.governance_profile_id
		WHERE p.repository_id = $1
		ORDER BY b.name ASC
	`, repositoryID)
	if err != nil {
		return nil, fmt.Errorf("%w: read global asset bindings: %w", governanceprofile.ErrGovernanceProfileReadFailed, err)
	}
	defer rows.Close()

	bindings := []governanceprofile.GlobalAssetBinding{}
	for rows.Next() {
		var b governanceprofile.GlobalAssetBinding
		if err := rows.Scan(&b.Name, &b.Kind, &b.EntryRef, &b.Role, &b.StructuredSummary); err != nil {
			return nil, fmt.Errorf("%w: scan global asset binding: %w", governanceprofile.ErrGovernanceProfileReadFailed, err)
		}
		b.MarkdownResolvable = governanceprofile.MarkdownResolvableForGlobalAsset(b.Name)
		bindings = append(bindings, b)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("%w: iterate global asset bindings: %w", governanceprofile.ErrGovernanceProfileReadFailed, err)
	}
	return bindings, nil
}

// replaceCanonicalRootFileBindings 在事务内全量替换 canonical 根级文件 bindings。
func replaceCanonicalRootFileBindings(ctx context.Context, tx pgx.Tx, repositoryID string, bindings []governanceprofile.CanonicalRootFileBinding) error {
	if _, err := tx.Exec(ctx, `
		DELETE FROM governance_canonical_root_file_bindings
		WHERE governance_profile_id = (SELECT id FROM governance_profiles WHERE repository_id = $1)
	`, repositoryID); err != nil {
		return fmt.Errorf("%w: delete canonical root file bindings: %w", governanceprofile.ErrGovernanceProfileSaveFailed, err)
	}

	for _, b := range bindings {
		if _, err := tx.Exec(ctx, `
			INSERT INTO governance_canonical_root_file_bindings (
				governance_profile_id, file_name, role, required
			) VALUES (
				(SELECT id FROM governance_profiles WHERE repository_id = $1), $2, $3, $4
			)
		`, repositoryID, b.FileName, b.Role, b.Required); err != nil {
			return fmt.Errorf("%w: insert canonical root file binding %s: %w", governanceprofile.ErrGovernanceProfileSaveFailed, b.FileName, err)
		}
	}
	return nil
}

// replaceGlobalAssetBindings 在事务内全量替换全局规范资产 bindings。
func replaceGlobalAssetBindings(ctx context.Context, tx pgx.Tx, repositoryID string, bindings []governanceprofile.GlobalAssetBinding) error {
	if _, err := tx.Exec(ctx, `
		DELETE FROM governance_global_asset_bindings
		WHERE governance_profile_id = (SELECT id FROM governance_profiles WHERE repository_id = $1)
	`, repositoryID); err != nil {
		return fmt.Errorf("%w: delete global asset bindings: %w", governanceprofile.ErrGovernanceProfileSaveFailed, err)
	}

	for _, b := range bindings {
		if _, err := tx.Exec(ctx, `
			INSERT INTO governance_global_asset_bindings (
				governance_profile_id, name, kind, entry_ref, role, structured_summary
			) VALUES (
				(SELECT id FROM governance_profiles WHERE repository_id = $1), $2, $3, $4, $5, $6
			)
		`, repositoryID, b.Name, b.Kind, b.EntryRef, b.Role, b.StructuredSummary); err != nil {
			return fmt.Errorf("%w: insert global asset binding %s: %w", governanceprofile.ErrGovernanceProfileSaveFailed, b.Name, err)
		}
	}
	return nil
}
