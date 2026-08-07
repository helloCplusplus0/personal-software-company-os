// Package productregistry — 跨模块关系摘要读取消费侧接口定义。
//
// ProductRepositorySummaryRead 的 owner 是 Repository Binding（phase04-07 §"跨模块已绑定关系摘要读取边界冻结" L162-181）。
// Product Registry 通过本接口注入消费，不得在 productregistry/repository/ 内复制第二套实现，
// 也不得在 service/ 或 handler/ 内直接写跨模块 SQL 读取 product_repositories 表。
//
// 文件落点：backend/internal/productregistry/bound_repository_reader.go
package productregistry

import "context"

// BoundRepositoryReader 承接 ProductRepositorySummaryRead 的消费侧接口。
//
// 实现由 Repository Binding 模块的 repository.BindingStore 提供（隐式实现），
// 在 backend/internal/platform/ 装配点注入到 Product Registry 的 QueryService。
//
// 接口归属 Product Registry（消费方），实现归属 Repository Binding（owner），
// 依赖方向为 repositorybinding/repository → productregistry，单向无循环。
type BoundRepositoryReader interface {
	// ListBoundRepositoriesByProduct 读取指定产品的已绑定 Repository 摘要列表。
	//
	// 语义：给定 productId，返回已绑定的 Repository 摘要。
	// 排序按 repositories.name 升序（与 RepositoryProductSummaryRead 一致）。
	ListBoundRepositoriesByProduct(ctx context.Context, productID string) ([]BoundRepositorySummary, error)
}
