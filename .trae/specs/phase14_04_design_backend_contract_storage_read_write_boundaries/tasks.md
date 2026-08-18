# Tasks

- [x] Task 1: 建立 phase14-04 后端合同、存储与读写边界设计 spec 工件
  - [x] SubTask 1.1: 冻结合同包结构（package / go_package / 4 枚举值序 / 4 核心消息与字段号；单一 Standard 消息决策留档；文件头四段注释风格对齐 governance_profile.proto）
    - 执行记录：合同包 Requirement 冻结（spec §ADDED-1）：syntax/package/go_package/import/4 枚举（含 proto 代码草案）/4 消息（含字段号与注释）；单一 Standard 消息决策（无 StandardSummary）附三点理由留档。
  - [x] SubTask 1.2: 冻结 8 个 RPC 三要素（CreateStandard / ListStandards 不分页 / GetStandard 含 bindings / UpdateStandard 整树原子替换 + change_summary 必填 + 单事务 / DeleteStandard active 拒绝 / BindStandard 五步校验链 / UnbindStandard 四元组定位 / ListStandardRevisions DESC 不分页；每个含请求/响应/错误语义）
    - 执行记录：8 RPC 三要素表冻结（spec §ADDED-2）+ 五条语义要点（Create 不记 revision / Update 无部分更新 / Create status 约束 / 不分页反假大空 / 重复绑定归类 InvalidArgument）；附两个 Scenario（RPC 实现判定 + 绑定校验链判定）。
  - [x] SubTask 1.3: 冻结 Go 模块分层（backend/internal/standard/ 8 文件职责单值映射 + 分层约束：connect 无 SQL / repository 无业务 / service 跨模块经 candidate）
    - 执行记录：8 文件职责表冻结（spec §ADDED-3）：connect/server.go、service/command_service.go、service/query_service.go、repository/standard_store.go、candidate/target_reader.go、errors.go、types.go、validate.go；三条分层约束 + router.go 装配顺序。
  - [x] SubTask 1.4: 冻结写路径归属矩阵（5 写 RPC web 唯一入口 + 3 读 RPC web/agent 共享；语义边界非技术强制说明）
    - 执行记录：写路径归属矩阵冻结（spec §ADDED-4）：5 写 RPC 仅 web（/standards 切片，裁决⑦）、agent 无写回（CON-09）；3 读 RPC 双方共享；语义边界非技术强制 + agent 消费主路径（brief）+ 直读补充场景说明。
  - [x] SubTask 1.5: 冻结 StandardReader 跨模块接口签名（消费方拥有模式 / 落点 context_readers.go / 实现位 query_service.go / 注入位 router.go / 空列表非错误语义）与 brief 装配演进点（phase14-07 新增接入 vs phase14-09 退役切换分工表 + 跨包引用活跃合同源的设计决策）
    - 执行记录：StandardReader 接口 Go 草案冻结（spec §ADDED-5，含完整接口文档注释）+ 三条装配约束；brief 两步分工表冻结（spec §ADDED-6：phase14-07 新增接入不动现状 / phase14-09 退役切换 reserved 3 + 内联化）+ 跨包引用设计决策留档；附三个 Scenario。

- [x] Task 2: 一致性核对
  - [x] SubTask 2.1: 与 phase14-03 单值一致（proto 对齐表字段名/值序逐字一致：NODE_TYPE_DIRECTORY=1/NODE_TYPE_FILE=2；jsonb 规范 6 字段 ↔ proto 6 字段；8 条校验规则在 Create/Update 错误语义中正确引用；DDL 三表列 ↔ 消息投影完备）
    - 执行记录：DirectoryTreeNode proto 6 字段与 phase14-03 对齐表逐字一致（含值序与 file children 约束注释）；Standard 消息 7 字段 ↔ standards 表列完备（id/name/description/status/directory_tree/created_at/updated_at）；StandardBinding 7 字段 ↔ bindings 表列 + id；StandardRevision 4 字段 ↔ revisions 表列 + id；R1-R8 在 Create（树校验失败）/ Update（树校验失败 + R2 状态时机）错误语义中正确引用。
  - [x] SubTask 2.2: 与 phase14-02 单值一致（八格矩阵校验链与 BindStandard 错误语义一致；brief 对照表 8 字段 ↔ 步骤 1/2 分工覆盖；StandardSummary 占位名单值化收敛理由闭环；跨包引用决策与 T6 性质区分成立）
    - 执行记录：BindStandard 校验链含八格矩阵判定（template_source 携非 repository → InvalidArgument）与 target 存在性（InvalidArgument），与 phase14-02 逐字一致；brief 两步分工表覆盖对照表全部演进字段（standards=8 新增于步骤 1；global_assets=3 BREAKING+reserved 与 governance_profile=2 内联化于步骤 2；1/4/5-7 不动）；StandardSummary 收敛在合同包 Requirement 显式声明；跨包引用活跃合同源与 T6"退役包依赖解除"性质区分在演进点 Requirement 留档。
  - [x] SubTask 2.3: 与 shared_baseline §2.3/§2.4/§3.2-3.4 一致（技术主线 ConnectRPC + chi 基础设施；整树原子替换 + revision 追加；绑定矩阵枚举与唯一约束；反假大空：不分页决策留档）
    - 执行记录：ConnectRPC 正式传输（connect/server.go handler + router 挂载）对齐 §2.3；UpdateStandard 整树替换 + revision 追加单事务对齐 §2.4"树整体原子替换"；枚举值域与 §3.4 绑定矩阵一致；ListStandards/ListStandardRevisions 不分页以 CON-07 反假大空为依据留档。
  - [x] SubTask 2.4: 与仓库既有模式零漂移（错误三类 code 对齐 governanceprofile errors.go 哨兵模式；candidate 模式对齐 phase13-10 GovernanceProfileReader 冻结口径；模块文件清单对齐工程约定单值化映射）
    - 执行记录：errors.go 哨兵清单（ErrStandardNotFound/ErrBindingNotFound → NotFound；ErrInvalidInput → InvalidArgument；ErrStandardReadFailed/ErrStandardSaveFailed → Internal）对齐 governanceprofile errors.go 结构与命名；StandardReader 接口文档结构逐字仿照 context_readers.go L26-36 GovernanceProfileReader 冻结口径；router.go 装配点经本轮 grep 实证（L405-407 buildProjectContext 接口注入模式），StandardReader 注入设计与现状一致；8 文件清单对齐工程约定支撑文件单值化映射。
  - [x] SubTask 2.5: dev_plan L34-37 范围与 DoD 逐条满足（8 RPC 三要素 ✓ / 合同对齐 .proto 主线 ✓ / canonical owner 单值化 ✓ / 跨模块经独立 Read 接口沿袭 candidate 模式 ✓）
    - 执行记录：范围五件全覆盖——psco.standard.v1 合同（消息 + 8 RPC 三要素，dev_plan 点名的 ListStandards 不分页 / UpdateStandard 整树替换 + change_summary 必填 / BindStandard target 不存在错误语义 / ListStandardRevisions 回看读取位四项显式落实）；Go 模块分层（connect/service/repository + candidate）；写路径归属（web 写 agent 只读矩阵）；StandardReader 签名（输入 repository_id、输出含整树列表）；brief 装配演进点（两步分工）。DoD 四条逐条满足。
  - [x] SubTask 2.6: 零代码改动、零三件套正文改动、零根级改动（git status 验证）
    - 执行记录：git status 验证本子任务仅新增 phase14_04 目录（phase14_03 已由用户提交），无其他变更。

- [x] Task 3: 独立复核与收口
  - [x] SubTask 3.1: 子代理独立复核（消息/枚举完备性与 phase14-03 对齐、8 RPC 错误语义无重叠遗漏、校验链顺序自洽、分层约束可执行、brief 分工与 phase14-02 对照表闭环、与上游零漂移）
    - 执行记录（2026-08-18 独立复核代理）：复核证据链 18 处全部为确认性判定，无 FAIL 项——DirectoryTreeNode 6 字段与 phase14-03 jsonb 规范逐字一致（spec L75-82 ↔ 03 L96-104）；4 枚举值序含 NODE_TYPE_DIRECTORY=1/FILE=2 对齐（L36-69）；8 RPC 错误三类互斥与重复绑定归类 InvalidArgument（L130-139）；Update 无部分更新 / Create 不记 revision / status 约束（L142-145）；BindStandard 五步校验链顺序固定（L159-160）与 phase14-02 八格矩阵一致（02 L33-50）；8 文件分层与 governanceprofile 现状模式一致（L165-172）；StandardReader 沿袭 GovernanceProfileReader 消费方拥有模式与空列表非错误（L207-220 ↔ context_readers.go L26-36）；哨兵错误三类对齐（errors.go L9-31）；brief 两步分工与 phase14-02 对照表完全覆盖且不改 1/4/5-7（L235-240 ↔ 02 L107-125）；dev_plan DoD 四条相符（L34-37）；写路径语义边界非技术强制（L184-196）；Create 不记 revision 与裁决⑧"迁移脚本不经 RPC"说明一致（L140-141）。特别核查点 1（GetStandard 响应含 bindings 免第 9 RPC）经 8 RPC 表内闭环确认。最终结论：PASS。
  - [x] SubTask 3.2: 复核通过后勾选 checklist 收口；变更保持未提交，待用户最终确认后手动提交
    - 执行记录：checklist.md 全部勾选；变更未提交（phase14_04 目录 untracked），待用户确认后手动提交。

# Task Dependencies

- Task 2 depends on Task 1
- Task 3 depends on Task 2
