// Package handler — HTTP 入口层共享工具。
//
// 只负责 HTTP 协议层事务：JSON 编解码、错误到状态码映射。
// 不承接业务语义，业务校验由 service 层完成（§9.3 分层语义）。
package handler

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"github.com/psco/backend/internal/moduleregistry"
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
// 映射规则：
//   - ErrModuleNotFound / ErrProductNotFound / ErrRepositoryNotFound → 404
//   - ErrDuplicateModuleName / ErrDuplicateBinding → 409
//   - ErrInvalidInput / ErrInvalidStatus / ErrInvalidReleaseStatus → 400
//   - 其他 → 500
func writeError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, moduleregistry.ErrModuleNotFound),
		errors.Is(err, moduleregistry.ErrProductNotFound),
		errors.Is(err, moduleregistry.ErrRepositoryNotFound):
		writeJSON(w, http.StatusNotFound, errorResponse{Error: err.Error()})
	case errors.Is(err, moduleregistry.ErrDuplicateModuleName),
		errors.Is(err, moduleregistry.ErrDuplicateBinding),
		errors.Is(err, moduleregistry.ErrDuplicateReleaseVersion):
		writeJSON(w, http.StatusConflict, errorResponse{Error: err.Error()})
	case errors.Is(err, moduleregistry.ErrInvalidInput),
		errors.Is(err, moduleregistry.ErrInvalidStatus),
		errors.Is(err, moduleregistry.ErrInvalidReleaseStatus):
		writeJSON(w, http.StatusBadRequest, errorResponse{Error: err.Error()})
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
