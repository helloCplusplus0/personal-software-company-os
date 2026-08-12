// Package templatereuse 承载 Template Reuse 模块的跨层共享哨兵错误。
package templatereuse

import "errors"

var (
	// ErrInvalidInput 表示 Template Reuse 请求参数不满足已冻结的合同约束。
	ErrInvalidInput = errors.New("invalid template reuse input")
)
