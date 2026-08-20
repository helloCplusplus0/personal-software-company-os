# phase15-08 Checklist

## 环境与数据

- [x] 固定样本仓库 `ca261521-8daf-4248-8f12-43525326e759` 经恢复脚本重建，`GetProjectBrief` 应答 repository + standards[] 非空（背景侧就绪）
- [x] main-repo 上 5 条 phase01 测试事件全部删除（API 路径），删除记录留档；main-repo List 返回空
- [x] 录入前全库 progress_events = 0 行

## dogfooding 固定录入集（附件 A）

- [x] 16 条全部经 web 表单路径录入（真实浏览器驱动 ProgressEventForm，零 API 直插）
- [x] occurred_at 逐条等于附件 A 本地时刻（+0800，补录历史；DP-3 转换后 DB 侧为对应 UTC）
- [x] title/task_key/detail/evidence_ref 逐条与附件 A 一致；note 行无 task_key；全部 source=manual
- [x] 录入顺序 = 附件 A 行序；完成后 List 返回 16 条且三键链倒序与附件 A 逆行序一致
- [x] web 取证：时间轴 16 条倒序 + 当前卡"暂无进行中 phase"完结态 + latest 行 phase14-11 + 三轨过滤 phase=16/audit=0/fix=0

## 固定问题取证

- [x] brief 一次调用同体含背景（repository/products/modules/decisions/standards[]）与进度（currentPhaseKey 空串 / latestTaskCompleted=phase14-11 / recentEvents 恰 10 条=行 7~16）
- [x] List 全量 16 条三键链序 + phase 过滤 16 / audit 0 / fix 0（冻结态）；临时取证对下 17/14/2/1 分布验证后删除恢复 16 条（删除留档）
- [x] append-only：proto 与 connect 生成物均无 Update；重取 List 前 15 条零变形零丢失；progress_events 无派生列（派生仅响应侧）
- [x] 派生正确性：全部完结态命中（phase_completed phase14 晚于 phase_started）/ latest=phase14-11 / recent 截断 10 与附件 A 注 3 一致

## 十一项裁决门禁矩阵

- [x] ① brief 一次调用背景+进度同体（应答体摘录）
- [x] ② 16 条 repository_id 全等 ca261521（SQL）+ main-repo 隔离空
- [x] ③ append-only + 派生不落库（复用固定问题取证 3）
- [x] ④ 三轨可录可滤（临时取证对 audit_001/fix_001 表单录入→过滤分布→删除恢复）
- [x] ⑤ task_key 格式逐条（phase14/phase14-NN/audit_NNN/fix_NNN）；audit/fix 轨无边界事件
- [x] ⑥ 维护入口仅 Repository detail（全站 grep 无 /progress 路由导航 + 浏览器确认）
- [x] ⑦ evidence_ref 全部 / 或 https:// 前缀；web 纯文本/外链双形态；零正文托管（DB 无正文列）
- [x] ⑧ 16 条 source 全 manual（SQL）+ Create 合同与前端不设 source 字段
- [x] ⑨ 无 Update RPC（proto 方法集 + 生成物方法集双侧）
- [x] ⑩ brief recentEvents=10 vs List=16 分层证据
- [x] ⑪ 三重边界：phase11 PhaseEntry 零触碰（grep）/ decision 零关联（SQL）/ plan.md 零复制（DB 无正文列）
- [x] 矩阵 11 行全部 PASS

## 工具链门禁

- [x] `proto/`：make lint / make build / make breaking 退出码全 0
- [x] `backend/`：go build / go vet / go test 退出码全 0（集成测试连库执行非 skip；在 16 条冻结态下跑——测试独立 fixture 不受影响）
- [x] `frontend/`：npx tsc -b 零错误

## 浏览器反回归矩阵

- [x] Repository detail（ca261521）：工作台右列四卡片（已绑定产品/已映射模块/相关决策/项目进度）+ Standard 摘要底部注释位——三轮 UI 返工最终形态（截图）
- [x] 进度区功能点：16 条倒序 + max-h-80 限高滚动 + evidence_ref 纯文本/外链双形态 + 三轨过滤四向 + 当前卡完结态
- [x] Standard 摘要区协排：样本 standards[] 展示正常，与进度区互不影响
- [x] 既有页面抽查 ≥6 页（Dashboard/Standards list/Standards detail/Decisions list/Products list/Onboarding）：零白屏 + 控制台零错误 + 导航无 progress 项
- [x] main-repo detail：进度区空态文案 + 当前卡空态组合正确

## 收口

- [x] acceptance_report.md 十三节齐全（沿 phase14-10 协议），冻结"达标"结论
- [x] 独立复核 PASS（附件 A 与 git log 逐条重新比对 / 证据复查 / 截图抽查 / 无偷渡）
- [x] 无偷渡：零产品代码改动（git status 仅本 spec 目录新增）/ 零根级文档回写 / dogfooding 16 条保持冻结 / main-repo 0 条 / 未重启服务器
- [x] tasks.md / checklist.md 全部勾选附执行记录
- [x] 变更未提交（spec 三件套 + acceptance_report + independent_review），待用户最终确认后手动提交
