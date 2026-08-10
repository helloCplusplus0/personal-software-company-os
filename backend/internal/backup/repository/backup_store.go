// Package repository — Backup 持久化层。
//
// 承接 instance_backups 表的读写，对齐 phase06-14 spec §"Backup 主线必须实现 read / verify 子路径与三类失败语义"。
// service 层不直接写 SQL，跨表读写由 repository 层承接。
//
// 文件落点：backend/internal/backup/repository/backup_store.go
package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/psco/backend/internal/backup"
)

// BackupStore 承接 instance_backups 表的读写。
type BackupStore struct {
	pool *pgxpool.Pool
}

// NewBackupStore 构造 BackupStore。
func NewBackupStore(pool *pgxpool.Pool) *BackupStore {
	return &BackupStore{pool: pool}
}

// BackupRecord instance_backups 表的记录模型。
type BackupRecord struct {
	ID                  string
	CreatedAt           time.Time
	ManifestJSON        json.RawMessage
	AssetCoverageJSON   json.RawMessage
	SchemaVersion       string
	InstanceVersion     string
	VerifiedStatus      backup.BackupVerifiedStatus
	VerifyFailureCode   *backup.VerifyFailureCode
	BackupPayload       json.RawMessage
}

// CreateLatest 写入一条新的备份快照记录。
// 初始写入时 verified_status 必须为 unverified（phase06-14 spec §"CreateInstanceBackup 持久化语义"）。
// created_at 由数据库默认值（NOW()）承接，返回写入后的完整记录。
func (s *BackupStore) CreateLatest(ctx context.Context, manifestJSON, assetCoverageJSON json.RawMessage, schemaVersion, instanceVersion string, backupPayload json.RawMessage) (*BackupRecord, error) {
	rec := &BackupRecord{
		ManifestJSON:      manifestJSON,
		AssetCoverageJSON: assetCoverageJSON,
		SchemaVersion:     schemaVersion,
		InstanceVersion:   instanceVersion,
		VerifiedStatus:    backup.BackupVerifiedStatusUnverified,
		BackupPayload:     backupPayload,
	}

	err := s.pool.QueryRow(ctx, `
		INSERT INTO instance_backups (manifest_json, asset_coverage_json, schema_version, instance_version, verified_status, backup_payload_json)
		VALUES ($1, $2, $3, $4, 'unverified', $5)
		RETURNING id, created_at`,
		manifestJSON, assetCoverageJSON, schemaVersion, instanceVersion, backupPayload,
	).Scan(&rec.ID, &rec.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("insert instance_backups: %w", err)
	}

	return rec, nil
}

// UpdateVerifiedStatus 更新最新一条备份记录的校验状态与失败原因代码。
// 用于 read / verify 子路径校验后回写校验结果。
func (s *BackupStore) UpdateVerifiedStatus(ctx context.Context, id string, status backup.BackupVerifiedStatus, failureCode *backup.VerifyFailureCode) error {
	var fc any
	if failureCode != nil {
		fc = string(*failureCode)
	}

	_, err := s.pool.Exec(ctx, `
		UPDATE instance_backups SET verified_status = $1, verify_failure_code = $2 WHERE id = $3`,
		string(status), fc, id)
	if err != nil {
		return fmt.Errorf("update instance_backups verified_status: %w", err)
	}
	return nil
}

// ReadLatest 读取最新一条备份快照记录。
// 若无历史记录，返回 nil + nil error（由 service 层判断为无备份状态）。
func (s *BackupStore) ReadLatest(ctx context.Context) (*BackupRecord, error) {
	rec := &BackupRecord{}
	var verifiedStatus string
	var verifyFailureCode *string

	err := s.pool.QueryRow(ctx, `
		SELECT id, created_at, manifest_json, asset_coverage_json, schema_version, instance_version, verified_status, verify_failure_code, backup_payload_json
		FROM instance_backups
		ORDER BY created_at DESC
		LIMIT 1`).Scan(
		&rec.ID, &rec.CreatedAt, &rec.ManifestJSON, &rec.AssetCoverageJSON,
		&rec.SchemaVersion, &rec.InstanceVersion, &verifiedStatus, &verifyFailureCode, &rec.BackupPayload)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("read latest instance_backups: %w", err)
	}

	rec.VerifiedStatus = backup.BackupVerifiedStatus(verifiedStatus)
	if verifyFailureCode != nil {
		fc := backup.VerifyFailureCode(*verifyFailureCode)
		rec.VerifyFailureCode = &fc
	}

	return rec, nil
}

// DeleteAll 清空 instance_backups 表的所有记录。
// 用于 reset_phase06_acceptance.sh 的清空阶段。
func (s *BackupStore) DeleteAll(ctx context.Context) error {
	if _, err := s.pool.Exec(ctx, `DELETE FROM instance_backups`); err != nil {
		return fmt.Errorf("delete instance_backups: %w", err)
	}
	return nil
}
