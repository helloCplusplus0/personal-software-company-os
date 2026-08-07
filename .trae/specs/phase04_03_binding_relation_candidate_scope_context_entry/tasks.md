# Tasks

- [x] Task 1: 冻结三类绑定关系的关系语义。将 `product_repositories / product_modules / module_repositories` 的多对多关系、语义解释与唯一性约束写成单值结论。
  - [x] SubTask 1.1: 明确 `product_repositories` 为 `Product` 与 `Repository` 之间的多对多绑定关系，语义为"Repository R 是 Product P 的实现锚点"，由 `BindRepositoryToProduct` 建立
    - 证据：`spec.md` §ADDED Requirements「三类绑定关系关系语义冻结」Scenario: 判断 product_repositories 关系语义 L31-35 — 多对多绑定关系，语义为"实现锚点"，由 `BindRepositoryToProduct` 建立
  - [x] SubTask 1.2: 明确 `product_modules` 为 `Product` 与 `Module` 之间的多对多绑定关系，语义为"Module M 被 Product P 使用"，由 `BindModuleToProduct` 建立
    - 证据：`spec.md` §ADDED Requirements「三类绑定关系关系语义冻结」Scenario: 判断 product_modules 关系语义 L39-43 — 多对多绑定关系，语义为"被使用"，由 `BindModuleToProduct` 建立
  - [x] SubTask 1.3: 明确 `module_repositories` 为 `Module` 与 `Repository` 之间的多对多映射关系，语义为"Module M 实现于 Repository R"，由 `MapModuleToRepository` 建立
    - 证据：`spec.md` §ADDED Requirements「三类绑定关系关系语义冻结」Scenario: 判断 module_repositories 关系语义 L47-51 — 多对多映射关系，语义为"实现于"，由 `MapModuleToRepository` 建立
  - [x] SubTask 1.4: 明确三类绑定关系的唯一性约束均为同一对组合只允许一条记录
    - 证据：`spec.md` §ADDED Requirements「三类绑定关系关系语义冻结」三个 Scenario 中"同一 (xxx_id, yyy_id) 对只允许存在一条绑定记录" L34 / L42 / L50

- [x] Task 2: 冻结三类绑定动作的 canonical owner 与 reread 承接页面。把绑定写入归属与成功后 reread 页面写成单值结论。
  - [x] SubTask 2.1: 明确 `BindModuleToProduct` → `Product Detail`（`Product Registry` 模块）
    - 证据：`spec.md` §ADDED Requirements「三类绑定动作 canonical owner 冻结」Scenario: 判断 canonical owner L135-140 — `BindModuleToProduct` → `Product Detail`（`Product Registry` 模块）
  - [x] SubTask 2.2: 明确 `BindRepositoryToProduct` → `Repository Binding Detail / Workspace`（`Repository Binding` 模块）
    - 证据：`spec.md` §ADDED Requirements「三类绑定动作 canonical owner 冻结」Scenario: 判断 canonical owner L135-140 — `BindRepositoryToProduct` → `Repository Binding Detail / Workspace`（`Repository Binding` 模块）
  - [x] SubTask 2.3: 明确 `MapModuleToRepository` → `Repository Binding Detail / Workspace`（`Repository Binding` 模块）
    - 证据：`spec.md` §ADDED Requirements「三类绑定动作 canonical owner 冻结」Scenario: 判断 canonical owner L135-140 — `MapModuleToRepository` → `Repository Binding Detail / Workspace`（`Repository Binding` 模块）
  - [x] SubTask 2.4: 明确绑定成功后必须回到 canonical owner 页面完成 reread，不得只靠 toast
    - 证据：`spec.md` §ADDED Requirements「绑定成功后 reread 承接页面冻结」Scenario: 判断 reread 承接页面 L148-153 — 回到 canonical owner 页面完成 reread，不得只靠 toast

- [x] Task 3: 冻结三类绑定动作的候选范围。把候选过滤策略、排序规则与已绑定排除写成单值结论。
  - [x] SubTask 3.1: 明确 `BindRepositoryToProduct` 候选为 `active` 状态 `Product`，排除已绑定到当前 `Repository` 的 `Product`，按 `created_at` 降序
    - 证据：`spec.md` §ADDED Requirements「BindRepositoryToProduct 候选范围冻结」Scenario: 判断 BindRepositoryToProduct 候选范围 L59-63 — `active` 状态、已绑定排除、`created_at` 降序、不得纳入 `archived`
  - [x] SubTask 3.2: 明确 `BindModuleToProduct` 候选为 `active` 状态 `Module`，排除已绑定到当前 `Product` 的 `Module`，按 `created_at` 降序
    - 证据：`spec.md` §ADDED Requirements「BindModuleToProduct 候选范围冻结」Scenario: 判断 BindModuleToProduct 候选范围 L71-75 — `active` 状态、已绑定排除、`created_at` 降序、不得纳入 `archived`
  - [x] SubTask 3.3: 明确 `MapModuleToRepository` 候选为 `active` 状态 `Module`，排除已映射到当前 `Repository` 的 `Module`，按 `created_at` 降序
    - 证据：`spec.md` §ADDED Requirements「MapModuleToRepository 候选范围冻结」Scenario: 判断 MapModuleToRepository 候选范围 L83-87 — `active` 状态、已映射排除、`created_at` 降序、不得纳入 `archived`
  - [x] SubTask 3.4: 明确 `archived` 状态记录不进入候选列表
    - 证据：`spec.md` §ADDED Requirements 三个候选范围 Requirement 中"不得将 `archived` 状态的 xxx 纳入候选列表" L63 / L75 / L87 + §REMOVED Requirements「基于 archived 记录建立新绑定」L291-294

- [x] Task 4: 冻结候选读取空状态、展示模型与重复绑定语义。把空候选返回值、候选字段集合与重复绑定返回语义写成单值结论。
  - [x] SubTask 4.1: 明确候选读取返回零条记录时必须返回空列表语义，不得映射为资源不存在或接口错误
    - 证据：`spec.md` §ADDED Requirements「候选读取空状态语义冻结」Scenario: 判断候选读取空状态 L95-99 — 返回空列表语义，不得映射为资源不存在或接口错误
  - [x] SubTask 4.2: 明确候选 `Product` 展示模型为 `product_id / product_name / product_status`
    - 证据：`spec.md` §ADDED Requirements「候选读取展示模型冻结」Scenario: 判断候选 Product 展示模型 L107-109 — `product_id / product_name / product_status`
  - [x] SubTask 4.3: 明确候选 `Module` 展示模型为 `module_id / module_name / module_status`
    - 证据：`spec.md` §ADDED Requirements「候选读取展示模型冻结」Scenario: 判断候选 Module 展示模型 L113-115 — `module_id / module_name / module_status`
  - [x] SubTask 4.4: 明确重复绑定必须返回明确的重复冲突语义，不得静默成功或降级为通用错误
    - 证据：`spec.md` §ADDED Requirements「重复绑定语义冻结」Scenario: 判断重复绑定语义 L123-127 — 返回重复冲突语义，不得静默成功，不得降级为通用错误，不得通过 `ON CONFLICT DO NOTHING` 隐式吞掉

- [x] Task 5: 冻结 Module Detail 旧入口兼容跳转与 Product Detail 上下文入口跳转参数。把旧入口保留级别、跳转参数与接收方预填行为写成单值结论。
  - [x] SubTask 5.1: 明确 `Module Detail` 只允许兼容跳转，不得直接提交绑定写入
    - 证据：`spec.md` §ADDED Requirements「Module Detail 旧入口兼容跳转与参数冻结」Scenario: 判断 Module Detail 旧入口保留级别 L161-164 — 只允许兼容跳转，不得直接提交绑定写入
  - [x] SubTask 5.2: 明确 `Module Detail` 兼容跳转进入对应绑定动作的正式主入口，上下文参数为 `moduleId / moduleName / fromModuleDetail`，与目标页身份参数 `productId / repositoryId` 拆开传递，接收方必须能预填候选 `Module` 选择
    - 证据：`spec.md` §ADDED Requirements「Module Detail 旧入口兼容跳转与参数冻结」Scenario: 判断 Module Detail 兼容跳转参数 L168-173 — 必须进入对应绑定动作的正式主入口，携带 `moduleId / moduleName / fromModuleDetail` 作为上下文参数；上下文参数只表示来源模块身份与来源页面标记，不表示目标实体身份；目标实体未确定时先进入对应列表页（`Product Registry / List` 或 `Repository Binding / List`）选择，已确定时额外携带 `productId` 或 `repositoryId` 作为目标页身份参数与上下文参数拆开传递；接收方页面必须能基于上下文参数预填候选 `Module` 选择
  - [x] SubTask 5.3: 明确 `Product Detail` 兼容跳转进入 `BindRepositoryToProduct` 的正式主入口，上下文参数为 `productId / productName / fromProductDetail`，与目标页身份参数 `repositoryId` 拆开传递，接收方必须能预填候选 `Product` 选择
    - 证据：`spec.md` §ADDED Requirements「Product Detail 上下文入口跳转参数冻结」Scenario: 判断 Product Detail 上下文跳转参数 L179-187 — 必须进入 `BindRepositoryToProduct` 的正式主入口，携带 `productId / productName / fromProductDetail` 作为上下文参数；上下文参数只表示来源产品身份与来源页面标记，不表示目标 `Repository` 身份；目标 `Repository` 未确定时先进入 `Repository Binding / List` 选择，已确定时额外携带 `repositoryId` 作为目标页身份参数与上下文参数拆开传递；接收方页面必须能基于上下文参数预填候选 `Product` 选择；`Product Detail` 自身不得承接第二套仓库绑定写入流程

- [x] Task 6: 冻结 phase02 临时承接迁移边界与候选读取接口归属。把接口迁移、数据访问迁移、历史数据兼容与候选读取归属写成单值结论。
  - [x] SubTask 6.1: 明确 `ProductBindingCandidateRead` 从 `Module Registry` 迁移到 `Repository Binding` 模块的 `candidate/` 子包
    - 证据：`spec.md` §ADDED Requirements「phase02 临时承接迁移边界冻结」Scenario: 判断 ProductBindingCandidateRead 迁移边界 L195-198 — 迁移到 `Repository Binding` 模块 `candidate/` 子包，接口契约由 `Repository Binding` 拥有
  - [x] SubTask 6.2: 明确 `RepositoryBindingCandidateRead` 标记为废弃，不在 `phase04` 保留并行实现
    - 证据：`spec.md` §ADDED Requirements「phase02 临时承接迁移边界冻结」Scenario: 判断 RepositoryBindingCandidateRead 迁移边界 L202-205 — 废弃，`MapModuleToRepository` 候选为 `Module` 而非 `Repository`
  - [x] SubTask 6.3: 明确 `ModuleBindingWrite` 拆分迁移：`BindModuleToProduct` → `Product Registry`，`MapModuleToRepository` → `Repository Binding`，`BindRepositoryToProduct` 作为 `Repository Binding` 新增能力
    - 证据：`spec.md` §ADDED Requirements「phase02 临时承接迁移边界冻结」Scenario: 判断 ModuleBindingWrite 迁移边界 L209-213 — 拆分迁移到 `Product Registry` 与 `Repository Binding`，`BindRepositoryToProduct` 作为新增能力
  - [x] SubTask 6.4: 明确 `binding_store` 拆分迁移：`product_modules` → `Product Registry`，`module_repositories` → `Repository Binding`，`product_repositories` 作为 `Repository Binding` 新增能力
    - 证据：`spec.md` §ADDED Requirements「phase02 临时承接迁移边界冻结」Scenario: 判断 binding_store 迁移边界 L217-221 — `product_modules` → `Product Registry`，`module_repositories` → `Repository Binding`，`product_repositories` 作为新增
  - [x] SubTask 6.5: 明确历史绑定数据必须保持可读，不得通过影子表或双写绕过迁移
    - 证据：`spec.md` §ADDED Requirements「phase02 临时承接迁移边界冻结」Scenario: 判断历史绑定数据兼容前提 L225-228 — 历史数据必须保持可读，不得通过影子表或双写绕过迁移
  - [x] SubTask 6.6: 明确 `Product Registry` 模块拥有 `BindModuleToProduct` 的候选 `Module` 读取接口
    - 证据：`spec.md` §ADDED Requirements「候选读取接口归属冻结」Scenario: 判断 Product Registry 模块候选读取归属 L236-239 — 由 `Product Registry` 模块 `candidate/` 子包拥有，service 不得直接写跨模块 SQL
  - [x] SubTask 6.7: 明确 `Repository Binding` 模块拥有 `BindRepositoryToProduct` 的候选 `Product` 读取接口与 `MapModuleToRepository` 的候选 `Module` 读取接口
    - 证据：`spec.md` §ADDED Requirements「候选读取接口归属冻结」Scenario: 判断 Repository Binding 模块候选读取归属 L243-246 — 由 `Repository Binding` 模块 `candidate/` 子包拥有，service 不得直接写跨模块 SQL

- [x] Task 7: 冻结非目标边界。明确不把 `Decision Center`、`Module Registry` 重新扩写为并列绑定主线。
  - [x] SubTask 7.1: 明确不把 `Decision Center` 扩写为 `Product / Repository` 绑定写入主线
    - 证据：`spec.md` §ADDED Requirements「非目标冻结」Scenario: 判断非目标边界 L254-258 — 不得把 `Decision Center` 扩写为绑定写入主线
  - [x] SubTask 7.2: 明确不把 `Module Registry` 重新扩写为绑定写入主线，`Module Detail` 只保留兼容跳转
    - 证据：`spec.md` §ADDED Requirements「非目标冻结」Scenario: 判断非目标边界 L254-258 — 不得把 `Module Registry` 重新扩写为绑定写入主线，`Module Detail` 只允许保留兼容跳转

- [x] Task 8: 完成规格校验。确认本次 `phase04-03` 规格可以直接作为后续页面、接口与合同设计的上游。
  - [x] SubTask 8.1: 验证三类绑定关系关系语义已经单值化
    - 证据：`spec.md` §ADDED Requirements「三类绑定关系关系语义冻结」— 三类关系各自单值化为多对多 + 唯一性约束 + 语义解释
  - [x] SubTask 8.2: 验证 canonical owner 与 reread 承接页面已经明确
    - 证据：`spec.md` §ADDED Requirements「三类绑定动作 canonical owner 冻结」+「绑定成功后 reread 承接页面冻结」— owner 与 reread 已单值化
  - [x] SubTask 8.3: 验证候选范围、空状态、展示模型与重复绑定语义已经明确
    - 证据：`spec.md` §ADDED Requirements「BindRepositoryToProduct 候选范围冻结」+「BindModuleToProduct 候选范围冻结」+「MapModuleToRepository 候选范围冻结」+「候选读取空状态语义冻结」+「候选读取展示模型冻结」+「重复绑定语义冻结」— 候选与重复语义已单值化
  - [x] SubTask 8.4: 验证 Module Detail 旧入口兼容跳转参数与 Product Detail 上下文入口跳转参数已经明确
    - 证据：`spec.md` §ADDED Requirements「Module Detail 旧入口兼容跳转与参数冻结」+「Product Detail 上下文入口跳转参数冻结」— 跳转参数已单值化
  - [x] SubTask 8.5: 验证 phase02 迁移边界与候选读取接口归属已经明确
    - 证据：`spec.md` §ADDED Requirements「phase02 临时承接迁移边界冻结」+「候选读取接口归属冻结」— 迁移边界与归属已单值化
  - [x] SubTask 8.6: 验证非目标边界已经明确
    - 证据：`spec.md` §ADDED Requirements「非目标冻结」+ §REMOVED Requirements「Module Detail 作为绑定写入主入口」L286-289 +「基于 archived 记录建立新绑定」L291-294 — 非目标已单值化

# Task Dependencies

- `Task 2` depends on `Task 1`
- `Task 3` depends on `Task 1`
- `Task 4` depends on `Task 3`
- `Task 5` depends on `Task 2`
- `Task 6` depends on `Task 2`, `Task 3`
- `Task 7` depends on `Task 2`
- `Task 8` depends on `Task 1`, `Task 2`, `Task 3`, `Task 4`, `Task 5`, `Task 6`, and `Task 7`
