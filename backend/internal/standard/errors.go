// Package standard — 业务层错误定义。
//
// 这些哨兵错误由 service / repository 层返回，由 Connect handler 统一
// 调用 connecterrors.MapToConnectError 映射为 Connect code。
package standard

import "errors"

// 业务错误哨兵值。
var (
	// ErrStandardNotFound 目标规范不存在。
	// 映射为 connect.CodeNotFound。
	ErrStandardNotFound = errors.New("standard not found")

	// ErrBindingNotFound 目标绑定关系不存在（四元组定位失败）。
	// 映射为 connect.CodeNotFound。
	// 与 ErrStandardNotFound 区分：规范存在但绑定未建立。
	ErrBindingNotFound = errors.New("standard binding not found")

	// ErrInvalidInput 规范写读输入违反字段约束或树校验规则 R1-R8。
	// 映射为 connect.CodeInvalidArgument。
	ErrInvalidInput = errors.New("invalid standard input")

	// ErrStandardReadFailed 规范结构化读取失败。
	// 映射为 connect.CodeInternal。
	ErrStandardReadFailed = errors.New("standard read failed")

	// ErrStandardSaveFailed 规范保存失败（整树替换 + revision 追加事务失败，不写入半套状态）。
	// 映射为 connect.CodeInternal。
	ErrStandardSaveFailed = errors.New("standard save failed")
)
