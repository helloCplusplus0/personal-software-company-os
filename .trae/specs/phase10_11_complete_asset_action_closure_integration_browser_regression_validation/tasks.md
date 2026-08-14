# Tasks

- [x] Task 1: 冻结 `phase10-11` 的单值验收前置数据与场景定义
  - [x] SubTask 1.1: 冻结“空态或近空态”样本，明确最小 Product / Repository / Module / Decision 数量与允许保留的历史数据范围
  - [x] SubTask 1.2: 冻结 `Product / Module / Repository Detail` 的最小 canonical 结构缺口样本，明确每条浏览器链路使用哪一个固定实体
  - [x] SubTask 1.3: 冻结 `Decision pending` 样本，明确唯一 `proposed` 样本、其来源入口与应出现的 pending 位置

- [x] Task 2: 冻结 `phase10-11` 的工具链验收矩阵
  - [x] SubTask 2.1: 明确 `buf` 相关合同验证的正式命令、执行顺序与通过标准
  - [x] SubTask 2.2: 明确 `go test` 相关后端验证的正式命令、执行顺序与通过标准
  - [x] SubTask 2.3: 明确 `frontend build` 相关前端构建验证的正式命令、执行顺序与通过标准
  - [x] SubTask 2.4: 明确工具链 warning、失败与环境异常的归类口径，不允许临场解释

- [x] Task 3: 冻结 `Onboarding` 首轮建链浏览器验收矩阵
  - [x] SubTask 3.1: 明确从 `Dashboard -> Onboarding` 的唯一正式入口与冷启动前置状态
  - [x] SubTask 3.2: 逐步冻结 `welcome -> product -> repository -> module -> decision -> complete` 的机械验收动作、预期结果与回看点
  - [x] SubTask 3.3: 冻结 `Onboarding` 每一步默认下一步动作与 canonical handoff 的验收口径
  - [x] SubTask 3.4: 冻结完成态与返回原入口后“用户能看懂刚刚发生了什么”的验收点

- [x] Task 4: 冻结 `Decision` 生命周期与 pending reread 的浏览器验收矩阵
  - [x] SubTask 4.1: 冻结从 `Dashboard / Daily Review / Current Focus` 进入冻结 pending 样本的入口顺序
  - [x] SubTask 4.2: 冻结 `Decision Detail` 从 `proposed` 推进后的页面级预期，包括 CTA 切换与返回来源
  - [x] SubTask 4.3: 冻结返回 `Dashboard / Daily Review / Current Focus` 后的 reread 回看点，确保 pending 不残留

- [x] Task 5: 冻结关键 detail pages 的浏览器验收矩阵
  - [x] SubTask 5.1: 冻结 `Product Detail` 的浏览器验收链，验证页面级主 CTA、canonical path 与返回来源后的 reread
  - [x] SubTask 5.2: 冻结 `Module Detail` 的浏览器验收链，验证页面级主 CTA、canonical handoff 与返回来源后的 reread
  - [x] SubTask 5.3: 冻结 `Repository Detail` 的浏览器验收链，验证页面级主 CTA、进入动作承接位与到 `Decision Detail` 的连续 handoff
  - [x] SubTask 5.4: 冻结 detail page 验收后的统一回看点，确保来源页能解释刚刚完成的动作

- [x] Task 6: 冻结 `Current Focus / pending signals` 的反回归验证矩阵
  - [x] SubTask 6.1: 明确关键动作完成后需要回看的 `Current Focus` 区块、主 CTA 与解释文案
  - [x] SubTask 6.2: 明确需要回看的 pending signals 区块、信号文案与跳转目标
  - [x] SubTask 6.3: 明确哪些回归现象应直接判定 `phase10-11` 未通过

- [x] Task 7: 留档当前阶段明确不做的边界证据
  - [x] SubTask 7.1: 明确 `Agent Consumption Layer` 不属于本轮机械验收范围
  - [x] SubTask 7.2: 明确新实体回归不属于本轮机械验收范围
  - [x] SubTask 7.3: 明确第五态状态机不属于本轮机械验收范围
  - [x] SubTask 7.4: 将这些边界与 `phase10-11` DoD 对齐，避免后续扩大范围

- [x] Task 8: 完成 `phase10-11` 验收规格自检
  - [x] SubTask 8.1: 复核 spec 是否已提供单值样本、单值入口、单值动作与单值预期结果
  - [x] SubTask 8.2: 复核 tasks 是否足以让后续独立验收者不再补造主测试路径
  - [x] SubTask 8.3: 复核 checklist 是否覆盖工具链、Onboarding、Decision、detail pages、反回归与边界证据

# Task Dependencies

- `Task 2` depends on `Task 1`
- `Task 3` depends on `Task 1`
- `Task 4` depends on `Task 1`
- `Task 5` depends on `Task 1`
- `Task 6` depends on `Task 4`
- `Task 6` depends on `Task 5`
- `Task 7` can run after `Task 1`
- `Task 8` depends on `Task 2`
- `Task 8` depends on `Task 3`
- `Task 8` depends on `Task 4`
- `Task 8` depends on `Task 5`
- `Task 8` depends on `Task 6`
- `Task 8` depends on `Task 7`
