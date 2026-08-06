#!/usr/bin/env bash
# reset_decision_mainline.sh — phase03-14 验收辅助：Decision 主线清空与基线恢复
#
# 职责：
#   为 phase03-14 空状态验收与基线复验提供可重复执行的"清空 -> 恢复基线"入口。
#   - 默认模式：清空 Decision 主线 + 恢复 phase03-14 验收基线数据
#   - --clean-only：仅清空 Decision 主线（用于验证空状态后保持空状态）
#   - --restore-only：仅恢复基线数据（跳过清空，依赖 ON CONFLICT DO NOTHING 保证幂等）
#
# 清空范围：decisions 及其级联表（decision_links），不影响只读前提数据
#   （products / repositories）与模块主线（modules / module_releases /
#   product_modules / module_repositories）。
#
# 幂等：可重复执行。清空依赖 ON DELETE CASCADE；恢复使用 ON CONFLICT DO NOTHING。
#
# 上游规格：phase03-09 spec §"Decision Center 重置脚本必须与 module_mainline 同构"
#           phase03-10 decision_center_spec_v0.1.md §11.2
# 依赖：reset_module_mainline.sh 已执行（提供 modules 基线，decision_links 依赖 modules）
#
# 用法：
#   # 自动检测 psql（宿主机 psql 优先，否则回退到容器内 psql）
#   ./database/scripts/reset_decision_mainline.sh                # 清空 + 恢复
#   ./database/scripts/reset_decision_mainline.sh --clean-only   # 仅清空
#   ./database/scripts/reset_decision_mainline.sh --restore-only # 仅恢复
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
      echo "[reset_decision_mainline] 错误：未知参数 '$1'" >&2
      echo "[reset_decision_mainline] 用法：$0 [--clean-only|--restore-only]" >&2
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
  echo "[reset_decision_mainline] 错误：未找到 psql，且容器 '$PG_CONTAINER' 未运行。" >&2
  echo "[reset_decision_mainline] 请安装 postgresql-client，或启动共享 PostgreSQL 容器。" >&2
  exit 1
}

resolve_psql

# —— 解析密码（仅宿主机模式需要）——
if [[ "$PSQL_MODE" == "host" ]]; then
  if [[ -z "${PG_PASSWORD:-}" ]]; then
    echo "[reset_decision_mainline] PG_PASSWORD 未设置，尝试从容器读取..." >&2
    if docker ps --format '{{.Names}}' 2>/dev/null | grep -qx "$PG_CONTAINER"; then
      PG_PASSWORD="$(docker exec "$PG_CONTAINER" printenv POSTGRES_PASSWORD)"
    elif podman ps --format '{{.Names}}' 2>/dev/null | grep -qx "$PG_CONTAINER"; then
      PG_PASSWORD="$(podman exec "$PG_CONTAINER" printenv POSTGRES_PASSWORD)"
    else
      echo "[reset_decision_mainline] 错误：无法获取密码，容器未运行且 PG_PASSWORD 未设置。" >&2
      exit 1
    fi
  fi
  export PGPASSWORD="${PG_PASSWORD:-}"
fi

# —— 前置校验：PSCO 数据库必须已存在 ——
exists=$("${PSQL[@]}" -d postgres \
  -tAc "SELECT 1 FROM pg_database WHERE datname = '$PSCO_DB'" 2>/dev/null || echo "")
if [[ "$exists" != "1" ]]; then
  echo "[reset_decision_mainline] 错误：数据库 '$PSCO_DB' 不存在，请先执行 ./database/scripts/init_db.sh" >&2
  exit 1
fi

# —— 前置校验：modules 基线数据必须已存在 ——
# 恢复阶段依赖 modules 表中已有基线记录（decision_links 依赖 modules）
if [[ "$MODE" == "all" || "$MODE" == "restore" ]]; then
  module_count=$("${PSQL[@]}" -d "$PSCO_DB" \
    -tAc "SELECT COUNT(*) FROM modules" 2>/dev/null || echo "0")
  if [[ "$module_count" -lt 1 ]]; then
    echo "[reset_decision_mainline] 错误：modules 基线数据不存在（COUNT < 1）。" >&2
    echo "[reset_decision_mainline] 请先执行 ./database/scripts/reset_module_mainline.sh" >&2
    exit 1
  fi
fi

BASELINE_SQL="$SEEDS_DIR/seed_decision_mainline_baseline.sql"

# —— 按模式执行 ——
case "$MODE" in
  clean)
    echo "[reset_decision_mainline] 模式：仅清空 Decision 主线..."
    "${PSQL[@]}" -d "$PSCO_DB" \
      -v ON_ERROR_STOP=1 \
      -c "DELETE FROM decisions;"
    echo "[reset_decision_mainline] 清空完成。当前 decisions / decision_links 计数："
    "${PSQL[@]}" -d "$PSCO_DB" -tAc "SELECT 'decisions', COUNT(*) FROM decisions UNION ALL SELECT 'decision_links', COUNT(*) FROM decision_links;"
    ;;
  restore)
    if [[ ! -f "$BASELINE_SQL" ]]; then
      echo "[reset_decision_mainline] 错误：基线 SQL 文件不存在：$BASELINE_SQL" >&2
      exit 1
    fi
    echo "[reset_decision_mainline] 模式：仅恢复基线数据（跳过清空）..."
    # 跳过 DELETE FROM decisions 行，依赖 UPDATE 收口 + WHERE NOT EXISTS + ON CONFLICT 保证幂等
    sed '/^DELETE FROM decisions;/d' "$BASELINE_SQL" | \
      "${PSQL[@]}" -d "$PSCO_DB" -v ON_ERROR_STOP=1
    echo "[reset_decision_mainline] 恢复完成。"
    ;;
  all)
    if [[ ! -f "$BASELINE_SQL" ]]; then
      echo "[reset_decision_mainline] 错误：基线 SQL 文件不存在：$BASELINE_SQL" >&2
      exit 1
    fi
    echo "[reset_decision_mainline] 模式：清空 Decision 主线 + 恢复基线数据..."
    "${PSQL[@]}" -d "$PSCO_DB" \
      -v ON_ERROR_STOP=1 \
      < "$BASELINE_SQL"
    echo "[reset_decision_mainline] 重置完成。"
    ;;
esac
