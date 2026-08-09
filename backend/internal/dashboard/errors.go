// Package dashboard — 业务层错误定义。
//
// 这些哨兵错误由 service / candidate 层返回，由 handler 层映射为 HTTP 状态码。
// 错误语义对齐 phase05-04 spec 的错误语义前提与 phase05-10 正式规格正文。
package dashboard

import "errors"

// 业务错误哨兵值。
var (
	// ErrOverviewReadFailed DashboardOverviewRead 整页失败。
	// 当任一计数 reader 失败时返回。
	ErrOverviewReadFailed = errors.New("dashboard overview read failed")

	// ErrFeedbackSignalReadFailed FeedbackSignalRead 局部失败。
	// 当 pending_decision_signals 或 product_asset_coverage reader 失败时返回。
	ErrFeedbackSignalReadFailed = errors.New("dashboard feedback signal read failed")

	// ErrRecentActivityReadFailed RecentActivityRead 局部失败。
	// 当任一活动 reader 失败时返回。
	ErrRecentActivityReadFailed = errors.New("dashboard recent activity read failed")
)
