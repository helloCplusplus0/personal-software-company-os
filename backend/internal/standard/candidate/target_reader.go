// Package candidate — standard 跨模块只读 reader。
//
// 跨模块读取全部通过本子包隔离，service 层不直接写跨模块 SQL
// （对齐工程约定：候选/前提读取接口由消费方模块的 candidate 子包定义和拥有）。
//
// 文件落点：backend/internal/standard/candidate/target_reader.go
package candidate

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/psco/backend/internal/standard"
)

// TargetReader 承接绑定目标实体（repository / product / decision / module）的存在性校验。
// 多态 target_id 无 DB 外键（0011 迁移冻结），存在性由应用层按 target_type 查对应实体表。
type TargetReader struct {
	pool *pgxpool.Pool
}

// NewTargetReader 构造 TargetReader。
func NewTargetReader(pool *pgxpool.Pool) *TargetReader {
	return &TargetReader{pool: pool}
}

// EnsureTargetExists 校验绑定目标实体已存在。
//
// 失败语义（phase14-04 冻结）：
//   - 未知 targetType → ErrInvalidInput（InvalidArgument）
//   - target 不存在   → ErrInvalidInput（InvalidArgument，target 不存在归 InvalidArgument）
//   - 查询失败        → 包装原始错误（Internal，由 connecterrors 兜底）
func (r *TargetReader) EnsureTargetExists(ctx context.Context, targetType standard.BindingTargetType, targetID string) error {
	// 目标实体表按受控枚举映射（常量拼接，无注入面）
	var table string
	switch targetType {
	case standard.BindingTargetRepository:
		table = "repositories"
	case standard.BindingTargetProduct:
		table = "products"
	case standard.BindingTargetDecision:
		table = "decisions"
	case standard.BindingTargetModule:
		table = "modules"
	default:
		return fmt.Errorf("%w: unknown target type %q", standard.ErrInvalidInput, targetType)
	}

	var exists bool
	err := r.pool.QueryRow(ctx, fmt.Sprintf(`
		SELECT EXISTS(SELECT 1 FROM %s WHERE id = $1)
	`, table), targetID).Scan(&exists)
	if err != nil {
		return fmt.Errorf("standard: check target exists: %w", err)
	}
	if !exists {
		return fmt.Errorf("%w: target %s %s not found", standard.ErrInvalidInput, targetType, targetID)
	}
	return nil
}
