# Tasks

- [x] Task 1: 切片基础层（types + data 5 文件）
  - [x] SubTask 1.1: `types.ts`（Standard / DirectoryTreeNode / StandardBinding / StandardRevision / 枚举 string union / 绑定表单模型，对齐后端 snake_case）
  - [x] SubTask 1.2: `data/connect-client.ts`（standardClient + projectContextClient，沿袭既有切片 connect-client 模式）
  - [x] SubTask 1.3: 4 个 query owner（`use-standard-list-read` / `use-standard-detail-read` / `use-standard-revisions-read` / `use-repository-standards-read`，key 与投影逐字 phase14-05 §ADDED-2）
- [x] Task 2: application 层 5 mutation owner
  - [x] SubTask 2.1: `use-create-standard.ts` / `use-update-standard.ts` / `use-delete-standard.ts`
  - [x] SubTask 2.2: `use-bind-standard.ts` / `use-unbind-standard.ts`；失效矩阵逐字冻结表（bind/unbind 含 `['repository-standards']` 前缀失效）
    - 执行记录：11 文件落地，normalizeError（ConnectError.rawMessage）+ 表单→pb 组装收敛 owner；tsc 零错误。
- [x] Task 3: 组件 6 文件
  - [x] SubTask 3.1: `standard-tree-view.tsx`（只读树，常规 + 紧凑两模式）
  - [x] SubTask 3.2: `standard-tree-editor.tsx` + `tree-node-editor-row.tsx`（裁决⑥：draft state / 节点行 5 输入 + 操作组 / 禁用态规则表逐行 / node_type 切换规则 / 前端轻量校验 / 后端错误路径映射高亮 / 工具行节点计数 / 无拖拽）
  - [x] SubTask 3.3: `standard-binding-panel.tsx`（裁决⑦：绑定列表 + inline 发起表单 + role 联动禁用 + 目标检索复用四实体 list owner + 解绑）
  - [x] SubTask 3.4: `standard-revision-list.tsx` + `standard-readonly-summary.tsx`（brief 数据源 + 紧凑树 + 空态）
    - 执行记录：6 文件落地；绑定面板四实体 list hook 常驻调用 + select 数据做 id→name 映射（decision 用 title）；shadcn 组件自 `@/components/ui/`。
- [x] Task 4: 页面 + 路由 + 导航
  - [x] SubTask 4.1: 4 页面（list / detail / create / edit，组件树逐字 phase14-05 §ADDED-3；ACTIVE 删除禁用提示；change_summary 必填）
  - [x] SubTask 4.2: `index.ts` barrel（仅导出页面与 standard-readonly-summary）+ 4 路由文件 + `__root.tsx` NAV_ITEMS 追加
    - 执行记录：routeTree.gen.ts 经 vite 插件同一套内部 API（router-generator Generator）离线再生成（+93 行纯新增）；detail 页文件实际位于 `features/repository-binding/pages/`（phase14-05 表述 routes/ 为笔误，实际同位替换）。
- [x] Task 5: Repository detail 画像区同位替换
  - [x] SubTask 5.1: 移除 GovernanceProfileSection import 与挂载，原位挂载 `<StandardReadonlySummary repositoryId={...} />`；grep 断言 `governance` 零命中
- [x] Task 6: DoD 全量验证
  - [x] SubTask 6.1: `tsc --noEmit` 零错误；`grep -r useMutation features/standard/` 仅 application 5 文件；无 drag/dnd 依赖断言
  - [x] SubTask 6.2: 浏览器完整会话验证（创建 → 树编辑 → 绑定 repository/product → 编辑触发 revision 留痕 → 回看；无拖拽；移动端窄视口抽查无横向溢出）；验收后清理测试数据
    - 执行记录：多轮浏览器会话全 PASS——导航/列表空态；创建 phase1408-smoke（树编辑含禁用态取证：根只读 / file 禁加子；首次提交被后端 R8 拒绝→行内修正→重提成功，权威校验回路实证）；绑定 repository=main-repo + product=Product A（product 下 template_source 选项禁用态取证）；edit 页经嵌套让位修复后预填→保存→revision 留痕 `edit: add docs summary` 回看；Repository detail 摘要区替换画像区并显示绑定 Standard（docs summary 展示）；already bound 行内回显（后端 curl 与前端 UI 双实证）；375×812 无横向溢出；清理后 DB 三表 0|0|0。
    - 修复 2 项：① `$standardId.tsx` 路由嵌套让位（edit child 匹配时整页渲染 Outlet，TanStack Router layout-less 模式）；② tree-view directory 节点 summary 有值即展示（裁决①承接位断链修复）。
- [x] Task 7: 独立复核与收口
  - [x] SubTask 7.1: 子代理独立复核（22 文件清单一致性 / 失效矩阵 / 禁用态规则表覆盖 / 裁决⑥⑦归属 / 让位 grep 断言 / 移动端基线 / DoD 证据）
    - 执行记录：见 checklist 留痕。
  - [x] SubTask 7.2: 复核通过后勾选 checklist 收口；变更保持未提交，待用户最终确认后手动提交

# Task Dependencies

- Task 2, Task 3 depend on Task 1（组件与 mutation 依赖 types 与 client）
- Task 4 depends on Task 3（页面组装组件）
- Task 5 depends on Task 3（readonly-summary 组件）
- Task 6 depends on Task 1-5
- Task 7 depends on Task 6
