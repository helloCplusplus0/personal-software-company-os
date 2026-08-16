# Tasks

- [x] Task 1: 盘点 `phase12-11` 的直接上游、固定样本与验收协议输入
  - [x] SubTask 1.1: 审阅 `dev_plan#L272-L334`，冻结 `phase12-11` 的范围、验收协议与 DoD
  - [x] SubTask 1.2: 审阅 `phase12-01 / 05 / 06 / 08 / 09 / 10` 的正式输入与当前实现状态，确认本轮只承接联调、dogfooding 与反回归验证
  - [x] SubTask 1.3: 继承 `phase11-09` 已冻结的 `PSCO` dogfooding 样本与 `repository_id` 证据，明确“样本 != 未来模板”的前置说明
  - [x] SubTask 1.4: 冻结 `Repository / repository_id / Product / Module / Decision` 的单值正式样本，禁止临场换样本

- [x] Task 2: 冻结最小工具链验证矩阵
  - [x] SubTask 2.1: 明确 `buf` 的正式命令、执行顺序与通过标准
  - [x] SubTask 2.2: 明确 `go test` 的正式命令、执行顺序与通过标准
  - [x] SubTask 2.3: 明确 `frontend build` 的正式命令、执行顺序与通过标准
  - [x] SubTask 2.4: 冻结 warning、环境异常与命令失败的归类口径，不允许临场解释

- [x] Task 3: 冻结同一 `repository_id` 锚点下的样本解析协议
  - [x] SubTask 3.1: 明确 `repository_id` 只允许使用固定样本中的正式值
  - [x] SubTask 3.2: 明确 `product_id / module_id / decision_id` 只允许从同一 `repository_id` 驱动的结构化只读结果或其受控派生视图解析
  - [x] SubTask 3.3: 明确解析链路必须留档：固定入口、解析出的 id、是否一次成功
  - [x] SubTask 3.4: 明确名称解析失败、不唯一或无法回到同一 `repository_id` 时的直接失败判定

- [x] Task 4: 冻结 Web 侧 primary owner 与跟随回归页面的浏览器验收矩阵
  - [x] SubTask 4.1: 冻结 `/repositories/$repositoryId`、`/products/$productId`、`/modules/$moduleId`、`/decisions/$decisionId` 的 primary owner 验证清单
  - [x] SubTask 4.2: 冻结 `/dashboard`、`/onboarding`、`/reviews/daily`、`/reviews/weekly` 的跟随回归验证清单
  - [x] SubTask 4.3: 明确每个页面需要验证的四实体解释、共享只读消费、返回链与 reread 断言
  - [x] SubTask 4.4: 明确 Web 侧详情页不得在页面内临场猜测 `repository_id`

- [x] Task 5: 冻结 Web / agent 共享只读 dogfooding 协议与固定 `6` 问
  - [x] SubTask 5.1: 冻结 agent 侧固定入口集合：`AGENTS.md`、`PSCO-mvp05-summarize-feedback.md`、同一 `repository_id` 的共享只读结果
  - [x] SubTask 5.2: 冻结 Web 侧回答固定 `6` 问时允许使用的固定页面与固定入口
  - [x] SubTask 5.3: 明确 `Module / Decision` 的误读风险必须作为固定提问显式复核
  - [x] SubTask 5.4: 明确 Web 与 agent 必须回到同一组规则、约束与文档入口，不允许长出双轨答案

- [x] Task 6: 冻结本阶段边界证据与失败判定
  - [x] SubTask 6.1: 明确本阶段不做 schema 重写的边界证据要求
  - [x] SubTask 6.2: 明确本阶段不做 MCP / CLI / agent 写回 / 对话入口的边界证据要求
  - [x] SubTask 6.3: 明确不允许通过额外第 `7` 个临时入口补齐回答
  - [x] SubTask 6.4: 明确工具链失败、样本解析失败、语义误读与双事实源出现时的统一失败判定

- [x] Task 7: 将 `phase12-11` 的冻结结果回写到 spec 包并完成自检
  - [x] SubTask 7.1: 复核 `spec.md` 是否已提供单值样本、单值入口、单值 `6` 问与单值失败判定
  - [x] SubTask 7.2: 复核 `tasks.md` 是否足以让后续执行者不再补造验收主路径
  - [x] SubTask 7.3: 复核 `checklist.md` 是否覆盖工具链、样本解析、浏览器、dogfooding、边界证据与 rerun 条件
  - [x] SubTask 7.4: 复核本 spec 包是否足以让不同执行者按相同样本、相同入口与相同问题复跑

- [x] Task 8: 修复 `phase12-11` 浏览器 / agent 联调中暴露的固定入口阻断
  - [x] SubTask 8.1: 让固定三入口中的至少一个共享只读结果能直接确认 `Product = 经营目标与交付容器`
  - [x] SubTask 8.2: 让固定三入口中的至少一个共享只读结果能直接确认 `Repository = 代码仓库身份对象与项目锚点`
  - [x] SubTask 8.3: 让 shared-readonly 显式带出 `plan.md / architecture_map.md / docs/README.md` 与当前 `phase12` 入口，消除 Q5 阻断
  - [x] SubTask 8.4: 重跑 fixed `6` 问与相关 shared-readonly 验收，确认 Q1 / Q2 / Q5 由“不达标”转为“达标”

# Task Dependencies

- `Task 2` depends on `Task 1`
- `Task 3` depends on `Task 1`
- `Task 4` depends on `Task 1`
- `Task 5` depends on `Task 1`
- `Task 6` depends on `Task 1`
- `Task 4` depends on `Task 3`
- `Task 5` depends on `Task 3`
- `Task 7` depends on `Task 2`
- `Task 7` depends on `Task 3`
- `Task 7` depends on `Task 4`
- `Task 7` depends on `Task 5`
- `Task 7` depends on `Task 6`
- `Task 8` depends on `Task 5`
- `Task 8` depends on `Task 6`
- `Task 7` depends on `Task 8`
