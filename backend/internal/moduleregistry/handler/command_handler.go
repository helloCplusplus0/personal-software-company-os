// Package handler — 写组入口层。
//
// 承接 ModuleCreateWrite / ModuleReleaseWrite / ModuleBindingWrite（兼容委派）。
//
// phase04-12 起，BindModuleToProduct 与 MapModuleToRepository 的业务 owner 已迁移到
// Product Registry 与 Repository Binding。旧 Module Registry 入口若保留，只能作为
// 兼容适配层委派到新的 canonical 实现，不在本 service 层继续保留长期 owner 逻辑
// （phase04-12 spec §"phase02 旧 transport 入口若保留，必须只做兼容委派"）。
//
// 文件落点：backend/internal/moduleregistry/handler/command_handler.go
package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/psco/backend/internal/moduleregistry"
	productservice "github.com/psco/backend/internal/productregistry/service"
	reposervice "github.com/psco/backend/internal/repositorybinding/service"
	"github.com/psco/backend/internal/moduleregistry/service"
)

// CommandHandler 承接写组 HTTP 入口。
//
// 依赖注入：
//   - command: Module Registry 自身的 ModuleCreateWrite / ModuleReleaseWrite
//   - productBindingSvc: Product Registry 的 BindModuleToProduct（兼容委派目标）
//   - repositoryMappingSvc: Repository Binding 的 MapModuleToRepository（兼容委派目标）
type CommandHandler struct {
	command              *service.CommandService
	productBindingSvc    *productservice.CommandService
	repositoryMappingSvc *reposervice.CommandService
}

// NewCommandHandler 构造 CommandHandler。
//
// productBindingSvc 与 repositoryMappingSvc 用于旧模块中心绑定入口的兼容委派。
func NewCommandHandler(c *service.CommandService, productBindingSvc *productservice.CommandService, repositoryMappingSvc *reposervice.CommandService) *CommandHandler {
	return &CommandHandler{
		command:              c,
		productBindingSvc:    productBindingSvc,
		repositoryMappingSvc: repositoryMappingSvc,
	}
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
// 兼容委派入口（phase04-12）。
//
// 旧 Module Registry 模块中心写入口，委派到 Product Registry 的 canonical 实现
// productregistry.CommandService.BindModuleToProduct(productID, moduleID)。
//
// 注意参数顺序：旧入口以 module 为中心（moduleId 在 URL），新 canonical 以 product 为中心
// （productID 在 URL），委派时需要交换参数顺序。
//
// 成功返回 204 No Content，前端停留 ModuleDetailPage 并重新读取绑定结果。
// reread owner 仍是 ProductDetailRead（canonical 语义），不制造第二套 reread owner。
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

	// 委派到 Product Registry canonical 实现（参数顺序：productID, moduleID）
	if err := h.productBindingSvc.BindModuleToProduct(r.Context(), req.ProductID, moduleID); err != nil {
		writeError(w, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// MapModuleToRepository POST /api/modules/{moduleId}/bindings/repositories
//
// 兼容委派入口（phase04-12）。
//
// 旧 Module Registry 模块中心写入口，委派到 Repository Binding 的 canonical 实现
// repositorybinding.CommandService.MapModuleToRepository(repositoryID, moduleID)。
//
// 注意参数顺序：旧入口以 module 为中心（moduleId 在 URL），新 canonical 以 repository 为中心
// （repositoryId 在 URL），委派时需要交换参数顺序。
//
// 成功返回 204 No Content，前端停留 ModuleDetailPage 并重新读取绑定结果。
// reread owner 仍是 RepositoryDetailRead（canonical 语义），不制造第二套 reread owner。
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

	// 委派到 Repository Binding canonical 实现（参数顺序：repositoryID, moduleID）
	if err := h.repositoryMappingSvc.MapModuleToRepository(r.Context(), req.RepositoryID, moduleID); err != nil {
		writeError(w, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
