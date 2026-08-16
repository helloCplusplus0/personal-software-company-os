# Tasks

- [x] Task 1: 盘点 `phase12-10` 的上游冻结输入与关键页面现状
  - [x] SubTask 1.1: 复核 `dev_plan#L254-L268` 中 `phase12-10` 的范围与 DoD
  - [x] SubTask 1.2: 复核 `phase12-05 / 06 / 08 / 09` 中共享入口、共享语义来源、reread 与衍生消费页边界
  - [x] SubTask 1.3: 盘点 `dashboard / onboarding / reviews/daily / reviews/weekly / detail pages` 当前的共享摘要、固定入口、返回链与旧文案 surface
  - [x] SubTask 1.4: 标记哪些页面已具备正式共享承接链，哪些页面仍存在重复解释、漂移文案或不稳定入口

- [x] Task 2: 收口关键页面的共享语义摘要与固定入口接入
  - [x] SubTask 2.1: 落实 `dashboard / onboarding / reviews/daily / reviews/weekly` 的共享语义摘要或固定入口消费矩阵
  - [x] SubTask 2.2: 落实相关 detail pages 对共享摘要、固定入口或受控派生摘要的正式接入矩阵
  - [x] SubTask 2.3: 回收关键页面中仍残留的页面私有四实体解释，避免继续保留第二套解释链
  - [x] SubTask 2.4: 确认衍生消费页没有新长出结构化主锚点、第二 query owner 或第二 adapter 主线

- [x] Task 3: 收口 reread、精确失效与返回链联动
  - [x] SubTask 3.1: 盘点影响共享摘要刷新的成功回调与当前精确失效实现
  - [x] SubTask 3.2: 明确关键页面返回链中的 reread 触发点、局部降级与静态共享语义消费边界
  - [x] SubTask 3.3: 修正会导致回流后继续显示旧摘要、旧入口或旧解释的页面读路径
  - [x] SubTask 3.4: 确认没有通过全量失效、页面私有补查询或额外扫描来伪造共享摘要闭环

- [x] Task 4: 完成旧文案与不稳定入口的回收
  - [x] SubTask 4.1: 逐页回收与四实体冻结语义冲突的旧解释文案（Weekly Review "本周资产盘点与整理" → "本周经营回顾与能力整理"）
  - [x] SubTask 4.2: 逐页回收不能稳定定位、不能稳定 reread 或来源不明的入口展示
  - [x] SubTask 4.3: 确认需要保留的解释 surface 已迁回共享语义来源、固定入口或受控派生摘要
  - [x] SubTask 4.4: 记录本子任务明确不做的项：不新增写路径、不新增服务、不新增结构化锚点

- [x] Task 5: 完成验收与回归验证
  - [x] SubTask 5.1: 对照 `spec.md` 验证关键页面都能通过共享语义摘要或固定入口解释四实体角色
  - [x] SubTask 5.2: 验证返回链、成功回流与 reread 不再放大语义漂移
  - [x] SubTask 5.3: 验证 `dashboard / onboarding / reviews/daily / reviews/weekly` 仍只作为衍生消费页接入共享摘要或固定入口
  - [x] SubTask 5.4: 运行相关前端校验并记录通过证据（tsc --noEmit 通过，oxlint 0 errors / 3 pre-existing warnings）

# Task Dependencies

- Task 2 depends on Task 1
- Task 3 depends on Task 1
- Task 4 depends on Task 1
- Task 3 depends on Task 2
- Task 4 depends on Task 2
- Task 5 depends on Task 2
- Task 5 depends on Task 3
- Task 5 depends on Task 4