// Package handler — Backup 读组入口层（read / verify 子路径）。
//
// 承接 GetBackupSnapshot 一个 GET endpoint。
// 对齐 phase06-14 spec §"Phase06 路由装配必须接入现有 chi 组合根"的路由注册矩阵：
//
//	GET /api/dashboard/backup
//
// 该入口是当前阶段正式 read / verify 子路径合同出口，
// 由独立读取 owner（BackupRead / QueryService）承接，不与 BackupWrite 写入响应耦合。
//
// 文件落点：backend/internal/backup/handler/query_handler.go
package handler

import (
	"net/http"

	"github.com/psco/backend/internal/backup"
	"github.com/psco/backend/internal/backup/service"
)

// QueryHandler 承接 Backup 读组 HTTP 入口。
type QueryHandler struct {
	query *service.QueryService
}

// NewQueryHandler 构造 QueryHandler。
func NewQueryHandler(q *service.QueryService) *QueryHandler {
	return &QueryHandler{query: q}
}

// GetBackupSnapshot GET /api/dashboard/backup
//
// 承接 BackupSnapshotRead（phase06-14 read / verify 子路径）。
// handler 显式组装空 Proto request 语义（直接调用 service 方法），不绕过 service 合同边界。
// 响应体必须为 GetBackupSnapshotResponse envelope {"snapshot": {...}}，
// 与 backup.proto 冻结的唯一合同源对齐，不得返回裸 BackupSnapshot。
//
// 返回语义：
//   - 有备份记录 → 200 + 校验后的 BackupSnapshot（verified / verify_failed）
//   - 无备份记录 → 200 + snapshot: null
//   - 读取失败 → 500
func (h *QueryHandler) GetBackupSnapshot(w http.ResponseWriter, r *http.Request) {
	snapshot, err := h.query.ReadBackupSnapshot(r.Context())
	if err != nil {
		writeError(w, err)
		return
	}

	// 无历史备份记录时返回 snapshot: null，由前端按正式空态消费。
	if snapshot == nil {
		writeJSON(w, http.StatusOK, backup.BackupSnapshotReadResult{
			Snapshot: nil,
		})
		return
	}

	writeJSON(w, http.StatusOK, backup.BackupSnapshotReadResult{
		Snapshot: snapshot,
	})
}
