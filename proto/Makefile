# proto/Makefile — PSCO Proto 合同生成与校验入口
#
# 上游规格：phase02-11A spec §"生成链与实现边界"
#           phase03-08 spec §"生成链与校验链必须冻结"（补全 build 受控入口）
#           phase03-11 spec §"buf 工具链入口必须可运行且复用现有入口"
#           phase04-08 spec §"合同演进与 breaking check 规则必须冻结"
#           phase04-11 spec §"buf 工具链入口必须可运行且复用现有入口"
#           phase05-08 spec §"合同演进与 Buf 校验规则必须冻结"
#           phase05-11 spec §"buf 工具链入口必须可运行且复用现有入口"
#
# 使用方式：
#   make build     — 编译 proto 文件为 buf image，验证合同源可被正确解析与编译
#   make gen       — 生成 Go + TypeScript 合同产物
#   make lint      — 校验 Proto 文件规范性
#   make breaking  — 检测破坏性变更（对比 main 分支，需要 git 基准）
#   make clean     — 清理生成产物

.PHONY: build gen lint breaking clean

# 编译 proto 文件为 buf image，验证合同源可被正确解析与编译。
# 这是 buf 校验链的最基础入口，lint / generate / breaking 都隐式依赖 build 成功。
build:
	buf build

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
#   - 仓库已有 main 分支且 proto/ 子目录已有 .proto 基准（phase02-11A 已落地）
#   - 新增 .proto 文件属于向后兼容变更，不触发 breaking 失败
#   - CI 中应在 main 分支基准存在后运行此 target
breaking:
	buf breaking --against '../.git#branch=main,subdir=proto'

# 清理生成产物
clean:
	rm -rf ../backend/internal/gen/proto
	rm -rf ../frontend/src/gen/proto
