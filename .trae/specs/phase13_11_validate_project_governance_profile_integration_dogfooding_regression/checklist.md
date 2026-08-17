# phase13-11 Checklist

## 验收环境与样本

- [x] 固定样本经 `database/scripts/restore_phase11_phase12_dogfooding_sample.sh` 正式脚本恢复，未使用手工 SQL 插样本或更换 `repository_id` 锚点
- [x] 恢复后固定 `repository_id` 一次解析成功（兼容层 `GetProjectContext` 或 `GetProjectBrief` repository 块留档）
- [x] 画像未创建时 `GetProjectBrief` 返回 `not_found` 的预期行为已留档

## 人类维护路径 dogfooding

- [x] 固定样本治理画像经 `/repositories/$repositoryId` 表单手工创建（非种子 SQL / 非后门脚本）
- [x] `canonical_root_files[]` 与当前项目范式 v1 根级文件集合一致
- [x] `global_asset_bindings[]` 覆盖 8 项资产，前 5 项摘要型资产已填写 `structured_summary`
- [x] 概览区只读字段与后端 `RootFrozen*` 常量一致（`track_type=durable_system` / `docs_workflow_layout=phase/fix/audit/review` / `current_phase_name=phase13_project_governance_profile_foundation` / `current_phase_ref=plan.md#phase13_project_governance_profile_foundation` / `current_phase_status=in_progress`）
- [x] 表单保存成功回流（精准失效刷新）且刷新后回看数据一致
- [x] 摘要回看区以 `structured_summary + entry_ref` 为主视觉，真实路径未成为主内容

## agent 读取路径 dogfooding

- [x] `GetProjectBrief` 实时返回 7 顶层块完整（repository / governance_profile / global_assets / current_phase / products[] / modules[] / decisions[]）
- [x] brief 治理画像字段与 `GetGovernanceProfile` 同源一致
- [x] `products / modules / decisions` 保持数组形态（长度为 1 也不退化）
- [x] `current_phase` 从治理画像主记录三 read-only 字段单向派生
- [x] brief 无目录扫描结果、无第二套事实源投影、无自然语言指导词字段
- [x] `AGENTS.md` 与 `plan.md` 固定入口可回答 `phase13` 当前阶段状态

## 固定 6 问

- [x] Q1 治理画像版本与技术路线：answer / direct entry refs / 达标 留档
- [x] Q2 canonical 根级文件集合承接：answer / direct entry refs / 达标 留档
- [x] Q3 全局规范资产结构化摘要 + 入口关系承接：answer / direct entry refs / 达标 留档
- [x] Q4 brief 由同一 `repository_id` 驱动且无第二套目录扫描：answer / direct entry refs / 达标 留档
- [x] Q5 前端唯一承接位仍是 Repository detail 无并列第二入口：answer / direct entry refs / 达标 留档
- [x] Q6 未进入 Git 推进跟踪 / 模板接入 / 自动同步 / agent 写回：answer / direct entry refs / 达标 留档
- [x] 固定 6 问最终结论为 6 / 6 达标

## 集成测试补齐

- [x] `backend/internal/projectcontext/connect` 新增 `GetProjectBrief` 集成用例（画像创建 → 7 顶层块 + 同源断言；画像未创建 → not_found）
- [x] 测试内画像准备走 service 正式写路径，未绕过校验构造半套状态
- [x] 测试自建 fixture，不依赖共享开发库 dogfooding 样本（不引入切库样本丢失问题）
- [x] 既有测试（含旧 RPC 4 场景）不回归

## 工具链

- [x] `proto/` 下 `buf build` 通过
- [x] `proto/` 下 `buf lint` 通过
- [x] `backend/` 下 `go test ./...` 通过
- [x] `frontend/` 下 `npm run build` 通过
- [x] 各步 warning 已单独记录且未篡改通过/失败归类

## 浏览器反回归矩阵

- [x] 主验证页 `/repositories/ca261521-8daf-4248-8f12-43525326e759` 治理画像三区齐全且数据正确
- [x] `/dashboard` 正常加载不回归
- [x] `/modules` 列表与固定模块详情正常加载不回归
- [x] `/decisions` 列表与固定决策详情正常加载不回归
- [x] `/products` 列表与固定产品详情正常加载不回归
- [x] `/repositories` 列表正常加载不回归
- [x] `/onboarding` 正常加载不回归
- [x] `/reviews/daily`、`/reviews/weekly` 正常加载不回归（weekly 首轮暴露 nil 解引用缺陷，修复后复验通过）
- [x] 全矩阵无 `not_found`、无第二套治理画像入口、四实体语义不回归

## 边界证据

- [x] 不做 Git 推进跟踪：证据留档
- [x] 不做模板仓库接入：证据留档
- [x] 不做自动同步 / 目录全文扫描入库：证据留档（RPC 清单 + 迁移清单）
- [x] 不做 agent 写回 / MCP / CLI / Draft / 审批流：证据留档

## 验收记录与复核

- [x] `acceptance_report.md` 已按固定 rerun 记录格式产出（样本与 `repository_id` / Web 页面与 agent 入口 / 恢复与维护方式 / 6 问逐题结果 / 工具链逐步结果 / 浏览器矩阵 / 边界证据 / 失败点与是否达标）
- [x] 子代理独立复核完成且全部阻断性问题已修复（复核结论：A-E 全 PASS，阻断问题无，非阻断建议 3 条已留档处置）
- [x] 不同执行者可按同一固定样本、固定入口、固定问题与固定工具链顺序 rerun
