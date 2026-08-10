// Package handler — Export HTTP 入口层共享工具。
//
// 只负责 HTTP 协议层事务：JSON 编解码、错误到状态码映射。
// 不承接业务语义，业务编排由 service 层完成（对齐 phase06-14 spec）。
//
// 文件落点：backend/internal/export/handler/response.go
package handler

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"github.com/psco/backend/internal/export"
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
//   - ErrAssetReadFailed → 500（资产装配失败）
//   - ErrExportPersistFailed → 500（持久化失败）
//   - ErrExportSnapshotReadFailed → 500（快照读取失败）
//   - 其他 → 500
func writeError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, export.ErrAssetReadFailed),
		errors.Is(err, export.ErrExportPersistFailed),
		errors.Is(err, export.ErrExportSnapshotReadFailed):
		slog.Error("export operation failed", "error", err)
		writeJSON(w, http.StatusInternalServerError, errorResponse{Error: "internal server error"})
	default:
		slog.Error("unhandled export error", "error", err)
		writeJSON(w, http.StatusInternalServerError, errorResponse{Error: "internal server error"})
	}
}
