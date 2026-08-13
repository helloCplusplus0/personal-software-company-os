# Fix 002 Analysis - Decision 待处理信号语义错位

## 1. 问题摘要
- 对应问题：`fix_002`
- 问题级别：`P0`
- 是否阻断修复：是

## 2. 根因结论
- 根因一：`Decision.status = proposed` 的 canonical 语义只是“已经被记录，但尚未成为当前生效结论”，但 `Dashboard` 与 `Daily Review` 当前直接把全部 `proposed` 决策解释成“待处理 / 待决策”，把“已留痕”与“待动作”混成了同一层语义。
- 根因二：当前系统不存在正式、单值的 `Decision` 状态推进出口；`LinkDecisionToTarget` 只建立关联，不推进 `status`，`SubmitReviewResult` 只记录轻量 `review_records`，也不推进 `status`，导致用户即使已经“消费过”一条决策，canonical facts 仍停留在 `proposed`。
- 根因三：`Dashboard` 与 `Review` 两条读链各自独立消费 `proposed` 语义，形成并行误报：`Dashboard` 的 `ReadPendingDecisions` 直接查 `decisions.status = 'proposed'`，`Daily Review` 又直接读取 `ListDecisions(status = proposed)`，把同一个语义错位复制到了多个入口。

## 3. 证据链
- 页面预展示链路：
  - `backend/internal/dashboard/candidate/feedback_readers.go:45-63`
  - `Dashboard` 的 pending decision reader 直接以 `WHERE d.status = 'proposed'` 读取，并仅用 `HasDecisionLink` 区分跳详情还是跳列表，不把“已关联”视为退出 pending。
  - `backend/internal/review/service/query_service.go:64-93`
  - `Daily Review` 直接把 `decisioncenter.QueryService.ListDecisions(status = proposed)` 作为 `pending_decisions` 来源。
- 服务端实际创建 / 状态承接链路：
  - `backend/internal/decisioncenter/service/command_service.go:57-64`
  - `Decision` 创建时，未指定状态默认补为 `proposed`。
  - `backend/internal/decisioncenter/service/command_service.go:111-152`
  - `LinkDecisionToTarget` 只写 `decision_links`，不更新 `Decision.status`。
- review 结果写回链路：
  - `frontend/src/features/review/application/use-review-action.ts:139-150`
  - `submit_next_step` 只调用 `reviewClient.submitReviewResult`。
  - `backend/internal/review/types.go:30-44`
  - `backend/internal/review/repository/review_record_store.go:27-74`
  - `review_records` 只承接轻量过程记录，不是 `Decision` canonical 状态字段。
- 详情承接位现状：
  - `frontend/src/features/decision-center/pages/decision-detail-page.tsx:116-168`
  - 当前详情页只承接来源上下文、已关联目标与候选关联，不存在正式状态推进 CTA。
- 规格冻结证据：
  - `phase03_02_decision_template_status_read_model/spec.md:74-82`
    - `proposed` 的冻结语义是“已经被记录，但尚未成为当前生效结论”。
  - `phase05_12_dashboard_feedback_backend_data_mainline/spec.md:179-184`
    - `pending_decision_signals` 当前明确沿用 `proposed` 语义。
  - `phase08_07_design_review_backend_contract_service_data_handoff/spec.md:80-84`
    - `Daily Review` 当前也明确以 `ListDecisions(status = proposed)` 作为 pending 来源。
  - `phase08_10_land_feedback_decision_update_closed_loop_result_writeback_cleanup/spec.md:136-160`
    - `SubmitReviewResult` 被明确冻结为轻量过程记录，不得代替实体 canonical update。
- 数据模型映射：
  - `proto/psco/decision_center/v1/decision_center.proto`
  - 当前正式枚举只有 `proposed / active / superseded / archived` 四态。
- 是否存在历史脏数据：
  - 当前判断：存在**潜在语义脏数据**，但不能自动安全判定。
  - 原因：历史上已被用户“处理过”的 `Decision` 可能仍停留在 `proposed`，但仅凭 `decision_links` 或 `review_records` 无法可靠推断它应变成 `active / superseded / archived` 中哪一种。

## 4. 影响面分析
- Dashboard / Current Focus：
  - 直接受影响。
  - 当前主行动队列会持续把部分已被消费的 `Decision` 作为 `pending_decision` 信号抬到 P1。
- Daily Review：
  - 直接受影响。
  - `pending_decisions` 列表会重复呈现这些 `proposed` 决策，形成“我已经处理过但仍被催促”的体感。
- Decision Detail：
  - 间接受影响。
  - 因为详情页没有正式状态推进出口，用户无法通过 canonical 行为让决策退出 pending。
- Review Loop：
  - 间接受影响。
  - `submit_next_step` 会形成“这次 review 有后续动作”的证据，但不会真正改变 `Decision` 的 pending 语义。
- 历史数据：
  - 可能受影响。
  - 存量 `proposed` 决策中，部分实际已不应继续被当作 pending；但当前没有安全自动迁移依据。

## 5. 候选方案对比
### 方案 A
- 做法：
  - 不改 `Decision` canonical 生命周期；
  - 仅在 `Dashboard` / `Review` 的 pending 派生中增加额外排除条件，例如：
    - 已存在 `decision_links` 的 `Decision` 不再算 pending；
    - 或已存在 `review_records(decision_id=...)` 的 `Decision` 不再算 pending。
- 优点：
  - 表面上改动最小；
  - 可以较快降低当前 UI 误报数量。
- 风险：
  - 把 `decision_links` 或 `review_records` 偷偷升级成“已处理”判定依据，形成第二套 canonical 语义；
  - 已关联目标并不等于“当前结论已经正式确认生效”；
  - `review_records` 按规格只是一条过程记录，明确不能被解释为实体更新完成；
  - 会让 `Dashboard / Review` 与 `Decision Detail` 的状态解释继续分叉，属于补丁式修复。

### 方案 B
- 做法：
  - 保持当前正式状态枚举仍为 `proposed / active / superseded / archived`；
  - 不把 `decision_links` 或 `review_records` 直接当作退出 pending 的代理条件；
  - 将 pending 的正式判定继续锚定在 canonical `Decision.status` 上；
  - 同时要求通过 `fix_003` 补齐正式状态推进入口，让用户能把已处理决策从 `proposed` 推进到合适的非 pending 状态；
  - 在修复实现中把 `Dashboard` 与 `Review` 的 pending 判定继续统一建立在同一套 canonical status 规则上，不各自长临时补丁。
- 优点：
  - 符合当前 phase03/05/08 已冻结的合同与状态模型；
  - 不会把过程记录、关联记录误升级成第二事实源；
  - 能把 `fix_002` 与 `fix_003` 串成真正的闭环：`fix_002` 解决“pending 应基于什么语义”，`fix_003` 解决“用户如何正式退出 pending”。
- 风险：
  - 不能靠单条 SQL 立刻彻底消除全部误报，必须与 `fix_003` 联动才能完整收口；
  - 对历史遗留的 `proposed` 决策，仍需要用户在新详情页中手动推进状态，不能盲目自动回填。

### 方案 C
- 做法：
  - 在当前枚举之外新增中间状态，例如 `acknowledged` / `resolved`，并用它区分“已留痕但未处理”“已确认但未生效”等语义。
- 优点：
  - 语义表达更细。
- 风险：
  - 需要同时改 `.proto`、Go domain、前端枚举、数据库约束、列表筛选、seed 与验收；
  - 超出本轮 fix 的最小修复边界；
  - 与现有 `phase03` 已冻结的四态状态集冲突，除非重新开一轮更大范围 spec，不适合当前 fix workflow。

## 6. 推荐方案
- 推荐原因：
  - 推荐 **方案 B**。
  - 当前问题的真正根因不是 reader 条件写得太宽，而是系统还没有一条正式、单值的 `Decision` 退出 pending 主线。
  - 只在 `Dashboard` / `Review` 上加排除条件会把 `decision_links` 或 `review_records` 偷偷变成第二事实源，反而会让状态语义更乱。
  - 更稳妥的做法是：
    1. 继续把 pending 判定锚定在 canonical `Decision.status`；
    2. 由 `fix_003` 补齐正式状态推进 CTA；
    3. 状态一旦从 `proposed` 退出，`Dashboard / Daily Review / Current Focus` 自然不再误报。
- 实施边界：
  - 本 fix 的正式结论应冻结为：`pending decision` 的 canonical 退出条件不能来自 `decision_links` 或 `review_records`，而必须来自 `Decision.status` 的正式推进；
  - 允许在实现层回收 / 统一 pending reader 的承接位，避免 Dashboard 与 Review 各自复制补丁逻辑；
  - 允许对前端 pending 文案、badge 与提示语做最小对齐，但不得创造第二套状态解释。
- 明确不在本次修复范围内的内容：
  - 不在 `fix_002` 中新增第五个 `DecisionStatus` 枚举；
  - 不在 `fix_002` 中把 `review_records` 升格为实体状态来源；
  - 不在 `fix_002` 中把 `decision_links` 直接解释为“已解决”；
  - 不在 `fix_002` 中单独重写 Dashboard / Review 的临时本地状态机；
  - 不在 `fix_002` 中处理前端聊天式工作台、Agent 写入或其他无关扩展。

## 7. 数据修复策略
- 是否需要修历史数据：需要，但仅允许**人工确认后最小修复**，不允许自动批量推断。
- 若需要，修复范围：
  - 历史 `proposed` 决策中，用户确认“实际上已生效 / 已替代 / 已归档”的条目；
  - 修复方式应在 `fix_003` 提供正式状态推进后，通过 canonical 写路径逐条修正。
- 若不需要自动修复，原因：
  - 仅凭 `decision_links`、`review_records`、是否进入过详情页，无法可靠判断一条 `Decision` 应转为 `active / superseded / archived` 中哪一种；
  - 自动回填会制造新的语义脏数据。

## 8. 验收标准
- `Dashboard / Current Focus / Daily Review` 不得再把“已被正式确认退出 pending”的 `Decision` 继续误报为待处理。
- `pending decision` 的判定规则必须仍然建立在单一 canonical `Decision.status` 语义上，而不是改由 `decision_links` 或 `review_records` 代理。
- `SubmitReviewResult` 的职责边界必须保持不变，仍然只记录轻量 `review record`。
- `LinkDecisionToTarget` 的职责边界必须保持清晰：它只建立关联，不得被偷偷解释为状态完成。
- fix 实现完成后，`fix_003` 必须能提供用户真正退出 pending 的正式动作入口；二者联动后，Dashboard / Review reread 结果应与 Decision Detail 状态一致。

## 9. 回滚条件
- 若修复后 `Dashboard` 或 `Daily Review` 通过 `decision_links` / `review_records` 直接隐藏仍处于 `proposed` 的 `Decision`，则必须回滚；
- 若修复后为了消除误报而引入新的隐式状态、第二套 pending 字段或页面局部“已处理”假象，则必须回滚。
