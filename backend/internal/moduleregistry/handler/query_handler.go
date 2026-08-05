// Package handler — 读组入口层。
//
// 承接 ModuleListRead / ModuleDetailRead / ProductBindingCandidateRead / RepositoryBindingCandidateRead。
// 文件落点：backend/internal/moduleregistry/handler/query_handler.go
package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/psco/backend/internal/moduleregistry"
	"github.com/psco/backend/internal/moduleregistry/candidate"
	"github.com/psco/backend/internal/moduleregistry/service"
)

// QueryHandler 承接读组 HTTP 入口。
type QueryHandler struct {
	query              *service.QueryService
	productCandidate   *candidate.ProductCandidateRead
	repositoryCandidate *candidate.RepositoryCandidateRead
}

// NewQueryHandler 构造 QueryHandler。
func NewQueryHandler(q *service.QueryService, pc *candidate.ProductCandidateRead, rc *candidate.RepositoryCandidateRead) *QueryHandler {
	return &QueryHandler{query: q, productCandidate: pc, repositoryCandidate: rc}
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

// ListProductCandidates GET /api/candidates/products
//
// 承接 ProductBindingCandidateRead（§6.2 候选读取，phase02 由 Module Registry 临时承接）。
func (h *QueryHandler) ListProductCandidates(w http.ResponseWriter, r *http.Request) {
	items, err := h.productCandidate.List(r.Context())
	if err != nil {
		writeError(w, err)
		return
	}
	if items == nil {
		items = []moduleregistry.ProductCandidate{}
	}
	writeJSON(w, http.StatusOK, items)
}

// ListRepositoryCandidates GET /api/candidates/repositories
//
// 承接 RepositoryBindingCandidateRead（§6.2 候选读取，phase02 由 Module Registry 临时承接）。
func (h *QueryHandler) ListRepositoryCandidates(w http.ResponseWriter, r *http.Request) {
	items, err := h.repositoryCandidate.List(r.Context())
	if err != nil {
		writeError(w, err)
		return
	}
	if items == nil {
		items = []moduleregistry.RepositoryCandidate{}
	}
	writeJSON(w, http.StatusOK, items)
}
