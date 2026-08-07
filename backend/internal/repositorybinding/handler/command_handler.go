// Package handler — 写组入口层。
//
// 承接 RepositoryCreateWrite / RepositoryProductBindingWrite / RepositoryModuleMappingWrite。
// 文件落点：backend/internal/repositorybinding/handler/command_handler.go
package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/psco/backend/internal/repositorybinding"
	"github.com/psco/backend/internal/repositorybinding/service"
)

// CommandHandler 承接写组 HTTP 入口。
type CommandHandler struct {
	command *service.CommandService
}

// NewCommandHandler 构造 CommandHandler。
func NewCommandHandler(c *service.CommandService) *CommandHandler {
	return &CommandHandler{command: c}
}

// CreateRepository POST /api/repositories
//
// 承接 RepositoryCreateWrite（phase04-07 写组）。
// 成功返回新建仓库 id（CreateRepositoryResponse），支撑前端回流到 Repository Binding Detail / Workspace。
func (h *CommandHandler) CreateRepository(w http.ResponseWriter, r *http.Request) {
	var req repositorybinding.CreateRepositoryRequest
	if !decodeJSON(w, r, &req) {
		return
	}

	resp, err := h.command.CreateRepository(r.Context(), req)
	if err != nil {
		writeError(w, err)
		return
	}
	writeJSON(w, http.StatusCreated, resp)
}

// BindRepositoryToProduct POST /api/repositories/{repositoryId}/bindings/products
//
// 承接 RepositoryProductBindingWrite（phase04-07 写组），repositoryId 由路径参数隐式承接。
// 成功返回 204 No Content，前端停留 Repository Binding Detail / Workspace 并重新读取绑定结果。
func (h *CommandHandler) BindRepositoryToProduct(w http.ResponseWriter, r *http.Request) {
	repositoryID := chi.URLParam(r, "repositoryId")
	if repositoryID == "" {
		writeJSON(w, http.StatusBadRequest, errorResponse{Error: "repositoryId is required"})
		return
	}

	var req repositorybinding.BindRepositoryToProductRequest
	if !decodeJSON(w, r, &req) {
		return
	}

	if err := h.command.BindRepositoryToProduct(r.Context(), repositoryID, req.ProductID); err != nil {
		writeError(w, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// MapModuleToRepository POST /api/repositories/{repositoryId}/bindings/modules
//
// 承接 RepositoryModuleMappingWrite（phase04-07 写组），repositoryId 由路径参数隐式承接。
// 成功返回 204 No Content，前端停留 Repository Binding Detail / Workspace 并重新读取绑定结果。
func (h *CommandHandler) MapModuleToRepository(w http.ResponseWriter, r *http.Request) {
	repositoryID := chi.URLParam(r, "repositoryId")
	if repositoryID == "" {
		writeJSON(w, http.StatusBadRequest, errorResponse{Error: "repositoryId is required"})
		return
	}

	var req repositorybinding.MapModuleToRepositoryRequest
	if !decodeJSON(w, r, &req) {
		return
	}

	if err := h.command.MapModuleToRepository(r.Context(), repositoryID, req.ModuleID); err != nil {
		writeError(w, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
