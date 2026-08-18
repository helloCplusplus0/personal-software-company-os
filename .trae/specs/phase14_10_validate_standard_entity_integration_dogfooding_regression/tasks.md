# Tasks

- [x] Task 1: 验收环境准备与固定样本补建
  - [x] SubTask 1.1: 启动验收环境（后端 8081 / 前端 5173 / psco_development 连通）；四实体基座确认（ca261521 / f0d034cc / 9b02e0ca / aa8ee5ad 在库 + `GetProjectContext` 对固定 repository_id 一次解析成功；缺失则执行 `database/scripts/restore_phase11_phase12_dogfooding_sample.sh` 兜底）
  - [x] SubTask 1.2: Standard 固定样本正式化（盘点既有 dogfooding Standard 含 `899ed9f3-614a-4f19-9c63-9da7a47c3db4`；满足 spec 语义集合 1-4 则正式化，否则 web 新建；核对迁移产物 count=0；固定 standard_id 留档）
  - [x] SubTask 1.3: web 维护会话补齐固定样本缺口（树语义集合对齐：嵌套 directory / ≥5 file 节点含四项根级关键文档 / role+summary+ref 实例含 https:// 形态 / status=active；5 行绑定覆盖八格矩阵 5 合法格——repository→template_source，product/module/decision→adopts；至少一次 Update 产生 revision）

  > 执行记录：环境三方连通（healthz=200 / 5173=200）；基座曾缺失（集成测试重置）经 `restore_phase11_phase12_dogfooding_sample.sh` 恢复，四 id 与 phase13-11 冻结值一致，`GetProjectContext` 解析成功；standards 三表曾全空（`899ed9f3` 已不存在，正式化路径=**web 新建**）；经浏览器 dogfooding 会话创建固定 Standard `85a5d8b7-f41a-44ed-8f6f-421e548b53ed`（Durable System Track 项目范式 / active），一次 Update 补全 ref/summary 并产生首条 revision；绑定 5 行（repository/template_source + repository/product/module/decision 各 adopts，5 合法格全覆盖）；迁移产物 count=0；截图归档 screenshots/（step1~step6 关键步骤）。
- [x] Task 2: agent 读取路径 dogfooding 与固定 6 问取证（执行于 Task 7 解耦完成后）
  - [x] SubTask 2.1: `GetProjectBrief`（固定 repository_id）实时取证：standards[] 含固定 Standard 全树（嵌套 + file 节点 role/summary/ref）；5 顶层块字段面核对（槽位 2/3/4 reserved；repository / products / modules / decisions / standards[]；无 global_assets / 无 governance_profile / 无 current_phase / 无扫描 / 无 Git / 无正文）；`GetStandard` 同源性核对
  - [x] SubTask 2.2: 固定 6 问（Q1 树形直答 / Q2 退役零丢失 / Q3 单规范多复用 / Q4 revision 留痕 / Q5 唯一维护入口 / Q6 画像残余彻底退役与时间轴未偷渡）逐题 answer + direct entry refs + 是否达标 留档

  > 执行记录（2026-08-18，Task 7 解耦后执行）：`GetProjectBrief`(ca261521) HTTP 200 × 2 次，顶层字段恰 5 个（decisions/modules/products/repository/standards），governanceProfile/currentPhase/globalAssets 全部 present=False；standards[] 含固定 Standard `85a5d8b7` 全树（根"." + docs 目录 + 6 file 节点，5 个 `/` ref + 1 个 `https://` ref）；`GetStandard` 同源性：directoryTree 规范化 JSON 全等 + 6 标量字段全 equal=True。Q1-Q6 = 6/6 达标（Q3 实测 5 行绑定 target_id 全匹配；Q4 实测 2→3 条 revision 含 change_summary；Q5 写 RPC 调用 100% 在 features/standard/ 切片 + backend 无 MCP/CLI/agent 写回命中；Q6 主表 NULL + 模块不存在 + migrations 仅到 0012 + 无时间轴活字段）。证据文件 /tmp/brief_task2.json、/tmp/getstandard_task2.json。
- [x] Task 3: 画像退役触点零丢失复验（T1-T7 + 裁决⑧两库形态）
  - [x] SubTask 3.1: T1 proto / T2 存储（三表 to_regclass NULL）/ T3 后端模块（目录不存在 + grep 零命中）/ T5 前端 / T7 画像残余（reserved + 消息删除 + 0012 应用 + 豁免留痕）机械检查逐条执行留痕（命令 + 结果）
  - [x] SubTask 3.2: T4 画像 RPC 404 实测（对照 ListStandards 200 / healthz 200）；T6 brief 字段面核对（5 顶层块 + 无画像派生消息）
  - [x] SubTask 3.3: 裁决⑧两库形态：验收库 0011 `psql -f` 重放幂等无产物；`默认项目范式（迁移自治理画像）` count=0；phase14-09 `migration_evidence.md` §3/§5 一次性验证库证据引用复核

  > 执行记录（2026-08-18）：T1-T7 + 裁决⑧全 PASS 无 FAIL。T1：proto 画像包目录不存在 + buf.yaml 豁免均为单文件粒度（ignore 单文件 + ignore_only 规则级单文件）无目录前缀扩大化 + 生成产物命中仅为 project_context 裁决注释与 reserved 名编码残留；T2：三表 to_regclass 均 NULL；T3：backend/internal/governanceprofile 不存在 + grep 零命中（仅 router.go:433 一行裁决留痕注释）；T4：画像 RPC 404 / ListStandards 200 / healthz 200；T5：前端目录不存在 + grep 零命中；T6：proto 画像派生消息仅 1 处裁决注释命中；T7：双 reserved（L198-199）+ schema_migrations 含 0012 + 豁免无新增。裁决⑧：0011 psql 重放 exit=0 全幂等 NOTICE + 迁移种子 count=0 + standards 总数 1 固定样本完好 + migration_evidence.md §3/§5 章节存在。
- [x] Task 4: 八项裁决验收门禁逐条取证（shared_baseline §4 ①-⑧矩阵）
  - [x] SubTask 4.1: ①混合式颗粒度 / ②主表+jsonb树+多态绑定 / ③正文零托管 取证（浏览器 + psql \d + brief JSON + grep）
  - [x] SubTask 4.2: ⑤无第五主实体扩散 / ⑥树编辑器无拖拽 / ⑦绑定唯一发起位 取证（grep 断言 + 浏览器抽查；④⑧由 Task 3 结果汇入）

  > 执行记录（2026-08-18）：①PASS（brief 树 JSON 仅六字段无正文承载；directory"docs"与 file 节点均带 summary；standard-tree-view.tsx L55-76 directory/file 双侧渲染 summary）；②PASS（directory_tree jsonb NOT NULL + 四元组唯一约束 + target_type CHECK 四类 + 4 类实例行 1/1/1/2）；③PASS（DirectoryTreeNode 仅 name/node_type/role/summary/ref/children + backend/internal/standard 抓取类符号零命中）；⑤PASS（dashboard 目录零命中 + 交叉引用仅 repository-binding detail 只读摘要 + __root.tsx NAV_ITEMS /standards 独立导航项）；⑥PASS（拖拽符号零命中 + 交互原语 handleAddRootChild/replaceChild/removeChild/swapChildren 按钮式）；⑦PASS（bind/unbind caller 4 文件全在 features/standard + 挂载点唯一 standard-detail-page L182-184 + 四实体页面零绑定入口）；④⑧由 Task 3 结果汇入（全 PASS）。
- [x] Task 5: 工具链四步门禁 + breaking 豁免专项（单值顺序完整执行）
  - [x] SubTask 5.1: proto/ `buf build` → `buf lint` → `buf breaking`（豁免留痕生效验证：ignore 单文件 + ignore_only 规则级单文件〔含 T7 触发的 brief 字段/消息移除最小豁免〕无扩大化，退出码 0）
  - [x] SubTask 5.2: backend/ `go build ./... && go vet ./... && go test ./...`（重跑前 pg_dump 备份验收库，重跑后复验固定样本在库）
  - [x] SubTask 5.3: frontend/ `npx tsc -b && npm run build`（含 UI 反馈轮两处修复的编译回归）

  > 执行记录（2026-08-18，7/7 步全 PASS）：步骤 0 备份 pg_dump 32595B（/tmp/psco_dev_backup_phase14_10_task5.sql）；步骤 1 buf build / 步骤 2 buf lint 均退出码 0 无 warning；步骤 2b make breaking（`buf breaking --against '../.git#branch=main,subdir=proto'`）退出码 0，buf.yaml 豁免确认两条均为精确单文件（ignore: governance_profile.proto 单文件 + ignore_only: FIELD_WIRE_JSON_COMPATIBLE_TYPE → project_context.proto 规则级单文件）无目录前缀扩大化；步骤 3 go build/vet/test 退出码 0（9 包 ok 0 FAIL，8 包缓存命中——代码自 Task 7 门禁实跑〔18.6s 集成全绿〕后未变更故缓存有效，standard/service 实跑 0.194s）；步骤 3b 固定样本复验 std_ok=1/bindings=5/revisions=3/repo/prod/mod/dec 各 1 全达标（未触发恢复协议）；步骤 4 tsc -b 零错误 + vite build 2462 modules 1.35s（含 standard 全部页面产物）。
- [x] Task 6: 浏览器反回归矩阵 16 页逐页验证
  - [x] SubTask 6.1: `/standards` 四页（列表 / 详情 / 创建 / 编辑——含输入焦点稳定回归与无拖拽验证）+ Repository detail 让位后回归（第 5 页，含 T7 解耦回归：Standard 摘要经 brief 正常加载无 404）
  - [x] SubTask 6.2: 四实体列表/详情抽查（第 7-13 页）+ 既有基线 8 页（dashboard / onboarding / reviews 双路径，第 6/14/15/16 页）；每页通用检查点（加载 / 无 not_found / 无画像残留）+ 专属检查点逐页留档

  > 执行记录（2026-08-18，两轮浏览器会话）：**前置修复**——Task 2 取证发现固定 Standard status=draft（违反 spec 语义集合 4），经 web 编辑会话正式修复为 active（change_summary="恢复固定样本 active 状态（phase14-10 验收口径）"，产生第 3 条 revision），详情页确认 active；16 页矩阵全 PASS：p01 列表含固定样本行（active）/ p02 详情五要素齐全（树形 + 绑定 5 行 + Revision ≥3 条含恢复条目 + 返回行/语义导语/左摘要右内容 grid）/ p03 创建表单树编辑器可用 / p04 编辑页连续输入不跳焦（UI 反馈修复回归通过）+ 无拖拽 / p05 Repository detail 画像已让位 + Standard 只读摘要正常加载（T7 解耦回归无 404）/ p06 dashboard 无 Standard 主卡片 / p07-13 四实体列表详情正常 / p14 onboarding 正常 / p15-16 review 双路径正常；全部页面无 not_found、无画像残留文案、无 console error。截图归档 screenshots/task6-p01~p16.png（16 张）。
- [x] Task 7: 裁决触发修复——brief 画像残余解耦与失败点处置（已触发：Task 2 首验实证画像残余依赖阻断 brief 双消费路径）
  - [x] SubTask 7.1: proto：`GetProjectBriefResponse` 移除 `governance_profile`(2)/`current_phase`(4) 并 reserved 号+名；删除 `BriefGovernanceProfile` / `BriefCurrentPhase` / `BriefTrackType` / `BriefPhaseStatus`；`make gen` 三端再生成
  - [x] SubTask 7.2: backend：删除 `backend/internal/governanceprofile/` 整目录；`projectcontext`（types / candidate / service / connect / 集成测试）与 `connecterrors` / `platform/server.go` 联动去引用；brief 装配移除画像读取与 current_phase 派生
  - [x] SubTask 7.3: database：新增 `0012_phase14_brief_profile_decoupling.sql`（`DROP TABLE IF EXISTS governance_profiles`，幂等，裁决留档注释）；`RunMigrations` 编入并在验收库应用（主表 `to_regclass` NULL + `schema_migrations` 含 0012）
  - [x] SubTask 7.4: `buf.yaml` breaking 豁免最小扩展（仅 `project_context.proto` 单文件规则级，无目录前缀扩大化）留痕
  - [x] SubTask 7.5: 修复后门禁完整重跑（buf build / lint / breaking + go build / vet / test〔pg_dump 保护协议〕+ tsc / build）+ 同端口重启验收环境
  - [x] SubTask 7.6: 失败点 / 恢复路径 / rerun 结果留档（沿袭 phase13-11 §9 格式）

  > 执行记录（2026-08-18，子代理实现 + 门禁全绿）：
  > - 变更：proto 双 reserved + 4 画像派生消息删除 + 注释裁决留痕；`make gen` 三端再生成零残留；`backend/internal/governanceprofile/` 整目录删除（5 文件）；projectcontext 五文件收缩（`NewQueryService` 签名减参、brief 编排 5 步、集成测试新增 `assertNoProfileRemnants` 断言 5 顶层块无画像残余）；connecterrors 删 2 哨兵；platform/router+server 去装配；新增 `0012` 迁移（幂等 drop 主表）。
  > - 0012 应用：后端启动 RunMigrations 正式入口，日志 `migration applied version=0012_...`，二次启动 `no new migrations`（幂等）；`to_regclass('governance_profiles')` = NULL；`schema_migrations` 含 0012（applied_at 2026-08-18 08:56:14+00）。
  > - 门禁重跑：buf build / lint / breaking（0，无 warning）、go build / vet / test（0，集成 18.6s 全绿）、tsc / vite build（0）；测试前 pg_dump 备份（/tmp/psco_dev_backup_phase14_10_t7.sql，34354B）。
  > - breaking 豁免：**实测无需追加**（WIRE_JSON 规则集下号+名双 reserved 已满足；MESSAGE_DELETE/ENUM_DELETE 不在该集合）——豁免面零扩大，优于指令预设；既有豁免原样。
  > - 失败点/恢复：go test 集成重置清空四实体基座（standards 三表完好）→ 备份 COPY 整段重放因同名异 id fixture 唯一键冲突中止 → 改 `INSERT ... ON CONFLICT DO NOTHING` 逐行恢复（`decisions.impact` 按 schema 语义空字符串）→ 最终库形态：四基座 1/1/1/1 + decision_link/product_repositories/module_repositories 各 1 + standard `85a5d8b7` 1 + bindings 5 + revisions 1。
  > - 冒烟（8081 同端口重启，healthz 200）：`GetProjectBrief`(ca261521) **200**，顶层 keys = decisions/modules/products/repository/standards（5 实际顶层字段，槽位 2/3/4 reserved，无 governanceProfile/currentPhase/globalAssets），standards[] 含固定 Standard 全树；`GetProjectContext` 200；画像 RPC 404（对照）。
  > - grep 复核：`governanceprofile` backend/internal 零命中（6 处注释措辞改写为中文描述，语义保留）；`Brief*` 画像消息仅 proto 裁决注释与 buf.yaml 既有注释 2 处历史留档命中。
- [x] Task 8: acceptance_report 冻结 + 独立复核 + 收口
  - [x] SubTask 8.1: 按沿袭结构冻结 `acceptance_report.md`（固定样本与 id / 固定入口 / 双侧 dogfooding / 固定 6 问 / 工具链逐步结果 / 浏览器矩阵 / 八项裁决门禁矩阵 / 边界证据 / 失败点与 rerun / 是否达标 / Rerun 指引）
  - [x] SubTask 8.2: 子代理独立复核（样本协议一致性 / 六触点证据真实性 / 八项裁决取证完整性 / 工具链顺序与结果 / 矩阵覆盖完整性 / 范围外改动）；阻断问题修复后复核通过
  - [x] SubTask 8.3: tasks.md / checklist.md 全部勾选附执行记录；变更保持未提交，待用户最终确认后手动提交

  > 执行记录（2026-08-18）：acceptance_report.md §1-§13 冻结（沿袭 phase13-11 结构 + 八项裁决矩阵 + T7 裁决与 CON-08 口径变更留痕）；独立复核（independent_review.md）**PASS：0 阻断 / 4 观察项**——R1-R7 全 PASS，关键佐证：`go test -count=1` 强制实跑 18.812s 集成全绿（缓存命中说明如实）、备份 32595B 同字节数（库形态未漂移）、git status 19 项变更全部落在 phase14 交付面；4 观察项处置——①checklist 已回填、②T3 注释命中时点差（现状更干净，报告如实记录取证时点）、③裁决⑦文件计数口径差异（实质结论一致：caller 全在切片内）、④go test 重跑清基座行为已入报告 §10 失败点表。DoD 四项全达标：6 问 6/6 + 八项裁决全绿 + 16 页矩阵全绿 + 报告冻结 + 复核通过。变更未提交，待用户确认后手动提交。

# Task Dependencies

- Task 2, Task 3, Task 4, Task 5, Task 6 depend on Task 1（固定样本就绪是全部取证前提）
- Task 2-6 的取证执行于 Task 7 解耦落地之后（2026-08-18 裁决触发，Task 7 已前置执行）
- Task 5 的 SubTask 5.2 依赖 Task 1 样本在库（pg_dump 保护协议需要样本基线）
- Task 8 depends on Task 2, Task 3, Task 4, Task 5, Task 6, Task 7
