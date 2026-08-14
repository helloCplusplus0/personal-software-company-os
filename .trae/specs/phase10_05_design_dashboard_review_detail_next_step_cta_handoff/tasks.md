# Tasks

- [x] Task 1: 冻结 `Dashboard Home / Daily Review` 的下一步动作承接矩阵
  - [x] SubTask 1.1: 明确 `Dashboard` 在空态、pending decision、结构缺口三类场景下的主 CTA 与次 CTA
  - [x] SubTask 1.2: 明确 `Dashboard / Daily Review / Current Focus` 在结构缺口场景下的单值默认跳转目标映射，不再保留 “detail 或 binding owner” 的并列默认目标
  - [x] SubTask 1.3: 明确 `Dashboard / Daily Review` 成功回流后的 reread 观察点与主 CTA 切换规则
  - [x] SubTask 1.4: 修复 `Dashboard` 在“结构缺口 + pending decision”并存场景下的主 CTA 优先级冲突，确保不需要执行者自行判断

- [x] Task 2: 冻结 `Product Detail / Module Detail / Repository Detail` 的 CTA inventory
  - [x] SubTask 2.1: 明确 `Product Detail` 允许的下一步动作类型、触发条件、默认跳转目标与 reread 落点
  - [x] SubTask 2.2: 明确 `Module Detail` 允许的下一步动作类型、触发条件、默认跳转目标与 reread 落点
  - [x] SubTask 2.3: 明确 `Repository Detail` 允许的下一步动作类型、触发条件、默认跳转目标与 reread 落点
  - [x] SubTask 2.4: 为三类 Detail 页分别冻结结构缺口内部排序，避免“当前最核心结构缺口”继续依赖人工判断
  - [x] SubTask 2.5: 为三类 Detail 页补齐结构 CTA 成功回流后的页面级 reread 结果，能机械回答“回来看什么”

- [x] Task 3: 冻结 `Decision Detail` 与 `Current Focus` 的 canonical 承接关系
  - [x] SubTask 3.1: 明确 `Decision Detail` 在 `proposed / active / superseded / archived` 四态下的页面级主 CTA / 次 CTA 规则
  - [x] SubTask 3.2: 明确 `Current Focus` 命中 `Decision` 或结构缺口时必须 handoff 到哪个 canonical owner
  - [x] SubTask 3.3: 明确 `Decision Detail` 成功回流后对 `Dashboard / Daily Review` 的 reread 影响

- [x] Task 4: 冻结跨页面主 CTA / 次 CTA 的统一优先级
  - [x] SubTask 4.1: 明确多个 CTA 同时成立时的统一优先级：结构缺口优先、正式状态推进次之、reread 返回兜底
  - [x] SubTask 4.2: 明确次 CTA 的并列边界，禁止与主 CTA 的目标 owner 冲突
  - [x] SubTask 4.3: 明确哪些场景允许把 reread 返回 CTA 升为兜底主动作

- [x] Task 5: 完成 `phase10-05` 三件套自检并对齐上游
  - [x] SubTask 5.1: 复核 `spec.md` 是否正确承接 `phase10-03 / phase10-04 / shared_baseline` 的冻结结论
  - [x] SubTask 5.2: 复核 `tasks.md` 与 `checklist.md` 是否完整覆盖 Dashboard、Review、三类 Detail、Current Focus、单值跳转目标映射、结构内部排序与优先级矩阵
  - [x] SubTask 5.3: 确认文档已能机械回答“何时显示什么 CTA、点完去哪、回来看什么”

# Task Dependencies

- `Task 2` depends on `Task 1`
- `Task 3` depends on `Task 1`
- `Task 4` depends on `Task 2` and `Task 3`
- `Task 5` depends on `Task 1` to `Task 4`
