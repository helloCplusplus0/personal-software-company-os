// Package repositorybinding 承载 Repository Binding 后端模块的全部业务实现。
//
// 分层语义（对齐 phase04-07 / phase04-10 已冻结结论）：
//   - handler/     入口层：只负责承接 HTTP 请求与返回结果
//   - service/     业务编排层：动作语义、校验顺序与 reread 所需读取聚合
//   - repository/  数据访问层：repositories / product_repositories / module_repositories
//   - candidate/   外部连接层：Product / Module 候选读取（绑定前提）
//
// 本文件定义跨层共享的 API 消息结构。
// 约束：消息结构从 proto/psco/repository_binding/v1/repository_binding.proto 单向派生或显式对齐，
// 不直接暴露存储模型，不在 types.go 或 handler DTO 中新增 .proto 中不存在的业务字段语义。
package repositorybinding

import (
	"time"

	"github.com/psco/backend/internal/moduleregistry"
)

// RepositoryStatus 冻结为 active / archived（phase04-04）。
// 与 psco.common.v1.ActiveArchivedStatus 单值对齐。
type RepositoryStatus string

const (
	RepositoryStatusActive   RepositoryStatus = "active"
	RepositoryStatusArchived RepositoryStatus = "archived"
)

// Repository 核心对象（phase04-04 冻结的数据范围）。
// 对齐 proto Repository：id / name / url / provider / status / created_at。
type Repository struct {
	ID        string           `json:"id"`
	Name      string           `json:"name"`
	URL       string           `json:"url"`
	Provider  string           `json:"provider"`
	Status    RepositoryStatus `json:"status"`
	CreatedAt time.Time        `json:"created_at"`
}

// RepositoryListItem 列表读取模型（phase04-04）。
// 对齐 proto RepositoryListItem：
// id / name / url / provider / status / created_at / product_bind_count / module_bind_count。
type RepositoryListItem struct {
	ID                 string           `json:"id"`
	Name               string           `json:"name"`
	URL                string           `json:"url"`
	Provider           string           `json:"provider"`
	Status             RepositoryStatus `json:"status"`
	CreatedAt          time.Time        `json:"created_at"`
	ProductBindCount   int              `json:"product_bind_count"`
	ModuleBindCount    int              `json:"module_bind_count"`
}

// BoundProductSummary 详情读取中的已绑定 Product 摘要。
// 对齐 proto BoundProductSummary：product_id / product_name / product_status。
// product_status 复用 ActiveArchivedStatus（psco.common.v1.ActiveArchivedStatus）。
type BoundProductSummary struct {
	ProductID     string           `json:"product_id"`
	ProductName   string           `json:"product_name"`
	ProductStatus RepositoryStatus `json:"product_status"`
}

// MappedModuleSummary 详情读取中的已映射 Module 摘要。
// 对齐 proto MappedModuleSummary：module_id / module_name / module_status。
// module_status 复用 moduleregistry.ModuleStatus（跨包 import psco.module_registry.v1.ModuleStatus）。
type MappedModuleSummary struct {
	ModuleID     string                      `json:"module_id"`
	ModuleName   string                      `json:"module_name"`
	ModuleStatus moduleregistry.ModuleStatus `json:"module_status"`
}

// RepositoryDetail 详情读取模型（phase04-04）。
// 对齐 proto RepositoryDetail：repository + repeated BoundProductSummary + repeated MappedModuleSummary。
// 候选读取结果不内嵌于此，必须通过 ListRepositoryProductCandidates /
// ListRepositoryModuleCandidates 独立承接。
type RepositoryDetail struct {
	Repository     Repository            `json:"repository"`
	BoundProducts  []BoundProductSummary `json:"bound_products"`
	MappedModules  []MappedModuleSummary `json:"mapped_modules"`
}

// RepositoryProductCandidate Product 候选读取返回项（phase04-04）。
// 对齐 proto RepositoryProductCandidate：product_id / product_name / product_status。
type RepositoryProductCandidate struct {
	ProductID     string           `json:"product_id"`
	ProductName   string           `json:"product_name"`
	ProductStatus RepositoryStatus `json:"product_status"`
}

// RepositoryModuleCandidate Module 候选读取返回项（phase04-04）。
// 对齐 proto RepositoryModuleCandidate：module_id / module_name / module_status。
type RepositoryModuleCandidate struct {
	ModuleID     string                      `json:"module_id"`
	ModuleName   string                      `json:"module_name"`
	ModuleStatus moduleregistry.ModuleStatus `json:"module_status"`
}

// --- 写组请求体 ---

// CreateRepositoryRequest 创建写入最小字段（phase04-04）。
// 对齐 proto CreateRepositoryRequest：name / url / provider / status。
// status 必填且必须属于 active / archived，UNSPECIFIED 不允许作为合法写入值。
type CreateRepositoryRequest struct {
	Name     string           `json:"name"`
	URL      string           `json:"url"`
	Provider string           `json:"provider"`
	Status   RepositoryStatus `json:"status"`
}

// CreateRepositoryResponse 创建响应，只返回新建 Repository 标识。
// 对齐 proto CreateRepositoryResponse：repository_id。
// 支撑前端成功后进入新 Repository Binding Detail / Workspace。
type CreateRepositoryResponse struct {
	RepositoryID string `json:"repository_id"`
}

// BindRepositoryToProductRequest 绑定 Repository 到 Product 写入最小字段（phase04-04）。
// 对齐 proto BindRepositoryToProductRequest：repository_id / product_id。
// repository_id 在 RPC 请求中显式承接；在 HTTP 过渡层中由 URL 路径参数隐式承接。
type BindRepositoryToProductRequest struct {
	ProductID string `json:"product_id"`
}

// MapModuleToRepositoryRequest 映射 Module 到 Repository 写入最小字段（phase04-04）。
// 对齐 proto MapModuleToRepositoryRequest：repository_id / module_id。
// repository_id 在 RPC 请求中显式承接；在 HTTP 过渡层中由 URL 路径参数隐式承接。
type MapModuleToRepositoryRequest struct {
	ModuleID string `json:"module_id"`
}

// --- 列表查询参数 ---

// ListRepositoriesQuery 列表读取筛选参数（phase04-04）。
// status_filter 为空 / "all" 时表示不过滤（对应 UI/路由层的 all）。
type ListRepositoriesQuery struct {
	QueryText    string
	StatusFilter RepositoryStatus // 空 / "all" 表示不过滤
}
