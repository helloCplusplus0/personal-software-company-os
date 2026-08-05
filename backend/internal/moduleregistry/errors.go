// Package moduleregistry — 业务层错误定义。
//
// 这些哨兵错误由 service / repository 层返回，由 handler 层映射为 HTTP 状态码。
// 错误语义对齐 phase02-09 spec 的准入规则与接口分组约束。
package moduleregistry

import "errors"

// 业务错误哨兵值。
var (
	// ErrModuleNotFound 模块不存在。
	ErrModuleNotFound = errors.New("module not found")

	// ErrDuplicateModuleName 模块名称冲突（§5.6 名称唯一）。
	ErrDuplicateModuleName = errors.New("module name already exists")

	// ErrProductNotFound 候选 Product 不存在（绑定前提校验失败）。
	ErrProductNotFound = errors.New("product not found")

	// ErrRepositoryNotFound 候选 Repository 不存在（映射前提校验失败）。
	ErrRepositoryNotFound = errors.New("repository not found")

	// ErrDuplicateBinding 绑定关系已存在（避免重复绑定）。
	ErrDuplicateBinding = errors.New("binding already exists")

	// ErrDuplicateReleaseVersion 同一模块下版本号已存在（UNIQUE(module_id, version) 冲突）。
	ErrDuplicateReleaseVersion = errors.New("release version already exists for this module")

	// ErrInvalidStatus 无效的状态值。
	ErrInvalidStatus = errors.New("invalid status")

	// ErrInvalidReleaseStatus 无效的版本状态值。
	ErrInvalidReleaseStatus = errors.New("invalid release status")

	// ErrInvalidInput 输入字段缺失或非法（除名称冲突外的其他输入校验失败）。
	ErrInvalidInput = errors.New("invalid input")
)
