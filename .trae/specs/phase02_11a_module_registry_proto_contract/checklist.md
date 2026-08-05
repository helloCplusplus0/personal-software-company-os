- [x] 已明确 `phase02-11A` 的直接上游是 `phase02-09` 的 `module_registry_spec_v0.1.md`
- [x] 已明确 `Module Registry` 当前阶段必须落地最小 `.proto` 合同源
- [x] 已明确 `.proto` 是当前阶段唯一合同源，`chi + JSON HTTP` 只能作为过渡传输层
- [x] 已明确 Proto 包名（`psco.module_registry.v1`）、版本语义（`v1`）与文件落点（`proto/psco/module_registry/v1/`）
- [x] 已明确 `ModuleListRead / ModuleDetailRead / ModuleCreateWrite / ModuleReleaseWrite / ModuleBindingWrite / Candidate Read` 的服务接口需要进入 Proto 合同
- [x] 已明确 `Decision` 只作为 `ModuleDetailRead` 的内嵌附属读取，不新增独立 Proto 服务
- [x] 已明确核心消息结构至少覆盖列表项、详情对象、版本列表、产品绑定、仓库映射与候选读取
- [x] 已明确字段编号必须稳定、不可复用，并允许后续兼容演进
- [x] 已明确当前阶段不要求完整传输层迁移，但必须冻结最小生成入口约定
- [x] 已明确 `backend/internal/moduleregistry/types.go` 与前端 `api-adapter` 后续必须对齐 `.proto` 消息语义
- [x] 已明确 `phase02-12` 联调与验收需要以 `.proto` 合同源为核对基准
- [x] 已明确后续实现不得再按"Protocol Buffers 只作为长期方向、不在当前阶段落地"的旧口径推进

# 额外校验项（实现过程中补充）

- [x] `buf lint` 通过（STANDARD 规则集）
- [x] `buf generate` 成功生成 Go + TypeScript 产物
- [x] `go build ./...` 通过（含生成代码）
- [x] `go vet ./...` 通过
- [x] `npm run build` 通过（生成代码已排除出 TS 编译范围）
- [x] `npm run lint` 通过（0 errors）
- [x] `.proto` 字段语义与 `types.go` / `types.ts` 语义对齐（HTTP DTO 通过显式映射保持语义一致，非字段严格一致）
- [x] `.proto` service 定义覆盖全部 8 个 RPC（读组 2 + 写组 4 + 候选读取 2）
- [x] 每个字段已分配稳定编号，无编号复用
- [x] 生成产物已加入 `.gitignore`，不进入版本控制
- [x] `proto/README.md` 已记录 RPC → HTTP 映射矩阵与字段映射约定
- [x] `proto/Makefile` 提供 `make gen / make lint / make breaking / make clean` 入口
- [x] 当前阶段只生成消息类型，不生成 gRPC 服务桩（符合 spec"不要求完整传输层迁移"）
