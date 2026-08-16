#!/usr/bin/env bash
# restore_phase11_phase12_dogfooding_sample.sh
# phase11 / phase12 固定 dogfooding 样本恢复入口
#
# 职责：
#   以与现有 reset/seed 脚本一致的 psql 解析方式，
#   执行 seed_phase11_phase12_dogfooding_sample.sql，
#   恢复 phase11 / phase12 验收冻结的 PSCO 固定样本。
#
# 范围：
#   - 不清空现有开发数据
#   - 仅补齐固定样本与 canonical 关系
#   - 若同名异 id 冲突，由 SQL 显式失败
#
# 用法：
#   ./database/scripts/restore_phase11_phase12_dogfooding_sample.sh
#
# 可覆盖参数（带默认值）：
#   PG_HOST (默认 127.0.0.1)
#   PG_PORT (默认 55432)
#   PG_USER (默认 rento)
#   PSCO_DB (默认 psco_development)
#   PG_CONTAINER (默认 rento-preview-postgres)
#   SEEDS_DIR (默认脚本同级的 ../seeds)

set -euo pipefail

PG_HOST="${PG_HOST:-127.0.0.1}"
PG_PORT="${PG_PORT:-55432}"
PG_USER="${PG_USER:-rento}"
PSCO_DB="${PSCO_DB:-psco_development}"
PG_CONTAINER="${PG_CONTAINER:-rento-preview-postgres}"

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
SEEDS_DIR="${SEEDS_DIR:-$SCRIPT_DIR/../seeds}"
SEED_FILE="$SEEDS_DIR/seed_phase11_phase12_dogfooding_sample.sql"

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
  echo "[restore_phase11_phase12_dogfooding_sample] 错误：未找到 psql，且容器 '$PG_CONTAINER' 未运行。" >&2
  echo "[restore_phase11_phase12_dogfooding_sample] 请安装 postgresql-client，或启动共享 PostgreSQL 容器。" >&2
  exit 1
}

resolve_psql

if [[ "$PSQL_MODE" == "host" ]]; then
  if [[ -z "${PG_PASSWORD:-}" ]]; then
    echo "[restore_phase11_phase12_dogfooding_sample] PG_PASSWORD 未设置，尝试从容器读取..." >&2
    if docker ps --format '{{.Names}}' 2>/dev/null | grep -qx "$PG_CONTAINER"; then
      PG_PASSWORD="$(docker exec "$PG_CONTAINER" printenv POSTGRES_PASSWORD)"
    elif podman ps --format '{{.Names}}' 2>/dev/null | grep -qx "$PG_CONTAINER"; then
      PG_PASSWORD="$(podman exec "$PG_CONTAINER" printenv POSTGRES_PASSWORD)"
    else
      echo "[restore_phase11_phase12_dogfooding_sample] 错误：无法获取密码，容器未运行且 PG_PASSWORD 未设置。" >&2
      exit 1
    fi
  fi
  export PGPASSWORD="${PG_PASSWORD:-}"
fi

if [[ ! -f "$SEED_FILE" ]]; then
  echo "[restore_phase11_phase12_dogfooding_sample] 错误：seed 文件不存在：$SEED_FILE" >&2
  exit 1
fi

exists=$("${PSQL[@]}" -d postgres -tAc "SELECT 1 FROM pg_database WHERE datname = '$PSCO_DB'" 2>/dev/null || echo "")
if [[ "$exists" != "1" ]]; then
  echo "[restore_phase11_phase12_dogfooding_sample] 错误：数据库 '$PSCO_DB' 不存在，请先执行 ./database/scripts/init_db.sh" >&2
  exit 1
fi

echo "[restore_phase11_phase12_dogfooding_sample] 恢复 phase11 / phase12 固定样本..."
"${PSQL[@]}" -d "$PSCO_DB" -v ON_ERROR_STOP=1 < "$SEED_FILE"
echo "[restore_phase11_phase12_dogfooding_sample] 恢复完成。"
