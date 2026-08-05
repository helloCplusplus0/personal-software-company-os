#!/usr/bin/env bash
# run_seeds.sh — PSCO 种子数据统一执行入口
#
# 职责：
#   按 phase02-11 spec 要求，统一执行 database/seeds/ 下的种子文件，
#   为 phase02-12 联调验收提供最小候选数据与示例数据。
#
# 幂等：种子文件自身通过 ON CONFLICT / 守卫块保证可重复执行。
#
# 上游规格：phase02-11 spec §"迁移与最小种子数据必须可支撑 phase02-12 验收"
#
# 用法：
#   # 自动检测 psql（宿主机 psql 优先，否则回退到容器内 psql）
#   ./database/scripts/run_seeds.sh
#
#   # 同时执行 decision_link fixture（需已存在模块，否则自动跳过）
#   RUN_DECISION_LINK_FIXTURE=1 ./database/scripts/run_seeds.sh
#
# 可覆盖参数（带默认值）：
#   PG_HOST (默认 127.0.0.1)
#   PG_PORT (默认 55432)
#   PG_USER (默认 rento)
#   PSCO_DB (默认 psco_development)
#   PG_CONTAINER (默认 rento-preview-postgres)
#   SEEDS_DIR (默认脚本同级的 ../seeds)
#   RUN_DECISION_LINK_FIXTURE (默认 0)

set -euo pipefail

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
  echo "[run_seeds] 错误：未找到 psql，且容器 '$PG_CONTAINER' 未运行。" >&2
  echo "[run_seeds] 请安装 postgresql-client，或启动共享 PostgreSQL 容器。" >&2
  exit 1
}

resolve_psql

# —— 解析密码（仅宿主机模式需要）——
if [[ "$PSQL_MODE" == "host" ]]; then
  if [[ -z "${PG_PASSWORD:-}" ]]; then
    echo "[run_seeds] PG_PASSWORD 未设置，尝试从容器读取..." >&2
    if docker ps --format '{{.Names}}' 2>/dev/null | grep -qx "$PG_CONTAINER"; then
      PG_PASSWORD="$(docker exec "$PG_CONTAINER" printenv POSTGRES_PASSWORD)"
    elif podman ps --format '{{.Names}}' 2>/dev/null | grep -qx "$PG_CONTAINER"; then
      PG_PASSWORD="$(podman exec "$PG_CONTAINER" printenv POSTGRES_PASSWORD)"
    else
      echo "[run_seeds] 错误：无法获取密码，容器未运行且 PG_PASSWORD 未设置。" >&2
      exit 1
    fi
  fi
  export PGPASSWORD="${PG_PASSWORD:-}"
fi

# —— 前置校验：PSCO 数据库必须已存在 ——
exists=$("${PSQL[@]}" -d postgres \
  -tAc "SELECT 1 FROM pg_database WHERE datname = '$PSCO_DB'" 2>/dev/null || echo "")
if [[ "$exists" != "1" ]]; then
  echo "[run_seeds] 错误：数据库 '$PSCO_DB' 不存在，请先执行 ./database/scripts/init_db.sh" >&2
  exit 1
fi

# —— 前置校验：种子目录存在 ——
if [[ ! -d "$SEEDS_DIR" ]]; then
  echo "[run_seeds] 错误：种子目录不存在：$SEEDS_DIR" >&2
  exit 1
fi

# —— 执行只读前提种子（products / repositories / decisions）——
# 注意：使用 stdin 重定向而非 -f，以兼容容器模式（容器内无法访问宿主机文件路径）。
SEED_PREREQ="$SEEDS_DIR/seed_readonly_prereqs.sql"
if [[ -f "$SEED_PREREQ" ]]; then
  echo "[run_seeds] 执行 seed_readonly_prereqs.sql ..."
  "${PSQL[@]}" -d "$PSCO_DB" \
    -v ON_ERROR_STOP=1 \
    < "$SEED_PREREQ"
  echo "[run_seeds] seed_readonly_prereqs.sql 完成。"
else
  echo "[run_seeds] 警告：未找到 seed_readonly_prereqs.sql，跳过。" >&2
fi

# —— 可选：执行 decision_link fixture ——
# 该 fixture 需已存在模块，否则内部守卫块自动跳过（不报错）。
# 通过 RUN_DECISION_LINK_FIXTURE=1 显式启用。
if [[ "${RUN_DECISION_LINK_FIXTURE:-0}" == "1" ]]; then
  SEED_FIXTURE="$SEEDS_DIR/seed_decision_link_fixture.sql"
  if [[ -f "$SEED_FIXTURE" ]]; then
    echo "[run_seeds] 执行 seed_decision_link_fixture.sql ..."
    "${PSQL[@]}" -d "$PSCO_DB" \
      -v ON_ERROR_STOP=1 \
      < "$SEED_FIXTURE"
    echo "[run_seeds] seed_decision_link_fixture.sql 完成。"
  else
    echo "[run_seeds] 警告：未找到 seed_decision_link_fixture.sql，跳过。" >&2
  fi
fi

echo "[run_seeds] 全部种子数据执行完成。"
