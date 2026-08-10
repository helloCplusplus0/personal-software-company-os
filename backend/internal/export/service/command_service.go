// Package service — Export 写编排层。
//
// CommandService 承接 ExportCoreAssets 写组，
// 对齐 phase06-14 spec §"Export 主线必须装配 9 类核心资产并形成可重复读取的快照"。
//
// 写入语义（phase06-14 spec §"ExportCoreAssets 数据装配"）：
//   - candidate.AssetReader 必须装配 9 类 canonical 数据
//   - command_service 必须把装配结果持久化到 instance_exports
//   - 不得只导出主实体而遗漏绑定关系
//
// 文件落点：backend/internal/export/service/command_service.go
package service

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/psco/backend/internal/export"
	"github.com/psco/backend/internal/export/candidate"
	"github.com/psco/backend/internal/export/repository"
)

// CommandService 承接 Export 导出执行编排。
//
// 依赖通过 platform 装配点注入：
//   - store：instance_exports 表读写
//   - assetReader：9 类核心资产装配
type CommandService struct {
	store       *repository.ExportStore
	assetReader *candidate.AssetReader
}

// NewCommandService 构造 CommandService。
func NewCommandService(store *repository.ExportStore, assetReader *candidate.AssetReader) *CommandService {
	return &CommandService{store: store, assetReader: assetReader}
}

// ExportCoreAssets 装配 9 类核心资产并持久化导出快照。
//
// 流程（phase06-14 spec §"ExportCoreAssets 数据装配"）：
//  1. 通过 candidate.AssetReader 装配 9 类 canonical 数据
//  2. 组装 asset_scope（9 类资产覆盖矩阵）
//  3. 组装 artifact_payload（9 类资产完整数据载荷）
//  4. 持久化到 instance_exports
//  5. 返回写入后的 ExportSnapshot
func (s *CommandService) ExportCoreAssets(ctx context.Context) (*export.ExportSnapshot, error) {
	// 1. 装配 9 类核心资产
	payload, err := s.assetReader.ReadCoreAssets(ctx)
	if err != nil {
		return nil, export.ErrAssetReadFailed
	}

	// 2. 组装 asset_scope JSON（9 类资产覆盖矩阵）
	scopes := export.AllExportAssetScopes()
	scopeStrings := make([]string, 0, len(scopes))
	for _, s := range scopes {
		scopeStrings = append(scopeStrings, string(s))
	}
	assetScopeJSON, err := json.Marshal(scopeStrings)
	if err != nil {
		return nil, fmt.Errorf("marshal asset_scope: %w", err)
	}

	// 3. 组装 artifact_payload JSON（9 类资产完整数据载荷）
	artifactPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("marshal artifact_payload: %w", err)
	}

	// 4. 持久化到 instance_exports
	resultSummary := fmt.Sprintf("exported %d core asset scopes at %s", len(scopes), time.Now().UTC().Format(time.RFC3339))
	rec, err := s.store.CreateLatest(ctx, export.ExportResultStatusSuccess, resultSummary, assetScopeJSON, artifactPayload)
	if err != nil {
		return nil, export.ErrExportPersistFailed
	}

	// 5. 返回写入后的 ExportSnapshot
	return &export.ExportSnapshot{
		AssetScope:    scopes,
		CreatedAt:     rec.CreatedAt,
		ResultStatus:  rec.ResultStatus,
		ResultSummary: rec.ResultSummary,
	}, nil
}
