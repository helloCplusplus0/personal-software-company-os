# Fix 002 Issue - Decision 待处理信号语义错位

## 1. 问题摘要
- 问题编号：`fix_002`
- 发现阶段：`Real-Project Dry-Run`
- 问题级别：`P0`
- 发现日期：2026-08-13
- 发现人：真实使用者 / GPT54 复核

## 2. 背景
- 业务背景：`Dashboard / Daily Review / Current Focus` 应只提示真正仍待处理的决策动作。
- 涉及模块：
  - `backend/internal/review`
  - `backend/internal/dashboard`
  - `backend/internal/decisioncenter`
  - 对应前端展示切片
- 涉及页面 / API：
  - `Dashboard`
  - `Daily Review`
  - `Current Focus`
  - `Decision Detail`
- 涉及样例数据：已创建、已被消费、已完成关联但仍显示为待处理的 `Decision`

## 3. 复现前提
- 环境：本地开发环境 / 真实 dry-run 环境
- 账号：默认测试账号或真实使用者账号
- 前置数据：存在至少一条已创建且已进入后续处理链路的 `Decision`
- 是否稳定复现：是

## 4. 复现步骤
1. 创建一条 `Decision`，并完成其实体关联或后续处理动作。
2. 返回 `Dashboard`、`Daily Review` 或 `Current Focus`。
3. 观察该 `Decision` 是否仍被标记为待处理 / 待决策。

## 5. 预期结果
- 只有真正未被确认、未被消费、仍需动作承接的 `Decision` 才应继续显示为待处理。

## 6. 实际结果
- 部分已经被消费或已经形成正式结论语义的 `Decision`，仍被持续显示为待处理。
- 用户体感上形成“我明明已经处理过，但系统仍在催促”的语义冲突。

## 7. 影响范围初判
- 是否影响主链：是
- 是否影响历史数据：否 / 待确认
- 是否影响其他入口：是

## 8. 初步观察
- 当前问题本质上不是单点 UI 文案错误，而是 `Decision` 状态与 pending signal 的领域语义没有真正分开。
- 现象集中表现为：`proposed` 同时承担“已留痕”与“待处理”双重语义。
- 现阶段仅冻结问题现象，不在 issue 文档中提前写实施方案。
