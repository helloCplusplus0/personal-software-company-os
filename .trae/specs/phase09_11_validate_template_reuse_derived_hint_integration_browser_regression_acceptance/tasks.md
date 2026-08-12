# Tasks

- [x] Task 1: 冻结 `phase09-11` 的统一验收输入边界与上游证据来源
  - [x] SubTask 1.1: 对齐 `docs/phase/phase09_template_reuse_derived_intelligence_foundation_dev_plan.md#L238-L264` 的范围、DoD 与非目标
  - [x] SubTask 1.2: 继承 `phase09-08 / 09 / 10` 的已交付能力，明确它们在本阶段只作为上游输入与待复验项
  - [x] SubTask 1.3: 明确本阶段统一验收只覆盖模板复用与派生提示当前主线，不扩写 `dry-run / AI Context Enhancement / Venture`

- [x] Task 2: 冻结 `phase09-11` 的环境准备、工具链验证与 API smoke 顺序
  - [x] SubTask 2.1: 明确前端、后端、数据库与 `/api` 主线的统一前置检查顺序
  - [x] SubTask 2.2: 明确 `buf / frontend type-check / build / backend build` 的执行顺序、通过标准与失败判定
  - [x] SubTask 2.3: 明确模板候选读取、模板预填详情读取、派生提示读取三类 API smoke 的最小验证矩阵

- [x] Task 3: 冻结 `Weekly Review -> Product Create 预填 -> Product Detail` 的浏览器验收矩阵
  - [x] SubTask 3.1: 为 `Weekly Review` 入口定义模板候选、active candidate 与派生提示的成功/空态/失败态判定标准
  - [x] SubTask 3.2: 为 `Product Create` 定义 `templateCandidateId` 进入、字段级预填可编辑性与创建成功判定标准
  - [x] SubTask 3.3: 为 `Product Detail` 定义模板来源摘要、候选 `Module` 组合摘要与 canonical binding CTA 的结果页断言

- [x] Task 4: 冻结 `Dashboard / Review / Product Detail / ReuseSummary` 的最小反回归与 reread 断言
  - [x] SubTask 4.1: 明确四个页面的最小反回归页面清单、进入方式与成功判定
  - [x] SubTask 4.2: 明确 `Dashboard / Review / ReuseSummary` 的成功 reread 观察口径，禁止把“无统计变化”误判为失败
  - [x] SubTask 4.3: 明确页面崩溃、来源链丢失、reread 语义漂移与 owner 越界四类阻断问题的判定规则

- [x] Task 5: 冻结本阶段边界证据与正式验收结论结构
  - [x] SubTask 5.1: 明确本阶段如何记录“不做 `dry-run / AI Context Enhancement / Venture`”的边界证据
  - [x] SubTask 5.2: 明确正式验收结论必须包含环境、步骤、结果、问题、复测与 DoD 判定
  - [x] SubTask 5.3: 明确 `phase09-11` 的统一结论只作为后续收口上游，不提前混入下一阶段正文

- [x] Task 6: 完成 `phase09-11` 规格一致性校验
  - [x] SubTask 6.1: 校验本规格与 `phase09-08 / 09 / 10` 的冻结口径一致
  - [x] SubTask 6.2: 校验本规格未把未来能力写成当前阶段交付
  - [x] SubTask 6.3: 校验本规格未把局部实现结论误写成统一最终验收结论

## 本轮验收结果

- 工具链验证：`proto / backend / frontend type-check / frontend build` 全部通过
- API smoke：模板候选读取、模板预填详情读取、派生提示读取全部通过
- 浏览器闭环：`Weekly Review -> Product Create 预填 -> Product Detail` 已真实走通
- 派生提示动作：`Weekly Review -> Module Registry -> 返回 Weekly Review` 已真实走通
- 最小反回归：`Dashboard / Review / Product Detail / ReuseSummary` reread 均通过
- 正式留档：见 `acceptance_report.md`

# Task Dependencies

- `Task 2` depends on `Task 1`
- `Task 3` depends on `Task 1`, `Task 2`
- `Task 4` depends on `Task 1`, `Task 3`
- `Task 5` depends on `Task 2`, `Task 3`, `Task 4`
- `Task 6` depends on `Task 1` through `Task 5`
