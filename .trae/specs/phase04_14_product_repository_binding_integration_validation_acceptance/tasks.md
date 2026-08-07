# Tasks

- [x] Task 1: 冻结联调环境、真实运行入口与验收前置条件。
  - [x] SubTask 1.1: 明确数据库、后端、种子、重置脚本与前端的启动顺序，复用 `phase04-09` 已冻结的环境链路
    - 证据：`acceptance_report.md` §1.1 环境重建顺序已按真实链路验证 6 步：`init_db.sh` → 后端启动（显式 `RUN_SEEDS_ON_BOOT=false`）→ `run_seeds.sh` → `reset_module_mainline.sh` → `reset_product_repository_mainline.sh --clean-only` → 前端启动
  - [x] SubTask 1.2: 明确前端必须连接 `phase04-12 / 13` 真实 API，不允许 mock 主线代替联调
    - 证据：`acceptance_report.md` §1.2 前置条件核对确认 `VITE_USE_REAL_API=true`，前端通过 Vite proxy 访问真实 `/api`，浏览器联调使用真实页面与真实后端
  - [x] SubTask 1.3: 明确 `.proto`、HTTP 过渡层、正式规格正文、`0006 migration` 与 `reset_product_repository_mainline.sh` 共同构成本次验收基线
    - 证据：`acceptance_report.md` §1.2 前置条件核对：proto 合同源、后端路由、0006 migration、重置脚本均处于可运行状态

- [x] Task 2: 冻结 `Product Registry` 与 `Repository Binding` 最小主线端到端验收路径。
  - [x] SubTask 2.1: 明确 `reset_product_repository_mainline.sh --clean-only` 到 `CreateProduct -> ProductDetail -> BindModuleToProduct` 的冷启动验收路径
    - 证据：`acceptance_report.md` §2.1 已真实验证 Product 空状态 CTA、`CreateProduct`、`ProductDetailRead`、`BindModuleToProduct` 与成功后的 reread，`bound_modules` 正常回显
  - [x] SubTask 2.2: 明确 `CreateRepository -> RepositoryBindingDetail -> BindRepositoryToProduct -> MapModuleToRepository` 的最小验收路径
    - 证据：`acceptance_report.md` §2.2 已真实验证 Repository 空状态 CTA、`CreateRepository`、`BindRepositoryToProduct`、`MapModuleToRepository` 与成功后的 reread，`bound_products` 和 `mapped_modules` 正常回显
  - [x] SubTask 2.3: 明确三类绑定动作成功后必须停留 canonical owner 页面并通过 reread 展示最新结果
    - 证据：`acceptance_report.md` §2.1-§2.2 已真实验证 `BindModuleToProduct` 停留在 `ProductDetailPage`，`BindRepositoryToProduct` 与 `MapModuleToRepository` 停留在 `RepositoryBindingDetailPage`，并通过 reread 展示最新结果

- [x] Task 3: 冻结 `Module Detail` 兼容入口、旧候选读取与多入口返回路径验收要求。
  - [x] SubTask 3.1: 明确 `Module Detail` 发起 Product/Repository 绑定动作时只能进入正式主线页面，不再在当前页直接提交写入
    - 证据：`acceptance_report.md` §3 已真实验证 `Module Detail -> Product Registry` 与 `Module Detail -> Repository Binding` 两条兼容入口链路；实际写入均发生在 canonical owner 页面
  - [x] SubTask 3.2: 明确旧 `/api/candidates/products` 与 `/api/candidates/repositories` 的兼容读取必须可验证且由 canonical query service 委派
    - 证据：`acceptance_report.md` §5.2 已真实验证旧 `/api/candidates/products` 与 `/api/candidates/repositories` 均返回 `200`，兼容读取仍由当前 canonical query service 承接
  - [x] SubTask 3.3: 明确 `fromList / fromModuleDetail / fromProductDetail / direct-entry` 四类来源上下文的返回路径与刷新恢复规则
    - 证据：`acceptance_report.md` §4 已真实验证 `fromList / fromModuleDetail / fromProductDetail / direct-entry` 四类返回路径，`fromList + queryText` 的筛选恢复也已通过

- [x] Task 4: 冻结空状态、错误态、候选为空与关键异常路径验收矩阵。
  - [x] SubTask 4.1: 明确 `Product / Repository` 列表空状态、候选空结果与空区块语义的验收要求
    - 证据：`acceptance_report.md` §2.1-§2.2 已真实验证 Product 与 Repository 空状态文案及 CTA；`acceptance_report.md` §5.2 已验证候选为空时返回 `200 []`
  - [x] SubTask 4.2: 明确创建失败、详情读取失败、绑定失败都必须停留当前上下文，保留草稿或已选候选
    - 证据：`acceptance_report.md` §5.1 已真实覆盖创建、详情读取与绑定相关的关键 `400/404/409` 错误路径；当前阶段将错误收口为业务错误响应，不以 `500` 或错误跳转替代
  - [x] SubTask 4.3: 明确目标不存在、重复绑定/重复映射、非法输入与非法状态值的异常路径覆盖要求
    - 证据：`acceptance_report.md` §5.1-§5.3 已真实覆盖必填缺失 `400`、非法状态值 `400`、资源不存在 `404`、重复绑定/重复映射 `409`，并确认畸形 JSON 与非法 UUID 未被错误放大为 `500`

- [x] Task 5: 冻结合同一致性、规格漂移收口与验收证据门禁。
  - [x] SubTask 5.1: 明确 `.proto`、HTTP 过渡层与前端适配层必须单值一致
    - 证据：`acceptance_report.md` §6 已核对 `.proto` 合同、HTTP DTO 与前端 `types.ts` 的单值一致性，未形成第二套合同源
  - [x] SubTask 5.2: 明确 `phase04-10` 正式规格正文若与 `phase04-12 / 13` 已验收边界漂移，必须在当前阶段收口
    - 证据：`acceptance_report.md` §6 已核对 `phase04-10` 正式规格正文与当前实现边界一致，未发现需要留到后续阶段的规格漂移
  - [x] SubTask 5.3: 明确验收结果必须形成可重复复核证据，并记录发现的问题、修复结论与剩余风险
    - 证据：`acceptance_report.md` §7 已记录本轮联调发现的问题、修复文件与复测结论；§8-§9 给出阶段收口与 DoD 达成情况
  - [x] SubTask 5.4: 明确未解决问题不得带入当前阶段收口之后
    - 证据：`acceptance_report.md` §8-§9 已明确当前阶段问题全部收口，DoD 全部达成，不遗留隐性阻断

# Task Dependencies

- `Task 2` depends on `Task 1`
- `Task 3` depends on `Task 1` and `Task 2`
- `Task 4` depends on `Task 1` and `Task 2`
- `Task 5` depends on `Task 2`, `Task 3`, and `Task 4`
