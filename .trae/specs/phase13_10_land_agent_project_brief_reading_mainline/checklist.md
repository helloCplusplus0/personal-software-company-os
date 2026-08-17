- [x] brief 合同落点已冻结为 `ProjectContextService` 内新增 `GetProjectBrief` RPC，未新建第二个只读聚合 service
- [x] brief proto schema 已冻结为 7 顶层字段，复用 `governance_profile.proto` 与既有 summary 消息
- [x] brief 读取编排已冻结：repository 存在校验 → 治理画像聚合 → 三数组摘要 → current_phase 派生
- [x] brief 失败语义已冻结：repository 不存在 / 画像未创建 → NotFound；数组空为合法，不做绑定完整性强制校验
- [x] `GetProjectContext / ExportProjectContext` 已冻结为 deprecated 兼容层，phase13 内行为不变，退役窗口后置
- [x] 跨模块承接已冻结：candidate 接口 + 平台装配注入，不在 projectcontext 内复制治理画像 SQL 或字段映射
- [x] brief 不混入硬编码投影、目录扫描、自然语言指导词或第 8 个顶层块

---

验收结论（2026-08-17）：

- buf lint / build / gen 通过；`go build ./...` / `go vet ./...` / `go test ./...`（含旧 RPC 集成反回归）通过；frontend `tsc --noEmit` 通过
- Connect 冒烟五场景全部符合预期（两类 NotFound 失败语义、7 顶层块完整同源、数组语义、旧 RPC 行为不变）
- 子代理独立复核 12 项清单全部 PASS，最终结论「通过」；3 条非阻断建议已留档至 tasks.md Task 4.3，其中注释精确化建议已当场修复
