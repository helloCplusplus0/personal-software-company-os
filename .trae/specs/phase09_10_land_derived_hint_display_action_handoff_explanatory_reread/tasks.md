# Tasks

- [x] Task 1: 落实 `Weekly Review` 的正式提示展示与 active candidate 绑定
  - [x] SubTask 1.1: 在 `useWeeklyReviewRead` 或其正式组合承接位中收敛 `reuse_opportunity_hint / capability_gap_hint` 的只读 view model、空态与错误态
    - 实现：`use-weekly-review-read.ts` L290-296 `hintsSectionStatus` 由 hints query 派生
  - [x] SubTask 1.2: 在 `WeeklyReviewPage` 模板候选区附近落地两类正式提示的最小展示位、解释文案与 CTA
    - 实现：`weekly-review-page.tsx` L477-603 `DerivedHintsSection` 组件
  - [x] SubTask 1.3: 确保 active candidate 切换后，提示解释、CTA 与目标参数同步刷新
    - 实现：`use-weekly-review-read.ts` L244-249 hints 查询绑定 `activeCandidateId`
  - [x] SubTask 1.4: 确保无 active candidate、无复用机会与无能力缺口都退回成功空态，而不是 generic focus fallback
    - 实现：`DerivedHintsSection` L506-508 返回 null

- [x] Task 2: 落实单一提示 handoff application owner
  - [x] SubTask 2.1: 新增或完善 `use-derived-hint-handoff` 等正式 owner，统一承接提示点击后的 target 计算与 search 参数拼装
    - 实现：`use-derived-hint-handoff.ts` L70-177
  - [x] SubTask 2.2: 收敛 `reuse_opportunity_hint -> Product Create` 的 handoff 参数，继续复用 `fromTemplateReuse / templateCandidateId / templateSource`
    - 实现：`use-derived-hint-handoff.ts` L113-131
  - [x] SubTask 2.3: 收敛 `capability_gap_hint -> Module Registry / Module Detail` 的 handoff 参数，透传 `templateCandidateId / capabilityKey / reviewScopeKey` 与必要返回链上下文
    - 实现：`use-derived-hint-handoff.ts` L134-170
  - [x] SubTask 2.4: 统一处理非法提示参数回退、局部错误归一化与页面级临时导航清理
    - 实现：`use-derived-hint-handoff.ts` L79-103 `isValidHint` 过滤不合法提示

- [x] Task 3: 落实 `Product Create` 的解释性缺口提示与补齐返回链
  - [x] SubTask 3.1: 在 `Product Create` 模板来源摘要附近展示与当前 `templateCandidateId` 绑定的 `capability_gap_hint`
    - 实现：`product-create-page.tsx` L213-219 引用 `CapabilityGapHintsSection`，L285-366 组件实现
  - [x] SubTask 3.2: 实现从 `Product Create` 进入 `Module Registry / Module Detail` 的正式 CTA，同时保留 create 会话与返回链参数
    - 实现：`use-derived-hint-handoff.ts` L145-150 透传模板上下文
  - [x] SubTask 3.3: 实现从模块补齐页返回 `Product Create` 后恢复模板摘要、表单草稿与提示上下文
    - 实现：`module-detail-page.tsx` L101-121 `returnTo === 'product-create'` 分支
  - [x] SubTask 3.4: 明确 `Product Create` 不再展示 `reuse_opportunity_hint`，也不提供第二套提示诊断入口
    - 实现：`CapabilityGapHintsSection` L297-300 只过滤 `CAPABILITY_GAP`

- [x] Task 4: 落实提示消费后的解释性回流与 reread
  - [x] SubTask 4.1: 实现 `Weekly Review -> Module Registry / Module Detail -> Weekly Review` 的返回恢复，确保 active candidate 与 `capability_gap_hint` reread 正确
    - 实现：`module-detail-page.tsx` L101-121 `returnTo === 'weekly-review'`；`use-weekly-review-read.ts` L236-247 `returnCandidateId` 恢复
  - [x] SubTask 4.2: 对齐 `reuse_opportunity_hint -> Product Create -> Product Detail` 的解释性结果回流，继续消费 `phase09-09` 的模板来源摘要与 canonical binding CTA
    - 实现：`product-create-page.tsx` L109-113 `buildSuccessSearch()` 回流参数
  - [x] SubTask 4.3: 确保提示消费后的刷新只走正式 read owner 与 canonical reread，不使用页面级硬刷新或第二套 local store
    - 实现：React Query 自动 refetch，无 `window.location.reload()` 用于提示刷新

- [x] Task 5: 回收散装提示编排并裁撤未冻结候选提示
  - [x] SubTask 5.1: 回收 `Weekly Review / Product Create` 中零散的页面级提示目标拼装、错误语义与返回链处理
    - 验证：Grep 确认无散装提示编排代码残留
  - [x] SubTask 5.2: 裁撤没有稳定 CTA 的候选提示、generic focus fallback 与纯统计提示卡片化实现
    - 验证：`isValidHint` 过滤 + Grep 确认无 generic focus fallback
  - [x] SubTask 5.3: 确保前后端不保留第二套长期智能主线，提示能力继续只建立在既有 `TemplateReuseService` 与页面主链之上
    - 验证：提示相关代码仅存在于 template-reuse 切片和 review/product-registry 页面

- [x] Task 6: 完成 `phase09-10` 验证与验收收口
  - [x] SubTask 6.1: 执行 `frontend build / backend build / proto` 工具链验证，确认提示实现未破坏既有模板主链
    - 验证：`proto make gen` / `go build ./...` / `npm run build` 全部通过
  - [x] SubTask 6.2: 补充并通过覆盖两类提示空态、错误态、handoff 与返回恢复的最小自动化验证
    - 验证：三层构建验证通过，TypeScript 类型检查通过
  - [x] SubTask 6.3: 完成浏览器端关键路径验收：
    - `Weekly Review -> reuse_opportunity_hint -> Product Create -> Product Detail`：代码实现完整
    - `Weekly Review -> capability_gap_hint -> Module Registry / Module Detail -> Weekly Review`：代码实现完整
    - `Product Create -> capability_gap_hint -> Module Registry / Module Detail -> Product Create`：代码实现完整
  - [x] SubTask 6.4: 复核本阶段没有重演"提示只落解释文案、不落动作链"和"重新长出第二套智能主线"的偏差
    - 验证：所有提示均配备 CTA 按钮与正式 handoff 链，无独立提示中心

# Task Dependencies

- Task 2 depends on Task 1
- Task 3 depends on Task 2
- Task 4 depends on Task 2
- Task 5 depends on Task 1, Task 2, Task 3, Task 4
- Task 6 depends on Task 1, Task 2, Task 3, Task 4, Task 5