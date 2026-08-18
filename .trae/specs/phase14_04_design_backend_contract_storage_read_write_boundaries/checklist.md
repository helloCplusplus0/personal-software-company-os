# phase14-04 Checklist

## 合同包结构冻结

- [x] package `psco.standard.v1` / go_package 落点 / import timestamp 与 governance_profile.proto 合同风格逐字对齐（文件头四段注释）
- [x] 4 枚举值序冻结（StandardStatus / NodeType / BindingTargetType / BindingRole，全部含 _UNSPECIFIED=0；NodeType 值序与 phase14-03 proto 对齐表一致）
- [x] 4 核心消息冻结（DirectoryTreeNode 6 字段 ↔ phase14-03 jsonb 规范逐字一致；Standard / StandardBinding / StandardRevision 与 DDL 三表投影完备）
- [x] 单一 Standard 消息决策留档（不另造 StandardSummary；理由：轻量全量投影 + 数量级可控 + 避免第二套字段语义；对 phase14-02 占位名的收敛显式声明）

## 8 个 RPC 三要素冻结

- [x] CreateStandard：status 默认 DRAFT / 拒绝 RETIRED / ACTIVE 需树满足 R2 / 不记 revision（决策留档）
- [x] ListStandards：无参数不分页 / updated_at DESC / CON-07 反假大空决策留档
- [x] GetStandard：响应含 bindings（绑定管理区直接消费，无第 9 RPC）
- [x] UpdateStandard：directory_tree 必带整树原子替换 / change_summary 必填 / optional name/description/status / 单事务（树替换 + updated_at + revision 追加）不写半套状态
- [x] DeleteStandard：ACTIVE 拒绝（先 retire 防误删）/ CASCADE 连带 bindings 与 revisions
- [x] BindStandard：五步校验链固定顺序（standard 存在 → 枚举 → 八格矩阵 → target 存在 → 唯一约束）；target 不存在 → InvalidArgument（phase14-02 冻结）
- [x] UnbindStandard：四元组定位（note 不参与）/ 不存在 → NotFound
- [x] ListStandardRevisions：created_at DESC 不分页
- [x] 错误 code 三类归一（NotFound / InvalidArgument / Internal），重复绑定归类 InvalidArgument 决策留档（不引入 AlreadyExists 第四类）

## Go 模块分层冻结

- [x] 8 文件清单职责单值映射（connect/server.go / service/command_service.go / service/query_service.go / repository/standard_store.go / candidate/target_reader.go / errors.go / types.go / validate.go）
- [x] 分层约束冻结：connect 无 SQL / repository 无业务判断 / service 跨模块一律经 candidate
- [x] errors.go 哨兵三类对齐 governanceprofile 模式（ErrStandardNotFound / ErrBindingNotFound / ErrInvalidInput / ErrStandardReadFailed / ErrStandardSaveFailed）

## 写路径归属冻结

- [x] 5 写 RPC web 唯一入口（/standards 切片，裁决⑦）/ agent 无写回承接位（CON-09）
- [x] 3 读 RPC web/agent 共享；agent 主路径 GetProjectBrief.standards[] + 直读补充场景说明
- [x] 语义边界非技术强制说明（当前无鉴权体系）

## StandardReader 与 brief 演进点冻结

- [x] 接口签名冻结（ListStandardsByRepository / 输入 repositoryID / 输出 []StandardReadResult 含全树 / 空列表非错误）
- [x] 消费方拥有模式沿袭 phase13-10 GovernanceProfileReader（落点 context_readers.go 追加 / 实现位 query_service.go / 注入位 router.go）
- [x] projectcontext 无 standard 表直接 SQL 约束
- [x] brief 两步分工表冻结（phase14-07 新增接入不动现状 / phase14-09 退役切换 reserved 3 + 内联化；过渡态字段面快照测试要求）
- [x] 跨包引用活跃合同源设计决策留档（与 T6 退役包依赖的性质区分）

## 一致性核对

- [x] 与 phase14-03 单值一致（proto 对齐表 / jsonb 6 字段 / 8 条校验规则引用 / DDL 三表投影）
- [x] 与 phase14-02 单值一致（八格矩阵 / brief 对照表 / StandardSummary 收敛闭环）
- [x] 与 shared_baseline §2.3/§2.4/§3.2-3.4 一致
- [x] 与仓库既有模式零漂移（错误哨兵 / candidate / 文件单值化；router.go L405-407 装配模式实证）
- [x] dev_plan L34-37 范围与 DoD 逐条满足
- [x] 零代码改动、零三件套正文改动、零根级改动（git status 验证）

## 复核与收口

- [x] 子代理独立复核通过（18 处证据链全部确认性判定、无 FAIL 项、特别核查点 4 项闭环）
- [x] tasks.md 全部勾选并附执行记录
- [x] 变更未提交，待用户最终确认后手动提交
