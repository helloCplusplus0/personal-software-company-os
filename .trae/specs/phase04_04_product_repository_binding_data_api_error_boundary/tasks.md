# Tasks

> 阶段分工说明：本规格只冻结 `phase04-04` 范围内的数据读写范围、最小接口边界与错误语义前提。原 Tasks 中的接口分组 / 方向级 API 矩阵 / 关系写入后读取语义 / 合同与存储解耦 / 候选读取归属与隔离 / 不冻结的实现工具已移出，分别由 `phase04-07`（后端模块边界与接口分组设计）与 `phase04-08`（`.proto` 合同设计）承接。

- [x] Task 1: 冻结 Product Registry 最小数据读写范围。将列表读取、详情读取、创建写入、绑定写入与候选读取的数据范围写成单值结论。
  - [x] SubTask 1.1: 明确 `ProductListRead` 数据范围（`name / description / status / created_at / module_bind_count / repository_bind_count`），筛选参数（`queryText + statusFilter`），`queryText` 只匹配 `name`，排序 `created_at` 降序，不引入分页
    - 证据：`spec.md` §ADDED Requirements「Product Registry 最小数据读写范围冻结」Scenario: 判断 Product List 读取数据范围 L32-38
  - [x] SubTask 1.2: 明确 `ProductDetailRead` 数据范围（核心字段 `id / name / description / status / created_at` + 已绑定模块列表 + 已绑定仓库列表）
    - 证据：`spec.md` §ADDED Requirements「Product Registry 最小数据读写范围冻结」Scenario: 判断 Product Detail 读取数据范围 L42-46
  - [x] SubTask 1.3: 明确 `ProductCreateWrite` 数据范围（`name / description / status` 必填，`status` 属于 `active / archived`，默认预填并显式提交 `active`）
    - 证据：`spec.md` §ADDED Requirements「Product Registry 最小数据读写范围冻结」Scenario: 判断 CreateProduct 写入数据范围 L50-55
  - [x] SubTask 1.4: 明确 `ProductModuleBindingWrite` 数据范围（`product_id / module_id` 必填，`product_id` 由上下文隐式承接，`module_id` 由候选选择）
    - 证据：`spec.md` §ADDED Requirements「Product Registry 最小数据读写范围冻结」Scenario: 判断 BindModuleToProduct 写入数据范围 L59-64
  - [x] SubTask 1.5: 明确 `ProductModuleCandidateRead` 数据范围（`module_id / module_name / module_status`，`active` 状态，已绑定排除，`created_at` 降序）
    - 证据：`spec.md` §ADDED Requirements「Product Registry 最小数据读写范围冻结」Scenario: 判断 Product Module 候选读取数据范围 L68-73

- [x] Task 2: 冻结 Repository Binding 最小数据读写范围。将列表读取、详情读取、创建写入、两类绑定写入与两类候选读取的数据范围写成单值结论。
  - [x] SubTask 2.1: 明确 `RepositoryListRead` 数据范围（`name / url / provider / status / created_at / product_bind_count / module_bind_count`），筛选参数（`queryText + statusFilter`），`queryText` 只匹配 `name`，排序 `created_at` 降序，不引入分页，不引入 `providerFilter`
    - 证据：`spec.md` §ADDED Requirements「Repository Binding 最小数据读写范围冻结」Scenario: 判断 Repository List 读取数据范围 L82-89
  - [x] SubTask 2.2: 明确 `RepositoryDetailRead` 数据范围（核心字段 `id / name / url / provider / status / created_at` + 已绑定产品列表 + 已映射模块列表）
    - 证据：`spec.md` §ADDED Requirements「Repository Binding 最小数据读写范围冻结」Scenario: 判断 Repository Detail 读取数据范围 L93-97
  - [x] SubTask 2.3: 明确 `RepositoryCreateWrite` 数据范围（`name / url / provider / status` 必填，`provider` 为必填字符串不采用受控枚举）
    - 证据：`spec.md` §ADDED Requirements「Repository Binding 最小数据读写范围冻结」Scenario: 判断 CreateRepository 写入数据范围 L101-107
  - [x] SubTask 2.4: 明确 `RepositoryProductBindingWrite` 数据范围（`repository_id / product_id` 必填，`repository_id` 由上下文隐式承接，`product_id` 由候选选择）
    - 证据：`spec.md` §ADDED Requirements「Repository Binding 最小数据读写范围冻结」Scenario: 判断 BindRepositoryToProduct 写入数据范围 L111-116
  - [x] SubTask 2.5: 明确 `RepositoryModuleMappingWrite` 数据范围（`repository_id / module_id` 必填，`repository_id` 由上下文隐式承接，`module_id` 由候选选择）
    - 证据：`spec.md` §ADDED Requirements「Repository Binding 最小数据读写范围冻结」Scenario: 判断 MapModuleToRepository 写入数据范围 L120-125
  - [x] SubTask 2.6: 明确 `ProductBindingCandidateRead` 数据范围（`product_id / product_name / product_status`，`active` 状态，已绑定排除，`created_at` 降序）
    - 证据：`spec.md` §ADDED Requirements「Repository Binding 最小数据读写范围冻结」Scenario: 判断 Repository Product 候选读取数据范围 L129-134
  - [x] SubTask 2.7: 明确 `RepositoryModuleCandidateRead` 数据范围（`module_id / module_name / module_status`，`active` 状态，已映射排除，`created_at` 降序）
    - 证据：`spec.md` §ADDED Requirements「Repository Binding 最小数据读写范围冻结」Scenario: 判断 Repository Module 候选读取数据范围 L138-143

- [x] Task 3: 冻结详情读取与候选读取边界。明确两者独立承接，不得合并或拆散。
  - [x] SubTask 3.1: 明确详情读取只承接详情本体与已绑定列表，候选读取必须通过独立 request / response 承接，不得并入详情读取或拆散
    - 证据：`spec.md` §ADDED Requirements「详情读取与候选读取边界冻结」Scenario: 判断详情读取与候选读取独立性 L151-157

- [x] Task 4: 冻结错误语义前提。把创建失败、目标不存在、重复绑定、候选空结果与列表空结果的错误语义写成单值结论。
  - [x] SubTask 4.1: 明确创建失败错误语义（`CreateProduct / CreateRepository` 缺失必填或非法 `status` → 校验失败，不得降级为通用错误或 `500` 级未收口错误）
    - 证据：`spec.md` §ADDED Requirements「错误语义前提冻结」Scenario: 判断创建失败错误语义 L165-174
  - [x] SubTask 4.2: 明确目标不存在错误语义（详情读取、三类绑定写入与候选读取依附实体的目标不存在 → 资源不存在，按归属接口单值化）
    - 证据：`spec.md` §ADDED Requirements「错误语义前提冻结」Scenario: 判断目标不存在错误语义 L178-199
  - [x] SubTask 4.3: 明确重复绑定错误语义（三类绑定写入重复 → 重复冲突，不得静默成功，不得 `ON CONFLICT DO NOTHING`）
    - 证据：`spec.md` §ADDED Requirements「错误语义前提冻结」Scenario: 判断重复绑定错误语义 L203-210
  - [x] SubTask 4.4: 明确候选读取空结果语义（返回空列表语义，不映射为资源不存在或接口错误）
    - 证据：`spec.md` §ADDED Requirements「错误语义前提冻结」Scenario: 判断候选读取空结果语义 L214-219
  - [x] SubTask 4.5: 明确列表读取空结果语义（返回空列表语义，不映射为资源不存在或接口错误）
    - 证据：`spec.md` §ADDED Requirements「错误语义前提冻结」Scenario: 判断列表读取空结果语义 L223-226
  - [x] SubTask 4.6: 明确三类绑定写入的校验失败类型（必填缺失 → 校验失败，目标不存在 → 资源不存在，目标非 `active` → 校验失败，重复绑定 → 重复冲突）
    - 证据：`spec.md` §ADDED Requirements「错误语义前提冻结」Scenario: 判断三类绑定写入的校验失败类型 L230-236

- [x] Task 5: 冻结非目标边界。明确不提前冻结 Dashboard 聚合、`product_asset_coverage`、`Decision -> Product / Repository` 写入与超出当前阶段的分页/复杂检索/批量操作。
  - [x] SubTask 5.1: 明确不提前冻结 Dashboard 聚合接口、`product_asset_coverage` 聚合读取、`Decision -> Product / Repository` 关联写入与超出当前阶段的分页/复杂检索/批量操作
    - 证据：`spec.md` §ADDED Requirements「非目标冻结」Scenario: 判断非目标边界 L244-248 + §REMOVED Requirements「Dashboard 聚合接口」L268-271 +「超出当前阶段的分页与复杂检索接口」L273-276

- [x] Task 6: 冻结 phase02 迁移边界解释。明确 `ProductBindingCandidateRead` 迁移、`RepositoryBindingCandidateRead` 废弃、`ModuleBindingWrite` 拆分迁移后的迁移边界，不冻结接口名是否沿用或拆分后的具体接口名。
  - [x] SubTask 6.1: 明确 `ProductBindingCandidateRead` 迁移到 `Repository Binding` 模块由其拥有，`RepositoryBindingCandidateRead` 废弃，`ModuleBindingWrite` 拆分迁移到 `Product Registry` 与 `Repository Binding` 模块；不冻结接口名是否沿用，不冻结拆分后的具体接口名
    - 证据：`spec.md` §MODIFIED Requirements「phase02 临时承接接口迁移后的迁移边界解释」Scenario L260-264

- [x] Task 7: 完成规格校验。确认本次 `phase04-04` 规格可以直接作为后续实现设计的上游。
  - [x] SubTask 7.1: 验证 `Product Registry` 与 `Repository Binding` 数据读写范围已经单值化
    - 证据：`spec.md` §ADDED Requirements「Product Registry 最小数据读写范围冻结」+「Repository Binding 最小数据读写范围冻结」— 两个模块各自单值化为列表/详情/创建/绑定/候选五类数据范围
  - [x] SubTask 7.2: 验证详情读取与候选读取边界已经单值化
    - 证据：`spec.md` §ADDED Requirements「详情读取与候选读取边界冻结」— 详情读取与候选读取独立性已单值化
  - [x] SubTask 7.3: 验证错误语义前提已经单值化
    - 证据：`spec.md` §ADDED Requirements「错误语义前提冻结」— 创建失败、目标不存在、重复绑定、候选空结果、列表空结果与三类绑定写入校验失败类型已单值化
  - [x] SubTask 7.4: 验证非目标与迁移边界解释已经单值化
    - 证据：`spec.md` §ADDED Requirements「非目标冻结」+ §MODIFIED Requirements「phase02 临时承接接口迁移后的迁移边界解释」+ §REMOVED Requirements — 全部已单值化
  - [x] SubTask 7.5: 验证本规格未越界冻结 `phase04-07`（接口分组、方向级 API 矩阵、canonical owner 兼容策略、候选读取接线位置）与 `phase04-08`（合同演进约束、`.proto` 服务接口命名）的设计职责
    - 证据：`spec.md` Why 部分阶段分工约束 + 数据读写范围冻结 Requirement 内引用说明 + MODIFIED Requirement 内引用说明 — 接口分组、方向级 API 矩阵、接线位置、合同演进约束均未在本规格冻结

# Task Dependencies

- `Task 3` depends on `Task 1`, `Task 2`
- `Task 4` depends on `Task 1`, `Task 2`
- `Task 5` depends on `Task 1`, `Task 2`
- `Task 6` depends on `Task 1`, `Task 2`
- `Task 7` depends on `Task 1`, `Task 2`, `Task 3`, `Task 4`, `Task 5`, and `Task 6`
