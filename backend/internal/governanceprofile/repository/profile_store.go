// Package repository — governanceprofile 数据访问层。
//
// phase14-09 收缩：画像写路径与两张 bindings 表
// （governance_canonical_root_file_bindings / governance_global_asset_bindings）
// 的读取已由 Standard 实体承接；本文件收敛为只读 governance_profiles 主表
// 三组字段（track_type / template_source / current_phase 三字段），
// 返回轻量核心结果 GovernanceProfileCoreReadResult。
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

// ProfileStore 承接治理画像主记录的只读访问。
type ProfileStore struct {
	pool *pgxpool.Pool
}

// NewProfileStore 构造 ProfileStore。
func NewProfileStore(pool *pgxpool.Pool) *ProfileStore {
	return &ProfileStore{pool: pool}
}

// ReadProfile 只读治理画像主表三组字段。
// 主记录不存在时返回 ErrGovernanceProfileNotFound。
func (s *ProfileStore) ReadProfile(ctx context.Context, repositoryID string) (*governanceprofile.GovernanceProfileCoreReadResult, error) {
	result := &governanceprofile.GovernanceProfileCoreReadResult{}

	err := s.pool.QueryRow(ctx, `
		SELECT track_type, template_source,
		       current_phase_name, current_phase_ref, current_phase_status
		FROM governance_profiles
		WHERE repository_id = $1
	`, repositoryID).Scan(
		&result.TrackType,
		&result.TemplateSource,
		&result.CurrentPhaseName,
		&result.CurrentPhaseRef,
		&result.CurrentPhaseStatus,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, governanceprofile.ErrGovernanceProfileNotFound
		}
		return nil, fmt.Errorf("%w: read governance profile: %w", governanceprofile.ErrGovernanceProfileReadFailed, err)
	}

	return result, nil
}
