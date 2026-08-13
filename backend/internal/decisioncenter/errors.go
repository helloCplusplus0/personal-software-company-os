// Package decisioncenter — 业务层错误定义。
//
// 这些哨兵错误由 service / repository 层返回，由 handler 层映射为 HTTP 状态码。
// 错误语义对齐 phase03-04 spec 冻结的全部错误语义与 phase03-10 §6.5 错误归属。
//
// 文件落点：backend/internal/decisioncenter/errors.go
package decisioncenter

import "errors"

// 业务错误哨兵值（覆盖 phase03-04 / phase03-10 §6.5 全部错误语义）。
var (
	// ErrDecisionNotFound Decision 不存在（资源不存在，404）。
	// 归属接口：DecisionDetailRead / DecisionLinkWrite。
	ErrDecisionNotFound = errors.New("decision not found")

	// ErrModuleNotFound Module 不存在（资源不存在，404）。
	// 归属接口：DecisionLinkWrite（关联目标不存在）。
	ErrModuleNotFound = errors.New("module not found")

	// ErrDuplicateLink 重复关联（重复冲突，409）。
	// 归属接口：DecisionLinkWrite。
	ErrDuplicateLink = errors.New("decision link already exists")

	// ErrInvalidInput 必填字段缺失或空白（校验失败，400）。
	// 归属接口：DecisionWrite（RecordDecision 必填校验）。
	ErrInvalidInput = errors.New("invalid input")

	// ErrInvalidStatus 非法 status 值（校验失败，400）。
	// 归属接口：DecisionWrite。
	ErrInvalidStatus = errors.New("invalid status")

	// ErrInvalidTargetType 目标类型越界，非 MODULE（校验失败，400）。
	// 归属接口：DecisionLinkWrite。
	ErrInvalidTargetType = errors.New("invalid target type")

	// ErrInvalidAlternatives alternatives 条目空白（校验失败，400）。
	// 归属接口：DecisionWrite。
	ErrInvalidAlternatives = errors.New("invalid alternatives: items must not be blank")

	// ErrInvalidStatusTransition 非法状态迁移（校验失败，400）。
	// 归属接口：UpdateDecisionStatus（fix_002_003）。
	ErrInvalidStatusTransition = errors.New("invalid status transition")
)
