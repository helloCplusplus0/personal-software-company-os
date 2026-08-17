# Tasks

- [x] Task 1: 冻结治理画像前端正式承接位与页面落点
  - [x] SubTask 1.1: 明确 `Repository detail` 是唯一正式承接位
  - [x] SubTask 1.2: 明确治理画像区位于仓库业务主内容后的 secondary 区域
  - [x] SubTask 1.3: 复核不会新增独立页面、第二入口或跨页并列主承接区

- [x] Task 2: 冻结治理画像区的内部结构与 UI 层级
  - [x] SubTask 2.1: 明确“概览 / 维护 / 摘要回看”三部分结构
  - [x] SubTask 2.2: 明确 `structured_summary` 优先于路径信息成为主阅读内容
  - [x] SubTask 2.3: 复核不会把真实路径做成大块重复说明

- [x] Task 3: 冻结前端读路径与唯一 query owner
  - [x] SubTask 3.1: 明确治理画像读取以 `repository_id` 为唯一锚点
  - [x] SubTask 3.2: 明确 query owner 只承接只读解包、缓存键与错误状态
  - [x] SubTask 3.3: 复核不会继续复用 `phase12 project-context` 读取主线

- [x] Task 4: 冻结前端写路径与唯一 mutation owner
  - [x] SubTask 4.1: 明确 mutation owner 只承接治理画像保存、成功回流与错误归一化
  - [x] SubTask 4.2: 明确保存成功后精准刷新当前 `repository_id` 的治理画像读取
  - [x] SubTask 4.3: 复核不会在 page / form / card 中散落多套 `useMutation`

- [x] Task 5: 冻结第一版人类手工维护表单范围
  - [x] SubTask 5.1: 明确 `template_source / canonical_root_files[] / global_asset_bindings[]` 的维护入口
  - [x] SubTask 5.2: 明确 8 项全局规范资产矩阵的前端受控承接方式
  - [x] SubTask 5.3: 复核不会把 markdown 正文、只读字段或矩阵外资产带进表单

- [x] Task 6: 冻结只读展示与摘要回看范围
  - [x] SubTask 6.1: 明确 `project_profile_version / track_type / docs_workflow_layout / current_phase_*` 的只读展示
  - [x] SubTask 6.2: 明确 `markdown_resolvable` 与顶层目录矩阵的轻量回看方式
  - [x] SubTask 6.3: 复核不会把 `entry_ref` 和真实路径放大成主内容

- [x] Task 7: 冻结 phase12 遗留 `project-context` UI 的退出规则
  - [x] SubTask 7.1: 明确 `Repository detail` 现有“项目上下文”区必须移除
  - [x] SubTask 7.2: 明确 `Decision detail` 现有“共享项目上下文入口”卡片必须移除
  - [x] SubTask 7.3: 复核不会通过换名、并区或弱化文案延续旧设计

- [x] Task 8: 冻结 UI 层级与产品定位一致性
  - [x] SubTask 8.1: 明确治理画像区属于 secondary governance 区
  - [x] SubTask 8.2: 明确不抢占四实体业务主内容与详情页结构
  - [x] SubTask 8.3: 复核不会把 agent-only 协议做成前端主内容

- [x] Task 9: 完成 spec 包与上游冻结结论的一致性校验
  - [x] SubTask 9.1: 校验本 spec 与 `phase13-06` 的前端边界保持单值一致
  - [x] SubTask 9.2: 校验本 spec 与 `phase13-08` 的治理画像后端读写主线保持单值一致
  - [x] SubTask 9.3: 校验本 spec 与当前真实前端残留问题保持对齐，不遗漏 `Repository detail / Decision detail` 退出项

# Task Dependencies

- `Task 2` depends on `Task 1`
- `Task 3` depends on `Task 1`, `Task 2`
- `Task 4` depends on `Task 1`, `Task 2`, and `Task 3`
- `Task 5` depends on `Task 3` and `Task 4`
- `Task 6` depends on `Task 2`, `Task 3`, and `Task 5`
- `Task 7` depends on `Task 1`, `Task 2`, and `Task 6`
- `Task 8` depends on `Task 1`, `Task 2`, `Task 6`, and `Task 7`
- `Task 9` depends on `Task 1`, `Task 2`, `Task 3`, `Task 4`, `Task 5`, `Task 6`, `Task 7`, and `Task 8`

# 执行记录

- SubTask 1.1 结论：治理画像第一版前端正式承接位已冻结为 `Repository detail`，不再允许独立页面或第二入口。
- SubTask 1.2 结论：治理画像区被明确定位为仓库业务主内容之后的 secondary governance 区，不改写四实体详情主语义。
- SubTask 1.3 结论：spec 已显式排除跨页并列主承接区，避免 `Dashboard / Review / Decision detail` 再次长出治理画像主内容。

- SubTask 2.1 结论：治理画像区内部结构已冻结为“概览 / 维护 / 摘要回看”三部分，后续实现不必再临场拆卡片。
- SubTask 2.2 结论：`structured_summary` 被明确冻结为主阅读内容，`entry_ref` 只承担 secondary locator 角色。
- SubTask 2.3 结论：spec 已显式禁止把真实文件路径与目录说明做成大块重复说明，直接接住当前 DoD。

- SubTask 3.1 结论：治理画像前端读取已冻结为以 `repository_id` 为唯一锚点的单一主线。
- SubTask 3.2 结论：query owner 只承接只读解包、缓存键与错误状态，符合项目规则中“query 层纯只读”的约束。
- SubTask 3.3 结论：spec 已显式要求前端直接消费 `phase13-08` 治理画像读取主线，不再沿用 `phase12 project-context` 读取主线。

- SubTask 4.1 结论：mutation owner 已冻结为单一正式承接位，页面与表单组件不得散落多套写路径。
- SubTask 4.2 结论：保存成功后的精准刷新语义已固定为当前 `repository_id` 对应治理画像读取结果的单点失效刷新。
- SubTask 4.3 结论：spec 已禁止在 page / form / card 中各自内联 `useMutation`，对齐项目级前端写路径约束。

- SubTask 5.1 结论：第一版表单范围已明确收敛为 `template_source / canonical_root_files[] / global_asset_bindings[]`。
- SubTask 5.2 结论：8 项全局规范资产矩阵的前端承接方式已明确为受控展示与结构化维护，不允许随意新增第 9 项资产。
- SubTask 5.3 结论：markdown 正文、只读字段与矩阵外资产都已被排除在表单之外。

- SubTask 6.1 结论：只读展示范围已接住 `project_profile_version / track_type / docs_workflow_layout / current_phase_*` 与 `phase13-08` 后端返回结构。
- SubTask 6.2 结论：`markdown_resolvable` 与顶层目录矩阵都被冻结为轻量回看内容，不提升为主交互区。
- SubTask 6.3 结论：`entry_ref` 与真实路径已被明确约束为 secondary metadata，不得放大成主内容。

- SubTask 7.1 结论：`Repository detail` 现有“项目上下文”区已被正式登记为必须移除的旧设计。
- SubTask 7.2 结论：`Decision detail` 现有“共享项目上下文入口”卡片已被正式登记为必须移除的旧设计。
- SubTask 7.3 结论：spec 已显式禁止通过换名、并区或弱化文案延续 `phase12 project-context` 叙事。

- SubTask 8.1 结论：治理画像区被明确为 secondary governance 区，与产品定位一致。
- SubTask 8.2 结论：四实体主线与详情页结构稳定性已被显式接住，治理画像不再抢占主视觉。
- SubTask 8.3 结论：agent-only 协议、IDE 目录能力与解释性规则已被明确排除出前端主内容。

- SubTask 9.1 结论：本 spec 与 `phase13-06` 的承接位、信息架构、旧 `project-context` 退出规则保持单值一致。
- SubTask 9.2 结论：本 spec 与 `phase13-08` 的治理画像后端读写主线保持单值一致，前端承接的正是后端同源结构化事实。
- SubTask 9.3 结论：当前源码中的两处旧残留 `[Repository detail 项目上下文区]` 与 `[Decision detail 共享项目上下文入口]` 已被完整接入实现退出范围，没有遗漏。

# 实现与验证执行记录（2026-08-17）

## 实现落点

- 前端切片：`frontend/src/features/governance-profile/`（10 文件）
  - `components/`：governance-profile-section（唯一承接区壳）/ governance-profile-overview（概览层）/ governance-profile-form（维护层）/ governance-profile-readonly-summary（摘要回看层）
  - `data/`：use-governance-profile-read（唯一 query owner）/ connect-client（切片局部 transport）/ governance-profile-baseline（范式 v1 基线 + 8 项资产矩阵 + 枚举展示映射）
  - `application/`：use-update-governance-profile（唯一 mutation owner，成功后精准 invalidate `['governance-profile', repositoryId]`）
  - `index.ts` 只导出 `GovernanceProfileSection`，不导出 owner 供页面散装拼装
- 挂载：`repository-binding-detail-page.tsx` L305-306，位于仓库业务主内容之后的 secondary governance 区
- 旧 UI 退出：Repository detail“项目上下文”区与 Decision detail“共享项目上下文入口”卡片均已整体移除；`project-context/` 切片收窄为四实体语义常量切片（合法保留）

## 验证证据

- 静态：`npx tsc --noEmit` 通过（exit 0）
- API 级（Connect JSON，经 `/api` 前缀，repository `acde5fad-0c1f-4ad9-8692-76e131b81e5f`）：
  - 已创建画像读取 → 200，返回完整画像（9 根级文件 + 8 资产承接 + 只读阶段字段）
  - 等值幂等更新 → 200，`trackType` / `currentPhaseName` 只读字段不被写路径改写，`updatedAt` 刷新
  - 空 `canonicalRootFiles` → 400 `invalid_argument`（与前端预校验同规则）
  - 仓库不存在（全零 UUID）→ 404 `not_found`
- 浏览器级（4 场景 PASS）：回看态展示、维护表单打开/预填/`docs_workflow_layout` 只读透传/取消恢复、旧 UI 退出（两处详情页均无“项目上下文”文案残留）、控制台无错误且无网络失败
- 独立复核：9 项复核清单（承接位唯一性/三层结构/读写路径/表单范围/只读展示/旧 UI 退出/UI 层级/导出约束）全部 PASS

## 复核发现（非阻断，留档）

1. 后端 `GetGovernanceProfile` / `UpdateGovernanceProfile` 传入非 UUID 格式 `repositoryId` 时返回 500 + 原始 SQL 错误而非 400/404；定位在 `backend/internal/governanceprofile/candidate/repository_reader.go` 直接将输入传入 UUID 列比较。前端真实链路 repositoryId 恒为有效 UUID，不影响本阶段收口；建议后续走 `fix*` workflow，参照 `repositorybinding/validate.go` 的 `ValidateRepositoryID` 前置校验模式修复。
2. mutation 成功回流可进一步用 `setQueryData` 立即回填保存响应，避免 invalidate refetch 间隙短暂显示旧值；当前实现已满足 spec“精准刷新”要求，列为后续 polish 候选。
3. 保存失败文案直接透出 ConnectError 原始 message，未做 code → 中文归一化映射；错误可见可重试已满足 spec 核心要求，列为后续 polish 候选。
