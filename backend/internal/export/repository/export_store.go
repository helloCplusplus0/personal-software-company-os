// Package repository — Export 持久化层。
//
// 承接 instance_exports 表的读写，对齐 phase06-14 spec §"Export 主线必须装配 9 类核心资产并形成可重复读取的快照"。
// service 层不直接写 SQL，跨表读写由 repository 层承接。
//
// 文件落点：backend/internal/export/repository/export_store.go
package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/psco/backend/internal/export"
)

// ExportStore 承接 instance_exports 表的读写。
type ExportStore struct {
	pool *pgxpool.Pool
}

// NewExportStore 构造 ExportStore。
func NewExportStore(pool *pgxpool.Pool) *ExportStore {
	return &ExportStore{pool: pool}
}

// ExportRecord instance_exports 表的记录模型。
type ExportRecord struct {
	ID                string
	CreatedAt         time.Time
	ResultStatus      export.ExportResultStatus
	ResultSummary     string
	AssetScopeJSON    json.RawMessage
	ArtifactPayload   json.RawMessage
}

// CreateLatest 写入一条新的导出快照记录。
// created_at 由数据库默认值（NOW()）承接，返回写入后的完整记录。
func (s *ExportStore) CreateLatest(ctx context.Context, resultStatus export.ExportResultStatus, resultSummary string, assetScopeJSON, artifactPayload json.RawMessage) (*ExportRecord, error) {
	rec := &ExportRecord{
		ResultStatus:   resultStatus,
		ResultSummary:  resultSummary,
		AssetScopeJSON: assetScopeJSON,
		ArtifactPayload: artifactPayload,
	}

	err := s.pool.QueryRow(ctx, `
		INSERT INTO instance_exports (result_status, result_summary, asset_scope_json, artifact_payload_json)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at`,
		string(resultStatus), resultSummary, assetScopeJSON, artifactPayload,
	).Scan(&rec.ID, &rec.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("insert instance_exports: %w", err)
	}

	return rec, nil
}

// ReadLatest 读取最新一条导出快照记录。
// 若无历史记录，返回 nil + nil error（由 service 层判断为预览态）。
func (s *ExportStore) ReadLatest(ctx context.Context) (*ExportRecord, error) {
	rec := &ExportRecord{}
	var resultStatus string
	err := s.pool.QueryRow(ctx, `
		SELECT id, created_at, result_status, result_summary, asset_scope_json, artifact_payload_json
		FROM instance_exports
		ORDER BY created_at DESC
		LIMIT 1`).Scan(
		&rec.ID, &rec.CreatedAt, &resultStatus, &rec.ResultSummary, &rec.AssetScopeJSON, &rec.ArtifactPayload)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("read latest instance_exports: %w", err)
	}
	rec.ResultStatus = export.ExportResultStatus(resultStatus)
	return rec, nil
}

// DeleteAll 清空 instance_exports 表的所有记录。
// 用于 reset_phase06_acceptance.sh 的清空阶段。
func (s *ExportStore) DeleteAll(ctx context.Context) error {
	if _, err := s.pool.Exec(ctx, `DELETE FROM instance_exports`); err != nil {
		return fmt.Errorf("delete instance_exports: %w", err)
	}
	return nil
}
