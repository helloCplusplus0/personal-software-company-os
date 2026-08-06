- [x] 已明确 `Decision Center` 路由文件、页面文件与根导航入口的实现落点
- [x] 已明确前端实现将使用 `React 19 + TanStack Router(createFileRoute/validateSearch) + TanStack Query(useQuery/useMutation/invalidateQueries) + Zustand(sessionStorage)` 主线
- [x] 已明确 `Decision Center` 数据适配层必须直接消费 `phase03-12` 真实 API，不保留并列 mock 主线
- [x] 已明确 `DecisionListRoute` 的 `queryText / statusFilter` 搜索参数、默认值与列表上下文恢复规则
- [x] 已明确 `DecisionListPage` 的读取态、空状态、错误态与进入创建/详情的最小行为
- [x] 已明确 `DecisionCreateRoute` 承接 `sourceModuleId / sourceModuleName` 搜索参数
- [x] 已明确 `DecisionCreatePage` 的来源上下文展示、草稿提交与成功回流规则
- [x] 已明确 `DecisionDetailPage` 通过 `source_context + linked_modules` 派生待关联目标展示
- [x] 已明确详情页候选读取、目标关联与 mutation 成功后的 reread / invalidateQueries 策略
- [x] 已明确 `ModuleDecisionEntryPanel` 必须升级为两个正式入口动作，并可直接进入详情页
- [x] 已明确 `Decision List / Create / Detail` 在 `PC / 移动浏览器` 下共用单一 `React Web` 页面主线
- [x] 已明确不引入第二套移动端 UI 架构、不引入独立 `React Native` 或完整 `PWA` 前提
- [x] 已明确 `phase03-13` 的最小验收证据至少包含 `npm run build` 与前端闭环走通验证

# 验收证据补充

## 编译验证
- `tsc -b` 通过（无类型错误）
- `vite build` 通过（2162 modules transformed）
- `oxlint` 通过（0 errors，4 warnings 均为预先存在）

## 路由树验证
- `routeTree.gen.ts` 由 `tanstackRouter` Vite 插件自动生成，包含 3 个新 decisions 路由
- `__root.tsx` 导航已增加 Decision Center 入口（`/decisions`）

## 数据适配验证
- `api-adapter.ts` 直接消费 phase03-12 真实 API，5 个函数对齐后端端点
- `decision-center-adapter.ts` 只导出真实 API 实现，不提供 mock-adapter.ts
- `types.ts` 从 .proto 语义派生，承接 source_module_id / linked_modules / source_context

## 前端闭环逻辑验证
- Decision List → Create → Detail → LinkDecisionToTarget 闭环代码路径完整
- Module Detail → Decision Create（带 sourceModuleId / sourceModuleName）→ Detail（source_context 贯通）代码路径完整
- 创建/关联 mutation 成功后均通过 invalidateQueries 驱动 reread，不用手工拼接假数据

## 布局降级验证
- DecisionListPage：PC 表格（md 以上）+ 移动卡片（md 以下）双场景
- DecisionCreatePage：max-w-2xl 单列布局，来源上下文/表单/动作垂直排列
- DecisionDetailPage：PC 三列分区（lg:grid-cols-3）+ 移动垂直重排

## 复核修复验证（GPT-5.4 第二轮独立复核）
- [x] 返回列表上下文“存在/不存在”已通过显式 `fromList` 搜索参数单值化，从 `Module Detail` / 外部直达返回列表落默认参数
- [x] `DecisionCreatePage / DecisionDetailPage` 返回列表不再无条件读取 `lastSearch`，改为 `fromList ? lastSearch : { statusFilter: 'all' }`
- [x] `DecisionListPage / DecisionListContent` 导航到 `Create / Detail` 显式带 `fromList: true`
- [x] `DecisionCreatePage` 创建成功回流 `Detail` 透传 `fromList`，保持列表上下文链路一致
- [x] `ModuleDecisionEntryPanel` 进入 `Create / Detail` 不带 `fromList`，返回列表落默认参数
- [x] 待关联目标结束条件收敛为“仅在正式 `LinkDecisionToTarget` 写入后消失”，删除未实现的“主动放弃关联”承诺
- [x] `spec.md` 两个 Scenario 已显式声明 `fromList` 单值化规则与待关联目标结束条件非目标
- [x] `npm run build` 通过（2162 modules transformed），`oxlint` 0 errors

## 上游规格同步验证（第三轮复核：phase03-10 口径漂移修复）
- [x] phase03-10 `decision_center_spec_v0.1.md` 三处“或主动放弃关联”旧承诺已对齐式更新为“不提供主动放弃出口，仅在正式关联后消失”
- [x] 三件套基线文档（shared_baseline / architecture_plan / dev_plan）全量扫描无“放弃”残留
- [x] `grep "或主动放弃" decision_center_spec_v0.1.md` 返回 0
- [x] phase03-10 spec §9.5「跨页面列表上下文单值承接」与 phase03-13 `fromList` 实现一致
- [x] phase03-10 spec 三处收敛语义与 phase03-13 spec.md L152-153、前端实现注释完全对齐
- [x] P2-1 已收口：phase03-12 spec L13+L216 `source_context` 旧表述同步为收敛语义，`grep` 返回 0
- [x] P2-2 已记录决策：前置子规格 phase03-03/05/06 共 6 处历史残留，依 phase03-10 L828 声明不作为长期入口，当前不修
- [x] 活动规格链（phase03-10 / phase03-12 / phase03-13）结束条件口径已完全对齐
