# Tasks

- [x] Task 1: 冻结 brief 合同落点与 proto schema
  - [x] SubTask 1.1: 在 `project_context.proto` 新增 `GetProjectBrief` RPC 与 brief 消息（7 顶层字段，import 复用 governance_profile 消息）
    - 执行记录：新增 `BriefCurrentPhase`（name / entry_ref / status，status 复用 `psco.governance_profile.v1.PhaseStatus`）、`GetProjectBriefRequest` / `GetProjectBriefResponse`（7 顶层字段：repository / governance_profile / global_assets / current_phase / products[] / modules[] / decisions[]）；`import "psco/governance_profile/v1/governance_profile.proto"` 复用既有消息，无第 8 顶层块、无自然语言指导词字段。
  - [x] SubTask 1.2: 为 `GetProjectContext / ExportProjectContext` 标注 deprecated 兼容窗口注释，行为保持不变
    - 执行记录：两个旧 RPC 增加 `option deprecated = true;` 与 Deprecated 注释；服务级注释冻结兼容窗口（phase13 内行为与响应结构不变、退役讨论后置到 phase13-11 验收后）；未改动旧 RPC 行为（含硬编码投影）。
  - [x] SubTask 1.3: 执行 `make gen` 完成 Go / TS 代码生成并确认 buf lint / build 通过
    - 执行记录：`make lint && make build && make gen` 全部通过；Go 产物（`backend/internal/gen/proto|connect/psco/project_context/v1/`）与 TS 产物（`frontend/src/gen/proto/psco/project_context/v1/project_context_pb.ts`）均含 `GetProjectBrief` / `BriefCurrentPhase`。

- [x] Task 2: 冻结跨模块读取承接位
  - [x] SubTask 2.1: `projectcontext/candidate` 新增 `GovernanceProfileReader` 接口
    - 执行记录：`context_readers.go` 新增消费方拥有的 `GovernanceProfileReader` 接口（`GetGovernanceProfile(ctx, repositoryID)`），注释冻结装配与禁止复制 SQL 约束。
  - [x] SubTask 2.2: `projectcontext/candidate` 新增 `ReadProducts`（数组版）；保留旧 `ReadProduct` 仅供兼容层
    - 执行记录：`ReadProducts` 数组版无 LIMIT 1、`ORDER BY p.name`、空结果返回空切片；旧 `ReadProduct`（singular LIMIT 1）保留并标注 Deprecated 仅供 `GetProjectContext` 兼容层。
  - [x] SubTask 2.3: `governanceprofile/connect` 导出 domain→proto 转换函数（`DomainResultToProto` 等）供复用
    - 执行记录：导出 `DomainResultToProto` / `DomainAssetBindingsToProto`（新提取）/ `DomainPhaseStatusToProto`，均注释 phase13-10 复用原因；`server_test.go` 同步引用更新。
  - [x] SubTask 2.4: `router.go` 装配点完成 governanceprofile 读取依赖注入
    - 执行记录：`buildProjectContext(pool, governanceReader)` 签名注入；`server.go` 装配顺序调整为先构建 `governanceProfileQuerySvc` 再传入 projectcontext（`projectcontextservice.NewQueryService(readers, governanceReader)`）。

- [x] Task 3: 实现 brief service 编排与 Connect handler
  - [x] SubTask 3.1: `types.go` 新增 brief domain DTO（BriefCurrentPhase 等）
    - 执行记录：新增 `BriefCurrentPhase`（name / entry_ref / status 字符串形式）与 `ProjectBriefReadResult`（7 顶层块，治理画像 domain 类型直接透传）。
  - [x] SubTask 3.2: `service/query_service.go` 新增 `GetProjectBrief` 编排（repository → 治理画像 → products → modules → decisions → current_phase 派生）
    - 执行记录：编排顺序与 spec 冻结步骤一一对应；治理画像错误原样透传（保留 `ErrGovernanceProfileNotFound` 哨兵语义）；不调用 `ValidateRepositoryBinding`（数组空合法）；`phaseStatusToString` 供 domain DTO 字符串形态。
  - [x] SubTask 3.3: `connect/server.go` 新增 `GetProjectBrief` handler（复用 governanceprofile 转换函数组装 7 顶层块）
    - 执行记录：handler 复用 `gpconnect.DomainResultToProto / DomainAssetBindingsToProto / DomainPhaseStatusToProto` 组装，无第二套治理画像字段映射；新增 `domainProductSummariesToProto` 数组转换（四实体自身 DTO）。
  - [x] SubTask 3.4: 失败语义验证：repository 不存在 → NotFound；画像未创建 → NotFound；数组空为合法
    - 执行记录：Connect 冒烟验证三类场景全部符合（详见 Task 4.2）；`connecterrors.MapToConnectError` 既有映射表已含 `projectcontext.ErrRepositoryNotFound` / `governanceprofile.ErrGovernanceProfileNotFound`，无需扩展。

- [x] Task 4: 验证与收口
  - [x] SubTask 4.1: `go build ./...` 与 `go vet ./...` 通过；既有测试不回归
    - 执行记录：build / vet 零错误；`go test ./...` 全部通过（含 `projectcontext/connect` 集成测试 9.5s，旧 RPC 4 场景反回归 PASS）；frontend `npx tsc --noEmit` 通过。
  - [x] SubTask 4.2: Connect 冒烟：brief 返回 7 顶层块完整、与治理画像同源、数组语义正确；旧 RPC 行为不变
    - 执行记录（真实数据库 psco_development + main-repo）：
      1. GetProjectBrief repository 不存在 → 404 `not_found` ✓
      2. GetProjectBrief 画像未创建 → 404 `not_found`（`governance profile not created`）✓
      3. UpdateGovernanceProfile 创建画像后 GetProjectBrief → 7 顶层块完整：repository ✓ / governanceProfile（与 GetGovernanceProfile 同源，含 9 类核心字段）✓ / globalAssets（与画像 globalAssetBindings 同源，markdown_resolvable=true）✓ / currentPhase（从主记录三 read-only 字段派生，PHASE_STATUS_IN_PROGRESS）✓ / products 长度 1 保持数组 ✓ / modules 数组 ✓ / decisions 数组（hitSources 双命中合并）✓
      4. 旧 GetProjectContext → 200，singular product + 7 条硬编码 rules + boundaries 保留，行为不变 ✓
      5. 旧 ExportProjectContext → 200 行为不变 ✓
  - [x] SubTask 4.3: 子代理独立复核并留档结论
    - 执行记录：独立复核 12 项清单全部 PASS，最终结论「通过」，无阻断性问题；3 条非阻断建议留档：
      1. GetProjectBrief 缺自动化集成测试用例（现有 fixture 不含治理画像数据），建议 phase13-11 验收阶段补充
      2. 本文件与 checklist.md 收口勾选（已在本轮完成）
      3. `BriefCurrentPhase.Status`（string DTO 字段）在 Connect 组装路径未被消费（handler 直接从枚举转换更安全），已按建议精确化 `types.go` 注释；`phaseStatusToString` 保留供 JSON domain DTO 消费

# Task Dependencies

- `Task 2` depends on `Task 1`
- `Task 3` depends on `Task 1`, `Task 2`
- `Task 4` depends on `Task 3`
