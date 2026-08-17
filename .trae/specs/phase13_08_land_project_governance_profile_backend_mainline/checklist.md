- [x] 项目治理画像第一版后端实现主线已明确冻结为同一 `repository_id` 驱动的 repository-scoped governance profile 写读主线
- [x] `.proto -> ConnectRPC -> Go service -> repository -> PostgreSQL` 的单一实现链路已明确冻结
- [x] Go 代码分层与 router 装配落点已明确冻结
- [x] 治理画像主记录、`canonical_root_file_bindings` 与 `global_asset_bindings` 的第一版数据库结构已明确冻结
- [x] 写路径可写字段、`read-only` 排除项与单一事务边界已明确冻结
- [x] 读取路径已明确区分结构化字段、结构化摘要与正文回源能力
- [x] 自动扫描、自动同步、全文入库与完整 brief 主线已明确冻结为 `phase13-08` 非目标
- [x] 现有 `ProjectContextService` 与治理画像后端主线的职责边界已明确冻结
- [x] 不产生第二套项目事实源或第二套并列项目上下文读取协议的约束已明确进入 formal spec
- [x] 本 spec 包与 `phase13-05 / phase13-07 / project_context.proto` 的上游口径保持单值一致

# 实现验证 Checklist（phase13-08 后端主线落地）

- [x] `proto/psco/governance_profile/v1/governance_profile.proto` 已落地并通过 `buf lint / build / gen`，生成 Go / Go Connect / TS 三链产物
- [x] `UpdateGovernanceProfileRequest` 不携带 `track_type / current_phase_* / project_profile_version`，read-only 排除在合同层成立
- [x] `database/migrations/0010_phase13_governance_profile.sql` 已创建三张正式表（主记录 + 两组 bindings），`repository_id` 唯一锚点与 bindings 层外键约束齐备，无 markdown 正文存储列
- [x] 迁移已由后端 `RunMigrations` 正式应用并记录版本（`0010_phase13_governance_profile`）
- [x] Go 模块按 `connect / service / repository / candidate / types / errors` 分层落地，写校验收敛在 `service/`，SQL 收敛在 `repository/`，跨模块读取经 `candidate/` 隔离
- [x] 保存入口为单一事务边界：主记录 UPSERT + 两组 bindings 全量替换，任一步失败整体回滚；`ON CONFLICT DO UPDATE` 分支不改写 read-only 列
- [x] read-only 字段由根级上游冻结结论投影受控初始化，不由维护写路径改写（冒烟已验证二次保存后保持不变）
- [x] 8 项资产矩阵校验生效：越出矩阵的 name 与前 5 项缺 `structured_summary` 均返回 `invalid_argument`
- [x] 读取失败语义区分 `repository not found` 与 `governance profile not created`
- [x] `go build ./... && go vet ./... && go test ./...` 全部通过，无四实体主线回归
- [x] 前端 `tsc --noEmit` 通过（TS 生成产物纯新增）
- [x] Connect 冒烟六场景全部符合预期，冒烟数据已恢复为完整画像（9 root files + 8 asset bindings）
- [x] 未引入自动扫描 / 自动同步 / 全文入库 / 完整 brief 主线，未产生第二套项目事实源
