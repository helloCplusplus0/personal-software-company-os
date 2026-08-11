// Package handler — 读组入口层。
//
// 承接 ModuleListRead / ModuleDetailRead。
//
// phase07-09：L1/L2 候选 compat 入口（ListProductCandidates / ListRepositoryCandidates）
// 已退场。canonical 候选读取已迁移到 Product Registry / Repository Binding 的 Connect handler。
//
// 文件落点：backend/internal/moduleregistry/handler/query_handler.go
package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/psco/backend/internal/moduleregistry"
	"github.com/psco/backend/internal/moduleregistry/service"
)

// QueryHandler 承接读组 HTTP 入口。
type QueryHandler struct {
	query *service.QueryService
}

// NewQueryHandler 构造 QueryHandler。
func NewQueryHandler(q *service.QueryService) *QueryHandler {
	return &QueryHandler{query: q}
}

// ListModules GET /api/modules?queryText=...&statusFilter=...
//
// 承接 ModuleListRead（§6.2 读组）。
func (h *QueryHandler) ListModules(w http.ResponseWriter, r *http.Request) {
	q := moduleregistry.ListQuery{
		QueryText:    r.URL.Query().Get("queryText"),
		StatusFilter: moduleregistry.ModuleStatus(r.URL.Query().Get("statusFilter")),
	}

	items, err := h.query.ListModules(r.Context(), q)
	if err != nil {
		writeError(w, err)
		return
	}
	if items == nil {
		items = []moduleregistry.ModuleListItem{}
	}
	writeJSON(w, http.StatusOK, items)
}

// GetModuleDetail GET /api/modules/{moduleId}
//
// 承接 ModuleDetailRead（§6.2 读组），内嵌 Decision 附属读取。
func (h *QueryHandler) GetModuleDetail(w http.ResponseWriter, r *http.Request) {
	moduleID := chi.URLParam(r, "moduleId")
	if moduleID == "" {
		writeJSON(w, http.StatusBadRequest, errorResponse{Error: "moduleId is required"})
		return
	}

	detail, err := h.query.GetModuleDetail(r.Context(), moduleID)
	if err != nil {
		writeError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, detail)
}
