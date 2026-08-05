// Package moduleregistry — 输入校验辅助。
package moduleregistry

import (
	"github.com/google/uuid"
)

// validateUUID 校验字符串是否为有效 UUID 格式。
//
// 用于 service 层在接收路径参数时，防止无效 ID 直接打到数据库导致 SQL 错误。
// 无效 ID 由调用方映射为对应的 Err*NotFound。
//
// 使用 google/uuid 解析，支持带连字符的标准 UUID 格式（PostgreSQL gen_random_uuid() 输出）。
func validateUUID(id string) bool {
	_, err := uuid.Parse(id)
	return err == nil
}

// ValidateModuleID 校验 module ID 是否为有效 UUID。无效时返回 ErrModuleNotFound。
func ValidateModuleID(id string) error {
	if !validateUUID(id) {
		return ErrModuleNotFound
	}
	return nil
}

// ValidateProductID 校验 product ID 是否为有效 UUID。无效时返回 ErrProductNotFound。
func ValidateProductID(id string) error {
	if !validateUUID(id) {
		return ErrProductNotFound
	}
	return nil
}

// ValidateRepositoryID 校验 repository ID 是否为有效 UUID。无效时返回 ErrRepositoryNotFound。
func ValidateRepositoryID(id string) error {
	if !validateUUID(id) {
		return ErrRepositoryNotFound
	}
	return nil
}
