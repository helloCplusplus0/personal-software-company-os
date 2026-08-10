// Package candidate — Export 跨模块 reader 接口定义与实现（由 Export 拥有）。
//
// phase06-14 spec §"Phase06 后端模块必须按现有主线结构落地"：
//   - reader 接口的定义与实现均由 Export 模块 candidate/ 子包自己拥有
//   - Export candidate/ 实现可以直接读取 canonical 模块的表，但必须在 candidate/ 子包内隔离
//   - Export service/ 层不得直接跨模块写 SQL
//
// 本文件承接 Export 所需的 9 类核心资产装配 reader。
// 文件落点：backend/internal/export/candidate/asset_reader.go
package candidate

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

// AssetReader 承接 Export 所需的 9 类核心资产装配。
//
// 由 platform 装配点构造并注入到 Export QueryService / CommandService。
// 9 类核心资产对齐 phase06-12 §"Export / Backup 正式语义与边界冻结"的最小覆盖矩阵。
type AssetReader struct {
	pool *pgxpool.Pool
}

// NewAssetReader 构造 AssetReader。
func NewAssetReader(pool *pgxpool.Pool) *AssetReader {
	return &AssetReader{pool: pool}
}

// AssetPayload 9 类核心资产的完整数据载荷。
// 每个字段以 JSON 原始字节承接，由 service 层统一打包为 artifact_payload_json。
// 空表返回 'null'（JSON null），不返回空数组，以区分"表存在但无数据"与"读取失败"。
type AssetPayload struct {
	Products            json.RawMessage `json:"products"`
	Modules             json.RawMessage `json:"modules"`
	Releases            json.RawMessage `json:"releases"`
	Repositories        json.RawMessage `json:"repositories"`
	Decisions           json.RawMessage `json:"decisions"`
	DecisionLinks       json.RawMessage `json:"decision_links"`
	ProductModules      json.RawMessage `json:"product_modules"`
	ProductRepositories json.RawMessage `json:"product_repositories"`
	ModuleRepositories  json.RawMessage `json:"module_repositories"`
}

// assetEntry 单类资产的表名与目标字段映射。
type assetEntry struct {
	table  string
	target *json.RawMessage
}

// ReadCoreAssets 装配 9 类核心资产的完整数据载荷。
//
// 覆盖矩阵（phase06-14 spec §"ExportCoreAssets 数据装配"）：
//   - products / modules / releases / repositories / decisions
//   - decision_links / product_modules / product_repositories / module_repositories
//
// 不得只导出主实体而遗漏绑定关系。
// 任一资产读取失败时返回 error，由 service 层归一化为 export.ErrAssetReadFailed。
//
// 每类资产使用 json_agg(row_to_json(t)) 聚合为 JSON 数组，
// 空表通过 COALESCE 返回 'null'，不返回空数组。
func (r *AssetReader) ReadCoreAssets(ctx context.Context) (*AssetPayload, error) {
	payload := &AssetPayload{}

	entries := []assetEntry{
		{"products", &payload.Products},
		{"modules", &payload.Modules},
		{"module_releases", &payload.Releases},
		{"repositories", &payload.Repositories},
		{"decisions", &payload.Decisions},
		{"decision_links", &payload.DecisionLinks},
		{"product_modules", &payload.ProductModules},
		{"product_repositories", &payload.ProductRepositories},
		{"module_repositories", &payload.ModuleRepositories},
	}

	for _, e := range entries {
		// 单行单列 JSON 聚合查询，使用 QueryRow 直接 Scan
		// 注意：json_agg 返回 json 类型，COALESCE 后备值必须使用 'null'::json（非 jsonb）以避免类型不匹配
		query := fmt.Sprintf(`SELECT COALESCE(json_agg(row_to_json(t)), 'null'::json) FROM (SELECT * FROM %s) t`, e.table)
		var raw json.RawMessage
		if err := r.pool.QueryRow(ctx, query).Scan(&raw); err != nil {
			return nil, fmt.Errorf("read %s: %w", e.table, err)
		}
		*e.target = raw
	}

	return payload, nil
}
