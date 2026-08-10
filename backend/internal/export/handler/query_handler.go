// Package handler — Export 读组入口层。
//
// 承接 GetExportSnapshot 一个 GET endpoint。
// 对齐 phase06-14 spec §"Phase06 路由装配必须接入现有 chi 组合根"的路由注册矩阵：
//
//	GET /api/dashboard/export
//
// 文件落点：backend/internal/export/handler/query_handler.go
package handler

import (
	"net/http"

	"github.com/psco/backend/internal/export"
	"github.com/psco/backend/internal/export/service"
)

// QueryHandler 承接 Export 读组 HTTP 入口。
type QueryHandler struct {
	query *service.QueryService
}

// NewQueryHandler 构造 QueryHandler。
func NewQueryHandler(q *service.QueryService) *QueryHandler {
	return &QueryHandler{query: q}
}

// GetExportSnapshot GET /api/dashboard/export
//
// 承接 ExportSnapshotRead（phase06-14 读组）。
// handler 显式组装空 Proto request 语义（直接调用 service 方法），不绕过 service 合同边界。
// 响应体必须为 GetExportSnapshotResponse envelope {"snapshot": {...}}，
// 与 export.proto 冻结的唯一合同源对齐，不得返回裸 ExportSnapshot。
// 整页失败时返回 500；预览态与历史快照都返回 200。
func (h *QueryHandler) GetExportSnapshot(w http.ResponseWriter, r *http.Request) {
	snapshot, err := h.query.ReadExportSnapshot(r.Context())
	if err != nil {
		writeError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, export.ExportSnapshotReadResult{
		Snapshot: snapshot,
	})
}
