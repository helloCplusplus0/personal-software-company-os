# phase14-08 Checklist

## 路由与导航

- [x] 4 路由文件（`standards/index.tsx` / `standards/new.tsx` / `standards/$standardId.tsx` / `standards/$standardId.edit.tsx`）URL 语义逐字 phase14-05 §ADDED-1
- [x] `__root.tsx` NAV_ITEMS 追加 Standards 项（Repository Binding 之后）；PC 常驻与移动端菜单同步出现
  - 留痕：浏览器会话 PASS（导航项可见并点击进入）；`$standardId.tsx` 增加嵌套让位（edit child 匹配时整页渲染 Outlet，layout-less 模式）——修复前 edit URL 渲染 detail 内容。

## 切片结构与纪律

- [x] 22 文件清单与 phase14-05 §ADDED-2 一一对应（data 5 / application 5 / pages 4 / components 6 / types / index；无多余无缺失）
- [x] data 层 4 query owner 的 query key 逐字（`['standard-list']` / `['standard-detail', id]` / `['standard-revisions', id]` / `['repository-standards', repositoryId]`）
- [x] `use-repository-standards-read` 数据源为 GetProjectBrief 投影 `standards[]`（与 agent 消费同源，无第 9 RPC）
- [x] `grep -r useMutation frontend/src/features/standard/` 仅命中 application 5 文件
- [x] application 5 owner 失效矩阵逐字冻结表（bind/unbind 含 `['repository-standards']` 前缀失效）
- [x] types.ts 对齐后端 snake_case；枚举为 string union
- [x] index.ts barrel 仅导出页面与 `standard-readonly-summary`；`standard-tree-view` / `tree-node-editor-row` 未晋升 shared

## 四页面

- [x] StandardListPage 组件树逐层一致（页壳 / 错误 / 加载 / 空态 / 摘要卡整卡 Link）
- [x] StandardDetailPage：头部操作组（编辑 / 删除确认弹窗 / ACTIVE 删除禁用提示先 Retire）+ TreeView + BindingPanel + RevisionList（`border-t pt-2`）
- [x] StandardCreatePage：status select 默认 draft 不含 retired；初始树单根 directory `name="."`；提交成功 → 详情页；取消 → `/standards`
- [x] StandardEditPage：预填 + change_summary 必填 + 整树预填；保存成功 → 详情页；取消 → 详情页
- [x] 共享壳层模式（`text-xl` 标题 / `text-xs` 导语 / `border-t` 分隔 / 无 Card 重型嵌套 / 容器仅 `:focus-visible`）
  - 留痕：浏览器会话 PASS——创建提交→详情导航、edit 预填（name/树全量）→保存→回流、取消路径均实测通过。

## 树编辑器（裁决⑥）

- [x] 整树 draft state 本地持有；不发起网络请求；提交整树单次调用
- [x] 节点行 5 输入 + 操作组（`h-7 px-2 text-xs variant=outline`）；层级缩进每层 `pl-4`；移动端字段网格折行
- [x] 禁用态规则表逐行落地（根只读与操作禁用 / file 添加子节点禁用 / directory 含 children 切 file 禁用 / 第 6 层 directory 禁用 / 同层首末移动禁用 / 删 directory 确认弹窗 / name 非法字符行内警告）
- [x] node_type 切换规则（file→directory 置空 children；directory→file 仅无 children）
- [x] 前端轻量校验不阻断输入；后端 R1-R8 错误节点路径映射回节点行高亮
- [x] 工具行：添加根级节点 + 节点计数
- [x] 无拖拽交互（无 drag 事件 / 无 dnd 依赖）
  - 留痕：浏览器取证——根行 name/node_type 只读、根删除/上移/下移禁用、file 行"添加子节点"禁用；首次提交被后端 R8 拒绝（`must equal node path`）→ 修正后重提成功，双层校验回路实证。复核 OBS-1 注释关键字已改写，drag grep 零命中。

## 绑定管理区（裁决⑦）

- [x] 绑定发起控件全站仅 StandardDetailPage 一处
- [x] 发起表单：target_type select → 目标检索复用四实体 owning 切片 list owner（全量 + 前端过滤）→ role 联动禁用（非 repository 时 template_source 禁用）→ note 可选
- [x] 绑定列表 target 名称经 owning 切片缓存解析（未命中显示 id 短版）；解绑走四元组
- [x] `invalid_argument`（含 "already bound"）行内错误回显
  - 留痕：浏览器会话 PASS——repository=main-repo + product=Product A 绑定成功（product 下 template_source 选项禁用态截图取证）；already bound 重复提交行内回显原文（后端 curl 与前端 UI 双实证）。

## Repository detail 让位

- [x] GovernanceProfileSection import 与挂载移除，原位挂载 `StandardReadonlySummary`
- [x] `grep governance repository-binding-detail-page.tsx` 零命中
- [x] 摘要：name + status Badge + 链接 `/standards/:id` + 紧凑树；空态文案；brief 数据源
  - 留痕：浏览器会话 PASS——main-repo 详情页显示"关联 Standard"区（含 phase1408-smoke 条目 + 紧凑树 + docs summary 展示）。让位目标文件实际位于 `features/repository-binding/pages/`（phase14-05 中 routes/ 为路径笔误）。修复：tree-view directory 节点 summary 有值即展示（裁决①承接位）。OBS-2 已改类型化 Link。

## 移动端基线

- [x] 页头响应式折行；主按钮 `w-full sm:w-auto`
- [x] 列表密度 / hover 限可交互元素 / 树缩进不收缩 / `min-w-0 truncate` 防溢出
- [x] 绑定面板字段 `grid-cols-1 sm:grid-cols-2`；列表行 `flex-wrap`
  - 留痕：375×812 视口实测 `/standards` 与详情页无横向溢出、新建按钮全宽、树节点行折行（step-mobile.png）。

## DoD 门禁

- [x] `tsc --noEmit` 零错误
- [x] 浏览器完整会话：创建 → 结构化树编辑 → 从详情页绑定 repository 与 product → revision 留痕 → 回看，全链可完成
- [x] 会话验证后测试数据已清理
  - 留痕：清理后 DB 三表 count = 0|0|0（standards / standard_bindings / standard_revisions）。

## 复核与收口

- [x] 子代理独立复核通过（发现项已修复回填）
  - 留痕：独立复核 12 项全 PASS、0 阻断、4 观察项。OBS-1（注释关键字假阳性）与 OBS-2（模板字符串 Link）已修复；OBS-3（detail 树区无 border-t，规格仅要求 RevisionList）与 OBS-4（面板独立请求，设计内行为）不阻断留痕。
- [x] tasks.md 全部勾选并附执行记录
- [x] 变更未提交，待用户最终确认后手动提交
