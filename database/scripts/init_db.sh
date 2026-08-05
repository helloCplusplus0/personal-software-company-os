#!/usr/bin/env bash
# init_db.sh — PSCO 独立数据库初始化入口
#
# 职责：
#   在本地共享 PostgreSQL 容器（rento-preview-postgres）中创建 PSCO 独立数据库，
#   不新开第二个数据库容器，不复用 rento_production。
#
# 幂等：数据库已存在时跳过，可重复执行。
#
# 上游规格：phase02-11 spec §"本地开发环境必须复用共享 PostgreSQL 容器"
#
# 用法：
#   # 自动检测 psql（宿主机 psql 优先，否则回退到容器内 psql）
#   ./database/scripts/init_db.sh
#
#   # 显式提供密码（仅宿主机 psql 模式需要）
#   PG_PASSWORD=<password> ./database/scripts/init_db.sh
#
# 可覆盖参数（带默认值）：
#   PG_HOST (默认 127.0.0.1)
#   PG_PORT (默认 55432)
#   PG_USER (默认 rento)
#   PSCO_DB (默认 psco_development)
#   PG_CONTAINER (默认 rento-preview-postgres)

set -euo pipefail

# —— 连接参数（允许环境变量覆盖）——
PG_HOST="${PG_HOST:-127.0.0.1}"
PG_PORT="${PG_PORT:-55432}"
PG_USER="${PG_USER:-rento}"
PSCO_DB="${PSCO_DB:-psco_development}"
PG_CONTAINER="${PG_CONTAINER:-rento-preview-postgres}"

# —— 解析 psql 执行方式（数组形式，正确处理参数边界）——
# 优先级：宿主机 psql > docker exec > podman exec
# 宿主机模式通过 TCP 连接映射端口；容器模式通过 socket 连接内部实例。
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
  echo "[init_db] 错误：未找到 psql，且容器 '$PG_CONTAINER' 未运行。" >&2
  echo "[init_db] 请安装 postgresql-client，或启动共享 PostgreSQL 容器。" >&2
  exit 1
}

resolve_psql

# —— 解析密码（仅宿主机模式需要）——
if [[ "$PSQL_MODE" == "host" ]]; then
  if [[ -z "${PG_PASSWORD:-}" ]]; then
    echo "[init_db] PG_PASSWORD 未设置，尝试从容器读取..." >&2
    if docker ps --format '{{.Names}}' 2>/dev/null | grep -qx "$PG_CONTAINER"; then
      PG_PASSWORD="$(docker exec "$PG_CONTAINER" printenv POSTGRES_PASSWORD)"
    elif podman ps --format '{{.Names}}' 2>/dev/null | grep -qx "$PG_CONTAINER"; then
      PG_PASSWORD="$(podman exec "$PG_CONTAINER" printenv POSTGRES_PASSWORD)"
    else
      echo "[init_db] 错误：无法获取密码。" >&2
      exit 1
    fi
  fi
  export PGPASSWORD="$PG_PASSWORD"
fi

# —— 幂等建库 ——
# 连接到共享实例的 postgres 维护库，检查 PSCO 数据库是否已存在
exists=$("${PSQL[@]}" -d postgres \
  -tAc "SELECT 1 FROM pg_database WHERE datname = '$PSCO_DB'" 2>/dev/null || echo "")

if [[ "$exists" == "1" ]]; then
  echo "[init_db] 数据库 '$PSCO_DB' 已存在，跳过创建。"
else
  echo "[init_db] 在共享 PostgreSQL 实例中创建独立数据库 '$PSCO_DB'..."
  "${PSQL[@]}" -d postgres \
    -v ON_ERROR_STOP=1 \
    -c "CREATE DATABASE \"$PSCO_DB\";"
  echo "[init_db] 数据库 '$PSCO_DB' 创建完成。"
fi

# —— 输出下一步提示 ——
echo "[init_db] 下一步："
echo "  1. 启动后端（自动运行迁移 + 可选种子）：在项目根目录执行 ./backend/bin/psco-server"
echo "  2. 或独立执行种子数据：./database/scripts/run_seeds.sh"

unset PGPASSWORD
