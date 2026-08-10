// Package moduleregistry 承载 Module Registry 后端模块的全部业务实现。
//
// 分层语义（对齐 phase02-08 / 09 已冻结结论）：
//   - handler/  入口层：只负责承接 HTTP 请求与返回结果
//   - service/  业务编排层：动作语义、校验顺序与跨连接口编排
//   - repository/ 数据访问层：modules / module_releases / product_modules / module_repositories
//   - candidate/  外部连接层：Product / Repository 候选读取（phase02 临时承接）
//
// 本文件定义跨层共享的 API 消息结构。
// 约束（§6.5 合同与存储解耦）：消息结构独立于数据库表结构，不直接暴露存储模型。
package moduleregistry

import "time"

// ModuleStatus 冻结为 active / archived（§5.5）。
type ModuleStatus string

const (
	ModuleStatusActive   ModuleStatus = "active"
	ModuleStatusArchived ModuleStatus = "archived"
)

// ReleaseStatus 冻结为 active / archived。
type ReleaseStatus string

const (
	ReleaseStatusActive   ReleaseStatus = "active"
	ReleaseStatusArchived ReleaseStatus = "archived"
)

// Module 核心对象（§5.4）。
type Module struct {
	ID          string       `json:"id"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Status      ModuleStatus `json:"status"`
	CreatedAt   time.Time    `json:"created_at"`
}

// Release 核心对象（§5.4）。
type Release struct {
	ID         string        `json:"id"`
	ModuleID   string        `json:"module_id"`
	Version    string        `json:"version"`
	Status     ReleaseStatus `json:"status"`
	ReleasedAt time.Time     `json:"released_at"`
}

// ModuleListItem 列表读取模型（§5.7）。
// 至少承接 name / description / status / latest_release / product_bind_count / repository_bind_count。
type ModuleListItem struct {
	ID                  string       `json:"id"`
	Name                string       `json:"name"`
	Description         string       `json:"description"`
	Status              ModuleStatus `json:"status"`
	LatestRelease       *string      `json:"latest_release"`
	ProductBindCount    int          `json:"product_bind_count"`
	RepositoryBindCount int          `json:"repository_bind_count"`
}

// ProductBinding 详情读取中的产品绑定关系（剥离 module_id）。
type ProductBinding struct {
	ProductID   string `json:"product_id"`
	ProductName string `json:"product_name"`
}

// RepositoryMapping 详情读取中的仓库映射关系（剥离 module_id）。
type RepositoryMapping struct {
	RepositoryID   string `json:"repository_id"`
	RepositoryName string `json:"repository_name"`
}

// DecisionLink Decision 附属读取模型（§6.3 内嵌于 ModuleDetailRead，不设独立读接口）。
type DecisionLink struct {
	DecisionID    string `json:"decision_id"`
	DecisionTitle string `json:"decision_title"`
}

// ModuleDetail 详情读取模型（§5.7）。
// 统一承接模块核心字段、版本列表、产品绑定、仓库映射与相关 Decision 入口。
type ModuleDetail struct {
	Module             Module              `json:"module"`
	Releases           []Release           `json:"releases"`
	ProductBindings    []ProductBinding    `json:"product_bindings"`
	RepositoryMappings []RepositoryMapping `json:"repository_mappings"`
	DecisionLinks      []DecisionLink      `json:"decision_links"`
}

// ProductCandidate 旧 ProductBindingCandidateRead 的兼容返回项。
//
// phase04-12 起，该 DTO 仅服务 Module Detail 历史入口的兼容适配，
// 不再承接 canonical Product Registry / Repository Binding 主线。
type ProductCandidate struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// RepositoryCandidate 旧 RepositoryBindingCandidateRead 的兼容返回项。
//
// phase04-12 起，该 DTO 仅服务 Module Detail 历史入口的兼容适配，
// 不再承接 canonical Repository Binding 主线。
type RepositoryCandidate struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// --- 写组请求体 ---

// CreateModuleRequest §5.7 创建写入最小字段：
// name 为最小人工必填；description 为空时由系统默认补为 ""；
// status 为空时由系统默认补为 active。
type CreateModuleRequest struct {
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Status      ModuleStatus `json:"status"`
}

// CreateReleaseRequest §5.7 版本写入最小字段：version / status / released_at。
// module_id 由 URL 路径参数隐式承接，不放在请求体。
type CreateReleaseRequest struct {
	Version    string        `json:"version"`
	Status     ReleaseStatus `json:"status"`
	ReleasedAt string        `json:"released_at"` // RFC3339 字符串
}

// BindModuleToProductRequest §4.1 BindModuleToProduct。
type BindModuleToProductRequest struct {
	ProductID string `json:"product_id"`
}

// MapModuleToRepositoryRequest §4.1 MapModuleToRepository。
type MapModuleToRepositoryRequest struct {
	RepositoryID string `json:"repository_id"`
}

// --- 列表查询参数 ---

// ListQuery 列表读取筛选参数（§8.4 路由搜索参数层 queryText / statusFilter）。
type ListQuery struct {
	QueryText    string       // 模糊匹配 name / description
	StatusFilter ModuleStatus // 空 / "all" 表示不过滤
}
