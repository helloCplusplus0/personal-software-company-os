// Package reusesummary — 业务层错误定义。
//
// 这些哨兵错误由 service / candidate 层返回，由 handler 层映射为 HTTP 状态码。
// 错误语义对齐 phase06-14 spec §"Reuse Summary 必须通过读时聚合返回最新已提交状态"。
package reusesummary

import "errors"

// 业务错误哨兵值。
var (
	// ErrReuseSummaryReadFailed 复用感知读取失败。
	// 当任一 candidate reader 失败时返回。
	ErrReuseSummaryReadFailed = errors.New("reuse summary read failed")

	// ErrInvalidScope 作用域参数非法。
	// 当 scope 不在 dashboard / module_detail / product_detail 范围内时返回。
	ErrInvalidScope = errors.New("invalid reuse summary scope")
)
