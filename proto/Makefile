# proto/Makefile — PSCO Proto 合同生成与校验入口
#
# 上游规格：phase07-07 formal spec §3.3（buf 3 插件正式矩阵）
#           phase07-02 frozen_scope.md §3（工具链入口）
#           历史上游：phase02-11A / phase03-08/11 / phase04-08/11 / phase05-08/11
#
# 使用方式：
#   make build     — 编译 proto 文件为 buf image，验证合同源可被正确解析与编译
#   make gen       — 生成 Go protobuf + Go Connect + TypeScript 合同产物
#   make lint      — 校验 Proto 文件规范性
#   make breaking  — 检测破坏性变更（对比 main 分支，需要 git 基准）
#   make clean     — 清理生成产物

.PHONY: build gen lint breaking clean

# 编译 proto 文件为 buf image，验证合同源可被正确解析与编译。
# 这是 buf 校验链的最基础入口，lint / generate / breaking 都隐式依赖 build 成功。
build:
	buf build

# 生成 Go protobuf + Go Connect + TypeScript 合同产物
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

# 清理生成产物（覆盖 Go protobuf + Go Connect + TypeScript 三个生成目录）
clean:
	rm -rf ../backend/internal/gen/proto
	rm -rf ../backend/internal/gen/connect
	rm -rf ../frontend/src/gen/proto
