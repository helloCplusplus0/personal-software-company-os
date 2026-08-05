// Package handler — 写组入口层。
//
// 承接 ModuleCreateWrite / ModuleReleaseWrite / ModuleBindingWrite。
// 文件落点：backend/internal/moduleregistry/handler/command_handler.go
package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/psco/backend/internal/moduleregistry"
	"github.com/psco/backend/internal/moduleregistry/service"
)

// CommandHandler 承接写组 HTTP 入口。
type CommandHandler struct {
	command *service.CommandService
}

// NewCommandHandler 构造 CommandHandler。
func NewCommandHandler(c *service.CommandService) *CommandHandler {
	return &CommandHandler{command: c}
}

// CreateModule POST /api/modules
//
// 承接 ModuleCreateWrite（§6.2 写组）。
// 成功返回新建模块对象（含 id），支持前端回流到 ModuleDetailPage。
func (h *CommandHandler) CreateModule(w http.ResponseWriter, r *http.Request) {
	var req moduleregistry.CreateModuleRequest
	if !decodeJSON(w, r, &req) {
		return
	}

	m, err := h.command.CreateModule(r.Context(), req)
	if err != nil {
		writeError(w, err)
		return
	}
	writeJSON(w, http.StatusCreated, m)
}

// CreateRelease POST /api/modules/{moduleId}/releases
//
// 承接 ModuleReleaseWrite（§6.2 写组），moduleId 由路径参数隐式承接。
// 成功返回新建版本对象，前端默认回流到当前模块的 ModuleDetailPage。
func (h *CommandHandler) CreateRelease(w http.ResponseWriter, r *http.Request) {
	moduleID := chi.URLParam(r, "moduleId")
	if moduleID == "" {
		writeJSON(w, http.StatusBadRequest, errorResponse{Error: "moduleId is required"})
		return
	}

	var req moduleregistry.CreateReleaseRequest
	if !decodeJSON(w, r, &req) {
		return
	}

	rel, err := h.command.CreateRelease(r.Context(), moduleID, req)
	if err != nil {
		writeError(w, err)
		return
	}
	writeJSON(w, http.StatusCreated, rel)
}

// BindModuleToProduct POST /api/modules/{moduleId}/bindings/products
//
// 承接 ModuleBindingWrite 的产品绑定子动作（§6.2 写组）。
// 成功返回 204 No Content，前端停留 ModuleDetailPage 并重新读取绑定结果。
func (h *CommandHandler) BindModuleToProduct(w http.ResponseWriter, r *http.Request) {
	moduleID := chi.URLParam(r, "moduleId")
	if moduleID == "" {
		writeJSON(w, http.StatusBadRequest, errorResponse{Error: "moduleId is required"})
		return
	}

	var req moduleregistry.BindModuleToProductRequest
	if !decodeJSON(w, r, &req) {
		return
	}

	if err := h.command.BindModuleToProduct(r.Context(), moduleID, req.ProductID); err != nil {
		writeError(w, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// MapModuleToRepository POST /api/modules/{moduleId}/bindings/repositories
//
// 承接 ModuleBindingWrite 的仓库映射子动作（§6.2 写组）。
// 成功返回 204 No Content，前端停留 ModuleDetailPage 并重新读取绑定结果。
func (h *CommandHandler) MapModuleToRepository(w http.ResponseWriter, r *http.Request) {
	moduleID := chi.URLParam(r, "moduleId")
	if moduleID == "" {
		writeJSON(w, http.StatusBadRequest, errorResponse{Error: "moduleId is required"})
		return
	}

	var req moduleregistry.MapModuleToRepositoryRequest
	if !decodeJSON(w, r, &req) {
		return
	}

	if err := h.command.MapModuleToRepository(r.Context(), moduleID, req.RepositoryID); err != nil {
		writeError(w, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
