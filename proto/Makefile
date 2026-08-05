# proto/Makefile — PSCO Proto 合同生成与校验入口
#
# 上游规格：phase02-11A spec §"生成链与实现边界"
#
# 使用方式：
#   make gen       — 生成 Go + TypeScript 合同产物
#   make lint      — 校验 Proto 文件规范性
#   make breaking  — 检测破坏性变更（对比 main 分支，需要 git 基准）
#   make clean     — 清理生成产物

.PHONY: gen lint breaking clean

# 生成 Go + TypeScript 合同产物
gen:
	buf generate

# 校验 Proto 文件规范性
lint:
	buf lint

# 检测破坏性变更（确保字段编号与语义不被意外删除或修改）
#
# 基准路径指向仓库根 .git（../.git），subdir=proto 指定 git 树中的 proto/ 子目录。
# 失败时直接返回非零退出码，不吞掉错误——这是 .proto 兼容性守卫的前提。
#
# 前置条件：
#   - 仓库已有 main 分支（首次提交前 make breaking 会失败，这是预期行为）
#   - CI 中应在 main 分支基准存在后运行此 target
breaking:
	buf breaking --against '../.git#branch=main,subdir=proto'

# 清理生成产物
clean:
	rm -rf ../backend/internal/gen/proto
	rm -rf ../frontend/src/gen/proto
