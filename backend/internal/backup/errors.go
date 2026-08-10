// Package backup — 业务层错误定义。
//
// 这些哨兵错误由 service / candidate / repository 层返回，由 handler 层映射为 HTTP 状态码。
// 错误语义对齐 phase06-14 spec §"Backup 主线必须实现 read / verify 子路径与三类失败语义"。
package backup

import "errors"

// 业务错误哨兵值。
var (
	// ErrAssetReadFailed 9 类核心资产装配失败。
	ErrAssetReadFailed = errors.New("backup asset read failed")

	// ErrBackupPersistFailed 备份结果持久化失败。
	ErrBackupPersistFailed = errors.New("backup persist failed")

	// ErrBackupSnapshotReadFailed 备份快照读取失败。
	// 当 instance_backups 读取失败时返回。
	ErrBackupSnapshotReadFailed = errors.New("backup snapshot read failed")

	// ErrSchemaVersionReadFailed 当前 schema_migrations 最新版本读取失败。
	ErrSchemaVersionReadFailed = errors.New("schema version read failed")
)
