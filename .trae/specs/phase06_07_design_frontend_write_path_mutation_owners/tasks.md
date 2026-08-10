# Tasks

- [x] Task 1: 产出 Application Owner 切片落点设计。把 4 个 feature slice 的 `application/` 目录结构、文件命名与落点约束冻结到可直接落地的程度。
  - [x] SubTask 1.1: 明确 4 个切片各自的 `application/` 目录结构与文件清单（`index.ts` / `normalize-error.ts` / `use-create-draft-*.ts` / `use-update-*.ts`）
  - [x] SubTask 1.2: 明确 `application/` 目录是该切片唯一正式 mutation 承接位，不得承接读取逻辑
  - [x] SubTask 1.3: 明确 `normalize-error.ts` 是切片级工具，延迟晋升到 `shared/lib/`
  - [x] SubTask 1.4: 明确 `use-update-*.ts` 在 phase06 只预留落点与接口签名，不要求实现
  - [x] SubTask 1.5: 明确 `application/` 与既有 `data/` 层的关系（`data/` 保持纯只读 + HTTP 调用，application owner 内部调用 `createXxx`）

- [x] Task 2: 产出当前漂移点全景。识别前端所有 page-level / panel-level / release 级 `useMutation` 漂移点。
  - [x] SubTask 2.1: 识别 4 个 create 页面的 page-level `useMutation`（文件、行号、mutationFn、失效 queryKey、回流、错误处理）
  - [x] SubTask 2.2: 识别 4 个 binding/link panel 的 panel-level `useMutation`（文件、行号、mutationFn、失效 queryKey）
  - [x] SubTask 2.3: 识别 release create page 的 page-level `useMutation`
  - [x] SubTask 2.4: 明确 4 个 create 页面共同模式（`useMutation` + `invalidateQueries` + `toast` + `navigate`）

- [x] Task 3: 产出 4 个 Create Draft Application Owner 设计。把接口契约、输入类型、系统填充、失效目标与职责边界写成单值结论。
  - [x] SubTask 3.1: 明确统一接口契约（`mutate / mutateAsync / isPending / isError / error / data`）
  - [x] SubTask 3.2: 明确 `useCreateDraftProduct` 的输入类型（`name` 必填 + `description? / status?` 可选）、系统填充（`description: ''`, `status: 'active'`）、失效目标
  - [x] SubTask 3.3: 明确 `useCreateDraftRepository` 的输入类型（`name + url` 必填 + `provider? / status?` 可选）、系统填充（`provider: 'manual'`, `status: 'active'`）、失效目标
  - [x] SubTask 3.4: 明确 `useCreateDraftModule` 的输入类型（`name` 必填 + `description? / status? / capability_key?` 可选）、系统填充、失效目标
  - [x] SubTask 3.5: 明确 `useCreateDraftDecision` 的输入类型（`title + choice + reason` 必填 + 其余可选）、系统填充、失效目标
  - [x] SubTask 3.6: 明确 Draft Input 与既有 `CreateXxxInput` 的关系（Draft 类型定义在 `application/`，不修改 `types.ts`）

- [x] Task 4: 产出 Update 写路径预留落点设计。
  - [x] SubTask 4.1: 明确 4 个 `use-update-*.ts` 的文件落点与接口签名
  - [x] SubTask 4.2: 明确 phase06 阶段只创建文件与签名，不要求实现 `mutationFn`
  - [x] SubTask 4.3: 明确后续实现时的失效目标（detail + list + onboarding read）

- [x] Task 5: 产出统一错误归一化策略。
  - [x] SubTask 5.1: 明确 `NormalizedError` 形状（`message` + 可选 `status` / `code`）
  - [x] SubTask 5.2: 明确 `normalizeApiError` 函数实现（`ApiError` → `NormalizedError`，`Error` → `NormalizedError`，`unknown` → 兜底）
  - [x] SubTask 5.3: 明确归一化约束（application owner 不得暴露原始错误，消费方只消费 `message`）
  - [x] SubTask 5.4: 明确切片级延迟晋升约束（先落切片，稳定后提升到 `shared/lib/`）
  - [x] SubTask 5.5: 明确消费方错误展示方式（create 页面 `toast.error`，Onboarding 步骤传入 `submitError`）

- [x] Task 6: 产出统一成功回流与 query 失效策略。
  - [x] SubTask 6.1: 明确 query 失效矩阵（4 个 owner 各自失效 canonical 列表 + `['onboarding', 'read']`）
  - [x] SubTask 6.2: 明确回流职责拆分（application owner: API + 填充 + 失效 + 归一化；消费方: toast + navigate + 来源上下文）
  - [x] SubTask 6.3: 明确 create 页面回流模式（`mutateAsync` + `try/catch` + `navigate` 携带来源上下文）
  - [x] SubTask 6.4: 明确 Onboarding 步骤回流模式（`mutate` fire-and-forget + query 失效驱动自动前进）
  - [x] SubTask 6.5: 明确回流约束（application owner 不得 `navigate` / `toast`，create 页面不得保留 `useMutation` / `useQueryClient` import）

- [x] Task 7: 产出旧模式回收清单。
  - [x] SubTask 7.1: 明确必须回收的 4 个 create 页面清单与对应 application owner
  - [x] SubTask 7.2: 明确允许过渡保留的 4 个 binding/link panel + 1 个 release create page 清单
  - [x] SubTask 7.3: 明确回收后 create 页面的职责（只承接来源上下文 + 消费 owner + toast + navigate + 表单壳层）
  - [x] SubTask 7.4: 明确回收后 create 页面的禁止事项（不得 `useMutation` / `createXxx` / `invalidateQueries` / `onSuccess` / `onError`）
  - [x] SubTask 7.5: 明确表单组件 `onSubmit / submitting / submitError` props 接口结构不变，但字段级必填约束必须按 §7.5 放宽
  - [x] SubTask 7.6: 明确 4 个表单组件的字段级放宽清单（Product: description 不阻断 + status 移除；Repository: provider 不阻断 + status 移除；Module: description 不阻断 + status 移除；Decision: context/problem 不阻断 + status 移除）
  - [x] SubTask 7.7: 明确 DecisionCreatePage 回收后完整保留 `sourceModuleId` + `sourceModuleName` 来源链路（search params → 面板展示 → 提交持久化）

- [x] Task 8: 产出回收迁移顺序。
  - [x] SubTask 8.1: 明确 6 个迁移阶段（目录创建 → 参考实现 → 验证回收+表单放宽 → 并行实现 → 并行回收+表单放宽 → 预留落点）
  - [x] SubTask 8.2: 明确阶段依赖关系与可并行性
  - [x] SubTask 8.3: 明确迁移约束（先验证模式再并行推广、行为不变、API 形状不变、不保留旧 import、表单字段放宽同步完成）

- [x] Task 9: 产出 phase06 收口标准。
  - [x] SubTask 9.1: 明确必须满足的 10 项收口条件（含表单字段放宽、Decision 来源链路保持）
  - [x] SubTask 9.2: 明确禁止事项（不得内联 `useMutation`、不得保留旧表单必填阻断与 status 选择器、不得承接读取、不得 `navigate` / `toast`、不得复制散装模式、不得跨切片引用 `normalize-error.ts`、不得丢失 Decision 来源链路）

- [x] Task 10: 完成规格一致性校验。验证本次设计与 phase06-01/02/05/06 已冻结语义、shared_baseline 保持一致。
  - [x] SubTask 10.1: 验证 4 个 owner 的系统填充默认值与 phase06-02 冻结的 draft-first 字段一致
  - [x] SubTask 10.2: 验证 application owner 命名（`useCreateDraftProduct` 等）与 phase06-06 §11 消费方式一致
  - [x] SubTask 10.3: 验证 query 失效包含 `['onboarding', 'read']` 与 phase06-06 §11.2 一致
  - [x] SubTask 10.4: 验证 `query` 纯只读、mutation 固定承接位与 phase06-05 §前端四条约束一致
  - [x] SubTask 10.5: 验证既有 create 页面回收范围与 phase06-02 §"既有 Create 页面回收范围冻结"一致
  - [x] SubTask 10.6: 验证过渡保留的 panel / release 旧模式与 phase06-05 §"旧实现过渡与新增门禁"一致
  - [x] SubTask 10.7: 验证表单字段级放宽与 phase06-02 冻结的 draft-first 最小必填字段一致（Product=name, Repository=name+url, Module=name, Decision=title+choice+reason）
  - [x] SubTask 10.8: 验证 Decision 来源 Module 链路（sourceModuleId + sourceModuleName）与既有 routes/decisions/new.tsx validateSearch 与 decision-context-source-panel.tsx props 一致

# Task Dependencies

- `Task 2` depends on `Task 1`
- `Task 3` depends on `Task 1`
- `Task 4` depends on `Task 1`
- `Task 5` depends on `Task 1`
- `Task 6` depends on `Task 3` and `Task 5`
- `Task 7` depends on `Task 2` and `Task 3`
- `Task 8` depends on `Task 7`
- `Task 9` depends on `Task 7` and `Task 8`
- `Task 10` depends on `Task 1` through `Task 9`
