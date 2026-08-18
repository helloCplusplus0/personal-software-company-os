# phase14-05 Checklist

## URL 与路由冻结

- [x] 4 路由文件映射冻结（standards/index.tsx、new.tsx、$standardId.tsx、$standardId.edit.tsx；列表无搜索参数决策留档）
- [x] NAV_ITEMS 追加 `{ to: '/standards', label: 'Standards' }` 于 Repository Binding 之后（PC 与移动端菜单同步）

## 切片结构冻结

- [x] 22 文件清单职责单值映射（data 5 / application 5 / pages 4 / components 6 / types / index）
- [x] data 层纯只读（5 query owner；query key 冻结：standard-list / standard-detail:id / standard-revisions:id / repository-standards:repositoryId）
- [x] application 层 5 mutation owner + 失效矩阵逐字冻结（bind/unbind 含 repository-standards 前缀失效）
- [x] 切片纪律三条（query 零写动作 / useMutation 仅 application 5 文件 / tree-view 与 node-row 不晋升 shared）

## 四页面组件树冻结

- [x] 列表页（标题行 / 错误区 / 空态 / 摘要卡列表整卡 Link）
- [x] 详情页（头部 + 删除 ACTIVE 禁用提示先 Retire + TreeView + BindingPanel + RevisionList border-t 分隔）
- [x] 创建页（status select 默认 draft 且含 active 选项无 retired / 初始单根空树 / 取消回列表）
- [x] 编辑页（预填 + change_summary 必填 / 整树保存 / 取消回详情）
- [x] 共享壳层模式（text-xl 标题 / text-xs 导语 / border-t / :focus-visible）

## 树编辑器交互规格冻结（裁决⑥）

- [x] draft state 本地整树承载 + 提交整树单次调用（无节点级增量协议）
- [x] 节点行 5 输入 + 4 操作按钮 + 移动端折行网格
- [x] 禁用态规则表 10 行逐行冻结（根只读集 / 根删移禁 / file 添加子节点禁 / 有 children 禁切 file / 第 6 层添加与 directory 选项禁 / 首末位移禁 / 删目录确认弹窗 / name 字符集警告）
- [x] node_type 切换规则（file→directory 直接 / directory→file 须无 children）
- [x] 校验反馈双层模型（前端轻量即时提示不阻断 + 后端权威 R1-R8 错误路径映射回节点行高亮）
- [x] 无拖拽 + 节点计数轻量提示（不做字节计数）

## 绑定管理区冻结（裁决⑦）

- [x] 仅 StandardDetailPage 承接（全站唯一绑定发起位）
- [x] 现有绑定列表（target_type / target 名称经 list owner 缓存映射 / role / note / created_at / 解绑）
- [x] inline 表单五步（target_type → 目标检索复用四实体 list owner 全量+前端过滤 → role 联动禁用（非 repository 禁 template_source）→ note → 提交）
- [x] invalid_argument（含 already bound）行内回显；解绑四元组确认

## Repository detail 让位方案冻结

- [x] 让位时序单值（08 同位替换 L27/L305-306 / 09 删切片目录；不并存双治理区）
- [x] StandardReadonlySummary 数据源 = GetProjectBrief.standards[]（无第 9 RPC 决策留档；与 agent 主路径同源）
- [x] 紧凑树形只读 + 链接详情 + 空态文案 + 过渡态空态说明（08~09 窗口设计内）
- [x] 实现判定 Scenario（grep governance 零命中 / tsc 零错误）

## 移动端适配基线

- [x] 页头响应式 / 按钮 w-full sm:w-auto / p-3 space-y-2 密度 / 树缩进折行 / min-w-0 truncate / 字段标签 text-xs / hover 仅可交互元素

## 一致性核对

- [x] 与 phase14-04 单值一致（RPC ↔ owner 覆盖 / 消息 ↔ types / 错误语义 ↔ 反馈设计）
- [x] 与 phase14-03 单值一致（6 字段 ↔ 5 输入+children / R1/R4/R5/R7/R8 ↔ 禁用与校验规则 / 根规范 ↔ 只读集）
- [x] 与 phase14-02 / shared_baseline §2.2 裁决⑥⑦ / §3.7 承接矩阵五项逐条覆盖（grep 实证）
- [x] 与前端既有模式零漂移（路由 / NAV_ITEMS L25-31 / owner 模式实证 / 四实体 list owner 存在性 / 挂载点 L27/L305-306 / Dashboard 基线）
- [x] dev_plan L39-42 范围与 DoD 逐条满足（四条均附实现判定 Scenario）
- [x] 零代码改动、零三件套正文改动、零根级改动（git status 验证仅新增 phase14_05 目录）

## 复核与收口

- [x] 子代理独立复核通过（A-H 全 PASS、特别核查点 4 项闭环；代理引用幻觉已识别并由 Task 2 直接实证交叉覆盖，结论采纳）
- [x] tasks.md 全部勾选并附执行记录
- [x] 变更未提交，待用户最终确认后手动提交
