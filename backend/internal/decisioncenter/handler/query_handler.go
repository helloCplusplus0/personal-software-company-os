// Package handler — 读组入口层。
//
// 承接 DecisionListRead / DecisionDetailRead / DecisionModuleCandidateRead
// （phase03-10 §10.4 读组单文件编排）。
//
// 文件落点：backend/internal/decisioncenter/handler/query_handler.go
package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/psco/backend/internal/decisioncenter"
	"github.com/psco/backend/internal/decisioncenter/service"
)

// QueryHandler 承接读组 HTTP 入口。
type QueryHandler struct {
	query *service.QueryService
}

// NewQueryHandler 构造 QueryHandler。
func NewQueryHandler(q *service.QueryService) *QueryHandler {
	return &QueryHandler{query: q}
}

// ListDecisions GET /api/decisions?queryText=...&statusFilter=...
//
// 承接 DecisionListRead（phase03-10 §6.2 读组 / §7.7 RPC→HTTP 映射）。
// queryText / statusFilter 由 URL query 参数承接，组装为 ListQuery 后进入业务层。
func (h *QueryHandler) ListDecisions(w http.ResponseWriter, r *http.Request) {
	q := decisioncenter.ListQuery{
		QueryText:    r.URL.Query().Get("queryText"),
		StatusFilter: decisioncenter.DecisionStatus(r.URL.Query().Get("statusFilter")),
	}

	items, err := h.query.ListDecisions(r.Context(), q)
	if err != nil {
		writeError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, items)
}

// GetDecisionDetail GET /api/decisions/{decisionId}
//
// 承接 DecisionDetailRead（phase03-10 §6.2 读组 / §7.7 RPC→HTTP 映射）。
// decisionId 由 URL 路径参数承接，handler 在进入业务层前显式组装为 service 调用参数，
// 不放在 JSON 请求体（phase03-10 §7.7 约束）。
func (h *QueryHandler) GetDecisionDetail(w http.ResponseWriter, r *http.Request) {
	decisionID := chi.URLParam(r, "decisionId")
	if decisionID == "" {
		writeJSON(w, http.StatusBadRequest, errorResponse{Error: "decisionId is required"})
		return
	}

	detail, err := h.query.GetDecisionDetail(r.Context(), decisionID)
	if err != nil {
		writeError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, detail)
}

// ListDecisionModuleCandidates GET /api/decisions/{decisionId}/candidates/modules
//
// 承接 DecisionModuleCandidateRead（phase03-10 §6.2 读组 / §7.7 RPC→HTTP 映射）。
// decisionId 由 URL 路径参数承接，用于排除已建立关联的目标。
//
// 无可关联候选时返回空列表（[]），不返回错误
// （phase03-10 §5.10 空候选结果语义）。
func (h *QueryHandler) ListDecisionModuleCandidates(w http.ResponseWriter, r *http.Request) {
	decisionID := chi.URLParam(r, "decisionId")
	if decisionID == "" {
		writeJSON(w, http.StatusBadRequest, errorResponse{Error: "decisionId is required"})
		return
	}

	items, err := h.query.ListModuleCandidates(r.Context(), decisionID)
	if err != nil {
		writeError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, items)
}
