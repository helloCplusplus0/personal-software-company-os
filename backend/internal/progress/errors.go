// Package progress — 业务层错误定义。
//
// 这些哨兵错误由 service / repository 层返回，由 Connect handler 统一
// 调用 connecterrors.MapToConnectError 映射为 Connect code。
package progress

import "errors"

// 业务错误哨兵值。
var (
	// ErrProgressEventNotFound 目标事件不存在（Delete 定位失败）。
	// 映射为 connect.CodeNotFound。
	ErrProgressEventNotFound = errors.New("progress event not found")

	// ErrRepositoryNotFound 读路径锚点仓库不存在（ListProgressEvents）。
	// 映射为 connect.CodeNotFound（读锚点语义，沿 GetProjectBrief）。
	ErrRepositoryNotFound = errors.New("repository not found")

	// ErrInvalidInput 事件输入违反校验规则（V1-V9 + envelope 前置 +
	// REPOSITORY_NOT_FOUND 跨模块引用校验统一包装）。
	// 映射为 connect.CodeInvalidArgument。
	ErrInvalidInput = errors.New("invalid progress input")

	// ErrProgressReadFailed 事件流读取失败。
	// 映射为 connect.CodeInternal。
	ErrProgressReadFailed = errors.New("progress read failed")

	// ErrProgressWriteFailed 事件写入/删除失败。
	// 映射为 connect.CodeInternal。
	ErrProgressWriteFailed = errors.New("progress write failed")
)
