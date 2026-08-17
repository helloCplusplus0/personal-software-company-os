// Package governanceprofile — 业务层错误定义。
//
// 这些哨兵错误由 service / repository 层返回，由 Connect handler 统一
// 调用 connecterrors.MapToConnectError 映射为 Connect code。
package governanceprofile

import "errors"

// 业务错误哨兵值。
var (
	// ErrRepositoryNotFound 目标仓库未在 PSCO 中完成登记。
	// 映射为 connect.CodeNotFound。
	ErrRepositoryNotFound = errors.New("repository not found in PSCO")

	// ErrGovernanceProfileNotFound 目标仓库尚未建立治理画像主记录。
	// 映射为 connect.CodeNotFound。
	// 与 ErrRepositoryNotFound 区分：仓库存在但画像未创建。
	ErrGovernanceProfileNotFound = errors.New("governance profile not created for repository")

	// ErrInvalidInput 治理画像保存输入违反字段分类约束或 8 项资产矩阵。
	// 映射为 connect.CodeInvalidArgument。
	ErrInvalidInput = errors.New("invalid governance profile input")

	// ErrGovernanceProfileReadFailed 治理画像结构化读取失败。
	// 映射为 connect.CodeInternal。
	ErrGovernanceProfileReadFailed = errors.New("governance profile read failed")

	// ErrGovernanceProfileSaveFailed 治理画像保存失败（整体事务失败，不写入半套状态）。
	// 映射为 connect.CodeInternal。
	ErrGovernanceProfileSaveFailed = errors.New("governance profile save failed")
)
