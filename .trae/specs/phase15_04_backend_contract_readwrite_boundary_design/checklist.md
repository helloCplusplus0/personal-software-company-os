# phase15-04 Checklist

- [x] proto 草案冻结：`psco.progress.v1` 注释版草案（3 枚举 UNSPECIFIED=0 + `ProgressEvent` 11 字段号 1-11 按 DDL 列序 + 3 RPC envelope + `ProgressService`）可逐字转写为 `proto/psco/progress/v1/progress.proto`（生成链：Go pb / Go Connect / TS 三端）
- [x] 7 项合同设计决策冻结且理由留档：字段号自然分配 / workflow_type 过滤零值=不过滤 / source optional 未设置归一 manual / 可空文本三字段空串语义（DB NULL ↔ 空串 repository 层转换）/ Create 响应返回完整事件 / Delete 响应空消息 / 时间字段 Timestamp
- [x] 无 Update 显式声明（proto service 定义处注释 + spec 声明；append-only 语义纯净，裁决⑨；误录修正 = Delete + 重新 Create）
- [x] 3 RPC 错误语义逐个三要素冻结：List（非 UUID → InvalidArgument / 不存在 → NotFound 读锚点 / 空 → 非错误 / Internal）；Create（执行序 6 步 → InvalidArgument / 写入失败 → Internal）；Delete（非 UUID → InvalidArgument / 不存在 → NotFound / Internal）
- [x] DP-2 承接位单值裁决：`progress/candidate/RepositoryReader.RepositoryExists(ctx, repositoryID) (bool, error)` 纯事实查询（沿 standard TargetReader 模式）；错误语义包装归 service（Create → ErrInvalidInput + REPOSITORY_NOT_FOUND / List → ErrRepositoryNotFound）；FK RESTRICT 为存储层兜底非校验承接位
- [x] `REPOSITORY_NOT_FOUND`（跨模块引用校验码）与 `INVALID_PROGRESS_EVENT_ID`（Delete id 格式码）新增留档，显式区分于 phase15-03 的 9 业务码 + 2 envelope 码；Create（InvalidArgument）与 List（NotFound）错误语义辨析留档（各自冻结上游不改写）
- [x] 哨兵错误清单冻结（5 个：ErrProgressEventNotFound / ErrRepositoryNotFound / ErrInvalidInput / ErrProgressReadFailed / ErrProgressWriteFailed）；connecterrors 登记点冻结（NotFound ×2 + InvalidArgument ×1 + Internal 兜底 ×2）
- [x] Go 模块分层冻结：根包 4 文件（types.go 含受控枚举 + ProgressEventReadResult + ProgressSummary + CreateProgressEventInput 草案 / errors.go / validate.go / derive.go）+ 四子包（connect 导出 DomainProgressEventToProto / service 双 service / repository 单查询 store / candidate RepositoryReader）；service 依赖注入形态冻结
- [x] platform 装配点冻结：buildProgress + mountProgressConnect（沿 buildStandard / mountStandardConnect 模式）+ buildProjectContext 新增 progressReader 参数
- [x] ProgressReader 接口签名冻结：落点 projectcontext/candidate/context_readers.go（与 StandardReader 同文件）；`GetProgressSummary(ctx, repositoryID) (progress.ProgressSummary, error)`；实现位 progress/service.QueryService；4 项设计决策留档（不含全量事件集 / RecentEvents 恒非 nil / LatestTaskCompleted 指针可空 / Reader 不做存在性校验）
- [x] brief progress = 9 装配演进点冻结：BriefProgress 字段号 1-4 自然序（current_phase_key / current_phase_label / latest_task_completed / recent_events）；project_context.proto 新增 import psco/progress/v1/progress.proto；GetProjectBriefResponse.progress = 9（槽位 2/3/4 保持 reserved）；Go 编排新增步骤 6（恒构造空态）；转换函数复用 progress/connect.DomainProgressEventToProto
- [x] 与 phase15-02/03 语义上游单值一致（枚举值域 / 11 字段对应 / 执行序零改写 / ProgressSummary 四字段 1:1 / BriefProgress 字段号为本 spec 冻结项）；无 phase15-05/06 偷渡（无组件树 / 无交互规格 / 无 DP-1 / 无 DP-3 / 无实现代码与 proto 文件实体）
- [x] 独立复核通过（0 阻断；复核维度：proto 可逐字转写性 / 错误语义完备性 / DP-2 自洽性 / 模式一致性 / 单值性 / 零漂移 / 无偷渡 / 勾选真实性）
- [x] tasks.md / checklist.md 全部勾选并附执行记录
- [x] 变更未提交，待用户最终确认后手动提交
