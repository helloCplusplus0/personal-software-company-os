// Package productregistry — 业务层错误定义。
//
// 这些哨兵错误由 service / repository 层返回，由 handler 层映射为 HTTP 状态码。
// 错误语义对齐 phase04-04 spec 的错误语义前提与 phase04-10 正式规格正文。
package productregistry

import "errors"

// 业务错误哨兵值。
var (
	// ErrProductNotFound Product 不存在。
	ErrProductNotFound = errors.New("product not found")

	// ErrModuleNotFound 候选 Module 不存在（绑定前提校验失败）。
	ErrModuleNotFound = errors.New("module not found")

	// ErrDuplicateBinding 绑定关系已存在（避免重复绑定）。
	ErrDuplicateBinding = errors.New("binding already exists")

	// ErrInvalidStatus 无效的状态值。
	ErrInvalidStatus = errors.New("invalid status")

	// ErrInvalidInput 输入字段缺失或非法（除名称冲突外的其他输入校验失败）。
	ErrInvalidInput = errors.New("invalid input")

	// ErrModuleNotActive 候选 Module 状态非 active（绑定前提校验失败）。
	ErrModuleNotActive = errors.New("module is not active")
)
