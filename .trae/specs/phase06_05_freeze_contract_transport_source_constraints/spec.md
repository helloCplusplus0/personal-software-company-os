# Phase06-05 当前阶段合同、传输与源码约束执行前提 Spec

## Why

`phase06` 已经分别冻结了 onboarding、draft-first、export/backup 和 reuse summary 的业务语义，但如果不继续把“合同源是谁、HTTP 适配只负责什么、前端写路径从哪里进、验收时如何判定一致”写成单值结论，后续实现仍然会在 `.proto`、HTTP DTO、前端 `useMutation` 落点和 `buf` 工具链入口上各自长出第二套口径。

## What Changes

- 冻结 `.proto` 作为 `phase06` 新增接口的唯一合同源与字段演进边界
- 冻结 `chi + HTTP JSON` 在当前阶段只承担传输适配职责，不承担第二套合同定义职责
- 冻结 `Export / Backup / Reuse Summary` 的 HTTP DTO、后端 `types.go` 与前端 `types.ts` 的单向派生约束
- 冻结前端 `application / query / mutation / shared` 四条约束在当前阶段的执行口径
- 冻结 `buf build / lint / generate / breaking` 的唯一工具链入口与 breaking check 前提
- 冻结当前阶段阶段验收必须覆盖的合同一致性与源码边界检查

## Impact

- Affected specs:
  - `phase06_onboarding_sovereignty_reuse_foundation`
  - `phase06-03` 导出 / 备份语义
  - `phase06-04` 复用感知读模型与挂接位
- Affected code:
  - 后续 `proto/` 下新增的 `phase06` 合同文件与 `buf.yaml / buf.gen.yaml / Makefile`
  - 后续 `backend/internal/export/`、`backend/internal/backup/`、`backend/internal/reusesummary/` 或等价模块
  - 后续 `backend/internal/*/types.go`、`handler/` 中的 HTTP DTO 与 envelope 映射
  - 后续 `frontend/src/features/onboarding/`、`export/`、`backup/`、`reuse-summary/` 或等价切片
  - 现有 `frontend/src/features/*/data/api-adapter.ts`
  - 现有 page / panel 级 `useMutation` 的回收范围与后续 phase06 新增写路径承接位

## ADDED Requirements

### Requirement: `phase06` 新增接口的唯一合同源冻结

系统 SHALL 将 `phase06` 中新增或扩展的正式接口统一冻结为“`.proto` 是唯一长期合同源”，不得再并列长出第二套 HTTP canonical contract。

#### Scenario: `phase06` 新增接口合同定义

- **WHEN** 接手者为 `Onboarding / Export / Backup / Reuse Summary` 新增或扩展正式接口
- **THEN** 字段、枚举、响应 envelope 与错误语义必须先在 `.proto` 中冻结
- **AND** 当前阶段不得先在 handler、`types.go`、`types.ts` 或 `api-adapter.ts` 中发明并列的业务字段语义
- **AND** 当前阶段允许保留 `chi + HTTP JSON` 作为传输层，但不得把 HTTP 形状误写成合同本体

#### Scenario: 合同演进边界

- **WHEN** 接手者演进 `phase06` 新增 `.proto`
- **THEN** 新增字段必须使用新的递增编号
- **AND** 删除字段、修改字段类型或修改已冻结字段语义时，必须按破坏性变更处理
- **AND** 删除或废弃字段后必须保留 `reserved` 字段号，必要时同时保留字段名

### Requirement: `chi + HTTP JSON` 传输适配职责冻结

系统 SHALL 将 `chi + HTTP JSON` 在当前阶段的职责冻结为传输适配，而不是第二套合同定义层。

#### Scenario: HTTP request 映射

- **WHEN** handler 承接 HTTP 请求
- **THEN** URL path、query params、JSON body 只允许作为 Proto request 的传输来源
- **AND** handler 在进入业务层前必须显式组装为对应 Proto request 或与之单向对齐的 DTO
- **AND** 当前阶段不得让 service / application 层直接依赖裸 `http.Request`、裸 query param 或隐式路径参数

#### Scenario: HTTP response 映射

- **WHEN** handler 返回 HTTP JSON 响应
- **THEN** 必须返回与 `.proto` 对齐的正式 response envelope
- **AND** 当前阶段不得把 `.proto` 已冻结为 envelope 的读取接口退回为裸对象、裸数组或临时拼装 map
- **AND** 当前阶段不得在 HTTP response 中额外加入 `.proto` 未定义的业务字段语义

### Requirement: HTTP DTO 单向派生冻结

系统 SHALL 将 `types.go`、`types.ts` 与 `api-adapter.ts` 中承接 `phase06` 接口的 HTTP DTO 冻结为“从 `.proto` 单向派生或显式对齐映射”的过渡传输层，而不是新的事实源。

#### Scenario: 后端 `types.go` 承接规则

- **WHEN** 接手者为 `Export / Backup / Reuse Summary` 新增后端消息结构
- **THEN** `types.go` 只能承接与 `.proto` 对齐的 JSON DTO
- **AND** 若存在传输层差异（如路径参数、时间字符串、状态码差异），必须在注释或显式映射中说明其仅为传输差异
- **AND** 当前阶段不得在 `types.go` 中新增 `.proto` 中不存在的业务字段语义、错误分类或成功态语义

#### Scenario: 前端 `types.ts` 承接规则

- **WHEN** 接手者为 `Export / Backup / Reuse Summary` 新增前端类型
- **THEN** 响应 / 域类型必须直接承接后端 JSON 形状，并与 `.proto` 语义单值一致
- **AND** 输入 / 写入参数若因前端表单使用 camelCase，可在 `api-adapter.ts` 中做显式字段转换
- **AND** 当前阶段不得在 `types.ts` 中增加 `.proto` 未定义、但会影响业务判断的私有字段

### Requirement: `Export / Backup / Reuse Summary` DTO 一致性冻结

系统 SHALL 将 `Export / Backup / Reuse Summary` 的 `.proto`、HTTP DTO 与前端消费模型之间的一致性冻结为当前阶段正式门禁。

#### Scenario: `Export / Backup` DTO 一致性

- **WHEN** 接手者实现 `Export / Backup` 的 HTTP 接口与前端消费模型
- **THEN** 资产覆盖矩阵、创建结果、`manifest`、`schema / version` 前提与错误语义必须与 `.proto` 字段单值一致
- **AND** 当前阶段不得在 HTTP 层自行新增另一套 `backup verified` 判定字段或错误归类
- **AND** 当前阶段不得让前端通过拼装多个无合同约束的字段自行推导第二套 canonical 语义

#### Scenario: `backup_snapshot` 读取侧一致性

- **WHEN** 接手者实现 `backup_snapshot` 的读取侧（用于 `backup verified` 校验中“可重新读取并解析备份 `manifest`”）
- **THEN** 读取侧必须由当前阶段 `Backup` 能力中的正式读时承接位负责，并与 shared baseline 已冻结的 `BackupWrite` owner 语义保持一致
- **AND** 当前阶段允许将该读取侧实现为 `BackupWrite` owner 内的 read / verify 子路径，或与上游冻结口径不冲突的等价读取承接位
- **AND** 该读取侧的 `.proto` 合同、HTTP DTO 与前端消费模型必须保持单值一致
- **AND** 该读取侧的 `manifest` 字段、覆盖矩阵字段与 `schema / version` 前提字段必须从 `.proto` 单向承接，不得在 HTTP 层或前端类型中补出第二套字段语义
- **AND** 当前阶段不得把“`BackupWrite` 响应中附带了一次 `manifest`”解释为已满足“可重新读取”要求

#### Scenario: `Reuse Summary` DTO 一致性

- **WHEN** 接手者实现 `ReuseSummaryRead`
- **THEN** `module_reuse_summary / capability_summary` 的字段、排序前提、空态语义与读取失败语义必须从 `.proto` 单向承接到 HTTP DTO 与前端消费模型
- **AND** 当前阶段不得在 HTTP DTO 或前端类型中额外补出与 `.proto` 不一致的排序字段、状态字段或解释字段

#### Scenario: `Onboarding`（含 `first_run_state`）DTO 一致性

- **WHEN** 接手者实现 `OnboardingRead`（含 `first_run_state` 的 `status / 是否首次进入 / 当前引导步骤 / 首轮完成条件`）
- **THEN** `first_run_state` 的字段、状态枚举（`not_started / in_progress / completed`）与跃迁语义必须从 `.proto` 单向承接到 HTTP DTO 与前端消费模型
- **AND** 当前阶段不得在 HTTP DTO 或前端类型中额外补出与 `.proto` 不一致的状态字段、跃迁判定字段或解释字段
- **AND** 当前阶段不得把 `first_run_state` 的状态机判定逻辑分散到多个 DTO 或前端类型中各自派生

### Requirement: 前端 `application / query / mutation / shared` 执行口径冻结

系统 SHALL 将前端四条约束在 `phase06` 的执行口径冻结为单值结论，避免新增写路径继续散落在页面与展示组件中。

#### Scenario: `query` 层职责

- **WHEN** 接手者实现 `OnboardingRead / ExportRead / backup_snapshot` 读取侧 / `ReuseSummaryRead` 等读取能力
- **THEN** `query` owner 只承接读取、`queryKey`、只读解包与 `queryOptions` 级别的配置
- **AND** 当前阶段不得在 `query` owner、读适配层或展示组件中夹带 `create / update / bind / link / export / backup write` 一类写动作
- **AND** 当前阶段允许 `query` owner 为页面与预取共享同一组只读配置，但不得吸收 mutation 逻辑

#### Scenario: `application` 与 `mutation` 承接位

- **WHEN** 接手者实现 `phase06` 的新增写路径
- **THEN** 正式 `useMutation`、失效刷新、成功回流、错误归一化必须收敛到切片内唯一 `application` 承接位
- **AND** 页面、表单、面板组件只保留字段收集、提交事件、局部 loading / error 展示与路由上下文展示职责
- **AND** 当前阶段不得继续在新增页面、展示组件或面板中直接内联正式 mutation 主线

#### Scenario: 旧实现过渡与新增门禁

- **WHEN** 仓库中存在既有 page-level 或 panel-level `useMutation`
- **THEN** 可以作为过渡现实继续存在，直到对应 `phase06` 实现任务回收
- **AND** 自 `phase06` 起新增写路径不得复制这种散装模式
- **AND** 若本阶段重构到相关写路径，优先向切片内固定 `application` 承接位回收

#### Scenario: `shared` 晋升边界

- **WHEN** 接手者抽取 `Export / Backup / Reuse Summary / Onboarding` 的共享代码
- **THEN** 默认先落在业务切片内部
- **AND** 只有在跨多个切片稳定复用、且语义清晰后才允许提升到 `shared`
- **AND** 当前阶段不得为了提早“抽象整洁”而把尚未稳定的 query key、DTO mapper、mutation owner 或页面状态机提前抽到 `shared`

### Requirement: `buf` 工具链入口与 breaking check 前提冻结

系统 SHALL 将 `buf build / lint / generate / breaking` 冻结为当前阶段 `.proto` 合同校验的唯一正式工具链入口。

#### Scenario: 唯一工具链入口

- **WHEN** 接手者校验或生成 `phase06` 的 `.proto`
- **THEN** 必须复用 `proto/Makefile` 与 `proto/buf.yaml / buf.gen.yaml` 既有入口
- **AND** 当前阶段不得为 `Export / Backup / Reuse Summary` 新增第二套 `buf.yaml`、`buf.gen.yaml`、并列 proto 根目录或私有生成脚本
- **AND** `make build / gen / lint / breaking` 或等价 `buf build / generate / lint / breaking` 必须保持可运行

#### Scenario: breaking check 基准

- **WHEN** 接手者执行 breaking 检查
- **THEN** 必须以仓库主线 `main` 分支中的 `proto/` 作为对比基准
- **AND** 当前阶段不得通过跳过 breaking、关闭规则或改写对比基准来绕过破坏性变更检查
- **AND** breaking check 失败必须被显式处理为阻断，不得在 CI 或本地流程中通过 `continue-on-error`、`allow-failure` 或等价机制降级为 warning

### Requirement: 当前阶段合同与源码约束验收口径冻结

系统 SHALL 将合同一致性、传输映射与前端承接位检查纳入 `phase06` 阶段验收口径，而不是实现完成后再补票。

#### Scenario: 合同一致性验收

- **WHEN** 当前阶段进入 `Onboarding / Export / Backup / Reuse Summary` 的验收
- **THEN** 至少必须验证：
  - `.proto` 可通过 `buf build / lint / generate / breaking`
  - HTTP DTO 与 `.proto` 的字段语义单值一致
  - 前端消费模型与 HTTP JSON 形状单值一致
  - `OnboardingRead`（含 `first_run_state` 状态枚举与跃迁语义）的 `.proto -> HTTP DTO -> 前端消费模型` 单值一致已被显式验证
  - `backup_snapshot` 读取侧（用于 `backup verified` 校验）的 `.proto -> HTTP DTO -> 前端消费模型` 单值一致已被显式验证
- **AND** 不得仅凭接口“能跑通”就判定合同一致性成立
- **AND** 不得只验证 `Export / Backup / Reuse Summary` 而遗漏 `Onboarding` 合同一致性
- **AND** 不得只验证 `BackupWrite` 写入响应而遗漏 `backup_snapshot` 读取侧合同一致性

#### Scenario: 前端承接位验收

- **WHEN** 当前阶段进入新增写路径验收
- **THEN** 至少必须验证：
  - 新增正式 mutation 已收敛到切片内固定 `application` 承接位
  - `query` owner 保持只读
  - 页面 / 面板 / 表单未再复制新的正式写路径
- **AND** 不得把“页面里 `useMutation` 也能提交成功”误判为满足阶段约束

## MODIFIED Requirements

### Requirement: 当前阶段过渡传输层承诺

当前项目在 `phase06` 中 SHALL 继续允许 `chi + HTTP JSON` 作为过渡传输层，但其正式承诺必须被收敛为“对 `.proto` 合同做单向适配与显式映射”，而不是自由演化的第二套 API 定义层。

#### Scenario: 过渡传输层语义回落

- **WHEN** 接手者设计 `phase06` 的 HTTP 接口
- **THEN** 必须把 `chi + HTTP JSON` 视为 transport adapter
- **AND** 必须把 `.proto` 视为唯一合同源
- **AND** 不得把“当前还没切到 gRPC / Connect”误解释为可以先放宽合同一致性要求

## REMOVED Requirements

### Requirement: `phase06` 新增接口可先在 HTTP DTO 或页面层定义业务语义，后续再回填 `.proto`

**Reason**: 这种做法会把 `.proto`、HTTP DTO、前端消费模型和错误语义重新拉回并列演化状态，直接破坏当前项目“唯一合同源 + 单一写路径”的阶段目标。

**Migration**: `phase06` 后续实现统一改为：先冻结 `.proto` 合同，再在 handler、`types.go`、`types.ts`、`api-adapter.ts` 中做单向派生或显式对齐映射；新增正式 mutation 统一收敛到切片内固定 `application` 承接位。
