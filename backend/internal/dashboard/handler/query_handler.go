// Package handler — Dashboard 读组入口层。
//
// 承接 DashboardOverviewRead / FeedbackSignalRead / RecentActivityRead 三个 GET endpoint。
// 对齐 phase05-08 已冻结的 RPC → HTTP 映射矩阵与 phase05-12 spec §"Dashboard handler 必须承接三个 GET endpoint"。
//
// 文件落点：backend/internal/dashboard/handler/query_handler.go
package handler

import (
	"net/http"

	"github.com/psco/backend/internal/dashboard"
	"github.com/psco/backend/internal/dashboard/service"
)

// QueryHandler 承接 Dashboard 读组 HTTP 入口。
type QueryHandler struct {
	query *service.QueryService
}

// NewQueryHandler 构造 QueryHandler。
func NewQueryHandler(q *service.QueryService) *QueryHandler {
	return &QueryHandler{query: q}
}

// GetOverview GET /api/dashboard/overview
//
// 承接 DashboardOverviewRead（phase05-07 读组）。
// handler 显式组装空 Proto request 语义（直接调用 service 方法），不绕过 service 合同边界。
// 响应体必须为 GetDashboardOverviewResponse envelope {"overview": {...}}，
// 与 dashboard.proto 冻结的唯一合同源对齐，不得返回裸 DashboardOverview。
// 整页失败时返回 500；空态成功时返回 200 + 零计数。
func (h *QueryHandler) GetOverview(w http.ResponseWriter, r *http.Request) {
	overview, err := h.query.ReadOverview(r.Context())
	if err != nil {
		writeError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, dashboard.DashboardOverviewReadResult{
		Overview: overview,
	})
}

// GetFeedbackSignals GET /api/dashboard/feedback-signals
//
// 承接 FeedbackSignalRead（phase05-07 读组）。
// handler 显式组装空 Proto request 语义（直接调用 service 方法），不绕过 service 合同边界。
// 局部失败时返回 500；空态成功时返回 200 + 空列表与零值计数。
func (h *QueryHandler) GetFeedbackSignals(w http.ResponseWriter, r *http.Request) {
	result, err := h.query.ReadFeedbackSignal(r.Context())
	if err != nil {
		writeError(w, err)
		return
	}
	// 确保空态返回非 nil 列表
	if result.CurrentFocusSignals == nil {
		result.CurrentFocusSignals = []dashboard.FeedbackSignal{}
	}
	if result.AssetFeedbackSummary.RepresentativeSignals == nil {
		result.AssetFeedbackSummary.RepresentativeSignals = []dashboard.FeedbackSignal{}
	}
	writeJSON(w, http.StatusOK, result)
}

// GetRecentActivities GET /api/dashboard/recent-activities
//
// 承接 RecentActivityRead（phase05-07 读组）。
// handler 显式组装空 Proto request 语义（直接调用 service 方法），不绕过 service 合同边界。
// 响应体必须为 GetRecentActivitiesResponse envelope {"activities": [...]}，
// 与 dashboard.proto 冻结的唯一合同源对齐，不得返回裸 []RecentActivityItem。
// 局部失败时返回 500；空态成功时返回 200 + 空列表。
func (h *QueryHandler) GetRecentActivities(w http.ResponseWriter, r *http.Request) {
	items, err := h.query.ReadRecentActivity(r.Context())
	if err != nil {
		writeError(w, err)
		return
	}
	if items == nil {
		items = []dashboard.RecentActivityItem{}
	}
	writeJSON(w, http.StatusOK, dashboard.RecentActivityReadResult{
		Activities: items,
	})
}
