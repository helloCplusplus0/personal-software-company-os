// Package candidate — Template Reuse 跨模块 reader 接口定义与实现（由 TemplateReuse 拥有）。
//
// phase09-08 spec §"candidate reader 的数据承接必须保持最小且单值"：
//   - reader 接口的定义与实现均由 TemplateReuse 模块 candidate/ 子包自己拥有
//   - TemplateReuse candidate/ 实现可以直接读取 products / modules / product_modules
//   - service 层不得直接写跨模块 SQL
//
// 本文件承接模板候选的读时派生与 templateCandidateId 解析。
// 文件落点：backend/internal/templatereuse/candidate/template_candidate_readers.go
package candidate

import (
	"context"
	"crypto/md5"
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// TemplateCandidateReaders 承接模板候选读时派生所需的跨模块 reader。
//
// 由 platform 装配点构造并注入到 TemplateReuse QueryService。
type TemplateCandidateReaders struct {
	pool *pgxpool.Pool
}

// NewTemplateCandidateReaders 构造 TemplateCandidateReaders。
func NewTemplateCandidateReaders(pool *pgxpool.Pool) *TemplateCandidateReaders {
	return &TemplateCandidateReaders{pool: pool}
}

// ProductModuleRow 单条 product-module 关联原始数据。
type ProductModuleRow struct {
	ProductID          string
	ProductName        string
	ProductDescription string
	ProductCreatedAt   time.Time
	ModuleID           string
	ModuleName         string
	CapabilityKey      *string
}

// TemplateCandidateData 模板候选派生结果。
type TemplateCandidateData struct {
	TemplateCandidateID         string
	TemplateTitle               string
	TemplateDescription         string
	Modules                     []TemplateModuleRefData
	SourceProductCount          int
	TotalReuseProductCount      int
	LatestSourceProductUpdatedAt time.Time
}

// TemplateModuleRefData 模板候选内的模块引用数据。
type TemplateModuleRefData struct {
	ModuleID        string
	ModuleName      string
	CapabilityKey   string
	CapabilityLabel string
}

// ReadAllProductModuleBindings 读取所有活跃 product-module 绑定关系。
//
// 只返回 status='active' 的 products 与 modules 之间的绑定关系，
// 用于后续在 service 层进行模板候选派生。
func (r *TemplateCandidateReaders) ReadAllProductModuleBindings(ctx context.Context) ([]ProductModuleRow, error) {
	rows, err := r.pool.Query(ctx, `
		SELECT
			p.id AS product_id,
			p.name AS product_name,
			COALESCE(p.description, '') AS product_description,
			p.created_at AS product_created_at,
			m.id AS module_id,
			m.name AS module_name,
			m.capability_key
		FROM product_modules pm
		INNER JOIN products p ON p.id = pm.product_id AND p.status = 'active'
		INNER JOIN modules m ON m.id = pm.module_id AND m.status = 'active'
		ORDER BY pm.product_id, pm.module_id`)
	if err != nil {
		return nil, fmt.Errorf("read all product module bindings: %w", err)
	}
	defer rows.Close()

	var results []ProductModuleRow
	for rows.Next() {
		var row ProductModuleRow
		var capabilityKey *string
		if err := rows.Scan(&row.ProductID, &row.ProductName, &row.ProductDescription,
			&row.ProductCreatedAt, &row.ModuleID, &row.ModuleName, &capabilityKey); err != nil {
			return nil, fmt.Errorf("scan product module row: %w", err)
		}
		row.CapabilityKey = capabilityKey
		results = append(results, row)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate product module rows: %w", err)
	}
	if results == nil {
		results = []ProductModuleRow{}
	}
	return results, nil
}

// ComputeTemplateCandidateID 从已排序的 module_id 集合派生 template_candidate_id。
//
// 派生规则（phase09-02 冻结）：
//   - 对 module_id 集合去重并按升序排序
//   - 以逗号连接生成 normalized key
//   - 对 normalized key 取 MD5 生成 template_candidate_id
func ComputeTemplateCandidateID(moduleIDs []string) string {
	sorted := make([]string, len(moduleIDs))
	copy(sorted, moduleIDs)
	sort.Strings(sorted)

	// 去重
	deduped := make([]string, 0, len(sorted))
	for i, id := range sorted {
		if i == 0 || id != sorted[i-1] {
			deduped = append(deduped, id)
		}
	}

	normalizedKey := strings.Join(deduped, ",")
	hash := md5.Sum([]byte(normalizedKey))
	return fmt.Sprintf("%x", hash)
}

// capabilityLabelMap 后端内置的单一 capability_key -> capability_label 映射。
//
// 与 ReuseSummary 模块保持一致的映射表（phase09-08 不新增第二套映射）。
var capabilityLabelMap = map[string]string{
	"web_frontend":     "Web Frontend",
	"backend_api":      "Backend API",
	"database":         "Database",
	"auth":             "Authentication",
	"cli_tool":         "CLI Tool",
	"build_system":     "Build System",
	"observability":    "Observability",
	"deployment":       "Deployment",
	"documentation":    "Documentation",
	"state_management": "State Management",
}

// CapabilityLabel 返回 capability_key 对应的 label，未在映射表中时回退为 key 本身。
func CapabilityLabel(key string) string {
	if label, ok := capabilityLabelMap[key]; ok {
		return label
	}
	return key
}