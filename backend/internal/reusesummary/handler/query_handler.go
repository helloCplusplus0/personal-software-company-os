// Package handler — ReuseSummary 读组入口层。
//
// 承接 GetReuseSummary 一个 GET endpoint。
// 对齐 phase06-14 spec §"Phase06 路由装配必须接入现有 chi 组合根"的路由注册矩阵：
//
//	GET /api/reuse-summary?scope=<dashboard|module_detail|product_detail>&module_id=<id>&product_id=<id>
//
// 文件落点：backend/internal/reusesummary/handler/query_handler.go
package handler

import (
	"net/http"

	"github.com/psco/backend/internal/reusesummary"
	"github.com/psco/backend/internal/reusesummary/service"
)

// QueryHandler 承接 ReuseSummary 读组 HTTP 入口。
type QueryHandler struct {
	query *service.QueryService
}

// NewQueryHandler 构造 QueryHandler。
func NewQueryHandler(q *service.QueryService) *QueryHandler {
	return &QueryHandler{query: q}
}

// GetReuseSummary GET /api/reuse-summary
//
// 承接 ReuseSummaryRead（phase06-14 读组）。
// handler 必须在进入业务层前显式组装对应的 Proto request 消息，
// 不得因 GET 入口无 body 就绕过 Proto request 这一合同边界。
//
// query 参数映射（phase06-14 spec §"Phase06 路由装配必须接入现有 chi 组合根"）：
//   - scope：dashboard / module_detail / product_detail（必填）
//   - module_id：module_detail 作用域使用
//   - product_id：product_detail 作用域使用
//
// 响应体必须为 GetReuseSummaryResponse envelope {"module_reuse_summary":[...], "capability_summary":[...]}，
// 与 reuse_summary.proto 冻结的唯一合同源对齐。
// 空结果返回 200 + 空列表；作用域参数非法返回 400；读取失败返回 500。
func (h *QueryHandler) GetReuseSummary(w http.ResponseWriter, r *http.Request) {
	// 显式组装 Proto request 语义：从 query 参数提取 scope / module_id / product_id
	scopeStr := r.URL.Query().Get("scope")
	moduleID := r.URL.Query().Get("module_id")
	productID := r.URL.Query().Get("product_id")

	scope := reusesummary.ReuseSummaryScope(scopeStr)

	result, err := h.query.ReadReuseSummary(r.Context(), scope, moduleID, productID)
	if err != nil {
		writeError(w, err)
		return
	}

	// 确保空态返回非 nil 列表
	if result.ModuleReuseSummary == nil {
		result.ModuleReuseSummary = []reusesummary.ModuleReuseSummary{}
	}
	if result.CapabilitySummary == nil {
		result.CapabilitySummary = []reusesummary.CapabilitySummary{}
	}

	writeJSON(w, http.StatusOK, result)
}
