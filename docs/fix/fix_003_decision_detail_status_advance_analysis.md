# Fix 003 Analysis - Decision Detail 缺少正式状态推进入口

## 1. 问题摘要
- 对应问题：`fix_003`
- 问题级别：`P0`
- 是否阻断修复：是

## 2. 根因结论
- 根因一：当前 `Decision Detail` 的前端与路由职责仍停留在 `详情读取 + 已关联目标展示 + 候选读取 + LinkDecisionToTarget`，并未承接任何 `Decision.status` 正式写入动作，因此用户进入详情页后只能“查看 / 关联”，不能“正式处理完成”。
- 根因二：`Decision Center` 的正式合同与后端写链当前只有 `CreateDecision` 与 `LinkDecisionToTarget` 两条写路径，没有 `UpdateDecisionStatus` 一类 canonical 写接口；这不是单纯漏放一个按钮，而是前后端整条状态推进写链都不存在。
- 根因三：`Review` 与 `Dashboard` 当前都把 `Decision Detail` 当作 pending 决策的 canonical 消费宿主，但详情页缺少正式退出 pending 的动作，导致 `fix_002` 中确认的语义问题无法真正闭环。

## 3. 证据链
- 页面预展示链路：
  - `frontend/src/features/decision-center/pages/decision-detail-page.tsx:124-168`
  - 当前页面只渲染：
    - `DecisionDetailSummaryCard`
    - `DecisionLinkedTargetsSection`
    - `DecisionPendingLinkTargetCard`
    - `DecisionModuleCandidatePanel`
  - 页面中不存在状态推进区、状态写入 CTA 或对应 mutation owner。
- 详情概要区现状：
  - `frontend/src/features/decision-center/components/decision-detail-summary-card.tsx:34-97`
  - 目前只展示状态 badge 与静态字段，没有任何“标记为 active / superseded / archived”之类的正式动作入口。
- 前端现有写路径边界：
  - `frontend/src/features/decision-center/application/use-link-decision-to-target.ts:10-24`
  - `frontend/src/features/decision-center/components/decision-module-candidate-panel.tsx:56-73`
  - 当前详情页唯一正式 mutation owner 是 `useLinkDecisionToTarget`，且成功后只失效 `decision-detail` 与 `decision-module-candidates`，页面包装层再额外失效 `decision-list`。
  - 该写路径只承接目标关联，不承接状态推进。
- Daily Review handoff 现状：
  - `frontend/src/features/review/pages/daily-review-page.tsx:97-112`
  - `frontend/src/features/review/pages/daily-review-page.tsx:227-244`
  - `pending decisions` 当前会把用户送入既有 `Decision Detail` canonical 路径，并透传 Dashboard 来源参数。
  - 但进入详情页后没有正式处理出口，用户只能查看或做目标关联，无法让 pending 语义真正退出。
- Review action owner 边界：
  - `frontend/src/features/review/application/use-review-action.ts:41-109`
  - `frontend/src/features/review/application/use-review-action.ts:139-163`
  - `go_to_decision` 只是 canonical handoff；`submit_next_step` 只写 review result，并统一失效 review/dashboard query。
  - 这再次证明“状态推进”必须在 `Decision Detail` canonical owner 内完成，而不是回落到 review 页。
- 后端正式写链现状：
  - `backend/internal/decisioncenter/connect/server.go:118-148`
  - `backend/internal/decisioncenter/service/command_service.go:34-109`
  - `backend/internal/decisioncenter/service/command_service.go:111-153`
  - Connect server 与 CommandService 当前只承接：
    - `CreateDecision`
    - `LinkDecisionToTarget`
  - 不存在任何状态更新 handler / service 方法。
- repository 能力现状：
  - `backend/internal/decisioncenter/repository/decision_store.go:41-74`
  - `backend/internal/decisioncenter/repository/decision_store.go:105-181`
  - `DecisionStore` 当前只有 `Create / GetByID / List / Exists`，没有 `UpdateStatus` 一类持久化入口。
- 正式合同现状：
  - `proto/psco/decision_center/v1/decision_center.proto:230-243`
  - `proto/psco/decision_center/v1/decision_center.proto:249-290`
  - 当前 `.proto` 写组只有 `CreateDecision` 与 `LinkDecisionToTarget`，没有正式状态推进 RPC。
- 规格冻结证据：
  - `phase03_02_decision_template_status_read_model/spec.md:74-91`
    - 已冻结四态语义：
      - `proposed`：已记录但未成为当前生效结论
      - `active`：当前仍然生效或正在执行
      - `superseded`：曾经生效但已被替代
      - `archived`：已归档保留，不再作为当前执行结论
  - `phase03_01_decision_center_pages_info_arch/spec.md:60-69`
    - `Decision Detail` 被冻结为最小承接页，但不得扩写成新的跨对象复合工作台。
  - `phase03_05_frontend_page_route_component_design/spec.md:202-210`
    - `Decision Detail` 必须继续在同一详情页壳层中整合能力，不得拆出第二套子工作台。
  - `phase08_03_freeze_feedback_decision_update_action_owner/spec.md:74-83`
    - review 与 dashboard 相关正式决策动作只能进入既有 `Decision Detail`，不得在 review 内联决策编辑工作台。
  - `phase08_10_land_feedback_decision_update_closed_loop_result_writeback_cleanup/spec.md:71-76`
    - review 消费决策信号后，若需形成正式经营判断，必须进入既有 `Decision Detail` 或 `Decision Create`，而不是把进入详情页本身视为闭环完成。
- 路由合同现状：
  - `frontend/src/routes/decisions/$decisionId.tsx:8-28`
  - 当前详情页搜索参数只承接 `fromList + dashboard 来源 + onboarding 来源`，没有 review-local 返回链。
  - 说明本次 fix 不应新长一套 review-local 状态推进通道，而应继续复用 canonical detail 路径与既有 Dashboard 返回语义。
- 是否存在历史脏数据：
  - 当前判断：存在**历史待修复语义积压**，但不属于自动数据迁移问题。
  - 原因：历史上停留在 `proposed` 的 `Decision` 并非都错，只是缺少正式推进出口；修复后应允许用户通过 canonical 详情页逐条推进，而不是自动批量改写。

## 4. 影响面分析
- Decision Detail：
  - 直接受影响。
  - 当前详情页没有“确认生效 / 已被替代 / 归档退出”动作，导致它无法承接当前版本的正式闭环。
- Dashboard / Current Focus：
  - 间接受影响。
  - pending 决策从 Dashboard 进入详情页后，没有 canonical 退出动作，用户返回后仍看到同一条信号。
- Daily Review：
  - 间接受影响。
  - pending decision 行能把用户送入详情页，但详情页内无法正式完成处理，形成“看到了 canonical 页面，但没法真正做完”的断裂体验。
- Review Loop：
  - 间接受影响。
  - `submit_next_step` 只记录 review 结果，不应承担状态推进；没有 detail 侧写入，review loop 只能停在“留痕已完成、决策未完成”的半闭环。
- 历史 `proposed` 决策：
  - 受影响。
  - 修复后需要通过 canonical 详情页逐条处理真实历史决策，才能让 `fix_002` 的误报逐步消失。

## 5. 候选方案对比
### 方案 A
- 做法：
  - 只在前端 `Decision Detail` 增加本地按钮或局部 UI 状态；
  - 点击后仅本地隐藏“待处理”提示，或依赖 `review_records` / `decision_links` 营造“已处理”假象；
  - 不新增后端正式写接口。
- 优点：
  - 表面改动最小；
  - UI 上能很快出现一个“完成”动作。
- 风险：
  - 没有 canonical 后端写入，详情页刷新后状态不会稳定存在；
  - 会把本地 UI 状态、`review_records` 或 `decision_links` 偷偷升级成第二事实源；
  - 无法真正解决 `fix_002` 的 pending 语义错位；
  - 明显违背“`.proto` 为唯一长期合同源、Go 后端承载核心能力”的项目基线。

### 方案 B
- 做法：
  - 在现有四态内新增正式状态推进写链：
    - `.proto` 新增 `UpdateDecisionStatus` 一类写接口
    - Connect server / CommandService / DecisionStore 新增对应承接位
    - 前端在 `Decision Detail` 内新增固定 mutation owner 与最小 CTA 组
  - CTA 只围绕现有四态承接最小闭环，不新增第五态；
  - 最小推进矩阵建议冻结为：
    - `proposed -> active`
    - `proposed -> superseded`
    - `proposed -> archived`
    - `active -> superseded`
    - `active -> archived`
    - `superseded / archived` 视为终态，不再提供继续推进 CTA
  - 写入成功后统一失效：
    - `decision-detail`
    - `decision-list`
    - `daily review / weekly review`
    - `dashboard-feedback-signals / dashboard-overview / dashboard-recent-activities`
- 优点：
  - 保持 canonical 状态仍锚定在既有四态，不引入新枚举；
  - 真正补齐 `Decision Detail` 作为正式处理承接位的闭环；
  - 与 `fix_002` 的推荐方案天然联动：pending 继续锚定 `Decision.status`，用户则通过 detail 正式退出 pending；
  - 符合“后端优先承载核心能力、前端只是消费与操作渠道”的项目方向。
- 风险：
  - 涉及 `.proto + Go + 前端` 联动修改，范围比纯 UI 按钮大；
  - 需要明确状态迁移矩阵，避免开放过宽导致语义飘移；
  - 需要小心保持 `Decision Detail` 仍是详情页，而不是新的决策工作台。

### 方案 C
- 做法：
  - 直接扩展状态机，引入 `acknowledged / resolved` 等中间态，再围绕这些中间态设计更完整的审批/确认流程。
- 优点：
  - 语义层次更细，看起来更贴近“待确认 / 已确认 / 已生效”的理想模型。
- 风险：
  - 需要改 `.proto` 枚举、数据库约束、domain types、前后端映射、筛选器与验收基线；
  - 超出当前 fix workflow 的最小修复边界；
  - 与 `phase03` 已冻结的四态集合冲突，需要新的结构化 spec，而不是 fix 级修复。

## 6. 推荐方案
- 推荐原因：
  - 推荐 **方案 B**。
  - 当前问题的本质不是“详情页视觉上少了个 CTA”，而是系统根本没有一条正式状态推进写链。
  - 只做前端 UI 补丁无法形成 canonical 闭环，也无法兑现 `fix_002` 对 pending 语义的修复前提。
  - 在既有四态内补一条最小状态推进主线，是当前阶段最小、最稳、也最符合项目基线的做法。
- 实施边界：
  - 状态推进必须走 `.proto -> Connect -> CommandService -> DecisionStore` 的正式后端写链；
  - 前端只能在既有 `Decision Detail` 壳层中承接状态推进 CTA，不得拆出第二套子工作台；
  - CTA 只围绕既有四态承接，不引入第五态；
  - 状态推进成功后，必须统一触发 detail / list / review / dashboard 的 reread，避免“详情页更新了、Dashboard 还在催促”的假闭环。
- 明确不在本次修复范围内的内容：
  - 不在 `fix_003` 中新增 `acknowledged / resolved` 等新状态；
  - 不在 `fix_003` 中把 `Review` 页面扩写成内联决策编辑工作台；
  - 不在 `fix_003` 中把 `decision_links` 或 `review_records` 直接解释为状态完成；
  - 不在 `fix_003` 中重写 `Decision Detail` 为跨对象复合工作台；
  - 不在 `fix_003` 中引入前端聊天式交互、Agent 写入或其他无关扩展。

## 7. 数据修复策略
- 是否需要修历史数据：需要，但仅允许**通过新的 canonical 写路径逐条修正**，不允许自动批量迁移。
- 若需要，修复范围：
  - 历史上仍停留在 `proposed`、但用户确认实际上已经生效 / 已被替代 / 应归档的 `Decision`。
- 若不需要自动修复，原因：
  - 当前缺的是正式推进动作，不是数据库损坏；
  - 自动批量把 `proposed` 改成某一终态没有可靠依据，容易制造新的语义脏数据。

## 8. 验收标准
- `Decision Detail` 必须提供正式、可执行的状态推进入口，而不是只有静态状态 badge。
- 状态推进必须通过单一 canonical 后端写链完成，不得只停留在前端局部状态。
- `fix_003` 完成后，用户从 `Daily Review / Dashboard` 进入某条 pending `Decision`，可以在详情页中把它正式推进到非 pending 状态。
- 状态推进成功后，`Decision Detail`、`Decision List`、`Dashboard / Current Focus`、`Daily Review` 的 reread 结果必须保持一致。
- `SubmitReviewResult` 与 `LinkDecisionToTarget` 的职责边界必须保持不变：
  - 前者仍只记录 review 结果；
  - 后者仍只承接目标关联；
  - 二者不得被偷渡解释为状态推进。

## 9. 回滚条件
- 若修复后状态推进只发生在前端局部 UI、刷新后丢失，必须回滚；
- 若修复后把 `Decision Detail` 扩写成新的跨对象编辑工作台、第二套 review-local decision owner 或第二状态机，必须回滚；
- 若修复后 `Dashboard / Daily Review` 与 `Decision Detail` 对同一条决策的状态解释再次不一致，必须回滚。
