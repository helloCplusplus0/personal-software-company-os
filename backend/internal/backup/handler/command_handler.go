// Package handler — Backup 写组入口层。
//
// 承接 CreateInstanceBackup 一个 POST endpoint。
// 对齐 phase06-14 spec §"Phase06 路由装配必须接入现有 chi 组合根"的路由注册矩阵：
//
//	POST /api/dashboard/backup
//
// 注意：写入响应不得替代 GetBackupSnapshot 的正式读取合同（phase06-05 / phase06-13 读取侧独立 owner 约束）。
//
// 文件落点：backend/internal/backup/handler/command_handler.go
package handler

import (
	"net/http"

	"github.com/psco/backend/internal/backup"
	"github.com/psco/backend/internal/backup/service"
)

// CommandHandler 承接 Backup 写组 HTTP 入口。
type CommandHandler struct {
	command *service.CommandService
}

// NewCommandHandler 构造 CommandHandler。
func NewCommandHandler(c *service.CommandService) *CommandHandler {
	return &CommandHandler{command: c}
}

// CreateInstanceBackup POST /api/dashboard/backup
//
// 承接 CreateInstanceBackupWrite（phase06-14 写组）。
// handler 显式组装空 Proto request 语义（直接调用 service 方法），不绕过 service 合同边界。
// 响应体必须为 CreateInstanceBackupResponse envelope {"snapshot": {...}}，
// 与 backup.proto 冻结的唯一合同源对齐。
// 资产装配失败或持久化失败时返回 500；成功时返回 200 + 写入后的 BackupSnapshot。
func (h *CommandHandler) CreateInstanceBackup(w http.ResponseWriter, r *http.Request) {
	snapshot, err := h.command.CreateInstanceBackup(r.Context())
	if err != nil {
		writeError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, backup.CreateBackupResult{
		Snapshot: snapshot,
	})
}
