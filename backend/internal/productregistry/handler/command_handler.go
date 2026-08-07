// Package handler — 写组入口层。
//
// 承接 ProductCreateWrite / ProductModuleBindingWrite。
// 文件落点：backend/internal/productregistry/handler/command_handler.go
package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/psco/backend/internal/productregistry"
	"github.com/psco/backend/internal/productregistry/service"
)

// CommandHandler 承接写组 HTTP 入口。
type CommandHandler struct {
	command *service.CommandService
}

// NewCommandHandler 构造 CommandHandler。
func NewCommandHandler(c *service.CommandService) *CommandHandler {
	return &CommandHandler{command: c}
}

// CreateProduct POST /api/products
//
// 承接 ProductCreateWrite（phase04-07 写组）。
// 成功返回新建产品 id（CreateProductResponse），支撑前端回流到 Product Detail。
func (h *CommandHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var req productregistry.CreateProductRequest
	if !decodeJSON(w, r, &req) {
		return
	}

	resp, err := h.command.CreateProduct(r.Context(), req)
	if err != nil {
		writeError(w, err)
		return
	}
	writeJSON(w, http.StatusCreated, resp)
}

// BindModuleToProduct POST /api/products/{productId}/bindings/modules
//
// 承接 ProductModuleBindingWrite（phase04-07 写组），productId 由路径参数隐式承接。
// 成功返回 204 No Content，前端停留 Product Detail 并重新读取绑定结果。
func (h *CommandHandler) BindModuleToProduct(w http.ResponseWriter, r *http.Request) {
	productID := chi.URLParam(r, "productId")
	if productID == "" {
		writeJSON(w, http.StatusBadRequest, errorResponse{Error: "productId is required"})
		return
	}

	var req productregistry.BindModuleToProductRequest
	if !decodeJSON(w, r, &req) {
		return
	}

	if err := h.command.BindModuleToProduct(r.Context(), productID, req.ModuleID); err != nil {
		writeError(w, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
