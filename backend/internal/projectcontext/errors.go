// Package projectcontext — 业务层错误定义。
//
// 这些哨兵错误由 service / candidate 层返回，由 Connect handler 映射为 Connect code。
package projectcontext

import "errors"

// 业务错误哨兵值。
var (
	// ErrRepositoryNotFound 目标仓库未在 PSCO 中完成登记/绑定。
	// 映射为 connect.CodeNotFound。
	ErrRepositoryNotFound = errors.New("repository not found in PSCO")

	// ErrRepositoryBindingIncomplete 目标仓库已存在，但尚未完成当前阶段要求的 Repository Binding。
	// 映射为 connect.CodeFailedPrecondition。
	ErrRepositoryBindingIncomplete = errors.New("repository binding incomplete in PSCO")

	// ErrProjectContextReadFailed 项目上下文聚合读取失败。
	// 映射为 connect.CodeInternal。
	ErrProjectContextReadFailed = errors.New("project context read failed")
)
