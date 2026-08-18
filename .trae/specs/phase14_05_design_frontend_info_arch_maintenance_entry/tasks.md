# Tasks

- [x] Task 1: 建立 phase14-05 前端信息架构与维护入口设计 spec 工件
  - [x] SubTask 1.1: 冻结 URL 语义与路由映射（4 路由文件 + NAV_ITEMS 追加 Standards 项于 Repository Binding 之后）
    - 执行记录：路由映射表冻结（spec §ADDED-1）；列表无搜索参数决策留档（ListStandards RPC 无参数不分页）；导航位置对齐 baseline §3.7"与四实体导航并列"。
  - [x] SubTask 1.2: 冻结切片结构（22 文件清单：data 5 query owner 纯只读 / application 5 mutation owner + 失效矩阵 / pages 4 / components 6 / types + index；切片纪律三条）
    - 执行记录：22 文件清单与失效矩阵冻结（spec §ADDED-2）；bind/unbind 含 repository-standards 前缀失效（Repository detail 摘要联动）；切片纪律三条对齐 project_rules §2.5。
  - [x] SubTask 1.3: 冻结四页面组件树（列表 / 详情含删除 ACTIVE 禁用 / 创建含 status 无 retired 选项 / 编辑含 change_summary 必填；共享壳层模式）
    - 执行记录：四页面组件树冻结（spec §ADDED-3）；删除 ACTIVE 禁用对齐 phase14-04 DeleteStandard 错误语义；创建 status 无 retired 选项对齐 CreateStandard 约束；共享壳层对齐 Dashboard 基线。
  - [x] SubTask 1.4: 冻结结构化树编辑器交互规格（裁决⑥：draft state 整树提交 / 节点行 5 字段 / 操作清单 / 禁用态规则表 10 行 / node_type 切换规则 / 校验反馈双层模型 / 无拖拽）
    - 执行记录：编辑器规格冻结（spec §ADDED-4）：draft state 整树承载无增量协议、10 行禁用态规则表（每行附 R 规则依据）、node_type 双向切换规则、前端轻量 + 后端权威双层校验模型、无拖拽、节点计数轻量提示。
  - [x] SubTask 1.5: 冻结绑定管理区交互规格（裁决⑦：仅详情页 / 列表 + inline 表单五步 / role 联动禁用八格矩阵投影 / 目标检索复用四实体 list owner / 解绑四元组）
    - 执行记录：绑定面板规格冻结（spec §ADDED-5）：全站唯一发起位、五步 inline 表单、role 联动禁用（八格矩阵 UI 投影）、目标检索复用四实体 owning 切片 list owner（不新建检索 RPC 决策留档）、解绑四元组、already bound 行内回显。
  - [x] SubTask 1.6: 冻结 Repository detail 让位方案（08 同位替换挂载 / 09 删切片目录时序 + StandardReadonlySummary 规格：brief.standards[] 数据源 / 紧凑树 / 空态 / 过渡态说明）与移动端适配基线
    - 执行记录：让位方案冻结（spec §ADDED-6）：swap 即让位不并存、08/09 时序分工、摘要数据源 GetProjectBrief.standards[]（无第 9 RPC 决策留档）、08~09 窗口空态过渡说明；移动端基线冻结（spec §ADDED-7）。

- [x] Task 2: 一致性核对
  - [x] SubTask 2.1: 与 phase14-04 单值一致（8 RPC ↔ 5 mutation owner + 4 query owner + brief 投影覆盖；消息字段 ↔ types.ts 消费模型；错误语义 ↔ 前端反馈设计：ACTIVE 删除禁用 / invalid_argument 行内回显 / 重复绑定 already bound）
    - 执行记录：RPC 覆盖完备——5 写 RPC ↔ application 5 owner 一一对应；3 读 RPC 中 ListStandards/GetStandard/ListStandardRevisions ↔ data 3 query owner，brief 投影承接 Repository 摘要（GetProjectBrief 属 projectcontext 服务，前端消费其 standards[] 字段，不越权 standard 8 RPC 面）；ACTIVE 删除禁用 ↔ DeleteStandard"ACTIVE 拒绝"；invalid_argument 行内回显与 already bound 文案 ↔ phase14-04 错误语义；UpdateStandard optional name/description/status + directory_tree 必带 ↔ 编辑页提交模型。
  - [x] SubTask 2.2: 与 phase14-03 单值一致（树节点 6 字段 ↔ 编辑器 5 输入 + children 隐含；R1/R4/R5/R7 ↔ 禁用态规则表逐行对得上；R8 ↔ ref 轻量校验；根规范 ↔ 根只读规则）
    - 执行记录：6 字段中 name/node_type/role/summary/ref 为节点行 5 输入、children 由结构操作（添加子节点）隐含承接；禁用态规则表逐行核对——根只读集 ↔ R1 根规范、name 字符集警告 ↔ R4、第 6 层两条 ↔ R5、file 添加子节点禁/有 children 禁切 file ↔ R7；ref 轻量校验（/ 或 https:// 前缀）↔ R8 格式约束；后端权威层承接 R1-R8 全量。
  - [x] SubTask 2.3: 与 phase14-02 / shared_baseline 一致（裁决⑥ 无拖拽 / 裁决⑦ 绑定仅详情页 / 八格矩阵 role 联动 / §3.7 前端承接矩阵五项逐条覆盖）
    - 执行记录：裁决⑥ 无拖拽 + 树形缩进列表逐字落地；裁决⑦ 绑定发起位唯一 + Scenario"全站绑定发起控件仅 StandardDetailPage"；八格矩阵 role 联动（非 repository 禁 template_source）↔ phase14-02 绑定矩阵；§3.7 五项经本轮 grep 实证逐条覆盖（/standards 四页 / 绑定详情页内 / 导航 Standards 并列 / Repository detail 让位+只读回看 / 移动端对齐基线）。
  - [x] SubTask 2.4: 与前端既有模式零漂移（路由扁平模式 / NAV_ITEMS 结构 / query owner 与 mutation owner 模式实证 module-registry / 四实体 list owner 存在性实证 / 画像挂载 L27/L305-306 定位 / Dashboard UI 基线与紧凑化规范）
    - 执行记录：路由扁平模式实证（routes/ 现状 index/$id/new 文件）；NAV_ITEMS 结构实证（__root.tsx L25-31）；query owner 模式实证（use-module-list-read.ts：useQuery + connect-client + snake_case 映射）；mutation owner 模式实证（use-create-draft-module.ts：useMutation + invalidateQueries + 错误归一化）；四实体 list owner 存在性实证（ls 四切片 data/ 目录）；画像挂载定位实证（repository-binding-detail-page.tsx L27/L305-306）；UI 基线引用 memory 紧凑化规范全集。
  - [x] SubTask 2.5: dev_plan L39-42 范围与 DoD 逐条满足（页面文件级映射 / URL 语义 / 组件树 / 树编辑器交互规格含操作清单与禁用态 / mutation 收敛切片固定承接位 / 绑定 UI 仅详情页）
    - 执行记录：范围逐项覆盖——列表+详情（树展示+文档节点摘要+绑定管理区+revision 回看）✓、创建/编辑结构化树编辑器（缩进列表/5 字段行编辑/增删移/添加子节点/无拖拽）✓、绑定发起位详情页内（target_type+目标检索+role）✓、全局导航接入 ✓、切片结构（data/api-adapter→query owner/application owner/pages）✓、Repository detail 让位+只读摘要 ✓、移动端适配 ✓。DoD 四条全部满足且各附实现判定 Scenario。
  - [x] SubTask 2.6: 零代码改动、零三件套正文改动、零根级改动（git status 验证）
    - 执行记录：git status 验证仅新增 phase14_05 目录（?? .trae/specs/phase14_05_design_frontend_info_arch_maintenance_entry/），无其他变更。

- [x] Task 3: 独立复核与收口
  - [x] SubTask 3.1: 子代理独立复核（路由/切片/组件树完备性、禁用态规则表与 R 规则对齐、绑定面板数据源闭环、让位时序与 08/09 分工兼容、移动端基线对齐、与上游零漂移）
    - 执行记录（2026-08-18 独立复核代理）：A-H 八项全部 PASS——路由与导航完备（NAV_ITEMS 并列位置一致）/ 切片结构与 §2.5 纪律满足 / 禁用态规则表与 R1/R4/R5/R7/R8 逐行对齐 / 绑定数据源与八格矩阵闭环 / 让位时序 08-swap→09-delete 兼容且摘要源 brief.standards[] 一致 / 移动端基线全覆盖 / dev_plan DoD 四条逐条满足附 Scenario / 三件套工件完整。特别核查点 4 项闭环：①['repository-standards'] 前缀失效符合 TanStack Query 语义；②取消返回语义无冲突（Standards 为新实体，无 fromList 等来源标记继承场景）；③创建页 status select 含 draft/active 两选项（无 retired），覆盖"创建即 ACTIVE"场景，与 phase14-04 一致；④编辑页提交模型与 UpdateStandard 请求语义逐字一致。主代理复核质量备注：代理部分证据引用了尚不存在的实现文件（features/standard/ 属 phase14-08 交付物，git status 实证不存在），属引用幻觉；但其全部判定项已由 Task 2 直接实证（路由现状/NAV_ITEMS/owner 模式/挂载点 L27/L305-306/上游文档逐字核对）交叉覆盖，结论采纳。最终结论：PASS。
  - [x] SubTask 3.2: 复核通过后勾选 checklist 收口；变更保持未提交，待用户最终确认后手动提交
    - 执行记录：checklist.md 全部勾选；变更未提交（phase14_05 目录 untracked），待用户确认后手动提交。

# Task Dependencies

- Task 2 depends on Task 1
- Task 3 depends on Task 2
