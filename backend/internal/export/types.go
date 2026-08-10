// Package export 承载导出（Export）后端模块的全部业务实现。
//
// 分层语义（对齐 phase06-14 spec §"Phase06 后端模块必须按现有主线结构落地"）：
//   - handler/     入口层：只负责承接 HTTP 请求与返回结果
//   - service/     业务编排层：导出快照读取与导出执行编排
//   - candidate/   外部连接层：跨模块 9 类核心资产 reader（Export 自拥有）
//   - repository/  持久化层：instance_exports 表读写
//
// 当前阶段 Export 实现为"读取快照 + 触发导出"双路径主线，
// 并通过数据库元数据表（instance_exports）形成最新可读取快照，
// 而不是只返回一次性导出响应。
//
// 本文件定义跨层共享的 API 消息结构。
// 约束：消息结构从 proto/psco/export/v1/export.proto 单向派生或显式对齐，
// 不直接暴露存储模型，不在 types.go 或 handler DTO 中新增 .proto 中不存在的业务字段语义。
package export

import "time"

// ============================================================================
// 枚举类型
// ============================================================================

// ExportAssetScope 导出 / 备份覆盖矩阵的核心资产类型。
// 对齐 proto ExportAssetScope：9 类核心资产。
type ExportAssetScope string

const (
	ExportAssetScopeUnspecified       ExportAssetScope = ""
	ExportAssetScopeProducts          ExportAssetScope = "products"
	ExportAssetScopeModules           ExportAssetScope = "modules"
	ExportAssetScopeReleases          ExportAssetScope = "releases"
	ExportAssetScopeRepositories      ExportAssetScope = "repositories"
	ExportAssetScopeDecisions         ExportAssetScope = "decisions"
	ExportAssetScopeDecisionLinks     ExportAssetScope = "decision_links"
	ExportAssetScopeProductModules    ExportAssetScope = "product_modules"
	ExportAssetScopeProductRepositories ExportAssetScope = "product_repositories"
	ExportAssetScopeModuleRepositories ExportAssetScope = "module_repositories"
)

// AllExportAssetScopes 返回当前阶段冻结的 9 类核心资产完整列表（不含 UNSPECIFIED）。
// 用于装配覆盖矩阵与校验完整性。
func AllExportAssetScopes() []ExportAssetScope {
	return []ExportAssetScope{
		ExportAssetScopeProducts,
		ExportAssetScopeModules,
		ExportAssetScopeReleases,
		ExportAssetScopeRepositories,
		ExportAssetScopeDecisions,
		ExportAssetScopeDecisionLinks,
		ExportAssetScopeProductModules,
		ExportAssetScopeProductRepositories,
		ExportAssetScopeModuleRepositories,
	}
}

// ExportResultStatus 导出执行的结果状态。
// 对齐 proto ExportResultStatus。
type ExportResultStatus string

const (
	ExportResultStatusUnspecified ExportResultStatus = ""
	ExportResultStatusSuccess     ExportResultStatus = "success"
	ExportResultStatusInProgress  ExportResultStatus = "in_progress"
	ExportResultStatusFailed      ExportResultStatus = "failed"
)

// ============================================================================
// 核心消息 DTO
// ============================================================================

// ExportSnapshot 导出快照主读模型。
// 对齐 proto ExportSnapshot。
// 该消息同时承接快照读取（GetExportSnapshot）与导出执行（ExportCoreAssets）的结果形状。
type ExportSnapshot struct {
	AssetScope    []ExportAssetScope `json:"asset_scope"`
	CreatedAt     time.Time          `json:"created_at"`
	ResultStatus  ExportResultStatus `json:"result_status"`
	ResultSummary string             `json:"result_summary"`
}

// ============================================================================
// 响应 DTO
// ============================================================================

// ExportSnapshotReadResult GetExportSnapshot 的响应结构。
// 对齐 proto GetExportSnapshotResponse：单一 snapshot 字段包装主读模型。
// handler 必须返回此包络结构，不得直接返回裸 ExportSnapshot。
type ExportSnapshotReadResult struct {
	Snapshot *ExportSnapshot `json:"snapshot"`
}

// ExportCoreAssetsResult ExportCoreAssets 的响应结构。
// 对齐 proto ExportCoreAssetsResponse：单一 snapshot 字段包装导出结果。
type ExportCoreAssetsResult struct {
	Snapshot *ExportSnapshot `json:"snapshot"`
}
