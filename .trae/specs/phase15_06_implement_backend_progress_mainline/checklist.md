# phase15-06 Checklist

- [x] `proto/psco/progress/v1/progress.proto` 与 phase15-04 草案逐字一致（3 枚举值域 / `ProgressEvent` 11 字段号 / 3 RPC envelope / 无 Update 显式声明注释 / 文件头文档定位）
- [x] `project_context.proto` 演进单值：import progress.proto + `BriefProgress` 字段号 1-4（latest_task_completed 与 recent_events 同型复用 `psco.progress.v1.ProgressEvent`）+ `progress = 9`；reserved 2/3/4 与既有字段 1、5-8 零改动（git diff 逐行核实）
- [x] `make gen` 三端产物齐备（Go pb / Go Connect / TS；progress 新包 + project_context 再生成）；`make lint` / `make build` / `make breaking` 全部零退出码
- [x] `0013_phase15_progress_timeline.sql` 与 phase15-03 DDL 草案逐字一致（11 列 / 三列 TEXT+CHECK / FK RESTRICT / `idx_progress_events_repository_sort` 三键形态 / IF NOT EXISTS 幂等 / 无 DO 块无 DROP 段）；零手工登记代码改动（RunMigrations 自动登记机制）
- [x] 0013 已应用到开发库（等效 psql 幂等执行〔本机无 psql，经 docker 容器同库执行〕；`progress_events` 表与索引存在且幂等重放空过；用户已启动的服务器未被重启、未在其他端口重复开启）
- [x] `types.go` 与 phase15-04 草案逐字一致（三受控枚举对齐 DDL CHECK / `ProgressEventReadResult` 11 字段 / `ProgressSummary` 四字段〔LatestTaskCompleted 指针可空、RecentEvents 恒非 nil〕/ `CreateProgressEventInput`）；不 import 生成 pb 包
- [x] `errors.go` 5 哨兵与 phase15-04 清单逐字一致；`connecterrors` 登记齐备（NotFound 组 2 + InvalidArgument 组 1 + Internal 兜底显式引用 2）
- [x] `validate.go` 与 phase15-03 冻结一致：11 错误码全量 + 执行序 6 步报第一个错误 + K-1~K-4 正则逐字 + V7 矩阵组合判定 + TrimSpace 边界（task_key trim 持久化 / title trim 判定原值入库 / rune 计数）+ `%w: [CODE] message` 格式
- [x] `derive.go` 三派生算法与 phase15-03 逐字一致（recent N=10 / latest_task_completed 首个命中 / current_phase 三态含全部完结同 key 匹配）；纯函数零 I/O 零时间函数
- [x] `repository/` 单一查询沿 phase15-03 冻结 SQL 形态（ORDER BY 三键链，Go 不重排；workflow_type 可选过滤）；NULL↔空串解码在 scan 层；Delete 0 行 → `ErrProgressEventNotFound`
- [x] `candidate/RepositoryReader.RepositoryExists` 纯存在性事实查询（DP-2 承接位，SELECT EXISTS 形态）；不承载业务错误语义
- [x] 3 RPC 错误语义逐项落地（phase15-04 三要素表）：List 不存在 → NotFound / Create 不存在 → InvalidArgument 含 `REPOSITORY_NOT_FOUND` / Delete id 非 UUID → `[INVALID_PROGRESS_EVENT_ID]` / Delete 不存在 → NotFound / List 过滤空 → 空列表非错误 / Create source nil 归一 manual、显式非 MANUAL → `INVALID_SOURCE`
- [x] `connect/` 导出 `DomainProgressEventToProto`（供 projectcontext 复用）；错误统一 `connecterrors.MapToConnectError`
- [x] projectcontext 接入：`ProgressReader` 接口与 StandardReader 同文件追加（签名逐字）；`ProjectBriefReadResult.Progress` 值类型字段；`GetProjectBrief` 编排步骤 6（失败透传 `ErrProgressReadFailed`）；connect 组装空态恒构造（progress 非 nil / recent_events 空数组 / latest_task_completed 不设置）
- [x] platform 装配：`buildProgress` / `mountProgressConnect`（沿 standard 模式）；`buildProjectContext` 三参签名演进；server.go 接线顺序 progress 先于 buildProjectContext；brief 的 standards 与 progress 双 reader 注入互不影响
- [x] 校验单元测试：合法矩阵 12 格全判定（8 格正例 + 4 禁止格 V7 反例——独立复核裁决：spec"12 格正例"为笔误级冲突，audit/fix × phase_started/phase_completed 按 phase15-02 语义矩阵为禁止格，以上游冻结口径为准）+ 11 错误码逐码反例（含 audit/fix × phase 边界两用例、K-1~K-4 逐正则不符值、title 空与超长）+ 执行序首错断言；按错误码断言
- [x] 派生单元测试：空集 / 仅 note / 从未开始 / 进行中 / 全部完结 / 补录同刻 tiebreak / recent 截断（>10）全边界纯内存覆盖
- [x] progress service 集成测试：Create→List round-trip（三键链倒序含补录与同刻 tiebreak）+ 三轨过滤 + Delete 后 NotFound + Create 各错误分支；沿 standard 模式（真实 DB / 无 DB skip / 独立 fixture + cleanup）
- [x] brief 集成测试：progress round-trip（current_phase_key / label / recent_events / latest_task_completed）+ **phase_completed 后 current_phase_key 为空串**（DoD 冻结断言）+ 空态恒构造
- [x] 门禁全绿：`make lint` / `make build` / `make breaking` + `go build ./...` / `go vet ./...` / `go test ./...` 全部零退出码（集成测试连开发库执行非 skip；独立复核亲自复验）
- [x] 零偷渡：无前端手写代码（frontend/ git status 为空，gen/ 生成物 gitignore）、无 phase15-08 dogfooding（progress_events 表 0 行）、无根级文档回写、无 UpdateProgressEvent、无 phase11 PhaseEntry 改动
- [x] git 工作区改动与 spec §What Changes 文件清单一一对应（+ 本 spec 三件套目录）
- [x] 独立复核通过（0 阻断：实现与上游零漂移 / 测试 DoD 全项 / 门禁实测 / 无偷渡 / 服务器未重启〔后端进程自 11:26 持续运行，全机无第二服务器进程〕）
- [x] tasks.md / checklist.md 全部勾选并附执行记录（见 tasks.md §执行记录）
- [x] 变更未提交，待用户最终确认后手动提交

## 非阻断观察（留档，供 phase15-09 澄清）

1. spec"合法矩阵 12 格正例"表述与 phase15-02 语义矩阵存在笔误级冲突——实际落地为"12 格全判定（8 正例 + 4 禁止格反例）"，建议 phase15-09 根级同步时澄清。
2. spec 运行时约束文中"后端 :8080"与用户环境实际监听端口（8081）不符——环境配置层面观察，恰好佐证无重复开服。
3. `schema_migrations` 尚未登记 0013 属 spec 预期内状态（服务器下次重启 RunMigrations 空过补登记）——phase15-07/08 联调前提依赖此补登记，需在后续子任务确认。
