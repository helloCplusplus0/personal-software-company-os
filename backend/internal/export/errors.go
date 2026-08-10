// Package export — 业务层错误定义。
//
// 这些哨兵错误由 service / candidate / repository 层返回，由 handler 层映射为 HTTP 状态码。
// 错误语义对齐 phase06-14 spec §"Export 主线必须装配 9 类核心资产并形成可重复读取的快照"。
package export

import "errors"

// 业务错误哨兵值。
var (
	// ErrAssetReadFailed 9 类核心资产装配失败。
	// 当任一 asset reader 失败时返回。
	ErrAssetReadFailed = errors.New("export asset read failed")

	// ErrExportPersistFailed 导出结果持久化失败。
	// 当 instance_exports 写入失败时返回。
	ErrExportPersistFailed = errors.New("export persist failed")

	// ErrExportSnapshotReadFailed 导出快照读取失败。
	// 当 instance_exports 读取失败时返回。
	ErrExportSnapshotReadFailed = errors.New("export snapshot read failed")
)
