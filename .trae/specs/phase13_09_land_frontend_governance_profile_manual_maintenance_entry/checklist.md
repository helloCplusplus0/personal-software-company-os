- [x] `Repository detail` 已明确冻结为治理画像第一版唯一正式前端承接位
- [x] 治理画像区的“概览 / 维护 / 摘要回看”三部分结构已明确冻结
- [x] 前端治理画像读取已明确冻结为以 `repository_id` 为锚点的单一 query owner
- [x] 前端治理画像保存已明确冻结为单一 mutation owner，并要求精准刷新当前 `repository_id` 的读取结果
- [x] 第一版手工维护表单范围已明确冻结为 `template_source / canonical_root_files[] / global_asset_bindings[]`
- [x] `project_profile_version / track_type / docs_workflow_layout / current_phase_* / markdown_resolvable` 等只读内容已明确冻结为只读展示或摘要回看
- [x] `entry_ref` 与真实文件路径已明确冻结为轻量 locator / secondary metadata，而不是大块主内容
- [x] `Repository detail` 现有“项目上下文”区已明确纳入必须移除的旧设计
- [x] `Decision detail` 现有“共享项目上下文入口”卡片已明确纳入必须移除的旧设计
- [x] 本 spec 包与 `phase13-06 / phase13-08` 以及当前真实前端残留问题保持单值一致

# 实现验收结论（2026-08-17）

- [x] 治理画像前端切片已落地：`frontend/src/features/governance-profile/`（唯一 query owner / 唯一 mutation owner / 受控导出 / 范式 v1 基线常量）
- [x] `GovernanceProfileSection` 已挂载于 Repository detail 业务主内容之后的 secondary governance 区，且为全站唯一挂载点
- [x] “概览 / 维护 / 摘要回看”三层结构同属单一治理画像区，`structured_summary` 为主阅读内容，`entry_ref` 为轻量 locator
- [x] 手工维护表单范围收敛为 `template_source / canonical_root_files[] / global_asset_bindings[]`；8 项资产矩阵受控展示，`docs_workflow_layout` 只读透传，只读字段不进入提交负载
- [x] phase12 遗留 UI 已退出：Repository detail“项目上下文”区与 Decision detail“共享项目上下文入口”卡片整体移除，无换名保留
- [x] 验证通过：`tsc --noEmit` 零错误；API 级四场景（读取 200 / 等值更新 200 且只读字段保持 / 空文件 400 / 仓库不存在 404）；浏览器级四场景（回看态 / 表单预填与取消 / 旧 UI 退出 / 控制台无错误）全部 PASS
- [x] 子代理独立复核通过（9 项清单全 PASS）；3 项非阻断发现已留档至 tasks.md（后端非 UUID 输入 500、setQueryData 回流优化、错误文案归一化），后续走 `fix*` / polish 处理
