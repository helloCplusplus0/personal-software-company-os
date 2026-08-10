// Package productregistry 承载 Product Registry 后端模块的全部业务实现。
//
// 分层语义（对齐 phase04-07 / phase04-10 已冻结结论）：
//   - handler/     入口层：只负责承接 HTTP 请求与返回结果
//   - service/     业务编排层：动作语义、校验顺序与 reread 所需读取聚合
//   - repository/  数据访问层：products / product_modules / product_repositories 摘要读取
//   - candidate/   外部连接层：Module 候选读取（BindModuleToProduct 前提）
//
// 本文件定义跨层共享的 API 消息结构。
// 约束：消息结构从 proto/psco/product_registry/v1/product_registry.proto 单向派生或显式对齐，
// 不直接暴露存储模型，不在 types.go 或 handler DTO 中新增 .proto 中不存在的业务字段语义。
package productregistry

import (
	"time"

	"github.com/psco/backend/internal/moduleregistry"
)

// ProductStatus 冻结为 active / archived（phase04-04）。
// 与 psco.common.v1.ActiveArchivedStatus 单值对齐。
type ProductStatus string

const (
	ProductStatusActive   ProductStatus = "active"
	ProductStatusArchived ProductStatus = "archived"
)

// Product 核心对象（phase04-04 冻结的数据范围）。
// 对齐 proto Product：id / name / description / status / created_at。
type Product struct {
	ID          string        `json:"id"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Status      ProductStatus `json:"status"`
	CreatedAt   time.Time     `json:"created_at"`
}

// ProductListItem 列表读取模型（phase04-04）。
// 对齐 proto ProductListItem：
// id / name / description / status / created_at / module_bind_count / repository_bind_count。
type ProductListItem struct {
	ID                  string        `json:"id"`
	Name                string        `json:"name"`
	Description         string        `json:"description"`
	Status              ProductStatus `json:"status"`
	CreatedAt           time.Time     `json:"created_at"`
	ModuleBindCount     int           `json:"module_bind_count"`
	RepositoryBindCount int           `json:"repository_bind_count"`
}

// BoundModuleSummary 详情读取中的已绑定 Module 摘要。
// 对齐 proto BoundModuleSummary：module_id / module_name / module_status。
// module_status 复用 moduleregistry.ModuleStatus（跨包 import psco.module_registry.v1.ModuleStatus）。
type BoundModuleSummary struct {
	ModuleID     string                       `json:"module_id"`
	ModuleName   string                       `json:"module_name"`
	ModuleStatus moduleregistry.ModuleStatus  `json:"module_status"`
}

// BoundRepositorySummary 详情读取中的已绑定 Repository 摘要。
// 对齐 proto BoundRepositorySummary：repository_id / repository_name / provider / repository_status。
// repository_status 复用 ActiveArchivedStatus（psco.common.v1.ActiveArchivedStatus）。
type BoundRepositorySummary struct {
	RepositoryID     string        `json:"repository_id"`
	RepositoryName   string        `json:"repository_name"`
	Provider         string        `json:"provider"`
	RepositoryStatus ProductStatus `json:"repository_status"`
}

// ProductDetail 详情读取模型（phase04-04）。
// 对齐 proto ProductDetail：product + repeated BoundModuleSummary + repeated BoundRepositorySummary。
// 候选读取结果不内嵌于此，必须通过 ListProductModuleCandidates 独立承接。
type ProductDetail struct {
	Product           Product                  `json:"product"`
	BoundModules      []BoundModuleSummary     `json:"bound_modules"`
	BoundRepositories []BoundRepositorySummary `json:"bound_repositories"`
}

// ProductModuleCandidate Module 候选读取返回项（phase04-04）。
// 对齐 proto ProductModuleCandidate：module_id / module_name / module_status。
type ProductModuleCandidate struct {
	ModuleID     string                      `json:"module_id"`
	ModuleName   string                      `json:"module_name"`
	ModuleStatus moduleregistry.ModuleStatus `json:"module_status"`
}

// --- 写组请求体 ---

// CreateProductRequest 创建写入最小字段。
// 对齐 proto CreateProductRequest：name 为最小人工必填；
// description 为空时由系统默认补为 ""，status 为空时由系统默认补为 active。
type CreateProductRequest struct {
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Status      ProductStatus `json:"status"`
}

// CreateProductResponse 创建响应，只返回新建 Product 标识。
// 对齐 proto CreateProductResponse：product_id。
// 支撑前端成功后进入新 Product Detail。
type CreateProductResponse struct {
	ProductID string `json:"product_id"`
}

// BindModuleToProductRequest 绑定 Module 到 Product 写入最小字段（phase04-04）。
// 对齐 proto BindModuleToProductRequest：product_id / module_id。
// product_id 在 RPC 请求中显式承接；在 HTTP 过渡层中由 URL 路径参数隐式承接。
type BindModuleToProductRequest struct {
	ModuleID string `json:"module_id"`
}

// --- 列表查询参数 ---

// ListProductsQuery 列表读取筛选参数（phase04-04）。
// status_filter 为空 / "all" 时表示不过滤（对应 UI/路由层的 all）。
type ListProductsQuery struct {
	QueryText    string
	StatusFilter ProductStatus // 空 / "all" 表示不过滤
}
