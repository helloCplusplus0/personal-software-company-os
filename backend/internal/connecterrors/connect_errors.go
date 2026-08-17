// Package connecterrors — Connect error 单值映射承接位。
//
// 本文件是 phase07-09 正式传输主线切换后，domain error → Connect code 的唯一映射入口。
// 各模块 Connect implementation 统一调用 MapToConnectError，不得各自维护独立错误映射表。
//
// 独立于 platform 包的原因：避免 platform（router.go）→ connect/* → platform 的循环导入。
//
// 映射规则对齐 phase07-07 formal spec §4.5：
//
//	资源不存在       → CodeNotFound
//	重复/冲突        → CodeAlreadyExists
//	非法输入         → CodeInvalidArgument
//	前置条件不满足    → CodeFailedPrecondition
//	内部错误         → CodeInternal
package connecterrors

import (
	"context"
	"errors"
	"strings"

	"connectrpc.com/connect"

	"github.com/psco/backend/internal/backup"
	"github.com/psco/backend/internal/dashboard"
	"github.com/psco/backend/internal/decisioncenter"
	"github.com/psco/backend/internal/export"
	"github.com/psco/backend/internal/governanceprofile"
	"github.com/psco/backend/internal/moduleregistry"
	"github.com/psco/backend/internal/onboarding"
	"github.com/psco/backend/internal/productregistry"
	"github.com/psco/backend/internal/projectcontext"
	"github.com/psco/backend/internal/repositorybinding"
	"github.com/psco/backend/internal/reusesummary"
	"github.com/psco/backend/internal/templatereuse"
)

// MapToConnectError 将 domain error 映射为 Connect error。
//
// 这是 phase07 正式传输主线的唯一错误映射入口。各模块 Connect implementation
// 必须统一调用此函数，不得各自维护独立错误映射表。
//
// 映射优先级：NotFound > AlreadyExists > InvalidArgument > FailedPrecondition > Internal。
// 每一层使用 errors.Is 进行哨兵匹配，支持 wrapped error。
func MapToConnectError(err error) *connect.Error {
	if isCancellationError(err) {
		return connect.NewError(connect.CodeCanceled, err)
	}

	if isDeadlineExceededError(err) {
		return connect.NewError(connect.CodeDeadlineExceeded, err)
	}

	// --- CodeNotFound ---
	if isAny(err,
		moduleregistry.ErrModuleNotFound,
		moduleregistry.ErrProductNotFound,
		moduleregistry.ErrRepositoryNotFound,
		decisioncenter.ErrDecisionNotFound,
		decisioncenter.ErrModuleNotFound,
		productregistry.ErrProductNotFound,
		productregistry.ErrModuleNotFound,
		repositorybinding.ErrRepositoryNotFound,
		repositorybinding.ErrProductNotFound,
		repositorybinding.ErrModuleNotFound,
		projectcontext.ErrRepositoryNotFound,
		governanceprofile.ErrRepositoryNotFound,
		governanceprofile.ErrGovernanceProfileNotFound,
	) {
		return connect.NewError(connect.CodeNotFound, err)
	}

	// --- CodeAlreadyExists ---
	if isAny(err,
		moduleregistry.ErrDuplicateModuleName,
		moduleregistry.ErrDuplicateBinding,
		moduleregistry.ErrDuplicateReleaseVersion,
		decisioncenter.ErrDuplicateLink,
		productregistry.ErrDuplicateBinding,
		repositorybinding.ErrDuplicateBinding,
		repositorybinding.ErrDuplicateMapping,
	) {
		return connect.NewError(connect.CodeAlreadyExists, err)
	}

	// --- CodeInvalidArgument ---
	if isAny(err,
		moduleregistry.ErrInvalidInput,
		moduleregistry.ErrInvalidStatus,
		moduleregistry.ErrInvalidReleaseStatus,
		decisioncenter.ErrInvalidInput,
		decisioncenter.ErrInvalidStatus,
		decisioncenter.ErrInvalidTargetType,
		decisioncenter.ErrInvalidAlternatives,
		decisioncenter.ErrInvalidStatusTransition,
		productregistry.ErrInvalidInput,
		productregistry.ErrInvalidStatus,
		repositorybinding.ErrInvalidInput,
		repositorybinding.ErrInvalidStatus,
		reusesummary.ErrInvalidScope,
		templatereuse.ErrInvalidInput,
		governanceprofile.ErrInvalidInput,
	) {
		return connect.NewError(connect.CodeInvalidArgument, err)
	}

	// --- CodeFailedPrecondition ---
	if isAny(err,
		productregistry.ErrModuleNotActive,
		repositorybinding.ErrProductNotActive,
		repositorybinding.ErrModuleNotActive,
		projectcontext.ErrRepositoryBindingIncomplete,
	) {
		return connect.NewError(connect.CodeFailedPrecondition, err)
	}

	// --- CodeInternal (兜底) ---
	// 包括：dashboard.ErrOverviewReadFailed, dashboard.ErrFeedbackSignalReadFailed,
	// dashboard.ErrRecentActivityReadFailed, onboarding.ErrFirstRunStateReadFailed,
	// export.ErrAssetReadFailed, export.ErrExportPersistFailed, export.ErrExportSnapshotReadFailed,
	// backup.ErrAssetReadFailed, backup.ErrBackupPersistFailed, backup.ErrBackupSnapshotReadFailed,
	// backup.ErrSchemaVersionReadFailed, reusesummary.ErrReuseSummaryReadFailed,
	// 以及所有未分类错误。
	_ = dashboard.ErrOverviewReadFailed
	_ = dashboard.ErrFeedbackSignalReadFailed
	_ = dashboard.ErrRecentActivityReadFailed
	_ = onboarding.ErrFirstRunStateReadFailed
	_ = export.ErrAssetReadFailed
	_ = export.ErrExportPersistFailed
	_ = export.ErrExportSnapshotReadFailed
	_ = backup.ErrAssetReadFailed
	_ = backup.ErrBackupPersistFailed
	_ = backup.ErrBackupSnapshotReadFailed
	_ = backup.ErrSchemaVersionReadFailed
	_ = reusesummary.ErrReuseSummaryReadFailed
	_ = projectcontext.ErrProjectContextReadFailed
	_ = governanceprofile.ErrGovernanceProfileReadFailed
	_ = governanceprofile.ErrGovernanceProfileSaveFailed

	return connect.NewError(connect.CodeInternal, err)
}

// isAny 检查 err 是否匹配任一哨兵错误（使用 errors.Is 支持 wrapped error）。
func isAny(err error, sentinels ...error) bool {
	for _, s := range sentinels {
		if errors.Is(err, s) {
			return true
		}
	}
	return false
}

func isCancellationError(err error) bool {
	if errors.Is(err, context.Canceled) {
		return true
	}

	message := strings.ToLower(err.Error())
	return strings.Contains(message, "context canceled") ||
		strings.Contains(message, "operation was canceled") ||
		strings.Contains(message, "request canceled")
}

func isDeadlineExceededError(err error) bool {
	if errors.Is(err, context.DeadlineExceeded) {
		return true
	}

	message := strings.ToLower(err.Error())
	return strings.Contains(message, "context deadline exceeded") ||
		strings.Contains(message, "deadline exceeded")
}
