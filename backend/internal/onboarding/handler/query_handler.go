// Package handler — Onboarding 读组入口层。
//
// 承接 GetFirstRunState 一个 GET endpoint。
// 对齐 phase06-14 spec §"Phase06 路由装配必须接入现有 chi 组合根"的路由注册矩阵：
//
//	GET /api/onboarding/state
//
// 文件落点：backend/internal/onboarding/handler/query_handler.go
package handler

import (
	"net/http"

	"github.com/psco/backend/internal/onboarding"
	"github.com/psco/backend/internal/onboarding/service"
)

// QueryHandler 承接 Onboarding 读组 HTTP 入口。
type QueryHandler struct {
	query *service.QueryService
}

// NewQueryHandler 构造 QueryHandler。
func NewQueryHandler(q *service.QueryService) *QueryHandler {
	return &QueryHandler{query: q}
}

// GetFirstRunState GET /api/onboarding/state
//
// 承接 FirstRunStateRead（phase06-14 读组）。
// handler 显式组装空 Proto request 语义（直接调用 service 方法），不绕过 service 合同边界。
// 响应体必须为 GetFirstRunStateResponse envelope {"first_run_state": {...}}，
// 与 onboarding.proto 冻结的唯一合同源对齐，不得返回裸 FirstRunState。
// 整页失败时返回 500；成功时返回 200 + 推导出的 FirstRunState。
func (h *QueryHandler) GetFirstRunState(w http.ResponseWriter, r *http.Request) {
	state, err := h.query.ReadFirstRunState(r.Context())
	if err != nil {
		writeError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, onboarding.FirstRunStateReadResult{
		FirstRunState: state,
	})
}
