#!/usr/bin/env bash
# reset_module_mainline.sh — phase02-12 验收辅助：模块主线清空与基线恢复
#
# 职责：
#   为 phase02-12 空状态验收与基线复验提供可重复执行的"清空 -> 恢复基线"入口。
#   - 默认模式：清空模块主线 + 恢复 phase02-12 验收基线数据
#   - --clean-only：仅清空模块主线（用于验证空状态后保持空状态）
#   - --restore-only：仅恢复基线数据（跳过清空，依赖 ON CONFLICT DO NOTHING 保证幂等）
#
# 清空范围：modules 及其级联表（module_releases / product_modules /
#   module_repositories / decision_links），不影响只读前提数据
#   （products / repositories / decisions）。
#
# 幂等：可重复执行。清空依赖 ON DELETE CASCADE；恢复使用 ON CONFLICT DO NOTHING。
#
# 上游规格：phase02-12 spec §"联调环境必须可重复建立"
# 依赖：seed_readonly_prereqs.sql 已执行（提供 products / repositories / decisions）
#
# 用法：
#   # 自动检测 psql（宿主机 psql 优先，否则回退到容器内 psql）
#   ./database/scripts/reset_module_mainline.sh                # 清空 + 恢复
#   ./database/scripts/reset_module_mainline.sh --clean-only   # 仅清空
#   ./database/scripts/reset_module_mainline.sh --restore-only # 仅恢复
#
# 可覆盖参数（带默认值）：
#   PG_HOST (默认 127.0.0.1)
#   PG_PORT (默认 55432)
#   PG_USER (默认 rento)
#   PSCO_DB (默认 psco_development)
#   PG_CONTAINER (默认 rento-preview-postgres)
#   SEEDS_DIR (默认脚本同级的 ../seeds)

set -euo pipefail

# —— 模式解析 ——
MODE="all"
if [[ $# -ge 1 ]]; then
  case "$1" in
    --clean-only)   MODE="clean" ;;
    --restore-only) MODE="restore" ;;
    --all|--reset)  MODE="all" ;;
    -h|--help)
      sed -n '2,/^$/p' "$0"
      exit 0
      ;;
    *)
      echo "[reset_module_mainline] 错误：未知参数 '$1'" >&2
      echo "[reset_module_mainline] 用法：$0 [--clean-only|--restore-only]" >&2
      exit 1
      ;;
  esac
fi

# —— 连接参数（允许环境变量覆盖）——
PG_HOST="${PG_HOST:-127.0.0.1}"
PG_PORT="${PG_PORT:-55432}"
PG_USER="${PG_USER:-rento}"
PSCO_DB="${PSCO_DB:-psco_development}"
PG_CONTAINER="${PG_CONTAINER:-rento-preview-postgres}"

# 种子文件目录默认指向脚本同级的 ../seeds
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
SEEDS_DIR="${SEEDS_DIR:-$SCRIPT_DIR/../seeds}"

# —— 解析 psql 执行方式（数组形式，正确处理参数边界）——
# 优先级：宿主机 psql > docker exec > podman exec
resolve_psql() {
  if command -v psql &>/dev/null; then
    PSQL=(psql -h "$PG_HOST" -p "$PG_PORT" -U "$PG_USER")
    PSQL_MODE="host"
    return
  fi
  if docker ps --format '{{.Names}}' 2>/dev/null | grep -qx "$PG_CONTAINER"; then
    PSQL=(docker exec -i "$PG_CONTAINER" psql -U "$PG_USER")
    PSQL_MODE="container"
    return
  fi
  if podman ps --format '{{.Names}}' 2>/dev/null | grep -qx "$PG_CONTAINER"; then
    PSQL=(podman exec -i "$PG_CONTAINER" psql -U "$PG_USER")
    PSQL_MODE="container"
    return
  fi
  echo "[reset_module_mainline] 错误：未找到 psql，且容器 '$PG_CONTAINER' 未运行。" >&2
  echo "[reset_module_mainline] 请安装 postgresql-client，或启动共享 PostgreSQL 容器。" >&2
  exit 1
}

resolve_psql

# —— 解析密码（仅宿主机模式需要）——
if [[ "$PSQL_MODE" == "host" ]]; then
  if [[ -z "${PG_PASSWORD:-}" ]]; then
    echo "[reset_module_mainline] PG_PASSWORD 未设置，尝试从容器读取..." >&2
    if docker ps --format '{{.Names}}' 2>/dev/null | grep -qx "$PG_CONTAINER"; then
      PG_PASSWORD="$(docker exec "$PG_CONTAINER" printenv POSTGRES_PASSWORD)"
    elif podman ps --format '{{.Names}}' 2>/dev/null | grep -qx "$PG_CONTAINER"; then
      PG_PASSWORD="$(podman exec "$PG_CONTAINER" printenv POSTGRES_PASSWORD)"
    else
      echo "[reset_module_mainline] 错误：无法获取密码，容器未运行且 PG_PASSWORD 未设置。" >&2
      exit 1
    fi
  fi
  export PGPASSWORD="${PG_PASSWORD:-}"
fi

# —— 前置校验：PSCO 数据库必须已存在 ——
exists=$("${PSQL[@]}" -d postgres \
  -tAc "SELECT 1 FROM pg_database WHERE datname = '$PSCO_DB'" 2>/dev/null || echo "")
if [[ "$exists" != "1" ]]; then
  echo "[reset_module_mainline] 错误：数据库 '$PSCO_DB' 不存在，请先执行 ./database/scripts/init_db.sh" >&2
  exit 1
fi

# —— 前置校验：只读前提数据已存在 ——
# 恢复阶段依赖 products / repositories / decisions 表中已有候选记录
if [[ "$MODE" == "all" || "$MODE" == "restore" ]]; then
  prereq_count=$("${PSQL[@]}" -d "$PSCO_DB" \
    -tAc "SELECT (SELECT COUNT(*) FROM products) + (SELECT COUNT(*) FROM repositories) + (SELECT COUNT(*) FROM decisions)" 2>/dev/null || echo "0")
  if [[ "$prereq_count" -lt 3 ]]; then
    echo "[reset_module_mainline] 错误：只读前提数据不足（products/repositories/decisions 总数 < 3）。" >&2
    echo "[reset_module_mainline] 请先执行 ./database/scripts/run_seeds.sh" >&2
    exit 1
  fi
fi

BASELINE_SQL="$SEEDS_DIR/seed_module_mainline_baseline.sql"

# —— 按模式执行 ——
case "$MODE" in
  clean)
    echo "[reset_module_mainline] 模式：仅清空模块主线..."
    "${PSQL[@]}" -d "$PSCO_DB" \
      -v ON_ERROR_STOP=1 \
      -c "DELETE FROM modules;"
    echo "[reset_module_mainline] 清空完成。当前模块数："
    "${PSQL[@]}" -d "$PSCO_DB" -tAc "SELECT COUNT(*) FROM modules;"
    ;;
  restore)
    if [[ ! -f "$BASELINE_SQL" ]]; then
      echo "[reset_module_mainline] 错误：基线 SQL 文件不存在：$BASELINE_SQL" >&2
      exit 1
    fi
    echo "[reset_module_mainline] 模式：仅恢复基线数据（跳过清空）..."
    # 跳过 DELETE FROM modules 行，依赖 ON CONFLICT DO NOTHING 保证幂等
    sed '/^DELETE FROM modules;/d' "$BASELINE_SQL" | \
      "${PSQL[@]}" -d "$PSCO_DB" -v ON_ERROR_STOP=1
    echo "[reset_module_mainline] 恢复完成。"
    ;;
  all)
    if [[ ! -f "$BASELINE_SQL" ]]; then
      echo "[reset_module_mainline] 错误：基线 SQL 文件不存在：$BASELINE_SQL" >&2
      exit 1
    fi
    echo "[reset_module_mainline] 模式：清空模块主线 + 恢复基线数据..."
    "${PSQL[@]}" -d "$PSCO_DB" \
      -v ON_ERROR_STOP=1 \
      < "$BASELINE_SQL"
    echo "[reset_module_mainline] 重置完成。"
    ;;
esac
