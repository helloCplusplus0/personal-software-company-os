#!/usr/bin/env bash
# reset_dashboard_acceptance.sh — phase05-12 验收辅助：Dashboard 验收统一入口
#
# 职责：
#   为 phase05 Dashboard + Feedback 联调验收提供可重复执行的统一入口。
#   编排既有 reset 脚本（module / decision / product_repository）的清空与恢复，
#   并提供 fixture 加载能力。
#
# 模式矩阵（phase05-09 spec §"reset_dashboard_acceptance.sh 必须作为 Dashboard 验收统一入口"）：
#   默认（无参数）           先清空所有 Dashboard 相关数据，再恢复完整基线
#   --clean-only             仅清空所有 Dashboard 相关数据（用于验证空系统状态）
#   --restore-only           仅恢复完整基线（用于验证有数据状态）
#   --fixture <name>         先清空，再加载指定 fixture（用于验证特定 CTA 或区块结果）
#
# 清空范围与顺序（phase05-09 spec §"清空范围"）：
#   按依赖逆序编排既有 reset 脚本的 --clean-only 模式：
#     1. reset_product_repository_mainline.sh --clean-only（依赖 modules）
#     2. reset_decision_mainline.sh --clean-only（decision_links 依赖 modules）
#     3. reset_module_mainline.sh --clean-only（modules 依赖 readonly prereqs）
#   覆盖表：decision_links / product_repositories / product_modules /
#          module_repositories / module_releases / modules / products /
#          repositories / decisions
#   不绕过既有脚本直接 DELETE 或 TRUNCATE 底层表。
#
# 恢复范围与顺序（phase05-09 spec §"恢复范围"）：
#   按依赖顺序执行：
#     1. seed_readonly_prereqs.sql（提供 products / repositories / decisions 只读前提）
#     2. reset_module_mainline.sh --restore-only（恢复 modules / module_releases /
#        product_modules / module_repositories / decision_links）
#     3. reset_decision_mainline.sh --restore-only（恢复 decisions 正式基线与 decision_links）
#     4. reset_product_repository_mainline.sh --restore-only（恢复 products / repositories 正式基线）
#     5. seed_dashboard_acceptance_baseline.sql（补齐 Dashboard 验收所需的额外基线数据）
#
# fixture 加载（phase05-09 spec §"fixture 加载"）：
#   --fixture <name> 先执行清空操作，再加载 seed_dashboard_fixture_<name>.sql。
#   所有 fixture 加载遵守"先清空再加载指定 fixture"的统一语义，禁止叠加已有数据。
#   <name> 只允许取九类 fixture 名称之一：
#     empty-system / modules-only / products-without-modules /
#     products-missing-repository / products-missing-module /
#     pending-decisions / recent-activities /
#     products-missing-all-repositories / products-missing-both-bindings
#
# 幂等：可重复执行。清空依赖既有脚本的受控 DELETE；恢复依赖 ON CONFLICT DO NOTHING + UPDATE + WHERE NOT EXISTS。
#
# 上游规格：phase05-09 spec §"reset_dashboard_acceptance.sh 必须作为 Dashboard 验收统一入口"
#           phase05-12 spec §"reset_dashboard_acceptance.sh 必须作为 Dashboard 验收统一入口"
# 依赖：既有 reset_module_mainline.sh / reset_decision_mainline.sh /
#       reset_product_repository_mainline.sh 与 seed_readonly_prereqs.sql
#
# 用法：
#   ./database/scripts/reset_dashboard_acceptance.sh                            # 清空 + 恢复基线
#   ./database/scripts/reset_dashboard_acceptance.sh --clean-only               # 仅清空
#   ./database/scripts/reset_dashboard_acceptance.sh --restore-only             # 仅恢复基线
#   ./database/scripts/reset_dashboard_acceptance.sh --fixture pending-decisions # 清空 + 加载指定 fixture
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
        echo "[reset_dashboard_acceptance] 错误：--fixture 需要指定 fixture 名称" >&2
        echo "[reset_dashboard_acceptance] 用法：$0 --fixture <name>" >&2
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
      echo "[reset_dashboard_acceptance] 错误：未知参数 '$1'" >&2
      echo "[reset_dashboard_acceptance] 用法：$0 [--clean-only|--restore-only|--fixture <name>]" >&2
      exit 1
      ;;
  esac
fi

# —— fixture 名称校验 ——
# 九类 fixture 名称清单（phase05-09 spec §"fixture 命名与 CTA 映射矩阵"）
VALID_FIXTURES=(
  empty-system
  modules-only
  products-without-modules
  products-missing-repository
  products-missing-module
  pending-decisions
  recent-activities
  products-missing-all-repositories
  products-missing-both-bindings
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
    echo "[reset_dashboard_acceptance] 错误：未知 fixture 名称 '$FIXTURE_NAME'" >&2
    echo "[reset_dashboard_acceptance] 允许的 fixture 名称：" >&2
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
MODULE_RESET="$SCRIPT_DIR/reset_module_mainline.sh"
DECISION_RESET="$SCRIPT_DIR/reset_decision_mainline.sh"
PRODUCT_RESET="$SCRIPT_DIR/reset_product_repository_mainline.sh"

# —— 校验既有 reset 脚本存在 ——
for script in "$MODULE_RESET" "$DECISION_RESET" "$PRODUCT_RESET"; do
  if [[ ! -x "$script" ]]; then
    echo "[reset_dashboard_acceptance] 错误：依赖脚本不存在或不可执行：$script" >&2
    exit 1
  fi
done

# —— 解析 psql 执行方式（数组形式，用于直接执行 SQL 文件）——
# 优先级：宿主机 psql > docker exec > podman exec
# 既有 reset 脚本各自独立解析 psql，本函数仅用于直接执行 seed SQL 文件。
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
  echo "[reset_dashboard_acceptance] 错误：未找到 psql，且容器 '$PG_CONTAINER' 未运行。" >&2
  echo "[reset_dashboard_acceptance] 请安装 postgresql-client，或启动共享 PostgreSQL 容器。" >&2
  exit 1
}

# 立即解析 psql 一次，PSQL 数组与 PSQL_MODE 在所有后续函数中可用。
# 既有 reset 脚本各自独立解析 psql；本脚本直接执行 seed SQL 文件时也需 PSQL 数组，
# 因此提前在主流程顶部解析，避免 do_restore / do_fixture 内重复解析。
resolve_psql

# —— 解析密码（仅宿主机模式需要）——
# 与既有 reset_module_mainline.sh / reset_decision_mainline.sh /
# reset_product_repository_mainline.sh 保持一致：
# 当检测到本机 psql 时，必须补 PGPASSWORD 或从容器读取 POSTGRES_PASSWORD，
# 否则 check_db_exists、seed_readonly_prereqs.sql、seed_dashboard_acceptance_baseline.sql、
# fixture SQL 等本脚本直接执行的步骤会在"本机装了 psql 但未配 PG_PASSWORD/.pgpass"场景下失败。
if [[ "$PSQL_MODE" == "host" ]]; then
  if [[ -z "${PG_PASSWORD:-}" ]]; then
    echo "[reset_dashboard_acceptance] PG_PASSWORD 未设置，尝试从容器读取..." >&2
    if docker ps --format '{{.Names}}' 2>/dev/null | grep -qx "$PG_CONTAINER"; then
      PG_PASSWORD="$(docker exec "$PG_CONTAINER" printenv POSTGRES_PASSWORD)"
    elif podman ps --format '{{.Names}}' 2>/dev/null | grep -qx "$PG_CONTAINER"; then
      PG_PASSWORD="$(podman exec "$PG_CONTAINER" printenv POSTGRES_PASSWORD)"
    else
      echo "[reset_dashboard_acceptance] 错误：无法获取密码，容器未运行且 PG_PASSWORD 未设置。" >&2
      exit 1
    fi
  fi
  export PGPASSWORD="${PG_PASSWORD:-}"
fi

# —— 前置校验：PSCO 数据库必须已存在 ——
# 既有 reset 脚本各自也有此校验，但提前校验可给出更清晰的错误。
# resolve_psql 已在主流程顶部完成，此处直接复用 PSQL 数组。
check_db_exists() {
  exists=$("${PSQL[@]}" -d postgres \
    -tAc "SELECT 1 FROM pg_database WHERE datname = '$PSCO_DB'" 2>/dev/null || echo "")
  if [[ "$exists" != "1" ]]; then
    echo "[reset_dashboard_acceptance] 错误：数据库 '$PSCO_DB' 不存在，请先执行 ./database/scripts/init_db.sh" >&2
    exit 1
  fi
}

# —— 清空函数：按依赖逆序编排既有 reset 脚本 ——
do_clean() {
  echo "[reset_dashboard_acceptance] 清空 Dashboard 相关数据（按依赖逆序）..."
  # 1. product_repository（依赖 modules，先清关系表与实体表）
  "$PRODUCT_RESET" --clean-only
  # 2. decision（decision_links 依赖 modules）
  "$DECISION_RESET" --clean-only
  # 3. module（modules 依赖 readonly prereqs 中的 products / decisions）
  "$MODULE_RESET" --clean-only
  echo "[reset_dashboard_acceptance] 清空完成。当前所有 Dashboard 相关表为空。"
}

# —— 恢复函数：按依赖顺序编排既有 reset 脚本 + dashboard baseline ——
do_restore() {
  local readonly_prereqs="$SEEDS_DIR/seed_readonly_prereqs.sql"
  local dashboard_baseline="$SEEDS_DIR/seed_dashboard_acceptance_baseline.sql"

  if [[ ! -f "$readonly_prereqs" ]]; then
    echo "[reset_dashboard_acceptance] 错误：只读前提 SQL 文件不存在：$readonly_prereqs" >&2
    exit 1
  fi
  if [[ ! -f "$dashboard_baseline" ]]; then
    echo "[reset_dashboard_acceptance] 错误：Dashboard 基线 SQL 文件不存在：$dashboard_baseline" >&2
    exit 1
  fi

  echo "[reset_dashboard_acceptance] 恢复 Dashboard 验收基线（按依赖顺序）..."
  # 1. 只读前提（products / repositories / decisions 占位）
  echo "[reset_dashboard_acceptance]   1/5 seed_readonly_prereqs.sql"
  "${PSQL[@]}" -d "$PSCO_DB" -v ON_ERROR_STOP=1 < "$readonly_prereqs"
  # 2. module 主线恢复
  echo "[reset_dashboard_acceptance]   2/5 reset_module_mainline.sh --restore-only"
  "$MODULE_RESET" --restore-only
  # 3. decision 主线恢复（依赖 modules）
  echo "[reset_dashboard_acceptance]   3/5 reset_decision_mainline.sh --restore-only"
  "$DECISION_RESET" --restore-only
  # 4. product_repository 主线恢复（依赖 modules）
  echo "[reset_dashboard_acceptance]   4/5 reset_product_repository_mainline.sh --restore-only"
  "$PRODUCT_RESET" --restore-only
  # 5. Dashboard 验收基线补齐
  echo "[reset_dashboard_acceptance]   5/5 seed_dashboard_acceptance_baseline.sql"
  "${PSQL[@]}" -d "$PSCO_DB" -v ON_ERROR_STOP=1 < "$dashboard_baseline"
  echo "[reset_dashboard_acceptance] 恢复完成。"
}

# —— fixture 加载函数：先清空，再加载指定 fixture ——
do_fixture() {
  # fixture 名称使用连字符（如 pending-decisions），文件名使用下划线（如 pending_decisions），
  # 对齐 phase05-09 spec Impact 部分冻结的文件落点。
  local fixture_file_name="${FIXTURE_NAME//-/_}"
  local fixture_sql="$SEEDS_DIR/seed_dashboard_fixture_${fixture_file_name}.sql"
  if [[ ! -f "$fixture_sql" ]]; then
    echo "[reset_dashboard_acceptance] 错误：fixture SQL 文件不存在：$fixture_sql" >&2
    exit 1
  fi

  echo "[reset_dashboard_acceptance] 加载 fixture '$FIXTURE_NAME'（先清空再加载）..."
  do_clean
  echo "[reset_dashboard_acceptance] 加载 seed_dashboard_fixture_${fixture_file_name}.sql"
  "${PSQL[@]}" -d "$PSCO_DB" -v ON_ERROR_STOP=1 < "$fixture_sql"
  echo "[reset_dashboard_acceptance] fixture '$FIXTURE_NAME' 加载完成。"
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
