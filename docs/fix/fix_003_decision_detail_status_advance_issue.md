# Fix 003 Issue - Decision Detail 缺少正式状态推进入口

## 1. 问题摘要
- 问题编号：`fix_003`
- 发现阶段：`Real-Project Dry-Run`
- 问题级别：`P0`
- 发现日期：2026-08-13
- 发现人：真实使用者 / GPT54 复核

## 2. 背景
- 业务背景：`Decision` 既然是 PSCO MVP 中的核心经营对象，就必须存在从“留痕”进入“正式确认 / 生效 / 退出待处理”的明确动作承接位。
- 涉及模块：
  - `frontend/src/features/decision-center`
  - `backend/internal/decisioncenter`
  - `backend/internal/review`
  - `backend/internal/dashboard`
- 涉及页面 / API：
  - `Decision Detail`
  - `Daily Review`
  - `Dashboard`
- 涉及样例数据：已创建但仍无正式完成动作的 `Decision`

## 3. 复现前提
- 环境：本地开发环境 / 真实 dry-run 环境
- 账号：默认测试账号或真实使用者账号
- 前置数据：存在一条已创建并可进入详情页查看的 `Decision`
- 是否稳定复现：是

## 4. 复现步骤
1. 从 `Dashboard` 或 `Review` 进入某条 `Decision` 详情页。
2. 尝试寻找“确认 / 生效 / 完成处理”一类正式状态推进入口。
3. 观察处理完成后，是否能够回流到 `Review`、详情页与 pending signal 展示。

## 5. 预期结果
- `Decision Detail` 应存在明确、可执行的正式状态推进入口。
- 状态推进后，详情页、review 语义与 dashboard 待读信号应保持一致。

## 6. 实际结果
- 详情页缺少正式状态推进 CTA。
- 用户即使完成了实体关联或理解上已经“处理完”，系统仍缺乏一个明确的完成出口。

## 7. 影响范围初判
- 是否影响主链：是
- 是否影响历史数据：否 / 待确认
- 是否影响其他入口：是

## 8. 初步观察
- 当前问题并非简单缺一个按钮，而是 `Decision Detail` 没有承接 “review 发现 -> decision 处理 -> 实体回流” 这条业务动作链。
- 它与 `fix_002` 高度相关，但仍应作为独立 fix 记录，因为它有独立的页面承接位和用户可感知断点。
- 现阶段仅冻结问题现象，不在 issue 文档中提前写实施方案。
