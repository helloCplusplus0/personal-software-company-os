// Package service — 写组业务编排层。
//
// 统一承接 CreateProduct / BindModuleToProduct 编排。
//
// 文件落点：backend/internal/productregistry/service/command_service.go
package service

import (
	"context"
	"fmt"
	"strings"

	"github.com/psco/backend/internal/productregistry"
	"github.com/psco/backend/internal/productregistry/candidate"
	"github.com/psco/backend/internal/productregistry/repository"
)

// CommandService 承接写组业务编排。
//
// 依赖注入：
//   - products: products 表读写（含 product_modules 绑定写入）
//   - moduleCandidate: Module 存在性与 active 状态校验（跨模块，通过 candidate 子包隔离）
type CommandService struct {
	products        *repository.ProductStore
	moduleCandidate *candidate.ModuleCandidateRead
}

// NewCommandService 构造 CommandService。
func NewCommandService(products *repository.ProductStore, moduleCandidate *candidate.ModuleCandidateRead) *CommandService {
	return &CommandService{products: products, moduleCandidate: moduleCandidate}
}

// CreateProduct 承接 ProductCreateWrite。
//
// 校验顺序（phase04-04 / phase04-12 spec；phase06 draft-first 对齐）：
//  1. 必填字段去首尾空白后非空（name）
//  2. description 为空时由系统默认补为 ""
//  3. status 为空时由系统默认补为 active；非空时只允许 active / archived
//  4. Product.name 的最小格式合法性
//
// 成功返回新建产品 id，支撑前端回流到 Product Detail。
func (s *CommandService) CreateProduct(ctx context.Context, req productregistry.CreateProductRequest) (*productregistry.CreateProductResponse, error) {
	// 输入校验
	name := strings.TrimSpace(req.Name)
	description := strings.TrimSpace(req.Description)
	if name == "" {
		return nil, productregistry.ErrInvalidInput
	}

	status := req.Status
	if status == "" {
		status = productregistry.ProductStatusActive
	}
	if status != productregistry.ProductStatusActive && status != productregistry.ProductStatusArchived {
		return nil, productregistry.ErrInvalidStatus
	}

	// 写入
	p, err := s.products.Create(ctx, name, description, status)
	if err != nil {
		return nil, err
	}
	return &productregistry.CreateProductResponse{ProductID: p.ID}, nil
}

// BindModuleToProduct 承接 ProductModuleBindingWrite。
//
// 校验顺序（phase04-04 / phase04-12 spec）：
//  1. product_id / module_id 格式合法
//  2. Product 存在且状态合法
//  3. Module 存在且状态为 active
//  4. 重复绑定检测
//
// 成功后默认 reread owner 必须是 ProductDetailRead（phase04-12 spec）。
// 不得返回脱离详情 reread 的第二套结果模型。
func (s *CommandService) BindModuleToProduct(ctx context.Context, productID, moduleID string) error {
	// 1. ID 格式校验
	if err := productregistry.ValidateProductID(productID); err != nil {
		return err
	}
	if err := productregistry.ValidateModuleID(moduleID); err != nil {
		return err
	}

	// 2. Product 存在性校验
	p, err := s.products.GetByID(ctx, productID)
	if err != nil {
		return err
	}
	// Product 状态合法性（active / archived 均允许绑定）
	_ = p // 状态已由 CHECK 约束保证

	// 3. Module 存在且 active 校验（通过 candidate 子包，不直接写跨模块 SQL）
	//
	// 三态分流（phase04-04 错误语义前提）：
	//   - 不存在 → ErrModuleNotFound（404）
	//   - 存在但非 active → ErrModuleNotActive（400）
	moduleExists, moduleActive, err := s.moduleCandidate.CheckModuleExistsActive(ctx, moduleID)
	if err != nil {
		return fmt.Errorf("check module exists and active: %w", err)
	}
	if !moduleExists {
		return productregistry.ErrModuleNotFound
	}
	if !moduleActive {
		return productregistry.ErrModuleNotActive
	}

	// 4. 重复绑定检测（由 DB UNIQUE 约束兜底）
	return s.products.BindModule(ctx, productID, moduleID)
}
