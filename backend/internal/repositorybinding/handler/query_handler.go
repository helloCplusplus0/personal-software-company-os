// Package handler — 读组入口层。
//
// 承接 RepositoryListRead / RepositoryDetailRead /
// ProductBindingCandidateRead / RepositoryModuleCandidateRead。
// 文件落点：backend/internal/repositorybinding/handler/query_handler.go
package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/psco/backend/internal/repositorybinding"
	"github.com/psco/backend/internal/repositorybinding/service"
)

// QueryHandler 承接读组 HTTP 入口。
type QueryHandler struct {
	query *service.QueryService
}

// NewQueryHandler 构造 QueryHandler。
func NewQueryHandler(q *service.QueryService) *QueryHandler {
	return &QueryHandler{query: q}
}

// ListRepositories GET /api/repositories?queryText=...&statusFilter=...
//
// 承接 RepositoryListRead（phase04-07 读组）。
func (h *QueryHandler) ListRepositories(w http.ResponseWriter, r *http.Request) {
	q := repositorybinding.ListRepositoriesQuery{
		QueryText:    r.URL.Query().Get("queryText"),
		StatusFilter: repositorybinding.RepositoryStatus(r.URL.Query().Get("statusFilter")),
	}

	items, err := h.query.ListRepositories(r.Context(), q)
	if err != nil {
		writeError(w, err)
		return
	}
	if items == nil {
		items = []repositorybinding.RepositoryListItem{}
	}
	writeJSON(w, http.StatusOK, items)
}

// GetRepositoryDetail GET /api/repositories/{repositoryId}
//
// 承接 RepositoryDetailRead（phase04-07 读组）。
func (h *QueryHandler) GetRepositoryDetail(w http.ResponseWriter, r *http.Request) {
	repositoryID := chi.URLParam(r, "repositoryId")
	if repositoryID == "" {
		writeJSON(w, http.StatusBadRequest, errorResponse{Error: "repositoryId is required"})
		return
	}

	detail, err := h.query.GetRepositoryDetail(r.Context(), repositoryID)
	if err != nil {
		writeError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, detail)
}

// ListRepositoryProductCandidates GET /api/repositories/{repositoryId}/candidates/products
//
// 承接 ProductBindingCandidateRead（phase04-07 候选读取）。
// 无可关联候选时返回空列表，不返回错误。
func (h *QueryHandler) ListRepositoryProductCandidates(w http.ResponseWriter, r *http.Request) {
	repositoryID := chi.URLParam(r, "repositoryId")
	if repositoryID == "" {
		writeJSON(w, http.StatusBadRequest, errorResponse{Error: "repositoryId is required"})
		return
	}

	items, err := h.query.ListRepositoryProductCandidates(r.Context(), repositoryID)
	if err != nil {
		writeError(w, err)
		return
	}
	if items == nil {
		items = []repositorybinding.RepositoryProductCandidate{}
	}
	writeJSON(w, http.StatusOK, items)
}

// ListRepositoryModuleCandidates GET /api/repositories/{repositoryId}/candidates/modules
//
// 承接 RepositoryModuleCandidateRead（phase04-07 候选读取）。
// 无可关联候选时返回空列表，不返回错误。
func (h *QueryHandler) ListRepositoryModuleCandidates(w http.ResponseWriter, r *http.Request) {
	repositoryID := chi.URLParam(r, "repositoryId")
	if repositoryID == "" {
		writeJSON(w, http.StatusBadRequest, errorResponse{Error: "repositoryId is required"})
		return
	}

	items, err := h.query.ListRepositoryModuleCandidates(r.Context(), repositoryID)
	if err != nil {
		writeError(w, err)
		return
	}
	if items == nil {
		items = []repositorybinding.RepositoryModuleCandidate{}
	}
	writeJSON(w, http.StatusOK, items)
}
