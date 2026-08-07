// Package handler — 读组入口层。
//
// 承接 ProductListRead / ProductDetailRead / ProductModuleCandidateRead。
// 文件落点：backend/internal/productregistry/handler/query_handler.go
package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/psco/backend/internal/productregistry"
	"github.com/psco/backend/internal/productregistry/service"
)

// QueryHandler 承接读组 HTTP 入口。
type QueryHandler struct {
	query *service.QueryService
}

// NewQueryHandler 构造 QueryHandler。
func NewQueryHandler(q *service.QueryService) *QueryHandler {
	return &QueryHandler{query: q}
}

// ListProducts GET /api/products?queryText=...&statusFilter=...
//
// 承接 ProductListRead（phase04-07 读组）。
func (h *QueryHandler) ListProducts(w http.ResponseWriter, r *http.Request) {
	q := productregistry.ListProductsQuery{
		QueryText:    r.URL.Query().Get("queryText"),
		StatusFilter: productregistry.ProductStatus(r.URL.Query().Get("statusFilter")),
	}

	items, err := h.query.ListProducts(r.Context(), q)
	if err != nil {
		writeError(w, err)
		return
	}
	if items == nil {
		items = []productregistry.ProductListItem{}
	}
	writeJSON(w, http.StatusOK, items)
}

// GetProductDetail GET /api/products/{productId}
//
// 承接 ProductDetailRead（phase04-07 读组）。
func (h *QueryHandler) GetProductDetail(w http.ResponseWriter, r *http.Request) {
	productID := chi.URLParam(r, "productId")
	if productID == "" {
		writeJSON(w, http.StatusBadRequest, errorResponse{Error: "productId is required"})
		return
	}

	detail, err := h.query.GetProductDetail(r.Context(), productID)
	if err != nil {
		writeError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, detail)
}

// ListProductModuleCandidates GET /api/products/{productId}/candidates/modules
//
// 承接 ProductModuleCandidateRead（phase04-07 候选读取）。
// 无可关联候选时返回空列表，不返回错误。
func (h *QueryHandler) ListProductModuleCandidates(w http.ResponseWriter, r *http.Request) {
	productID := chi.URLParam(r, "productId")
	if productID == "" {
		writeJSON(w, http.StatusBadRequest, errorResponse{Error: "productId is required"})
		return
	}

	items, err := h.query.ListProductModuleCandidates(r.Context(), productID)
	if err != nil {
		writeError(w, err)
		return
	}
	if items == nil {
		items = []productregistry.ProductModuleCandidate{}
	}
	writeJSON(w, http.StatusOK, items)
}
