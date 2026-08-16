# Tasks

- [x] Task 1: 盘点 `phase11` 三件套中与 `phase11-05` 直接相关的项目上下文导出表达
  - [x] SubTask 1.1: 审阅 `phase11_project_context_foundation_dev_plan.md#L113-L161` 中当前 `phase11-05` 的范围、正式产物与 DoD
  - [x] SubTask 1.2: 审阅 `phase11_project_context_foundation_architecture_plan.md` 中关于输入锚点、聚合边界、Decision 聚合与 Markdown 导出职责的正式表达
  - [x] SubTask 1.3: 审阅 `phase11_project_context_foundation_shared_baseline.md` 中项目上下文导出矩阵、Decision 聚合口径与导出形式冻结

- [x] Task 2: 冻结最小只读导出的输入锚点与失败语义
  - [x] SubTask 2.1: 固定 `repository_id` 为当前阶段唯一正式结构化输入锚点
  - [x] SubTask 2.2: 固定不以本地路径、Git remote URL、`product_id` 或工作区扫描作为并列主锚点
  - [x] SubTask 2.3: 固定只承接"已完成 Repository Binding"的仓库上下文读取与未绑定仓库失败语义

- [x] Task 3: 冻结聚合内容范围与 canonical 投影边界
  - [x] SubTask 3.1: 明确聚合哪些 `Repository / Product / Module / Decision / 规则信息`
  - [x] SubTask 3.2: 明确哪些信息属于结构化只读输出
  - [x] SubTask 3.3: 明确哪些边界依赖 `PSCO` canonical 关系、哪些不依赖消费侧项目目录结构

- [x] Task 4: 冻结 `Decision` 聚合口径、去重与过滤规则
  - [x] SubTask 4.1: 固定三类直接 canonical 命中的 `Decision` 聚合范围
  - [x] SubTask 4.2: 固定不得沿 `Product -> Module -> 其他 Repository` 做递归扩张
  - [x] SubTask 4.3: 固定 `decision_id` 去重、命中来源摘要与 `archived` 过滤规则

- [x] Task 5: 冻结结构化只读输出与 Markdown 导出的职责边界
  - [x] SubTask 5.1: 固定结构化只读读取继续落在 Go backend 的 `.proto + ConnectRPC` 正式主线
  - [x] SubTask 5.2: 固定 Markdown 导出字段边界与 AGENTS 风格导出职责
  - [x] SubTask 5.3: 固定"结构化读取 -> Markdown 导出"的单向派生关系，不形成第二套事实源

- [x] Task 6: 冻结当前阶段明确不做的协议、写路径与目录依赖
  - [x] SubTask 6.1: 明确当前阶段不承接 `MCP / CLI` 协议层
  - [x] SubTask 6.2: 明确当前阶段不承接 agent 写回、审批流、主动注入或第二套 canonical API
  - [x] SubTask 6.3: 明确当前阶段不把消费侧固定文件名、固定目录结构或统一项目模板作为前置依赖

- [x] Task 7: 冻结 `phase11-05` 的成功标准、DoD 与收口口径
  - [x] SubTask 7.1: 将"何时算完成、何时不得判定完成"写成单值口径
  - [x] SubTask 7.2: 保证后续执行者无需再猜"上下文到底聚合到什么程度"
  - [x] SubTask 7.3: 固定超出范围的协议与写路径扩张必须后移

- [x] Task 8: 完成三件套一致性校验
  - [x] SubTask 8.1: 校验 `architecture_plan`、`dev_plan`、`shared_baseline` 对输入锚点、聚合范围与导出职责边界的表述单值一致
  - [x] SubTask 8.2: 校验 `Decision` 聚合规则、失败语义与非目标口径在三件套中单值一致
  - [x] SubTask 8.3: 校验后续执行者不会再把最小只读导出误解为协议层、写路径或目录扫描方案

- [x] Task 9: 将 `phase11-05` 的冻结结果显式回写到目标源文档
  - [x] SubTask 9.1: 在 `phase11_project_context_foundation_dev_plan.md` 中冻结输入锚点、聚合边界、正式产物与 DoD 口径
  - [x] SubTask 9.2: 在 `phase11_project_context_foundation_architecture_plan.md` 中补齐结构化只读输出字段边界、Markdown 导出字段边界与单向派生约束
  - [x] SubTask 9.3: 在 `phase11_project_context_foundation_shared_baseline.md` 中补齐导出矩阵对应的字段边界、失败语义承接与单向派生约束

- [x] Task 10: 校验本 spec 包的验收对象已经落到 `phase11` 三件套实际回写结果
  - [x] SubTask 10.1: 确认验收不只检查 spec 包自身，而是检查三件套是否已完成正式冻结
  - [x] SubTask 10.2: 确认输入锚点、失败语义、聚合范围与导出职责边界已在目标源文档中可直接引用
  - [x] SubTask 10.3: 确认"只做最小只读导出、不做协议与写路径扩张"的边界与收口口径已在目标源文档中可直接引用

# Task Dependencies

- Task 2 depends on Task 1
- Task 3 depends on Task 1
- Task 4 depends on Task 1
- Task 5 depends on Task 1
- Task 6 depends on Task 1
- Task 7 depends on Task 2
- Task 7 depends on Task 3
- Task 7 depends on Task 4
- Task 7 depends on Task 5
- Task 7 depends on Task 6
- Task 8 depends on Task 2
- Task 8 depends on Task 3
- Task 8 depends on Task 4
- Task 8 depends on Task 5
- Task 8 depends on Task 6
- Task 9 depends on Task 7
- Task 10 depends on Task 8
- Task 10 depends on Task 9
