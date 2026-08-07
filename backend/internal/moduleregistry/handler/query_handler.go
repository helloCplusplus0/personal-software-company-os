// Package handler — 读组入口层。
//
// 承接 ModuleListRead / ModuleDetailRead 与历史候选读取兼容入口。
//
// phase04-12 起，canonical 候选读取已迁移到 Product Registry / Repository Binding；
// 这里仅保留 Module Detail 历史入口所需的兼容适配层，不再承接业务 owner。
// 文件落点：backend/internal/moduleregistry/handler/query_handler.go
package handler

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/psco/backend/internal/moduleregistry"
	"github.com/psco/backend/internal/moduleregistry/service"
)

// ProductCandidateReader 承接旧 ProductBindingCandidateRead 的兼容委派。
type ProductCandidateReader func(context.Context) ([]moduleregistry.ProductCandidate, error)

// RepositoryCandidateReader 承接旧 RepositoryBindingCandidateRead 的兼容委派。
type RepositoryCandidateReader func(context.Context) ([]moduleregistry.RepositoryCandidate, error)

// QueryHandler 承接读组 HTTP 入口。
type QueryHandler struct {
	query               *service.QueryService
	productCandidate    ProductCandidateReader
	repositoryCandidate RepositoryCandidateReader
}

// NewQueryHandler 构造 QueryHandler。
func NewQueryHandler(q *service.QueryService, pc ProductCandidateReader, rc RepositoryCandidateReader) *QueryHandler {
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
// 兼容旧 Module Detail 的 Product 候选入口；实际数据由 canonical Product Registry 主线提供。
func (h *QueryHandler) ListProductCandidates(w http.ResponseWriter, r *http.Request) {
	items, err := h.productCandidate(r.Context())
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
// 兼容旧 Module Detail 的 Repository 候选入口；实际数据由 canonical Repository Binding 主线提供。
func (h *QueryHandler) ListRepositoryCandidates(w http.ResponseWriter, r *http.Request) {
	items, err := h.repositoryCandidate(r.Context())
	if err != nil {
		writeError(w, err)
		return
	}
	if items == nil {
		items = []moduleregistry.RepositoryCandidate{}
	}
	writeJSON(w, http.StatusOK, items)
}
