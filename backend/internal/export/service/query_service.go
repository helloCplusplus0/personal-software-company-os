// Package service — Export 读编排层。
//
// QueryService 承接 GetExportSnapshot 读组，
// 对齐 phase06-14 spec §"Export 主线必须装配 9 类核心资产并形成可重复读取的快照"。
//
// 读取语义（phase06-14 spec §"GetExportSnapshot 读取语义"）：
//   - 若已存在 instance_exports 记录，返回最新一条快照摘要
//   - 若尚无历史导出记录，返回基于当前 canonical 数据现算的预览态 ExportSnapshot
//   - 预览态不得被误判为错误
//
// 文件落点：backend/internal/export/service/query_service.go
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

// QueryService 承接 Export 快照读取编排。
//
// 依赖通过 platform 装配点注入：
//   - store：instance_exports 表读写
//   - assetReader：9 类核心资产装配（用于预览态现算）
type QueryService struct {
	store       *repository.ExportStore
	assetReader *candidate.AssetReader
}

// NewQueryService 构造 QueryService。
func NewQueryService(store *repository.ExportStore, assetReader *candidate.AssetReader) *QueryService {
	return &QueryService{store: store, assetReader: assetReader}
}

// ReadExportSnapshot 读取最新导出快照。
//
// 语义（phase06-14 spec §"GetExportSnapshot 读取语义"）：
//   - 若已存在 instance_exports 记录，返回最新一条快照摘要
//   - 若尚无历史导出记录，返回基于当前 canonical 数据现算的预览态 ExportSnapshot
//   - 预览态不得被误判为错误
func (s *QueryService) ReadExportSnapshot(ctx context.Context) (*export.ExportSnapshot, error) {
	rec, err := s.store.ReadLatest(ctx)
	if err != nil {
		return nil, export.ErrExportSnapshotReadFailed
	}

	// 已有历史导出记录 → 返回最新快照
	if rec != nil {
		snapshot, err := buildSnapshotFromRecord(rec)
		if err != nil {
			return nil, export.ErrExportSnapshotReadFailed
		}
		return snapshot, nil
	}

	// 无历史记录 → 返回预览态快照（基于当前 canonical 数据现算）
	return s.buildPreviewSnapshot(ctx)
}

// buildSnapshotFromRecord 从数据库记录构建 ExportSnapshot。
func buildSnapshotFromRecord(rec *repository.ExportRecord) (*export.ExportSnapshot, error) {
	var scopes []string
	if err := json.Unmarshal(rec.AssetScopeJSON, &scopes); err != nil {
		return nil, fmt.Errorf("unmarshal asset_scope_json: %w", err)
	}

	assetScopes := make([]export.ExportAssetScope, 0, len(scopes))
	for _, s := range scopes {
		assetScopes = append(assetScopes, export.ExportAssetScope(s))
	}

	return &export.ExportSnapshot{
		AssetScope:    assetScopes,
		CreatedAt:     rec.CreatedAt,
		ResultStatus:  rec.ResultStatus,
		ResultSummary: rec.ResultSummary,
	}, nil
}

// buildPreviewSnapshot 基于当前 canonical 数据现算预览态快照。
//
// 预览态不持久化到 instance_exports，只返回当前资产覆盖矩阵与时间戳。
// result_status 固定为 success（预览态不是错误），result_summary 标注为预览态。
func (s *QueryService) buildPreviewSnapshot(ctx context.Context) (*export.ExportSnapshot, error) {
	// 预览态也需要装配 9 类资产以确认覆盖矩阵完整
	if _, err := s.assetReader.ReadCoreAssets(ctx); err != nil {
		return nil, export.ErrAssetReadFailed
	}

	scopes := export.AllExportAssetScopes()
	return &export.ExportSnapshot{
		AssetScope:    scopes,
		CreatedAt:     time.Now().UTC(),
		ResultStatus:  export.ExportResultStatusSuccess,
		ResultSummary: "preview: snapshot computed from current canonical data (no persisted export record)",
	}, nil
}
