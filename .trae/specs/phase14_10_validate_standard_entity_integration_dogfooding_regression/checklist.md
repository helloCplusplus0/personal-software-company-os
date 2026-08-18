# phase14-10 Checklist

## 固定样本与验收环境（Task 1）

- [x] 四实体基座在库（ca261521 / f0d034cc / 9b02e0ca / aa8ee5ad）且 `GetProjectContext` 对固定 repository_id 一次解析成功（缺失时经 `restore_phase11_phase12_dogfooding_sample.sh` 兜底，无手工 SQL 插基座）
- [x] Standard 固定样本满足语义集合：单根 directory + ≥1 嵌套 directory；≥5 file 节点含 AGENTS.md / plan.md / project_rules.md / TECH_STACK_BASELINE.md；role 必填；≥4 节点 summary + `/` ref；≥1 节点 `https://` ref；status=active
- [x] 固定绑定 5 行在库：repository→template_source；product / module / decision→adopts（八格矩阵 5 合法格全覆盖）
- [x] `ListStandardRevisions` ≥1 条且 change_summary 为人工一句话（经 web Update 产生）——实测 3 条
- [x] Standard 与绑定 100% 经 `/standards` web 维护会话补建/正式化（无种子 SQL、无后门脚本）
- [x] 固定 standard_id 留档；正式化路径（既有正式化 / 新建）与依据在报告中说明——85a5d8b7，web 新建（报告 §1）
- [x] 验收库 `SELECT count(*) FROM standards WHERE name = '默认项目范式（迁移自治理画像）'` = 0

## agent 读取路径 dogfooding（Task 2）

- [x] `GetProjectBrief`（固定 repository_id）HTTP 200；standards[] 含固定 Standard 全树（嵌套 directory / file + role / summary / ref 完整，无转译直读）
- [x] 无画像残余块：`governance_profile` / `current_phase` 已移除（字段号 2/4 reserved）；template_source 语义经 Standard 绑定（role=template_source）承接
- [x] brief 字段面 = 5 顶层块（槽位 2/3/4 reserved；repository / products / modules / decisions / standards[]）；无 global_assets / 无 governance_profile / 无 current_phase / 无目录扫描 / 无 Git 状态 / 无第二套事实源投影 / 无正文内容字段
- [x] `GetStandard(固定 id)` 树与 brief 内该 Standard 树逐字段一致（同源性）——规范化 JSON 全等 + 6 标量字段 equal
- [x] 固定 6 问（Q1-Q6）逐题 answer / direct entry refs / 是否达标 留档，6/6 达标

## 画像退役六触点复验（Task 3）

- [x] T1 proto：画像包目录不存在；`buf.yaml` 豁免留痕在且无扩大化（ignore 单文件 + ignore_only 规则级单文件）；三端生成产物无画像包（命中仅为裁决注释与 reserved 名编码残留）
- [x] T2 存储：三张画像表（主表 + 两张 bindings 表）`to_regclass` 均 NULL（主表经 0012 drop）
- [x] T3 后端：`backend/internal/governanceprofile` 目录不存在；`grep -r "governanceprofile" backend/internal` 零命中；platform 层无画像 RPC 挂载
- [x] T4 RPC：画像 Connect 请求实测 404（对照 ListStandards 200 / healthz 200）
- [x] T5 前端：`frontend/src/features/governance-profile/` 不存在；`grep -r "governance-profile" frontend/src` 零命中（合法注释命中逐条说明）
- [x] T6 brief：两清单信息唯一来源 standards[]；5 顶层块（槽位 2/3/4 reserved）；无 BriefGovernanceProfile / BriefCurrentPhase 等画像派生消息（proto 仅 1 处裁决留痕注释命中）
- [x] T7 画像残余（裁决触发 2026-08-18）：proto 字段号 2/4 + 名 reserved 且画像派生消息删除；`0012` 已应用（主表 `to_regclass` NULL + `schema_migrations` 含 0012）；buf breaking 豁免最小扩展（仅 project_context.proto 单文件）留痕无扩大化——实测号+名双 reserved 已满足 WIRE_JSON 规则，豁免面零新增；Repository detail Standard 摘要经 brief 正常加载（无画像前提不再 404）
- [x] 裁决⑧验收库形态：0011 `psql -f` 重放幂等（无报错无产物）；迁移产物 count=0（无源数据分支）
- [x] 裁决⑧一次性验证库证据：phase14-09 `migration_evidence.md` §3（恰 1 条 + revision 含 N/M 与源 repository）与 §5（零丢失对照）引用复核通过

## 八项裁决验收门禁（Task 4）

- [x] ① 混合式颗粒度：directory 与 file 节点 summary Web + brief 双侧可见；无正文字段（正文以模板仓库为唯一事实源）
- [x] ② 主表 + jsonb 树 + 多态绑定：`\d standards` directory_tree jsonb；standard_bindings 4 类 target_type 实例行 + 四元组唯一约束
- [x] ③ 正文零托管：file 节点仅 ref 定位引用；树 JSON / brief 响应无正文承载字段；代码无正文抓取逻辑
- [x] ④ 画像系统性退役：T1-T7 全绿（Task 3 结果汇入，含裁决触发 T7）
- [x] ⑤ 无第五主实体扩散：Dashboard 无 Standard 主卡片；四实体页面无 Standard CRUD 侵入（仅 Repository detail 只读摘要）；导航并列但独立
- [x] ⑥ 树编辑器无拖拽：`grep -rn "draggable\|onDragStart\|dnd" frontend/src/features/standard` 零命中；编辑会话全程经增删/上移下移/添加子节点完成
- [x] ⑦ 绑定仅 Standard 详情页发起：bind/unbind mutation caller 唯一（StandardBindingPanel）；四类目标实体页面无 Standard 绑定入口
- [x] ⑧ 存量合并迁移：两库形态复核通过（Task 3 结果汇入）；验收库无第二条迁移产物

## 工具链四步门禁 + breaking 专项（Task 5）

- [x] 步骤 1（proto/）`buf build` 退出码 0 无 warning
- [x] 步骤 2（proto/）`buf lint` 退出码 0 无 warning
- [x] 步骤 2b（proto/）`buf breaking`（.git 基准）退出码 0；豁免留痕生效且无目录前缀扩大化（含 T7 触发的最小规则级豁免）——实测豁免面零新增（号+名双 reserved 已满足）
- [x] 步骤 3（backend/）`go build ./... && go vet ./... && go test ./...` 全绿；重跑前 pg_dump 备份、重跑后固定样本复验在库（32595B 备份；1/5/3/1/1/1/1 全达标未触发恢复）
- [x] 步骤 4（frontend/）`npx tsc -b && npm run build` 零错误（含 UI 反馈轮两处修复的编译回归）
- [x] 工具链按单值顺序完整执行（缺陷修复后如需重跑，完整重跑一遍并留痕）——Task 7 修复后完整重跑一次 + Task 5 验收轮完整执行；独立复核实跑 `go test -count=1` 18.812s 集成全绿佐证

## 浏览器反回归矩阵（Task 6）

- [x] 第 1-4 页 `/standards` 列表 / 详情 / 创建 / 编辑：加载正常；列表含固定样本行；详情树形 + 绑定 5 行 + Revision 区 + 一致性布局在位；编辑页输入焦点稳定（连续输入不跳焦）+ 无拖拽
- [x] 第 5 页 Repository detail：画像区已让位（无"维护治理信息"入口）；Standard 只读摘要入口在且 compact 树正确；T7 解耦回归：摘要经 brief 正常加载无 404
- [x] 第 6 页 `/dashboard`：四区块正常，无 Standard 主卡片
- [x] 第 7-13 页 四实体列表/详情抽查 + `/repositories`：加载正常、数据行在、无画像残留
- [x] 第 14-16 页 `/onboarding` / `/reviews/daily` / `/reviews/weekly`：正常
- [x] 16/16 页通过；每页通用检查点（加载 / 无 not_found / 无画像残留）逐页留档（报告 §7）
- [x] web 维护会话全程无控制台错误；关键步骤截图留档（step-N 命名）——step*.png 会话截图 + task6-p01~p16.png 矩阵截图

## 裁决触发修复与失败点处置（Task 7，已触发）

- [x] brief 画像残余解耦落地：proto（reserved + 消息删除 + make gen）/ backend（模块删除 + 去引用）/ 0012 迁移（应用 + 幂等重放）/ buf 豁免最小扩展——逐项附证据（tasks.md Task 7 执行记录）
- [x] 修复后门禁完整重跑（四步 + 2b 专项）全绿留痕；同端口重启验收环境
- [x] 失败点 / 恢复路径 / rerun 结果留档（沿袭 phase13-11 §9 格式）；后续若再发现阻断缺陷沿用本节处置（报告 §10 三项全留档）

## 边界证据（四类非目标，随报告留档）

- [x] 无 Git 推进跟踪 / 无模板仓库接入 / 无自动同步与目录扫描入库 / 无 agent 写回（MCP / CLI / Draft / 审批流）——各附 grep / RPC 清单 / 表列检查证据（报告 §9；含 CON-08 时间轴口径变更留痕）

## DoD 门禁与收口（Task 8）

- [x] 固定问题全达标（固定 6 问 6/6）
- [x] 八项裁决验收门禁全绿（①-⑧ 逐条取证留档）
- [x] 反回归矩阵全绿（16/16 页）
- [x] `acceptance_report.md` 冻结验收结论（结构沿袭 phase13-11，含 Rerun 指引）
- [x] 子代理独立复核通过（阻断问题已修复回填）——independent_review.md：PASS（0 阻断 / 4 观察项，观察项均如实留档或现状更优）
- [x] tasks.md / checklist.md 全部勾选附执行记录
- [x] 变更保持未提交，待用户最终确认后手动提交
