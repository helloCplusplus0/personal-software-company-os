# Checklist

- [x] `Module Registry` 当前阶段的数据读写范围已明确
- [x] 当前阶段最小接口承接前提已明确
- [x] `ModuleListRead` 与 `ModuleDetailRead` 的最小读接口分组已明确
- [x] `ModuleCreateWrite`、`ModuleReleaseWrite`、`ModuleBindingWrite` 的最小写接口分组已明确
- [x] 绑定动作候选目标读取前提（`ProductBindingCandidateRead` / `RepositoryBindingCandidateRead`）已明确只服务当前绑定动作，不扩写为独立主线
- [x] `Decision` 在当前阶段已明确只作为只读/跳转接口边界
- [x] `Decision` 入口已明确作为 `ModuleDetailRead` 的附属读取承接，不设独立读接口组
- [x] 当前阶段未把独立 `RecordDecision` 写接口主线提前并入
- [x] 当前阶段未提前冻结完整查询矩阵
- [x] 当前阶段未提前冻结 `Dashboard` 聚合接口
- [x] 本次规格与 `Contract First` 方向一致
- [x] 本次规格与 `phase01-06` 正式 MVP 规格正文保持一致