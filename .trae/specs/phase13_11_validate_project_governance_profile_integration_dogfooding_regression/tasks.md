# Tasks

- [x] Task 1: 验收环境与固定样本恢复
  - [x] SubTask 1.1: 执行 `database/scripts/restore_phase11_phase12_dogfooding_sample.sh` 恢复固定样本（`personal-software-company-os` / `ca261521-8daf-4248-8f12-43525326e759` / Product `PSCO`）
    - 执行记录：固定样本曾因 `projectcontext` 集成测试切库丢失，本轮已先执行正式恢复脚本补齐；本轮 API 复验确认四实体样本在库。
  - [x] SubTask 1.2: 恢复后验证固定 `repository_id` 一次解析成功（`GetProjectContext` 兼容层或 `GetProjectBrief` repository 块），并验证画像未创建时 `GetProjectBrief` 返回 `not_found`（预期行为留档）
    - 执行记录：`GetProjectContext`（兼容层）一次解析成功，repository 块返回 `{"id":"ca261521-8daf-4248-8f12-43525326e759","name":"personal-software-company-os","provider":"github",...}`，product 块为 `PSCO`，modules/decisions/rules/phases 各块完整；画像未创建状态下 `GetProjectBrief` 返回 HTTP 404 `{"code":"not_found","message":"governanceprofile: get governance profile: governance profile not created for repository"}`，与 spec 冻结的预期行为一致。

- [x] Task 2: 人类维护路径 dogfooding（Web 表单创建画像）
  - [x] SubTask 2.1: 通过 `/repositories/ca261521-8daf-4248-8f12-43525326e759` 治理画像表单创建画像：维护 `template_source`、`canonical_root_files[]`（对齐项目范式 v1 根级文件集合）、`global_asset_bindings[]`（8 项资产全覆盖，前 5 项填写 `structured_summary`）
    - 执行记录：浏览器代理经真实表单操作完成 —— 空态确认（"初始化治理画像"按钮）→ 表单预填 9 项范式 v1 根级文件（.env/AGENTS.md/architecture_map.md/plan.md/project_rules.md/README.md/TECH_STACK_BASELINE.md/global_skills.md/project_skills.md，未修改）→ template_source 填 `manual://psco-phase13-dogfooding` → 8 项资产全承接（entry_ref 均为资产自身根级路径，前 5 项填结构化摘要）→ 保存成功回流回看态；全程无控制台错误，关键步骤截图留档（step1~step7）。
  - [x] SubTask 2.2: 验证概览区只读字段与后端 `RootFrozen*` 常量一致（`track_type=durable_system`、`docs_workflow_layout=phase/fix/audit/review`、`current_phase_name=phase13_project_governance_profile_foundation`、`current_phase_ref=plan.md#phase13_project_governance_profile_foundation`、`current_phase_status=in_progress`）
    - 执行记录：概览区实际显示 track_type=Durable System Track、docs_workflow_layout=phase/fix/audit/review、current_phase_name=phase13_project_governance_profile_foundation、current_phase_ref=plan.md#phase13_project_governance_profile_foundation、current_phase_status=进行中（in_progress），五项全部与后端冻结常量一致。
  - [x] SubTask 2.3: 验证表单保存成功回流（精准失效刷新）与刷新后回看数据一致；摘要回看区以 `structured_summary + entry_ref` 为主，真实路径不成为主视觉
    - 执行记录：保存后无整页刷新直接回到回看态；重新导航刷新后概览只读字段、8 项资产承接、9 项 canonical 根级文件、template_source 全部与保存时一致；摘要回看区以"名称/角色/可回源摘要/入口链接"呈现，structured_summary 为主视觉，真实路径非主内容。

- [x] Task 3: agent 读取路径 dogfooding 与固定 `6` 问取证
  - [x] SubTask 3.1: 对固定 `repository_id` 实时调用 `GetProjectBrief` 取证：7 顶层块完整、与 `GetGovernanceProfile` 同源、`products/modules/decisions` 保持数组形态、`current_phase` 从主记录派生、无目录扫描字段与自然语言指导词
    - 执行记录：`GetProjectBrief`（画像创建后）HTTP 200，7 顶层块完整：`repository`（ca261521.../personal-software-company-os）、`governanceProfile`（projectProfileVersion=project_governance_profile_v1、trackType=TRACK_TYPE_DURABLE_SYSTEM、templateSource=manual://psco-phase13-dogfooding、canonicalRootFiles[9]、globalAssetBindings[8]、currentPhaseName/Ref/Status）、`globalAssets[8]`、`currentPhase`、`products[1]`（PSCO）、`modules[1]`（project-context-foundation）、`decisions[1]`；`GetGovernanceProfile.profile` 与 brief 的 governanceProfile 逐字段一致（repositoryId/version/trackType/templateSource/docsWorkflowLayout/canonicalRootFiles/globalAssetBindings/currentPhase*/timestamps）；三数组长度为 1 仍保持数组形态未退化；`currentPhase.{name,entryRef,status}` 与主记录三 read-only 字段单向派生一致；响应字段面无目录扫描、无 Git 状态、无自然语言指导词字段。
  - [x] SubTask 3.2: 核对 `AGENTS.md`（推荐阅读顺序 / 当前阶段状态）与 `plan.md`（`phase13` 阶段路线）固定入口可回答当前阶段状态
    - 执行记录：`AGENTS.md` §1 项目定位声明"当前阶段：phase13_project_governance_profile_foundation（已建立正式 /plan 入口）"且主目标与画像交付一致，§5 推荐阅读顺序完整；`plan.md` L7 当前状态、L10 当前目标、L184-191 phase13 路线块（目标/进入条件/范围约束/交付要求/状态）均可直接回答 phase13 当前状态。
  - [x] SubTask 3.3: 固定 `6` 问逐题回答并按 `answer / direct entry refs / 是否达标` 格式留档，形成 `6 / 6` 达标结论
    - 执行记录：6 问逐题回答已留档于 `acceptance_report.md` §5（Q1 版本/路线 → brief.governanceProfile.project_profile_version+track_type；Q2 根级文件集合 → brief.canonical_root_files[9] 与表单回看；Q3 规范资产结构化承接 → brief.global_assets[8] 前 5 项含 structured_summary；Q4 同一 repository_id 驱动无第二套扫描 → 请求体唯一参数 + 响应字段面；Q5 唯一前端承接位 → repository-binding-detail-page.tsx L306 唯一挂载 + 浏览器矩阵 12 页无第二入口；Q6 未进入四类非目标 → RPC 清单/迁移清单/plan.md L188 范围约束），结论 6/6 达标。

- [x] Task 4: 补齐 `GetProjectBrief` 自动化集成测试（承接 phase13-10 复核留档建议 1）
  - [x] SubTask 4.1: 在 `backend/internal/projectcontext/connect` 既有集成测试中新增 `GetProjectBrief` 用例：画像已创建（走 service 正式写路径准备）→ 7 顶层块完整 + 同源断言；画像未创建 → `not_found`；测试自建 fixture 不依赖共享开发库 dogfooding 样本
    - 执行记录：核对发现 brief 4 场景（repository not found / profile not found / 7 顶层块完整 / 空数组合法）已在 phase13-10 实现中落地于 `server_integration_test.go`，自建 fixture（`reset_phase06_acceptance.sh --fixture`）且画像准备走 `store.SaveProfile` 正式入口，符合 spec 冻结口径；本轮补齐缺口：新增 phase13-11 同源断言（template_source / track_type / canonical_root_files / global_assets round-trip 一致 + current_phase 从主记录三 read-only 字段单向派生断言），新增 `canonicalRootFilesContain` / `globalAssetsContain` 取证 helper。
  - [x] SubTask 4.2: 运行新增测试通过，且既有测试（含旧 RPC 4 场景反回归）不回归
    - 执行记录：`go test ./internal/projectcontext/connect/...` 通过（19.217s），旧 RPC 4 场景与 brief 4 场景全部 PASS。

- [x] Task 5: 最小工具链验证（单值顺序）
  - [x] SubTask 5.1: `proto/` 下 `buf build` + `buf lint` 通过
    - 执行记录：`buf build` 与 `buf lint` 均通过，无输出、无 warning（Task 6 缺陷修复后重跑：proto 无改动，再次通过）。
  - [x] SubTask 5.2: `backend/` 下 `go test ./...` 通过（含 Task 4 新增用例）
    - 执行记录：首轮全部 package `ok`（`projectcontext/connect` 含新增同源断言，实跑 19.217s 通过）；Task 6 缺陷修复后完整重跑：全部 package `ok`（`projectcontext/connect` 等集成包 cache 命中、代码输入未变，`review` 相关包 `go build`+`go vet`+`go test` 通过）；重跑前已对共享开发库做 pg_dump 备份，重跑后复验 `GetProjectBrief` 与 `GetWeeklyReviewContext` 均 200，画像与 dogfooding 样本未受影响。
  - [x] SubTask 5.3: `frontend/` 下 `npm run build` 通过
    - 执行记录：`vite build` 成功（首轮 built in 1.83s；缺陷修复后重跑 built in 1.45s），产物体积正常，无 warning。

- [x] Task 6: 浏览器反回归矩阵
  - [x] SubTask 6.1: 主验证页 `/repositories/ca261521-8daf-4248-8f12-43525326e759`：治理画像三区（概览 / 维护表单 / 摘要回看）齐全且数据正确
    - 执行记录：主验证页治理画像区三部分齐全（概览卡片 / 编辑表单可进入 / 摘要回看），"维护治理信息"按钮可进入编辑态、"取消"可退回；截图 page1-repository-detail.png。
  - [x] SubTask 6.2: 跟随回归页逐页验证正常加载与不回归：`/dashboard`、`/modules` 列表 + 固定模块详情、`/decisions` 列表 + 固定决策详情、`/products` 列表 + 固定产品详情、`/repositories` 列表、`/onboarding`、`/reviews/daily`、`/reviews/weekly`
    - 执行记录：首轮矩阵 12 页中 11 页通过；`/reviews/weekly` 暴露回归缺陷（"Review 上下文加载失败"，后端 `GetWeeklyReviewContext` HTTP 500 空 body）。
    - 缺陷定位与修复：根因是 phase08 遗留 nil 解引用 —— `backend/internal/review/connect/server.go` 的 `domainModuleReuseSummariesToProto` / `domainCapabilitySummariesToProto` 直接解引用 `*item.LatestReuseAt` / `*item.LatestCapabilityUpdateAt`（均为 `*time.Time` 可空指针），当库中存在"从未被复用的模块"（phase06 fixture 的 auth-service，`latest_reuse_at` 为 nil）时 panic；此前验收 fixture 恰好全部模块均有复用时间故未暴露。修复：对齐 `reusesummary/connect/server.go` 既有 nil 保护口径，两处改为判空后赋值（最小修复，无行为变更，nil 时 proto 字段省略与 reuse_summary 主线一致）。
    - 修复验证：后端同端口 8081 重启（kill 旧进程 → 重新编译启动，未新开端口）后 `GetWeeklyReviewContext` HTTP 200；浏览器复验 `/reviews/weekly` 正常渲染全部区块（系统概览/模板候选/派生智能提示/最近活动/代表性反馈信号/复用感知快照），无错误态；修复后完整重跑工具链三步全部通过（见 Task 5）；复验画像与样本未受影响。
  - [x] SubTask 6.3: 验证无第二套治理画像入口（治理画像主内容区仅存在于 Repository detail）
    - 执行记录：矩阵全部 12 页检查，"项目治理画像"区域仅出现在 Repository detail 主验证页；代码层 grep 佐证：`GovernanceProfileSection` 唯一页面消费入口为 `repository-binding-detail-page.tsx`（L27 import / L306 挂载），切片内其余为组件/数据/应用层内部文件。

- [x] Task 7: 边界证据留档与验收记录收口
  - [x] SubTask 7.1: 留档四类非目标边界证据（不做 Git 推进跟踪 / 模板仓库接入 / 自动同步与目录扫描 / agent 写回与 MCP / CLI），证据来源：RPC 清单、brief 字段面、数据库迁移、根级与阶段文档明确不做清单
    - 执行记录：四类边界证据已留档于 `acceptance_report.md` §8 —— RPC 清单（governance_profile 仅 Get/手工 Update；project_context 纯读）、brief 字段面（无扫描/Git/指导词）、迁移 0010 关键字检索零命中（git/template_repo/scan/auto_sync/bootstrap/clone/pull）、`plan.md` L188 与 `AGENTS.md` §1/§4 明确不做清单。
  - [x] SubTask 7.2: 按固定 rerun 记录格式产出本目录 `acceptance_report.md`（样本与 `repository_id`、Web 页面与 agent 入口、样本恢复与画像维护方式、`6` 问逐题结果、工具链逐步结果、浏览器矩阵、边界证据、失败点与是否达标）
    - 执行记录：`acceptance_report.md` 已按 spec 冻结的 8 段 rerun 记录格式产出（§1 样本 / §2 入口 / §3 恢复与维护方式 / §4 双侧 dogfooding / §5 六问逐题 / §6 工具链逐步 / §7 浏览器矩阵 / §8 边界证据 / §9 失败点与 rerun / §10 达标结论 / §11 rerun 指引）。
  - [x] SubTask 7.3: 子代理独立复核（覆盖固定 `6` 问可复验性、工具链结果真实性、dogfooding 双侧取证完整性、边界证据充分性），修复全部阻断性问题后勾选收口
    - 执行记录（2026-08-17 独立复核代理）：全部结论基于独立实调（GetProjectBrief/GetGovernanceProfile/GetWeeklyReviewContext 实调、buf build+lint 重跑、go build+vet 重跑、proto/迁移/前端源码逐项核对），未采信原执行记录自述。结果：A 固定 6 问可复验性 PASS（6/6 逐题实调复验，同源 12 字段独立比对一致）；B 工具链真实性 PASS（nil 判空修复 L212-216/L231-234 在位，weekly 实调 200）；C 双侧 dogfooding 完整性 PASS（RootFrozen 常量于 types.go L52-67 独立核实五项一致）；D 边界证据充分性 PASS（四类证据逐项交叉核实）；E 报告一致性 PASS（rerun 8 项内容全覆盖）。阻断性问题：无。非阻断建议 3 条：① plan.md L190 状态 planned 与画像 in_progress 语义位不同步，留待 phase13-12 根级收口时同步；② 报告 §6 补 review 包三重复核说明（已补）；③ checklist 整体勾选 + 本记录留档（本条即执行）。总结论：通过。

# Task Dependencies

- `Task 2` depends on `Task 1`
- `Task 3` depends on `Task 2`
- `Task 5` depends on `Task 4`
- `Task 6` depends on `Task 2`（画像维护完成后主验证页才有完整数据）
- `Task 7` depends on `Task 3`, `Task 5`, `Task 6`
