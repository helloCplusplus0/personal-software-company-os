// Package handler — HTTP 入口层共享工具。
//
// 只负责 HTTP 协议层事务：JSON 编解码、错误到状态码映射。
// 不承接业务语义，业务校验由 service 层完成（§9.3 分层语义）。
//
// phase04-12 起，本 handler 承接旧模块中心绑定入口的兼容委派，
// 因此 writeError 需要同时处理 moduleregistry / productregistry / repositorybinding
// 三个模块的哨兵错误，确保兼容委派不制造第二套错误语义
// （phase04-12 spec §"兼容委派不得制造第二套 reread owner 或第二套错误语义"）。
package handler

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"github.com/psco/backend/internal/moduleregistry"
	"github.com/psco/backend/internal/productregistry"
	"github.com/psco/backend/internal/repositorybinding"
)

// writeJSON 以 application/json 写入响应体。
func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	if v == nil {
		return
	}
	if err := json.NewEncoder(w).Encode(v); err != nil {
		slog.Error("encode json response failed", "error", err)
	}
}

// errorResponse 标准错误响应体。
type errorResponse struct {
	Error string `json:"error"`
}

// writeError 将业务错误映射为 HTTP 状态码并写入错误响应。
//
// 映射规则（对齐 phase04-04 错误语义，跨三模块单值稳定）：
//   - 资源不存在（Module / Product / Repository NotFound）→ 404
//   - 重复绑定 / 重复映射 / 名称冲突 / 版本冲突 → 409
//   - 非法输入 / 非法状态 / 非 active 候选 → 400
//   - 其他 → 500
//
// 兼容委派场景：旧 Module Registry 入口委派到 Product Registry / Repository Binding 后，
// 返回的错误来自 canonical owner，本函数统一映射，不产生第二套错误语义。
func writeError(w http.ResponseWriter, err error) {
	switch {
	// --- 404 资源不存在 ---
	case errors.Is(err, moduleregistry.ErrModuleNotFound),
		errors.Is(err, moduleregistry.ErrProductNotFound),
		errors.Is(err, moduleregistry.ErrRepositoryNotFound),
		errors.Is(err, productregistry.ErrProductNotFound),
		errors.Is(err, productregistry.ErrModuleNotFound),
		errors.Is(err, repositorybinding.ErrRepositoryNotFound),
		errors.Is(err, repositorybinding.ErrProductNotFound),
		errors.Is(err, repositorybinding.ErrModuleNotFound):
		writeJSON(w, http.StatusNotFound, errorResponse{Error: err.Error()})

	// --- 409 冲突 ---
	case errors.Is(err, moduleregistry.ErrDuplicateModuleName),
		errors.Is(err, moduleregistry.ErrDuplicateBinding),
		errors.Is(err, moduleregistry.ErrDuplicateReleaseVersion),
		errors.Is(err, productregistry.ErrDuplicateBinding),
		errors.Is(err, repositorybinding.ErrDuplicateBinding),
		errors.Is(err, repositorybinding.ErrDuplicateMapping):
		writeJSON(w, http.StatusConflict, errorResponse{Error: err.Error()})

	// --- 400 非法输入 ---
	case errors.Is(err, moduleregistry.ErrInvalidInput),
		errors.Is(err, moduleregistry.ErrInvalidStatus),
		errors.Is(err, moduleregistry.ErrInvalidReleaseStatus),
		errors.Is(err, productregistry.ErrInvalidInput),
		errors.Is(err, productregistry.ErrInvalidStatus),
		errors.Is(err, productregistry.ErrModuleNotActive),
		errors.Is(err, repositorybinding.ErrInvalidInput),
		errors.Is(err, repositorybinding.ErrInvalidStatus),
		errors.Is(err, repositorybinding.ErrProductNotActive),
		errors.Is(err, repositorybinding.ErrModuleNotActive):
		writeJSON(w, http.StatusBadRequest, errorResponse{Error: err.Error()})

	// --- 500 兜底 ---
	default:
		slog.Error("unhandled service error", "error", err)
		writeJSON(w, http.StatusInternalServerError, errorResponse{Error: "internal server error"})
	}
}

// decodeJSON 解码请求体到目标对象。返回 true 表示成功，false 表示已写入错误响应。
func decodeJSON(w http.ResponseWriter, r *http.Request, dst any) bool {
	if err := json.NewDecoder(r.Body).Decode(dst); err != nil {
		writeJSON(w, http.StatusBadRequest, errorResponse{Error: "invalid json body"})
		return false
	}
	return true
}
