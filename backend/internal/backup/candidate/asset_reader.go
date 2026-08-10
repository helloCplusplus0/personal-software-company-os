// Package candidate — Backup 跨模块 reader 接口定义与实现（由 Backup 拥有）。
//
// phase06-14 spec §"Phase06 后端模块必须按现有主线结构落地"：
//   - reader 接口的定义与实现均由 Backup 模块 candidate/ 子包自己拥有
//   - Backup 与 Export 模块边界互不耦合（phase06-08 模块分离冻结），
//     因此 Backup 拥有独立的 AssetReader，不 import Export 模块
//   - Backup candidate/ 实现可以直接读取 canonical 模块的表
//   - Backup service/ 层不得直接跨模块写 SQL
//
// 本文件承接 Backup 所需的 9 类核心资产装配 reader 与 schema 版本 reader。
// 文件落点：backend/internal/backup/candidate/asset_reader.go
package candidate

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

// AssetReader 承接 Backup 所需的 9 类核心资产装配。
//
// 由 platform 装配点构造并注入到 Backup CommandService。
// 9 类核心资产对齐 phase06-12 §"Export / Backup 正式语义与边界冻结"的最小覆盖矩阵。
type AssetReader struct {
	pool *pgxpool.Pool
}

// NewAssetReader 构造 AssetReader。
func NewAssetReader(pool *pgxpool.Pool) *AssetReader {
	return &AssetReader{pool: pool}
}

// AssetPayload 9 类核心资产的完整数据载荷。
// 每个字段以 JSON 原始字节承接，由 service 层统一打包为 backup_payload_json。
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
// 覆盖矩阵（phase06-14 spec §"CreateInstanceBackup 持久化语义"）：
//   - products / modules / releases / repositories / decisions
//   - decision_links / product_modules / product_repositories / module_repositories
//
// 任一资产读取失败时返回 error，由 service 层归一化为 backup.ErrAssetReadFailed。
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

// ReadLatestSchemaVersion 读取 schema_migrations 表中的最新版本号。
// 用于备份写入时记录 schema_version，以及 read / verify 子路径校验 schema 前提。
// 版本号按字符串降序排序取第一条（migration 文件名格式为 0001_xxx，自然字典序即时间序）。
func (r *AssetReader) ReadLatestSchemaVersion(ctx context.Context) (string, error) {
	var version string
	err := r.pool.QueryRow(ctx,
		`SELECT version FROM schema_migrations ORDER BY version DESC LIMIT 1`).Scan(&version)
	if err != nil {
		return "", fmt.Errorf("read latest schema_version: %w", err)
	}
	return version, nil
}
