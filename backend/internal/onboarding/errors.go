// Package onboarding — 业务层错误定义。
//
// 这些哨兵错误由 service / candidate 层返回，由 handler 层映射为 HTTP 状态码。
// 错误语义对齐 phase06-14 spec §"GetFirstRunState 必须由 canonical 数据读时派生"。
package onboarding

import "errors"

// 业务错误哨兵值。
var (
	// ErrFirstRunStateReadFailed GetFirstRunState 读取失败。
	// 当任一 canonical 计数 reader 失败时返回。
	ErrFirstRunStateReadFailed = errors.New("onboarding first run state read failed")
)
