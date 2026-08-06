// Package handler — 写组入口层。
//
// 承接 DecisionWrite (RecordDecision) 与 DecisionLinkWrite (LinkDecisionToTarget)
// （phase03-10 §10.4 写组单文件编排）。
//
// 文件落点：backend/internal/decisioncenter/handler/command_handler.go
package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/psco/backend/internal/decisioncenter"
	"github.com/psco/backend/internal/decisioncenter/service"
)

// CommandHandler 承接写组 HTTP 入口。
type CommandHandler struct {
	command *service.CommandService
}

// NewCommandHandler 构造 CommandHandler。
func NewCommandHandler(c *service.CommandService) *CommandHandler {
	return &CommandHandler{command: c}
}

// CreateDecision POST /api/decisions
//
// 承接 DecisionWrite (RecordDecision)（phase03-10 §6.2 写组 / §7.7 RPC→HTTP 映射）。
// 成功返回 201 + decision_id，支持前端回流到 DecisionDetailPage
// （phase03-10 §6.4 不返回完整 Decision 对象）。
func (h *CommandHandler) CreateDecision(w http.ResponseWriter, r *http.Request) {
	var req decisioncenter.CreateDecisionRequest
	if !decodeJSON(w, r, &req) {
		return
	}

	resp, err := h.command.CreateDecision(r.Context(), req)
	if err != nil {
		writeError(w, err)
		return
	}
	writeJSON(w, http.StatusCreated, resp)
}

// LinkDecisionToTarget POST /api/decisions/{decisionId}/links
//
// 承接 DecisionLinkWrite (LinkDecisionToTarget)（phase03-10 §6.2 写组 / §7.7 RPC→HTTP 映射）。
// decisionId 由 URL 路径参数承接，handler 在进入业务层前显式组装为 service 调用参数，
// 不放在 JSON 请求体（phase03-10 §7.7 约束）。
//
// 成功返回 204 No Content，前端通过 DecisionDetailRead 重新读取
// （phase03-10 §6.4 不返回 link 结果或 detail reread 标识）。
func (h *CommandHandler) LinkDecisionToTarget(w http.ResponseWriter, r *http.Request) {
	decisionID := chi.URLParam(r, "decisionId")
	if decisionID == "" {
		writeJSON(w, http.StatusBadRequest, errorResponse{Error: "decisionId is required"})
		return
	}

	var req decisioncenter.LinkDecisionToTargetRequest
	if !decodeJSON(w, r, &req) {
		return
	}

	if err := h.command.LinkDecisionToTarget(r.Context(), decisionID, req); err != nil {
		writeError(w, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
