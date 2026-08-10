# Audit 001 Analysis - Go 传输层单一真相与 ConnectRPC 主线迁移

## 1. 审计摘要
- 对应问题：`audit_001`
- 审计级别：`P1`
- 是否建议进入后续动作：是
- 最终去向：`escalate-phase`

## 2. 模块目标复述
- 该模块要解决什么问题：
  在 Go 后端中保证 `.proto` 是唯一长期合同源，同时避免 `chi + JSON HTTP` 在后续演进中变成第二套事实源。
- 该模块服务哪些角色：
  后端实现者、前端消费方、合同维护者、阶段规划者。
- 当前正式定位是什么：
  `project_rules.md` 已冻结：`.proto` 是唯一长期合同源，`chi` 只承担路由与中间件装配职责，业务接口默认优先采用 `ConnectRPC` 作为正式传输层。

## 3. 当前现状结论
- 当前设计中已经合理的部分：
  - `proto/` 已成为统一合同源
  - `buf build / lint / breaking / gen` 已成为统一工具链入口
  - `chi` 作为顶层装配根、middleware 与子路由挂载层，工程上稳定且简单
  - 非业务端点如 `healthz / metrics / debug` 保持 `chi + net/http` 原生处理是合理的
- 当前设计中存在疑问的部分：
  - 当前业务接口仍以手工 JSON 映射矩阵承接
  - 生成链目前只生成消息类型，尚未生成 Connect 相关 handler/client
  - 若后续新增业务接口继续复制该模式，`.proto canonical` 与 `JSON transport contract` 之间的漂移风险会持续上升
- 当前实现中需要重点关注的边界：
  - `chi` 可以保留，但不能继续定义业务合同
  - 存量 JSON 业务接口只允许在迁移过程中短时存在，不应作为 phase 收口后的长期稳态
  - 迁移必须是“业务接口切到 ConnectRPC，基础设施接口继续保留 chi 原生端点”

## 4. 开发者视角审计
- 问题是否成立：
  成立，但要精确表述为：**当前实现尚未失控，但继续以 `chi + JSON HTTP` 作为新增业务接口默认模式，会系统性放大双真相风险。**
- 当前实现是否符合工程目标：
  在 `phase06` 之前符合“最小落地 + 过渡实现”的工程目标；从 `mvp0.3` 开始，不再是最佳长期形态。
- 当前实现是否符合既有冻结原则：
  当前阶段总体符合，因为 `.proto` 仍是 canonical source，JSON 仍被定义为过渡层。
- 当前实现距离最佳实践的主要差距：
  业务传输层没有直接由 `.proto` 生成的正式 handler/client 主线，仍依赖手工 DTO 与显式映射矩阵。
- 关键代码位置：
  - 路由装配：[backend/internal/platform/router.go](file:///home/dell/Projects/personal-software-company-os/backend/internal/platform/router.go)
  - 合同入口：[proto/README.md](file:///home/dell/Projects/personal-software-company-os/proto/README.md)
  - 工具链入口：[proto/Makefile](file:///home/dell/Projects/personal-software-company-os/proto/Makefile)

## 5. 用户视角审计
- 问题是否成立：
  成立。用户不关心底层框架，但会直接承担“接口语义漂移、前后端类型不同步、错误码不一致”的成本。
- 当前交互是否顺手：
  目前可用，但“可用”并不等于“可长期扩展”。
- 当前信息表达是否清晰：
  对维护者来说不够清晰，因为“`.proto` 是唯一真相”与“业务接口仍手工 JSON 承接”之间存在长期紧张关系。
- 当前反馈是否可信、完整：
  对已完成阶段是可信的；对未来新增接口来说，可信度会随手工映射扩大而下降。
- 当前体验距离最佳实践的主要差距：
  业务合同没有直接在传输层闭合，必须靠文档纪律与人工映射来维持一致性。

## 6. 综合审计结论
- 最终判断：
  **PSCO 的最佳实践应冻结为：保留 `chi` 作为 HTTP 装配根与非业务端点承载层，同时将业务接口正式收敛到 `.proto + ConnectRPC` 主线。**
- 判断依据：
  1. 这与当前根级规则已经对齐
  2. 这避免了“全盘重写成重型 gRPC 体系”的过度动作
  3. 这同时阻断了手工 JSON DTO 继续演化为第二套合同源
- 当前是否值得改动：
  值得，但不是以局部修补方式，而是作为下一阶段正式规划议题。
- 若不改，原因：
  不适用。
- 若要改，原因：
  `mvp0.3` 将继续新增业务接口与写路径；如果不在这一轮把正式业务传输主线收紧，后续每个 phase 都会继续累积过渡写法。
  同时，既然该问题已升级为独立前置 phase，就不应把“旧 JSON 业务接口长期兼容”当作 phase 收口后的合理状态。

## 7. 最佳实践对标
- 最佳实践应是什么：
  - `.proto`：唯一长期合同源
  - `ConnectRPC`：业务接口正式传输层
  - `chi`：顶层 router / middleware / 非业务端点 mount 层
  - 存量 JSON 业务接口：仅允许作为迁移过程中的临时兼容适配层，而不是 phase 收口后的长期 canonical API
- 当前方案与最佳实践的差异：
  当前仓库仍停留在“合同 canonical 已冻结，但业务正式传输尚未切换”的过渡阶段。
- 是否需要完全看齐：是
- 若不完全看齐，保留差异的理由：
  只允许在**迁移节奏**上分阶段，不允许在**最终结构**上继续模糊。

## 8. 候选方案对比
### 方案 A
- 做法：
  保持现状，继续用 `chi + JSON HTTP` 承接新增业务接口，只靠规则约束 `.proto` 对齐。
- 优点：
  - 短期改动最少
  - 不需要调整生成链
- 风险：
  - 新接口越多，人工映射越多
  - `proto canonical` 与 JSON DTO 漂移风险持续累积
  - 后续再迁移成本只会更高

### 方案 B
- 做法：
  立即推翻 `chi`，全盘改为重型 gRPC / Protobuf 工具链。
- 优点：
  - 业务合同理论上最统一
- 风险：
  - 与当前仓库结构和单人维护现实不匹配
  - 前后端与运维成本明显上升
  - 会把“收敛正式业务传输主线”升级成“全面重写后端传输框架”

### 方案 C
- 做法：
  保留 `chi` 作为装配层与非业务端点承载层，在独立前置 phase 中将 `phase01 ~ phase06` canonical 业务接口一次性切到 `ConnectRPC`，旧 JSON 业务接口仅在迁移过程短时存在并在 phase 收口前退场。
- 优点：
  - 保住当前路由与 middleware 结构
  - 让业务合同直接收敛到 `.proto`
  - 迁移成本可控，可按 phase 分段推进
- 风险：
  - 需要调整 `buf.gen.yaml`、生成链与前后端调用方式
  - 迁移期会短暂并存“正式 ConnectRPC + 兼容 JSON adapter”

## 9. 推荐方案
- 推荐原因：
  推荐 **方案 C**。它是当前 PSCO 在工程现实、合同单一真相与长期维护成本之间的最佳平衡点。
- 最小实施边界：
  1. `chi` 不退场，但正式退回 transport shell 身份
  2. `phase01 ~ phase06` 已交付 canonical 业务接口在该前置 phase 内全部切到 `ConnectRPC`
  3. `healthz / readyz / metrics / debug` 继续保留 `chi + net/http`
  4. 存量 JSON 业务接口只允许在迁移过程中短时并存，phase 收口前必须退场
- 需要保持不变的内容：
  - `.proto` 唯一长期合同源
  - `buf build / lint / breaking` 工具链入口
  - `chi` 作为顶层路由与 middleware 装配层
- 明确不在本次改进范围内的内容：
  - 一次性迁移所有 HTTP 端点（非业务基础设施端点仍保留在 `chi + net/http`）
  - 在当前审计文档里直接给出具体 phase 名称
  - 引入第二套微服务或重型 gRPC 基础设施

## 10. 后续去向说明
### 若结论为 `keep-as-is`
- 保持原样的原因：
  不适用。

### 若结论为 `enter-fix`
- 为什么它本质上是 bug / 错误行为 / 明确偏差：
  不适用；这不是单点错误，而是结构性演进议题。
- 后续应转入的 `fix` 编号：
  不适用。

### 若结论为 `enter-improvement`
- 为什么它属于优化而非 bug：
  部分成立，但不足以覆盖其跨模块、跨阶段的规划性质。
- 后续建议以何种实现方式承接：
  不优先采用本路径。

### 若结论为 `escalate-phase`
- 为什么它已超出局部优化范围：
  因为它同时影响：
  - `proto` 合同与生成链
  - Go 后端业务传输层
  - 前端客户端生成与调用方式
  - 存量业务接口兼容策略
  - 下一阶段新增业务接口默认模式
- 需要升级到什么级别的阶段议题：
  升级为下一阶段 `/plan` 的结构性议题，至少应在后续 phase 文档中覆盖：
  1. 传输层主线切换原则
  2. `buf.gen.yaml` 与生成链补齐 Connect 相关产物
  3. `phase01 ~ phase06` 业务接口一次性迁移到 ConnectRPC 的范围、顺序与验收口径
  4. 旧 JSON 业务接口仅作为迁移过程短时适配层并在 phase 收口前退场

## 11. 验收标准
- 根级规则、技术基线与审计结论之间对 ConnectRPC 主线的口径一致
- 下一阶段正式 `/plan` 明确承接“业务接口走 ConnectRPC、非业务端点保留 chi”这一结构
- 新增业务接口不再复制手写 `chi + JSON HTTP` DTO 模式
- `phase01 ~ phase06` canonical 业务接口已完成正式传输主线切换
- phase 收口后不再保留手写 `chi + JSON HTTP` 业务主线

## 12. 风险与回退条件
- 主要风险：
  - 迁移期短时同时存在正式 ConnectRPC 与兼容 JSON adapter，若边界不清会重新长出双真相
  - 生成链改造若不受控，可能影响现有前后端构建流程
- 禁止引入的新问题：
  - 不得把 `chi` 误删成“必须彻底移除”
  - 不得把所有基础设施端点一并纳入 `.proto`
  - 不得在迁移期新增第二套 hand-written business DTO 规范
- 若后续实施不成立，回退条件：
  - 若某一批迁移无法稳定落地，可在开发过程短时保留对应 JSON adapter 以完成排障，但该状态不得作为 phase 收口结果通过验收
