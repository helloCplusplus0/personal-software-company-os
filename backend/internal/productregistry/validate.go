// Package productregistry — 输入校验辅助。
package productregistry

import (
	"github.com/google/uuid"
)

// validateUUID 校验字符串是否为有效 UUID 格式。
//
// 用于 service 层在接收路径参数时，防止无效 ID 直接打到数据库导致 SQL 错误。
// 无效 ID 由调用方映射为对应的 Err*NotFound。
func validateUUID(id string) bool {
	_, err := uuid.Parse(id)
	return err == nil
}

// ValidateProductID 校验 product ID 是否为有效 UUID。无效时返回 ErrProductNotFound。
func ValidateProductID(id string) error {
	if !validateUUID(id) {
		return ErrProductNotFound
	}
	return nil
}

// ValidateModuleID 校验 module ID 是否为有效 UUID。无效时返回 ErrModuleNotFound。
func ValidateModuleID(id string) error {
	if !validateUUID(id) {
		return ErrModuleNotFound
	}
	return nil
}
