// Package handler — Onboarding HTTP 入口层共享工具。
//
// 只负责 HTTP 协议层事务：JSON 编解码、错误到状态码映射。
// 不承接业务语义，业务编排由 service 层完成（对齐 phase06-14 spec §"Phase06 后端模块必须按现有主线结构落地"）。
//
// 文件落点：backend/internal/onboarding/handler/response.go
package handler

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"github.com/psco/backend/internal/onboarding"
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
// 映射规则（对齐 phase06-14 spec）：
//   - ErrFirstRunStateReadFailed → 500（整页失败）
//   - 其他 → 500
func writeError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, onboarding.ErrFirstRunStateReadFailed):
		slog.Error("onboarding first run state read failed", "error", err)
		writeJSON(w, http.StatusInternalServerError, errorResponse{Error: "internal server error"})
	default:
		slog.Error("unhandled onboarding error", "error", err)
		writeJSON(w, http.StatusInternalServerError, errorResponse{Error: "internal server error"})
	}
}
