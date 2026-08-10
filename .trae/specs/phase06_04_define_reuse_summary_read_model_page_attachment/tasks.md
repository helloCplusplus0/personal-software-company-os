# Tasks

- [x] Task 1: 冻结 `module_reuse_summary` 的最小读模型。把字段、统计口径、排序与空态写成单值结论。
  - [x] SubTask 1.1: 明确 `module_reuse_summary` 至少包含 `module_id / reuse_product_count / latest_reuse_at / explanation_text`
  - [x] SubTask 1.2: 明确最小统计口径固定为“一个 Module 当前被多少 Product 直接复用”
  - [x] SubTask 1.3: 明确 Dashboard / Product Detail 中的最小排序与裁剪规则
  - [x] SubTask 1.4: 验证当前规格没有把 `Repository` 映射数、`Decision` 链接数混入复用计数

- [x] Task 2: 冻结 `capability_summary` 的最小读模型。把聚合单位、字段集合与事实来源压缩成可执行前提。
  - [x] SubTask 2.1: 明确 `capability_summary` 至少包含 `capability_key / capability_label / supporting_module_count / latest_capability_update_at / empty_state_text`
  - [x] SubTask 2.2: 明确最小聚合单位固定为 `capability_key`
  - [x] SubTask 2.3: 明确唯一事实来源固定为 `Module.capability_key + 系统内置 capability_label 映射`
  - [x] SubTask 2.4: 验证当前规格没有引入独立 `Capability` 重实体或第二套事实源
  - [x] SubTask 2.5: 明确 `capability_summary` 在 Dashboard（最多 5 条）与 Product Detail（全量）中的排序规则与裁剪上限

- [x] Task 3: 冻结未声明 capability 的 Module 处理规则与空聚合语义。避免实现阶段再临时猜测。
  - [x] SubTask 3.1: 明确未填写 `capability_key` 的 `Module` 不参与当前阶段 `capability_summary` 聚合
  - [x] SubTask 3.2: 明确该行为不影响 `module_reuse_summary` 与首轮成功会话
  - [x] SubTask 3.3: 明确“全部 Module 都未填写 capability”时返回成功空态而不是错误态

- [x] Task 4: 冻结 Dashboard、Module Detail、Product Detail 的最小挂接位。确保复用反馈成为正式页面能力而不是新开一套中心。
  - [x] SubTask 4.1: 明确 Dashboard 正式挂接位位于现有 `Asset Feedback` 区块内部
  - [x] SubTask 4.2: 明确 Module Detail 正式挂接位位于 `Module Summary` 邻近区域
  - [x] SubTask 4.3: 明确 Product Detail 正式挂接位位于已绑定模块相关区域附近
  - [x] SubTask 4.4: 验证当前规格没有新增独立一级“复用中心”或 `/reuse` 页面
  - [x] SubTask 4.5: 明确 `ReuseSummaryRead` owner 单值化、复用感知不合并到 `FeedbackSignalRead`、`Asset Feedback` 区块内读取状态分层独立

- [x] Task 5: 冻结解释文案、空状态与新鲜度语义。让页面在成功空态、失败态和最新状态下保持单值口径。
  - [x] SubTask 5.1: 明确两类派生读模型的最小解释文案
  - [x] SubTask 5.2: 明确成功空态与读取失败态的区分规则
  - [x] SubTask 5.3: 明确新鲜度口径固定为“读取时反映最新已提交状态”
  - [x] SubTask 5.4: 验证当前规格不依赖异步离线刷新或后台统计表

- [x] Task 6: 完成规格一致性校验。验证本次 `phase06-04` 规格与三件套、phase05 Dashboard 基线和现有页面结构保持一致。
  - [x] SubTask 6.1: 验证本次规格与 `phase06` shared baseline、architecture plan、dev plan 保持一致
  - [x] SubTask 6.2: 验证 Dashboard 挂接位沿用现有 `Asset Feedback` 主线，不破坏 phase05 已冻结的一级结构
  - [x] SubTask 6.3: 验证本次规格足以支撑后续 `ReuseSummaryRead`、`.proto`、页面实现与验收 fixture

# Task Dependencies

- `Task 2` depends on `Task 1`
- `Task 3` depends on `Task 2`
- `Task 4` depends on `Task 1` and `Task 2`
- `Task 5` depends on `Task 1`, `Task 2`, and `Task 3`
- `Task 6` depends on `Task 1`, `Task 2`, `Task 3`, `Task 4`, and `Task 5`
