# Tasks

- [x] Task 1: 实现 `template_reuse.proto` 与生成链接入
  - [x] SubTask 1.1: 新增 `proto/psco/template_reuse/v1/template_reuse.proto`
  - [x] SubTask 1.2: 落实 `TemplateReuseService` 的四个正式读取 RPC 与 `phase09-07` 已冻结的 request/response 合同
  - [x] SubTask 1.3: 接入现有 `buf build / gen / lint / breaking` 主线，生成 Go proto、Go Connect 与前端 TypeScript 产物

- [x] Task 2: 实现 template reuse 后端 query owner 与 Connect 承接位
  - [x] SubTask 2.1: 新增 `backend/internal/templatereuse/service/query_service.go`
  - [x] SubTask 2.2: 新增 `backend/internal/templatereuse/candidate/template_candidate_readers.go`
  - [x] SubTask 2.3: 新增 `backend/internal/templatereuse/connect/server.go`
  - [x] SubTask 2.4: 在 `backend/internal/platform/router.go` 挂载 `TemplateReuseService`
  - [x] SubTask 2.5: 落实 `RESOLVED / UNAVAILABLE` 成功态与 Connect 错误映射边界

- [x] Task 3: 实现 template reuse 前端只读切片与四个 read owner
  - [x] SubTask 3.1: 新增 `frontend/src/features/template-reuse/data/connect-client.ts`
  - [x] SubTask 3.2: 新增 `frontend/src/features/template-reuse/data/template-reuse-query-options.ts`
  - [x] SubTask 3.3: 新增 `use-template-candidates-read.ts`
  - [x] SubTask 3.4: 新增 `use-template-prefill-read.ts`
  - [x] SubTask 3.5: 新增 `use-derived-insight-hints-read.ts`
  - [x] SubTask 3.6: 新增 `use-template-source-read.ts`
  - [x] SubTask 3.7: 确保四个 read owner 只承接 queryKey、读取、解包与 `empty / unavailable / error` 派生

- [x] Task 4: 清理 enablement 边界并预留后续消费位
  - [x] SubTask 4.1: 确保 `Weekly Review / Product Create / Product Detail` 未直接接 generated client 或底层 template reuse query
  - [x] SubTask 4.2: 确保模板读取没有落进 `review/data/` 或 `product-registry/data/`
  - [x] SubTask 4.3: 仅在必要位置预留后续消费口，不提前实现 `phase09-09 / 10` 的 handoff、页面展示与 create 回流

- [x] Task 5: 完成 phase09-08 的工具链验证与最小 smoke
  - [x] SubTask 5.1: 执行 `(cd proto && make build && make gen && make lint)`
  - [x] SubTask 5.2: 在有基准分支条件下执行 `(cd proto && make breaking)`
  - [x] SubTask 5.3: 执行 `(cd backend && go build ./...)`
  - [x] SubTask 5.4: 执行 `(cd frontend && npm run build)`
  - [x] SubTask 5.5: 验证四个 `TemplateReuseService` RPC 的最小 API smoke
  - [x] SubTask 5.6: 在 smoke 中覆盖至少一种 `UNAVAILABLE` 成功态

- [x] Task 6: 完成 phase09-08 实现收口与一致性复核
  - [x] SubTask 6.1: 复核支撑能力相关合同与 owner 已单值一致
  - [x] SubTask 6.2: 复核前后端已具备进入模板候选与提示消费的正式承接位
  - [x] SubTask 6.3: 复核本阶段未以页面级临时拼装作为长期稳态
  - [x] SubTask 6.4: 复核 `phase09-09 / 10` 的页面级逻辑没有被提前偷跑

# Task Dependencies

- Task 2 depends on Task 1
- Task 3 depends on Task 1
- Task 4 depends on Task 2, Task 3
- Task 5 depends on Task 1, Task 2, Task 3, Task 4
- Task 6 depends on Task 1, Task 2, Task 3, Task 4, Task 5
