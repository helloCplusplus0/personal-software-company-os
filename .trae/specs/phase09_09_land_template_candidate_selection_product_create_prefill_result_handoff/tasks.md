# Tasks

- [x] Task 1: 实现 `use-product-create-template-handoff` application owner
  - [x] SubTask 1.1: 新增 `frontend/src/features/product-registry/application/use-product-create-template-handoff.ts`
  - [x] SubTask 1.2: 实现搜索参数解析（`fromTemplateReuse`, `templateCandidateId`, `templateSource`）
  - [x] SubTask 1.3: 组合 `use-template-prefill-read` 结果，导出 `templateSummary`, `prefillInitialValues`, `resolutionStatus`
  - [x] SubTask 1.4: 实现 `handleReturn()` 基于 `templateSource` 的返回路径
  - [x] SubTask 1.5: 实现 `buildSuccessSearch()` 创建成功回流参数
  - [x] SubTask 1.6: 实现非法参数回退（`fromTemplateReuse=true` 但 `templateCandidateId` 为空 → direct-entry）
  - [x] SubTask 1.7: 在 `templateSource=product-detail` 时补充 `templateSourceProductId`，确保取消返回能回到原 `Product Detail`

- [x] Task 2: 实现 `use-product-create-form-state` 正式 form state owner
  - [x] SubTask 2.1: 新增 `frontend/src/features/product-registry/application/use-product-create-form-state.ts`
  - [x] SubTask 2.2: 实现 `name / description / status` 单值状态与 `initialValues` 接受
  - [x] SubTask 2.3: 实现 `isDirty`, `buildSubmitInput()` 导出

- [x] Task 3: 重构 `ProductCreateForm` 为 props-driven
  - [x] SubTask 3.1: 移除组件本地 `useState`，改为接收 `name / description / status / onChangeName / onChangeDescription / onChangeStatus` props
  - [x] SubTask 3.2: 新增 `isFromTemplate` prop 控制预填标记展示

- [x] Task 4: 重构 `ProductCreatePage` 接入模板 handoff 与 form state
  - [x] SubTask 4.1: 消费 `use-product-create-form-state` 并传入 `initialValues`
  - [x] SubTask 4.2: 消费 `use-product-create-template-handoff`
  - [x] SubTask 4.3: 在表单上方展示模板来源摘要区（`border-t pt-2`）
  - [x] SubTask 4.4: 实现模板 unavailable 成功态与请求失败态
  - [x] SubTask 4.5: 实现 `templateSource` 驱动的取消返回
  - [x] SubTask 4.6: 实现创建成功回流携带模板参数

- [x] Task 5: 更新 `useWeeklyReviewRead` 内部组合模板读能力
  - [x] SubTask 5.1: 内部组合 `use-template-candidates-read` 与 `use-derived-insight-hints-read`
  - [x] SubTask 5.2: 导出 `templateCandidates`, `activeCandidateId`, `setActiveCandidateId`, `templateSectionStatus`, `hints`, `hintsSectionStatus`

- [x] Task 6: 在 `WeeklyReviewPage` 新增模板候选选择区
  - [x] SubTask 6.1: 新增模板候选区 UI 区块（位于 Overview 与 Recent Activity 之间）
  - [x] SubTask 6.2: 实现候选卡片列表 + 单选切换 + active 高亮
  - [x] SubTask 6.3: 实现 ready / empty / error 三态
  - [x] SubTask 6.4: 实现"以该模板创建产品"CTA 按钮与导航

- [x] Task 7: 在 `ProductDetailPage` 新增模板来源摘要区
  - [x] SubTask 7.1: 消费 `use-template-source-read`（`enabled = fromTemplateReuse && templateCandidateId !== ''`）
  - [x] SubTask 7.2: 在 `ProductSummaryCard` 与 `ReuseSummaryInline` 之间展示模板来源摘要
  - [x] SubTask 7.3: 实现 canonical binding CTA（滚动到 `ProductModuleBindingPanel`）
  - [x] SubTask 7.4: 实现 unavailable 可恢复空态

- [x] Task 8: 更新路由搜索参数
  - [x] SubTask 8.1: `/products/new` 新增 `fromTemplateReuse / templateCandidateId / templateSource`
  - [x] SubTask 8.2: `/products/$productId` 新增 `fromTemplateReuse / templateCandidateId / templateSource`
  - [x] SubTask 8.3: `/products/new` 在 `templateSource=product-detail` 场景新增 `templateSourceProductId`

- [x] Task 9: 完成构建验证与浏览器验收
  - [x] SubTask 9.1: `(cd frontend && npm run build)` 通过
  - [x] SubTask 9.2: `(cd backend && go build ./...)` 通过

# Task Dependencies

- Task 2 depends on Task 1 (form state 依赖 template handoff 的 initialValues 语义)
- Task 3 depends on Task 2 (form 重构依赖 form state owner)
- Task 4 depends on Task 1, Task 2, Task 3 (page 重构依赖所有三个 owner)
- Task 5, Task 6 与 Task 1-4 并行（review 页面独立）
- Task 7 独立（Product Detail 不依赖其他 Task）
- Task 8 与 Task 4, Task 7 并行（路由更新可提前）
- Task 9 depends on Task 1-8
