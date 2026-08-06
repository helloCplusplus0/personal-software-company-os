// Package decisioncenter — 输入校验辅助。
//
// 用于 service 层在接收路径参数时，防止无效 ID 直接打到数据库导致 SQL 错误。
// 无效 ID 由调用方映射为对应的 Err*NotFound。
//
// 文件落点：backend/internal/decisioncenter/validate.go
package decisioncenter

import "github.com/google/uuid"

// validateUUID 校验字符串是否为有效 UUID 格式。
//
// 使用 google/uuid 解析，支持带连字符的标准 UUID 格式
// （PostgreSQL gen_random_uuid() 输出）。
func validateUUID(id string) bool {
	_, err := uuid.Parse(id)
	return err == nil
}

// ValidateDecisionID 校验 decision ID 是否为有效 UUID。无效时返回 ErrDecisionNotFound。
func ValidateDecisionID(id string) error {
	if !validateUUID(id) {
		return ErrDecisionNotFound
	}
	return nil
}

// ValidateModuleID 校验 module ID 是否为有效 UUID。无效时返回 ErrModuleNotFound。
//
// 用于 LinkDecisionToTarget 的目标 Module 存在性校验前提，
// 防止无效 module_id 直接打到数据库。
func ValidateModuleID(id string) error {
	if !validateUUID(id) {
		return ErrModuleNotFound
	}
	return nil
}
