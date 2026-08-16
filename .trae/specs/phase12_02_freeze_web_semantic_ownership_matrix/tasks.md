# Tasks

- [x] Task 1: 盘点 `phase12-02` 直接上游中的四实体语义与 Web 承接口径
  - [x] SubTask 1.1: 审阅 `phase12_semantic_alignment_and_readonly_consumption_foundation_dev_plan.md#L44-L58` 中 `phase12-02` 的范围与 DoD
  - [x] SubTask 1.2: 审阅 `architecture_plan` 中四实体语义边界、前端承接策略与默认非 primary owner 口径
  - [x] SubTask 1.3: 审阅 `shared_baseline` 中四实体语义矩阵、primary owner / 跟随回归清单与设计模板

- [x] Task 2: 冻结四实体在 Web 端的正式解释口径
  - [x] SubTask 2.1: 将 `Product / Repository / Module / Decision` 的正式 Web 解释写成单值口径
  - [x] SubTask 2.2: 明确 `Module` 与 `Decision` 在 Web 端不得再回落为旧的弱语义解释
  - [x] SubTask 2.3: 明确当前阶段只承接表达层与消费层对齐，不承接结构重构

- [x] Task 3: 冻结 Web 端语义承接矩阵
  - [x] SubTask 3.1: 固定四类详情页为四实体语义的 primary owner 页面
  - [x] SubTask 3.2: 固定四类摘要卡片为 primary owner 摘要组件
  - [x] SubTask 3.3: 固定 `dashboard / onboarding / daily review / weekly review` 页面及其容器组件为跟随回归对象
  - [x] SubTask 3.4: 固定 route shell、list pages、toolbar 与搜索 store 不作为独立语义 primary owner

- [x] Task 4: 冻结共享承接位与禁止散落重复解释的边界
  - [x] SubTask 4.1: 明确切片页面、切片组件与 `frontend/src/features/project-context/` 的语义承接边界
  - [x] SubTask 4.2: 明确只有 `3+` 页面稳定复用时才允许晋升到跨切片共享只读承接位
  - [x] SubTask 4.3: 明确不得在页面本地拼出第二套跨切片语义摘要或第二套事实源

- [x] Task 5: 冻结 `no-change`、`follow-regression` 与审计留档规则
  - [x] SubTask 5.1: 明确所有进入 `phase12-04` 影响面的对象都必须被分类为 `must-change / follow-regression / no-change`
  - [x] SubTask 5.2: 明确 `no-change` 必须附带"不改仍满足冻结口径"的理由
  - [x] SubTask 5.3: 明确 `follow-regression` 不是跳过审计，而是跟随 primary owner 或共享承接位回归

- [x] Task 5a: 冻结空态、提示文案与下一步动作说明的显式审计面
  - [x] SubTask 5a.1: 明确空态（List 页空态、Dashboard 无数据空态、Review 无待处理项空态、Onboarding 各步骤空态）为独立审计面，必须逐项标记 `must-change / follow-regression / no-change`
  - [x] SubTask 5a.2: 明确提示文案（Onboarding WelcomeStep、Dashboard 区块说明、Review 头部说明、Detail 页描述性标题）为独立审计面，必须逐项标记并留档
  - [x] SubTask 5a.3: 明确下一步动作说明（ModuleNextActionBar、Decision CTA、DashboardPrimaryActionPanel、OnboardingCtaButton、ReviewActionFooter）为独立审计面，必须逐项标记并留档
  - [x] SubTask 5a.4: 明确不允许用"页面已经在 owner 矩阵里了"跳过对这三类 surface 的独立审计

- [x] Task 6: 完成三件套一致性校验
  - [x] SubTask 6.1: 校验 `dev_plan`、`architecture_plan` 与 `shared_baseline` 对四实体正式语义的表达单值一致
  - [x] SubTask 6.2: 校验三件套对 primary owner、跟随回归与默认非 primary owner 的表达单值一致
  - [x] SubTask 6.3: 校验三件套对共享承接位与 `no-change` 留档规则的表达单值一致

- [x] Task 7: 校验本 spec 包可直接作为 `phase12-04` 的输入前提
  - [x] SubTask 7.1: 确认 spec 已回答“哪些页面必须展示什么语义”
  - [x] SubTask 7.2: 确认 spec 已回答“哪些页面是 primary owner，哪些只是跟随回归”
  - [x] SubTask 7.3: 确认 spec 没有把 `phase12-03` 的共享只读 owner、后端合同或更重通道边界偷渡为本任务新增主范围

# Task Dependencies

- Task 2 depends on Task 1
- Task 3 depends on Task 1
- Task 4 depends on Task 1
- Task 5 depends on Task 1
- Task 5a depends on Task 1
- Task 6 depends on Task 2
- Task 6 depends on Task 3
- Task 6 depends on Task 4
- Task 6 depends on Task 5
- Task 6 depends on Task 5a
- Task 7 depends on Task 6
