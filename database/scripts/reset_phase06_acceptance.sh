#!/usr/bin/env bash
# reset_phase06_acceptance.sh — phase06-14 验收辅助：Phase06 联调验收统一入口
#
# 职责：
#   为 phase06 Onboarding + Sovereignty + Reuse 联调验收提供可重复执行的统一入口。
#   编排既有 reset 脚本（module / decision / product_repository / dashboard）的清空与恢复，
#   并提供 phase06 专属数据（instance_exports / instance_backups）的清理与 fixture 加载能力。
#
# 模式矩阵（phase06-11 spec §"reset_phase06_acceptance.sh 必须作为统一验收入口"）：
#   默认（无参数）           先清空所有 phase06 相关数据，再恢复完整基线
#   --clean-only             仅清空所有 phase06 相关数据（用于验证空系统状态）
#   --restore-only           仅恢复完整基线（用于验证有数据状态）
#   --fixture <name>         先清空，再加载指定 fixture（用于验证特定状态）
#
# 清空范围与顺序（phase06-14 spec §"Phase06 清空与恢复边界"）：
#   1. reset_dashboard_acceptance.sh --clean-only（编排 module / decision / product_repository 清空）
#   2. 清理 instance_exports 与 instance_backups（phase06 专属元数据表）
#   不清理 schema_migrations 或 migration 基线。
#   不绕过既有脚本直接 DELETE 或 TRUNCATE canonical 表。
#
# 恢复范围与顺序：
#   1. reset_dashboard_acceptance.sh --restore-only（恢复 phase05 Dashboard / Feedback 验收基线）
#   2. seed_phase06_acceptance_baseline.sql（补齐 phase06 验收所需的额外基线数据）
#
# fixture 加载：
#   --fixture <name> 先执行清空操作，再加载 seed_phase06_fixture_<name>.sql。
#   所有 fixture 加载遵守"先清空再加载指定 fixture"的统一语义，禁止叠加已有数据。
#   <name> 只允许取 11 类 fixture 名称之一（phase06-11 spec §"Fixture 白名单"）。
#
# 幂等：可重复执行。清空依赖既有脚本的受控 DELETE + phase06 专属表 TRUNCATE；
#   恢复依赖 ON CONFLICT DO NOTHING + UPDATE + WHERE NOT EXISTS。
#
# 上游规格：phase06-11 spec §"reset_phase06_acceptance.sh 必须作为统一验收入口"
#           phase06-14 spec §"reset_phase06_acceptance.sh 必须把 11 个 fixture 变成可重复执行的正式主线"
# 依赖：既有 reset_dashboard_acceptance.sh / reset_module_mainline.sh /
#       reset_decision_mainline.sh / reset_product_repository_mainline.sh
#
# 用法：
#   ./database/scripts/reset_phase06_acceptance.sh                            # 清空 + 恢复基线
#   ./database/scripts/reset_phase06_acceptance.sh --clean-only               # 仅清空
#   ./database/scripts/reset_phase06_acceptance.sh --restore-only             # 仅恢复基线
#   ./database/scripts/reset_phase06_acceptance.sh --fixture cold-start-empty # 清空 + 加载指定 fixture
#
# 可覆盖参数（带默认值，与既有 reset 脚本一致）：
#   PG_HOST (默认 127.0.0.1)
#   PG_PORT (默认 55432)
#   PG_USER (默认 rento)
#   PSCO_DB (默认 psco_development)
#   PG_CONTAINER (默认 rento-preview-postgres)
#   SEEDS_DIR (默认脚本同级的 ../seeds)

set -euo pipefail

# —— 模式解析 ——
MODE="all"
FIXTURE_NAME=""

if [[ $# -ge 1 ]]; then
  case "$1" in
    --clean-only)   MODE="clean" ;;
    --restore-only) MODE="restore" ;;
    --fixture)
      MODE="fixture"
      if [[ $# -lt 2 ]]; then
        echo "[reset_phase06_acceptance] 错误：--fixture 需要指定 fixture 名称" >&2
        echo "[reset_phase06_acceptance] 用法：$0 --fixture <name>" >&2
        exit 1
      fi
      FIXTURE_NAME="$2"
      ;;
    --all|--reset)  MODE="all" ;;
    -h|--help)
      sed -n '2,/^$/p' "$0"
      exit 0
      ;;
    *)
      echo "[reset_phase06_acceptance] 错误：未知参数 '$1'" >&2
      echo "[reset_phase06_acceptance] 用法：$0 [--clean-only|--restore-only|--fixture <name>]" >&2
      exit 1
      ;;
  esac
fi

# —— fixture 名称校验 ——
# 11 类 fixture 名称清单（phase06-11 spec §"Fixture 白名单"）
VALID_FIXTURES=(
  cold-start-empty
  in-progress-partial-entry
  completed-unbound
  completed-bound
  export-ready
  backup-verified
  backup-manifest-missing
  backup-coverage-incomplete
  backup-schema-mismatch
  reuse-latest
  reuse-latest-after-binding
)

if [[ "$MODE" == "fixture" ]]; then
  fixture_valid=false
  for f in "${VALID_FIXTURES[@]}"; do
    if [[ "$FIXTURE_NAME" == "$f" ]]; then
      fixture_valid=true
      break
    fi
  done
  if [[ "$fixture_valid" != "true" ]]; then
    echo "[reset_phase06_acceptance] 错误：未知 fixture 名称 '$FIXTURE_NAME'" >&2
    echo "[reset_phase06_acceptance] 允许的 fixture 名称：" >&2
    printf '  %s\n' "${VALID_FIXTURES[@]}" >&2
    exit 1
  fi
fi

# —— 连接参数（允许环境变量覆盖，与既有 reset 脚本一致）——
PG_HOST="${PG_HOST:-127.0.0.1}"
PG_PORT="${PG_PORT:-55432}"
PG_USER="${PG_USER:-rento}"
PSCO_DB="${PSCO_DB:-psco_development}"
PG_CONTAINER="${PG_CONTAINER:-rento-preview-postgres}"

# 脚本与种子文件目录
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
SEEDS_DIR="${SEEDS_DIR:-$SCRIPT_DIR/../seeds}"

# 既有 reset 脚本路径
DASHBOARD_RESET="$SCRIPT_DIR/reset_dashboard_acceptance.sh"

# —— 校验既有 reset 脚本存在 ——
if [[ ! -x "$DASHBOARD_RESET" ]]; then
  echo "[reset_phase06_acceptance] 错误：依赖脚本不存在或不可执行：$DASHBOARD_RESET" >&2
  exit 1
fi

# —— 解析 psql 执行方式（数组形式，用于直接执行 SQL 文件）——
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
  echo "[reset_phase06_acceptance] 错误：未找到 psql，且容器 '$PG_CONTAINER' 未运行。" >&2
  echo "[reset_phase06_acceptance] 请安装 postgresql-client，或启动共享 PostgreSQL 容器。" >&2
  exit 1
}

resolve_psql

# —— 解析密码（仅宿主机模式需要）——
if [[ "$PSQL_MODE" == "host" ]]; then
  if [[ -z "${PG_PASSWORD:-}" ]]; then
    echo "[reset_phase06_acceptance] PG_PASSWORD 未设置，尝试从容器读取..." >&2
    if docker ps --format '{{.Names}}' 2>/dev/null | grep -qx "$PG_CONTAINER"; then
      PG_PASSWORD="$(docker exec "$PG_CONTAINER" printenv POSTGRES_PASSWORD)"
    elif podman ps --format '{{.Names}}' 2>/dev/null | grep -qx "$PG_CONTAINER"; then
      PG_PASSWORD="$(podman exec "$PG_CONTAINER" printenv POSTGRES_PASSWORD)"
    else
      echo "[reset_phase06_acceptance] 错误：无法获取密码，容器未运行且 PG_PASSWORD 未设置。" >&2
      exit 1
    fi
  fi
  export PGPASSWORD="${PG_PASSWORD:-}"
fi

# —— 前置校验：PSCO 数据库必须已存在 ——
check_db_exists() {
  exists=$("${PSQL[@]}" -d postgres \
    -tAc "SELECT 1 FROM pg_database WHERE datname = '$PSCO_DB'" 2>/dev/null || echo "")
  if [[ "$exists" != "1" ]]; then
    echo "[reset_phase06_acceptance] 错误：数据库 '$PSCO_DB' 不存在，请先执行 ./database/scripts/init_db.sh" >&2
    exit 1
  fi
}

# —— 清空函数：编排既有 dashboard reset 脚本 + 清理 phase06 专属表 ——
do_clean() {
  echo "[reset_phase06_acceptance] 清空 phase06 相关数据（编排既有 reset + 专属表清理）..."
  # 1. dashboard reset --clean-only（编排 module / decision / product_repository 清空）
  "$DASHBOARD_RESET" --clean-only
  # 2. 清理 phase06 专属元数据表（instance_exports / instance_backups）
  #    不清理 schema_migrations
  "${PSQL[@]}" -d "$PSCO_DB" -v ON_ERROR_STOP=1 \
    -c "DELETE FROM instance_backups;" \
    -c "DELETE FROM instance_exports;"
  echo "[reset_phase06_acceptance] 清空完成。"
}

# —— 恢复函数：编排既有 dashboard reset 恢复 + phase06 baseline ——
do_restore() {
  local phase06_baseline="$SEEDS_DIR/seed_phase06_acceptance_baseline.sql"

  if [[ ! -f "$phase06_baseline" ]]; then
    echo "[reset_phase06_acceptance] 错误：phase06 基线 SQL 文件不存在：$phase06_baseline" >&2
    exit 1
  fi

  echo "[reset_phase06_acceptance] 恢复 phase06 验收基线（按依赖顺序）..."
  # 1. 恢复 phase05 Dashboard / Feedback 验收基线
  echo "[reset_phase06_acceptance]   1/2 reset_dashboard_acceptance.sh --restore-only"
  "$DASHBOARD_RESET" --restore-only
  # 2. 补齐 phase06 验收基线
  echo "[reset_phase06_acceptance]   2/2 seed_phase06_acceptance_baseline.sql"
  "${PSQL[@]}" -d "$PSCO_DB" -v ON_ERROR_STOP=1 < "$phase06_baseline"
  echo "[reset_phase06_acceptance] 恢复完成。"
}

# —— fixture 加载函数：先清空，再加载指定 fixture ——
do_fixture() {
  # fixture 名称使用连字符（如 cold-start-empty），文件名使用下划线（如 cold_start_empty）
  local fixture_file_name="${FIXTURE_NAME//-/_}"
  local fixture_sql="$SEEDS_DIR/seed_phase06_fixture_${fixture_file_name}.sql"
  if [[ ! -f "$fixture_sql" ]]; then
    echo "[reset_phase06_acceptance] 错误：fixture SQL 文件不存在：$fixture_sql" >&2
    exit 1
  fi

  echo "[reset_phase06_acceptance] 加载 fixture '$FIXTURE_NAME'（先清空再加载）..."
  do_clean
  echo "[reset_phase06_acceptance] 加载 seed_phase06_fixture_${fixture_file_name}.sql"
  "${PSQL[@]}" -d "$PSCO_DB" -v ON_ERROR_STOP=1 < "$fixture_sql"
  echo "[reset_phase06_acceptance] fixture '$FIXTURE_NAME' 加载完成。"
}

# —— 主流程 ——
check_db_exists

case "$MODE" in
  clean)
    do_clean
    ;;
  restore)
    do_restore
    ;;
  all)
    do_clean
    do_restore
    ;;
  fixture)
    do_fixture
    ;;
esac
