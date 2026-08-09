// Package handler — Dashboard HTTP 入口层共享工具。
//
// 只负责 HTTP 协议层事务：JSON 编解码、错误到状态码映射。
// 不承接业务语义，业务编排由 service 层完成（对齐 phase05-07 分层语义）。
//
// 文件落点：backend/internal/dashboard/handler/response.go
package handler

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"github.com/psco/backend/internal/dashboard"
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
// 映射规则（对齐 phase05-04 错误语义与 phase05-12 spec §"整页失败响应 / 局部失败响应"）：
//   - ErrOverviewReadFailed → 500（整页失败）
//   - ErrFeedbackSignalReadFailed → 500（局部失败，但当前阶段不在响应包络里发明局部失败标记）
//   - ErrRecentActivityReadFailed → 500（局部失败，同上）
//   - 其他 → 500
//
// 当前阶段不在附属读 endpoint 响应包络里额外发明"局部失败标记"，
// Dashboard 页面级"主聚合成功、附属聚合局部失败"语义由前端基于三次独立请求结果组合派生。
func writeError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, dashboard.ErrOverviewReadFailed),
		errors.Is(err, dashboard.ErrFeedbackSignalReadFailed),
		errors.Is(err, dashboard.ErrRecentActivityReadFailed):
		slog.Error("dashboard read failed", "error", err)
		writeJSON(w, http.StatusInternalServerError, errorResponse{Error: "internal server error"})
	default:
		slog.Error("unhandled dashboard error", "error", err)
		writeJSON(w, http.StatusInternalServerError, errorResponse{Error: "internal server error"})
	}
}
